package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ord_eqRec gopurs_runtime.Value
var once_Data_Ord_eqRec sync.Once
func Get_Data_Ord_eqRec() gopurs_runtime.Value {
	once_Data_Ord_eqRec.Do(func() {
		cache_Data_Ord_eqRec = gopurs_runtime.Func(func(dictEqRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_eqRec(dictEqRecord_0_box)
})
	})
	return cache_Data_Ord_eqRec
}

var cache_Data_Ord_OrdRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Ord_OrdRecord_dollar_Dict sync.Once
func Get_Data_Ord_OrdRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ord_OrdRecord_dollar_Dict.Do(func() {
		cache_Data_Ord_OrdRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer(Call_Data_Ord_OrdRecord_dollar_Dict(func() struct{
	EqRecord0 gopurs_runtime.Value
	compareRecord gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	EqRecord0 gopurs_runtime.Value
	compareRecord gopurs_runtime.Value
}{}
					clone.EqRecord0 = gopurs_runtime.RecordGet(orig, "EqRecord0")
					clone.compareRecord = gopurs_runtime.RecordGet(orig, "compareRecord")
					return clone
				}()))}
})
	})
	return cache_Data_Ord_OrdRecord_dollar_Dict
}

var cache_Data_Ord_Ord_dollar_Dict gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict sync.Once
func Get_Data_Ord_Ord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict(func() struct{
	Eq0 gopurs_runtime.Value
	compare gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Eq0 gopurs_runtime.Value
	compare gopurs_runtime.Value
}{}
					clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
					clone.compare = gopurs_runtime.RecordGet(orig, "compare")
					return clone
				}()))}
})
	})
	return cache_Data_Ord_Ord_dollar_Dict
}

var cache_Data_Ord_Ord1_dollar_Dict gopurs_runtime.Value
var once_Data_Ord_Ord1_dollar_Dict sync.Once
func Get_Data_Ord_Ord1_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ord_Ord1_dollar_Dict.Do(func() {
		cache_Data_Ord_Ord1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord1_dollar_Dict(func() struct{
	Eq10 gopurs_runtime.Value
	compare1 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Eq10 gopurs_runtime.Value
	compare1 gopurs_runtime.Value
}{}
					clone.Eq10 = gopurs_runtime.RecordGet(orig, "Eq10")
					clone.compare1 = gopurs_runtime.RecordGet(orig, "compare1")
					return clone
				}()))}
})
	})
	return cache_Data_Ord_Ord1_dollar_Dict
}

