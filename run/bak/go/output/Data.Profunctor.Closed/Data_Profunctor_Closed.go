package Data_Profunctor_Closed

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
)

var closedFunction gopurs_runtime.Value
var once_closedFunction sync.Once
func Get_closedFunction() gopurs_runtime.Value {
	once_closedFunction.Do(func() {
		closedFunction = gopurs_runtime.RecordDict2("closed", "Profunctor0", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
}))
	})
	return closedFunction
}

var closed gopurs_runtime.Value
var once_closed sync.Once
func Get_closed() gopurs_runtime.Value {
	once_closed.Do(func() {
		closed = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "closed")
}()
})
	})
	return closed
}




