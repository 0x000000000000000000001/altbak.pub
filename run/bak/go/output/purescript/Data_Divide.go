package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Divide_identity gopurs_runtime.Value
var once_Data_Divide_identity sync.Once
func Get_Data_Divide_identity() gopurs_runtime.Value {
	once_Data_Divide_identity.Do(func() {
		cache_Data_Divide_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Divide_identity(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](x_0_box)))}
})
	})
	return cache_Data_Divide_identity
}

var cache_Data_Divide_Divide_dollarDict gopurs_runtime.Value
var once_Data_Divide_Divide_dollarDict sync.Once
func Get_Data_Divide_Divide_dollarDict() gopurs_runtime.Value {
	once_Data_Divide_Divide_dollarDict.Do(func() {
		cache_Data_Divide_Divide_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divide_Divide_dollarDict(x_0_box)
})
	})
	return cache_Data_Divide_Divide_dollarDict
}

var cache_Data_Divide_dividePredicate gopurs_runtime.Value
var once_Data_Divide_dividePredicate sync.Once
func Get_Data_Divide_dividePredicate() gopurs_runtime.Value {
	once_Data_Divide_dividePredicate.Do(func() {
		cache_Data_Divide_dividePredicate = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(&Constructor_Data_Divide_Divide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](Get_Data_Predicate_contravariantPredicate()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_4_0 -> *Constructor_Data_Tuple_Tuple
v2_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, a_3))
_ = v2_4_0
return gopurs_runtime.Bool(((gopurs_runtime.Apply(v_1, (v2_4_0).V0).IntVal) != (0)) && ((gopurs_runtime.Apply(v1_2, (v2_4_0).V1).IntVal) != (0)))
})
})
})
})})}
	})
	return cache_Data_Divide_dividePredicate
}

var cache_Data_Divide_divideOp gopurs_runtime.Value
var once_Data_Divide_divideOp sync.Once
func Get_Data_Divide_divideOp() gopurs_runtime.Value {
	once_Data_Divide_divideOp.Do(func() {
		cache_Data_Divide_divideOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divide_divideOp(dictSemigroup_0_box)
})
	})
	return cache_Data_Divide_divideOp
}

var cache_Data_Divide_divideEquivalence gopurs_runtime.Value
var once_Data_Divide_divideEquivalence sync.Once
func Get_Data_Divide_divideEquivalence() gopurs_runtime.Value {
	once_Data_Divide_divideEquivalence.Do(func() {
		cache_Data_Divide_divideEquivalence = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(&Constructor_Data_Divide_Divide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](Get_Data_Equivalence_contravariantEquivalence()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> *Constructor_Data_Tuple_Tuple
v2_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, a_3))
_ = v2_5_0
// TAST (Let): v3_6_1 -> *Constructor_Data_Tuple_Tuple
v3_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, b_4))
_ = v3_6_1
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_1, (v2_5_0).V0, (v3_6_1).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(v1_2, (v2_5_0).V1, (v3_6_1).V1).IntVal) != (0)))
})
})
})
})
})})}
	})
	return cache_Data_Divide_divideEquivalence
}

var cache_Data_Divide_divideComparison gopurs_runtime.Value
var once_Data_Divide_divideComparison sync.Once
func Get_Data_Divide_divideComparison() gopurs_runtime.Value {
	once_Data_Divide_divideComparison.Do(func() {
		cache_Data_Divide_divideComparison = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(&Constructor_Data_Divide_Divide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](Get_Data_Comparison_contravariantComparison()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> *Constructor_Data_Tuple_Tuple
v2_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, a_3))
_ = v2_5_0
// TAST (Let): v3_6_1 -> *Constructor_Data_Tuple_Tuple
v3_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, b_4))
_ = v3_6_1
// TAST (Let): __local_var_7_2 -> uint32
__local_var_7_2 := uint32(gopurs_runtime.Apply2(v_1, (v2_5_0).V0, (v3_6_1).V0).IntVal)
_ = __local_var_7_2
// TAST (Let): __local_var_8_3 -> uint32
__local_var_8_3 := uint32(gopurs_runtime.Apply2(v1_2, (v2_5_0).V1, (v3_6_1).V1).IntVal)
_ = __local_var_8_3
var __t4 uint32
{
if (__local_var_7_2 == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (__local_var_7_2 == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if (__local_var_7_2 == 902936544) {
__t4 = __local_var_8_3
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
})
})
})
})})}
	})
	return cache_Data_Divide_divideComparison
}

var cache_Data_Divide_divide gopurs_runtime.Value
var once_Data_Divide_divide sync.Once
func Get_Data_Divide_divide() gopurs_runtime.Value {
	once_Data_Divide_divide.Do(func() {
		cache_Data_Divide_divide = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divide_divide(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](dict_0_box))
})
	})
	return cache_Data_Divide_divide
}

var cache_Data_Divide_divided gopurs_runtime.Value
var once_Data_Divide_divided sync.Once
func Get_Data_Divide_divided() gopurs_runtime.Value {
	once_Data_Divide_divided.Do(func() {
		cache_Data_Divide_divided = gopurs_runtime.Func(func(dictDivide_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divide_divided(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](dictDivide_0_box))
})
	})
	return cache_Data_Divide_divided
}

var cache_Data_Divide_divide__1446725958 gopurs_runtime.Value
var once_Data_Divide_divide__1446725958 sync.Once
func Get_Data_Divide_divide__1446725958() gopurs_runtime.Value {
	once_Data_Divide_divide__1446725958.Do(func() {
		cache_Data_Divide_divide__1446725958 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divide_divide__1446725958(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](dict_0_box))
})
	})
	return cache_Data_Divide_divide__1446725958
}

