package Data_Interval_Duration_Iso

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Interval_Duration "gopurs/output/Data.Interval.Duration"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	unsafe "unsafe"
)

var cache_lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		cache_lookup = gopurs_runtime.Func(func(k_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
if (k_0.Type == 9 && k_0.IntVal == 3908053364) {
var __t3 gopurs_runtime.Value
{
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 3908053364) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3})}
goto end_branch_3
} else {

}
}
{
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
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
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 3908053364) {
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (k_0.Type == 9 && k_0.IntVal == 217821258) {
var __t4 gopurs_runtime.Value
{
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 217821258) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3})}
goto end_branch_4
} else {

}
}
{
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
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
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 217821258) {
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (k_0.Type == 9 && k_0.IntVal == 1292308612) {
var __t5 gopurs_runtime.Value
{
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 1292308612) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3})}
goto end_branch_5
} else {

}
}
{
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
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
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 1292308612) {
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (k_0.Type == 9 && k_0.IntVal == 2311060696) {
var __t6 gopurs_runtime.Value
{
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 2311060696) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3})}
goto end_branch_6
} else {

}
}
{
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
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
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 2311060696) {
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (k_0.Type == 9 && k_0.IntVal == 401302776) {
var __t7 gopurs_runtime.Value
{
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 401302776) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3})}
goto end_branch_7
} else {

}
}
{
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
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
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 401302776) {
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (k_0.Type == 9 && k_0.IntVal == 3327533908) {
var __t8 gopurs_runtime.Value
{
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 3327533908) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3})}
goto end_branch_8
} else {

}
}
{
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
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
if ((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 3327533908) {
v_2_loop = (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if ((k_0.Type == 9 && k_0.IntVal == 3631736139)) && (((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.Type == 9 && (*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal == 3631736139)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3})}
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
	return cache_lookup
}

var cache_foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		cache_foldMap1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), pkg_Data_List_Types.Get_monoidList())
	})
	return cache_foldMap1
}

var cache_foldMap2 gopurs_runtime.Value
var once_foldMap2 sync.Once
func Get_foldMap2() gopurs_runtime.Value {
	once_foldMap2.Do(func() {
		cache_foldMap2 = func() gopurs_runtime.Value {
semigroupAdditive1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
}))
_ = semigroupAdditive1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_0_0
})))
}()
	})
	return cache_foldMap2
}

var cache_fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		cache_fold = func() gopurs_runtime.Value {
semigroupFn_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(f_0, x_2))
}))
_ = semigroupFn_0_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_0_0
})), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}()
	})
	return cache_fold
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if (v_4_1.Type == 9 && v_4_1.IntVal == 3589588149) {
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
if (v_7.Type == 9 && v_7.IntVal == 786377863) {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_7.UnsafePtr).V0, b_6})}
v_7_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_7.UnsafePtr).V1
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
__t2 = gopurs_runtime.Apply2(go__5_3, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, memo_3)
goto end_branch_2
} else {

}
}
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136) {
source_2_loop = (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4_1.UnsafePtr).V0.UnsafePtr).V1
memo_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4_1.UnsafePtr).V0.UnsafePtr).V0, memo_3})}
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
return gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_IterNode{x_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_IterLeaf{})}})}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}()
})
	})
	return cache_toUnfoldable
}

var cache_IsEmpty gopurs_runtime.Value
var once_IsEmpty sync.Once
func Get_IsEmpty() gopurs_runtime.Value {
	once_IsEmpty.Do(func() {
		cache_IsEmpty = gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Iso_IsEmpty{})}
	})
	return cache_IsEmpty
}

var cache_InvalidWeekComponentUsage gopurs_runtime.Value
var once_InvalidWeekComponentUsage sync.Once
func Get_InvalidWeekComponentUsage() gopurs_runtime.Value {
	once_InvalidWeekComponentUsage.Do(func() {
		cache_InvalidWeekComponentUsage = gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Iso_InvalidWeekComponentUsage{})}
	})
	return cache_InvalidWeekComponentUsage
}

var cache_ContainsNegativeValue gopurs_runtime.Value
var once_ContainsNegativeValue sync.Once
func Get_ContainsNegativeValue() gopurs_runtime.Value {
	once_ContainsNegativeValue.Do(func() {
		cache_ContainsNegativeValue = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Iso_ContainsNegativeValue{value0})}
})
	})
	return cache_ContainsNegativeValue
}

