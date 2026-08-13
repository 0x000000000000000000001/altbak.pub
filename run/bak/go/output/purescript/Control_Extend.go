package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Extend_identity gopurs_runtime.Value
var once_Control_Extend_identity sync.Once
func Get_Control_Extend_identity() gopurs_runtime.Value {
	once_Control_Extend_identity.Do(func() {
		cache_Control_Extend_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_identity(x_0_box)
})
	})
	return cache_Control_Extend_identity
}

var cache_Control_Extend_Extend_dollarDict gopurs_runtime.Value
var once_Control_Extend_Extend_dollarDict sync.Once
func Get_Control_Extend_Extend_dollarDict() gopurs_runtime.Value {
	once_Control_Extend_Extend_dollarDict.Do(func() {
		cache_Control_Extend_Extend_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_Extend_dollarDict(x_0_box)
})
	})
	return cache_Control_Extend_Extend_dollarDict
}

var cache_Control_Extend_extendFn gopurs_runtime.Value
var once_Control_Extend_extendFn sync.Once
func Get_Control_Extend_extendFn() gopurs_runtime.Value {
	once_Control_Extend_extendFn.Do(func() {
		cache_Control_Extend_extendFn = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_extendFn(dictSemigroup_0_box)
})
	})
	return cache_Control_Extend_extendFn
}

var cache_Control_Extend_extendArray gopurs_runtime.Value
var once_Control_Extend_extendArray sync.Once
func Get_Control_Extend_extendArray() gopurs_runtime.Value {
	once_Control_Extend_extendArray.Do(func() {
		cache_Control_Extend_extendArray = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), Get_Control_Extend_arrayExtend())
	})
	return cache_Control_Extend_extendArray
}

var cache_Control_Extend_extend gopurs_runtime.Value
var once_Control_Extend_extend sync.Once
func Get_Control_Extend_extend() gopurs_runtime.Value {
	once_Control_Extend_extend.Do(func() {
		cache_Control_Extend_extend = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_extend(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dict_0_box))
})
	})
	return cache_Control_Extend_extend
}

var cache_Control_Extend_extendFlipped gopurs_runtime.Value
var once_Control_Extend_extendFlipped sync.Once
func Get_Control_Extend_extendFlipped() gopurs_runtime.Value {
	once_Control_Extend_extendFlipped.Do(func() {
		cache_Control_Extend_extendFlipped = gopurs_runtime.Func3(func(dictExtend_0_box gopurs_runtime.Value, w_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_extendFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dictExtend_0_box), w_1_box, f_2_box)
})
	})
	return cache_Control_Extend_extendFlipped
}

var cache_Control_Extend_duplicate gopurs_runtime.Value
var once_Control_Extend_duplicate sync.Once
func Get_Control_Extend_duplicate() gopurs_runtime.Value {
	once_Control_Extend_duplicate.Do(func() {
		cache_Control_Extend_duplicate = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_duplicate(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dictExtend_0_box))
})
	})
	return cache_Control_Extend_duplicate
}

var cache_Control_Extend_composeCoKleisliFlipped gopurs_runtime.Value
var once_Control_Extend_composeCoKleisliFlipped sync.Once
func Get_Control_Extend_composeCoKleisliFlipped() gopurs_runtime.Value {
	once_Control_Extend_composeCoKleisliFlipped.Do(func() {
		cache_Control_Extend_composeCoKleisliFlipped = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_composeCoKleisliFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dictExtend_0_box), f_1_box, g_2_box, w_3_box)
})
	})
	return cache_Control_Extend_composeCoKleisliFlipped
}

var cache_Control_Extend_composeCoKleisli gopurs_runtime.Value
var once_Control_Extend_composeCoKleisli sync.Once
func Get_Control_Extend_composeCoKleisli() gopurs_runtime.Value {
	once_Control_Extend_composeCoKleisli.Do(func() {
		cache_Control_Extend_composeCoKleisli = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_composeCoKleisli(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dictExtend_0_box), f_1_box, g_2_box, w_3_box)
})
	})
	return cache_Control_Extend_composeCoKleisli
}

var cache_Control_Extend_composeCoKleisliFlipped__1582554720 gopurs_runtime.Value
var once_Control_Extend_composeCoKleisliFlipped__1582554720 sync.Once
func Get_Control_Extend_composeCoKleisliFlipped__1582554720() gopurs_runtime.Value {
	once_Control_Extend_composeCoKleisliFlipped__1582554720.Do(func() {
		cache_Control_Extend_composeCoKleisliFlipped__1582554720 = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_composeCoKleisliFlipped__1582554720(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dictExtend_0_box), f_1_box, g_2_box, w_3_box)
})
	})
	return cache_Control_Extend_composeCoKleisliFlipped__1582554720
}

var cache_Control_Extend_extend__1264481661 gopurs_runtime.Value
var once_Control_Extend_extend__1264481661 sync.Once
func Get_Control_Extend_extend__1264481661() gopurs_runtime.Value {
	once_Control_Extend_extend__1264481661.Do(func() {
		cache_Control_Extend_extend__1264481661 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_extend__1264481661(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dict_0_box))
})
	})
	return cache_Control_Extend_extend__1264481661
}

