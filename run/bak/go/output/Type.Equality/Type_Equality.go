package Type_Equality

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_refl gopurs_runtime.Value
var once_refl sync.Once
func Get_refl() gopurs_runtime.Value {
	once_refl.Do(func() {
		cache_refl = gopurs_runtime.RecordDict2("proof", "Coercible0", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_refl
}

var cache_proof gopurs_runtime.Value
var once_proof sync.Once
func Get_proof() gopurs_runtime.Value {
	once_proof.Do(func() {
		cache_proof = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "proof")
}()
})
	})
	return cache_proof
}

var cache_to gopurs_runtime.Value
var once_to sync.Once
func Get_to() gopurs_runtime.Value {
	once_to.Do(func() {
		cache_to = gopurs_runtime.Func(func(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0, "proof"), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}()
})
	})
	return cache_to
}

var cache_from gopurs_runtime.Value
var once_from sync.Once
func Get_from() gopurs_runtime.Value {
	once_from.Do(func() {
		cache_from = gopurs_runtime.Func(func(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0, "proof"), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}()
})
	})
	return cache_from
}




