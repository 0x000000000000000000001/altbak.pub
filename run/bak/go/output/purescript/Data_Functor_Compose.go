package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Compose_Compose gopurs_runtime.Value
var once_Data_Functor_Compose_Compose sync.Once
func Get_Data_Functor_Compose_Compose() gopurs_runtime.Value {
	once_Data_Functor_Compose_Compose.Do(func() {
		cache_Data_Functor_Compose_Compose = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_Compose(x_0_box)
})
	})
	return cache_Data_Functor_Compose_Compose
}

var cache_Data_Functor_Compose_showCompose gopurs_runtime.Value
var once_Data_Functor_Compose_showCompose sync.Once
func Get_Data_Functor_Compose_showCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_showCompose.Do(func() {
		cache_Data_Functor_Compose_showCompose = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_showCompose(dictShow_0_box)
})
	})
	return cache_Data_Functor_Compose_showCompose
}

var cache_Data_Functor_Compose_newtypeCompose gopurs_runtime.Value
var once_Data_Functor_Compose_newtypeCompose sync.Once
func Get_Data_Functor_Compose_newtypeCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_newtypeCompose.Do(func() {
		cache_Data_Functor_Compose_newtypeCompose = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Functor_Compose_newtypeCompose
}

var cache_Data_Functor_Compose_functorCompose gopurs_runtime.Value
var once_Data_Functor_Compose_functorCompose sync.Once
func Get_Data_Functor_Compose_functorCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_functorCompose.Do(func() {
		cache_Data_Functor_Compose_functorCompose = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_functorCompose(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_Data_Functor_Compose_functorCompose
}

var cache_Data_Functor_Compose_eqCompose gopurs_runtime.Value
var once_Data_Functor_Compose_eqCompose sync.Once
func Get_Data_Functor_Compose_eqCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_eqCompose.Do(func() {
		cache_Data_Functor_Compose_eqCompose = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_eqCompose(dictEq1_0_box, dictEq11_1_box, dictEq_2_box)
})
	})
	return cache_Data_Functor_Compose_eqCompose
}

var cache_Data_Functor_Compose_ordCompose gopurs_runtime.Value
var once_Data_Functor_Compose_ordCompose sync.Once
func Get_Data_Functor_Compose_ordCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_ordCompose.Do(func() {
		cache_Data_Functor_Compose_ordCompose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_ordCompose(dictOrd1_0_box)
})
	})
	return cache_Data_Functor_Compose_ordCompose
}

var cache_Data_Functor_Compose_eq1Compose gopurs_runtime.Value
var once_Data_Functor_Compose_eq1Compose sync.Once
func Get_Data_Functor_Compose_eq1Compose() gopurs_runtime.Value {
	once_Data_Functor_Compose_eq1Compose.Do(func() {
		cache_Data_Functor_Compose_eq1Compose = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_eq1Compose(dictEq1_0_box, dictEq11_1_box)
})
	})
	return cache_Data_Functor_Compose_eq1Compose
}

var cache_Data_Functor_Compose_ord1Compose gopurs_runtime.Value
var once_Data_Functor_Compose_ord1Compose sync.Once
func Get_Data_Functor_Compose_ord1Compose() gopurs_runtime.Value {
	once_Data_Functor_Compose_ord1Compose.Do(func() {
		cache_Data_Functor_Compose_ord1Compose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_ord1Compose(dictOrd1_0_box)
})
	})
	return cache_Data_Functor_Compose_ord1Compose
}

var cache_Data_Functor_Compose_bihoistCompose gopurs_runtime.Value
var once_Data_Functor_Compose_bihoistCompose sync.Once
func Get_Data_Functor_Compose_bihoistCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_bihoistCompose.Do(func() {
		cache_Data_Functor_Compose_bihoistCompose = gopurs_runtime.Func4(func(dictFunctor_0_box gopurs_runtime.Value, natF_1_box gopurs_runtime.Value, natG_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_bihoistCompose(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), natF_1_box, natG_2_box, v_3_box)
})
	})
	return cache_Data_Functor_Compose_bihoistCompose
}

var cache_Data_Functor_Compose_applyCompose gopurs_runtime.Value
var once_Data_Functor_Compose_applyCompose sync.Once
func Get_Data_Functor_Compose_applyCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_applyCompose.Do(func() {
		cache_Data_Functor_Compose_applyCompose = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_applyCompose(dictApply_0_box)
})
	})
	return cache_Data_Functor_Compose_applyCompose
}

