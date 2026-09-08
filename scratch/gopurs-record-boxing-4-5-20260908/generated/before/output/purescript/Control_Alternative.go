package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Alternative_Alternative_dollar_Dict gopurs_runtime.Value
var once_Control_Alternative_Alternative_dollar_Dict sync.Once
func Get_Control_Alternative_Alternative_dollar_Dict() gopurs_runtime.Value {
	once_Control_Alternative_Alternative_dollar_Dict.Do(func() {
		cache_Control_Alternative_Alternative_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Call_Control_Alternative_Alternative_dollar_Dict(func() struct{
	Applicative0 gopurs_runtime.Value
	Plus1 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Applicative0 gopurs_runtime.Value
	Plus1 gopurs_runtime.Value
}{}
					clone.Applicative0 = gopurs_runtime.RecordGet(orig, "Applicative0")
					clone.Plus1 = gopurs_runtime.RecordGet(orig, "Plus1")
					return clone
				}()))}
})
	})
	return cache_Control_Alternative_Alternative_dollar_Dict
}

var cache_Control_Alternative_guard gopurs_runtime.Value
var once_Control_Alternative_guard sync.Once
func Get_Control_Alternative_guard() gopurs_runtime.Value {
	once_Control_Alternative_guard.Do(func() {
		cache_Control_Alternative_guard = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Alternative_guard(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_Control_Alternative_guard
}

var cache_Control_Alternative_guard__2562009031 gopurs_runtime.Value
var once_Control_Alternative_guard__2562009031 sync.Once
func Get_Control_Alternative_guard__2562009031() gopurs_runtime.Value {
	once_Control_Alternative_guard__2562009031.Do(func() {
		cache_Control_Alternative_guard__2562009031 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Alternative_guard__2562009031(__eta_norm_0_0_box)
})
	})
	return cache_Control_Alternative_guard__2562009031
}

var cache_Control_Alternative_guard__1365085559 gopurs_runtime.Value
var once_Control_Alternative_guard__1365085559 sync.Once
func Get_Control_Alternative_guard__1365085559() gopurs_runtime.Value {
	once_Control_Alternative_guard__1365085559.Do(func() {
		cache_Control_Alternative_guard__1365085559 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Alternative_guard__1365085559(__eta_norm_0_0_box)
})
	})
	return cache_Control_Alternative_guard__1365085559
}

var cache_Control_Alternative_alternativeArray gopurs_runtime.Value
var once_Control_Alternative_alternativeArray sync.Once
func Get_Control_Alternative_alternativeArray() gopurs_runtime.Value {
	once_Control_Alternative_alternativeArray.Do(func() {
		cache_Control_Alternative_alternativeArray = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Control_Plus_plusArray()))}
})}))}
	})
	return cache_Control_Alternative_alternativeArray
}

type Constructor_Control_Alternative_Alternative[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[397869517] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Alternative_Alternative[any])(ptr)
		_ = c
		switch key {
		case "Applicative0": return gopurs_runtime.Box(c.V0)
		case "Plus1": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Alternative_Alternative: " + key)
		}
	}
}


func Call_Control_Alternative_Alternative_dollar_Dict(x_0_loop struct{
	Applicative0 gopurs_runtime.Value
	Plus1 gopurs_runtime.Value
}) *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] {
var x_0 struct{
	Applicative0 gopurs_runtime.Value
	Plus1 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Applicative0", "Plus1"}, []gopurs_runtime.Value{orig.Applicative0, orig.Plus1})
				}())
}

func Call_Control_Alternative_guard(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): empty_2_1 shape=Other bindingType=(TypeApp (TypeVar m) [Unit])
empty_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "empty")
_ = empty_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
goto end_branch_2
} else {

}
}
{
__t2 = empty_2_1
}
end_branch_2:
return __t2
})
}

func Call_Control_Alternative_guard__2562009031(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
guard__2562009031:
for {
if false { continue guard__2562009031 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__eta_norm_0_0.IntVal) != (0) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{Get_Data_Unit_unit(), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}
}

func Call_Control_Alternative_guard__1365085559(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
guard__1365085559:
for {
if false { continue guard__1365085559 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0.IntVal) != (0) {
__t0 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Unit_unit()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}
end_branch_0:
return __t0
}
}


