package Data_Interval_Duration_Iso

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Interval_Duration "gopurs/output/Data.Interval.Duration"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		lookup = gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(k_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__1_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(k_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
goto end_branch_4
} else {

}
}
{
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__1_0
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t2 = __t4
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(k_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
goto end_branch_5
} else {

}
}
{
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__1_0
__t5 = gopurs_runtime.Value{}
}
end_branch_5:
__t2 = __t5
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(k_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
goto end_branch_6
} else {

}
}
{
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__1_0
__t6 = gopurs_runtime.Value{}
}
end_branch_6:
__t2 = __t6
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(k_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
goto end_branch_7
} else {

}
}
{
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__1_0
__t7 = gopurs_runtime.Value{}
}
end_branch_7:
__t2 = __t7
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(k_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
goto end_branch_8
} else {

}
}
{
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__1_0
__t8 = gopurs_runtime.Value{}
}
end_branch_8:
__t2 = __t8
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(k_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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
}()
})
return go__1_0
})
	})
	return lookup
}

var foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		foldMap1 = gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], pkg_Data_List_Types.Get_monoidList())
	})
	return foldMap1
}

var foldMap2 gopurs_runtime.Value
var once_foldMap2 sync.Once
func Get_foldMap2() gopurs_runtime.Value {
	once_foldMap2.Do(func() {
		foldMap2 = func() gopurs_runtime.Value {
semigroupAdditive1_0_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), v_0), v1_1)
})
})})
return gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Float(0.0), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_0_0
})}))
}()
	})
	return foldMap2
}

var fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		fold = func() gopurs_runtime.Value {
semigroupFn_0_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Apply(g_1, x_2)), gopurs_runtime.Apply(f_0, x_2))
})
})
})})
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableArray().PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_0_0
})})), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}()
	})
	return fold
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(source_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var source_2 = source_2_loop
_ = source_2
var memo_3 = memo_3_loop
_ = memo_3
v_4_1 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepUnfoldr(), source_2)
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var go__5_3 gopurs_runtime.Value
go__5_3 = gopurs_runtime.Func(func(b_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_3:
for {
if false { continue go__5_3 }
var b_6 = b_6_loop
_ = b_6
var v_7 = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
b_6_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_6})
v_7_loop = v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
continue go__5_3
__t4 = gopurs_runtime.Value{}
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
}()
})
})
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), memo_3)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
source_2_loop = v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]
memo_3_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": memo_3})
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": x_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
	})
	return toUnfoldable
}

var IsEmpty gopurs_runtime.Value
var once_IsEmpty sync.Once
func Get_IsEmpty() gopurs_runtime.Value {
	once_IsEmpty.Do(func() {
		IsEmpty = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IsEmpty")})
	})
	return IsEmpty
}

var InvalidWeekComponentUsage gopurs_runtime.Value
var once_InvalidWeekComponentUsage sync.Once
func Get_InvalidWeekComponentUsage() gopurs_runtime.Value {
	once_InvalidWeekComponentUsage.Do(func() {
		InvalidWeekComponentUsage = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("InvalidWeekComponentUsage")})
	})
	return InvalidWeekComponentUsage
}

var ContainsNegativeValue gopurs_runtime.Value
var once_ContainsNegativeValue sync.Once
func Get_ContainsNegativeValue() gopurs_runtime.Value {
	once_ContainsNegativeValue.Do(func() {
		ContainsNegativeValue = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("ContainsNegativeValue"), "value0": value0})
})
	})
	return ContainsNegativeValue
}

var InvalidFractionalUse gopurs_runtime.Value
var once_InvalidFractionalUse sync.Once
func Get_InvalidFractionalUse() gopurs_runtime.Value {
	once_InvalidFractionalUse.Do(func() {
		InvalidFractionalUse = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("InvalidFractionalUse"), "value0": value0})
})
	})
	return InvalidFractionalUse
}

