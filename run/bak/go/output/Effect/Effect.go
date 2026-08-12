package Effect

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_monadEffect_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Any gopurs_runtime.Value
var once_monadEffect_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Any sync.Once
func Get_monadEffect_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Any() gopurs_runtime.Value {
	once_monadEffect_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Any.Do(func() {
		cache_monadEffect_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Any = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEffect()
}))
	})
	return cache_monadEffect_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any_Any
}

var cache_monadEffect gopurs_runtime.Value
var once_monadEffect sync.Once
func Get_monadEffect() gopurs_runtime.Value {
	once_monadEffect.Do(func() {
		cache_monadEffect = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEffect()
}))
	})
	return cache_monadEffect
}

var cache_monadEffect__gopurs_runtime_Value_3527935219 gopurs_runtime.Value
var once_monadEffect__gopurs_runtime_Value_3527935219 sync.Once
func Get_monadEffect__gopurs_runtime_Value_3527935219() gopurs_runtime.Value {
	once_monadEffect__gopurs_runtime_Value_3527935219.Do(func() {
		cache_monadEffect__gopurs_runtime_Value_3527935219 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEffect()
}))
	})
	return cache_monadEffect__gopurs_runtime_Value_3527935219
}

var cache_bindEffect_ADT_Control_Bind_Bind_ADT_Effect_Effect gopurs_runtime.Value
var once_bindEffect_ADT_Control_Bind_Bind_ADT_Effect_Effect sync.Once
func Get_bindEffect_ADT_Control_Bind_Bind_ADT_Effect_Effect() gopurs_runtime.Value {
	once_bindEffect_ADT_Control_Bind_Bind_ADT_Effect_Effect.Do(func() {
		cache_bindEffect_ADT_Control_Bind_Bind_ADT_Effect_Effect = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_bindE())
	})
	return cache_bindEffect_ADT_Control_Bind_Bind_ADT_Effect_Effect
}

var cache_bindEffect_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any gopurs_runtime.Value
var once_bindEffect_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any sync.Once
func Get_bindEffect_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any() gopurs_runtime.Value {
	once_bindEffect_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any.Do(func() {
		cache_bindEffect_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_bindE())
	})
	return cache_bindEffect_Record_Row_bind_ForAll_a_b_Func_ADT_Effect_Effect_Any_Func_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any
}

var cache_bindEffect gopurs_runtime.Value
var once_bindEffect sync.Once
func Get_bindEffect() gopurs_runtime.Value {
	once_bindEffect.Do(func() {
		cache_bindEffect = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_bindE())
	})
	return cache_bindEffect
}

var cache_bindEffect__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2113658466 gopurs_runtime.Value
var once_bindEffect__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2113658466 sync.Once
func Get_bindEffect__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2113658466() gopurs_runtime.Value {
	once_bindEffect__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2113658466.Do(func() {
		cache_bindEffect__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2113658466 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_bindE()})}
	})
	return cache_bindEffect__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2113658466
}

var cache_bindEffect__gopurs_runtime_Value_3856311079 gopurs_runtime.Value
var once_bindEffect__gopurs_runtime_Value_3856311079 sync.Once
func Get_bindEffect__gopurs_runtime_Value_3856311079() gopurs_runtime.Value {
	once_bindEffect__gopurs_runtime_Value_3856311079.Do(func() {
		cache_bindEffect__gopurs_runtime_Value_3856311079 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_bindE())
	})
	return cache_bindEffect__gopurs_runtime_Value_3856311079
}

var cache_applyEffect_ADT_Control_Apply_Apply_ADT_Effect_Effect gopurs_runtime.Value
var once_applyEffect_ADT_Control_Apply_Apply_ADT_Effect_Effect sync.Once
func Get_applyEffect_ADT_Control_Apply_Apply_ADT_Effect_Effect() gopurs_runtime.Value {
	once_applyEffect_ADT_Control_Apply_Apply_ADT_Effect_Effect.Do(func() {
		cache_applyEffect_ADT_Control_Apply_Apply_ADT_Effect_Effect = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect_ADT_Control_Apply_Apply_ADT_Effect_Effect
}

var cache_applyEffect_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any gopurs_runtime.Value
var once_applyEffect_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any sync.Once
func Get_applyEffect_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any() gopurs_runtime.Value {
	once_applyEffect_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any.Do(func() {
		cache_applyEffect_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any
}

var cache_applyEffect gopurs_runtime.Value
var once_applyEffect sync.Once
func Get_applyEffect() gopurs_runtime.Value {
	once_applyEffect.Do(func() {
		cache_applyEffect = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect
}

var cache_applyEffect__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__1723132130 gopurs_runtime.Value
var once_applyEffect__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__1723132130 sync.Once
func Get_applyEffect__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__1723132130() gopurs_runtime.Value {
	once_applyEffect__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__1723132130.Do(func() {
		cache_applyEffect__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__1723132130 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
})})}
}()
	})
	return cache_applyEffect__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__1723132130
}

