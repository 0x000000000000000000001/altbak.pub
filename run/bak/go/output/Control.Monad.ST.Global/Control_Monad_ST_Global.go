package Control_Monad_ST_Global

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var cache_toEffect gopurs_runtime.Value
var once_toEffect sync.Once
func Get_toEffect() gopurs_runtime.Value {
	once_toEffect.Do(func() {
		cache_toEffect = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func() interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0())
})), nil))
}
}(func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, nil))
})())
})
})
	})
	return cache_toEffect
}


