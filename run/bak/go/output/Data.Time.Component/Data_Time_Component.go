package Data_Time_Component

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
)

var showSecond gopurs_runtime.Value
var once_showSecond sync.Once
func Get_showSecond() gopurs_runtime.Value {
	once_showSecond.Do(func() {
		showSecond = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Second ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showSecond
}

var showMinute gopurs_runtime.Value
var once_showMinute sync.Once
func Get_showMinute() gopurs_runtime.Value {
	once_showMinute.Do(func() {
		showMinute = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Minute ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showMinute
}

var showMillisecond gopurs_runtime.Value
var once_showMillisecond sync.Once
func Get_showMillisecond() gopurs_runtime.Value {
	once_showMillisecond.Do(func() {
		showMillisecond = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Millisecond ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showMillisecond
}

var showHour gopurs_runtime.Value
var once_showHour sync.Once
func Get_showHour() gopurs_runtime.Value {
	once_showHour.Do(func() {
		showHour = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Hour ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showHour
}

var ordSecond gopurs_runtime.Value
var once_ordSecond sync.Once
func Get_ordSecond() gopurs_runtime.Value {
	once_ordSecond.Do(func() {
		ordSecond = pkg_Data_Ord.Get_ordInt()
	})
	return ordSecond
}

var ordMinute gopurs_runtime.Value
var once_ordMinute sync.Once
func Get_ordMinute() gopurs_runtime.Value {
	once_ordMinute.Do(func() {
		ordMinute = pkg_Data_Ord.Get_ordInt()
	})
	return ordMinute
}

var ordMillisecond gopurs_runtime.Value
var once_ordMillisecond sync.Once
func Get_ordMillisecond() gopurs_runtime.Value {
	once_ordMillisecond.Do(func() {
		ordMillisecond = pkg_Data_Ord.Get_ordInt()
	})
	return ordMillisecond
}

var ordHour gopurs_runtime.Value
var once_ordHour sync.Once
func Get_ordHour() gopurs_runtime.Value {
	once_ordHour.Do(func() {
		ordHour = pkg_Data_Ord.Get_ordInt()
	})
	return ordHour
}

var eqSecond gopurs_runtime.Value
var once_eqSecond sync.Once
func Get_eqSecond() gopurs_runtime.Value {
	once_eqSecond.Do(func() {
		eqSecond = pkg_Data_Eq.Get_eqInt()
	})
	return eqSecond
}

var eqMinute gopurs_runtime.Value
var once_eqMinute sync.Once
func Get_eqMinute() gopurs_runtime.Value {
	once_eqMinute.Do(func() {
		eqMinute = pkg_Data_Eq.Get_eqInt()
	})
	return eqMinute
}

var eqMillisecond gopurs_runtime.Value
var once_eqMillisecond sync.Once
func Get_eqMillisecond() gopurs_runtime.Value {
	once_eqMillisecond.Do(func() {
		eqMillisecond = pkg_Data_Eq.Get_eqInt()
	})
	return eqMillisecond
}

var eqHour gopurs_runtime.Value
var once_eqHour sync.Once
func Get_eqHour() gopurs_runtime.Value {
	once_eqHour.Do(func() {
		eqHour = pkg_Data_Eq.Get_eqInt()
	})
	return eqHour
}

var boundedSecond gopurs_runtime.Value
var once_boundedSecond sync.Once
func Get_boundedSecond() gopurs_runtime.Value {
	once_boundedSecond.Do(func() {
		boundedSecond = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(59), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return boundedSecond
}

var boundedMinute gopurs_runtime.Value
var once_boundedMinute sync.Once
func Get_boundedMinute() gopurs_runtime.Value {
	once_boundedMinute.Do(func() {
		boundedMinute = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(59), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return boundedMinute
}

var boundedMillisecond gopurs_runtime.Value
var once_boundedMillisecond sync.Once
func Get_boundedMillisecond() gopurs_runtime.Value {
	once_boundedMillisecond.Do(func() {
		boundedMillisecond = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(999), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return boundedMillisecond
}

var boundedHour gopurs_runtime.Value
var once_boundedHour sync.Once
func Get_boundedHour() gopurs_runtime.Value {
	once_boundedHour.Do(func() {
		boundedHour = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(23), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return boundedHour
}

var boundedEnumSecond gopurs_runtime.Value
var once_boundedEnumSecond sync.Once
func Get_boundedEnumSecond() gopurs_runtime.Value {
	once_boundedEnumSecond.Do(func() {
		boundedEnumSecond = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(60), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(n_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(n_0.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", n_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedSecond()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumSecond()
}))
	})
	return boundedEnumSecond
}

var enumSecond gopurs_runtime.Value
var once_enumSecond sync.Once
func Get_enumSecond() gopurs_runtime.Value {
	once_enumSecond.Do(func() {
		enumSecond = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Int(x_0.IntVal + gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_0.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", __local_var_1_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Int(x_0.IntVal - gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_2.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_2.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", __local_var_1_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return enumSecond
}

var boundedEnumMinute gopurs_runtime.Value
var once_boundedEnumMinute sync.Once
func Get_boundedEnumMinute() gopurs_runtime.Value {
	once_boundedEnumMinute.Do(func() {
		boundedEnumMinute = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(60), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(n_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(n_0.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", n_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMinute()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMinute()
}))
	})
	return boundedEnumMinute
}

var enumMinute gopurs_runtime.Value
var once_enumMinute sync.Once
func Get_enumMinute() gopurs_runtime.Value {
	once_enumMinute.Do(func() {
		enumMinute = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Int(x_0.IntVal + gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_0.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", __local_var_1_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Int(x_0.IntVal - gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_2.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_2.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", __local_var_1_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return enumMinute
}

var boundedEnumMillisecond gopurs_runtime.Value
var once_boundedEnumMillisecond sync.Once
func Get_boundedEnumMillisecond() gopurs_runtime.Value {
	once_boundedEnumMillisecond.Do(func() {
		boundedEnumMillisecond = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(1000), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(n_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(n_0.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", n_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMillisecond()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMillisecond()
}))
	})
	return boundedEnumMillisecond
}

var enumMillisecond gopurs_runtime.Value
var once_enumMillisecond sync.Once
func Get_enumMillisecond() gopurs_runtime.Value {
	once_enumMillisecond.Do(func() {
		enumMillisecond = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Int(x_0.IntVal + gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_0.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", __local_var_1_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Int(x_0.IntVal - gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_2.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_2.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", __local_var_1_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return enumMillisecond
}

var boundedEnumHour gopurs_runtime.Value
var once_boundedEnumHour sync.Once
func Get_boundedEnumHour() gopurs_runtime.Value {
	once_boundedEnumHour.Do(func() {
		boundedEnumHour = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(24), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(n_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(n_0.IntVal <= gopurs_runtime.Int(23).IntVal).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", n_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedHour()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumHour()
}))
	})
	return boundedEnumHour
}

var enumHour gopurs_runtime.Value
var once_enumHour sync.Once
func Get_enumHour() gopurs_runtime.Value {
	once_enumHour.Do(func() {
		enumHour = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Int(x_0.IntVal + gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_0.IntVal <= gopurs_runtime.Int(23).IntVal).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", __local_var_1_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Int(x_0.IntVal - gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_2.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_2.IntVal <= gopurs_runtime.Int(23).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", __local_var_1_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return enumHour
}


