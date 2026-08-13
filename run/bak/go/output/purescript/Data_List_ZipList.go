package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_ZipList_ZipList gopurs_runtime.Value
var once_Data_List_ZipList_ZipList sync.Once
func Get_Data_List_ZipList_ZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_ZipList.Do(func() {
		cache_Data_List_ZipList_ZipList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_ZipList(x_0_box)
})
	})
	return cache_Data_List_ZipList_ZipList
}

var cache_Data_List_ZipList_traversableZipList gopurs_runtime.Value
var once_Data_List_ZipList_traversableZipList sync.Once
func Get_Data_List_ZipList_traversableZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_traversableZipList.Do(func() {
		cache_Data_List_ZipList_traversableZipList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Lazy_Types_traversableList()))}
	})
	return cache_Data_List_ZipList_traversableZipList
}

var cache_Data_List_ZipList_showZipList gopurs_runtime.Value
var once_Data_List_ZipList_showZipList sync.Once
func Get_Data_List_ZipList_showZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_showZipList.Do(func() {
		cache_Data_List_ZipList_showZipList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_showZipList(dictShow_0_box)
})
	})
	return cache_Data_List_ZipList_showZipList
}

var cache_Data_List_ZipList_semigroupZipList gopurs_runtime.Value
var once_Data_List_ZipList_semigroupZipList sync.Once
func Get_Data_List_ZipList_semigroupZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_semigroupZipList.Do(func() {
		cache_Data_List_ZipList_semigroupZipList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()))}
	})
	return cache_Data_List_ZipList_semigroupZipList
}

var cache_Data_List_ZipList_ordZipList gopurs_runtime.Value
var once_Data_List_ZipList_ordZipList sync.Once
func Get_Data_List_ZipList_ordZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_ordZipList.Do(func() {
		cache_Data_List_ZipList_ordZipList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_ordZipList(dictOrd_0_box)
})
	})
	return cache_Data_List_ZipList_ordZipList
}

var cache_Data_List_ZipList_newtypeZipList gopurs_runtime.Value
var once_Data_List_ZipList_newtypeZipList sync.Once
func Get_Data_List_ZipList_newtypeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_newtypeZipList.Do(func() {
		cache_Data_List_ZipList_newtypeZipList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_List_ZipList_newtypeZipList
}

var cache_Data_List_ZipList_monoidZipList gopurs_runtime.Value
var once_Data_List_ZipList_monoidZipList sync.Once
func Get_Data_List_ZipList_monoidZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_monoidZipList.Do(func() {
		cache_Data_List_ZipList_monoidZipList = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_List_Lazy_Types_monoidList()))}
	})
	return cache_Data_List_ZipList_monoidZipList
}

var cache_Data_List_ZipList_functorZipList gopurs_runtime.Value
var once_Data_List_ZipList_functorZipList sync.Once
func Get_Data_List_ZipList_functorZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_functorZipList.Do(func() {
		cache_Data_List_ZipList_functorZipList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
	})
	return cache_Data_List_ZipList_functorZipList
}

var cache_Data_List_ZipList_foldableZipList gopurs_runtime.Value
var once_Data_List_ZipList_foldableZipList sync.Once
func Get_Data_List_ZipList_foldableZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_foldableZipList.Do(func() {
		cache_Data_List_ZipList_foldableZipList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()))}
	})
	return cache_Data_List_ZipList_foldableZipList
}

var cache_Data_List_ZipList_eqZipList gopurs_runtime.Value
var once_Data_List_ZipList_eqZipList sync.Once
func Get_Data_List_ZipList_eqZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_eqZipList.Do(func() {
		cache_Data_List_ZipList_eqZipList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_eqZipList(dictEq_0_box)
})
	})
	return cache_Data_List_ZipList_eqZipList
}

var cache_Data_List_ZipList_applyZipList gopurs_runtime.Value
var once_Data_List_ZipList_applyZipList sync.Once
func Get_Data_List_ZipList_applyZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_applyZipList.Do(func() {
		cache_Data_List_ZipList_applyZipList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Data_List_Lazy_zipWith(), Get_Data_Function_apply(), v_0, v1_1)
})
})})}
	})
	return cache_Data_List_ZipList_applyZipList
}

