package Data_Profunctor_Closed

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
return Call_closed(gopurs_runtime.CoerceToStruct[Constructor_Closed[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_closed
}

type Constructor_Closed[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[768764671] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Closed[gopurs_runtime.Value])(ptr)
		switch key {
		case "Profunctor0": return c.V0
		case "closed": return c.V1
		default: panic("Key not found in dictionary Constructor_Closed: " + key)
		}
	}
}


func Call_closed(dict_0_loop *Constructor_Closed[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Closed[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


