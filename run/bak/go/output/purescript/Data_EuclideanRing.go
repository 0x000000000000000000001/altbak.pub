package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_EuclideanRing_EuclideanRing_dollarDict gopurs_runtime.Value
var once_Data_EuclideanRing_EuclideanRing_dollarDict sync.Once
func Get_Data_EuclideanRing_EuclideanRing_dollarDict() gopurs_runtime.Value {
	once_Data_EuclideanRing_EuclideanRing_dollarDict.Do(func() {
		cache_Data_EuclideanRing_EuclideanRing_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_EuclideanRing_dollarDict(x_0_box)
})
	})
	return cache_Data_EuclideanRing_EuclideanRing_dollarDict
}

var cache_Data_EuclideanRing_mod gopurs_runtime.Value
var once_Data_EuclideanRing_mod sync.Once
func Get_Data_EuclideanRing_mod() gopurs_runtime.Value {
	once_Data_EuclideanRing_mod.Do(func() {
		cache_Data_EuclideanRing_mod = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_mod(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dict_0_box))
})
	})
	return cache_Data_EuclideanRing_mod
}

var cache_Data_EuclideanRing_gcd gopurs_runtime.Value
var once_Data_EuclideanRing_gcd sync.Once
func Get_Data_EuclideanRing_gcd() gopurs_runtime.Value {
	once_Data_EuclideanRing_gcd.Do(func() {
		cache_Data_EuclideanRing_gcd = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_gcd(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dictEuclideanRing_1_box))
})
	})
	return cache_Data_EuclideanRing_gcd
}

var cache_Data_EuclideanRing_euclideanRingNumber gopurs_runtime.Value
var once_Data_EuclideanRing_euclideanRingNumber sync.Once
func Get_Data_EuclideanRing_euclideanRingNumber() gopurs_runtime.Value {
	once_Data_EuclideanRing_euclideanRingNumber.Do(func() {
		cache_Data_EuclideanRing_euclideanRingNumber = gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(&Constructor_Data_EuclideanRing_EuclideanRing{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing](Get_Data_CommutativeRing_commutativeRingNumber()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(1)
}), Get_Data_EuclideanRing_numDiv(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(0.0)
})
})})}
	})
	return cache_Data_EuclideanRing_euclideanRingNumber
}

var cache_Data_EuclideanRing_euclideanRingInt gopurs_runtime.Value
var once_Data_EuclideanRing_euclideanRingInt sync.Once
func Get_Data_EuclideanRing_euclideanRingInt() gopurs_runtime.Value {
	once_Data_EuclideanRing_euclideanRingInt.Do(func() {
		cache_Data_EuclideanRing_euclideanRingInt = gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(&Constructor_Data_EuclideanRing_EuclideanRing{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing](Get_Data_CommutativeRing_commutativeRingInt()))}
}), Get_Data_EuclideanRing_intDegree(), Get_Data_EuclideanRing_intDiv(), Get_Data_EuclideanRing_intMod()})}
	})
	return cache_Data_EuclideanRing_euclideanRingInt
}

var cache_Data_EuclideanRing_div gopurs_runtime.Value
var once_Data_EuclideanRing_div sync.Once
func Get_Data_EuclideanRing_div() gopurs_runtime.Value {
	once_Data_EuclideanRing_div.Do(func() {
		cache_Data_EuclideanRing_div = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_div(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dict_0_box))
})
	})
	return cache_Data_EuclideanRing_div
}

var cache_Data_EuclideanRing_lcm gopurs_runtime.Value
var once_Data_EuclideanRing_lcm sync.Once
func Get_Data_EuclideanRing_lcm() gopurs_runtime.Value {
	once_Data_EuclideanRing_lcm.Do(func() {
		cache_Data_EuclideanRing_lcm = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_lcm(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dictEuclideanRing_1_box))
})
	})
	return cache_Data_EuclideanRing_lcm
}

var cache_Data_EuclideanRing_degree gopurs_runtime.Value
var once_Data_EuclideanRing_degree sync.Once
func Get_Data_EuclideanRing_degree() gopurs_runtime.Value {
	once_Data_EuclideanRing_degree.Do(func() {
		cache_Data_EuclideanRing_degree = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_degree(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dict_0_box))
})
	})
	return cache_Data_EuclideanRing_degree
}