var cache_Data_List_ZipList_zipListIsNotBind gopurs_runtime.Value
var once_Data_List_ZipList_zipListIsNotBind sync.Once
func Get_Data_List_ZipList_zipListIsNotBind() gopurs_runtime.Value {
	once_Data_List_ZipList_zipListIsNotBind.Do(func() {
		cache_Data_List_ZipList_zipListIsNotBind = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_zipListIsNotBind(_dollar__unused_0_box)
})
	})
	return cache_Data_List_ZipList_zipListIsNotBind
}

var cache_Data_List_ZipList_applicativeZipList gopurs_runtime.Value
var once_Data_List_ZipList_applicativeZipList sync.Once
func Get_Data_List_ZipList_applicativeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_applicativeZipList.Do(func() {
		cache_Data_List_ZipList_applicativeZipList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_ZipList_applyZipList()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_List_Lazy_repeat(), x_0)
})})}
	})
	return cache_Data_List_ZipList_applicativeZipList
}

var cache_Data_List_ZipList_altZipList gopurs_runtime.Value
var once_Data_List_ZipList_altZipList sync.Once
func Get_Data_List_ZipList_altZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_altZipList.Do(func() {
		cache_Data_List_ZipList_altZipList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_List_Lazy_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_List_Lazy_length(), v_0).IntVal), v1_1)
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = __local_var_4_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0))
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, __local_var_2_0)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
})
})})}
	})
	return cache_Data_List_ZipList_altZipList
}

var cache_Data_List_ZipList_plusZipList gopurs_runtime.Value
var once_Data_List_ZipList_plusZipList sync.Once
func Get_Data_List_ZipList_plusZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_plusZipList.Do(func() {
		cache_Data_List_ZipList_plusZipList = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Data_List_ZipList_altZipList()))}
}), Get_Data_List_Lazy_Types_nil()})}
	})
	return cache_Data_List_ZipList_plusZipList
}

var cache_Data_List_ZipList_alternativeZipList gopurs_runtime.Value
var once_Data_List_ZipList_alternativeZipList sync.Once
func Get_Data_List_ZipList_alternativeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_alternativeZipList.Do(func() {
		cache_Data_List_ZipList_alternativeZipList = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_ZipList_applicativeZipList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_List_ZipList_plusZipList()))}
})})}
	})
	return cache_Data_List_ZipList_alternativeZipList
}

var cache_Data_List_ZipList_altZipList__3343465296 gopurs_runtime.Value
var once_Data_List_ZipList_altZipList__3343465296 sync.Once
func Get_Data_List_ZipList_altZipList__3343465296() gopurs_runtime.Value {
	once_Data_List_ZipList_altZipList__3343465296.Do(func() {
		cache_Data_List_ZipList_altZipList__3343465296 = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_List_Lazy_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_List_Lazy_length(), v_0).IntVal), v1_1)
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = __local_var_4_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0))
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, __local_var_2_0)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
})
})})}
	})
	return cache_Data_List_ZipList_altZipList__3343465296
}

var cache_Data_List_ZipList_applicativeZipList__4243212624 gopurs_runtime.Value
var once_Data_List_ZipList_applicativeZipList__4243212624 sync.Once
func Get_Data_List_ZipList_applicativeZipList__4243212624() gopurs_runtime.Value {
	once_Data_List_ZipList_applicativeZipList__4243212624.Do(func() {
		cache_Data_List_ZipList_applicativeZipList__4243212624 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_ZipList_applyZipList()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_List_Lazy_repeat(), x_0)
})})}
	})
	return cache_Data_List_ZipList_applicativeZipList__4243212624
}

var cache_Data_List_ZipList_applyZipList__1470982352 gopurs_runtime.Value
var once_Data_List_ZipList_applyZipList__1470982352 sync.Once
func Get_Data_List_ZipList_applyZipList__1470982352() gopurs_runtime.Value {
	once_Data_List_ZipList_applyZipList__1470982352.Do(func() {
		cache_Data_List_ZipList_applyZipList__1470982352 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Data_List_Lazy_zipWith(), Get_Data_Function_apply(), v_0, v1_1)
})
})})}
	})
	return cache_Data_List_ZipList_applyZipList__1470982352
}

