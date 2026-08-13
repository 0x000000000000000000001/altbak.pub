package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Show_ShowRecordFields_dollarDict gopurs_runtime.Value
var once_Data_Show_ShowRecordFields_dollarDict sync.Once
func Get_Data_Show_ShowRecordFields_dollarDict() gopurs_runtime.Value {
	once_Data_Show_ShowRecordFields_dollarDict.Do(func() {
		cache_Data_Show_ShowRecordFields_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_ShowRecordFields_dollarDict(x_0_box)
})
	})
	return cache_Data_Show_ShowRecordFields_dollarDict
}

var cache_Data_Show_Show_dollarDict gopurs_runtime.Value
var once_Data_Show_Show_dollarDict sync.Once
func Get_Data_Show_Show_dollarDict() gopurs_runtime.Value {
	once_Data_Show_Show_dollarDict.Do(func() {
		cache_Data_Show_Show_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_Show_dollarDict(x_0_box)
})
	})
	return cache_Data_Show_Show_dollarDict
}

var cache_Data_Show_showVoid gopurs_runtime.Value
var once_Data_Show_showVoid sync.Once
func Get_Data_Show_showVoid() gopurs_runtime.Value {
	once_Data_Show_showVoid.Do(func() {
		cache_Data_Show_showVoid = gopurs_runtime.RecordDict1("show", Get_Data_Void_absurd())
	})
	return cache_Data_Show_showVoid
}

var cache_Data_Show_showUnit gopurs_runtime.Value
var once_Data_Show_showUnit sync.Once
func Get_Data_Show_showUnit() gopurs_runtime.Value {
	once_Data_Show_showUnit.Do(func() {
		cache_Data_Show_showUnit = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unit")
}))
	})
	return cache_Data_Show_showUnit
}

var cache_Data_Show_showString gopurs_runtime.Value
var once_Data_Show_showString sync.Once
func Get_Data_Show_showString() gopurs_runtime.Value {
	once_Data_Show_showString.Do(func() {
		cache_Data_Show_showString = gopurs_runtime.RecordDict1("show", Get_Data_Show_showStringImpl())
	})
	return cache_Data_Show_showString
}

var cache_Data_Show_showRecordFieldsNil gopurs_runtime.Value
var once_Data_Show_showRecordFieldsNil sync.Once
func Get_Data_Show_showRecordFieldsNil() gopurs_runtime.Value {
	once_Data_Show_showRecordFieldsNil.Do(func() {
		cache_Data_Show_showRecordFieldsNil = gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
})
}))
	})
	return cache_Data_Show_showRecordFieldsNil
}

var cache_Data_Show_showRecordFields gopurs_runtime.Value
var once_Data_Show_showRecordFields sync.Once
func Get_Data_Show_showRecordFields() gopurs_runtime.Value {
	once_Data_Show_showRecordFields.Do(func() {
		cache_Data_Show_showRecordFields = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_showRecordFields(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_showRecordFields
}

var cache_Data_Show_showRecord gopurs_runtime.Value
var once_Data_Show_showRecord sync.Once
func Get_Data_Show_showRecord() gopurs_runtime.Value {
	once_Data_Show_showRecord.Do(func() {
		cache_Data_Show_showRecord = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictShowRecordFields_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_showRecord(_dollar__unused_0_box, _dollar__unused_1_box, dictShowRecordFields_2_box)
})
	})
	return cache_Data_Show_showRecord
}

var cache_Data_Show_showProxy gopurs_runtime.Value
var once_Data_Show_showProxy sync.Once
func Get_Data_Show_showProxy() gopurs_runtime.Value {
	once_Data_Show_showProxy.Do(func() {
		cache_Data_Show_showProxy = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("Proxy")
}))
	})
	return cache_Data_Show_showProxy
}

var cache_Data_Show_showNumber gopurs_runtime.Value
var once_Data_Show_showNumber sync.Once
func Get_Data_Show_showNumber() gopurs_runtime.Value {
	once_Data_Show_showNumber.Do(func() {
		cache_Data_Show_showNumber = gopurs_runtime.RecordDict1("show", Get_Data_Show_showNumberImpl())
	})
	return cache_Data_Show_showNumber
}

