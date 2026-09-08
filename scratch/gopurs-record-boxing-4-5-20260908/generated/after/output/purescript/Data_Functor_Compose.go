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
		cache_Data_Functor_Compose_newtypeCompose = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
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
return Call_Data_Functor_Compose_bihoistCompose(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), natF_1_box, natG_2_box, v_3_box)
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
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Compose ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
})}))}
}

func Call_Data_Functor_Compose_functorCompose(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2), v_3)
})}))}
}

func Call_Data_Functor_Compose_eqCompose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
// TAST (Let): eqApp1_3_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g), (TypeVar a)])])
eqApp1_3_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))}, x_3, y_4).IntVal) != (0))
})})
_ = eqApp1_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_3_0)}, v_4, v1_5).IntVal) != (0))
})}))}
}

func Call_Data_Functor_Compose_ordCompose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqCompose1__193435443_1_0 shape=Let(Abs(Abs(Let(LitRecord)))) bindingType=Any
eqCompose1__193435443_1_0 := gopurs_runtime.Func2(func(dictEq11_2 gopurs_runtime.Value, dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqApp1_4_2 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g), (TypeVar a)])])
eqApp1_4_2 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_3))}, x_4, y_5).IntVal) != (0))
})})
_ = eqApp1_4_2
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_4_2)}, v_5, v1_6).IntVal) != (0))
})}))}
})
_ = eqCompose1__193435443_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): eqApp1__193435443_3_4 shape=Let(Abs(LitRecord)) bindingType=Any
eqApp1__193435443_3_4 := gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_5, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_4))}, x_5, y_6).IntVal) != (0))
})}))}
})
_ = eqApp1__193435443_3_4
// TAST (Let): ordApp__193435443_3_3 shape=Let(Abs(Let(LitRecord))) bindingType=Any
ordApp__193435443_3_3 := gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqApp2_5_6 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f), (TypeVar a)])])
eqApp2_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqApp1__193435443_3_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_4, "Eq0"), gopurs_runtime.Value{})))
_ = eqApp2_5_6
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp2_5_6)}
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, x_6, y_7).IntVal)), UnsafePtr: nil}
})}))}
})
_ = ordApp__193435443_3_3
// TAST (Let): eqCompose2__193435443_4_7 shape=App(Other) bindingType=Any
eqCompose2__193435443_4_7 := gopurs_runtime.Apply(eqCompose1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{}))
_ = eqCompose2__193435443_4_7
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordApp1_6_8 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g), (TypeVar a)])])
ordApp1_6_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordApp__193435443_3_3, dictOrd_5))
_ = ordApp1_6_8
// TAST (Let): eqCompose3_7_9 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar a)])])
eqCompose3_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqCompose2__193435443_4_7, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})))
_ = eqCompose3_7_9
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqCompose3_7_9)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordApp1_6_8)}, v_8, v1_9).IntVal)), UnsafePtr: nil}
})}))}
})
})
}

func Call_Data_Functor_Compose_eq1Compose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqApp1_3_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g), (TypeVar a)])])
eqApp1_3_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))}, x_3, y_4).IntVal) != (0))
})})
_ = eqApp1_3_0
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_3_0)}, v_4, v1_5).IntVal) != (0))
})
})}))}
}

