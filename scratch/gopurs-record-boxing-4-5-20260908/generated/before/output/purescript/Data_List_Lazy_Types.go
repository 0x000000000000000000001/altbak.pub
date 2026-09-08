package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Lazy_Types_identity gopurs_runtime.Value
var once_Data_List_Lazy_Types_identity sync.Once
func Get_Data_List_Lazy_Types_identity() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_identity.Do(func() {
		cache_Data_List_Lazy_Types_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_List_Lazy_Types_identity
}

var cache_Data_List_Lazy_Types_unwrap gopurs_runtime.Value
var once_Data_List_Lazy_Types_unwrap sync.Once
func Get_Data_List_Lazy_Types_unwrap() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unwrap.Do(func() {
		cache_Data_List_Lazy_Types_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_List_Lazy_Types_unwrap
}

var cache_Data_List_Lazy_Types_List gopurs_runtime.Value
var once_Data_List_Lazy_Types_List sync.Once
func Get_Data_List_Lazy_Types_List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_List.Do(func() {
		cache_Data_List_Lazy_Types_List = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_List(x_0_box)
})
	})
	return cache_Data_List_Lazy_Types_List
}

var cache_Data_List_Lazy_Types_Nil gopurs_runtime.Value
var once_Data_List_Lazy_Types_Nil sync.Once
func Get_Data_List_Lazy_Types_Nil() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_Nil.Do(func() {
		cache_Data_List_Lazy_Types_Nil = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Data_List_Lazy_Types_Nil
}

var cache_Data_List_Lazy_Types_Cons gopurs_runtime.Value
var once_Data_List_Lazy_Types_Cons sync.Once
func Get_Data_List_Lazy_Types_Cons() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_Cons.Do(func() {
		cache_Data_List_Lazy_Types_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, value0, value1}))}
})
})
	})
	return cache_Data_List_Lazy_Types_Cons
}

var cache_Data_List_Lazy_Types_NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_NonEmptyList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_NonEmptyList(x_0_box)
})
	})
	return cache_Data_List_Lazy_Types_NonEmptyList
}

var cache_Data_List_Lazy_Types_nil gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil sync.Once
func Get_Data_List_Lazy_Types_nil() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil.Do(func() {
		cache_Data_List_Lazy_Types_nil = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}))}
}))
	})
	return cache_Data_List_Lazy_Types_nil
}

var cache_Data_List_Lazy_Types_newtypeNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_newtypeNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_newtypeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_newtypeNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_newtypeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Data_List_Lazy_Types_newtypeNonEmptyList
}

var cache_Data_List_Lazy_Types_newtypeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_newtypeList sync.Once
func Get_Data_List_Lazy_Types_newtypeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_newtypeList.Do(func() {
		cache_Data_List_Lazy_Types_newtypeList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Data_List_Lazy_Types_newtypeList
}

var cache_Data_List_Lazy_Types_step gopurs_runtime.Value
var once_Data_List_Lazy_Types_step sync.Once
func Get_Data_List_Lazy_Types_step() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step.Do(func() {
		cache_Data_List_Lazy_Types_step = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step
}

var cache_Data_List_Lazy_Types_semigroupList gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList sync.Once
func Get_Data_List_Lazy_Types_semigroupList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V1, ys_1)})
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_semigroupList
}

var cache_Data_List_Lazy_Types_monoidList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monoidList sync.Once
func Get_Data_List_Lazy_Types_monoidList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monoidList.Do(func() {
		cache_Data_List_Lazy_Types_monoidList = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_semigroupList()))}
}), Get_Data_List_Lazy_Types_nil()}))}
	})
	return cache_Data_List_Lazy_Types_monoidList
}

var cache_Data_List_Lazy_Types_lazyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_lazyList sync.Once
func Get_Data_List_Lazy_Types_lazyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_lazyList.Do(func() {
		cache_Data_List_Lazy_Types_lazyList = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, x_1))
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_lazyList
}

var cache_Data_List_Lazy_Types_functorList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList sync.Once
func Get_Data_List_Lazy_Types_functorList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList.Do(func() {
		cache_Data_List_Lazy_Types_functorList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V1)})
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_functorList
}

var cache_Data_List_Lazy_Types_functorNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmpty sync.Once
func Get_Data_List_Lazy_Types_functorNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1468200963_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=Other bindingType=(TypeApp (TypeVar f) [(TypeVar a)])
__local_var_2_0 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
_ = __local_var_4_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))}))}
})})))}
	})
	return cache_Data_List_Lazy_Types_functorNonEmpty
}

var cache_Data_List_Lazy_Types_functorNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_functorNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_functorNonEmpty()).V0), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_functorNonEmptyList
}

var cache_Data_List_Lazy_Types_eq1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1List sync.Once
func Get_Data_List_Lazy_Types_eq1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1List.Do(func() {
		cache_Data_List_Lazy_Types_eq1List = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_0 gopurs_runtime.Value
_ = go__go_3_0_0
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_0 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = (__t_tag_5 == nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})}))}
	})
	return cache_Data_List_Lazy_Types_eq1List
}

var cache_Data_List_Lazy_Types_eqNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_eqNonEmpty sync.Once
func Get_Data_List_Lazy_Types_eqNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eqNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_eqNonEmpty = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_eqNonEmpty(dictEq_0_box)
})
	})
	return cache_Data_List_Lazy_Types_eqNonEmpty
}

var cache_Data_List_Lazy_Types_eq1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1 sync.Once
func Get_Data_List_Lazy_Types_eq1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1.Do(func() {
		cache_Data_List_Lazy_Types_eq1 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_Lazy_Types_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), xs_1_box, ys_2_box))
})
	})
	return cache_Data_List_Lazy_Types_eq1
}

var cache_Data_List_Lazy_Types_eq1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_eq1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_eq1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqNonEmpty1_1_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])])
eqNonEmpty1_1_0 := (&Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_8 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_1_3 gopurs_runtime.Value
_ = go__go_3_1_3
// FALLBACK TCO: isLoop=false len=1
go__go_3_1_3 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 bool
{
var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_5 == nil) {
var __t_tag_6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t7 = (__t_tag_6 == nil)
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_4 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_4 = ((__t_tag_3 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_3_1_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t7 = __t_and_4
}
end_branch_7:
return gopurs_runtime.Bool(__t7)
})
__t_and_8 = (gopurs_runtime.Apply2(go__go_3_1_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_8)
})})
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqNonEmpty1_1_0.V0), gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_3)).IntVal) != (0))
})
})}))}
	})
	return cache_Data_List_Lazy_Types_eq1NonEmptyList
}

var cache_Data_List_Lazy_Types_eqList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eqList sync.Once
func Get_Data_List_Lazy_Types_eqList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eqList.Do(func() {
		cache_Data_List_Lazy_Types_eqList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_eqList(dictEq_0_box)
})
	})
	return cache_Data_List_Lazy_Types_eqList
}

var cache_Data_List_Lazy_Types_eqNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eqNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_eqNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eqNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_eqNonEmptyList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_eqNonEmptyList(dictEq_0_box)
})
	})
	return cache_Data_List_Lazy_Types_eqNonEmptyList
}

var cache_Data_List_Lazy_Types_ord1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_ord1List sync.Once
func Get_Data_List_Lazy_Types_ord1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ord1List.Do(func() {
		cache_Data_List_Lazy_Types_ord1List = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))}
}), gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_3_0_6 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Lazy_Types_go__go_3_0_6
var go__go_3_0_6 gopurs_runtime.Value
_ = go__go_3_0_6
Call_local_Data_List_Lazy_Types_go__go_3_0_6 = func(v_4_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_3_0_6:
for {
if false { continue go__go_3_0_6 }
var v_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t4 uint32
{
if (v_4 == nil) {
var __t1 uint32
{
if (v1_5 == nil) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t4 = __t1
goto end_branch_4
} else {

}
}
{
if (v1_5 == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((v_4 != nil)) && ((v1_5 != nil)) {
// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_6_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v_4).V0, (v1_5).V0).IntVal)
_ = v2_6_2
var __t3 uint32
{
if (v2_6_2 == 902936544) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_4).V1))
v1_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v1_5).V1))
continue go__go_3_0_6
__t3 = func() uint32 { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = v2_6_2
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_3_0_6 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_3_0_6(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val))), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_3_0_6(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))), UnsafePtr: nil}
})}))}
	})
	return cache_Data_List_Lazy_Types_ord1List
}

var cache_Data_List_Lazy_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordNonEmpty sync.Once
func Get_Data_List_Lazy_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_ordNonEmpty = gopurs_runtime.Apply(Get_Data_NonEmpty_ordNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()))})
	})
	return cache_Data_List_Lazy_Types_ordNonEmpty
}

var cache_Data_List_Lazy_Types_compare1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compare1 sync.Once
func Get_Data_List_Lazy_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compare1.Do(func() {
		cache_Data_List_Lazy_Types_compare1 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_List_Lazy_Types_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), xs_1_box, ys_2_box)), UnsafePtr: nil}
})
	})
	return cache_Data_List_Lazy_Types_compare1
}

var cache_Data_List_Lazy_Types_ord1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_ord1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_ord1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ord1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_ord1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1NonEmptyList()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordNonEmpty1_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])])
ordNonEmpty1_1_0 := Rebox_Data_List_Lazy_Types_4177771502_167870147(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_NonEmpty_ordNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()))}, dictOrd_0)))
_ = ordNonEmpty1_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordNonEmpty1_1_0.V1), gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_3)).IntVal)), UnsafePtr: nil}
})
})}))}
	})
	return cache_Data_List_Lazy_Types_ord1NonEmptyList
}

