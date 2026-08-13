package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Writer_Trans_WriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_WriterT sync.Once
func Get_Control_Monad_Writer_Trans_WriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_WriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_WriterT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_WriterT(x_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_WriterT
}

var cache_Control_Monad_Writer_Trans_runWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_runWriterT sync.Once
func Get_Control_Monad_Writer_Trans_runWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_runWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_runWriterT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_runWriterT(v_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_runWriterT
}

var cache_Control_Monad_Writer_Trans_newtypeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_newtypeWriterT sync.Once
func Get_Control_Monad_Writer_Trans_newtypeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_newtypeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_newtypeWriterT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Monad_Writer_Trans_newtypeWriterT
}

var cache_Control_Monad_Writer_Trans_monadTransWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadTransWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadTransWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadTransWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadTransWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadTransWriterT
}

var cache_Control_Monad_Writer_Trans_mapWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_mapWriterT sync.Once
func Get_Control_Monad_Writer_Trans_mapWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_mapWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_mapWriterT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_mapWriterT(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_mapWriterT
}

var cache_Control_Monad_Writer_Trans_functorWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_functorWriterT sync.Once
func Get_Control_Monad_Writer_Trans_functorWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_functorWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_functorWriterT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_functorWriterT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_functorWriterT
}

var cache_Control_Monad_Writer_Trans_execWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_execWriterT sync.Once
func Get_Control_Monad_Writer_Trans_execWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_execWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_execWriterT = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_execWriterT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), v_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_execWriterT
}

var cache_Control_Monad_Writer_Trans_applyWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_applyWriterT sync.Once
func Get_Control_Monad_Writer_Trans_applyWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_applyWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_applyWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_0_box, dictApply_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_applyWriterT
}

var cache_Control_Monad_Writer_Trans_bindWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_bindWriterT sync.Once
func Get_Control_Monad_Writer_Trans_bindWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_bindWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_bindWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictBind_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_bindWriterT(dictSemigroup_0_box, dictBind_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_bindWriterT
}

var cache_Control_Monad_Writer_Trans_semigroupWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_semigroupWriterT sync.Once
func Get_Control_Monad_Writer_Trans_semigroupWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_semigroupWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_semigroupWriterT = gopurs_runtime.Func2(func(dictApply_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_semigroupWriterT(dictApply_0_box, dictSemigroup_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_semigroupWriterT
}

var cache_Control_Monad_Writer_Trans_applicativeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_applicativeWriterT sync.Once
func Get_Control_Monad_Writer_Trans_applicativeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_applicativeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_applicativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_applicativeWriterT
}

var cache_Control_Monad_Writer_Trans_monadWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadWriterT
}

var cache_Control_Monad_Writer_Trans_monadAskWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadAskWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadAskWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadAskWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadAskWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadAskWriterT
}

var cache_Control_Monad_Writer_Trans_monadReaderWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadReaderWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadReaderWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadReaderWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadReaderWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadReaderWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadReaderWriterT
}

var cache_Control_Monad_Writer_Trans_monadContWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadContWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadContWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadContWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadContWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadContWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadContWriterT
}

var cache_Control_Monad_Writer_Trans_monadEffectWriter gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadEffectWriter sync.Once
func Get_Control_Monad_Writer_Trans_monadEffectWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadEffectWriter.Do(func() {
		cache_Control_Monad_Writer_Trans_monadEffectWriter = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadEffectWriter
}

var cache_Control_Monad_Writer_Trans_monadRecWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadRecWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadRecWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadRecWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadRecWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadRecWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadRecWriterT
}

var cache_Control_Monad_Writer_Trans_monadStateWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadStateWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadStateWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadStateWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadStateWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadStateWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadStateWriterT
}

var cache_Control_Monad_Writer_Trans_monadTellWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadTellWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadTellWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadTellWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadTellWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadTellWriterT
}

var cache_Control_Monad_Writer_Trans_monadWriterWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadWriterWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadWriterWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadWriterWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadWriterWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadWriterWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadWriterWriterT
}

var cache_Control_Monad_Writer_Trans_monadThrowWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadThrowWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadThrowWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadThrowWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadThrowWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadThrowWriterT
}

var cache_Control_Monad_Writer_Trans_monadErrorWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadErrorWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadErrorWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadErrorWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadErrorWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadErrorWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadErrorWriterT
}

var cache_Control_Monad_Writer_Trans_monadSTWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadSTWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadSTWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadSTWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadSTWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadSTWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadSTWriterT
}

var cache_Control_Monad_Writer_Trans_monoidWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monoidWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monoidWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monoidWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monoidWriterT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monoidWriterT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monoidWriterT
}

var cache_Control_Monad_Writer_Trans_altWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_altWriterT sync.Once
func Get_Control_Monad_Writer_Trans_altWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_altWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_altWriterT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_altWriterT(dictAlt_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_altWriterT
}

var cache_Control_Monad_Writer_Trans_plusWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_plusWriterT sync.Once
func Get_Control_Monad_Writer_Trans_plusWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_plusWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_plusWriterT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_plusWriterT(dictPlus_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_plusWriterT
}

var cache_Control_Monad_Writer_Trans_alternativeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_alternativeWriterT sync.Once
func Get_Control_Monad_Writer_Trans_alternativeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_alternativeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_alternativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_alternativeWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_alternativeWriterT
}

var cache_Control_Monad_Writer_Trans_monadPlusWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadPlusWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadPlusWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadPlusWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadPlusWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadPlusWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadPlusWriterT
}

var cache_Control_Monad_Writer_Trans_mapWriterT__2842489082 gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_mapWriterT__2842489082 sync.Once
func Get_Control_Monad_Writer_Trans_mapWriterT__2842489082() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_mapWriterT__2842489082.Do(func() {
		cache_Control_Monad_Writer_Trans_mapWriterT__2842489082 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_mapWriterT__2842489082(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_mapWriterT__2842489082
}

var cache_Control_Monad_Writer_Trans_mapWriterT__4072164636 gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_mapWriterT__4072164636 sync.Once
func Get_Control_Monad_Writer_Trans_mapWriterT__4072164636() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_mapWriterT__4072164636.Do(func() {
		cache_Control_Monad_Writer_Trans_mapWriterT__4072164636 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_Trans_mapWriterT__4072164636(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box)))}
})
	})
	return cache_Control_Monad_Writer_Trans_mapWriterT__4072164636
}

var cache_Control_Monad_Writer_Trans_mapWriterT__77717660 gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_mapWriterT__77717660 sync.Once
func Get_Control_Monad_Writer_Trans_mapWriterT__77717660() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_mapWriterT__77717660.Do(func() {
		cache_Control_Monad_Writer_Trans_mapWriterT__77717660 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_mapWriterT__77717660(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_mapWriterT__77717660
}

var cache_Control_Monad_Writer_Trans_runWriterT__4273258459 gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_runWriterT__4273258459 sync.Once
func Get_Control_Monad_Writer_Trans_runWriterT__4273258459() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_runWriterT__4273258459.Do(func() {
		cache_Control_Monad_Writer_Trans_runWriterT__4273258459 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_Trans_runWriterT__4273258459(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Control_Monad_Writer_Trans_runWriterT__4273258459
}

func Call_Control_Monad_Writer_Trans_WriterT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Writer_Trans_runWriterT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_0 -> *Constructor_Control_Bind_Bind
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): pure_3_1 -> gopurs_runtime.Value
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
}))
}

func Call_Control_Monad_Writer_Trans_mapWriterT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Writer_Trans_functorWriterT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1})}
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, v_3)
})
}))
}