var cache_applyEffect__gopurs_runtime_Value_2014400020 gopurs_runtime.Value
var once_applyEffect__gopurs_runtime_Value_2014400020 sync.Once
func Get_applyEffect__gopurs_runtime_Value_2014400020() gopurs_runtime.Value {
	once_applyEffect__gopurs_runtime_Value_2014400020.Do(func() {
		cache_applyEffect__gopurs_runtime_Value_2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__gopurs_runtime_Value_2014400020
}

var cache_applicativeEffect_ADT_Control_Applicative_Applicative_ADT_Effect_Effect gopurs_runtime.Value
var once_applicativeEffect_ADT_Control_Applicative_Applicative_ADT_Effect_Effect sync.Once
func Get_applicativeEffect_ADT_Control_Applicative_Applicative_ADT_Effect_Effect() gopurs_runtime.Value {
	once_applicativeEffect_ADT_Control_Applicative_Applicative_ADT_Effect_Effect.Do(func() {
		cache_applicativeEffect_ADT_Control_Applicative_Applicative_ADT_Effect_Effect = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_pureE())
	})
	return cache_applicativeEffect_ADT_Control_Applicative_Applicative_ADT_Effect_Effect
}

var cache_applicativeEffect_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any gopurs_runtime.Value
var once_applicativeEffect_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any sync.Once
func Get_applicativeEffect_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any() gopurs_runtime.Value {
	once_applicativeEffect_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any.Do(func() {
		cache_applicativeEffect_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_pureE())
	})
	return cache_applicativeEffect_Record_Row_pure_ForAll_a_Func_Any_ADT_Effect_Effect_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Effect_Effect_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any_Any_Any
}

var cache_applicativeEffect gopurs_runtime.Value
var once_applicativeEffect sync.Once
func Get_applicativeEffect() gopurs_runtime.Value {
	once_applicativeEffect.Do(func() {
		cache_applicativeEffect = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_pureE())
	})
	return cache_applicativeEffect
}

var cache_applicativeEffect__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__284161122 gopurs_runtime.Value
var once_applicativeEffect__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__284161122 sync.Once
func Get_applicativeEffect__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__284161122() gopurs_runtime.Value {
	once_applicativeEffect__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__284161122.Do(func() {
		cache_applicativeEffect__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__284161122 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_pureE()})}
	})
	return cache_applicativeEffect__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__284161122
}

var cache_applicativeEffect__gopurs_runtime_Value_1969567048 gopurs_runtime.Value
var once_applicativeEffect__gopurs_runtime_Value_1969567048 sync.Once
func Get_applicativeEffect__gopurs_runtime_Value_1969567048() gopurs_runtime.Value {
	once_applicativeEffect__gopurs_runtime_Value_1969567048.Do(func() {
		cache_applicativeEffect__gopurs_runtime_Value_1969567048 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_pureE())
	})
	return cache_applicativeEffect__gopurs_runtime_Value_1969567048
}

var cache_functorEffect_ADT_Data_Functor_Functor_ADT_Effect_Effect gopurs_runtime.Value
var once_functorEffect_ADT_Data_Functor_Functor_ADT_Effect_Effect sync.Once
func Get_functorEffect_ADT_Data_Functor_Functor_ADT_Effect_Effect() gopurs_runtime.Value {
	once_functorEffect_ADT_Data_Functor_Functor_ADT_Effect_Effect.Do(func() {
		cache_functorEffect_ADT_Data_Functor_Functor_ADT_Effect_Effect = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect_ADT_Data_Functor_Functor_ADT_Effect_Effect
}

var cache_functorEffect_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any gopurs_runtime.Value
var once_functorEffect_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any sync.Once
func Get_functorEffect_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any() gopurs_runtime.Value {
	once_functorEffect_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any.Do(func() {
		cache_functorEffect_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Effect_Effect_Any_ADT_Effect_Effect_Any_Any
}

var cache_functorEffect gopurs_runtime.Value
var once_functorEffect sync.Once
func Get_functorEffect() gopurs_runtime.Value {
	once_functorEffect.Do(func() {
		cache_functorEffect = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect
}

var cache_functorEffect__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__347161653 gopurs_runtime.Value
var once_functorEffect__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__347161653 sync.Once
func Get_functorEffect__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__347161653() gopurs_runtime.Value {
	once_functorEffect__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__347161653.Do(func() {
		cache_functorEffect__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__347161653 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), f_1), a_2)
})
})})}
}()
	})
	return cache_functorEffect__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__347161653
}

var cache_functorEffect__gopurs_runtime_Value_3107547953 gopurs_runtime.Value
var once_functorEffect__gopurs_runtime_Value_3107547953 sync.Once
func Get_functorEffect__gopurs_runtime_Value_3107547953() gopurs_runtime.Value {
	once_functorEffect__gopurs_runtime_Value_3107547953.Do(func() {
		cache_functorEffect__gopurs_runtime_Value_3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__gopurs_runtime_Value_3107547953
}

var cache_semigroupEffect gopurs_runtime.Value
var once_semigroupEffect sync.Once
func Get_semigroupEffect() gopurs_runtime.Value {
	once_semigroupEffect.Do(func() {
		cache_semigroupEffect = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupEffect(dictSemigroup_0_box)
})
	})
	return cache_semigroupEffect
}

var cache_monoidEffect gopurs_runtime.Value
var once_monoidEffect sync.Once
func Get_monoidEffect() gopurs_runtime.Value {
	once_monoidEffect.Do(func() {
		cache_monoidEffect = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidEffect(dictMonoid_0_box)
})
	})
	return cache_monoidEffect
}

func Call_semigroupEffect(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyEffect(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_5_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_2
Applicative0_6_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_3
return gopurs_runtime.Apply2(Bind1_5_2.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, __local_var_2_1, a_3), gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_5_2.V1, b_4, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_6_3.V1, gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_monoidEffect(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
semigroupEffect1_1_0 := Call_semigroupEffect(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupEffect1_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffect1_1_0
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
}

func Get_bindE() gopurs_runtime.Value {
	return _Gopurs_BindE
}

func Get_forE() gopurs_runtime.Value {
	return _Gopurs_ForE
}

func Get_foreachE() gopurs_runtime.Value {
	return _Gopurs_ForeachE
}

func Get_pureE() gopurs_runtime.Value {
	return _Gopurs_PureE
}

func Get_untilE() gopurs_runtime.Value {
	return _Gopurs_UntilE
}

func Get_whileE() gopurs_runtime.Value {
	return _Gopurs_WhileE
}
