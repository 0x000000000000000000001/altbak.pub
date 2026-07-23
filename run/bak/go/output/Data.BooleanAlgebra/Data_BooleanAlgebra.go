package Data_BooleanAlgebra

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var booleanAlgebraUnit gopurs_runtime.Value
var once_booleanAlgebraUnit sync.Once
func Get_booleanAlgebraUnit() gopurs_runtime.Value {
	once_booleanAlgebraUnit.Do(func() {
		booleanAlgebraUnit = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraUnit()
}))
	})
	return booleanAlgebraUnit
}

var booleanAlgebraRecordNil gopurs_runtime.Value
var once_booleanAlgebraRecordNil sync.Once
func Get_booleanAlgebraRecordNil() gopurs_runtime.Value {
	once_booleanAlgebraRecordNil.Do(func() {
		booleanAlgebraRecordNil = gopurs_runtime.RecordDict1("HeytingAlgebraRecord0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraRecordNil()
}))
	})
	return booleanAlgebraRecordNil
}

var booleanAlgebraRecordCons gopurs_runtime.Value
var once_booleanAlgebraRecordCons sync.Once
func Get_booleanAlgebraRecordCons() gopurs_runtime.Value {
	once_booleanAlgebraRecordCons.Do(func() {
		booleanAlgebraRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0 gopurs_runtime.Value, _dollar__unused_1 gopurs_runtime.Value, dictBooleanAlgebraRecord_2 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraRecordCons1_3_0 := gopurs_runtime.Apply3(pkg_Data_HeytingAlgebra.Get_heytingAlgebraRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebraRecord_2, "HeytingAlgebraRecord0"), gopurs_runtime.Value{}))
_ = heytingAlgebraRecordCons1_3_0
return gopurs_runtime.Func(func(dictBooleanAlgebra_4 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraRecordCons2_5_1 := gopurs_runtime.Apply(heytingAlgebraRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_4, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraRecordCons2_5_1
return gopurs_runtime.RecordDict1("HeytingAlgebraRecord0", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraRecordCons2_5_1
}))
})
})
	})
	return booleanAlgebraRecordCons
}

var booleanAlgebraRecord gopurs_runtime.Value
var once_booleanAlgebraRecord sync.Once
func Get_booleanAlgebraRecord() gopurs_runtime.Value {
	once_booleanAlgebraRecord.Do(func() {
		booleanAlgebraRecord = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, dictBooleanAlgebraRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebraRecord_1, "HeytingAlgebraRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_0
heytingAlgebraRecord1_3_1 := gopurs_runtime.RecordDict([]string{"ff", "tt", "conj", "disj", "implies", "not"}, []gopurs_runtime.Value{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "ffRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "ttRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "conjRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "disjRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "impliesRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "notRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy")))})
_ = heytingAlgebraRecord1_3_1
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraRecord1_3_1
}))
})
	})
	return booleanAlgebraRecord
}

var booleanAlgebraProxy gopurs_runtime.Value
var once_booleanAlgebraProxy sync.Once
func Get_booleanAlgebraProxy() gopurs_runtime.Value {
	once_booleanAlgebraProxy.Do(func() {
		booleanAlgebraProxy = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraProxy()
}))
	})
	return booleanAlgebraProxy
}

var booleanAlgebraFn gopurs_runtime.Value
var once_booleanAlgebraFn sync.Once
func Get_booleanAlgebraFn() gopurs_runtime.Value {
	once_booleanAlgebraFn.Do(func() {
		booleanAlgebraFn = gopurs_runtime.Func(func(dictBooleanAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{})
_ = __local_var_1_0
ff1_2_1 := gopurs_runtime.RecordGet(__local_var_1_0, "ff")
_ = ff1_2_1
tt1_3_3 := gopurs_runtime.RecordGet(__local_var_1_0, "tt")
_ = tt1_3_3
heytingAlgebraFunction_3_2 := gopurs_runtime.RecordDict([]string{"ff", "tt", "implies", "conj", "disj", "not"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ff1_2_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return tt1_3_3
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "implies"), gopurs_runtime.Apply(f_4, a_6), gopurs_runtime.Apply(g_5, a_6))
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "conj"), gopurs_runtime.Apply(f_4, a_6), gopurs_runtime.Apply(g_5, a_6))
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "disj"), gopurs_runtime.Apply(f_4, a_6), gopurs_runtime.Apply(g_5, a_6))
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "not"), gopurs_runtime.Apply(f_4, a_5))
})})
_ = heytingAlgebraFunction_3_2
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraFunction_3_2
}))
})
	})
	return booleanAlgebraFn
}

var booleanAlgebraBoolean gopurs_runtime.Value
var once_booleanAlgebraBoolean sync.Once
func Get_booleanAlgebraBoolean() gopurs_runtime.Value {
	once_booleanAlgebraBoolean.Do(func() {
		booleanAlgebraBoolean = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean()
}))
	})
	return booleanAlgebraBoolean
}


