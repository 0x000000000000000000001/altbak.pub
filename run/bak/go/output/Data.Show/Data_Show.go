package Data_Show

import (
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Void "gopurs/output/Data.Void"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_showVoid gopurs_runtime.Value
var once_showVoid sync.Once
func Get_showVoid() gopurs_runtime.Value {
	once_showVoid.Do(func() {
		cache_showVoid = gopurs_runtime.RecordDict1("show", pkg_Data_Void.Get_absurd__gopurs_runtime_Value_1771830288())
	})
	return cache_showVoid
}

var cache_showUnit gopurs_runtime.Value
var once_showUnit sync.Once
func Get_showUnit() gopurs_runtime.Value {
	once_showUnit.Do(func() {
		cache_showUnit = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unit")
}))
	})
	return cache_showUnit
}

var cache_showString gopurs_runtime.Value
var once_showString sync.Once
func Get_showString() gopurs_runtime.Value {
	once_showString.Do(func() {
		cache_showString = gopurs_runtime.RecordDict1("show", Get_showStringImpl())
	})
	return cache_showString
}

var cache_showRecordFieldsNil gopurs_runtime.Value
var once_showRecordFieldsNil sync.Once
func Get_showRecordFieldsNil() gopurs_runtime.Value {
	once_showRecordFieldsNil.Do(func() {
		cache_showRecordFieldsNil = gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
})
}))
	})
	return cache_showRecordFieldsNil
}

var cache_showRecordFields gopurs_runtime.Value
var once_showRecordFields sync.Once
func Get_showRecordFields() gopurs_runtime.Value {
	once_showRecordFields.Do(func() {
		cache_showRecordFields = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showRecordFields(gopurs_runtime.CoerceToStruct[Constructor_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_showRecordFields
}

var cache_showRecordFields__gopurs_runtime_Value_2713688005 gopurs_runtime.Value
var once_showRecordFields__gopurs_runtime_Value_2713688005 sync.Once
func Get_showRecordFields__gopurs_runtime_Value_2713688005() gopurs_runtime.Value {
	once_showRecordFields__gopurs_runtime_Value_2713688005.Do(func() {
		cache_showRecordFields__gopurs_runtime_Value_2713688005 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showRecordFields__gopurs_runtime_Value_2713688005(gopurs_runtime.CoerceToStruct[Constructor_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_showRecordFields__gopurs_runtime_Value_2713688005
}

var cache_showRecord gopurs_runtime.Value
var once_showRecord sync.Once
func Get_showRecord() gopurs_runtime.Value {
	once_showRecord.Do(func() {
		cache_showRecord = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictShowRecordFields_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showRecord(_dollar__unused_0_box, _dollar__unused_1_box, dictShowRecordFields_2_box)
})
	})
	return cache_showRecord
}

var cache_showProxy gopurs_runtime.Value
var once_showProxy sync.Once
func Get_showProxy() gopurs_runtime.Value {
	once_showProxy.Do(func() {
		cache_showProxy = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("Proxy")
}))
	})
	return cache_showProxy
}

var cache_showNumber gopurs_runtime.Value
var once_showNumber sync.Once
func Get_showNumber() gopurs_runtime.Value {
	once_showNumber.Do(func() {
		cache_showNumber = gopurs_runtime.RecordDict1("show", Get_showNumberImpl())
	})
	return cache_showNumber
}

var cache_showInt gopurs_runtime.Value
var once_showInt sync.Once
func Get_showInt() gopurs_runtime.Value {
	once_showInt.Do(func() {
		cache_showInt = gopurs_runtime.RecordDict1("show", Get_showIntImpl())
	})
	return cache_showInt
}

var cache_showChar gopurs_runtime.Value
var once_showChar sync.Once
func Get_showChar() gopurs_runtime.Value {
	once_showChar.Do(func() {
		cache_showChar = gopurs_runtime.RecordDict1("show", Get_showCharImpl())
	})
	return cache_showChar
}

var cache_showBoolean gopurs_runtime.Value
var once_showBoolean sync.Once
func Get_showBoolean() gopurs_runtime.Value {
	once_showBoolean.Do(func() {
		cache_showBoolean = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) != (0) {
__t0 = gopurs_runtime.Str("true")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Str("false")
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
	})
	return cache_showBoolean
}

var cache_show gopurs_runtime.Value
var once_show sync.Once
func Get_show() gopurs_runtime.Value {
	once_show.Do(func() {
		cache_show = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show(gopurs_runtime.CoerceToStruct[Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show
}

var cache_show__gopurs_runtime_Value_2742601362 gopurs_runtime.Value
var once_show__gopurs_runtime_Value_2742601362 sync.Once
func Get_show__gopurs_runtime_Value_2742601362() gopurs_runtime.Value {
	once_show__gopurs_runtime_Value_2742601362.Do(func() {
		cache_show__gopurs_runtime_Value_2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__gopurs_runtime_Value_2742601362(gopurs_runtime.CoerceToStruct[Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__gopurs_runtime_Value_2742601362
}

var cache_showArray gopurs_runtime.Value
var once_showArray sync.Once
func Get_showArray() gopurs_runtime.Value {
	once_showArray.Do(func() {
		cache_showArray = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showArray(dictShow_0_box)
})
	})
	return cache_showArray
}

var cache_showRecordFieldsCons gopurs_runtime.Value
var once_showRecordFieldsCons sync.Once
func Get_showRecordFieldsCons() gopurs_runtime.Value {
	once_showRecordFieldsCons.Do(func() {
		cache_showRecordFieldsCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, dictShowRecordFields_1_box gopurs_runtime.Value, dictShow_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showRecordFieldsCons(dictIsSymbol_0_box, dictShowRecordFields_1_box, dictShow_2_box)
})
	})
	return cache_showRecordFieldsCons
}

var cache_showRecordFieldsConsNil gopurs_runtime.Value
var once_showRecordFieldsConsNil sync.Once
func Get_showRecordFieldsConsNil() gopurs_runtime.Value {
	once_showRecordFieldsConsNil.Do(func() {
		cache_showRecordFieldsConsNil = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showRecordFieldsConsNil(dictIsSymbol_0_box, dictShow_1_box)
})
	})
	return cache_showRecordFieldsConsNil
}

type Constructor_ShowRecordFields[T_rowlist any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2498393510] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "showRecordFields": return c.V0
		default: panic("Key not found in dictionary Constructor_ShowRecordFields: " + key)
		}
	}
}


type Constructor_Show[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1835580986] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Show[gopurs_runtime.Value])(ptr)
		switch key {
		case "show": return c.V0
		default: panic("Key not found in dictionary Constructor_Show: " + key)
		}
	}
}


