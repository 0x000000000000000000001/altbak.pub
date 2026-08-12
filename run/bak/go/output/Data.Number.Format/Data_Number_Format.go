package Data_Number_Format

import (
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
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(nil)})
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_1 := gopurs_runtime.Apply2(__local_var_0_0, gopurs_runtime.Int(1), x_1)
_ = v_2_1
var __t3 gopurs_runtime.Value
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 1527465420) {
__t3 = x_1
goto end_branch_3
} else {

}
}
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 902936544) {
__t3 = gopurs_runtime.Int(1)
goto end_branch_3
} else {

}
}
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 380165415) {
__t3 = gopurs_runtime.Int(1)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__local_var_3_2 := __t3
_ = __local_var_3_2
v_4_4 := gopurs_runtime.Apply2(__local_var_0_0, gopurs_runtime.Int(21), __local_var_3_2)
_ = v_4_4
var __t5 gopurs_runtime.Value
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 1527465420) {
__t5 = gopurs_runtime.Int(21)
goto end_branch_5
} else {

}
}
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 902936544) {
__t5 = gopurs_runtime.Int(21)
goto end_branch_5
} else {

}
}
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 380165415) {
__t5 = __local_var_3_2
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1786680275, UnsafePtr: unsafe.Pointer(&Constructor_Precision{1, __t5.IntVal})}
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
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(nil)})
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_1 := gopurs_runtime.Apply2(__local_var_0_0, gopurs_runtime.Int(0), x_1)
_ = v_2_1
var __t3 gopurs_runtime.Value
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 1527465420) {
__t3 = x_1
goto end_branch_3
} else {

}
}
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 902936544) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 380165415) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__local_var_3_2 := __t3
_ = __local_var_3_2
v_4_4 := gopurs_runtime.Apply2(__local_var_0_0, gopurs_runtime.Int(20), __local_var_3_2)
_ = v_4_4
var __t5 gopurs_runtime.Value
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 1527465420) {
__t5 = gopurs_runtime.Int(20)
goto end_branch_5
} else {

}
}
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 902936544) {
__t5 = gopurs_runtime.Int(20)
goto end_branch_5
} else {

}
}
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 380165415) {
__t5 = __local_var_3_2
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1390206259, UnsafePtr: unsafe.Pointer(&Constructor_Fixed{1, __t5.IntVal})}
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
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(nil)})
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_1 := gopurs_runtime.Apply2(__local_var_0_0, gopurs_runtime.Int(0), x_1)
_ = v_2_1
var __t3 gopurs_runtime.Value
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 1527465420) {
__t3 = x_1
goto end_branch_3
} else {

}
}
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 902936544) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
if (v_2_1.Type == 9 && v_2_1.IntVal == 380165415) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__local_var_3_2 := __t3
_ = __local_var_3_2
v_4_4 := gopurs_runtime.Apply2(__local_var_0_0, gopurs_runtime.Int(20), __local_var_3_2)
_ = v_4_4
var __t5 gopurs_runtime.Value
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 1527465420) {
__t5 = gopurs_runtime.Int(20)
goto end_branch_5
} else {

}
}
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 902936544) {
__t5 = gopurs_runtime.Int(20)
goto end_branch_5
} else {

}
}
{
if (v_4_4.Type == 9 && v_4_4.IntVal == 380165415) {
__t5 = __local_var_3_2
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1734244434, UnsafePtr: unsafe.Pointer(&Constructor_Exponential{1, __t5.IntVal})}
})
}()
	})
	return cache_exponential
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
