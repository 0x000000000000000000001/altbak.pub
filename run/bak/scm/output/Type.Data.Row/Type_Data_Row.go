package Type_Data_Row

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var RProxy gopurs_runtime.Value
var once_RProxy sync.Once
func Get_RProxy() gopurs_runtime.Value {
	once_RProxy.Do(func() {
		RProxy = gopurs_runtime.Constructor0("RProxy")
	})
	return RProxy
}