func Call_Control_Monad_Writer_Trans_execWriterT(dictFunctor_0_loop *Constructor_Data_Functor_Functor, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), Get_Data_Tuple_snd(), v_1)
}

func Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
// TAST (Let): Functor0_2_0 -> *Constructor_Data_Functor_Functor
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorWriterT1_3_1 -> gopurs_runtime.Value
functorWriterT1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1})}
}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, v_6)
})
}))
_ = functorWriterT1_3_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_7.UnsafePtr).V1)})}
})
}), v_4), v1_5)
})
}))
}

func Call_Control_Monad_Writer_Trans_bindWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictBind_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictBind_1 gopurs_runtime.Value = dictBind_1_loop
_ = dictBind_1
// TAST (Let): Apply0_2_0 -> gopurs_runtime.Value
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_1, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_2_0
// TAST (Let): Functor0_3_1 -> *Constructor_Data_Functor_Functor
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorWriterT1_5_4 -> gopurs_runtime.Value
functorWriterT1_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1})}
}))
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
}))
_ = functorWriterT1_5_4
// TAST (Let): applyWriterT2_4_2 -> gopurs_runtime.Value
applyWriterT2_4_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), gopurs_runtime.Func(func(v3_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_9.UnsafePtr).V1)})}
})
}), v_6), v1_7)
})
}))
_ = applyWriterT2_4_2
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_4_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_1, "bind"), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_7 -> gopurs_runtime.Value
__local_var_8_7 := (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1
_ = __local_var_8_7
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_9.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_8_7, (*Constructor_Data_Tuple_Tuple)(v3_9.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_6, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0))
}))
})
}))
}

func Call_Control_Monad_Writer_Trans_semigroupWriterT(dictApply_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorWriterT1_3_2 -> gopurs_runtime.Value
functorWriterT1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1})}
}))
_ = __local_var_5_4
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, v_6)
})
}))
_ = functorWriterT1_3_2
// TAST (Let): applyWriterT1_2_0 -> *Constructor_Control_Apply_Apply
applyWriterT1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_7.UnsafePtr).V1)})}
})
}), v_4), v1_5)
})
})))
_ = applyWriterT1_2_0
return gopurs_runtime.Func(func(dictSemigroup1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 -> *Constructor_Data_Functor_Functor
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyWriterT1_2_0.V0), gopurs_runtime.Value{}))
_ = Functor0_4_5
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.RecordGet(dictSemigroup1_3, "append")
_ = __local_var_5_6
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyWriterT1_2_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), __local_var_5_6, a_6), b_7)
})
}))
})
}

func Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorWriterT1_5_4 -> gopurs_runtime.Value
functorWriterT1_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1})}
}))
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
}))
_ = functorWriterT1_5_4
// TAST (Let): applyWriterT2_3_1 -> gopurs_runtime.Value
applyWriterT2_3_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), gopurs_runtime.Func(func(v3_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_9.UnsafePtr).V1)})}
})
}), v_6), v1_7)
})
}))
_ = applyWriterT2_3_1
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_3_1
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
}

func Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): Functor0_6_6 -> *Constructor_Data_Functor_Functor
Functor0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_6
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): functorWriterT1_7_7 -> gopurs_runtime.Value
functorWriterT1_7_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1})}
}))
_ = __local_var_9_9
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_9, v_10)
})
}))
_ = functorWriterT1_7_7
// TAST (Let): applyWriterT2_5_4 -> gopurs_runtime.Value
applyWriterT2_5_4 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_7_7
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_6.V0), gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_10.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_11.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_10.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_11.UnsafePtr).V1)})}
})
}), v_8), v1_9)
})
}))
_ = applyWriterT2_5_4
// TAST (Let): applicativeWriterT2_4_2 -> gopurs_runtime.Value
applicativeWriterT2_4_2 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_5_4
}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_6, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_4_2
// TAST (Let): __local_var_5_11 -> gopurs_runtime.Value
__local_var_5_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_11
// TAST (Let): Apply0_6_12 -> gopurs_runtime.Value
Apply0_6_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_11, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_12
// TAST (Let): Functor0_7_13 -> *Constructor_Data_Functor_Functor
Functor0_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_13
// TAST (Let): Functor0_8_15 -> *Constructor_Data_Functor_Functor
Functor0_8_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_15
// TAST (Let): __local_var_9_17 -> gopurs_runtime.Value
__local_var_9_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_17
// TAST (Let): functorWriterT1_9_16 -> gopurs_runtime.Value
functorWriterT1_9_16 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_18 -> gopurs_runtime.Value
__local_var_11_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_17, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1})}
}))
_ = __local_var_11_18
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_18, v_12)
})
}))
_ = functorWriterT1_9_16
// TAST (Let): applyWriterT2_8_14 -> gopurs_runtime.Value
applyWriterT2_8_14 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_9_16
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_6_12, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_15.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_13.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_12.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_13.UnsafePtr).V1)})}
})
}), v_10), v1_11)
})
}))
_ = applyWriterT2_8_14
// TAST (Let): bindWriterT2_5_10 -> gopurs_runtime.Value
bindWriterT2_5_10 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_8_14
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_11, "bind"), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_19 -> gopurs_runtime.Value
__local_var_12_19 := (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1
_ = __local_var_12_19
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_13.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_12_19, (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_10, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_5_10
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_5_10
}))
})
}

func Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 -> *Constructor_Control_Monad_Trans_Class_MonadTrans
monadTransWriterT1_1_0 := &Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_4
return gopurs_runtime.Func(func(dictMonadAsk_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_4, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): __local_var_6_8 -> gopurs_runtime.Value
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_6_8
// TAST (Let): __local_var_7_10 -> gopurs_runtime.Value
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_7_10
// TAST (Let): Functor0_8_11 -> *Constructor_Data_Functor_Functor
Functor0_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_11
// TAST (Let): __local_var_9_13 -> gopurs_runtime.Value
__local_var_9_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_13
// TAST (Let): functorWriterT1_9_12 -> gopurs_runtime.Value
functorWriterT1_9_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_14 -> gopurs_runtime.Value
__local_var_11_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1})}
}))
_ = __local_var_11_14
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_14, v_12)
})
}))
_ = functorWriterT1_9_12
// TAST (Let): applyWriterT2_7_9 -> gopurs_runtime.Value
applyWriterT2_7_9 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_9_12
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_10, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_11.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_13.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "append"), (*Constructor_Data_Tuple_Tuple)(v3_12.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_13.UnsafePtr).V1)})}
})
}), v_10), v1_11)
})
}))
_ = applyWriterT2_7_9
// TAST (Let): applicativeWriterT2_6_7 -> gopurs_runtime.Value
applicativeWriterT2_6_7 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_7_9
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_6_7
// TAST (Let): __local_var_7_16 -> gopurs_runtime.Value
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_7_16
// TAST (Let): Apply0_8_17 -> gopurs_runtime.Value
Apply0_8_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_16, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_8_17
// TAST (Let): Functor0_9_18 -> *Constructor_Data_Functor_Functor
Functor0_9_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_18
// TAST (Let): Functor0_10_20 -> *Constructor_Data_Functor_Functor
Functor0_10_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_20
// TAST (Let): __local_var_11_22 -> gopurs_runtime.Value
__local_var_11_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_17, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_22
// TAST (Let): functorWriterT1_11_21 -> gopurs_runtime.Value
functorWriterT1_11_21 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_23 -> gopurs_runtime.Value
__local_var_13_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_22, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_13.UnsafePtr).V1})}
}))
_ = __local_var_13_23
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_23, v_14)
})
}))
_ = functorWriterT1_11_21
// TAST (Let): applyWriterT2_10_19 -> gopurs_runtime.Value
applyWriterT2_10_19 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_11_21
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_8_17, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_20.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_15.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "append"), (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_15.UnsafePtr).V1)})}
})
}), v_12), v1_13)
})
}))
_ = applyWriterT2_10_19
// TAST (Let): bindWriterT2_7_15 -> gopurs_runtime.Value
bindWriterT2_7_15 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_10_19
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_16, "bind"), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_24 -> gopurs_runtime.Value
__local_var_14_24 := (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1
_ = __local_var_14_24
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_18.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "append"), __local_var_14_24, (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_12, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_7_15
// TAST (Let): monadWriterT2_5_5 -> gopurs_runtime.Value
monadWriterT2_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_6_7
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_7_15
}))
_ = monadWriterT2_5_5
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_5
}), gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_4, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_4, "ask")))
})
}

