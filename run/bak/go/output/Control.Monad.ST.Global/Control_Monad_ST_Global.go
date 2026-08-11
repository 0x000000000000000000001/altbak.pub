package Control_Monad_ST_Global

import (
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_toEffect gopurs_runtime.Value
var once_toEffect sync.Once
func Get_toEffect() gopurs_runtime.Value {
	once_toEffect.Do(func() {
		cache_toEffect = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_toEffect
}

var cache_toEffect__gopurs_runtime_Value_4169273813 gopurs_runtime.Value
var once_toEffect__gopurs_runtime_Value_4169273813 sync.Once
func Get_toEffect__gopurs_runtime_Value_4169273813() gopurs_runtime.Value {
	once_toEffect__gopurs_runtime_Value_4169273813.Do(func() {
		cache_toEffect__gopurs_runtime_Value_4169273813 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_toEffect__gopurs_runtime_Value_4169273813
}