var cache_Data_Show_showInt gopurs_runtime.Value
var once_Data_Show_showInt sync.Once
func Get_Data_Show_showInt() gopurs_runtime.Value {
	once_Data_Show_showInt.Do(func() {
		cache_Data_Show_showInt = gopurs_runtime.RecordDict1("show", Get_Data_Show_showIntImpl())
	})
	return cache_Data_Show_showInt
}

var cache_Data_Show_showChar gopurs_runtime.Value
var once_Data_Show_showChar sync.Once
func Get_Data_Show_showChar() gopurs_runtime.Value {
	once_Data_Show_showChar.Do(func() {
		cache_Data_Show_showChar = gopurs_runtime.RecordDict1("show", Get_Data_Show_showCharImpl())
	})
	return cache_Data_Show_showChar
}

var cache_Data_Show_showBoolean gopurs_runtime.Value
var once_Data_Show_showBoolean sync.Once
func Get_Data_Show_showBoolean() gopurs_runtime.Value {
	once_Data_Show_showBoolean.Do(func() {
		cache_Data_Show_showBoolean = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_0.IntVal) != (0) {
__t0 = "true"
goto end_branch_0
} else {

}
}
{
__t0 = "false"
}
end_branch_0:
return gopurs_runtime.Str(__t0)
}))
	})
	return cache_Data_Show_showBoolean
}

var cache_Data_Show_show gopurs_runtime.Value
var once_Data_Show_show sync.Once
func Get_Data_Show_show() gopurs_runtime.Value {
	once_Data_Show_show.Do(func() {
		cache_Data_Show_show = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show
}

var cache_Data_Show_showArray gopurs_runtime.Value
var once_Data_Show_showArray sync.Once
func Get_Data_Show_showArray() gopurs_runtime.Value {
	once_Data_Show_showArray.Do(func() {
		cache_Data_Show_showArray = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_showArray(dictShow_0_box)
})
	})
	return cache_Data_Show_showArray
}

var cache_Data_Show_showRecordFieldsCons gopurs_runtime.Value
var once_Data_Show_showRecordFieldsCons sync.Once
func Get_Data_Show_showRecordFieldsCons() gopurs_runtime.Value {
	once_Data_Show_showRecordFieldsCons.Do(func() {
		cache_Data_Show_showRecordFieldsCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, dictShowRecordFields_1_box gopurs_runtime.Value, dictShow_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_showRecordFieldsCons(dictIsSymbol_0_box, dictShowRecordFields_1_box, dictShow_2_box)
})
	})
	return cache_Data_Show_showRecordFieldsCons
}

var cache_Data_Show_showRecordFieldsConsNil gopurs_runtime.Value
var once_Data_Show_showRecordFieldsConsNil sync.Once
func Get_Data_Show_showRecordFieldsConsNil() gopurs_runtime.Value {
	once_Data_Show_showRecordFieldsConsNil.Do(func() {
		cache_Data_Show_showRecordFieldsConsNil = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_showRecordFieldsConsNil(dictIsSymbol_0_box, dictShow_1_box)
})
	})
	return cache_Data_Show_showRecordFieldsConsNil
}

var cache_Data_Show_show__857859206 gopurs_runtime.Value
var once_Data_Show_show__857859206 sync.Once
func Get_Data_Show_show__857859206() gopurs_runtime.Value {
	once_Data_Show_show__857859206.Do(func() {
		cache_Data_Show_show__857859206 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__857859206(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](dict_0_box))
})
	})
	return cache_Data_Show_show__857859206
}

var cache_Data_Show_show__587512715 gopurs_runtime.Value
var once_Data_Show_show__587512715 sync.Once
func Get_Data_Show_show__587512715() gopurs_runtime.Value {
	once_Data_Show_show__587512715.Do(func() {
		cache_Data_Show_show__587512715 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__587512715(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Date_Date]](dict_0_box))
})
	})
	return cache_Data_Show_show__587512715
}

