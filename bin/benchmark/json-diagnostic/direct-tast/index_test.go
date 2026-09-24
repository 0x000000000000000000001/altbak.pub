package purescript

import (
	"encoding/json"
	"fmt"
	rt "gopurs/output/gopurs_runtime"
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

// Bridge used by the preserved generated-PureScript/native differential tests.
func DirectDecodeModuleText(text string) rt.Value {
	cursor, ok := directIndex(text)
	if !ok {
		panic("expected valid indexed JSON")
	}
	return PureScript_Backend_Optimizer_CoreFn_Json_DecodeModuleImpl(rt.Value{}, Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), rt.Any(cursor))
}

func TestDirectIndexContract(t *testing.T) {
	cases := []string{`null`, `true`, `false`, `-0`, `1e-300`, `1.7e308`, `""`, `"é🙂"`, `"\ud800"`, `"\udfff"`, `"\ud834\udd1e"`, `"\u0000\t\n\\\"\/"`, "\"\xff\"", `[]`, `{}`, `{"x":1,"x":2,"\u0078":3}`, `{"\ufffd":1,"\ud800":2}`, `{"a":[1,null,{"b":true}],"z":{}}`,
		``, ` `, `01`, `1e`, `1e999`, `-1e999`, `"\x00"`, `"unterminated`, `{"x":1,}`, `[1,]`, `{"ignored":[true, false, }`, `{} false`, strings.Repeat("[", 10001) + "0" + strings.Repeat("]", 10001)}
	rng := rand.New(rand.NewSource(20260924))
	for i := 0; i < 1000; i++ {
		value := map[string]any{"n": rng.Float64() * 1e20, "a": []any{rng.Intn(100000), nil, fmt.Sprintf("key:%d\n🙂", i)}}
		bytes, _ := json.Marshal(value)
		cases = append(cases, string(bytes))
		if len(bytes) > 1 {
			cases = append(cases, string(bytes[:rng.Intn(len(bytes))]))
		}
	}
	for i, text := range cases {
		var expected any
		err := json.Unmarshal([]byte(text), &expected)
		cursor, ok := directIndex(text)
		if ok != (err == nil) {
			t.Fatalf("case %d accepted=%v error=%v: %q", i, ok, err, text)
		}
		if ok {
			got := directMaterialize(cursor)
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("case %d: got %#v want %#v", i, got, expected)
			}
		}
	}
	t.Logf("%d parser/index cases", len(cases))
}

func TestDirectStringsOwnStorage(t *testing.T) {
	for _, text := range []string{`"owned text"`, `"é🙂"`, `"escaped\ntext"`, `"\ud800"`} {
		source := strings.Clone(text)
		cursor, ok := directIndex(source)
		if !ok {
			t.Fatal(text)
		}
		value := cursor.native().(string)
		start := uintptr(unsafe.Pointer(unsafe.StringData(source)))
		end := start + uintptr(len(source))
		pointer := uintptr(unsafe.Pointer(unsafe.StringData(value)))
		if pointer >= start && pointer < end {
			t.Fatal("published string retains source")
		}
	}
}

func TestDirectIndexConcurrentAndIndependent(t *testing.T) {
	const text = `{"a":[1,2,{"x":"owned"}],"duplicate":false,"duplicate":true}`
	var group sync.WaitGroup
	for i := 0; i < 16; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for n := 0; n < 50; n++ {
				cursor, ok := directIndex(text)
				if !ok {
					t.Error("index rejected")
				}
				out := directMaterialize(cursor).(map[string]any)
				if out["duplicate"] != true {
					t.Error("last key lost")
				}
				out["a"] = nil
			}
		}()
	}
	group.Wait()
}
