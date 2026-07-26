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
		cache_toEffect = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_toEffect
}




