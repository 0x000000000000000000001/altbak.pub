package main

import (
	"testing"
)

// 1. Generic Approach (Similar to gopurs_runtime)
type Value interface{}
type Func func(Value) Value

func Apply(f Value, arg Value) Value {
	return f.(Func)(arg)
}

func mapGeneric(f Value, arr []Value) []Value {
	res := make([]Value, len(arr))
	for i, v := range arr {
		res[i] = Apply(f, v)
	}
	return res
}

// 2. Devirtualized & Unboxed Approach (Monomorphization)
func mapSpecializedAdd1(arr []int64) []int64 {
	res := make([]int64, len(arr))
	for i, v := range arr {
		res[i] = v + 1
	}
	return res
}

func BenchmarkGeneric(b *testing.B) {
	arr := make([]Value, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = int64(i)
	}
	f := Func(func(x Value) Value {
		return x.(int64) + 1
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mapGeneric(f, arr)
	}
}

func BenchmarkSpecialized(b *testing.B) {
	arr := make([]int64, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = int64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mapSpecializedAdd1(arr)
	}
}
