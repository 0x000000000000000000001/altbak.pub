package Control_Comonad_Store_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var StoreT gopurs_runtime.Value
var once_StoreT sync.Once
func Get_StoreT() gopurs_runtime.Value {
	once_StoreT.Do(func() {
		StoreT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return StoreT
}

var runStoreT gopurs_runtime.Value
var once_runStoreT sync.Once
func Get_runStoreT() gopurs_runtime.Value {
	once_runStoreT.Do(func() {
		runStoreT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runStoreT
}

var newtypeStoreT gopurs_runtime.Value
var once_newtypeStoreT sync.Once
func Get_newtypeStoreT() gopurs_runtime.Value {
	once_newtypeStoreT.Do(func() {
		newtypeStoreT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeStoreT
}

var functorStoreT gopurs_runtime.Value
var once_functorStoreT sync.Once
func Get_functorStoreT() gopurs_runtime.Value {
	once_functorStoreT.Do(func() {
		functorStoreT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func2(func(h_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(h_3, x_4))
}), gopurs_runtime.ConstructorGet(v_2, 0)), gopurs_runtime.ConstructorGet(v_2, 1))
}))
})
	})
	return functorStoreT
}

var extendStoreT gopurs_runtime.Value
var once_extendStoreT sync.Once
func Get_extendStoreT() gopurs_runtime.Value {
	once_extendStoreT.Do(func() {
		extendStoreT = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorStoreT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func2(func(h_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(h_4, x_5))
}), gopurs_runtime.ConstructorGet(v_3, 0)), gopurs_runtime.ConstructorGet(v_3, 1))
}))
_ = functorStoreT1_2_1
return gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func2(func(w_prime_5 gopurs_runtime.Value, s_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Constructor2("Tuple", w_prime_5, s_prime_6))
}), gopurs_runtime.ConstructorGet(v_4, 0)), gopurs_runtime.ConstructorGet(v_4, 1))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStoreT1_2_1
}))
})
	})
	return extendStoreT
}

var comonadTransStoreT gopurs_runtime.Value
var once_comonadTransStoreT sync.Once
func Get_comonadTransStoreT() gopurs_runtime.Value {
	once_comonadTransStoreT.Do(func() {
		comonadTransStoreT = gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func2(func(dictComonad_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.ConstructorGet(v_1, 1)
_ = __local_var_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, __local_var_2_0)
}), gopurs_runtime.ConstructorGet(v_1, 0))
}))
	})
	return comonadTransStoreT
}

var comonadStoreT gopurs_runtime.Value
var once_comonadStoreT sync.Once
func Get_comonadStoreT() gopurs_runtime.Value {
	once_comonadStoreT.Do(func() {
		comonadStoreT = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorStoreT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func2(func(h_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Apply(h_5, x_6))
}), gopurs_runtime.ConstructorGet(v_4, 0)), gopurs_runtime.ConstructorGet(v_4, 1))
}))
_ = functorStoreT1_3_3
extendStoreT1_3_2 := gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "extend"), gopurs_runtime.Func2(func(w_prime_6 gopurs_runtime.Value, s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Constructor2("Tuple", w_prime_6, s_prime_7))
}), gopurs_runtime.ConstructorGet(v_5, 0)), gopurs_runtime.ConstructorGet(v_5, 1))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStoreT1_3_3
}))
_ = extendStoreT1_3_2
return gopurs_runtime.RecordDict2("extract", "Extend0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), gopurs_runtime.ConstructorGet(v_4, 0), gopurs_runtime.ConstructorGet(v_4, 1))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendStoreT1_3_2
}))
})
	})
	return comonadStoreT
}


