// Hand-written Go decoder for the JSON Decoding diagnostic schema.
//
// Audit purpose: place a point between the gopurs-generated decoder (generic
// Argonaut machinery, boxed values) and the hand-written C++ reference
// (simdjson + arena), to separate "Go as a language/representation" from "the
// current generic decoding path".
//
// Modes:
//   dom  - one json.Unmarshal into any per document (parse phase), then a
//          hand-written typed walk over that DOM (decode phase). Mirrors the
//          C++ driver's two-phase structure and the gopurs architecture.
//          -clone copies field strings during the walk, mirroring the C++
//          arena copy; without it the walk reuses the DOM strings, mirroring
//          the Go/Argonaut lifetime.
//   std  - encoding/json directly into concrete structs (what a Go
//          application writes); one step, reported as "combined".
//
// Fingerprints use the same canonical serialization as the C++ reference
// (decoding.cc) so the frozen oracle in expected.json applies unchanged.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------------------------
// Concrete structures (mirror decoding.cc)
// ---------------------------------------------------------------------------

type profile struct {
	city   string
	note   *string
	scores []float64
}

type user struct {
	id      int64
	name    string
	active  bool
	tags    []string
	profile *profile
}

type item struct {
	sku      string
	quantity int64
	price    float64
}

type event struct {
	kind        int // 0 = view, 1 = purchase
	path        string
	hasDuration bool
	duration    int64
	orderID     int64
	items       []item
}

type payload struct {
	version int64
	next    *string
	users   []user
	events  []event
}

// ---------------------------------------------------------------------------
// Strict decoding from a decoded DOM (map[string]any / []any / ...)
// ---------------------------------------------------------------------------

type decodeError struct{}

func (decodeError) Error() string { return "decode error" }

func asObject(v any) map[string]any {
	if m, ok := v.(map[string]any); ok {
		return m
	}
	panic(decodeError{})
}

func asArray(v any) []any {
	if a, ok := v.([]any); ok {
		return a
	}
	panic(decodeError{})
}

func asString(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	panic(decodeError{})
}

func asBool(v any) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	panic(decodeError{})
}

func asInt(v any) int64 {
	f, ok := v.(float64)
	if !ok {
		panic(decodeError{})
	}
	if math.IsNaN(f) || math.IsInf(f, 0) || math.Trunc(f) != f || math.Abs(f) >= 9.0e15 {
		panic(decodeError{})
	}
	return int64(f)
}

func asDouble(v any) float64 {
	f, ok := v.(float64)
	if !ok {
		panic(decodeError{})
	}
	return f
}

func requireField(obj map[string]any, key string) any {
	v, ok := obj[key]
	if !ok {
		panic(decodeError{})
	}
	return v
}

func optionalField(obj map[string]any, key string) (any, bool) {
	v, ok := obj[key]
	return v, ok
}

type stringCopier struct{ clone bool }

func (c stringCopier) copy(s string) string {
	if c.clone {
		return strings.Clone(s)
	}
	return s
}

func decodeProfileDOM(raw any, copier stringCopier) *profile {
	obj := asObject(raw)
	out := &profile{}
	out.city = copier.copy(asString(requireField(obj, "city")))
	scores := asArray(requireField(obj, "scores"))
	out.scores = make([]float64, len(scores))
	for i, score := range scores {
		out.scores[i] = asDouble(score)
	}
	if note, present := optionalField(obj, "note"); present && note != nil {
		text := copier.copy(asString(note))
		out.note = &text
	}
	return out
}

func decodeUserDOM(raw any, copier stringCopier) user {
	obj := asObject(raw)
	out := user{}
	out.id = asInt(requireField(obj, "id"))
	out.name = copier.copy(asString(requireField(obj, "name")))
	out.active = asBool(requireField(obj, "active"))
	tags := asArray(requireField(obj, "tags"))
	out.tags = make([]string, len(tags))
	for i, tag := range tags {
		out.tags[i] = copier.copy(asString(tag))
	}
	if profileRaw, present := optionalField(obj, "profile"); present && profileRaw != nil {
		out.profile = decodeProfileDOM(profileRaw, copier)
	}
	return out
}

