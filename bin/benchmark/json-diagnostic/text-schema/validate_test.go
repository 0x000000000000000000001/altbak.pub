package purescript

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"unsafe"

	rt "gopurs/output/gopurs_runtime"
)

func textSchemaFingerprint(value rt.Value) string {
	return rt.Apply(Get_Test_JsonDecoding_fingerprint(), value).StrVal()
}

func textSchemaOrdinary(text string) string {
	parsed, err := argonautParseJSON(text)
	if err != nil {
		return textSchemaFingerprint(rt.Apply(Get_Data_Argonaut_Decode_Parser_parseJson(), rt.Str(text)))
	}
	return textSchemaFingerprint(rt.Apply(Get_Test_JsonDecoding_decode(), rt.Box(parsed)))
}

func textSchemaPublic(text string) string {
	return textSchemaFingerprint(rt.Apply(Get_Test_JsonDecoding_decodeText(), rt.Str(text)))
}

func textSchemaCases(t *testing.T) []string {
	t.Helper()
	bytes, err := os.ReadFile(os.Getenv("DIAG_CORPUS"))
	if err != nil {
		t.Fatal(err)
	}
	var files []struct{ Contents string }
	if err := json.Unmarshal(bytes, &files); err != nil {
		t.Fatal(err)
	}
	cases := make([]string, 0, len(files)+2020)
	for _, file := range files {
		cases = append(cases, file.Contents)
	}
	rng := rand.New(rand.NewSource(2026092405))
	values := []any{nil, true, false, "wrong", float64(-2147483649), float64(2147483648), 1.5, []any{}, map[string]any{}}
	for n := 0; n < 2000; n++ {
		user := map[string]any{"id": 1, "active": true, "name": "owned-é🙂", "tags": []any{"saved"}, "profile": map[string]any{"city": "Paris", "note": nil, "scores": []any{1.5, -0.0, 2}}}
		event := map[string]any{"tag": "view", "path": "/home", "duration": 12}
		item := map[string]any{"sku": "owned-sku", "price": 12.75, "quantity": 2}
		payload := map[string]any{"version": 1, "next": nil, "users": []any{user}, "events": []any{event}}
		if n%2 == 0 {
			event["tag"] = "purchase"
			event["items"] = []any{item}
			event["orderId"] = 3
		}
		target := []map[string]any{payload, user, event, item, user["profile"].(map[string]any)}[rng.Intn(5)]
		keys := []string{"version", "next", "users", "events", "id", "active", "name", "tags", "profile", "tag", "path", "duration", "items", "orderId", "sku", "price", "quantity", "city", "note", "scores"}
		for mutations := 0; mutations < 1+rng.Intn(4); mutations++ {
			key := keys[rng.Intn(len(keys))]
			if rng.Intn(3) == 0 {
				delete(target, key)
			} else {
				target[key] = values[rng.Intn(len(values))]
			}
		}
		text, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		cases = append(cases, string(text))
	}
	cases = append(cases,
		`{"version":1,"users":[],"events":[],"version":2,"ver\u0073ion":3}`,
		`{"version":1,"users":[],"events":[{"tag":"wrong","tag":"view","path":"kept"}]}`,
		`{"version":1,"users":[],"events":[{"tag":"view","duration":false}]}`,
		`{"version":false,"users":[],"events":[{"tag":"wrong"}]}`,
		`{"version":1,"users":[],"events":[],"ignored":{"x":[true,"\ud800",-0]}}`,
	)
	return cases
}

func TestTextSchemaPublicDifferential(t *testing.T) {
	tag, ok := rt.FunctionData[*typedKind](Get_Test_JsonDecoding_decode())
	if !ok || tag == nil {
		t.Fatal("missing original tagged decoder")
	}
	worker, ok := tag.textWorker.(argonautTextWorker)
	if !ok || worker == nil {
		t.Fatal("text schema was not installed")
	}
	cases := textSchemaCases(t)
	for i, text := range cases {
		want, got := textSchemaOrdinary(text), textSchemaPublic(text)
		if want != got {
			t.Fatalf("case %d differs:\n%s\nwant %s\ngot %s", i, text, want, got)
		}
	}
	t.Logf("%d public text decodes match the original decoder's complete values and exact errors", len(cases))
}

