package Control_Monad_Trans_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_lift gopurs_runtime.Value
var once_lift sync.Once
func Get_lift() gopurs_runtime.Value {
	once_lift.Do(func() {
		cache_lift = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "lift")
}()
})
	})
	return cache_lift
}




