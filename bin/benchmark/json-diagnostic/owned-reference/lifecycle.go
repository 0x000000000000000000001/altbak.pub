package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"runtime"
	"time"

	rt "gopurs/output/gopurs_runtime"
	p "gopurs/output/purescript"
)

type sample struct {
	TimeUS         float64 `json:"time_us"`
	AllocatedBytes uint64  `json:"allocated_bytes"`
}

func measure(action func()) sample {
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	start := time.Now()
	action()
	elapsed := time.Since(start).Nanoseconds()
	runtime.ReadMemStats(&after)
	return sample{float64(elapsed) / 1000, after.TotalAlloc - before.TotalAlloc}
}

func main() {
	var files []struct {
		Name, Contents string
		Benchmark      bool
	}
	bytes, err := os.ReadFile(os.Getenv("DIAG_CORPUS"))
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bytes, &files); err != nil {
		panic(err)
	}
	inputs, names := []string{}, []string{}
	for _, file := range files {
		if file.Benchmark {
			inputs = append(inputs, file.Contents)
			names = append(names, file.Name)
		}
	}
	// Clear input-loader scratch before constructing any decoder or result.
	runtime.GC()
	var combined func(string) rt.Value
	construction := measure(func() {
		// BUILD_DECODER
	})
	first, second := make([]rt.Value, len(inputs)), make([]rt.Value, len(inputs))
	pass := func(results []rt.Value) sample {
		return measure(func() {
			for i, input := range inputs {
				results[i] = combined(input)
			}
		})
	}
	firstSample, secondSample := pass(first), pass(second)
	// Charge deferred parser/decoder garbage before creating any fingerprint
	// scratch. Both complete batches remain live through this timed collection.
	retainedCollection := measure(runtime.GC)
	var retained, released runtime.MemStats
	runtime.ReadMemStats(&retained)
	fingerprint := p.Get_Test_JsonDecoding_fingerprint()
	hashes := func(results []rt.Value) []string {
		out := make([]string, len(results))
		for i, result := range results {
			var value any
			if err := json.Unmarshal([]byte(rt.Apply(fingerprint, result).StrVal()), &value); err != nil {
				panic(err)
			}
			bytes, err := json.Marshal(value)
			if err != nil {
				panic(err)
			}
			hash := sha256.Sum256(bytes)
			out[i] = hex.EncodeToString(hash[:])
		}
		return out
	}
	firstHashes, secondHashes := hashes(first), hashes(second)
	// Only validation scratch is now excluded. Decoding garbage was charged
	// above, and result reclamation is charged below after dropping both roots.
	runtime.GC()
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)
	first, second = nil, nil
	drain := measure(runtime.GC)
	runtime.ReadMemStats(&released)
	if err := json.NewEncoder(os.Stdout).Encode(map[string]any{
		"names": names, "construction": construction,
		"first_combined": firstSample, "second_combined": secondSample,
		"retained_collection": retainedCollection, "release_collection": drain,
		"retained_heap_bytes": retained.HeapAlloc, "released_heap_bytes": released.HeapAlloc,
		"first_fingerprints": firstHashes, "second_fingerprints": secondHashes,
	}); err != nil {
		panic(err)
	}
	runtime.KeepAlive(inputs)
	runtime.KeepAlive(combined)
}
