package Data_Number_Format

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Precision gopurs_runtime.Value
var once_Precision sync.Once
func Get_Precision() gopurs_runtime.Value {
	once_Precision.Do(func() {
		cache_Precision = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1786680275, UnsafePtr: unsafe.Pointer(&Constructor_Precision{1, value0.IntVal})}
})
	})
	return cache_Precision
}

var cache_Fixed gopurs_runtime.Value
var once_Fixed sync.Once
func Get_Fixed() gopurs_runtime.Value {
	once_Fixed.Do(func() {
		cache_Fixed = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1390206259, UnsafePtr: unsafe.Pointer(&Constructor_Fixed{1, value0.IntVal})}
})
	})
	return cache_Fixed
}

var cache_Exponential gopurs_runtime.Value
var once_Exponential sync.Once
func Get_Exponential() gopurs_runtime.Value {
	once_Exponential.Do(func() {
		cache_Exponential = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1734244434, UnsafePtr: unsafe.Pointer(&Constructor_Exponential{1, value0.IntVal})}
})
	})
	return cache_Exponential
}

var cache_toStringWith gopurs_runtime.Value
var once_toStringWith sync.Once
func Get_toStringWith() gopurs_runtime.Value {
	once_toStringWith.Do(func() {
		cache_toStringWith = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toStringWith(v_0_box)
})
	})
	return cache_toStringWith
}

var cache_precision gopurs_runtime.Value
var once_precision sync.Once
func Get_precision() gopurs_runtime.Value {
	once_precision.Do(func() {
		cache_precision = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply2(Get_clamp__1512183668(), gopurs_runtime.Int(1), gopurs_runtime.Int(21))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1786680275, UnsafePtr: unsafe.Pointer(&Constructor_Precision{1, gopurs_runtime.Apply(__local_var_0_0, x_1).IntVal})}
})
}()
	})
	return cache_precision
}

var cache_fixed gopurs_runtime.Value
var once_fixed sync.Once
func Get_fixed() gopurs_runtime.Value {
	once_fixed.Do(func() {
		cache_fixed = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply2(Get_clamp__1512183668(), gopurs_runtime.Int(0), gopurs_runtime.Int(20))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1390206259, UnsafePtr: unsafe.Pointer(&Constructor_Fixed{1, gopurs_runtime.Apply(__local_var_0_0, x_1).IntVal})}
})
}()
	})
	return cache_fixed
}

var cache_exponential gopurs_runtime.Value
var once_exponential sync.Once
func Get_exponential() gopurs_runtime.Value {
	once_exponential.Do(func() {
		cache_exponential = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply2(Get_clamp__1512183668(), gopurs_runtime.Int(0), gopurs_runtime.Int(20))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1734244434, UnsafePtr: unsafe.Pointer(&Constructor_Exponential{1, gopurs_runtime.Apply(__local_var_0_0, x_1).IntVal})}
})
}()
	})
	return cache_exponential
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_clamp__1512183668 gopurs_runtime.Value
var once_clamp__1512183668 sync.Once
func Get_clamp__1512183668() gopurs_runtime.Value {
	once_clamp__1512183668.Do(func() {
		cache_clamp__1512183668 = gopurs_runtime.Func3(func(low_0_box gopurs_runtime.Value, hi_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_clamp__1512183668(low_0_box, hi_1_box, x_2_box)
})
	})
	return cache_clamp__1512183668
}

var cache_clamp__709576177 gopurs_runtime.Value
var once_clamp__709576177 sync.Once
func Get_clamp__709576177() gopurs_runtime.Value {
	once_clamp__709576177.Do(func() {
		cache_clamp__709576177 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_clamp__709576177(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), low_1_box, hi_2_box, x_3_box)
})
	})
	return cache_clamp__709576177
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_max__2767602680 gopurs_runtime.Value
var once_max__2767602680 sync.Once
func Get_max__2767602680() gopurs_runtime.Value {
	once_max__2767602680.Do(func() {
		cache_max__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max__2767602680(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_max__2767602680
}

var cache_min__2767602680 gopurs_runtime.Value
var once_min__2767602680 sync.Once
func Get_min__2767602680() gopurs_runtime.Value {
	once_min__2767602680.Do(func() {
		cache_min__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_min__2767602680(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_min__2767602680
}

type Constructor_Precision struct {
	Rc uint32
	V0 int64
}


type Constructor_Fixed struct {
	Rc uint32
	V0 int64
}


type Constructor_Exponential struct {
	Rc uint32
	V0 int64
}


func Call_toStringWith(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1786680275) {
__t0 = gopurs_runtime.Apply(Get_toPrecisionNative(), gopurs_runtime.Int((*Constructor_Precision)(v_0.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1390206259) {
__t0 = gopurs_runtime.Apply(Get_toFixedNative(), gopurs_runtime.Int((*Constructor_Fixed)(v_0.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1734244434) {
__t0 = gopurs_runtime.Apply(Get_toExponentialNative(), gopurs_runtime.Int((*Constructor_Exponential)(v_0.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_clamp__1512183668(low_0_loop gopurs_runtime.Value, hi_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var low_0 gopurs_runtime.Value = low_0_loop
_ = low_0
var hi_1 gopurs_runtime.Value = hi_1_loop
_ = hi_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
v_3_1 := gopurs_runtime.Apply5(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, low_0, x_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_3_1.IntVal) == 1527465420) {
__t2 = x_2
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 902936544) {
__t2 = low_0
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 380165415) {
__t2 = low_0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__local_var_3_0 := __t2
_ = __local_var_3_0
v_4_3 := gopurs_runtime.Apply5(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, hi_1, __local_var_3_0)
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (uint32(v_4_3.IntVal) == 1527465420) {
__t4 = hi_1
goto end_branch_4
} else {

}
}
{
if (uint32(v_4_3.IntVal) == 902936544) {
__t4 = hi_1
goto end_branch_4
} else {

}
}
{
if (uint32(v_4_3.IntVal) == 380165415) {
__t4 = __local_var_3_0
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}

func Call_clamp__709576177(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
v_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, low_1, x_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_4_1.IntVal) == 1527465420) {
__t2 = x_3
goto end_branch_2
} else {

}
}
{
if (uint32(v_4_1.IntVal) == 902936544) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_4_1.IntVal) == 380165415) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__local_var_4_0 := __t2
_ = __local_var_4_0
v_5_3 := gopurs_runtime.Apply2(dictOrd_0.V1, hi_2, __local_var_4_0)
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (uint32(v_5_3.IntVal) == 1527465420) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (uint32(v_5_3.IntVal) == 902936544) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (uint32(v_5_3.IntVal) == 380165415) {
__t4 = __local_var_4_0
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_max__2767602680(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_min__2767602680(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Get_toExponentialNative() gopurs_runtime.Value {
	return _Gopurs_ToExponentialNative
}

func Get_toFixedNative() gopurs_runtime.Value {
	return _Gopurs_ToFixedNative
}

func Get_toPrecisionNative() gopurs_runtime.Value {
	return _Gopurs_ToPrecisionNative
}

func Get_toString() gopurs_runtime.Value {
	return _Gopurs_ToString
}
