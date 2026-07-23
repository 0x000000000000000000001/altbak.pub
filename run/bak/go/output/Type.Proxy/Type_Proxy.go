package Type_Proxy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Proxy gopurs_runtime.Value
var once_Proxy sync.Once
func Get_Proxy() gopurs_runtime.Value {
	once_Proxy.Do(func() {
		Proxy = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))
	})
	return Proxy
}


