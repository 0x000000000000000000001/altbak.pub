package Control_Comonad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var extract gopurs_runtime.Value
var once_extract sync.Once
func Get_extract() gopurs_runtime.Value {
	once_extract.Do(func() {
		extract = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "extract")
})
	})
	return extract
}




