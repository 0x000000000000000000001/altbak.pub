package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_Tiny_Circle gopurs_runtime.Value
var once_Test_Tiny_Circle sync.Once
func Get_Test_Tiny_Circle() gopurs_runtime.Value {
	once_Test_Tiny_Circle.Do(func() {
		cache_Test_Tiny_Circle = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 248718980, UnsafePtr: unsafe.Pointer(&Constructor_Test_Tiny_Circle{1, value0.IntVal})}
})
	})
	return cache_Test_Tiny_Circle
}

var cache_Test_Tiny_Rect gopurs_runtime.Value
var once_Test_Tiny_Rect sync.Once
func Get_Test_Tiny_Rect() gopurs_runtime.Value {
	once_Test_Tiny_Rect.Do(func() {
		cache_Test_Tiny_Rect = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 261969494, UnsafePtr: unsafe.Pointer(&Constructor_Test_Tiny_Rect{1, value0.IntVal, value1.IntVal})}
})
})
	})
	return cache_Test_Tiny_Rect
}

var cache_Test_Tiny_area gopurs_runtime.Value
var once_Test_Tiny_area sync.Once
func Get_Test_Tiny_area() gopurs_runtime.Value {
	once_Test_Tiny_area.Do(func() {
		cache_Test_Tiny_area = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Tiny_area(v_0_box))
})
	})
	return cache_Test_Tiny_area
}

type Constructor_Test_Tiny_Circle struct {
	Rc uint32
	V0 int64
}


type Constructor_Test_Tiny_Rect struct {
	Rc uint32
	V0 int64
	V1 int64
}


func Call_Test_Tiny_area(v_0_loop gopurs_runtime.Value) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 248718980) {
__t0 = gopurs_runtime.Int(((*Constructor_Test_Tiny_Circle)(v_0.UnsafePtr).V0) * ((*Constructor_Test_Tiny_Circle)(v_0.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 261969494) {
__t0 = gopurs_runtime.Int(((*Constructor_Test_Tiny_Rect)(v_0.UnsafePtr).V0) * ((*Constructor_Test_Tiny_Rect)(v_0.UnsafePtr).V1))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}


