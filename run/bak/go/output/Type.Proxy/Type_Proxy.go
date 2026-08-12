package Type_Proxy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Proxy gopurs_runtime.Value
var once_Proxy sync.Once
func Get_Proxy() gopurs_runtime.Value {
	once_Proxy.Do(func() {
		cache_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Proxy
}

type Constructor_Proxy[T_a any] struct {
	Rc uint32
}



