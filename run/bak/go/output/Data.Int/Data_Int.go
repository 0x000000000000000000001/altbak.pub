package Data_Int

import (
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Number "gopurs/output/Data.Number"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Even gopurs_runtime.Value
var once_Even sync.Once
func Get_Even() gopurs_runtime.Value {
	once_Even.Do(func() {
		cache_Even = gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}
	})
	return cache_Even
}

var cache_Odd gopurs_runtime.Value
var once_Odd sync.Once
func Get_Odd() gopurs_runtime.Value {
	once_Odd.Do(func() {
		cache_Odd = gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}
	})
	return cache_Odd
}

var cache_showParity gopurs_runtime.Value
var once_showParity sync.Once
func Get_showParity() gopurs_runtime.Value {
	once_showParity.Do(func() {
		cache_showParity = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(v_0.IntVal) == 2591059121) {
__t0 = gopurs_runtime.Str("Even")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 658452902) {
__t0 = gopurs_runtime.Str("Odd")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
	})
	return cache_showParity
}

var cache_radix gopurs_runtime.Value
var once_radix sync.Once
func Get_radix() gopurs_runtime.Value {
	once_radix.Do(func() {
		cache_radix = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_radix(n_0_box.IntVal))}
})
	})
	return cache_radix
}

var cache_odd gopurs_runtime.Value
var once_odd sync.Once
func Get_odd() gopurs_runtime.Value {
	once_odd.Do(func() {
		cache_odd = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_odd(x_0_box.IntVal))
})
	})
	return cache_odd
}

var cache_octal gopurs_runtime.Value
var once_octal sync.Once
func Get_octal() gopurs_runtime.Value {
	once_octal.Do(func() {
		cache_octal = gopurs_runtime.Int(8)
	})
	return cache_octal
}

var cache_hexadecimal gopurs_runtime.Value
var once_hexadecimal sync.Once
func Get_hexadecimal() gopurs_runtime.Value {
	once_hexadecimal.Do(func() {
		cache_hexadecimal = gopurs_runtime.Int(16)
	})
	return cache_hexadecimal
}

var cache_fromStringAs gopurs_runtime.Value
var once_fromStringAs sync.Once
func Get_fromStringAs() gopurs_runtime.Value {
	once_fromStringAs.Do(func() {
		cache_fromStringAs = gopurs_runtime.Apply2(Get_fromStringAsImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
	})
	return cache_fromStringAs
}

var cache_fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		cache_fromString = gopurs_runtime.Apply(Get_fromStringAs(), gopurs_runtime.Int(10))
	})
	return cache_fromString
}

var cache_fromNumber gopurs_runtime.Value
var once_fromNumber sync.Once
func Get_fromNumber() gopurs_runtime.Value {
	once_fromNumber.Do(func() {
		cache_fromNumber = gopurs_runtime.Apply2(Get_fromNumberImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
	})
	return cache_fromNumber
}

var cache_unsafeClamp gopurs_runtime.Value
var once_unsafeClamp sync.Once
func Get_unsafeClamp() gopurs_runtime.Value {
	once_unsafeClamp.Do(func() {
		cache_unsafeClamp = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_unsafeClamp(x_0_box.FloatVal()))
})
	})
	return cache_unsafeClamp
}

var cache_round gopurs_runtime.Value
var once_round sync.Once
func Get_round() gopurs_runtime.Value {
	once_round.Do(func() {
		cache_round = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_round(x_0_box.FloatVal()))
})
	})
	return cache_round
}

var cache_trunc gopurs_runtime.Value
var once_trunc sync.Once
func Get_trunc() gopurs_runtime.Value {
	once_trunc.Do(func() {
		cache_trunc = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_trunc(x_0_box.FloatVal()))
})
	})
	return cache_trunc
}

var cache_floor gopurs_runtime.Value
var once_floor sync.Once
func Get_floor() gopurs_runtime.Value {
	once_floor.Do(func() {
		cache_floor = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_floor(x_0_box.FloatVal()))
})
	})
	return cache_floor
}