var cache_Data_List_ZipList_functorZipList__699353223 gopurs_runtime.Value
var once_Data_List_ZipList_functorZipList__699353223 sync.Once
func Get_Data_List_ZipList_functorZipList__699353223() gopurs_runtime.Value {
	once_Data_List_ZipList_functorZipList__699353223.Do(func() {
		cache_Data_List_ZipList_functorZipList__699353223 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
	})
	return cache_Data_List_ZipList_functorZipList__699353223
}

var cache_Data_List_ZipList_plusZipList__2873873584 gopurs_runtime.Value
var once_Data_List_ZipList_plusZipList__2873873584 sync.Once
func Get_Data_List_ZipList_plusZipList__2873873584() gopurs_runtime.Value {
	once_Data_List_ZipList_plusZipList__2873873584.Do(func() {
		cache_Data_List_ZipList_plusZipList__2873873584 = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Data_List_ZipList_altZipList()))}
}), Get_Data_List_Lazy_Types_nil()})}
	})
	return cache_Data_List_ZipList_plusZipList__2873873584
}

func Call_Data_List_ZipList_ZipList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_ZipList_showZipList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList_1_0 -> *Constructor_Data_Show_Show
showList_1_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
_ = v_2_1
var __t5 string
{
if (v_2_1 == nil) {
__t5 = "(fromFoldable [])"
goto end_branch_5
} else {

}
}
{
if (v_2_1 != nil) {
var go__go_3_2_0 gopurs_runtime.Value
go__go_3_2_0 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_2_0:
for {
if false { continue go__go_3_2_0 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_3
var __t4 gopurs_runtime.Value
{
if (v_6_3 == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_6_3 != nil) {
b_4_loop = gopurs_runtime.Str(((b_4.StrVal()) + (",")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_6_3).V0).StrVal()))
xs_5_loop = (v_6_3).V1
continue go__go_3_2_0
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
__t5 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_1).V0).StrVal())) + (gopurs_runtime.Apply2(go__go_3_2_0, gopurs_runtime.Str(""), (v_2_1).V1).StrVal())) + ("])")
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_5:
return gopurs_runtime.Str(__t5)
})}
_ = showList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(ZipList ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showList_1_0.V0), v_2).StrVal())) + (")"))
})})}
}

func Call_Data_List_ZipList_ordZipList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqList1_1_0 -> *Constructor_Data_Eq_Eq
eqList1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_1 gopurs_runtime.Value
go__go_4_2_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_1:
for {
if false { continue go__go_4_2_1 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t4 bool
{
if (v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr == nil) {
var __t3 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr == nil) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr != nil)) && (((v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V0).IntVal) != (0))) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V1)))}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V1)))}
continue go__go_4_2_1
__t4 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
}
}()
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_4_2_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))}).IntVal) != (0))
})
})))
_ = eqList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_0)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_2 gopurs_runtime.Value
go__go_4_5_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_2:
for {
if false { continue go__go_4_5_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t9 uint32
{
if (v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr == nil) {
var __t6 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr == nil) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t9 = __t6
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr == nil) {
__t9 = 380165415
goto end_branch_9
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr != nil)) && ((v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr != nil)) {
// TAST (Let): v2_7_7 -> gopurs_runtime.Value
v2_7_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V0)
_ = v2_7_7
var __t8 uint32
{
if (uint32(v2_7_7.IntVal) == 902936544) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V1)))}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V1)))}
continue go__go_4_5_2
__t8 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_8
} else {

}
}
{
__t8 = uint32(v2_7_7.IntVal)
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t9), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_4_5_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))}).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_List_ZipList_eqZipList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_3 gopurs_runtime.Value
go__go_3_0_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_3:
for {
if false { continue go__go_3_0_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t2 bool
{
if (v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr == nil) {
var __t1 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)))}
continue go__go_3_0_3
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}
}()
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})
})})}
}

func Call_Data_List_ZipList_zipListIsNotBind(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_ZipList_applyZipList()))}
}), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("bind: unreachable"))
}))})}
}


