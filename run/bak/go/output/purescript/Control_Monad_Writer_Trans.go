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
return Call_Control_Monad_Writer_Trans_execWriterT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_Trans_mapWriterT__4072164636(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_Trans_runWriterT__4273258459(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
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
// TAST (Let): Bind1_2_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): pure_3_1 -> gopurs_runtime.Value
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1})}
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, v_3)
})
}))
}

func Call_Control_Monad_Writer_Trans_execWriterT(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
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
// TAST (Let): Functor0_2_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorWriterT1_3_1 -> gopurs_runtime.Value
functorWriterT1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_7.UnsafePtr).V1)})}
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
// TAST (Let): Functor0_3_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
// TAST (Let): applyWriterT2_4_2 -> gopurs_runtime.Value
applyWriterT2_4_2 := Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_0, Apply0_2_0)
_ = applyWriterT2_4_2
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_4_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_1, "bind"), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_3 -> gopurs_runtime.Value
__local_var_8_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1
_ = __local_var_8_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_9.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_8_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_9.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0))
}))
})
}))
}

func Call_Control_Monad_Writer_Trans_semigroupWriterT(dictApply_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): applyWriterT1_2_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
applyWriterT1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_1, dictApply_0))
_ = applyWriterT1_2_0
return gopurs_runtime.Func(func(dictSemigroup1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyWriterT1_2_0.V0), gopurs_runtime.Value{}))
_ = Functor0_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.RecordGet(dictSemigroup1_3, "append")
_ = __local_var_5_2
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyWriterT1_2_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_1.V0), __local_var_5_2, a_6), b_7)
})
}))
})
}

func Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): applyWriterT1_1_0 -> gopurs_runtime.Value
applyWriterT1_1_0 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_applyWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyWriterT1_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_3_1 -> gopurs_runtime.Value
applyWriterT2_3_1 := gopurs_runtime.Apply(applyWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = applyWriterT2_3_1
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_3_1
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
}

func Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): applicativeWriterT1_1_0 -> gopurs_runtime.Value
applicativeWriterT1_1_0 := Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0)
_ = applicativeWriterT1_1_0
// TAST (Let): bindWriterT1_2_1 -> gopurs_runtime.Value
bindWriterT1_2_1 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_bindWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = bindWriterT1_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_4_2 -> gopurs_runtime.Value
applicativeWriterT2_4_2 := gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeWriterT2_4_2
// TAST (Let): bindWriterT2_5_3 -> gopurs_runtime.Value
bindWriterT2_5_3 := gopurs_runtime.Apply(bindWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = bindWriterT2_5_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_5_3
}))
})
}

func Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 -> *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]
monadTransWriterT1_1_0 := &Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
// TAST (Let): monadWriterT1_2_3 -> gopurs_runtime.Value
monadWriterT1_2_3 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_3
return gopurs_runtime.Func(func(dictMonadAsk_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_4_4 -> gopurs_runtime.Value
monadWriterT2_4_4 := gopurs_runtime.Apply(monadWriterT1_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_4_4
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_4
}), gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_3, "ask")))
})
}

func Call_Control_Monad_Writer_Trans_monadReaderWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadAskWriterT1_1_0 -> gopurs_runtime.Value
monadAskWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0)
_ = monadAskWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadReader_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadAskWriterT2_3_1 -> gopurs_runtime.Value
monadAskWriterT2_3_1 := gopurs_runtime.Apply(monadAskWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskWriterT2_3_1
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskWriterT2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "local"), f_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, v_6)
})
}))
})
}