var cache_Data_List_Lazy_Types_ordList gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordList sync.Once
func Get_Data_List_Lazy_Types_ordList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordList.Do(func() {
		cache_Data_List_Lazy_Types_ordList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_ordList(dictOrd_0_box)
})
	})
	return cache_Data_List_Lazy_Types_ordList
}

var cache_Data_List_Lazy_Types_ordNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_ordNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_ordNonEmptyList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_ordNonEmptyList(dictOrd_0_box)
})
	})
	return cache_Data_List_Lazy_Types_ordNonEmptyList
}

var cache_Data_List_Lazy_Types_cons gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons sync.Once
func Get_Data_List_Lazy_Types_cons() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons.Do(func() {
		cache_Data_List_Lazy_Types_cons = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons(x_0_box, xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons
}

var cache_Data_List_Lazy_Types_foldableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList sync.Once
func Get_Data_List_Lazy_Types_foldableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList.Do(func() {
		cache_Data_List_Lazy_Types_foldableList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1), gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_4, gopurs_runtime.Apply(f_3, a_5))
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_1_2_10 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_1_2_10
var go__go_1_2_10 gopurs_runtime.Value
_ = go__go_1_2_10
Call_local_Data_List_Lazy_Types_go__go_1_2_10 = func(b_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_1_2_10:
for {
if false { continue go__go_1_2_10 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (v_4_3 == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (v_4_3 != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (v_4_3).V0)
xs_3_loop = (v_4_3).V1
continue go__go_1_2_10
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_1_2_10 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_1_2_10(b_2_loop_val, xs_3_loop_val)
})
})
return go__go_1_2_10
}), gopurs_runtime.Func3(func(op_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, b_3}))}
}))
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})}))}
	})
	return cache_Data_List_Lazy_Types_foldableList
}

var cache_Data_List_Lazy_Types_foldableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmpty sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_3071895939_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_1
var Call_local_Data_List_Lazy_Types_go__go_5_2_11 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_5_2_11
var go__go_5_2_11 gopurs_runtime.Value
_ = go__go_5_2_11
Call_local_Data_List_Lazy_Types_go__go_5_2_11 = func(b_6_loop gopurs_runtime.Value, xs_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_5_2_11:
for {
if false { continue go__go_5_2_11 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_8_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_3
var __t4 gopurs_runtime.Value
{
if (v_8_3 == nil) {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if (v_8_3 != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_1.V0), b_6, gopurs_runtime.Apply(f_2, (v_8_3).V0))
xs_7_loop = (v_8_3).V1
continue go__go_5_2_11
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_5_2_11 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_5_2_11(b_6_loop_val, xs_7_loop_val)
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), Call_local_Data_List_Lazy_Types_go__go_5_2_11(gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_3_5_12 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_5_12
var go__go_3_5_12 gopurs_runtime.Value
_ = go__go_3_5_12
Call_local_Data_List_Lazy_Types_go__go_3_5_12 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_5_12:
for {
if false { continue go__go_3_5_12 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_6 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_6
var __t7 gopurs_runtime.Value
{
if (v_6_6 == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_6_6 != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (v_6_6).V0)
xs_5_loop = (v_6_6).V1
continue go__go_3_5_12
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_3_5_12 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_5_12(b_4_loop_val, xs_5_loop_val)
})
})
return Call_local_Data_List_Lazy_Types_go__go_3_5_12(gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_3_8_13 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_8_13
var go__go_3_8_13 gopurs_runtime.Value
_ = go__go_3_8_13
Call_local_Data_List_Lazy_Types_go__go_3_8_13 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_8_13:
for {
if false { continue go__go_3_8_13 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_9 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_9
var __t10 gopurs_runtime.Value
{
if (v_6_9 == nil) {
__t10 = b_4
goto end_branch_10
} else {

}
}
{
if (v_6_9 != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (v_6_9).V0, b_4)
xs_5_loop = (v_6_9).V1
continue go__go_3_8_13
__t10 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}
}
go__go_3_8_13 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_8_13(b_4_loop_val, xs_5_loop_val)
})
})
var Call_local_Data_List_Lazy_Types_go__go_4_11_14 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_11_14
var go__go_4_11_14 gopurs_runtime.Value
_ = go__go_4_11_14
Call_local_Data_List_Lazy_Types_go__go_4_11_14 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_11_14:
for {
if false { continue go__go_4_11_14 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_12
var __t14 gopurs_runtime.Value
{
if (v_7_12 == nil) {
__t14 = b_5
goto end_branch_14
} else {

}
}
{
if (v_7_12 != nil) {
// TAST (Let): __local_var_8_13 shape=Other bindingType=(TypeVar a)
__local_var_8_13 := (v_7_12).V0
_ = __local_var_8_13
b_5_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_8_13, b_5}))}
}))
xs_6_loop = (v_7_12).V1
continue go__go_4_11_14
__t14 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}
}
go__go_4_11_14 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_11_14(b_5_loop_val, xs_6_loop_val)
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, Call_local_Data_List_Lazy_Types_go__go_3_8_13(b_1, Call_local_Data_List_Lazy_Types_go__go_4_11_14(Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)))
})})))}
	})
	return cache_Data_List_Lazy_Types_foldableNonEmpty
}

var cache_Data_List_Lazy_Types_extendList gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendList sync.Once
func Get_Data_List_Lazy_Types_extendList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendList.Do(func() {
		cache_Data_List_Lazy_Types_extendList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
_ = v_2_0
var __t14 gopurs_runtime.Value
{
if (v_2_0 == nil) {
__t14 = Get_Data_List_Lazy_Types_nil()
goto end_branch_14
} else {

}
}
{
if (v_2_0 != nil) {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(f_0, l_1)
_ = __local_var_3_1
var Call_local_Data_List_Lazy_Types_go__go_4_2_15 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_2_15
var go__go_4_2_15 gopurs_runtime.Value
_ = go__go_4_2_15
Call_local_Data_List_Lazy_Types_go__go_4_2_15 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_2_15:
for {
if false { continue go__go_4_2_15 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_3
var __t8 gopurs_runtime.Value
{
if (v_7_3 == nil) {
__t8 = b_5
goto end_branch_8
} else {

}
}
{
if (v_7_3 != nil) {
// TAST (Let): __local_var_8_4 shape=Other bindingType=Any
__local_var_8_4 := gopurs_runtime.RecordGet(b_5, "acc")
_ = __local_var_8_4
// TAST (Let): __local_var_9_5 shape=Other bindingType=Any
__local_var_9_5 := gopurs_runtime.RecordGet(b_5, "val")
_ = __local_var_9_5
// TAST (Let): acc_prime__10_6 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
acc_prime__10_6 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v_7_3).V0, __local_var_8_4}))}
}))
_ = acc_prime__10_6
// TAST (Let): __local_var_11_7 shape=App(Other) bindingType=Any
__local_var_11_7 := gopurs_runtime.Apply(f_0, acc_prime__10_6)
_ = __local_var_11_7
b_5_loop = func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	val gopurs_runtime.Value
}{acc_prime__10_6, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_11_7, __local_var_9_5}))}
}))}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "val"}, []gopurs_runtime.Value{orig.acc, orig.val})
				}()
xs_6_loop = (v_7_3).V1
continue go__go_4_2_15
__t8 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}
go__go_4_2_15 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_2_15(b_5_loop_val, xs_6_loop_val)
})
})
var Call_local_Data_List_Lazy_Types_go__go_5_10_16 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_5_10_16
var go__go_5_10_16 gopurs_runtime.Value
_ = go__go_5_10_16
Call_local_Data_List_Lazy_Types_go__go_5_10_16 = func(b_6_loop gopurs_runtime.Value, xs_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_5_10_16:
for {
if false { continue go__go_5_10_16 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_11 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_11
var __t13 gopurs_runtime.Value
{
if (v_8_11 == nil) {
__t13 = b_6
goto end_branch_13
} else {

}
}
{
if (v_8_11 != nil) {
// TAST (Let): __local_var_9_12 shape=Other bindingType=(TypeVar a)
__local_var_9_12 := (v_8_11).V0
_ = __local_var_9_12
b_6_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_9_12, b_6}))}
}))
xs_7_loop = (v_8_11).V1
continue go__go_5_10_16
__t13 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}
go__go_5_10_16 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_5_10_16(b_6_loop_val, xs_7_loop_val)
})
})
// TAST (Let): __local_var_5_9 shape=Other bindingType=Any
__local_var_5_9 := gopurs_runtime.RecordGet(Call_local_Data_List_Lazy_Types_go__go_4_2_15(func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	val gopurs_runtime.Value
}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "val"}, []gopurs_runtime.Value{orig.acc, orig.val})
				}(), Call_local_Data_List_Lazy_Types_go__go_5_10_16(Get_Data_List_Lazy_Types_nil(), (v_2_0).V1)), "val")
_ = __local_var_5_9
__t14 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_5_9}))}
}))
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
})}))}
	})
	return cache_Data_List_Lazy_Types_extendList
}

