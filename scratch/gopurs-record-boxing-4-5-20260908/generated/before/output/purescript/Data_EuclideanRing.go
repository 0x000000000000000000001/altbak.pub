package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_EuclideanRing_EuclideanRing_dollar_Dict gopurs_runtime.Value
var once_Data_EuclideanRing_EuclideanRing_dollar_Dict sync.Once
func Get_Data_EuclideanRing_EuclideanRing_dollar_Dict() gopurs_runtime.Value {
	once_Data_EuclideanRing_EuclideanRing_dollar_Dict.Do(func() {
		cache_Data_EuclideanRing_EuclideanRing_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(Call_Data_EuclideanRing_EuclideanRing_dollar_Dict(func() struct{
	CommutativeRing0 gopurs_runtime.Value
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	CommutativeRing0 gopurs_runtime.Value
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
}{}
					clone.CommutativeRing0 = gopurs_runtime.RecordGet(orig, "CommutativeRing0")
					clone.degree = gopurs_runtime.RecordGet(orig, "degree")
					clone.div = gopurs_runtime.RecordGet(orig, "div")
					clone.mod = gopurs_runtime.RecordGet(orig, "mod")
					return clone
				}()))}
})
	})
	return cache_Data_EuclideanRing_EuclideanRing_dollar_Dict
}

var cache_Data_EuclideanRing_mod gopurs_runtime.Value
var once_Data_EuclideanRing_mod sync.Once
func Get_Data_EuclideanRing_mod() gopurs_runtime.Value {
	once_Data_EuclideanRing_mod.Do(func() {
		cache_Data_EuclideanRing_mod = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_mod(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_EuclideanRing_mod
}

var cache_Data_EuclideanRing_mod__3052583336 gopurs_runtime.Value
var once_Data_EuclideanRing_mod__3052583336 sync.Once
func Get_Data_EuclideanRing_mod__3052583336() gopurs_runtime.Value {
	once_Data_EuclideanRing_mod__3052583336.Do(func() {
		cache_Data_EuclideanRing_mod__3052583336 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_mod__3052583336(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_EuclideanRing_mod__3052583336
}

var cache_Data_EuclideanRing_gcd gopurs_runtime.Value
var once_Data_EuclideanRing_gcd sync.Once
func Get_Data_EuclideanRing_gcd() gopurs_runtime.Value {
	once_Data_EuclideanRing_gcd.Do(func() {
		cache_Data_EuclideanRing_gcd = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_gcd(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]](dictEuclideanRing_1_box))
})
	})
	return cache_Data_EuclideanRing_gcd
}

var cache_Data_EuclideanRing_euclideanRingNumber gopurs_runtime.Value
var once_Data_EuclideanRing_euclideanRingNumber sync.Once
func Get_Data_EuclideanRing_euclideanRingNumber() gopurs_runtime.Value {
	once_Data_EuclideanRing_euclideanRingNumber.Do(func() {
		cache_Data_EuclideanRing_euclideanRingNumber = gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(Rebox_Data_EuclideanRing_2521013494_1774031598((&Constructor_Data_EuclideanRing_EuclideanRing[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Rebox_Data_EuclideanRing_2766724982_1073849710(gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing[float64]](Get_Data_CommutativeRing_commutativeRingNumber())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(int64(1))
}), Get_Data_EuclideanRing_numDiv(), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(0.0)
})})))}
	})
	return cache_Data_EuclideanRing_euclideanRingNumber
}

var cache_Data_EuclideanRing_euclideanRingInt gopurs_runtime.Value
var once_Data_EuclideanRing_euclideanRingInt sync.Once
func Get_Data_EuclideanRing_euclideanRingInt() gopurs_runtime.Value {
	once_Data_EuclideanRing_euclideanRingInt.Do(func() {
		cache_Data_EuclideanRing_euclideanRingInt = gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(Rebox_Data_EuclideanRing_166963253_1774031598((&Constructor_Data_EuclideanRing_EuclideanRing[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Rebox_Data_EuclideanRing_357988789_1073849710(gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing[int64]](Get_Data_CommutativeRing_commutativeRingInt())))}
}), Get_Data_EuclideanRing_intDegree(), Get_Data_EuclideanRing_intDiv(), Get_Data_EuclideanRing_intMod()})))}
	})
	return cache_Data_EuclideanRing_euclideanRingInt
}

