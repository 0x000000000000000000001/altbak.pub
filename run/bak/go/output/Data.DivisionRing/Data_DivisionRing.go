package Data_DivisionRing

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
)

var cache_recip gopurs_runtime.Value
var once_recip sync.Once
func Get_recip() gopurs_runtime.Value {
	once_recip.Do(func() {
		cache_recip = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_recip(dict_0_box)
})
	})
	return cache_recip
}

var cache_recip__gopurs_runtime_Value_3654627386 gopurs_runtime.Value
var once_recip__gopurs_runtime_Value_3654627386 sync.Once
func Get_recip__gopurs_runtime_Value_3654627386() gopurs_runtime.Value {
	once_recip__gopurs_runtime_Value_3654627386.Do(func() {
		cache_recip__gopurs_runtime_Value_3654627386 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_recip__gopurs_runtime_Value_3654627386(dict_0_box)
})
	})
	return cache_recip__gopurs_runtime_Value_3654627386
}

var cache_rightDiv gopurs_runtime.Value
var once_rightDiv sync.Once
func Get_rightDiv() gopurs_runtime.Value {
	once_rightDiv.Do(func() {
		cache_rightDiv = gopurs_runtime.Func3(func(dictDivisionRing_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rightDiv(dictDivisionRing_0_box, a_1_box, b_2_box)
})
	})
	return cache_rightDiv
}

var cache_leftDiv gopurs_runtime.Value
var once_leftDiv sync.Once
func Get_leftDiv() gopurs_runtime.Value {
	once_leftDiv.Do(func() {
		cache_leftDiv = gopurs_runtime.Func3(func(dictDivisionRing_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_leftDiv(dictDivisionRing_0_box, a_1_box, b_2_box)
})
	})
	return cache_leftDiv
}

var cache_divisionringNumber gopurs_runtime.Value
var once_divisionringNumber sync.Once
func Get_divisionringNumber() gopurs_runtime.Value {
	once_divisionringNumber.Do(func() {
		cache_divisionringNumber = gopurs_runtime.RecordDict2("Ring0", "recip", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringNumber()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), gopurs_runtime.Float(1.0), x_0).FloatVal())
}))
	})
	return cache_divisionringNumber
}

func Call_recip(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "recip")
}

func Call_recip__gopurs_runtime_Value_3654627386(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "recip")
}

func Call_rightDiv(dictDivisionRing_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDivisionRing_0 gopurs_runtime.Value = dictDivisionRing_0_loop
_ = dictDivisionRing_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivisionRing_0, "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "mul"), a_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivisionRing_0, "recip"), b_2))
}

func Call_leftDiv(dictDivisionRing_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDivisionRing_0 gopurs_runtime.Value = dictDivisionRing_0_loop
_ = dictDivisionRing_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivisionRing_0, "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "mul"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivisionRing_0, "recip"), b_2), a_1)
}