var cache_Data_Show_show__1306042987 gopurs_runtime.Value
var once_Data_Show_show__1306042987 sync.Once
func Get_Data_Show_show__1306042987() gopurs_runtime.Value {
	once_Data_Show_show__1306042987.Do(func() {
		cache_Data_Show_show__1306042987 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__1306042987(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Time_Time]](dict_0_box))
})
	})
	return cache_Data_Show_show__1306042987
}

var cache_Data_Show_show__1488465650 gopurs_runtime.Value
var once_Data_Show_show__1488465650 sync.Once
func Get_Data_Show_show__1488465650() gopurs_runtime.Value {
	once_Data_Show_show__1488465650.Do(func() {
		cache_Data_Show_show__1488465650 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__1488465650(__eta0_0_box)
})
	})
	return cache_Data_Show_show__1488465650
}

var cache_Data_Show_show__3380206610 gopurs_runtime.Value
var once_Data_Show_show__3380206610 sync.Once
func Get_Data_Show_show__3380206610() gopurs_runtime.Value {
	once_Data_Show_show__3380206610.Do(func() {
		cache_Data_Show_show__3380206610 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__3380206610(__eta0_0_box)
})
	})
	return cache_Data_Show_show__3380206610
}

var cache_Data_Show_show__3756561682 gopurs_runtime.Value
var once_Data_Show_show__3756561682 sync.Once
func Get_Data_Show_show__3756561682() gopurs_runtime.Value {
	once_Data_Show_show__3756561682.Do(func() {
		cache_Data_Show_show__3756561682 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__3756561682(__eta0_0_box)
})
	})
	return cache_Data_Show_show__3756561682
}

var cache_Data_Show_show__2742601362 gopurs_runtime.Value
var once_Data_Show_show__2742601362 sync.Once
func Get_Data_Show_show__2742601362() gopurs_runtime.Value {
	once_Data_Show_show__2742601362.Do(func() {
		cache_Data_Show_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__2742601362(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show__2742601362
}

var cache_Data_Show_show__4031350738 gopurs_runtime.Value
var once_Data_Show_show__4031350738 sync.Once
func Get_Data_Show_show__4031350738() gopurs_runtime.Value {
	once_Data_Show_show__4031350738.Do(func() {
		cache_Data_Show_show__4031350738 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__4031350738(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show__4031350738
}

var cache_Data_Show_show__2696961938 gopurs_runtime.Value
var once_Data_Show_show__2696961938 sync.Once
func Get_Data_Show_show__2696961938() gopurs_runtime.Value {
	once_Data_Show_show__2696961938.Do(func() {
		cache_Data_Show_show__2696961938 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__2696961938(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_Data_Show_show__2696961938
}

var cache_Data_Show_show__1626410898 gopurs_runtime.Value
var once_Data_Show_show__1626410898 sync.Once
func Get_Data_Show_show__1626410898() gopurs_runtime.Value {
	once_Data_Show_show__1626410898.Do(func() {
		cache_Data_Show_show__1626410898 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Show_show__1626410898(uint32(v_0_box.IntVal)))
})
	})
	return cache_Data_Show_show__1626410898
}

var cache_Data_Show_show__1723386194 gopurs_runtime.Value
var once_Data_Show_show__1723386194 sync.Once
func Get_Data_Show_show__1723386194() gopurs_runtime.Value {
	once_Data_Show_show__1723386194.Do(func() {
		cache_Data_Show_show__1723386194 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Show_show__1723386194(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box)))
})
	})
	return cache_Data_Show_show__1723386194
}

var cache_Data_Show_show__1261750354 gopurs_runtime.Value
var once_Data_Show_show__1261750354 sync.Once
func Get_Data_Show_show__1261750354() gopurs_runtime.Value {
	once_Data_Show_show__1261750354.Do(func() {
		cache_Data_Show_show__1261750354 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Show_show__1261750354(uint32(v_0_box.IntVal)))
})
	})
	return cache_Data_Show_show__1261750354
}

