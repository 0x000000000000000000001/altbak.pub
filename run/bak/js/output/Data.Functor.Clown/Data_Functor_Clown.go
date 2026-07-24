package Data_Functor_Clown

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Clown gopurs_runtime.Value
var once_Clown sync.Once
func Get_Clown() gopurs_runtime.Value {
	once_Clown.Do(func() {
		Clown = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Clown
}

var showClown gopurs_runtime.Value
var once_showClown sync.Once
func Get_showClown() gopurs_runtime.Value {
	once_showClown.Do(func() {
		showClown = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Clown " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal + ")")
}))
})
	})
	return showClown
}

var profunctorClown gopurs_runtime.Value
var once_profunctorClown sync.Once
func Get_profunctorClown() gopurs_runtime.Value {
	once_profunctorClown.Do(func() {
		profunctorClown = gopurs_runtime.Func(func(dictContravariant_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_1, v1_3)
}))
})
	})
	return profunctorClown
}

var ordClown gopurs_runtime.Value
var once_ordClown sync.Once
func Get_ordClown() gopurs_runtime.Value {
	once_ordClown.Do(func() {
		ordClown = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordClown
}

var newtypeClown gopurs_runtime.Value
var once_newtypeClown sync.Once
func Get_newtypeClown() gopurs_runtime.Value {
	once_newtypeClown.Do(func() {
		newtypeClown = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeClown
}

var hoistClown gopurs_runtime.Value
var once_hoistClown sync.Once
func Get_hoistClown() gopurs_runtime.Value {
	once_hoistClown.Do(func() {
		hoistClown = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
	})
	return hoistClown
}

var functorClown gopurs_runtime.Value
var once_functorClown sync.Once
func Get_functorClown() gopurs_runtime.Value {
	once_functorClown.Do(func() {
		functorClown = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return functorClown
}

var eqClown gopurs_runtime.Value
var once_eqClown sync.Once
func Get_eqClown() gopurs_runtime.Value {
	once_eqClown.Do(func() {
		eqClown = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqClown
}

var bifunctorClown gopurs_runtime.Value
var once_bifunctorClown sync.Once
func Get_bifunctorClown() gopurs_runtime.Value {
	once_bifunctorClown.Do(func() {
		bifunctorClown = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, v1_3)
}))
})
	})
	return bifunctorClown
}

var biapplyClown gopurs_runtime.Value
var once_biapplyClown sync.Once
func Get_biapplyClown() gopurs_runtime.Value {
	once_biapplyClown.Do(func() {
		biapplyClown = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
	})
	return biapplyClown
}

var biapplicativeClown gopurs_runtime.Value
var once_biapplicativeClown sync.Once
func Get_biapplicativeClown() gopurs_runtime.Value {
	once_biapplicativeClown.Do(func() {
		biapplicativeClown = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
	})
	return biapplicativeClown
}




