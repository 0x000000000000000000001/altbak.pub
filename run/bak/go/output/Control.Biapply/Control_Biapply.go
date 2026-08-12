package Control_Biapply

import (
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_biapplyTuple gopurs_runtime.Value
var once_biapplyTuple sync.Once
func Get_biapplyTuple() gopurs_runtime.Value {
	once_biapplyTuple.Do(func() {
		cache_biapplyTuple = gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}))}
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
return Call_biapply(gopurs_runtime.CoerceToStruct[Constructor_Biapply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_biapply
}

var cache_biapply__gopurs_runtime_Value_3394381979 gopurs_runtime.Value
var once_biapply__gopurs_runtime_Value_3394381979 sync.Once
func Get_biapply__gopurs_runtime_Value_3394381979() gopurs_runtime.Value {
	once_biapply__gopurs_runtime_Value_3394381979.Do(func() {
		cache_biapply__gopurs_runtime_Value_3394381979 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapply__gopurs_runtime_Value_3394381979(gopurs_runtime.CoerceToStruct[Constructor_Biapply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_biapply__gopurs_runtime_Value_3394381979
}

var cache_biapplyFirst gopurs_runtime.Value
var once_biapplyFirst sync.Once
func Get_biapplyFirst() gopurs_runtime.Value {
	once_biapplyFirst.Do(func() {
		cache_biapplyFirst = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplyFirst(gopurs_runtime.CoerceToStruct[Constructor_Biapply[gopurs_runtime.Value]](dictBiapply_0_box))
})
	})
	return cache_biapplyFirst
}

var cache_biapplySecond gopurs_runtime.Value
var once_biapplySecond sync.Once
func Get_biapplySecond() gopurs_runtime.Value {
	once_biapplySecond.Do(func() {
		cache_biapplySecond = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplySecond(gopurs_runtime.CoerceToStruct[Constructor_Biapply[gopurs_runtime.Value]](dictBiapply_0_box))
})
	})
	return cache_biapplySecond
}

var cache_bilift2 gopurs_runtime.Value
var once_bilift2 sync.Once
func Get_bilift2() gopurs_runtime.Value {
	once_bilift2.Do(func() {
		cache_bilift2 = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bilift2(gopurs_runtime.CoerceToStruct[Constructor_Biapply[gopurs_runtime.Value]](dictBiapply_0_box))
})
	})
	return cache_bilift2
}

var cache_bilift3 gopurs_runtime.Value
var once_bilift3 sync.Once
func Get_bilift3() gopurs_runtime.Value {
	once_bilift3.Do(func() {
		cache_bilift3 = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bilift3(gopurs_runtime.CoerceToStruct[Constructor_Biapply[gopurs_runtime.Value]](dictBiapply_0_box))
})
	})
	return cache_bilift3
}

type Constructor_Biapply[T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3774602829] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Biapply[gopurs_runtime.Value])(ptr)
		switch key {
		case "Bifunctor0": return c.V0
		case "biapply": return c.V1
		default: panic("Key not found in dictionary Constructor_Biapply: " + key)
		}
	}
}


func Call_biapply(dict_0_loop *Constructor_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Biapply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_biapply__gopurs_runtime_Value_3394381979(dict_0_loop *Constructor_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Biapply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_biapplyFirst(dictBiapply_0_loop *Constructor_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Biapply[gopurs_runtime.Value] = dictBiapply_0_loop
_ = dictBiapply_0
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictBiapply_0.V0, gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictBiapply_0.V1, gopurs_runtime.Apply3(Bifunctor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), a_2), b_3)
})
})
}

func Call_biapplySecond(dictBiapply_0_loop *Constructor_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Biapply[gopurs_runtime.Value] = dictBiapply_0_loop
_ = dictBiapply_0
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictBiapply_0.V0, gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictBiapply_0.V1, gopurs_runtime.Apply3(Bifunctor0_1_0.V0, pkg_Data_Function.Get_go__const(), pkg_Data_Function.Get_go__const(), a_2), b_3)
})
})
}

func Call_bilift2(dictBiapply_0_loop *Constructor_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Biapply[gopurs_runtime.Value] = dictBiapply_0_loop
_ = dictBiapply_0
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictBiapply_0.V0, gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictBiapply_0.V1, gopurs_runtime.Apply3(Bifunctor0_1_0.V0, f_2, g_3, a_4), b_5)
})
})
})
})
}

func Call_bilift3(dictBiapply_0_loop *Constructor_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Biapply[gopurs_runtime.Value] = dictBiapply_0_loop
_ = dictBiapply_0
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictBiapply_0.V0, gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictBiapply_0.V1, gopurs_runtime.Apply2(dictBiapply_0.V1, gopurs_runtime.Apply3(Bifunctor0_1_0.V0, f_2, g_3, a_4), b_5), c_6)
})
})
})
})
})
}


