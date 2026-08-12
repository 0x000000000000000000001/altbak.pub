package Data_Functor_Coproduct

import (
	pkg_Control_Extend "gopurs/output/Control.Extend"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Coproduct gopurs_runtime.Value
var once_Coproduct sync.Once
func Get_Coproduct() gopurs_runtime.Value {
	once_Coproduct.Do(func() {
		cache_Coproduct = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Coproduct(x_0_box)
})
	})
	return cache_Coproduct
}

var cache_showCoproduct gopurs_runtime.Value
var once_showCoproduct sync.Once
func Get_showCoproduct() gopurs_runtime.Value {
	once_showCoproduct.Do(func() {
		cache_showCoproduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showCoproduct(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showCoproduct
}

var cache_right gopurs_runtime.Value
var once_right sync.Once
func Get_right() gopurs_runtime.Value {
	once_right.Do(func() {
		cache_right = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right(ga_0_box)
})
	})
	return cache_right
}

var cache_newtypeCoproduct gopurs_runtime.Value
var once_newtypeCoproduct sync.Once
func Get_newtypeCoproduct() gopurs_runtime.Value {
	once_newtypeCoproduct.Do(func() {
		cache_newtypeCoproduct = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCoproduct
}

var cache_left gopurs_runtime.Value
var once_left sync.Once
func Get_left() gopurs_runtime.Value {
	once_left.Do(func() {
		cache_left = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_left(fa_0_box)
})
	})
	return cache_left
}

var cache_functorCoproduct gopurs_runtime.Value
var once_functorCoproduct sync.Once
func Get_functorCoproduct() gopurs_runtime.Value {
	once_functorCoproduct.Do(func() {
		cache_functorCoproduct = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorCoproduct(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_functorCoproduct
}

var cache_eq1Coproduct gopurs_runtime.Value
var once_eq1Coproduct sync.Once
func Get_eq1Coproduct() gopurs_runtime.Value {
	once_eq1Coproduct.Do(func() {
		cache_eq1Coproduct = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Coproduct(dictEq1_0_box, dictEq11_1_box)
})
	})
	return cache_eq1Coproduct
}

var cache_eqCoproduct gopurs_runtime.Value
var once_eqCoproduct sync.Once
func Get_eqCoproduct() gopurs_runtime.Value {
	once_eqCoproduct.Do(func() {
		cache_eqCoproduct = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqCoproduct(dictEq1_0_box, dictEq11_1_box, dictEq_2_box)
})
	})
	return cache_eqCoproduct
}

var cache_ord1Coproduct gopurs_runtime.Value
var once_ord1Coproduct sync.Once
func Get_ord1Coproduct() gopurs_runtime.Value {
	once_ord1Coproduct.Do(func() {
		cache_ord1Coproduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Coproduct(dictOrd1_0_box)
})
	})
	return cache_ord1Coproduct
}

var cache_ordCoproduct gopurs_runtime.Value
var once_ordCoproduct sync.Once
func Get_ordCoproduct() gopurs_runtime.Value {
	once_ordCoproduct.Do(func() {
		cache_ordCoproduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordCoproduct(dictOrd1_0_box)
})
	})
	return cache_ordCoproduct
}

var cache_coproduct gopurs_runtime.Value
var once_coproduct sync.Once
func Get_coproduct() gopurs_runtime.Value {
	once_coproduct.Do(func() {
		cache_coproduct = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct
}

var cache_extendCoproduct gopurs_runtime.Value
var once_extendCoproduct sync.Once
func Get_extendCoproduct() gopurs_runtime.Value {
	once_extendCoproduct.Do(func() {
		cache_extendCoproduct = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendCoproduct(dictExtend_0_box)
})
	})
	return cache_extendCoproduct
}

var cache_comonadCoproduct gopurs_runtime.Value
var once_comonadCoproduct sync.Once
func Get_comonadCoproduct() gopurs_runtime.Value {
	once_comonadCoproduct.Do(func() {
		cache_comonadCoproduct = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadCoproduct(dictComonad_0_box)
})
	})
	return cache_comonadCoproduct
}

var cache_bihoistCoproduct gopurs_runtime.Value
var once_bihoistCoproduct sync.Once
func Get_bihoistCoproduct() gopurs_runtime.Value {
	once_bihoistCoproduct.Do(func() {
		cache_bihoistCoproduct = gopurs_runtime.Func3(func(natF_0_box gopurs_runtime.Value, natG_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bihoistCoproduct(natF_0_box, natG_1_box, v_2_box)
})
	})
	return cache_bihoistCoproduct
}

