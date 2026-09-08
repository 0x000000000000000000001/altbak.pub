package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Show_ShowRecordFields_dollar_Dict gopurs_runtime.Value
var once_Data_Show_ShowRecordFields_dollar_Dict sync.Once
func Get_Data_Show_ShowRecordFields_dollar_Dict() gopurs_runtime.Value {
	once_Data_Show_ShowRecordFields_dollar_Dict.Do(func() {
		cache_Data_Show_ShowRecordFields_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2498393510, UnsafePtr: unsafe.Pointer(Call_Data_Show_ShowRecordFields_dollar_Dict(func() struct{
	showRecordFields gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	showRecordFields gopurs_runtime.Value
}{}
					clone.showRecordFields = gopurs_runtime.RecordGet(orig, "showRecordFields")
					return clone
				}()))}
})
	})
	return cache_Data_Show_ShowRecordFields_dollar_Dict
}

var cache_Data_Show_Show_dollar_Dict gopurs_runtime.Value
var once_Data_Show_Show_dollar_Dict sync.Once
func Get_Data_Show_Show_dollar_Dict() gopurs_runtime.Value {
	once_Data_Show_Show_dollar_Dict.Do(func() {
		cache_Data_Show_Show_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Call_Data_Show_Show_dollar_Dict(func() struct{
	show gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	show gopurs_runtime.Value
}{}
					clone.show = gopurs_runtime.RecordGet(orig, "show")
					return clone
				}()))}
})
	})
	return cache_Data_Show_Show_dollar_Dict
}

var cache_Data_Show_showVoid gopurs_runtime.Value
var once_Data_Show_showVoid sync.Once
func Get_Data_Show_showVoid() gopurs_runtime.Value {
	once_Data_Show_showVoid.Do(func() {
		cache_Data_Show_showVoid = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, Get_Data_Void_absurd()}))}
	})
	return cache_Data_Show_showVoid
}

var cache_Data_Show_showUnit gopurs_runtime.Value
var once_Data_Show_showUnit sync.Once
func Get_Data_Show_showUnit() gopurs_runtime.Value {
	once_Data_Show_showUnit.Do(func() {
		cache_Data_Show_showUnit = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unit")
})}))}
	})
	return cache_Data_Show_showUnit
}

var cache_Data_Show_showString gopurs_runtime.Value
var once_Data_Show_showString sync.Once
func Get_Data_Show_showString() gopurs_runtime.Value {
	once_Data_Show_showString.Do(func() {
		cache_Data_Show_showString = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Show_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, Get_Data_Show_showStringImpl()})))}
	})
	return cache_Data_Show_showString
}

var cache_Data_Show_showRecordFieldsNil gopurs_runtime.Value
var once_Data_Show_showRecordFieldsNil sync.Once
func Get_Data_Show_showRecordFieldsNil() gopurs_runtime.Value {
	once_Data_Show_showRecordFieldsNil.Do(func() {
		cache_Data_Show_showRecordFieldsNil = gopurs_runtime.Value{Type: 9, IntVal: 2498393510, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
})}))}
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
		cache_Data_Show_showRecord = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, dictShowRecordFields_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_showRecord(_dollar___unused_0_box, _dollar___unused_1_box, dictShowRecordFields_2_box)
})
	})
	return cache_Data_Show_showRecord
}

var cache_Data_Show_showProxy gopurs_runtime.Value
var once_Data_Show_showProxy sync.Once
func Get_Data_Show_showProxy() gopurs_runtime.Value {
	once_Data_Show_showProxy.Do(func() {
		cache_Data_Show_showProxy = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Show_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("Proxy")
})})))}
	})
	return cache_Data_Show_showProxy
}

var cache_Data_Show_showNumber gopurs_runtime.Value
var once_Data_Show_showNumber sync.Once
func Get_Data_Show_showNumber() gopurs_runtime.Value {
	once_Data_Show_showNumber.Do(func() {
		cache_Data_Show_showNumber = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Show_3263178038_1386611502((&Constructor_Data_Show_Show[float64]{1, Get_Data_Show_showNumberImpl()})))}
	})
	return cache_Data_Show_showNumber
}

var cache_Data_Show_showInt gopurs_runtime.Value
var once_Data_Show_showInt sync.Once
func Get_Data_Show_showInt() gopurs_runtime.Value {
	once_Data_Show_showInt.Do(func() {
		cache_Data_Show_showInt = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Show_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, Get_Data_Show_showIntImpl()})))}
	})
	return cache_Data_Show_showInt
}

var cache_Data_Show_showChar gopurs_runtime.Value
var once_Data_Show_showChar sync.Once
func Get_Data_Show_showChar() gopurs_runtime.Value {
	once_Data_Show_showChar.Do(func() {
		cache_Data_Show_showChar = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Show_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, Get_Data_Show_showCharImpl()})))}
	})
	return cache_Data_Show_showChar
}

