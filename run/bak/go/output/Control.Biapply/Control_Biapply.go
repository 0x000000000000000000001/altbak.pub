package Control_Biapply

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Function "gopurs/output/Data.Function"
	unsafe "unsafe"
)

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity1(x_0_box)
})
	})
	return cache_identity1
}

var cache_identity2 gopurs_runtime.Value
var once_identity2 sync.Once
func Get_identity2() gopurs_runtime.Value {
	once_identity2.Do(func() {
		cache_identity2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity2(x_0_box)
})
	})
	return cache_identity2
}

var cache_biapplyTuple gopurs_runtime.Value
var once_biapplyTuple sync.Once
func Get_biapplyTuple() gopurs_runtime.Value {
	once_biapplyTuple.Do(func() {
		cache_biapplyTuple = gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_biapplyTuple
}

var cache_biapplyTuple__gopurs_runtime_Value_355763440 gopurs_runtime.Value
var once_biapplyTuple__gopurs_runtime_Value_355763440 sync.Once
func Get_biapplyTuple__gopurs_runtime_Value_355763440() gopurs_runtime.Value {
	once_biapplyTuple__gopurs_runtime_Value_355763440.Do(func() {
		cache_biapplyTuple__gopurs_runtime_Value_355763440 = gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_biapplyTuple__gopurs_runtime_Value_355763440
}

var cache_biapply gopurs_runtime.Value
var once_biapply sync.Once
func Get_biapply() gopurs_runtime.Value {
	once_biapply.Do(func() {
		cache_biapply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapply(dict_0_box)
})
	})
	return cache_biapply
}

var cache_biapply__gopurs_runtime_Value_3394381979 gopurs_runtime.Value
var once_biapply__gopurs_runtime_Value_3394381979 sync.Once
func Get_biapply__gopurs_runtime_Value_3394381979() gopurs_runtime.Value {
	once_biapply__gopurs_runtime_Value_3394381979.Do(func() {
		cache_biapply__gopurs_runtime_Value_3394381979 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapply__gopurs_runtime_Value_3394381979(dict_0_box)
})
	})
	return cache_biapply__gopurs_runtime_Value_3394381979
}

var cache_biapplyFirst gopurs_runtime.Value
var once_biapplyFirst sync.Once
func Get_biapplyFirst() gopurs_runtime.Value {
	once_biapplyFirst.Do(func() {
		cache_biapplyFirst = gopurs_runtime.Func3(func(dictBiapply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplyFirst(dictBiapply_0_box, a_1_box, b_2_box)
})
	})
	return cache_biapplyFirst
}

var cache_biapplySecond gopurs_runtime.Value
var once_biapplySecond sync.Once
func Get_biapplySecond() gopurs_runtime.Value {
	once_biapplySecond.Do(func() {
		cache_biapplySecond = gopurs_runtime.Func3(func(dictBiapply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplySecond(dictBiapply_0_box, a_1_box, b_2_box)
})
	})
	return cache_biapplySecond
}

var cache_bilift2 gopurs_runtime.Value
var once_bilift2 sync.Once
func Get_bilift2() gopurs_runtime.Value {
	once_bilift2.Do(func() {
		cache_bilift2 = gopurs_runtime.Func5(func(dictBiapply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value, b_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bilift2(dictBiapply_0_box, f_1_box, g_2_box, a_3_box, b_4_box)
})
	})
	return cache_bilift2
}

var cache_bilift3 gopurs_runtime.Value
var once_bilift3 sync.Once
func Get_bilift3() gopurs_runtime.Value {
	once_bilift3.Do(func() {
		cache_bilift3 = gopurs_runtime.Func6(func(dictBiapply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value, b_4_box gopurs_runtime.Value, c_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bilift3(dictBiapply_0_box, f_1_box, g_2_box, a_3_box, b_4_box, c_5_box)
})
	})
	return cache_bilift3
}

func Call_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_biapply(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "biapply")
}

func Call_biapply__gopurs_runtime_Value_3394381979(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "biapply")
}

func Call_biapplyFirst(dictBiapply_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_identity1()
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_identity2()
}), a_1), b_2)
}

func Call_biapplySecond(dictBiapply_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), pkg_Data_Function.Get_go__const(), pkg_Data_Function.Get_go__const(), a_1), b_2)
}

func Call_bilift2(dictBiapply_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value, b_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), f_1, g_2, a_3), b_4)
}

func Call_bilift3(dictBiapply_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value, b_4_loop gopurs_runtime.Value, c_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var c_5 gopurs_runtime.Value = c_5_loop
_ = c_5
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), f_1, g_2, a_3), b_4), c_5)
}


