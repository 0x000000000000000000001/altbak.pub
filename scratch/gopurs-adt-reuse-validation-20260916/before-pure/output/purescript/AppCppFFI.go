package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_AppCppFFI_main gopurs_runtime.Value
var once_AppCppFFI_main sync.Once
func Get_AppCppFFI_main() gopurs_runtime.Value {
	once_AppCppFFI_main.Do(func() {
		cache_AppCppFFI_main = Get_AppFFI_main()
	})
	return cache_AppCppFFI_main
}




