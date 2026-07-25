package Data_Int

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Control_Category "gopurs/output/Control.Category"
	unsafe "unsafe"
)

var cache_Even gopurs_runtime.Value
var once_Even sync.Once
func Get_Even() gopurs_runtime.Value {
	once_Even.Do(func() {
		cache_Even = gopurs_runtime.Value{Type: 9, IntVal: 2591059121, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Even{})}
	})
	return cache_Even
}

var cache_Odd gopurs_runtime.Value
var once_Odd sync.Once
func Get_Odd() gopurs_runtime.Value {
	once_Odd.Do(func() {
		cache_Odd = gopurs_runtime.Value{Type: 9, IntVal: 658452902, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Odd{})}
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
if (v_0.Type == 9 && v_0.IntVal == 2591059121) {
__t0 = gopurs_runtime.Str("Even")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 658452902) {
__t0 = gopurs_runtime.Str("Odd")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return cache_showParity
}

var cache_radix gopurs_runtime.Value
var once_radix sync.Once
func Get_radix() gopurs_runtime.Value {
	once_radix.Do(func() {
		cache_radix = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (2)) && ((n_0.IntVal) <= (36)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{n_0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}()
})
	})
	return cache_radix
}

var cache_odd gopurs_runtime.Value
var once_odd sync.Once
func Get_odd() gopurs_runtime.Value {
	once_odd.Do(func() {
		cache_odd = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Bool(((x_0.IntVal) & (1)) != (0))
}()
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
		cache_fromStringAs = gopurs_runtime.Apply2(Get_fromStringAsImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})})
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
		cache_fromNumber = gopurs_runtime.Apply2(Get_fromNumberImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})})
	})
	return cache_fromNumber
}

var cache_unsafeClamp gopurs_runtime.Value
var once_unsafeClamp sync.Once
func Get_unsafeClamp() gopurs_runtime.Value {
	once_unsafeClamp.Do(func() {
		cache_unsafeClamp = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var __t2 gopurs_runtime.Value
{
if ((gopurs_runtime.Apply(pkg_Data_Number.Get_isFinite(), x_0).IntVal) != (0)) != (true) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (x_0.FloatVal()) >= (gopurs_runtime.Apply(Get_toNumber(), pkg_Data_Bounded.Get_topInt()).FloatVal()) {
__t2 = pkg_Data_Bounded.Get_topInt()
goto end_branch_2
} else {

}
}
{
if (x_0.FloatVal()) <= (gopurs_runtime.Apply(Get_toNumber(), pkg_Data_Bounded.Get_bottomInt()).FloatVal()) {
__t2 = pkg_Data_Bounded.Get_bottomInt()
goto end_branch_2
} else {

}
}
{
__local_var_1_0 := gopurs_runtime.Apply(Get_fromNumber(), x_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 3589588149) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136) {
__t1 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_1_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return __t2
}()
})
	})
	return cache_unsafeClamp
}

var cache_round gopurs_runtime.Value
var once_round sync.Once
func Get_round() gopurs_runtime.Value {
	once_round.Do(func() {
		cache_round = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_round(), x_0))
}()
})
	})
	return cache_round
}

var cache_trunc gopurs_runtime.Value
var once_trunc sync.Once
func Get_trunc() gopurs_runtime.Value {
	once_trunc.Do(func() {
		cache_trunc = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_trunc(), x_0))
}()
})
	})
	return cache_trunc
}

var cache_floor gopurs_runtime.Value
var once_floor sync.Once
func Get_floor() gopurs_runtime.Value {
	once_floor.Do(func() {
		cache_floor = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), x_0))
}()
})
	})
	return cache_floor
}

var cache_even gopurs_runtime.Value
var once_even sync.Once
func Get_even() gopurs_runtime.Value {
	once_even.Do(func() {
		cache_even = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Bool(((x_0.IntVal) & (1)) == (0))
}()
})
	})
	return cache_even
}

