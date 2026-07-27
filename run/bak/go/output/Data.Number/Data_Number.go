package Data_Number

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_tau gopurs_runtime.Value
var once_tau sync.Once
func Get_tau() gopurs_runtime.Value {
	once_tau.Do(func() {
		cache_tau = gopurs_runtime.Float(6.283185307179586)
	})
	return cache_tau
}

var cache_sqrt2 gopurs_runtime.Value
var once_sqrt2 sync.Once
func Get_sqrt2() gopurs_runtime.Value {
	once_sqrt2.Do(func() {
		cache_sqrt2 = gopurs_runtime.Float(1.4142135623730951)
	})
	return cache_sqrt2
}

var cache_sqrt1_2 gopurs_runtime.Value
var once_sqrt1_2 sync.Once
func Get_sqrt1_2() gopurs_runtime.Value {
	once_sqrt1_2.Do(func() {
		cache_sqrt1_2 = gopurs_runtime.Float(0.7071067811865476)
	})
	return cache_sqrt1_2
}

var cache_pi gopurs_runtime.Value
var once_pi sync.Once
func Get_pi() gopurs_runtime.Value {
	once_pi.Do(func() {
		cache_pi = gopurs_runtime.Float(3.141592653589793)
	})
	return cache_pi
}

var cache_log2e gopurs_runtime.Value
var once_log2e sync.Once
func Get_log2e() gopurs_runtime.Value {
	once_log2e.Do(func() {
		cache_log2e = gopurs_runtime.Float(1.4426950408889634)
	})
	return cache_log2e
}

var cache_log10e gopurs_runtime.Value
var once_log10e sync.Once
func Get_log10e() gopurs_runtime.Value {
	once_log10e.Do(func() {
		cache_log10e = gopurs_runtime.Float(0.4342944819032518)
	})
	return cache_log10e
}

var cache_ln2 gopurs_runtime.Value
var once_ln2 sync.Once
func Get_ln2() gopurs_runtime.Value {
	once_ln2.Do(func() {
		cache_ln2 = gopurs_runtime.Float(0.6931471805599453)
	})
	return cache_ln2
}

var cache_ln10 gopurs_runtime.Value
var once_ln10 sync.Once
func Get_ln10() gopurs_runtime.Value {
	once_ln10.Do(func() {
		cache_ln10 = gopurs_runtime.Float(2.302585092994046)
	})
	return cache_ln10
}

var cache_fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		cache_fromString = gopurs_runtime.Func(func(str_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromString(str_0_box.StrVal()))}
})
	})
	return cache_fromString
}

var cache_e gopurs_runtime.Value
var once_e sync.Once
func Get_e() gopurs_runtime.Value {
	once_e.Do(func() {
		cache_e = gopurs_runtime.Float(2.718281828459045)
	})
	return cache_e
}

var cache_abs gopurs_runtime.Value
var once_abs sync.Once
func Get_abs() gopurs_runtime.Value {
	once_abs.Do(func() {
		cache_abs = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Abs(arg0.FloatVal()))
})
	})
	return cache_abs
}

var cache_acos gopurs_runtime.Value
var once_acos sync.Once
func Get_acos() gopurs_runtime.Value {
	once_acos.Do(func() {
		cache_acos = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Acos(arg0.FloatVal()))
})
	})
	return cache_acos
}

var cache_asin gopurs_runtime.Value
var once_asin sync.Once
func Get_asin() gopurs_runtime.Value {
	once_asin.Do(func() {
		cache_asin = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Asin(arg0.FloatVal()))
})
	})
	return cache_asin
}

var cache_atan gopurs_runtime.Value
var once_atan sync.Once
func Get_atan() gopurs_runtime.Value {
	once_atan.Do(func() {
		cache_atan = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Atan(arg0.FloatVal()))
})
	})
	return cache_atan
}

var cache_atan2 gopurs_runtime.Value
var once_atan2 sync.Once
func Get_atan2() gopurs_runtime.Value {
	once_atan2.Do(func() {
		cache_atan2 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Atan2(arg0.FloatVal(), arg1.FloatVal()))
})
	})
	return cache_atan2
}

var cache_ceil gopurs_runtime.Value
var once_ceil sync.Once
func Get_ceil() gopurs_runtime.Value {
	once_ceil.Do(func() {
		cache_ceil = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Ceil(arg0.FloatVal()))
})
	})
	return cache_ceil
}

var cache_cos gopurs_runtime.Value
var once_cos sync.Once
func Get_cos() gopurs_runtime.Value {
	once_cos.Do(func() {
		cache_cos = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Cos(arg0.FloatVal()))
})
	})
	return cache_cos
}

var cache_exp gopurs_runtime.Value
var once_exp sync.Once
func Get_exp() gopurs_runtime.Value {
	once_exp.Do(func() {
		cache_exp = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Exp(arg0.FloatVal()))
})
	})
	return cache_exp
}

