package Data_Functor_Clown

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Clown gopurs_runtime.Value
var once_Clown sync.Once
func Get_Clown() gopurs_runtime.Value {
	once_Clown.Do(func() {
		cache_Clown = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Clown
}

var cache_showClown gopurs_runtime.Value
var once_showClown sync.Once
func Get_showClown() gopurs_runtime.Value {
	once_showClown.Do(func() {
		cache_showClown = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Clown ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showClown
}

var cache_profunctorClown gopurs_runtime.Value
var once_profunctorClown sync.Once
func Get_profunctorClown() gopurs_runtime.Value {
	once_profunctorClown.Do(func() {
		cache_profunctorClown = gopurs_runtime.Func(func(dictContravariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_1, v1_3)
}))
}()
})
	})
	return cache_profunctorClown
}

var cache_ordClown gopurs_runtime.Value
var once_ordClown sync.Once
func Get_ordClown() gopurs_runtime.Value {
	once_ordClown.Do(func() {
		cache_ordClown = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}()
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
		cache_functorClown = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return cache_functorClown
}

var cache_eqClown gopurs_runtime.Value
var once_eqClown sync.Once
func Get_eqClown() gopurs_runtime.Value {
	once_eqClown.Do(func() {
		cache_eqClown = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}()
})
	})
	return cache_eqClown
}

var cache_bifunctorClown gopurs_runtime.Value
var once_bifunctorClown sync.Once
func Get_bifunctorClown() gopurs_runtime.Value {
	once_bifunctorClown.Do(func() {
		cache_bifunctorClown = gopurs_runtime.Func(func(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, v1_3)
}))
}()
})
	})
	return cache_bifunctorClown
}

var cache_biapplyClown gopurs_runtime.Value
var once_biapplyClown sync.Once
func Get_biapplyClown() gopurs_runtime.Value {
	once_biapplyClown.Do(func() {
		cache_biapplyClown = gopurs_runtime.Func(func(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorClown1_2_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, v1_4)
}))
_ = bifunctorClown1_2_1
return gopurs_runtime.RecordDict2("biapply", "Bifunctor0", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), v_3, v1_4)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown1_2_1
}))
}()
})
	})
	return cache_biapplyClown
}

var cache_biapplicativeClown gopurs_runtime.Value
var once_biapplicativeClown sync.Once
func Get_biapplicativeClown() gopurs_runtime.Value {
	once_biapplicativeClown.Do(func() {
		cache_biapplicativeClown = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
bifunctorClown1_3_3 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_3, v1_5)
}))
_ = bifunctorClown1_3_3
biapplyClown1_3_2 := gopurs_runtime.RecordDict2("biapply", "Bifunctor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), v_4, v1_5)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown1_3_3
}))
_ = biapplyClown1_3_2
return gopurs_runtime.RecordDict2("bipure", "Biapply0", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), a_4)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyClown1_3_2
}))
}()
})
	})
	return cache_biapplicativeClown
}

func Call_hoistClown(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}