func Call_Data_Functor_Compose_ord1Compose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqCompose1__193435443_1_1 shape=Let(Abs(Abs(Let(LitRecord)))) bindingType=Any
eqCompose1__193435443_1_1 := gopurs_runtime.Func2(func(dictEq11_2 gopurs_runtime.Value, dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqApp1_4_3 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g), (TypeVar a)])])
eqApp1_4_3 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_3))}, x_4, y_5).IntVal) != (0))
})})
_ = eqApp1_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_4_3)}, v_5, v1_6).IntVal) != (0))
})}))}
})
_ = eqCompose1__193435443_1_1
// TAST (Let): ordCompose1__193435443_1_0 shape=Let(Abs(Let(Let(Abs(Let(Let(LitRecord))))))) bindingType=Any
ordCompose1__193435443_1_0 := gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_6 shape=App(Other) bindingType=Any
__local_var_3_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_6
// TAST (Let): eqApp1__193435443_3_5 shape=Let(Abs(LitRecord)) bindingType=Any
eqApp1__193435443_3_5 := gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_6, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_4))}, x_5, y_6).IntVal) != (0))
})}))}
})
_ = eqApp1__193435443_3_5
// TAST (Let): ordApp__193435443_3_4 shape=Let(Abs(Let(LitRecord))) bindingType=Any
ordApp__193435443_3_4 := gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqApp2_5_7 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f), (TypeVar a)])])
eqApp2_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqApp1__193435443_3_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_4, "Eq0"), gopurs_runtime.Value{})))
_ = eqApp2_5_7
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp2_5_7)}
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, x_6, y_7).IntVal)), UnsafePtr: nil}
})}))}
})
_ = ordApp__193435443_3_4
// TAST (Let): eqCompose2__193435443_4_8 shape=App(Other) bindingType=Any
eqCompose2__193435443_4_8 := gopurs_runtime.Apply(eqCompose1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{}))
_ = eqCompose2__193435443_4_8
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordApp1_6_9 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g), (TypeVar a)])])
ordApp1_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordApp__193435443_3_4, dictOrd_5))
_ = ordApp1_6_9
// TAST (Let): eqCompose3_7_10 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar a)])])
eqCompose3_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqCompose2__193435443_4_8, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})))
_ = eqCompose3_7_10
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqCompose3_7_10)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordApp1_6_9)}, v_8, v1_9).IntVal)), UnsafePtr: nil}
})}))}
})
})
_ = ordCompose1__193435443_1_0
// TAST (Let): __local_var_2_12 shape=App(Other) bindingType=Any
__local_var_2_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_12
// TAST (Let): eq1Compose1__193435443_2_11 shape=Let(Abs(LitRecord)) bindingType=Any
eq1Compose1__193435443_2_11 := gopurs_runtime.Func(func(dictEq11_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqApp1_5_13 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g), (TypeVar a)])])
eqApp1_5_13 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_3, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_4))}, x_5, y_6).IntVal) != (0))
})})
_ = eqApp1_5_13
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_12, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_5_13)}, v_6, v1_7).IntVal) != (0))
})
})}))}
})
_ = eq1Compose1__193435443_2_11
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordCompose2__193435443_4_14 shape=App(Other) bindingType=Any
ordCompose2__193435443_4_14 := gopurs_runtime.Apply(ordCompose1__193435443_1_0, dictOrd11_3)
_ = ordCompose2__193435443_4_14
// TAST (Let): eq1Compose2_5_15 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
eq1Compose2_5_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](gopurs_runtime.Apply(eq1Compose1__193435443_2_11, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})))
_ = eq1Compose2_5_15
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Compose2_5_15)}
}), gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordCompose2__193435443_4_14, dictOrd_6), "compare")
})}))}
})
}

func Call_Data_Functor_Compose_bihoistCompose(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], natF_1_loop gopurs_runtime.Value, natG_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
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
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorCompose1__193435443_2_1 shape=Let(Abs(LitRecord)) bindingType=Any
functorCompose1__193435443_2_1 := gopurs_runtime.Func(func(dictFunctor1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_3, "map"), f_4), v_5)
})}))}
})
_ = functorCompose1__193435443_2_1
return gopurs_runtime.Func(func(dictApply1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): apply_4_3 shape=Other bindingType=(Func [(TypeApp (TypeVar g) [(Func [(TypeVar a)] (TypeVar b))]), (TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
apply_4_3 := gopurs_runtime.RecordGet(dictApply1_3, "apply")
_ = apply_4_3
// TAST (Let): functorCompose2_5_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose2_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose1__193435443_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_3, "Functor0"), gopurs_runtime.Value{})))
_ = functorCompose2_5_4
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_5_4)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), apply_4_3, v_6), v1_7)
})}))}
})
}

func Call_Data_Functor_Compose_applicativeCompose(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorCompose1__193435443_3_3 shape=Let(Abs(LitRecord)) bindingType=Any
functorCompose1__193435443_3_3 := gopurs_runtime.Func(func(dictFunctor1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_4, "map"), f_5), v_6)
})}))}
})
_ = functorCompose1__193435443_3_3
// TAST (Let): applyCompose1__193435443_1_0 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
applyCompose1__193435443_1_0 := gopurs_runtime.Func(func(dictApply1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): apply_5_5 shape=Other bindingType=(Func [(TypeApp (TypeVar g) [(Func [(TypeVar a)] (TypeVar b))]), (TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
apply_5_5 := gopurs_runtime.RecordGet(dictApply1_4, "apply")
_ = apply_5_5
// TAST (Let): functorCompose2_6_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose2_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose1__193435443_3_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_4, "Functor0"), gopurs_runtime.Value{})))
_ = functorCompose2_6_6
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_6_6)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), apply_5_5, v_7), v1_8)
})}))}
})
_ = applyCompose1__193435443_1_0
return gopurs_runtime.Func(func(dictApplicative1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyCompose2_3_7 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
applyCompose2_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyCompose1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyCompose2_3_7
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyCompose2_3_7)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_2, "pure"), x_4))
})}))}
})
}

func Call_Data_Functor_Compose_altCompose(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorCompose1__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
functorCompose1__193435443_1_0 := gopurs_runtime.Func(func(dictFunctor1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_2, "map"), f_3), v_4)
})}))}
})
_ = functorCompose1__193435443_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCompose2_3_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose2_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose1__193435443_1_0, dictFunctor_2))
_ = functorCompose2_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_3_2)}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_4, v1_5)
})}))}
})
}

