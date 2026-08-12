package Data_Ord

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Symbol "gopurs/output/Data.Symbol"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_ordVoid gopurs_runtime.Value
var once_ordVoid sync.Once
func Get_ordVoid() gopurs_runtime.Value {
	once_ordVoid.Do(func() {
		cache_ordVoid = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqVoid()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
}))
	})
	return cache_ordVoid
}

var cache_ordUnit gopurs_runtime.Value
var once_ordUnit sync.Once
func Get_ordUnit() gopurs_runtime.Value {
	once_ordUnit.Do(func() {
		cache_ordUnit = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqUnit()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
}))
	})
	return cache_ordUnit
}

var cache_ordString gopurs_runtime.Value
var once_ordString sync.Once
func Get_ordString() gopurs_runtime.Value {
	once_ordString.Do(func() {
		cache_ordString = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqString()
}), gopurs_runtime.Apply3(Get_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordString
}

var cache_ordRecordNil gopurs_runtime.Value
var once_ordRecordNil sync.Once
func Get_ordRecordNil() gopurs_runtime.Value {
	once_ordRecordNil.Do(func() {
		cache_ordRecordNil = gopurs_runtime.RecordDict2("EqRecord0", "compareRecord", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqRowNil()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
})
}))
	})
	return cache_ordRecordNil
}

var cache_ordProxy gopurs_runtime.Value
var once_ordProxy sync.Once
func Get_ordProxy() gopurs_runtime.Value {
	once_ordProxy.Do(func() {
		cache_ordProxy = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
}))
	})
	return cache_ordProxy
}

var cache_ordOrdering gopurs_runtime.Value
var once_ordOrdering sync.Once
func Get_ordOrdering() gopurs_runtime.Value {
	once_ordOrdering.Do(func() {
		cache_ordOrdering = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ordering.Get_eqOrdering()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
var __t1 gopurs_runtime.Value
{
if (uint32(v1_1.IntVal) == 902936544) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (uint32(v1_1.IntVal) == 1527465420) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (uint32(v1_1.IntVal) == 380165415) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t2), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t3.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordOrdering
}

var cache_ordNumber gopurs_runtime.Value
var once_ordNumber sync.Once
func Get_ordNumber() gopurs_runtime.Value {
	once_ordNumber.Do(func() {
		cache_ordNumber = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqNumber()
}), gopurs_runtime.Apply3(Get_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordNumber
}

var cache_ordInt gopurs_runtime.Value
var once_ordInt sync.Once
func Get_ordInt() gopurs_runtime.Value {
	once_ordInt.Do(func() {
		cache_ordInt = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordInt
}

var cache_ordChar gopurs_runtime.Value
var once_ordChar sync.Once
func Get_ordChar() gopurs_runtime.Value {
	once_ordChar.Do(func() {
		cache_ordChar = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqChar()
}), gopurs_runtime.Apply3(Get_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordChar
}

var cache_ordBoolean gopurs_runtime.Value
var once_ordBoolean sync.Once
func Get_ordBoolean() gopurs_runtime.Value {
	once_ordBoolean.Do(func() {
		cache_ordBoolean = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqBoolean()
}), gopurs_runtime.Apply3(Get_ordBooleanImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordBoolean
}

var cache_compareRecord gopurs_runtime.Value
var once_compareRecord sync.Once
func Get_compareRecord() gopurs_runtime.Value {
	once_compareRecord.Do(func() {
		cache_compareRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compareRecord(gopurs_runtime.CoerceToStruct[Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compareRecord
}

var cache_ordRecord gopurs_runtime.Value
var once_ordRecord sync.Once
func Get_ordRecord() gopurs_runtime.Value {
	once_ordRecord.Do(func() {
		cache_ordRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrdRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordRecord(_dollar__unused_0_box, dictOrdRecord_1_box)
})
	})
	return cache_ordRecord
}

var cache_compare1 gopurs_runtime.Value
var once_compare1 sync.Once
func Get_compare1() gopurs_runtime.Value {
	once_compare1.Do(func() {
		cache_compare1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare1(gopurs_runtime.CoerceToStruct[Constructor_Ord1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare1
}

var cache_compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		cache_compare = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare
}

var cache_comparing gopurs_runtime.Value
var once_comparing sync.Once
func Get_comparing() gopurs_runtime.Value {
	once_comparing.Do(func() {
		cache_comparing = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_comparing(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
})
	})
	return cache_comparing
}

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan
}

var cache_greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		cache_greaterThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq
}

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan
}

var cache_signum gopurs_runtime.Value
var once_signum sync.Once
func Get_signum() gopurs_runtime.Value {
	once_signum.Do(func() {
		cache_signum = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_signum(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_1_box))
})
	})
	return cache_signum
}

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq
}

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_max
}

var cache_min gopurs_runtime.Value
var once_min sync.Once
func Get_min() gopurs_runtime.Value {
	once_min.Do(func() {
		cache_min = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_min(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_min
}

var cache_ordArray gopurs_runtime.Value
var once_ordArray sync.Once
func Get_ordArray() gopurs_runtime.Value {
	once_ordArray.Do(func() {
		cache_ordArray = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordArray(dictOrd_0_box)
})
	})
	return cache_ordArray
}

var cache_ord1Array gopurs_runtime.Value
var once_ord1Array sync.Once
func Get_ord1Array() gopurs_runtime.Value {
	once_ord1Array.Do(func() {
		cache_ord1Array = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eq1Array()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_ordArray(dictOrd_0), "compare")
}))
	})
	return cache_ord1Array
}