func Call_Control_Monad_Writer_Trans_monadReaderWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 -> *Constructor_Control_Monad_Trans_Class_MonadTrans
monadTransWriterT1_1_0 := &Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): monadAskWriterT1_3_4 -> gopurs_runtime.Value
monadAskWriterT1_3_4 := gopurs_runtime.Func(func(dictMonadAsk_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_7 -> gopurs_runtime.Value
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_4, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_5_7
// TAST (Let): __local_var_6_9 -> gopurs_runtime.Value
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_6_9
// TAST (Let): __local_var_7_11 -> gopurs_runtime.Value
__local_var_7_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_9, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_7_11
// TAST (Let): Functor0_8_12 -> *Constructor_Data_Functor_Functor
Functor0_8_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_12
// TAST (Let): __local_var_9_14 -> gopurs_runtime.Value
__local_var_9_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_14
// TAST (Let): functorWriterT1_9_13 -> gopurs_runtime.Value
functorWriterT1_9_13 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_15 -> gopurs_runtime.Value
__local_var_11_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_14, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1})}
}))
_ = __local_var_11_15
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_15, v_12)
})
}))
_ = functorWriterT1_9_13
// TAST (Let): applyWriterT2_7_10 -> gopurs_runtime.Value
applyWriterT2_7_10 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_9_13
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_11, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_12.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_13.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "append"), (*Constructor_Data_Tuple_Tuple)(v3_12.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_13.UnsafePtr).V1)})}
})
}), v_10), v1_11)
})
}))
_ = applyWriterT2_7_10
// TAST (Let): applicativeWriterT2_6_8 -> gopurs_runtime.Value
applicativeWriterT2_6_8 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_7_10
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_9, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_6_8
// TAST (Let): __local_var_7_17 -> gopurs_runtime.Value
__local_var_7_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_7_17
// TAST (Let): Apply0_8_18 -> gopurs_runtime.Value
Apply0_8_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_8_18
// TAST (Let): Functor0_9_19 -> *Constructor_Data_Functor_Functor
Functor0_9_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_18, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_19
// TAST (Let): Functor0_10_21 -> *Constructor_Data_Functor_Functor
Functor0_10_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_18, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_21
// TAST (Let): __local_var_11_23 -> gopurs_runtime.Value
__local_var_11_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_18, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_23
// TAST (Let): functorWriterT1_11_22 -> gopurs_runtime.Value
functorWriterT1_11_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_24 -> gopurs_runtime.Value
__local_var_13_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_23, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_13.UnsafePtr).V1})}
}))
_ = __local_var_13_24
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_24, v_14)
})
}))
_ = functorWriterT1_11_22
// TAST (Let): applyWriterT2_10_20 -> gopurs_runtime.Value
applyWriterT2_10_20 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_11_22
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_8_18, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_21.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_15.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "append"), (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_15.UnsafePtr).V1)})}
})
}), v_12), v1_13)
})
}))
_ = applyWriterT2_10_20
// TAST (Let): bindWriterT2_7_16 -> gopurs_runtime.Value
bindWriterT2_7_16 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_10_20
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_17, "bind"), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_25 -> gopurs_runtime.Value
__local_var_14_25 := (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1
_ = __local_var_14_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_19.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "append"), __local_var_14_25, (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_12, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_7_16
// TAST (Let): monadWriterT2_5_6 -> gopurs_runtime.Value
monadWriterT2_5_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_6_8
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_7_16
}))
_ = monadWriterT2_5_6
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_6
}), gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_4, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_4, "ask")))
})
_ = monadAskWriterT1_3_4
return gopurs_runtime.Func(func(dictMonadReader_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadAskWriterT2_5_26 -> gopurs_runtime.Value
monadAskWriterT2_5_26 := gopurs_runtime.Apply(monadAskWriterT1_3_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_4, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskWriterT2_5_26
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskWriterT2_5_26
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_4, "local"), f_6)
_ = __local_var_7_27
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_27, v_8)
})
}))
})
}

