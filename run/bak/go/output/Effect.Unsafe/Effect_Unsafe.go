package Effect_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_unsafePerformEffect gopurs_runtime.Value
var once_unsafePerformEffect sync.Once
func Get_unsafePerformEffect() gopurs_runtime.Value {
	once_unsafePerformEffect.Do(func() {
		cache_unsafePerformEffect = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(UnsafePerformEffect(func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, nil))
}))
})
	})
	return cache_unsafePerformEffect
}


