package Data_Unit

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_unit gopurs_runtime.Value
var once_unit sync.Once
func Get_unit() gopurs_runtime.Value {
	once_unit.Do(func() {
		cache_unit = Unit
	})
	return cache_unit
}