var cache_Data_Ord_ordVoid gopurs_runtime.Value
var once_Data_Ord_ordVoid sync.Once
func Get_Data_Ord_ordVoid() gopurs_runtime.Value {
	once_Data_Ord_ordVoid.Do(func() {
		cache_Data_Ord_ordVoid = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqVoid()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Ord_ordVoid
}

var cache_Data_Ord_ordUnit gopurs_runtime.Value
var once_Data_Ord_ordUnit sync.Once
func Get_Data_Ord_ordUnit() gopurs_runtime.Value {
	once_Data_Ord_ordUnit.Do(func() {
		cache_Data_Ord_ordUnit = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqUnit()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Ord_ordUnit
}

var cache_Data_Ord_ordString gopurs_runtime.Value
var once_Data_Ord_ordString sync.Once
func Get_Data_Ord_ordString() gopurs_runtime.Value {
	once_Data_Ord_ordString.Do(func() {
		cache_Data_Ord_ordString = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordString
}

var cache_Data_Ord_ordRecordNil gopurs_runtime.Value
var once_Data_Ord_ordRecordNil sync.Once
func Get_Data_Ord_ordRecordNil() gopurs_runtime.Value {
	once_Data_Ord_ordRecordNil.Do(func() {
		cache_Data_Ord_ordRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Eq_eqRowNil()))}
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Ord_ordRecordNil
}

var cache_Data_Ord_ordProxy gopurs_runtime.Value
var once_Data_Ord_ordProxy sync.Once
func Get_Data_Ord_ordProxy() gopurs_runtime.Value {
	once_Data_Ord_ordProxy.Do(func() {
		cache_Data_Ord_ordProxy = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3768443459_3790796878(Rebox_Data_Ord_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqProxy()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Ord_ordProxy
}

var cache_Data_Ord_ordOrdering gopurs_runtime.Value
var once_Data_Ord_ordOrdering sync.Once
func Get_Data_Ord_ordOrdering() gopurs_runtime.Value {
	once_Data_Ord_ordOrdering.Do(func() {
		cache_Data_Ord_ordOrdering = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3768443459_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[uint32]](Get_Data_Ordering_eqOrdering())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 uint32
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 1527465420) {
var __t2 uint32
{
var __t_tag_1 uint32 = uint32(v1_1.IntVal)
if (uint32(__t_tag_1) == 1527465420) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t11 = __t2
goto end_branch_11
} else {

}
}
{
var __t_tag_3 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_3) == 902936544) {
var __t7 uint32
{
var __t_tag_4 uint32 = uint32(v1_1.IntVal)
if (uint32(__t_tag_4) == 902936544) {
__t7 = 902936544
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = uint32(v1_1.IntVal)
if (uint32(__t_tag_5) == 1527465420) {
__t7 = 380165415
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = uint32(v1_1.IntVal)
if (uint32(__t_tag_6) == 380165415) {
__t7 = 1527465420
goto end_branch_7
} else {

}
}
{
__t7 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_7:
__t11 = __t7
goto end_branch_11
} else {

}
}
{
var __t_tag_8 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_8) == 380165415) {
var __t10 uint32
{
var __t_tag_9 uint32 = uint32(v1_1.IntVal)
if (uint32(__t_tag_9) == 380165415) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 380165415
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Ord_ordOrdering
}

var cache_Data_Ord_ordNumber gopurs_runtime.Value
var once_Data_Ord_ordNumber sync.Once
func Get_Data_Ord_ordNumber() gopurs_runtime.Value {
	once_Data_Ord_ordNumber.Do(func() {
		cache_Data_Ord_ordNumber = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3047586294_4177771502((&Constructor_Data_Ord_Ord[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_687527510_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[float64]](Get_Data_Eq_eqNumber())))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordNumber
}

var cache_Data_Ord_ordInt gopurs_runtime.Value
var once_Data_Ord_ordInt sync.Once
func Get_Data_Ord_ordInt() gopurs_runtime.Value {
	once_Data_Ord_ordInt.Do(func() {
		cache_Data_Ord_ordInt = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3308271157_4177771502((&Constructor_Data_Ord_Ord[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordInt
}

var cache_Data_Ord_ordChar gopurs_runtime.Value
var once_Data_Ord_ordChar sync.Once
func Get_Data_Ord_ordChar() gopurs_runtime.Value {
	once_Data_Ord_ordChar.Do(func() {
		cache_Data_Ord_ordChar = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqChar())))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordChar
}

var cache_Data_Ord_ordBoolean gopurs_runtime.Value
var once_Data_Ord_ordBoolean sync.Once
func Get_Data_Ord_ordBoolean() gopurs_runtime.Value {
	once_Data_Ord_ordBoolean.Do(func() {
		cache_Data_Ord_ordBoolean = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_219188042_4177771502((&Constructor_Data_Ord_Ord[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_2737952170_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[bool]](Get_Data_Eq_eqBoolean())))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordBooleanImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordBoolean
}

var cache_Data_Ord_compareRecord gopurs_runtime.Value
var once_Data_Ord_compareRecord sync.Once
func Get_Data_Ord_compareRecord() gopurs_runtime.Value {
	once_Data_Ord_compareRecord.Do(func() {
		cache_Data_Ord_compareRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compareRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Ord_compareRecord
}

var cache_Data_Ord_ordRecord gopurs_runtime.Value
var once_Data_Ord_ordRecord sync.Once
func Get_Data_Ord_ordRecord() gopurs_runtime.Value {
	once_Data_Ord_ordRecord.Do(func() {
		cache_Data_Ord_ordRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictOrdRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_ordRecord(_dollar___unused_0_box, dictOrdRecord_1_box)
})
	})
	return cache_Data_Ord_ordRecord
}

var cache_Data_Ord_compare1 gopurs_runtime.Value
var once_Data_Ord_compare1 sync.Once
func Get_Data_Ord_compare1() gopurs_runtime.Value {
	once_Data_Ord_compare1.Do(func() {
		cache_Data_Ord_compare1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Ord_compare1
}

var cache_Data_Ord_compare gopurs_runtime.Value
var once_Data_Ord_compare sync.Once
func Get_Data_Ord_compare() gopurs_runtime.Value {
	once_Data_Ord_compare.Do(func() {
		cache_Data_Ord_compare = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Ord_compare
}

var cache_Data_Ord_compare__378087589 gopurs_runtime.Value
var once_Data_Ord_compare__378087589 sync.Once
func Get_Data_Ord_compare__378087589() gopurs_runtime.Value {
	once_Data_Ord_compare__378087589.Do(func() {
		cache_Data_Ord_compare__378087589 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__378087589(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__378087589
}

var cache_Data_Ord_compare__3355668421 gopurs_runtime.Value
var once_Data_Ord_compare__3355668421 sync.Once
func Get_Data_Ord_compare__3355668421() gopurs_runtime.Value {
	once_Data_Ord_compare__3355668421.Do(func() {
		cache_Data_Ord_compare__3355668421 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3355668421(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__3355668421
}

var cache_Data_Ord_compare__3348508517 gopurs_runtime.Value
var once_Data_Ord_compare__3348508517 sync.Once
func Get_Data_Ord_compare__3348508517() gopurs_runtime.Value {
	once_Data_Ord_compare__3348508517.Do(func() {
		cache_Data_Ord_compare__3348508517 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3348508517(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__3348508517
}

var cache_Data_Ord_compare__3625676717 gopurs_runtime.Value
var once_Data_Ord_compare__3625676717 sync.Once
func Get_Data_Ord_compare__3625676717() gopurs_runtime.Value {
	once_Data_Ord_compare__3625676717.Do(func() {
		cache_Data_Ord_compare__3625676717 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3625676717(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__3625676717
}

var cache_Data_Ord_compare__3009960869 gopurs_runtime.Value
var once_Data_Ord_compare__3009960869 sync.Once
func Get_Data_Ord_compare__3009960869() gopurs_runtime.Value {
	once_Data_Ord_compare__3009960869.Do(func() {
		cache_Data_Ord_compare__3009960869 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3009960869(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__3009960869
}

var cache_Data_Ord_compare__738992261 gopurs_runtime.Value
var once_Data_Ord_compare__738992261 sync.Once
func Get_Data_Ord_compare__738992261() gopurs_runtime.Value {
	once_Data_Ord_compare__738992261.Do(func() {
		cache_Data_Ord_compare__738992261 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__738992261(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__738992261
}

var cache_Data_Ord_compare__3745222789 gopurs_runtime.Value
var once_Data_Ord_compare__3745222789 sync.Once
func Get_Data_Ord_compare__3745222789() gopurs_runtime.Value {
	once_Data_Ord_compare__3745222789.Do(func() {
		cache_Data_Ord_compare__3745222789 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3745222789(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__3745222789
}

var cache_Data_Ord_compare__2876810981 gopurs_runtime.Value
var once_Data_Ord_compare__2876810981 sync.Once
func Get_Data_Ord_compare__2876810981() gopurs_runtime.Value {
	once_Data_Ord_compare__2876810981.Do(func() {
		cache_Data_Ord_compare__2876810981 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__2876810981(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__2876810981
}

var cache_Data_Ord_compare__3195466625 gopurs_runtime.Value
var once_Data_Ord_compare__3195466625 sync.Once
func Get_Data_Ord_compare__3195466625() gopurs_runtime.Value {
	once_Data_Ord_compare__3195466625.Do(func() {
		cache_Data_Ord_compare__3195466625 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3195466625(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Ord_compare__3195466625
}

var cache_Data_Ord_compare__1173898388 gopurs_runtime.Value
var once_Data_Ord_compare__1173898388 sync.Once
func Get_Data_Ord_compare__1173898388() gopurs_runtime.Value {
	once_Data_Ord_compare__1173898388.Do(func() {
		cache_Data_Ord_compare__1173898388 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__1173898388(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Ord_compare__1173898388
}

var cache_Data_Ord_compare__786888001 gopurs_runtime.Value
var once_Data_Ord_compare__786888001 sync.Once
func Get_Data_Ord_compare__786888001() gopurs_runtime.Value {
	once_Data_Ord_compare__786888001.Do(func() {
		cache_Data_Ord_compare__786888001 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__786888001(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Ord_compare__786888001
}

var cache_Data_Ord_compare__2451891644 gopurs_runtime.Value
var once_Data_Ord_compare__2451891644 sync.Once
func Get_Data_Ord_compare__2451891644() gopurs_runtime.Value {
	once_Data_Ord_compare__2451891644.Do(func() {
		cache_Data_Ord_compare__2451891644 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__2451891644(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Ord_compare__2451891644
}

var cache_Data_Ord_compare__282706981 gopurs_runtime.Value
var once_Data_Ord_compare__282706981 sync.Once
func Get_Data_Ord_compare__282706981() gopurs_runtime.Value {
	once_Data_Ord_compare__282706981.Do(func() {
		cache_Data_Ord_compare__282706981 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__282706981(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__282706981
}

var cache_Data_Ord_compare__3469724627 gopurs_runtime.Value
var once_Data_Ord_compare__3469724627 sync.Once
func Get_Data_Ord_compare__3469724627() gopurs_runtime.Value {
	once_Data_Ord_compare__3469724627.Do(func() {
		cache_Data_Ord_compare__3469724627 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3469724627(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Ord_compare__3469724627
}

var cache_Data_Ord_compare__1942422199 gopurs_runtime.Value
var once_Data_Ord_compare__1942422199 sync.Once
func Get_Data_Ord_compare__1942422199() gopurs_runtime.Value {
	once_Data_Ord_compare__1942422199.Do(func() {
		cache_Data_Ord_compare__1942422199 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__1942422199(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Ord_compare__1942422199
}

var cache_Data_Ord_compare__4035906730 gopurs_runtime.Value
var once_Data_Ord_compare__4035906730 sync.Once
func Get_Data_Ord_compare__4035906730() gopurs_runtime.Value {
	once_Data_Ord_compare__4035906730.Do(func() {
		cache_Data_Ord_compare__4035906730 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__4035906730(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Ord_compare__4035906730
}

var cache_Data_Ord_compare__3399466676 gopurs_runtime.Value
var once_Data_Ord_compare__3399466676 sync.Once
func Get_Data_Ord_compare__3399466676() gopurs_runtime.Value {
	once_Data_Ord_compare__3399466676.Do(func() {
		cache_Data_Ord_compare__3399466676 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3399466676(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Ord_compare__3399466676
}

var cache_Data_Ord_compare__3926472293 gopurs_runtime.Value
var once_Data_Ord_compare__3926472293 sync.Once
func Get_Data_Ord_compare__3926472293() gopurs_runtime.Value {
	once_Data_Ord_compare__3926472293.Do(func() {
		cache_Data_Ord_compare__3926472293 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3926472293(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ord_compare__3926472293
}

var cache_Data_Ord_comparing gopurs_runtime.Value
var once_Data_Ord_comparing sync.Once
func Get_Data_Ord_comparing() gopurs_runtime.Value {
	once_Data_Ord_comparing.Do(func() {
		cache_Data_Ord_comparing = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_comparing(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_comparing
}

var cache_Data_Ord_greaterThan gopurs_runtime.Value
var once_Data_Ord_greaterThan sync.Once
func Get_Data_Ord_greaterThan() gopurs_runtime.Value {
	once_Data_Ord_greaterThan.Do(func() {
		cache_Data_Ord_greaterThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_greaterThan
}

var cache_Data_Ord_greaterThan__777679551 gopurs_runtime.Value
var once_Data_Ord_greaterThan__777679551 sync.Once
func Get_Data_Ord_greaterThan__777679551() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__777679551.Do(func() {
		cache_Data_Ord_greaterThan__777679551 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__777679551(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](a1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](a2_1_box)))
})
	})
	return cache_Data_Ord_greaterThan__777679551
}

var cache_Data_Ord_greaterThan__1265619615 gopurs_runtime.Value
var once_Data_Ord_greaterThan__1265619615 sync.Once
func Get_Data_Ord_greaterThan__1265619615() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__1265619615.Do(func() {
		cache_Data_Ord_greaterThan__1265619615 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__1265619615(a1_0_box.IntVal, a2_1_box.IntVal))
})
	})
	return cache_Data_Ord_greaterThan__1265619615
}

var cache_Data_Ord_greaterThan__954999263 gopurs_runtime.Value
var once_Data_Ord_greaterThan__954999263 sync.Once
func Get_Data_Ord_greaterThan__954999263() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__954999263.Do(func() {
		cache_Data_Ord_greaterThan__954999263 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__954999263(a1_0_box.FloatVal(), a2_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_greaterThan__954999263
}

var cache_Data_Ord_greaterThan__294895920 gopurs_runtime.Value
var once_Data_Ord_greaterThan__294895920 sync.Once
func Get_Data_Ord_greaterThan__294895920() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__294895920.Do(func() {
		cache_Data_Ord_greaterThan__294895920 = gopurs_runtime.Func2(func(a2_unused_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__294895920(a2_unused_0_box.FloatVal(), a1_1_box))
})
	})
	return cache_Data_Ord_greaterThan__294895920
}

var cache_Data_Ord_greaterThanOrEq gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq sync.Once
func Get_Data_Ord_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq.Do(func() {
		cache_Data_Ord_greaterThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_greaterThanOrEq
}

var cache_Data_Ord_greaterThanOrEq__1265619615 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__1265619615 sync.Once
func Get_Data_Ord_greaterThanOrEq__1265619615() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__1265619615.Do(func() {
		cache_Data_Ord_greaterThanOrEq__1265619615 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__1265619615(a1_0_box.IntVal, a2_1_box.IntVal))
})
	})
	return cache_Data_Ord_greaterThanOrEq__1265619615
}

var cache_Data_Ord_greaterThanOrEq__1859892494 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__1859892494 sync.Once
func Get_Data_Ord_greaterThanOrEq__1859892494() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__1859892494.Do(func() {
		cache_Data_Ord_greaterThanOrEq__1859892494 = gopurs_runtime.Func2(func(a2_unused_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__1859892494(a2_unused_0_box.IntVal, a1_1_box))
})
	})
	return cache_Data_Ord_greaterThanOrEq__1859892494
}

var cache_Data_Ord_greaterThanOrEq__2271442022 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__2271442022 sync.Once
func Get_Data_Ord_greaterThanOrEq__2271442022() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__2271442022.Do(func() {
		cache_Data_Ord_greaterThanOrEq__2271442022 = gopurs_runtime.Func2(func(a2_unused_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__2271442022(a2_unused_0_box.IntVal, a1_1_box))
})
	})
	return cache_Data_Ord_greaterThanOrEq__2271442022
}

var cache_Data_Ord_greaterThanOrEq__954999263 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__954999263 sync.Once
func Get_Data_Ord_greaterThanOrEq__954999263() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__954999263.Do(func() {
		cache_Data_Ord_greaterThanOrEq__954999263 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__954999263(a1_0_box.FloatVal(), a2_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_greaterThanOrEq__954999263
}

var cache_Data_Ord_greaterThanOrEq__1036186893 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__1036186893 sync.Once
func Get_Data_Ord_greaterThanOrEq__1036186893() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__1036186893.Do(func() {
		cache_Data_Ord_greaterThanOrEq__1036186893 = gopurs_runtime.Func2(func(a2_unused_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__1036186893(a2_unused_0_box.FloatVal(), a1_1_box))
})
	})
	return cache_Data_Ord_greaterThanOrEq__1036186893
}

var cache_Data_Ord_lessThan gopurs_runtime.Value
var once_Data_Ord_lessThan sync.Once
func Get_Data_Ord_lessThan() gopurs_runtime.Value {
	once_Data_Ord_lessThan.Do(func() {
		cache_Data_Ord_lessThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_lessThan
}

var cache_Data_Ord_lessThan__1265619615 gopurs_runtime.Value
var once_Data_Ord_lessThan__1265619615 sync.Once
func Get_Data_Ord_lessThan__1265619615() gopurs_runtime.Value {
	once_Data_Ord_lessThan__1265619615.Do(func() {
		cache_Data_Ord_lessThan__1265619615 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan__1265619615(a1_0_box.IntVal, a2_1_box.IntVal))
})
	})
	return cache_Data_Ord_lessThan__1265619615
}

var cache_Data_Ord_lessThan__854894971 gopurs_runtime.Value
var once_Data_Ord_lessThan__854894971 sync.Once
func Get_Data_Ord_lessThan__854894971() gopurs_runtime.Value {
	once_Data_Ord_lessThan__854894971.Do(func() {
		cache_Data_Ord_lessThan__854894971 = gopurs_runtime.Func2(func(a2_unused_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan__854894971(a2_unused_0_box.IntVal, a1_1_box))
})
	})
	return cache_Data_Ord_lessThan__854894971
}

var cache_Data_Ord_lessThan__954999263 gopurs_runtime.Value
var once_Data_Ord_lessThan__954999263 sync.Once
func Get_Data_Ord_lessThan__954999263() gopurs_runtime.Value {
	once_Data_Ord_lessThan__954999263.Do(func() {
		cache_Data_Ord_lessThan__954999263 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan__954999263(a1_0_box.FloatVal(), a2_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_lessThan__954999263
}

var cache_Data_Ord_lessThan__2434180334 gopurs_runtime.Value
var once_Data_Ord_lessThan__2434180334 sync.Once
func Get_Data_Ord_lessThan__2434180334() gopurs_runtime.Value {
	once_Data_Ord_lessThan__2434180334.Do(func() {
		cache_Data_Ord_lessThan__2434180334 = gopurs_runtime.Func2(func(a2_unused_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan__2434180334(a2_unused_0_box.FloatVal(), a1_1_box))
})
	})
	return cache_Data_Ord_lessThan__2434180334
}

var cache_Data_Ord_signum gopurs_runtime.Value
var once_Data_Ord_signum sync.Once
func Get_Data_Ord_signum() gopurs_runtime.Value {
	once_Data_Ord_signum.Do(func() {
		cache_Data_Ord_signum = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_signum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dictRing_1_box))
})
	})
	return cache_Data_Ord_signum
}

var cache_Data_Ord_lessThanOrEq gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq sync.Once
func Get_Data_Ord_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq.Do(func() {
		cache_Data_Ord_lessThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_lessThanOrEq
}

var cache_Data_Ord_lessThanOrEq__1265619615 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__1265619615 sync.Once
func Get_Data_Ord_lessThanOrEq__1265619615() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__1265619615.Do(func() {
		cache_Data_Ord_lessThanOrEq__1265619615 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__1265619615(a1_0_box.IntVal, a2_1_box.IntVal))
})
	})
	return cache_Data_Ord_lessThanOrEq__1265619615
}

var cache_Data_Ord_lessThanOrEq__1401012923 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__1401012923 sync.Once
func Get_Data_Ord_lessThanOrEq__1401012923() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__1401012923.Do(func() {
		cache_Data_Ord_lessThanOrEq__1401012923 = gopurs_runtime.Func2(func(a2_unused_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__1401012923(a2_unused_0_box.IntVal, a1_1_box))
})
	})
	return cache_Data_Ord_lessThanOrEq__1401012923
}

var cache_Data_Ord_lessThanOrEq__954999263 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__954999263 sync.Once
func Get_Data_Ord_lessThanOrEq__954999263() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__954999263.Do(func() {
		cache_Data_Ord_lessThanOrEq__954999263 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__954999263(a1_0_box.FloatVal(), a2_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_lessThanOrEq__954999263
}

var cache_Data_Ord_lessThanOrEq__110343017 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__110343017 sync.Once
func Get_Data_Ord_lessThanOrEq__110343017() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__110343017.Do(func() {
		cache_Data_Ord_lessThanOrEq__110343017 = gopurs_runtime.Func2(func(a2_unused_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__110343017(a2_unused_0_box.FloatVal(), a1_1_box))
})
	})
	return cache_Data_Ord_lessThanOrEq__110343017
}

var cache_Data_Ord_max gopurs_runtime.Value
var once_Data_Ord_max sync.Once
func Get_Data_Ord_max() gopurs_runtime.Value {
	once_Data_Ord_max.Do(func() {
		cache_Data_Ord_max = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_max(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_Data_Ord_max
}

var cache_Data_Ord_max__3052583336 gopurs_runtime.Value
var once_Data_Ord_max__3052583336 sync.Once
func Get_Data_Ord_max__3052583336() gopurs_runtime.Value {
	once_Data_Ord_max__3052583336.Do(func() {
		cache_Data_Ord_max__3052583336 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Ord_max__3052583336(x_0_box.IntVal, y_1_box.IntVal))
})
	})
	return cache_Data_Ord_max__3052583336
}

var cache_Data_Ord_min gopurs_runtime.Value
var once_Data_Ord_min sync.Once
func Get_Data_Ord_min() gopurs_runtime.Value {
	once_Data_Ord_min.Do(func() {
		cache_Data_Ord_min = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_min(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_Data_Ord_min
}

var cache_Data_Ord_min__3052583336 gopurs_runtime.Value
var once_Data_Ord_min__3052583336 sync.Once
func Get_Data_Ord_min__3052583336() gopurs_runtime.Value {
	once_Data_Ord_min__3052583336.Do(func() {
		cache_Data_Ord_min__3052583336 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Ord_min__3052583336(x_0_box.IntVal, y_1_box.IntVal))
})
	})
	return cache_Data_Ord_min__3052583336
}

var cache_Data_Ord_min__294757560 gopurs_runtime.Value
var once_Data_Ord_min__294757560 sync.Once
func Get_Data_Ord_min__294757560() gopurs_runtime.Value {
	once_Data_Ord_min__294757560.Do(func() {
		cache_Data_Ord_min__294757560 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Ord_min__294757560(x_0_box.FloatVal(), y_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_min__294757560
}

var cache_Data_Ord_ordArray gopurs_runtime.Value
var once_Data_Ord_ordArray sync.Once
func Get_Data_Ord_ordArray() gopurs_runtime.Value {
	once_Data_Ord_ordArray.Do(func() {
		cache_Data_Ord_ordArray = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_ordArray(dictOrd_0_box)
})
	})
	return cache_Data_Ord_ordArray
}

var cache_Data_Ord_ord1Array gopurs_runtime.Value
var once_Data_Ord_ord1Array sync.Once
func Get_Data_Ord_ord1Array() gopurs_runtime.Value {
	once_Data_Ord_ord1Array.Do(func() {
		cache_Data_Ord_ord1Array = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}
}), gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(int64(0)), gopurs_runtime.Apply3(Get_Data_Ord_ordArrayImpl(), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_5_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_3, y_4).IntVal)
_ = v_5_0
var __t1 int64
{
if (v_5_0 == 902936544) {
__t1 = int64(0)
goto end_branch_1
} else {

}
}
{
if (v_5_0 == 1527465420) {
__t1 = int64(1)
goto end_branch_1
} else {

}
}
{
if (v_5_0 == 380165415) {
__t1 = int64(-1)
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Int(__t1)
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal)), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Ord_ord1Array
}

var cache_Data_Ord_ordRecordCons gopurs_runtime.Value
var once_Data_Ord_ordRecordCons sync.Once
func Get_Data_Ord_ordRecordCons() gopurs_runtime.Value {
	once_Data_Ord_ordRecordCons.Do(func() {
		cache_Data_Ord_ordRecordCons = gopurs_runtime.Func(func(dictOrdRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_ordRecordCons(dictOrdRecord_0_box)
})
	})
	return cache_Data_Ord_ordRecordCons
}

var cache_Data_Ord_clamp gopurs_runtime.Value
var once_Data_Ord_clamp sync.Once
func Get_Data_Ord_clamp() gopurs_runtime.Value {
	once_Data_Ord_clamp.Do(func() {
		cache_Data_Ord_clamp = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_clamp(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), low_1_box, hi_2_box, x_3_box)
})
	})
	return cache_Data_Ord_clamp
}

var cache_Data_Ord_clamp__3109832964 gopurs_runtime.Value
var once_Data_Ord_clamp__3109832964 sync.Once
func Get_Data_Ord_clamp__3109832964() gopurs_runtime.Value {
	once_Data_Ord_clamp__3109832964.Do(func() {
		cache_Data_Ord_clamp__3109832964 = gopurs_runtime.Func3(func(low_0_box gopurs_runtime.Value, hi_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Ord_clamp__3109832964(low_0_box.IntVal, hi_1_box.IntVal, x_2_box.IntVal))
})
	})
	return cache_Data_Ord_clamp__3109832964
}

var cache_Data_Ord_between gopurs_runtime.Value
var once_Data_Ord_between sync.Once
func Get_Data_Ord_between() gopurs_runtime.Value {
	once_Data_Ord_between.Do(func() {
		cache_Data_Ord_between = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_between(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), low_1_box, hi_2_box, x_3_box))
})
	})
	return cache_Data_Ord_between
}

var cache_Data_Ord_abs gopurs_runtime.Value
var once_Data_Ord_abs sync.Once
func Get_Data_Ord_abs() gopurs_runtime.Value {
	once_Data_Ord_abs.Do(func() {
		cache_Data_Ord_abs = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_abs(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dictRing_1_box))
})
	})
	return cache_Data_Ord_abs
}

var cache_Data_Ord_abs__3883318852 gopurs_runtime.Value
var once_Data_Ord_abs__3883318852 sync.Once
func Get_Data_Ord_abs__3883318852() gopurs_runtime.Value {
	once_Data_Ord_abs__3883318852.Do(func() {
		cache_Data_Ord_abs__3883318852 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_abs__3883318852(__eta_norm_0_0_box)
})
	})
	return cache_Data_Ord_abs__3883318852
}

type Constructor_Data_Ord_OrdRecord[T_rowlist any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4162894775] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_OrdRecord[any, any])(ptr)
		_ = c
		switch key {
		case "EqRecord0": return gopurs_runtime.Box(c.V0)
		case "compareRecord": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ord_OrdRecord: " + key)
		}
	}
}


type Constructor_Data_Ord_Ord[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1435789946] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_Ord[any])(ptr)
		_ = c
		switch key {
		case "Eq0": return gopurs_runtime.Box(c.V0)
		case "compare": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ord_Ord: " + key)
		}
	}
}


type Constructor_Data_Ord_Ord1[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1632188299] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_Ord1[any])(ptr)
		_ = c
		switch key {
		case "Eq10": return gopurs_runtime.Box(c.V0)
		case "compare1": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ord_Ord1: " + key)
		}
	}
}


func Call_Data_Ord_eqRec(dictEqRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEqRecord_0 gopurs_runtime.Value = dictEqRecord_0_loop
_ = dictEqRecord_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEqRecord_0, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Ord_OrdRecord_dollar_Dict(x_0_loop struct{
	EqRecord0 gopurs_runtime.Value
	compareRecord gopurs_runtime.Value
}) *Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	EqRecord0 gopurs_runtime.Value
	compareRecord gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("EqRecord0", "compareRecord", orig.EqRecord0, orig.compareRecord)
				}())
}

func Call_Data_Ord_Ord_dollar_Dict(x_0_loop struct{
	Eq0 gopurs_runtime.Value
	compare gopurs_runtime.Value
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
var x_0 struct{
	Eq0 gopurs_runtime.Value
	compare gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
				}())
}

func Call_Data_Ord_Ord1_dollar_Dict(x_0_loop struct{
	Eq10 gopurs_runtime.Value
	compare1 gopurs_runtime.Value
}) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
var x_0 struct{
	Eq10 gopurs_runtime.Value
	compare1 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Eq10", "compare1", orig.Eq10, orig.compare1)
				}())
}

func Call_Data_Ord_compareRecord(dict_0_loop *Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_ordRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictOrdRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictOrdRecord_1 gopurs_runtime.Value = dictOrdRecord_1_loop
_ = dictOrdRecord_1
// TAST (Let): eqRec1_2_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Record (Row [] (TypeVar row)))])
eqRec1_2_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = eqRec1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqRec1_2_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Ord_compare1(dict_0_loop *Constructor_Data_Ord_Ord1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare(dict_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__378087589(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__378087589:
for {
if false { continue compare__378087589 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Date_Component_ordMonth()).V1), __eta_norm_1_0, __eta_norm_0_1).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3355668421(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3355668421:
for {
if false { continue compare__3355668421 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]](Get_Data_Date_ordDate()).V1), __eta_norm_1_0, __eta_norm_0_1).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3348508517(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3348508517:
for {
if false { continue compare__3348508517 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t27 uint32
{
var __t_tag_0 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_0) == 3908053364) {
var __t2 uint32
{
var __t_tag_1 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_1) == 3908053364) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t27 = __t2
goto end_branch_27
} else {

}
}
{
var __t_tag_3 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_3) == 3908053364) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_4 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_4) == 217821258) {
var __t6 uint32
{
var __t_tag_5 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_5) == 217821258) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t27 = __t6
goto end_branch_27
} else {

}
}
{
var __t_tag_7 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_7) == 217821258) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_8 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_8) == 1292308612) {
var __t10 uint32
{
var __t_tag_9 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_9) == 1292308612) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t27 = __t10
goto end_branch_27
} else {

}
}
{
var __t_tag_11 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_11) == 1292308612) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_12 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_12) == 2311060696) {
var __t14 uint32
{
var __t_tag_13 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_13) == 2311060696) {
__t14 = 902936544
goto end_branch_14
} else {

}
}
{
__t14 = 1527465420
}
end_branch_14:
__t27 = __t14
goto end_branch_27
} else {

}
}
{
var __t_tag_15 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_15) == 2311060696) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_16 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_16) == 401302776) {
var __t18 uint32
{
var __t_tag_17 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_17) == 401302776) {
__t18 = 902936544
goto end_branch_18
} else {

}
}
{
__t18 = 1527465420
}
end_branch_18:
__t27 = __t18
goto end_branch_27
} else {

}
}
{
var __t_tag_19 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_19) == 401302776) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_20 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_20) == 3327533908) {
var __t22 uint32
{
var __t_tag_21 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_21) == 3327533908) {
__t22 = 902936544
goto end_branch_22
} else {

}
}
{
__t22 = 1527465420
}
end_branch_22:
__t27 = __t22
goto end_branch_27
} else {

}
}
{
var __t_tag_23 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_23) == 3327533908) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_24 uint32 = uint32(__eta_norm_1_0.IntVal)
var __t_and_26 bool = false
if (uint32(__t_tag_24) == 3631736139) {

var __t_tag_25 uint32 = uint32(__eta_norm_0_1.IntVal)
__t_and_26 = (uint32(__t_tag_25) == 3631736139)
}
if __t_and_26 {
__t27 = 902936544
goto end_branch_27
} else {

}
}
{
__t27 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t27), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3625676717(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3625676717:
for {
if false { continue compare__3625676717 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t5 uint32
{
var __t_tag_0 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_0) == 3908053364) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
var __t_tag_1 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_1) == 217821258) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
var __t_tag_2 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_2) == 1292308612) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
var __t_tag_3 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_3) == 2311060696) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
var __t_tag_4 uint32 = uint32(__eta_norm_0_1.IntVal)
if (uint32(__t_tag_4) == 401302776) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3009960869(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3009960869:
for {
if false { continue compare__3009960869 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_ordMap()).V1), __eta_norm_1_0, __eta_norm_0_1).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__738992261(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__738992261:
for {
if false { continue compare__738992261 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_ordMaybe()).V1), __eta_norm_1_0, __eta_norm_0_1).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3745222789(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3745222789:
for {
if false { continue compare__3745222789 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time]](Get_Data_Time_ordTime()).V1), __eta_norm_1_0, __eta_norm_0_1).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__2876810981(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__2876810981:
for {
if false { continue compare__2876810981 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_0, __eta_norm_0_1).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3195466625(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3195466625:
for {
if false { continue compare__3195466625 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_1, gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()).StrVal())).IntVal)).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__1173898388(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__1173898388:
for {
if false { continue compare__1173898388 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_1, gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal)).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__786888001(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__786888001:
for {
if false { continue compare__786888001 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_1, gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())).IntVal)).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__2451891644(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__2451891644:
for {
if false { continue compare__2451891644 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_1, gopurs_runtime.Int(gopurs_runtime.Int(int64(0)).IntVal)).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__282706981(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__282706981:
for {
if false { continue compare__282706981 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_0, __eta_norm_0_1).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3469724627(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3469724627:
for {
if false { continue compare__3469724627 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_1, gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal)).FloatVal())).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__1942422199(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__1942422199:
for {
if false { continue compare__1942422199 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_1, gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal)).FloatVal())).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__4035906730(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__4035906730:
for {
if false { continue compare__4035906730 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_1, gopurs_runtime.Float(Get_Data_Time_maxTime().FloatVal())).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3399466676(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3399466676:
for {
if false { continue compare__3399466676 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_1, gopurs_runtime.Float(Get_Data_Time_minTime().FloatVal())).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_compare__3926472293(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compare__3926472293:
for {
if false { continue compare__3926472293 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta_norm_1_0, __eta_norm_0_1).IntVal)), UnsafePtr: nil}
}
}

func Call_Data_Ord_comparing(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)
}

func Call_Data_Ord_greaterThan(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2).IntVal)
return (uint32(__t_tag_0) == 380165415)
}

