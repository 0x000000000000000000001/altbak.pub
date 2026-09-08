package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_Main_main gopurs_runtime.Value
var once_Test_Main_main sync.Once
func Get_Test_Main_main() gopurs_runtime.Value {
	once_Test_Main_main.Do(func() {
		cache_Test_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
	})
	return cache_Test_Main_main
}




