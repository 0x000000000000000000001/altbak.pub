package Data_Semigroup

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Void "gopurs/output/Data.Void"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var semigroupVoid gopurs_runtime.Value
var once_semigroupVoid sync.Once
func Get_semigroupVoid() gopurs_runtime.Value {
	once_semigroupVoid.Do(func() {
		semigroupVoid = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Void.Get_absurd()
}))
	})
	return semigroupVoid
}

var semigroupUnit gopurs_runtime.Value
var once_semigroupUnit sync.Once
func Get_semigroupUnit() gopurs_runtime.Value {
	once_semigroupUnit.Do(func() {
		semigroupUnit = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return semigroupUnit
}

var semigroupString gopurs_runtime.Value
var once_semigroupString sync.Once
func Get_semigroupString() gopurs_runtime.Value {
	once_semigroupString.Do(func() {
		semigroupString = gopurs_runtime.RecordDict1("append", Get_concatString())
	})
	return semigroupString
}

var semigroupRecordNil gopurs_runtime.Value
var once_semigroupRecordNil sync.Once
func Get_semigroupRecordNil() gopurs_runtime.Value {
	once_semigroupRecordNil.Do(func() {
		semigroupRecordNil = gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}))
	})
	return semigroupRecordNil
}

var semigroupProxy gopurs_runtime.Value
var once_semigroupProxy sync.Once
func Get_semigroupProxy() gopurs_runtime.Value {
	once_semigroupProxy.Do(func() {
		semigroupProxy = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))
}))
	})
	return semigroupProxy
}

var semigroupArray gopurs_runtime.Value
var once_semigroupArray sync.Once
func Get_semigroupArray() gopurs_runtime.Value {
	once_semigroupArray.Do(func() {
		semigroupArray = gopurs_runtime.RecordDict1("append", Get_concatArray())
	})
	return semigroupArray
}

var appendRecord gopurs_runtime.Value
var once_appendRecord sync.Once
func Get_appendRecord() gopurs_runtime.Value {
	once_appendRecord.Do(func() {
		appendRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "appendRecord")
})
	})
	return appendRecord
}

var semigroupRecord gopurs_runtime.Value
var once_semigroupRecord sync.Once
func Get_semigroupRecord() gopurs_runtime.Value {
	once_semigroupRecord.Do(func() {
		semigroupRecord = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, dictSemigroupRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroupRecord_1, "appendRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))))
})
	})
	return semigroupRecord
}

var append_ gopurs_runtime.Value
var once_append_ sync.Once
func Get_append_() gopurs_runtime.Value {
	once_append_.Do(func() {
		append_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "append")
})
	})
	return append_
}

var semigroupFn gopurs_runtime.Value
var once_semigroupFn sync.Once
func Get_semigroupFn() gopurs_runtime.Value {
	once_semigroupFn.Do(func() {
		semigroupFn = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
}))
})
	})
	return semigroupFn
}

var semigroupRecordCons gopurs_runtime.Value
var once_semigroupRecordCons sync.Once
func Get_semigroupRecordCons() gopurs_runtime.Value {
	once_semigroupRecordCons.Do(func() {
		semigroupRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0 gopurs_runtime.Value, _dollar__unused_1 gopurs_runtime.Value, dictSemigroupRecord_2 gopurs_runtime.Value, dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, ra_5 gopurs_runtime.Value, rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
key_7_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy")))
_ = key_7_0
get_8_1 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_7_0)
_ = get_8_1
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_7_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_3, "append"), gopurs_runtime.Apply(get_8_1, ra_5), gopurs_runtime.Apply(get_8_1, rb_6)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemigroupRecord_2, "appendRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy")), ra_5, rb_6))
}))
})
	})
	return semigroupRecordCons
}

func Get_concatArray() gopurs_runtime.Value {
	return _Gopurs_ConcatArray
}

func Get_concatString() gopurs_runtime.Value {
	return _Gopurs_ConcatString
}
