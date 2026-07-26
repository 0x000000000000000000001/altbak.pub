package Data_Predicate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_BooleanAlgebra "gopurs/output/Data.BooleanAlgebra"
)

var cache_Predicate gopurs_runtime.Value
var once_Predicate sync.Once
func Get_Predicate() gopurs_runtime.Value {
	once_Predicate.Do(func() {
		cache_Predicate = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Predicate(x_0_box)
})
	})
	return cache_Predicate
}

var cache_newtypePredicate gopurs_runtime.Value
var once_newtypePredicate sync.Once
func Get_newtypePredicate() gopurs_runtime.Value {
	once_newtypePredicate.Do(func() {
		cache_newtypePredicate = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypePredicate
}

var cache_heytingAlgebraPredicate gopurs_runtime.Value
var once_heytingAlgebraPredicate sync.Once
func Get_heytingAlgebraPredicate() gopurs_runtime.Value {
	once_heytingAlgebraPredicate.Do(func() {
		cache_heytingAlgebraPredicate = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply(f_0, a_2), gopurs_runtime.Apply(g_1, a_2))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(f_0, a_2), gopurs_runtime.Apply(g_1, a_2))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(f_0, a_2)), gopurs_runtime.Apply(g_1, a_2))
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(f_0, a_1))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "tt")
})})
	})
	return cache_heytingAlgebraPredicate
}

var cache_contravariantPredicate gopurs_runtime.Value
var once_contravariantPredicate sync.Once
func Get_contravariantPredicate() gopurs_runtime.Value {
	once_contravariantPredicate.Do(func() {
		cache_contravariantPredicate = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), v_1, f_0)
}))
	})
	return cache_contravariantPredicate
}

var cache_booleanAlgebraPredicate gopurs_runtime.Value
var once_booleanAlgebraPredicate sync.Once
func Get_booleanAlgebraPredicate() gopurs_runtime.Value {
	once_booleanAlgebraPredicate.Do(func() {
		cache_booleanAlgebraPredicate = gopurs_runtime.Apply(pkg_Data_BooleanAlgebra.Get_booleanAlgebraFn(), pkg_Data_BooleanAlgebra.Get_booleanAlgebraBoolean())
	})
	return cache_booleanAlgebraPredicate
}

func Call_Predicate(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}


