package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once
func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = Get_AppFFICheatcode_main()
	})
	return cache_Main_main
}




