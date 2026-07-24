package Data_Number_Format

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var clamp gopurs_runtime.Value
var once_clamp sync.Once
func Get_clamp() gopurs_runtime.Value {
	once_clamp.Do(func() {
		clamp = gopurs_runtime.Func3(func(low_0_box gopurs_runtime.Value, hi_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_clamp(low_0_box, hi_1_box, x_2_box)
})
	})
	return clamp
}

var Precision gopurs_runtime.Value
var once_Precision sync.Once
func Get_Precision() gopurs_runtime.Value {
	once_Precision.Do(func() {
		Precision = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Precision", value0)
})
	})
	return Precision
}

var Fixed gopurs_runtime.Value
var once_Fixed sync.Once
func Get_Fixed() gopurs_runtime.Value {
	once_Fixed.Do(func() {
		Fixed = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Fixed", value0)
})
	})
	return Fixed
}

var Exponential gopurs_runtime.Value
var once_Exponential sync.Once
func Get_Exponential() gopurs_runtime.Value {
	once_Exponential.Do(func() {
		Exponential = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Exponential", value0)
})
	})
	return Exponential
}

var toStringWith gopurs_runtime.Value
var once_toStringWith sync.Once
func Get_toStringWith() gopurs_runtime.Value {
	once_toStringWith.Do(func() {
		toStringWith = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Precision").IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_toPrecisionNative(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Fixed").IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_toFixedNative(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Exponential").IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_toExponentialNative(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return toStringWith
}

var precision gopurs_runtime.Value
var once_precision sync.Once
func Get_precision() gopurs_runtime.Value {
	once_precision.Do(func() {
		precision = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Constructor1("Precision", Call_clamp(gopurs_runtime.Int(1), gopurs_runtime.Int(21), x_0))
}()
})
	})
	return precision
}

var fixed gopurs_runtime.Value
var once_fixed sync.Once
func Get_fixed() gopurs_runtime.Value {
	once_fixed.Do(func() {
		fixed = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Constructor1("Fixed", Call_clamp(gopurs_runtime.Int(0), gopurs_runtime.Int(20), x_0))
}()
})
	})
	return fixed
}

var exponential gopurs_runtime.Value
var once_exponential sync.Once
func Get_exponential() gopurs_runtime.Value {
	once_exponential.Do(func() {
		exponential = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Constructor1("Exponential", Call_clamp(gopurs_runtime.Int(0), gopurs_runtime.Int(20), x_0))
}()
})
	})
	return exponential
}

func Call_clamp(low_0_loop gopurs_runtime.Value, hi_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var low_0 gopurs_runtime.Value = low_0_loop
_ = low_0
var hi_1 gopurs_runtime.Value = hi_1_loop
_ = hi_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), low_0, x_2)
_ = v_3_0
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3_0.StrVal == "LT").IntVal != 0 {
__t2 = x_2
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_3_0.StrVal == "EQ").IntVal != 0 {
__t2 = low_0
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_3_0.StrVal == "GT").IntVal != 0 {
__t2 = low_0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__local_var_4_1 := __t2
_ = __local_var_4_1
v_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), hi_1, __local_var_4_1)
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5_3.StrVal == "LT").IntVal != 0 {
__t4 = hi_1
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_5_3.StrVal == "EQ").IntVal != 0 {
__t4 = hi_1
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_5_3.StrVal == "GT").IntVal != 0 {
__t4 = __local_var_4_1
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
