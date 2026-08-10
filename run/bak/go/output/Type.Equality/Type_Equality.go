package Type_Equality

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
return Call_proof(dict_0_box)
})
	})
	return cache_proof
}

var cache_proof__gopurs_runtime_Value_3363032129 gopurs_runtime.Value
var once_proof__gopurs_runtime_Value_3363032129 sync.Once
func Get_proof__gopurs_runtime_Value_3363032129() gopurs_runtime.Value {
	once_proof__gopurs_runtime_Value_3363032129.Do(func() {
		cache_proof__gopurs_runtime_Value_3363032129 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_proof__gopurs_runtime_Value_3363032129(dict_0_box)
})
	})
	return cache_proof__gopurs_runtime_Value_3363032129
}

var cache_to gopurs_runtime.Value
var once_to sync.Once
func Get_to() gopurs_runtime.Value {
	once_to.Do(func() {
		cache_to = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_to(dictTypeEquals_0_box)
})
	})
	return cache_to
}

var cache_from gopurs_runtime.Value
var once_from sync.Once
func Get_from() gopurs_runtime.Value {
	once_from.Do(func() {
		cache_from = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_from(dictTypeEquals_0_box)
})
	})
	return cache_from
}

var cache_from__gopurs_runtime_Value_2366809570 gopurs_runtime.Value
var once_from__gopurs_runtime_Value_2366809570 sync.Once
func Get_from__gopurs_runtime_Value_2366809570() gopurs_runtime.Value {
	once_from__gopurs_runtime_Value_2366809570.Do(func() {
		cache_from__gopurs_runtime_Value_2366809570 = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_from__gopurs_runtime_Value_2366809570(dictTypeEquals_0_box)
})
	})
	return cache_from__gopurs_runtime_Value_2366809570
}

func Call_proof(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "proof")
}

func Call_proof__gopurs_runtime_Value_3363032129(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "proof")
}

func Call_to(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0, "proof"), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_from(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0, "proof"), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_from__gopurs_runtime_Value_2366809570(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0, "proof"), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}


