package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Int_top gopurs_runtime.Value
var once_Data_Int_top sync.Once
func Get_Data_Int_top() gopurs_runtime.Value {
	once_Data_Int_top.Do(func() {
		cache_Data_Int_top = gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal)
	})
	return cache_Data_Int_top
}

var cache_Data_Int_bottom gopurs_runtime.Value
var once_Data_Int_bottom sync.Once
func Get_Data_Int_bottom() gopurs_runtime.Value {
	once_Data_Int_bottom.Do(func() {
		cache_Data_Int_bottom = gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal)
	})
	return cache_Data_Int_bottom
}

var cache_Data_Int_Radix gopurs_runtime.Value
var once_Data_Int_Radix sync.Once
func Get_Data_Int_Radix() gopurs_runtime.Value {
	once_Data_Int_Radix.Do(func() {
		cache_Data_Int_Radix = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Int_Radix(x_0_box.IntVal))
})
	})
	return cache_Data_Int_Radix
}

var cache_Data_Int_Even gopurs_runtime.Value
var once_Data_Int_Even sync.Once
func Get_Data_Int_Even() gopurs_runtime.Value {
	once_Data_Int_Even.Do(func() {
		cache_Data_Int_Even = gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}
	})
	return cache_Data_Int_Even
}

var cache_Data_Int_Odd gopurs_runtime.Value
var once_Data_Int_Odd sync.Once
func Get_Data_Int_Odd() gopurs_runtime.Value {
	once_Data_Int_Odd.Do(func() {
		cache_Data_Int_Odd = gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}
	})
	return cache_Data_Int_Odd
}

var cache_Data_Int_showParity gopurs_runtime.Value
var once_Data_Int_showParity sync.Once
func Get_Data_Int_showParity() gopurs_runtime.Value {
	once_Data_Int_showParity.Do(func() {
		cache_Data_Int_showParity = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 string
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 2591059121) {
__t2 = "Even"
goto end_branch_2
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 658452902) {
__t2 = "Odd"
goto end_branch_2
} else {

}
}
{
__t2 = func() string { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Str(__t2)
})})))}
	})
	return cache_Data_Int_showParity
}

var cache_Data_Int_radix gopurs_runtime.Value
var once_Data_Int_radix sync.Once
func Get_Data_Int_radix() gopurs_runtime.Value {
	once_Data_Int_radix.Do(func() {
		cache_Data_Int_radix = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Int_radix(n_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Int_radix
}

var cache_Data_Int_odd gopurs_runtime.Value
var once_Data_Int_odd sync.Once
func Get_Data_Int_odd() gopurs_runtime.Value {
	once_Data_Int_odd.Do(func() {
		cache_Data_Int_odd = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Int_odd(x_0_box.IntVal))
})
	})
	return cache_Data_Int_odd
}

var cache_Data_Int_octal gopurs_runtime.Value
var once_Data_Int_octal sync.Once
func Get_Data_Int_octal() gopurs_runtime.Value {
	once_Data_Int_octal.Do(func() {
		cache_Data_Int_octal = gopurs_runtime.Int(int64(8))
	})
	return cache_Data_Int_octal
}

var cache_Data_Int_hexadecimal gopurs_runtime.Value
var once_Data_Int_hexadecimal sync.Once
func Get_Data_Int_hexadecimal() gopurs_runtime.Value {
	once_Data_Int_hexadecimal.Do(func() {
		cache_Data_Int_hexadecimal = gopurs_runtime.Int(int64(16))
	})
	return cache_Data_Int_hexadecimal
}

var cache_Data_Int_fromStringAs gopurs_runtime.Value
var once_Data_Int_fromStringAs sync.Once
func Get_Data_Int_fromStringAs() gopurs_runtime.Value {
	once_Data_Int_fromStringAs.Do(func() {
		cache_Data_Int_fromStringAs = gopurs_runtime.Apply2(Get_Data_Int_fromStringAsImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))})
	})
	return cache_Data_Int_fromStringAs
}

