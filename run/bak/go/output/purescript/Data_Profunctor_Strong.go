package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Strong_Strong_dollarDict gopurs_runtime.Value
var once_Data_Profunctor_Strong_Strong_dollarDict sync.Once
func Get_Data_Profunctor_Strong_Strong_dollarDict() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_Strong_dollarDict.Do(func() {
		cache_Data_Profunctor_Strong_Strong_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Strong_Strong_dollarDict(x_0_box)
})
	})
	return cache_Data_Profunctor_Strong_Strong_dollarDict
}

var cache_Data_Profunctor_Strong_strongFn gopurs_runtime.Value
var once_Data_Profunctor_Strong_strongFn sync.Once
func Get_Data_Profunctor_Strong_strongFn() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_strongFn.Do(func() {
		cache_Data_Profunctor_Strong_strongFn = gopurs_runtime.RecordDict3("Profunctor0", "first", "second", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Profunctor_profunctorFn()
}), gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(a2b_0, (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V1})}
})
}), gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map"))
	})
	return cache_Data_Profunctor_Strong_strongFn
}

var cache_Data_Profunctor_Strong_second gopurs_runtime.Value
var once_Data_Profunctor_Strong_second sync.Once
func Get_Data_Profunctor_Strong_second() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_second.Do(func() {
		cache_Data_Profunctor_Strong_second = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Strong_second(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong](dict_0_box))
})
	})
	return cache_Data_Profunctor_Strong_second
}

var cache_Data_Profunctor_Strong_first gopurs_runtime.Value
var once_Data_Profunctor_Strong_first sync.Once
func Get_Data_Profunctor_Strong_first() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_first.Do(func() {
		cache_Data_Profunctor_Strong_first = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Strong_first(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong](dict_0_box))
})
	})
	return cache_Data_Profunctor_Strong_first
}

var cache_Data_Profunctor_Strong_splitStrong gopurs_runtime.Value
var once_Data_Profunctor_Strong_splitStrong sync.Once
func Get_Data_Profunctor_Strong_splitStrong() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_splitStrong.Do(func() {
		cache_Data_Profunctor_Strong_splitStrong = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictStrong_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Strong_splitStrong(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong](dictStrong_1_box), l_2_box, r_3_box)
})
	})
	return cache_Data_Profunctor_Strong_splitStrong
}

var cache_Data_Profunctor_Strong_fanout gopurs_runtime.Value
var once_Data_Profunctor_Strong_fanout sync.Once
func Get_Data_Profunctor_Strong_fanout() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_fanout.Do(func() {
		cache_Data_Profunctor_Strong_fanout = gopurs_runtime.Func2(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictStrong_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Strong_fanout(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong](dictStrong_1_box))
})
	})
	return cache_Data_Profunctor_Strong_fanout
}

var cache_Data_Profunctor_Strong_first__1843542330 gopurs_runtime.Value
var once_Data_Profunctor_Strong_first__1843542330 sync.Once
func Get_Data_Profunctor_Strong_first__1843542330() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_first__1843542330.Do(func() {
		cache_Data_Profunctor_Strong_first__1843542330 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Strong_first__1843542330(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong](dict_0_box))
})
	})
	return cache_Data_Profunctor_Strong_first__1843542330
}

var cache_Data_Profunctor_Strong_second__1843542330 gopurs_runtime.Value
var once_Data_Profunctor_Strong_second__1843542330 sync.Once
func Get_Data_Profunctor_Strong_second__1843542330() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_second__1843542330.Do(func() {
		cache_Data_Profunctor_Strong_second__1843542330 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Strong_second__1843542330(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong](dict_0_box))
})
	})
	return cache_Data_Profunctor_Strong_second__1843542330
}

var cache_Data_Profunctor_Strong_splitStrong__2623652703 gopurs_runtime.Value
var once_Data_Profunctor_Strong_splitStrong__2623652703 sync.Once
func Get_Data_Profunctor_Strong_splitStrong__2623652703() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_splitStrong__2623652703.Do(func() {
		cache_Data_Profunctor_Strong_splitStrong__2623652703 = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictStrong_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Strong_splitStrong__2623652703(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong](dictStrong_1_box), l_2_box, r_3_box)
})
	})
	return cache_Data_Profunctor_Strong_splitStrong__2623652703
}

type Constructor_Data_Profunctor_Strong_Strong struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1323482783] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Strong_Strong)(ptr)
		_ = c
		switch key {
		case "Profunctor0": return gopurs_runtime.Box(c.V0)
		case "first": return gopurs_runtime.Box(c.V1)
		case "second": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Profunctor_Strong_Strong: " + key)
		}
	}
}


func Call_Data_Profunctor_Strong_Strong_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Strong_second(dict_0_loop *Constructor_Data_Profunctor_Strong_Strong) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Strong_Strong = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Profunctor_Strong_first(dict_0_loop *Constructor_Data_Profunctor_Strong_Strong) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Strong_Strong = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Profunctor_Strong_splitStrong(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid, dictStrong_1_loop *Constructor_Data_Profunctor_Strong_Strong, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictStrong_1 *Constructor_Data_Profunctor_Strong_Strong = dictStrong_1_loop
_ = dictStrong_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictStrong_1.V2), r_3), gopurs_runtime.Apply(gopurs_runtime.Box(dictStrong_1.V1), l_2))
}

func Call_Data_Profunctor_Strong_fanout(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid, dictStrong_1_loop *Constructor_Data_Profunctor_Strong_Strong) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictStrong_1 *Constructor_Data_Profunctor_Strong_Strong = dictStrong_1_loop
_ = dictStrong_1
// TAST (Let): Profunctor0_2_0 -> *Constructor_Data_Profunctor_Profunctor
Profunctor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor](gopurs_runtime.Apply(gopurs_runtime.Box(dictStrong_1.V0), gopurs_runtime.Value{}))
_ = Profunctor0_2_0
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(Profunctor0_2_0.V0), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, a_5})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
}), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictStrong_1.V2), r_4), gopurs_runtime.Apply(gopurs_runtime.Box(dictStrong_1.V1), l_3)))
})
})
}

func Call_Data_Profunctor_Strong_first__1843542330(dict_0_loop *Constructor_Data_Profunctor_Strong_Strong) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Strong_Strong = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Profunctor_Strong_second__1843542330(dict_0_loop *Constructor_Data_Profunctor_Strong_Strong) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Strong_Strong = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Profunctor_Strong_splitStrong__2623652703(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid, dictStrong_1_loop *Constructor_Data_Profunctor_Strong_Strong, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictStrong_1 *Constructor_Data_Profunctor_Strong_Strong = dictStrong_1_loop
_ = dictStrong_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictStrong_1.V2), r_3), gopurs_runtime.Apply(gopurs_runtime.Box(dictStrong_1.V1), l_2))
}