var cache_ordRecordCons gopurs_runtime.Value
var once_ordRecordCons sync.Once
func Get_ordRecordCons() gopurs_runtime.Value {
	once_ordRecordCons.Do(func() {
		cache_ordRecordCons = gopurs_runtime.Func(func(dictOrdRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordRecordCons(dictOrdRecord_0_box)
})
	})
	return cache_ordRecordCons
}

var cache_clamp gopurs_runtime.Value
var once_clamp sync.Once
func Get_clamp() gopurs_runtime.Value {
	once_clamp.Do(func() {
		cache_clamp = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_clamp(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), low_1_box, hi_2_box, x_3_box)
})
	})
	return cache_clamp
}

var cache_between gopurs_runtime.Value
var once_between sync.Once
func Get_between() gopurs_runtime.Value {
	once_between.Do(func() {
		cache_between = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_between(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), low_1_box, hi_2_box, x_3_box))
})
	})
	return cache_between
}

var cache_abs gopurs_runtime.Value
var once_abs sync.Once
func Get_abs() gopurs_runtime.Value {
	once_abs.Do(func() {
		cache_abs = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_abs(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_1_box))
})
	})
	return cache_abs
}

var cache_eq__2276491096 gopurs_runtime.Value
var once_eq__2276491096 sync.Once
func Get_eq__2276491096() gopurs_runtime.Value {
	once_eq__2276491096.Do(func() {
		cache_eq__2276491096 = gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq")
	})
	return cache_eq__2276491096
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq1Array__2389734302 gopurs_runtime.Value
var once_eq1Array__2389734302 sync.Once
func Get_eq1Array__2389734302() gopurs_runtime.Value {
	once_eq1Array__2389734302.Do(func() {
		cache_eq1Array__2389734302 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}))
	})
	return cache_eq1Array__2389734302
}

var cache_eqProxy__2077298405 gopurs_runtime.Value
var once_eqProxy__2077298405 sync.Once
func Get_eqProxy__2077298405() gopurs_runtime.Value {
	once_eqProxy__2077298405.Do(func() {
		cache_eqProxy__2077298405 = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_eqProxy__2077298405
}

var cache_eqRowNil__2234696885 gopurs_runtime.Value
var once_eqRowNil__2234696885 sync.Once
func Get_eqRowNil__2234696885() gopurs_runtime.Value {
	once_eqRowNil__2234696885.Do(func() {
		cache_eqRowNil__2234696885 = gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})
}))
	})
	return cache_eqRowNil__2234696885
}

var cache_notEq__2384498378 gopurs_runtime.Value
var once_notEq__2384498378 sync.Once
func Get_notEq__2384498378() gopurs_runtime.Value {
	once_notEq__2384498378.Do(func() {
		cache_notEq__2384498378 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_notEq__2384498378
}

var cache_notEq__1272715810 gopurs_runtime.Value
var once_notEq__1272715810 sync.Once
func Get_notEq__1272715810() gopurs_runtime.Value {
	once_notEq__1272715810.Do(func() {
		cache_notEq__1272715810 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__1272715810(x_0_box, y_1_box))
})
	})
	return cache_notEq__1272715810
}

