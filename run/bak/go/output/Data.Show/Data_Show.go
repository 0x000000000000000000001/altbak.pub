package Data_Show

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Void "gopurs/output/Data.Void"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var cache_showVoid gopurs_runtime.Value
var once_showVoid sync.Once
func Get_showVoid() gopurs_runtime.Value {
	once_showVoid.Do(func() {
		cache_showVoid = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", pkg_Data_Void.Get_absurd())))
	})
	return cache_showVoid
}

var cache_showUnit gopurs_runtime.Value
var once_showUnit sync.Once
func Get_showUnit() gopurs_runtime.Value {
	once_showUnit.Do(func() {
		cache_showUnit = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unit")
}))))
	})
	return cache_showUnit
}

var cache_showString gopurs_runtime.Value
var once_showString sync.Once
func Get_showString() gopurs_runtime.Value {
	once_showString.Do(func() {
		cache_showString = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", Get_showStringImpl())))
	})
	return cache_showString
}

var cache_showRecordFieldsNil gopurs_runtime.Value
var once_showRecordFieldsNil sync.Once
func Get_showRecordFieldsNil() gopurs_runtime.Value {
	once_showRecordFieldsNil.Do(func() {
		cache_showRecordFieldsNil = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
}))))
	})
	return cache_showRecordFieldsNil
}

var cache_showRecordFields gopurs_runtime.Value
var once_showRecordFields sync.Once
func Get_showRecordFields() gopurs_runtime.Value {
	once_showRecordFields.Do(func() {
		cache_showRecordFields = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showRecordFields(dict_0_box)
})
	})
	return cache_showRecordFields
}

var cache_showRecord gopurs_runtime.Value
var once_showRecord sync.Once
func Get_showRecord() gopurs_runtime.Value {
	once_showRecord.Do(func() {
		cache_showRecord = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictShowRecordFields_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_showRecord(_dollar__unused_0_box, _dollar__unused_1_box, dictShowRecordFields_2_box))
})
	})
	return cache_showRecord
}

var cache_showProxy gopurs_runtime.Value
var once_showProxy sync.Once
func Get_showProxy() gopurs_runtime.Value {
	once_showProxy.Do(func() {
		cache_showProxy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("Proxy")
}))))
	})
	return cache_showProxy
}

var cache_showNumber gopurs_runtime.Value
var once_showNumber sync.Once
func Get_showNumber() gopurs_runtime.Value {
	once_showNumber.Do(func() {
		cache_showNumber = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", Get_showNumberImpl())))
	})
	return cache_showNumber
}

var cache_showInt gopurs_runtime.Value
var once_showInt sync.Once
func Get_showInt() gopurs_runtime.Value {
	once_showInt.Do(func() {
		cache_showInt = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", Get_showIntImpl())))
	})
	return cache_showInt
}

var cache_showChar gopurs_runtime.Value
var once_showChar sync.Once
func Get_showChar() gopurs_runtime.Value {
	once_showChar.Do(func() {
		cache_showChar = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", Get_showCharImpl())))
	})
	return cache_showChar
}

var cache_showBoolean gopurs_runtime.Value
var once_showBoolean sync.Once
func Get_showBoolean() gopurs_runtime.Value {
	once_showBoolean.Do(func() {
		cache_showBoolean = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
return __t0
}))))
	})
	return cache_showBoolean
}

var cache_show gopurs_runtime.Value
var once_show sync.Once
func Get_show() gopurs_runtime.Value {
	once_show.Do(func() {
		cache_show = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show(dict_0_box)
})
	})
	return cache_show
}

var cache_show__func_gopurs_runtime_Value__interface____string_2425962676 gopurs_runtime.Value
var once_show__func_gopurs_runtime_Value__interface____string_2425962676 sync.Once
func Get_show__func_gopurs_runtime_Value__interface____string_2425962676() gopurs_runtime.Value {
	once_show__func_gopurs_runtime_Value__interface____string_2425962676.Do(func() {
		cache_show__func_gopurs_runtime_Value__interface____string_2425962676 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__func_gopurs_runtime_Value__interface____string_2425962676(dict_0_box)
})
	})
	return cache_show__func_gopurs_runtime_Value__interface____string_2425962676
}

var cache_showArray gopurs_runtime.Value
var once_showArray sync.Once
func Get_showArray() gopurs_runtime.Value {
	once_showArray.Do(func() {
		cache_showArray = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_showArray(dictShow_0_box))
})
	})
	return cache_showArray
}

