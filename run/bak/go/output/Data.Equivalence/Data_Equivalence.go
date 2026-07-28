package Data_Equivalence

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
)

var cache_Equivalence gopurs_runtime.Value
var once_Equivalence sync.Once
func Get_Equivalence() gopurs_runtime.Value {
	once_Equivalence.Do(func() {
		cache_Equivalence = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Equivalence(x_0_box)
})
	})
	return cache_Equivalence
}

var cache_semigroupEquivalence gopurs_runtime.Value
var once_semigroupEquivalence sync.Once
func Get_semigroupEquivalence() gopurs_runtime.Value {
	once_semigroupEquivalence.Do(func() {
		cache_semigroupEquivalence = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(v_0, a_2, b_3), gopurs_runtime.Apply2(v1_1, a_2, b_3))
}))
	})
	return cache_semigroupEquivalence
}

var cache_newtypeEquivalence gopurs_runtime.Value
var once_newtypeEquivalence sync.Once
func Get_newtypeEquivalence() gopurs_runtime.Value {
	once_newtypeEquivalence.Do(func() {
		cache_newtypeEquivalence = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeEquivalence
}

var cache_monoidEquivalence gopurs_runtime.Value
var once_monoidEquivalence sync.Once
func Get_monoidEquivalence() gopurs_runtime.Value {
	once_monoidEquivalence.Do(func() {
		cache_monoidEquivalence = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupEquivalence()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return cache_monoidEquivalence
}

var cache_defaultEquivalence gopurs_runtime.Value
var once_defaultEquivalence sync.Once
func Get_defaultEquivalence() gopurs_runtime.Value {
	once_defaultEquivalence.Do(func() {
		cache_defaultEquivalence = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultEquivalence(dictEq_0_box)
})
	})
	return cache_defaultEquivalence
}

var cache_contravariantEquivalence gopurs_runtime.Value
var once_contravariantEquivalence sync.Once
func Get_contravariantEquivalence() gopurs_runtime.Value {
	once_contravariantEquivalence.Do(func() {
		cache_contravariantEquivalence = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
}))
	})
	return cache_contravariantEquivalence
}

var cache_comparisonEquivalence gopurs_runtime.Value
var once_comparisonEquivalence sync.Once
func Get_comparisonEquivalence() gopurs_runtime.Value {
	once_comparisonEquivalence.Do(func() {
		cache_comparisonEquivalence = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_comparisonEquivalence(v_0_box, a_1_box, b_2_box))
})
	})
	return cache_comparisonEquivalence
}

func Call_Equivalence(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_defaultEquivalence(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}

func Call_comparisonEquivalence(v_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) bool {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(v_0, a_1, b_2), gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}).IntVal) != (0)
}