var cache_Data_List_Lazy_Types_extendNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_extendNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_extendNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
__local_var_2_0 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1).UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_4_1_17 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_1_17
var go__go_4_1_17 gopurs_runtime.Value
_ = go__go_4_1_17
Call_local_Data_List_Lazy_Types_go__go_4_1_17 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_1_17:
for {
if false { continue go__go_4_1_17 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_7_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_2
var __t6 gopurs_runtime.Value
{
if (v_7_2 == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_7_2 != nil) {
// TAST (Let): __local_var_8_3 shape=Other bindingType=Any
__local_var_8_3 := gopurs_runtime.RecordGet(b_5, "acc")
_ = __local_var_8_3
// TAST (Let): __local_var_9_4 shape=Other bindingType=Any
__local_var_9_4 := gopurs_runtime.RecordGet(b_5, "val")
_ = __local_var_9_4
// TAST (Let): __local_var_10_5 shape=App(Other) bindingType=(TypeVar b)
__local_var_10_5 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_7_2).V0, __local_var_8_3}))}
})))
_ = __local_var_10_5
b_5_loop = func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	val gopurs_runtime.Value
}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v_7_2).V0, __local_var_8_3}))}
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_10_5, __local_var_9_4}))}
}))}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "val"}, []gopurs_runtime.Value{orig.acc, orig.val})
				}()
xs_6_loop = (v_7_2).V1
continue go__go_4_1_17
__t6 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__go_4_1_17 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_1_17(b_5_loop_val, xs_6_loop_val)
})
})
var Call_local_Data_List_Lazy_Types_go__go_5_7_18 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_5_7_18
var go__go_5_7_18 gopurs_runtime.Value
_ = go__go_5_7_18
Call_local_Data_List_Lazy_Types_go__go_5_7_18 = func(b_6_loop gopurs_runtime.Value, xs_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_5_7_18:
for {
if false { continue go__go_5_7_18 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_8 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_8
var __t10 gopurs_runtime.Value
{
if (v_8_8 == nil) {
__t10 = b_6
goto end_branch_10
} else {

}
}
{
if (v_8_8 != nil) {
// TAST (Let): __local_var_9_9 shape=Other bindingType=(TypeVar a)
__local_var_9_9 := (v_8_8).V0
_ = __local_var_9_9
b_6_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_9_9, b_6}))}
}))
xs_7_loop = (v_8_8).V1
continue go__go_5_7_18
__t10 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}
}
go__go_5_7_18 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_5_7_18(b_6_loop_val, xs_7_loop_val)
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(Call_local_Data_List_Lazy_Types_go__go_4_1_17(func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	val gopurs_runtime.Value
}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "val"}, []gopurs_runtime.Value{orig.acc, orig.val})
				}(), Call_local_Data_List_Lazy_Types_go__go_5_7_18(Get_Data_List_Lazy_Types_nil(), __local_var_2_0)), "val")}))}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_extendNonEmptyList
}

var cache_Data_List_Lazy_Types_foldableNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V1), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V2), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})}))}
	})
	return cache_Data_List_Lazy_Types_foldableNonEmptyList
}

var cache_Data_List_Lazy_Types_showList gopurs_runtime.Value
var once_Data_List_Lazy_Types_showList sync.Once
func Get_Data_List_Lazy_Types_showList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_showList.Do(func() {
		cache_Data_List_Lazy_Types_showList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_showList(dictShow_0_box)
})
	})
	return cache_Data_List_Lazy_Types_showList
}

var cache_Data_List_Lazy_Types_showNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_showNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_showNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_showNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_showNonEmptyList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_showNonEmptyList(dictShow_0_box)
})
	})
	return cache_Data_List_Lazy_Types_showNonEmptyList
}

var cache_Data_List_Lazy_Types_showStep gopurs_runtime.Value
var once_Data_List_Lazy_Types_showStep sync.Once
func Get_Data_List_Lazy_Types_showStep() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_showStep.Do(func() {
		cache_Data_List_Lazy_Types_showStep = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_showStep(dictShow_0_box)
})
	})
	return cache_Data_List_Lazy_Types_showStep
}

var cache_Data_List_Lazy_Types_foldableWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexList sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()).V2), gopurs_runtime.Func2(func(i_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=Any
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_5)
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=App(Other) bindingType=(Func [(TypeVar a)] (TypeVar m))
__local_var_7_3 := gopurs_runtime.Apply(f_3, gopurs_runtime.Int(i_4.IntVal))
_ = __local_var_7_3
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_7_3, x_8))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_2_5_22 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_2_5_22
var go__go_2_5_22 gopurs_runtime.Value
_ = go__go_2_5_22
Call_local_Data_List_Lazy_Types_go__go_2_5_22 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_5_22:
for {
if false { continue go__go_2_5_22 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_6 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_6
var __t7 gopurs_runtime.Value
{
if (v_5_6 == nil) {
__t7 = b_3
goto end_branch_7
} else {

}
}
{
if (v_5_6 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V1, (v_5_6).V0)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_4_loop = (v_5_6).V1
continue go__go_2_5_22
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_2_5_22 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_2_5_22(b_3_loop_val, xs_4_loop_val)
})
})
// TAST (Let): __local_var_2_4 shape=LetRec(App(Other)) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] (ADT ["Data","Tuple","Tuple"] [Int, (TypeVar b)]))
__local_var_2_4 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_5_22, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), acc_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))})))
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(__local_var_2_4))}, x_3).UnsafePtr).V1
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_3_9_23 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_9_23
var go__go_3_9_23 gopurs_runtime.Value
_ = go__go_3_9_23
Call_local_Data_List_Lazy_Types_go__go_3_9_23 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_9_23:
for {
if false { continue go__go_3_9_23 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_10 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_10
var __t13 gopurs_runtime.Value
{
if (v_6_10 == nil) {
__t13 = b_4
goto end_branch_13
} else {

}
}
{
if (v_6_10 != nil) {
// TAST (Let): __local_var_7_11 shape=Other bindingType=Any
__local_var_7_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1
_ = __local_var_7_11
// TAST (Let): __local_var_8_12 shape=Other bindingType=(TypeVar a)
__local_var_8_12 := (v_6_10).V0
_ = __local_var_8_12
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_8_12, __local_var_7_11}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_5_loop = (v_6_10).V1
continue go__go_3_9_23
__t13 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}
go__go_3_9_23 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_9_23(b_4_loop_val, xs_5_loop_val)
})
})
// TAST (Let): v_3_8 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
v_3_8 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_3_9_23(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, xs_2)))
_ = v_3_8
var Call_local_Data_List_Lazy_Types_go__go_4_14_24 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_14_24
var go__go_4_14_24 gopurs_runtime.Value
_ = go__go_4_14_24
Call_local_Data_List_Lazy_Types_go__go_4_14_24 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_14_24:
for {
if false { continue go__go_4_14_24 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_15 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_15
var __t16 gopurs_runtime.Value
{
if (v_7_15 == nil) {
__t16 = b_5
goto end_branch_16
} else {

}
}
{
if (v_7_15 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), (v_7_15).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_6_loop = (v_7_15).V1
continue go__go_4_14_24
__t16 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}
}
go__go_4_14_24 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_14_24(b_5_loop_val, xs_6_loop_val)
})
})
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Lazy_Types_go__go_4_14_24(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_3_8).V0), b_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_3_8).V1).UnsafePtr).V1
})})))}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexList
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmpty sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2048287455_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2048287455(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())))})))))}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableNonEmptyList()))}
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())))}), "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t0 = int64(0)
goto end_branch_0
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t0 = (int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_1, gopurs_runtime.Int(__t0))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())))}), "foldlWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 int64
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t1 = int64(0)
goto end_branch_1
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t1 = (int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t1))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())))}), "foldrWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 int64
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t2 = int64(0)
goto end_branch_2
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t2 = (int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t2))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})})))}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_functorWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexList sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_2_1_25 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_2_1_25
var go__go_2_1_25 gopurs_runtime.Value
_ = go__go_2_1_25
Call_local_Data_List_Lazy_Types_go__go_2_1_25 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_1_25:
for {
if false { continue go__go_2_1_25 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_2
var __t5 gopurs_runtime.Value
{
if (v_5_2 == nil) {
__t5 = b_3
goto end_branch_5
} else {

}
}
{
if (v_5_2 != nil) {
// TAST (Let): __local_var_6_3 shape=Other bindingType=Any
__local_var_6_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V1
_ = __local_var_6_3
// TAST (Let): __local_var_7_4 shape=Other bindingType=(TypeVar a)
__local_var_7_4 := (v_5_2).V0
_ = __local_var_7_4
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_4, __local_var_6_3}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_4_loop = (v_5_2).V1
continue go__go_2_1_25
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_2_1_25 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_2_1_25(b_3_loop_val, xs_4_loop_val)
})
})
// TAST (Let): v_2_0 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
v_2_0 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_2_1_25(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, xs_1)))
_ = v_2_0
var Call_local_Data_List_Lazy_Types_go__go_3_6_26 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_6_26
var go__go_3_6_26 gopurs_runtime.Value
_ = go__go_3_6_26
Call_local_Data_List_Lazy_Types_go__go_3_6_26 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_6_26:
for {
if false { continue go__go_3_6_26 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_7 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_7
var __t10 gopurs_runtime.Value
{
if (v_6_7 == nil) {
__t10 = b_4
goto end_branch_10
} else {

}
}
{
if (v_6_7 != nil) {
// TAST (Let): __local_var_7_8 shape=Other bindingType=(TypeVar b)
__local_var_7_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1
_ = __local_var_7_8
// TAST (Let): __local_var_8_9 shape=App(Other) bindingType=Any
__local_var_8_9 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), (v_6_7).V0)
_ = __local_var_8_9
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_8_9, __local_var_7_8}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_5_loop = (v_6_7).V1
continue go__go_3_6_26
__t10 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}
}
go__go_3_6_26 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_6_26(b_4_loop_val, xs_5_loop_val)
})
})
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Lazy_Types_go__go_3_6_26(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_2_0).V0), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_2_0).V1).UnsafePtr).V1
})})))}
	})
	return cache_Data_List_Lazy_Types_functorWithIndexList
}

