package main

import (
	"fmt"
	"time"
)

// ==========================================
// 1. Avant Go 1.27 : Value Boxing (Simulation)
// ==========================================
type Value interface{}

// Dictionnaire Functor avant Go 1.27 : Map doit utiliser Value car on n'a pas de generic methods
type DictFunctorOld struct {
	Map func(f func(Value) Value, fa Value) Value
}

func testOld() time.Duration {
	dict := DictFunctorOld{
		Map: func(f func(Value) Value, fa Value) Value {
			// Box/Unbox overhead
			val := fa.(int)
			return f(val).(int)
		},
	}

	start := time.Now()
	res := 0
	for i := 0; i < 10000000; i++ {
		out := dict.Map(func(v Value) Value {
			return v.(int) + 1
		}, res)
		res = out.(int)
	}
	elapsed := time.Since(start)
	if res != 10000000 {
		panic("error old")
	}
	return elapsed
}

// ==========================================
// 2. Avec Go 1.27 : Generic Methods sur Dictionnaires
// ==========================================
// Simulation de Go 1.27 : en vrai on utiliserait une methode générique sur le struct,
// mais pour le simuler en Go 1.24 on met les types A et B au niveau du struct juste pour le test.
// Le vrai code Go 1.27 permettrait : func (d Dict) Map[A, B any](f func(A)B, fa A) B
type DictFunctorNew[A any, B any] struct {
	Map func(f func(A) B, fa A) B
}

func testNew() time.Duration {
	dict := DictFunctorNew[int, int]{
		Map: func(f func(int) int, fa int) int {
			return f(fa)
		},
	}

	start := time.Now()
	res := 0
	for i := 0; i < 10000000; i++ {
		out := dict.Map(func(v int) int {
			return v + 1
		}, res)
		res = out
	}
	elapsed := time.Since(start)
	if res != 10000000 {
		panic("error new")
	}
	return elapsed
}

func main() {
	fmt.Println("=== Benchmark : Go < 1.27 (Boxing) vs Go 1.27 (Generic Methods) ===")
	tOld := testOld()
	fmt.Printf("Avant Go 1.27 (Value Boxing)     : %v\n", tOld)
	
	tNew := testNew()
	fmt.Printf("Avec Go 1.27  (Generic Methods)  : %v\n", tNew)
	
	fmt.Printf("Gain de performance (Ratio)      : x%.2f\n", float64(tOld)/float64(tNew))
}
