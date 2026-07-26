package Control_Comonad_Store_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	unsafe "unsafe"
)

var cache_StoreT gopurs_runtime.Value
var once_StoreT sync.Once
func Get_StoreT() gopurs_runtime.Value {
	once_StoreT.Do(func() {
		cache_StoreT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_StoreT(x_0_box)
})
	})
	return cache_StoreT
}

var cache_runStoreT gopurs_runtime.Value
var once_runStoreT sync.Once
func Get_runStoreT() gopurs_runtime.Value {
	once_runStoreT.Do(func() {
		cache_runStoreT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runStoreT(v_0_box)
})
	})
	return cache_runStoreT
}

var cache_newtypeStoreT gopurs_runtime.Value
var once_newtypeStoreT sync.Once
func Get_newtypeStoreT() gopurs_runtime.Value {
	once_newtypeStoreT.Do(func() {
		cache_newtypeStoreT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeStoreT
}

var cache_functorStoreT gopurs_runtime.Value
var once_functorStoreT sync.Once
func Get_functorStoreT() gopurs_runtime.Value {
	once_functorStoreT.Do(func() {
		cache_functorStoreT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorStoreT(dictFunctor_0_box)
})
	})
	return cache_functorStoreT
}

var cache_extendStoreT gopurs_runtime.Value
var once_extendStoreT sync.Once
func Get_extendStoreT() gopurs_runtime.Value {
	once_extendStoreT.Do(func() {
		cache_extendStoreT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendStoreT(dictExtend_0_box)
})
	})
	return cache_extendStoreT
}

var cache_comonadTransStoreT gopurs_runtime.Value
var once_comonadTransStoreT sync.Once
func Get_comonadTransStoreT() gopurs_runtime.Value {
	once_comonadTransStoreT.Do(func() {
		cache_comonadTransStoreT = gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func2(func(dictComonad_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, __local_var_2_0)
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)
}))
	})
	return cache_comonadTransStoreT
}

var cache_comonadStoreT gopurs_runtime.Value
var once_comonadStoreT sync.Once
func Get_comonadStoreT() gopurs_runtime.Value {
	once_comonadStoreT.Do(func() {
		cache_comonadStoreT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadStoreT(dictComonad_0_box)
})
	})
	return cache_comonadStoreT
}

func Call_StoreT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_runStoreT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_functorStoreT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), f_1, h_3)
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1})}
}))
}

func Call_extendStoreT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
functorStoreT1_1_0 := gopurs_runtime.Apply(Get_functorStoreT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = functorStoreT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStoreT1_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictExtend_0.UnsafePtr)).V0, gopurs_runtime.Func2(func(w_prime_4 gopurs_runtime.Value, s_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{w_prime_4, s_prime_5})})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1})}
}))
}

func Call_comonadStoreT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
extendStoreT1_1_0 := gopurs_runtime.Apply(Get_extendStoreT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = extendStoreT1_1_0
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return extendStoreT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictComonad_0.UnsafePtr)).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}))
}


