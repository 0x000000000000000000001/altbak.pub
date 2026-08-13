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
		cache_Effect_monadEffect = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Effect_applicativeEffect()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Effect_bindEffect()))}
})})}
	})
	return cache_Effect_monadEffect
}

var cache_Effect_bindEffect gopurs_runtime.Value
var once_Effect_bindEffect sync.Once
func Get_Effect_bindEffect() gopurs_runtime.Value {
	once_Effect_bindEffect.Do(func() {
		cache_Effect_bindEffect = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Effect_applyEffect()))}
}), Get_Effect_bindE()})}
	})
	return cache_Effect_bindEffect
}

var cache_Effect_applyEffect gopurs_runtime.Value
var once_Effect_applyEffect sync.Once
func Get_Effect_applyEffect() gopurs_runtime.Value {
	once_Effect_applyEffect.Do(func() {
		cache_Effect_applyEffect = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Effect_functorEffect()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{})
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Effect_applicativeEffect()).V1), gopurs_runtime.Apply(__local_var_2_0, __local_var_3_1)), gopurs_runtime.Value{})
})
})
})})}
	})
	return cache_Effect_applyEffect
}

var cache_Effect_applicativeEffect gopurs_runtime.Value
var once_Effect_applicativeEffect sync.Once
func Get_Effect_applicativeEffect() gopurs_runtime.Value {
	once_Effect_applicativeEffect.Do(func() {
		cache_Effect_applicativeEffect = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Effect_applyEffect()))}
}), Get_Effect_pureE()})}
	})
	return cache_Effect_applicativeEffect
}

var cache_Effect_functorEffect gopurs_runtime.Value
var once_Effect_functorEffect sync.Once
func Get_Effect_functorEffect() gopurs_runtime.Value {
	once_Effect_functorEffect.Do(func() {
		cache_Effect_functorEffect = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := f_0
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Apply(__local_var_2_0, __local_var_3_1)
})
})
})})}
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
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := __local_var_1_0
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(a_2, gopurs_runtime.Value{})
_ = __local_var_5_3
__local_var_4_1 := gopurs_runtime.Apply(__local_var_4_2, __local_var_5_3)
_ = __local_var_4_1
__local_var_5_4 := gopurs_runtime.Apply(b_3, gopurs_runtime.Value{})
_ = __local_var_5_4
return gopurs_runtime.Apply(__local_var_4_1, __local_var_5_4)
})
})
})})}
}

func Call_Effect_monoidEffect(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_1_1
// TAST (Let): semigroupEffect1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupEffect1_1_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := __local_var_1_1
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply(a_2, gopurs_runtime.Value{})
_ = __local_var_5_4
__local_var_4_2 := gopurs_runtime.Apply(__local_var_4_3, __local_var_5_4)
_ = __local_var_4_2
__local_var_5_5 := gopurs_runtime.Apply(b_3, gopurs_runtime.Value{})
_ = __local_var_5_5
return gopurs_runtime.Apply(__local_var_4_2, __local_var_5_5)
})
})
})}
_ = semigroupEffect1_1_0
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_6
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEffect1_1_0)}
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_6
})})}
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
