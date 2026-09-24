package purescript

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"gopurs/output/gopurs_runtime"
)

var zzCoverage = map[string]int{}

func zzCountSize(name string, size int) {
	zzCoverage[name+"_"+strconv.Itoa(size)]++
}

// This binary only counts paths. No timings from it enter a performance table.
func ZzCoverageMain() {
	var files []struct {
		Name      string `json:"name"`
		Contents  string `json:"contents"`
		Benchmark bool   `json:"benchmark"`
	}
	var oracle struct {
		Fingerprints []string `json:"fingerprints"`
	}
	read := func(path string, target any) {
		data, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(data, target); err != nil {
			panic(err)
		}
	}
	read(os.Getenv("DIAG_CORPUS"), &files)
	read(os.Getenv("DIAG_ORACLE"), &oracle)
	parse, decode, fingerprint := Get_Test_JsonDecoding_parse(), Get_Test_JsonDecoding_decode(), Get_Test_JsonDecoding_fingerprint()
	parsed := make([]gopurs_runtime.Value, len(files))
	for i, file := range files {
		parsed[i] = gopurs_runtime.Apply(parse, gopurs_runtime.Str(file.Contents))
		result := gopurs_runtime.Apply(decode, parsed[i])
		if canonicalHash(gopurs_runtime.Apply(fingerprint, result).StrVal()) != oracle.Fingerprints[i] {
			panic("coverage setup differs from oracle: " + file.Name)
		}
	}
	counts := map[string]map[string]int{}
	total := map[string]int{}
	retained := make([]gopurs_runtime.Value, len(files))
	for i, file := range files {
		if !file.Benchmark {
			continue
		}
		zzCoverage = map[string]int{}
		retained[i] = gopurs_runtime.Apply(decode, parsed[i])
		counts[file.Name] = zzCoverage
		for key, count := range zzCoverage {
			total[key] += count
		}
	}
	for i, file := range files {
		if file.Benchmark && canonicalHash(gopurs_runtime.Apply(fingerprint, retained[i]).StrVal()) != oracle.Fingerprints[i] {
			panic("retained coverage result differs from oracle: " + file.Name)
		}
	}
	data, err := json.Marshal(map[string]any{"cases": counts, "total": total, "timing": false})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
