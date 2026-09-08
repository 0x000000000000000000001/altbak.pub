// Inject this harness only into frozen copies of the generated Go module.
// The kernel is the existing Test.StateMonad.runManyTimes used by its act:
// 20 evaluations of the original 60-deep chain, with accumulator zero.
package purescript

import (
	"fmt"
	rt "gopurs/output/gopurs_runtime"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"testing"
)

var coreSink int64
var actSink string

var sizes = []struct {
	name        string
	n, expected int64
}{
	{"n20", 20, 1200},
}

// These outputs were checked for both variants and the real application in
// step 3.5. A deterministic but wrong A/B result must not become the reference.
var acts = []struct {
	name     string
	get      func() rt.Value
	expected string
}{
	{"AstTree", Get_Test_AstTree_act, "7"},
	{"Fib", Get_Test_Fib_act, "55"},
	{"ListOps", Get_Test_ListOps_act, "202950"},
	{"TCO", Get_Test_TCO_act, "100000"},
	{"Records", Get_Test_Records_act, "20000"},
	{"Ackermann", Get_Test_Ackermann_act, "125"},
	{"Church", Get_Test_Church_act, "100000"},
	{"Primes", Get_Test_Primes_act, "21536"},
	{"RBTree", Get_Test_RBTree_act, "22"},
	{"Polymorphism", Get_Test_Polymorphism_act, "10000000"},
	{"StateMonad", Get_Test_StateMonad_act, "1200"},
	{"LazyEvaluation", Get_Test_LazyEvaluation_act, "1000000"},
	{"ArrayOps", Get_Test_ArrayOps_act, "202950"},
	{"RowToList", Get_Test_RowToList_act, "5"},
}

func TestResults(t *testing.T) {
	for _, s := range sizes {
		for i := 0; i < 6; i++ {
			if got := Call_Test_StateMonad_runManyTimes(s.n, 0); got != s.expected {
				t.Fatalf("%s: got %d, want %d", s.name, got, s.expected)
			}
		}
		fmt.Printf("RESULT core %s %d\n", s.name, s.expected)
	}
	for _, tc := range acts {
		act := tc.get()
		for i := 0; i < 6; i++ {
			if got := rt.Apply(act, rt.Value{}).StrVal(); got != tc.expected {
				t.Fatalf("%s: got %q, want %q", tc.name, got, tc.expected)
			}
		}
		fmt.Printf("RESULT act %s %q\n", tc.name, tc.expected)
	}
}

func BenchmarkStateMonad(b *testing.B) {
	for _, s := range sizes {
		b.Run(s.name, func(b *testing.B) {
			for i := 0; i < 6; i++ {
				coreSink = Call_Test_StateMonad_runManyTimes(s.n, 0)
			}
			if coreSink != s.expected {
				b.Fatal("wrong warmup result")
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				coreSink = Call_Test_StateMonad_runManyTimes(s.n, 0)
			}
			b.StopTimer()
			if coreSink != s.expected {
				b.Fatal("wrong measured result")
			}
		})
	}
}

func BenchmarkAct(b *testing.B) {
	for _, tc := range acts {
		b.Run(tc.name, func(b *testing.B) {
			act := tc.get()
			for i := 0; i < 6; i++ {
				actSink = rt.Apply(act, rt.Value{}).StrVal()
			}
			if actSink != tc.expected {
				b.Fatal("wrong warmup result")
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				actSink = rt.Apply(act, rt.Value{}).StrVal()
			}
			b.StopTimer()
			if actSink != tc.expected {
				b.Fatal("wrong measured result")
			}
		})
	}
}

