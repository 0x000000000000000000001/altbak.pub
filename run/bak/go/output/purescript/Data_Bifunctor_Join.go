package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Bifunctor_Join_Join gopurs_runtime.Value
var once_Data_Bifunctor_Join_Join sync.Once
func Get_Data_Bifunctor_Join_Join() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_Join.Do(func() {
		cache_Data_Bifunctor_Join_Join = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_Join_Join(x_0_box)
})
	})
	return cache_Data_Bifunctor_Join_Join
}

var cache_Data_Bifunctor_Join_showJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_showJoin sync.Once
func Get_Data_Bifunctor_Join_showJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_showJoin.Do(func() {
		cache_Data_Bifunctor_Join_showJoin = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_Join_showJoin(dictShow_0_box)
})
	})
	return cache_Data_Bifunctor_Join_showJoin
}

var cache_Data_Bifunctor_Join_ordJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_ordJoin sync.Once
func Get_Data_Bifunctor_Join_ordJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_ordJoin.Do(func() {
		cache_Data_Bifunctor_Join_ordJoin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_Join_ordJoin(dictOrd_0_box)
})
	})
	return cache_Data_Bifunctor_Join_ordJoin
}

var cache_Data_Bifunctor_Join_newtypeJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_newtypeJoin sync.Once
func Get_Data_Bifunctor_Join_newtypeJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_newtypeJoin.Do(func() {
		cache_Data_Bifunctor_Join_newtypeJoin = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Bifunctor_Join_newtypeJoin
}

var cache_Data_Bifunctor_Join_eqJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_eqJoin sync.Once
func Get_Data_Bifunctor_Join_eqJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_eqJoin.Do(func() {
		cache_Data_Bifunctor_Join_eqJoin = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_Join_eqJoin(dictEq_0_box)
})
	})
	return cache_Data_Bifunctor_Join_eqJoin
}

var cache_Data_Bifunctor_Join_bifunctorJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_bifunctorJoin sync.Once
func Get_Data_Bifunctor_Join_bifunctorJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_bifunctorJoin.Do(func() {
		cache_Data_Bifunctor_Join_bifunctorJoin = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_Join_bifunctorJoin(dictBifunctor_0_box)
})
	})
	return cache_Data_Bifunctor_Join_bifunctorJoin
}

var cache_Data_Bifunctor_Join_biapplyJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_biapplyJoin sync.Once
func Get_Data_Bifunctor_Join_biapplyJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_biapplyJoin.Do(func() {
		cache_Data_Bifunctor_Join_biapplyJoin = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_Join_biapplyJoin(dictBiapply_0_box)
})
	})
	return cache_Data_Bifunctor_Join_biapplyJoin
}

var cache_Data_Bifunctor_Join_biapplicativeJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_biapplicativeJoin sync.Once
func Get_Data_Bifunctor_Join_biapplicativeJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_biapplicativeJoin.Do(func() {
		cache_Data_Bifunctor_Join_biapplicativeJoin = gopurs_runtime.Func(func(dictBiapplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_Join_biapplicativeJoin(dictBiapplicative_0_box)
})
	})
	return cache_Data_Bifunctor_Join_biapplicativeJoin
}

func Call_Data_Bifunctor_Join_Join(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bifunctor_Join_showJoin(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Join ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Bifunctor_Join_ordJoin(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Bifunctor_Join_eqJoin(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Bifunctor_Join_bifunctorJoin(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, f_1, v_2)
})
}))
}

func Call_Data_Bifunctor_Join_biapplyJoin(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorJoin1_1_0 -> gopurs_runtime.Value
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

func Call_Data_Bifunctor_Join_biapplicativeJoin(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
_ = dictBiapplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifunctorJoin1_2_2 -> gopurs_runtime.Value
bifunctorJoin1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "bimap"), f_3, f_3, v_4)
})
}))
_ = bifunctorJoin1_2_2
// TAST (Let): biapplyJoin1_1_0 -> gopurs_runtime.Value
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