var cache_Data_Show_showBoolean gopurs_runtime.Value
var once_Data_Show_showBoolean sync.Once
func Get_Data_Show_showBoolean() gopurs_runtime.Value {
	once_Data_Show_showBoolean.Do(func() {
		cache_Data_Show_showBoolean = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Show_2735895690_1386611502((&Constructor_Data_Show_Show[bool]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})})))}
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

var cache_Data_Show_show__4098857531 gopurs_runtime.Value
var once_Data_Show_show__4098857531 sync.Once
func Get_Data_Show_show__4098857531() gopurs_runtime.Value {
	once_Data_Show_show__4098857531.Do(func() {
		cache_Data_Show_show__4098857531 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__4098857531(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__4098857531
}

var cache_Data_Show_show__575334575 gopurs_runtime.Value
var once_Data_Show_show__575334575 sync.Once
func Get_Data_Show_show__575334575() gopurs_runtime.Value {
	once_Data_Show_show__575334575.Do(func() {
		cache_Data_Show_show__575334575 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__575334575(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__575334575
}

var cache_Data_Show_show__1490398944 gopurs_runtime.Value
var once_Data_Show_show__1490398944 sync.Once
func Get_Data_Show_show__1490398944() gopurs_runtime.Value {
	once_Data_Show_show__1490398944.Do(func() {
		cache_Data_Show_show__1490398944 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__1490398944(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__1490398944
}

var cache_Data_Show_show__3482260741 gopurs_runtime.Value
var once_Data_Show_show__3482260741 sync.Once
func Get_Data_Show_show__3482260741() gopurs_runtime.Value {
	once_Data_Show_show__3482260741.Do(func() {
		cache_Data_Show_show__3482260741 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__3482260741(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__3482260741
}

var cache_Data_Show_show__3604015523 gopurs_runtime.Value
var once_Data_Show_show__3604015523 sync.Once
func Get_Data_Show_show__3604015523() gopurs_runtime.Value {
	once_Data_Show_show__3604015523.Do(func() {
		cache_Data_Show_show__3604015523 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__3604015523(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__3604015523
}

var cache_Data_Show_show__999399311 gopurs_runtime.Value
var once_Data_Show_show__999399311 sync.Once
func Get_Data_Show_show__999399311() gopurs_runtime.Value {
	once_Data_Show_show__999399311.Do(func() {
		cache_Data_Show_show__999399311 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__999399311(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__999399311
}

var cache_Data_Show_show__2808757794 gopurs_runtime.Value
var once_Data_Show_show__2808757794 sync.Once
func Get_Data_Show_show__2808757794() gopurs_runtime.Value {
	once_Data_Show_show__2808757794.Do(func() {
		cache_Data_Show_show__2808757794 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__2808757794(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__2808757794
}

var cache_Data_Show_show__2850478386 gopurs_runtime.Value
var once_Data_Show_show__2850478386 sync.Once
func Get_Data_Show_show__2850478386() gopurs_runtime.Value {
	once_Data_Show_show__2850478386.Do(func() {
		cache_Data_Show_show__2850478386 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__2850478386(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__2850478386
}

var cache_Data_Show_show__1943266116 gopurs_runtime.Value
var once_Data_Show_show__1943266116 sync.Once
func Get_Data_Show_show__1943266116() gopurs_runtime.Value {
	once_Data_Show_show__1943266116.Do(func() {
		cache_Data_Show_show__1943266116 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Show_show__1943266116(__eta_norm_0_0_box)
})
	})
	return cache_Data_Show_show__1943266116
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

type Constructor_Data_Show_ShowRecordFields[T_rowlist any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2498393510] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Show_ShowRecordFields[any, any])(ptr)
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
		c := (*Constructor_Data_Show_Show[any])(ptr)
		_ = c
		switch key {
		case "show": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Show_Show: " + key)
		}
	}
}


func Call_Data_Show_ShowRecordFields_dollar_Dict(x_0_loop struct{
	showRecordFields gopurs_runtime.Value
}) *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	showRecordFields gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"showRecordFields"}, []gopurs_runtime.Value{orig.showRecordFields})
				}())
}

func Call_Data_Show_Show_dollar_Dict(x_0_loop struct{
	show gopurs_runtime.Value
}) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
var x_0 struct{
	show gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"show"}, []gopurs_runtime.Value{orig.show})
				}())
}

func Call_Data_Show_showRecordFields(dict_0_loop *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_showRecord(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, dictShowRecordFields_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
_ = _dollar___unused_1
var dictShowRecordFields_2 gopurs_runtime.Value = dictShowRecordFields_2_loop
_ = dictShowRecordFields_2
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("{") + (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_2, "showRecordFields"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, record_3).StrVal())) + ("}"))
})}))}
}

func Call_Data_Show_show(dict_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Show_show__4098857531(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__4098857531:
for {
if false { continue show__4098857531 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t12 string
{
var __t_tag_0 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_0) == 1908470532) {
__t12 = "January"
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_1) == 2455627378) {
__t12 = "February"
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_2) == 4162469099) {
__t12 = "March"
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_3) == 1692989816) {
__t12 = "April"
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_4) == 330658827) {
__t12 = "May"
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_5) == 4067355978) {
__t12 = "June"
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_6) == 2276710548) {
__t12 = "July"
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_7) == 243771071) {
__t12 = "August"
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_8) == 215731793) {
__t12 = "September"
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_9) == 8639228) {
__t12 = "October"
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_10) == 49471444) {
__t12 = "November"
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_11) == 3889233761) {
__t12 = "December"
goto end_branch_12
} else {

}
}
{
__t12 = func() string { panic("Failed pattern match") }()
}
end_branch_12:
return gopurs_runtime.Str(__t12)
}
}