func TestTextSchemaArchivedApplicationCases(t *testing.T) {
	bytes, err := os.ReadFile(os.Getenv("DIAG_EXTRA_CORPUS"))
	if err != nil {
		t.Fatal(err)
	}
	var cases []struct{ Name, Contents string }
	if err = json.Unmarshal(bytes, &cases); err != nil {
		t.Fatal(err)
	}
	bytes, err = os.ReadFile(os.Getenv("DIAG_EXTRA_EXPECTED"))
	if err != nil {
		t.Fatal(err)
	}
	var oracle struct{ Names, Fingerprints []string }
	if err = json.Unmarshal(bytes, &oracle); err != nil {
		t.Fatal(err)
	}
	if len(cases) != len(oracle.Fingerprints) || len(cases) != len(oracle.Names) {
		t.Fatal("archived corpus length")
	}
	for i, input := range cases {
		if input.Name != oracle.Names[i] || canonicalHash(textSchemaPublic(input.Contents)) != oracle.Fingerprints[i] {
			t.Fatalf("archived case %d %s differs", i, input.Name)
		}
	}
	t.Logf("%d text decodes match the preserved pre-fusion native application's fingerprints", len(cases))
}

func TestTextSchemaCompleteOwnershipAndConcurrency(t *testing.T) {
	cases := textSchemaCases(t)[:29]
	for i, text := range cases {
		expected := textSchemaPublic(text)
		bytes := []byte(text)
		value := rt.Apply(Get_Test_JsonDecoding_decodeText(), rt.Str(unsafe.String(unsafe.SliceData(bytes), len(bytes))))
		for i := range bytes {
			bytes[i] = 'x'
		}
		if textSchemaFingerprint(value) != expected {
			t.Fatalf("result %d retained input", i)
		}
	}
	var group sync.WaitGroup
	for worker := 0; worker < 8; worker++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for _, text := range cases {
				if textSchemaPublic(text) != textSchemaOrdinary(text) {
					t.Error("concurrent result differs")
				}
			}
		}()
	}
	group.Wait()
}

// Read through the exact typed cursor accessors used by emitted constructors.
func textSchemaMaterialize(c argonautTextCursor) any {
	switch c.kind() {
	case '{':
		out := map[string]any{}
		for at := c.index + 1; at < c.document.tokens[c.index].next; {
			key := (argonautTextCursor{c.document, at}).string(true)
			value, _ := c.Lookup(key)
			out[key] = textSchemaMaterialize(value)
			at = c.document.tokens[at+1].next
		}
		return out
	case '[':
		a, _ := argonautTextArrayValues(c)
		out := make([]any, a.length())
		at := c.index + 1
		for i := range out {
			out[i] = textSchemaMaterialize(argonautTextCursor{c.document, at})
			at = c.document.tokens[at].next
		}
		return out
	case '"':
		value, _ := argonautTextString(c)
		return value
	case 'n':
		return nil
	case 't', 'f':
		value, _ := argonautTextBool(c)
		return value
	default:
		value, _ := argonautTextNumber(c)
		return value
	}
}

