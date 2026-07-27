package Control_Apply

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Function "gopurs/output/Data.Function"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_identity(gopurs_runtime.UnboxAny(x_0_box)))
})
	})
	return cache_identity
}

var cache_applyProxy gopurs_runtime.Value
var once_applyProxy sync.Once
func Get_applyProxy() gopurs_runtime.Value {
	once_applyProxy.Do(func() {
		cache_applyProxy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorProxy()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})
}))))
	})
	return cache_applyProxy
}

var cache_applyFn gopurs_runtime.Value
var once_applyFn sync.Once
func Get_applyFn() gopurs_runtime.Value {
	once_applyFn.Do(func() {
		cache_applyFn = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
}))))
	})
	return cache_applyFn
}

var cache_applyArray gopurs_runtime.Value
var once_applyArray sync.Once
func Get_applyArray() gopurs_runtime.Value {
	once_applyArray.Do(func() {
		cache_applyArray = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), Get_arrayApply())))
	})
	return cache_applyArray
}

var cache_apply gopurs_runtime.Value
var once_apply sync.Once
func Get_apply() gopurs_runtime.Value {
	once_apply.Do(func() {
		cache_apply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply(dict_0_box)
})
	})
	return cache_apply
}

var cache_apply__func_gopurs_runtime_Value__interface____interface____interface___3235407395 gopurs_runtime.Value
var once_apply__func_gopurs_runtime_Value__interface____interface____interface___3235407395 sync.Once
func Get_apply__func_gopurs_runtime_Value__interface____interface____interface___3235407395() gopurs_runtime.Value {
	once_apply__func_gopurs_runtime_Value__interface____interface____interface___3235407395.Do(func() {
		cache_apply__func_gopurs_runtime_Value__interface____interface____interface___3235407395 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__func_gopurs_runtime_Value__interface____interface____interface___3235407395(dict_0_box)
})
	})
	return cache_apply__func_gopurs_runtime_Value__interface____interface____interface___3235407395
}

var cache_applyFirst gopurs_runtime.Value
var once_applyFirst sync.Once
func Get_applyFirst() gopurs_runtime.Value {
	once_applyFirst.Do(func() {
		cache_applyFirst = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applyFirst(dictApply_0_box, gopurs_runtime.UnboxAny(a_1_box), gopurs_runtime.UnboxAny(b_2_box)))
})
	})
	return cache_applyFirst
}

var cache_applySecond gopurs_runtime.Value
var once_applySecond sync.Once
func Get_applySecond() gopurs_runtime.Value {
	once_applySecond.Do(func() {
		cache_applySecond = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applySecond(dictApply_0_box, gopurs_runtime.UnboxAny(a_1_box), gopurs_runtime.UnboxAny(b_2_box)))
})
	})
	return cache_applySecond
}

var cache_applySecond__func_gopurs_runtime_Value__interface____interface____interface___3235407395 gopurs_runtime.Value
var once_applySecond__func_gopurs_runtime_Value__interface____interface____interface___3235407395 sync.Once
func Get_applySecond__func_gopurs_runtime_Value__interface____interface____interface___3235407395() gopurs_runtime.Value {
	once_applySecond__func_gopurs_runtime_Value__interface____interface____interface___3235407395.Do(func() {
		cache_applySecond__func_gopurs_runtime_Value__interface____interface____interface___3235407395 = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applySecond__func_gopurs_runtime_Value__interface____interface____interface___3235407395(dictApply_0_box, gopurs_runtime.UnboxAny(a_1_box), gopurs_runtime.UnboxAny(b_2_box)))
})
	})
	return cache_applySecond__func_gopurs_runtime_Value__interface____interface____interface___3235407395
}

var cache_lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		cache_lift2 = gopurs_runtime.Func4(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_lift2(dictApply_0_box, func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box)))
})
	})
	return cache_lift2
}

var cache_lift2__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1674249415 gopurs_runtime.Value
var once_lift2__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1674249415 sync.Once
func Get_lift2__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1674249415() gopurs_runtime.Value {
	once_lift2__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1674249415.Do(func() {
		cache_lift2__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1674249415 = gopurs_runtime.Func4(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_lift2__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1674249415(dictApply_0_box, func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box)))
})
	})
	return cache_lift2__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1674249415
}

var cache_lift3 gopurs_runtime.Value
var once_lift3 sync.Once
func Get_lift3() gopurs_runtime.Value {
	once_lift3.Do(func() {
		cache_lift3 = gopurs_runtime.Func5(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_lift3(dictApply_0_box, func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(f_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2)))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box)))
})
	})
	return cache_lift3
}

