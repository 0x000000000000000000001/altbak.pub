package Data_Divide

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Predicate "gopurs/output/Data.Predicate"
	pkg_Data_Op "gopurs/output/Data.Op"
	pkg_Data_Equivalence "gopurs/output/Data.Equivalence"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Comparison "gopurs/output/Data.Comparison"
	pkg_Control_Category "gopurs/output/Control.Category"
	unsafe "unsafe"
)

var cache_dividePredicate gopurs_runtime.Value
var once_dividePredicate sync.Once
func Get_dividePredicate() gopurs_runtime.Value {
	once_dividePredicate.Do(func() {
		cache_dividePredicate = gopurs_runtime.RecordDict2("divide", "Contravariant0", gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_4_0
return gopurs_runtime.Bool(((gopurs_runtime.Apply(v_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v2_4_0.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply(v1_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v2_4_0.UnsafePtr).V1).IntVal) != (0)))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Predicate.Get_contravariantPredicate()
}))
	})
	return cache_dividePredicate
}

var cache_divideOp gopurs_runtime.Value
var once_divideOp sync.Once
func Get_divideOp() gopurs_runtime.Value {
	once_divideOp.Do(func() {
		cache_divideOp = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("divide", "Contravariant0", gopurs_runtime.Func4(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_1, a_4)
_ = v2_5_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(v_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v2_5_0.UnsafePtr).V0), gopurs_runtime.Apply(v1_3, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v2_5_0.UnsafePtr).V1))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Op.Get_contravariantOp()
}))
}()
})
	})
	return cache_divideOp
}

var cache_divideEquivalence gopurs_runtime.Value
var once_divideEquivalence sync.Once
func Get_divideEquivalence() gopurs_runtime.Value {
	once_divideEquivalence.Do(func() {
		cache_divideEquivalence = gopurs_runtime.RecordDict2("divide", "Contravariant0", gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v2_5_0.UnsafePtr).V0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v3_6_1.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(v1_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v2_5_0.UnsafePtr).V1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v3_6_1.UnsafePtr).V1).IntVal) != (0)))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Equivalence.Get_contravariantEquivalence()
}))
	})
	return cache_divideEquivalence
}

var cache_divideComparison gopurs_runtime.Value
var once_divideComparison sync.Once
func Get_divideComparison() gopurs_runtime.Value {
	once_divideComparison.Do(func() {
		cache_divideComparison = gopurs_runtime.RecordDict2("divide", "Contravariant0", gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
_ = v2_5_0
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
_ = v3_6_1
__local_var_7_2 := gopurs_runtime.Apply2(v_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v2_5_0.UnsafePtr).V0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v3_6_1.UnsafePtr).V0)
_ = __local_var_7_2
__local_var_8_3 := gopurs_runtime.Apply2(v1_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v2_5_0.UnsafePtr).V1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v3_6_1.UnsafePtr).V1)
_ = __local_var_8_3
var __t4 gopurs_runtime.Value
{
if (__local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_4
} else {

}
}
{
if (__local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
if (__local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 902936544) {
__t4 = __local_var_8_3
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Comparison.Get_contravariantComparison()
}))
	})
	return cache_divideComparison
}

var cache_divide gopurs_runtime.Value
var once_divide sync.Once
func Get_divide() gopurs_runtime.Value {
	once_divide.Do(func() {
		cache_divide = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "divide")
}()
})
	})
	return cache_divide
}

var cache_divided gopurs_runtime.Value
var once_divided sync.Once
func Get_divided() gopurs_runtime.Value {
	once_divided.Do(func() {
		cache_divided = gopurs_runtime.Func(func(dictDivide_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictDivide_0 gopurs_runtime.Value = dictDivide_0_loop
_ = dictDivide_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivide_0, "divide"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}()
})
	})
	return cache_divided
}