func Call_Data_Ord_greaterThan__777679551(a1_0_loop *Constructor_Data_Maybe_Just[int64], a2_1_loop *Constructor_Data_Maybe_Just[int64]) bool {
greaterThan__777679551:
for {
if false { continue greaterThan__777679551 }
var a1_0 *Constructor_Data_Maybe_Just[int64] = a1_0_loop
_ = a1_0
var a2_1 *Constructor_Data_Maybe_Just[int64] = a2_1_loop
_ = a2_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_ordMaybe()).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1170268447_3094389156(a1_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1170268447_3094389156(a2_1))})
return (uint32(__t_tag_0.IntVal) == 380165415)
}
}

func Call_Data_Ord_greaterThan__1265619615(a1_0_loop int64, a2_1_loop int64) bool {
greaterThan__1265619615:
for {
if false { continue greaterThan__1265619615 }
var a1_0 int64 = a1_0_loop
_ = a1_0
var a2_1 int64 = a2_1_loop
_ = a2_1
return (a1_0) > (a2_1)
}
}

func Call_Data_Ord_greaterThan__954999263(a1_0_loop float64, a2_1_loop float64) bool {
greaterThan__954999263:
for {
if false { continue greaterThan__954999263 }
var a1_0 float64 = a1_0_loop
_ = a1_0
var a2_1 float64 = a2_1_loop
_ = a2_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(a1_0), gopurs_runtime.Float(a2_1))
return (uint32(__t_tag_0.IntVal) == 380165415)
}
}

