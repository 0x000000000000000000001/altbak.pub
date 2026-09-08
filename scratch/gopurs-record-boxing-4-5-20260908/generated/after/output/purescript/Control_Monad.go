package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Monad_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Monad_dollar_Dict sync.Once
func Get_Control_Monad_Monad_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Monad_dollar_Dict.Do(func() {
		cache_Control_Monad_Monad_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Monad_dollar_Dict(func() struct{
	Applicative0 gopurs_runtime.Value
	Bind1 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Applicative0 gopurs_runtime.Value
	Bind1 gopurs_runtime.Value
}{}
					clone.Applicative0 = gopurs_runtime.RecordGet(orig, "Applicative0")
					clone.Bind1 = gopurs_runtime.RecordGet(orig, "Bind1")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Monad_dollar_Dict
}

var cache_Control_Monad_whenM gopurs_runtime.Value
var once_Control_Monad_whenM sync.Once
func Get_Control_Monad_whenM() gopurs_runtime.Value {
	once_Control_Monad_whenM.Do(func() {
		cache_Control_Monad_whenM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_whenM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_whenM
}

var cache_Control_Monad_unlessM gopurs_runtime.Value
var once_Control_Monad_unlessM sync.Once
func Get_Control_Monad_unlessM() gopurs_runtime.Value {
	once_Control_Monad_unlessM.Do(func() {
		cache_Control_Monad_unlessM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_unlessM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_unlessM
}

var cache_Control_Monad_monadProxy gopurs_runtime.Value
var once_Control_Monad_monadProxy sync.Once
func Get_Control_Monad_monadProxy() gopurs_runtime.Value {
	once_Control_Monad_monadProxy.Do(func() {
		cache_Control_Monad_monadProxy = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_1564826004_2568689657((&Constructor_Control_Monad_Monad[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_3924528404_1439734649(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[uint32]](Get_Control_Applicative_applicativeProxy())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_4014717844_2748095225(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[uint32]](Get_Control_Bind_bindProxy())))}
})})))}
	})
	return cache_Control_Monad_monadProxy
}

var cache_Control_Monad_monadFn gopurs_runtime.Value
var once_Control_Monad_monadFn sync.Once
func Get_Control_Monad_monadFn() gopurs_runtime.Value {
	once_Control_Monad_monadFn.Do(func() {
		cache_Control_Monad_monadFn = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeFn()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Bind_bindFn()))}
})}))}
	})
	return cache_Control_Monad_monadFn
}

var cache_Control_Monad_monadArray gopurs_runtime.Value
var once_Control_Monad_monadArray sync.Once
func Get_Control_Monad_monadArray() gopurs_runtime.Value {
	once_Control_Monad_monadArray.Do(func() {
		cache_Control_Monad_monadArray = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Bind_bindArray()))}
})}))}
	})
	return cache_Control_Monad_monadArray
}

var cache_Control_Monad_liftM1 gopurs_runtime.Value
var once_Control_Monad_liftM1 sync.Once
func Get_Control_Monad_liftM1() gopurs_runtime.Value {
	once_Control_Monad_liftM1.Do(func() {
		cache_Control_Monad_liftM1 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_liftM1(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_liftM1
}

var cache_Control_Monad_ap gopurs_runtime.Value
var once_Control_Monad_ap sync.Once
func Get_Control_Monad_ap() gopurs_runtime.Value {
	once_Control_Monad_ap.Do(func() {
		cache_Control_Monad_ap = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_ap
}

type Constructor_Control_Monad_Monad[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[778916621] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Monad[any])(ptr)
		_ = c
		switch key {
		case "Applicative0": return gopurs_runtime.Box(c.V0)
		case "Bind1": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Monad: " + key)
		}
	}
}


func Call_Control_Monad_Monad_dollar_Dict(x_0_loop struct{
	Applicative0 gopurs_runtime.Value
	Bind1 gopurs_runtime.Value
}) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
var x_0 struct{
	Applicative0 gopurs_runtime.Value
	Bind1 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Applicative0", "Bind1", orig.Applicative0, orig.Bind1)
				}())
}

func Call_Control_Monad_whenM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func2(func(mb_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), mb_3, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (b_5.IntVal) != (0) {
__t2 = m_4
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), Get_Data_Unit_unit())
}
end_branch_2:
return __t2
}))
})
}

func Call_Control_Monad_unlessM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func2(func(mb_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), mb_3, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if ((b_5.IntVal) != (0)) != (true) {
__t2 = m_4
goto end_branch_2
} else {

}
}
{
if (b_5.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), Get_Data_Unit_unit())
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
}

func Call_Control_Monad_liftM1(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), a_4, gopurs_runtime.Func(func(a_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Apply(f_3, a_prime__5))
}))
})
}

func Call_Control_Monad_ap(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), f_3, gopurs_runtime.Func(func(f_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), a_4, gopurs_runtime.Func(func(a_prime__6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Apply(f_prime__5, a_prime__6))
}))
}))
})
}

func Rebox_Control_Monad_1564826004_2568689657(in *Constructor_Control_Monad_Monad[uint32]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_3924528404_1439734649(in *Constructor_Control_Applicative_Applicative[uint32]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_4014717844_2748095225(in *Constructor_Control_Bind_Bind[uint32]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


