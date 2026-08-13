package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Type_Equality_TypeEquals_dollarDict gopurs_runtime.Value
var once_Type_Equality_TypeEquals_dollarDict sync.Once
func Get_Type_Equality_TypeEquals_dollarDict() gopurs_runtime.Value {
	once_Type_Equality_TypeEquals_dollarDict.Do(func() {
		cache_Type_Equality_TypeEquals_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_TypeEquals_dollarDict(x_0_box)
})
	})
	return cache_Type_Equality_TypeEquals_dollarDict
}

var cache_Type_Equality_To gopurs_runtime.Value
var once_Type_Equality_To sync.Once
func Get_Type_Equality_To() gopurs_runtime.Value {
	once_Type_Equality_To.Do(func() {
		cache_Type_Equality_To = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_To(x_0_box)
})
	})
	return cache_Type_Equality_To
}

var cache_Type_Equality_From gopurs_runtime.Value
var once_Type_Equality_From sync.Once
func Get_Type_Equality_From() gopurs_runtime.Value {
	once_Type_Equality_From.Do(func() {
		cache_Type_Equality_From = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_From(x_0_box)
})
	})
	return cache_Type_Equality_From
}

var cache_Type_Equality_refl gopurs_runtime.Value
var once_Type_Equality_refl sync.Once
func Get_Type_Equality_refl() gopurs_runtime.Value {
	once_Type_Equality_refl.Do(func() {
		cache_Type_Equality_refl = gopurs_runtime.RecordDict2("Coercible0", "proof", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
	})
	return cache_Type_Equality_refl
}

var cache_Type_Equality_proof gopurs_runtime.Value
var once_Type_Equality_proof sync.Once
func Get_Type_Equality_proof() gopurs_runtime.Value {
	once_Type_Equality_proof.Do(func() {
		cache_Type_Equality_proof = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_proof(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Type_Equality_proof
}

var cache_Type_Equality_to gopurs_runtime.Value
var once_Type_Equality_to sync.Once
func Get_Type_Equality_to() gopurs_runtime.Value {
	once_Type_Equality_to.Do(func() {
		cache_Type_Equality_to = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_to(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
})
	})
	return cache_Type_Equality_to
}

var cache_Type_Equality_from gopurs_runtime.Value
var once_Type_Equality_from sync.Once
func Get_Type_Equality_from() gopurs_runtime.Value {
	once_Type_Equality_from.Do(func() {
		cache_Type_Equality_from = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_from(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
})
	})
	return cache_Type_Equality_from
}

var cache_Type_Equality_from__4089948322 gopurs_runtime.Value
var once_Type_Equality_from__4089948322 sync.Once
func Get_Type_Equality_from__4089948322() gopurs_runtime.Value {
	once_Type_Equality_from__4089948322.Do(func() {
		cache_Type_Equality_from__4089948322 = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_from__4089948322(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
})
	})
	return cache_Type_Equality_from__4089948322
}

var cache_Type_Equality_from__2366809570 gopurs_runtime.Value
var once_Type_Equality_from__2366809570 sync.Once
func Get_Type_Equality_from__2366809570() gopurs_runtime.Value {
	once_Type_Equality_from__2366809570.Do(func() {
		cache_Type_Equality_from__2366809570 = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_from__2366809570(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
})
	})
	return cache_Type_Equality_from__2366809570
}

var cache_Type_Equality_proof__3118023201 gopurs_runtime.Value
var once_Type_Equality_proof__3118023201 sync.Once
func Get_Type_Equality_proof__3118023201() gopurs_runtime.Value {
	once_Type_Equality_proof__3118023201.Do(func() {
		cache_Type_Equality_proof__3118023201 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_proof__3118023201(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Type_Equality_proof__3118023201
}

var cache_Type_Equality_proof__3363032129 gopurs_runtime.Value
var once_Type_Equality_proof__3363032129 sync.Once
func Get_Type_Equality_proof__3363032129() gopurs_runtime.Value {
	once_Type_Equality_proof__3363032129.Do(func() {
		cache_Type_Equality_proof__3363032129 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Type_Equality_proof__3363032129(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Type_Equality_proof__3363032129
}

type Constructor_Type_Equality_TypeEquals[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3275391293] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Coercible0": return gopurs_runtime.Box(c.V0)
		case "proof": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Type_Equality_TypeEquals: " + key)
		}
	}
}


func Call_Type_Equality_TypeEquals_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Type_Equality_To(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Type_Equality_From(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Type_Equality_proof(dict_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Type_Equality_to(dictTypeEquals_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTypeEquals_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictTypeEquals_0.V1), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_Type_Equality_from(dictTypeEquals_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTypeEquals_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictTypeEquals_0.V1), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_Type_Equality_from__4089948322(dictTypeEquals_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTypeEquals_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictTypeEquals_0.V1), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_Type_Equality_from__2366809570(dictTypeEquals_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTypeEquals_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictTypeEquals_0.V1), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_Type_Equality_proof__3118023201(dict_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Type_Equality_proof__3363032129(dict_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}


