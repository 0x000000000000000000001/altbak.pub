package Data_Interval_Duration_Iso

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Interval_Duration "gopurs/output/Data.Interval.Duration"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_Number "gopurs/output/Data.Number"
)

var lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		lookup = gopurs_runtime.Func(func(k_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Node").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(k_0.StrVal == "Second").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Second").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3])
goto end_branch_3
} else {

}
}
{
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Second").IntVal != 0 {
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(k_0.StrVal == "Minute").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Minute").IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3])
goto end_branch_4
} else {

}
}
{
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Minute").IntVal != 0 {
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(k_0.StrVal == "Hour").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Hour").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3])
goto end_branch_5
} else {

}
}
{
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Hour").IntVal != 0 {
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(k_0.StrVal == "Day").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Day").IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3])
goto end_branch_6
} else {

}
}
{
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Day").IntVal != 0 {
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(k_0.StrVal == "Week").IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Week").IntVal != 0 {
__t7 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3])
goto end_branch_7
} else {

}
}
{
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Week").IntVal != 0 {
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(k_0.StrVal == "Month").IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Month").IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3])
goto end_branch_8
} else {

}
}
{
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Month").IntVal != 0 {
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(k_0.StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2].StrVal == "Year").IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3])
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
}()
})
	})
	return lookup
}

var foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		foldMap1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), pkg_Data_List_Types.Get_monoidList())
	})
	return foldMap1
}

var foldMap2 gopurs_runtime.Value
var once_foldMap2 sync.Once
func Get_foldMap2() gopurs_runtime.Value {
	once_foldMap2.Do(func() {
		foldMap2 = func() gopurs_runtime.Value {
semigroupAdditive1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() + v1_1.FloatVal())
}))
_ = semigroupAdditive1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_0_0
})))
}()
	})
	return foldMap2
}

var fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		fold = func() gopurs_runtime.Value {
semigroupFn_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(f_0, x_2))
}))
_ = semigroupFn_0_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nil")
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_0_0
})), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}()
	})
	return fold
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(source_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var source_2 gopurs_runtime.Value = source_2_loop
_ = source_2
var memo_3 gopurs_runtime.Value = memo_3_loop
_ = memo_3
v_4_1 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepUnfoldr(), source_2)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4_1.StrVal == "Nothing").IntVal != 0 {
var go__5_3 gopurs_runtime.Value
go__5_3 = gopurs_runtime.Func(func(b_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_3:
for {
if false { continue go__5_3 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_7.StrVal == "Nil").IntVal != 0 {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_7.StrVal == "Cons").IntVal != 0 {
b_6_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0], b_6)
v_7_loop = (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1]
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
__t2 = gopurs_runtime.Apply2(go__5_3, gopurs_runtime.Constructor0("Nil"), memo_3)
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_4_1.StrVal == "Just").IntVal != 0 {
source_2_loop = (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].UnsafePtr)[1]
memo_3_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].UnsafePtr)[0], memo_3)
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
return gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Constructor2("IterNode", x_0, gopurs_runtime.Constructor0("IterLeaf")), gopurs_runtime.Constructor0("Nil"))
}()
})
	})
	return toUnfoldable
}

var IsEmpty gopurs_runtime.Value
var once_IsEmpty sync.Once
func Get_IsEmpty() gopurs_runtime.Value {
	once_IsEmpty.Do(func() {
		IsEmpty = gopurs_runtime.Constructor0("IsEmpty")
	})
	return IsEmpty
}

var InvalidWeekComponentUsage gopurs_runtime.Value
var once_InvalidWeekComponentUsage sync.Once
func Get_InvalidWeekComponentUsage() gopurs_runtime.Value {
	once_InvalidWeekComponentUsage.Do(func() {
		InvalidWeekComponentUsage = gopurs_runtime.Constructor0("InvalidWeekComponentUsage")
	})
	return InvalidWeekComponentUsage
}

var ContainsNegativeValue gopurs_runtime.Value
var once_ContainsNegativeValue sync.Once
func Get_ContainsNegativeValue() gopurs_runtime.Value {
	once_ContainsNegativeValue.Do(func() {
		ContainsNegativeValue = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("ContainsNegativeValue", value0)
})
	})
	return ContainsNegativeValue
}

