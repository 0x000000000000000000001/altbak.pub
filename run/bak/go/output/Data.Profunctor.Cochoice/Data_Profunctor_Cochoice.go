package Data_Profunctor_Cochoice

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_unright gopurs_runtime.Value
var once_unright sync.Once
func Get_unright() gopurs_runtime.Value {
	once_unright.Do(func() {
		cache_unright = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "unright")
}()
})
	})
	return cache_unright
}

var cache_unleft gopurs_runtime.Value
var once_unleft sync.Once
func Get_unleft() gopurs_runtime.Value {
	once_unleft.Do(func() {
		cache_unleft = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "unleft")
}()
})
	})
	return cache_unleft
}




