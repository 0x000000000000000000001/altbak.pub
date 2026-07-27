package Data_Char

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_toCharCode gopurs_runtime.Value
var once_toCharCode sync.Once
func Get_toCharCode() gopurs_runtime.Value {
	once_toCharCode.Do(func() {
		cache_toCharCode = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 string) int64 {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), gopurs_runtime.Str(inner_arg0)).IntVal
}(arg0.StrVal()))
})
	})
	return cache_toCharCode
}

var cache_fromCharCode gopurs_runtime.Value
var once_fromCharCode sync.Once
func Get_fromCharCode() gopurs_runtime.Value {
	once_fromCharCode.Do(func() {
		cache_fromCharCode = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 int64) *pkg_Data_Maybe.Constructor_Just[string] {
return (*pkg_Data_Maybe.Constructor_Just[string])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), gopurs_runtime.Int(inner_arg0)).UnsafePtr)
}(arg0.IntVal))}
})
	})
	return cache_fromCharCode
}


