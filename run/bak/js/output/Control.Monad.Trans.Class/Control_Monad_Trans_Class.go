package Control_Monad_Trans_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var lift gopurs_runtime.Value
var once_lift sync.Once
func Get_lift() gopurs_runtime.Value {
	once_lift.Do(func() {
		lift = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "lift")
})
	})
	return lift
}




