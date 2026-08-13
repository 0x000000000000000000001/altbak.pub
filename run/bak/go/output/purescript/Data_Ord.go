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
		cache_Data_Ord_eqRec = gopurs_runtime.Apply(Get_Data_Eq_eqRec(), gopurs_runtime.Value{})
	})
	return cache_Data_Ord_eqRec
}

var cache_Data_Ord_OrdRecord_dollarDict gopurs_runtime.Value
var once_Data_Ord_OrdRecord_dollarDict sync.Once
func Get_Data_Ord_OrdRecord_dollarDict() gopurs_runtime.Value {
	once_Data_Ord_OrdRecord_dollarDict.Do(func() {
		cache_Data_Ord_OrdRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_OrdRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_Ord_OrdRecord_dollarDict
}

var cache_Data_Ord_Ord_dollarDict gopurs_runtime.Value
var once_Data_Ord_Ord_dollarDict sync.Once
func Get_Data_Ord_Ord_dollarDict() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollarDict.Do(func() {
		cache_Data_Ord_Ord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Ord_dollarDict(x_0_box)
})
	})
	return cache_Data_Ord_Ord_dollarDict
}

var cache_Data_Ord_Ord1_dollarDict gopurs_runtime.Value
var once_Data_Ord_Ord1_dollarDict sync.Once
func Get_Data_Ord_Ord1_dollarDict() gopurs_runtime.Value {
	once_Data_Ord_Ord1_dollarDict.Do(func() {
		cache_Data_Ord_Ord1_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Ord1_dollarDict(x_0_box)
})
	})
	return cache_Data_Ord_Ord1_dollarDict
}

var cache_Data_Ord_ordVoid gopurs_runtime.Value
var once_Data_Ord_ordVoid sync.Once
func Get_Data_Ord_ordVoid() gopurs_runtime.Value {
	once_Data_Ord_ordVoid.Do(func() {
		cache_Data_Ord_ordVoid = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqVoid()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Ord_ordVoid
}

var cache_Data_Ord_ordUnit gopurs_runtime.Value
var once_Data_Ord_ordUnit sync.Once
func Get_Data_Ord_ordUnit() gopurs_runtime.Value {
	once_Data_Ord_ordUnit.Do(func() {
		cache_Data_Ord_ordUnit = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqUnit()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Ord_ordUnit
}

var cache_Data_Ord_ordString gopurs_runtime.Value
var once_Data_Ord_ordString sync.Once
func Get_Data_Ord_ordString() gopurs_runtime.Value {
	once_Data_Ord_ordString.Do(func() {
		cache_Data_Ord_ordString = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqString()))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}
	})
	return cache_Data_Ord_ordString
}

var cache_Data_Ord_ordRecordNil gopurs_runtime.Value
var once_Data_Ord_ordRecordNil sync.Once
func Get_Data_Ord_ordRecordNil() gopurs_runtime.Value {
	once_Data_Ord_ordRecordNil.Do(func() {
		cache_Data_Ord_ordRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_OrdRecord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord](Get_Data_Eq_eqRowNil()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
})
})})}
	})
	return cache_Data_Ord_ordRecordNil
}

var cache_Data_Ord_ordProxy gopurs_runtime.Value
var once_Data_Ord_ordProxy sync.Once
func Get_Data_Ord_ordProxy() gopurs_runtime.Value {
	once_Data_Ord_ordProxy.Do(func() {
		cache_Data_Ord_ordProxy = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqProxy()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Ord_ordProxy
}

var cache_Data_Ord_ordOrdering gopurs_runtime.Value
var once_Data_Ord_ordOrdering sync.Once
func Get_Data_Ord_ordOrdering() gopurs_runtime.Value {
	once_Data_Ord_ordOrdering.Do(func() {
		cache_Data_Ord_ordOrdering = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Ordering_eqOrdering()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 uint32
{
if (uint32(v_0.IntVal) == 1527465420) {
var __t0 uint32
{
if (uint32(v1_1.IntVal) == 1527465420) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t3 = __t0
goto end_branch_3
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
var __t1 uint32
{
if (uint32(v1_1.IntVal) == 902936544) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
if (uint32(v1_1.IntVal) == 1527465420) {
__t1 = 380165415
goto end_branch_1
} else {

}
}
{
if (uint32(v1_1.IntVal) == 380165415) {
__t1 = 1527465420
goto end_branch_1
} else {

}
}
{
__t1 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_1:
__t3 = __t1
goto end_branch_3
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
var __t2 uint32
{
if (uint32(v1_1.IntVal) == 380165415) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 380165415
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Ord_ordOrdering
}

var cache_Data_Ord_ordNumber gopurs_runtime.Value
var once_Data_Ord_ordNumber sync.Once
func Get_Data_Ord_ordNumber() gopurs_runtime.Value {
	once_Data_Ord_ordNumber.Do(func() {
		cache_Data_Ord_ordNumber = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqNumber()))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}
	})
	return cache_Data_Ord_ordNumber
}

var cache_Data_Ord_ordInt gopurs_runtime.Value
var once_Data_Ord_ordInt sync.Once
func Get_Data_Ord_ordInt() gopurs_runtime.Value {
	once_Data_Ord_ordInt.Do(func() {
		cache_Data_Ord_ordInt = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}
	})
	return cache_Data_Ord_ordInt
}

var cache_Data_Ord_ordChar gopurs_runtime.Value
var once_Data_Ord_ordChar sync.Once
func Get_Data_Ord_ordChar() gopurs_runtime.Value {
	once_Data_Ord_ordChar.Do(func() {
		cache_Data_Ord_ordChar = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqChar()))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}
	})
	return cache_Data_Ord_ordChar
}

var cache_Data_Ord_ordBoolean gopurs_runtime.Value
var once_Data_Ord_ordBoolean sync.Once
func Get_Data_Ord_ordBoolean() gopurs_runtime.Value {
	once_Data_Ord_ordBoolean.Do(func() {
		cache_Data_Ord_ordBoolean = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqBoolean()))}
}), gopurs_runtime.Apply3(Get_Data_Ord_ordBooleanImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}
	})
	return cache_Data_Ord_ordBoolean
}

var cache_Data_Ord_compareRecord gopurs_runtime.Value
var once_Data_Ord_compareRecord sync.Once
func Get_Data_Ord_compareRecord() gopurs_runtime.Value {
	once_Data_Ord_compareRecord.Do(func() {
		cache_Data_Ord_compareRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compareRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord](dict_0_box))
})
	})
	return cache_Data_Ord_compareRecord
}

var cache_Data_Ord_ordRecord gopurs_runtime.Value
var once_Data_Ord_ordRecord sync.Once
func Get_Data_Ord_ordRecord() gopurs_runtime.Value {
	once_Data_Ord_ordRecord.Do(func() {
		cache_Data_Ord_ordRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrdRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_ordRecord(_dollar__unused_0_box, dictOrdRecord_1_box)
})
	})
	return cache_Data_Ord_ordRecord
}

var cache_Data_Ord_compare1 gopurs_runtime.Value
var once_Data_Ord_compare1 sync.Once
func Get_Data_Ord_compare1() gopurs_runtime.Value {
	once_Data_Ord_compare1.Do(func() {
		cache_Data_Ord_compare1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1](dict_0_box))
})
	})
	return cache_Data_Ord_compare1
}