var cache_even gopurs_runtime.Value
var once_even sync.Once
func Get_even() gopurs_runtime.Value {
	once_even.Do(func() {
		cache_even = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_even(x_0_box.IntVal))
})
	})
	return cache_even
}

var cache_parity gopurs_runtime.Value
var once_parity sync.Once
func Get_parity() gopurs_runtime.Value {
	once_parity.Do(func() {
		cache_parity = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_parity(n_0_box.IntVal)), UnsafePtr: nil}
})
	})
	return cache_parity
}

var cache_eqParity gopurs_runtime.Value
var once_eqParity sync.Once
func Get_eqParity() gopurs_runtime.Value {
	once_eqParity.Do(func() {
		cache_eqParity = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(x_0.IntVal) == 2591059121) {
var __t1 gopurs_runtime.Value
{
if (uint32(y_1.IntVal) == 2591059121) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if ((uint32(x_0.IntVal) == 658452902)) && ((uint32(y_1.IntVal) == 658452902)) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
})
}))
	})
	return cache_eqParity
}

var cache_ordParity gopurs_runtime.Value
var once_ordParity sync.Once
func Get_ordParity() gopurs_runtime.Value {
	once_ordParity.Do(func() {
		cache_ordParity = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqParity()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(x_0.IntVal) == 2591059121) {
var __t1 gopurs_runtime.Value
{
if (uint32(y_1.IntVal) == 2591059121) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (uint32(y_1.IntVal) == 2591059121) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if ((uint32(x_0.IntVal) == 658452902)) && ((uint32(y_1.IntVal) == 658452902)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t0.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordParity
}

var cache_semiringParity gopurs_runtime.Value
var once_semiringParity sync.Once
func Get_semiringParity() gopurs_runtime.Value {
	once_semiringParity.Do(func() {
		cache_semiringParity = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqParity(), "eq"), x_0, y_1).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t0.IntVal)), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((uint32(v_0.IntVal) == 658452902)) && ((uint32(v1_1.IntVal) == 658452902)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t1.IntVal)), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}.IntVal)), UnsafePtr: nil})
	})
	return cache_semiringParity
}

var cache_ringParity gopurs_runtime.Value
var once_ringParity sync.Once
func Get_ringParity() gopurs_runtime.Value {
	once_ringParity.Do(func() {
		cache_ringParity = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semiringParity()
}), gopurs_runtime.RecordGet(Get_semiringParity(), "add"))
	})
	return cache_ringParity
}

var cache_divisionRingParity gopurs_runtime.Value
var once_divisionRingParity sync.Once
func Get_divisionRingParity() gopurs_runtime.Value {
	once_divisionRingParity.Do(func() {
		cache_divisionRingParity = gopurs_runtime.RecordDict2("Ring0", "recip", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ringParity()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_divisionRingParity
}

var cache_decimal gopurs_runtime.Value
var once_decimal sync.Once
func Get_decimal() gopurs_runtime.Value {
	once_decimal.Do(func() {
		cache_decimal = gopurs_runtime.Int(10)
	})
	return cache_decimal
}

var cache_commutativeRingParity gopurs_runtime.Value
var once_commutativeRingParity sync.Once
func Get_commutativeRingParity() gopurs_runtime.Value {
	once_commutativeRingParity.Do(func() {
		cache_commutativeRingParity = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ringParity()
}))
	})
	return cache_commutativeRingParity
}

var cache_euclideanRingParity gopurs_runtime.Value
var once_euclideanRingParity sync.Once
func Get_euclideanRingParity() gopurs_runtime.Value {
	once_euclideanRingParity.Do(func() {
		cache_euclideanRingParity = gopurs_runtime.RecordDict4("CommutativeRing0", "degree", "div", "mod", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_commutativeRingParity()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(v_0.IntVal) == 2591059121) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 658452902) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0.IntVal)
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(x_0.IntVal)), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_euclideanRingParity
}

var cache_ceil gopurs_runtime.Value
var once_ceil sync.Once
func Get_ceil() gopurs_runtime.Value {
	once_ceil.Do(func() {
		cache_ceil = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_ceil(x_0_box.FloatVal()))
})
	})
	return cache_ceil
}

