package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_TC_MyClass_dollar_Dict gopurs_runtime.Value
var once_Test_TC_MyClass_dollar_Dict sync.Once
func Get_Test_TC_MyClass_dollar_Dict() gopurs_runtime.Value {
	once_Test_TC_MyClass_dollar_Dict.Do(func() {
		cache_Test_TC_MyClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1432427153, UnsafePtr: unsafe.Pointer(Call_Test_TC_MyClass_dollar_Dict(func() struct{
	myMethod gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	myMethod gopurs_runtime.Value
}{}
					clone.myMethod = gopurs_runtime.RecordGet(orig, "myMethod")
					return clone
				}()))}
})
	})
	return cache_Test_TC_MyClass_dollar_Dict
}

var cache_Test_TC_myClassInt gopurs_runtime.Value
var once_Test_TC_myClassInt sync.Once
func Get_Test_TC_myClassInt() gopurs_runtime.Value {
	once_Test_TC_myClassInt.Do(func() {
		cache_Test_TC_myClassInt = gopurs_runtime.Value{Type: 9, IntVal: 1432427153, UnsafePtr: unsafe.Pointer(Rebox_Test_TC_173323230_3370946277((&Constructor_Test_TC_MyClass[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("Int")
})})))}
	})
	return cache_Test_TC_myClassInt
}

var cache_Test_TC_myMethod gopurs_runtime.Value
var once_Test_TC_myMethod sync.Once
func Get_Test_TC_myMethod() gopurs_runtime.Value {
	once_Test_TC_myMethod.Do(func() {
		cache_Test_TC_myMethod = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_TC_myMethod(gopurs_runtime.CoerceToStruct[Constructor_Test_TC_MyClass[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Test_TC_myMethod
}

var cache_Test_TC_myMethod__2808757794 gopurs_runtime.Value
var once_Test_TC_myMethod__2808757794 sync.Once
func Get_Test_TC_myMethod__2808757794() gopurs_runtime.Value {
	once_Test_TC_myMethod__2808757794.Do(func() {
		cache_Test_TC_myMethod__2808757794 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_TC_myMethod__2808757794(__eta_norm_0_0_box)
})
	})
	return cache_Test_TC_myMethod__2808757794
}

var cache_Test_TC_describe gopurs_runtime.Value
var once_Test_TC_describe sync.Once
func Get_Test_TC_describe() gopurs_runtime.Value {
	once_Test_TC_describe.Do(func() {
		cache_Test_TC_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("TC"))
	})
	return cache_Test_TC_describe
}

var cache_Test_TC_act gopurs_runtime.Value
var once_Test_TC_act sync.Once
func Get_Test_TC_act() gopurs_runtime.Value {
	once_Test_TC_act.Do(func() {
		cache_Test_TC_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("Int")
})
	})
	return cache_Test_TC_act
}

type Constructor_Test_TC_MyClass[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1432427153] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Test_TC_MyClass[any])(ptr)
		_ = c
		switch key {
		case "myMethod": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Test_TC_MyClass: " + key)
		}
	}
}


func Call_Test_TC_MyClass_dollar_Dict(x_0_loop struct{
	myMethod gopurs_runtime.Value
}) *Constructor_Test_TC_MyClass[gopurs_runtime.Value] {
var x_0 struct{
	myMethod gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Test_TC_MyClass[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("myMethod", orig.myMethod)
				}())
}

func Call_Test_TC_myMethod(dict_0_loop *Constructor_Test_TC_MyClass[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Test_TC_MyClass[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Test_TC_myMethod__2808757794(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
myMethod__2808757794:
for {
if false { continue myMethod__2808757794 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Str("Int")
}
}

func Rebox_Test_TC_173323230_3370946277(in *Constructor_Test_TC_MyClass[int64]) *Constructor_Test_TC_MyClass[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Test_TC_MyClass[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}


