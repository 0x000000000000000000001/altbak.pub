package Data_Ord

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var ordVoid gopurs_runtime.Value
var once_ordVoid sync.Once
func Get_ordVoid() gopurs_runtime.Value {
	once_ordVoid.Do(func() {
		ordVoid = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
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
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
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
		ordString = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordStringImpl(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
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
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
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
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "LT")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT"))
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "EQ")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "EQ")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT"))
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
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "GT")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))
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
		ordNumber = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordNumberImpl(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqNumber()
}))
	})
	return ordNumber
}

var ordInt gopurs_runtime.Value
var once_ordInt sync.Once
func Get_ordInt() gopurs_runtime.Value {
	once_ordInt.Do(func() {
		ordInt = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordIntImpl(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqInt()
}))
	})
	return ordInt
}

var ordChar gopurs_runtime.Value
var once_ordChar sync.Once
func Get_ordChar() gopurs_runtime.Value {
	once_ordChar.Do(func() {
		ordChar = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordCharImpl(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqChar()
}))
	})
	return ordChar
}

var ordBoolean gopurs_runtime.Value
var once_ordBoolean sync.Once
func Get_ordBoolean() gopurs_runtime.Value {
	once_ordBoolean.Do(func() {
		ordBoolean = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply3(Get_ordBooleanImpl(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqBoolean()
}))
	})
	return ordBoolean
}

var compareRecord gopurs_runtime.Value
var once_compareRecord sync.Once
func Get_compareRecord() gopurs_runtime.Value {
	once_compareRecord.Do(func() {
		compareRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "compareRecord")
})
	})
	return compareRecord
}

var ordRecord gopurs_runtime.Value
var once_ordRecord sync.Once
func Get_ordRecord() gopurs_runtime.Value {
	once_ordRecord.Do(func() {
		ordRecord = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, dictOrdRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
eqRec1_2_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))))
_ = eqRec1_2_0
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "compareRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_2_0
}))
})
	})
	return ordRecord
}

var compare1 gopurs_runtime.Value
var once_compare1 sync.Once
func Get_compare1() gopurs_runtime.Value {
	once_compare1.Do(func() {
		compare1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "compare1")
})
	})
	return compare1
}

var compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		compare = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "compare")
})
	})
	return compare
}

var comparing gopurs_runtime.Value
var once_comparing sync.Once
func Get_comparing() gopurs_runtime.Value {
	once_comparing.Do(func() {
		comparing = gopurs_runtime.Func4(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3))
})
	})
	return comparing
}

var greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		greaterThan = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2), "_tag").StrVal == "GT")
})
	})
	return greaterThan
}

var greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		greaterThanOrEq = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2), "_tag").StrVal == "LT").IntVal == 0)
})
	})
	return greaterThanOrEq
}

var lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		lessThan = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2), "_tag").StrVal == "LT")
})
	})
	return lessThan
}

var signum gopurs_runtime.Value
var once_signum sync.Once
func Get_signum() gopurs_runtime.Value {
	once_signum.Do(func() {
		signum = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, dictRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_6, zero_3_1), "_tag").StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_1, "sub"), zero_4_2, one_5_3)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_6, zero_3_1), "_tag").StrVal == "GT")).IntVal != 0 {
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
})
	})
	return signum
}

var lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		lessThanOrEq = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), a1_1, a2_2), "_tag").StrVal == "GT").IntVal == 0)
})
	})
	return lessThanOrEq
}

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3_0, "_tag").StrVal == "LT")).IntVal != 0 {
__t1 = y_2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3_0, "_tag").StrVal == "EQ")).IntVal != 0 {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3_0, "_tag").StrVal == "GT")).IntVal != 0 {
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
})
	})
	return max
}

var min gopurs_runtime.Value
var once_min sync.Once
func Get_min() gopurs_runtime.Value {
	once_min.Do(func() {
		min = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_1, y_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3_0, "_tag").StrVal == "LT")).IntVal != 0 {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3_0, "_tag").StrVal == "EQ")).IntVal != 0 {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3_0, "_tag").StrVal == "GT")).IntVal != 0 {
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
})
	})
	return min
}

var ordArray gopurs_runtime.Value
var once_ordArray sync.Once
func Get_ordArray() gopurs_runtime.Value {
	once_ordArray.Do(func() {
		ordArray = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqArray_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}), "eq")))