var cache_InvalidFractionalUse gopurs_runtime.Value
var once_InvalidFractionalUse sync.Once
func Get_InvalidFractionalUse() gopurs_runtime.Value {
	once_InvalidFractionalUse.Do(func() {
		cache_InvalidFractionalUse = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Iso_InvalidFractionalUse{value0})}
})
	})
	return cache_InvalidFractionalUse
}

var cache_unIsoDuration gopurs_runtime.Value
var once_unIsoDuration sync.Once
func Get_unIsoDuration() gopurs_runtime.Value {
	once_unIsoDuration.Do(func() {
		cache_unIsoDuration = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}()
})
	})
	return cache_unIsoDuration
}

var cache_showIsoDuration gopurs_runtime.Value
var once_showIsoDuration sync.Once
func Get_showIsoDuration() gopurs_runtime.Value {
	once_showIsoDuration.Do(func() {
		cache_showIsoDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(IsoDuration (Duration ") + (gopurs_runtime.Apply(pkg_Data_Interval_Duration.Get_show(), v_0).StrVal())) + ("))"))
}))
	})
	return cache_showIsoDuration
}

var cache_showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		cache_showError = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1422140417) {
__t0 = gopurs_runtime.Str("(IsEmpty)")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1775501833) {
__t0 = gopurs_runtime.Str("(InvalidWeekComponentUsage)")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3224543173) {
var __t1 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 217821258) {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Minute)")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 3908053364) {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Second)")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 1292308612) {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Hour)")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 2311060696) {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Day)")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 401302776) {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Week)")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 3327533908) {
__t1 = gopurs_runtime.Str("(ContainsNegativeValue Month)")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 3631736139) {
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
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
var __t2 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 217821258) {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Minute)")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 3908053364) {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Second)")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 1292308612) {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Hour)")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 2311060696) {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Day)")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 401302776) {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Week)")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 3327533908) {
__t2 = gopurs_runtime.Str("(InvalidFractionalUse Month)")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 3631736139) {
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
	return cache_showError
}

var cache_prettyError gopurs_runtime.Value
var once_prettyError sync.Once
func Get_prettyError() gopurs_runtime.Value {
	once_prettyError.Do(func() {
		cache_prettyError = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1422140417) {
__t0 = gopurs_runtime.Str("Duration is empty (has no components)")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1775501833) {
__t0 = gopurs_runtime.Str("Week component of Duration is used with other components")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3224543173) {
var __t1 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 217821258) {
__t1 = gopurs_runtime.Str("Component `Minute` contains negative value")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 3908053364) {
__t1 = gopurs_runtime.Str("Component `Second` contains negative value")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 1292308612) {
__t1 = gopurs_runtime.Str("Component `Hour` contains negative value")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 2311060696) {
__t1 = gopurs_runtime.Str("Component `Day` contains negative value")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 401302776) {
__t1 = gopurs_runtime.Str("Component `Week` contains negative value")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 3327533908) {
__t1 = gopurs_runtime.Str("Component `Month` contains negative value")
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0.IntVal == 3631736139) {
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
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
var __t2 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 217821258) {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Minute`")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 3908053364) {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Second`")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 1292308612) {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Hour`")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 2311060696) {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Day`")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 401302776) {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Week`")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 3327533908) {
__t2 = gopurs_runtime.Str("Invalid usage of Fractional value at component `Month`")
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0.IntVal == 3631736139) {
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
	return cache_prettyError
}

var cache_eqIsoDuration gopurs_runtime.Value
var once_eqIsoDuration sync.Once
func Get_eqIsoDuration() gopurs_runtime.Value {
	once_eqIsoDuration.Do(func() {
		cache_eqIsoDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Interval_Duration.Get_eq(), x_0, y_1)
}))
	})
	return cache_eqIsoDuration
}

var cache_ordIsoDuration gopurs_runtime.Value
var once_ordIsoDuration sync.Once
func Get_ordIsoDuration() gopurs_runtime.Value {
	once_ordIsoDuration.Do(func() {
		cache_ordIsoDuration = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Interval_Duration.Get_compare(), x_0, y_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqIsoDuration()
}))
	})
	return cache_ordIsoDuration
}

