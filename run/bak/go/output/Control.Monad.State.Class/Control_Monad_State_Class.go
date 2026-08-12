package Control_Monad_State_Class

import (
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_state gopurs_runtime.Value
var once_state sync.Once
func Get_state() gopurs_runtime.Value {
	once_state.Do(func() {
		cache_state = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_state(gopurs_runtime.CoerceToStruct[Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_state
}

var cache_state__gopurs_runtime_Value_3572857840 gopurs_runtime.Value
var once_state__gopurs_runtime_Value_3572857840 sync.Once
func Get_state__gopurs_runtime_Value_3572857840() gopurs_runtime.Value {
	once_state__gopurs_runtime_Value_3572857840.Do(func() {
		cache_state__gopurs_runtime_Value_3572857840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_state__gopurs_runtime_Value_3572857840(gopurs_runtime.CoerceToStruct[Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_state__gopurs_runtime_Value_3572857840
}

var cache_put gopurs_runtime.Value
var once_put sync.Once
func Get_put() gopurs_runtime.Value {
	once_put.Do(func() {
		cache_put = gopurs_runtime.Func2(func(dictMonadState_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_put(gopurs_runtime.CoerceToStruct[Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0_box), s_1_box)
})
	})
	return cache_put
}

var cache_modify_ gopurs_runtime.Value
var once_modify_ sync.Once
func Get_modify_() gopurs_runtime.Value {
	once_modify_.Do(func() {
		cache_modify_ = gopurs_runtime.Func2(func(dictMonadState_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify_(gopurs_runtime.CoerceToStruct[Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0_box), f_1_box)
})
	})
	return cache_modify_
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func2(func(dictMonadState_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(gopurs_runtime.CoerceToStruct[Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0_box), f_1_box)
})
	})
	return cache_modify
}

var cache_gets gopurs_runtime.Value
var once_gets sync.Once
func Get_gets() gopurs_runtime.Value {
	once_gets.Do(func() {
		cache_gets = gopurs_runtime.Func2(func(dictMonadState_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_gets(gopurs_runtime.CoerceToStruct[Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0_box), f_1_box)
})
	})
	return cache_gets
}

var cache_get gopurs_runtime.Value
var once_get sync.Once
func Get_get() gopurs_runtime.Value {
	once_get.Do(func() {
		cache_get = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get(dictMonadState_0_box)
})
	})
	return cache_get
}

type Constructor_MonadState[T_s any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2100320995] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "state": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadState: " + key)
		}
	}
}


func Call_state(dict_0_loop *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_state__gopurs_runtime_Value_3572857840(dict_0_loop *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_put(dictMonadState_0_loop *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value], s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_0_loop
_ = dictMonadState_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(dictMonadState_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit(), s_1})}))}
}))
}

func Call_modify_(dictMonadState_0_loop *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_0_loop
_ = dictMonadState_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(dictMonadState_0.V1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit(), gopurs_runtime.Apply(f_1, s_2)})}))}
}))
}

func Call_modify(dictMonadState_0_loop *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_0_loop
_ = dictMonadState_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(dictMonadState_0.V1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_3_0 := gopurs_runtime.Apply(f_1, s_2)
_ = s_prime_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_prime_3_0, s_prime_3_0})}))}
}))
}

func Call_gets(dictMonadState_0_loop *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 *Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_0_loop
_ = dictMonadState_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(dictMonadState_0.V1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, s_2), s_2})}))}
}))
}

func Call_get(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_1, s_1})}))}
}))
}


