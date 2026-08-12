package Data_Int

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Ord "gopurs/output/Data.Ord"
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
		cache_fromStringAs = gopurs_runtime.Apply2(Get_fromStringAsImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
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
		cache_fromNumber = gopurs_runtime.Apply2(Get_fromNumberImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
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
var __t1 gopurs_runtime.Value
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (uint32(y_1.IntVal) == 2591059121) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if ((uint32(x_0.IntVal) == 658452902)) && ((uint32(y_1.IntVal) == 658452902)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t1.IntVal)), UnsafePtr: nil}
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
var __t0 uint32
{
if (gopurs_runtime.Apply2(Get_eq__196302102(), x_0, y_1).IntVal) != (0) {
__t0 = 2591059121
goto end_branch_0
} else {

}
}
{
__t0 = 658452902
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 uint32
{
if ((uint32(v_0.IntVal) == 658452902)) && ((uint32(v1_1.IntVal) == 658452902)) {
__t1 = 658452902
goto end_branch_1
} else {

}
}
{
__t1 = 2591059121
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil})
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}
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
}), gopurs_runtime.Value{Type: 9, IntVal: int64(2591059121), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(658452902), UnsafePtr: nil})
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

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = pkg_Data_Eq.Get_eqIntImpl()
	})
	return cache_eq__2843686287
}

var cache_eq__2276491096 gopurs_runtime.Value
var once_eq__2276491096 sync.Once
func Get_eq__2276491096() gopurs_runtime.Value {
	once_eq__2276491096.Do(func() {
		cache_eq__2276491096 = gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq")
	})
	return cache_eq__2276491096
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq__196302102 gopurs_runtime.Value
var once_eq__196302102 sync.Once
func Get_eq__196302102() gopurs_runtime.Value {
	once_eq__196302102.Do(func() {
		cache_eq__196302102 = gopurs_runtime.RecordGet(Get_eqParity(), "eq")
	})
	return cache_eq__196302102
}

var cache_notEq__2843686287 gopurs_runtime.Value
var once_notEq__2843686287 sync.Once
func Get_notEq__2843686287() gopurs_runtime.Value {
	once_notEq__2843686287.Do(func() {
		cache_notEq__2843686287 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2843686287(x_0_box, y_1_box))
})
	})
	return cache_notEq__2843686287
}

var cache_notEq__2384498378 gopurs_runtime.Value
var once_notEq__2384498378 sync.Once
func Get_notEq__2384498378() gopurs_runtime.Value {
	once_notEq__2384498378.Do(func() {
		cache_notEq__2384498378 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_notEq__2384498378
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj")
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_fromMaybe__1972796397 gopurs_runtime.Value
var once_fromMaybe__1972796397 sync.Once
func Get_fromMaybe__1972796397() gopurs_runtime.Value {
	once_fromMaybe__1972796397.Do(func() {
		cache_fromMaybe__1972796397 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__1972796397(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__1972796397
}

var cache_fromMaybe__430429096 gopurs_runtime.Value
var once_fromMaybe__430429096 sync.Once
func Get_fromMaybe__430429096() gopurs_runtime.Value {
	once_fromMaybe__430429096.Do(func() {
		cache_fromMaybe__430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__430429096(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__430429096
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_greaterThanOrEq__4087042607 gopurs_runtime.Value
var once_greaterThanOrEq__4087042607 sync.Once
func Get_greaterThanOrEq__4087042607() gopurs_runtime.Value {
	once_greaterThanOrEq__4087042607.Do(func() {
		cache_greaterThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__4087042607
}

var cache_greaterThanOrEq__1061005983 gopurs_runtime.Value
var once_greaterThanOrEq__1061005983 sync.Once
func Get_greaterThanOrEq__1061005983() gopurs_runtime.Value {
	once_greaterThanOrEq__1061005983.Do(func() {
		cache_greaterThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__1061005983
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1061005983 gopurs_runtime.Value
var once_lessThanOrEq__1061005983 sync.Once
func Get_lessThanOrEq__1061005983() gopurs_runtime.Value {
	once_lessThanOrEq__1061005983.Do(func() {
		cache_lessThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__1061005983
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
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
if (gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(gopurs_runtime.Int(n_0), gopurs_runtime.Int(2))), gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(n_0), gopurs_runtime.Int(36)))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(n_0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0)
}

func Call_odd(x_0_loop int64) bool {
var x_0 int64 = x_0_loop
_ = x_0
return (gopurs_runtime.Bool(Call_notEq__2843686287(gopurs_runtime.Int((x_0) & (1)), gopurs_runtime.Int(0))).IntVal) != (0)
}

func Call_unsafeClamp(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
var __t0 int64
{
if (gopurs_runtime.Apply(Get_not__3201284355(), gopurs_runtime.Apply(pkg_Data_Number.Get_isFinite(), gopurs_runtime.Float(x_0))).IntVal) != (0) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(gopurs_runtime.Float(x_0), gopurs_runtime.Apply(Get_toNumber(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "top")))).IntVal) != (0) {
__t0 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "top").IntVal
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(gopurs_runtime.Float(x_0), gopurs_runtime.Apply(Get_toNumber(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "bottom")))).IntVal) != (0) {
__t0 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "bottom").IntVal
goto end_branch_0
} else {

}
}
{
__t0 = Call_fromMaybe__1972796397(gopurs_runtime.Int(0), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_fromNumber(), gopurs_runtime.Float(x_0)))).IntVal
}
end_branch_0:
return __t0
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
return (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int((x_0) & (1)), gopurs_runtime.Int(0)).IntVal) != (0)
}

func Call_parity(n_0_loop int64) uint32 {
var n_0 int64 = n_0_loop
_ = n_0
var __t0 uint32
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int((n_0) & (1)), gopurs_runtime.Int(0)).IntVal) != (0) {
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

func Call_ceil(x_0_loop float64) int64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(Call_unsafeClamp(gopurs_runtime.Apply(pkg_Data_Number.Get_ceil(), gopurs_runtime.Float(x_0)).FloatVal())).IntVal
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_notEq__2843686287(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((x_0.IntVal) == (y_1.IntVal)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_notEq__2384498378(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(dictEq_0.V0, x_1, y_2), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_fromMaybe__1972796397(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromMaybe__430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) < (a2_1.FloatVal()) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) > (a2_1.FloatVal()) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
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
