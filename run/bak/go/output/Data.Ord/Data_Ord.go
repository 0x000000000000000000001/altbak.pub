package Data_Ord

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Type_Proxy "gopurs/output/Type.Proxy"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	unsafe "unsafe"
)

var ordVoid gopurs_runtime.Value
var once_ordVoid sync.Once
func Get_ordVoid() gopurs_runtime.Value {
	once_ordVoid.Do(func() {
		ordVoid = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqVoid()
}))
	})
	return ordVoid
}

var ordUnit gopurs_runtime.Value
var once_ordUnit sync.Once
func Get_ordUnit() gopurs_runtime.Value {
	once_ordUnit.Do(func() {
		ordUnit = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqUnit()
}))
	})
	return ordUnit
}

var ordString gopurs_runtime.Value
var once_ordString sync.Once
func Get_ordString() gopurs_runtime.Value {
	once_ordString.Do(func() {
		ordString = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}, gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}, gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqString()
}))
	})
	return ordString
}

var ordRecordNil gopurs_runtime.Value
var once_ordRecordNil sync.Once
func Get_ordRecordNil() gopurs_runtime.Value {
	once_ordRecordNil.Do(func() {
		ordRecordNil = gopurs_runtime.RecordDict2("compareRecord", "EqRecord0", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqRowNil()
}))
	})
	return ordRecordNil
}

var ordProxy gopurs_runtime.Value
var once_ordProxy sync.Once
func Get_ordProxy() gopurs_runtime.Value {
	once_ordProxy.Do(func() {
		ordProxy = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqProxy()
}))
	})
	return ordProxy
}

var ordOrdering gopurs_runtime.Value
var once_ordOrdering sync.Once
func Get_ordOrdering() gopurs_runtime.Value {
	once_ordOrdering.Do(func() {
		ordOrdering = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3866105248) {
var __t1 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3866105248) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1111389260) {
var __t2 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 1111389260) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_2
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 3866105248) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_2
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2098047435) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2098047435) {
var __t3 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 2098047435) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ordering.Get_eqOrdering()
}))
	})
	return ordOrdering
}

var ordNumber gopurs_runtime.Value
var once_ordNumber sync.Once
func Get_ordNumber() gopurs_runtime.Value {
	once_ordNumber.Do(func() {
		ordNumber = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}, gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}, gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqNumber()
}))
	})
	return ordNumber
}

var ordInt gopurs_runtime.Value
var once_ordInt sync.Once
func Get_ordInt() gopurs_runtime.Value {
	once_ordInt.Do(func() {
		ordInt = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}, gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}, gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqInt()
}))
	})
	return ordInt
}

var ordChar gopurs_runtime.Value
var once_ordChar sync.Once
func Get_ordChar() gopurs_runtime.Value {
	once_ordChar.Do(func() {
		ordChar = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}, gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}, gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqChar()
}))
	})
	return ordChar
}

var ordBoolean gopurs_runtime.Value
var once_ordBoolean sync.Once
func Get_ordBoolean() gopurs_runtime.Value {
	once_ordBoolean.Do(func() {
		ordBoolean = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordBooleanImpl(), gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}, gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}, gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqBoolean()
}))
	})
	return ordBoolean
}

var compareRecord gopurs_runtime.Value
var once_compareRecord sync.Once
func Get_compareRecord() gopurs_runtime.Value {
	once_compareRecord.Do(func() {
		compareRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compareRecord")
}()
})
	})
	return compareRecord
}

var ordRecord gopurs_runtime.Value
var once_ordRecord sync.Once
func Get_ordRecord() gopurs_runtime.Value {
	once_ordRecord.Do(func() {
		ordRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrdRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordRecord(_dollar__unused_0_box, dictOrdRecord_1_box)
})
	})
	return ordRecord
}

var compare1 gopurs_runtime.Value
var once_compare1 sync.Once
func Get_compare1() gopurs_runtime.Value {
	once_compare1.Do(func() {
		compare1 = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compare1")
}()
})
	})
	return compare1
}

var compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		compare = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compare")
}()
})
	})
	return compare
}

var comparing gopurs_runtime.Value
var once_comparing sync.Once
func Get_comparing() gopurs_runtime.Value {
	once_comparing.Do(func() {
		comparing = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comparing(dictOrd_0_box, f_1_box, x_2_box, y_3_box)
})
	})
	return comparing
}

var greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		greaterThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan(dictOrd_0_box, a1_1_box, a2_2_box))
})
	})
	return greaterThan
}

var greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		greaterThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq(dictOrd_0_box, a1_1_box, a2_2_box))
})
	})
	return greaterThanOrEq
}

var lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		lessThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan(dictOrd_0_box, a1_1_box, a2_2_box))
})
	})
	return lessThan
}

var signum gopurs_runtime.Value
var once_signum sync.Once
func Get_signum() gopurs_runtime.Value {
	once_signum.Do(func() {
		signum = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_signum(dictOrd_0_box, dictRing_1_box)
})
	})
	return signum
}

var lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		lessThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq(dictOrd_0_box, a1_1_box, a2_2_box))
})
	})
	return lessThanOrEq
}

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max(dictOrd_0_box, x_1_box, y_2_box)
})
	})
	return max
}

var min gopurs_runtime.Value
var once_min sync.Once
func Get_min() gopurs_runtime.Value {
	once_min.Do(func() {
		min = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_min(dictOrd_0_box, x_1_box, y_2_box)
})
	})
	return min
}

var ordArray gopurs_runtime.Value
var once_ordArray sync.Once
func Get_ordArray() gopurs_runtime.Value {
	once_ordArray.Do(func() {
		ordArray = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqArray_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}), "eq")))
_ = eqArray_1_0
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_ordInt(), "compare"), gopurs_runtime.Int(0), gopurs_runtime.Apply3(Get_ordArrayImpl(), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, y_5)
_ = v_6_1
var __t2 gopurs_runtime.Value
{
if (v_6_1.Type == 9 && v_6_1.IntVal == 1111389260) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (v_6_1.Type == 9 && v_6_1.IntVal == 3866105248) {
__t2 = gopurs_runtime.Int(1)
goto end_branch_2
} else {

}
}
{
if (v_6_1.Type == 9 && v_6_1.IntVal == 2098047435) {
__t2 = gopurs_runtime.Int(-1)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), xs_2, ys_3))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqArray_1_0
}))
}()
})
	})
	return ordArray
}

var ord1Array gopurs_runtime.Value
var once_ord1Array sync.Once
func Get_ord1Array() gopurs_runtime.Value {
	once_ord1Array.Do(func() {
		ord1Array = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_ordArray(), dictOrd_0), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eq1Array()
}))
	})
	return ord1Array
}

var ordRecordCons gopurs_runtime.Value
var once_ordRecordCons sync.Once
func Get_ordRecordCons() gopurs_runtime.Value {
	once_ordRecordCons.Do(func() {
		ordRecordCons = gopurs_runtime.Func(func(dictOrdRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrdRecord_0 gopurs_runtime.Value = dictOrdRecord_0_loop
_ = dictOrdRecord_0
eqRowCons_1_0 := gopurs_runtime.Apply2(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "EqRecord0"), gopurs_runtime.Value{}), gopurs_runtime.Value{})
_ = eqRowCons_1_0
return gopurs_runtime.Func2(func(_dollar__unused_2 gopurs_runtime.Value, dictIsSymbol_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons1_4_1 := gopurs_runtime.Apply(eqRowCons_1_0, dictIsSymbol_3)
_ = eqRowCons1_4_1
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons2_6_2 := gopurs_runtime.Apply(eqRowCons1_4_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{}))
_ = eqRowCons2_6_2
return gopurs_runtime.RecordDict2("compareRecord", "EqRecord0", gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, ra_8 gopurs_runtime.Value, rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
key_10_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_3, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 3178699476, UnsafePtr: unsafe.Pointer(&pkg_Type_Proxy.Data_Type_Proxy_Proxy{})})
_ = key_10_3
left_11_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_5, "compare"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_10_3, ra_8), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_10_3, rb_9))
_ = left_11_4
var __t5 gopurs_runtime.Value
{
if (left_11_4.Type == 9 && left_11_4.IntVal == 3866105248) || (left_11_4.Type == 9 && left_11_4.IntVal == 2098047435) || (left_11_4.Type == 9 && left_11_4.IntVal == 1111389260) != true {
__t5 = left_11_4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrdRecord_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: 3178699476, UnsafePtr: unsafe.Pointer(&pkg_Type_Proxy.Data_Type_Proxy_Proxy{})}, ra_8, rb_9)
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRowCons2_6_2
}))
})
})
}()
})
	})
	return ordRecordCons
}

