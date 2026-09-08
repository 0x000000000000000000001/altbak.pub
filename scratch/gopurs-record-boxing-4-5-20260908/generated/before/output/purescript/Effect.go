package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_monadEffect gopurs_runtime.Value
var once_Effect_monadEffect sync.Once
func Get_Effect_monadEffect() gopurs_runtime.Value {
	once_Effect_monadEffect.Do(func() {
		cache_Effect_monadEffect = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()))}
})}))}
	})
	return cache_Effect_monadEffect
}

var cache_Effect_bindEffect gopurs_runtime.Value
var once_Effect_bindEffect sync.Once
func Get_Effect_bindEffect() gopurs_runtime.Value {
	once_Effect_bindEffect.Do(func() {
		cache_Effect_bindEffect = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()))}
}), Get_Effect_bindE()}))}
	})
	return cache_Effect_bindEffect
}

var cache_Effect_applyEffect gopurs_runtime.Value
var once_Effect_applyEffect sync.Once
func Get_Effect_applyEffect() gopurs_runtime.Value {
	once_Effect_applyEffect.Do(func() {
		cache_Effect_applyEffect = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()).V1), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()).V0), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()))}
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime__4, a_prime__5))
}))
}))
})}))}
}()
	})
	return cache_Effect_applyEffect
}

var cache_Effect_applicativeEffect gopurs_runtime.Value
var once_Effect_applicativeEffect sync.Once
func Get_Effect_applicativeEffect() gopurs_runtime.Value {
	once_Effect_applicativeEffect.Do(func() {
		cache_Effect_applicativeEffect = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()))}
}), Get_Effect_pureE()}))}
	})
	return cache_Effect_applicativeEffect
}

var cache_Effect_functorEffect gopurs_runtime.Value
var once_Effect_functorEffect sync.Once
func Get_Effect_functorEffect() gopurs_runtime.Value {
	once_Effect_functorEffect.Do(func() {
		cache_Effect_functorEffect = func() gopurs_runtime.Value {
// TAST (Let): Apply0_0_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V0), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_0_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1), f_1), a_2)
})}))}
}()
	})
	return cache_Effect_functorEffect
}

var cache_Effect_semigroupEffect gopurs_runtime.Value
var once_Effect_semigroupEffect sync.Once
func Get_Effect_semigroupEffect() gopurs_runtime.Value {
	once_Effect_semigroupEffect.Do(func() {
		cache_Effect_semigroupEffect = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_semigroupEffect(dictSemigroup_0_box)
})
	})
	return cache_Effect_semigroupEffect
}

var cache_Effect_monoidEffect gopurs_runtime.Value
var once_Effect_monoidEffect sync.Once
func Get_Effect_monoidEffect() gopurs_runtime.Value {
	once_Effect_monoidEffect.Do(func() {
		cache_Effect_monoidEffect = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_monoidEffect(dictMonoid_0_box)
})
	})
	return cache_Effect_monoidEffect
}

func Call_Effect_semigroupEffect(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): __local_var_1_0 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
f_prime__4_2 := __local_var_1_0
_ = f_prime__4_2
a_prime__5_3 := gopurs_runtime.Apply(a_2, gopurs_runtime.Value{})
_ = a_prime__5_3
f_prime__4_1 := gopurs_runtime.Apply(f_prime__4_2, a_prime__5_3)
_ = f_prime__4_1
a_prime__5_4 := gopurs_runtime.Apply(b_3, gopurs_runtime.Value{})
_ = a_prime__5_4
return gopurs_runtime.Apply(f_prime__4_1, a_prime__5_4)
})
})}))}
}

func Call_Effect_monoidEffect(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_1_1
// TAST (Let): semigroupEffect1_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Effect","Effect"] [(TypeVar a)])])
semigroupEffect1_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
f_prime__4_3 := __local_var_1_1
_ = f_prime__4_3
a_prime__5_4 := gopurs_runtime.Apply(a_2, gopurs_runtime.Value{})
_ = a_prime__5_4
f_prime__4_2 := gopurs_runtime.Apply(f_prime__4_3, a_prime__5_4)
_ = f_prime__4_2
a_prime__5_5 := gopurs_runtime.Apply(b_3, gopurs_runtime.Value{})
_ = a_prime__5_5
return gopurs_runtime.Apply(f_prime__4_2, a_prime__5_5)
})
})})
_ = semigroupEffect1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEffect1_1_0)}
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_6 shape=Other bindingType=Any
__local_var_2_6 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_6
return __local_var_2_6
})}))}
}

func Get_Effect_bindE() gopurs_runtime.Value {
	return _Gopurs_Effect_BindE
}

func Get_Effect_forE() gopurs_runtime.Value {
	return _Gopurs_Effect_ForE
}

func Get_Effect_foreachE() gopurs_runtime.Value {
	return _Gopurs_Effect_ForeachE
}

func Get_Effect_pureE() gopurs_runtime.Value {
	return _Gopurs_Effect_PureE
}

func Get_Effect_untilE() gopurs_runtime.Value {
	return _Gopurs_Effect_UntilE
}

func Get_Effect_whileE() gopurs_runtime.Value {
	return _Gopurs_Effect_WhileE
}
