// Audit: what the Go TAST decode phase actually contains.
//
// The diagnostic times `Apply(Test.JsonTypedAst.decode, json)`. That entry
// point returns the PS record representation: the native decoder's struct is
// materialised into a boxed record at the boundary. The C++ reference keeps
// native structures. This program times each layer separately over the whole
// corpus:
//
//   validated    - DecodeModuleImpl with the real usage validation
//   unvalidated  - DecodeModuleImpl with a no-op validate (same decoder)
//   worker       - Call_Test_JsonTypedAst_decode: native + validation, no
//                  record materialisation
//   boundary     - the driver's decode entry point: worker + materialisation
//
// and authenticates the boundary result against the frozen fingerprints.
//
// Lives in the generated module (gopurs/output) so the imports resolve.
// Run with DIAG_CORPUS=<decompressed corpus> and DIAG_EXPECTED=<expected.json>.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gopurs/output/gopurs_runtime"
	"gopurs/output/purescript"
)

type corpusCase struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
}

func nowUS() float64 { return float64(time.Now().UnixNano()) / 1000.0 }

// Same canonical hash the diagnostic drivers use.
func canonicalHash(text string) string {
	var value any
	if err := json.Unmarshal([]byte(text), &value); err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}

var sink []gopurs_runtime.Value

func main() {
	raw, err := os.ReadFile(os.Getenv("DIAG_CORPUS"))
	if err != nil {
		panic(err)
	}
	var cases []corpusCase
	if err := json.Unmarshal(raw, &cases); err != nil {
		panic(err)
	}

	parse := purescript.Get_Test_JsonTypedAst_parse()
	decode := purescript.Get_Test_JsonTypedAst_decode()
	fingerprint := purescript.Get_Test_JsonTypedAst_fingerprint()
	realValidate := purescript.Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule()
	stubValidate := gopurs_runtime.Func(func(gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 1}
	})
	fallback := gopurs_runtime.Value{Type: 1}

	parsed := make([]gopurs_runtime.Value, len(cases))
	for i, entry := range cases {
		parsed[i] = gopurs_runtime.Apply(parse, gopurs_runtime.Str(entry.Contents))
	}

	// Authenticate the entry point the diagnostic measures.
	hashes := make([]string, len(cases))
	for i := range cases {
		result := gopurs_runtime.Apply(decode, parsed[i])
		hashes[i] = canonicalHash(gopurs_runtime.Apply(fingerprint, result).StrVal())
	}
	oracleMatches := true
	if expected := os.Getenv("DIAG_EXPECTED"); expected != "" {
		text, err := os.ReadFile(expected)
		if err != nil {
			panic(err)
		}
		var oracle struct {
			Fingerprints []string `json:"fingerprints"`
		}
		if err := json.Unmarshal(text, &oracle); err != nil {
			panic(err)
		}
		for i := range hashes {
			if hashes[i] != oracle.Fingerprints[i] {
				oracleMatches = false
			}
		}
	}

	measure := func(runs int, body func()) float64 {
		var best float64
		for run := 0; run < runs; run++ {
			begin := nowUS()
			body()
			elapsed := nowUS() - begin
			if run == 0 || elapsed < best {
				best = elapsed
			}
		}
		return best
	}

	validated := measure(5, func() {
		var last gopurs_runtime.Value
		for i := range cases {
			last = purescript.PureScript_Backend_Optimizer_CoreFn_Json_DecodeModuleImpl(fallback, realValidate, parsed[i])
		}
		sink = append(sink[:0], last)
	})
	unvalidated := measure(5, func() {
		var last gopurs_runtime.Value
		for i := range cases {
			last = purescript.PureScript_Backend_Optimizer_CoreFn_Json_DecodeModuleImpl(fallback, stubValidate, parsed[i])
		}
		sink = append(sink[:0], last)
	})
	worker := measure(7, func() {
		for i := range cases {
			purescript.Call_Test_JsonTypedAst_decode(parsed[i])
		}
	})
	boundary := measure(5, func() {
		var last gopurs_runtime.Value
		for i := range cases {
			last = gopurs_runtime.Apply(decode, parsed[i])
		}
		sink = append(sink[:0], last)
	})

	report := map[string]any{
		"modules":         len(cases),
		"validated_us":    validated,
		"unvalidated_us":  unvalidated,
		"worker_us":       worker,
		"boundary_us":     boundary,
		"oracle_matches":  oracleMatches,
	}
	encoded, err := json.Marshal(report)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))
}