var cache_Data_Ord_compare gopurs_runtime.Value
var once_Data_Ord_compare sync.Once
func Get_Data_Ord_compare() gopurs_runtime.Value {
	once_Data_Ord_compare.Do(func() {
		cache_Data_Ord_compare = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare
}

var cache_Data_Ord_comparing gopurs_runtime.Value
var once_Data_Ord_comparing sync.Once
func Get_Data_Ord_comparing() gopurs_runtime.Value {
	once_Data_Ord_comparing.Do(func() {
		cache_Data_Ord_comparing = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_comparing(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_comparing
}

var cache_Data_Ord_greaterThan gopurs_runtime.Value
var once_Data_Ord_greaterThan sync.Once
func Get_Data_Ord_greaterThan() gopurs_runtime.Value {
	once_Data_Ord_greaterThan.Do(func() {
		cache_Data_Ord_greaterThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_greaterThan
}

var cache_Data_Ord_greaterThanOrEq gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq sync.Once
func Get_Data_Ord_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq.Do(func() {
		cache_Data_Ord_greaterThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_greaterThanOrEq
}

var cache_Data_Ord_lessThan gopurs_runtime.Value
var once_Data_Ord_lessThan sync.Once
func Get_Data_Ord_lessThan() gopurs_runtime.Value {
	once_Data_Ord_lessThan.Do(func() {
		cache_Data_Ord_lessThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_lessThan
}

var cache_Data_Ord_signum gopurs_runtime.Value
var once_Data_Ord_signum sync.Once
func Get_Data_Ord_signum() gopurs_runtime.Value {
	once_Data_Ord_signum.Do(func() {
		cache_Data_Ord_signum = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_signum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dictRing_1_box))
})
	})
	return cache_Data_Ord_signum
}

var cache_Data_Ord_lessThanOrEq gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq sync.Once
func Get_Data_Ord_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq.Do(func() {
		cache_Data_Ord_lessThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_lessThanOrEq
}

var cache_Data_Ord_max gopurs_runtime.Value
var once_Data_Ord_max sync.Once
func Get_Data_Ord_max() gopurs_runtime.Value {
	once_Data_Ord_max.Do(func() {
		cache_Data_Ord_max = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_max(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_Data_Ord_max
}

var cache_Data_Ord_min gopurs_runtime.Value
var once_Data_Ord_min sync.Once
func Get_Data_Ord_min() gopurs_runtime.Value {
	once_Data_Ord_min.Do(func() {
		cache_Data_Ord_min = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_min(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_Data_Ord_min
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
		cache_Data_Ord_ord1Array = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Eq_eq1Array()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_Data_Ord_compare__372254389(gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Data_Ord_ordArrayImpl(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_0 -> gopurs_runtime.Value
v_5_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_3, y_4)
_ = v_5_0
var __t1 int64
{
if (uint32(v_5_0.IntVal) == 902936544) {
__t1 = 0
goto end_branch_1
} else {

}
}
{
if (uint32(v_5_0.IntVal) == 1527465420) {
__t1 = 1
goto end_branch_1
} else {

}
}
{
if (uint32(v_5_0.IntVal) == 380165415) {
__t1 = -1
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_1:
return gopurs_runtime.Int(__t1)
})
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
				}())).IntVal)).IntVal)), UnsafePtr: nil}
})
})
})})}
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
return Call_Data_Ord_clamp(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), low_1_box, hi_2_box, x_3_box)
})
	})
	return cache_Data_Ord_clamp
}

var cache_Data_Ord_between gopurs_runtime.Value
var once_Data_Ord_between sync.Once
func Get_Data_Ord_between() gopurs_runtime.Value {
	once_Data_Ord_between.Do(func() {
		cache_Data_Ord_between = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_between(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), low_1_box, hi_2_box, x_3_box))
})
	})
	return cache_Data_Ord_between
}

var cache_Data_Ord_abs gopurs_runtime.Value
var once_Data_Ord_abs sync.Once
func Get_Data_Ord_abs() gopurs_runtime.Value {
	once_Data_Ord_abs.Do(func() {
		cache_Data_Ord_abs = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_abs(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dictRing_1_box))
})
	})
	return cache_Data_Ord_abs
}

var cache_Data_Ord_abs__1599282999 gopurs_runtime.Value
var once_Data_Ord_abs__1599282999 sync.Once
func Get_Data_Ord_abs__1599282999() gopurs_runtime.Value {
	once_Data_Ord_abs__1599282999.Do(func() {
		cache_Data_Ord_abs__1599282999 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_abs__1599282999(__eta0_0_box)
})
	})
	return cache_Data_Ord_abs__1599282999
}

var cache_Data_Ord_abs__2515802711 gopurs_runtime.Value
var once_Data_Ord_abs__2515802711 sync.Once
func Get_Data_Ord_abs__2515802711() gopurs_runtime.Value {
	once_Data_Ord_abs__2515802711.Do(func() {
		cache_Data_Ord_abs__2515802711 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_abs__2515802711(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dictRing_1_box))
})
	})
	return cache_Data_Ord_abs__2515802711
}

var cache_Data_Ord_clamp__1512183668 gopurs_runtime.Value
var once_Data_Ord_clamp__1512183668 sync.Once
func Get_Data_Ord_clamp__1512183668() gopurs_runtime.Value {
	once_Data_Ord_clamp__1512183668.Do(func() {
		cache_Data_Ord_clamp__1512183668 = gopurs_runtime.Func3(func(low_0_box gopurs_runtime.Value, hi_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Ord_clamp__1512183668(low_0_box.IntVal, hi_1_box.IntVal, x_2_box.IntVal))
})
	})
	return cache_Data_Ord_clamp__1512183668
}

var cache_Data_Ord_clamp__709576177 gopurs_runtime.Value
var once_Data_Ord_clamp__709576177 sync.Once
func Get_Data_Ord_clamp__709576177() gopurs_runtime.Value {
	once_Data_Ord_clamp__709576177.Do(func() {
		cache_Data_Ord_clamp__709576177 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_clamp__709576177(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), low_1_box, hi_2_box, x_3_box)
})
	})
	return cache_Data_Ord_clamp__709576177
}