var cache_showRecordFieldsCons gopurs_runtime.Value
var once_showRecordFieldsCons sync.Once
func Get_showRecordFieldsCons() gopurs_runtime.Value {
	once_showRecordFieldsCons.Do(func() {
		cache_showRecordFieldsCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, dictShowRecordFields_1_box gopurs_runtime.Value, dictShow_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_showRecordFieldsCons(dictIsSymbol_0_box, dictShowRecordFields_1_box, dictShow_2_box))
})
	})
	return cache_showRecordFieldsCons
}

var cache_showRecordFieldsConsNil gopurs_runtime.Value
var once_showRecordFieldsConsNil sync.Once
func Get_showRecordFieldsConsNil() gopurs_runtime.Value {
	once_showRecordFieldsConsNil.Do(func() {
		cache_showRecordFieldsConsNil = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_showRecordFieldsConsNil(dictIsSymbol_0_box, dictShow_1_box))
})
	})
	return cache_showRecordFieldsConsNil
}

var cache_showArrayImpl gopurs_runtime.Value
var once_showArrayImpl sync.Once
func Get_showArrayImpl() gopurs_runtime.Value {
	once_showArrayImpl.Do(func() {
		cache_showArrayImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ShowArrayImpl(func(inner_arg0 interface{}) string {
return gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).StrVal()
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_showArrayImpl
}

var cache_showCharImpl gopurs_runtime.Value
var once_showCharImpl sync.Once
func Get_showCharImpl() gopurs_runtime.Value {
	once_showCharImpl.Do(func() {
		cache_showCharImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ShowCharImpl(arg0.StrVal()))
})
	})
	return cache_showCharImpl
}

var cache_showIntImpl gopurs_runtime.Value
var once_showIntImpl sync.Once
func Get_showIntImpl() gopurs_runtime.Value {
	once_showIntImpl.Do(func() {
		cache_showIntImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ShowIntImpl(arg0.IntVal))
})
	})
	return cache_showIntImpl
}

var cache_showNumberImpl gopurs_runtime.Value
var once_showNumberImpl sync.Once
func Get_showNumberImpl() gopurs_runtime.Value {
	once_showNumberImpl.Do(func() {
		cache_showNumberImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ShowNumberImpl(arg0.FloatVal()))
})
	})
	return cache_showNumberImpl
}

var cache_showStringImpl gopurs_runtime.Value
var once_showStringImpl sync.Once
func Get_showStringImpl() gopurs_runtime.Value {
	once_showStringImpl.Do(func() {
		cache_showStringImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ShowStringImpl(arg0.StrVal()))
})
	})
	return cache_showStringImpl
}

func Call_showRecordFields(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "showRecordFields")
}

func Call_showRecord(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictShowRecordFields_2_loop gopurs_runtime.Value) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictShowRecordFields_2 gopurs_runtime.Value = dictShowRecordFields_2_loop
_ = dictShowRecordFields_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("{"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_2, "showRecordFields"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), record_3), gopurs_runtime.Str("}")))
})))
}

func Call_show(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "show")
}

func Call_show__func_gopurs_runtime_Value__interface____string_2425962676(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "show")
}

func Call_showArray(dictShow_0_loop gopurs_runtime.Value) interface{} {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Apply(Get_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))))
}

func Call_showRecordFieldsCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictShowRecordFields_1_loop gopurs_runtime.Value, dictShow_2_loop gopurs_runtime.Value) interface{} {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShowRecordFields_1 gopurs_runtime.Value = dictShowRecordFields_1_loop
_ = dictShowRecordFields_1
var dictShow_2 gopurs_runtime.Value = dictShow_2_loop
_ = dictShow_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, record_4 gopurs_runtime.Value) gopurs_runtime.Value {
key_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = key_5_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), key_5_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(": "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_2, "show"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_5_0, record_4)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(","), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_1, "showRecordFields"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), record_4))))))
})))
}

func Call_showRecordFieldsConsNil(dictIsSymbol_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value) interface{} {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, record_3 gopurs_runtime.Value) gopurs_runtime.Value {
key_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = key_4_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), key_4_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(": "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_4_0, record_3)), gopurs_runtime.Str(" ")))))
})))
}
