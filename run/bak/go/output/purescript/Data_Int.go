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
		cache_Data_Int_top = gopurs_runtime.Int(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "top").IntVal)
	})
	return cache_Data_Int_top
}

var cache_Data_Int_bottom gopurs_runtime.Value
var once_Data_Int_bottom sync.Once
func Get_Data_Int_bottom() gopurs_runtime.Value {
	once_Data_Int_bottom.Do(func() {
		cache_Data_Int_bottom = gopurs_runtime.Int(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "bottom").IntVal)
	})
	return cache_Data_Int_bottom
}

var cache_Data_Int_Radix gopurs_runtime.Value
var once_Data_Int_Radix sync.Once
func Get_Data_Int_Radix() gopurs_runtime.Value {
	once_Data_Int_Radix.Do(func() {
		cache_Data_Int_Radix = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Int_Radix(x_0_box)
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
		cache_Data_Int_showParity = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (uint32(v_0.IntVal) == 2591059121) {
__t0 = "Even"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 658452902) {
__t0 = "Odd"
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
}))
	})
	return cache_Data_Int_showParity
}

var cache_Data_Int_radix gopurs_runtime.Value
var once_Data_Int_radix sync.Once
func Get_Data_Int_radix() gopurs_runtime.Value {
	once_Data_Int_radix.Do(func() {
		cache_Data_Int_radix = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Int_radix(n_0_box.IntVal))}
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
		cache_Data_Int_octal = gopurs_runtime.Int(8)
	})
	return cache_Data_Int_octal
}

var cache_Data_Int_hexadecimal gopurs_runtime.Value
var once_Data_Int_hexadecimal sync.Once
func Get_Data_Int_hexadecimal() gopurs_runtime.Value {
	once_Data_Int_hexadecimal.Do(func() {
		cache_Data_Int_hexadecimal = gopurs_runtime.Int(16)
	})
	return cache_Data_Int_hexadecimal
}

var cache_Data_Int_fromStringAs gopurs_runtime.Value
var once_Data_Int_fromStringAs sync.Once
func Get_Data_Int_fromStringAs() gopurs_runtime.Value {
	once_Data_Int_fromStringAs.Do(func() {
		cache_Data_Int_fromStringAs = gopurs_runtime.Apply2(Get_Data_Int_fromStringAsImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
	})
	return cache_Data_Int_fromStringAs
}

var cache_Data_Int_fromString gopurs_runtime.Value
var once_Data_Int_fromString sync.Once
func Get_Data_Int_fromString() gopurs_runtime.Value {
	once_Data_Int_fromString.Do(func() {
		cache_Data_Int_fromString = gopurs_runtime.Apply(Get_Data_Int_fromStringAs(), gopurs_runtime.Int(10))
	})
	return cache_Data_Int_fromString
}

var cache_Data_Int_fromNumber gopurs_runtime.Value
var once_Data_Int_fromNumber sync.Once
func Get_Data_Int_fromNumber() gopurs_runtime.Value {
	once_Data_Int_fromNumber.Do(func() {
		cache_Data_Int_fromNumber = gopurs_runtime.Apply2(Get_Data_Int_fromNumberImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
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
		cache_Data_Int_eqParity = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (uint32(x_0.IntVal) == 2591059121) {
var __t0 bool
{
if (uint32(y_1.IntVal) == 2591059121) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((uint32(x_0.IntVal) == 658452902)) && ((uint32(y_1.IntVal) == 658452902)) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
}))
	})
	return cache_Data_Int_eqParity
}

var cache_Data_Int_ordParity gopurs_runtime.Value
var once_Data_Int_ordParity sync.Once
func Get_Data_Int_ordParity() gopurs_runtime.Value {
	once_Data_Int_ordParity.Do(func() {
		cache_Data_Int_ordParity = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Int_eqParity()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 uint32
{
if (uint32(x_0.IntVal) == 2591059121) {
var __t0 uint32
{
if (uint32(y_1.IntVal) == 2591059121) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if (uint32(y_1.IntVal) == 2591059121) {
__t1 = 380165415
goto end_branch_1
} else {

}
}
{
if ((uint32(x_0.IntVal) == 658452902)) && ((uint32(y_1.IntVal) == 658452902)) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Int_ordParity
}

var cache_Data_Int_semiringParity gopurs_runtime.Value
var once_Data_Int_semiringParity sync.Once
func Get_Data_Int_semiringParity() gopurs_runtime.Value {
	once_Data_Int_semiringParity.Do(func() {
		cache_Data_Int_semiringParity = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 uint32
{
var __t6 bool
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_0) == 2591059121) {
var __t2 bool
{
var __t_tag_1 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_1) == 2591059121) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t6 = __t2
goto end_branch_6
} else {

}
}
{
var __t_tag_3 uint32 = uint32(x_0.IntVal)
var __t_and_5 bool = false
if (uint32(__t_tag_3) == 658452902) {

var __t_tag_4 uint32 = uint32(y_1.IntVal)
__t_and_5 = (uint32(__t_tag_4) == 658452902)
}
if __t_and_5 {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
if __t6 {
__t7 = 2591059121
goto end_branch_7
} else {

}
}
{
__t7 = 658452902
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 uint32
{
if ((uint32(v_0.IntVal) == 658452902)) && ((uint32(v1_1.IntVal) == 658452902)) {
__t8 = 658452902
goto end_branch_8
} else {

}
}
{
__t8 = 2591059121
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t8), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil})
	})
	return cache_Data_Int_semiringParity
}

var cache_Data_Int_ringParity gopurs_runtime.Value
var once_Data_Int_ringParity sync.Once
func Get_Data_Int_ringParity() gopurs_runtime.Value {
	once_Data_Int_ringParity.Do(func() {
		cache_Data_Int_ringParity = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Int_semiringParity()
}), gopurs_runtime.RecordGet(Get_Data_Int_semiringParity(), "add"))
	})
	return cache_Data_Int_ringParity
}

