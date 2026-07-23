package Control_Comonad_Store

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var store gopurs_runtime.Value
var once_store sync.Once
func Get_store() gopurs_runtime.Value {
	once_store.Do(func() {
		store = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), f_0, x_1)
})
	})
	return store
}

var runStore gopurs_runtime.Value
var once_runStore sync.Once
func Get_runStore() gopurs_runtime.Value {
	once_runStore.Do(func() {
		runStore = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordGet(v_0, "value0"), gopurs_runtime.RecordGet(v_0, "value1"))
})
	})
	return runStore
}


