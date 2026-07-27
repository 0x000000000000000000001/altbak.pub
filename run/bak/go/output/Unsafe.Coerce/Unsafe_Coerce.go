package Unsafe_Coerce

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_unsafeCoerce gopurs_runtime.Value
var once_unsafeCoerce sync.Once
func Get_unsafeCoerce() gopurs_runtime.Value {
	once_unsafeCoerce.Do(func() {
		cache_unsafeCoerce = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(UnsafeCoerce(gopurs_runtime.UnboxAny(arg0)))
})
	})
	return cache_unsafeCoerce
}


