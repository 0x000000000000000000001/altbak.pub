package Data_EuclideanRing

import (
	pkg_Data_CommutativeRing "gopurs/output/Data.CommutativeRing"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_mod gopurs_runtime.Value
var once_mod sync.Once
func Get_mod() gopurs_runtime.Value {
	once_mod.Do(func() {
		cache_mod = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod(gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod
}

var cache_mod__gopurs_runtime_Value_2579358968 gopurs_runtime.Value
var once_mod__gopurs_runtime_Value_2579358968 sync.Once
func Get_mod__gopurs_runtime_Value_2579358968() gopurs_runtime.Value {
	once_mod__gopurs_runtime_Value_2579358968.Do(func() {
		cache_mod__gopurs_runtime_Value_2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__gopurs_runtime_Value_2579358968(gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod__gopurs_runtime_Value_2579358968
}

var cache_gcd_ForAll_a_ConstrainedType_Data_Eq_Eq_Any_Data_EuclideanRing_EuclideanRing_Any_Func_Any_Any_Any gopurs_runtime.Value
var once_gcd_ForAll_a_ConstrainedType_Data_Eq_Eq_Any_Data_EuclideanRing_EuclideanRing_Any_Func_Any_Any_Any sync.Once
func Get_gcd_ForAll_a_ConstrainedType_Data_Eq_Eq_Any_Data_EuclideanRing_EuclideanRing_Any_Func_Any_Any_Any() gopurs_runtime.Value {
	once_gcd_ForAll_a_ConstrainedType_Data_Eq_Eq_Any_Data_EuclideanRing_EuclideanRing_Any_Func_Any_Any_Any.Do(func() {
		cache_gcd_ForAll_a_ConstrainedType_Data_Eq_Eq_Any_Data_EuclideanRing_EuclideanRing_Any_Func_Any_Any_Any = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_gcd_ForAll_a_ConstrainedType_Data_Eq_Eq_Any_Data_EuclideanRing_EuclideanRing_Any_Func_Any_Any_Any(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dictEuclideanRing_1_box))
})
	})
	return cache_gcd_ForAll_a_ConstrainedType_Data_Eq_Eq_Any_Data_EuclideanRing_EuclideanRing_Any_Func_Any_Any_Any
}

var cache_gcd gopurs_runtime.Value
var once_gcd sync.Once
func Get_gcd() gopurs_runtime.Value {
	once_gcd.Do(func() {
		cache_gcd = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_gcd(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dictEuclideanRing_1_box))
})
	})
	return cache_gcd
}

var cache_gcd__gopurs_runtime_Value_3697052990 gopurs_runtime.Value
var once_gcd__gopurs_runtime_Value_3697052990 sync.Once
func Get_gcd__gopurs_runtime_Value_3697052990() gopurs_runtime.Value {
	once_gcd__gopurs_runtime_Value_3697052990.Do(func() {
		cache_gcd__gopurs_runtime_Value_3697052990 = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_gcd__gopurs_runtime_Value_3697052990(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dictEuclideanRing_1_box))
})
	})
	return cache_gcd__gopurs_runtime_Value_3697052990
}

var cache_euclideanRingNumber gopurs_runtime.Value
var once_euclideanRingNumber sync.Once
func Get_euclideanRingNumber() gopurs_runtime.Value {
	once_euclideanRingNumber.Do(func() {
		cache_euclideanRingNumber = gopurs_runtime.RecordDict4("CommutativeRing0", "degree", "div", "mod", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_CommutativeRing.Get_commutativeRingNumber()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(1)
}), Get_numDiv(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(0.0)
})
}))
	})
	return cache_euclideanRingNumber
}

var cache_euclideanRingInt gopurs_runtime.Value
var once_euclideanRingInt sync.Once
func Get_euclideanRingInt() gopurs_runtime.Value {
	once_euclideanRingInt.Do(func() {
		cache_euclideanRingInt = gopurs_runtime.RecordDict4("CommutativeRing0", "degree", "div", "mod", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_CommutativeRing.Get_commutativeRingInt()
}), Get_intDegree(), Get_intDiv(), Get_intMod())
	})
	return cache_euclideanRingInt
}

