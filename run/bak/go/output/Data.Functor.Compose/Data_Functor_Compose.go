package Data_Functor_Compose

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Compose gopurs_runtime.Value
var once_Compose sync.Once
func Get_Compose() gopurs_runtime.Value {
	once_Compose.Do(func() {
		cache_Compose = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Compose(x_0_box)
})
	})
	return cache_Compose
}

var cache_showCompose gopurs_runtime.Value
var once_showCompose sync.Once
func Get_showCompose() gopurs_runtime.Value {
	once_showCompose.Do(func() {
		cache_showCompose = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showCompose(dictShow_0_box)
})
	})
	return cache_showCompose
}

var cache_newtypeCompose gopurs_runtime.Value
var once_newtypeCompose sync.Once
func Get_newtypeCompose() gopurs_runtime.Value {
	once_newtypeCompose.Do(func() {
		cache_newtypeCompose = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCompose
}

var cache_functorCompose gopurs_runtime.Value
var once_functorCompose sync.Once
func Get_functorCompose() gopurs_runtime.Value {
	once_functorCompose.Do(func() {
		cache_functorCompose = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorCompose(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_functorCompose
}

var cache_eqCompose gopurs_runtime.Value
var once_eqCompose sync.Once
func Get_eqCompose() gopurs_runtime.Value {
	once_eqCompose.Do(func() {
		cache_eqCompose = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqCompose(dictEq1_0_box, dictEq11_1_box, dictEq_2_box)
})
	})
	return cache_eqCompose
}

var cache_ordCompose gopurs_runtime.Value
var once_ordCompose sync.Once
func Get_ordCompose() gopurs_runtime.Value {
	once_ordCompose.Do(func() {
		cache_ordCompose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordCompose(dictOrd1_0_box)
})
	})
	return cache_ordCompose
}

var cache_eq1Compose gopurs_runtime.Value
var once_eq1Compose sync.Once
func Get_eq1Compose() gopurs_runtime.Value {
	once_eq1Compose.Do(func() {
		cache_eq1Compose = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Compose(dictEq1_0_box, dictEq11_1_box)
})
	})
	return cache_eq1Compose
}

var cache_ord1Compose gopurs_runtime.Value
var once_ord1Compose sync.Once
func Get_ord1Compose() gopurs_runtime.Value {
	once_ord1Compose.Do(func() {
		cache_ord1Compose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Compose(dictOrd1_0_box)
})
	})
	return cache_ord1Compose
}

var cache_bihoistCompose gopurs_runtime.Value
var once_bihoistCompose sync.Once
func Get_bihoistCompose() gopurs_runtime.Value {
	once_bihoistCompose.Do(func() {
		cache_bihoistCompose = gopurs_runtime.Func4(func(dictFunctor_0_box gopurs_runtime.Value, natF_1_box gopurs_runtime.Value, natG_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bihoistCompose(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), natF_1_box, natG_2_box, v_3_box)
})
	})
	return cache_bihoistCompose
}

var cache_applyCompose gopurs_runtime.Value
var once_applyCompose sync.Once
func Get_applyCompose() gopurs_runtime.Value {
	once_applyCompose.Do(func() {
		cache_applyCompose = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyCompose(dictApply_0_box)
})
	})
	return cache_applyCompose
}

var cache_applicativeCompose gopurs_runtime.Value
var once_applicativeCompose sync.Once
func Get_applicativeCompose() gopurs_runtime.Value {
	once_applicativeCompose.Do(func() {
		cache_applicativeCompose = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeCompose(dictApplicative_0_box)
})
	})
	return cache_applicativeCompose
}

var cache_altCompose gopurs_runtime.Value
var once_altCompose sync.Once
func Get_altCompose() gopurs_runtime.Value {
	once_altCompose.Do(func() {
		cache_altCompose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altCompose(dictAlt_0_box)
})
	})
	return cache_altCompose
}

var cache_plusCompose gopurs_runtime.Value
var once_plusCompose sync.Once
func Get_plusCompose() gopurs_runtime.Value {
	once_plusCompose.Do(func() {
		cache_plusCompose = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusCompose(dictPlus_0_box)
})
	})
	return cache_plusCompose
}

var cache_alternativeCompose gopurs_runtime.Value
var once_alternativeCompose sync.Once
func Get_alternativeCompose() gopurs_runtime.Value {
	once_alternativeCompose.Do(func() {
		cache_alternativeCompose = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeCompose(dictAlternative_0_box)
})
	})
	return cache_alternativeCompose
}