var InvalidFractionalUse gopurs_runtime.Value
var once_InvalidFractionalUse sync.Once
func Get_InvalidFractionalUse() gopurs_runtime.Value {
	once_InvalidFractionalUse.Do(func() {
		InvalidFractionalUse = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("InvalidFractionalUse", value0)
})
	})
	return InvalidFractionalUse
}

var unIsoDuration gopurs_runtime.Value
var once_unIsoDuration sync.Once
func Get_unIsoDuration() gopurs_runtime.Value {
	once_unIsoDuration.Do(func() {
		unIsoDuration = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}()
})
	})
	return unIsoDuration
}

var showIsoDuration gopurs_runtime.Value
var once_showIsoDuration sync.Once
func Get_showIsoDuration() gopurs_runtime.Value {
	once_showIsoDuration.Do(func() {
		showIsoDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(IsoDuration (Duration " + gopurs_runtime.Apply(pkg_Data_Interval_Duration.Get_show(), v_0).StrVal + "))")
}))
	})
	return showIsoDuration
}

var showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		showError = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "IsEmpty").IntVal != 0 {
__t0 = gopurs_runtime.Str("(IsEmpty)")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "InvalidWeekComponentUsage").IntVal != 0 {
__t0 = gopurs_runtime.Str("(InvalidWeekComponentUsage)")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "ContainsNegativeValue").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Minute)")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Second)")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Hour)")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Day)")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Week)")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Month)")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Year").IntVal != 0 {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Year)")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "InvalidFractionalUse").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Minute)")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Second)")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Hour)")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Day)")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Week)")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Month)")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Year").IntVal != 0 {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Year)")
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
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return showError
}