var cache_Data_List_Lazy_Types_functorWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndex sync.Once
func Get_Data_List_Lazy_Types_functorWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndex = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_0 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f)])])
functorNonEmpty1_0_0 := (&Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 shape=Other bindingType=(TypeApp (TypeVar f) [(TypeVar a)])
__local_var_2_1 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_1)
_ = __local_var_4_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 218341868 && __local_var_4_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 218341868 && __local_var_4_2.UnsafePtr != nil) {
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_2.UnsafePtr).V1)})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))}))}
})})
_ = functorNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_80275103_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1468200963_2812149806(functorNonEmpty1_0_0))}
}), gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_3_5_27 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_5_27
var go__go_3_5_27 gopurs_runtime.Value
_ = go__go_3_5_27
Call_local_Data_List_Lazy_Types_go__go_3_5_27 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_5_27:
for {
if false { continue go__go_3_5_27 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_6 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_6
var __t9 gopurs_runtime.Value
{
if (v_6_6 == nil) {
__t9 = b_4
goto end_branch_9
} else {

}
}
{
if (v_6_6 != nil) {
// TAST (Let): __local_var_7_7 shape=Other bindingType=Any
__local_var_7_7 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1
_ = __local_var_7_7
// TAST (Let): __local_var_8_8 shape=Other bindingType=(TypeVar a)
__local_var_8_8 := (v_6_6).V0
_ = __local_var_8_8
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_8_8, __local_var_7_7}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_5_loop = (v_6_6).V1
continue go__go_3_5_27
__t9 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}
go__go_3_5_27 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_5_27(b_4_loop_val, xs_5_loop_val)
})
})
// TAST (Let): v_3_4 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
v_3_4 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_3_5_27(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)))
_ = v_3_4
var Call_local_Data_List_Lazy_Types_go__go_4_10_28 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_10_28
var go__go_4_10_28 gopurs_runtime.Value
_ = go__go_4_10_28
Call_local_Data_List_Lazy_Types_go__go_4_10_28 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_10_28:
for {
if false { continue go__go_4_10_28 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_11 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_11
var __t14 gopurs_runtime.Value
{
if (v_7_11 == nil) {
__t14 = b_5
goto end_branch_14
} else {

}
}
{
if (v_7_11 != nil) {
// TAST (Let): __local_var_8_12 shape=Other bindingType=(TypeVar b)
__local_var_8_12 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1
_ = __local_var_8_12
// TAST (Let): __local_var_9_13 shape=App(Other) bindingType=(TypeVar b)
__local_var_9_13 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, (v_7_11).V0)
_ = __local_var_9_13
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_9_13, __local_var_8_12}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_6_loop = (v_7_11).V1
continue go__go_4_10_28
__t14 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}
}
go__go_4_10_28 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_10_28(b_5_loop_val, xs_6_loop_val)
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Lazy_Types_go__go_4_10_28(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_3_4).V0), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_3_4).V1).UnsafePtr).V1}))}
})})))}
}()
	})
	return cache_Data_List_Lazy_Types_functorWithIndex
}

var cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
_ = __local_var_3_0
var __t3 int64
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
if (__t_tag_1 == nil) {
__t3 = int64(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
if (__t_tag_2 != nil) {
__t3 = (int64(1)) + (((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = func() int64 { panic("Failed pattern match") }()
}
end_branch_3:
var Call_local_Data_List_Lazy_Types_go__go_4_5_29 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_5_29
var go__go_4_5_29 gopurs_runtime.Value
_ = go__go_4_5_29
Call_local_Data_List_Lazy_Types_go__go_4_5_29 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_5_29:
for {
if false { continue go__go_4_5_29 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_6 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_6
var __t9 gopurs_runtime.Value
{
if (v_7_6 == nil) {
__t9 = b_5
goto end_branch_9
} else {

}
}
{
if (v_7_6 != nil) {
// TAST (Let): __local_var_8_7 shape=Other bindingType=Any
__local_var_8_7 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1
_ = __local_var_8_7
// TAST (Let): __local_var_9_8 shape=Other bindingType=(TypeVar a)
__local_var_9_8 := (v_7_6).V0
_ = __local_var_9_8
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_9_8, __local_var_8_7}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_6_loop = (v_7_6).V1
continue go__go_4_5_29
__t9 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}
go__go_4_5_29 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_5_29(b_5_loop_val, xs_6_loop_val)
})
})
// TAST (Let): v_4_4 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
v_4_4 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_4_5_29(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (__local_var_3_0).V1)))
_ = v_4_4
var Call_local_Data_List_Lazy_Types_go__go_5_10_30 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_5_10_30
var go__go_5_10_30 gopurs_runtime.Value
_ = go__go_5_10_30
Call_local_Data_List_Lazy_Types_go__go_5_10_30 = func(b_6_loop gopurs_runtime.Value, xs_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_5_10_30:
for {
if false { continue go__go_5_10_30 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_11 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_11
var __t14 gopurs_runtime.Value
{
if (v_8_11 == nil) {
__t14 = b_6
goto end_branch_14
} else {

}
}
{
if (v_8_11 != nil) {
// TAST (Let): __local_var_9_12 shape=Other bindingType=(TypeVar b)
__local_var_9_12 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V1
_ = __local_var_9_12
// TAST (Let): __local_var_10_13 shape=App(Other) bindingType=(TypeVar b)
__local_var_10_13 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int((int64(1)) + (((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V0.IntVal) - (int64(1)))), (v_8_11).V0)
_ = __local_var_10_13
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_10_13, __local_var_9_12}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_7_loop = (v_8_11).V1
continue go__go_5_10_30
__t14 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}
}
go__go_5_10_30 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_5_10_30(b_6_loop_val, xs_7_loop_val)
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(__t3), (__local_var_3_0).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Lazy_Types_go__go_5_10_30(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_4_4).V0), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_4_4).V1).UnsafePtr).V1}))}
}))
})})))}
	})
	return cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_toList gopurs_runtime.Value
var once_Data_List_Lazy_Types_toList sync.Once
func Get_Data_List_Lazy_Types_toList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_toList.Do(func() {
		cache_Data_List_Lazy_Types_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_toList(v_0_box)
})
	})
	return cache_Data_List_Lazy_Types_toList
}

var cache_Data_List_Lazy_Types_semigroupNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_semigroupNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_semigroupNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, as_prime__1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
_ = v1_2_0
// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
__local_var_3_1 := (v1_2_0).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
__local_var_4_2 := (v1_2_0).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_6_3 := Call_Data_List_Lazy_Types_toList(as_prime__1)
_ = __local_var_6_3
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_4 shape=App(Var) bindingType=(TypeVar a)
__local_var_8_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2)
_ = __local_var_8_4
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
goto end_branch_5
} else {

}
}
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr != nil) {
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_4.UnsafePtr).V1, __local_var_6_3)})
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
}))}))}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_semigroupNonEmptyList
}

var cache_Data_List_Lazy_Types_traversableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableList sync.Once
func Get_Data_List_Lazy_Types_traversableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableList.Do(func() {
		cache_Data_List_Lazy_Types_traversableList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(TypeApp (TypeVar m) [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])])
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_4_2
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_6_3_31 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_6_3_31
var go__go_6_3_31 gopurs_runtime.Value
_ = go__go_6_3_31
Call_local_Data_List_Lazy_Types_go__go_6_3_31 = func(b_7_loop gopurs_runtime.Value, xs_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_6_3_31:
for {
if false { continue go__go_6_3_31 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_4 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_9_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_4
var __t5 gopurs_runtime.Value
{
if (v_9_4 == nil) {
__t5 = b_7
goto end_branch_5
} else {

}
}
{
if (v_9_4 != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, (v_9_4).V0)), b_7)
xs_8_loop = (v_9_4).V1
continue go__go_6_3_31
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_6_3_31 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_6_3_31(b_7_loop_val, xs_8_loop_val)
})
})
var Call_local_Data_List_Lazy_Types_go__go_7_6_32 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_7_6_32
var go__go_7_6_32 gopurs_runtime.Value
_ = go__go_7_6_32
Call_local_Data_List_Lazy_Types_go__go_7_6_32 = func(b_8_loop gopurs_runtime.Value, xs_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_7_6_32:
for {
if false { continue go__go_7_6_32 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_7 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_10_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_7
var __t9 gopurs_runtime.Value
{
if (v_10_7 == nil) {
__t9 = b_8
goto end_branch_9
} else {

}
}
{
if (v_10_7 != nil) {
// TAST (Let): __local_var_11_8 shape=Other bindingType=(TypeVar a)
__local_var_11_8 := (v_10_7).V0
_ = __local_var_11_8
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_11_8, b_8}))}
}))
xs_9_loop = (v_10_7).V1
continue go__go_7_6_32
__t9 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}
go__go_7_6_32 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_7_6_32(b_8_loop_val, xs_9_loop_val)
})
})
return Call_local_Data_List_Lazy_Types_go__go_6_3_31(__local_var_4_2, Call_local_Data_List_Lazy_Types_go__go_7_6_32(Get_Data_List_Lazy_Types_nil(), xs_5))
})
})
})}))}
	})
	return cache_Data_List_Lazy_Types_traversableList
}

var cache_Data_List_Lazy_Types_traversableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmpty sync.Once
func Get_Data_List_Lazy_Types_traversableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_4249989635_3043886126(Rebox_Data_List_Lazy_Types_3043886126_4249989635(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))})))))}
	})
	return cache_Data_List_Lazy_Types_traversableNonEmpty
}

var cache_Data_List_Lazy_Types_traversableNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_traversableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_3))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
})
})}))}
	})
	return cache_Data_List_Lazy_Types_traversableNonEmptyList
}

