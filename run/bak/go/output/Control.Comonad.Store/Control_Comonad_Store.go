package Control_Comonad_Store

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var store gopurs_runtime.Value
var once_store sync.Once
func Get_store() gopurs_runtime.Value {
	once_store.Do(func() {
		store = gopurs_runtime.Func2(Call_store)
	})
	return store
}

var runStore gopurs_runtime.Value
var once_runStore sync.Once
func Get_runStore() gopurs_runtime.Value {
	once_runStore.Do(func() {
		runStore = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[1])
}()
})
	})
	return runStore
}

func Call_store(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Constructor2("Tuple", f_0_loop, x_1_loop)
}


