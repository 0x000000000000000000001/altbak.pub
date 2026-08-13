package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Predicate_Predicate gopurs_runtime.Value
var once_Data_Predicate_Predicate sync.Once
func Get_Data_Predicate_Predicate() gopurs_runtime.Value {
	once_Data_Predicate_Predicate.Do(func() {
		cache_Data_Predicate_Predicate = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Predicate_Predicate(x_0_box)
})
	})
	return cache_Data_Predicate_Predicate
}

var cache_Data_Predicate_newtypePredicate gopurs_runtime.Value
var once_Data_Predicate_newtypePredicate sync.Once
func Get_Data_Predicate_newtypePredicate() gopurs_runtime.Value {
	once_Data_Predicate_newtypePredicate.Do(func() {
		cache_Data_Predicate_newtypePredicate = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Predicate_newtypePredicate
}

var cache_Data_Predicate_heytingAlgebraPredicate gopurs_runtime.Value
var once_Data_Predicate_heytingAlgebraPredicate sync.Once
func Get_Data_Predicate_heytingAlgebraPredicate() gopurs_runtime.Value {
	once_Data_Predicate_heytingAlgebraPredicate.Do(func() {
		cache_Data_Predicate_heytingAlgebraPredicate = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) && ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) || ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) != (true)) || ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_1).IntVal) != (0)) != (true))
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})})
	})
	return cache_Data_Predicate_heytingAlgebraPredicate
}

var cache_Data_Predicate_contravariantPredicate gopurs_runtime.Value
var once_Data_Predicate_contravariantPredicate sync.Once
func Get_Data_Predicate_contravariantPredicate() gopurs_runtime.Value {
	once_Data_Predicate_contravariantPredicate.Do(func() {
		cache_Data_Predicate_contravariantPredicate = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}))
	})
	return cache_Data_Predicate_contravariantPredicate
}

var cache_Data_Predicate_booleanAlgebraPredicate gopurs_runtime.Value
var once_Data_Predicate_booleanAlgebraPredicate sync.Once
func Get_Data_Predicate_booleanAlgebraPredicate() gopurs_runtime.Value {
	once_Data_Predicate_booleanAlgebraPredicate.Do(func() {
		cache_Data_Predicate_booleanAlgebraPredicate = gopurs_runtime.Apply(Get_Data_BooleanAlgebra_booleanAlgebraFn(), Get_Data_BooleanAlgebra_booleanAlgebraBoolean())
	})
	return cache_Data_Predicate_booleanAlgebraPredicate
}

var cache_Data_Predicate_contravariantPredicate__2354513683 gopurs_runtime.Value
var once_Data_Predicate_contravariantPredicate__2354513683 sync.Once
func Get_Data_Predicate_contravariantPredicate__2354513683() gopurs_runtime.Value {
	once_Data_Predicate_contravariantPredicate__2354513683.Do(func() {
		cache_Data_Predicate_contravariantPredicate__2354513683 = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}))
	})
	return cache_Data_Predicate_contravariantPredicate__2354513683
}

func Call_Data_Predicate_Predicate(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}


