package Data_Bifunctor_Join

import (
	pkg_Control_Biapplicative "gopurs/output/Control.Biapplicative"
	pkg_Control_Biapply "gopurs/output/Control.Biapply"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Join gopurs_runtime.Value
var once_Join sync.Once
func Get_Join() gopurs_runtime.Value {
	once_Join.Do(func() {
		cache_Join = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Join(x_0_box)
})
	})
	return cache_Join
}

var cache_showJoin gopurs_runtime.Value
var once_showJoin sync.Once
func Get_showJoin() gopurs_runtime.Value {
	once_showJoin.Do(func() {
		cache_showJoin = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showJoin(dictShow_0_box)
})
	})
	return cache_showJoin
}

var cache_ordJoin gopurs_runtime.Value
var once_ordJoin sync.Once
func Get_ordJoin() gopurs_runtime.Value {
	once_ordJoin.Do(func() {
		cache_ordJoin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordJoin(dictOrd_0_box)
})
	})
	return cache_ordJoin
}

var cache_newtypeJoin gopurs_runtime.Value
var once_newtypeJoin sync.Once
func Get_newtypeJoin() gopurs_runtime.Value {
	once_newtypeJoin.Do(func() {
		cache_newtypeJoin = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeJoin
}

var cache_eqJoin gopurs_runtime.Value
var once_eqJoin sync.Once
func Get_eqJoin() gopurs_runtime.Value {
	once_eqJoin.Do(func() {
		cache_eqJoin = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqJoin(dictEq_0_box)
})
	})
	return cache_eqJoin
}

var cache_bifunctorJoin gopurs_runtime.Value
var once_bifunctorJoin sync.Once
func Get_bifunctorJoin() gopurs_runtime.Value {
	once_bifunctorJoin.Do(func() {
		cache_bifunctorJoin = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifunctorJoin(dictBifunctor_0_box)
})
	})
	return cache_bifunctorJoin
}

var cache_biapplyJoin gopurs_runtime.Value
var once_biapplyJoin sync.Once
func Get_biapplyJoin() gopurs_runtime.Value {
	once_biapplyJoin.Do(func() {
		cache_biapplyJoin = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplyJoin(dictBiapply_0_box)
})
	})
	return cache_biapplyJoin
}

var cache_biapplicativeJoin gopurs_runtime.Value
var once_biapplicativeJoin sync.Once
func Get_biapplicativeJoin() gopurs_runtime.Value {
	once_biapplicativeJoin.Do(func() {
		cache_biapplicativeJoin = gopurs_runtime.Func(func(dictBiapplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplicativeJoin(dictBiapplicative_0_box)
})
	})
	return cache_biapplicativeJoin
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

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
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

func Call_Join(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showJoin(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Join "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_ordJoin(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_eqJoin(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_bifunctorJoin(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, f_1, v_2)
})
}))
}

func Call_biapplyJoin(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
bifunctorJoin1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), f_2, f_2, v_3)
})
}))
_ = bifunctorJoin1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoin1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), v_2, v1_3)
})
}))
}

func Call_biapplicativeJoin(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
_ = dictBiapplicative_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
bifunctorJoin1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "bimap"), f_3, f_3, v_4)
})
}))
_ = bifunctorJoin1_2_2
biapplyJoin1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoin1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "biapply"), v_3, v1_4)
})
}))
_ = biapplyJoin1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyJoin1_1_0
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), a_2, a_2)
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

func Call_bimap__132457202(dict_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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


