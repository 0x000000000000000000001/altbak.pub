package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Traced_Trans_TracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_TracedT sync.Once
func Get_Control_Comonad_Traced_Trans_TracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_TracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_TracedT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_TracedT(x_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_TracedT
}

var cache_Control_Comonad_Traced_Trans_runTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_runTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_runTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_runTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_runTracedT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_runTracedT(v_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_runTracedT
}

var cache_Control_Comonad_Traced_Trans_newtypeTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_newtypeTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_newtypeTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_newtypeTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_newtypeTracedT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Control_Comonad_Traced_Trans_newtypeTracedT
}

var cache_Control_Comonad_Traced_Trans_functorTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_functorTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_functorTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_functorTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_functorTracedT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_functorTracedT(dictFunctor_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_functorTracedT
}

var cache_Control_Comonad_Traced_Trans_extendTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_extendTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_extendTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_extendTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_extendTracedT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_extendTracedT(dictExtend_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_extendTracedT
}

var cache_Control_Comonad_Traced_Trans_comonadTransTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_comonadTransTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_comonadTransTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_comonadTransTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_comonadTransTracedT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_comonadTransTracedT(dictMonoid_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_comonadTransTracedT
}

var cache_Control_Comonad_Traced_Trans_comonadTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_comonadTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_comonadTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_comonadTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_comonadTracedT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_comonadTracedT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_comonadTracedT
}

func Call_Control_Comonad_Traced_Trans_TracedT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Traced_Trans_runTracedT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Comonad_Traced_Trans_functorTracedT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func2(func(g_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(g_3, t_4))
}), v_2)
})}))}
}

func Call_Control_Comonad_Traced_Trans_extendTracedT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorTracedT1_2_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])])
functorTracedT1_2_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func2(func(g_5 gopurs_runtime.Value, t_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Apply(g_5, t_6))
}), v_4)
})})
_ = functorTracedT1_2_1
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorTracedT1_2_1)}
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func2(func(w_prime__6 gopurs_runtime.Value, t_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func2(func(h_8 gopurs_runtime.Value, t_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_8, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_3, "append"), t_7, t_prime__9))
}), w_prime__6))
}), v_5)
})}))}
})
}

func Call_Control_Comonad_Traced_Trans_comonadTransTracedT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Value{Type: 9, IntVal: 3399197123, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictComonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_2_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_1, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
}), v_3)
})
})}))}
}

func Call_Control_Comonad_Traced_Trans_comonadTracedT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorTracedT1_3_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])])
functorTracedT1_3_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func2(func(g_6 gopurs_runtime.Value, t_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Apply(g_6, t_7))
}), v_5)
})})
_ = functorTracedT1_3_3
// TAST (Let): extendTracedT1__193435443_1_0 shape=Let(Let(Let(Abs(LitRecord)))) bindingType=Any
extendTracedT1__193435443_1_0 := gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorTracedT1_3_3)}
}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "extend"), gopurs_runtime.Func2(func(w_prime__7 gopurs_runtime.Value, t_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Func2(func(h_9 gopurs_runtime.Value, t_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_9, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_4, "append"), t_8, t_prime__10))
}), w_prime__7))
}), v_6)
})}))}
})
_ = extendTracedT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): extendTracedT2_3_5 shape=App(Other) bindingType=(ADT ["Control","Extend","Extend"] [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])])
extendTracedT2_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](gopurs_runtime.Apply(extendTracedT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})))
_ = extendTracedT2_3_5
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendTracedT2_3_5)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), v_4, gopurs_runtime.RecordGet(dictMonoid_2, "mempty"))
})}))}
})
}