func Call_Control_Monad_Writer_Trans_monadContWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadWriterT1_1_0 -> gopurs_runtime.Value
monadWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadCont_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_3_1 -> gopurs_runtime.Value
monadWriterT2_3_1 := gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_2, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_3_1
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_2, "callCC"), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
}))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadWriterT1_1_0 -> gopurs_runtime.Value
monadWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadEffect_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_3_1 -> gopurs_runtime.Value
Monad0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_1
// TAST (Let): monadWriterT2_4_2 -> gopurs_runtime.Value
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_1_0, Monad0_3_1)
_ = monadWriterT2_4_2
// TAST (Let): Bind1_5_4 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_4
// TAST (Let): pure_6_5 -> gopurs_runtime.Value
pure_6_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_5
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_4.V1), m_7, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
_ = __local_var_5_3
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "liftEffect"), x_6))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadRecWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): monadWriterT1_2_1 -> gopurs_runtime.Value
monadWriterT1_2_1 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadRec_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_2 -> gopurs_runtime.Value
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
// TAST (Let): Bind1_5_3 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
// TAST (Let): Applicative0_6_4 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
// TAST (Let): monadWriterT2_7_5 -> gopurs_runtime.Value
monadWriterT2_7_5 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
_ = monadWriterT2_7_5
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_7_5
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_3, "tailRecM"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_6 -> gopurs_runtime.Value
__local_var_11_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_11_6
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 525585346) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), __local_var_11_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1)})}})}
goto end_branch_9
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 60402430) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), __local_var_11_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1)})}})}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), __t9)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
})
}))
})
}

func Call_Control_Monad_Writer_Trans_monadStateWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 -> *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]
monadTransWriterT1_1_0 := &Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
// TAST (Let): monadWriterT1_2_3 -> gopurs_runtime.Value
monadWriterT1_2_3 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_3
return gopurs_runtime.Func(func(dictMonadState_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_4 -> *Constructor_Control_Monad_Monad[gopurs_runtime.Value]
Monad0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_4_4
// TAST (Let): monadWriterT2_5_5 -> gopurs_runtime.Value
monadWriterT2_5_5 := gopurs_runtime.Apply(monadWriterT1_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_5_5
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_5
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_4_4)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "state"), f_6))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 -> gopurs_runtime.Value
Semigroup0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_1_0
// TAST (Let): monadWriterT1_2_1 -> gopurs_runtime.Value
monadWriterT1_2_1 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_4_2 -> gopurs_runtime.Value
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_2_1, dictMonad_3)
_ = monadWriterT2_4_2
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_4
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), Get_Data_Unit_unit())
_ = __local_var_6_5
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(__local_var_6_5, x_7))
})
_ = __local_var_5_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_1_0
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
}))
})
}

func Call_Control_Monad_Writer_Trans_monadWriterWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTellWriterT1_1_0 -> gopurs_runtime.Value
monadTellWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0)
_ = monadTellWriterT1_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Applicative0_4_2 -> gopurs_runtime.Value
Applicative0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_4_2
// TAST (Let): monadTellWriterT2_5_3 -> gopurs_runtime.Value
monadTellWriterT2_5_3 := gopurs_runtime.Apply(monadTellWriterT1_1_0, dictMonad_2)
_ = monadTellWriterT2_5_3
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellWriterT2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_0
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), v_6, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1})})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), v_6, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1)})})
}))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 -> *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]
monadTransWriterT1_1_0 := &Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
// TAST (Let): monadWriterT1_2_3 -> gopurs_runtime.Value
monadWriterT1_2_3 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_3
return gopurs_runtime.Func(func(dictMonadThrow_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_4 -> *Constructor_Control_Monad_Monad[gopurs_runtime.Value]
Monad0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_4_4
// TAST (Let): monadWriterT2_5_5 -> gopurs_runtime.Value
monadWriterT2_5_5 := gopurs_runtime.Apply(monadWriterT1_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_5_5
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_5
}), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_4_4)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "throwError"), e_6))
}))
})
}

