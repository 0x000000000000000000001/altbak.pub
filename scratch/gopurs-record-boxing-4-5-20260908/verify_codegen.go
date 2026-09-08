// Compare generated Go trees after normalizing compact record constructors.
// Usage: go run verify_codegen.go BEFORE_ROOT AFTER_ROOT > codegen-verification.json
// Reads the trees without modifying them. Comments and source positions are
// ignored; every other AST distinction remains visible in the comparison.
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

type counts map[string]int

type tree struct {
	Files    map[string]string
	Excluded []string
}

type parsedFile struct {
	Raw        []byte
	Normalized []byte
	Helpers    counts
}

type difference struct {
	Path             string `json:"path"`
	BeforeSHA256     string `json:"before_normalized_sha256"`
	AfterSHA256      string `json:"after_normalized_sha256"`
	FirstChangedLine int    `json:"first_changed_line"`
	BeforeLine       string `json:"before_line"`
	AfterLine        string `json:"after_line"`
}

type changedFile struct {
	Path            string `json:"path"`
	BeforeHelpers   counts `json:"before_helpers"`
	AfterHelpers    counts `json:"after_helpers"`
	NormalizedEqual bool   `json:"normalized_equal"`
}

type report struct {
	OK                      bool          `json:"ok"`
	BeforeRoot              string        `json:"before_root"`
	AfterRoot               string        `json:"after_root"`
	Normalization           string        `json:"normalization"`
	Ignored                 []string      `json:"ignored"`
	BeforeFiles             int           `json:"before_go_files"`
	AfterFiles              int           `json:"after_go_files"`
	FilesCompared           int           `json:"files_compared"`
	ExcludedBefore          []string      `json:"excluded_before"`
	ExcludedAfter           []string      `json:"excluded_after"`
	OnlyBefore              []string      `json:"only_before"`
	OnlyAfter               []string      `json:"only_after"`
	BeforeHelpers           counts        `json:"before_helpers_common_files"`
	AfterHelpers            counts        `json:"after_helpers_common_files"`
	RawChangedFiles         []changedFile `json:"raw_changed_files"`
	UnexpectedChangedFiles  []difference  `json:"unexpected_changed_files"`
	Errors                  []string      `json:"errors"`
}

func newCounts() counts {
	out := counts{"RecordDict": 0}
	for size := 0; size <= 5; size++ {
		out[fmt.Sprintf("RecordDict%d", size)] = 0
	}
	return out
}

func addCounts(total, next counts) {
	for name, amount := range next {
		total[name] += amount
	}
}

func collect(root string) (tree, error) {
	out := tree{Files: map[string]string{}, Excluded: []string{}}
	info, err := os.Stat(root)
	if err != nil {
		return out, err
	}
	if !info.IsDir() {
		return out, fmt.Errorf("not a directory: %s", root)
	}
	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			return nil
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		relative = filepath.ToSlash(relative)
		if entry.Name() == "bench_test.go" || strings.HasSuffix(entry.Name(), "_bench_test.go") {
			out.Excluded = append(out.Excluded, relative)
			return nil
		}
		out.Files[relative] = path
		return nil
	})
	sort.Strings(out.Excluded)
	return out, err
}

// ast.Inspect reaches the nodes that own all token.Pos fields. Two position
// fields also encode syntax: CallExpr.Ellipsis marks variadic expansion and
// TypeSpec.Assign marks a type alias. Preserve those flags at a fixed position;
// erase their original offsets along with every other source position.
func clearPositions(file *ast.File) {
	positionType := reflect.TypeOf(token.NoPos)
	ast.Inspect(file, func(node ast.Node) bool {
		if node == nil {
			return true
		}
		value := reflect.ValueOf(node).Elem()
		for index := 0; index < value.NumField(); index++ {
			field := value.Field(index)
			if field.Type() == positionType && field.CanSet() {
				keepSyntaxFlag := false
				switch node.(type) {
				case *ast.CallExpr:
					keepSyntaxFlag = value.Type().Field(index).Name == "Ellipsis"
				case *ast.TypeSpec:
					keepSyntaxFlag = value.Type().Field(index).Name == "Assign"
				}
				if keepSyntaxFlag && field.Int() != 0 {
					field.SetInt(1)
				} else {
					field.SetInt(0)
				}
			}
		}
		return true
	})
}

