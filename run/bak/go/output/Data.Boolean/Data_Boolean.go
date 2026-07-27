package Data_Boolean

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_otherwise gopurs_runtime.Value
var once_otherwise sync.Once
func Get_otherwise() gopurs_runtime.Value {
	once_otherwise.Do(func() {
		cache_otherwise = gopurs_runtime.Bool(true)
	})
	return cache_otherwise
}