var clamp gopurs_runtime.Value
var once_clamp sync.Once
func Get_clamp() gopurs_runtime.Value {
	once_clamp.Do(func() {
		clamp = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_clamp(dictOrd_0_box, low_1_box, hi_2_box, x_3_box)
})
	})
	return clamp
}

var between gopurs_runtime.Value
var once_between sync.Once
func Get_between() gopurs_runtime.Value {
	once_between.Do(func() {
		between = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_between(dictOrd_0_box, low_1_box, hi_2_box, x_3_box)
})
	})
	return between
}

var abs gopurs_runtime.Value
var once_abs sync.Once
func Get_abs() gopurs_runtime.Value {
	once_abs.Do(func() {
		abs = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_abs(dictOrd_0_box, dictRing_1_box)
})
	})
	return abs
}

func Call_ordRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictOrdRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrdRecord_1 gopurs_runtime.Value = dictOrdRecord_1_loop
_ = dictOrdRecord_1
eqRec1_2_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: 3178699476, UnsafePtr: unsafe.Pointer(&pkg_Type_Proxy.Data_Type_Proxy_Proxy{})}))
_ = eqRec1_2_0
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: 3178699476, UnsafePtr: unsafe.Pointer(&pkg_Type_Proxy.Data_Type_Proxy_Proxy{})}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_2_0
}))
}

func Call_comparing(dictOrd_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3))
}

func Call_greaterThan(dictOrd_0_loop gopurs_runtime.Value, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2).IntVal == 2098047435)
}

func Call_greaterThanOrEq(dictOrd_0_loop gopurs_runtime.Value, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2).IntVal == 3866105248) != true
}

func Call_lessThan(dictOrd_0_loop gopurs_runtime.Value, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2).IntVal == 3866105248)
}

func Call_signum(dictOrd_0_loop gopurs_runtime.Value, dictRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 gopurs_runtime.Value = dictRing_1_loop
_ = dictRing_1
Semiring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_1, "Semiring0"), gopurs_runtime.Value{})
_ = Semiring0_2_0
zero_3_1 := gopurs_runtime.RecordGet(Semiring0_2_0, "zero")
_ = zero_3_1
zero_4_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_1, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_4_2
one_5_3 := gopurs_runtime.RecordGet(Semiring0_2_0, "one")
_ = one_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_6, zero_3_1).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_6, zero_3_1).IntVal == 3866105248) {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_1, "sub"), zero_4_2, one_5_3)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_6, zero_3_1).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_6, zero_3_1).IntVal == 2098047435) {
__t4 = one_5_3
goto end_branch_4
} else {

}
}
{
__t4 = x_6
}
end_branch_4:
return __t4
})
}

func Call_lessThanOrEq(dictOrd_0_loop gopurs_runtime.Value, a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2).IntVal == 2098047435) != true
}

func Call_max(dictOrd_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 3866105248) {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 1111389260) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 2098047435) {
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

func Call_min(dictOrd_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 3866105248) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 1111389260) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 2098047435) {
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

func Call_clamp(dictOrd_0_loop gopurs_runtime.Value, low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
v_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), low_1, x_3)
_ = v_4_0
var __t2 gopurs_runtime.Value
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 3866105248) {
__t2 = x_3
goto end_branch_2
} else {

}
}
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 1111389260) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 2098047435) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__local_var_5_1 := __t2
_ = __local_var_5_1
v_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), hi_2, __local_var_5_1)
_ = v_6_3
var __t4 gopurs_runtime.Value
{
if (v_6_3.Type == 9 && v_6_3.IntVal == 3866105248) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (v_6_3.Type == 9 && v_6_3.IntVal == 1111389260) {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (v_6_3.Type == 9 && v_6_3.IntVal == 2098047435) {
__t4 = __local_var_5_1
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

func Call_between(dictOrd_0_loop gopurs_runtime.Value, low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var low_1 gopurs_runtime.Value = low_1_loop
_ = low_1
var hi_2 gopurs_runtime.Value = hi_2_loop
_ = hi_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_3, low_1).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_3, low_1).IntVal == 3866105248) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_3, hi_2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_3, hi_2).IntVal == 2098047435) != true)
}
end_branch_0:
return __t0
}

func Call_abs(dictOrd_0_loop gopurs_runtime.Value, dictRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 gopurs_runtime.Value = dictRing_1_loop
_ = dictRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_1, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
zero_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_1, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, zero_2_0).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, zero_2_0).IntVal == 3866105248) != true {
__t2 = x_4
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_1, "sub"), zero_3_1, x_4)
}
end_branch_2:
return __t2
})
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