var cache_compare__372254389 gopurs_runtime.Value
var once_compare__372254389 sync.Once
func Get_compare__372254389() gopurs_runtime.Value {
	once_compare__372254389.Do(func() {
		cache_compare__372254389 = gopurs_runtime.Apply3(Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})
	})
	return cache_compare__372254389
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_compareRecord__1222555784 gopurs_runtime.Value
var once_compareRecord__1222555784 sync.Once
func Get_compareRecord__1222555784() gopurs_runtime.Value {
	once_compareRecord__1222555784.Do(func() {
		cache_compareRecord__1222555784 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compareRecord__1222555784(gopurs_runtime.CoerceToStruct[Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compareRecord__1222555784
}

var cache_compareRecord__2984072590 gopurs_runtime.Value
var once_compareRecord__2984072590 sync.Once
func Get_compareRecord__2984072590() gopurs_runtime.Value {
	once_compareRecord__2984072590.Do(func() {
		cache_compareRecord__2984072590 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compareRecord__2984072590(gopurs_runtime.CoerceToStruct[Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compareRecord__2984072590
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_max__2767602680 gopurs_runtime.Value
var once_max__2767602680 sync.Once
func Get_max__2767602680() gopurs_runtime.Value {
	once_max__2767602680.Do(func() {
		cache_max__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max__2767602680(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_max__2767602680
}

var cache_min__2767602680 gopurs_runtime.Value
var once_min__2767602680 sync.Once
func Get_min__2767602680() gopurs_runtime.Value {
	once_min__2767602680.Do(func() {
		cache_min__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_min__2767602680(gopurs_runtime.CoerceToStruct[Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
})
	})
	return cache_min__2767602680
}

var cache_negate__2635823316 gopurs_runtime.Value
var once_negate__2635823316 sync.Once
func Get_negate__2635823316() gopurs_runtime.Value {
	once_negate__2635823316.Do(func() {
		cache_negate__2635823316 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2635823316(a_0_box)
})
	})
	return cache_negate__2635823316
}

var cache_negate__1364373265 gopurs_runtime.Value
var once_negate__1364373265 sync.Once
func Get_negate__1364373265() gopurs_runtime.Value {
	once_negate__1364373265.Do(func() {
		cache_negate__1364373265 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__1364373265(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate__1364373265
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_one__1204848985 gopurs_runtime.Value
var once_one__1204848985 sync.Once
func Get_one__1204848985() gopurs_runtime.Value {
	once_one__1204848985.Do(func() {
		cache_one__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one__1204848985(dict_0_box)
})
	})
	return cache_one__1204848985
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

var cache_reflectSymbol__3416619207 gopurs_runtime.Value
var once_reflectSymbol__3416619207 sync.Once
func Get_reflectSymbol__3416619207() gopurs_runtime.Value {
	once_reflectSymbol__3416619207.Do(func() {
		cache_reflectSymbol__3416619207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__3416619207(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__3416619207
}

var cache_reflectSymbol__1166932993 gopurs_runtime.Value
var once_reflectSymbol__1166932993 sync.Once
func Get_reflectSymbol__1166932993() gopurs_runtime.Value {
	once_reflectSymbol__1166932993.Do(func() {
		cache_reflectSymbol__1166932993 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__1166932993(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__1166932993
}

type Constructor_OrdRecord[T_rowlist any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4162894775] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "EqRecord0": return c.V0
		case "compareRecord": return c.V1
		default: panic("Key not found in dictionary Constructor_OrdRecord: " + key)
		}
	}
}


type Constructor_Ord[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1435789946] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Ord[gopurs_runtime.Value])(ptr)
		switch key {
		case "Eq0": return c.V0
		case "compare": return c.V1
		default: panic("Key not found in dictionary Constructor_Ord: " + key)
		}
	}
}


type Constructor_Ord1[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1632188299] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Ord1[gopurs_runtime.Value])(ptr)
		switch key {
		case "Eq10": return c.V0
		case "compare1": return c.V1
		default: panic("Key not found in dictionary Constructor_Ord1: " + key)
		}
	}
}


func Call_compareRecord(dict_0_loop *Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_ordRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictOrdRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrdRecord_1 gopurs_runtime.Value = dictOrdRecord_1_loop
_ = dictOrdRecord_1
eqRec1_2_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = eqRec1_2_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_2_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_compare1(dict_0_loop *Constructor_Ord1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Ord1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare(dict_0_loop *Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_comparing(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(dictOrd_0.V1, gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)
}

func Call_greaterThan(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
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

func Call_greaterThanOrEq(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
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

func Call_lessThan(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
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

func Call_signum(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], dictRing_1_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_1_loop
_ = dictRing_1
Semiring0_2_0 := gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{})
_ = Semiring0_2_0
zero_3_1 := gopurs_runtime.RecordGet(Semiring0_2_0, "zero")
_ = zero_3_1
Semiring01_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}))
_ = Semiring01_4_2
one_5_3 := gopurs_runtime.RecordGet(Semiring0_2_0, "one")
_ = one_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_6, zero_3_1)
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
__t9 = gopurs_runtime.Apply2(dictRing_1.V1, gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}), "zero"), Semiring01_4_2.V2)
goto end_branch_9
} else {

}
}
{
var __t6 gopurs_runtime.Value
{
var __t5 bool
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_6, zero_3_1)
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

func Call_lessThanOrEq(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
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

func Call_max(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2)
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

func Call_min(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2)
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

func Call_ordArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqArray_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}), "eq")))
_ = eqArray_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqArray_1_0
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(Get_compare__372254389(), gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_ordArrayImpl(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, y_5)
_ = v_6_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_6_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (uint32(v_6_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Int(1)
goto end_branch_2
} else {

}
}
{
if (uint32(v_6_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Int(Call_negate__2635823316(gopurs_runtime.Int(1)).IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Int(__t2.IntVal)
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
}))
}

func Call_ordRecordCons(dictOrdRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrdRecord_0 gopurs_runtime.Value = dictOrdRecord_0_loop
_ = dictOrdRecord_0
eqRowCons_1_0 := gopurs_runtime.Apply2(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "EqRecord0"), gopurs_runtime.Value{}), gopurs_runtime.Value{})
_ = eqRowCons_1_0
return gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons1_4_1 := gopurs_runtime.Apply(eqRowCons_1_0, dictIsSymbol_3)
_ = eqRowCons1_4_1
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons2_6_2 := gopurs_runtime.Apply(eqRowCons1_4_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{}))
_ = eqRowCons2_6_2
return gopurs_runtime.RecordDict2("EqRecord0", "compareRecord", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRowCons2_6_2
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
key_10_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_3, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_10_3
left_11_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_5, "compare"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Str(key_10_3.StrVal()), ra_8), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Str(key_10_3.StrVal()), rb_9))
_ = left_11_4
var __t5 uint32
{
if (gopurs_runtime.Bool(Call_notEq__1272715810(gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(left_11_4.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil})).IntVal) != (0) {
__t5 = uint32(left_11_4.IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrdRecord_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9).IntVal)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})
})
}))
})
})
})
}