// Run in a fresh process under /usr/bin/time -l, with the same positive
// GOPURS_RSS_ITERATIONS in A and B and GOPURS_RSS_SIZE=n20. This fixed work
// avoids adaptive benchmark calibration changing the amount of RSS work.
// Peak RSS includes runtime/test-runner initialization and six warmups;
// it is not the peak of the full altbak application.
func TestFixedRSS(t *testing.T) {
	size := os.Getenv("GOPURS_RSS_SIZE")
	if size == "" {
		t.Skip("RSS workload not requested")
	}
	if os.Getenv("GOPURS_ALLOC_PROFILE") != "" {
		t.Fatal("run RSS and allocation profiling in separate processes")
	}
	repetitions, err := strconv.Atoi(os.Getenv("GOPURS_RSS_ITERATIONS"))
	if err != nil || repetitions < 1 {
		t.Fatal("invalid RSS iteration count")
	}
	for _, s := range sizes {
		if s.name != size {
			continue
		}
		for i := 0; i < 6; i++ {
			coreSink = Call_Test_StateMonad_runManyTimes(s.n, 0)
			if coreSink != s.expected {
				t.Fatal("wrong RSS warmup result")
			}
		}
		for i := 0; i < repetitions; i++ {
			coreSink = Call_Test_StateMonad_runManyTimes(s.n, 0)
			if coreSink != s.expected {
				t.Fatal("wrong RSS workload result")
			}
		}
		fmt.Printf("RSS workload %s iterations=%d result=%d\n", size, repetitions, coreSink)
		return
	}
	t.Fatal("unknown RSS size")
}

// Separate diagnostic run only, never use its elapsed time as a benchmark.
// Select -test.run=^TestAllocationProfile$ -test.memprofilerate=1, set
// GOPURS_ALLOC_PROFILE to the output path and GOPURS_ALLOC_ITERATIONS to a
// positive fixed count. Both profiles are cumulative; subtract PATH.before
// from PATH and focus on generated StateMonad/runtime allocation stacks.
// Profile serialization itself can appear outside those stacks. The printed
// MemStats delta isolates hot kernel calls, excluding serialization and the
// explicit GC calls; automatic GC activity during the kernel remains included.
func TestAllocationProfile(t *testing.T) {
	path := os.Getenv("GOPURS_ALLOC_PROFILE")
	if path == "" {
		t.Skip("allocation profile not requested")
	}
	if os.Getenv("GOPURS_RSS_SIZE") != "" {
		t.Fatal("run RSS and allocation profiling in separate processes")
	}
	if runtime.MemProfileRate != 1 {
		t.Fatal("exact allocation profile requires -test.memprofilerate=1")
	}
	repetitions, err := strconv.Atoi(os.Getenv("GOPURS_ALLOC_ITERATIONS"))
	if err != nil || repetitions < 1 {
		t.Fatal("invalid allocation profile iteration count")
	}
	s := sizes[0]
	for i := 0; i < 6; i++ {
		coreSink = Call_Test_StateMonad_runManyTimes(s.n, 0)
		if coreSink != s.expected {
			t.Fatal("wrong allocation profile warmup result")
		}
	}
	writeProfile := func(name string) {
		f, err := os.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		writeErr := pprof.Lookup("allocs").WriteTo(f, 0)
		closeErr := f.Close()
		if writeErr != nil {
			t.Fatal(writeErr)
		}
		if closeErr != nil {
			t.Fatal(closeErr)
		}
	}
	runtime.GC()
	runtime.GC()
	writeProfile(path + ".before")
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	for i := 0; i < repetitions; i++ {
		coreSink = Call_Test_StateMonad_runManyTimes(s.n, 0)
		if coreSink != s.expected {
			t.Fatal("wrong allocation profile result")
		}
	}
	runtime.ReadMemStats(&after)
	runtime.GC()
	runtime.GC()
	writeProfile(path)
	bytes := after.TotalAlloc - before.TotalAlloc
	allocs := after.Mallocs - before.Mallocs
	fmt.Printf("ALLOC StateMonad iterations=%d result=%d bytes=%d allocs=%d gcs=%d bytes/op=%.2f allocs/op=%.2f\n",
		repetitions, coreSink, bytes, allocs, after.NumGC-before.NumGC,
		float64(bytes)/float64(repetitions), float64(allocs)/float64(repetitions))
}
