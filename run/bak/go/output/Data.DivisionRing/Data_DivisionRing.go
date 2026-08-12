package Data_DivisionRing

import (
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_recip gopurs_runtime.Value
var once_recip sync.Once
func Get_recip() gopurs_runtime.Value {
	once_recip.Do(func() {
		cache_recip = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_recip(gopurs_runtime.CoerceToStruct[Constructor_DivisionRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_recip
}

var cache_rightDiv gopurs_runtime.Value
var once_rightDiv sync.Once
func Get_rightDiv() gopurs_runtime.Value {
	once_rightDiv.Do(func() {
		cache_rightDiv = gopurs_runtime.Func(func(dictDivisionRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rightDiv(gopurs_runtime.CoerceToStruct[Constructor_DivisionRing[gopurs_runtime.Value]](dictDivisionRing_0_box))
})
	})
	return cache_rightDiv
}

var cache_leftDiv gopurs_runtime.Value
var once_leftDiv sync.Once
func Get_leftDiv() gopurs_runtime.Value {
	once_leftDiv.Do(func() {
		cache_leftDiv = gopurs_runtime.Func(func(dictDivisionRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_leftDiv(gopurs_runtime.CoerceToStruct[Constructor_DivisionRing[gopurs_runtime.Value]](dictDivisionRing_0_box))
})
	})
	return cache_leftDiv
}

var cache_divisionringNumber gopurs_runtime.Value
var once_divisionringNumber sync.Once
func Get_divisionringNumber() gopurs_runtime.Value {
	once_divisionringNumber.Do(func() {
		cache_divisionringNumber = gopurs_runtime.RecordDict2("Ring0", "recip", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", pkg_Data_Semiring.Get_numAdd(), pkg_Data_Semiring.Get_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
}), pkg_Data_Ring.Get_numSub())
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(gopurs_runtime.Apply2(Get_div__1002719800(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(x_0.FloatVal())).FloatVal())
}))
	})
	return cache_divisionringNumber
}

var cache_recip__1644564657 gopurs_runtime.Value
var once_recip__1644564657 sync.Once
func Get_recip__1644564657() gopurs_runtime.Value {
	once_recip__1644564657.Do(func() {
		cache_recip__1644564657 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_recip__1644564657(gopurs_runtime.CoerceToStruct[Constructor_DivisionRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_recip__1644564657
}

var cache_div__1002719800 gopurs_runtime.Value
var once_div__1002719800 sync.Once
func Get_div__1002719800() gopurs_runtime.Value {
	once_div__1002719800.Do(func() {
		cache_div__1002719800 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div")
	})
	return cache_div__1002719800
}

var cache_div__2579358968 gopurs_runtime.Value
var once_div__2579358968 sync.Once
func Get_div__2579358968() gopurs_runtime.Value {
	once_div__2579358968.Do(func() {
		cache_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__2579358968
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

type Constructor_DivisionRing[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2548491258] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_DivisionRing[gopurs_runtime.Value])(ptr)
		switch key {
		case "Ring0": return c.V0
		case "recip": return c.V1
		default: panic("Key not found in dictionary Constructor_DivisionRing: " + key)
		}
	}
}


func Call_recip(dict_0_loop *Constructor_DivisionRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_DivisionRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_rightDiv(dictDivisionRing_0_loop *Constructor_DivisionRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDivisionRing_0 *Constructor_DivisionRing[gopurs_runtime.Value] = dictDivisionRing_0_loop
_ = dictDivisionRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictDivisionRing_0.V0, gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semiring0_1_0.V1, a_2, gopurs_runtime.Apply(dictDivisionRing_0.V1, b_3))
})
})
}

func Call_leftDiv(dictDivisionRing_0_loop *Constructor_DivisionRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDivisionRing_0 *Constructor_DivisionRing[gopurs_runtime.Value] = dictDivisionRing_0_loop
_ = dictDivisionRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictDivisionRing_0.V0, gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semiring0_1_0.V1, gopurs_runtime.Apply(dictDivisionRing_0.V1, b_3), a_2)
})
})
}

func Call_recip__1644564657(dict_0_loop *Constructor_DivisionRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_DivisionRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


