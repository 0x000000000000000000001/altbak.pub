package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Reflectable_Reifiable_dollar_Dict gopurs_runtime.Value
var once_Data_Reflectable_Reifiable_dollar_Dict sync.Once
func Get_Data_Reflectable_Reifiable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Reflectable_Reifiable_dollar_Dict.Do(func() {
		cache_Data_Reflectable_Reifiable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Reflectable_Reifiable_dollar_Dict(func() struct{

} {
					orig := x_0_box
					_ = orig
					clone := struct{

}{}

					return clone
				}())), UnsafePtr: nil}
})
	})
	return cache_Data_Reflectable_Reifiable_dollar_Dict
}

var cache_Data_Reflectable_Reflectable_dollar_Dict gopurs_runtime.Value
var once_Data_Reflectable_Reflectable_dollar_Dict sync.Once
func Get_Data_Reflectable_Reflectable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Reflectable_Reflectable_dollar_Dict.Do(func() {
		cache_Data_Reflectable_Reflectable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 19771322, UnsafePtr: unsafe.Pointer(Call_Data_Reflectable_Reflectable_dollar_Dict(func() struct{
	reflectType gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	reflectType gopurs_runtime.Value
}{}
					clone.reflectType = gopurs_runtime.RecordGet(orig, "reflectType")
					return clone
				}()))}
})
	})
	return cache_Data_Reflectable_Reflectable_dollar_Dict
}

var cache_Data_Reflectable_reifiableString gopurs_runtime.Value
var once_Data_Reflectable_reifiableString sync.Once
func Get_Data_Reflectable_reifiableString() gopurs_runtime.Value {
	once_Data_Reflectable_reifiableString.Do(func() {
		cache_Data_Reflectable_reifiableString = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}().IntVal)), UnsafePtr: nil}
	})
	return cache_Data_Reflectable_reifiableString
}

var cache_Data_Reflectable_reifiableOrdering gopurs_runtime.Value
var once_Data_Reflectable_reifiableOrdering sync.Once
func Get_Data_Reflectable_reifiableOrdering() gopurs_runtime.Value {
	once_Data_Reflectable_reifiableOrdering.Do(func() {
		cache_Data_Reflectable_reifiableOrdering = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}().IntVal)), UnsafePtr: nil}
	})
	return cache_Data_Reflectable_reifiableOrdering
}

var cache_Data_Reflectable_reifiableInt gopurs_runtime.Value
var once_Data_Reflectable_reifiableInt sync.Once
func Get_Data_Reflectable_reifiableInt() gopurs_runtime.Value {
	once_Data_Reflectable_reifiableInt.Do(func() {
		cache_Data_Reflectable_reifiableInt = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}().IntVal)), UnsafePtr: nil}
	})
	return cache_Data_Reflectable_reifiableInt
}

var cache_Data_Reflectable_reifiableBoolean gopurs_runtime.Value
var once_Data_Reflectable_reifiableBoolean sync.Once
func Get_Data_Reflectable_reifiableBoolean() gopurs_runtime.Value {
	once_Data_Reflectable_reifiableBoolean.Do(func() {
		cache_Data_Reflectable_reifiableBoolean = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}().IntVal)), UnsafePtr: nil}
	})
	return cache_Data_Reflectable_reifiableBoolean
}

var cache_Data_Reflectable_reifyType gopurs_runtime.Value
var once_Data_Reflectable_reifyType sync.Once
func Get_Data_Reflectable_reifyType() gopurs_runtime.Value {
	once_Data_Reflectable_reifyType.Do(func() {
		cache_Data_Reflectable_reifyType = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Reflectable_reifyType(uint32(_dollar___unused_0_box.IntVal), s_1_box, f_2_box)
})
	})
	return cache_Data_Reflectable_reifyType
}

var cache_Data_Reflectable_reflectType gopurs_runtime.Value
var once_Data_Reflectable_reflectType sync.Once
func Get_Data_Reflectable_reflectType() gopurs_runtime.Value {
	once_Data_Reflectable_reflectType.Do(func() {
		cache_Data_Reflectable_reflectType = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Reflectable_reflectType(gopurs_runtime.CoerceToStruct[Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Reflectable_reflectType
}

type Constructor_Data_Reflectable_Reifiable[T_t any] struct {
	Rc uint32
}


func init() {
	gopurs_runtime.StructGetters[3870607684] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Reflectable_Reifiable[any])(ptr)
		_ = c
		switch key {

		default: panic("Key not found in dictionary Constructor_Data_Reflectable_Reifiable: " + key)
		}
	}
}


type Constructor_Data_Reflectable_Reflectable[T_v any, T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[19771322] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Reflectable_Reflectable[any, any])(ptr)
		_ = c
		switch key {
		case "reflectType": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Reflectable_Reflectable: " + key)
		}
	}
}


func Call_Data_Reflectable_Reifiable_dollar_Dict(x_0_loop struct{

}) uint32 {
var x_0 struct{

} = x_0_loop
_ = x_0
return uint32(func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}().IntVal)
}

func Call_Data_Reflectable_Reflectable_dollar_Dict(x_0_loop struct{
	reflectType gopurs_runtime.Value
}) *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	reflectType gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"reflectType"}, []gopurs_runtime.Value{orig.reflectType})
				}())
}

func Call_Data_Reflectable_reifyType(_dollar___unused_0_loop uint32, s_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 uint32 = _dollar___unused_0_loop
_ = _dollar___unused_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply3(Get_Data_Reflectable_unsafeCoerce(), gopurs_runtime.Func(func(dictReflectable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, dictReflectable_3)
}), func() gopurs_runtime.Value {
				orig := struct{
	reflectType gopurs_runtime.Value
}{gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return s_1
})}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"reflectType"}, []gopurs_runtime.Value{orig.reflectType})
				}(), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
}

func Call_Data_Reflectable_reflectType(dict_0_loop *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Get_Data_Reflectable_unsafeCoerce() gopurs_runtime.Value {
	return _Gopurs_Data_Reflectable_UnsafeCoerce
}
