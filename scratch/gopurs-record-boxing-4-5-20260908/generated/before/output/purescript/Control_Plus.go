package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Plus_Plus_dollar_Dict gopurs_runtime.Value
var once_Control_Plus_Plus_dollar_Dict sync.Once
func Get_Control_Plus_Plus_dollar_Dict() gopurs_runtime.Value {
	once_Control_Plus_Plus_dollar_Dict.Do(func() {
		cache_Control_Plus_Plus_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Call_Control_Plus_Plus_dollar_Dict(func() struct{
	Alt0 gopurs_runtime.Value
	empty gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Alt0 gopurs_runtime.Value
	empty gopurs_runtime.Value
}{}
					clone.Alt0 = gopurs_runtime.RecordGet(orig, "Alt0")
					clone.empty = gopurs_runtime.RecordGet(orig, "empty")
					return clone
				}()))}
})
	})
	return cache_Control_Plus_Plus_dollar_Dict
}

var cache_Control_Plus_plusArray gopurs_runtime.Value
var once_Control_Plus_plusArray sync.Once
func Get_Control_Plus_plusArray() gopurs_runtime.Value {
	once_Control_Plus_plusArray.Do(func() {
		cache_Control_Plus_plusArray = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Control_Alt_altArray()))}
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}))}
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

var cache_Control_Plus_empty__1269207005 gopurs_runtime.Value
var once_Control_Plus_empty__1269207005 sync.Once
func Get_Control_Plus_empty__1269207005() gopurs_runtime.Value {
	once_Control_Plus_empty__1269207005.Do(func() {
		cache_Control_Plus_empty__1269207005 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Control_Plus_empty__1269207005
}

type Constructor_Control_Plus_Plus[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3709470893] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Plus_Plus[any])(ptr)
		_ = c
		switch key {
		case "Alt0": return gopurs_runtime.Box(c.V0)
		case "empty": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Plus_Plus: " + key)
		}
	}
}


func Call_Control_Plus_Plus_dollar_Dict(x_0_loop struct{
	Alt0 gopurs_runtime.Value
	empty gopurs_runtime.Value
}) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
var x_0 struct{
	Alt0 gopurs_runtime.Value
	empty gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Alt0", "empty"}, []gopurs_runtime.Value{orig.Alt0, orig.empty})
				}())
}

func Call_Control_Plus_empty(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "empty")
}


