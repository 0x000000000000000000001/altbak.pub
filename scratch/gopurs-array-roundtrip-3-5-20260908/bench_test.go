// Measurement harness injected into frozen copies of the generated Go module.
package purescript

import (
	"fmt"
	rt "gopurs/output/gopurs_runtime"
	"os"
	"strconv"
	"testing"
)

var coreSink int64
var actSink string

var sizes = []struct {
	name        string
	n, expected int64
}{
	{"n900", 900, 202950}, {"n90000", 90000, 2025045000},
}

var acts = []struct {
	name string
	get  func() rt.Value
}{
	{"AstTree", Get_Test_AstTree_act}, {"Fib", Get_Test_Fib_act},
	{"ListOps", Get_Test_ListOps_act}, {"TCO", Get_Test_TCO_act},
	{"Records", Get_Test_Records_act}, {"Ackermann", Get_Test_Ackermann_act},
	{"Church", Get_Test_Church_act}, {"Primes", Get_Test_Primes_act},
	{"RBTree", Get_Test_RBTree_act}, {"Polymorphism", Get_Test_Polymorphism_act},
	{"StateMonad", Get_Test_StateMonad_act}, {"LazyEvaluation", Get_Test_LazyEvaluation_act},
	{"ArrayOps", Get_Test_ArrayOps_act}, {"RowToList", Get_Test_RowToList_act},
}

func TestResults(t *testing.T) {
	for _, s := range sizes {
		for i := 0; i < 6; i++ {
			if got := Call_Test_ArrayOps_sumEvens(s.n); got != s.expected {
				t.Fatalf("%s: got %d, want %d", s.name, got, s.expected)
			}
		}
		fmt.Printf("RESULT core %s %d\n", s.name, s.expected)
	}
	for _, tc := range acts {
		act := tc.get()
		expected := rt.Apply(act, rt.Value{}).StrVal()
		for i := 0; i < 6; i++ {
			if got := rt.Apply(act, rt.Value{}).StrVal(); got != expected {
				t.Fatalf("%s: unstable result", tc.name)
			}
		}
		fmt.Printf("RESULT act %s %q\n", tc.name, expected)
	}
}

func BenchmarkSumEvens(b *testing.B) {
	for _, s := range sizes {
		b.Run(s.name, func(b *testing.B) {
			for i := 0; i < 6; i++ {
				coreSink = Call_Test_ArrayOps_sumEvens(s.n)
			}
			if coreSink != s.expected {
				b.Fatal("wrong warmup result")
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				coreSink = Call_Test_ArrayOps_sumEvens(s.n)
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
			expected := rt.Apply(act, rt.Value{}).StrVal()
			for i := 0; i < 6; i++ {
				actSink = rt.Apply(act, rt.Value{}).StrVal()
			}
			if actSink != expected {
				b.Fatal("wrong warmup result")
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				actSink = rt.Apply(act, rt.Value{}).StrVal()
			}
			b.StopTimer()
			if actSink != expected {
				b.Fatal("wrong measured result")
			}
		})
	}
}

// Fixed work per fresh process: adaptive benchmark calibration is unsuitable
// for comparing peak RSS because it performs different operation counts.
func TestFixedRSS(t *testing.T) {
	size := os.Getenv("GOPURS_RSS_SIZE")
	if size == "" {
		t.Skip("RSS workload not requested")
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
			coreSink = Call_Test_ArrayOps_sumEvens(s.n)
		}
		for i := 0; i < repetitions; i++ {
			coreSink = Call_Test_ArrayOps_sumEvens(s.n)
		}
		if coreSink != s.expected {
			t.Fatal("wrong RSS workload result")
		}
		fmt.Printf("RSS workload %s iterations=%d result=%d\n", size, repetitions, coreSink)
		return
	}
	t.Fatal("unknown RSS size")
}