var cache_lift4 gopurs_runtime.Value
var once_lift4 sync.Once
func Get_lift4() gopurs_runtime.Value {
	once_lift4.Do(func() {
		cache_lift4 = gopurs_runtime.Func6(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_lift4(dictApply_0_box, func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply4(f_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3)))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box), gopurs_runtime.UnboxAny(d_5_box)))
})
	})
	return cache_lift4
}

var cache_lift5 gopurs_runtime.Value
var once_lift5 sync.Once
func Get_lift5() gopurs_runtime.Value {
	once_lift5.Do(func() {
		cache_lift5 = gopurs_runtime.Func7(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value, e_6_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_lift5(dictApply_0_box, func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply5(f_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4)))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box), gopurs_runtime.UnboxAny(d_5_box), gopurs_runtime.UnboxAny(e_6_box)))
})
	})
	return cache_lift5
}

var cache_arrayApply gopurs_runtime.Value
var once_arrayApply sync.Once
func Get_arrayApply() gopurs_runtime.Value {
	once_arrayApply.Do(func() {
		cache_arrayApply = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := ArrayApply(func() []func(interface{}) interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]func(interface{}) interface{}, len(arr))
					for i, v := range arr { unboxed[i] = func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v, gopurs_runtime.Any(inner_arg0)))
} }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_arrayApply
}

func Call_identity(x_0_loop interface{}) interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
return x_0
}

func Call_apply(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "apply")
}

func Call_apply__func_gopurs_runtime_Value__interface____interface____interface___3235407395(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "apply")
}

func Call_applyFirst(dictApply_0_loop gopurs_runtime.Value, a_1_loop interface{}, b_2_loop interface{}) interface{} {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var a_1 interface{} = a_1_loop
_ = a_1
var b_2 interface{} = b_2_loop
_ = b_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Function.Get_const_(), gopurs_runtime.Any(a_1)), gopurs_runtime.Any(b_2)))
}

func Call_applySecond(dictApply_0_loop gopurs_runtime.Value, a_1_loop interface{}, b_2_loop interface{}) interface{} {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var a_1 interface{} = a_1_loop
_ = a_1
var b_2 interface{} = b_2_loop
_ = b_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_identity()
}), gopurs_runtime.Any(a_1)), gopurs_runtime.Any(b_2)))
}

func Call_applySecond__func_gopurs_runtime_Value__interface____interface____interface___3235407395(dictApply_0_loop gopurs_runtime.Value, a_1_loop interface{}, b_2_loop interface{}) interface{} {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var a_1 interface{} = a_1_loop
_ = a_1
var b_2 interface{} = b_2_loop
_ = b_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_identity()
}), gopurs_runtime.Any(a_1)), gopurs_runtime.Any(b_2)))
}

func Call_lift2(dictApply_0_loop gopurs_runtime.Value, f_1_loop func(interface{}, interface{}) interface{}, a_2_loop interface{}, b_3_loop interface{}) interface{} {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 func(interface{}, interface{}) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Any(a_2)), gopurs_runtime.Any(b_3)))
}

func Call_lift2__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1674249415(dictApply_0_loop gopurs_runtime.Value, f_1_loop func(interface{}, interface{}) interface{}, a_2_loop interface{}, b_3_loop interface{}) interface{} {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 func(interface{}, interface{}) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Any(a_2)), gopurs_runtime.Any(b_3)))
}

func Call_lift3(dictApply_0_loop gopurs_runtime.Value, f_1_loop func(interface{}, interface{}, interface{}) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}) interface{} {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 func(interface{}, interface{}, interface{}) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2)))
}), gopurs_runtime.Any(a_2)), gopurs_runtime.Any(b_3)), gopurs_runtime.Any(c_4)))
}

func Call_lift4(dictApply_0_loop gopurs_runtime.Value, f_1_loop func(interface{}, interface{}, interface{}, interface{}) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}, d_5_loop interface{}) interface{} {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 func(interface{}, interface{}, interface{}, interface{}) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3)))
}), gopurs_runtime.Any(a_2)), gopurs_runtime.Any(b_3)), gopurs_runtime.Any(c_4)), gopurs_runtime.Any(d_5)))
}

func Call_lift5(dictApply_0_loop gopurs_runtime.Value, f_1_loop func(interface{}, interface{}, interface{}, interface{}, interface{}) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}, d_5_loop interface{}, e_6_loop interface{}) interface{} {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 func(interface{}, interface{}, interface{}, interface{}, interface{}) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
var e_6 interface{} = e_6_loop
_ = e_6
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4)))
}), gopurs_runtime.Any(a_2)), gopurs_runtime.Any(b_3)), gopurs_runtime.Any(c_4)), gopurs_runtime.Any(d_5)), gopurs_runtime.Any(e_6)))
}