var cache_Data_Int_fromString gopurs_runtime.Value
var once_Data_Int_fromString sync.Once
func Get_Data_Int_fromString() gopurs_runtime.Value {
	once_Data_Int_fromString.Do(func() {
		cache_Data_Int_fromString = gopurs_runtime.Apply(Get_Data_Int_fromStringAs(), gopurs_runtime.Int(int64(10)))
	})
	return cache_Data_Int_fromString
}

var cache_Data_Int_fromNumber gopurs_runtime.Value
var once_Data_Int_fromNumber sync.Once
func Get_Data_Int_fromNumber() gopurs_runtime.Value {
	once_Data_Int_fromNumber.Do(func() {
		cache_Data_Int_fromNumber = gopurs_runtime.Apply2(Get_Data_Int_fromNumberImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))})
	})
	return cache_Data_Int_fromNumber
}

var cache_Data_Int_unsafeClamp gopurs_runtime.Value
var once_Data_Int_unsafeClamp sync.Once
func Get_Data_Int_unsafeClamp() gopurs_runtime.Value {
	once_Data_Int_unsafeClamp.Do(func() {
		cache_Data_Int_unsafeClamp = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(x_0_box.FloatVal()))
})
	})
	return cache_Data_Int_unsafeClamp
}

var cache_Data_Int_round gopurs_runtime.Value
var once_Data_Int_round sync.Once
func Get_Data_Int_round() gopurs_runtime.Value {
	once_Data_Int_round.Do(func() {
		cache_Data_Int_round = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Int_round(x_0_box.FloatVal()))
})
	})
	return cache_Data_Int_round
}

var cache_Data_Int_trunc gopurs_runtime.Value
var once_Data_Int_trunc sync.Once
func Get_Data_Int_trunc() gopurs_runtime.Value {
	once_Data_Int_trunc.Do(func() {
		cache_Data_Int_trunc = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Int_trunc(x_0_box.FloatVal()))
})
	})
	return cache_Data_Int_trunc
}

var cache_Data_Int_floor gopurs_runtime.Value
var once_Data_Int_floor sync.Once
func Get_Data_Int_floor() gopurs_runtime.Value {
	once_Data_Int_floor.Do(func() {
		cache_Data_Int_floor = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Int_floor(x_0_box.FloatVal()))
})
	})
	return cache_Data_Int_floor
}

var cache_Data_Int_even gopurs_runtime.Value
var once_Data_Int_even sync.Once
func Get_Data_Int_even() gopurs_runtime.Value {
	once_Data_Int_even.Do(func() {
		cache_Data_Int_even = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Int_even(x_0_box.IntVal))
})
	})
	return cache_Data_Int_even
}

var cache_Data_Int_parity gopurs_runtime.Value
var once_Data_Int_parity sync.Once
func Get_Data_Int_parity() gopurs_runtime.Value {
	once_Data_Int_parity.Do(func() {
		cache_Data_Int_parity = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Int_parity(n_0_box.IntVal)), UnsafePtr: nil}
})
	})
	return cache_Data_Int_parity
}

var cache_Data_Int_eqParity gopurs_runtime.Value
var once_Data_Int_eqParity sync.Once
func Get_Data_Int_eqParity() gopurs_runtime.Value {
	once_Data_Int_eqParity.Do(func() {
		cache_Data_Int_eqParity = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 bool
{
var __t_tag_3 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_3) == 2591059121) {
var __t_tag_4 uint32 = uint32(y_1.IntVal)
__t5 = (uint32(__t_tag_4) == 2591059121)
goto end_branch_5
} else {

}
}
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 658452902) {

var __t_tag_1 uint32 = uint32(y_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 658452902)
}
__t5 = __t_and_2
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})})))}
	})
	return cache_Data_Int_eqParity
}