var unIsoDuration gopurs_runtime.Value
var once_unIsoDuration sync.Once
func Get_unIsoDuration() gopurs_runtime.Value {
	once_unIsoDuration.Do(func() {
		unIsoDuration = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return unIsoDuration
}

var showIsoDuration gopurs_runtime.Value
var once_showIsoDuration sync.Once
func Get_showIsoDuration() gopurs_runtime.Value {
	once_showIsoDuration.Do(func() {
		showIsoDuration = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(IsoDuration ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Duration ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Interval_Duration.Get_show(), v_0)), gopurs_runtime.Str(")")))), gopurs_runtime.Str(")")))
})})
	})
	return showIsoDuration
}

var showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		showError = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IsEmpty")).IntVal != 0 {
__t0 = gopurs_runtime.Str("(IsEmpty)")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidWeekComponentUsage")).IntVal != 0 {
__t0 = gopurs_runtime.Str("(InvalidWeekComponentUsage)")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ContainsNegativeValue")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Minute")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Second")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Hour")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Day")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Week")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Month")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Year")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(ContainsNegativeValue ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), __t1), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidFractionalUse")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Minute")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Second")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Hour")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Day")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Week")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Month")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Year")
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(InvalidFractionalUse ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), __t2), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})})
	})
	return showError
}

var prettyError gopurs_runtime.Value
var once_prettyError sync.Once
func Get_prettyError() gopurs_runtime.Value {
	once_prettyError.Do(func() {
		prettyError = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IsEmpty")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Duration is empty (has no components)")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidWeekComponentUsage")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Week component of Duration is used with other components")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ContainsNegativeValue")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Minute")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Second")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Hour")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Day")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Week")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Month")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Year")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("Component `")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), __t1), gopurs_runtime.Str("` contains negative value")))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidFractionalUse")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Minute")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Second")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Hour")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Day")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Week")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Month")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year")).IntVal != 0 {
__t2 = gopurs_runtime.Str("Year")
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("Invalid usage of Fractional value at component `")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), __t2), gopurs_runtime.Str("`")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return prettyError
}

var eqIsoDuration gopurs_runtime.Value
var once_eqIsoDuration sync.Once
func Get_eqIsoDuration() gopurs_runtime.Value {
	once_eqIsoDuration.Do(func() {
		eqIsoDuration = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Interval_Duration.Get_eq(), x_0), y_1)
})
})})
	})
	return eqIsoDuration
}

var ordIsoDuration gopurs_runtime.Value
var once_ordIsoDuration sync.Once
func Get_ordIsoDuration() gopurs_runtime.Value {
	once_ordIsoDuration.Do(func() {
		ordIsoDuration = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Interval_Duration.Get_compare(), x_0), y_1)
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqIsoDuration()
})})
	})
	return ordIsoDuration
}

var eqError gopurs_runtime.Value
var once_eqError sync.Once
func Get_eqError() gopurs_runtime.Value {
	once_eqError.Do(func() {
		eqError = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IsEmpty")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IsEmpty")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidWeekComponentUsage")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidWeekComponentUsage")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ContainsNegativeValue")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0)
}
end_branch_2:
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ContainsNegativeValue").IntVal != 0 && __t2.IntVal != 0)
goto end_branch_1
} else {

}
}
{
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0)
}
end_branch_0:
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidFractionalUse").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidFractionalUse").IntVal != 0 && __t0.IntVal != 0).IntVal != 0)
}
end_branch_1:
return __t1
})
})})
	})
	return eqError
}

var ordError gopurs_runtime.Value
var once_ordError sync.Once
func Get_ordError() gopurs_runtime.Value {
	once_ordError.Do(func() {
		ordError = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IsEmpty")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IsEmpty")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IsEmpty")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidWeekComponentUsage")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidWeekComponentUsage")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidWeekComponentUsage")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ContainsNegativeValue")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ContainsNegativeValue")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_8:
