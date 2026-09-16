package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	rt "gopurs/output/gopurs_runtime"
	ps "gopurs/output/purescript"
)

type sample struct {
	Mode           string `json:"mode"`
	Iteration      int    `json:"iteration"`
	Input          int64  `json:"input"`
	Depth          int64  `json:"depth"`
	Nanoseconds    int64  `json:"nanoseconds"`
	AllocatedBytes uint64 `json:"allocated_bytes"`
	Allocations    uint64 `json:"allocations"`
}

var treeSink *ps.Constructor_Test_RBTree_T
var depthSink int64

func run(mode string, count int64) int64 {
	if mode == "pure" {
		treeSink = buildFresh(count)
		return ps.Call_Test_RBTree_depth(treeSink)
	}
	return rt.Apply(ps.Get_Test_RBTreeFFICheatcode_runRBTreeFFICheatcode(), rt.Int(count)).IntVal
}

func main() {
	mode := flag.String("mode", "pure", "pure or fficc")
	count := flag.Int64("n", 100000, "number of distinct descending keys")
	iterations := flag.Int("iterations", 5, "number of measured calls")
	flag.Parse()
	if (*mode != "pure" && *mode != "fficc") || *count < 0 || *iterations < 1 {
		fmt.Fprintln(os.Stderr, "invalid mode, count, or iteration count")
		os.Exit(2)
	}
	// Warm the wrappers and both runtime and application paths before sampling.
	for i := 0; i < 3; i++ {
		depthSink = run(*mode, *count)
	}
	encoder := json.NewEncoder(os.Stdout)
	for i := 1; i <= *iterations; i++ {
		treeSink = nil
		runtime.GC()
		var before, after runtime.MemStats
		runtime.ReadMemStats(&before)
		start := time.Now()
		depthSink = run(*mode, *count)
		elapsed := time.Since(start)
		runtime.ReadMemStats(&after)
		runtime.KeepAlive(treeSink)
		if *count == 100000 && depthSink != 22 {
			fmt.Fprintf(os.Stderr, "unexpected depth: %d\n", depthSink)
			os.Exit(1)
		}
		if err := encoder.Encode(sample{*mode, i, *count, depthSink, elapsed.Nanoseconds(), after.TotalAlloc - before.TotalAlloc, after.Mallocs - before.Mallocs}); err != nil {
			panic(err)
		}
	}
}