var cache_eqError gopurs_runtime.Value
var once_eqError sync.Once
func Get_eqError() gopurs_runtime.Value {
	once_eqError.Do(func() {
		cache_eqError = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
__t1 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 1422140417))
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
__t1 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 1775501833))
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t2 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 3908053364) {
__t2 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 3908053364))
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 217821258) {
__t2 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 217821258))
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 1292308612) {
__t2 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 1292308612))
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 2311060696) {
__t2 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 2311060696))
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 401302776) {
__t2 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 401302776))
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 3327533908) {
__t2 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 3327533908))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool((((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 3631736139)) && (((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 3631736139)))
}
end_branch_2:
__t1 = gopurs_runtime.Bool(((y_1.Type == 9 && y_1.IntVal == 3224543173)) && ((__t2.IntVal) != (0)))
goto end_branch_1
} else {

}
}
{
var __t0 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 3908053364) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 3908053364))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 217821258) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 217821258))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 1292308612) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 1292308612))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 2311060696) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 2311060696))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 401302776) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 401302776))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 3327533908) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 3327533908))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool((((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 3631736139)) && (((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 3631736139)))
}
end_branch_0:
__t1 = gopurs_runtime.Bool(((x_0.Type == 9 && x_0.IntVal == 574232667)) && (((y_1.Type == 9 && y_1.IntVal == 574232667)) && ((__t0.IntVal) != (0))))
}
end_branch_1:
return __t1
}))
	})
	return cache_eqError
}

var cache_ordError gopurs_runtime.Value
var once_ordError sync.Once
func Get_ordError() gopurs_runtime.Value {
	once_ordError.Do(func() {
		cache_ordError = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
var __t4 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 3908053364) {
var __t5 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 3908053364) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 3908053364) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 217821258) {
var __t6 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 217821258) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 217821258) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 1292308612) {
var __t7 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 1292308612) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 1292308612) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 2311060696) {
var __t8 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 2311060696) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_8:
__t4 = __t8
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 2311060696) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 401302776) {
var __t9 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 401302776) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_9:
__t4 = __t9
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 401302776) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 3327533908) {
var __t10 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 3327533908) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_10:
__t4 = __t10
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 3327533908) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
if (((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0.IntVal == 3631736139)) && (((*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0.IntVal == 3631736139)) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 574232667)) && ((y_1.Type == 9 && y_1.IntVal == 574232667)) {
var __t11 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 3908053364) {
var __t12 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 3908053364) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_12:
__t11 = __t12
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 3908053364) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 217821258) {
var __t13 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 217821258) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_13:
__t11 = __t13
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 217821258) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 1292308612) {
var __t14 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 1292308612) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_14:
__t11 = __t14
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 1292308612) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 2311060696) {
var __t15 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 2311060696) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_15:
__t11 = __t15
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 2311060696) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 401302776) {
var __t16 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 401302776) {
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_16:
__t11 = __t16
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 401302776) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 3327533908) {
var __t17 gopurs_runtime.Value
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 3327533908) {
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_17:
__t11 = __t17
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 3327533908) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_11
} else {

}
}
{
if (((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0.IntVal == 3631736139)) && (((*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0.IntVal == 3631736139)) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
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
	return cache_ordError
}

var cache_checkWeekUsage gopurs_runtime.Value
var once_checkWeekUsage sync.Once
func Get_checkWeekUsage() gopurs_runtime.Value {
	once_checkWeekUsage.Do(func() {
		cache_checkWeekUsage = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
__local_var_1_1 := gopurs_runtime.Apply2(Get_lookup(), gopurs_runtime.Value{Type: 9, IntVal: 401302776, UnsafePtr: unsafe.Pointer(&pkg_Data_Interval_Duration.Data_Data_Interval_Duration_Week{})}, gopurs_runtime.RecordGet(v_0, "asMap"))
_ = __local_var_1_1
var __t2 gopurs_runtime.Value
{
if (__local_var_1_1.Type == 9 && __local_var_1_1.IntVal == 3589588149) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
if (__local_var_1_1.Type == 9 && __local_var_1_1.IntVal == 930809136) {
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
if (gopurs_runtime.RecordGet(v_0, "asMap").Type == 9 && gopurs_runtime.RecordGet(v_0, "asMap").IntVal == 687041424) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.RecordGet(v_0, "asMap").Type == 9 && gopurs_runtime.RecordGet(v_0, "asMap").IntVal == 324739070) {
__t3 = gopurs_runtime.Bool(((*pkg_Data_Map_Internal.Data_Data_Map_Internal_Node)(gopurs_runtime.RecordGet(v_0, "asMap").UnsafePtr).V1.IntVal) > (1))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
if ((__t2.IntVal) != (0)) && ((__t3.IntVal) != (0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Iso_InvalidWeekComponentUsage{})}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
}
end_branch_0:
return __t0
}()
})
	})
	return cache_checkWeekUsage
}

