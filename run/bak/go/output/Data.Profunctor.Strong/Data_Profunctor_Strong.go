package Data_Profunctor_Strong

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_strongFn gopurs_runtime.Value
var once_strongFn sync.Once
func Get_strongFn() gopurs_runtime.Value {
	once_strongFn.Do(func() {
		cache_strongFn = gopurs_runtime.RecordDict3("Profunctor0", "first", "second", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
}), gopurs_runtime.Func2(func(a2b_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(a2b_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1})}
}), gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"))
	})
	return cache_strongFn
}

var cache_second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		cache_second = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_second(dict_0_box)
})
	})
	return cache_second
}

var cache_first gopurs_runtime.Value
var once_first sync.Once
func Get_first() gopurs_runtime.Value {
	once_first.Do(func() {
		cache_first = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_first(dict_0_box)
})
	})
	return cache_first
}

var cache_splitStrong gopurs_runtime.Value
var once_splitStrong sync.Once
func Get_splitStrong() gopurs_runtime.Value {
	once_splitStrong.Do(func() {
		cache_splitStrong = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictStrong_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitStrong(dictSemigroupoid_0_box, dictStrong_1_box, l_2_box, r_3_box)
})
	})
	return cache_splitStrong
}

var cache_fanout gopurs_runtime.Value
var once_fanout sync.Once
func Get_fanout() gopurs_runtime.Value {
	once_fanout.Do(func() {
		cache_fanout = gopurs_runtime.Func2(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictStrong_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fanout(dictSemigroupoid_0_box, dictStrong_1_box)
})
	})
	return cache_fanout
}

func Call_second(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData2)(dict_0.UnsafePtr)).V1
}

func Call_first(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData2)(dict_0.UnsafePtr)).V0
}

func Call_splitStrong(dictSemigroupoid_0_loop gopurs_runtime.Value, dictStrong_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictStrong_1 gopurs_runtime.Value = dictStrong_1_loop
_ = dictStrong_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictSemigroupoid_0.UnsafePtr)).V0, gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictStrong_1.UnsafePtr)).V1, r_3), gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictStrong_1.UnsafePtr)).V0, l_2))
}

func Call_fanout(dictSemigroupoid_0_loop gopurs_runtime.Value, dictStrong_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictStrong_1 gopurs_runtime.Value = dictStrong_1_loop
_ = dictStrong_1
lcmap_2_0 := gopurs_runtime.Apply(pkg_Data_Profunctor.Get_lcmap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_1, "Profunctor0"), gopurs_runtime.Value{}))
_ = lcmap_2_0
return gopurs_runtime.Func2(func(l_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(lcmap_2_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_5, a_5})}
}), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictSemigroupoid_0.UnsafePtr)).V0, gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictStrong_1.UnsafePtr)).V1, r_4), gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictStrong_1.UnsafePtr)).V0, l_3)))
})
}


