package Control_MonadPlus

import (
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_monadPlusArray gopurs_runtime.Value
var once_monadPlusArray sync.Once
func Get_monadPlusArray() gopurs_runtime.Value {
	once_monadPlusArray.Do(func() {
		cache_monadPlusArray = gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alternative.Get_alternativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadArray()
}))
	})
	return cache_monadPlusArray
}