func decodeItemDOM(raw any, copier stringCopier) item {
	obj := asObject(raw)
	return item{
		sku:      copier.copy(asString(requireField(obj, "sku"))),
		quantity: asInt(requireField(obj, "quantity")),
		price:    asDouble(requireField(obj, "price")),
	}
}

func decodeEventDOM(raw any, copier stringCopier) event {
	obj := asObject(raw)
	out := event{}
	tag := asString(requireField(obj, "tag"))
	switch tag {
	case "view":
		out.kind = 0
		out.path = copier.copy(asString(requireField(obj, "path")))
		if duration, present := optionalField(obj, "duration"); present && duration != nil {
			out.hasDuration = true
			out.duration = asInt(duration)
		}
	case "purchase":
		out.kind = 1
		out.orderID = asInt(requireField(obj, "orderId"))
		items := asArray(requireField(obj, "items"))
		out.items = make([]item, len(items))
		for i, raw := range items {
			out.items[i] = decodeItemDOM(raw, copier)
		}
	default:
		panic(decodeError{})
	}
	return out
}

func decodePayloadDOM(raw any, copier stringCopier) payload {
	obj := asObject(raw)
	out := payload{}
	out.version = asInt(requireField(obj, "version"))
	users := asArray(requireField(obj, "users"))
	out.users = make([]user, len(users))
	for i, raw := range users {
		out.users[i] = decodeUserDOM(raw, copier)
	}
	events := asArray(requireField(obj, "events"))
	out.events = make([]event, len(events))
	for i, raw := range events {
		out.events[i] = decodeEventDOM(raw, copier)
	}
	if next, present := optionalField(obj, "next"); present && next != nil {
		text := copier.copy(asString(next))
		out.next = &text
	}
	return out
}

// ---------------------------------------------------------------------------
// Canonical serialization (same conventions as the fixture oracle and the C++
// reference: sorted object keys, HTML escaping, integral numbers as integers).
// ---------------------------------------------------------------------------

func appendEscaped(out *strings.Builder, text string) {
	out.WriteByte('"')
	for index := 0; index < len(text); {
		b := text[index]
		if b == 0xe2 && index+2 < len(text) && text[index+1] == 0x80 {
			third := text[index+2]
			if third == 0xa8 || third == 0xa9 {
				if third == 0xa8 {
					out.WriteString("\\u2028")
				} else {
					out.WriteString("\\u2029")
				}
				index += 3
				continue
			}
		}
		index++
		switch b {
		case '"':
			out.WriteString("\\\"")
		case '\\':
			out.WriteString("\\\\")
		case '\n':
			out.WriteString("\\n")
		case '\r':
			out.WriteString("\\r")
		case '\t':
			out.WriteString("\\t")
		case '<':
			out.WriteString("\\u003c")
		case '>':
			out.WriteString("\\u003e")
		case '&':
			out.WriteString("\\u0026")
		default:
			if b < 0x20 {
				fmt.Fprintf(out, "\\u%04x", b)
			} else {
				out.WriteByte(b)
			}
		}
	}
	out.WriteByte('"')
}

func appendNumber(out *strings.Builder, value float64) {
	if math.IsInf(value, 0) || math.IsNaN(value) {
		out.WriteString("null")
		return
	}
	if math.Trunc(value) == value && math.Abs(value) < 9.0e15 {
		out.WriteString(strconv.FormatInt(int64(value), 10))
		return
	}
	for precision := 15; precision <= 17; precision++ {
		scratch := strconv.FormatFloat(value, 'g', precision, 64)
		if parsed, err := strconv.ParseFloat(scratch, 64); err == nil && parsed == value {
			out.WriteString(scratch)
			return
		}
	}
	out.WriteString(strconv.FormatFloat(value, 'g', 17, 64))
}