_ = eqArray_1_0
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_ordInt(), "compare"), gopurs_runtime.Int(0), gopurs_runtime.Apply3(Get_ordArrayImpl(), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, y_5)
_ = v_6_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6_1, "_tag").StrVal == "EQ")).IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6_1, "_tag").StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Int(1)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6_1, "_tag").StrVal == "GT")).IntVal != 0 {
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
		ordRecordCons = gopurs_runtime.Func(func(dictOrdRecord_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons_1_0 := gopurs_runtime.Apply2(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "EqRecord0"), gopurs_runtime.Value{}), gopurs_runtime.Value{})
_ = eqRowCons_1_0
return gopurs_runtime.Func2(func(_dollar__unused_2 gopurs_runtime.Value, dictIsSymbol_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons1_4_1 := gopurs_runtime.Apply(eqRowCons_1_0, dictIsSymbol_3)
_ = eqRowCons1_4_1
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons2_6_2 := gopurs_runtime.Apply(eqRowCons1_4_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{}))
_ = eqRowCons2_6_2
return gopurs_runtime.RecordDict2("compareRecord", "EqRecord0", gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, ra_8 gopurs_runtime.Value, rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
key_10_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_3, "reflectSymbol"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy")))
_ = key_10_3
left_11_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_5, "compare"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_10_3, ra_8), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_10_3, rb_9))
_ = left_11_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(left_11_4, "_tag").StrVal == "LT").IntVal != 0 || gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(left_11_4, "_tag").StrVal == "GT").IntVal != 0 || gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(left_11_4, "_tag").StrVal == "EQ").IntVal == 0).IntVal != 0).IntVal != 0)).IntVal != 0 {
__t5 = left_11_4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrdRecord_0, "compareRecord"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy")), ra_8, rb_9)
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRowCons2_6_2
}))
})
})
})
	})
	return ordRecordCons
}

var clamp gopurs_runtime.Value
var once_clamp sync.Once
func Get_clamp() gopurs_runtime.Value {
	once_clamp.Do(func() {
		clamp = gopurs_runtime.Func4(func(dictOrd_0 gopurs_runtime.Value, low_1 gopurs_runtime.Value, hi_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), low_1, x_3)
_ = v_4_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4_0, "_tag").StrVal == "LT")).IntVal != 0 {
__t2 = x_3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4_0, "_tag").StrVal == "EQ")).IntVal != 0 {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4_0, "_tag").StrVal == "GT")).IntVal != 0 {
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
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6_3, "_tag").StrVal == "LT")).IntVal != 0 {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6_3, "_tag").StrVal == "EQ")).IntVal != 0 {
__t4 = hi_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6_3, "_tag").StrVal == "GT")).IntVal != 0 {
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
})
	})
	return clamp
}

var between gopurs_runtime.Value
var once_between sync.Once
func Get_between() gopurs_runtime.Value {
	once_between.Do(func() {
		between = gopurs_runtime.Func4(func(dictOrd_0 gopurs_runtime.Value, low_1 gopurs_runtime.Value, hi_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_3, low_1), "_tag").StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_3, hi_2), "_tag").StrVal == "GT").IntVal == 0)
}
end_branch_0:
return __t0
})
	})
	return between
}

var abs gopurs_runtime.Value
var once_abs sync.Once
func Get_abs() gopurs_runtime.Value {
	once_abs.Do(func() {
		abs = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, dictRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_1, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_2_0
zero_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_1, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, zero_2_0), "_tag").StrVal == "LT").IntVal == 0)).IntVal != 0 {
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
})
	})
	return abs
}

func Get_ordArrayImpl() gopurs_runtime.Value {
	return OrdArrayImpl
}

func Get_ordBooleanImpl() gopurs_runtime.Value {
	return OrdBooleanImpl
}

func Get_ordCharImpl() gopurs_runtime.Value {
	return OrdCharImpl
}

func Get_ordIntImpl() gopurs_runtime.Value {
	return OrdIntImpl
}

func Get_ordNumberImpl() gopurs_runtime.Value {
	return OrdNumberImpl
}

func Get_ordStringImpl() gopurs_runtime.Value {
	return OrdStringImpl
}