var cache_boundedParity gopurs_runtime.Value
var once_boundedParity sync.Once
func Get_boundedParity() gopurs_runtime.Value {
	once_boundedParity.Do(func() {
		cache_boundedParity = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordParity()
}), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}.IntVal)), UnsafePtr: nil})
	})
	return cache_boundedParity
}

var cache_binary gopurs_runtime.Value
var once_binary sync.Once
func Get_binary() gopurs_runtime.Value {
	once_binary.Do(func() {
		cache_binary = gopurs_runtime.Int(2)
	})
	return cache_binary
}

var cache_base36 gopurs_runtime.Value
var once_base36 sync.Once
func Get_base36() gopurs_runtime.Value {
	once_base36.Do(func() {
		cache_base36 = gopurs_runtime.Int(36)
	})
	return cache_base36
}

type Constructor_Even struct {
	Rc uint32
}


type Constructor_Odd struct {
	Rc uint32
}


func Call_radix(n_0_loop int64) *pkg_Data_Maybe.Constructor_Just[int64] {
var n_0 int64 = n_0_loop
_ = n_0
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (n_0) < (2) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (n_0) > (36) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(true)
}
end_branch_2:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t1, __t2).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(n_0)})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0)
}

func Call_odd(x_0_loop int64) bool {
var x_0 int64 = x_0_loop
_ = x_0
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool(((x_0) & (1)) == (0)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_unsafeClamp(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(pkg_Data_Number.Get_isFinite(), gopurs_runtime.Float(x_0))).IntVal) != (0) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
if (x_0) < (gopurs_runtime.Apply(Get_toNumber(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "top")).FloatVal()) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(true)
}
end_branch_3:
if (__t3.IntVal) != (0) {
__t2 = gopurs_runtime.Int(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "top").IntVal)
goto end_branch_2
} else {

}
}
{
var __t4 gopurs_runtime.Value
{
if (x_0) > (gopurs_runtime.Apply(Get_toNumber(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "bottom")).FloatVal()) {
__t4 = gopurs_runtime.Bool(false)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(true)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t2 = gopurs_runtime.Int(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "bottom").IntVal)
goto end_branch_2
} else {

}
}
{
__local_var_1_0 := gopurs_runtime.Apply(Get_fromNumber(), gopurs_runtime.Float(x_0))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Int(__t1.IntVal)
}
end_branch_2:
return __t2.IntVal
}

func Call_round(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_unsafeClamp(gopurs_runtime.Apply(pkg_Data_Number.Get_round(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
}

func Call_trunc(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_unsafeClamp(gopurs_runtime.Apply(pkg_Data_Number.Get_trunc(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
}

func Call_floor(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_unsafeClamp(gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
}

func Call_even(x_0_loop int64) bool {
var x_0 int64 = x_0_loop
_ = x_0
return ((x_0) & (1)) == (0)
}

func Call_parity(n_0_loop int64) uint32 {
var n_0 int64 = n_0_loop
_ = n_0
var __t0 gopurs_runtime.Value
{
if ((n_0) & (1)) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_0:
return uint32(__t0.IntVal)
}

func Call_ceil(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_unsafeClamp(gopurs_runtime.Apply(pkg_Data_Number.Get_ceil(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
}

func Get_fromNumberImpl() gopurs_runtime.Value {
	return _Gopurs_FromNumberImpl
}

func Get_fromStringAsImpl() gopurs_runtime.Value {
	return _Gopurs_FromStringAsImpl
}

func Get_pow() gopurs_runtime.Value {
	return _Gopurs_Pow
}

func Get_quot() gopurs_runtime.Value {
	return _Gopurs_Quot
}

func Get_rem() gopurs_runtime.Value {
	return _Gopurs_Rem
}

func Get_toNumber() gopurs_runtime.Value {
	return _Gopurs_ToNumber
}

func Get_toStringAs() gopurs_runtime.Value {
	return _Gopurs_ToStringAs
}
