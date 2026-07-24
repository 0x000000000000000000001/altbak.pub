package Control_Apply

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var applyProxy gopurs_runtime.Value
var once_applyProxy sync.Once
func Get_applyProxy() gopurs_runtime.Value {
	once_applyProxy.Do(func() {
		applyProxy = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorProxy()
}))
	})
	return applyProxy
}

var applyFn gopurs_runtime.Value
var once_applyFn sync.Once
func Get_applyFn() gopurs_runtime.Value {
	once_applyFn.Do(func() {
		applyFn = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}))
	})
	return applyFn
}

var applyArray gopurs_runtime.Value
var once_applyArray sync.Once
func Get_applyArray() gopurs_runtime.Value {
	once_applyArray.Do(func() {
		applyArray = gopurs_runtime.RecordDict2("apply", "Functor0", Get_arrayApply(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}))
	})
	return applyArray
}

var apply gopurs_runtime.Value
var once_apply sync.Once
func Get_apply() gopurs_runtime.Value {
	once_apply.Do(func() {
		apply = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "apply")
}()
})
	})
	return apply
}

var applyFirst gopurs_runtime.Value
var once_applyFirst sync.Once
func Get_applyFirst() gopurs_runtime.Value {
	once_applyFirst.Do(func() {
		applyFirst = gopurs_runtime.Func3(Call_applyFirst)
	})
	return applyFirst
}

var applySecond gopurs_runtime.Value
var once_applySecond sync.Once
func Get_applySecond() gopurs_runtime.Value {
	once_applySecond.Do(func() {
		applySecond = gopurs_runtime.Func3(Call_applySecond)
	})
	return applySecond
}

var lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		lift2 = gopurs_runtime.Func4(Call_lift2)
	})
	return lift2
}

var lift3 gopurs_runtime.Value
var once_lift3 sync.Once
func Get_lift3() gopurs_runtime.Value {
	once_lift3.Do(func() {
		lift3 = gopurs_runtime.Func5(Call_lift3)
	})
	return lift3
}

var lift4 gopurs_runtime.Value
var once_lift4 sync.Once
func Get_lift4() gopurs_runtime.Value {
	once_lift4.Do(func() {
		lift4 = gopurs_runtime.Func6(Call_lift4)
	})
	return lift4
}

var lift5 gopurs_runtime.Value
var once_lift5 sync.Once
func Get_lift5() gopurs_runtime.Value {
	once_lift5.Do(func() {
		lift5 = gopurs_runtime.Func7(Call_lift5)
	})
	return lift5
}

func Call_applyFirst(dictApply_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0_loop, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Function.Get_const_(), a_1_loop), b_2_loop)
}

func Call_applySecond(dictApply_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0_loop, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity")
}), a_1_loop), b_2_loop)
}

func Call_lift2(dictApply_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0_loop, "Functor0"), gopurs_runtime.Value{}), "map"), f_1_loop, a_2_loop), b_3_loop)
}

func Call_lift3(dictApply_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var c_4 gopurs_runtime.Value = c_4_loop
_ = c_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0_loop, "Functor0"), gopurs_runtime.Value{}), "map"), f_1_loop, a_2_loop), b_3_loop), c_4_loop)
}

func Call_lift4(dictApply_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value, d_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var c_4 gopurs_runtime.Value = c_4_loop
_ = c_4
var d_5 gopurs_runtime.Value = d_5_loop
_ = d_5
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0_loop, "Functor0"), gopurs_runtime.Value{}), "map"), f_1_loop, a_2_loop), b_3_loop), c_4_loop), d_5_loop)
}

func Call_lift5(dictApply_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value, d_5_loop gopurs_runtime.Value, e_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var c_4 gopurs_runtime.Value = c_4_loop
_ = c_4
var d_5 gopurs_runtime.Value = d_5_loop
_ = d_5
var e_6 gopurs_runtime.Value = e_6_loop
_ = e_6
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0_loop, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0_loop, "Functor0"), gopurs_runtime.Value{}), "map"), f_1_loop, a_2_loop), b_3_loop), c_4_loop), d_5_loop), e_6_loop)
}

func Get_arrayApply() gopurs_runtime.Value {
	return _Gopurs_ArrayApply
}
