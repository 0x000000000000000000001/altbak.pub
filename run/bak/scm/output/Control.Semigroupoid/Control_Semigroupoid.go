package Control_Semigroupoid

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var semigroupoidFn gopurs_runtime.Value
var once_semigroupoidFn sync.Once
func Get_semigroupoidFn() gopurs_runtime.Value {
	once_semigroupoidFn.Do(func() {
		semigroupoidFn = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
}))
	})
	return semigroupoidFn
}

var compose gopurs_runtime.Value
var once_compose sync.Once
func Get_compose() gopurs_runtime.Value {
	once_compose.Do(func() {
		compose = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "compose")
})
	})
	return compose
}

var composeFlipped gopurs_runtime.Value
var once_composeFlipped sync.Once
func Get_composeFlipped() gopurs_runtime.Value {
	once_composeFlipped.Do(func() {
		composeFlipped = gopurs_runtime.Func3(func(dictSemigroupoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), g_2, f_1)
})
	})
	return composeFlipped
}