var cache_Data_EuclideanRing_div__2185172824 gopurs_runtime.Value
var once_Data_EuclideanRing_div__2185172824 sync.Once
func Get_Data_EuclideanRing_div__2185172824() gopurs_runtime.Value {
	once_Data_EuclideanRing_div__2185172824.Do(func() {
		cache_Data_EuclideanRing_div__2185172824 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_div__2185172824(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_EuclideanRing_div__2185172824
}

var cache_Data_EuclideanRing_div__1002719800 gopurs_runtime.Value
var once_Data_EuclideanRing_div__1002719800 sync.Once
func Get_Data_EuclideanRing_div__1002719800() gopurs_runtime.Value {
	once_Data_EuclideanRing_div__1002719800.Do(func() {
		cache_Data_EuclideanRing_div__1002719800 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_div__1002719800(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_EuclideanRing_div__1002719800
}

var cache_Data_EuclideanRing_div__2579358968 gopurs_runtime.Value
var once_Data_EuclideanRing_div__2579358968 sync.Once
func Get_Data_EuclideanRing_div__2579358968() gopurs_runtime.Value {
	once_Data_EuclideanRing_div__2579358968.Do(func() {
		cache_Data_EuclideanRing_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_div__2579358968(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dict_0_box))
})
	})
	return cache_Data_EuclideanRing_div__2579358968
}

var cache_Data_EuclideanRing_gcd__3697052990 gopurs_runtime.Value
var once_Data_EuclideanRing_gcd__3697052990 sync.Once
func Get_Data_EuclideanRing_gcd__3697052990() gopurs_runtime.Value {
	once_Data_EuclideanRing_gcd__3697052990.Do(func() {
		cache_Data_EuclideanRing_gcd__3697052990 = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_gcd__3697052990(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dictEuclideanRing_1_box))
})
	})
	return cache_Data_EuclideanRing_gcd__3697052990
}

var cache_Data_EuclideanRing_mod__2185172824 gopurs_runtime.Value
var once_Data_EuclideanRing_mod__2185172824 sync.Once
func Get_Data_EuclideanRing_mod__2185172824() gopurs_runtime.Value {
	once_Data_EuclideanRing_mod__2185172824.Do(func() {
		cache_Data_EuclideanRing_mod__2185172824 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_mod__2185172824(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_EuclideanRing_mod__2185172824
}

var cache_Data_EuclideanRing_mod__2579358968 gopurs_runtime.Value
var once_Data_EuclideanRing_mod__2579358968 sync.Once
func Get_Data_EuclideanRing_mod__2579358968() gopurs_runtime.Value {
	once_Data_EuclideanRing_mod__2579358968.Do(func() {
		cache_Data_EuclideanRing_mod__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_mod__2579358968(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dict_0_box))
})
	})
	return cache_Data_EuclideanRing_mod__2579358968
}

type Constructor_Data_EuclideanRing_EuclideanRing struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3214993658] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_EuclideanRing_EuclideanRing)(ptr)
		_ = c
		switch key {
		case "CommutativeRing0": return gopurs_runtime.Box(c.V0)
		case "degree": return gopurs_runtime.Box(c.V1)
		case "div": return gopurs_runtime.Box(c.V2)
		case "mod": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_EuclideanRing_EuclideanRing: " + key)
		}
	}
}


func Call_Data_EuclideanRing_EuclideanRing_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_EuclideanRing_mod(dict_0_loop *Constructor_Data_EuclideanRing_EuclideanRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_EuclideanRing_EuclideanRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_EuclideanRing_gcd(dictEq_0_loop *Constructor_Data_Eq_Eq, dictEuclideanRing_1_loop *Constructor_Data_EuclideanRing_EuclideanRing) gopurs_runtime.Value {
gcd:
for {
if false { continue gcd }
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_Data_EuclideanRing_EuclideanRing = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
// TAST (Let): zero_2_0 -> gopurs_runtime.Value
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEuclideanRing_1.V0), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), b_4, zero_2_0).IntVal) != (0) {
__t1 = a_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_Data_EuclideanRing_gcd(dictEq_0, dictEuclideanRing_1), b_4, gopurs_runtime.Apply2(gopurs_runtime.Box(dictEuclideanRing_1.V3), a_3, b_4))
}
end_branch_1:
return __t1
})
})
}
}

