package Data_Show

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Void "gopurs/output/Data.Void"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var showVoid gopurs_runtime.Value
var once_showVoid sync.Once
func Get_showVoid() gopurs_runtime.Value {
	once_showVoid.Do(func() {
		showVoid = gopurs_runtime.RecordDict1("show", pkg_Data_Void.Get_absurd())
	})
	return showVoid
}

var showUnit gopurs_runtime.Value
var once_showUnit sync.Once
func Get_showUnit() gopurs_runtime.Value {
	once_showUnit.Do(func() {
		showUnit = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unit")
}))
	})
	return showUnit
}

var showString gopurs_runtime.Value
var once_showString sync.Once
func Get_showString() gopurs_runtime.Value {
	once_showString.Do(func() {
		showString = gopurs_runtime.RecordDict1("show", Get_showStringImpl())
	})
	return showString
}

var showRecordFieldsNil gopurs_runtime.Value
var once_showRecordFieldsNil sync.Once
func Get_showRecordFieldsNil() gopurs_runtime.Value {
	once_showRecordFieldsNil.Do(func() {
		showRecordFieldsNil = gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
}))
	})
	return showRecordFieldsNil
}

var showRecordFields gopurs_runtime.Value
var once_showRecordFields sync.Once
func Get_showRecordFields() gopurs_runtime.Value {
	once_showRecordFields.Do(func() {
		showRecordFields = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "showRecordFields")
}()
})
	})
	return showRecordFields
}

var showRecord gopurs_runtime.Value
var once_showRecord sync.Once
func Get_showRecord() gopurs_runtime.Value {
	once_showRecord.Do(func() {
		showRecord = gopurs_runtime.Func3(Call_showRecord)
	})
	return showRecord
}

var showProxy gopurs_runtime.Value
var once_showProxy sync.Once
func Get_showProxy() gopurs_runtime.Value {
	once_showProxy.Do(func() {
		showProxy = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("Proxy")
}))
	})
	return showProxy
}

var showNumber gopurs_runtime.Value
var once_showNumber sync.Once
func Get_showNumber() gopurs_runtime.Value {
	once_showNumber.Do(func() {
		showNumber = gopurs_runtime.RecordDict1("show", Get_showNumberImpl())
	})
	return showNumber
}

var showInt gopurs_runtime.Value
var once_showInt sync.Once
func Get_showInt() gopurs_runtime.Value {
	once_showInt.Do(func() {
		showInt = gopurs_runtime.RecordDict1("show", Get_showIntImpl())
	})
	return showInt
}

var showChar gopurs_runtime.Value
var once_showChar sync.Once
func Get_showChar() gopurs_runtime.Value {
	once_showChar.Do(func() {
		showChar = gopurs_runtime.RecordDict1("show", Get_showCharImpl())
	})
	return showChar
}

var showBoolean gopurs_runtime.Value
var once_showBoolean sync.Once
func Get_showBoolean() gopurs_runtime.Value {
	once_showBoolean.Do(func() {
		showBoolean = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if v_0.IntVal != 0 {
__t0 = gopurs_runtime.Str("true")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Str("false")
}
end_branch_0:
return __t0
}))
	})
	return showBoolean
}

var show gopurs_runtime.Value
var once_show sync.Once
func Get_show() gopurs_runtime.Value {
	once_show.Do(func() {
		show = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "show")
}()
})
	})
	return show
}

var showArray gopurs_runtime.Value
var once_showArray sync.Once
func Get_showArray() gopurs_runtime.Value {
	once_showArray.Do(func() {
		showArray = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Apply(Get_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0_loop, "show")))
}()
})
	})
	return showArray
}

var showRecordFieldsCons gopurs_runtime.Value
var once_showRecordFieldsCons sync.Once
func Get_showRecordFieldsCons() gopurs_runtime.Value {
	once_showRecordFieldsCons.Do(func() {
		showRecordFieldsCons = gopurs_runtime.Func3(Call_showRecordFieldsCons)
	})
	return showRecordFieldsCons
}

var showRecordFieldsConsNil gopurs_runtime.Value
var once_showRecordFieldsConsNil sync.Once
func Get_showRecordFieldsConsNil() gopurs_runtime.Value {
	once_showRecordFieldsConsNil.Do(func() {
		showRecordFieldsConsNil = gopurs_runtime.Func2(Call_showRecordFieldsConsNil)
	})
	return showRecordFieldsConsNil
}

func Call_showRecord(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictShowRecordFields_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictShowRecordFields_2 gopurs_runtime.Value = dictShowRecordFields_2_loop
_ = dictShowRecordFields_2
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("{" + gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_2_loop, "showRecordFields"), gopurs_runtime.Constructor0("Proxy"), record_3).StrVal + "}")
}))
}

func Call_showRecordFieldsCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictShowRecordFields_1_loop gopurs_runtime.Value, dictShow_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShowRecordFields_1 gopurs_runtime.Value = dictShowRecordFields_1_loop
_ = dictShowRecordFields_1
var dictShow_2 gopurs_runtime.Value = dictShow_2_loop
_ = dictShow_2
return gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, record_4 gopurs_runtime.Value) gopurs_runtime.Value {
key_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0_loop, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_5_0
return gopurs_runtime.Str(" " + key_5_0.StrVal + ": " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_2_loop, "show"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_5_0, record_4)).StrVal + "," + gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_1_loop, "showRecordFields"), gopurs_runtime.Constructor0("Proxy"), record_4).StrVal)
}))
}

func Call_showRecordFieldsConsNil(dictIsSymbol_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
return gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, record_3 gopurs_runtime.Value) gopurs_runtime.Value {
key_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0_loop, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_4_0
return gopurs_runtime.Str(" " + key_4_0.StrVal + ": " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1_loop, "show"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_4_0, record_3)).StrVal + " ")
}))
}

func Get_showArrayImpl() gopurs_runtime.Value {
	return _Gopurs_ShowArrayImpl
}

func Get_showCharImpl() gopurs_runtime.Value {
	return _Gopurs_ShowCharImpl
}

func Get_showIntImpl() gopurs_runtime.Value {
	return _Gopurs_ShowIntImpl
}

func Get_showNumberImpl() gopurs_runtime.Value {
	return _Gopurs_ShowNumberImpl
}

func Get_showStringImpl() gopurs_runtime.Value {
	return _Gopurs_ShowStringImpl
}