var cache_floor gopurs_runtime.Value
var once_floor sync.Once
func Get_floor() gopurs_runtime.Value {
	once_floor.Do(func() {
		cache_floor = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Floor(arg0.FloatVal()))
})
	})
	return cache_floor
}

var cache_fromStringImpl gopurs_runtime.Value
var once_fromStringImpl sync.Once
func Get_fromStringImpl() gopurs_runtime.Value {
	once_fromStringImpl.Do(func() {
		cache_fromStringImpl = FromStringImpl
	})
	return cache_fromStringImpl
}

var cache_infinity gopurs_runtime.Value
var once_infinity sync.Once
func Get_infinity() gopurs_runtime.Value {
	once_infinity.Do(func() {
		cache_infinity = gopurs_runtime.Float(Infinity)
	})
	return cache_infinity
}

var cache_isFinite gopurs_runtime.Value
var once_isFinite sync.Once
func Get_isFinite() gopurs_runtime.Value {
	once_isFinite.Do(func() {
		cache_isFinite = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(IsFinite(arg0.FloatVal()))
})
	})
	return cache_isFinite
}

var cache_isNaN gopurs_runtime.Value
var once_isNaN sync.Once
func Get_isNaN() gopurs_runtime.Value {
	once_isNaN.Do(func() {
		cache_isNaN = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(IsNaN(arg0.FloatVal()))
})
	})
	return cache_isNaN
}

var cache_log gopurs_runtime.Value
var once_log sync.Once
func Get_log() gopurs_runtime.Value {
	once_log.Do(func() {
		cache_log = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Log(arg0.FloatVal()))
})
	})
	return cache_log
}

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Max(arg0.FloatVal(), arg1.FloatVal()))
})
	})
	return cache_max
}

var cache_min gopurs_runtime.Value
var once_min sync.Once
func Get_min() gopurs_runtime.Value {
	once_min.Do(func() {
		cache_min = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Min(arg0.FloatVal(), arg1.FloatVal()))
})
	})
	return cache_min
}

var cache_nan gopurs_runtime.Value
var once_nan sync.Once
func Get_nan() gopurs_runtime.Value {
	once_nan.Do(func() {
		cache_nan = gopurs_runtime.Float(Nan)
	})
	return cache_nan
}

var cache_pow gopurs_runtime.Value
var once_pow sync.Once
func Get_pow() gopurs_runtime.Value {
	once_pow.Do(func() {
		cache_pow = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Pow(arg0.FloatVal(), arg1.FloatVal()))
})
	})
	return cache_pow
}

var cache_remainder gopurs_runtime.Value
var once_remainder sync.Once
func Get_remainder() gopurs_runtime.Value {
	once_remainder.Do(func() {
		cache_remainder = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Remainder(arg0.FloatVal(), arg1.FloatVal()))
})
	})
	return cache_remainder
}

var cache_round gopurs_runtime.Value
var once_round sync.Once
func Get_round() gopurs_runtime.Value {
	once_round.Do(func() {
		cache_round = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Round(arg0.FloatVal()))
})
	})
	return cache_round
}

var cache_sign gopurs_runtime.Value
var once_sign sync.Once
func Get_sign() gopurs_runtime.Value {
	once_sign.Do(func() {
		cache_sign = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Sign(arg0.FloatVal()))
})
	})
	return cache_sign
}

var cache_sin gopurs_runtime.Value
var once_sin sync.Once
func Get_sin() gopurs_runtime.Value {
	once_sin.Do(func() {
		cache_sin = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Sin(arg0.FloatVal()))
})
	})
	return cache_sin
}

var cache_sqrt gopurs_runtime.Value
var once_sqrt sync.Once
func Get_sqrt() gopurs_runtime.Value {
	once_sqrt.Do(func() {
		cache_sqrt = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Sqrt(arg0.FloatVal()))
})
	})
	return cache_sqrt
}

var cache_tan gopurs_runtime.Value
var once_tan sync.Once
func Get_tan() gopurs_runtime.Value {
	once_tan.Do(func() {
		cache_tan = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Tan(arg0.FloatVal()))
})
	})
	return cache_tan
}

var cache_trunc gopurs_runtime.Value
var once_trunc sync.Once
func Get_trunc() gopurs_runtime.Value {
	once_trunc.Do(func() {
		cache_trunc = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Trunc(arg0.FloatVal()))
})
	})
	return cache_trunc
}

func Call_fromString(str_0_loop string) *pkg_Data_Maybe.Constructor_Just[float64] {
var str_0 string = str_0_loop
_ = str_0
return (*pkg_Data_Maybe.Constructor_Just[float64])(gopurs_runtime.UncurriedApp4(Get_fromStringImpl(), gopurs_runtime.Str(str_0), Get_isFinite(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})).UnsafePtr)
}
