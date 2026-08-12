package Data_Exists

import (
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_runExists gopurs_runtime.Value
var once_runExists sync.Once
func Get_runExists() gopurs_runtime.Value {
	once_runExists.Do(func() {
		cache_runExists = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_runExists
}

var cache_mkExists gopurs_runtime.Value
var once_mkExists sync.Once
func Get_mkExists() gopurs_runtime.Value {
	once_mkExists.Do(func() {
		cache_mkExists = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_mkExists
}




