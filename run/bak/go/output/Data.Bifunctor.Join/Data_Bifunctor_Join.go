package Data_Bifunctor_Join

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Join gopurs_runtime.Value
var once_Join sync.Once
func Get_Join() gopurs_runtime.Value {
	once_Join.Do(func() {
		Join = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Join
}

var showJoin gopurs_runtime.Value
var once_showJoin sync.Once
func Get_showJoin() gopurs_runtime.Value {
	once_showJoin.Do(func() {
		showJoin = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Join " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), v_1).StrVal + ")")
}))
}()
})
	})
	return showJoin
}

var ordJoin gopurs_runtime.Value
var once_ordJoin sync.Once
func Get_ordJoin() gopurs_runtime.Value {
	once_ordJoin.Do(func() {
		ordJoin = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0_loop
}()
})
	})
	return ordJoin
}

var newtypeJoin gopurs_runtime.Value
var once_newtypeJoin sync.Once
func Get_newtypeJoin() gopurs_runtime.Value {
	once_newtypeJoin.Do(func() {
		newtypeJoin = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeJoin
}

var eqJoin gopurs_runtime.Value
var once_eqJoin sync.Once
func Get_eqJoin() gopurs_runtime.Value {
	once_eqJoin.Do(func() {
		eqJoin = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0_loop
}()
})
	})
	return eqJoin
}

var bifunctorJoin gopurs_runtime.Value
var once_bifunctorJoin sync.Once
func Get_bifunctorJoin() gopurs_runtime.Value {
	once_bifunctorJoin.Do(func() {
		bifunctorJoin = gopurs_runtime.Func(func(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0_loop, "bimap"), f_1, f_1, v_2)
}))
}()
})
	})
	return bifunctorJoin
}

var biapplyJoin gopurs_runtime.Value
var once_biapplyJoin sync.Once
func Get_biapplyJoin() gopurs_runtime.Value {
	once_biapplyJoin.Do(func() {
		biapplyJoin = gopurs_runtime.Func(func(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0_loop, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorJoin1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_2, f_2, v_3)
}))
_ = bifunctorJoin1_2_1
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0_loop, "biapply"), v_3, v1_4)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoin1_2_1
}))
}()
})
	})
	return biapplyJoin
}

var biapplicativeJoin gopurs_runtime.Value
var once_biapplicativeJoin sync.Once
func Get_biapplicativeJoin() gopurs_runtime.Value {
	once_biapplicativeJoin.Do(func() {
		biapplicativeJoin = gopurs_runtime.Func(func(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
_ = dictBiapplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0_loop, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
bifunctorJoin1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "bimap"), f_3, f_3, v_4)
}))
_ = bifunctorJoin1_3_3
biapplyJoin1_3_2 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "biapply"), v_4, v1_5)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoin1_3_3
}))
_ = biapplyJoin1_3_2
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0_loop, "bipure"), a_4, a_4)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyJoin1_3_2
}))
}()
})
	})
	return biapplicativeJoin
}