var cache_Control_Extend_extend__3641500541 gopurs_runtime.Value
var once_Control_Extend_extend__3641500541 sync.Once
func Get_Control_Extend_extend__3641500541() gopurs_runtime.Value {
	once_Control_Extend_extend__3641500541.Do(func() {
		cache_Control_Extend_extend__3641500541 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_extend__3641500541(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dict_0_box))
})
	})
	return cache_Control_Extend_extend__3641500541
}

var cache_Control_Extend_extend__267444733 gopurs_runtime.Value
var once_Control_Extend_extend__267444733 sync.Once
func Get_Control_Extend_extend__267444733() gopurs_runtime.Value {
	once_Control_Extend_extend__267444733.Do(func() {
		cache_Control_Extend_extend__267444733 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_extend__267444733(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dict_0_box))
})
	})
	return cache_Control_Extend_extend__267444733
}

var cache_Control_Extend_extend__1965081501 gopurs_runtime.Value
var once_Control_Extend_extend__1965081501 sync.Once
func Get_Control_Extend_extend__1965081501() gopurs_runtime.Value {
	once_Control_Extend_extend__1965081501.Do(func() {
		cache_Control_Extend_extend__1965081501 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_extend__1965081501(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](dict_0_box))
})
	})
	return cache_Control_Extend_extend__1965081501
}

var cache_Control_Extend_extend__4254185051 gopurs_runtime.Value
var once_Control_Extend_extend__4254185051 sync.Once
func Get_Control_Extend_extend__4254185051() gopurs_runtime.Value {
	once_Control_Extend_extend__4254185051.Do(func() {
		cache_Control_Extend_extend__4254185051 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Extend_extend__4254185051(v_0_box, v1_1_box)
})
	})
	return cache_Control_Extend_extend__4254185051
}

type Constructor_Control_Extend_Extend struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3028639021] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Extend_Extend)(ptr)
		_ = c
		switch key {
		case "Functor0": return gopurs_runtime.Box(c.V0)
		case "extend": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Extend_Extend: " + key)
		}
	}
}


func Call_Control_Extend_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Extend_Extend_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Extend_extendFn(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorFn()
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Func(func(w_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), w_3, w_prime_4))
}))
})
})
}))
}

func Call_Control_Extend_extend(dict_0_loop *Constructor_Control_Extend_Extend) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Extend_Extend = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Extend_extendFlipped(dictExtend_0_loop *Constructor_Control_Extend_Extend, w_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 *Constructor_Control_Extend_Extend = dictExtend_0_loop
_ = dictExtend_0
var w_1 gopurs_runtime.Value = w_1_loop
_ = w_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictExtend_0.V1), f_2, w_1)
}

func Call_Control_Extend_duplicate(dictExtend_0_loop *Constructor_Control_Extend_Extend) gopurs_runtime.Value {
var dictExtend_0 *Constructor_Control_Extend_Extend = dictExtend_0_loop
_ = dictExtend_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictExtend_0.V1), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Control_Extend_composeCoKleisliFlipped(dictExtend_0_loop *Constructor_Control_Extend_Extend, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, w_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 *Constructor_Control_Extend_Extend = dictExtend_0_loop
_ = dictExtend_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var w_3 gopurs_runtime.Value = w_3_loop
_ = w_3
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(gopurs_runtime.Box(dictExtend_0.V1), g_2, w_3))
}

func Call_Control_Extend_composeCoKleisli(dictExtend_0_loop *Constructor_Control_Extend_Extend, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, w_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 *Constructor_Control_Extend_Extend = dictExtend_0_loop
_ = dictExtend_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var w_3 gopurs_runtime.Value = w_3_loop
_ = w_3
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply2(gopurs_runtime.Box(dictExtend_0.V1), f_1, w_3))
}

func Call_Control_Extend_composeCoKleisliFlipped__1582554720(dictExtend_0_loop *Constructor_Control_Extend_Extend, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, w_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 *Constructor_Control_Extend_Extend = dictExtend_0_loop
_ = dictExtend_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var w_3 gopurs_runtime.Value = w_3_loop
_ = w_3
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(gopurs_runtime.Box(dictExtend_0.V1), g_2, w_3))
}

func Call_Control_Extend_extend__1264481661(dict_0_loop *Constructor_Control_Extend_Extend) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Extend_Extend = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Extend_extend__3641500541(dict_0_loop *Constructor_Control_Extend_Extend) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Extend_Extend = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Extend_extend__267444733(dict_0_loop *Constructor_Control_Extend_Extend) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Extend_Extend = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Extend_extend__1965081501(dict_0_loop *Constructor_Control_Extend_Extend) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Extend_Extend = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Extend_extend__4254185051(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Apply2(Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1))
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Apply2(Get_Data_Interval_DurationEnd(), (*Constructor_Data_Interval_DurationEnd)(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1))
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Apply2(Get_Data_Interval_StartDuration(), gopurs_runtime.Apply(v_0, v1_1), (*Constructor_Data_Interval_StartDuration)(v1_1.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Apply(Get_Data_Interval_DurationOnly(), (*Constructor_Data_Interval_DurationOnly)(v1_1.UnsafePtr).V0)
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

func Get_Control_Extend_arrayExtend() gopurs_runtime.Value {
	return _Gopurs_Control_Extend_ArrayExtend
}
