package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Plus_Plus_dollarDict gopurs_runtime.Value
var once_Control_Plus_Plus_dollarDict sync.Once
func Get_Control_Plus_Plus_dollarDict() gopurs_runtime.Value {
	once_Control_Plus_Plus_dollarDict.Do(func() {
		cache_Control_Plus_Plus_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Plus_Plus_dollarDict(x_0_box)
})
	})
	return cache_Control_Plus_Plus_dollarDict
}

var cache_Control_Plus_plusArray gopurs_runtime.Value
var once_Control_Plus_plusArray sync.Once
func Get_Control_Plus_plusArray() gopurs_runtime.Value {
	once_Control_Plus_plusArray.Do(func() {
		cache_Control_Plus_plusArray = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Alt_altArray()
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
	})
	return cache_Control_Plus_plusArray
}

var cache_Control_Plus_empty gopurs_runtime.Value
var once_Control_Plus_empty sync.Once
func Get_Control_Plus_empty() gopurs_runtime.Value {
	once_Control_Plus_empty.Do(func() {
		cache_Control_Plus_empty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Plus_empty(dict_0_box)
})
	})
	return cache_Control_Plus_empty
}

var cache_Control_Plus_empty__932402776 gopurs_runtime.Value
var once_Control_Plus_empty__932402776 sync.Once
func Get_Control_Plus_empty__932402776() gopurs_runtime.Value {
	once_Control_Plus_empty__932402776.Do(func() {
		cache_Control_Plus_empty__932402776 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Plus_empty__932402776(dict_0_box)
})
	})
	return cache_Control_Plus_empty__932402776
}

var cache_Control_Plus_plusArray__4260531026 gopurs_runtime.Value
var once_Control_Plus_plusArray__4260531026 sync.Once
func Get_Control_Plus_plusArray__4260531026() gopurs_runtime.Value {
	once_Control_Plus_plusArray__4260531026.Do(func() {
		cache_Control_Plus_plusArray__4260531026 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Alt_altArray()
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
	})
	return cache_Control_Plus_plusArray__4260531026
}

type Constructor_Control_Plus_Plus struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3709470893] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Plus_Plus)(ptr)
		_ = c
		switch key {
		case "Alt0": return gopurs_runtime.Box(c.V0)
		case "empty": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Plus_Plus: " + key)
		}
	}
}


func Call_Control_Plus_Plus_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Plus_empty(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "empty")
}

func Call_Control_Plus_empty__932402776(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "empty")
}