func Call_clamp(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
v_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, low_1, x_3)
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
__local_var_4_0 := __t2
_ = __local_var_4_0
v_5_3 := gopurs_runtime.Apply2(dictOrd_0.V1, hi_2, __local_var_4_0)
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

func Call_between(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
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
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_3, low_1)
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
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_3, hi_2)
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

func Call_abs(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], dictRing_1_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_1_loop
_ = dictRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_3, zero_2_0)
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
__t3 = gopurs_runtime.Apply2(dictRing_1.V1, gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}), "zero"), x_3)
}
end_branch_3:
return __t3
})
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_notEq__2384498378(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, x_1, y_2).IntVal) != (0)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_notEq__1272715810(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), x_0, y_1).IntVal) != (0)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_compare__821463600(dict_0_loop *Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compareRecord__1222555784(dict_0_loop *Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compareRecord__2984072590(dict_0_loop *Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__1409282474(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
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

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
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

func Call_lessThan__1409282474(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
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

func Call_max__2767602680(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2)
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

func Call_min__2767602680(dictOrd_0_loop *Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2)
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

func Call_negate__2635823316(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Int(-(a_0.IntVal))
}

func Call_negate__1364373265(dictRing_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_one__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_reflectSymbol__3416619207(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_reflectSymbol__1166932993(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_ordArrayImpl() gopurs_runtime.Value {
	return _Gopurs_OrdArrayImpl
}

func Get_ordBooleanImpl() gopurs_runtime.Value {
	return _Gopurs_OrdBooleanImpl
}

func Get_ordCharImpl() gopurs_runtime.Value {
	return _Gopurs_OrdCharImpl
}

func Get_ordIntImpl() gopurs_runtime.Value {
	return _Gopurs_OrdIntImpl
}

func Get_ordNumberImpl() gopurs_runtime.Value {
	return _Gopurs_OrdNumberImpl
}

func Get_ordStringImpl() gopurs_runtime.Value {
	return _Gopurs_OrdStringImpl
}
