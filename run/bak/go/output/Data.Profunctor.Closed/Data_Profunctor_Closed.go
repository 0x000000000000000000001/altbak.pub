package Data_Profunctor_Closed

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
)

var cache_closedFunction gopurs_runtime.Value
var once_closedFunction sync.Once
func Get_closedFunction() gopurs_runtime.Value {
	once_closedFunction.Do(func() {
		cache_closedFunction = gopurs_runtime.RecordDict2("Profunctor0", "closed", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
}), gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_closedFunction
}

var cache_closed gopurs_runtime.Value
var once_closed sync.Once
func Get_closed() gopurs_runtime.Value {
	once_closed.Do(func() {
		cache_closed = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_closed(dict_0_box)
})
	})
	return cache_closed
}

func Call_closed(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "closed")
}