var cache_Data_List_Lazy_Types_traversableWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexList sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(TypeApp (TypeVar m) [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])])
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_4_2
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_6_4_33 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_6_4_33
var go__go_6_4_33 gopurs_runtime.Value
_ = go__go_6_4_33
Call_local_Data_List_Lazy_Types_go__go_6_4_33 = func(b_7_loop gopurs_runtime.Value, xs_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_6_4_33:
for {
if false { continue go__go_6_4_33 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_5 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_9_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_5
var __t8 gopurs_runtime.Value
{
if (v_9_5 == nil) {
__t8 = b_7
goto end_branch_8
} else {

}
}
{
if (v_9_5 != nil) {
// TAST (Let): __local_var_10_6 shape=Other bindingType=Any
__local_var_10_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1
_ = __local_var_10_6
// TAST (Let): __local_var_11_7 shape=Other bindingType=(TypeVar a)
__local_var_11_7 := (v_9_5).V0
_ = __local_var_11_7
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_11_7, __local_var_10_6}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_8_loop = (v_9_5).V1
continue go__go_6_4_33
__t8 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}
go__go_6_4_33 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_6_4_33(b_7_loop_val, xs_8_loop_val)
})
})
// TAST (Let): v_6_3 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
v_6_3 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_6_4_33(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, xs_5)))
_ = v_6_3
var Call_local_Data_List_Lazy_Types_go__go_7_9_34 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_7_9_34
var go__go_7_9_34 gopurs_runtime.Value
_ = go__go_7_9_34
Call_local_Data_List_Lazy_Types_go__go_7_9_34 = func(b_8_loop gopurs_runtime.Value, xs_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_7_9_34:
for {
if false { continue go__go_7_9_34 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_10 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_10
var __t11 gopurs_runtime.Value
{
if (v_10_10 == nil) {
__t11 = b_8
goto end_branch_11
} else {

}
}
{
if (v_10_10 != nil) {
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_8.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply2(f_3, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_8.UnsafePtr).V0.IntVal) - (int64(1))), (v_10_10).V0)), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
xs_9_loop = (v_10_10).V1
continue go__go_7_9_34
__t11 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}
go__go_7_9_34 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_7_9_34(b_8_loop_val, xs_9_loop_val)
})
})
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Lazy_Types_go__go_7_9_34(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_6_3).V0), __local_var_4_2}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_6_3).V1).UnsafePtr).V1
})
})
})})))}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexList
}

var cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexNonEmpty sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_525353375_1812164904(Rebox_Data_List_Lazy_Types_1812164904_525353375(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableWithIndexList())))})))))}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty
}

var cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableNonEmptyList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_traversableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableWithIndexList())))}), "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 int64
{
if (x_4.Type == 9 && x_4.IntVal == 930809136 && x_4.UnsafePtr == nil) {
__t1 = int64(0)
goto end_branch_1
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 930809136 && x_4.UnsafePtr != nil) {
__t1 = (int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_4.UnsafePtr).V0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, gopurs_runtime.Int(__t1))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
})
})})))}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_unfoldable1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1List sync.Once
func Get_Data_List_Lazy_Types_unfoldable1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1List.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1List = func() gopurs_runtime.Value {
var go__go_0_0_35 gopurs_runtime.Value
_ = go__go_0_0_35
// FALLBACK TCO: isLoop=false len=1
go__go_0_0_35 = gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])])
v1_4_1 := Rebox_Data_List_Lazy_Types_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, b_2)))
_ = v1_4_1
var __t7 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v1_4_1).V1
if (__t_tag_2 != nil) {
// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
__local_var_5_3 := (v1_4_1).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=Any
__local_var_6_4 := gopurs_runtime.Apply2(go__go_0_0_35, f_1, ((v1_4_1).V1).V0)
_ = __local_var_6_4
__t7 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_3, __local_var_6_4}))}
}))
goto end_branch_7
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v1_4_1).V1
if (__t_tag_5 == nil) {
// TAST (Let): __local_var_5_6 shape=Other bindingType=Any
__local_var_5_6 := (v1_4_1).V0
_ = __local_var_5_6
__t7 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_6, Get_Data_List_Lazy_Types_nil()}))}
}))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t7)
}))
})
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, go__go_0_0_35}))}
}()
	})
	return cache_Data_List_Lazy_Types_unfoldable1List
}

var cache_Data_List_Lazy_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldableList sync.Once
func Get_Data_List_Lazy_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldableList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldableList = func() gopurs_runtime.Value {
var go__go_0_0_36 gopurs_runtime.Value
_ = go__go_0_0_36
// FALLBACK TCO: isLoop=false len=1
go__go_0_0_36 = gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
v1_4_1 := Rebox_Data_List_Lazy_Types_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, b_2)))
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (v1_4_1 == nil) {
__t4 = Get_Data_List_Lazy_Types_nil()
goto end_branch_4
} else {

}
}
{
if (v1_4_1 != nil) {
// TAST (Let): __local_var_5_2 shape=Other bindingType=Any
__local_var_5_2 := ((v1_4_1).V0).V0
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=Any
__local_var_6_3 := gopurs_runtime.Apply2(go__go_0_0_36, f_1, ((v1_4_1).V0).V1)
_ = __local_var_6_3
__t4 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_2, __local_var_6_3}))}
}))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t4)
}))
})
return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldable1List()))}
}), go__go_0_0_36}))}
}()
	})
	return cache_Data_List_Lazy_Types_unfoldableList
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmpty sync.Once
func Get_Data_List_Lazy_Types_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2751875267_2187088110((&Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])])
__local_var_2_1 := Rebox_Data_List_Lazy_Types_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, b_1)))
_ = __local_var_2_1
var go__go_3_2_37 gopurs_runtime.Value
_ = go__go_3_2_37
// FALLBACK TCO: isLoop=false len=1
go__go_3_2_37 = gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_7_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
v1_7_3 := Rebox_Data_List_Lazy_Types_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_4, b_5)))
_ = v1_7_3
var __t6 gopurs_runtime.Value
{
if (v1_7_3 == nil) {
__t6 = Get_Data_List_Lazy_Types_nil()
goto end_branch_6
} else {

}
}
{
if (v1_7_3 != nil) {
// TAST (Let): __local_var_8_4 shape=Other bindingType=Any
__local_var_8_4 := ((v1_7_3).V0).V0
_ = __local_var_8_4
// TAST (Let): __local_var_9_5 shape=App(Other) bindingType=Any
__local_var_9_5 := gopurs_runtime.Apply2(go__go_3_2_37, f_4, ((v1_7_3).V0).V1)
_ = __local_var_9_5
__t6 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_8_4, __local_var_9_5}))}
}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t6)
}))
})
// TAST (Let): __local_var_2_0 shape=Let(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeApp (TypeVar f) [(TypeVar a)])])
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(__local_var_2_1).V0, gopurs_runtime.Apply2(go__go_3_2_37, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_4)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((__local_var_2_1).V1)})}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_2_0).V0, (__local_var_2_0).V1}))}
})})))}
	})
	return cache_Data_List_Lazy_Types_unfoldable1NonEmpty
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_unfoldable1NonEmpty()).V0), f_0, b_1)))}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_unfoldable1NonEmptyList
}

var cache_Data_List_Lazy_Types_comonadNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_comonadNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_comonadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_comonadNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_comonadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_extendNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V0
})}))}
	})
	return cache_Data_List_Lazy_Types_comonadNonEmptyList
}

var cache_Data_List_Lazy_Types_monadList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadList sync.Once
func Get_Data_List_Lazy_Types_monadList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadList.Do(func() {
		cache_Data_List_Lazy_Types_monadList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applicativeList()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()))}
})}))}
	})
	return cache_Data_List_Lazy_Types_monadList
}

var cache_Data_List_Lazy_Types_bindList gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList sync.Once
func Get_Data_List_Lazy_Types_bindList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList.Do(func() {
		cache_Data_List_Lazy_Types_bindList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyList()))}
}), gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_4_1 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_5_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V1, f_1)
_ = __local_var_5_2
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_3 shape=App(Var) bindingType=(TypeVar a)
__local_var_7_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_7_3
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_7_3.Type == 9 && __local_var_7_3.IntVal == 218341868 && __local_var_7_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2))
goto end_branch_4
} else {

}
}
{
if (__local_var_7_3.Type == 9 && __local_var_7_3.IntVal == 218341868 && __local_var_7_3.UnsafePtr != nil) {
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_3.UnsafePtr).V1, __local_var_5_2)})
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))))
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_bindList
}

var cache_Data_List_Lazy_Types_applyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList sync.Once
func Get_Data_List_Lazy_Types_applyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList.Do(func() {
		cache_Data_List_Lazy_Types_applyList = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monadList()).V1), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monadList()).V0), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime__4, a_prime__5))
}))
}))
})}))}
}()
	})
	return cache_Data_List_Lazy_Types_applyList
}

var cache_Data_List_Lazy_Types_applicativeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeList sync.Once
func Get_Data_List_Lazy_Types_applicativeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeList.Do(func() {
		cache_Data_List_Lazy_Types_applicativeList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyList()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()}))}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_applicativeList
}