var cache_Data_Functor_Compose_applicativeCompose gopurs_runtime.Value
var once_Data_Functor_Compose_applicativeCompose sync.Once
func Get_Data_Functor_Compose_applicativeCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_applicativeCompose.Do(func() {
		cache_Data_Functor_Compose_applicativeCompose = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_applicativeCompose(dictApplicative_0_box)
})
	})
	return cache_Data_Functor_Compose_applicativeCompose
}

var cache_Data_Functor_Compose_altCompose gopurs_runtime.Value
var once_Data_Functor_Compose_altCompose sync.Once
func Get_Data_Functor_Compose_altCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_altCompose.Do(func() {
		cache_Data_Functor_Compose_altCompose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_altCompose(dictAlt_0_box)
})
	})
	return cache_Data_Functor_Compose_altCompose
}

var cache_Data_Functor_Compose_plusCompose gopurs_runtime.Value
var once_Data_Functor_Compose_plusCompose sync.Once
func Get_Data_Functor_Compose_plusCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_plusCompose.Do(func() {
		cache_Data_Functor_Compose_plusCompose = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_plusCompose(dictPlus_0_box)
})
	})
	return cache_Data_Functor_Compose_plusCompose
}

var cache_Data_Functor_Compose_alternativeCompose gopurs_runtime.Value
var once_Data_Functor_Compose_alternativeCompose sync.Once
func Get_Data_Functor_Compose_alternativeCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_alternativeCompose.Do(func() {
		cache_Data_Functor_Compose_alternativeCompose = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Compose_alternativeCompose(dictAlternative_0_box)
})
	})
	return cache_Data_Functor_Compose_alternativeCompose
}

func Call_Data_Functor_Compose_Compose(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Compose_showCompose(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Compose ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Functor_Compose_functorCompose(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2), v_3)
})
}))
}

func Call_Data_Functor_Compose_eqCompose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
// TAST (Let): eqApp1_3_0 -> *Constructor_Data_Eq_Eq
eqApp1_3_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_2))}, x_3, y_4).IntVal) != (0))
})
})}
_ = eqApp1_3_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_3_0)}, v_4, v1_5).IntVal) != (0))
})
}))
}

func Call_Data_Functor_Compose_ordCompose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): eqApp2_6_4 -> gopurs_runtime.Value
eqApp2_6_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](__local_var_6_5))}, x_7, y_8).IntVal) != (0))
})
}))
_ = eqApp2_6_4
// TAST (Let): ordApp1_6_3 -> *Constructor_Data_Ord_Ord
ordApp1_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_6_4
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_5))}, x_7, y_8).IntVal)), UnsafePtr: nil}
})
})))
_ = ordApp1_6_3
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): eqApp1_8_8 -> *Constructor_Data_Eq_Eq
eqApp1_8_8 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](__local_var_7_7))}, x_8, y_9).IntVal) != (0))
})
})}
_ = eqApp1_8_8
// TAST (Let): eqCompose3_7_6 -> gopurs_runtime.Value
eqCompose3_7_6 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_8_8)}, v_9, v1_10).IntVal) != (0))
})
}))
_ = eqCompose3_7_6
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCompose3_7_6
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordApp1_6_3)}, v_8, v1_9).IntVal)), UnsafePtr: nil}
})
}))
})
})
}

func Call_Data_Functor_Compose_eq1Compose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqApp1_3_0 -> *Constructor_Data_Eq_Eq
eqApp1_3_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_2))}, x_3, y_4).IntVal) != (0))
})
})}
_ = eqApp1_3_0
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_3_0)}, v_4, v1_5).IntVal) != (0))
})
})
}))
}

func Call_Data_Functor_Compose_ord1Compose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): ordCompose1_1_0 -> gopurs_runtime.Value
ordCompose1_1_0 := Call_Data_Functor_Compose_ordCompose(dictOrd1_0)
_ = ordCompose1_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordCompose2_4_2 -> gopurs_runtime.Value
ordCompose2_4_2 := gopurs_runtime.Apply(ordCompose1_1_0, dictOrd11_3)
_ = ordCompose2_4_2
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): eq1Compose2_5_3 -> gopurs_runtime.Value
eq1Compose2_5_3 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqApp1_7_5 -> *Constructor_Data_Eq_Eq
eqApp1_7_5 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_4, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_6))}, x_7, y_8).IntVal) != (0))
})
})}
_ = eqApp1_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_7_5)}, v_8, v1_9).IntVal) != (0))
})
})
}))
_ = eq1Compose2_5_3
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Compose2_5_3
}), gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordCompose2_4_2, dictOrd_6), "compare")
}))
})
}