func Call_Data_Ord_greaterThan__294895920(a2_unused_0_loop float64, a1_1_loop gopurs_runtime.Value) bool {
greaterThan__294895920:
for {
if false { continue greaterThan__294895920 }
var a2_unused_0 float64 = a2_unused_0_loop
_ = a2_unused_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(Get_Data_Time_maxTime().FloatVal()), gopurs_runtime.Float(Get_Data_Time_maxTime().FloatVal()))
return (uint32(__t_tag_0.IntVal) == 380165415)
}
}

func Call_Data_Ord_greaterThanOrEq(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2).IntVal)
return ((uint32(__t_tag_0) == 1527465420)) != (true)
}

func Call_Data_Ord_greaterThanOrEq__1265619615(a1_0_loop int64, a2_1_loop int64) bool {
greaterThanOrEq__1265619615:
for {
if false { continue greaterThanOrEq__1265619615 }
var a1_0 int64 = a1_0_loop
_ = a1_0
var a2_1 int64 = a2_1_loop
_ = a2_1
return ((a1_0) < (a2_1)) != (true)
}
}

func Call_Data_Ord_greaterThanOrEq__1859892494(a2_unused_0_loop int64, a1_1_loop gopurs_runtime.Value) bool {
greaterThanOrEq__1859892494:
for {
if false { continue greaterThanOrEq__1859892494 }
var a2_unused_0 int64 = a2_unused_0_loop
_ = a2_unused_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
return (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal) >= (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal)
}
}