var cache_Data_List_Lazy_Types_applyNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_applyNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_applyNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])
v2_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_1))
_ = v2_2_0
// TAST (Let): v3_3_1 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (Func [(TypeVar a)] (TypeVar b))])
v3_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
_ = v3_3_1
// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
__local_var_4_2 := (v2_2_0).V0
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
__local_var_5_3 := (v2_2_0).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 shape=Other bindingType=Any
__local_var_6_4 := (v3_3_1).V0
_ = __local_var_6_4
// TAST (Let): __local_var_7_5 shape=Other bindingType=Any
__local_var_7_5 := (v3_3_1).V1
_ = __local_var_7_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_7 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_9_7 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_2, Get_Data_List_Lazy_Types_nil()}))}
}))
_ = __local_var_9_7
var Call_local_Data_List_Lazy_Types___local_var_10_8 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types___local_var_10_8
var __local_var_10_8 gopurs_runtime.Value
_ = __local_var_10_8
Call_local_Data_List_Lazy_Types___local_var_10_8 = func(f_prime__10_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_prime__10 gopurs_runtime.Value = f_prime__10_loop
_ = f_prime__10
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_9 shape=App(Var) bindingType=(TypeVar a)
__local_var_12_9 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_7)
_ = __local_var_12_9
var __t19 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_12_9.Type == 9 && __local_var_12_9.IntVal == 218341868 && __local_var_12_9.UnsafePtr == nil) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_19
} else {

}
}
{
if (__local_var_12_9.Type == 9 && __local_var_12_9.IntVal == 218341868 && __local_var_12_9.UnsafePtr != nil) {
// TAST (Let): __local_var_13_11 shape=App(Other) bindingType=(TypeVar b)
__local_var_13_11 := gopurs_runtime.Apply(f_prime__10, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_12_9.UnsafePtr).V0)
_ = __local_var_13_11
// TAST (Let): __local_var_13_10 shape=Let(App(Var)) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_13_10 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_13_11, Get_Data_List_Lazy_Types_nil()}))}
}))
_ = __local_var_13_10
// TAST (Let): __local_var_14_12 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_14_12 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_12_9.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_13 shape=App(Other) bindingType=(TypeVar b)
__local_var_15_13 := gopurs_runtime.Apply(f_prime__10, a_prime__14)
_ = __local_var_15_13
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_15_13, Get_Data_List_Lazy_Types_nil()}))}
}))
}))
_ = __local_var_14_12
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_14 shape=App(Var) bindingType=(TypeVar a)
__local_var_16_14 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_10)
_ = __local_var_16_14
var __t18 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_16_14.Type == 9 && __local_var_16_14.IntVal == 218341868 && __local_var_16_14.UnsafePtr == nil) {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_12))
goto end_branch_18
} else {

}
}
{
if (__local_var_16_14.Type == 9 && __local_var_16_14.IntVal == 218341868 && __local_var_16_14.UnsafePtr != nil) {
// TAST (Let): __local_var_17_15 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_17_15 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_16_14.UnsafePtr).V1
_ = __local_var_17_15
__t18 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_16_14.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_16 shape=App(Var) bindingType=(TypeVar a)
__local_var_19_16 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_15)
_ = __local_var_19_16
var __t17 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_19_16.Type == 9 && __local_var_19_16.IntVal == 218341868 && __local_var_19_16.UnsafePtr == nil) {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_12))
goto end_branch_17
} else {

}
}
{
if (__local_var_19_16.Type == 9 && __local_var_19_16.IntVal == 218341868 && __local_var_19_16.UnsafePtr != nil) {
__t17 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_19_16.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_19_16.UnsafePtr).V1, __local_var_14_12)})
goto end_branch_17
} else {

}
}
{
__t17 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t17)}
}))})
goto end_branch_18
} else {

}
}
{
__t18 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t18)}
}))))
goto end_branch_19
} else {

}
}
{
__t19 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t19)}
}))
}
__local_var_10_8 = gopurs_runtime.Func(func(f_prime__10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types___local_var_10_8(f_prime__10_loop_val)
})
// TAST (Let): __local_var_9_6 shape=Let(Let(App(Var))) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_9_6 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_20 shape=App(Var) bindingType=(TypeVar a)
__local_var_12_20 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_5)
_ = __local_var_12_20
var __t28 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_12_20.Type == 9 && __local_var_12_20.IntVal == 218341868 && __local_var_12_20.UnsafePtr == nil) {
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_28
} else {

}
}
{
if (__local_var_12_20.Type == 9 && __local_var_12_20.IntVal == 218341868 && __local_var_12_20.UnsafePtr != nil) {
// TAST (Let): __local_var_13_21 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_13_21 := Call_local_Data_List_Lazy_Types___local_var_10_8((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_12_20.UnsafePtr).V0)
_ = __local_var_13_21
// TAST (Let): __local_var_14_22 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_14_22 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_12_20.UnsafePtr).V1, __local_var_10_8)
_ = __local_var_14_22
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_23 shape=App(Var) bindingType=(TypeVar a)
__local_var_16_23 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_21)
_ = __local_var_16_23
var __t27 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_16_23.Type == 9 && __local_var_16_23.IntVal == 218341868 && __local_var_16_23.UnsafePtr == nil) {
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_22))
goto end_branch_27
} else {

}
}
{
if (__local_var_16_23.Type == 9 && __local_var_16_23.IntVal == 218341868 && __local_var_16_23.UnsafePtr != nil) {
// TAST (Let): __local_var_17_24 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_17_24 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_16_23.UnsafePtr).V1
_ = __local_var_17_24
__t27 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_16_23.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_25 shape=App(Var) bindingType=(TypeVar a)
__local_var_19_25 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_24)
_ = __local_var_19_25
var __t26 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_19_25.Type == 9 && __local_var_19_25.IntVal == 218341868 && __local_var_19_25.UnsafePtr == nil) {
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_22))
goto end_branch_26
} else {

}
}
{
if (__local_var_19_25.Type == 9 && __local_var_19_25.IntVal == 218341868 && __local_var_19_25.UnsafePtr != nil) {
__t26 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_19_25.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_19_25.UnsafePtr).V1, __local_var_14_22)})
goto end_branch_26
} else {

}
}
{
__t26 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_26:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t26)}
}))})
goto end_branch_27
} else {

}
}
{
__t27 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t27)}
}))))
goto end_branch_28
} else {

}
}
{
__t28 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t28)}
}))
_ = __local_var_9_6
// TAST (Let): __local_var_10_30 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(Func [(TypeVar a)] (TypeVar b))])])
__local_var_10_30 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_4, __local_var_7_5}))}
}))
_ = __local_var_10_30
var Call_local_Data_List_Lazy_Types___local_var_11_31 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types___local_var_11_31
var __local_var_11_31 gopurs_runtime.Value
_ = __local_var_11_31
Call_local_Data_List_Lazy_Types___local_var_11_31 = func(f_prime__11_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_prime__11 gopurs_runtime.Value = f_prime__11_loop
_ = f_prime__11
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_32 shape=App(Var) bindingType=(TypeVar a)
__local_var_13_32 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3)
_ = __local_var_13_32
var __t42 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_13_32.Type == 9 && __local_var_13_32.IntVal == 218341868 && __local_var_13_32.UnsafePtr == nil) {
__t42 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_42
} else {

}
}
{
if (__local_var_13_32.Type == 9 && __local_var_13_32.IntVal == 218341868 && __local_var_13_32.UnsafePtr != nil) {
// TAST (Let): __local_var_14_34 shape=App(Other) bindingType=(TypeVar b)
__local_var_14_34 := gopurs_runtime.Apply(f_prime__11, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_13_32.UnsafePtr).V0)
_ = __local_var_14_34
// TAST (Let): __local_var_14_33 shape=Let(App(Var)) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_14_33 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_14_34, Get_Data_List_Lazy_Types_nil()}))}
}))
_ = __local_var_14_33
// TAST (Let): __local_var_15_35 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_15_35 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_13_32.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime__15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_36 shape=App(Other) bindingType=(TypeVar b)
__local_var_16_36 := gopurs_runtime.Apply(f_prime__11, a_prime__15)
_ = __local_var_16_36
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_16_36, Get_Data_List_Lazy_Types_nil()}))}
}))
}))
_ = __local_var_15_35
__t42 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_37 shape=App(Var) bindingType=(TypeVar a)
__local_var_17_37 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_33)
_ = __local_var_17_37
var __t41 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_17_37.Type == 9 && __local_var_17_37.IntVal == 218341868 && __local_var_17_37.UnsafePtr == nil) {
__t41 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_35))
goto end_branch_41
} else {

}
}
{
if (__local_var_17_37.Type == 9 && __local_var_17_37.IntVal == 218341868 && __local_var_17_37.UnsafePtr != nil) {
// TAST (Let): __local_var_18_38 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_18_38 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_17_37.UnsafePtr).V1
_ = __local_var_18_38
__t41 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_17_37.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_39 shape=App(Var) bindingType=(TypeVar a)
__local_var_20_39 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_38)
_ = __local_var_20_39
var __t40 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_20_39.Type == 9 && __local_var_20_39.IntVal == 218341868 && __local_var_20_39.UnsafePtr == nil) {
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_35))
goto end_branch_40
} else {

}
}
{
if (__local_var_20_39.Type == 9 && __local_var_20_39.IntVal == 218341868 && __local_var_20_39.UnsafePtr != nil) {
__t40 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_20_39.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_20_39.UnsafePtr).V1, __local_var_15_35)})
goto end_branch_40
} else {

}
}
{
__t40 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_40:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t40)}
}))})
goto end_branch_41
} else {

}
}
{
__t41 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_41:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t41)}
}))))
goto end_branch_42
} else {

}
}
{
__t42 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_42:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t42)}
}))
}
__local_var_11_31 = gopurs_runtime.Func(func(f_prime__11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types___local_var_11_31(f_prime__11_loop_val)
})
// TAST (Let): __local_var_10_29 shape=Let(Let(App(Var))) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_10_29 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_43 shape=App(Var) bindingType=(TypeVar a)
__local_var_13_43 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_30)
_ = __local_var_13_43
var __t51 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_13_43.Type == 9 && __local_var_13_43.IntVal == 218341868 && __local_var_13_43.UnsafePtr == nil) {
__t51 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_51
} else {

}
}
{
if (__local_var_13_43.Type == 9 && __local_var_13_43.IntVal == 218341868 && __local_var_13_43.UnsafePtr != nil) {
// TAST (Let): __local_var_14_44 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_14_44 := Call_local_Data_List_Lazy_Types___local_var_11_31((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_13_43.UnsafePtr).V0)
_ = __local_var_14_44
// TAST (Let): __local_var_15_45 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_15_45 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_13_43.UnsafePtr).V1, __local_var_11_31)
_ = __local_var_15_45
__t51 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_46 shape=App(Var) bindingType=(TypeVar a)
__local_var_17_46 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_44)
_ = __local_var_17_46
var __t50 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_17_46.Type == 9 && __local_var_17_46.IntVal == 218341868 && __local_var_17_46.UnsafePtr == nil) {
__t50 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_45))
goto end_branch_50
} else {

}
}
{
if (__local_var_17_46.Type == 9 && __local_var_17_46.IntVal == 218341868 && __local_var_17_46.UnsafePtr != nil) {
// TAST (Let): __local_var_18_47 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_18_47 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_17_46.UnsafePtr).V1
_ = __local_var_18_47
__t50 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_17_46.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_48 shape=App(Var) bindingType=(TypeVar a)
__local_var_20_48 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_47)
_ = __local_var_20_48
var __t49 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_20_48.Type == 9 && __local_var_20_48.IntVal == 218341868 && __local_var_20_48.UnsafePtr == nil) {
__t49 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_45))
goto end_branch_49
} else {

}
}
{
if (__local_var_20_48.Type == 9 && __local_var_20_48.IntVal == 218341868 && __local_var_20_48.UnsafePtr != nil) {
__t49 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_20_48.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_20_48.UnsafePtr).V1, __local_var_15_45)})
goto end_branch_49
} else {

}
}
{
__t49 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_49:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t49)}
}))})
goto end_branch_50
} else {

}
}
{
__t50 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_50:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t50)}
}))))
goto end_branch_51
} else {

}
}
{
__t51 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_51:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t51)}
}))
_ = __local_var_10_29
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_52 shape=App(Var) bindingType=(TypeVar a)
__local_var_12_52 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_6)
_ = __local_var_12_52
var __t53 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_12_52.Type == 9 && __local_var_12_52.IntVal == 218341868 && __local_var_12_52.UnsafePtr == nil) {
__t53 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_29))
goto end_branch_53
} else {

}
}
{
if (__local_var_12_52.Type == 9 && __local_var_12_52.IntVal == 218341868 && __local_var_12_52.UnsafePtr != nil) {
__t53 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_12_52.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_12_52.UnsafePtr).V1, __local_var_10_29)})
goto end_branch_53
} else {

}
}
{
__t53 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_53:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t53)}
}))}))}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_applyNonEmptyList
}