var cache_Data_Int_ordParity gopurs_runtime.Value
var once_Data_Int_ordParity sync.Once
func Get_Data_Int_ordParity() gopurs_runtime.Value {
	once_Data_Int_ordParity.Do(func() {
		cache_Data_Int_ordParity = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_3768443459_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[uint32]](Get_Data_Int_eqParity())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 uint32
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_0) == 2591059121) {
var __t2 uint32
{
var __t_tag_1 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_1) == 2591059121) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t7 = __t2
goto end_branch_7
} else {

}
}
{
var __t_tag_3 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_3) == 2591059121) {
__t7 = 380165415
goto end_branch_7
} else {

}
}
{
var __t_tag_4 uint32 = uint32(x_0.IntVal)
var __t_and_6 bool = false
if (uint32(__t_tag_4) == 658452902) {

var __t_tag_5 uint32 = uint32(y_1.IntVal)
__t_and_6 = (uint32(__t_tag_5) == 658452902)
}
if __t_and_6 {
__t7 = 902936544
goto end_branch_7
} else {

}
}
{
__t7 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Int_ordParity
}

var cache_Data_Int_semiringParity gopurs_runtime.Value
var once_Data_Int_semiringParity sync.Once
func Get_Data_Int_semiringParity() gopurs_runtime.Value {
	once_Data_Int_semiringParity.Do(func() {
		cache_Data_Int_semiringParity = gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_2722024963_2826095630((&Constructor_Data_Semiring_Semiring[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 uint32
{
var __t5 bool
{
var __t_tag_3 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_3) == 2591059121) {
var __t_tag_4 uint32 = uint32(y_1.IntVal)
__t5 = (uint32(__t_tag_4) == 2591059121)
goto end_branch_5
} else {

}
}
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 658452902) {

var __t_tag_1 uint32 = uint32(y_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 658452902)
}
__t5 = __t_and_2
}
end_branch_5:
if __t5 {
__t6 = 2591059121
goto end_branch_6
} else {

}
}
{
__t6 = 658452902
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t6), UnsafePtr: nil}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 uint32
{
var __t_tag_7 uint32 = uint32(v_0.IntVal)
var __t_and_9 bool = false
if (uint32(__t_tag_7) == 658452902) {

var __t_tag_8 uint32 = uint32(v1_1.IntVal)
__t_and_9 = (uint32(__t_tag_8) == 658452902)
}
if __t_and_9 {
__t10 = 658452902
goto end_branch_10
} else {

}
}
{
__t10 = 2591059121
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t10), UnsafePtr: nil}
}), 658452902, 2591059121})))}
	})
	return cache_Data_Int_semiringParity
}

var cache_Data_Int_ringParity gopurs_runtime.Value
var once_Data_Int_ringParity sync.Once
func Get_Data_Int_ringParity() gopurs_runtime.Value {
	once_Data_Int_ringParity.Do(func() {
		cache_Data_Int_ringParity = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_2370240579_3060961102((&Constructor_Data_Ring_Ring[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_2722024963_2826095630(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[uint32]](Get_Data_Int_semiringParity())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 uint32
{
var __t5 bool
{
var __t_tag_3 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_3) == 2591059121) {
var __t_tag_4 uint32 = uint32(y_1.IntVal)
__t5 = (uint32(__t_tag_4) == 2591059121)
goto end_branch_5
} else {

}
}
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 658452902) {

var __t_tag_1 uint32 = uint32(y_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 658452902)
}
__t5 = __t_and_2
}
end_branch_5:
if __t5 {
__t6 = 2591059121
goto end_branch_6
} else {

}
}
{
__t6 = 658452902
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t6), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Int_ringParity
}

var cache_Data_Int_divisionRingParity gopurs_runtime.Value
var once_Data_Int_divisionRingParity sync.Once
func Get_Data_Int_divisionRingParity() gopurs_runtime.Value {
	once_Data_Int_divisionRingParity.Do(func() {
		cache_Data_Int_divisionRingParity = gopurs_runtime.Value{Type: 9, IntVal: 2548491258, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_2384868323_2381053422((&Constructor_Data_DivisionRing_DivisionRing[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_2370240579_3060961102(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[uint32]](Get_Data_Int_ringParity())))}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)})))}
	})
	return cache_Data_Int_divisionRingParity
}

