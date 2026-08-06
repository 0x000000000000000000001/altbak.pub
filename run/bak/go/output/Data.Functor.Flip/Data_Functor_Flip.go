package Data_Functor_Flip

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
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

func Call_Flip(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showFlip(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Flip "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")"))).StrVal())
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