func Call_Data_Ord_greaterThanOrEq__2271442022(a2_unused_0_loop int64, a1_1_loop gopurs_runtime.Value) bool {
greaterThanOrEq__2271442022:
for {
if false { continue greaterThanOrEq__2271442022 }
var a2_unused_0 int64 = a2_unused_0_loop
_ = a2_unused_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
return true
}
}

func Call_Data_Ord_greaterThanOrEq__954999263(a1_0_loop float64, a2_1_loop float64) bool {
greaterThanOrEq__954999263:
for {
if false { continue greaterThanOrEq__954999263 }
var a1_0 float64 = a1_0_loop
_ = a1_0
var a2_1 float64 = a2_1_loop
_ = a2_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(a1_0), gopurs_runtime.Float(a2_1))
return ((uint32(__t_tag_0.IntVal) == 1527465420)) != (true)
}
}

func Call_Data_Ord_greaterThanOrEq__1036186893(a2_unused_0_loop float64, a1_1_loop gopurs_runtime.Value) bool {
greaterThanOrEq__1036186893:
for {
if false { continue greaterThanOrEq__1036186893 }
var a2_unused_0 float64 = a2_unused_0_loop
_ = a2_unused_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal)), gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal)).FloatVal()))
return ((uint32(__t_tag_0.IntVal) == 1527465420)) != (true)
}
}

