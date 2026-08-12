package Data_Functor_Flip

import (
	pkg_Control_Biapplicative "gopurs/output/Control.Biapplicative"
	pkg_Control_Biapply "gopurs/output/Control.Biapply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Flip gopurs_runtime.Value
var once_Flip sync.Once
func Get_Flip() gopurs_runtime.Value {
	once_Flip.Do(func() {
		cache_Flip = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Flip(x_0_box)
})
	})
	return cache_Flip
}

var cache_showFlip gopurs_runtime.Value
var once_showFlip sync.Once
func Get_showFlip() gopurs_runtime.Value {
	once_showFlip.Do(func() {
		cache_showFlip = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showFlip(dictShow_0_box)
})
	})
	return cache_showFlip
}

var cache_semigroupoidFlip gopurs_runtime.Value
var once_semigroupoidFlip sync.Once
func Get_semigroupoidFlip() gopurs_runtime.Value {
	once_semigroupoidFlip.Do(func() {
		cache_semigroupoidFlip = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupoidFlip(dictSemigroupoid_0_box)
})
	})
	return cache_semigroupoidFlip
}

var cache_ordFlip gopurs_runtime.Value
var once_ordFlip sync.Once
func Get_ordFlip() gopurs_runtime.Value {
	once_ordFlip.Do(func() {
		cache_ordFlip = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordFlip(dictOrd_0_box)
})
	})
	return cache_ordFlip
}

var cache_newtypeFlip gopurs_runtime.Value
var once_newtypeFlip sync.Once
func Get_newtypeFlip() gopurs_runtime.Value {
	once_newtypeFlip.Do(func() {
		cache_newtypeFlip = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeFlip
}

var cache_functorFlip gopurs_runtime.Value
var once_functorFlip sync.Once
func Get_functorFlip() gopurs_runtime.Value {
	once_functorFlip.Do(func() {
		cache_functorFlip = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorFlip(dictBifunctor_0_box)
})
	})
	return cache_functorFlip
}

var cache_eqFlip gopurs_runtime.Value
var once_eqFlip sync.Once
func Get_eqFlip() gopurs_runtime.Value {
	once_eqFlip.Do(func() {
		cache_eqFlip = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqFlip(dictEq_0_box)
})
	})
	return cache_eqFlip
}

var cache_contravariantFlip gopurs_runtime.Value
var once_contravariantFlip sync.Once
func Get_contravariantFlip() gopurs_runtime.Value {
	once_contravariantFlip.Do(func() {
		cache_contravariantFlip = gopurs_runtime.Func(func(dictProfunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_contravariantFlip(dictProfunctor_0_box)
})
	})
	return cache_contravariantFlip
}

var cache_categoryFlip gopurs_runtime.Value
var once_categoryFlip sync.Once
func Get_categoryFlip() gopurs_runtime.Value {
	once_categoryFlip.Do(func() {
		cache_categoryFlip = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_categoryFlip(dictCategory_0_box)
})
	})
	return cache_categoryFlip
}

var cache_bifunctorFlip gopurs_runtime.Value
var once_bifunctorFlip sync.Once
func Get_bifunctorFlip() gopurs_runtime.Value {
	once_bifunctorFlip.Do(func() {
		cache_bifunctorFlip = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifunctorFlip(dictBifunctor_0_box)
})
	})
	return cache_bifunctorFlip
}

var cache_biapplyFlip gopurs_runtime.Value
var once_biapplyFlip sync.Once
func Get_biapplyFlip() gopurs_runtime.Value {
	once_biapplyFlip.Do(func() {
		cache_biapplyFlip = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplyFlip(dictBiapply_0_box)
})
	})
	return cache_biapplyFlip
}

var cache_biapplicativeFlip gopurs_runtime.Value
var once_biapplicativeFlip sync.Once
func Get_biapplicativeFlip() gopurs_runtime.Value {
	once_biapplicativeFlip.Do(func() {
		cache_biapplicativeFlip = gopurs_runtime.Func(func(dictBiapplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplicativeFlip(dictBiapplicative_0_box)
})
	})
	return cache_biapplicativeFlip
}