var cache_div gopurs_runtime.Value
var once_div sync.Once
func Get_div() gopurs_runtime.Value {
	once_div.Do(func() {
		cache_div = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div(gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div
}

var cache_div__gopurs_runtime_Value_2579358968 gopurs_runtime.Value
var once_div__gopurs_runtime_Value_2579358968 sync.Once
func Get_div__gopurs_runtime_Value_2579358968() gopurs_runtime.Value {
	once_div__gopurs_runtime_Value_2579358968.Do(func() {
		cache_div__gopurs_runtime_Value_2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__gopurs_runtime_Value_2579358968(gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__gopurs_runtime_Value_2579358968
}

var cache_lcm gopurs_runtime.Value
var once_lcm sync.Once
func Get_lcm() gopurs_runtime.Value {
	once_lcm.Do(func() {
		cache_lcm = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lcm(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dictEuclideanRing_1_box))
})
	})
	return cache_lcm
}

var cache_degree gopurs_runtime.Value
var once_degree sync.Once
func Get_degree() gopurs_runtime.Value {
	once_degree.Do(func() {
		cache_degree = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_degree(gopurs_runtime.CoerceToStruct[Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_degree
}

type Constructor_EuclideanRing[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3214993658] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_EuclideanRing[gopurs_runtime.Value])(ptr)
		switch key {
		case "CommutativeRing0": return c.V0
		case "degree": return c.V1
		case "div": return c.V2
		case "mod": return c.V3
		default: panic("Key not found in dictionary Constructor_EuclideanRing: " + key)
		}
	}
}


func Call_mod(dict_0_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_mod__gopurs_runtime_Value_2579358968(dict_0_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_gcd_ForAll_a_ConstrainedType_Data_Eq_Eq_Any_Data_EuclideanRing_EuclideanRing_Any_Func_Any_Any_Any(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], dictEuclideanRing_1_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_EuclideanRing[gopurs_runtime.Value] = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEuclideanRing_1.V0, gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(dictEq_0.V0, b_4, zero_2_0).IntVal) != (0) {
__t1 = a_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_gcd(dictEq_0, dictEuclideanRing_1), b_4, gopurs_runtime.Apply2(dictEuclideanRing_1.V3, a_3, b_4))
}
end_branch_1:
return __t1
})
})
}

func Call_gcd(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], dictEuclideanRing_1_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_EuclideanRing[gopurs_runtime.Value] = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEuclideanRing_1.V0, gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(dictEq_0.V0, b_4, zero_2_0).IntVal) != (0) {
__t1 = a_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_gcd(dictEq_0, dictEuclideanRing_1), b_4, gopurs_runtime.Apply2(dictEuclideanRing_1.V3, a_3, b_4))
}
end_branch_1:
return __t1
})
})
}

func Call_gcd__gopurs_runtime_Value_3697052990(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], dictEuclideanRing_1_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_EuclideanRing[gopurs_runtime.Value] = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEuclideanRing_1.V0, gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(dictEq_0.V0, b_4, zero_2_0).IntVal) != (0) {
__t1 = a_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_gcd(dictEq_0, dictEuclideanRing_1), b_4, gopurs_runtime.Apply2(dictEuclideanRing_1.V3, a_3, b_4))
}
end_branch_1:
return __t1
})
})
}

func Call_div(dict_0_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_div__gopurs_runtime_Value_2579358968(dict_0_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_lcm(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], dictEuclideanRing_1_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_EuclideanRing[gopurs_runtime.Value] = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
Ring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEuclideanRing_1.V0, gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{})
_ = Ring0_2_0
zero_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ring0_2_0, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_3_1
Semiring0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ring0_2_0, "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply2(dictEq_0.V0, a_5, zero_3_1), gopurs_runtime.Apply2(dictEq_0.V0, b_6, zero_3_1)).IntVal) != (0) {
__t4 = zero_3_1
goto end_branch_4
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(dictEq_0.V0, b_6, gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictEuclideanRing_1.V0, gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")).IntVal) != (0) {
__t3 = a_5
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Call_gcd(dictEq_0, dictEuclideanRing_1), b_6, gopurs_runtime.Apply2(dictEuclideanRing_1.V3, a_5, b_6))
}
end_branch_3:
__t4 = gopurs_runtime.Apply2(dictEuclideanRing_1.V2, gopurs_runtime.Apply2(Semiring0_4_2.V1, a_5, b_6), __t3)
}
end_branch_4:
return __t4
})
})
}

func Call_degree(dict_0_loop *Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Get_intDegree() gopurs_runtime.Value {
	return _Gopurs_IntDegree
}

func Get_intDiv() gopurs_runtime.Value {
	return _Gopurs_IntDiv
}

func Get_intMod() gopurs_runtime.Value {
	return _Gopurs_IntMod
}

func Get_numDiv() gopurs_runtime.Value {
	return _Gopurs_NumDiv
}