var prettyError gopurs_runtime.Value
var once_prettyError sync.Once
func Get_prettyError() gopurs_runtime.Value {
	once_prettyError.Do(func() {
		prettyError = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "IsEmpty").IntVal != 0 {
__t0 = gopurs_runtime.Str("Duration is empty (has no components)")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "InvalidWeekComponentUsage").IntVal != 0 {
__t0 = gopurs_runtime.Str("Week component of Duration is used with other components")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "ContainsNegativeValue").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t1 = gopurs_runtime.Str("Component `Minute` contains negative value")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t1 = gopurs_runtime.Str("Component `Second` contains negative value")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t1 = gopurs_runtime.Str("Component `Hour` contains negative value")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t1 = gopurs_runtime.Str("Component `Day` contains negative value")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t1 = gopurs_runtime.Str("Component `Week` contains negative value")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t1 = gopurs_runtime.Str("Component `Month` contains negative value")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Year").IntVal != 0 {
__t1 = gopurs_runtime.Str("Component `Year` contains negative value")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "InvalidFractionalUse").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Minute`")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Second`")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Hour`")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Day`")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Week`")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Month`")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].StrVal == "Year").IntVal != 0 {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Year`")
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
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return prettyError
}

var eqIsoDuration gopurs_runtime.Value
var once_eqIsoDuration sync.Once
func Get_eqIsoDuration() gopurs_runtime.Value {
	once_eqIsoDuration.Do(func() {
		eqIsoDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Interval_Duration.Get_eq(), x_0, y_1)
}))
	})
	return eqIsoDuration
}

var ordIsoDuration gopurs_runtime.Value
var once_ordIsoDuration sync.Once
func Get_ordIsoDuration() gopurs_runtime.Value {
	once_ordIsoDuration.Do(func() {
		ordIsoDuration = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Interval_Duration.Get_compare(), x_0, y_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqIsoDuration()
}))
	})
	return ordIsoDuration
}

var eqError gopurs_runtime.Value
var once_eqError sync.Once
func Get_eqError() gopurs_runtime.Value {
	once_eqError.Do(func() {
		eqError = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_0.StrVal == "IsEmpty").IntVal != 0 {
__t1 = gopurs_runtime.Bool(y_1.StrVal == "IsEmpty")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "InvalidWeekComponentUsage").IntVal != 0 {
__t1 = gopurs_runtime.Bool(y_1.StrVal == "InvalidWeekComponentUsage")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "ContainsNegativeValue").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t2 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Second")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t2 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Minute")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t2 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Hour")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t2 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Day")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t2 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Week")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t2 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Month")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Year").IntVal != 0)
}
end_branch_2:
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_1.StrVal == "ContainsNegativeValue").IntVal != 0 && __t2.IntVal != 0)
goto end_branch_1
} else {

}
}
{
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Second")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Minute")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Hour")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Day")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Week")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Month")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Year").IntVal != 0)
}
end_branch_0:
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.StrVal == "InvalidFractionalUse").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "InvalidFractionalUse").IntVal != 0 && __t0.IntVal != 0)
}
end_branch_1:
return __t1
}))
	})
	return eqError
}

var ordError gopurs_runtime.Value
var once_ordError sync.Once
func Get_ordError() gopurs_runtime.Value {
	once_ordError.Do(func() {
		ordError = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_0.StrVal == "IsEmpty").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_1.StrVal == "IsEmpty").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("EQ")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("LT")
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_1.StrVal == "IsEmpty").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "InvalidWeekComponentUsage").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_1.StrVal == "InvalidWeekComponentUsage").IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("EQ")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("LT")
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_1.StrVal == "InvalidWeekComponentUsage").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "ContainsNegativeValue").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_1.StrVal == "ContainsNegativeValue").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("EQ")
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("LT")
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("EQ")
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Constructor0("LT")
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t7 = gopurs_runtime.Constructor0("EQ")
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Constructor0("LT")
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t8 = gopurs_runtime.Constructor0("EQ")
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Constructor0("LT")
}
end_branch_8:
__t4 = __t8
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
var __t9 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t9 = gopurs_runtime.Constructor0("EQ")
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Constructor0("LT")
}
end_branch_9:
__t4 = __t9
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
var __t10 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t10 = gopurs_runtime.Constructor0("EQ")
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Constructor0("LT")
}
end_branch_10:
__t4 = __t10
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Year").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("EQ")
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
__t3 = gopurs_runtime.Constructor0("LT")
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_1.StrVal == "ContainsNegativeValue").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "InvalidFractionalUse").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "InvalidFractionalUse").IntVal != 0 {
var __t11 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
var __t12 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t12 = gopurs_runtime.Constructor0("EQ")
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Constructor0("LT")
}
end_branch_12:
__t11 = __t12
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Second").IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("GT")
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
var __t13 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t13 = gopurs_runtime.Constructor0("EQ")
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Constructor0("LT")
}
end_branch_13:
__t11 = __t13
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Minute").IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("GT")
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
var __t14 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t14 = gopurs_runtime.Constructor0("EQ")
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Constructor0("LT")
}
end_branch_14:
__t11 = __t14
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Hour").IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("GT")
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
var __t15 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t15 = gopurs_runtime.Constructor0("EQ")
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Constructor0("LT")
}
end_branch_15:
__t11 = __t15
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Day").IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("GT")
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
var __t16 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t16 = gopurs_runtime.Constructor0("EQ")
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Constructor0("LT")
}
end_branch_16:
__t11 = __t16
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Week").IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("GT")
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
var __t17 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t17 = gopurs_runtime.Constructor0("EQ")
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Constructor0("LT")
}
end_branch_17:
__t11 = __t17
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Month").IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("GT")
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].StrVal == "Year").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].StrVal == "Year").IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("EQ")
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
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqError()
}))
	})
	return ordError
}

var checkWeekUsage gopurs_runtime.Value
var once_checkWeekUsage sync.Once
func Get_checkWeekUsage() gopurs_runtime.Value {
	once_checkWeekUsage.Do(func() {
		checkWeekUsage = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
__local_var_1_1 := gopurs_runtime.Apply2(Get_lookup(), gopurs_runtime.Constructor0("Week"), gopurs_runtime.RecordGet(v_0, "asMap"))
_ = __local_var_1_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_1_1.StrVal == "Nothing").IntVal != 0 {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_1_1.StrVal == "Just").IntVal != 0 {
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
if gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "asMap").StrVal == "Leaf").IntVal != 0 {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "asMap").StrVal == "Node").IntVal != 0 {
__t3 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(gopurs_runtime.RecordGet(v_0, "asMap").UnsafePtr)[1].IntVal > 1)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
if __t2.IntVal != 0 && __t3.IntVal != 0 {
__t0 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor0("InvalidWeekComponentUsage"), gopurs_runtime.Constructor0("Nil"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nil")
}
end_branch_0:
return __t0
}()
})
	})
	return checkWeekUsage
}

var checkNegativeValues gopurs_runtime.Value
var once_checkNegativeValues sync.Once
func Get_checkNegativeValues() gopurs_runtime.Value {
	once_checkNegativeValues.Do(func() {
		checkNegativeValues = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(Get_foldMap1(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[1].FloatVal() >= 0.0 {
__t0 = gopurs_runtime.Constructor0("Nil")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor1("ContainsNegativeValue", (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0]), gopurs_runtime.Constructor0("Nil"))
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordGet(v_0, "asList"))
}()
})
	})
	return checkNegativeValues
}

var checkFractionalUse gopurs_runtime.Value
var once_checkFractionalUse sync.Once
func Get_checkFractionalUse() gopurs_runtime.Value {
	once_checkFractionalUse.Do(func() {
		checkFractionalUse = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[1]).FloatVal() == (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[1].FloatVal())
}), gopurs_runtime.RecordGet(v_0, "asList"))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_1_0, "rest").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply2(Get_foldMap2(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[1])
}), (*[1024]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_1_0, "rest").UnsafePtr)[1]).FloatVal() > 0.0 {
__t1 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor1("InvalidFractionalUse", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_1_0, "rest").UnsafePtr)[0].UnsafePtr)[0]), gopurs_runtime.Constructor0("Nil"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nil")
}
end_branch_1:
return __t1
}()
})
	})
	return checkFractionalUse
}

var checkEmptiness gopurs_runtime.Value
var once_checkEmptiness sync.Once
func Get_checkEmptiness() gopurs_runtime.Value {
	once_checkEmptiness.Do(func() {
		checkEmptiness = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "asList").StrVal == "Nil").IntVal != 0 {
__t0 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Constructor0("IsEmpty"), gopurs_runtime.Constructor0("Nil"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nil")
}
end_branch_0:
return __t0
}()
})
	})
	return checkEmptiness
}

var checkValidIsoDuration gopurs_runtime.Value
var once_checkValidIsoDuration sync.Once
func Get_checkValidIsoDuration() gopurs_runtime.Value {
	once_checkValidIsoDuration.Do(func() {
		checkValidIsoDuration = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Nil").IntVal != 0 {
__t1 = v_2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_3.StrVal == "Cons").IntVal != 0 {
v_2_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0], v_2)
v1_3_loop = (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]
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
return gopurs_runtime.Apply2(Get_fold(), gopurs_runtime.Array([]gopurs_runtime.Value{Get_checkWeekUsage(), Get_checkEmptiness(), Get_checkFractionalUse(), Get_checkNegativeValues()}), gopurs_runtime.RecordDict2("asList", "asMap", gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Constructor0("Nil"), gopurs_runtime.Apply(Get_toUnfoldable(), v_0)), v_0))
}()
})
	})
	return checkValidIsoDuration
}

var mkIsoDuration gopurs_runtime.Value
var once_mkIsoDuration sync.Once
func Get_mkIsoDuration() gopurs_runtime.Value {
	once_mkIsoDuration.Do(func() {
		mkIsoDuration = gopurs_runtime.Func(func(d_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var d_0 gopurs_runtime.Value = d_0_loop
_ = d_0
__local_var_1_0 := gopurs_runtime.Apply(Get_checkValidIsoDuration(), d_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_1_0.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Right", d_0)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_1_0.StrVal == "Cons").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)[1]))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}()
})
	})
	return mkIsoDuration
}