var cache_bipure__1449949980 gopurs_runtime.Value
var once_bipure__1449949980 sync.Once
func Get_bipure__1449949980() gopurs_runtime.Value {
	once_bipure__1449949980.Do(func() {
		cache_bipure__1449949980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bipure__1449949980(gopurs_runtime.CoerceToStruct[pkg_Control_Biapplicative.Constructor_Biapplicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bipure__1449949980
}

var cache_biapply__3394381979 gopurs_runtime.Value
var once_biapply__3394381979 sync.Once
func Get_biapply__3394381979() gopurs_runtime.Value {
	once_biapply__3394381979.Do(func() {
		cache_biapply__3394381979 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapply__3394381979(gopurs_runtime.CoerceToStruct[pkg_Control_Biapply.Constructor_Biapply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_biapply__3394381979
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

var cache_lmap__2196160232 gopurs_runtime.Value
var once_lmap__2196160232 sync.Once
func Get_lmap__2196160232() gopurs_runtime.Value {
	once_lmap__2196160232.Do(func() {
		cache_lmap__2196160232 = gopurs_runtime.Func2(func(dictBifunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lmap__2196160232(gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](dictBifunctor_0_box), f_1_box)
})
	})
	return cache_lmap__2196160232
}

var cache_dimap__1466332548 gopurs_runtime.Value
var once_dimap__1466332548 sync.Once
func Get_dimap__1466332548() gopurs_runtime.Value {
	once_dimap__1466332548.Do(func() {
		cache_dimap__1466332548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dimap__1466332548(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_dimap__1466332548
}

var cache_lcmap__1762133278 gopurs_runtime.Value
var once_lcmap__1762133278 sync.Once
func Get_lcmap__1762133278() gopurs_runtime.Value {
	once_lcmap__1762133278.Do(func() {
		cache_lcmap__1762133278 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, a2b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lcmap__1762133278(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), a2b_1_box)
})
	})
	return cache_lcmap__1762133278
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

var cache_show__255526802 gopurs_runtime.Value
var once_show__255526802 sync.Once
func Get_show__255526802() gopurs_runtime.Value {
	once_show__255526802.Do(func() {
		cache_show__255526802 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__255526802(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__255526802
}

func Call_Flip(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showFlip(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Flip "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_semigroupoidFlip(dictSemigroupoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
return gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), v1_2, v_1)
})
}))
}

func Call_ordFlip(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_functorFlip(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, pkg_Data_Bifunctor.Get_identity(), v_2)
})
}))
}

func Call_eqFlip(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_contravariantFlip(dictProfunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
return gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_1, pkg_Data_Profunctor.Get_identity(), v_2)
})
}))
}

func Call_categoryFlip(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
_ = dictCategory_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_1_1
semigroupoidFlip1_1_0 := gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compose"), v1_3, v_2)
})
}))
_ = semigroupoidFlip1_1_0
return gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupoidFlip1_1_0
}), gopurs_runtime.RecordGet(dictCategory_0, "identity"))
}

func Call_bifunctorFlip(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), g_2, f_1, v_3)
})
})
}))
}

func Call_biapplyFlip(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
bifunctorFlip1_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), g_3, f_2, v_4)
})
})
}))
_ = bifunctorFlip1_1_0
return gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), v_2, v1_3)
})
}))
}

func Call_biapplicativeFlip(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
_ = dictBiapplicative_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
bifunctorFlip1_2_2 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "bimap"), g_4, f_3, v_5)
})
})
}))
_ = bifunctorFlip1_2_2
biapplyFlip1_1_0 := gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "biapply"), v_3, v1_4)
})
}))
_ = biapplyFlip1_1_0
return gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyFlip1_1_0
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), b_3, a_2)
})
}))
}

func Call_bipure__1449949980(dict_0_loop *pkg_Control_Biapplicative.Constructor_Biapplicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Biapplicative.Constructor_Biapplicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_biapply__3394381979(dict_0_loop *pkg_Control_Biapply.Constructor_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Biapply.Constructor_Biapply[gopurs_runtime.Value] = dict_0_loop
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

func Call_lmap__2196160232(dictBifunctor_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dictBifunctor_0_loop
_ = dictBifunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(dictBifunctor_0.V0, f_1, pkg_Data_Bifunctor.Get_identity())
}

func Call_dimap__1466332548(dict_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_lcmap__1762133278(dictProfunctor_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value], a2b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var a2b_1 gopurs_runtime.Value = a2b_1_loop
_ = a2b_1
return gopurs_runtime.Apply2(dictProfunctor_0.V0, a2b_1, pkg_Data_Profunctor.Get_identity())
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

func Call_show__255526802(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


