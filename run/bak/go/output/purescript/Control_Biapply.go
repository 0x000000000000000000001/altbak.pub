package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Biapply_Biapply_dollarDict gopurs_runtime.Value
var once_Control_Biapply_Biapply_dollarDict sync.Once
func Get_Control_Biapply_Biapply_dollarDict() gopurs_runtime.Value {
	once_Control_Biapply_Biapply_dollarDict.Do(func() {
		cache_Control_Biapply_Biapply_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_Biapply_dollarDict(x_0_box)
})
	})
	return cache_Control_Biapply_Biapply_dollarDict
}

var cache_Control_Biapply_biapplyTuple gopurs_runtime.Value
var once_Control_Biapply_biapplyTuple sync.Once
func Get_Control_Biapply_biapplyTuple() gopurs_runtime.Value {
	once_Control_Biapply_biapplyTuple.Do(func() {
		cache_Control_Biapply_biapplyTuple = gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bifunctor_bifunctorTuple()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Control_Biapply_biapplyTuple
}

var cache_Control_Biapply_biapply gopurs_runtime.Value
var once_Control_Biapply_biapply sync.Once
func Get_Control_Biapply_biapply() gopurs_runtime.Value {
	once_Control_Biapply_biapply.Do(func() {
		cache_Control_Biapply_biapply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_biapply(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply](dict_0_box))
})
	})
	return cache_Control_Biapply_biapply
}

var cache_Control_Biapply_biapplyFirst gopurs_runtime.Value
var once_Control_Biapply_biapplyFirst sync.Once
func Get_Control_Biapply_biapplyFirst() gopurs_runtime.Value {
	once_Control_Biapply_biapplyFirst.Do(func() {
		cache_Control_Biapply_biapplyFirst = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_biapplyFirst(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply](dictBiapply_0_box))
})
	})
	return cache_Control_Biapply_biapplyFirst
}

var cache_Control_Biapply_biapplySecond gopurs_runtime.Value
var once_Control_Biapply_biapplySecond sync.Once
func Get_Control_Biapply_biapplySecond() gopurs_runtime.Value {
	once_Control_Biapply_biapplySecond.Do(func() {
		cache_Control_Biapply_biapplySecond = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_biapplySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply](dictBiapply_0_box))
})
	})
	return cache_Control_Biapply_biapplySecond
}

var cache_Control_Biapply_bilift2 gopurs_runtime.Value
var once_Control_Biapply_bilift2 sync.Once
func Get_Control_Biapply_bilift2() gopurs_runtime.Value {
	once_Control_Biapply_bilift2.Do(func() {
		cache_Control_Biapply_bilift2 = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_bilift2(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply](dictBiapply_0_box))
})
	})
	return cache_Control_Biapply_bilift2
}

var cache_Control_Biapply_bilift3 gopurs_runtime.Value
var once_Control_Biapply_bilift3 sync.Once
func Get_Control_Biapply_bilift3() gopurs_runtime.Value {
	once_Control_Biapply_bilift3.Do(func() {
		cache_Control_Biapply_bilift3 = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_bilift3(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply](dictBiapply_0_box))
})
	})
	return cache_Control_Biapply_bilift3
}

var cache_Control_Biapply_biapply__3394381979 gopurs_runtime.Value
var once_Control_Biapply_biapply__3394381979 sync.Once
func Get_Control_Biapply_biapply__3394381979() gopurs_runtime.Value {
	once_Control_Biapply_biapply__3394381979.Do(func() {
		cache_Control_Biapply_biapply__3394381979 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_biapply__3394381979(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply](dict_0_box))
})
	})
	return cache_Control_Biapply_biapply__3394381979
}

var cache_Control_Biapply_biapply__2409699611 gopurs_runtime.Value
var once_Control_Biapply_biapply__2409699611 sync.Once
func Get_Control_Biapply_biapply__2409699611() gopurs_runtime.Value {
	once_Control_Biapply_biapply__2409699611.Do(func() {
		cache_Control_Biapply_biapply__2409699611 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_biapply__2409699611(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply](dict_0_box))
})
	})
	return cache_Control_Biapply_biapply__2409699611
}

var cache_Control_Biapply_biapplyTuple__355763440 gopurs_runtime.Value
var once_Control_Biapply_biapplyTuple__355763440 sync.Once
func Get_Control_Biapply_biapplyTuple__355763440() gopurs_runtime.Value {
	once_Control_Biapply_biapplyTuple__355763440.Do(func() {
		cache_Control_Biapply_biapplyTuple__355763440 = gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bifunctor_bifunctorTuple()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Control_Biapply_biapplyTuple__355763440
}

type Constructor_Control_Biapply_Biapply struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3774602829] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Biapply_Biapply)(ptr)
		_ = c
		switch key {
		case "Bifunctor0": return gopurs_runtime.Box(c.V0)
		case "biapply": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Biapply_Biapply: " + key)
		}
	}
}


func Call_Control_Biapply_Biapply_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Biapply_biapply(dict_0_loop *Constructor_Control_Biapply_Biapply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Biapply_Biapply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Biapply_biapplyFirst(dictBiapply_0_loop *Constructor_Control_Biapply_Biapply) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Control_Biapply_Biapply = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): Bifunctor0_1_0 -> *Constructor_Data_Bifunctor_Bifunctor
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](gopurs_runtime.Apply(gopurs_runtime.Box(dictBiapply_0.V0), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Control_Biapply_biapplySecond(dictBiapply_0_loop *Constructor_Control_Biapply_Biapply) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Control_Biapply_Biapply = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): Bifunctor0_1_0 -> *Constructor_Data_Bifunctor_Bifunctor
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](gopurs_runtime.Apply(gopurs_runtime.Box(dictBiapply_0.V0), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), Get_Data_Function_go__const(), Get_Data_Function_go__const(), a_2), b_3)
})
})
}

func Call_Control_Biapply_bilift2(dictBiapply_0_loop *Constructor_Control_Biapply_Biapply) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Control_Biapply_Biapply = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): Bifunctor0_1_0 -> *Constructor_Data_Bifunctor_Bifunctor
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](gopurs_runtime.Apply(gopurs_runtime.Box(dictBiapply_0.V0), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), f_2, g_3, a_4), b_5)
})
})
})
})
}

func Call_Control_Biapply_bilift3(dictBiapply_0_loop *Constructor_Control_Biapply_Biapply) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Control_Biapply_Biapply = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): Bifunctor0_1_0 -> *Constructor_Data_Bifunctor_Bifunctor
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](gopurs_runtime.Apply(gopurs_runtime.Box(dictBiapply_0.V0), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), f_2, g_3, a_4), b_5), c_6)
})
})
})
})
})
}

func Call_Control_Biapply_biapply__3394381979(dict_0_loop *Constructor_Control_Biapply_Biapply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Biapply_Biapply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Biapply_biapply__2409699611(dict_0_loop *Constructor_Control_Biapply_Biapply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Biapply_Biapply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}


