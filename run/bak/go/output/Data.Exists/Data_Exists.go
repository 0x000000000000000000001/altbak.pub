package Data_Exists

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var cache_runExists gopurs_runtime.Value
var once_runExists sync.Once
func Get_runExists() gopurs_runtime.Value {
	once_runExists.Do(func() {
		cache_runExists = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_runExists
}

var cache_runExists__gopurs_runtime_Value_2047851920 gopurs_runtime.Value
var once_runExists__gopurs_runtime_Value_2047851920 sync.Once
func Get_runExists__gopurs_runtime_Value_2047851920() gopurs_runtime.Value {
	once_runExists__gopurs_runtime_Value_2047851920.Do(func() {
		cache_runExists__gopurs_runtime_Value_2047851920 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_runExists__gopurs_runtime_Value_2047851920
}

var cache_mkExists gopurs_runtime.Value
var once_mkExists sync.Once
func Get_mkExists() gopurs_runtime.Value {
	once_mkExists.Do(func() {
		cache_mkExists = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_mkExists
}

var cache_mkExists__gopurs_runtime_Value_2238680401 gopurs_runtime.Value
var once_mkExists__gopurs_runtime_Value_2238680401 sync.Once
func Get_mkExists__gopurs_runtime_Value_2238680401() gopurs_runtime.Value {
	once_mkExists__gopurs_runtime_Value_2238680401.Do(func() {
		cache_mkExists__gopurs_runtime_Value_2238680401 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_mkExists__gopurs_runtime_Value_2238680401
}




