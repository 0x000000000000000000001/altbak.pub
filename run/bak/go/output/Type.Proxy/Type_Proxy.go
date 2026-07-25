package Type_Proxy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Proxy gopurs_runtime.Value
var once_Proxy sync.Once
func Get_Proxy() gopurs_runtime.Value {
	once_Proxy.Do(func() {
		cache_Proxy = gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(&Data_Type_Proxy_Proxy{})}
	})
	return cache_Proxy
}

type Data_Type_Proxy_Proxy struct {
	
}
func Is_Data_Type_Proxy_Proxy(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 513803634
}


