package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Store_Trans_StoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_StoreT sync.Once
func Get_Control_Comonad_Store_Trans_StoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_StoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_StoreT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Trans_StoreT(x_0_box)
})
	})
	return cache_Control_Comonad_Store_Trans_StoreT
}

var cache_Control_Comonad_Store_Trans_runStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_runStoreT sync.Once
func Get_Control_Comonad_Store_Trans_runStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_runStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_runStoreT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Store_Trans_runStoreT(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Control_Comonad_Store_Trans_runStoreT
}

var cache_Control_Comonad_Store_Trans_newtypeStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_newtypeStoreT sync.Once
func Get_Control_Comonad_Store_Trans_newtypeStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_newtypeStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_newtypeStoreT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Control_Comonad_Store_Trans_newtypeStoreT
}

var cache_Control_Comonad_Store_Trans_functorStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_functorStoreT sync.Once
func Get_Control_Comonad_Store_Trans_functorStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_functorStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_functorStoreT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Trans_functorStoreT(dictFunctor_0_box)
})
	})
	return cache_Control_Comonad_Store_Trans_functorStoreT
}

var cache_Control_Comonad_Store_Trans_extendStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_extendStoreT sync.Once
func Get_Control_Comonad_Store_Trans_extendStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_extendStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_extendStoreT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Trans_extendStoreT(dictExtend_0_box)
})
	})
	return cache_Control_Comonad_Store_Trans_extendStoreT
}

var cache_Control_Comonad_Store_Trans_comonadTransStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_comonadTransStoreT sync.Once
func Get_Control_Comonad_Store_Trans_comonadTransStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_comonadTransStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_comonadTransStoreT = gopurs_runtime.Value{Type: 9, IntVal: 3399197123, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Trans_Class_ComonadTrans{1, gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1
_ = __local_var_3_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, __local_var_3_1)
}), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0)
})
})})}
	})
	return cache_Control_Comonad_Store_Trans_comonadTransStoreT
}

var cache_Control_Comonad_Store_Trans_comonadStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_comonadStoreT sync.Once
func Get_Control_Comonad_Store_Trans_comonadStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_comonadStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_comonadStoreT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Trans_comonadStoreT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Store_Trans_comonadStoreT
}

func Call_Control_Comonad_Store_Trans_StoreT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Store_Trans_runStoreT(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Comonad_Store_Trans_functorStoreT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(h_3, x_4))
})
}), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1})}
})
})})}
}

func Call_Control_Comonad_Store_Trans_extendStoreT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorStoreT1_1_0 -> *Constructor_Data_Functor_Functor
functorStoreT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(h_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(h_4, x_5))
})
}), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1})}
})
})))
_ = functorStoreT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStoreT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(w_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, w_prime_4, s_prime_5})})
})
}), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1})}
})
})})}
}

func Call_Control_Comonad_Store_Trans_comonadStoreT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorStoreT1_2_2 -> *Constructor_Data_Functor_Functor
functorStoreT1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), gopurs_runtime.Func(func(h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Apply(h_5, x_6))
})
}), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1})}
})
})))
_ = functorStoreT1_2_2
// TAST (Let): extendStoreT1_1_0 -> *Constructor_Control_Extend_Extend
extendStoreT1_1_0 := &Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStoreT1_2_2)}
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "extend"), gopurs_runtime.Func(func(w_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, w_prime_5, s_prime_6})})
})
}), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1})}
})
})}
_ = extendStoreT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendStoreT1_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})})}
}


