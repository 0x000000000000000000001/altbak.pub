package TestDump

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
)

var cache_main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		cache_main = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Hello"))
	})
	return cache_main
}