var cache_Data_Ord_compare__1787266401 gopurs_runtime.Value
var once_Data_Ord_compare__1787266401 sync.Once
func Get_Data_Ord_compare__1787266401() gopurs_runtime.Value {
	once_Data_Ord_compare__1787266401.Do(func() {
		cache_Data_Ord_compare__1787266401 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__1787266401(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__1787266401
}

var cache_Data_Ord_compare__2286295841 gopurs_runtime.Value
var once_Data_Ord_compare__2286295841 sync.Once
func Get_Data_Ord_compare__2286295841() gopurs_runtime.Value {
	once_Data_Ord_compare__2286295841.Do(func() {
		cache_Data_Ord_compare__2286295841 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__2286295841(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__2286295841
}

var cache_Data_Ord_compare__669572705 gopurs_runtime.Value
var once_Data_Ord_compare__669572705 sync.Once
func Get_Data_Ord_compare__669572705() gopurs_runtime.Value {
	once_Data_Ord_compare__669572705.Do(func() {
		cache_Data_Ord_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__669572705(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__669572705
}

var cache_Data_Ord_compare__1110679617 gopurs_runtime.Value
var once_Data_Ord_compare__1110679617 sync.Once
func Get_Data_Ord_compare__1110679617() gopurs_runtime.Value {
	once_Data_Ord_compare__1110679617.Do(func() {
		cache_Data_Ord_compare__1110679617 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__1110679617(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__1110679617
}

var cache_Data_Ord_compare__45059489 gopurs_runtime.Value
var once_Data_Ord_compare__45059489 sync.Once
func Get_Data_Ord_compare__45059489() gopurs_runtime.Value {
	once_Data_Ord_compare__45059489.Do(func() {
		cache_Data_Ord_compare__45059489 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__45059489(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__45059489
}

var cache_Data_Ord_compare__2790853377 gopurs_runtime.Value
var once_Data_Ord_compare__2790853377 sync.Once
func Get_Data_Ord_compare__2790853377() gopurs_runtime.Value {
	once_Data_Ord_compare__2790853377.Do(func() {
		cache_Data_Ord_compare__2790853377 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__2790853377(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__2790853377
}

var cache_Data_Ord_compare__3635905793 gopurs_runtime.Value
var once_Data_Ord_compare__3635905793 sync.Once
func Get_Data_Ord_compare__3635905793() gopurs_runtime.Value {
	once_Data_Ord_compare__3635905793.Do(func() {
		cache_Data_Ord_compare__3635905793 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3635905793(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__3635905793
}

var cache_Data_Ord_compare__372254389 gopurs_runtime.Value
var once_Data_Ord_compare__372254389 sync.Once
func Get_Data_Ord_compare__372254389() gopurs_runtime.Value {
	once_Data_Ord_compare__372254389.Do(func() {
		cache_Data_Ord_compare__372254389 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__372254389(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Ord_compare__372254389
}

var cache_Data_Ord_compare__882312371 gopurs_runtime.Value
var once_Data_Ord_compare__882312371 sync.Once
func Get_Data_Ord_compare__882312371() gopurs_runtime.Value {
	once_Data_Ord_compare__882312371.Do(func() {
		cache_Data_Ord_compare__882312371 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__882312371(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Ord_compare__882312371
}

var cache_Data_Ord_compare__472859678 gopurs_runtime.Value
var once_Data_Ord_compare__472859678 sync.Once
func Get_Data_Ord_compare__472859678() gopurs_runtime.Value {
	once_Data_Ord_compare__472859678.Do(func() {
		cache_Data_Ord_compare__472859678 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__472859678(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__472859678
}

var cache_Data_Ord_compare__821463600 gopurs_runtime.Value
var once_Data_Ord_compare__821463600 sync.Once
func Get_Data_Ord_compare__821463600() gopurs_runtime.Value {
	once_Data_Ord_compare__821463600.Do(func() {
		cache_Data_Ord_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__821463600(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__821463600
}

var cache_Data_Ord_compare__4035831926 gopurs_runtime.Value
var once_Data_Ord_compare__4035831926 sync.Once
func Get_Data_Ord_compare__4035831926() gopurs_runtime.Value {
	once_Data_Ord_compare__4035831926.Do(func() {
		cache_Data_Ord_compare__4035831926 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__4035831926(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__4035831926
}

var cache_Data_Ord_compare__696857420 gopurs_runtime.Value
var once_Data_Ord_compare__696857420 sync.Once
func Get_Data_Ord_compare__696857420() gopurs_runtime.Value {
	once_Data_Ord_compare__696857420.Do(func() {
		cache_Data_Ord_compare__696857420 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__696857420(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_compare__696857420
}

var cache_Data_Ord_compare__146529112 gopurs_runtime.Value
var once_Data_Ord_compare__146529112 sync.Once
func Get_Data_Ord_compare__146529112() gopurs_runtime.Value {
	once_Data_Ord_compare__146529112.Do(func() {
		cache_Data_Ord_compare__146529112 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__146529112(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](x_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](y_1_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_compare__146529112
}

var cache_Data_Ord_compare__3077449111 gopurs_runtime.Value
var once_Data_Ord_compare__3077449111 sync.Once
func Get_Data_Ord_compare__3077449111() gopurs_runtime.Value {
	once_Data_Ord_compare__3077449111.Do(func() {
		cache_Data_Ord_compare__3077449111 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__3077449111(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_compare__3077449111
}

var cache_Data_Ord_compare__738396984 gopurs_runtime.Value
var once_Data_Ord_compare__738396984 sync.Once
func Get_Data_Ord_compare__738396984() gopurs_runtime.Value {
	once_Data_Ord_compare__738396984.Do(func() {
		cache_Data_Ord_compare__738396984 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__738396984(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__738396984
}

var cache_Data_Ord_compare__2349537221 gopurs_runtime.Value
var once_Data_Ord_compare__2349537221 sync.Once
func Get_Data_Ord_compare__2349537221() gopurs_runtime.Value {
	once_Data_Ord_compare__2349537221.Do(func() {
		cache_Data_Ord_compare__2349537221 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__2349537221(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__2349537221
}

var cache_Data_Ord_compare__3029065925 gopurs_runtime.Value
var once_Data_Ord_compare__3029065925 sync.Once
func Get_Data_Ord_compare__3029065925() gopurs_runtime.Value {
	once_Data_Ord_compare__3029065925.Do(func() {
		cache_Data_Ord_compare__3029065925 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3029065925(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__3029065925
}

var cache_Data_Ord_compare__231252914 gopurs_runtime.Value
var once_Data_Ord_compare__231252914 sync.Once
func Get_Data_Ord_compare__231252914() gopurs_runtime.Value {
	once_Data_Ord_compare__231252914.Do(func() {
		cache_Data_Ord_compare__231252914 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__231252914(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Ord_compare__231252914
}

var cache_Data_Ord_compare__2802126154 gopurs_runtime.Value
var once_Data_Ord_compare__2802126154 sync.Once
func Get_Data_Ord_compare__2802126154() gopurs_runtime.Value {
	once_Data_Ord_compare__2802126154.Do(func() {
		cache_Data_Ord_compare__2802126154 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__2802126154(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__2802126154
}

var cache_Data_Ord_compare__2740609364 gopurs_runtime.Value
var once_Data_Ord_compare__2740609364 sync.Once
func Get_Data_Ord_compare__2740609364() gopurs_runtime.Value {
	once_Data_Ord_compare__2740609364.Do(func() {
		cache_Data_Ord_compare__2740609364 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__2740609364(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Ord_compare__2740609364
}

var cache_Data_Ord_compare__1746579729 gopurs_runtime.Value
var once_Data_Ord_compare__1746579729 sync.Once
func Get_Data_Ord_compare__1746579729() gopurs_runtime.Value {
	once_Data_Ord_compare__1746579729.Do(func() {
		cache_Data_Ord_compare__1746579729 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__1746579729(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__1746579729
}

var cache_Data_Ord_compare__463614392 gopurs_runtime.Value
var once_Data_Ord_compare__463614392 sync.Once
func Get_Data_Ord_compare__463614392() gopurs_runtime.Value {
	once_Data_Ord_compare__463614392.Do(func() {
		cache_Data_Ord_compare__463614392 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__463614392(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](x_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](y_1_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_compare__463614392
}

var cache_Data_Ord_compare__1965400253 gopurs_runtime.Value
var once_Data_Ord_compare__1965400253 sync.Once
func Get_Data_Ord_compare__1965400253() gopurs_runtime.Value {
	once_Data_Ord_compare__1965400253.Do(func() {
		cache_Data_Ord_compare__1965400253 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__1965400253(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__1965400253
}

var cache_Data_Ord_compare__2107160184 gopurs_runtime.Value
var once_Data_Ord_compare__2107160184 sync.Once
func Get_Data_Ord_compare__2107160184() gopurs_runtime.Value {
	once_Data_Ord_compare__2107160184.Do(func() {
		cache_Data_Ord_compare__2107160184 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__2107160184(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__2107160184
}

var cache_Data_Ord_compare__3175661023 gopurs_runtime.Value
var once_Data_Ord_compare__3175661023 sync.Once
func Get_Data_Ord_compare__3175661023() gopurs_runtime.Value {
	once_Data_Ord_compare__3175661023.Do(func() {
		cache_Data_Ord_compare__3175661023 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3175661023(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__3175661023
}

var cache_Data_Ord_compare__3215000822 gopurs_runtime.Value
var once_Data_Ord_compare__3215000822 sync.Once
func Get_Data_Ord_compare__3215000822() gopurs_runtime.Value {
	once_Data_Ord_compare__3215000822.Do(func() {
		cache_Data_Ord_compare__3215000822 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare__3215000822(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dict_0_box))
})
	})
	return cache_Data_Ord_compare__3215000822
}

var cache_Data_Ord_compare1__650153534 gopurs_runtime.Value
var once_Data_Ord_compare1__650153534 sync.Once
func Get_Data_Ord_compare1__650153534() gopurs_runtime.Value {
	once_Data_Ord_compare1__650153534.Do(func() {
		cache_Data_Ord_compare1__650153534 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare1__650153534(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1](dict_0_box))
})
	})
	return cache_Data_Ord_compare1__650153534
}

var cache_Data_Ord_compare1__3498430039 gopurs_runtime.Value
var once_Data_Ord_compare1__3498430039 sync.Once
func Get_Data_Ord_compare1__3498430039() gopurs_runtime.Value {
	once_Data_Ord_compare1__3498430039.Do(func() {
		cache_Data_Ord_compare1__3498430039 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare1__3498430039(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1](dict_0_box))
})
	})
	return cache_Data_Ord_compare1__3498430039
}

var cache_Data_Ord_compare1__3282065035 gopurs_runtime.Value
var once_Data_Ord_compare1__3282065035 sync.Once
func Get_Data_Ord_compare1__3282065035() gopurs_runtime.Value {
	once_Data_Ord_compare1__3282065035.Do(func() {
		cache_Data_Ord_compare1__3282065035 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare1__3282065035(dictOrd_0_box)
})
	})
	return cache_Data_Ord_compare1__3282065035
}

var cache_Data_Ord_compareRecord__1222555784 gopurs_runtime.Value
var once_Data_Ord_compareRecord__1222555784 sync.Once
func Get_Data_Ord_compareRecord__1222555784() gopurs_runtime.Value {
	once_Data_Ord_compareRecord__1222555784.Do(func() {
		cache_Data_Ord_compareRecord__1222555784 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compareRecord__1222555784(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord](dict_0_box))
})
	})
	return cache_Data_Ord_compareRecord__1222555784
}

var cache_Data_Ord_compareRecord__2984072590 gopurs_runtime.Value
var once_Data_Ord_compareRecord__2984072590 sync.Once
func Get_Data_Ord_compareRecord__2984072590() gopurs_runtime.Value {
	once_Data_Ord_compareRecord__2984072590.Do(func() {
		cache_Data_Ord_compareRecord__2984072590 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compareRecord__2984072590(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord](dict_0_box))
})
	})
	return cache_Data_Ord_compareRecord__2984072590
}

var cache_Data_Ord_comparing__3783120632 gopurs_runtime.Value
var once_Data_Ord_comparing__3783120632 sync.Once
func Get_Data_Ord_comparing__3783120632() gopurs_runtime.Value {
	once_Data_Ord_comparing__3783120632.Do(func() {
		cache_Data_Ord_comparing__3783120632 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_comparing__3783120632(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_comparing__3783120632
}

var cache_Data_Ord_comparing__1990975733 gopurs_runtime.Value
var once_Data_Ord_comparing__1990975733 sync.Once
func Get_Data_Ord_comparing__1990975733() gopurs_runtime.Value {
	once_Data_Ord_comparing__1990975733.Do(func() {
		cache_Data_Ord_comparing__1990975733 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_comparing__1990975733(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](x_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](y_3_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_comparing__1990975733
}

var cache_Data_Ord_comparing__3506074860 gopurs_runtime.Value
var once_Data_Ord_comparing__3506074860 sync.Once
func Get_Data_Ord_comparing__3506074860() gopurs_runtime.Value {
	once_Data_Ord_comparing__3506074860.Do(func() {
		cache_Data_Ord_comparing__3506074860 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_comparing__3506074860(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_comparing__3506074860
}

var cache_Data_Ord_greaterThan__3259097883 gopurs_runtime.Value
var once_Data_Ord_greaterThan__3259097883 sync.Once
func Get_Data_Ord_greaterThan__3259097883() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__3259097883.Do(func() {
		cache_Data_Ord_greaterThan__3259097883 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__3259097883(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a2_2_box)))
})
	})
	return cache_Data_Ord_greaterThan__3259097883
}

var cache_Data_Ord_greaterThan__4087042607 gopurs_runtime.Value
var once_Data_Ord_greaterThan__4087042607 sync.Once
func Get_Data_Ord_greaterThan__4087042607() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__4087042607.Do(func() {
		cache_Data_Ord_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__4087042607(a1_0_box.IntVal, a2_1_box.IntVal))
})
	})
	return cache_Data_Ord_greaterThan__4087042607
}

var cache_Data_Ord_greaterThan__1061005983 gopurs_runtime.Value
var once_Data_Ord_greaterThan__1061005983 sync.Once
func Get_Data_Ord_greaterThan__1061005983() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__1061005983.Do(func() {
		cache_Data_Ord_greaterThan__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__1061005983(a1_0_box.FloatVal(), a2_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_greaterThan__1061005983
}

var cache_Data_Ord_greaterThan__3448835524 gopurs_runtime.Value
var once_Data_Ord_greaterThan__3448835524 sync.Once
func Get_Data_Ord_greaterThan__3448835524() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__3448835524.Do(func() {
		cache_Data_Ord_greaterThan__3448835524 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__3448835524(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box.StrVal(), a2_2_box.StrVal()))
})
	})
	return cache_Data_Ord_greaterThan__3448835524
}

var cache_Data_Ord_greaterThan__1409282474 gopurs_runtime.Value
var once_Data_Ord_greaterThan__1409282474 sync.Once
func Get_Data_Ord_greaterThan__1409282474() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__1409282474.Do(func() {
		cache_Data_Ord_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_greaterThan__1409282474
}

var cache_Data_Ord_greaterThan__2157625836 gopurs_runtime.Value
var once_Data_Ord_greaterThan__2157625836 sync.Once
func Get_Data_Ord_greaterThan__2157625836() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__2157625836.Do(func() {
		cache_Data_Ord_greaterThan__2157625836 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__2157625836(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a1_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a2_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Ord_greaterThan__2157625836
}

var cache_Data_Ord_greaterThan__2400628110 gopurs_runtime.Value
var once_Data_Ord_greaterThan__2400628110 sync.Once
func Get_Data_Ord_greaterThan__2400628110() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__2400628110.Do(func() {
		cache_Data_Ord_greaterThan__2400628110 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__2400628110(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a2_1_box)))
})
	})
	return cache_Data_Ord_greaterThan__2400628110
}

var cache_Data_Ord_greaterThanOrEq__1710332219 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__1710332219 sync.Once
func Get_Data_Ord_greaterThanOrEq__1710332219() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__1710332219.Do(func() {
		cache_Data_Ord_greaterThanOrEq__1710332219 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__1710332219(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box.IntVal, a2_2_box.IntVal))
})
	})
	return cache_Data_Ord_greaterThanOrEq__1710332219
}

var cache_Data_Ord_greaterThanOrEq__4087042607 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__4087042607 sync.Once
func Get_Data_Ord_greaterThanOrEq__4087042607() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__4087042607.Do(func() {
		cache_Data_Ord_greaterThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__4087042607(a1_0_box.IntVal, a2_1_box.IntVal))
})
	})
	return cache_Data_Ord_greaterThanOrEq__4087042607
}

var cache_Data_Ord_greaterThanOrEq__1061005983 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__1061005983 sync.Once
func Get_Data_Ord_greaterThanOrEq__1061005983() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__1061005983.Do(func() {
		cache_Data_Ord_greaterThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__1061005983(a1_0_box.FloatVal(), a2_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_greaterThanOrEq__1061005983
}

var cache_Data_Ord_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__1409282474 sync.Once
func Get_Data_Ord_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__1409282474.Do(func() {
		cache_Data_Ord_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_greaterThanOrEq__1409282474
}

var cache_Data_Ord_greaterThanOrEq__2065354949 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__2065354949 sync.Once
func Get_Data_Ord_greaterThanOrEq__2065354949() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__2065354949.Do(func() {
		cache_Data_Ord_greaterThanOrEq__2065354949 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__2065354949(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_greaterThanOrEq__2065354949
}

var cache_Data_Ord_lessThan__828536027 gopurs_runtime.Value
var once_Data_Ord_lessThan__828536027 sync.Once
func Get_Data_Ord_lessThan__828536027() gopurs_runtime.Value {
	once_Data_Ord_lessThan__828536027.Do(func() {
		cache_Data_Ord_lessThan__828536027 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan__828536027(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box.StrVal(), a2_2_box.StrVal()))
})
	})
	return cache_Data_Ord_lessThan__828536027
}

var cache_Data_Ord_lessThan__4087042607 gopurs_runtime.Value
var once_Data_Ord_lessThan__4087042607 sync.Once
func Get_Data_Ord_lessThan__4087042607() gopurs_runtime.Value {
	once_Data_Ord_lessThan__4087042607.Do(func() {
		cache_Data_Ord_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan__4087042607(a1_0_box.IntVal, a2_1_box.IntVal))
})
	})
	return cache_Data_Ord_lessThan__4087042607
}

var cache_Data_Ord_lessThan__1061005983 gopurs_runtime.Value
var once_Data_Ord_lessThan__1061005983 sync.Once
func Get_Data_Ord_lessThan__1061005983() gopurs_runtime.Value {
	once_Data_Ord_lessThan__1061005983.Do(func() {
		cache_Data_Ord_lessThan__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan__1061005983(a1_0_box.FloatVal(), a2_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_lessThan__1061005983
}

var cache_Data_Ord_lessThan__1409282474 gopurs_runtime.Value
var once_Data_Ord_lessThan__1409282474 sync.Once
func Get_Data_Ord_lessThan__1409282474() gopurs_runtime.Value {
	once_Data_Ord_lessThan__1409282474.Do(func() {
		cache_Data_Ord_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThan__1409282474(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_lessThan__1409282474
}

var cache_Data_Ord_lessThanOrEq__1710332219 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__1710332219 sync.Once
func Get_Data_Ord_lessThanOrEq__1710332219() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__1710332219.Do(func() {
		cache_Data_Ord_lessThanOrEq__1710332219 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__1710332219(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box.IntVal, a2_2_box.IntVal))
})
	})
	return cache_Data_Ord_lessThanOrEq__1710332219
}

var cache_Data_Ord_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__4087042607 sync.Once
func Get_Data_Ord_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__4087042607.Do(func() {
		cache_Data_Ord_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__4087042607(a1_0_box.IntVal, a2_1_box.IntVal))
})
	})
	return cache_Data_Ord_lessThanOrEq__4087042607
}

var cache_Data_Ord_lessThanOrEq__1061005983 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__1061005983 sync.Once
func Get_Data_Ord_lessThanOrEq__1061005983() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__1061005983.Do(func() {
		cache_Data_Ord_lessThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__1061005983(a1_0_box.FloatVal(), a2_1_box.FloatVal()))
})
	})
	return cache_Data_Ord_lessThanOrEq__1061005983
}

var cache_Data_Ord_lessThanOrEq__3448835524 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__3448835524 sync.Once
func Get_Data_Ord_lessThanOrEq__3448835524() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__3448835524.Do(func() {
		cache_Data_Ord_lessThanOrEq__3448835524 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__3448835524(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box.StrVal(), a2_2_box.StrVal()))
})
	})
	return cache_Data_Ord_lessThanOrEq__3448835524
}

var cache_Data_Ord_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__1409282474 sync.Once
func Get_Data_Ord_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__1409282474.Do(func() {
		cache_Data_Ord_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_lessThanOrEq__1409282474
}

var cache_Data_Ord_lessThanOrEq__1395963554 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__1395963554 sync.Once
func Get_Data_Ord_lessThanOrEq__1395963554() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__1395963554.Do(func() {
		cache_Data_Ord_lessThanOrEq__1395963554 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__1395963554(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](a1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](a2_2_box)))
})
	})
	return cache_Data_Ord_lessThanOrEq__1395963554
}

var cache_Data_Ord_lessThanOrEq__2065354949 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__2065354949 sync.Once
func Get_Data_Ord_lessThanOrEq__2065354949() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__2065354949.Do(func() {
		cache_Data_Ord_lessThanOrEq__2065354949 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__2065354949(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_Data_Ord_lessThanOrEq__2065354949
}

var cache_Data_Ord_max__2927892844 gopurs_runtime.Value
var once_Data_Ord_max__2927892844 sync.Once
func Get_Data_Ord_max__2927892844() gopurs_runtime.Value {
	once_Data_Ord_max__2927892844.Do(func() {
		cache_Data_Ord_max__2927892844 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Ord_max__2927892844(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box.IntVal, y_2_box.IntVal))
})
	})
	return cache_Data_Ord_max__2927892844
}

var cache_Data_Ord_max__2538992856 gopurs_runtime.Value
var once_Data_Ord_max__2538992856 sync.Once
func Get_Data_Ord_max__2538992856() gopurs_runtime.Value {
	once_Data_Ord_max__2538992856.Do(func() {
		cache_Data_Ord_max__2538992856 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Ord_max__2538992856(x_0_box.IntVal, y_1_box.IntVal))
})
	})
	return cache_Data_Ord_max__2538992856
}

var cache_Data_Ord_max__2767602680 gopurs_runtime.Value
var once_Data_Ord_max__2767602680 sync.Once
func Get_Data_Ord_max__2767602680() gopurs_runtime.Value {
	once_Data_Ord_max__2767602680.Do(func() {
		cache_Data_Ord_max__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_max__2767602680(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_Data_Ord_max__2767602680
}

var cache_Data_Ord_min__2927892844 gopurs_runtime.Value
var once_Data_Ord_min__2927892844 sync.Once
func Get_Data_Ord_min__2927892844() gopurs_runtime.Value {
	once_Data_Ord_min__2927892844.Do(func() {
		cache_Data_Ord_min__2927892844 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Ord_min__2927892844(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box.IntVal, y_2_box.IntVal))
})
	})
	return cache_Data_Ord_min__2927892844
}

var cache_Data_Ord_min__2767602680 gopurs_runtime.Value
var once_Data_Ord_min__2767602680 sync.Once
func Get_Data_Ord_min__2767602680() gopurs_runtime.Value {
	once_Data_Ord_min__2767602680.Do(func() {
		cache_Data_Ord_min__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_min__2767602680(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_Data_Ord_min__2767602680
}

var cache_Data_Ord_ordProxy__1361862173 gopurs_runtime.Value
var once_Data_Ord_ordProxy__1361862173 sync.Once
func Get_Data_Ord_ordProxy__1361862173() gopurs_runtime.Value {
	once_Data_Ord_ordProxy__1361862173.Do(func() {
		cache_Data_Ord_ordProxy__1361862173 = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqProxy()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Ord_ordProxy__1361862173
}

var cache_Data_Ord_ordRecordNil__2428069645 gopurs_runtime.Value
var once_Data_Ord_ordRecordNil__2428069645 sync.Once
func Get_Data_Ord_ordRecordNil__2428069645() gopurs_runtime.Value {
	once_Data_Ord_ordRecordNil__2428069645.Do(func() {
		cache_Data_Ord_ordRecordNil__2428069645 = gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_OrdRecord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord](Get_Data_Eq_eqRowNil()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
})
})})}
	})
	return cache_Data_Ord_ordRecordNil__2428069645
}

type Constructor_Data_Ord_OrdRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4162894775] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_OrdRecord)(ptr)
		_ = c
		switch key {
		case "EqRecord0": return gopurs_runtime.Box(c.V0)
		case "compareRecord": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ord_OrdRecord: " + key)
		}
	}
}


type Constructor_Data_Ord_Ord struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1435789946] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_Ord)(ptr)
		_ = c
		switch key {
		case "Eq0": return gopurs_runtime.Box(c.V0)
		case "compare": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ord_Ord: " + key)
		}
	}
}


type Constructor_Data_Ord_Ord1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1632188299] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_Ord1)(ptr)
		_ = c
		switch key {
		case "Eq10": return gopurs_runtime.Box(c.V0)
		case "compare1": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ord_Ord1: " + key)
		}
	}
}


func Call_Data_Ord_OrdRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Ord_Ord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Ord_Ord1_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Ord_compareRecord(dict_0_loop *Constructor_Data_Ord_OrdRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_OrdRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_ordRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictOrdRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrdRecord_1 gopurs_runtime.Value = dictOrdRecord_1_loop
_ = dictOrdRecord_1
// TAST (Let): eqRec1_2_0 -> *Constructor_Data_Eq_Eq
eqRec1_2_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}
_ = eqRec1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqRec1_2_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})}
}

func Call_Data_Ord_compare1(dict_0_loop *Constructor_Data_Ord_Ord1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_comparing(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)
}

func Call_Data_Ord_greaterThan(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThanOrEq(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThan(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_signum(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictRing_1_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 *Constructor_Data_Ring_Ring = dictRing_1_loop
_ = dictRing_1
// TAST (Let): Semiring0_2_0 -> gopurs_runtime.Value
Semiring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{})
_ = Semiring0_2_0
// TAST (Let): zero_3_1 -> gopurs_runtime.Value
zero_3_1 := gopurs_runtime.RecordGet(Semiring0_2_0, "zero")
_ = zero_3_1
// TAST (Let): Semiring01_4_2 -> *Constructor_Data_Semiring_Semiring
Semiring01_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}))
_ = Semiring01_4_2
// TAST (Let): one_5_3 -> gopurs_runtime.Value
one_5_3 := gopurs_runtime.RecordGet(Semiring0_2_0, "one")
_ = one_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_6, zero_3_1)
if (uint32(__t_tag_7.IntVal) == 1527465420) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
if __t8 {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_1.V1), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}), "zero"), gopurs_runtime.Box(Semiring01_4_2.V2))
goto end_branch_9
} else {

}
}
{
var __t6 gopurs_runtime.Value
{
var __t5 bool
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_6, zero_3_1)
if (uint32(__t_tag_4.IntVal) == 380165415) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
if __t5 {
__t6 = one_5_3
goto end_branch_6
} else {

}
}
{
__t6 = x_6
}
end_branch_6:
__t9 = __t6
}
end_branch_9:
return __t9
})
}

func Call_Data_Ord_lessThanOrEq(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_max(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
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

func Call_Data_Ord_min(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
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

func Call_Data_Ord_ordArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqArray_1_0 -> *Constructor_Data_Eq_Eq
eqArray_1_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}), "eq"))}
_ = eqArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqArray_1_0)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_Data_Ord_compare__372254389(gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Data_Ord_ordArrayImpl(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_1 -> gopurs_runtime.Value
v_6_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, y_5)
_ = v_6_1
var __t2 int64
{
if (uint32(v_6_1.IntVal) == 902936544) {
__t2 = 0
goto end_branch_2
} else {

}
}
{
if (uint32(v_6_1.IntVal) == 1527465420) {
__t2 = 1
goto end_branch_2
} else {

}
}
{
if (uint32(v_6_1.IntVal) == 380165415) {
__t2 = -1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return gopurs_runtime.Int(__t2)
})
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
				}())).IntVal)).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_Ord_ordRecordCons(dictOrdRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrdRecord_0 gopurs_runtime.Value = dictOrdRecord_0_loop
_ = dictOrdRecord_0
// TAST (Let): eqRowCons_1_0 -> gopurs_runtime.Value
eqRowCons_1_0 := gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "EqRecord0"), gopurs_runtime.Value{}), gopurs_runtime.Value{})
_ = eqRowCons_1_0
return gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqRowCons1_4_1 -> gopurs_runtime.Value
eqRowCons1_4_1 := gopurs_runtime.Apply(eqRowCons_1_0, dictIsSymbol_3)
_ = eqRowCons1_4_1
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqRowCons2_6_2 -> *Constructor_Data_Eq_EqRecord
eqRowCons2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord](gopurs_runtime.Apply(eqRowCons1_4_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})))
_ = eqRowCons2_6_2
return gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_OrdRecord{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(eqRowCons2_6_2)}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_10_3 -> gopurs_runtime.Value
key_10_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_3, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_10_3
// TAST (Let): left_11_4 -> gopurs_runtime.Value
left_11_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_5, "compare"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_10_3.StrVal()), ra_8), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_10_3.StrVal()), rb_9))
_ = left_11_4
var __t9 uint32
{
var __t8 bool
{
var __t_tag_5 uint32 = uint32(left_11_4.IntVal)
if (uint32(__t_tag_5) == 1527465420) {
__t8 = false
goto end_branch_8
} else {

}
}
{
var __t_tag_6 uint32 = uint32(left_11_4.IntVal)
if (uint32(__t_tag_6) == 380165415) {
__t8 = false
goto end_branch_8
} else {

}
}
{
var __t_tag_7 uint32 = uint32(left_11_4.IntVal)
if (uint32(__t_tag_7) == 902936544) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
if (__t8) != (true) {
__t9 = uint32(left_11_4.IntVal)
goto end_branch_9
} else {

}
}
{
__t9 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrdRecord_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9).IntVal)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t9), UnsafePtr: nil}
})
})
})})}
})
})
})
}