var cache_Data_Int_divisionRingParity gopurs_runtime.Value
var once_Data_Int_divisionRingParity sync.Once
func Get_Data_Int_divisionRingParity() gopurs_runtime.Value {
	once_Data_Int_divisionRingParity.Do(func() {
		cache_Data_Int_divisionRingParity = gopurs_runtime.RecordDict2("Ring0", "recip", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Int_ringParity()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Int_divisionRingParity
}

var cache_Data_Int_decimal gopurs_runtime.Value
var once_Data_Int_decimal sync.Once
func Get_Data_Int_decimal() gopurs_runtime.Value {
	once_Data_Int_decimal.Do(func() {
		cache_Data_Int_decimal = gopurs_runtime.Int(10)
	})
	return cache_Data_Int_decimal
}

var cache_Data_Int_commutativeRingParity gopurs_runtime.Value
var once_Data_Int_commutativeRingParity sync.Once
func Get_Data_Int_commutativeRingParity() gopurs_runtime.Value {
	once_Data_Int_commutativeRingParity.Do(func() {
		cache_Data_Int_commutativeRingParity = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Int_ringParity()
}))
	})
	return cache_Data_Int_commutativeRingParity
}

var cache_Data_Int_euclideanRingParity gopurs_runtime.Value
var once_Data_Int_euclideanRingParity sync.Once
func Get_Data_Int_euclideanRingParity() gopurs_runtime.Value {
	once_Data_Int_euclideanRingParity.Do(func() {
		cache_Data_Int_euclideanRingParity = gopurs_runtime.RecordDict4("CommutativeRing0", "degree", "div", "mod", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Int_commutativeRingParity()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if (uint32(v_0.IntVal) == 2591059121) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 658452902) {
__t0 = 1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(x_0.IntVal)), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}
})
}))
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
		cache_Data_Int_boundedParity = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Int_ordParity()
}), gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil})
	})
	return cache_Data_Int_boundedParity
}

var cache_Data_Int_binary gopurs_runtime.Value
var once_Data_Int_binary sync.Once
func Get_Data_Int_binary() gopurs_runtime.Value {
	once_Data_Int_binary.Do(func() {
		cache_Data_Int_binary = gopurs_runtime.Int(2)
	})
	return cache_Data_Int_binary
}

var cache_Data_Int_base36 gopurs_runtime.Value
var once_Data_Int_base36 sync.Once
func Get_Data_Int_base36() gopurs_runtime.Value {
	once_Data_Int_base36.Do(func() {
		cache_Data_Int_base36 = gopurs_runtime.Int(36)
	})
	return cache_Data_Int_base36
}

type Constructor_Data_Int_Even struct {
	Rc uint32
}


type Constructor_Data_Int_Odd struct {
	Rc uint32
}


func Call_Data_Int_Radix(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Int_radix(n_0_loop int64) *Constructor_Data_Maybe_Just {
var n_0 int64 = n_0_loop
_ = n_0
var __t3 gopurs_runtime.Value
{
var __t0 bool
{
if (n_0) < (2) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
var __t_and_2 bool = false
if __t0 {

var __t1 bool
{
if (n_0) > (36) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
__t_and_2 = __t1
}
if __t_and_2 {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_0)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t3)
}

func Call_Data_Int_odd(x_0_loop int64) bool {
var x_0 int64 = x_0_loop
_ = x_0
return (((x_0) & (1)) == (0)) != (true)
}

func Call_Data_Int_unsafeClamp(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
var __t4 int64
{
if ((gopurs_runtime.Apply(Get_Data_Number_isFinite(), gopurs_runtime.Float(x_0)).IntVal) != (0)) != (true) {
__t4 = 0
goto end_branch_4
} else {

}
}
{
var __t2 bool
{
if (x_0) < (gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "top").IntVal)).FloatVal()) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
if __t2 {
__t4 = gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "top").IntVal
goto end_branch_4
} else {

}
}
{
var __t3 bool
{
if (x_0) > (gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "bottom").IntVal)).FloatVal()) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "bottom").IntVal
goto end_branch_4
} else {

}
}
{
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Int_fromNumber(), gopurs_runtime.Float(x_0)))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0 == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0 != nil) {
__t1 = (__local_var_1_0).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t4 = __t1.IntVal
}
end_branch_4:
return __t4
}

func Call_Data_Int_round(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(gopurs_runtime.Apply(Get_Data_Number_round(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
}

func Call_Data_Int_trunc(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(gopurs_runtime.Apply(Get_Data_Number_trunc(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
}

func Call_Data_Int_floor(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
}

func Call_Data_Int_even(x_0_loop int64) bool {
var x_0 int64 = x_0_loop
_ = x_0
return ((x_0) & (1)) == (0)
}

func Call_Data_Int_parity(n_0_loop int64) uint32 {
var n_0 int64 = n_0_loop
_ = n_0
var __t0 uint32
{
if ((n_0) & (1)) == (0) {
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
return gopurs_runtime.Int(Call_Data_Int_unsafeClamp(gopurs_runtime.Apply(Get_Data_Number_ceil(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
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
