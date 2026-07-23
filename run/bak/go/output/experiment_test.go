package main

import (
	"testing"
)

// 1. Polymorphisme et Records (Dictionaries)
type MapRecord map[string]int

type StructRecord struct {
	a int
	b int
}

func BenchmarkDynamicMapRecord(b *testing.B) {
	r := MapRecord{"a": 1, "b": 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r["a"] = r["a"] + r["b"]
	}
}

func BenchmarkNativeStructRecord(b *testing.B) {
	r := StructRecord{a: 1, b: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.a = r.a + r.b
	}
}

// 2. Currying et Dispatch dynamique (Apply)
type Value interface{}
type Func func(Value) Value

func Apply(f Func, x Value) Value {
	return f(x)
}

func BenchmarkBoxedCurriedApply(b *testing.B) {
	add := Func(func(x Value) Value {
		return Func(func(y Value) Value {
			return x.(int) + y.(int)
		})
	})
	// Initialiser la fonction partiellement appliquée
	add10 := Apply(add, 10).(Func)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Apply(add10, i)
	}
}

func BenchmarkNativeFunctionCall(b *testing.B) {
	add := func(x, y int) int {
		return x + y
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = add(10, i)
	}
}

// 3. Boucles avec Boxing vs Type Primitif
func BenchmarkBoxedLoop(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var acc Value = 0
		var n Value = 1000
		for n.(int) > 0 {
			acc = acc.(int) + n.(int)%3
			n = n.(int) - 1
		}
	}
}

func BenchmarkNativeLoop(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		acc := 0
		n := 1000
		for n > 0 {
			acc = acc + n%3
			n = n - 1
		}
	}
}
