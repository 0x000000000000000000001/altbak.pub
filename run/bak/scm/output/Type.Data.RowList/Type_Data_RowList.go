package Type_Data_RowList

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var RLProxy gopurs_runtime.Value
var once_RLProxy sync.Once
func Get_RLProxy() gopurs_runtime.Value {
	once_RLProxy.Do(func() {
		RLProxy = gopurs_runtime.Constructor0("RLProxy")
	})
	return RLProxy
}