var cache_extend__1264481661 gopurs_runtime.Value
var once_extend__1264481661 sync.Once
func Get_extend__1264481661() gopurs_runtime.Value {
	once_extend__1264481661.Do(func() {
		cache_extend__1264481661 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__1264481661(gopurs_runtime.CoerceToStruct[pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extend__1264481661
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

var cache_bifunctorEither__3558063994 gopurs_runtime.Value
var once_bifunctorEither__3558063994 sync.Once
func Get_bifunctorEither__3558063994() gopurs_runtime.Value {
	once_bifunctorEither__3558063994.Do(func() {
		cache_bifunctorEither__3558063994 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
}))
	})
	return cache_bifunctorEither__3558063994
}

var cache_bimap__132457202 gopurs_runtime.Value
var once_bimap__132457202 sync.Once
func Get_bimap__132457202() gopurs_runtime.Value {
	once_bimap__132457202.Do(func() {
		cache_bimap__132457202 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__132457202(gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__132457202
}

var cache_bimap__1783967194 gopurs_runtime.Value
var once_bimap__1783967194 sync.Once
func Get_bimap__1783967194() gopurs_runtime.Value {
	once_bimap__1783967194.Do(func() {
		cache_bimap__1783967194 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__1783967194(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_bimap__1783967194
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

var cache_coproduct__79520197 gopurs_runtime.Value
var once_coproduct__79520197 sync.Once
func Get_coproduct__79520197() gopurs_runtime.Value {
	once_coproduct__79520197.Do(func() {
		cache_coproduct__79520197 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct__79520197(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct__79520197
}

var cache_coproduct__413515331 gopurs_runtime.Value
var once_coproduct__413515331 sync.Once
func Get_coproduct__413515331() gopurs_runtime.Value {
	once_coproduct__413515331.Do(func() {
		cache_coproduct__413515331 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct__413515331(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct__413515331
}

var cache_coproduct__829064685 gopurs_runtime.Value
var once_coproduct__829064685 sync.Once
func Get_coproduct__829064685() gopurs_runtime.Value {
	once_coproduct__829064685.Do(func() {
		cache_coproduct__829064685 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct__829064685(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct__829064685
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

var cache_show__3978978930 gopurs_runtime.Value
var once_show__3978978930 sync.Once
func Get_show__3978978930() gopurs_runtime.Value {
	once_show__3978978930.Do(func() {
		cache_show__3978978930 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3978978930(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__3978978930
}

func Call_Coproduct(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showCoproduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(left "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(right "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
}

func Call_right(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_left(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0})}
}

func Call_functorCoproduct(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorEither(), "bimap"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2), v_3)
})
}))
}

func Call_eq1Coproduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (v_3.Type == 9 && v_3.IntVal == 3711209382) {
var __t0 bool
{
if (v1_4.Type == 9 && v1_4.IntVal == 3711209382) {
__t0 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_2))}, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 2465973597)) && ((v1_4.Type == 9 && v1_4.IntVal == 2465973597)) {
__t1 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_2))}, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
})
}))
}

func Call_eqCoproduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (v_3.Type == 9 && v_3.IntVal == 3711209382) {
var __t0 bool
{
if (v1_4.Type == 9 && v1_4.IntVal == 3711209382) {
__t0 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_2))}, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 2465973597)) && ((v1_4.Type == 9 && v1_4.IntVal == 2465973597)) {
__t1 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_2))}, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
}))
}

func Call_ord1Coproduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_2
eq1Coproduct2_3_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 bool
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
var __t3 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t3 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_4))}, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 2465973597)) && ((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) {
__t4 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_4))}, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
})
})
}))
_ = eq1Coproduct2_3_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Coproduct2_3_1
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
var __t5 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t5 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 2465973597)) && ((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t6.IntVal)), UnsafePtr: nil}
})
})
}))
})
}

func Call_ordCoproduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
ord1Coproduct1_1_0 := Call_ord1Coproduct(dictOrd1_0)
_ = ord1Coproduct1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_6_4
eqCoproduct3_6_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if (v_7.Type == 9 && v_7.IntVal == 3711209382) {
var __t5 bool
{
if (v1_8.Type == 9 && v1_8.IntVal == 3711209382) {
__t5 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](__local_var_6_4))}, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0).IntVal) != (0)
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if ((v_7.Type == 9 && v_7.IntVal == 2465973597)) && ((v1_8.Type == 9 && v1_8.IntVal == 2465973597)) {
__t6 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](__local_var_6_4))}, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0).IntVal) != (0)
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
}))
_ = eqCoproduct3_6_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCoproduct3_6_3
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ord1Coproduct1_1_0, dictOrd11_3), "compare1"), dictOrd_5))
})
})
}

func Call_coproduct(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_extendCoproduct(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
functorCoproduct1_1_0 := gopurs_runtime.Apply(Get_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct1_1_0
return gopurs_runtime.Func(func(dictExtend1_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorCoproduct2_3_1 := gopurs_runtime.Apply(functorCoproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct2_3_1
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_5})})
}))
_ = __local_var_5_4
__local_var_5_3 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_5_4, x_6)})}
})
_ = __local_var_5_3
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_6})})
}))
_ = __local_var_6_6
__local_var_6_5 := gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_6, x_7)})}
})
_ = __local_var_6_5
__local_var_5_2 := gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(__local_var_5_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(__local_var_6_5, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, x_6)
})
}))
})
}

func Call_comonadCoproduct(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
extendCoproduct1_1_0 := Call_extendCoproduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}))
_ = extendCoproduct1_1_0
return gopurs_runtime.Func(func(dictComonad1_2 gopurs_runtime.Value) gopurs_runtime.Value {
extendCoproduct2_3_1 := gopurs_runtime.Apply(extendCoproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "Extend0"), gopurs_runtime.Value{}))
_ = extendCoproduct2_3_1
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendCoproduct2_3_1
}), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "extract"), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
}

func Call_bihoistCoproduct(natF_0_loop gopurs_runtime.Value, natG_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var natF_0 gopurs_runtime.Value = natF_0_loop
_ = natF_0
var natG_1 gopurs_runtime.Value = natG_1_loop
_ = natG_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorEither(), "bimap"), natF_0, natG_1, v_2)
}

func Call_extend__1264481661(dict_0_loop *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bimap__132457202(dict_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bimap__1783967194(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_eq1__1773593252(dict_0_loop *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_coproduct__79520197(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_coproduct__413515331(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_coproduct__829064685(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compare1__650153534(dict_0_loop *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_show__3978978930(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