func Call_Control_Monad_Writer_Trans_monadContWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonadCont_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_3, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): Functor0_7_8 -> *Constructor_Data_Functor_Functor
Functor0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_8
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorWriterT1_8_9 -> gopurs_runtime.Value
functorWriterT1_8_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1})}
}))
_ = __local_var_10_11
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_11, v_11)
})
}))
_ = functorWriterT1_8_9
// TAST (Let): applyWriterT2_6_6 -> gopurs_runtime.Value
applyWriterT2_6_6 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_8_9
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_8.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V1)})}
})
}), v_9), v1_10)
})
}))
_ = applyWriterT2_6_6
// TAST (Let): applicativeWriterT2_5_4 -> gopurs_runtime.Value
applicativeWriterT2_5_4 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_6_6
}), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_5_4
// TAST (Let): __local_var_6_13 -> gopurs_runtime.Value
__local_var_6_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_6_13
// TAST (Let): Apply0_7_14 -> gopurs_runtime.Value
Apply0_7_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_13, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_14
// TAST (Let): Functor0_8_15 -> *Constructor_Data_Functor_Functor
Functor0_8_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_15
// TAST (Let): Functor0_9_17 -> *Constructor_Data_Functor_Functor
Functor0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_17
// TAST (Let): __local_var_10_19 -> gopurs_runtime.Value
__local_var_10_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_19
// TAST (Let): functorWriterT1_10_18 -> gopurs_runtime.Value
functorWriterT1_10_18 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_20 -> gopurs_runtime.Value
__local_var_12_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_19, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_20
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_20, v_13)
})
}))
_ = functorWriterT1_10_18
// TAST (Let): applyWriterT2_9_16 -> gopurs_runtime.Value
applyWriterT2_9_16 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_18
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_14, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_17.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_9_16
// TAST (Let): bindWriterT2_6_12 -> gopurs_runtime.Value
bindWriterT2_6_12 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_9_16
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_13, "bind"), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_21 -> gopurs_runtime.Value
__local_var_13_21 := (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1
_ = __local_var_13_21
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_15.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_13_21, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_11, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_6_12
// TAST (Let): monadWriterT2_4_2 -> gopurs_runtime.Value
monadWriterT2_4_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_5_4
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_6_12
}))
_ = monadWriterT2_4_2
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_3, "callCC"), gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
}))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonadEffect_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_2 -> gopurs_runtime.Value
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): Functor0_7_8 -> *Constructor_Data_Functor_Functor
Functor0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_8
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorWriterT1_8_9 -> gopurs_runtime.Value
functorWriterT1_8_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1})}
}))
_ = __local_var_10_11
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_11, v_11)
})
}))
_ = functorWriterT1_8_9
// TAST (Let): applyWriterT2_6_6 -> gopurs_runtime.Value
applyWriterT2_6_6 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_8_9
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_8.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V1)})}
})
}), v_9), v1_10)
})
}))
_ = applyWriterT2_6_6
// TAST (Let): applicativeWriterT2_5_4 -> gopurs_runtime.Value
applicativeWriterT2_5_4 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_6_6
}), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_5_4
// TAST (Let): __local_var_6_13 -> gopurs_runtime.Value
__local_var_6_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_6_13
// TAST (Let): Apply0_7_14 -> gopurs_runtime.Value
Apply0_7_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_13, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_14
// TAST (Let): Functor0_8_15 -> *Constructor_Data_Functor_Functor
Functor0_8_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_15
// TAST (Let): Functor0_9_17 -> *Constructor_Data_Functor_Functor
Functor0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_17
// TAST (Let): __local_var_10_19 -> gopurs_runtime.Value
__local_var_10_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_19
// TAST (Let): functorWriterT1_10_18 -> gopurs_runtime.Value
functorWriterT1_10_18 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_20 -> gopurs_runtime.Value
__local_var_12_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_19, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_20
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_20, v_13)
})
}))
_ = functorWriterT1_10_18
// TAST (Let): applyWriterT2_9_16 -> gopurs_runtime.Value
applyWriterT2_9_16 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_18
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_14, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_17.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_9_16
// TAST (Let): bindWriterT2_6_12 -> gopurs_runtime.Value
bindWriterT2_6_12 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_9_16
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_13, "bind"), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_21 -> gopurs_runtime.Value
__local_var_13_21 := (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1
_ = __local_var_13_21
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_15.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_13_21, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_11, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_6_12
// TAST (Let): monadWriterT2_5_3 -> gopurs_runtime.Value
monadWriterT2_5_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_5_4
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_6_12
}))
_ = monadWriterT2_5_3
// TAST (Let): Bind1_6_23 -> *Constructor_Control_Bind_Bind
Bind1_6_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_23
// TAST (Let): pure_7_24 -> gopurs_runtime.Value
pure_7_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_24
// TAST (Let): __local_var_6_22 -> gopurs_runtime.Value
__local_var_6_22 := gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_23.V1), m_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_24, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
_ = __local_var_6_22
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_22, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_3, "liftEffect"), x_7))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadRecWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictMonadRec_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_5_3 -> gopurs_runtime.Value
Monad0_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_4, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_5_3
// TAST (Let): Bind1_6_4 -> *Constructor_Control_Bind_Bind
Bind1_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_5_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_4
// TAST (Let): Applicative0_7_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_5_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_5
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_5_3, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_8_8
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_9_10
// TAST (Let): Functor0_10_11 -> *Constructor_Data_Functor_Functor
Functor0_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_11
// TAST (Let): __local_var_11_13 -> gopurs_runtime.Value
__local_var_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_13
// TAST (Let): functorWriterT1_11_12 -> gopurs_runtime.Value
functorWriterT1_11_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_14 -> gopurs_runtime.Value
__local_var_13_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_13, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_13.UnsafePtr).V1})}
}))
_ = __local_var_13_14
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_14, v_14)
})
}))
_ = functorWriterT1_11_12
// TAST (Let): applyWriterT2_9_9 -> gopurs_runtime.Value
applyWriterT2_9_9 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_11_12
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_10, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_11.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_15.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_15.UnsafePtr).V1)})}
})
}), v_12), v1_13)
})
}))
_ = applyWriterT2_9_9
// TAST (Let): applicativeWriterT2_8_7 -> gopurs_runtime.Value
applicativeWriterT2_8_7 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_9_9
}), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_10, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_8_7
// TAST (Let): __local_var_9_16 -> gopurs_runtime.Value
__local_var_9_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_5_3, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_9_16
// TAST (Let): Apply0_10_17 -> gopurs_runtime.Value
Apply0_10_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_16, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_10_17
// TAST (Let): Functor0_11_18 -> *Constructor_Data_Functor_Functor
Functor0_11_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_11_18
// TAST (Let): Functor0_12_20 -> *Constructor_Data_Functor_Functor
Functor0_12_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_12_20
// TAST (Let): __local_var_13_22 -> gopurs_runtime.Value
__local_var_13_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_17, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_22
// TAST (Let): functorWriterT1_13_21 -> gopurs_runtime.Value
functorWriterT1_13_21 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_23 -> gopurs_runtime.Value
__local_var_15_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_22, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Tuple_Tuple)(v_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_15.UnsafePtr).V1})}
}))
_ = __local_var_15_23
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_23, v_16)
})
}))
_ = functorWriterT1_13_21
// TAST (Let): applyWriterT2_12_19 -> gopurs_runtime.Value
applyWriterT2_12_19 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_13_21
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_10_17, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_12_20.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_17.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "append"), (*Constructor_Data_Tuple_Tuple)(v3_16.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_17.UnsafePtr).V1)})}
})
}), v_14), v1_15)
})
}))
_ = applyWriterT2_12_19
// TAST (Let): bindWriterT2_9_15 -> gopurs_runtime.Value
bindWriterT2_9_15 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_12_19
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_16, "bind"), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_24 -> gopurs_runtime.Value
__local_var_16_24 := (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1
_ = __local_var_16_24
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_18.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_17.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "append"), __local_var_16_24, (*Constructor_Data_Tuple_Tuple)(v3_17.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_14, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_9_15
// TAST (Let): monadWriterT2_8_6 -> gopurs_runtime.Value
monadWriterT2_8_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_8_7
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_9_15
}))
_ = monadWriterT2_8_6
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_8_6
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_4, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_25 -> gopurs_runtime.Value
__local_var_12_25 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_12_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_4.V1), gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 gopurs_runtime.Value
{
var __t_tag_26 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v2_13.UnsafePtr).V0
if (__t_tag_26.Type == 9 && __t_tag_26.IntVal == 525585346) {
__t28 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Tuple_Tuple)(v2_13.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), __local_var_12_25, (*Constructor_Data_Tuple_Tuple)(v2_13.UnsafePtr).V1)})}})}
goto end_branch_28
} else {

}
}
{
var __t_tag_27 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v2_13.UnsafePtr).V0
if (__t_tag_27.Type == 9 && __t_tag_27.IntVal == 60402430) {
__t28 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Tuple_Tuple)(v2_13.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), __local_var_12_25, (*Constructor_Data_Tuple_Tuple)(v2_13.UnsafePtr).V1)})}})}
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_5.V1), __t28)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_10, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
})
}))
})
}

