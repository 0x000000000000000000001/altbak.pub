package Control_Comonad_Store

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var store gopurs_runtime.Value
var once_store sync.Once
func Get_store() gopurs_runtime.Value {
	once_store.Do(func() {
		store = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": f_0, "value1": x_1})
})
})
	})
	return store
}

var runStore gopurs_runtime.Value
var once_runStore sync.Once
func Get_runStore() gopurs_runtime.Value {
	once_runStore.Do(func() {
		runStore = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
	})
	return runStore
}