var cache_Data_EuclideanRing_div gopurs_runtime.Value
var once_Data_EuclideanRing_div sync.Once
func Get_Data_EuclideanRing_div() gopurs_runtime.Value {
	once_Data_EuclideanRing_div.Do(func() {
		cache_Data_EuclideanRing_div = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_div(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_EuclideanRing_div
}

var cache_Data_EuclideanRing_div__3052583336 gopurs_runtime.Value
var once_Data_EuclideanRing_div__3052583336 sync.Once
func Get_Data_EuclideanRing_div__3052583336() gopurs_runtime.Value {
	once_Data_EuclideanRing_div__3052583336.Do(func() {
		cache_Data_EuclideanRing_div__3052583336 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_div__3052583336(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_EuclideanRing_div__3052583336
}

var cache_Data_EuclideanRing_div__294757560 gopurs_runtime.Value
var once_Data_EuclideanRing_div__294757560 sync.Once
func Get_Data_EuclideanRing_div__294757560() gopurs_runtime.Value {
	once_Data_EuclideanRing_div__294757560.Do(func() {
		cache_Data_EuclideanRing_div__294757560 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_div__294757560(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_EuclideanRing_div__294757560
}

var cache_Data_EuclideanRing_lcm gopurs_runtime.Value
var once_Data_EuclideanRing_lcm sync.Once
func Get_Data_EuclideanRing_lcm() gopurs_runtime.Value {
	once_Data_EuclideanRing_lcm.Do(func() {
		cache_Data_EuclideanRing_lcm = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEuclideanRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_lcm(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]](dictEuclideanRing_1_box))
})
	})
	return cache_Data_EuclideanRing_lcm
}

var cache_Data_EuclideanRing_degree gopurs_runtime.Value
var once_Data_EuclideanRing_degree sync.Once
func Get_Data_EuclideanRing_degree() gopurs_runtime.Value {
	once_Data_EuclideanRing_degree.Do(func() {
		cache_Data_EuclideanRing_degree = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_EuclideanRing_degree(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_EuclideanRing_degree
}

type Constructor_Data_EuclideanRing_EuclideanRing[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3214993658] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_EuclideanRing_EuclideanRing[any])(ptr)
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


func Call_Data_EuclideanRing_EuclideanRing_dollar_Dict(x_0_loop struct{
	CommutativeRing0 gopurs_runtime.Value
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
}) *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] {
var x_0 struct{
	CommutativeRing0 gopurs_runtime.Value
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"CommutativeRing0", "degree", "div", "mod"}, []gopurs_runtime.Value{orig.CommutativeRing0, orig.degree, orig.div, orig.mod})
				}())
}

func Call_Data_EuclideanRing_mod(dict_0_loop *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_EuclideanRing_mod__3052583336(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
mod__3052583336:
for {
if false { continue mod__3052583336 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Int((__eta_norm_1_0.IntVal) % (__eta_norm_0_1.IntVal))
}
}

func Call_Data_EuclideanRing_gcd(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], dictEuclideanRing_1_loop *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
gcd:
for {
if false { continue gcd }
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
// TAST (Let): zero_2_0 shape=Other bindingType=(TypeVar a)
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEuclideanRing_1.V0), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
}
}

func Call_Data_EuclideanRing_div(dict_0_loop *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_EuclideanRing_div__3052583336(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
div__3052583336:
for {
if false { continue div__3052583336 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Int((__eta_norm_1_0.IntVal) / (__eta_norm_0_1.IntVal))
}
}

func Call_Data_EuclideanRing_div__294757560(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
div__294757560:
for {
if false { continue div__294757560 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Float((__eta_norm_1_0.FloatVal()) / (__eta_norm_0_1.FloatVal()))
}
}

func Call_Data_EuclideanRing_lcm(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], dictEuclideanRing_1_loop *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var dictEuclideanRing_1 *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] = dictEuclideanRing_1_loop
_ = dictEuclideanRing_1
// TAST (Let): Ring0_2_0 shape=App(Other) bindingType=Any
Ring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEuclideanRing_1.V0), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{})
_ = Ring0_2_0
// TAST (Let): zero_3_1 shape=Other bindingType=(TypeVar a)
zero_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ring0_2_0, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_3_1
// TAST (Let): Semiring0_4_2 shape=App(Other) bindingType=(ADT ["Data","Semiring","Semiring"] [(TypeVar a)])
Semiring0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ring0_2_0, "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_4_2
return gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), a_5, zero_3_1).IntVal) != (0)) || ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), b_6, zero_3_1).IntVal) != (0)) {
__t3 = zero_3_1
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictEuclideanRing_1.V2), gopurs_runtime.Apply2(gopurs_runtime.Box(Semiring0_4_2.V1), a_5, b_6), gopurs_runtime.Apply2(Call_Data_EuclideanRing_gcd(dictEq_0, dictEuclideanRing_1), a_5, b_6))
}
end_branch_3:
return __t3
})
}

func Call_Data_EuclideanRing_degree(dict_0_loop *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Rebox_Data_EuclideanRing_166963253_1774031598(in *Constructor_Data_EuclideanRing_EuclideanRing[int64]) *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_EuclideanRing_2521013494_1774031598(in *Constructor_Data_EuclideanRing_EuclideanRing[float64]) *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_EuclideanRing_2766724982_1073849710(in *Constructor_Data_CommutativeRing_CommutativeRing[float64]) *Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_EuclideanRing_357988789_1073849710(in *Constructor_Data_CommutativeRing_CommutativeRing[int64]) *Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
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