func Call_Control_Monad_Writer_Trans_monadStateWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 -> *Constructor_Control_Monad_Trans_Class_MonadTrans
monadTransWriterT1_1_0 := &Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_4
return gopurs_runtime.Func(func(dictMonadState_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_5_5 -> *Constructor_Control_Monad_Monad
Monad0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_4, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_4, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): __local_var_7_9 -> gopurs_runtime.Value
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_7_9
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_9, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_8_11
// TAST (Let): Functor0_9_12 -> *Constructor_Data_Functor_Functor
Functor0_9_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_12
// TAST (Let): __local_var_10_14 -> gopurs_runtime.Value
__local_var_10_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_14
// TAST (Let): functorWriterT1_10_13 -> gopurs_runtime.Value
functorWriterT1_10_13 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_15 -> gopurs_runtime.Value
__local_var_12_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_14, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_15
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_15, v_13)
})
}))
_ = functorWriterT1_10_13
// TAST (Let): applyWriterT2_8_10 -> gopurs_runtime.Value
applyWriterT2_8_10 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_13
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_11, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_12.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_8_10
// TAST (Let): applicativeWriterT2_7_8 -> gopurs_runtime.Value
applicativeWriterT2_7_8 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_8_10
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_9, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_7_8
// TAST (Let): __local_var_8_17 -> gopurs_runtime.Value
__local_var_8_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_8_17
// TAST (Let): Apply0_9_18 -> gopurs_runtime.Value
Apply0_9_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_17, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_9_18
// TAST (Let): Functor0_10_19 -> *Constructor_Data_Functor_Functor
Functor0_10_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_18, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_19
// TAST (Let): Functor0_11_21 -> *Constructor_Data_Functor_Functor
Functor0_11_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_18, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_11_21
// TAST (Let): __local_var_12_23 -> gopurs_runtime.Value
__local_var_12_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_18, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_23
// TAST (Let): functorWriterT1_12_22 -> gopurs_runtime.Value
functorWriterT1_12_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_24 -> gopurs_runtime.Value
__local_var_14_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_23, "map"), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_13, (*Constructor_Data_Tuple_Tuple)(v_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_14.UnsafePtr).V1})}
}))
_ = __local_var_14_24
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_24, v_15)
})
}))
_ = functorWriterT1_12_22
// TAST (Let): applyWriterT2_11_20 -> gopurs_runtime.Value
applyWriterT2_11_20 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_12_22
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_9_18, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_21.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_16.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "append"), (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_16.UnsafePtr).V1)})}
})
}), v_13), v1_14)
})
}))
_ = applyWriterT2_11_20
// TAST (Let): bindWriterT2_8_16 -> gopurs_runtime.Value
bindWriterT2_8_16 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_11_20
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_17, "bind"), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_25 -> gopurs_runtime.Value
__local_var_15_25 := (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1
_ = __local_var_15_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_19.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_16.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "append"), __local_var_15_25, (*Constructor_Data_Tuple_Tuple)(v3_16.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_13, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_8_16
// TAST (Let): monadWriterT2_6_6 -> gopurs_runtime.Value
monadWriterT2_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_7_8
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_8_16
}))
_ = monadWriterT2_6_6
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_6_6
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_5_5)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_4, "state"), f_7))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 -> gopurs_runtime.Value
Semigroup0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): Functor0_7_8 -> *Constructor_Data_Functor_Functor
Functor0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_8
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorWriterT1_8_9 -> gopurs_runtime.Value
functorWriterT1_8_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1})}
}))
_ = __local_var_10_11
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_11, v_11)
})
}))
_ = functorWriterT1_8_9
// TAST (Let): applyWriterT2_6_6 -> gopurs_runtime.Value
applyWriterT2_6_6 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_8_9
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_8.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V1)})}
})
}), v_9), v1_10)
})
}))
_ = applyWriterT2_6_6
// TAST (Let): applicativeWriterT2_5_4 -> gopurs_runtime.Value
applicativeWriterT2_5_4 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_6_6
}), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_5_4
// TAST (Let): __local_var_6_13 -> gopurs_runtime.Value
__local_var_6_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_6_13
// TAST (Let): Apply0_7_14 -> gopurs_runtime.Value
Apply0_7_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_13, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_14
// TAST (Let): Functor0_8_15 -> *Constructor_Data_Functor_Functor
Functor0_8_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_15
// TAST (Let): Functor0_9_17 -> *Constructor_Data_Functor_Functor
Functor0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_17
// TAST (Let): __local_var_10_19 -> gopurs_runtime.Value
__local_var_10_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_19
// TAST (Let): functorWriterT1_10_18 -> gopurs_runtime.Value
functorWriterT1_10_18 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_20 -> gopurs_runtime.Value
__local_var_12_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_19, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_20
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_20, v_13)
})
}))
_ = functorWriterT1_10_18
// TAST (Let): applyWriterT2_9_16 -> gopurs_runtime.Value
applyWriterT2_9_16 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_18
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_14, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_17.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_9_16
// TAST (Let): bindWriterT2_6_12 -> gopurs_runtime.Value
bindWriterT2_6_12 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_9_16
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_13, "bind"), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_21 -> gopurs_runtime.Value
__local_var_13_21 := (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1
_ = __local_var_13_21
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_15.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "append"), __local_var_13_21, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_11, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_6_12
// TAST (Let): monadWriterT2_5_3 -> gopurs_runtime.Value
monadWriterT2_5_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_5_4
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_6_12
}))
_ = monadWriterT2_5_3
// TAST (Let): __local_var_6_23 -> gopurs_runtime.Value
__local_var_6_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_6_23
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), Get_Data_Unit_unit())
_ = __local_var_7_24
// TAST (Let): __local_var_6_22 -> gopurs_runtime.Value
__local_var_6_22 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_23, gopurs_runtime.Apply(__local_var_7_24, x_8))
})
_ = __local_var_6_22
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_1_0
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_22, x_7)
}))
})
}