var cache_Data_Show_show__2929403666 gopurs_runtime.Value
var once_Data_Show_show__2929403666 sync.Once
func Get_Data_Show_show__2929403666() gopurs_runtime.Value {
	once_Data_Show_show__2929403666.Do(func() {
		cache_Data_Show_show__2929403666 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__2929403666(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show__2929403666
}

var cache_Data_Show_show__3316320786 gopurs_runtime.Value
var once_Data_Show_show__3316320786 sync.Once
func Get_Data_Show_show__3316320786() gopurs_runtime.Value {
	once_Data_Show_show__3316320786.Do(func() {
		cache_Data_Show_show__3316320786 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__3316320786(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show__3316320786
}

var cache_Data_Show_show__1092279890 gopurs_runtime.Value
var once_Data_Show_show__1092279890 sync.Once
func Get_Data_Show_show__1092279890() gopurs_runtime.Value {
	once_Data_Show_show__1092279890.Do(func() {
		cache_Data_Show_show__1092279890 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__1092279890(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show__1092279890
}

var cache_Data_Show_show__2130238610 gopurs_runtime.Value
var once_Data_Show_show__2130238610 sync.Once
func Get_Data_Show_show__2130238610() gopurs_runtime.Value {
	once_Data_Show_show__2130238610.Do(func() {
		cache_Data_Show_show__2130238610 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__2130238610(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_Data_Show_show__2130238610
}

var cache_Data_Show_show__2896747026 gopurs_runtime.Value
var once_Data_Show_show__2896747026 sync.Once
func Get_Data_Show_show__2896747026() gopurs_runtime.Value {
	once_Data_Show_show__2896747026.Do(func() {
		cache_Data_Show_show__2896747026 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__2896747026(__eta0_0_box)
})
	})
	return cache_Data_Show_show__2896747026
}

var cache_Data_Show_show__3698026194 gopurs_runtime.Value
var once_Data_Show_show__3698026194 sync.Once
func Get_Data_Show_show__3698026194() gopurs_runtime.Value {
	once_Data_Show_show__3698026194.Do(func() {
		cache_Data_Show_show__3698026194 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__3698026194(__eta0_0_box)
})
	})
	return cache_Data_Show_show__3698026194
}

var cache_Data_Show_show__2183537362 gopurs_runtime.Value
var once_Data_Show_show__2183537362 sync.Once
func Get_Data_Show_show__2183537362() gopurs_runtime.Value {
	once_Data_Show_show__2183537362.Do(func() {
		cache_Data_Show_show__2183537362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__2183537362(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_Data_Show_show__2183537362
}

var cache_Data_Show_show__842947602 gopurs_runtime.Value
var once_Data_Show_show__842947602 sync.Once
func Get_Data_Show_show__842947602() gopurs_runtime.Value {
	once_Data_Show_show__842947602.Do(func() {
		cache_Data_Show_show__842947602 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__842947602(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_Data_Show_show__842947602
}

var cache_Data_Show_show__1073032466 gopurs_runtime.Value
var once_Data_Show_show__1073032466 sync.Once
func Get_Data_Show_show__1073032466() gopurs_runtime.Value {
	once_Data_Show_show__1073032466.Do(func() {
		cache_Data_Show_show__1073032466 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Show_show__1073032466(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_0_box)))
})
	})
	return cache_Data_Show_show__1073032466
}

var cache_Data_Show_show__3978978930 gopurs_runtime.Value
var once_Data_Show_show__3978978930 sync.Once
func Get_Data_Show_show__3978978930() gopurs_runtime.Value {
	once_Data_Show_show__3978978930.Do(func() {
		cache_Data_Show_show__3978978930 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__3978978930(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show__3978978930
}

var cache_Data_Show_show__255526802 gopurs_runtime.Value
var once_Data_Show_show__255526802 sync.Once
func Get_Data_Show_show__255526802() gopurs_runtime.Value {
	once_Data_Show_show__255526802.Do(func() {
		cache_Data_Show_show__255526802 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__255526802(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show__255526802
}

var cache_Data_Show_show__3448840338 gopurs_runtime.Value
var once_Data_Show_show__3448840338 sync.Once
func Get_Data_Show_show__3448840338() gopurs_runtime.Value {
	once_Data_Show_show__3448840338.Do(func() {
		cache_Data_Show_show__3448840338 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__3448840338(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_show__3448840338
}

var cache_Data_Show_showRecordFields__3450865987 gopurs_runtime.Value
var once_Data_Show_showRecordFields__3450865987 sync.Once
func Get_Data_Show_showRecordFields__3450865987() gopurs_runtime.Value {
	once_Data_Show_showRecordFields__3450865987.Do(func() {
		cache_Data_Show_showRecordFields__3450865987 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_showRecordFields__3450865987(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_showRecordFields__3450865987
}

var cache_Data_Show_showRecordFields__2713688005 gopurs_runtime.Value
var once_Data_Show_showRecordFields__2713688005 sync.Once
func Get_Data_Show_showRecordFields__2713688005() gopurs_runtime.Value {
	once_Data_Show_showRecordFields__2713688005.Do(func() {
		cache_Data_Show_showRecordFields__2713688005 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_showRecordFields__2713688005(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Show_showRecordFields__2713688005
}

type Constructor_Data_Show_ShowRecordFields[T_rowlist any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2498393510] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "showRecordFields": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Show_ShowRecordFields: " + key)
		}
	}
}


type Constructor_Data_Show_Show[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1835580986] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Show_Show[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "show": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Show_Show: " + key)
		}
	}
}


func Call_Data_Show_ShowRecordFields_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Show_Show_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Show_showRecordFields(dict_0_loop *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_showRecord(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictShowRecordFields_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictShowRecordFields_2 gopurs_runtime.Value = dictShowRecordFields_2_loop
_ = dictShowRecordFields_2
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("{") + (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_2, "showRecordFields"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, record_3).StrVal())) + ("}"))
}))
}

func Call_Data_Show_show(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_showArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show")))
}

func Call_Data_Show_showRecordFieldsCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictShowRecordFields_1_loop gopurs_runtime.Value, dictShow_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShowRecordFields_1 gopurs_runtime.Value = dictShowRecordFields_1_loop
_ = dictShowRecordFields_1
var dictShow_2 gopurs_runtime.Value = dictShow_2_loop
_ = dictShow_2
return gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(record_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_5_0 -> gopurs_runtime.Value
key_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_5_0
return gopurs_runtime.Str((((((" ") + (key_5_0.StrVal())) + (": ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_2, "show"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_5_0.StrVal()), record_4)).StrVal())) + (",")) + (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_1, "showRecordFields"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, record_4).StrVal()))
})
}))
}

func Call_Data_Show_showRecordFieldsConsNil(dictIsSymbol_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
return gopurs_runtime.RecordDict1("showRecordFields", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_4_0 -> gopurs_runtime.Value
key_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_4_0
return gopurs_runtime.Str(((((" ") + (key_4_0.StrVal())) + (": ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_4_0.StrVal()), record_3)).StrVal())) + (" "))
})
}))
}

func Call_Data_Show_show__857859206(dict_0_loop *Constructor_Data_Show_Show[int64]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[int64] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__587512715(dict_0_loop *Constructor_Data_Show_Show[*Constructor_Data_Date_Date]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[*Constructor_Data_Date_Date] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__1306042987(dict_0_loop *Constructor_Data_Show_Show[*Constructor_Data_Time_Time]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[*Constructor_Data_Time_Time] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__1488465650(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), __eta0_0).StrVal())
}

func Call_Data_Show_show__3380206610(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), __eta0_0).StrVal())
}

func Call_Data_Show_show__3756561682(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), __eta0_0).StrVal())
}

func Call_Data_Show_show__2742601362(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__4031350738(dict_0_loop *Constructor_Data_Show_Show[[]gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[[]gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__2696961938(dict_0_loop *Constructor_Data_Show_Show[[]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[[]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__1626410898(v_0_loop uint32) string {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == 1908470532) {
__t0 = gopurs_runtime.Str("January")
goto end_branch_0
} else {

}
}
{
if (v_0 == 2455627378) {
__t0 = gopurs_runtime.Str("February")
goto end_branch_0
} else {

}
}
{
if (v_0 == 4162469099) {
__t0 = gopurs_runtime.Str("March")
goto end_branch_0
} else {

}
}
{
if (v_0 == 1692989816) {
__t0 = gopurs_runtime.Str("April")
goto end_branch_0
} else {

}
}
{
if (v_0 == 330658827) {
__t0 = gopurs_runtime.Str("May")
goto end_branch_0
} else {

}
}
{
if (v_0 == 4067355978) {
__t0 = gopurs_runtime.Str("June")
goto end_branch_0
} else {

}
}
{
if (v_0 == 2276710548) {
__t0 = gopurs_runtime.Str("July")
goto end_branch_0
} else {

}
}
{
if (v_0 == 243771071) {
__t0 = gopurs_runtime.Str("August")
goto end_branch_0
} else {

}
}
{
if (v_0 == 215731793) {
__t0 = gopurs_runtime.Str("September")
goto end_branch_0
} else {

}
}
{
if (v_0 == 8639228) {
__t0 = gopurs_runtime.Str("October")
goto end_branch_0
} else {

}
}
{
if (v_0 == 49471444) {
__t0 = gopurs_runtime.Str("November")
goto end_branch_0
} else {

}
}
{
if (v_0 == 3889233761) {
__t0 = gopurs_runtime.Str("December")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.StrVal()
}

func Call_Data_Show_show__1723386194(v_0_loop *Constructor_Data_Date_Date) string {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str("(Date "), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_showYear(), "show"), gopurs_runtime.Int((*Constructor_Data_Date_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_showMonth(), "show"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1), UnsafePtr: nil}).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_showDay(), "show"), gopurs_runtime.Int((*Constructor_Data_Date_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2)).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal()
}

func Call_Data_Show_show__1261750354(v_0_loop uint32) string {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == 217821258) {
__t0 = gopurs_runtime.Str("Minute")
goto end_branch_0
} else {

}
}
{
if (v_0 == 3908053364) {
__t0 = gopurs_runtime.Str("Second")
goto end_branch_0
} else {

}
}
{
if (v_0 == 1292308612) {
__t0 = gopurs_runtime.Str("Hour")
goto end_branch_0
} else {

}
}
{
if (v_0 == 2311060696) {
__t0 = gopurs_runtime.Str("Day")
goto end_branch_0
} else {

}
}
{
if (v_0 == 401302776) {
__t0 = gopurs_runtime.Str("Week")
goto end_branch_0
} else {

}
}
{
if (v_0 == 3327533908) {
__t0 = gopurs_runtime.Str("Month")
goto end_branch_0
} else {

}
}
{
if (v_0 == 3631736139) {
__t0 = gopurs_runtime.Str("Year")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.StrVal()
}

func Call_Data_Show_show__2929403666(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__3316320786(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__1092279890(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__2130238610(dict_0_loop *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__2896747026(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_showMap()).V0), __eta0_0)
}

func Call_Data_Show_show__3698026194(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_showMaybe()).V0), __eta0_0)
}

func Call_Data_Show_show__2183537362(dict_0_loop *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__842947602(dict_0_loop *Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__1073032466(v_0_loop *Constructor_Data_Time_Time) string {
var v_0 *Constructor_Data_Time_Time = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str("(Time "), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Time_Component_showHour(), "show"), gopurs_runtime.Int((*Constructor_Data_Time_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Time_Component_showMinute(), "show"), gopurs_runtime.Int((*Constructor_Data_Time_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Time_Component_showSecond(), "show"), gopurs_runtime.Int((*Constructor_Data_Time_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupString(), "append"), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Time_Component_showMillisecond(), "show"), gopurs_runtime.Int((*Constructor_Data_Time_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal()
}

func Call_Data_Show_show__3978978930(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__255526802(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__3448840338(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_showRecordFields__3450865987(dict_0_loop *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_showRecordFields__2713688005(dict_0_loop *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Get_Data_Show_showArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Show_ShowArrayImpl
}

func Get_Data_Show_showCharImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Show_ShowCharImpl
}

func Get_Data_Show_showIntImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Show_ShowIntImpl
}

func Get_Data_Show_showNumberImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Show_ShowNumberImpl
}

func Get_Data_Show_showStringImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Show_ShowStringImpl
}
