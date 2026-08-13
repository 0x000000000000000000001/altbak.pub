package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Alternative_Alternative_dollarDict gopurs_runtime.Value
var once_Control_Alternative_Alternative_dollarDict sync.Once
func Get_Control_Alternative_Alternative_dollarDict() gopurs_runtime.Value {
	once_Control_Alternative_Alternative_dollarDict.Do(func() {
		cache_Control_Alternative_Alternative_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Alternative_Alternative_dollarDict(x_0_box)
})
	})
	return cache_Control_Alternative_Alternative_dollarDict
}

var cache_Control_Alternative_guard gopurs_runtime.Value
var once_Control_Alternative_guard sync.Once
func Get_Control_Alternative_guard() gopurs_runtime.Value {
	once_Control_Alternative_guard.Do(func() {
		cache_Control_Alternative_guard = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Alternative_guard(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Control_Alternative_guard
}

var cache_Control_Alternative_alternativeArray gopurs_runtime.Value
var once_Control_Alternative_alternativeArray sync.Once
func Get_Control_Alternative_alternativeArray() gopurs_runtime.Value {
	once_Control_Alternative_alternativeArray.Do(func() {
		cache_Control_Alternative_alternativeArray = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Control_Plus_plusArray()))}
})})}
	})
	return cache_Control_Alternative_alternativeArray
}

var cache_Control_Alternative_guard__2168855335 gopurs_runtime.Value
var once_Control_Alternative_guard__2168855335 sync.Once
func Get_Control_Alternative_guard__2168855335() gopurs_runtime.Value {
	once_Control_Alternative_guard__2168855335.Do(func() {
		cache_Control_Alternative_guard__2168855335 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Alternative_guard__2168855335(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Control_Alternative_guard__2168855335
}

var cache_Control_Alternative_guard__666381313 gopurs_runtime.Value
var once_Control_Alternative_guard__666381313 sync.Once
func Get_Control_Alternative_guard__666381313() gopurs_runtime.Value {
	once_Control_Alternative_guard__666381313.Do(func() {
		cache_Control_Alternative_guard__666381313 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Alternative_guard__666381313(__eta0_0_box)
})
	})
	return cache_Control_Alternative_guard__666381313
}

type Constructor_Control_Alternative_Alternative struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[397869517] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Alternative_Alternative)(ptr)
		_ = c
		switch key {
		case "Applicative0": return gopurs_runtime.Box(c.V0)
		case "Plus1": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Alternative_Alternative: " + key)
		}
	}
}


func Call_Control_Alternative_Alternative_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Alternative_guard(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): empty_2_1 -> gopurs_runtime.Value
empty_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "empty")
_ = empty_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
goto end_branch_2
} else {

}
}
{
__t2 = empty_2_1
}
end_branch_2:
return __t2
})
}

func Call_Control_Alternative_guard__2168855335(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): empty_2_1 -> gopurs_runtime.Value
empty_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "empty")
_ = empty_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
goto end_branch_2
} else {

}
}
{
__t2 = empty_2_1
}
end_branch_2:
return __t2
})
}

func Call_Control_Alternative_guard__666381313(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __t0 *Constructor_Data_Maybe_Just
{
if (__eta0_0.IntVal) != (0) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Maybe_applicativeMaybe()).V1), Get_Data_Unit_unit()))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_Maybe_plusMaybe()).V1))
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}