func appendCanonical(out *strings.Builder, value any) {
	switch typed := value.(type) {
	case map[string]any:
		keys := make([]string, 0, len(typed))
		for key := range typed {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		out.WriteByte('{')
		for index, key := range keys {
			if index > 0 {
				out.WriteByte(',')
			}
			appendEscaped(out, key)
			out.WriteByte(':')
			appendCanonical(out, typed[key])
		}
		out.WriteByte('}')
	case []any:
		out.WriteByte('[')
		for index, entry := range typed {
			if index > 0 {
				out.WriteByte(',')
			}
			appendCanonical(out, entry)
		}
		out.WriteByte(']')
	case string:
		appendEscaped(out, typed)
	case float64:
		appendNumber(out, typed)
	case json.Number:
		if parsed, err := strconv.ParseFloat(string(typed), 64); err == nil {
			appendNumber(out, parsed)
		} else {
			out.WriteString(string(typed))
		}
	case bool:
		if typed {
			out.WriteString("true")
		} else {
			out.WriteString("false")
		}
	case nil:
		out.WriteString("null")
	default:
		out.WriteString("null")
	}
}

func fingerprintOf(value any) string {
	var out strings.Builder
	appendCanonical(&out, value)
	sum := sha256.Sum256([]byte(out.String()))
	return hex.EncodeToString(sum[:])
}

func payloadFingerprint(p *payload) string {
	var out strings.Builder
	out.WriteString("{\"value\":{\"events\":[")
	for index := range p.events {
		if index > 0 {
			out.WriteByte(',')
		}
		ev := &p.events[index]
		if ev.kind == 0 {
			out.WriteString("{\"duration\":")
			if ev.hasDuration {
				out.WriteString(strconv.FormatInt(ev.duration, 10))
			} else {
				out.WriteString("null")
			}
			out.WriteString(",\"path\":")
			appendEscaped(&out, ev.path)
			out.WriteString(",\"tag\":\"view\"}")
			continue
		}
		out.WriteString("{\"items\":[")
		for itemIndex := range ev.items {
			if itemIndex > 0 {
				out.WriteByte(',')
			}
			it := &ev.items[itemIndex]
			out.WriteString("{\"price\":")
			appendNumber(&out, it.price)
			out.WriteString(",\"quantity\":")
			out.WriteString(strconv.FormatInt(it.quantity, 10))
			out.WriteString(",\"sku\":")
			appendEscaped(&out, it.sku)
			out.WriteByte('}')
		}
		out.WriteString("],\"orderId\":")
		out.WriteString(strconv.FormatInt(ev.orderID, 10))
		out.WriteString(",\"tag\":\"purchase\"}")
	}
	out.WriteString("],\"next\":")
	if p.next != nil {
		appendEscaped(&out, *p.next)
	} else {
		out.WriteString("null")
	}
	out.WriteString(",\"users\":[")
	for index := range p.users {
		if index > 0 {
			out.WriteByte(',')
		}
		u := &p.users[index]
		out.WriteString("{\"active\":")
		if u.active {
			out.WriteString("true")
		} else {
			out.WriteString("false")
		}
		out.WriteString(",\"id\":")
		out.WriteString(strconv.FormatInt(u.id, 10))
		out.WriteString(",\"name\":")
		appendEscaped(&out, u.name)
		out.WriteString(",\"profile\":")
		if u.profile != nil {
			out.WriteString("{\"city\":")
			appendEscaped(&out, u.profile.city)
			out.WriteString(",\"note\":")
			if u.profile.note != nil {
				appendEscaped(&out, *u.profile.note)
			} else {
				out.WriteString("null")
			}
			out.WriteString(",\"scores\":[")
			for scoreIndex, score := range u.profile.scores {
				if scoreIndex > 0 {
					out.WriteByte(',')
				}
				appendNumber(&out, score)
			}
			out.WriteString("]}")
		} else {
			out.WriteString("null")
		}
		out.WriteString(",\"tags\":[")
		for tagIndex, tag := range u.tags {
			if tagIndex > 0 {
				out.WriteByte(',')
			}
			appendEscaped(&out, tag)
		}
		out.WriteString("]}")
	}
	out.WriteString("],\"version\":")
	out.WriteString(strconv.FormatInt(p.version, 10))
	out.WriteString("}}")
	sum := sha256.Sum256([]byte(out.String()))
	return hex.EncodeToString(sum[:])
}

// ---------------------------------------------------------------------------
// standard-library direct struct decoding ("what an application writes")
// ---------------------------------------------------------------------------

type stdProfile struct {
	City   *string    `json:"city"`
	Note   *string    `json:"note"`
	Scores *[]float64 `json:"scores"`
}

type stdUser struct {
	ID      *int64      `json:"id"`
	Name    *string     `json:"name"`
	Active  *bool       `json:"active"`
	Tags    *[]string   `json:"tags"`
	Profile *stdProfile `json:"profile"`
}

type stdItem struct {
	SKU      *string  `json:"sku"`
	Quantity *int64   `json:"quantity"`
	Price    *float64 `json:"price"`
}

type stdEvent struct {
	Tag     *string    `json:"tag"`
	Path    *string    `json:"path"`
	Duration *int64    `json:"duration"`
	OrderID *int64     `json:"orderId"`
	Items   *[]stdItem `json:"items"`
}

type stdPayload struct {
	Version *int64      `json:"version"`
	Users   *[]stdUser  `json:"users"`
	Events  *[]stdEvent `json:"events"`
	Next    *string     `json:"next"`
}

func stdProfileOf(raw *stdProfile) *profile {
	if raw == nil {
		return nil
	}
	if raw.City == nil || raw.Scores == nil {
		panic(decodeError{})
	}
	out := &profile{city: *raw.City, note: raw.Note, scores: *raw.Scores}
	if out.note != nil {
		text := *out.note
		out.note = &text
	}
	return out
}

func stdPayloadOf(raw *stdPayload) payload {
	if raw.Version == nil || raw.Users == nil || raw.Events == nil {
		panic(decodeError{})
	}
	out := payload{version: *raw.Version, next: raw.Next}
	if out.next != nil {
		text := *out.next
		out.next = &text
	}
	out.users = make([]user, len(*raw.Users))
	for index, rawUser := range *raw.Users {
		if rawUser.ID == nil || rawUser.Name == nil || rawUser.Active == nil || rawUser.Tags == nil {
			panic(decodeError{})
		}
		entry := user{id: *rawUser.ID, name: *rawUser.Name, active: *rawUser.Active, tags: *rawUser.Tags}
		entry.profile = stdProfileOf(rawUser.Profile)
		out.users[index] = entry
	}
	out.events = make([]event, len(*raw.Events))
	for index, rawEvent := range *raw.Events {
		if rawEvent.Tag == nil {
			panic(decodeError{})
		}
		entry := event{}
		switch *rawEvent.Tag {
		case "view":
			if rawEvent.Path == nil {
				panic(decodeError{})
			}
			entry.kind = 0
			entry.path = *rawEvent.Path
			if rawEvent.Duration != nil {
				entry.hasDuration = true
				entry.duration = *rawEvent.Duration
			}
		case "purchase":
			if rawEvent.OrderID == nil || rawEvent.Items == nil {
				panic(decodeError{})
			}
			entry.kind = 1
			entry.orderID = *rawEvent.OrderID
			entry.items = make([]item, len(*rawEvent.Items))
			for itemIndex, rawItem := range *rawEvent.Items {
				if rawItem.SKU == nil || rawItem.Quantity == nil || rawItem.Price == nil {
					panic(decodeError{})
				}
				entry.items[itemIndex] = item{sku: *rawItem.SKU, quantity: *rawItem.Quantity, price: *rawItem.Price}
			}
		default:
			panic(decodeError{})
		}
		out.events[index] = entry
	}
	return out
}

// ---------------------------------------------------------------------------
// Corpus, protocol, report
// ---------------------------------------------------------------------------

type corpusCase struct {
	Name      string `json:"name"`
	Benchmark bool   `json:"benchmark"`
	Contents  string `json:"contents"`
}

type sample struct {
	AllocatedBytes int64   `json:"allocated_bytes"`
	TimeUS         float64 `json:"time_us"`
}

type phaseReport struct {
	Samples []sample `json:"samples"`
	TimeUS  float64  `json:"time_us"`
}

func nowUS() float64 { return float64(time.Now().UnixNano()) / 1000.0 }

func main() {
	mode := flag.String("mode", "dom", "dom | std | goccy | dom-clone | std-goccy")
	flag.Parse()

	corpusPath := os.Getenv("DIAG_CORPUS")
	if corpusPath == "" {
		fmt.Fprintln(os.Stderr, "DIAG_CORPUS is required")
		os.Exit(2)
	}
	raw, err := os.ReadFile(corpusPath)
	if err != nil {
		panic(err)
	}
	var cases []corpusCase
	if err := json.Unmarshal(raw, &cases); err != nil {
		panic(err)
	}

	useGoccy := strings.HasSuffix(*mode, "goccy")
	clone := *mode == "dom-clone"

	// --- oracle material: fingerprints for every case, outside timing -------
	names := make([]string, len(cases))
	fingerprints := make([]string, len(cases))
	jsonFingerprints := make([]string, len(cases))
	timed := []int{}
	domCache := make([]any, len(cases))
	failures := 0
	for index, entry := range cases {
		names[index] = entry.Name
		var dom any
		var parseErr error
		if useGoccy {
			parseErr = goccyUnmarshal([]byte(entry.Contents), &dom)
		} else {
			parseErr = json.Unmarshal([]byte(entry.Contents), &dom)
		}
		if parseErr != nil {
			fmt.Fprintf(os.Stderr, "parse failed: %s: %v\n", entry.Name, parseErr)
			os.Exit(3)
		}
		domCache[index] = dom
		jsonFingerprints[index] = fingerprintOf(dom)
		fp, ok := decodeFingerprint(*mode, entry.Contents, dom)
		if ok {
			fingerprints[index] = fp
		} else {
			failures++
		}
		if entry.Benchmark {
			timed = append(timed, index)
		}
	}
	if len(timed) == 0 {
		fmt.Fprintln(os.Stderr, "no timed cases")
		os.Exit(2)
	}

	// --- phases --------------------------------------------------------------
	phaseNames := []string{"parse", "decode", "combined"}
	if *mode == "std" || *mode == "goccy" || *mode == "std-goccy" {
		phaseNames = []string{"parse", "combined"}
	}
	report := map[string]phaseReport{}
	var sink any
	for _, phase := range phaseNames {
		var best float64
		haveBest := false
		samples := []sample{}
		for pass := 0; pass < 7; pass++ {
			var before, after runtime.MemStats
			runtime.ReadMemStats(&before)
			begin := nowUS()
			switch phase {
			case "parse":
				results := make([]any, len(timed))
				for slot, index := range timed {
					var dom any
					if useGoccy {
						if err := goccyUnmarshal([]byte(cases[index].Contents), &dom); err != nil {
							panic(err)
						}
					} else {
						if err := json.Unmarshal([]byte(cases[index].Contents), &dom); err != nil {
							panic(err)
						}
					}
					results[slot] = dom
				}
				sink = results
			case "decode":
				results := make([]payload, len(timed))
				for slot, index := range timed {
					results[slot] = decodePayloadDOM(domCache[index], stringCopier{clone: clone})
				}
				sink = results
			case "combined":
				results := make([]payload, len(timed))
				for slot, index := range timed {
					var p payload
					if *mode == "std" {
						var dom stdPayload
						if err := json.Unmarshal([]byte(cases[index].Contents), &dom); err != nil {
							panic(err)
						}
						p = stdPayloadOf(&dom)
					} else if *mode == "goccy" || *mode == "std-goccy" {
						var dom stdPayload
						if err := goccyUnmarshal([]byte(cases[index].Contents), &dom); err != nil {
							panic(err)
						}
						p = stdPayloadOf(&dom)
					} else {
						var dom any
						if err := json.Unmarshal([]byte(cases[index].Contents), &dom); err != nil {
							panic(err)
						}
						p = decodePayloadDOM(dom, stringCopier{clone: clone})
					}
					results[slot] = p
				}
				sink = results
			}
			elapsed := nowUS() - begin
			runtime.ReadMemStats(&after)
			if pass >= 2 {
				if !haveBest || elapsed < best {
					best = elapsed
					haveBest = true
				}
				samples = append(samples, sample{
					AllocatedBytes: int64(after.TotalAlloc - before.TotalAlloc),
					TimeUS:         elapsed,
				})
			}
		}
		report[phase] = phaseReport{Samples: samples, TimeUS: best}
	}
	runtime.KeepAlive(sink)
	_ = failures

	// --- report --------------------------------------------------------------
	var out strings.Builder
	out.WriteString("{\"backend\":\"native-go\",\"mode\":\"")
	out.WriteString(*mode)
	out.WriteString("\",\"modules\":")
	out.WriteString(strconv.Itoa(len(cases)))
	out.WriteString(",\"timed_cases\":")
	out.WriteString(strconv.Itoa(len(timed)))
	out.WriteString(",\"names\":")
	namesJSON, _ := json.Marshal(names)
	out.Write(namesJSON)
	out.WriteString(",\"fingerprints\":")
	fpJSON, _ := json.Marshal(fingerprints)
	out.Write(fpJSON)
	out.WriteString(",\"json_fingerprints\":")
	jfpJSON, _ := json.Marshal(jsonFingerprints)
	out.Write(jfpJSON)
	out.WriteString(",\"phases\":{")
	for index, phase := range phaseNames {
		if index > 0 {
			out.WriteByte(',')
		}
		phaseJSON, _ := json.Marshal(report[phase])
		out.WriteString(strconv.Quote(phase))
		out.WriteByte(':')
		out.Write(phaseJSON)
	}
	out.WriteString("},\"go\":")
	goJSON, _ := json.Marshal(runtime.Version())
	out.Write(goJSON)
	out.WriteString(",\"gomaxprocs\":")
	out.WriteString(strconv.Itoa(runtime.GOMAXPROCS(0)))
	out.WriteString(",\"failures\":")
	out.WriteString(strconv.Itoa(failures))
	out.WriteString("}")
	fmt.Println(out.String())
}

// decodeFingerprint runs the typed decode for one document and returns the
// payload fingerprint; ok=false when the document must fail to decode.
func decodeFingerprint(mode string, contents string, dom any) (string, bool) {
	ok := true
	var result payload
	func() {
		defer func() {
			if recovered := recover(); recovered != nil {
				ok = false
			}
		}()
		if mode == "std" || mode == "goccy" || mode == "std-goccy" {
			var decoded stdPayload
			var err error
			if strings.HasSuffix(mode, "goccy") {
				err = goccyUnmarshal([]byte(contents), &decoded)
			} else {
				err = json.Unmarshal([]byte(contents), &decoded)
			}
			if err != nil {
				ok = false
				return
			}
			result = stdPayloadOf(&decoded)
			return
		}
		result = decodePayloadDOM(dom, stringCopier{clone: false})
	}()
	if !ok {
		return "", false
	}
	return payloadFingerprint(&result), true
}
