package Data_Int

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var Even gopurs_runtime.Value
var once_Even sync.Once
func Get_Even() gopurs_runtime.Value {
	once_Even.Do(func() {
		Even = gopurs_runtime.Constructor0("Even")
	})
	return Even
}

var Odd gopurs_runtime.Value
var once_Odd sync.Once
func Get_Odd() gopurs_runtime.Value {
	once_Odd.Do(func() {
		Odd = gopurs_runtime.Constructor0("Odd")
	})
	return Odd
}

var showParity gopurs_runtime.Value
var once_showParity sync.Once
func Get_showParity() gopurs_runtime.Value {
	once_showParity.Do(func() {
		showParity = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Even").IntVal != 0 {
__t0 = gopurs_runtime.Str("Even")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Odd").IntVal != 0 {
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
	return showParity
}

var radix gopurs_runtime.Value
var once_radix sync.Once
func Get_radix() gopurs_runtime.Value {
	once_radix.Do(func() {
		radix = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if n_0.IntVal >= 2 && n_0.IntVal <= 36 {
__t0 = gopurs_runtime.Constructor1("Just", n_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
})
	})
	return radix
}

var odd gopurs_runtime.Value
var once_odd sync.Once
func Get_odd() gopurs_runtime.Value {
	once_odd.Do(func() {
		odd = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(x_0.IntVal & 1 != 0)
})
	})
	return odd
}

var octal gopurs_runtime.Value
var once_octal sync.Once
func Get_octal() gopurs_runtime.Value {
	once_octal.Do(func() {
		octal = gopurs_runtime.Int(8)
	})
	return octal
}

var hexadecimal gopurs_runtime.Value
var once_hexadecimal sync.Once
func Get_hexadecimal() gopurs_runtime.Value {
	once_hexadecimal.Do(func() {
		hexadecimal = gopurs_runtime.Int(16)
	})
	return hexadecimal
}

var fromStringAs gopurs_runtime.Value
var once_fromStringAs sync.Once
func Get_fromStringAs() gopurs_runtime.Value {
	once_fromStringAs.Do(func() {
		fromStringAs = gopurs_runtime.Apply2(Get_fromStringAsImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"))
	})
	return fromStringAs
}

var fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		fromString = gopurs_runtime.Apply(Get_fromStringAs(), gopurs_runtime.Int(10))
	})
	return fromString
}

var fromNumber gopurs_runtime.Value
var once_fromNumber sync.Once
func Get_fromNumber() gopurs_runtime.Value {
	once_fromNumber.Do(func() {
		fromNumber = gopurs_runtime.Apply2(Get_fromNumberImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"))
	})
	return fromNumber
}

var unsafeClamp gopurs_runtime.Value
var once_unsafeClamp sync.Once
func Get_unsafeClamp() gopurs_runtime.Value {
	once_unsafeClamp.Do(func() {
		unsafeClamp = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply(pkg_Data_Number.Get_isFinite(), x_0).IntVal != 0 != true {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if x_0.FloatVal() >= gopurs_runtime.Apply(Get_toNumber(), pkg_Data_Bounded.Get_topInt()).FloatVal() {
__t2 = pkg_Data_Bounded.Get_topInt()
goto end_branch_2
} else {

}
}
{
if x_0.FloatVal() <= gopurs_runtime.Apply(Get_toNumber(), pkg_Data_Bounded.Get_bottomInt()).FloatVal() {
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
if gopurs_runtime.Bool(__local_var_1_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_1_0.StrVal == "Just").IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)[0]
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
})
	})
	return unsafeClamp
}

var round gopurs_runtime.Value
var once_round sync.Once
func Get_round() gopurs_runtime.Value {
	once_round.Do(func() {
		round = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_round(), x_0))
})
	})
	return round
}

var trunc gopurs_runtime.Value
var once_trunc sync.Once
func Get_trunc() gopurs_runtime.Value {
	once_trunc.Do(func() {
		trunc = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_trunc(), x_0))
})
	})
	return trunc
}

var floor gopurs_runtime.Value
var once_floor sync.Once
func Get_floor() gopurs_runtime.Value {
	once_floor.Do(func() {
		floor = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), x_0))
})
	})
	return floor
}

var even gopurs_runtime.Value
var once_even sync.Once
func Get_even() gopurs_runtime.Value {
	once_even.Do(func() {
		even = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(x_0.IntVal & 1 == 0)
})
	})
	return even
}

