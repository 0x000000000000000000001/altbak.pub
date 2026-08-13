package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_DivisionRing_DivisionRing_dollarDict gopurs_runtime.Value
var once_Data_DivisionRing_DivisionRing_dollarDict sync.Once
func Get_Data_DivisionRing_DivisionRing_dollarDict() gopurs_runtime.Value {
	once_Data_DivisionRing_DivisionRing_dollarDict.Do(func() {
		cache_Data_DivisionRing_DivisionRing_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DivisionRing_DivisionRing_dollarDict(x_0_box)
})
	})
	return cache_Data_DivisionRing_DivisionRing_dollarDict
}

var cache_Data_DivisionRing_recip gopurs_runtime.Value
var once_Data_DivisionRing_recip sync.Once
func Get_Data_DivisionRing_recip() gopurs_runtime.Value {
	once_Data_DivisionRing_recip.Do(func() {
		cache_Data_DivisionRing_recip = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DivisionRing_recip(gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing](dict_0_box))
})
	})
	return cache_Data_DivisionRing_recip
}

var cache_Data_DivisionRing_rightDiv gopurs_runtime.Value
var once_Data_DivisionRing_rightDiv sync.Once
func Get_Data_DivisionRing_rightDiv() gopurs_runtime.Value {
	once_Data_DivisionRing_rightDiv.Do(func() {
		cache_Data_DivisionRing_rightDiv = gopurs_runtime.Func(func(dictDivisionRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DivisionRing_rightDiv(gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing](dictDivisionRing_0_box))
})
	})
	return cache_Data_DivisionRing_rightDiv
}

var cache_Data_DivisionRing_leftDiv gopurs_runtime.Value
var once_Data_DivisionRing_leftDiv sync.Once
func Get_Data_DivisionRing_leftDiv() gopurs_runtime.Value {
	once_Data_DivisionRing_leftDiv.Do(func() {
		cache_Data_DivisionRing_leftDiv = gopurs_runtime.Func(func(dictDivisionRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DivisionRing_leftDiv(gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing](dictDivisionRing_0_box))
})
	})
	return cache_Data_DivisionRing_leftDiv
}

var cache_Data_DivisionRing_divisionringNumber gopurs_runtime.Value
var once_Data_DivisionRing_divisionringNumber sync.Once
func Get_Data_DivisionRing_divisionringNumber() gopurs_runtime.Value {
	once_Data_DivisionRing_divisionringNumber.Do(func() {
		cache_Data_DivisionRing_divisionringNumber = gopurs_runtime.RecordDict2("Ring0", "recip", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ring_ringNumber()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((1.0) / (x_0.FloatVal()))
}))
	})
	return cache_Data_DivisionRing_divisionringNumber
}

var cache_Data_DivisionRing_recip__1644564657 gopurs_runtime.Value
var once_Data_DivisionRing_recip__1644564657 sync.Once
func Get_Data_DivisionRing_recip__1644564657() gopurs_runtime.Value {
	once_Data_DivisionRing_recip__1644564657.Do(func() {
		cache_Data_DivisionRing_recip__1644564657 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DivisionRing_recip__1644564657(gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing](dict_0_box))
})
	})
	return cache_Data_DivisionRing_recip__1644564657
}

type Constructor_Data_DivisionRing_DivisionRing struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2548491258] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_DivisionRing_DivisionRing)(ptr)
		_ = c
		switch key {
		case "Ring0": return gopurs_runtime.Box(c.V0)
		case "recip": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_DivisionRing_DivisionRing: " + key)
		}
	}
}


func Call_Data_DivisionRing_DivisionRing_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_DivisionRing_recip(dict_0_loop *Constructor_Data_DivisionRing_DivisionRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_DivisionRing_DivisionRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_DivisionRing_rightDiv(dictDivisionRing_0_loop *Constructor_Data_DivisionRing_DivisionRing) gopurs_runtime.Value {
var dictDivisionRing_0 *Constructor_Data_DivisionRing_DivisionRing = dictDivisionRing_0_loop
_ = dictDivisionRing_0
// TAST (Let): Semiring0_1_0 -> *Constructor_Data_Semiring_Semiring
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictDivisionRing_0.V0), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semiring0_1_0.V1), a_2, gopurs_runtime.Apply(gopurs_runtime.Box(dictDivisionRing_0.V1), b_3))
})
})
}

func Call_Data_DivisionRing_leftDiv(dictDivisionRing_0_loop *Constructor_Data_DivisionRing_DivisionRing) gopurs_runtime.Value {
var dictDivisionRing_0 *Constructor_Data_DivisionRing_DivisionRing = dictDivisionRing_0_loop
_ = dictDivisionRing_0
// TAST (Let): Semiring0_1_0 -> *Constructor_Data_Semiring_Semiring
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictDivisionRing_0.V0), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semiring0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictDivisionRing_0.V1), b_3), a_2)
})
})
}

func Call_Data_DivisionRing_recip__1644564657(dict_0_loop *Constructor_Data_DivisionRing_DivisionRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_DivisionRing_DivisionRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}


