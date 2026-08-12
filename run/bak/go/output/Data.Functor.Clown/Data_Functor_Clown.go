package Data_Functor_Clown

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_Contravariant "gopurs/output/Data.Functor.Contravariant"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Clown gopurs_runtime.Value
var once_Clown sync.Once
func Get_Clown() gopurs_runtime.Value {
	once_Clown.Do(func() {
		cache_Clown = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Clown(x_0_box)
})
	})
	return cache_Clown
}

var cache_showClown gopurs_runtime.Value
var once_showClown sync.Once
func Get_showClown() gopurs_runtime.Value {
	once_showClown.Do(func() {
		cache_showClown = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showClown(dictShow_0_box)
})
	})
	return cache_showClown
}

var cache_profunctorClown gopurs_runtime.Value
var once_profunctorClown sync.Once
func Get_profunctorClown() gopurs_runtime.Value {
	once_profunctorClown.Do(func() {
		cache_profunctorClown = gopurs_runtime.Func(func(dictContravariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_profunctorClown(dictContravariant_0_box)
})
	})
	return cache_profunctorClown
}

var cache_ordClown gopurs_runtime.Value
var once_ordClown sync.Once
func Get_ordClown() gopurs_runtime.Value {
	once_ordClown.Do(func() {
		cache_ordClown = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordClown(dictOrd_0_box)
})
	})
	return cache_ordClown
}

var cache_newtypeClown gopurs_runtime.Value
var once_newtypeClown sync.Once
func Get_newtypeClown() gopurs_runtime.Value {
	once_newtypeClown.Do(func() {
		cache_newtypeClown = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeClown
}

var cache_hoistClown gopurs_runtime.Value
var once_hoistClown sync.Once
func Get_hoistClown() gopurs_runtime.Value {
	once_hoistClown.Do(func() {
		cache_hoistClown = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistClown(f_0_box, v_1_box)
})
	})
	return cache_hoistClown
}

var cache_functorClown gopurs_runtime.Value
var once_functorClown sync.Once
func Get_functorClown() gopurs_runtime.Value {
	once_functorClown.Do(func() {
		cache_functorClown = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_functorClown
}

var cache_eqClown gopurs_runtime.Value
var once_eqClown sync.Once
func Get_eqClown() gopurs_runtime.Value {
	once_eqClown.Do(func() {
		cache_eqClown = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqClown(dictEq_0_box)
})
	})
	return cache_eqClown
}

var cache_bifunctorClown gopurs_runtime.Value
var once_bifunctorClown sync.Once
func Get_bifunctorClown() gopurs_runtime.Value {
	once_bifunctorClown.Do(func() {
		cache_bifunctorClown = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifunctorClown(dictFunctor_0_box)
})
	})
	return cache_bifunctorClown
}

var cache_biapplyClown gopurs_runtime.Value
var once_biapplyClown sync.Once
func Get_biapplyClown() gopurs_runtime.Value {
	once_biapplyClown.Do(func() {
		cache_biapplyClown = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplyClown(dictApply_0_box)
})
	})
	return cache_biapplyClown
}

var cache_biapplicativeClown gopurs_runtime.Value
var once_biapplicativeClown sync.Once
func Get_biapplicativeClown() gopurs_runtime.Value {
	once_biapplicativeClown.Do(func() {
		cache_biapplicativeClown = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplicativeClown(dictApplicative_0_box)
})
	})
	return cache_biapplicativeClown
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
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

var cache_cmap__326373820 gopurs_runtime.Value
var once_cmap__326373820 sync.Once
func Get_cmap__326373820() gopurs_runtime.Value {
	once_cmap__326373820.Do(func() {
		cache_cmap__326373820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cmap__326373820(gopurs_runtime.CoerceToStruct[pkg_Data_Functor_Contravariant.Constructor_Contravariant[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_cmap__326373820
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

func Call_Clown(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showClown(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Clown "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_profunctorClown(dictContravariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_1, v1_3)
})
})
}))
}

func Call_ordClown(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_hoistClown(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_eqClown(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_bifunctorClown(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, v1_3)
})
})
}))
}

func Call_biapplyClown(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
bifunctorClown1_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, v1_4)
})
})
}))
_ = bifunctorClown1_1_0
return gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), v_2, v1_3)
})
}))
}

func Call_biapplicativeClown(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
bifunctorClown1_2_2 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, v1_5)
})
})
}))
_ = bifunctorClown1_2_2
biapplyClown1_1_0 := gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), v_3, v1_4)
})
}))
_ = biapplyClown1_1_0
return gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyClown1_1_0
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), a_2)
})
}))
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_cmap__326373820(dict_0_loop *pkg_Data_Functor_Contravariant.Constructor_Contravariant[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor_Contravariant.Constructor_Contravariant[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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


