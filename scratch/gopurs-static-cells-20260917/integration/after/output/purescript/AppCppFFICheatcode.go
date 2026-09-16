package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_AppCppFFICheatcode_main gopurs_runtime.Value
var once_AppCppFFICheatcode_main sync.Once
func Get_AppCppFFICheatcode_main() gopurs_runtime.Value {
	once_AppCppFFICheatcode_main.Do(func() {
		cache_AppCppFFICheatcode_main = Get_AppFFICheatcode_main()
	})
	return cache_AppCppFFICheatcode_main
}




