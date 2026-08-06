package Data_EuclideanRing

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_CommutativeRing "gopurs/output/Data.CommutativeRing"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var cache_mod gopurs_runtime.Value
var once_mod sync.Once
func Get_mod() gopurs_runtime.Value {
	once_mod.Do(func() {
		cache_mod = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod(dict_0_box)
})
	})
	return cache_mod
}

var cache_mod__gopurs_runtime_Value_2724609554 gopurs_runtime.Value
var once_mod__gopurs_runtime_Value_2724609554 sync.Once
func Get_mod__gopurs_runtime_Value_2724609554() gopurs_runtime.Value {
	once_mod__gopurs_runtime_Value_2724609554.Do(func() {
		cache_mod__gopurs_runtime_Value_2724609554 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__gopurs_runtime_Value_2724609554(dict_0_box)
})
	})
	return cache_mod__gopurs_runtime_Value_2724609554
}

var cache_gcd gopurs_runtime.Value
var once_gcd sync.Once
func Get_gcd() gopurs_runtime.Value {
	once_gcd.Do(func() {
		cache_gcd = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_gcd(dictEq_0_box, dictEuclideanRing_1_box)
})
	})
	return cache_gcd
}

var cache_gcd__gopurs_runtime_Value_3616179153 gopurs_runtime.Value
var once_gcd__gopurs_runtime_Value_3616179153 sync.Once
func Get_gcd__gopurs_runtime_Value_3616179153() gopurs_runtime.Value {
	once_gcd__gopurs_runtime_Value_3616179153.Do(func() {
		cache_gcd__gopurs_runtime_Value_3616179153 = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_gcd__gopurs_runtime_Value_3616179153(dictEq_0_box, dictEuclideanRing_1_box)
})
	})
	return cache_gcd__gopurs_runtime_Value_3616179153
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
return Call_div(dict_0_box)
})
	})
	return cache_div
}

var cache_div__gopurs_runtime_Value_2724609554 gopurs_runtime.Value
var once_div__gopurs_runtime_Value_2724609554 sync.Once
func Get_div__gopurs_runtime_Value_2724609554() gopurs_runtime.Value {
	once_div__gopurs_runtime_Value_2724609554.Do(func() {
		cache_div__gopurs_runtime_Value_2724609554 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__gopurs_runtime_Value_2724609554(dict_0_box)
})
	})
	return cache_div__gopurs_runtime_Value_2724609554
}

var cache_lcm gopurs_runtime.Value
var once_lcm sync.Once
func Get_lcm() gopurs_runtime.Value {
	once_lcm.Do(func() {
		cache_lcm = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lcm(dictEq_0_box, dictEuclideanRing_1_box)
})
	})
	return cache_lcm
}

var cache_degree gopurs_runtime.Value
var once_degree sync.Once
func Get_degree() gopurs_runtime.Value {
	once_degree.Do(func() {
		cache_degree = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_degree(dict_0_box)
})
	})
	return cache_degree
}

func Call_mod(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mod")
}

func Call_mod__gopurs_runtime_Value_2724609554(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mod")
}

func Call_gcd(dictEq_0_loop gopurs_runtime.Value, dictEuclideanRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 gopurs_runtime.Value = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_1, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), b_4, zero_2_0).IntVal) != (0) {
__t1 = a_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_gcd(dictEq_0, dictEuclideanRing_1), b_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_1, "mod"), a_3, b_4))
}
end_branch_1:
return __t1
})
})
}

func Call_gcd__gopurs_runtime_Value_3616179153(dictEq_0_loop gopurs_runtime.Value, dictEuclideanRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 gopurs_runtime.Value = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_1, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), b_4, zero_2_0).IntVal) != (0) {
__t1 = a_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_gcd(dictEq_0, dictEuclideanRing_1), b_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_1, "mod"), a_3, b_4))
}
end_branch_1:
return __t1
})
})
}

func Call_div(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "div")
}

func Call_div__gopurs_runtime_Value_2724609554(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "div")
}

func Call_lcm(dictEq_0_loop gopurs_runtime.Value, dictEuclideanRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 gopurs_runtime.Value = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
Semiring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_1, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{})
_ = Semiring0_2_0
zero_3_1 := gopurs_runtime.RecordGet(Semiring0_2_0, "zero")
_ = zero_3_1
gcd2_4_2 := Call_gcd(dictEq_0, dictEuclideanRing_1)
_ = gcd2_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), a_5, zero_3_1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), b_6, zero_3_1)).IntVal) != (0) {
__t3 = zero_3_1
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_1, "div"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Semiring0_2_0, "mul"), a_5, b_6), gopurs_runtime.Apply2(gcd2_4_2, a_5, b_6))
}
end_branch_3:
return __t3
})
})
}

func Call_degree(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "degree")
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