func Call_Data_Ord_lessThan(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2).IntVal)
return (uint32(__t_tag_0) == 1527465420)
}

func Call_Data_Ord_lessThan__1265619615(a1_0_loop int64, a2_1_loop int64) bool {
lessThan__1265619615:
for {
if false { continue lessThan__1265619615 }
var a1_0 int64 = a1_0_loop
_ = a1_0
var a2_1 int64 = a2_1_loop
_ = a2_1
return (a1_0) < (a2_1)
}
}

func Call_Data_Ord_lessThan__854894971(a2_unused_0_loop int64, a1_1_loop gopurs_runtime.Value) bool {
lessThan__854894971:
for {
if false { continue lessThan__854894971 }
var a2_unused_0 int64 = a2_unused_0_loop
_ = a2_unused_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
return (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()).StrVal())).IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()).StrVal())).IntVal)
}
}

func Call_Data_Ord_lessThan__954999263(a1_0_loop float64, a2_1_loop float64) bool {
lessThan__954999263:
for {
if false { continue lessThan__954999263 }
var a1_0 float64 = a1_0_loop
_ = a1_0
var a2_1 float64 = a2_1_loop
_ = a2_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(a1_0), gopurs_runtime.Float(a2_1))
return (uint32(__t_tag_0.IntVal) == 1527465420)
}
}