func Call_Data_Show_show__575334575(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__575334575:
for {
if false { continue show__575334575 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Date_Date]](Get_Data_Date_showDate()).V0), __eta_norm_0_0).StrVal())
}
}

func Call_Data_Show_show__1490398944(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__1490398944:
for {
if false { continue show__1490398944 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t7 string
{
var __t_tag_0 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_0) == 217821258) {
__t7 = "Minute"
goto end_branch_7
} else {

}
}
{
var __t_tag_1 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_1) == 3908053364) {
__t7 = "Second"
goto end_branch_7
} else {

}
}
{
var __t_tag_2 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_2) == 1292308612) {
__t7 = "Hour"
goto end_branch_7
} else {

}
}
{
var __t_tag_3 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_3) == 2311060696) {
__t7 = "Day"
goto end_branch_7
} else {

}
}
{
var __t_tag_4 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_4) == 401302776) {
__t7 = "Week"
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_5) == 3327533908) {
__t7 = "Month"
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_6) == 3631736139) {
__t7 = "Year"
goto end_branch_7
} else {

}
}
{
__t7 = func() string { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Str(__t7)
}
}

func Call_Data_Show_show__3482260741(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__3482260741:
for {
if false { continue show__3482260741 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_showMap()).V0), __eta_norm_0_0).StrVal())
}
}

func Call_Data_Show_show__3604015523(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__3604015523:
for {
if false { continue show__3604015523 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_showMaybe()).V0), __eta_norm_0_0).StrVal())
}
}

func Call_Data_Show_show__999399311(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__999399311:
for {
if false { continue show__999399311 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Time_Time]](Get_Data_Time_showTime()).V0), __eta_norm_0_0).StrVal())
}
}

func Call_Data_Show_show__2808757794(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__2808757794:
for {
if false { continue show__2808757794 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), __eta_norm_0_0).StrVal())
}
}

func Call_Data_Show_show__2850478386(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__2850478386:
for {
if false { continue show__2850478386 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), __eta_norm_0_0).StrVal())
}
}

func Call_Data_Show_show__1943266116(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
show__1943266116:
for {
if false { continue show__1943266116 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), __eta_norm_0_0).StrVal())
}
}

func Call_Data_Show_showArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Show_1356436936_1386611502((&Constructor_Data_Show_Show[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))})))}
}

func Call_Data_Show_showRecordFieldsCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictShowRecordFields_1_loop gopurs_runtime.Value, dictShow_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShowRecordFields_1 gopurs_runtime.Value = dictShowRecordFields_1_loop
_ = dictShowRecordFields_1
var dictShow_2 gopurs_runtime.Value = dictShow_2_loop
_ = dictShow_2
return gopurs_runtime.Value{Type: 9, IntVal: 2498393510, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, record_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_5_0 shape=App(Other) bindingType=String
key_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
_ = key_5_0
return gopurs_runtime.Str((((((" ") + (key_5_0)) + (": ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_2, "show"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_5_0), record_4)).StrVal())) + (",")) + (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictShowRecordFields_1, "showRecordFields"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, record_4).StrVal()))
})}))}
}

func Call_Data_Show_showRecordFieldsConsNil(dictIsSymbol_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
return gopurs_runtime.Value{Type: 9, IntVal: 2498393510, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_ShowRecordFields[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, record_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_4_0 shape=App(Other) bindingType=String
key_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
_ = key_4_0
return gopurs_runtime.Str(((((" ") + (key_4_0)) + (": ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_4_0), record_3)).StrVal())) + (" "))
})}))}
}

func Rebox_Data_Show_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Show_1356436936_1386611502(in *Constructor_Data_Show_Show[[]gopurs_runtime.Value]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Show_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Show_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Show_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Show_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
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
