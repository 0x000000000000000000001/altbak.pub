package Data_EuclideanRing

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_CommutativeRing "gopurs/output/Data.CommutativeRing"
)

var mod gopurs_runtime.Value
var once_mod sync.Once
func Get_mod() gopurs_runtime.Value {
	once_mod.Do(func() {
		mod = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "mod")
})
	})
	return mod
}

var gcd gopurs_runtime.Value
var once_gcd sync.Once
func Get_gcd() gopurs_runtime.Value {
	once_gcd.Do(func() {
		gcd = gopurs_runtime.Func2(Call_gcd)
	})
	return gcd
}

var euclideanRingNumber gopurs_runtime.Value
var once_euclideanRingNumber sync.Once
func Get_euclideanRingNumber() gopurs_runtime.Value {
	once_euclideanRingNumber.Do(func() {
		euclideanRingNumber = gopurs_runtime.RecordDict4("degree", "div", "mod", "CommutativeRing0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(1)
}), Get_numDiv(), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(0.0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_CommutativeRing.Get_commutativeRingNumber()
}))
	})
	return euclideanRingNumber
}

var euclideanRingInt gopurs_runtime.Value
var once_euclideanRingInt sync.Once
func Get_euclideanRingInt() gopurs_runtime.Value {
	once_euclideanRingInt.Do(func() {
		euclideanRingInt = gopurs_runtime.RecordDict4("degree", "div", "mod", "CommutativeRing0", Get_intDegree(), Get_intDiv(), Get_intMod(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_CommutativeRing.Get_commutativeRingInt()
}))
	})
	return euclideanRingInt
}

var div gopurs_runtime.Value
var once_div sync.Once
func Get_div() gopurs_runtime.Value {
	once_div.Do(func() {
		div = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "div")
})
	})
	return div
}

var lcm gopurs_runtime.Value
var once_lcm sync.Once
func Get_lcm() gopurs_runtime.Value {
	once_lcm.Do(func() {
		lcm = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, dictEuclideanRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
Semiring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_1, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{})
_ = Semiring0_2_0
zero_3_1 := gopurs_runtime.RecordGet(Semiring0_2_0, "zero")
_ = zero_3_1
gcd2_4_2 := gopurs_runtime.Apply2(Get_gcd(), dictEq_0, dictEuclideanRing_1)
_ = gcd2_4_2
return gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), a_5, zero_3_1).IntVal != 0 || gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), b_6, zero_3_1).IntVal != 0 {
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
	})
	return lcm
}

var degree gopurs_runtime.Value
var once_degree sync.Once
func Get_degree() gopurs_runtime.Value {
	once_degree.Do(func() {
		degree = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "degree")
})
	})
	return degree
}

func Call_gcd(dictEq_0_loop gopurs_runtime.Value, dictEuclideanRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
gcd:
for {
if false { continue gcd }
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 gopurs_runtime.Value = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_1_loop, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0_loop, "eq"), b_4, zero_2_0).IntVal != 0 {
__t1 = a_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply4(Get_gcd(), dictEq_0_loop, dictEuclideanRing_1_loop, b_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_1_loop, "mod"), a_3, b_4))
}
end_branch_1:
return __t1
})
}
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
