package scratch

import (
	"iter"
	"testing"
	"unsafe"
)

type Value struct {
	Type      uint8
	IntVal    int64
	UnsafePtr unsafe.Pointer
}

func BoxInt(i int64) Value {
	return Value{Type: 1, IntVal: i}
}

type Cons struct {
	Rc uint32
	V0 Value
	V1 *Cons
}

// 1. Current Generated Code for filtering evens (simulated)
func FilterEvensBoxed(lst *Cons) *Cons {
	var go_loop func(Value, *Cons) Value
	go_loop = func(v_2_loop Value, v1_3_loop *Cons) Value {
	loop:
		for {
			v_2 := v_2_loop
			v1_3 := v1_3_loop
			var t2 *Cons

			if v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr == nil {
				t2 = v1_3
				goto end_branch
			}

			if v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr != nil {
				c := (*Cons)(v_2.UnsafePtr)
				if c.V0.IntVal%2 == 0 {
					v_2_loop = Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(c.V1)}
					v1_3_loop = &Cons{1, c.V0, v1_3}
					continue loop
				}
				v_2_loop = Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(c.V1)}
				v1_3_loop = v1_3
				continue loop
			}
			t2 = nil
		end_branch:
			return Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(t2)}
		}
	}
	res := go_loop(Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(lst)}, nil)
	return (*Cons)(res.UnsafePtr)
}

// 2. Unboxed (Monomorphized) Loop variables
func FilterEvensUnboxed(lst *Cons) *Cons {
	var go_loop func(*Cons, *Cons) *Cons
	go_loop = func(v_2_loop *Cons, v1_3_loop *Cons) *Cons {
	loop:
		for {
			v_2 := v_2_loop
			v1_3 := v1_3_loop

			if v_2 == nil {
				return v1_3
			}

			if v_2.V0.IntVal%2 == 0 {
				v_2_loop = v_2.V1
				v1_3_loop = &Cons{1, v_2.V0, v1_3}
				continue loop
			}

			v_2_loop = v_2.V1
			v1_3_loop = v1_3
			continue loop
		}
	}
	return go_loop(lst, nil)
}

// 3. Fully Monomorphized ADT (IntCons)
type IntCons struct {
	Rc uint32
	V0 int64
	V1 *IntCons
}

func FilterEvensIntCons(lst *IntCons) *IntCons {
	var go_loop func(*IntCons, *IntCons) *IntCons
	go_loop = func(v_2_loop *IntCons, v1_3_loop *IntCons) *IntCons {
	loop:
		for {
			v_2 := v_2_loop
			v1_3 := v1_3_loop

			if v_2 == nil {
				return v1_3
			}

			if v_2.V0%2 == 0 {
				v_2_loop = v_2.V1
				v1_3_loop = &IntCons{1, v_2.V0, v1_3}
				continue loop
			}

			v_2_loop = v_2.V1
			v1_3_loop = v1_3
			continue loop
		}
	}
	return go_loop(lst, nil)
}

// 4. Using Go 1.23+ Iterators
func (l *IntCons) All() iter.Seq[int64] {
	return func(yield func(int64) bool) {
		for c := l; c != nil; c = c.V1 {
			if !yield(c.V0) {
				break
			}
		}
	}
}

func FilterEvensIter(lst *IntCons) *IntCons {
	var res *IntCons
	for v := range lst.All() {
		if v%2 == 0 {
			res = &IntCons{1, v, res}
		}
	}
	return res
}

// Benchmarks
func BenchmarkFilterEvensBoxed(b *testing.B) {
	var lst *Cons
	for i := int64(0); i < 1000; i++ {
		lst = &Cons{1, BoxInt(i), lst}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterEvensBoxed(lst)
	}
}

func BenchmarkFilterEvensUnboxed(b *testing.B) {
	var lst *Cons
	for i := int64(0); i < 1000; i++ {
		lst = &Cons{1, BoxInt(i), lst}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterEvensUnboxed(lst)
	}
}

func BenchmarkFilterEvensIntCons(b *testing.B) {
	var lst *IntCons
	for i := int64(0); i < 1000; i++ {
		lst = &IntCons{1, i, lst}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterEvensIntCons(lst)
	}
}

func BenchmarkFilterEvensIter(b *testing.B) {
	var lst *IntCons
	for i := int64(0); i < 1000; i++ {
		lst = &IntCons{1, i, lst}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterEvensIter(lst)
	}
}