func Call_Control_Monad_Writer_Trans_monadWriterWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 -> gopurs_runtime.Value
Semigroup0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): monadTellWriterT1_3_2 -> gopurs_runtime.Value
monadTellWriterT1_3_2 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): __local_var_6_8 -> gopurs_runtime.Value
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_6_8
// TAST (Let): Functor0_7_9 -> *Constructor_Data_Functor_Functor
Functor0_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_9
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_11
// TAST (Let): functorWriterT1_8_10 -> gopurs_runtime.Value
functorWriterT1_8_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_12 -> gopurs_runtime.Value
__local_var_10_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_11, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1})}
}))
_ = __local_var_10_12
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_12, v_11)
})
}))
_ = functorWriterT1_8_10
// TAST (Let): applyWriterT2_6_7 -> gopurs_runtime.Value
applyWriterT2_6_7 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_8_10
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_8, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_9.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V1)})}
})
}), v_9), v1_10)
})
}))
_ = applyWriterT2_6_7
// TAST (Let): applicativeWriterT2_5_5 -> gopurs_runtime.Value
applicativeWriterT2_5_5 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_6_7
}), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_5_5
// TAST (Let): __local_var_6_14 -> gopurs_runtime.Value
__local_var_6_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_6_14
// TAST (Let): Apply0_7_15 -> gopurs_runtime.Value
Apply0_7_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_14, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_15
// TAST (Let): Functor0_8_16 -> *Constructor_Data_Functor_Functor
Functor0_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_16
// TAST (Let): Functor0_9_18 -> *Constructor_Data_Functor_Functor
Functor0_9_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_18
// TAST (Let): __local_var_10_20 -> gopurs_runtime.Value
__local_var_10_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_15, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_20
// TAST (Let): functorWriterT1_10_19 -> gopurs_runtime.Value
functorWriterT1_10_19 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_21 -> gopurs_runtime.Value
__local_var_12_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_20, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_21
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_21, v_13)
})
}))
_ = functorWriterT1_10_19
// TAST (Let): applyWriterT2_9_17 -> gopurs_runtime.Value
applyWriterT2_9_17 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_19
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_15, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_18.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_9_17
// TAST (Let): bindWriterT2_6_13 -> gopurs_runtime.Value
bindWriterT2_6_13 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_9_17
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_14, "bind"), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_22 -> gopurs_runtime.Value
__local_var_13_22 := (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1
_ = __local_var_13_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_16.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "append"), __local_var_13_22, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_11, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_6_13
// TAST (Let): monadWriterT2_5_4 -> gopurs_runtime.Value
monadWriterT2_5_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_5_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_6_13
}))
_ = monadWriterT2_5_4
// TAST (Let): __local_var_6_24 -> gopurs_runtime.Value
__local_var_6_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_6_24
// TAST (Let): __local_var_7_25 -> gopurs_runtime.Value
__local_var_7_25 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), Get_Data_Unit_unit())
_ = __local_var_7_25
// TAST (Let): __local_var_6_23 -> gopurs_runtime.Value
__local_var_6_23 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_24, gopurs_runtime.Apply(__local_var_7_25, x_8))
})
_ = __local_var_6_23
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_1_0
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_23, x_7)
}))
})
_ = monadTellWriterT1_3_2
return gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_26 -> *Constructor_Control_Bind_Bind
Bind1_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_26
// TAST (Let): Applicative0_6_27 -> gopurs_runtime.Value
Applicative0_6_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_6_27
// TAST (Let): monadTellWriterT2_7_28 -> gopurs_runtime.Value
monadTellWriterT2_7_28 := gopurs_runtime.Apply(monadTellWriterT1_3_2, dictMonad_4)
_ = monadTellWriterT2_7_28
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellWriterT2_7_28
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_0
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_26.V1), v_8, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_6_27, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_26.V1), v_8, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_6_27, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1)})})
}))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 -> *Constructor_Control_Monad_Trans_Class_MonadTrans
monadTransWriterT1_1_0 := &Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_4
return gopurs_runtime.Func(func(dictMonadThrow_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_5_5 -> *Constructor_Control_Monad_Monad
Monad0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_4, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_4, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): __local_var_7_9 -> gopurs_runtime.Value
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_7_9
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_9, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_8_11
// TAST (Let): Functor0_9_12 -> *Constructor_Data_Functor_Functor
Functor0_9_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_12
// TAST (Let): __local_var_10_14 -> gopurs_runtime.Value
__local_var_10_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_14
// TAST (Let): functorWriterT1_10_13 -> gopurs_runtime.Value
functorWriterT1_10_13 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_15 -> gopurs_runtime.Value
__local_var_12_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_14, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_15
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_15, v_13)
})
}))
_ = functorWriterT1_10_13
// TAST (Let): applyWriterT2_8_10 -> gopurs_runtime.Value
applyWriterT2_8_10 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_13
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_11, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_12.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_8_10
// TAST (Let): applicativeWriterT2_7_8 -> gopurs_runtime.Value
applicativeWriterT2_7_8 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_8_10
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_9, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_7_8
// TAST (Let): __local_var_8_17 -> gopurs_runtime.Value
__local_var_8_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_8_17
// TAST (Let): Apply0_9_18 -> gopurs_runtime.Value
Apply0_9_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_17, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_9_18
// TAST (Let): Functor0_10_19 -> *Constructor_Data_Functor_Functor
Functor0_10_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_18, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_19
// TAST (Let): Functor0_11_21 -> *Constructor_Data_Functor_Functor
Functor0_11_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_18, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_11_21
// TAST (Let): __local_var_12_23 -> gopurs_runtime.Value
__local_var_12_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_18, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_23
// TAST (Let): functorWriterT1_12_22 -> gopurs_runtime.Value
functorWriterT1_12_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_24 -> gopurs_runtime.Value
__local_var_14_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_23, "map"), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_13, (*Constructor_Data_Tuple_Tuple)(v_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_14.UnsafePtr).V1})}
}))
_ = __local_var_14_24
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_24, v_15)
})
}))
_ = functorWriterT1_12_22
// TAST (Let): applyWriterT2_11_20 -> gopurs_runtime.Value
applyWriterT2_11_20 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_12_22
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_9_18, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_21.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_16.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "append"), (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_16.UnsafePtr).V1)})}
})
}), v_13), v1_14)
})
}))
_ = applyWriterT2_11_20
// TAST (Let): bindWriterT2_8_16 -> gopurs_runtime.Value
bindWriterT2_8_16 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_11_20
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_17, "bind"), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_25 -> gopurs_runtime.Value
__local_var_15_25 := (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1
_ = __local_var_15_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_19.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_16.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "append"), __local_var_15_25, (*Constructor_Data_Tuple_Tuple)(v3_16.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_13, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_8_16
// TAST (Let): monadWriterT2_6_6 -> gopurs_runtime.Value
monadWriterT2_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_7_8
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_8_16
}))
_ = monadWriterT2_6_6
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_6_6
}), gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_5_5)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_4, "throwError"), e_7))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadErrorWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 -> *Constructor_Control_Monad_Trans_Class_MonadTrans
monadTransWriterT1_1_0 := &Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): monadThrowWriterT1_3_4 -> gopurs_runtime.Value
monadThrowWriterT1_3_4 := gopurs_runtime.Func(func(dictMonadThrow_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_5_6 -> *Constructor_Control_Monad_Monad
Monad0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_4, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_5_6
// TAST (Let): __local_var_6_8 -> gopurs_runtime.Value
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_4, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_6_8
// TAST (Let): __local_var_7_10 -> gopurs_runtime.Value
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_7_10
// TAST (Let): __local_var_8_12 -> gopurs_runtime.Value
__local_var_8_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_10, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_8_12
// TAST (Let): Functor0_9_13 -> *Constructor_Data_Functor_Functor
Functor0_9_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_13
// TAST (Let): __local_var_10_15 -> gopurs_runtime.Value
__local_var_10_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_15
// TAST (Let): functorWriterT1_10_14 -> gopurs_runtime.Value
functorWriterT1_10_14 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_16 -> gopurs_runtime.Value
__local_var_12_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_15, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_16
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_16, v_13)
})
}))
_ = functorWriterT1_10_14
// TAST (Let): applyWriterT2_8_11 -> gopurs_runtime.Value
applyWriterT2_8_11 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_14
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_12, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_13.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_8_11
// TAST (Let): applicativeWriterT2_7_9 -> gopurs_runtime.Value
applicativeWriterT2_7_9 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_8_11
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_10, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_7_9
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_8_18
// TAST (Let): Apply0_9_19 -> gopurs_runtime.Value
Apply0_9_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_18, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_9_19
// TAST (Let): Functor0_10_20 -> *Constructor_Data_Functor_Functor
Functor0_10_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_19, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_20
// TAST (Let): Functor0_11_22 -> *Constructor_Data_Functor_Functor
Functor0_11_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_19, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_11_22
// TAST (Let): __local_var_12_24 -> gopurs_runtime.Value
__local_var_12_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_19, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_24
// TAST (Let): functorWriterT1_12_23 -> gopurs_runtime.Value
functorWriterT1_12_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_25 -> gopurs_runtime.Value
__local_var_14_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_24, "map"), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_13, (*Constructor_Data_Tuple_Tuple)(v_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_14.UnsafePtr).V1})}
}))
_ = __local_var_14_25
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_25, v_15)
})
}))
_ = functorWriterT1_12_23
// TAST (Let): applyWriterT2_11_21 -> gopurs_runtime.Value
applyWriterT2_11_21 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_12_23
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_9_19, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_22.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_16.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "append"), (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_16.UnsafePtr).V1)})}
})
}), v_13), v1_14)
})
}))
_ = applyWriterT2_11_21
// TAST (Let): bindWriterT2_8_17 -> gopurs_runtime.Value
bindWriterT2_8_17 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_11_21
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_18, "bind"), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_26 -> gopurs_runtime.Value
__local_var_15_26 := (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1
_ = __local_var_15_26
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_20.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_16.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "append"), __local_var_15_26, (*Constructor_Data_Tuple_Tuple)(v3_16.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_13, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_8_17
// TAST (Let): monadWriterT2_6_7 -> gopurs_runtime.Value
monadWriterT2_6_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_7_9
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_8_17
}))
_ = monadWriterT2_6_7
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_6_7
}), gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_5_6)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_4, "throwError"), e_7))
}))
})
_ = monadThrowWriterT1_3_4
return gopurs_runtime.Func(func(dictMonadError_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadThrowWriterT2_5_27 -> gopurs_runtime.Value
monadThrowWriterT2_5_27 := gopurs_runtime.Apply(monadThrowWriterT1_3_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_4, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowWriterT2_5_27
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowWriterT2_5_27
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_4, "catchError"), v_6, gopurs_runtime.Func(func(e_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_7, e_8)
}))
})
}))
})
}