var cache_Data_Int_decimal gopurs_runtime.Value
var once_Data_Int_decimal sync.Once
func Get_Data_Int_decimal() gopurs_runtime.Value {
	once_Data_Int_decimal.Do(func() {
		cache_Data_Int_decimal = gopurs_runtime.Int(int64(10))
	})
	return cache_Data_Int_decimal
}

var cache_Data_Int_commutativeRingParity gopurs_runtime.Value
var once_Data_Int_commutativeRingParity sync.Once
func Get_Data_Int_commutativeRingParity() gopurs_runtime.Value {
	once_Data_Int_commutativeRingParity.Do(func() {
		cache_Data_Int_commutativeRingParity = gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_2537586851_1073849710((&Constructor_Data_CommutativeRing_CommutativeRing[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_2370240579_3060961102(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[uint32]](Get_Data_Int_ringParity())))}
})})))}
	})
	return cache_Data_Int_commutativeRingParity
}

var cache_Data_Int_euclideanRingParity gopurs_runtime.Value
var once_Data_Int_euclideanRingParity sync.Once
func Get_Data_Int_euclideanRingParity() gopurs_runtime.Value {
	once_Data_Int_euclideanRingParity.Do(func() {
		cache_Data_Int_euclideanRingParity = gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_3263107363_1774031598((&Constructor_Data_EuclideanRing_EuclideanRing[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_2537586851_1073849710(gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing[uint32]](Get_Data_Int_commutativeRingParity())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 int64
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 2591059121) {
__t2 = int64(0)
goto end_branch_2
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 658452902) {
__t2 = int64(1)
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Int(__t2)
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(x_0.IntVal)), UnsafePtr: nil}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Int_euclideanRingParity
}

var cache_Data_Int_ceil gopurs_runtime.Value
var once_Data_Int_ceil sync.Once
func Get_Data_Int_ceil() gopurs_runtime.Value {
	once_Data_Int_ceil.Do(func() {
		cache_Data_Int_ceil = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Int_ceil(x_0_box.FloatVal()))
})
	})
	return cache_Data_Int_ceil
}

var cache_Data_Int_boundedParity gopurs_runtime.Value
var once_Data_Int_boundedParity sync.Once
func Get_Data_Int_boundedParity() gopurs_runtime.Value {
	once_Data_Int_boundedParity.Do(func() {
		cache_Data_Int_boundedParity = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_832288803_2094947566((&Constructor_Data_Bounded_Bounded[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Int_ordParity())))}
}), 2591059121, 658452902})))}
	})
	return cache_Data_Int_boundedParity
}

var cache_Data_Int_binary gopurs_runtime.Value
var once_Data_Int_binary sync.Once
func Get_Data_Int_binary() gopurs_runtime.Value {
	once_Data_Int_binary.Do(func() {
		cache_Data_Int_binary = gopurs_runtime.Int(int64(2))
	})
	return cache_Data_Int_binary
}

var cache_Data_Int_base36 gopurs_runtime.Value
var once_Data_Int_base36 sync.Once
func Get_Data_Int_base36() gopurs_runtime.Value {
	once_Data_Int_base36.Do(func() {
		cache_Data_Int_base36 = gopurs_runtime.Int(int64(36))
	})
	return cache_Data_Int_base36
}

type Constructor_Data_Int_Even struct {
	Rc uint32
}


type Constructor_Data_Int_Odd struct {
	Rc uint32
}


