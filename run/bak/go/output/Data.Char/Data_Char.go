package Data_Char

import (
	pkg_Data_Enum "gopurs/output/Data.Enum"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_toCharCode gopurs_runtime.Value
var once_toCharCode sync.Once
func Get_toCharCode() gopurs_runtime.Value {
	once_toCharCode.Do(func() {
		cache_toCharCode = gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum")
	})
	return cache_toCharCode
}

var cache_fromCharCode gopurs_runtime.Value
var once_fromCharCode sync.Once
func Get_fromCharCode() gopurs_runtime.Value {
	once_fromCharCode.Do(func() {
		cache_fromCharCode = gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum")
	})
	return cache_fromCharCode
}