var cache_alt__267341625 gopurs_runtime.Value
var once_alt__267341625 sync.Once
func Get_alt__267341625() gopurs_runtime.Value {
	once_alt__267341625.Do(func() {
		cache_alt__267341625 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__267341625(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__267341625
}

var cache_alt__3108117974 gopurs_runtime.Value
var once_alt__3108117974 sync.Once
func Get_alt__3108117974() gopurs_runtime.Value {
	once_alt__3108117974.Do(func() {
		cache_alt__3108117974 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__3108117974(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__3108117974
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__75278732 gopurs_runtime.Value
var once_apply__75278732 sync.Once
func Get_apply__75278732() gopurs_runtime.Value {
	once_apply__75278732.Do(func() {
		cache_apply__75278732 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__75278732(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__75278732
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_eq1__1773593252 gopurs_runtime.Value
var once_eq1__1773593252 sync.Once
func Get_eq1__1773593252() gopurs_runtime.Value {
	once_eq1__1773593252.Do(func() {
		cache_eq1__1773593252 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1__1773593252(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq1__1773593252
}

var cache_eq1__3199040333 gopurs_runtime.Value
var once_eq1__3199040333 sync.Once
func Get_eq1__3199040333() gopurs_runtime.Value {
	once_eq1__3199040333.Do(func() {
		cache_eq1__3199040333 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1__3199040333(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq1__3199040333
}

var cache_hoistLiftApp__1666147871 gopurs_runtime.Value
var once_hoistLiftApp__1666147871 sync.Once
func Get_hoistLiftApp__1666147871() gopurs_runtime.Value {
	once_hoistLiftApp__1666147871.Do(func() {
		cache_hoistLiftApp__1666147871 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_hoistLiftApp__1666147871
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__4258206196 gopurs_runtime.Value
var once_map__4258206196 sync.Once
func Get_map__4258206196() gopurs_runtime.Value {
	once_map__4258206196.Do(func() {
		cache_map__4258206196 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__4258206196(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__4258206196
}

var cache_map__1974414836 gopurs_runtime.Value
var once_map__1974414836 sync.Once
func Get_map__1974414836() gopurs_runtime.Value {
	once_map__1974414836.Do(func() {
		cache_map__1974414836 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1974414836(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1974414836
}

var cache_compare1__650153534 gopurs_runtime.Value
var once_compare1__650153534 sync.Once
func Get_compare1__650153534() gopurs_runtime.Value {
	once_compare1__650153534.Do(func() {
		cache_compare1__650153534 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare1__650153534(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare1__650153534
}

var cache_compare1__3498430039 gopurs_runtime.Value
var once_compare1__3498430039 sync.Once
func Get_compare1__3498430039() gopurs_runtime.Value {
	once_compare1__3498430039.Do(func() {
		cache_compare1__3498430039 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare1__3498430039(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare1__3498430039
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__3448840338 gopurs_runtime.Value
var once_show__3448840338 sync.Once
func Get_show__3448840338() gopurs_runtime.Value {
	once_show__3448840338.Do(func() {
		cache_show__3448840338 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3448840338(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__3448840338
}

func Call_Compose(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showCompose(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Compose "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_functorCompose(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_eqCompose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
eqApp1_3_0 := &pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_2))}, x_3, y_4).IntVal) != (0))
})
})}
_ = eqApp1_3_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp1_3_0)}, v_4, v1_5).IntVal) != (0))
})
}))
}

func Call_ordCompose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_6_5
eqApp2_6_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](__local_var_6_5))}, x_7, y_8).IntVal) != (0))
})
}))
_ = eqApp2_6_4
ordApp1_6_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_6_4
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_5))}, x_7, y_8).IntVal)), UnsafePtr: nil}
})
})))
_ = ordApp1_6_3
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_7_7
eqApp1_8_8 := &pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](__local_var_7_7))}, x_8, y_9).IntVal) != (0))
})
})}
_ = eqApp1_8_8
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

func Call_eq1Compose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqApp1_3_0 := &pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_2))}, x_3, y_4).IntVal) != (0))
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

func Call_ord1Compose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
ordCompose1_1_0 := Call_ordCompose(dictOrd1_0)
_ = ordCompose1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordCompose2_4_2 := gopurs_runtime.Apply(ordCompose1_1_0, dictOrd11_3)
_ = ordCompose2_4_2
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_5_4
eq1Compose2_5_3 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_6 gopurs_runtime.Value) gopurs_runtime.Value {
eqApp1_7_5 := &pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_4, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_6))}, x_7, y_8).IntVal) != (0))
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

func Call_bihoistCompose(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], natF_1_loop gopurs_runtime.Value, natG_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var natF_1 gopurs_runtime.Value = natF_1_loop
_ = natF_1
var natG_2 gopurs_runtime.Value = natG_2_loop
_ = natG_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
return gopurs_runtime.Apply(natF_1, gopurs_runtime.Apply2(dictFunctor_0.V0, natG_2, v_3))
}

func Call_applyCompose(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictApply1_3 gopurs_runtime.Value) gopurs_runtime.Value {
apply_4_2 := gopurs_runtime.RecordGet(dictApply1_3, "apply")
_ = apply_4_2
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(Functor0_1_0.V0, apply_4_2, v_6), v1_7)
})
}))
})
}

func Call_applicativeCompose(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictApplicative1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_4, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_4
apply_6_5 := gopurs_runtime.RecordGet(__local_var_5_4, "apply")
_ = apply_6_5
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
functorCompose2_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "map"), f_8), v_9)
})
}))
_ = functorCompose2_7_6
applyCompose2_5_3 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(Functor0_2_1.V0, apply_6_5, v_8), v1_9)
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

func Call_altCompose(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_plusCompose(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_4, "map"), f_5), v_6)
})
}))
_ = functorCompose2_5_4
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

func Call_alternativeCompose(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
applicativeCompose1_1_0 := Call_applicativeCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeCompose1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_1
empty_3_2 := gopurs_runtime.RecordGet(__local_var_2_1, "empty")
_ = empty_3_2
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_4
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
plusCompose1_4_3 := gopurs_runtime.Func(func(dictFunctor_6 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_7_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_6, "map"), f_7), v_8)
})
}))
_ = functorCompose2_7_7
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
applicativeCompose2_6_8 := gopurs_runtime.Apply(applicativeCompose1_1_0, dictApplicative_5)
_ = applicativeCompose2_6_8
plusCompose2_7_9 := gopurs_runtime.Apply(plusCompose1_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = plusCompose2_7_9
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeCompose2_6_8
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return plusCompose2_7_9
}))
})
}

func Call_alt__267341625(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_alt__3108117974(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__75278732(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq1__1773593252(dict_0_loop *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq1__3199040333(dict_0_loop *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__4258206196(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1974414836(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compare1__650153534(dict_0_loop *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare1__3498430039(dict_0_loop *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__3448840338(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


