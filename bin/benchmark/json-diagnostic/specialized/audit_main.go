// Audit-only A/B runner: generated Argonaut decoder against the hand-written
// schema-specialised decoder, in the same process, alternating passes.
//
// Copied into the audit workspace's output/purescript/ next to decode.go.
// It validates both decoders against the frozen oracle before timing, and
// validates every timed pass outside the timed interval, like the official
// diagnostic runner.
package purescript

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"

	"gopurs/output/gopurs_runtime"
)

type zzCorpusFile struct {
	Name      string `json:"name"`
	Contents  string `json:"contents"`
	Benchmark bool   `json:"benchmark"`
}

type zzOracle struct {
	Modules          int      `json:"modules"`
	TimedCases       int      `json:"timed_cases"`
	Names            []string `json:"names"`
	Fingerprints     []string `json:"fingerprints"`
	JSONFingerprints []string `json:"json_fingerprints"`
}

type zzSample struct {
	TimeUS     float64 `json:"time_us"`
	AllocBytes float64 `json:"allocated_bytes"`
}

type zzPhaseResult struct {
	TimeUS float64    `json:"time_us"`
	Samples []zzSample `json:"samples"`
}

type zzMode struct {
	Name   string
	Decode gopurs_runtime.Value
}

func zzReadJSON(path string, target any) {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(raw, target); err != nil {
		panic(err)
	}
}

func zzRunSegment(phase string, decode, parse gopurs_runtime.Value, files []zzCorpusFile, parsed []gopurs_runtime.Value, indices []int) (float64, float64) {
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	start := time.Now()
	for _, i := range indices {
		switch phase {
		case "parse":
			diagnosticSink = gopurs_runtime.Apply(parse, gopurs_runtime.Str(files[i].Contents))
		case "decode":
			diagnosticSink = gopurs_runtime.Apply(decode, parsed[i])
		case "combined":
			diagnosticSink = gopurs_runtime.Apply(decode, gopurs_runtime.Apply(parse, gopurs_runtime.Str(files[i].Contents)))
		default:
			panic("unknown phase " + phase)
		}
	}
	elapsed := float64(time.Since(start).Nanoseconds()) / 1000.0
	runtime.ReadMemStats(&after)
	return elapsed, float64(after.TotalAlloc - before.TotalAlloc)
}

// ZzAuditMain runs the A/B campaign and prints one JSON report.
func ZzAuditMain() {
	var files []zzCorpusFile
	zzReadJSON(os.Getenv("DIAG_CORPUS"), &files)
	var oracle zzOracle
	zzReadJSON(os.Getenv("DIAG_ORACLE"), &oracle)

	parse := Get_Test_JsonDecoding_parse()
	fingerprint := Get_Test_JsonDecoding_fingerprint()
	modes := []zzMode{
		{Name: "generated", Decode: Get_Test_JsonDecoding_decode()},
		{Name: "specialized", Decode: gopurs_runtime.Func(func(json gopurs_runtime.Value) gopurs_runtime.Value {
			return ZzSpecializedDecode(json)
		})},
	}

	parsed := make([]gopurs_runtime.Value, len(files))
	for i, file := range files {
		parsed[i] = gopurs_runtime.Apply(parse, gopurs_runtime.Str(file.Contents))
	}

	names := make([]string, len(files))
	jsonHashes := make([]string, len(files))
	hashes := map[string][]string{}
	for i, file := range files {
		names[i] = file.Name
		jsonHashes[i] = canonicalHash(file.Contents)
		for _, mode := range modes {
			decoded := gopurs_runtime.Apply(mode.Decode, parsed[i])
			hashes[mode.Name] = append(hashes[mode.Name], canonicalHash(gopurs_runtime.Apply(fingerprint, decoded).StrVal()))
		}
	}
	check := func(name string, got, want []string) {
		if len(got) != len(want) {
			panic(fmt.Sprintf("%s: length mismatch: %d vs %d", name, len(got), len(want)))
		}
		for i := range got {
			if got[i] != want[i] {
				panic(fmt.Sprintf("%s[%d] (%s): %s != oracle %s", name, i, names[i], got[i], want[i]))
			}
		}
	}
	if len(files) != oracle.Modules {
		panic("corpus module count does not match the oracle")
	}
	check("json", jsonHashes, oracle.JSONFingerprints)
	for _, mode := range modes {
		check(mode.Name, hashes[mode.Name], oracle.Fingerprints)
	}
	if len(names) != len(oracle.Names) {
		panic("corpus names do not match the oracle")
	}
	for i := range names {
		if names[i] != oracle.Names[i] {
			panic("corpus order does not match the oracle")
		}
	}

	indices := []int{}
	for i, file := range files {
		if file.Benchmark {
			indices = append(indices, i)
		}
	}
	if len(indices) != oracle.TimedCases {
		panic("timed case count does not match the oracle")
	}

	warmups := envInt("DIAG_WARMUPS", 2)
	samples := envInt("DIAG_SAMPLES", 5)
	phases := []string{}
	for _, name := range splitPhases(os.Getenv("DIAG_PHASES")) {
		phases = append(phases, name)
	}
	if len(phases) == 0 {
		phases = []string{"decode", "combined"}
	}

	report := map[string]any{
		"audit":             "specialized-decoder",
		"backend":           "go",
		"modules":           len(files),
		"timed_cases":       len(indices),
		"names":             names,
		"json_fingerprints": jsonHashes,
		"fingerprints":      hashes,
		"gomaxprocs":        runtime.GOMAXPROCS(0),
		"go":                runtime.Version(),
	}
	phaseReport := map[string]any{}
	for _, phase := range phases {
		results := map[string]any{}
		collected := map[string][]zzSample{}
		for _, mode := range modes {
			collected[mode.Name] = []zzSample{}
		}
		for pass := 0; pass < warmups+samples; pass++ {
			for k := 0; k < len(modes); k++ {
				index := k
				if pass%2 == 1 {
					index = len(modes) - 1 - k
				}
				active := modes[index]
				elapsed, alloc := zzRunSegment(phase, active.Decode, parse, files, parsed, indices)
				if pass >= warmups {
					collected[active.Name] = append(collected[active.Name], zzSample{TimeUS: elapsed, AllocBytes: alloc})
				}
			}
		}
		for _, mode := range modes {
			all := collected[mode.Name]
			best := 1e100
			for _, sample := range all {
				if sample.TimeUS < best {
					best = sample.TimeUS
				}
			}
			results[mode.Name] = zzPhaseResult{TimeUS: best, Samples: all}
		}
		phaseReport[phase] = results
	}
	report["phases"] = phaseReport
	report["warmups"] = warmups
	report["samples"] = samples
	data, err := json.Marshal(report)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func splitPhases(requested string) []string {
	result := []string{}
	if requested == "" {
		return result
	}
	start := 0
	for i := 0; i <= len(requested); i++ {
		if i == len(requested) || requested[i] == ',' {
			if i > start {
				result = append(result, requested[start:i])
			}
			start = i + 1
		}
	}
	return result
}
