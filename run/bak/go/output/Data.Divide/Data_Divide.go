package Data_Divide

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Predicate "gopurs/output/Data.Predicate"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Equivalence "gopurs/output/Data.Equivalence"
	pkg_Data_Comparison "gopurs/output/Data.Comparison"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Op "gopurs/output/Data.Op"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_dividePredicate gopurs_runtime.Value
var once_dividePredicate sync.Once
func Get_dividePredicate() gopurs_runtime.Value {
	once_dividePredicate.Do(func() {
		cache_dividePredicate = gopurs_runtime.RecordDict2("Contravariant0", "divide", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Predicate.Get_contravariantPredicate()
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_4_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4_0.UnsafePtr).V0), gopurs_runtime.Apply(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4_0.UnsafePtr).V1))
}))
	})
	return cache_dividePredicate
}

var cache_divideOp gopurs_runtime.Value
var once_divideOp sync.Once
func Get_divideOp() gopurs_runtime.Value {
	once_divideOp.Do(func() {
		cache_divideOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_divideOp(dictSemigroup_0_box)
})
	})
	return cache_divideOp
}

var cache_divideEquivalence gopurs_runtime.Value
var once_divideEquivalence sync.Once
func Get_divideEquivalence() gopurs_runtime.Value {
	once_divideEquivalence.Do(func() {
		cache_divideEquivalence = gopurs_runtime.RecordDict2("Contravariant0", "divide", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Equivalence.Get_contravariantEquivalence()
}), gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_1.UnsafePtr).V0), gopurs_runtime.Apply2(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_1.UnsafePtr).V1))
}))
	})
	return cache_divideEquivalence
}

var cache_divideComparison gopurs_runtime.Value
var once_divideComparison sync.Once
func Get_divideComparison() gopurs_runtime.Value {
	once_divideComparison.Do(func() {
		cache_divideComparison = gopurs_runtime.RecordDict2("Contravariant0", "divide", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Comparison.Get_contravariantComparison()
}), gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_semigroupOrdering(), "append"), gopurs_runtime.Apply2(v_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_1.UnsafePtr).V0), gopurs_runtime.Apply2(v1_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6_1.UnsafePtr).V1))
}))
	})
	return cache_divideComparison
}

var cache_divide gopurs_runtime.Value
var once_divide sync.Once
func Get_divide() gopurs_runtime.Value {
	once_divide.Do(func() {
		cache_divide = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_divide(dict_0_box)
})
	})
	return cache_divide
}

var cache_divide__gopurs_runtime_Value_220145316 gopurs_runtime.Value
var once_divide__gopurs_runtime_Value_220145316 sync.Once
func Get_divide__gopurs_runtime_Value_220145316() gopurs_runtime.Value {
	once_divide__gopurs_runtime_Value_220145316.Do(func() {
		cache_divide__gopurs_runtime_Value_220145316 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_divide__gopurs_runtime_Value_220145316(dict_0_box)
})
	})
	return cache_divide__gopurs_runtime_Value_220145316
}

var cache_divided gopurs_runtime.Value
var once_divided sync.Once
func Get_divided() gopurs_runtime.Value {
	once_divided.Do(func() {
		cache_divided = gopurs_runtime.Func(func(dictDivide_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_divided(dictDivide_0_box)
})
	})
	return cache_divided
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_divideOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("Contravariant0", "divide", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Op.Get_contravariantOp()
}), gopurs_runtime.Func4(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_1, a_4)
_ = v2_5_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(v_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V0), gopurs_runtime.Apply(v1_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5_0.UnsafePtr).V1))
}))
}

func Call_divide(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "divide")
}

func Call_divide__gopurs_runtime_Value_220145316(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "divide")
}

func Call_divided(dictDivide_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDivide_0 gopurs_runtime.Value = dictDivide_0_loop
_ = dictDivide_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivide_0, "divide"), Get_identity())
}