var cache_Data_List_Lazy_Types_bindNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_bindNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_bindNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
_ = v1_2_0
// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
__local_var_3_1 := (v1_2_0).V1
_ = __local_var_3_1
// TAST (Let): v2_4_2 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar b)])
v2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_1, (v1_2_0).V0)))
_ = v2_4_2
// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
__local_var_5_3 := (v2_4_2).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 shape=Other bindingType=Any
__local_var_6_4 := (v2_4_2).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_8_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_6 shape=App(Var) bindingType=(TypeVar a)
__local_var_9_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1)
_ = __local_var_9_6
var __t14 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_9_6.Type == 9 && __local_var_9_6.IntVal == 218341868 && __local_var_9_6.UnsafePtr == nil) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_14
} else {

}
}
{
if (__local_var_9_6.Type == 9 && __local_var_9_6.IntVal == 218341868 && __local_var_9_6.UnsafePtr != nil) {
// TAST (Let): __local_var_10_7 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_10_7 := Call_Data_List_Lazy_Types_toList(gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_6.UnsafePtr).V0))
_ = __local_var_10_7
// TAST (Let): __local_var_11_8 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_11_8 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_6.UnsafePtr).V1, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_toList(gopurs_runtime.Apply(f_1, x_11))
}))
_ = __local_var_11_8
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_9 shape=App(Var) bindingType=(TypeVar a)
__local_var_13_9 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_7)
_ = __local_var_13_9
var __t13 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_13_9.Type == 9 && __local_var_13_9.IntVal == 218341868 && __local_var_13_9.UnsafePtr == nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_13
} else {

}
}
{
if (__local_var_13_9.Type == 9 && __local_var_13_9.IntVal == 218341868 && __local_var_13_9.UnsafePtr != nil) {
// TAST (Let): __local_var_14_10 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_14_10 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_13_9.UnsafePtr).V1
_ = __local_var_14_10
__t13 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_13_9.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_11 shape=App(Var) bindingType=(TypeVar a)
__local_var_16_11 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_10)
_ = __local_var_16_11
var __t12 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_16_11.Type == 9 && __local_var_16_11.IntVal == 218341868 && __local_var_16_11.UnsafePtr == nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_12
} else {

}
}
{
if (__local_var_16_11.Type == 9 && __local_var_16_11.IntVal == 218341868 && __local_var_16_11.UnsafePtr != nil) {
__t12 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_16_11.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_16_11.UnsafePtr).V1, __local_var_11_8)})
goto end_branch_12
} else {

}
}
{
__t12 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t12)}
}))})
goto end_branch_13
} else {

}
}
{
__t13 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t13)}
}))))
goto end_branch_14
} else {

}
}
{
__t14 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t14)}
}))
_ = __local_var_8_5
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_3, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_15 shape=App(Var) bindingType=(TypeVar a)
__local_var_10_15 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4)
_ = __local_var_10_15
var __t16 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_10_15.Type == 9 && __local_var_10_15.IntVal == 218341868 && __local_var_10_15.UnsafePtr == nil) {
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_5))
goto end_branch_16
} else {

}
}
{
if (__local_var_10_15.Type == 9 && __local_var_10_15.IntVal == 218341868 && __local_var_10_15.UnsafePtr != nil) {
__t16 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_10_15.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_10_15.UnsafePtr).V1, __local_var_8_5)})
goto end_branch_16
} else {

}
}
{
__t16 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t16)}
}))}))}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_bindNonEmptyList
}

var cache_Data_List_Lazy_Types_altNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_altNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_altNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, as_prime__1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
_ = v1_2_0
// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
__local_var_3_1 := (v1_2_0).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
__local_var_4_2 := (v1_2_0).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_6_3 := Call_Data_List_Lazy_Types_toList(as_prime__1)
_ = __local_var_6_3
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_4 shape=App(Var) bindingType=(TypeVar a)
__local_var_8_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2)
_ = __local_var_8_4
var __t8 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
goto end_branch_8
} else {

}
}
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr != nil) {
// TAST (Let): __local_var_9_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_9_5 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_4.UnsafePtr).V1
_ = __local_var_9_5
__t8 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_4.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_6 shape=App(Var) bindingType=(TypeVar a)
__local_var_11_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_5)
_ = __local_var_11_6
var __t7 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_11_6.Type == 9 && __local_var_11_6.IntVal == 218341868 && __local_var_11_6.UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
goto end_branch_7
} else {

}
}
{
if (__local_var_11_6.Type == 9 && __local_var_11_6.IntVal == 218341868 && __local_var_11_6.UnsafePtr != nil) {
__t7 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_11_6.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_11_6.UnsafePtr).V1, __local_var_6_3)})
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t7)}
}))})
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t8)}
}))}))}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_altNonEmptyList
}

var cache_Data_List_Lazy_Types_altList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altList sync.Once
func Get_Data_List_Lazy_Types_altList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altList.Do(func() {
		cache_Data_List_Lazy_Types_altList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V1, ys_1)})
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_altList
}

var cache_Data_List_Lazy_Types_plusList gopurs_runtime.Value
var once_Data_List_Lazy_Types_plusList sync.Once
func Get_Data_List_Lazy_Types_plusList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_plusList.Do(func() {
		cache_Data_List_Lazy_Types_plusList = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_altList()))}
}), Get_Data_List_Lazy_Types_nil()}))}
	})
	return cache_Data_List_Lazy_Types_plusList
}

var cache_Data_List_Lazy_Types_alternativeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_alternativeList sync.Once
func Get_Data_List_Lazy_Types_alternativeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_alternativeList.Do(func() {
		cache_Data_List_Lazy_Types_alternativeList = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applicativeList()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_plusList()))}
})}))}
	})
	return cache_Data_List_Lazy_Types_alternativeList
}

var cache_Data_List_Lazy_Types_monadPlusList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadPlusList sync.Once
func Get_Data_List_Lazy_Types_monadPlusList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadPlusList.Do(func() {
		cache_Data_List_Lazy_Types_monadPlusList = gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_alternativeList()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monadList()))}
})}))}
	})
	return cache_Data_List_Lazy_Types_monadPlusList
}

var cache_Data_List_Lazy_Types_applicativeNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_applicativeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_applicativeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()}))}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_applicativeNonEmptyList
}

var cache_Data_List_Lazy_Types_monadNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_monadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_monadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applicativeNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindNonEmptyList()))}
})}))}
	})
	return cache_Data_List_Lazy_Types_monadNonEmptyList
}

type Constructor_Data_List_Lazy_Types_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Data_List_Lazy_Types_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}


