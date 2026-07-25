package Control_Category

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "identity")
}()
})
	})
	return cache_identity
}

var cache_categoryFn gopurs_runtime.Value
var once_categoryFn sync.Once
func Get_categoryFn() gopurs_runtime.Value {
	once_categoryFn.Do(func() {
		cache_categoryFn = gopurs_runtime.RecordDict2("identity", "Semigroupoid0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Semigroupoid.Get_semigroupoidFn()
}))
	})
	return cache_categoryFn
}




