package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type Value struct {
	Type      uint8
	IntVal    int64
	UnsafePtr unsafe.Pointer
}

func Int(v int64) Value { return Value{Type: 1, IntVal: v} }

type FuncType func(Value) Value
func Func(f FuncType) Value { return Value{Type: 2, UnsafePtr: unsafe.Pointer(&f)} }

func Apply(f Value, arg Value) Value { 
	return (*(*FuncType)(f.UnsafePtr))(arg) 
}
func Apply2(f Value, a Value, b Value) Value { 
	return Apply(Apply(f, a), b) 
}

// ---------------------------------------------------------
// 1. Polymorphism / Type Class Inlining Benchmark
// ---------------------------------------------------------
type Dict struct {
	V0 Value // mappend
	V1 Value // mempty
}

// Current: dynamic dispatch via dictionary
func polyLoopGo_Current(dict *Dict, n int64, acc Value) Value {
	if n == 0 {
		return acc
	}
	newAcc := Apply2(dict.V0, acc, dict.V1)
	return polyLoopGo_Current(dict, n-1, newAcc)
}

// Optimized: Inlined Type Class for Int (Point 1)
func polyLoopGo_Inlined(n int64, acc int64, mempty int64) int64 {
	if n == 0 {
		return acc
	}
	return polyLoopGo_Inlined(n-1, acc+mempty, mempty)
}

// ---------------------------------------------------------
// 2. Uncurrying Benchmark
// ---------------------------------------------------------
// Current: nested closures
func curriedAdd() Value {
	return Func(func(a Value) Value {
		return Func(func(b Value) Value {
			return Int(a.IntVal + b.IntVal)
		})
	})
}

// Optimized: native uncurried
func uncurriedAdd(a int64, b int64) int64 {
	return a + b
}

func loopCurried(n int64) {
	f := curriedAdd()
	acc := Int(0)
	for i := int64(0); i < n; i++ {
		acc = Apply2(f, acc, Int(1))
	}
}

func loopUncurried(n int64) {
	acc := int64(0)
	for i := int64(0); i < n; i++ {
		acc = uncurriedAdd(acc, 1)
	}
}

// ---------------------------------------------------------
// 3. Records Benchmark
// ---------------------------------------------------------
// Current: Dynamic map-like access
type Record struct {
	fields map[string]Value
}
func RecordGet(r *Record, key string) Value {
	return r.fields[key]
}

func loopRecordCurrent(n int64, r *Record) {
	acc := int64(0)
	for i := int64(0); i < n; i++ {
		acc += RecordGet(r, "val").IntVal
	}
}

// Optimized: Native struct
type NativeRecord struct {
	val int64
}
func loopRecordOptimized(n int64, r *NativeRecord) {
	acc := int64(0)
	for i := int64(0); i < n; i++ {
		acc += r.val
	}
}

func main() {
	fmt.Println("=====================================================")
	fmt.Println(" 🚀 GOPURS OPTIMIZATIONS BENCHMARK (Proof of Concept)")
	fmt.Println("=====================================================\n")

	res1 := testing.Benchmark(func(b *testing.B) {
		d := &Dict{
			V0: Func(func(a Value) Value {
				return Func(func(c Value) Value {
					return Int(a.IntVal + c.IntVal)
				})
			}),
			V1: Int(1),
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			polyLoopGo_Current(d, 1000, Int(0))
		}
	})
	res2 := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			polyLoopGo_Inlined(1000, 0, 1)
		}
	})
	fmt.Println("[1] Inlining Type Classes (Polymorphism)")
	fmt.Printf("Current (Boxing + Apply2) : %10d ns/op\n", res1.NsPerOp())
	fmt.Printf("Optimized (Inlined Ints)  : %10d ns/op\n", res2.NsPerOp())
	fmt.Printf("Speedup                   : %.2fx\n\n", float64(res1.NsPerOp())/float64(res2.NsPerOp()))

	res3 := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loopCurried(1000)
		}
	})
	res4 := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loopUncurried(1000)
		}
	})
	fmt.Println("[2] Uncurrying (Closures imbriquées vs Native Func)")
	fmt.Printf("Current (Nested Closures) : %10d ns/op\n", res3.NsPerOp())
	fmt.Printf("Optimized (Uncurried Native): %10d ns/op\n", res4.NsPerOp())
	fmt.Printf("Speedup                   : %.2fx\n\n", float64(res3.NsPerOp())/float64(res4.NsPerOp()))

	res5 := testing.Benchmark(func(b *testing.B) {
		r := &Record{fields: map[string]Value{"val": Int(42)}}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			loopRecordCurrent(1000, r)
		}
	})
	res6 := testing.Benchmark(func(b *testing.B) {
		r := &NativeRecord{val: 42}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			loopRecordOptimized(1000, r)
		}
	})
	fmt.Println("[3] Records (Dynamic map vs Native Structs)")
	fmt.Printf("Current (Dynamic Get)     : %10d ns/op\n", res5.NsPerOp())
	fmt.Printf("Optimized (Native Struct) : %10d ns/op\n", res6.NsPerOp())
	fmt.Printf("Speedup                   : %.2fx\n", float64(res5.NsPerOp())/float64(res6.NsPerOp()))
}
