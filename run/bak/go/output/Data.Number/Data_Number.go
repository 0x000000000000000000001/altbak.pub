package Data_Number

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var tau gopurs_runtime.Value
var once_tau sync.Once
func Get_tau() gopurs_runtime.Value {
	once_tau.Do(func() {
		tau = gopurs_runtime.Float(6.283185307179586)
	})
	return tau
}

var sqrt2 gopurs_runtime.Value
var once_sqrt2 sync.Once
func Get_sqrt2() gopurs_runtime.Value {
	once_sqrt2.Do(func() {
		sqrt2 = gopurs_runtime.Float(1.4142135623730951)
	})
	return sqrt2
}

var sqrt1_2 gopurs_runtime.Value
var once_sqrt1_2 sync.Once
func Get_sqrt1_2() gopurs_runtime.Value {
	once_sqrt1_2.Do(func() {
		sqrt1_2 = gopurs_runtime.Float(0.7071067811865476)
	})
	return sqrt1_2
}

var pi gopurs_runtime.Value
var once_pi sync.Once
func Get_pi() gopurs_runtime.Value {
	once_pi.Do(func() {
		pi = gopurs_runtime.Float(3.141592653589793)
	})
	return pi
}

var log2e gopurs_runtime.Value
var once_log2e sync.Once
func Get_log2e() gopurs_runtime.Value {
	once_log2e.Do(func() {
		log2e = gopurs_runtime.Float(1.4426950408889634)
	})
	return log2e
}

var log10e gopurs_runtime.Value
var once_log10e sync.Once
func Get_log10e() gopurs_runtime.Value {
	once_log10e.Do(func() {
		log10e = gopurs_runtime.Float(0.4342944819032518)
	})
	return log10e
}

var ln2 gopurs_runtime.Value
var once_ln2 sync.Once
func Get_ln2() gopurs_runtime.Value {
	once_ln2.Do(func() {
		ln2 = gopurs_runtime.Float(0.6931471805599453)
	})
	return ln2
}

var ln10 gopurs_runtime.Value
var once_ln10 sync.Once
func Get_ln10() gopurs_runtime.Value {
	once_ln10.Do(func() {
		ln10 = gopurs_runtime.Float(2.302585092994046)
	})
	return ln10
}

var fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		fromString = gopurs_runtime.Func(func(str_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var str_0 gopurs_runtime.Value = str_0_loop
_ = str_0
return gopurs_runtime.UncurriedApp4(Get_fromStringImpl(), str_0_loop, Get_isFinite(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"))
}()
})
	})
	return fromString
}

var e gopurs_runtime.Value
var once_e sync.Once
func Get_e() gopurs_runtime.Value {
	once_e.Do(func() {
		e = gopurs_runtime.Float(2.718281828459045)
	})
	return e
}



func Get_abs() gopurs_runtime.Value {
	return _Gopurs_Abs
}

func Get_acos() gopurs_runtime.Value {
	return _Gopurs_Acos
}

func Get_asin() gopurs_runtime.Value {
	return _Gopurs_Asin
}

func Get_atan() gopurs_runtime.Value {
	return _Gopurs_Atan
}

func Get_atan2() gopurs_runtime.Value {
	return _Gopurs_Atan2
}

func Get_ceil() gopurs_runtime.Value {
	return _Gopurs_Ceil
}

func Get_cos() gopurs_runtime.Value {
	return _Gopurs_Cos
}

func Get_exp() gopurs_runtime.Value {
	return _Gopurs_Exp
}

func Get_floor() gopurs_runtime.Value {
	return _Gopurs_Floor
}

func Get_fromStringImpl() gopurs_runtime.Value {
	return _Gopurs_FromStringImpl
}

func Get_infinity() gopurs_runtime.Value {
	return _Gopurs_Infinity
}

func Get_isFinite() gopurs_runtime.Value {
	return _Gopurs_IsFinite
}

func Get_isNaN() gopurs_runtime.Value {
	return _Gopurs_IsNaN
}

func Get_log() gopurs_runtime.Value {
	return _Gopurs_Log
}

func Get_max() gopurs_runtime.Value {
	return _Gopurs_Max
}

func Get_min() gopurs_runtime.Value {
	return _Gopurs_Min
}

func Get_nan() gopurs_runtime.Value {
	return _Gopurs_Nan
}

func Get_pow() gopurs_runtime.Value {
	return _Gopurs_Pow
}

func Get_remainder() gopurs_runtime.Value {
	return _Gopurs_Remainder
}

func Get_round() gopurs_runtime.Value {
	return _Gopurs_Round
}

func Get_sign() gopurs_runtime.Value {
	return _Gopurs_Sign
}

func Get_sin() gopurs_runtime.Value {
	return _Gopurs_Sin
}

func Get_sqrt() gopurs_runtime.Value {
	return _Gopurs_Sqrt
}

func Get_tan() gopurs_runtime.Value {
	return _Gopurs_Tan
}

func Get_trunc() gopurs_runtime.Value {
	return _Gopurs_Trunc
}