func Call_Data_List_Lazy_Types_List(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Lazy_Types_NonEmptyList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Lazy_Types_step(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_eqNonEmpty(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2176830691_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_7 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_1 gopurs_runtime.Value
_ = go__go_3_0_1
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_1 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = (__t_tag_5 == nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
__t_and_7 = (gopurs_runtime.Apply2(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_7)
})})))}
}

func Call_Data_List_Lazy_Types_eq1(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var go__go_3_0_2 gopurs_runtime.Value
_ = go__go_3_0_2
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_2 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = (__t_tag_5 == nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_3_0_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
return (gopurs_runtime.Apply2(go__go_3_0_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0)
}

func Call_Data_List_Lazy_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_4 gopurs_runtime.Value
_ = go__go_3_0_4
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_4 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = (__t_tag_5 == nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_3_0_4, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_4, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})}))}
}

func Call_Data_List_Lazy_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_eqLazy(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2176830691_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_7 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_5 gopurs_runtime.Value
_ = go__go_3_0_5
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_5 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = (__t_tag_5 == nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_3_0_5, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
__t_and_7 = (gopurs_runtime.Apply2(go__go_3_0_5, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_7)
})})))})))}
}

func Call_Data_List_Lazy_Types_compare1(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var Call_local_Data_List_Lazy_Types_go__go_3_0_7 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Lazy_Types_go__go_3_0_7
var go__go_3_0_7 gopurs_runtime.Value
_ = go__go_3_0_7
Call_local_Data_List_Lazy_Types_go__go_3_0_7 = func(v_4_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_3_0_7:
for {
if false { continue go__go_3_0_7 }
var v_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t4 uint32
{
if (v_4 == nil) {
var __t1 uint32
{
if (v1_5 == nil) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t4 = __t1
goto end_branch_4
} else {

}
}
{
if (v1_5 == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((v_4 != nil)) && ((v1_5 != nil)) {
// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_6_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (v_4).V0, (v1_5).V0).IntVal)
_ = v2_6_2
var __t3 uint32
{
if (v2_6_2 == 902936544) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_4).V1))
v1_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v1_5).V1))
continue go__go_3_0_7
__t3 = func() uint32 { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = v2_6_2
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_3_0_7 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_3_0_7(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val))), UnsafePtr: nil}
})
})
return Call_local_Data_List_Lazy_Types_go__go_3_0_7(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))
}

func Call_Data_List_Lazy_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqList1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
eqList1_1_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_8 gopurs_runtime.Value
_ = go__go_4_2_8
// FALLBACK TCO: isLoop=false len=1
go__go_4_2_8 = gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
var __t_tag_6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_5)
if (__t_tag_6 == nil) {
var __t_tag_7 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_6)
__t8 = (__t_tag_7 == nil)
goto end_branch_8
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_5)
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_6)
__t_and_5 = ((__t_tag_4 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_4_2_8, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t8 = __t_and_5
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_4_2_8, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))}).IntVal) != (0))
})})
_ = eqList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_0)}
}), gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_4_9_9 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Lazy_Types_go__go_4_9_9
var go__go_4_9_9 gopurs_runtime.Value
_ = go__go_4_9_9
Call_local_Data_List_Lazy_Types_go__go_4_9_9 = func(v_5_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_4_9_9:
for {
if false { continue go__go_4_9_9 }
var v_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t13 uint32
{
if (v_5 == nil) {
var __t10 uint32
{
if (v1_6 == nil) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t13 = __t10
goto end_branch_13
} else {

}
}
{
if (v1_6 == nil) {
__t13 = 380165415
goto end_branch_13
} else {

}
}
{
if ((v_5 != nil)) && ((v1_6 != nil)) {
// TAST (Let): v2_7_11 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_7_11 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v_5).V0, (v1_6).V0).IntVal)
_ = v2_7_11
var __t12 uint32
{
if (v2_7_11 == 902936544) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_5).V1))
v1_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v1_6).V1))
continue go__go_4_9_9
__t12 = func() uint32 { panic("unreachable") }()
goto end_branch_12
} else {

}
}
{
__t12 = v2_7_11
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
__t13 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}
go__go_4_9_9 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_4_9_9(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val))), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_4_9_9(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))), UnsafePtr: nil}
})}))}
}

func Call_Data_List_Lazy_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_ordLazy(), gopurs_runtime.Apply2(Get_Data_NonEmpty_ordNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()))}, dictOrd_0))))}
}

func Call_Data_List_Lazy_Types_cons(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, xs_1}))}
}))
}

func Call_Data_List_Lazy_Types_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
_ = v_2_0
var __t4 string
{
if (v_2_0 == nil) {
__t4 = "(fromFoldable [])"
goto end_branch_4
} else {

}
}
{
if (v_2_0 != nil) {
var Call_local_Data_List_Lazy_Types_go__go_3_1_19 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_1_19
var go__go_3_1_19 gopurs_runtime.Value
_ = go__go_3_1_19
Call_local_Data_List_Lazy_Types_go__go_3_1_19 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_1_19:
for {
if false { continue go__go_3_1_19 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if (v_6_2 == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_6_2 != nil) {
b_4_loop = gopurs_runtime.Str(((b_4.StrVal()) + (",")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_6_2).V0).StrVal()))
xs_5_loop = (v_6_2).V1
continue go__go_3_1_19
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_3_1_19 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_1_19(b_4_loop_val, xs_5_loop_val)
})
})
__t4 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_0).V0).StrVal())) + (Call_local_Data_List_Lazy_Types_go__go_3_1_19(gopurs_runtime.Str(""), (v_2_0).V1).StrVal())) + ("])")
goto end_branch_4
} else {

}
}
{
__t4 = func() string { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Str(__t4)
})}))}
}

func Call_Data_List_Lazy_Types_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): __local_var_1_1 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
__local_var_1_1 := (&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
_ = v_2_2
var __t6 string
{
if (v_2_2 == nil) {
__t6 = "(fromFoldable [])"
goto end_branch_6
} else {

}
}
{
if (v_2_2 != nil) {
var Call_local_Data_List_Lazy_Types_go__go_3_3_20 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_3_20
var go__go_3_3_20 gopurs_runtime.Value
_ = go__go_3_3_20
Call_local_Data_List_Lazy_Types_go__go_3_3_20 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_3_20:
for {
if false { continue go__go_3_3_20 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_4 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_4
var __t5 gopurs_runtime.Value
{
if (v_6_4 == nil) {
__t5 = b_4
goto end_branch_5
} else {

}
}
{
if (v_6_4 != nil) {
b_4_loop = gopurs_runtime.Str(((b_4.StrVal()) + (",")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_6_4).V0).StrVal()))
xs_5_loop = (v_6_4).V1
continue go__go_3_3_20
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_3_3_20 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_3_20(b_4_loop_val, xs_5_loop_val)
})
})
__t6 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_2).V0).StrVal())) + (Call_local_Data_List_Lazy_Types_go__go_3_3_20(gopurs_runtime.Str(""), (v_2_2).V1).StrVal())) + ("])")
goto end_branch_6
} else {

}
}
{
__t6 = func() string { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Str(__t6)
})})
_ = __local_var_1_1
// TAST (Let): __local_var_2_7 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f), (TypeVar a)])])
__local_var_2_7 := (&Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(NonEmpty ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_1.V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
})})
_ = __local_var_2_7
// TAST (Let): showLazy_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])])])
showLazy_1_0 := (&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(defer \\_ -> ") + (gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_7.V0), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_3)).StrVal())) + (")"))
})})
_ = showLazy_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyList ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showLazy_1_0.V0), v_2).StrVal())) + (")"))
})}))}
}

func Call_Data_List_Lazy_Types_showStep(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList1_1_0 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
showList1_1_0 := (&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
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
var Call_local_Data_List_Lazy_Types_go__go_3_2_21 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_2_21
var go__go_3_2_21 gopurs_runtime.Value
_ = go__go_3_2_21
Call_local_Data_List_Lazy_Types_go__go_3_2_21 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_2_21:
for {
if false { continue go__go_3_2_21 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
continue go__go_3_2_21
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_2_21 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_2_21(b_4_loop_val, xs_5_loop_val)
})
})
__t5 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_1).V0).StrVal())) + (Call_local_Data_List_Lazy_Types_go__go_3_2_21(gopurs_runtime.Str(""), (v_2_1).V1).StrVal())) + ("])")
goto end_branch_5
} else {

}
}
{
__t5 = func() string { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Str(__t5)
})})
_ = showList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_928333203_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 string
{
var __t_tag_6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2)
if (__t_tag_6 == nil) {
__t8 = "Nil"
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2)
if (__t_tag_7 != nil) {
__t8 = (((("(") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" : ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(showList1_1_0.V0), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_8
} else {

}
}
{
__t8 = func() string { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Str(__t8)
})})))}
}

func Call_Data_List_Lazy_Types_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])
v2_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
_ = v2_2_0
// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
__local_var_3_1 := (v2_2_0).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
__local_var_4_2 := (v2_2_0).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_4_2}))}
})))
}))
}

func Rebox_Data_List_Lazy_Types_138441832_1728839155(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]{}
		out.V0 = in.V0.IntVal
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_138441832_3804580809(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_List_Lazy_Types_1468200963_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Lazy_Types_1728839155_138441832(in *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_1812164904_525353375(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_2048287455_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_2176830691_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Lazy_Types_2491554675_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_2751875267_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Lazy_Types_2773701683_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_2955889203_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_3043886126_4249989635(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_3071895939_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_List_Lazy_Types_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_List_Lazy_Types_3725484264_2048287455(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_4177771502_167870147(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_4249989635_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_525353375_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_80275103_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_928333203_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}