var cache_parity gopurs_runtime.Value
var once_parity sync.Once
func Get_parity() gopurs_runtime.Value {
	once_parity.Do(func() {
		cache_parity = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) & (1)) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2591059121, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Even{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 658452902, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Odd{})}
}
end_branch_0:
return __t0
}()
})
	})
	return cache_parity
}

var cache_eqParity gopurs_runtime.Value
var once_eqParity sync.Once
func Get_eqParity() gopurs_runtime.Value {
	once_eqParity.Do(func() {
		cache_eqParity = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 2591059121) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 2591059121))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((x_0.Type == 9 && x_0.IntVal == 658452902)) && ((y_1.Type == 9 && y_1.IntVal == 658452902)))
}
end_branch_0:
return __t0
}))
	})
	return cache_eqParity
}

var cache_ordParity gopurs_runtime.Value
var once_ordParity sync.Once
func Get_ordParity() gopurs_runtime.Value {
	once_ordParity.Do(func() {
		cache_ordParity = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 2591059121) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 2591059121) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 2591059121) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 658452902)) && ((y_1.Type == 9 && y_1.IntVal == 658452902)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqParity()
}))
	})
	return cache_ordParity
}

var cache_semiringParity gopurs_runtime.Value
var once_semiringParity sync.Once
func Get_semiringParity() gopurs_runtime.Value {
	once_semiringParity.Do(func() {
		cache_semiringParity = gopurs_runtime.RecordDict4("zero", "add", "one", "mul", gopurs_runtime.Value{Type: 9, IntVal: 2591059121, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Even{})}, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 2591059121) {
__t1 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 2591059121))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(((x_0.Type == 9 && x_0.IntVal == 658452902)) && ((y_1.Type == 9 && y_1.IntVal == 658452902)))
}
end_branch_1:
if (__t1.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2591059121, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Even{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 658452902, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Odd{})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 658452902, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Odd{})}, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if ((v_0.Type == 9 && v_0.IntVal == 658452902)) && ((v1_1.Type == 9 && v1_1.IntVal == 658452902)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 658452902, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Odd{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2591059121, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Even{})}
}
end_branch_2:
return __t2
}))
	})
	return cache_semiringParity
}

var cache_ringParity gopurs_runtime.Value
var once_ringParity sync.Once
func Get_ringParity() gopurs_runtime.Value {
	once_ringParity.Do(func() {
		cache_ringParity = gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.RecordGet(Get_semiringParity(), "add"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semiringParity()
}))
	})
	return cache_ringParity
}

var cache_divisionRingParity gopurs_runtime.Value
var once_divisionRingParity sync.Once
func Get_divisionRingParity() gopurs_runtime.Value {
	once_divisionRingParity.Do(func() {
		cache_divisionRingParity = gopurs_runtime.RecordDict2("recip", "Ring0", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ringParity()
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
		cache_euclideanRingParity = gopurs_runtime.RecordDict4("degree", "div", "mod", "CommutativeRing0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 2591059121) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 658452902) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2591059121, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Even{})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_commutativeRingParity()
}))
	})
	return cache_euclideanRingParity
}

var cache_ceil gopurs_runtime.Value
var once_ceil sync.Once
func Get_ceil() gopurs_runtime.Value {
	once_ceil.Do(func() {
		cache_ceil = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_ceil(), x_0))
}()
})
	})
	return cache_ceil
}

var cache_boundedParity gopurs_runtime.Value
var once_boundedParity sync.Once
func Get_boundedParity() gopurs_runtime.Value {
	once_boundedParity.Do(func() {
		cache_boundedParity = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Value{Type: 9, IntVal: 2591059121, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Even{})}, gopurs_runtime.Value{Type: 9, IntVal: 658452902, UnsafePtr: unsafe.Pointer(&Data_Data_Int_Odd{})}, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordParity()
}))
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

type Data_Data_Int_Even struct {
	
}
func Is_Data_Data_Int_Even(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2591059121
}

type Data_Data_Int_Odd struct {
	
}
func Is_Data_Data_Int_Odd(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 658452902
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