func Call_Control_Monad_Writer_Trans_monadErrorWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadThrowWriterT1_1_0 -> gopurs_runtime.Value
monadThrowWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0)
_ = monadThrowWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadError_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadThrowWriterT2_3_1 -> gopurs_runtime.Value
monadThrowWriterT2_3_1 := gopurs_runtime.Apply(monadThrowWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_2, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowWriterT2_3_1
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowWriterT2_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_2, "catchError"), v_4, gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_5, e_6)
}))
})
}))
})
}

func Call_Control_Monad_Writer_Trans_monadSTWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadWriterT1_1_0 -> gopurs_runtime.Value
monadWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadST_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_3_1 -> gopurs_runtime.Value
Monad0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_1
// TAST (Let): monadWriterT2_4_2 -> gopurs_runtime.Value
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_1_0, Monad0_3_1)
_ = monadWriterT2_4_2
// TAST (Let): Bind1_5_4 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_4
// TAST (Let): pure_6_5 -> gopurs_runtime.Value
pure_6_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_5
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_4.V1), m_7, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
_ = __local_var_5_3
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "liftST"), x_6))
}))
})
}

func Call_Control_Monad_Writer_Trans_monoidWriterT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): semigroupWriterT1_1_0 -> gopurs_runtime.Value
semigroupWriterT1_1_0 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_semigroupWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = semigroupWriterT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT1_3_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
applicativeWriterT1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_2), dictApplicative_0))
_ = applicativeWriterT1_3_1
// TAST (Let): semigroupWriterT2_4_2 -> gopurs_runtime.Value
semigroupWriterT2_4_2 := gopurs_runtime.Apply(semigroupWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupWriterT2_4_2
return gopurs_runtime.Func(func(dictMonoid1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupWriterT3_6_3 -> gopurs_runtime.Value
semigroupWriterT3_6_3 := gopurs_runtime.Apply(semigroupWriterT2_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupWriterT3_6_3
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupWriterT3_6_3
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeWriterT1_3_1.V1), gopurs_runtime.RecordGet(dictMonoid1_5, "mempty")))
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1})}
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
// TAST (Let): applicativeWriterT1_1_0 -> gopurs_runtime.Value
applicativeWriterT1_1_0 := Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0)
_ = applicativeWriterT1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_3_1 -> gopurs_runtime.Value
applicativeWriterT2_3_1 := gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeWriterT2_3_1
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorWriterT1_6_6 -> gopurs_runtime.Value
functorWriterT1_6_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1})}
}))
_ = __local_var_8_8
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_8, v_9)
})
}))
_ = functorWriterT1_6_6
// TAST (Let): altWriterT1_5_4 -> gopurs_runtime.Value
altWriterT1_5_4 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_6_6
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "alt"), v_7, v1_8)
})
}))
_ = altWriterT1_5_4
// TAST (Let): plusWriterT1_4_2 -> gopurs_runtime.Value
plusWriterT1_4_2 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_5_4
}), gopurs_runtime.RecordGet(__local_var_4_3, "empty"))
_ = plusWriterT1_4_2
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusWriterT1_4_2
}))
})
}

func Call_Control_Monad_Writer_Trans_monadPlusWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadWriterT1_1_0 -> gopurs_runtime.Value
monadWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
_ = monadWriterT1_1_0
// TAST (Let): alternativeWriterT1_2_1 -> gopurs_runtime.Value
alternativeWriterT1_2_1 := Call_Control_Monad_Writer_Trans_alternativeWriterT(dictMonoid_0)
_ = alternativeWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadPlus_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_4_2 -> gopurs_runtime.Value
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_4_2
// TAST (Let): alternativeWriterT2_5_3 -> gopurs_runtime.Value
alternativeWriterT2_5_3 := gopurs_runtime.Apply(alternativeWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeWriterT2_5_3
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeWriterT2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
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

func Call_Control_Monad_Writer_Trans_mapWriterT__4072164636(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}))
}

func Call_Control_Monad_Writer_Trans_mapWriterT__77717660(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Writer_Trans_runWriterT__4273258459(v_0_loop *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
var v_0 *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return v_0
}


