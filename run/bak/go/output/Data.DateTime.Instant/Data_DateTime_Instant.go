package Data_DateTime_Instant

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
)

var negateDuration gopurs_runtime.Value
var once_negateDuration sync.Once
func Get_negateDuration() gopurs_runtime.Value {
	once_negateDuration.Do(func() {
		negateDuration = gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negateDuration(), pkg_Data_Time_Duration.Get_durationMilliseconds())
	})
	return negateDuration
}

var unInstant gopurs_runtime.Value
var once_unInstant sync.Once
func Get_unInstant() gopurs_runtime.Value {
	once_unInstant.Do(func() {
		unInstant = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return unInstant
}

var toDateTime gopurs_runtime.Value
var once_toDateTime sync.Once
func Get_toDateTime() gopurs_runtime.Value {
	once_toDateTime.Do(func() {
		toDateTime = gopurs_runtime.Apply(Get_toDateTimeImpl(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mo_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mi_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ms_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("January")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("February")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("March")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("April")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("May")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("June")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("July")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(8).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("August")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(9).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("September")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(10).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("October")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(11).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("November")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(mo_2.IntVal == gopurs_runtime.Int(12).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("December")})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Date.Get_canonicalDate(), y_1), __t0), d_3), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": h_4, "value1": mi_5, "value2": s_6, "value3": ms_7})})
})
})
})
})
})
})
})
})))
	})
	return toDateTime
}

var showInstant gopurs_runtime.Value
var once_showInstant sync.Once
func Get_showInstant() gopurs_runtime.Value {
	once_showInstant.Do(func() {
		showInstant = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Instant ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Milliseconds ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0)), gopurs_runtime.Str(")")))), gopurs_runtime.Str(")")))
})})
	})
	return showInstant
}

var ordDateTime gopurs_runtime.Value
var once_ordDateTime sync.Once
func Get_ordDateTime() gopurs_runtime.Value {
	once_ordDateTime.Do(func() {
		ordDateTime = pkg_Data_Ord.Get_ordNumber()
	})
	return ordDateTime
}

var instant gopurs_runtime.Value
var once_instant sync.Once
func Get_instant() gopurs_runtime.Value {
	once_instant.Do(func() {
		instant = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), gopurs_runtime.Float(0.0)), gopurs_runtime.Float(8639977881600000.0))).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Float(8639977881599999.0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_0:
return __t0
})
	})
	return instant
}

var fromDateTime gopurs_runtime.Value
var once_fromDateTime sync.Once
func Get_fromDateTime() gopurs_runtime.Value {
	once_fromDateTime.Do(func() {
		fromDateTime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn7(), Get_fromDateTimeImpl()), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __t0), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"])
})
	})
	return fromDateTime
}

var fromDate gopurs_runtime.Value
var once_fromDate sync.Once
func Get_fromDate() gopurs_runtime.Value {
	once_fromDate.Do(func() {
		fromDate = gopurs_runtime.Func(func(d_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(d_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn7(), Get_fromDateTimeImpl()), d_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __t0), d_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), gopurs_runtime.Int(0)), gopurs_runtime.Int(0)), gopurs_runtime.Int(0)), gopurs_runtime.Int(0))
})
	})
	return fromDate
}

var eqDateTime gopurs_runtime.Value
var once_eqDateTime sync.Once
func Get_eqDateTime() gopurs_runtime.Value {
	once_eqDateTime.Do(func() {
		eqDateTime = pkg_Data_Eq.Get_eqNumber()
	})
	return eqDateTime
}

var diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		diff = gopurs_runtime.Func(func(dictDuration_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dt1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dt2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictDuration_0.PtrVal.(map[string]gopurs_runtime.Value)["toDuration"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), dt1_1), gopurs_runtime.Apply(Get_negateDuration(), dt2_2)))
})
})
})
	})
	return diff
}

var boundedInstant gopurs_runtime.Value
var once_boundedInstant sync.Once
func Get_boundedInstant() gopurs_runtime.Value {
	once_boundedInstant.Do(func() {
		boundedInstant = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), gopurs_runtime.Float(0.0)), gopurs_runtime.Float(8639977881600000.0)), "top": gopurs_runtime.Float(8639977881599999.0), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordNumber()
})})
	})
	return boundedInstant
}

func Get_fromDateTimeImpl() gopurs_runtime.Value {
	return FromDateTimeImpl
}

func Get_toDateTimeImpl() gopurs_runtime.Value {
	return ToDateTimeImpl
}
