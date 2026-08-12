package Data_Profunctor_Strong

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_strongFn gopurs_runtime.Value
var once_strongFn sync.Once
func Get_strongFn() gopurs_runtime.Value {
	once_strongFn.Do(func() {
		cache_strongFn = gopurs_runtime.RecordDict3("Profunctor0", "first", "second", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
}), gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(a2b_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1})}))}
})
}), gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"))
	})
	return cache_strongFn
}

var cache_second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		cache_second = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_second(gopurs_runtime.CoerceToStruct[Constructor_Strong[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_second
}

var cache_second__gopurs_runtime_Value_1843542330 gopurs_runtime.Value
var once_second__gopurs_runtime_Value_1843542330 sync.Once
func Get_second__gopurs_runtime_Value_1843542330() gopurs_runtime.Value {
	once_second__gopurs_runtime_Value_1843542330.Do(func() {
		cache_second__gopurs_runtime_Value_1843542330 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_second__gopurs_runtime_Value_1843542330(gopurs_runtime.CoerceToStruct[Constructor_Strong[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_second__gopurs_runtime_Value_1843542330
}

var cache_first gopurs_runtime.Value
var once_first sync.Once
func Get_first() gopurs_runtime.Value {
	once_first.Do(func() {
		cache_first = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_first(gopurs_runtime.CoerceToStruct[Constructor_Strong[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_first
}

var cache_first__gopurs_runtime_Value_1843542330 gopurs_runtime.Value
var once_first__gopurs_runtime_Value_1843542330 sync.Once
func Get_first__gopurs_runtime_Value_1843542330() gopurs_runtime.Value {
	once_first__gopurs_runtime_Value_1843542330.Do(func() {
		cache_first__gopurs_runtime_Value_1843542330 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_first__gopurs_runtime_Value_1843542330(gopurs_runtime.CoerceToStruct[Constructor_Strong[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_first__gopurs_runtime_Value_1843542330
}

var cache_splitStrong gopurs_runtime.Value
var once_splitStrong sync.Once
func Get_splitStrong() gopurs_runtime.Value {
	once_splitStrong.Do(func() {
		cache_splitStrong = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictStrong_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitStrong(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Strong[gopurs_runtime.Value]](dictStrong_1_box), l_2_box, r_3_box)
})
	})
	return cache_splitStrong
}

var cache_splitStrong__gopurs_runtime_Value_2623652703 gopurs_runtime.Value
var once_splitStrong__gopurs_runtime_Value_2623652703 sync.Once
func Get_splitStrong__gopurs_runtime_Value_2623652703() gopurs_runtime.Value {
	once_splitStrong__gopurs_runtime_Value_2623652703.Do(func() {
		cache_splitStrong__gopurs_runtime_Value_2623652703 = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictStrong_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitStrong__gopurs_runtime_Value_2623652703(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Strong[gopurs_runtime.Value]](dictStrong_1_box), l_2_box, r_3_box)
})
	})
	return cache_splitStrong__gopurs_runtime_Value_2623652703
}

var cache_fanout gopurs_runtime.Value
var once_fanout sync.Once
func Get_fanout() gopurs_runtime.Value {
	once_fanout.Do(func() {
		cache_fanout = gopurs_runtime.Func2(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictStrong_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fanout(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Strong[gopurs_runtime.Value]](dictStrong_1_box))
})
	})
	return cache_fanout
}

type Constructor_Strong[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1323482783] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Strong[gopurs_runtime.Value])(ptr)
		switch key {
		case "Profunctor0": return c.V0
		case "first": return c.V1
		case "second": return c.V2
		default: panic("Key not found in dictionary Constructor_Strong: " + key)
		}
	}
}


func Call_second(dict_0_loop *Constructor_Strong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Strong[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_second__gopurs_runtime_Value_1843542330(dict_0_loop *Constructor_Strong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Strong[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_first(dict_0_loop *Constructor_Strong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Strong[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_first__gopurs_runtime_Value_1843542330(dict_0_loop *Constructor_Strong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Strong[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_splitStrong(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], dictStrong_1_loop *Constructor_Strong[gopurs_runtime.Value], l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictStrong_1 *Constructor_Strong[gopurs_runtime.Value] = dictStrong_1_loop
_ = dictStrong_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, gopurs_runtime.Apply(dictStrong_1.V2, r_3), gopurs_runtime.Apply(dictStrong_1.V1, l_2))
}

func Call_splitStrong__gopurs_runtime_Value_2623652703(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], dictStrong_1_loop *Constructor_Strong[gopurs_runtime.Value], l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictStrong_1 *Constructor_Strong[gopurs_runtime.Value] = dictStrong_1_loop
_ = dictStrong_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, gopurs_runtime.Apply(dictStrong_1.V2, r_3), gopurs_runtime.Apply(dictStrong_1.V1, l_2))
}

func Call_fanout(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], dictStrong_1_loop *Constructor_Strong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictStrong_1 *Constructor_Strong[gopurs_runtime.Value] = dictStrong_1_loop
_ = dictStrong_1
Profunctor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictStrong_1.V0, gopurs_runtime.Value{}))
_ = Profunctor0_2_0
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Profunctor0_2_0.V0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, a_5})}))}
}), pkg_Data_Profunctor.Get_identity(), gopurs_runtime.Apply2(dictSemigroupoid_0.V0, gopurs_runtime.Apply(dictStrong_1.V2, r_4), gopurs_runtime.Apply(dictStrong_1.V1, l_3)))
})
})
}


