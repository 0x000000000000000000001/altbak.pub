package Data_Time_Component

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Ring "gopurs/output/Data.Ring"
)

var showSecond gopurs_runtime.Value
var once_showSecond sync.Once
func Get_showSecond() gopurs_runtime.Value {
	once_showSecond.Do(func() {
		showSecond = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Second ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0)), gopurs_runtime.Str(")")))
})})
	})
	return showSecond
}

var showMinute gopurs_runtime.Value
var once_showMinute sync.Once
func Get_showMinute() gopurs_runtime.Value {
	once_showMinute.Do(func() {
		showMinute = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Minute ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0)), gopurs_runtime.Str(")")))
})})
	})
	return showMinute
}

var showMillisecond gopurs_runtime.Value
var once_showMillisecond sync.Once
func Get_showMillisecond() gopurs_runtime.Value {
	once_showMillisecond.Do(func() {
		showMillisecond = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Millisecond ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0)), gopurs_runtime.Str(")")))
})})
	})
	return showMillisecond
}

var showHour gopurs_runtime.Value
var once_showHour sync.Once
func Get_showHour() gopurs_runtime.Value {
	once_showHour.Do(func() {
		showHour = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Hour ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0)), gopurs_runtime.Str(")")))
})})
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
		boundedSecond = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Int(0), "top": gopurs_runtime.Int(59), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return boundedSecond
}

var boundedMinute gopurs_runtime.Value
var once_boundedMinute sync.Once
func Get_boundedMinute() gopurs_runtime.Value {
	once_boundedMinute.Do(func() {
		boundedMinute = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Int(0), "top": gopurs_runtime.Int(59), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return boundedMinute
}

var boundedMillisecond gopurs_runtime.Value
var once_boundedMillisecond sync.Once
func Get_boundedMillisecond() gopurs_runtime.Value {
	once_boundedMillisecond.Do(func() {
		boundedMillisecond = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Int(0), "top": gopurs_runtime.Int(999), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return boundedMillisecond
}

var boundedHour gopurs_runtime.Value
var once_boundedHour sync.Once
func Get_boundedHour() gopurs_runtime.Value {
	once_boundedHour.Do(func() {
		boundedHour = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Int(0), "top": gopurs_runtime.Int(23), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return boundedHour
}

var boundedEnumSecond gopurs_runtime.Value
var once_boundedEnumSecond sync.Once
func Get_boundedEnumSecond() gopurs_runtime.Value {
	once_boundedEnumSecond.Do(func() {
		boundedEnumSecond = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Int(60), "toEnum": gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(59)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": n_0})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_0:
return __t0
}), "fromEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedSecond()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumSecond()
})})
	})
	return boundedEnumSecond
}

var enumSecond gopurs_runtime.Value
var once_enumSecond sync.Once
func Get_enumSecond() gopurs_runtime.Value {
	once_enumSecond.Do(func() {
		enumSecond = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), x_0), gopurs_runtime.Int(1)))
}), "pred": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), x_0), gopurs_runtime.Int(1)))
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return enumSecond
}

var boundedEnumMinute gopurs_runtime.Value
var once_boundedEnumMinute sync.Once
func Get_boundedEnumMinute() gopurs_runtime.Value {
	once_boundedEnumMinute.Do(func() {
		boundedEnumMinute = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Int(60), "toEnum": gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(59)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": n_0})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_0:
return __t0
}), "fromEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMinute()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMinute()
})})
	})
	return boundedEnumMinute
}

var enumMinute gopurs_runtime.Value
var once_enumMinute sync.Once
func Get_enumMinute() gopurs_runtime.Value {
	once_enumMinute.Do(func() {
		enumMinute = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), x_0), gopurs_runtime.Int(1)))
}), "pred": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), x_0), gopurs_runtime.Int(1)))
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return enumMinute
}

var boundedEnumMillisecond gopurs_runtime.Value
var once_boundedEnumMillisecond sync.Once
func Get_boundedEnumMillisecond() gopurs_runtime.Value {
	once_boundedEnumMillisecond.Do(func() {
		boundedEnumMillisecond = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Int(1000), "toEnum": gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(999)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": n_0})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_0:
return __t0
}), "fromEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMillisecond()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMillisecond()
})})
	})
	return boundedEnumMillisecond
}

var enumMillisecond gopurs_runtime.Value
var once_enumMillisecond sync.Once
func Get_enumMillisecond() gopurs_runtime.Value {
	once_enumMillisecond.Do(func() {
		enumMillisecond = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), x_0), gopurs_runtime.Int(1)))
}), "pred": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), x_0), gopurs_runtime.Int(1)))
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return enumMillisecond
}

var boundedEnumHour gopurs_runtime.Value
var once_boundedEnumHour sync.Once
func Get_boundedEnumHour() gopurs_runtime.Value {
	once_boundedEnumHour.Do(func() {
		boundedEnumHour = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Int(24), "toEnum": gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(23)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": n_0})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_0:
return __t0
}), "fromEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedHour()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumHour()
})})
	})
	return boundedEnumHour
}

var enumHour gopurs_runtime.Value
var once_enumHour sync.Once
func Get_enumHour() gopurs_runtime.Value {
	once_enumHour.Do(func() {
		enumHour = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), x_0), gopurs_runtime.Int(1)))
}), "pred": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), x_0), gopurs_runtime.Int(1)))
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return enumHour
}