func Call_Data_Int_Radix(x_0_loop int64) int64 {
var x_0 int64 = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Int_radix(n_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var n_0 int64 = n_0_loop
_ = n_0
var __t0 gopurs_runtime.Value
{
if ((n_0) >= (int64(2))) && ((n_0) <= (int64(36))) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(n_0)}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Int_1170268447_3094389156(Rebox_Data_Int_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Int_odd(x_0_loop int64) bool {
var x_0 int64 = x_0_loop
_ = x_0
return (gopurs_runtime.Bool(((x_0) & (int64(1))) == (int64(0))).IntVal) == (gopurs_runtime.Bool(false).IntVal)
}

func Call_Data_Int_unsafeClamp(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
var __t2 int64
{
if ((gopurs_runtime.Apply(Get_Data_Number_isFinite(), gopurs_runtime.Float(x_0)).IntVal) != (0)) != (true) {
__t2 = int64(0)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Apply2(Get_Data_Ord_greaterThanOrEq__1036186893(), gopurs_runtime.Float(x_0), gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal)).FloatVal())).IntVal) != (0) {
__t2 = Get_Data_Bounded_topInt().IntVal
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Apply2(Get_Data_Ord_lessThanOrEq__110343017(), gopurs_runtime.Float(x_0), gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal)).FloatVal())).IntVal) != (0) {
__t2 = Get_Data_Bounded_bottomInt().IntVal
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_1_0 := Rebox_Data_Int_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Int_fromNumber(), gopurs_runtime.Float(x_0))))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0 == nil) {
__t1 = gopurs_runtime.Int(int64(0))
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.Int((__local_var_1_0).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = __t1.IntVal
}
end_branch_2:
return __t2
}

func Call_Data_Int_round(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Number_round(), gopurs_runtime.Float(x_0)).FloatVal()).FloatVal())).IntVal
}

func Call_Data_Int_trunc(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Number_trunc(), gopurs_runtime.Float(x_0)).FloatVal()).FloatVal())).IntVal
}

func Call_Data_Int_floor(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(x_0)).FloatVal()).FloatVal())).IntVal
}

func Call_Data_Int_even(x_0_loop int64) bool {
var x_0 int64 = x_0_loop
_ = x_0
return ((x_0) & (int64(1))) == (int64(0))
}

func Call_Data_Int_parity(n_0_loop int64) uint32 {
var n_0 int64 = n_0_loop
_ = n_0
var __t0 uint32
{
if ((n_0) & (int64(1))) == (int64(0)) {
__t0 = 2591059121
goto end_branch_0
} else {

}
}
{
__t0 = 658452902
}
end_branch_0:
return __t0
}

func Call_Data_Int_ceil(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Number_ceil(), gopurs_runtime.Float(x_0)).FloatVal()).FloatVal())).IntVal
}

func Rebox_Data_Int_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Int_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Int_2370240579_3060961102(in *Constructor_Data_Ring_Ring[uint32]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Ring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Int_2384868323_2381053422(in *Constructor_Data_DivisionRing_DivisionRing[uint32]) *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Int_2537586851_1073849710(in *Constructor_Data_CommutativeRing_CommutativeRing[uint32]) *Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Int_2722024963_2826095630(in *Constructor_Data_Semiring_Semiring[uint32]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
		out.V3 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V3), UnsafePtr: nil}
	return out
}

func Rebox_Data_Int_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Int_3263107363_1774031598(in *Constructor_Data_EuclideanRing_EuclideanRing[uint32]) *Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Int_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Int_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Int_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	return out
}

func Get_Data_Int_fromNumberImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Int_FromNumberImpl
}

func Get_Data_Int_fromStringAsImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Int_FromStringAsImpl
}

func Get_Data_Int_pow() gopurs_runtime.Value {
	return _Gopurs_Data_Int_Pow
}

func Get_Data_Int_quot() gopurs_runtime.Value {
	return _Gopurs_Data_Int_Quot
}

func Get_Data_Int_rem() gopurs_runtime.Value {
	return _Gopurs_Data_Int_Rem
}

func Get_Data_Int_toNumber() gopurs_runtime.Value {
	return _Gopurs_Data_Int_ToNumber
}

func Get_Data_Int_toStringAs() gopurs_runtime.Value {
	return _Gopurs_Data_Int_ToStringAs
}
