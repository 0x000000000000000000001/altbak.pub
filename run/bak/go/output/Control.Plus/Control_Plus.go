package Control_Plus

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_plusArray gopurs_runtime.Value
var once_plusArray sync.Once
func Get_plusArray() gopurs_runtime.Value {
	once_plusArray.Do(func() {
		cache_plusArray = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alt.Get_altArray()
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
	})
	return cache_plusArray
}

var cache_plusArray__gopurs_runtime_Value_4260531026 gopurs_runtime.Value
var once_plusArray__gopurs_runtime_Value_4260531026 sync.Once
func Get_plusArray__gopurs_runtime_Value_4260531026() gopurs_runtime.Value {
	once_plusArray__gopurs_runtime_Value_4260531026.Do(func() {
		cache_plusArray__gopurs_runtime_Value_4260531026 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alt.Get_altArray__gopurs_runtime_Value_2010533188()
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
	})
	return cache_plusArray__gopurs_runtime_Value_4260531026
}

var cache_empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		cache_empty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_empty(dict_0_box)
})
	})
	return cache_empty
}

var cache_empty__gopurs_runtime_Value_932402776 gopurs_runtime.Value
var once_empty__gopurs_runtime_Value_932402776 sync.Once
func Get_empty__gopurs_runtime_Value_932402776() gopurs_runtime.Value {
	once_empty__gopurs_runtime_Value_932402776.Do(func() {
		cache_empty__gopurs_runtime_Value_932402776 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_empty__gopurs_runtime_Value_932402776(dict_0_box)
})
	})
	return cache_empty__gopurs_runtime_Value_932402776
}

type Constructor_Plus[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3709470893] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Plus[gopurs_runtime.Value])(ptr)
		switch key {
		case "Alt0": return c.V0
		case "empty": return c.V1
		default: panic("Key not found in dictionary Constructor_Plus: " + key)
		}
	}
}


func Call_empty(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "empty")
}

func Call_empty__gopurs_runtime_Value_932402776(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "empty")
}