func Call_Data_EuclideanRing_div(dict_0_loop *Constructor_Data_EuclideanRing_EuclideanRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_EuclideanRing_EuclideanRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_EuclideanRing_lcm(dictEq_0_loop *Constructor_Data_Eq_Eq, dictEuclideanRing_1_loop *Constructor_Data_EuclideanRing_EuclideanRing) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_Data_EuclideanRing_EuclideanRing = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
// TAST (Let): Ring0_2_0 -> gopurs_runtime.Value
Ring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEuclideanRing_1.V0), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{})
_ = Ring0_2_0
// TAST (Let): zero_3_1 -> gopurs_runtime.Value
zero_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ring0_2_0, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_3_1
// TAST (Let): Semiring0_4_2 -> *Constructor_Data_Semiring_Semiring
Semiring0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ring0_2_0, "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), a_5, zero_3_1).IntVal) != (0)) || ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), b_6, zero_3_1).IntVal) != (0)) {
__t4 = zero_3_1
goto end_branch_4
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), b_6, gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEuclideanRing_1.V0), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")).IntVal) != (0) {
__t3 = a_5
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Call_Data_EuclideanRing_gcd(dictEq_0, dictEuclideanRing_1), b_6, gopurs_runtime.Apply2(gopurs_runtime.Box(dictEuclideanRing_1.V3), a_5, b_6))
}
end_branch_3:
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictEuclideanRing_1.V2), gopurs_runtime.Apply2(gopurs_runtime.Box(Semiring0_4_2.V1), a_5, b_6), __t3)
}
end_branch_4:
return __t4
})
})
}

func Call_Data_EuclideanRing_degree(dict_0_loop *Constructor_Data_EuclideanRing_EuclideanRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_EuclideanRing_EuclideanRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_EuclideanRing_div__2185172824(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) / (__eta1_1.IntVal))
}

func Call_Data_EuclideanRing_div__1002719800(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) / (__eta1_1.FloatVal()))
}

func Call_Data_EuclideanRing_div__2579358968(dict_0_loop *Constructor_Data_EuclideanRing_EuclideanRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_EuclideanRing_EuclideanRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_EuclideanRing_gcd__3697052990(dictEq_0_loop *Constructor_Data_Eq_Eq, dictEuclideanRing_1_loop *Constructor_Data_EuclideanRing_EuclideanRing) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_Data_EuclideanRing_EuclideanRing = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
// TAST (Let): zero_2_0 -> gopurs_runtime.Value
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEuclideanRing_1.V0), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), b_4, zero_2_0).IntVal) != (0) {
__t3 = a_3
goto end_branch_3
} else {

}
}
{
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictEuclideanRing_1.V3), a_3, b_4)
_ = __local_var_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), __local_var_5_1, gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEuclideanRing_1.V0), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")).IntVal) != (0) {
__t2 = b_4
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(Call_Data_EuclideanRing_gcd(dictEq_0, dictEuclideanRing_1), __local_var_5_1, gopurs_runtime.Apply2(gopurs_runtime.Box(dictEuclideanRing_1.V3), b_4, __local_var_5_1))
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
})
})
}

func Call_Data_EuclideanRing_mod__2185172824(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), __eta0_0, __eta1_1)
}

func Call_Data_EuclideanRing_mod__2579358968(dict_0_loop *Constructor_Data_EuclideanRing_EuclideanRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_EuclideanRing_EuclideanRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Get_Data_EuclideanRing_intDegree() gopurs_runtime.Value {
	return _Gopurs_Data_EuclideanRing_IntDegree
}

func Get_Data_EuclideanRing_intDiv() gopurs_runtime.Value {
	return _Gopurs_Data_EuclideanRing_IntDiv
}

func Get_Data_EuclideanRing_intMod() gopurs_runtime.Value {
	return _Gopurs_Data_EuclideanRing_IntMod
}

func Get_Data_EuclideanRing_numDiv() gopurs_runtime.Value {
	return _Gopurs_Data_EuclideanRing_NumDiv
}
