package JsonDecoding

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
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

func envInt(name string, fallback int) int {
	text := os.Getenv(name)
	if text == "" {
		return fallback
	}
	value, err := strconv.Atoi(text)
	if err != nil || value < 0 {
		panic("invalid " + name)
	}
	return value
}

// phaseList selects the measured phases. DIAG_PHASES is reserved for focused
// runs; a profile run selects its single phase through DIAG_PROFILE_PHASE.
func phaseList() []string {
	requested := os.Getenv("DIAG_PHASES")
	if requested == "" {
		return []string{"parse", "decode", "combined"}
	}
	phases := []string{}
	for _, name := range strings.Split(requested, ",") {
		phases = append(phases, strings.TrimSpace(name))
	}
	return phases
}

func Drive(parse, decode, decodeText, encode, fingerprint gopurs_runtime.Value) func() any {
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
		// runPass executes one timed pass and optionally validates the produced
		// results against the oracle. Validation is outside the timed interval
		// but part of a plain process profile, so profile runs leave it to the
		// single pass that precedes the CPU profile.
		runPass := func(phase string, validate bool) (float64, float64) {
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
					results[slot] = gopurs_runtime.Apply(decodeText, gopurs_runtime.Str(f.Contents))
				default:
					panic("unknown phase " + phase)
				}
				diagnosticSink = results[slot]
			}
			elapsed := float64(time.Since(start).Nanoseconds()) / 1000
			runtime.ReadMemStats(&after)
			if validate {
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
			}
			return elapsed, float64(after.TotalAlloc - before.TotalAlloc)
		}
		profilePhase := os.Getenv("DIAG_PROFILE_PHASE")
		phases := phaseList()
		if profilePhase != "" {
			phases = []string{profilePhase}
		}
		phaseReport := map[string]any{}
		for _, phase := range phases {
			switch phase {
			case "parse", "decode", "combined":
			default:
				panic("unknown phase " + phase)
			}
			if profilePhase != "" {
				// One validating pass outside the profile, then many unvalidated
				// passes so the samples describe the timed work, not the oracle.
				runPass(phase, true)
				var cpuFile *os.File
				if cpuPath := os.Getenv("DIAG_CPU_PROFILE"); cpuPath != "" {
					cpuFile, err = os.Create(cpuPath)
					if err != nil {
						panic(err)
					}
					if err = pprof.StartCPUProfile(cpuFile); err != nil {
						panic(err)
					}
				}
				passes := envInt("DIAG_PROFILE_PASSES", 200)
				best := 1e100
				for pass := 0; pass < passes; pass++ {
					elapsed, _ := runPass(phase, false)
					if elapsed < best {
						best = elapsed
					}
				}
				if cpuFile != nil {
					pprof.StopCPUProfile()
					cpuFile.Close()
				}
				if memPath := os.Getenv("DIAG_MEM_PROFILE"); memPath != "" {
					memFile, err := os.Create(memPath)
					if err != nil {
						panic(err)
					}
					if err = pprof.WriteHeapProfile(memFile); err != nil {
						panic(err)
					}
					memFile.Close()
				}
				phaseReport[phase] = map[string]any{"samples": []map[string]float64{}, "time_us": best}
				continue
			}
			samples := []map[string]float64{}
			best := 1e100
			for pass := 0; pass < 7; pass++ {
				elapsed, allocated := runPass(phase, true)
				if pass >= 2 {
					if elapsed < best {
						best = elapsed
					}
					samples = append(samples, map[string]float64{"time_us": elapsed, "allocated_bytes": allocated})
				}
			}
			phaseReport[phase] = map[string]any{"samples": samples, "time_us": best}
		}
		report := map[string]any{"backend": "go", "modules": len(files), "timed_cases": len(indices), "names": names, "fingerprints": expectedAST, "json_fingerprints": expectedJSON, "phases": phaseReport, "go": runtime.Version(), "gomaxprocs": runtime.GOMAXPROCS(0)}
		if profilePhase != "" {
			report["profile_phase"] = profilePhase
		}
		result, err := json.Marshal(report)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(result))
		return nil
	}
}
