package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
		cache_Data_Predicate_newtypePredicate = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_Predicate_newtypePredicate
}

var cache_Data_Predicate_heytingAlgebraPredicate gopurs_runtime.Value
var once_Data_Predicate_heytingAlgebraPredicate sync.Once
func Get_Data_Predicate_heytingAlgebraPredicate() gopurs_runtime.Value {
	once_Data_Predicate_heytingAlgebraPredicate.Do(func() {
		cache_Data_Predicate_heytingAlgebraPredicate = gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(&Constructor_Data_HeytingAlgebra_HeytingAlgebra{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})})}
	})
	return cache_Data_Predicate_heytingAlgebraPredicate
}

var cache_Data_Predicate_contravariantPredicate gopurs_runtime.Value
var once_Data_Predicate_contravariantPredicate sync.Once
func Get_Data_Predicate_contravariantPredicate() gopurs_runtime.Value {
	once_Data_Predicate_contravariantPredicate.Do(func() {
		cache_Data_Predicate_contravariantPredicate = gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Contravariant_Contravariant{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
})})}
	})
	return cache_Data_Predicate_contravariantPredicate
}

var cache_Data_Predicate_booleanAlgebraPredicate gopurs_runtime.Value
var once_Data_Predicate_booleanAlgebraPredicate sync.Once
func Get_Data_Predicate_booleanAlgebraPredicate() gopurs_runtime.Value {
	once_Data_Predicate_booleanAlgebraPredicate.Do(func() {
		cache_Data_Predicate_booleanAlgebraPredicate = func() gopurs_runtime.Value {
// TAST (Let): heytingAlgebraFunction_0_0 -> *Constructor_Data_HeytingAlgebra_HeytingAlgebra
heytingAlgebraFunction_0_0 := &Constructor_Data_HeytingAlgebra_HeytingAlgebra{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})}
_ = heytingAlgebraFunction_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 3257204378, UnsafePtr: unsafe.Pointer(&Constructor_Data_BooleanAlgebra_BooleanAlgebra{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(heytingAlgebraFunction_0_0)}
})})}
}()
	})
	return cache_Data_Predicate_booleanAlgebraPredicate
}

var cache_Data_Predicate_contravariantPredicate__2916341568 gopurs_runtime.Value
var once_Data_Predicate_contravariantPredicate__2916341568 sync.Once
func Get_Data_Predicate_contravariantPredicate__2916341568() gopurs_runtime.Value {
	once_Data_Predicate_contravariantPredicate__2916341568.Do(func() {
		cache_Data_Predicate_contravariantPredicate__2916341568 = gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Contravariant_Contravariant{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
})})}
	})
	return cache_Data_Predicate_contravariantPredicate__2916341568
}

func Call_Data_Predicate_Predicate(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}


