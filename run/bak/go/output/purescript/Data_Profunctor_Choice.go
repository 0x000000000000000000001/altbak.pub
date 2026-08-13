package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Choice_identity gopurs_runtime.Value
var once_Data_Profunctor_Choice_identity sync.Once
func Get_Data_Profunctor_Choice_identity() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_identity.Do(func() {
		cache_Data_Profunctor_Choice_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_identity(x_0_box)
})
	})
	return cache_Data_Profunctor_Choice_identity
}

var cache_Data_Profunctor_Choice_Choice_dollarDict gopurs_runtime.Value
var once_Data_Profunctor_Choice_Choice_dollarDict sync.Once
func Get_Data_Profunctor_Choice_Choice_dollarDict() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_Choice_dollarDict.Do(func() {
		cache_Data_Profunctor_Choice_Choice_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_Choice_dollarDict(x_0_box)
})
	})
	return cache_Data_Profunctor_Choice_Choice_dollarDict
}

var cache_Data_Profunctor_Choice_right gopurs_runtime.Value
var once_Data_Profunctor_Choice_right sync.Once
func Get_Data_Profunctor_Choice_right() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_right.Do(func() {
		cache_Data_Profunctor_Choice_right = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_right(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_Choice_right
}

var cache_Data_Profunctor_Choice_left gopurs_runtime.Value
var once_Data_Profunctor_Choice_left sync.Once
func Get_Data_Profunctor_Choice_left() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_left.Do(func() {
		cache_Data_Profunctor_Choice_left = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_left(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_Choice_left
}

var cache_Data_Profunctor_Choice_splitChoice gopurs_runtime.Value
var once_Data_Profunctor_Choice_splitChoice sync.Once
func Get_Data_Profunctor_Choice_splitChoice() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_splitChoice.Do(func() {
		cache_Data_Profunctor_Choice_splitChoice = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictChoice_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_splitChoice(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dictChoice_1_box), l_2_box, r_3_box)
})
	})
	return cache_Data_Profunctor_Choice_splitChoice
}

var cache_Data_Profunctor_Choice_fanin gopurs_runtime.Value
var once_Data_Profunctor_Choice_fanin sync.Once
func Get_Data_Profunctor_Choice_fanin() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_fanin.Do(func() {
		cache_Data_Profunctor_Choice_fanin = gopurs_runtime.Func2(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictChoice_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_fanin(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dictChoice_1_box))
})
	})
	return cache_Data_Profunctor_Choice_fanin
}

var cache_Data_Profunctor_Choice_choiceFn gopurs_runtime.Value
var once_Data_Profunctor_Choice_choiceFn sync.Once
func Get_Data_Profunctor_Choice_choiceFn() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_choiceFn.Do(func() {
		cache_Data_Profunctor_Choice_choiceFn = gopurs_runtime.RecordDict3("Profunctor0", "left", "right", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Profunctor_profunctorFn()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"))
	})
	return cache_Data_Profunctor_Choice_choiceFn
}

var cache_Data_Profunctor_Choice_left__3408498362 gopurs_runtime.Value
var once_Data_Profunctor_Choice_left__3408498362 sync.Once
func Get_Data_Profunctor_Choice_left__3408498362() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_left__3408498362.Do(func() {
		cache_Data_Profunctor_Choice_left__3408498362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_left__3408498362(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_Choice_left__3408498362
}

var cache_Data_Profunctor_Choice_right__3408498362 gopurs_runtime.Value
var once_Data_Profunctor_Choice_right__3408498362 sync.Once
func Get_Data_Profunctor_Choice_right__3408498362() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_right__3408498362.Do(func() {
		cache_Data_Profunctor_Choice_right__3408498362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_right__3408498362(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_Choice_right__3408498362
}

var cache_Data_Profunctor_Choice_splitChoice__1826977503 gopurs_runtime.Value
var once_Data_Profunctor_Choice_splitChoice__1826977503 sync.Once
func Get_Data_Profunctor_Choice_splitChoice__1826977503() gopurs_runtime.Value {
	once_Data_Profunctor_Choice_splitChoice__1826977503.Do(func() {
		cache_Data_Profunctor_Choice_splitChoice__1826977503 = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictChoice_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Choice_splitChoice__1826977503(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]](dictChoice_1_box), l_2_box, r_3_box)
})
	})
	return cache_Data_Profunctor_Choice_splitChoice__1826977503
}

type Constructor_Data_Profunctor_Choice_Choice[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3666633887] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Profunctor0": return gopurs_runtime.Box(c.V0)
		case "left": return gopurs_runtime.Box(c.V1)
		case "right": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Profunctor_Choice_Choice: " + key)
		}
	}
}


func Call_Data_Profunctor_Choice_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Choice_Choice_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Choice_right(dict_0_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Profunctor_Choice_left(dict_0_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Profunctor_Choice_splitChoice(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value], dictChoice_1_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value], l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictChoice_1 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dictChoice_1_loop
_ = dictChoice_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictChoice_1.V2), r_3), gopurs_runtime.Apply(gopurs_runtime.Box(dictChoice_1.V1), l_2))
}

func Call_Data_Profunctor_Choice_fanin(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value], dictChoice_1_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictChoice_1 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dictChoice_1_loop
_ = dictChoice_1
// TAST (Let): Profunctor0_2_0 -> *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]
Profunctor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictChoice_1.V0), gopurs_runtime.Value{}))
_ = Profunctor0_2_0
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(Profunctor0_2_0.V0), Get_Data_Profunctor_identity1(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 3711209382) {
__t1 = (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2465973597) {
__t1 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictChoice_1.V2), r_4), gopurs_runtime.Apply(gopurs_runtime.Box(dictChoice_1.V1), l_3)))
})
})
}

func Call_Data_Profunctor_Choice_left__3408498362(dict_0_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Profunctor_Choice_right__3408498362(dict_0_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Profunctor_Choice_splitChoice__1826977503(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value], dictChoice_1_loop *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value], l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictChoice_1 *Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value] = dictChoice_1_loop
_ = dictChoice_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictChoice_1.V2), r_3), gopurs_runtime.Apply(gopurs_runtime.Box(dictChoice_1.V1), l_2))
}