func Call_Control_Monad_Writer_Trans_monadSTWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonadST_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_2 -> gopurs_runtime.Value
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): Functor0_7_8 -> *Constructor_Data_Functor_Functor
Functor0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_8
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorWriterT1_8_9 -> gopurs_runtime.Value
functorWriterT1_8_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_10.UnsafePtr).V1})}
}))
_ = __local_var_10_11
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_11, v_11)
})
}))
_ = functorWriterT1_8_9
// TAST (Let): applyWriterT2_6_6 -> gopurs_runtime.Value
applyWriterT2_6_6 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_8_9
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_8.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_12.UnsafePtr).V1)})}
})
}), v_9), v1_10)
})
}))
_ = applyWriterT2_6_6
// TAST (Let): applicativeWriterT2_5_4 -> gopurs_runtime.Value
applicativeWriterT2_5_4 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_6_6
}), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_5_4
// TAST (Let): __local_var_6_13 -> gopurs_runtime.Value
__local_var_6_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_6_13
// TAST (Let): Apply0_7_14 -> gopurs_runtime.Value
Apply0_7_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_13, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_14
// TAST (Let): Functor0_8_15 -> *Constructor_Data_Functor_Functor
Functor0_8_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_15
// TAST (Let): Functor0_9_17 -> *Constructor_Data_Functor_Functor
Functor0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_17
// TAST (Let): __local_var_10_19 -> gopurs_runtime.Value
__local_var_10_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_19
// TAST (Let): functorWriterT1_10_18 -> gopurs_runtime.Value
functorWriterT1_10_18 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_20 -> gopurs_runtime.Value
__local_var_12_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_19, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_20
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_20, v_13)
})
}))
_ = functorWriterT1_10_18
// TAST (Let): applyWriterT2_9_16 -> gopurs_runtime.Value
applyWriterT2_9_16 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_18
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_14, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_17.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_9_16
// TAST (Let): bindWriterT2_6_12 -> gopurs_runtime.Value
bindWriterT2_6_12 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_9_16
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_13, "bind"), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_21 -> gopurs_runtime.Value
__local_var_13_21 := (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1
_ = __local_var_13_21
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_15.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_13_21, (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_11, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_6_12
// TAST (Let): monadWriterT2_5_3 -> gopurs_runtime.Value
monadWriterT2_5_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_5_4
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_6_12
}))
_ = monadWriterT2_5_3
// TAST (Let): Bind1_6_23 -> *Constructor_Control_Bind_Bind
Bind1_6_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_23
// TAST (Let): pure_7_24 -> gopurs_runtime.Value
pure_7_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_24
// TAST (Let): __local_var_6_22 -> gopurs_runtime.Value
__local_var_6_22 := gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_23.V1), m_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_24, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
_ = __local_var_6_22
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_22, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_3, "liftST"), x_7))
}))
})
}

func Call_Control_Monad_Writer_Trans_monoidWriterT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): Functor0_5_5 -> *Constructor_Data_Functor_Functor
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorWriterT1_6_6 -> gopurs_runtime.Value
functorWriterT1_6_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1})}
}))
_ = __local_var_8_8
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_8, v_9)
})
}))
_ = functorWriterT1_6_6
// TAST (Let): applyWriterT2_4_3 -> gopurs_runtime.Value
applyWriterT2_4_3 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_6_6
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_10.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "append"), (*Constructor_Data_Tuple_Tuple)(v3_9.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_10.UnsafePtr).V1)})}
})
}), v_7), v1_8)
})
}))
_ = applyWriterT2_4_3
// TAST (Let): applicativeWriterT1_3_1 -> *Constructor_Control_Applicative_Applicative
applicativeWriterT1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_4_3
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, gopurs_runtime.RecordGet(dictMonoid_2, "mempty")})})
})))
_ = applicativeWriterT1_3_1
// TAST (Let): __local_var_4_9 -> gopurs_runtime.Value
__local_var_4_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_4_9
// TAST (Let): Functor0_5_11 -> *Constructor_Data_Functor_Functor
Functor0_5_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_11
// TAST (Let): __local_var_6_13 -> gopurs_runtime.Value
__local_var_6_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_13
// TAST (Let): functorWriterT1_6_12 -> gopurs_runtime.Value
functorWriterT1_6_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_14 -> gopurs_runtime.Value
__local_var_8_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_13, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1})}
}))
_ = __local_var_8_14
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_14, v_9)
})
}))
_ = functorWriterT1_6_12
// TAST (Let): applyWriterT1_5_10 -> *Constructor_Control_Apply_Apply
applyWriterT1_5_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_6_12
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_11.V0), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_10.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_9, "append"), (*Constructor_Data_Tuple_Tuple)(v3_9.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_10.UnsafePtr).V1)})}
})
}), v_7), v1_8)
})
})))
_ = applyWriterT1_5_10
return gopurs_runtime.Func(func(dictMonoid1_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_16 -> *Constructor_Data_Functor_Functor
Functor0_7_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyWriterT1_5_10.V0), gopurs_runtime.Value{}))
_ = Functor0_7_16
// TAST (Let): __local_var_8_17 -> gopurs_runtime.Value
__local_var_8_17 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_6, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_8_17
// TAST (Let): semigroupWriterT3_7_15 -> gopurs_runtime.Value
semigroupWriterT3_7_15 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyWriterT1_5_10.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_16.V0), __local_var_8_17, a_9), b_10)
})
}))
_ = semigroupWriterT3_7_15
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupWriterT3_7_15
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeWriterT1_3_1.V1), gopurs_runtime.RecordGet(dictMonoid1_6, "mempty")))
})
})
}

func Call_Control_Monad_Writer_Trans_altWriterT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorWriterT1_1_0 -> gopurs_runtime.Value
functorWriterT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1})}
}))
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, v_4)
})
}))
_ = functorWriterT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_2, v1_3)
})
}))
}

func Call_Control_Monad_Writer_Trans_plusWriterT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorWriterT1_2_2 -> gopurs_runtime.Value
functorWriterT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1})}
}))
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, v_5)
})
}))
_ = functorWriterT1_2_2
// TAST (Let): altWriterT1_1_0 -> gopurs_runtime.Value
altWriterT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "alt"), v_3, v1_4)
})
}))
_ = altWriterT1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_1_0
}), gopurs_runtime.RecordGet(dictPlus_0, "empty"))
}

