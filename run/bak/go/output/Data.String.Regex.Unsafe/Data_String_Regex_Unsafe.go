package Data_String_Regex_Unsafe

import (
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_String_Regex "gopurs/output/Data.String.Regex"
	pkg_Partial "gopurs/output/Partial"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_unsafeRegex gopurs_runtime.Value
var once_unsafeRegex sync.Once
func Get_unsafeRegex() gopurs_runtime.Value {
	once_unsafeRegex.Do(func() {
		cache_unsafeRegex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeRegex(s_0_box.StrVal(), f_1_box)
})
	})
	return cache_unsafeRegex
}

var cache_either__1999147371 gopurs_runtime.Value
var once_either__1999147371 sync.Once
func Get_either__1999147371() gopurs_runtime.Value {
	once_either__1999147371.Do(func() {
		cache_either__1999147371 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__1999147371(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__1999147371
}

var cache_either__2158544585 gopurs_runtime.Value
var once_either__2158544585 sync.Once
func Get_either__2158544585() gopurs_runtime.Value {
	once_either__2158544585.Do(func() {
		cache_either__2158544585 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__2158544585(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__2158544585
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_unsafeCrashWith__3091512314 gopurs_runtime.Value
var once_unsafeCrashWith__3091512314 sync.Once
func Get_unsafeCrashWith__3091512314() gopurs_runtime.Value {
	once_unsafeCrashWith__3091512314.Do(func() {
		cache_unsafeCrashWith__3091512314 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__3091512314(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__3091512314
}

var cache_unsafePartial__1978895005 gopurs_runtime.Value
var once_unsafePartial__1978895005 sync.Once
func Get_unsafePartial__1978895005() gopurs_runtime.Value {
	once_unsafePartial__1978895005.Do(func() {
		cache_unsafePartial__1978895005 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1978895005
}

var cache_crashWith__3377663964 gopurs_runtime.Value
var once_crashWith__3377663964 sync.Once
func Get_crashWith__3377663964() gopurs_runtime.Value {
	once_crashWith__3377663964.Do(func() {
		cache_crashWith__3377663964 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_crashWith__3377663964(_dollar__unused_0_box)
})
	})
	return cache_crashWith__3377663964
}

func Call_unsafeRegex(s_0_loop string, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
__local_var_2_0 := gopurs_runtime.Apply4(pkg_Data_String_Regex.Get_regexImpl(), pkg_Data_Either.Get_Left(), pkg_Data_Either.Get_Right(), gopurs_runtime.Str(s_0), gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_Regex.Get_renderFlags(), f_1).StrVal()))
_ = __local_var_2_0
var __t2 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3711209382) {
__local_var_3_1 := (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0
_ = __local_var_3_1
__t2 = gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(__local_var_3_1.StrVal()))
}))
goto end_branch_2
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 2465973597) {
__t2 = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}

func Call_either__1999147371(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_either__2158544585(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unsafeCrashWith__3091512314(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}))
}

func Call_crashWith__3377663964(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Partial.Get__crashWith()
}