var cache_Data_Divide_divide__3365952934 gopurs_runtime.Value
var once_Data_Divide_divide__3365952934 sync.Once
func Get_Data_Divide_divide__3365952934() gopurs_runtime.Value {
	once_Data_Divide_divide__3365952934.Do(func() {
		cache_Data_Divide_divide__3365952934 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divide_divide__3365952934(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide](dict_0_box))
})
	})
	return cache_Data_Divide_divide__3365952934
}

var cache_Data_Divide_divideComparison__1969551483 gopurs_runtime.Value
var once_Data_Divide_divideComparison__1969551483 sync.Once
func Get_Data_Divide_divideComparison__1969551483() gopurs_runtime.Value {
	once_Data_Divide_divideComparison__1969551483.Do(func() {
		cache_Data_Divide_divideComparison__1969551483 = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(&Constructor_Data_Divide_Divide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](Get_Data_Comparison_contravariantComparison()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> *Constructor_Data_Tuple_Tuple
v2_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, a_3))
_ = v2_5_0
// TAST (Let): v3_6_1 -> *Constructor_Data_Tuple_Tuple
v3_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, b_4))
_ = v3_6_1
// TAST (Let): __local_var_7_2 -> uint32
__local_var_7_2 := uint32(gopurs_runtime.Apply2(v_1, (v2_5_0).V0, (v3_6_1).V0).IntVal)
_ = __local_var_7_2
// TAST (Let): __local_var_8_3 -> uint32
__local_var_8_3 := uint32(gopurs_runtime.Apply2(v1_2, (v2_5_0).V1, (v3_6_1).V1).IntVal)
_ = __local_var_8_3
var __t4 uint32
{
if (__local_var_7_2 == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (__local_var_7_2 == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if (__local_var_7_2 == 902936544) {
__t4 = __local_var_8_3
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
})
})
})
})})}
	})
	return cache_Data_Divide_divideComparison__1969551483
}

var cache_Data_Divide_divideEquivalence__1166972705 gopurs_runtime.Value
var once_Data_Divide_divideEquivalence__1166972705 sync.Once
func Get_Data_Divide_divideEquivalence__1166972705() gopurs_runtime.Value {
	once_Data_Divide_divideEquivalence__1166972705.Do(func() {
		cache_Data_Divide_divideEquivalence__1166972705 = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(&Constructor_Data_Divide_Divide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](Get_Data_Equivalence_contravariantEquivalence()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> *Constructor_Data_Tuple_Tuple
v2_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, a_3))
_ = v2_5_0
// TAST (Let): v3_6_1 -> *Constructor_Data_Tuple_Tuple
v3_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, b_4))
_ = v3_6_1
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_1, (v2_5_0).V0, (v3_6_1).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(v1_2, (v2_5_0).V1, (v3_6_1).V1).IntVal) != (0)))
})
})
})
})
})})}
	})
	return cache_Data_Divide_divideEquivalence__1166972705
}

var cache_Data_Divide_dividePredicate__2324723848 gopurs_runtime.Value
var once_Data_Divide_dividePredicate__2324723848 sync.Once
func Get_Data_Divide_dividePredicate__2324723848() gopurs_runtime.Value {
	once_Data_Divide_dividePredicate__2324723848.Do(func() {
		cache_Data_Divide_dividePredicate__2324723848 = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(&Constructor_Data_Divide_Divide{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](Get_Data_Predicate_contravariantPredicate()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_4_0 -> *Constructor_Data_Tuple_Tuple
v2_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, a_3))
_ = v2_4_0
return gopurs_runtime.Bool(((gopurs_runtime.Apply(v_1, (v2_4_0).V0).IntVal) != (0)) && ((gopurs_runtime.Apply(v1_2, (v2_4_0).V1).IntVal) != (0)))
})
})
})
})})}
	})
	return cache_Data_Divide_dividePredicate__2324723848
}

type Constructor_Data_Divide_Divide struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2642321722] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Divide_Divide)(ptr)
		_ = c
		switch key {
		case "Contravariant0": return gopurs_runtime.Box(c.V0)
		case "divide": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Divide_Divide: " + key)
		}
	}
}


func Call_Data_Divide_identity(x_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var x_0 *Constructor_Data_Tuple_Tuple = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Divide_Divide_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Divide_divideOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(&Constructor_Data_Divide_Divide{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](Get_Data_Op_contravariantOp()))}
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_5_0 -> *Constructor_Data_Tuple_Tuple
v2_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_1, a_4))
_ = v2_5_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(v_2, (v2_5_0).V0), gopurs_runtime.Apply(v1_3, (v2_5_0).V1))
})
})
})
})})}
}

func Call_Data_Divide_divide(dict_0_loop *Constructor_Data_Divide_Divide) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Divide_Divide = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Divide_divided(dictDivide_0_loop *Constructor_Data_Divide_Divide) gopurs_runtime.Value {
var dictDivide_0 *Constructor_Data_Divide_Divide = dictDivide_0_loop
_ = dictDivide_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDivide_0.V1), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Divide_divide__1446725958(dict_0_loop *Constructor_Data_Divide_Divide) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Divide_Divide = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Divide_divide__3365952934(dict_0_loop *Constructor_Data_Divide_Divide) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Divide_Divide = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}