func Call_Data_Ord_clamp(dictOrd_0_loop *Constructor_Data_Ord_Ord, low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
// TAST (Let): v_4_1 -> gopurs_runtime.Value
v_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), low_1, x_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_4_1.IntVal) == 1527465420) {
__t2 = x_3
goto end_branch_2
} else {

}
}
{
if (uint32(v_4_1.IntVal) == 902936544) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_4_1.IntVal) == 380165415) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := __t2
_ = __local_var_4_0
// TAST (Let): v_5_3 -> gopurs_runtime.Value
v_5_3 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), hi_2, __local_var_4_0)
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (uint32(v_5_3.IntVal) == 1527465420) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (uint32(v_5_3.IntVal) == 902936544) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (uint32(v_5_3.IntVal) == 380165415) {
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

func Call_Data_Ord_between(dictOrd_0_loop *Constructor_Data_Ord_Ord, low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
var __t4 bool
{
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_3, low_1)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t4 = false
goto end_branch_4
} else {

}
}
{
var __t3 bool
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_3, hi_2)
if (uint32(__t_tag_2.IntVal) == 380165415) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
if __t3 {
__t4 = false
goto end_branch_4
} else {

}
}
{
__t4 = true
}
end_branch_4:
return __t4
}

func Call_Data_Ord_abs(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictRing_1_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 *Constructor_Data_Ring_Ring = dictRing_1_loop
_ = dictRing_1
// TAST (Let): zero_2_0 -> gopurs_runtime.Value
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_3, zero_2_0)
if (uint32(__t_tag_1.IntVal) == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
if __t2 {
__t3 = x_3
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_1.V1), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}), "zero"), x_3)
}
end_branch_3:
return __t3
})
}

