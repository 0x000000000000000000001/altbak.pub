package main

import (
	"math"
	"testing"
	"unsafe"
)

const (
	TypeFunc = 1
	TypeFunc2 = 2
	TypeInt = 6
	TypeString = 7
	TypeAny = 13
)

type Value struct {
	Type   uint8
	IntVal int64
	StrVal string
	PtrVal any
	Func   unsafe.Pointer
}

func Int(v int64) Value { return Value{Type: TypeInt, IntVal: v} }
func Str(v string) Value { return Value{Type: TypeString, StrVal: v} }
func Any(v any) Value { return Value{Type: TypeAny, PtrVal: v} }

func Func(f func(Value) Value) Value {
	return Value{Type: TypeFunc, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}
func Func2(f func(Value, Value) Value) Value {
	return Value{Type: TypeFunc2, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}

func Apply(f Value, arg Value) Value {
	if f.Type == TypeFunc {
		return (*(*func(Value) Value)(unsafe.Pointer(&f.Func)))(arg)
	}
	fn := *(*func(Value, Value) Value)(unsafe.Pointer(&f.Func))
	return Func(func(a Value) Value { return fn(arg, a) })
}

func Apply2(f Value, arg1, arg2 Value) Value {
	if f.Type == TypeFunc2 {
		return (*(*func(Value, Value) Value)(unsafe.Pointer(&f.Func)))(arg1, arg2)
	}
	return Apply(Apply(f, arg1), arg2)
}

func Unbox[T any](v Value) T {
	var t any = *new(T)
	switch t.(type) {
	case int64: return any(v.IntVal).(T)
	case int: return any(int(v.IntVal)).(T)
	case string: return any(v.StrVal).(T)
	case float64: return any(math.Float64frombits(uint64(v.IntVal))).(T)
	case bool: return any(v.IntVal == 1).(T)
	case Value: return any(v).(T)
	default: return v.PtrVal.(T)
	}
}

func Box[T any](val T) Value {
	switch v := any(val).(type) {
	case int64: return Int(v)
	case int: return Int(int64(v))
	case string: return Str(v)
	case float64: return Value{Type: 10, IntVal: int64(math.Float64bits(v))}
	case bool: 
		if v { return Value{Type: 11, IntVal: 1} }
		return Value{Type: 11, IntVal: 0}
	case Value: return v
	default: return Any(v)
	}
}

func Wrap2[A, B, R any](f func(A, B) R) Value {
	return Func2(func(a, b Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b)))
	})
}

func nativeAdd(a, b int) int { return a + b }

var WrappedAdd = Wrap2(nativeAdd)

func TestWrap(t *testing.T) {
	res := Apply2(WrappedAdd, Int(10), Int(20))
	if res.IntVal != 30 {
		t.Fatalf("Expected 30, got %d", res.IntVal)
	}
}

func BenchmarkWrapAdd(b *testing.B) {
	v1, v2 := Int(10), Int(20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Apply2(WrappedAdd, v1, v2)
	}
}
