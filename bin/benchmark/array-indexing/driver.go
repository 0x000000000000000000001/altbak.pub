// Timing/input harness only: both indexed-read loops come from ArrayIndexing.purs.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"gopurs/output/gopurs_runtime"
	"gopurs/output/purescript"
)

var sink int64

func main() {
	accesses := flag.Int("accesses", 1<<23, "indexed reads per batch")
	batches := flag.Int("batches", 10, "measured batches per case")
	seed := flag.Int("seed", 5, "runtime input seed")
	flag.Parse()
	if *accesses < 16384 || *accesses > 1<<23 || *batches < 3 || *seed < 0 || *seed > 1000 {
		panic("invalid accesses, batches or seed")
	}
	encoder := json.NewEncoder(os.Stdout)
	for _, size := range []int{16, 1024, 16384} {
		native := make([]int64, size)
		boxedElements := make([]gopurs_runtime.Value, size)
		for i := range native {
			native[i] = int64((i*17+*seed*31)%251 + 1)
			boxedElements[i] = gopurs_runtime.Int(native[i])
		}
		boxed := gopurs_runtime.Array(boxedElements)
		start := *seed % size
		oracle := func(count int) int64 {
			var total int64
			for _, value := range native {
				total += value
			}
			result := total * int64(count/size)
			for i := 0; i < count%size; i++ {
				result += native[(start+i)%size]
			}
			return result
		}
		for _, representation := range []string{"native", "boxed"} {
			call := func(count int) int64 {
				if representation == "native" {
					return purescript.Call_Test_ArrayIndexing_nativeReads(native, int64(size), int64(count), int64(start))
				}
				return purescript.Call_Test_ArrayIndexing_boxedReads(boxed, int64(size), int64(count), int64(start))
			}
			// A bounded preflight catches the former whole-array copy before a
			// long batch could allocate terabytes on the largest input.
			sink = call(1)
			runtime.GC()
			var before, after runtime.MemStats
			runtime.ReadMemStats(&before)
			sink = call(1024)
			runtime.ReadMemStats(&after)
			probeBytes := float64(after.TotalAlloc-before.TotalAlloc) / 1024
			if sink != oracle(1024) || probeBytes > 256 {
				panic(fmt.Sprintf("indexing regression: %s size=%d checksum=%d bytes/access=%.3f", representation, size, sink, probeBytes))
			}
			expected := oracle(*accesses)
			for warm := 0; warm < 3; warm++ {
				sink = call(*accesses)
				if sink != expected {
					panic("warm-up checksum mismatch")
				}
			}
			nanoseconds := make([]float64, 0, *batches)
			allocated := make([]float64, 0, *batches)
			for batch := 0; batch < *batches; batch++ {
				runtime.ReadMemStats(&before)
				begin := time.Now()
				sink = call(*accesses)
				elapsed := time.Since(begin).Nanoseconds()
				runtime.ReadMemStats(&after)
				if sink != expected || elapsed <= 0 {
					panic("invalid measured result")
				}
				nanoseconds = append(nanoseconds, float64(elapsed)/float64(*accesses))
				allocated = append(allocated, float64(after.TotalAlloc-before.TotalAlloc)/float64(*accesses))
			}
			if err := encoder.Encode(map[string]any{
				"runtime": "go", "representation": representation, "size": size,
				"accesses": *accesses, "seed": *seed, "checksum": sink,
				"warmups": 3, "ns_per_access": nanoseconds, "bytes_per_access": allocated,
				"probe_bytes_per_access": probeBytes,
			}); err != nil {
				panic(err)
			}
		}
		runtime.KeepAlive(native)
		runtime.KeepAlive(boxed)
	}
}
