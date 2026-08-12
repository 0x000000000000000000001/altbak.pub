package Control_Alternative

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		cache_guard = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_guard(gopurs_runtime.CoerceToStruct[Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_guard
}

var cache_guard__gopurs_runtime_Value_2168855335 gopurs_runtime.Value
var once_guard__gopurs_runtime_Value_2168855335 sync.Once
func Get_guard__gopurs_runtime_Value_2168855335() gopurs_runtime.Value {
	once_guard__gopurs_runtime_Value_2168855335.Do(func() {
		cache_guard__gopurs_runtime_Value_2168855335 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_guard__gopurs_runtime_Value_2168855335(gopurs_runtime.CoerceToStruct[Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_guard__gopurs_runtime_Value_2168855335
}

var cache_alternativeArray gopurs_runtime.Value
var once_alternativeArray sync.Once
func Get_alternativeArray() gopurs_runtime.Value {
	once_alternativeArray.Do(func() {
		cache_alternativeArray = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray()
}))
	})
	return cache_alternativeArray
}

var cache_alternativeArray__ptrConstructor_Alternative_gopurs_runtime_Value__2415002109 gopurs_runtime.Value
var once_alternativeArray__ptrConstructor_Alternative_gopurs_runtime_Value__2415002109 sync.Once
func Get_alternativeArray__ptrConstructor_Alternative_gopurs_runtime_Value__2415002109() gopurs_runtime.Value {
	once_alternativeArray__ptrConstructor_Alternative_gopurs_runtime_Value__2415002109.Do(func() {
		cache_alternativeArray__ptrConstructor_Alternative_gopurs_runtime_Value__2415002109 = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray()
})})}
	})
	return cache_alternativeArray__ptrConstructor_Alternative_gopurs_runtime_Value__2415002109
}

var cache_alternativeArray__gopurs_runtime_Value_1996030013 gopurs_runtime.Value
var once_alternativeArray__gopurs_runtime_Value_1996030013 sync.Once
func Get_alternativeArray__gopurs_runtime_Value_1996030013() gopurs_runtime.Value {
	once_alternativeArray__gopurs_runtime_Value_1996030013.Do(func() {
		cache_alternativeArray__gopurs_runtime_Value_1996030013 = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray__gopurs_runtime_Value_1604836744()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray__gopurs_runtime_Value_4260531026()
}))
	})
	return cache_alternativeArray__gopurs_runtime_Value_1996030013
}

type Constructor_Alternative[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[397869517] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Alternative[gopurs_runtime.Value])(ptr)
		switch key {
		case "Applicative0": return c.V0
		case "Plus1": return c.V1
		default: panic("Key not found in dictionary Constructor_Alternative: " + key)
		}
	}
}


func Call_guard(dictAlternative_0_loop *Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
empty_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "empty")
_ = empty_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_Unit.Get_unit())
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

func Call_guard__gopurs_runtime_Value_2168855335(dictAlternative_0_loop *Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
empty_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "empty")
_ = empty_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_Unit.Get_unit())
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