__t4 = __t8
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_9:
__t4 = __t9
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_10:
__t4 = __t10
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0)).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ContainsNegativeValue")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidFractionalUse").IntVal != 0 && gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "InvalidFractionalUse").IntVal != 0)).IntVal != 0 {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_12:
__t11 = __t12
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Second")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t13 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_13:
__t11 = __t13
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Minute")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_14:
__t11 = __t14
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Hour")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
var __t15 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t15 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_15:
__t11 = __t15
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Day")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
var __t16 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t16 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_16:
__t11 = __t16
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Week")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
var __t17 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t17 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_17:
__t11 = __t17
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Month")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Year").IntVal != 0)).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t0 = __t11
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqError()
})})
	})
	return ordError
}

var checkWeekUsage gopurs_runtime.Value
var once_checkWeekUsage sync.Once
func Get_checkWeekUsage() gopurs_runtime.Value {
	once_checkWeekUsage.Do(func() {
		checkWeekUsage = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_lookup(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Week")})), v_0.PtrVal.(map[string]gopurs_runtime.Value)["asMap"])
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["asMap"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["asMap"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t3 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["asMap"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), __t2), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], __t3), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT"))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("InvalidWeekComponentUsage")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
}
end_branch_0:
return __t0
})
	})
	return checkWeekUsage
}

var checkNegativeValues gopurs_runtime.Value
var once_checkNegativeValues sync.Once
func Get_checkNegativeValues() gopurs_runtime.Value {
	once_checkNegativeValues.Do(func() {
		checkNegativeValues = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldMap1(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Float(0.0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("ContainsNegativeValue"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
}
end_branch_0:
return __t0
})), v_0.PtrVal.(map[string]gopurs_runtime.Value)["asList"])
})
	})
	return checkNegativeValues
}

var checkFractionalUse gopurs_runtime.Value
var once_checkFractionalUse sync.Once
func Get_checkFractionalUse() gopurs_runtime.Value {
	once_checkFractionalUse.Do(func() {
		checkFractionalUse = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_span(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolNot(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqBooleanImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqNumberImpl(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), x_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), x_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Bool(false)))
})), v_0.PtrVal.(map[string]gopurs_runtime.Value)["asList"])
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["rest"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldMap2(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})), __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["rest"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]))), gopurs_runtime.Float(0.0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("InvalidFractionalUse"), "value0": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["rest"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
}
end_branch_1:
return __t1
})
	})
	return checkFractionalUse
}

var checkEmptiness gopurs_runtime.Value
var once_checkEmptiness sync.Once
func Get_checkEmptiness() gopurs_runtime.Value {
	once_checkEmptiness.Do(func() {
		checkEmptiness = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["asList"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IsEmpty")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
}
end_branch_0:
return __t0
})
	})
	return checkEmptiness
}

var checkValidIsoDuration gopurs_runtime.Value
var once_checkValidIsoDuration sync.Once
func Get_checkValidIsoDuration() gopurs_runtime.Value {
	once_checkValidIsoDuration.Do(func() {
		checkValidIsoDuration = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v_2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
v_2_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2})
v1_3_loop = v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
continue go__1_0
__t1 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fold(), gopurs_runtime.Array([]gopurs_runtime.Value{Get_checkWeekUsage(), Get_checkEmptiness(), Get_checkFractionalUse(), Get_checkNegativeValues()})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"asList": gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), gopurs_runtime.Apply(Get_toUnfoldable(), v_0)), "asMap": v_0}))
})
	})
	return checkValidIsoDuration
}

var mkIsoDuration gopurs_runtime.Value
var once_mkIsoDuration sync.Once
func Get_mkIsoDuration() gopurs_runtime.Value {
	once_mkIsoDuration.Do(func() {
		mkIsoDuration = gopurs_runtime.Func(func(d_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_checkValidIsoDuration(), d_0)
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": d_0})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
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
	return mkIsoDuration
}


