package Test_Tiny

import (
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Circle gopurs_runtime.Value
var once_Circle sync.Once
func Get_Circle() gopurs_runtime.Value {
	once_Circle.Do(func() {
		cache_Circle = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 248718980, UnsafePtr: unsafe.Pointer(&Constructor_Circle{1, value0.IntVal})}
})
	})
	return cache_Circle
}

var cache_Rect gopurs_runtime.Value
var once_Rect sync.Once
func Get_Rect() gopurs_runtime.Value {
	once_Rect.Do(func() {
		cache_Rect = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 261969494, UnsafePtr: unsafe.Pointer(&Constructor_Rect{1, value0.IntVal, value1.IntVal})}
})
})
	})
	return cache_Rect
}

var cache_area gopurs_runtime.Value
var once_area sync.Once
func Get_area() gopurs_runtime.Value {
	once_area.Do(func() {
		cache_area = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_area(v_0_box))
})
	})
	return cache_area
}

var cache_mul__560788792 gopurs_runtime.Value
var once_mul__560788792 sync.Once
func Get_mul__560788792() gopurs_runtime.Value {
	once_mul__560788792.Do(func() {
		cache_mul__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mul__560788792
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

type Constructor_Circle struct {
	Rc uint32
	V0 int64
}


type Constructor_Rect struct {
	Rc uint32
	V0 int64
	V1 int64
}


func Call_area(v_0_loop gopurs_runtime.Value) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 248718980) {
__t0 = gopurs_runtime.Int(Call_mul__560788792(gopurs_runtime.Int((*Constructor_Circle)(v_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Circle)(v_0.UnsafePtr).V0)).IntVal)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 261969494) {
__t0 = gopurs_runtime.Int(Call_mul__560788792(gopurs_runtime.Int((*Constructor_Rect)(v_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Rect)(v_0.UnsafePtr).V1)).IntVal)
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

func Call_mul__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) * (__eta1_1.IntVal))
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