func TestTextSchemaIndexSyntaxAndScalars(t *testing.T) {
	cases := []string{`null`, `[]`, `{}`, `{"x":1,"x":2,"\u0078":3}`, `{"\ufffd":1,"\ud800":2}`, `["\ud800","\ud834\udd1e",-0,1e-300]`, `{"ignored":1e400}`, `{"x":false,"ignored":"\q"}`, `01`, `[1,]`, `true false`, string([]byte{'"', 0xff, '"'})}
	rng := rand.New(rand.NewSource(2026092406))
	for i := 0; i < 2000; i++ {
		text, _ := json.Marshal(map[string]any{"text": fmt.Sprintf("owned %d\né🙂", i), "nested": []any{rng.Float64(), nil, i%2 == 0}, "n": rng.Int63()})
		cases = append(cases, string(text), string(text[:rng.Intn(len(text))]))
		for j := 0; j < 3; j++ {
			bytes := append([]byte(nil), text...)
			bytes[rng.Intn(len(bytes))] = byte(rng.Intn(256))
			cases = append(cases, string(bytes))
		}
	}
	for i, text := range cases {
		var expected any
		err := json.Unmarshal([]byte(text), &expected)
		c, ok := argonautTextIndex(text)
		if ok != (err == nil) {
			t.Fatalf("syntax case %d: %q, accepted=%v, encoding/json=%v", i, text, ok, err)
		}
		if ok && !reflect.DeepEqual(textSchemaMaterialize(c), expected) {
			t.Fatalf("cursor value case %d differs: %q", i, text)
		}
	}
	for _, text := range []string{"-0", "0", "1e-300", "-1e-999", "0.1", "9007199254740993", "1.7976931348623157e308", "-999999999999999", "2.2250738585072014e-308"} {
		var expected float64
		if err := json.Unmarshal([]byte(text), &expected); err != nil {
			t.Fatal(err)
		}
		c, ok := argonautTextIndex(text)
		if !ok {
			t.Fatal(text)
		}
		got, ok := argonautTextNumber(c)
		if !ok || math.Float64bits(got) != math.Float64bits(expected) {
			t.Fatalf("numeric bits differ: %s", text)
		}
	}
	for _, depth := range []int{9999, 10000, 10001} {
		for _, pair := range [][2]string{{"[", "]"}, {`{"x":`, `}`}} {
			text := strings.Repeat(pair[0], depth) + "0" + strings.Repeat(pair[1], depth)
			var expected any
			err := json.Unmarshal([]byte(text), &expected)
			_, ok := argonautTextIndex(text)
			if ok != (err == nil) {
				t.Fatalf("depth %d differs", depth)
			}
		}
	}
	t.Logf("%d syntax/mutation cases, numeric bits and six depth boundaries agree with encoding/json", len(cases))
}

func TestTextSchemaSyntaxPrecedesCallbacks(t *testing.T) {
	decoded := 0
	decoder := rt.Func(func(_ rt.Value) rt.Value { decoded++; return rt.Int(99) })
	fallback := rt.Func(func(input rt.Value) rt.Value {
		if _, err := argonautParseJSON(input.StrVal()); err != nil {
			return rt.Int(-1)
		}
		return rt.Apply(decoder, rt.Value{})
	})
	for _, text := range []string{`{"a":false,"ignored":1e400}`, `{"a":false,"ignored":"\q"}`, `{} trailing`} {
		if got := Data_Argonaut_Decode_Parser_DecodeJsonStringImpl(1, fallback, decoder, text); got.IntVal != -1 {
			t.Fatal("fallback error changed")
		}
	}
	if decoded != 0 {
		t.Fatal("custom decoder ran on malformed input")
	}
	if got := Data_Argonaut_Decode_Parser_DecodeJsonStringImpl(1, fallback, decoder, `{}`); got.IntVal != 99 || decoded != 1 {
		t.Fatal("custom callback fallback changed")
	}
	for _, text := range []string{`{"version":false,"ignored":1e400}`, `{"events":[{"tag":"wrong"}],"ignored":"\q"}`, `{} trailing`} {
		got := textSchemaPublic(text)
		want := textSchemaFingerprint(rt.Apply(Get_Data_Argonaut_Decode_Parser_parseJson(), rt.Str(text)))
		if got != want {
			t.Fatalf("syntax error lost precedence: %s != %s", got, want)
		}
	}
}