func Call_Data_Functor_Compose_plusCompose(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 shape=Other bindingType=(TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])])
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorCompose1__193435443_3_3 shape=Let(Abs(LitRecord)) bindingType=Any
functorCompose1__193435443_3_3 := gopurs_runtime.Func(func(dictFunctor1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_4, "map"), f_5), v_6)
})}))}
})
_ = functorCompose1__193435443_3_3
// TAST (Let): altCompose1__193435443_2_1 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
altCompose1__193435443_2_1 := gopurs_runtime.Func(func(dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCompose2_5_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose2_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose1__193435443_3_3, dictFunctor_4))
_ = functorCompose2_5_5
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_5_5)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "alt"), v_6, v1_7)
})}))}
})
_ = altCompose1__193435443_2_1
return gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): altCompose2_4_6 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
altCompose2_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(altCompose1__193435443_2_1, dictFunctor_3))
_ = altCompose2_4_6
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altCompose2_4_6)}
}), empty_1_0}))}
})
}

func Call_Data_Functor_Compose_alternativeCompose(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): Functor0_3_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_4
// TAST (Let): __local_var_4_6 shape=App(Other) bindingType=Any
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): functorCompose1__193435443_4_5 shape=Let(Abs(LitRecord)) bindingType=Any
functorCompose1__193435443_4_5 := gopurs_runtime.Func(func(dictFunctor1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_6, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_5, "map"), f_6), v_7)
})}))}
})
_ = functorCompose1__193435443_4_5
// TAST (Let): applyCompose1__193435443_2_2 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
applyCompose1__193435443_2_2 := gopurs_runtime.Func(func(dictApply1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): apply_6_7 shape=Other bindingType=(Func [(TypeApp (TypeVar g) [(Func [(TypeVar a)] (TypeVar b))]), (TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
apply_6_7 := gopurs_runtime.RecordGet(dictApply1_5, "apply")
_ = apply_6_7
// TAST (Let): functorCompose2_7_8 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose2_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose1__193435443_4_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_5, "Functor0"), gopurs_runtime.Value{})))
_ = functorCompose2_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_7_8)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_4.V0), apply_6_7, v_8), v1_9)
})}))}
})
_ = applyCompose1__193435443_2_2
// TAST (Let): applicativeCompose1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeCompose1__193435443_1_0 := gopurs_runtime.Func(func(dictApplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyCompose2_4_9 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
applyCompose2_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyCompose1__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyCompose2_4_9
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyCompose2_4_9)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "pure"), x_5))
})}))}
})
_ = applicativeCompose1__193435443_1_0
// TAST (Let): __local_var_2_11 shape=App(Other) bindingType=Any
__local_var_2_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_11
// TAST (Let): empty_3_12 shape=Other bindingType=(TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])])
empty_3_12 := gopurs_runtime.RecordGet(__local_var_2_11, "empty")
_ = empty_3_12
// TAST (Let): __local_var_4_14 shape=App(Other) bindingType=Any
__local_var_4_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_11, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_14
// TAST (Let): __local_var_5_16 shape=App(Other) bindingType=Any
__local_var_5_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_16
// TAST (Let): functorCompose1__193435443_5_15 shape=Let(Abs(LitRecord)) bindingType=Any
functorCompose1__193435443_5_15 := gopurs_runtime.Func(func(dictFunctor1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_16, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_6, "map"), f_7), v_8)
})}))}
})
_ = functorCompose1__193435443_5_15
// TAST (Let): altCompose1__193435443_4_13 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
altCompose1__193435443_4_13 := gopurs_runtime.Func(func(dictFunctor_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCompose2_7_17 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose2_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose1__193435443_5_15, dictFunctor_6))
_ = functorCompose2_7_17
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_7_17)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_14, "alt"), v_8, v1_9)
})}))}
})
_ = altCompose1__193435443_4_13
// TAST (Let): plusCompose1__193435443_2_10 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
plusCompose1__193435443_2_10 := gopurs_runtime.Func(func(dictFunctor_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): altCompose2_6_18 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
altCompose2_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(altCompose1__193435443_4_13, dictFunctor_5))
_ = altCompose2_6_18
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altCompose2_6_18)}
}), empty_3_12}))}
})
_ = plusCompose1__193435443_2_10
return gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeCompose2_4_19 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
applicativeCompose2_4_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeCompose1__193435443_1_0, dictApplicative_3))
_ = applicativeCompose2_4_19
// TAST (Let): plusCompose2_5_20 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
plusCompose2_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(plusCompose1__193435443_2_10, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
_ = plusCompose2_5_20
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeCompose2_4_19)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusCompose2_5_20)}
})}))}
})
}