func Call_Data_Ord_abs__1599282999(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __t1 int64
{
var __t0 bool
{
if (gopurs_runtime.Int(__eta0_0.IntVal).IntVal) < (gopurs_runtime.Int(0).IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = __eta0_0.IntVal
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Int(-(gopurs_runtime.Int(__eta0_0.IntVal).IntVal)).IntVal
}
end_branch_1:
return gopurs_runtime.Int(__t1)
}

func Call_Data_Ord_abs__2515802711(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictRing_1_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 *Constructor_Data_Ring_Ring = dictRing_1_loop
_ = dictRing_1
// TAST (Let): zero_2_0 -> gopurs_runtime.Value
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_3, zero_2_0)
if (uint32(__t_tag_1.IntVal) == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
if __t2 {
__t3 = x_3
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_1.V1), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_1.V0), gopurs_runtime.Value{}), "zero"), x_3)
}
end_branch_3:
return __t3
})
}

func Call_Data_Ord_clamp__1512183668(low_0_loop int64, hi_1_loop int64, x_2_loop int64) int64 {
var low_0 int64 = low_0_loop
_ = low_0
var hi_1 int64 = hi_1_loop
_ = hi_1
var x_2 int64 = x_2_loop
_ = x_2
// TAST (Let): v_3_1 -> gopurs_runtime.Value
v_3_1 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(low_0), gopurs_runtime.Int(x_2))
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_3_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Int(x_2)
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Int(low_0)
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Int(low_0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
// TAST (Let): __local_var_3_0 -> int64
__local_var_3_0 := __t2.IntVal
_ = __local_var_3_0
// TAST (Let): v_4_3 -> gopurs_runtime.Value
v_4_3 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(hi_1), gopurs_runtime.Int(__local_var_3_0))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (uint32(v_4_3.IntVal) == 1527465420) {
__t4 = gopurs_runtime.Int(hi_1)
goto end_branch_4
} else {

}
}
{
if (uint32(v_4_3.IntVal) == 902936544) {
__t4 = gopurs_runtime.Int(hi_1)
goto end_branch_4
} else {

}
}
{
if (uint32(v_4_3.IntVal) == 380165415) {
__t4 = gopurs_runtime.Int(__local_var_3_0)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4.IntVal
}

func Call_Data_Ord_clamp__709576177(dictOrd_0_loop *Constructor_Data_Ord_Ord, low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
// TAST (Let): v_4_1 -> gopurs_runtime.Value
v_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), low_1, x_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_4_1.IntVal) == 1527465420) {
__t2 = x_3
goto end_branch_2
} else {

}
}
{
if (uint32(v_4_1.IntVal) == 902936544) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_4_1.IntVal) == 380165415) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := __t2
_ = __local_var_4_0
// TAST (Let): v_5_3 -> gopurs_runtime.Value
v_5_3 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), hi_2, __local_var_4_0)
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (uint32(v_5_3.IntVal) == 1527465420) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (uint32(v_5_3.IntVal) == 902936544) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (uint32(v_5_3.IntVal) == 380165415) {
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

func Call_Data_Ord_compare__1787266401(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__2286295841(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__669572705(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__1110679617(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__45059489(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__2790853377(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__3635905793(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__372254389(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta0_0, __eta1_1)
}

func Call_Data_Ord_compare__882312371(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta0_0, __eta1_1)
}

func Call_Data_Ord_compare__472859678(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__821463600(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__4035831926(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__696857420(x_0_loop uint32, y_1_loop uint32) uint32 {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t11 uint32
{
if (x_0 == 1908470532) {
var __t0 uint32
{
if (y_1 == 1908470532) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t11 = __t0
goto end_branch_11
} else {

}
}
{
if (y_1 == 1908470532) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 2455627378) {
var __t1 uint32
{
if (y_1 == 2455627378) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t11 = __t1
goto end_branch_11
} else {

}
}
{
if (y_1 == 2455627378) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 4162469099) {
var __t2 uint32
{
if (y_1 == 4162469099) {
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
if (y_1 == 4162469099) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 1692989816) {
var __t3 uint32
{
if (y_1 == 1692989816) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t11 = __t3
goto end_branch_11
} else {

}
}
{
if (y_1 == 1692989816) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 330658827) {
var __t4 uint32
{
if (y_1 == 330658827) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t11 = __t4
goto end_branch_11
} else {

}
}
{
if (y_1 == 330658827) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 4067355978) {
var __t5 uint32
{
if (y_1 == 4067355978) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t11 = __t5
goto end_branch_11
} else {

}
}
{
if (y_1 == 4067355978) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 2276710548) {
var __t6 uint32
{
if (y_1 == 2276710548) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t11 = __t6
goto end_branch_11
} else {

}
}
{
if (y_1 == 2276710548) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 243771071) {
var __t7 uint32
{
if (y_1 == 243771071) {
__t7 = 902936544
goto end_branch_7
} else {

}
}
{
__t7 = 1527465420
}
end_branch_7:
__t11 = __t7
goto end_branch_11
} else {

}
}
{
if (y_1 == 243771071) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 215731793) {
var __t8 uint32
{
if (y_1 == 215731793) {
__t8 = 902936544
goto end_branch_8
} else {

}
}
{
__t8 = 1527465420
}
end_branch_8:
__t11 = __t8
goto end_branch_11
} else {

}
}
{
if (y_1 == 215731793) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 8639228) {
var __t9 uint32
{
if (y_1 == 8639228) {
__t9 = 902936544
goto end_branch_9
} else {

}
}
{
__t9 = 1527465420
}
end_branch_9:
__t11 = __t9
goto end_branch_11
} else {

}
}
{
if (y_1 == 8639228) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if (x_0 == 49471444) {
var __t10 uint32
{
if (y_1 == 49471444) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
if (y_1 == 49471444) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if ((x_0 == 3889233761)) && ((y_1 == 3889233761)) {
__t11 = 902936544
goto end_branch_11
} else {

}
}
{
__t11 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_11:
return __t11
}

func Call_Data_Ord_compare__146529112(x_0_loop *Constructor_Data_Date_Date, y_1_loop *Constructor_Data_Date_Date) uint32 {
var x_0 *Constructor_Data_Date_Date = x_0_loop
_ = x_0
var y_1 *Constructor_Data_Date_Date = y_1_loop
_ = y_1
// TAST (Let): v_2_0 -> uint32
v_2_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Date_Component_ordYear()).V1), gopurs_runtime.Int((x_0).V0), gopurs_runtime.Int((y_1).V0)).IntVal)
_ = v_2_0
var __t3 uint32
{
if (v_2_0 == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (v_2_0 == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
// TAST (Let): v1_3_1 -> uint32
v1_3_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Date_Component_ordMonth()).V1), gopurs_runtime.Value{Type: 9, IntVal: int64((x_0).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((y_1).V1), UnsafePtr: nil}).IntVal)
_ = v1_3_1
var __t2 uint32
{
if (v1_3_1 == 1527465420) {
__t2 = 1527465420
goto end_branch_2
} else {

}
}
{
if (v1_3_1 == 380165415) {
__t2 = 380165415
goto end_branch_2
} else {

}
}
{
__t2 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Date_Component_ordDay()).V1), gopurs_runtime.Int((x_0).V2), gopurs_runtime.Int((y_1).V2)).IntVal)
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
}

func Call_Data_Ord_compare__3077449111(x_0_loop uint32, y_1_loop uint32) uint32 {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t6 uint32
{
if (x_0 == 3908053364) {
var __t0 uint32
{
if (y_1 == 3908053364) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t6 = __t0
goto end_branch_6
} else {

}
}
{
if (y_1 == 3908053364) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (x_0 == 217821258) {
var __t1 uint32
{
if (y_1 == 217821258) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t6 = __t1
goto end_branch_6
} else {

}
}
{
if (y_1 == 217821258) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (x_0 == 1292308612) {
var __t2 uint32
{
if (y_1 == 1292308612) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t6 = __t2
goto end_branch_6
} else {

}
}
{
if (y_1 == 1292308612) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (x_0 == 2311060696) {
var __t3 uint32
{
if (y_1 == 2311060696) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (y_1 == 2311060696) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (x_0 == 401302776) {
var __t4 uint32
{
if (y_1 == 401302776) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
if (y_1 == 401302776) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (x_0 == 3327533908) {
var __t5 uint32
{
if (y_1 == 3327533908) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (y_1 == 3327533908) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if ((x_0 == 3631736139)) && ((y_1 == 3631736139)) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_6:
return __t6
}

func Call_Data_Ord_compare__738396984(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__2349537221(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__3029065925(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__231252914(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Interval_Duration_ordMap()).V1), __eta0_0, __eta1_1)
}

func Call_Data_Ord_compare__2802126154(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__2740609364(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Interval_ordMaybe()).V1), __eta0_0, __eta1_1)
}

func Call_Data_Ord_compare__1746579729(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__463614392(x_0_loop *Constructor_Data_Time_Time, y_1_loop *Constructor_Data_Time_Time) uint32 {
var x_0 *Constructor_Data_Time_Time = x_0_loop
_ = x_0
var y_1 *Constructor_Data_Time_Time = y_1_loop
_ = y_1
// TAST (Let): v_2_0 -> uint32
v_2_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Time_Component_ordHour()).V1), gopurs_runtime.Int((x_0).V0), gopurs_runtime.Int((y_1).V0)).IntVal)
_ = v_2_0
var __t5 uint32
{
if (v_2_0 == 1527465420) {
__t5 = 1527465420
goto end_branch_5
} else {

}
}
{
if (v_2_0 == 380165415) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
// TAST (Let): v1_3_1 -> uint32
v1_3_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Time_Component_ordMinute()).V1), gopurs_runtime.Int((x_0).V1), gopurs_runtime.Int((y_1).V1)).IntVal)
_ = v1_3_1
var __t4 uint32
{
if (v1_3_1 == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (v1_3_1 == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
// TAST (Let): v2_4_2 -> uint32
v2_4_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Time_Component_ordSecond()).V1), gopurs_runtime.Int((x_0).V2), gopurs_runtime.Int((y_1).V2)).IntVal)
_ = v2_4_2
var __t3 uint32
{
if (v2_4_2 == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (v2_4_2 == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Time_Component_ordMillisecond()).V1), gopurs_runtime.Int((x_0).V3), gopurs_runtime.Int((y_1).V3)).IntVal)
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return __t5
}

func Call_Data_Ord_compare__1965400253(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__2107160184(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__3175661023(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare__3215000822(dict_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare1__650153534(dict_0_loop *Constructor_Data_Ord_Ord1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare1__3498430039(dict_0_loop *Constructor_Data_Ord_Ord1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Ord1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compare1__3282065035(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_Lazy_ordLazy(), dictOrd_0), "compare")
}

func Call_Data_Ord_compareRecord__1222555784(dict_0_loop *Constructor_Data_Ord_OrdRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_OrdRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_compareRecord__2984072590(dict_0_loop *Constructor_Data_Ord_OrdRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_OrdRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ord_comparing__3783120632(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Int(gopurs_runtime.Apply(f_1, x_2).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(f_1, y_3).IntVal)).IntVal)
}

func Call_Data_Ord_comparing__1990975733(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value, x_2_loop *Constructor_Data_Tuple_Tuple, y_3_loop *Constructor_Data_Tuple_Tuple) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 *Constructor_Data_Tuple_Tuple = x_2_loop
_ = x_2
var y_3 *Constructor_Data_Tuple_Tuple = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Int(gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(x_2)}).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(y_3)}).IntVal)).IntVal)
}

func Call_Data_Ord_comparing__3506074860(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)
}

func Call_Data_Ord_greaterThan__3259097883(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop *Constructor_Data_Date_Date, a2_2_loop *Constructor_Data_Date_Date) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 *Constructor_Data_Date_Date = a1_1_loop
_ = a1_1
var a2_2 *Constructor_Data_Date_Date = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(a1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(a2_2)})
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThan__4087042607(a1_0_loop int64, a2_1_loop int64) bool {
var a1_0 int64 = a1_0_loop
_ = a1_0
var a2_1 int64 = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0) > (a2_1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_Data_Ord_greaterThan__1061005983(a1_0_loop float64, a2_1_loop float64) bool {
var a1_0 float64 = a1_0_loop
_ = a1_0
var a2_1 float64 = a2_1_loop
_ = a2_1
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Time_Duration_ordMilliseconds()).V1), gopurs_runtime.Float(a1_0), gopurs_runtime.Float(a2_1))
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThan__3448835524(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop string, a2_2_loop string) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 string = a1_1_loop
_ = a1_1
var a2_2 string = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Str(a1_1), gopurs_runtime.Str(a2_2))
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThan__1409282474(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThan__2157625836(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop []gopurs_runtime.Value, a2_2_loop []gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 []gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 []gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Array(a1_1), gopurs_runtime.Array(a2_2))
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThan__2400628110(a1_0_loop *Constructor_Data_Maybe_Just, a2_1_loop *Constructor_Data_Maybe_Just) bool {
var a1_0 *Constructor_Data_Maybe_Just = a1_0_loop
_ = a1_0
var a2_1 *Constructor_Data_Maybe_Just = a2_1_loop
_ = a2_1
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Date_ordMaybe()).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a2_1)})
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThanOrEq__1710332219(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop int64, a2_2_loop int64) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 int64 = a1_1_loop
_ = a1_1
var a2_2 int64 = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Int(a1_1), gopurs_runtime.Int(a2_2))
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThanOrEq__4087042607(a1_0_loop int64, a2_1_loop int64) bool {
var a1_0 int64 = a1_0_loop
_ = a1_0
var a2_1 int64 = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0) < (a2_1) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_Data_Ord_greaterThanOrEq__1061005983(a1_0_loop float64, a2_1_loop float64) bool {
var a1_0 float64 = a1_0_loop
_ = a1_0
var a2_1 float64 = a2_1_loop
_ = a2_1
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(a1_0), gopurs_runtime.Float(a2_1))
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThanOrEq__1409282474(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_greaterThanOrEq__2065354949(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThan__828536027(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop string, a2_2_loop string) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 string = a1_1_loop
_ = a1_1
var a2_2 string = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Str(a1_1), gopurs_runtime.Str(a2_2))
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThan__4087042607(a1_0_loop int64, a2_1_loop int64) bool {
var a1_0 int64 = a1_0_loop
_ = a1_0
var a2_1 int64 = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0) < (a2_1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_Data_Ord_lessThan__1061005983(a1_0_loop float64, a2_1_loop float64) bool {
var a1_0 float64 = a1_0_loop
_ = a1_0
var a2_1 float64 = a2_1_loop
_ = a2_1
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(a1_0), gopurs_runtime.Float(a2_1))
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThan__1409282474(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThanOrEq__1710332219(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop int64, a2_2_loop int64) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 int64 = a1_1_loop
_ = a1_1
var a2_2 int64 = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Int(a1_1), gopurs_runtime.Int(a2_2))
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThanOrEq__4087042607(a1_0_loop int64, a2_1_loop int64) bool {
var a1_0 int64 = a1_0_loop
_ = a1_0
var a2_1 int64 = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0) > (a2_1) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_Data_Ord_lessThanOrEq__1061005983(a1_0_loop float64, a2_1_loop float64) bool {
var a1_0 float64 = a1_0_loop
_ = a1_0
var a2_1 float64 = a2_1_loop
_ = a2_1
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(a1_0), gopurs_runtime.Float(a2_1))
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThanOrEq__3448835524(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop string, a2_2_loop string) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 string = a1_1_loop
_ = a1_1
var a2_2 string = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Str(a1_1), gopurs_runtime.Str(a2_2))
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThanOrEq__1409282474(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThanOrEq__1395963554(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop *Constructor_Data_Tuple_Tuple, a2_2_loop *Constructor_Data_Tuple_Tuple) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 *Constructor_Data_Tuple_Tuple = a1_1_loop
_ = a1_1
var a2_2 *Constructor_Data_Tuple_Tuple = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(a1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(a2_2)})
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_lessThanOrEq__2065354949(dictOrd_0_loop *Constructor_Data_Ord_Ord, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_Data_Ord_max__2927892844(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop int64, y_2_loop int64) int64 {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 int64 = x_1_loop
_ = x_1
var y_2 int64 = y_2_loop
_ = y_2
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Int(x_1), gopurs_runtime.Int(y_2))
_ = v_3_0
var __t1 int64
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_1:
return __t1
}

func Call_Data_Ord_max__2538992856(x_0_loop int64, y_1_loop int64) int64 {
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
// TAST (Let): v_2_0 -> gopurs_runtime.Value
v_2_0 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(x_0), gopurs_runtime.Int(y_1))
_ = v_2_0
var __t1 int64
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t1 = y_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 902936544) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_1:
return __t1
}

func Call_Data_Ord_max__2767602680(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
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

func Call_Data_Ord_min__2927892844(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop int64, y_2_loop int64) int64 {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 int64 = x_1_loop
_ = x_1
var y_2 int64 = y_2_loop
_ = y_2
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Int(x_1), gopurs_runtime.Int(y_2))
_ = v_3_0
var __t1 int64
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_1:
return __t1
}

func Call_Data_Ord_min__2767602680(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
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
