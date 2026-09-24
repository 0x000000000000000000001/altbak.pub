package purescript

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	rt "gopurs/output/gopurs_runtime"
)

// Test-only materialization traverses precisely the typed read APIs used by
// final constructors. Production materializes only the cold sourceSpan boundary.
func typedMaterialize(c tcCursor) any {
	switch c.kind() {
	case '{':
		out := map[string]any{}
		for _, key := range c.Keys() {
			v, ok := c.Lookup(key)
			if !ok {
				panic(key)
			}
			out[key] = typedMaterialize(v)
		}
		return out
	case '[':
		a, _ := c.array()
		out := make([]any, a.count)
		for it := a.iter(); it.more(); {
			i, v := it.next()
			out[i] = typedMaterialize(v)
		}
		return out
	case '"':
		v, _ := c.text()
		return v
	case 'n':
		return nil
	case 't', 'f':
		v, _ := c.boolean()
		return v
	default:
		v, _ := c.number()
		return v
	}
}

func TestTypedCursorReadsAndOwnership(t *testing.T) {
	cases := []string{`null`, `[]`, `{}`, `{"x":1,"x":2,"\u0078":3}`, `{"\ufffd":1,"\ud800":2}`, `["\ud800","\ud834\udd1e",-0,1e-300]`}
	rng := rand.New(rand.NewSource(20260924))
	for i := 0; i < 1000; i++ {
		text, _ := json.Marshal(map[string]any{"nested": []any{rng.Float64(), fmt.Sprintf("owned %d\né🙂", i), nil}, "flag": i%2 == 0})
		cases = append(cases, string(text))
	}
	for _, text := range cases {
		cursor, ok := directIndex(text)
		if !ok {
			t.Fatal(text)
		}
		var expected any
		if err := json.Unmarshal([]byte(text), &expected); err != nil {
			t.Fatal(err)
		}
		if got := typedMaterialize(tcCursor(cursor)); !reflect.DeepEqual(got, expected) {
			t.Fatalf("%s: %#v / %#v", text, got, expected)
		}
	}
	for _, text := range []string{`"owned text"`, `"é🙂"`, `"escaped\ntext"`, `"\ud800"`} {
		source := strings.Clone(text)
		cursor, _ := directIndex(source)
		value, _ := tcCursor(cursor).text()
		start := uintptr(unsafe.Pointer(unsafe.StringData(source)))
		pointer := uintptr(unsafe.Pointer(unsafe.StringData(value)))
		if pointer >= start && pointer < start+uintptr(len(source)) {
			t.Fatal("final string retained input")
		}
	}
	t.Logf("%d typed cursor cases", len(cases))
}

func TestTypedIndexFullSyntaxValidation(t *testing.T) {
	cases := []string{
		`{"ignored":"\x00"}`, `{"ignored":"\u000z"}`, `{"ignored":1e999}`,
		`{"ignored":-1e999}`, `{"ignored":1e-999}`, `{"ignored":[0,]}`,
		`[+1]`, `[-]`, `[00]`, `[1.]`, `[1e+]`, `true false`, `nullx`,
		`{"\ud800":0,"\ufffd":1,"\udfff":2}`, "{\"\xff\":1,\"�\":2}",
	}
	rng := rand.New(rand.NewSource(2026092402))
	seeds := []string{`{"key":[true,false,null,-0,3.2e-15,"é🙂\ud800\u0000"],"nested":{"x":1}}`, `"\"\\\/\b\f\n\r\t\u0041"`, `[{},[],123,"plain"]`}
	for i := 0; i < 20000; i++ {
		text := []byte(seeds[rng.Intn(len(seeds))])
		for n := 0; n < 1+rng.Intn(3); n++ {
			at := rng.Intn(len(text))
			text[at] = byte(rng.Intn(256))
		}
		cases = append(cases, string(text))
	}
	for i, text := range cases {
		var expected any
		err := json.Unmarshal([]byte(text), &expected)
		cursor, ok := directIndex(text)
		if ok != (err == nil) {
			t.Fatalf("case %d: accepted=%v error=%v input=%q", i, ok, err, text)
		}
		if ok && !reflect.DeepEqual(typedMaterialize(tcCursor(cursor)), expected) {
			t.Fatalf("case %d: decoded value differs: %q", i, text)
		}
	}
	for _, depth := range []int{9999, 10000, 10001} {
		for _, leaf := range []string{"", "0"} {
			text := strings.Repeat("[", depth) + leaf + strings.Repeat("]", depth)
			_, ok := directIndex(text)
			if ok != json.Valid([]byte(text)) {
				t.Fatalf("depth %d leaf %q: accepted=%v", depth, leaf, ok)
			}
		}
	}
	t.Logf("%d syntax/mutation cases plus six depth boundaries", len(cases))
}

func TestTypedNumbersPreserveBits(t *testing.T) {
	for _, text := range []string{"0", "-0", "-0.0", "1", "-1", "999999999999999", "-99999999999999", "9007199254740993", "-9007199254740993", "1e-999", "-1e-999", "1.7e308", "1.0000000000000002"} {
		var expected float64
		if err := json.Unmarshal([]byte(text), &expected); err != nil {
			t.Fatal(err)
		}
		cursor, ok := directIndex(text)
		if !ok {
			t.Fatal(text)
		}
		actual, ok := tcCursor(cursor).number()
		if !ok || math.Float64bits(actual) != math.Float64bits(expected) {
			t.Fatalf("%s: actual %x expected %x", text, math.Float64bits(actual), math.Float64bits(expected))
		}
	}
}

func TestTypedFinalModuleOwnsStrings(t *testing.T) {
	bytes, err := os.ReadFile(os.Getenv("DIAG_CORPUS"))
	if err != nil {
		t.Fatal(err)
	}
	var files []struct{ Name, Contents string }
	if err = json.Unmarshal(bytes, &files); err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		// A private mutable backing buffer makes leaked source views observable.
		// It is overwritten only after complete decoding and usage validation.
		buffer := []byte(file.Contents)
		text := unsafe.String(unsafe.SliceData(buffer), len(buffer))
		cursor, ok := directIndex(text)
		if !ok {
			t.Fatal(file.Name)
		}
		result := tcDecodeModule(rt.Value{}, Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), tcCursor(cursor))
		if tc_cndIsLeft(result) {
			t.Fatal("decode failed", file.Name)
		}
		value := (*Constructor_Data_Either_Right[rt.Value, rt.Value])(result.UnsafePtr).V0
		fingerprint := Get_Test_JsonTypedAst_fingerprint()
		before := rt.Apply(fingerprint, value).StrVal()
		for i := range buffer {
			buffer[i] = '?'
		}
		after := rt.Apply(fingerprint, value).StrVal()
		if before != after {
			t.Fatal("final module retained input bytes", file.Name)
		}
	}
	t.Logf("%d complete modules survive input overwrite", len(files))
}
