package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Closed_Closed_dollarDict gopurs_runtime.Value
var once_Data_Profunctor_Closed_Closed_dollarDict sync.Once
func Get_Data_Profunctor_Closed_Closed_dollarDict() gopurs_runtime.Value {
	once_Data_Profunctor_Closed_Closed_dollarDict.Do(func() {
		cache_Data_Profunctor_Closed_Closed_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Closed_Closed_dollarDict(x_0_box)
})
	})
	return cache_Data_Profunctor_Closed_Closed_dollarDict
}

var cache_Data_Profunctor_Closed_closedFunction gopurs_runtime.Value
var once_Data_Profunctor_Closed_closedFunction sync.Once
func Get_Data_Profunctor_Closed_closedFunction() gopurs_runtime.Value {
	once_Data_Profunctor_Closed_closedFunction.Do(func() {
		cache_Data_Profunctor_Closed_closedFunction = gopurs_runtime.RecordDict2("Profunctor0", "closed", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Profunctor_profunctorFn()
}), gopurs_runtime.RecordGet(Get_Control_Semigroupoid_semigroupoidFn(), "compose"))
	})
	return cache_Data_Profunctor_Closed_closedFunction
}

var cache_Data_Profunctor_Closed_closed gopurs_runtime.Value
var once_Data_Profunctor_Closed_closed sync.Once
func Get_Data_Profunctor_Closed_closed() gopurs_runtime.Value {
	once_Data_Profunctor_Closed_closed.Do(func() {
		cache_Data_Profunctor_Closed_closed = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Closed_closed(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Closed_Closed](dict_0_box))
})
	})
	return cache_Data_Profunctor_Closed_closed
}

type Constructor_Data_Profunctor_Closed_Closed struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[768764671] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Closed_Closed)(ptr)
		_ = c
		switch key {
		case "Profunctor0": return gopurs_runtime.Box(c.V0)
		case "closed": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Profunctor_Closed_Closed: " + key)
		}
	}
}


func Call_Data_Profunctor_Closed_Closed_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Closed_closed(dict_0_loop *Constructor_Data_Profunctor_Closed_Closed) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Closed_Closed = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}


