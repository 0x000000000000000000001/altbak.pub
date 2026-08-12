package Type_Equality

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_refl gopurs_runtime.Value
var once_refl sync.Once
func Get_refl() gopurs_runtime.Value {
	once_refl.Do(func() {
		cache_refl = gopurs_runtime.RecordDict2("Coercible0", "proof", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
	})
	return cache_refl
}

var cache_proof gopurs_runtime.Value
var once_proof sync.Once
func Get_proof() gopurs_runtime.Value {
	once_proof.Do(func() {
		cache_proof = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_proof(gopurs_runtime.CoerceToStruct[Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_proof
}

var cache_to gopurs_runtime.Value
var once_to sync.Once
func Get_to() gopurs_runtime.Value {
	once_to.Do(func() {
		cache_to = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_to(gopurs_runtime.CoerceToStruct[Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
})
	})
	return cache_to
}

var cache_from gopurs_runtime.Value
var once_from sync.Once
func Get_from() gopurs_runtime.Value {
	once_from.Do(func() {
		cache_from = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_from(gopurs_runtime.CoerceToStruct[Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
})
	})
	return cache_from
}

var cache_proof__3363032129 gopurs_runtime.Value
var once_proof__3363032129 sync.Once
func Get_proof__3363032129() gopurs_runtime.Value {
	once_proof__3363032129.Do(func() {
		cache_proof__3363032129 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_proof__3363032129(gopurs_runtime.CoerceToStruct[Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_proof__3363032129
}

type Constructor_TypeEquals[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3275391293] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Coercible0": return c.V0
		case "proof": return c.V1
		default: panic("Key not found in dictionary Constructor_TypeEquals: " + key)
		}
	}
}


func Call_proof(dict_0_loop *Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_to(dictTypeEquals_0_loop *Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTypeEquals_0 *Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(dictTypeEquals_0.V1, gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_from(dictTypeEquals_0_loop *Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTypeEquals_0 *Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(dictTypeEquals_0.V1, gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_proof__3363032129(dict_0_loop *Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


