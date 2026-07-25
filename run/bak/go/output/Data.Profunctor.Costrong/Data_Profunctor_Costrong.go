package Data_Profunctor_Costrong

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_unsecond gopurs_runtime.Value
var once_unsecond sync.Once
func Get_unsecond() gopurs_runtime.Value {
	once_unsecond.Do(func() {
		cache_unsecond = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "unsecond")
}()
})
	})
	return cache_unsecond
}

var cache_unfirst gopurs_runtime.Value
var once_unfirst sync.Once
func Get_unfirst() gopurs_runtime.Value {
	once_unfirst.Do(func() {
		cache_unfirst = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "unfirst")
}()
})
	})
	return cache_unfirst
}