func Call_Data_Functor_Compose_bihoistCompose(dictFunctor_0_loop *Constructor_Data_Functor_Functor, natF_1_loop gopurs_runtime.Value, natG_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var natF_1 gopurs_runtime.Value = natF_1_loop
_ = natF_1
var natG_2 gopurs_runtime.Value = natG_2_loop
_ = natG_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
return gopurs_runtime.Apply(natF_1, gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), natG_2, v_3))
}

func Call_Data_Functor_Compose_applyCompose(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictApply1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): apply_4_2 -> gopurs_runtime.Value
apply_4_2 := gopurs_runtime.RecordGet(dictApply1_3, "apply")
_ = apply_4_2
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorCompose2_5_3 -> gopurs_runtime.Value
functorCompose2_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), f_6), v_7)
})
}))
_ = functorCompose2_5_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), apply_4_2, v_6), v1_7)
})
}))
})
}

func Call_Data_Functor_Compose_applicativeCompose(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictApplicative1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_4, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): apply_6_5 -> gopurs_runtime.Value
apply_6_5 := gopurs_runtime.RecordGet(__local_var_5_4, "apply")
_ = apply_6_5
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorCompose2_7_6 -> gopurs_runtime.Value
functorCompose2_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "map"), f_8), v_9)
})
}))
_ = functorCompose2_7_6
// TAST (Let): applyCompose2_5_3 -> gopurs_runtime.Value
applyCompose2_5_3 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), apply_6_5, v_8), v1_9)
})
}))
_ = applyCompose2_5_3
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose2_5_3
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_4, "pure"), x_6))
}))
})
}

func Call_Data_Functor_Compose_altCompose(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCompose2_3_1 -> gopurs_runtime.Value
functorCompose2_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3), v_4)
})
}))
_ = functorCompose2_3_1
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_4, v1_5)
})
}))
})
}

func Call_Data_Functor_Compose_plusCompose(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 -> gopurs_runtime.Value
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCompose2_5_4 -> gopurs_runtime.Value
functorCompose2_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_4, "map"), f_5), v_6)
})
}))
_ = functorCompose2_5_4
// TAST (Let): altCompose2_5_3 -> gopurs_runtime.Value
altCompose2_5_3 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "alt"), v_6, v1_7)
})
}))
_ = altCompose2_5_3
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_5_3
}), empty_1_0)
})
}

func Call_Data_Functor_Compose_alternativeCompose(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): applicativeCompose1_1_0 -> gopurs_runtime.Value
applicativeCompose1_1_0 := Call_Data_Functor_Compose_applicativeCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeCompose1_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): empty_3_2 -> gopurs_runtime.Value
empty_3_2 := gopurs_runtime.RecordGet(__local_var_2_1, "empty")
_ = empty_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): plusCompose1_4_3 -> gopurs_runtime.Value
plusCompose1_4_3 := gopurs_runtime.Func(func(dictFunctor_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCompose2_7_7 -> gopurs_runtime.Value
functorCompose2_7_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_6, "map"), f_7), v_8)
})
}))
_ = functorCompose2_7_7
// TAST (Let): altCompose2_7_6 -> gopurs_runtime.Value
altCompose2_7_6 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_7
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "alt"), v_8, v1_9)
})
}))
_ = altCompose2_7_6
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_7_6
}), empty_3_2)
})
_ = plusCompose1_4_3
return gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeCompose2_6_8 -> gopurs_runtime.Value
applicativeCompose2_6_8 := gopurs_runtime.Apply(applicativeCompose1_1_0, dictApplicative_5)
_ = applicativeCompose2_6_8
// TAST (Let): plusCompose2_7_9 -> gopurs_runtime.Value
plusCompose2_7_9 := gopurs_runtime.Apply(plusCompose1_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = plusCompose2_7_9
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeCompose2_6_8
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return plusCompose2_7_9
}))
})
}


