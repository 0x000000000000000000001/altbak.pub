package JsonDecoding

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"

	"gopurs/output/gopurs_runtime"
)

var diagnosticSink gopurs_runtime.Value

func canonicalHash(text string) string {
	var value any
	if err := json.Unmarshal([]byte(text), &value); err != nil {
		panic(err)
	}
	// The JS driver uses the same HTML and line-separator escaping.
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}

func Drive(parse, decode, encode, fingerprint gopurs_runtime.Value) func() any {
	return func() any {
		bytes, err := os.ReadFile(os.Getenv("DIAG_CORPUS"))
		if err != nil {
			panic(err)
		}
		var files []struct {
			Name      string `json:"name"`
			Contents  string `json:"contents"`
			Benchmark bool   `json:"benchmark"`
		}
		if err = json.Unmarshal(bytes, &files); err != nil {
			panic(err)
		}
		parsed := make([]gopurs_runtime.Value, len(files))
		expectedJSON := make([]string, len(files))
		expectedAST := make([]string, len(files))
		for i, f := range files {
			parsed[i] = gopurs_runtime.Apply(parse, gopurs_runtime.Str(f.Contents))
			expectedJSON[i] = canonicalHash(f.Contents)
			result := gopurs_runtime.Apply(decode, parsed[i])
			raw := gopurs_runtime.Apply(fingerprint, result).StrVal()
			expectedAST[i] = canonicalHash(raw)
		}
		// Error cases are validated against the fixed oracle, outside timed work.
		indices := []int{}
		names := make([]string, len(files))
		for i, f := range files {
			names[i] = f.Name
			if f.Benchmark {
				indices = append(indices, i)
			}
		}
		if len(indices) == 0 {
			panic("No timed cases")
		}
		phases := map[string]any{}
		for _, phase := range []string{"parse", "decode", "combined"} {
			samples := []map[string]float64{}
			best := 1e100
			for pass := 0; pass < 7; pass++ {
				results := make([]gopurs_runtime.Value, len(indices))
				var before, after runtime.MemStats
				runtime.ReadMemStats(&before)
				start := time.Now()
				for slot, i := range indices {
					f := files[i]
					switch phase {
					case "parse":
						results[slot] = gopurs_runtime.Apply(parse, gopurs_runtime.Str(f.Contents))
					case "decode":
						results[slot] = gopurs_runtime.Apply(decode, parsed[i])
					case "combined":
						results[slot] = gopurs_runtime.Apply(decode, gopurs_runtime.Apply(parse, gopurs_runtime.Str(f.Contents)))
					}
					diagnosticSink = results[slot]
				}
				elapsed := float64(time.Since(start).Nanoseconds()) / 1000
				runtime.ReadMemStats(&after)
				for slot, result := range results {
					i := indices[slot]
					callback := fingerprint
					expected := expectedAST[i]
					if phase == "parse" {
						callback = encode
						expected = expectedJSON[i]
					}
					if canonicalHash(gopurs_runtime.Apply(callback, result).StrVal()) != expected {
						panic("Unstable " + phase + " output")
					}
				}
				if pass >= 2 {
					if elapsed < best {
						best = elapsed
					}
					samples = append(samples, map[string]float64{"time_us": elapsed, "allocated_bytes": float64(after.TotalAlloc - before.TotalAlloc)})
				}
			}
			phases[phase] = map[string]any{"samples": samples, "time_us": best}
		}
		report := map[string]any{"backend": "go", "modules": len(files), "timed_cases": len(indices), "names": names, "fingerprints": expectedAST, "json_fingerprints": expectedJSON, "phases": phases, "go": runtime.Version(), "gomaxprocs": runtime.GOMAXPROCS(0)}
		result, err := json.Marshal(report)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(result))
		return nil
	}
}
