package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Class_MonadEffect_dollar_Dict gopurs_runtime.Value
var once_Effect_Class_MonadEffect_dollar_Dict sync.Once
func Get_Effect_Class_MonadEffect_dollar_Dict() gopurs_runtime.Value {
	once_Effect_Class_MonadEffect_dollar_Dict.Do(func() {
		cache_Effect_Class_MonadEffect_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(Call_Effect_Class_MonadEffect_dollar_Dict(func() struct{
	Monad0 gopurs_runtime.Value
	liftEffect gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Monad0 gopurs_runtime.Value
	liftEffect gopurs_runtime.Value
}{}
					clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
					clone.liftEffect = gopurs_runtime.RecordGet(orig, "liftEffect")
					return clone
				}()))}
})
	})
	return cache_Effect_Class_MonadEffect_dollar_Dict
}

var cache_Effect_Class_monadEffectEffect gopurs_runtime.Value
var once_Effect_Class_monadEffectEffect sync.Once
func Get_Effect_Class_monadEffectEffect() gopurs_runtime.Value {
	once_Effect_Class_monadEffectEffect.Do(func() {
		cache_Effect_Class_monadEffectEffect = gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)}))}
	})
	return cache_Effect_Class_monadEffectEffect
}

var cache_Effect_Class_liftEffect gopurs_runtime.Value
var once_Effect_Class_liftEffect sync.Once
func Get_Effect_Class_liftEffect() gopurs_runtime.Value {
	once_Effect_Class_liftEffect.Do(func() {
		cache_Effect_Class_liftEffect = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect
}

var cache_Effect_Class_liftEffect__2047119308 gopurs_runtime.Value
var once_Effect_Class_liftEffect__2047119308 sync.Once
func Get_Effect_Class_liftEffect__2047119308() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__2047119308.Do(func() {
		cache_Effect_Class_liftEffect__2047119308 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__2047119308(__eta_norm_0_0_box)
})
	})
	return cache_Effect_Class_liftEffect__2047119308
}

var cache_Effect_Class_liftEffect__2912063852 gopurs_runtime.Value
var once_Effect_Class_liftEffect__2912063852 sync.Once
func Get_Effect_Class_liftEffect__2912063852() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__2912063852.Do(func() {
		cache_Effect_Class_liftEffect__2912063852 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__2912063852(__eta_norm_0_0_box)
})
	})
	return cache_Effect_Class_liftEffect__2912063852
}

type Constructor_Effect_Class_MonadEffect[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2217729261] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Effect_Class_MonadEffect[any])(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "liftEffect": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Effect_Class_MonadEffect: " + key)
		}
	}
}


func Call_Effect_Class_MonadEffect_dollar_Dict(x_0_loop struct{
	Monad0 gopurs_runtime.Value
	liftEffect gopurs_runtime.Value
}) *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value] {
var x_0 struct{
	Monad0 gopurs_runtime.Value
	liftEffect gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Monad0", "liftEffect", orig.Monad0, orig.liftEffect)
				}())
}

func Call_Effect_Class_liftEffect(dict_0_loop *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__2047119308(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
liftEffect__2047119308:
for {
if false { continue liftEffect__2047119308 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), __eta_norm_0_0)
}
}

func Call_Effect_Class_liftEffect__2912063852(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
liftEffect__2912063852:
for {
if false { continue liftEffect__2912063852 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), __eta_norm_0_0)
}
}


