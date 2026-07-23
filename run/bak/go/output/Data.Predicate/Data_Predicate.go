package Data_Predicate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_BooleanAlgebra "gopurs/output/Data.BooleanAlgebra"
)

var Predicate gopurs_runtime.Value
var once_Predicate sync.Once
func Get_Predicate() gopurs_runtime.Value {
	once_Predicate.Do(func() {
		Predicate = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Predicate
}

var newtypePredicate gopurs_runtime.Value
var once_newtypePredicate sync.Once
func Get_newtypePredicate() gopurs_runtime.Value {
	once_newtypePredicate.Do(func() {
		newtypePredicate = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypePredicate
}

var heytingAlgebraPredicate gopurs_runtime.Value
var once_heytingAlgebraPredicate sync.Once
func Get_heytingAlgebraPredicate() gopurs_runtime.Value {
	once_heytingAlgebraPredicate.Do(func() {
		heytingAlgebraPredicate = gopurs_runtime.RecordDict([]string{"ff", "tt", "implies", "conj", "disj", "not"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(f_0, a_2).IntVal == 0).IntVal != 0 || gopurs_runtime.Apply(g_1, a_2).IntVal != 0)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(f_0, a_2).IntVal != 0 && gopurs_runtime.Apply(g_1, a_2).IntVal != 0)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(f_0, a_2).IntVal != 0 || gopurs_runtime.Apply(g_1, a_2).IntVal != 0)
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(f_0, a_1).IntVal == 0)
})})
	})
	return heytingAlgebraPredicate
}

var contravariantPredicate gopurs_runtime.Value
var once_contravariantPredicate sync.Once
func Get_contravariantPredicate() gopurs_runtime.Value {
	once_contravariantPredicate.Do(func() {
		contravariantPredicate = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}))
	})
	return contravariantPredicate
}

var booleanAlgebraPredicate gopurs_runtime.Value
var once_booleanAlgebraPredicate sync.Once
func Get_booleanAlgebraPredicate() gopurs_runtime.Value {
	once_booleanAlgebraPredicate.Do(func() {
		booleanAlgebraPredicate = gopurs_runtime.Apply(pkg_Data_BooleanAlgebra.Get_booleanAlgebraFn(), pkg_Data_BooleanAlgebra.Get_booleanAlgebraBoolean())
	})
	return booleanAlgebraPredicate
}


