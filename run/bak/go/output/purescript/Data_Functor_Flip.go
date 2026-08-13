package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Functor_Flip_Flip gopurs_runtime.Value
var once_Data_Functor_Flip_Flip sync.Once
func Get_Data_Functor_Flip_Flip() gopurs_runtime.Value {
	once_Data_Functor_Flip_Flip.Do(func() {
		cache_Data_Functor_Flip_Flip = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_Flip(x_0_box)
})
	})
	return cache_Data_Functor_Flip_Flip
}

var cache_Data_Functor_Flip_showFlip gopurs_runtime.Value
var once_Data_Functor_Flip_showFlip sync.Once
func Get_Data_Functor_Flip_showFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_showFlip.Do(func() {
		cache_Data_Functor_Flip_showFlip = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_showFlip(dictShow_0_box)
})
	})
	return cache_Data_Functor_Flip_showFlip
}

var cache_Data_Functor_Flip_semigroupoidFlip gopurs_runtime.Value
var once_Data_Functor_Flip_semigroupoidFlip sync.Once
func Get_Data_Functor_Flip_semigroupoidFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_semigroupoidFlip.Do(func() {
		cache_Data_Functor_Flip_semigroupoidFlip = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_semigroupoidFlip(dictSemigroupoid_0_box)
})
	})
	return cache_Data_Functor_Flip_semigroupoidFlip
}

var cache_Data_Functor_Flip_ordFlip gopurs_runtime.Value
var once_Data_Functor_Flip_ordFlip sync.Once
func Get_Data_Functor_Flip_ordFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_ordFlip.Do(func() {
		cache_Data_Functor_Flip_ordFlip = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_ordFlip(dictOrd_0_box)
})
	})
	return cache_Data_Functor_Flip_ordFlip
}

var cache_Data_Functor_Flip_newtypeFlip gopurs_runtime.Value
var once_Data_Functor_Flip_newtypeFlip sync.Once
func Get_Data_Functor_Flip_newtypeFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_newtypeFlip.Do(func() {
		cache_Data_Functor_Flip_newtypeFlip = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Functor_Flip_newtypeFlip
}

var cache_Data_Functor_Flip_functorFlip gopurs_runtime.Value
var once_Data_Functor_Flip_functorFlip sync.Once
func Get_Data_Functor_Flip_functorFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_functorFlip.Do(func() {
		cache_Data_Functor_Flip_functorFlip = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_functorFlip(dictBifunctor_0_box)
})
	})
	return cache_Data_Functor_Flip_functorFlip
}

var cache_Data_Functor_Flip_eqFlip gopurs_runtime.Value
var once_Data_Functor_Flip_eqFlip sync.Once
func Get_Data_Functor_Flip_eqFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_eqFlip.Do(func() {
		cache_Data_Functor_Flip_eqFlip = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_eqFlip(dictEq_0_box)
})
	})
	return cache_Data_Functor_Flip_eqFlip
}

var cache_Data_Functor_Flip_contravariantFlip gopurs_runtime.Value
var once_Data_Functor_Flip_contravariantFlip sync.Once
func Get_Data_Functor_Flip_contravariantFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_contravariantFlip.Do(func() {
		cache_Data_Functor_Flip_contravariantFlip = gopurs_runtime.Func(func(dictProfunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_contravariantFlip(dictProfunctor_0_box)
})
	})
	return cache_Data_Functor_Flip_contravariantFlip
}

var cache_Data_Functor_Flip_categoryFlip gopurs_runtime.Value
var once_Data_Functor_Flip_categoryFlip sync.Once
func Get_Data_Functor_Flip_categoryFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_categoryFlip.Do(func() {
		cache_Data_Functor_Flip_categoryFlip = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_categoryFlip(dictCategory_0_box)
})
	})
	return cache_Data_Functor_Flip_categoryFlip
}

var cache_Data_Functor_Flip_bifunctorFlip gopurs_runtime.Value
var once_Data_Functor_Flip_bifunctorFlip sync.Once
func Get_Data_Functor_Flip_bifunctorFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_bifunctorFlip.Do(func() {
		cache_Data_Functor_Flip_bifunctorFlip = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_bifunctorFlip(dictBifunctor_0_box)
})
	})
	return cache_Data_Functor_Flip_bifunctorFlip
}

var cache_Data_Functor_Flip_biapplyFlip gopurs_runtime.Value
var once_Data_Functor_Flip_biapplyFlip sync.Once
func Get_Data_Functor_Flip_biapplyFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_biapplyFlip.Do(func() {
		cache_Data_Functor_Flip_biapplyFlip = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_biapplyFlip(dictBiapply_0_box)
})
	})
	return cache_Data_Functor_Flip_biapplyFlip
}

var cache_Data_Functor_Flip_biapplicativeFlip gopurs_runtime.Value
var once_Data_Functor_Flip_biapplicativeFlip sync.Once
func Get_Data_Functor_Flip_biapplicativeFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_biapplicativeFlip.Do(func() {
		cache_Data_Functor_Flip_biapplicativeFlip = gopurs_runtime.Func(func(dictBiapplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Flip_biapplicativeFlip(dictBiapplicative_0_box)
})
	})
	return cache_Data_Functor_Flip_biapplicativeFlip
}

func Call_Data_Functor_Flip_Flip(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Flip_showFlip(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Flip ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Functor_Flip_semigroupoidFlip(dictSemigroupoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
return gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), v1_2, v_1)
})
}))
}

func Call_Data_Functor_Flip_ordFlip(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Functor_Flip_functorFlip(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}), v_2)
})
}))
}

func Call_Data_Functor_Flip_eqFlip(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Functor_Flip_contravariantFlip(dictProfunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
return gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}), v_2)
})
}))
}

func Call_Data_Functor_Flip_categoryFlip(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
_ = dictCategory_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupoidFlip1_1_0 -> gopurs_runtime.Value
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

func Call_Data_Functor_Flip_bifunctorFlip(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Functor_Flip_biapplyFlip(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorFlip1_1_0 -> gopurs_runtime.Value
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

func Call_Data_Functor_Flip_biapplicativeFlip(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
_ = dictBiapplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifunctorFlip1_2_2 -> gopurs_runtime.Value
bifunctorFlip1_2_2 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "bimap"), g_4, f_3, v_5)
})
})
}))
_ = bifunctorFlip1_2_2
// TAST (Let): biapplyFlip1_1_0 -> gopurs_runtime.Value
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