func Call_Control_Monad_Writer_Trans_alternativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): Functor0_5_5 -> *Constructor_Data_Functor_Functor
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorWriterT1_6_6 -> gopurs_runtime.Value
functorWriterT1_6_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1})}
}))
_ = __local_var_8_8
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_8, v_9)
})
}))
_ = functorWriterT1_6_6
// TAST (Let): applyWriterT2_4_3 -> gopurs_runtime.Value
applyWriterT2_4_3 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_6_6
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_10.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_9.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_10.UnsafePtr).V1)})}
})
}), v_7), v1_8)
})
}))
_ = applyWriterT2_4_3
// TAST (Let): applicativeWriterT2_3_1 -> gopurs_runtime.Value
applicativeWriterT2_3_1 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_4_3
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_3_1
// TAST (Let): __local_var_4_10 -> gopurs_runtime.Value
__local_var_4_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_10
// TAST (Let): __local_var_5_12 -> gopurs_runtime.Value
__local_var_5_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_10, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_12
// TAST (Let): __local_var_6_14 -> gopurs_runtime.Value
__local_var_6_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_14
// TAST (Let): functorWriterT1_6_13 -> gopurs_runtime.Value
functorWriterT1_6_13 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_15 -> gopurs_runtime.Value
__local_var_8_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_14, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1})}
}))
_ = __local_var_8_15
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_15, v_9)
})
}))
_ = functorWriterT1_6_13
// TAST (Let): altWriterT1_5_11 -> gopurs_runtime.Value
altWriterT1_5_11 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_6_13
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_12, "alt"), v_7, v1_8)
})
}))
_ = altWriterT1_5_11
// TAST (Let): plusWriterT1_4_9 -> gopurs_runtime.Value
plusWriterT1_4_9 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_5_11
}), gopurs_runtime.RecordGet(__local_var_4_10, "empty"))
_ = plusWriterT1_4_9
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusWriterT1_4_9
}))
})
}

func Call_Control_Monad_Writer_Trans_monadPlusWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictMonadPlus_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_4, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_6_6
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): Functor0_8_9 -> *Constructor_Data_Functor_Functor
Functor0_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_9
// TAST (Let): __local_var_9_11 -> gopurs_runtime.Value
__local_var_9_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_11
// TAST (Let): functorWriterT1_9_10 -> gopurs_runtime.Value
functorWriterT1_9_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1})}
}))
_ = __local_var_11_12
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_12, v_12)
})
}))
_ = functorWriterT1_9_10
// TAST (Let): applyWriterT2_7_7 -> gopurs_runtime.Value
applyWriterT2_7_7 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_9_10
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_9.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_13.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Data_Tuple_Tuple)(v3_12.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_13.UnsafePtr).V1)})}
})
}), v_10), v1_11)
})
}))
_ = applyWriterT2_7_7
// TAST (Let): applicativeWriterT2_6_5 -> gopurs_runtime.Value
applicativeWriterT2_6_5 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_7_7
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_6_5
// TAST (Let): __local_var_7_14 -> gopurs_runtime.Value
__local_var_7_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_7_14
// TAST (Let): Apply0_8_15 -> gopurs_runtime.Value
Apply0_8_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_14, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_8_15
// TAST (Let): Functor0_9_16 -> *Constructor_Data_Functor_Functor
Functor0_9_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_16
// TAST (Let): Functor0_10_18 -> *Constructor_Data_Functor_Functor
Functor0_10_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_18
// TAST (Let): __local_var_11_20 -> gopurs_runtime.Value
__local_var_11_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_15, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_20
// TAST (Let): functorWriterT1_11_19 -> gopurs_runtime.Value
functorWriterT1_11_19 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_21 -> gopurs_runtime.Value
__local_var_13_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_20, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_13.UnsafePtr).V1})}
}))
_ = __local_var_13_21
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_21, v_14)
})
}))
_ = functorWriterT1_11_19
// TAST (Let): applyWriterT2_10_17 -> gopurs_runtime.Value
applyWriterT2_10_17 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_11_19
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_8_15, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_18.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_15.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Constructor_Data_Tuple_Tuple)(v3_14.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_15.UnsafePtr).V1)})}
})
}), v_12), v1_13)
})
}))
_ = applyWriterT2_10_17
// TAST (Let): bindWriterT2_7_13 -> gopurs_runtime.Value
bindWriterT2_7_13 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_10_17
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_14, "bind"), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_22 -> gopurs_runtime.Value
__local_var_14_22 := (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1
_ = __local_var_14_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_16.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_14_22, (*Constructor_Data_Tuple_Tuple)(v3_15.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_12, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0))
}))
})
}))
_ = bindWriterT2_7_13
// TAST (Let): monadWriterT2_5_3 -> gopurs_runtime.Value
monadWriterT2_5_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_6_5
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_7_13
}))
_ = monadWriterT2_5_3
// TAST (Let): __local_var_6_24 -> gopurs_runtime.Value
__local_var_6_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_4, "Alternative1"), gopurs_runtime.Value{})
_ = __local_var_6_24
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_24, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_7_26
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_26, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_8_28
// TAST (Let): Functor0_9_29 -> *Constructor_Data_Functor_Functor
Functor0_9_29 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_29
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_31
// TAST (Let): functorWriterT1_10_30 -> gopurs_runtime.Value
functorWriterT1_10_30 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_32 -> gopurs_runtime.Value
__local_var_12_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_32
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_32, v_13)
})
}))
_ = functorWriterT1_10_30
// TAST (Let): applyWriterT2_8_27 -> gopurs_runtime.Value
applyWriterT2_8_27 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_30
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_28, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_29.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "append"), (*Constructor_Data_Tuple_Tuple)(v3_13.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v4_14.UnsafePtr).V1)})}
})
}), v_11), v1_12)
})
}))
_ = applyWriterT2_8_27
// TAST (Let): applicativeWriterT2_7_25 -> gopurs_runtime.Value
applicativeWriterT2_7_25 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_8_27
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_26, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
_ = applicativeWriterT2_7_25
// TAST (Let): __local_var_8_34 -> gopurs_runtime.Value
__local_var_8_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_24, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_8_34
// TAST (Let): __local_var_9_36 -> gopurs_runtime.Value
__local_var_9_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_34, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_9_36
// TAST (Let): __local_var_10_38 -> gopurs_runtime.Value
__local_var_10_38 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_36, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_38
// TAST (Let): functorWriterT1_10_37 -> gopurs_runtime.Value
functorWriterT1_10_37 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_39 -> gopurs_runtime.Value
__local_var_12_39 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_38, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_12.UnsafePtr).V1})}
}))
_ = __local_var_12_39
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_39, v_13)
})
}))
_ = functorWriterT1_10_37
// TAST (Let): altWriterT1_9_35 -> gopurs_runtime.Value
altWriterT1_9_35 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_10_37
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_36, "alt"), v_11, v1_12)
})
}))
_ = altWriterT1_9_35
// TAST (Let): plusWriterT1_8_33 -> gopurs_runtime.Value
plusWriterT1_8_33 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_9_35
}), gopurs_runtime.RecordGet(__local_var_8_34, "empty"))
_ = plusWriterT1_8_33
// TAST (Let): alternativeWriterT2_6_23 -> gopurs_runtime.Value
alternativeWriterT2_6_23 := gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_7_25
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return plusWriterT1_8_33
}))
_ = alternativeWriterT2_6_23
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeWriterT2_6_23
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}))
})
}

func Call_Control_Monad_Writer_Trans_mapWriterT__2842489082(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Writer_Trans_mapWriterT__4072164636(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}))
}

func Call_Control_Monad_Writer_Trans_mapWriterT__77717660(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Writer_Trans_runWriterT__4273258459(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return v_0
}


