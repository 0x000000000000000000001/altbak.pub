package Data_DivisionRing

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
)

var recip gopurs_runtime.Value
var once_recip sync.Once
func Get_recip() gopurs_runtime.Value {
	once_recip.Do(func() {
		recip = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "recip")
})
	})
	return recip
}

var rightDiv gopurs_runtime.Value
var once_rightDiv sync.Once
func Get_rightDiv() gopurs_runtime.Value {
	once_rightDiv.Do(func() {
		rightDiv = gopurs_runtime.Func3(func(dictDivisionRing_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivisionRing_0, "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "mul"), a_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivisionRing_0, "recip"), b_2))
})
	})
	return rightDiv
}

var leftDiv gopurs_runtime.Value
var once_leftDiv sync.Once
func Get_leftDiv() gopurs_runtime.Value {
	once_leftDiv.Do(func() {
		leftDiv = gopurs_runtime.Func3(func(dictDivisionRing_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivisionRing_0, "Ring0"), gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "mul"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDivisionRing_0, "recip"), b_2), a_1)
})
	})
	return leftDiv
}

var divisionringNumber gopurs_runtime.Value
var once_divisionringNumber sync.Once
func Get_divisionringNumber() gopurs_runtime.Value {
	once_divisionringNumber.Do(func() {
		divisionringNumber = gopurs_runtime.RecordDict2("recip", "Ring0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(1.0 / x_0.FloatVal())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringNumber()
}))
	})
	return divisionringNumber
}