var cache_checkNegativeValues gopurs_runtime.Value
var once_checkNegativeValues sync.Once
func Get_checkNegativeValues() gopurs_runtime.Value {
	once_checkNegativeValues.Do(func() {
		cache_checkNegativeValues = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(Get_foldMap1(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v1_1.UnsafePtr).V1.FloatVal()) >= (0.0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Iso_ContainsNegativeValue{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v1_1.UnsafePtr).V0})}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordGet(v_0, "asList"))
}()
})
	})
	return cache_checkNegativeValues
}

var cache_checkFractionalUse gopurs_runtime.Value
var once_checkFractionalUse sync.Once
func Get_checkFractionalUse() gopurs_runtime.Value {
	once_checkFractionalUse.Do(func() {
		cache_checkFractionalUse = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(x_1.UnsafePtr).V1).FloatVal()) == ((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(x_1.UnsafePtr).V1.FloatVal()))
}), gopurs_runtime.RecordGet(v_0, "asList"))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet(__local_var_1_0, "rest").Type == 9 && gopurs_runtime.RecordGet(__local_var_1_0, "rest").IntVal == 1358893437)) && ((gopurs_runtime.Apply2(Get_foldMap2(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(x_2.UnsafePtr).V1)
}), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(gopurs_runtime.RecordGet(__local_var_1_0, "rest").UnsafePtr).V1).FloatVal()) > (0.0)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Iso_InvalidFractionalUse{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(gopurs_runtime.RecordGet(__local_var_1_0, "rest").UnsafePtr).V0.UnsafePtr).V0})}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
}
end_branch_1:
return __t1
}()
})
	})
	return cache_checkFractionalUse
}

var cache_checkEmptiness gopurs_runtime.Value
var once_checkEmptiness sync.Once
func Get_checkEmptiness() gopurs_runtime.Value {
	once_checkEmptiness.Do(func() {
		cache_checkEmptiness = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_0, "asList").Type == 9 && gopurs_runtime.RecordGet(v_0, "asList").IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Iso_IsEmpty{})}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
}
end_branch_0:
return __t0
}()
})
	})
	return cache_checkEmptiness
}

var cache_checkValidIsoDuration gopurs_runtime.Value
var once_checkValidIsoDuration sync.Once
func Get_checkValidIsoDuration() gopurs_runtime.Value {
	once_checkValidIsoDuration.Do(func() {
		cache_checkValidIsoDuration = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
__t1 = v_2
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
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
return gopurs_runtime.Apply2(Get_fold(), gopurs_runtime.Array([]gopurs_runtime.Value{Get_checkWeekUsage(), Get_checkEmptiness(), Get_checkFractionalUse(), Get_checkNegativeValues()}), gopurs_runtime.RecordDict2("asList", "asMap", gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, gopurs_runtime.Apply(Get_toUnfoldable(), v_0)), v_0))
}()
})
	})
	return cache_checkValidIsoDuration
}

var cache_mkIsoDuration gopurs_runtime.Value
var once_mkIsoDuration sync.Once
func Get_mkIsoDuration() gopurs_runtime.Value {
	once_mkIsoDuration.Do(func() {
		cache_mkIsoDuration = gopurs_runtime.Func(func(d_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var d_0 gopurs_runtime.Value = d_0_loop
_ = d_0
__local_var_1_0 := gopurs_runtime.Apply(Get_checkValidIsoDuration(), d_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{d_0})}
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 1358893437) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(__local_var_1_0.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(__local_var_1_0.UnsafePtr).V1})}})}
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
	return cache_mkIsoDuration
}

type Data_Data_Interval_Duration_Iso_IsEmpty struct {
	
}
func Is_Data_Data_Interval_Duration_Iso_IsEmpty(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1422140417
}

type Data_Data_Interval_Duration_Iso_InvalidWeekComponentUsage struct {
	
}
func Is_Data_Data_Interval_Duration_Iso_InvalidWeekComponentUsage(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1775501833
}

type Data_Data_Interval_Duration_Iso_ContainsNegativeValue struct {
	V0 gopurs_runtime.Value
}
func Is_Data_Data_Interval_Duration_Iso_ContainsNegativeValue(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3224543173
}

type Data_Data_Interval_Duration_Iso_InvalidFractionalUse struct {
	V0 gopurs_runtime.Value
}
func Is_Data_Data_Interval_Duration_Iso_InvalidFractionalUse(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 574232667
}