func parseAndNormalize(path string) (parsedFile, error) {
	out := parsedFile{Helpers: newCounts()}
	var err error
	out.Raw, err = os.ReadFile(path)
	if err != nil {
		return out, err
	}
	// Object resolution is unnecessary and would introduce AST back-references.
	// Do not parse comments: generated TAST comments are outside this code audit.
	file, err := parser.ParseFile(token.NewFileSet(), path, out.Raw, parser.SkipObjectResolution)
	if err != nil {
		return out, err
	}
	invalid := []string{}
	ast.Inspect(file, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		qualifier, ok := selector.X.(*ast.Ident)
		if !ok || qualifier.Name != "gopurs_runtime" {
			return true
		}
		name := selector.Sel.Name
		if name == "RecordDict" {
			out.Helpers[name]++
			return true
		}
		if len(name) != len("RecordDict0") || !strings.HasPrefix(name, "RecordDict") || name[len(name)-1] < '0' || name[len(name)-1] > '5' {
			return true
		}
		out.Helpers[name]++
		size := int(name[len(name)-1] - '0')
		if len(call.Args) != 2*size || call.Ellipsis.IsValid() {
			invalid = append(invalid, fmt.Sprintf("%s requires %d ordinary arguments, got %d (variadic=%t)", name, 2*size, len(call.Args), call.Ellipsis.IsValid()))
			return true
		}
		// Compact helpers take all keys first, then all boxed values. This also
		// turns the zero-argument helper into two empty composite literals.
		keys := &ast.CompositeLit{
			Type: &ast.ArrayType{Elt: ast.NewIdent("string")},
			Elts: append([]ast.Expr{}, call.Args[:size]...),
		}
		values := &ast.CompositeLit{
			Type: &ast.ArrayType{Elt: &ast.SelectorExpr{X: ast.NewIdent("gopurs_runtime"), Sel: ast.NewIdent("Value")}},
			Elts: append([]ast.Expr{}, call.Args[size:]...),
		}
		selector.Sel = ast.NewIdent("RecordDict")
		call.Args = []ast.Expr{keys, values}
		return true
	})
	if len(invalid) > 0 {
		return out, fmt.Errorf("invalid compact calls: %s", strings.Join(invalid, "; "))
	}
	clearPositions(file)
	var rendered bytes.Buffer
	if err := format.Node(&rendered, token.NewFileSet(), file); err != nil {
		return out, err
	}
	out.Normalized = rendered.Bytes()
	return out, nil
}

func digest(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func firstDifference(path string, before, after []byte) difference {
	out := difference{Path: path, BeforeSHA256: digest(before), AfterSHA256: digest(after)}
	left, right := strings.Split(string(before), "\n"), strings.Split(string(after), "\n")
	for index := 0; index < len(left) || index < len(right); index++ {
		leftLine, rightLine := "<end of file>", "<end of file>"
		if index < len(left) {
			leftLine = left[index]
		}
		if index < len(right) {
			rightLine = right[index]
		}
		if leftLine != rightLine {
			out.FirstChangedLine, out.BeforeLine, out.AfterLine = index+1, leftLine, rightLine
			break
		}
	}
	return out
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: verify_codegen BEFORE_ROOT AFTER_ROOT")
		os.Exit(2)
	}
	beforeRoot, beforeErr := filepath.Abs(os.Args[1])
	afterRoot, afterErr := filepath.Abs(os.Args[2])
	if beforeErr != nil || afterErr != nil {
		fmt.Fprintln(os.Stderr, "invalid root path:", beforeErr, afterErr)
		os.Exit(2)
	}
	out := report{
		BeforeRoot: beforeRoot, AfterRoot: afterRoot,
		Normalization: "gopurs_runtime.RecordDictN(k0,...,kN-1,v0,...,vN-1), N=0..5 -> RecordDict([]string{keys}, []gopurs_runtime.Value{values})",
		Ignored: []string{"comments", "source positions and formatting", "bench_test.go and *_bench_test.go"},
		OnlyBefore: []string{}, OnlyAfter: []string{},
		BeforeHelpers: newCounts(), AfterHelpers: newCounts(),
		RawChangedFiles: []changedFile{}, UnexpectedChangedFiles: []difference{}, Errors: []string{},
	}
	before, beforeErr := collect(beforeRoot)
	after, afterErr := collect(afterRoot)
	out.BeforeFiles, out.AfterFiles = len(before.Files), len(after.Files)
	out.ExcludedBefore, out.ExcludedAfter = before.Excluded, after.Excluded
	if beforeErr != nil {
		out.Errors = append(out.Errors, "before tree: "+beforeErr.Error())
	}
	if afterErr != nil {
		out.Errors = append(out.Errors, "after tree: "+afterErr.Error())
	}
	common := []string{}
	for path := range before.Files {
		if _, exists := after.Files[path]; exists {
			common = append(common, path)
		} else {
			out.OnlyBefore = append(out.OnlyBefore, path)
		}
	}
	for path := range after.Files {
		if _, exists := before.Files[path]; !exists {
			out.OnlyAfter = append(out.OnlyAfter, path)
		}
	}
	sort.Strings(common)
	sort.Strings(out.OnlyBefore)
	sort.Strings(out.OnlyAfter)
	if len(common) == 0 {
		out.Errors = append(out.Errors, "no common Go source files to compare")
	}
	for _, path := range common {
		left, leftErr := parseAndNormalize(before.Files[path])
		right, rightErr := parseAndNormalize(after.Files[path])
		addCounts(out.BeforeHelpers, left.Helpers)
		addCounts(out.AfterHelpers, right.Helpers)
		if leftErr != nil {
			out.Errors = append(out.Errors, path+" (before): "+leftErr.Error())
		}
		if rightErr != nil {
			out.Errors = append(out.Errors, path+" (after): "+rightErr.Error())
		}
		if leftErr != nil || rightErr != nil {
			continue
		}
		out.FilesCompared++
		equal := bytes.Equal(left.Normalized, right.Normalized)
		if !bytes.Equal(left.Raw, right.Raw) {
			out.RawChangedFiles = append(out.RawChangedFiles, changedFile{
				Path: path, BeforeHelpers: left.Helpers, AfterHelpers: right.Helpers, NormalizedEqual: equal,
			})
		}
		if !equal {
			out.UnexpectedChangedFiles = append(out.UnexpectedChangedFiles, firstDifference(path, left.Normalized, right.Normalized))
		}
	}
	out.OK = len(out.Errors) == 0 && len(out.OnlyBefore) == 0 && len(out.OnlyAfter) == 0 && len(out.UnexpectedChangedFiles) == 0
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if !out.OK {
		os.Exit(1)
	}
}