var parity gopurs_runtime.Value
var once_parity sync.Once
func Get_parity() gopurs_runtime.Value {
	once_parity.Do(func() {
		parity = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if n_0.IntVal & 1 == 0 {
__t0 = gopurs_runtime.Constructor0("Even")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Odd")
}
end_branch_0:
return __t0
})
	})
	return parity
}

var eqParity gopurs_runtime.Value
var once_eqParity sync.Once
func Get_eqParity() gopurs_runtime.Value {
	once_eqParity.Do(func() {
		eqParity = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_0.StrVal == "Even").IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "Even")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.StrVal == "Odd").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "Odd").IntVal != 0)
}
end_branch_0:
return __t0
}))
	})
	return eqParity
}

var ordParity gopurs_runtime.Value
var once_ordParity sync.Once
func Get_ordParity() gopurs_runtime.Value {
	once_ordParity.Do(func() {
		ordParity = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_0.StrVal == "Even").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_1.StrVal == "Even").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("EQ")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("LT")
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_1.StrVal == "Even").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "Odd").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "Odd").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("EQ")
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
	return ordParity
}

var semiringParity gopurs_runtime.Value
var once_semiringParity sync.Once
func Get_semiringParity() gopurs_runtime.Value {
	once_semiringParity.Do(func() {
		semiringParity = gopurs_runtime.RecordDict4("zero", "add", "one", "mul", gopurs_runtime.Constructor0("Even"), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_0.StrVal == "Even").IntVal != 0 {
__t1 = gopurs_runtime.Bool(y_1.StrVal == "Even")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.StrVal == "Odd").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "Odd").IntVal != 0)
}
end_branch_1:
if __t1.IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Even")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Odd")
}
end_branch_0:
return __t0
}), gopurs_runtime.Constructor0("Odd"), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Odd").IntVal != 0 && gopurs_runtime.Bool(v1_1.StrVal == "Odd").IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Odd")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Even")
}
end_branch_2:
return __t2
}))
	})
	return semiringParity
}

var ringParity gopurs_runtime.Value
var once_ringParity sync.Once
func Get_ringParity() gopurs_runtime.Value {
	once_ringParity.Do(func() {
		ringParity = gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.RecordGet(Get_semiringParity(), "add"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semiringParity()
}))
	})
	return ringParity
}

var divisionRingParity gopurs_runtime.Value
var once_divisionRingParity sync.Once
func Get_divisionRingParity() gopurs_runtime.Value {
	once_divisionRingParity.Do(func() {
		divisionRingParity = gopurs_runtime.RecordDict2("recip", "Ring0", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ringParity()
}))
	})
	return divisionRingParity
}

var decimal gopurs_runtime.Value
var once_decimal sync.Once
func Get_decimal() gopurs_runtime.Value {
	once_decimal.Do(func() {
		decimal = gopurs_runtime.Int(10)
	})
	return decimal
}

var commutativeRingParity gopurs_runtime.Value
var once_commutativeRingParity sync.Once
func Get_commutativeRingParity() gopurs_runtime.Value {
	once_commutativeRingParity.Do(func() {
		commutativeRingParity = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ringParity()
}))
	})
	return commutativeRingParity
}

var euclideanRingParity gopurs_runtime.Value
var once_euclideanRingParity sync.Once
func Get_euclideanRingParity() gopurs_runtime.Value {
	once_euclideanRingParity.Do(func() {
		euclideanRingParity = gopurs_runtime.RecordDict4("degree", "div", "mod", "CommutativeRing0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Even").IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Odd").IntVal != 0 {
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
return gopurs_runtime.Constructor0("Even")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_commutativeRingParity()
}))
	})
	return euclideanRingParity
}

var ceil gopurs_runtime.Value
var once_ceil sync.Once
func Get_ceil() gopurs_runtime.Value {
	once_ceil.Do(func() {
		ceil = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_ceil(), x_0))
})
	})
	return ceil
}

var boundedParity gopurs_runtime.Value
var once_boundedParity sync.Once
func Get_boundedParity() gopurs_runtime.Value {
	once_boundedParity.Do(func() {
		boundedParity = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Constructor0("Even"), gopurs_runtime.Constructor0("Odd"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordParity()
}))
	})
	return boundedParity
}

var binary gopurs_runtime.Value
var once_binary sync.Once
func Get_binary() gopurs_runtime.Value {
	once_binary.Do(func() {
		binary = gopurs_runtime.Int(2)
	})
	return binary
}

var base36 gopurs_runtime.Value
var once_base36 sync.Once
func Get_base36() gopurs_runtime.Value {
	once_base36.Do(func() {
		base36 = gopurs_runtime.Int(36)
	})
	return base36
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