func Call_Data_Ord_lessThan__2434180334(a2_unused_0_loop float64, a1_1_loop gopurs_runtime.Value) bool {
lessThan__2434180334:
for {
if false { continue lessThan__2434180334 }
var a2_unused_0 float64 = a2_unused_0_loop
_ = a2_unused_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(Get_Data_Time_minTime().FloatVal()), gopurs_runtime.Float(Get_Data_Time_minTime().FloatVal()))
return (uint32(__t_tag_0.IntVal) == 1527465420)
}
}

func Call_Data_Ord_signum(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictRing_1_loop *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 *Constructor_Data_Ring_Ring[gopurs_runtime.Value] = dictRing_1_loop
_ = dictRing_1
// TAST (Let): Semiring0_2_0 shape=App(Other) bindingType=Any
Semiring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{})
_ = Semiring0_2_0
// TAST (Let): zero_3_1 shape=Other bindingType=(TypeVar a)
zero_3_1 := gopurs_runtime.RecordGet(Semiring0_2_0, "zero")
_ = zero_3_1
// TAST (Let): Semiring01_4_2 shape=App(Other) bindingType=(ADT ["Data","Semiring","Semiring"] [(TypeVar a)])
Semiring01_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}))
_ = Semiring01_4_2
// TAST (Let): one_5_3 shape=Other bindingType=(TypeVar a)
one_5_3 := gopurs_runtime.RecordGet(Semiring0_2_0, "one")
_ = one_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_6 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_6, zero_3_1).IntVal)
if (uint32(__t_tag_6) == 1527465420) {
__t7 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_1.V1), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}), "zero"), gopurs_runtime.Box(Semiring01_4_2.V2))
goto end_branch_7
} else {

}
}
{
var __t5 gopurs_runtime.Value
{
var __t_tag_4 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_6, zero_3_1).IntVal)
if (uint32(__t_tag_4) == 380165415) {
__t5 = one_5_3
goto end_branch_5
} else {

}
}
{
__t5 = x_6
}
end_branch_5:
__t7 = __t5
}
end_branch_7:
return __t7
})
}

func Call_Data_Ord_lessThanOrEq(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2).IntVal)
return ((uint32(__t_tag_0) == 380165415)) != (true)
}

func Call_Data_Ord_lessThanOrEq__1265619615(a1_0_loop int64, a2_1_loop int64) bool {
lessThanOrEq__1265619615:
for {
if false { continue lessThanOrEq__1265619615 }
var a1_0 int64 = a1_0_loop
_ = a1_0
var a2_1 int64 = a2_1_loop
_ = a2_1
return ((a1_0) > (a2_1)) != (true)
}
}

func Call_Data_Ord_lessThanOrEq__1401012923(a2_unused_0_loop int64, a1_1_loop gopurs_runtime.Value) bool {
lessThanOrEq__1401012923:
for {
if false { continue lessThanOrEq__1401012923 }
var a2_unused_0 int64 = a2_unused_0_loop
_ = a2_unused_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
return (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())).IntVal) <= (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())).IntVal)
}
}

func Call_Data_Ord_lessThanOrEq__954999263(a1_0_loop float64, a2_1_loop float64) bool {
lessThanOrEq__954999263:
for {
if false { continue lessThanOrEq__954999263 }
var a1_0 float64 = a1_0_loop
_ = a1_0
var a2_1 float64 = a2_1_loop
_ = a2_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(a1_0), gopurs_runtime.Float(a2_1))
return ((uint32(__t_tag_0.IntVal) == 380165415)) != (true)
}
}

func Call_Data_Ord_lessThanOrEq__110343017(a2_unused_0_loop float64, a1_1_loop gopurs_runtime.Value) bool {
lessThanOrEq__110343017:
for {
if false { continue lessThanOrEq__110343017 }
var a2_unused_0 float64 = a2_unused_0_loop
_ = a2_unused_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal)), gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal)).FloatVal()))
return ((uint32(__t_tag_0.IntVal) == 380165415)) != (true)
}
}