func Call_showRecordFields(dict_0_loop *Constructor_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_showRecordFields__gopurs_runtime_Value_2713688005(dict_0_loop *Constructor_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_showRecord(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictShowRecordFields_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictShowRecordFields_2 gopurs_runtime.Value = dictShowRecordFields_2_loop
_ = dictShowRecordFields_2
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("{"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_2, "showRecordFields"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, record_3), gopurs_runtime.Str("}"))).StrVal())
}))
}

func Call_show(dict_0_loop *Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__gopurs_runtime_Value_2742601362(dict_0_loop *Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_showArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Apply(Get_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show")))
}

func Call_showRecordFieldsCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictShowRecordFields_1_loop gopurs_runtime.Value, dictShow_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShowRecordFields_1 gopurs_runtime.Value = dictShowRecordFields_1_loop
_ = dictShowRecordFields_1
var dictShow_2 gopurs_runtime.Value = dictShow_2_loop
_ = dictShow_2
return gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(record_4 gopurs_runtime.Value) gopurs_runtime.Value {
key_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)})
_ = key_5_0
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), key_5_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(": "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_2, "show"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_5_0, record_4)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(","), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_1, "showRecordFields"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, record_4)))))).StrVal())
})
}))
}

func Call_showRecordFieldsConsNil(dictIsSymbol_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
return gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
key_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)})
_ = key_4_0
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), key_4_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(": "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_4_0, record_3)), gopurs_runtime.Str(" "))))).StrVal())
})
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