func Call_Data_Ord_max(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_3_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_1, y_2).IntVal)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (v_3_0 == 1527465420) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
if (v_3_0 == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (v_3_0 == 380165415) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_Data_Ord_max__3052583336(x_0_loop int64, y_1_loop int64) int64 {
max__3052583336:
for {
if false { continue max__3052583336 }
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_2_0 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(x_0), gopurs_runtime.Int(y_1)).IntVal)
_ = v_2_0
var __t1 int64
{
if (v_2_0 == 1527465420) {
__t1 = y_1
goto end_branch_1
} else {

}
}
{
if (v_2_0 == 902936544) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
if (v_2_0 == 380165415) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}

func Call_Data_Ord_min(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_3_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_1, y_2).IntVal)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (v_3_0 == 1527465420) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (v_3_0 == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (v_3_0 == 380165415) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_Data_Ord_min__3052583336(x_0_loop int64, y_1_loop int64) int64 {
min__3052583336:
for {
if false { continue min__3052583336 }
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_2_0 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(x_0), gopurs_runtime.Int(y_1)).IntVal)
_ = v_2_0
var __t1 int64
{
if (v_2_0 == 1527465420) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
if (v_2_0 == 902936544) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
if (v_2_0 == 380165415) {
__t1 = y_1
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}

func Call_Data_Ord_min__294757560(x_0_loop float64, y_1_loop float64) float64 {
min__294757560:
for {
if false { continue min__294757560 }
var x_0 float64 = x_0_loop
_ = x_0
var y_1 float64 = y_1_loop
_ = y_1
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_2_0 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(x_0), gopurs_runtime.Float(y_1)).IntVal)
_ = v_2_0
var __t1 float64
{
if (v_2_0 == 1527465420) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
if (v_2_0 == 902936544) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
if (v_2_0 == 380165415) {
__t1 = y_1
goto end_branch_1
} else {

}
}
{
__t1 = func() float64 { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}

func Call_Data_Ord_ordArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqArray_1_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
eqArray_1_0 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}), "eq"))})
_ = eqArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_4088624520_4177771502((&Constructor_Data_Ord_Ord[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1939691112_3790796878(eqArray_1_0))}
}), gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(int64(0)), gopurs_runtime.Apply3(Get_Data_Ord_ordArrayImpl(), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_6_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, y_5).IntVal)
_ = v_6_1
var __t2 int64
{
if (v_6_1 == 902936544) {
__t2 = int64(0)
goto end_branch_2
} else {

}
}
{
if (v_6_1 == 1527465420) {
__t2 = int64(1)
goto end_branch_2
} else {

}
}
{
if (v_6_1 == 380165415) {
__t2 = int64(-1)
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Int(__t2)
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal)), UnsafePtr: nil}
})})))}
}

func Call_Data_Ord_ordRecordCons(dictOrdRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrdRecord_0 gopurs_runtime.Value = dictOrdRecord_0_loop
_ = dictOrdRecord_0
// TAST (Let): eqRowCons__193435443_1_0 shape=App(Var) bindingType=Any
eqRowCons__193435443_1_0 := gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "EqRecord0"), gopurs_runtime.Value{}), gopurs_runtime.Value{})
_ = eqRowCons__193435443_1_0
return gopurs_runtime.Func2(func(_dollar___unused_2 gopurs_runtime.Value, dictIsSymbol_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqRowCons1__193435443_4_1 shape=App(Other) bindingType=Any
eqRowCons1__193435443_4_1 := gopurs_runtime.Apply(eqRowCons__193435443_1_0, dictIsSymbol_3)
_ = eqRowCons1__193435443_4_1
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqRowCons2_6_2 shape=App(Other) bindingType=(TypeApp (ADT ["Data","Eq","EqRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)]), (TypeVar row)])
eqRowCons2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(eqRowCons1__193435443_4_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})))
_ = eqRowCons2_6_2
return gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(eqRowCons2_6_2)}
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, ra_8 gopurs_runtime.Value, rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrdRecord_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9).IntVal)), UnsafePtr: nil}
})}))}
})
})
}

func Call_Data_Ord_clamp(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
// TAST (Let): v_4_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_4_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), low_1, x_3).IntVal)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (v_4_1 == 1527465420) {
__t2 = x_3
goto end_branch_2
} else {

}
}
{
if (v_4_1 == 902936544) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
if (v_4_1 == 380165415) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
// TAST (Let): __local_var_4_0 shape=Let(Branch(Other, Other, Other, def=Other)) bindingType=(TypeVar a)
__local_var_4_0 := __t2
_ = __local_var_4_0
// TAST (Let): v_5_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_5_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), hi_2, __local_var_4_0).IntVal)
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (v_5_3 == 1527465420) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (v_5_3 == 902936544) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (v_5_3 == 380165415) {
__t4 = __local_var_4_0
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}

func Call_Data_Ord_clamp__3109832964(low_0_loop int64, hi_1_loop int64, x_2_loop int64) int64 {
clamp__3109832964:
for {
if false { continue clamp__3109832964 }
var low_0 int64 = low_0_loop
_ = low_0
var hi_1 int64 = hi_1_loop
_ = hi_1
var x_2 int64 = x_2_loop
_ = x_2
return Call_Data_Ord_min__3052583336(hi_1, Call_Data_Ord_max__3052583336(low_0, x_2))
}
}

func Call_Data_Ord_between(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
var __t2 bool
{
var __t_tag_1 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_3, low_1).IntVal)
if (uint32(__t_tag_1) == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_3, hi_2).IntVal)
__t2 = ((uint32(__t_tag_0) == 380165415)) != (true)
}
end_branch_2:
return __t2
}

func Call_Data_Ord_abs(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictRing_1_loop *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 *Constructor_Data_Ring_Ring[gopurs_runtime.Value] = dictRing_1_loop
_ = dictRing_1
// TAST (Let): zero_2_0 shape=Other bindingType=(TypeVar a)
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_3, zero_2_0).IntVal)
if ((uint32(__t_tag_1) == 1527465420)) != (true) {
__t2 = x_3
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_1.V1), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}), "zero"), x_3)
}
end_branch_2:
return __t2
})
}

func Call_Data_Ord_abs__3883318852(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
abs__3883318852:
for {
if false { continue abs__3883318852 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Int(__eta_norm_0_0.IntVal)
}
}

func Rebox_Data_Ord_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Ord_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Ord_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Ord_1939691112_3790796878(in *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Ord_219188042_4177771502(in *Constructor_Data_Ord_Ord[bool]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ord_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ord_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Ord_3047586294_4177771502(in *Constructor_Data_Ord_Ord[float64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ord_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ord_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ord_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Ord_3790796878_3768443459(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[uint32]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Ord_4088624520_4177771502(in *Constructor_Data_Ord_Ord[[]gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ord_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Get_Data_Ord_ordArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdArrayImpl
}

func Get_Data_Ord_ordBooleanImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdBooleanImpl
}

func Get_Data_Ord_ordCharImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdCharImpl
}

func Get_Data_Ord_ordIntImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdIntImpl
}

func Get_Data_Ord_ordNumberImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdNumberImpl
}

func Get_Data_Ord_ordStringImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdStringImpl
}
