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
		cache_Data_List_Lazy_Types_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_List_Lazy_Types_identity
}

var cache_Data_List_Lazy_Types_unwrap gopurs_runtime.Value
var once_Data_List_Lazy_Types_unwrap sync.Once
func Get_Data_List_Lazy_Types_unwrap() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unwrap.Do(func() {
		cache_Data_List_Lazy_Types_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
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
		cache_Data_List_Lazy_Types_step = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Lazy_force(), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})))
	})
	return cache_Data_List_Lazy_Types_step
}

var cache_Data_List_Lazy_Types_semigroupList gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList sync.Once
func Get_Data_List_Lazy_Types_semigroupList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_0)
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
_ = __local_var_4_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), ys_1))
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, ys_1)})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
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
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), f_0))
})}))}
	})
	return cache_Data_List_Lazy_Types_lazyList
}

var cache_Data_List_Lazy_Types_functorList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList sync.Once
func Get_Data_List_Lazy_Types_functorList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList.Do(func() {
		cache_Data_List_Lazy_Types_functorList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope114)])])
__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_1)
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a$scope106)
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
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0, f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_functorList
}

var cache_Data_List_Lazy_Types_functorNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmpty sync.Once
func Get_Data_List_Lazy_Types_functorNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1468200963_2812149806(Rebox_Data_List_Lazy_Types_2812149806_1468200963(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))})))))}
	})
	return cache_Data_List_Lazy_Types_functorNonEmpty
}

var cache_Data_List_Lazy_Types_functorNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_functorNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(Func [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope108)])] (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar b$scope109)]))
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}), "map"), f_0)
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
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
var go__go_3_0_0_cell *gopurs_runtime.Value
_ = go__go_3_0_0_cell
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_0 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
_ = __t_tag_4
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
_ = __t_tag_5
__t6 = (__t_tag_5 == nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
_ = __t_tag_1
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
_ = __t_tag_2
__t_and_3 = ((__t_tag_2 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2((*go__go_3_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
go__go_3_0_0_cell = &go__go_3_0_0
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), ys_2)))}).IntVal) != (0))
})}))}
	})
	return cache_Data_List_Lazy_Types_eq1List
}

var cache_Data_List_Lazy_Types_eqNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_eqNonEmpty sync.Once
func Get_Data_List_Lazy_Types_eqNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eqNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_eqNonEmpty = gopurs_runtime.Apply(Get_Data_NonEmpty_eqNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))})
	})
	return cache_Data_List_Lazy_Types_eqNonEmpty
}

var cache_Data_List_Lazy_Types_eq1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1 sync.Once
func Get_Data_List_Lazy_Types_eq1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1.Do(func() {
		cache_Data_List_Lazy_Types_eq1 = Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))
	})
	return cache_Data_List_Lazy_Types_eq1
}

var cache_Data_List_Lazy_Types_eq1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_eq1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_eq1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqNonEmpty1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope162)])])
eqNonEmpty1_1_0 := Rebox_Data_List_Lazy_Types_3790796878_2176830691(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_NonEmpty_eqNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))}, dictEq_0)))
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Lazy_eqLazy(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2176830691_3790796878(eqNonEmpty1_1_0))}))), v_2, v1_3).IntVal) != (0))
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
var Call_local_Data_List_Lazy_Types_go__go_3_0_1 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Lazy_Types_go__go_3_0_1
var go__go_3_0_1 gopurs_runtime.Value
_ = go__go_3_0_1
Call_local_Data_List_Lazy_Types_go__go_3_0_1 = func(v_4_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_3_0_1:
for {
if false { continue go__go_3_0_1 }
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
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_4).V1))
v1_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v1_5).V1))
continue go__go_3_0_1
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
go__go_3_0_1 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_3_0_1(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val))), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_3_0_1(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_1)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), ys_2)))), UnsafePtr: nil}
})}))}
	})
	return cache_Data_List_Lazy_Types_ord1List
}

var cache_Data_List_Lazy_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordNonEmpty sync.Once
func Get_Data_List_Lazy_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_ordNonEmpty = Call_Data_NonEmpty_ordNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()))})
	})
	return cache_Data_List_Lazy_Types_ordNonEmpty
}

var cache_Data_List_Lazy_Types_compare1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compare1 sync.Once
func Get_Data_List_Lazy_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compare1.Do(func() {
		cache_Data_List_Lazy_Types_compare1 = Call_Data_Ord_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()))
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
// TAST (Let): ordNonEmpty1_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope61)])])
ordNonEmpty1_1_0 := Rebox_Data_List_Lazy_Types_4177771502_167870147(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_NonEmpty_ordNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()))}), dictOrd_0)))
_ = ordNonEmpty1_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Lazy_ordLazy(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_167870147_4177771502(ordNonEmpty1_1_0))}))), v_2, v1_3).IntVal)), UnsafePtr: nil}
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
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope202)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=App(Var) bindingType=(TypeVar m$scope202)
mempty_2_1 := Call_Data_Monoid_mempty(dictMonoid_0)
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_1_2_2 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_1_2_2
var go__go_1_2_2 gopurs_runtime.Value
_ = go__go_1_2_2
Call_local_Data_List_Lazy_Types_go__go_1_2_2 = func(b_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_1_2_2:
for {
if false { continue go__go_1_2_2 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
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
continue go__go_1_2_2
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
go__go_1_2_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_1_2_2(b_2_loop_val, xs_3_loop_val)
})
})
return go__go_1_2_2
}), gopurs_runtime.Func3(func(op_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons(a_4, b_3)
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})}))}
	})
	return cache_Data_List_Lazy_Types_foldableList
}

var cache_Data_List_Lazy_Types_foldableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmpty sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_3071895939_1680800814(Rebox_Data_List_Lazy_Types_1680800814_3071895939(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))})))))}
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
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope175)])
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), l_1))
_ = v_2_0
var __t13 gopurs_runtime.Value
{
if (v_2_0 == nil) {
__t13 = Get_Data_List_Lazy_Types_nil()
goto end_branch_13
} else {

}
}
{
if (v_2_0 != nil) {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=(TypeVar b$scope174)
__local_var_3_1 := gopurs_runtime.Apply(f_0, l_1)
_ = __local_var_3_1
var Call_local_Data_List_Lazy_Types_go__go_4_2_3 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_2_3
var go__go_4_2_3 gopurs_runtime.Value
_ = go__go_4_2_3
Call_local_Data_List_Lazy_Types_go__go_4_2_3 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_2_3:
for {
if false { continue go__go_4_2_3 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_6))
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
// TAST (Let): acc_prime__10_6 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope175)])])
acc_prime__10_6 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v_7_3).V0, __local_var_8_4}))}
}))
_ = acc_prime__10_6
// TAST (Let): __local_var_11_7 shape=App(Other) bindingType=(TypeVar b$scope174)
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
				return gopurs_runtime.RecordDict2("acc", "val", orig.acc, orig.val)
				}()
xs_6_loop = (v_7_3).V1
continue go__go_4_2_3
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
go__go_4_2_3 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_2_3(b_5_loop_val, xs_6_loop_val)
})
})
var Call_local_Data_List_Lazy_Types_go__go_5_10_4 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_5_10_4
var go__go_5_10_4 gopurs_runtime.Value
_ = go__go_5_10_4
Call_local_Data_List_Lazy_Types_go__go_5_10_4 = func(b_6_loop gopurs_runtime.Value, xs_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_5_10_4:
for {
if false { continue go__go_5_10_4 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_11 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_7))
_ = v_8_11
var __t12 gopurs_runtime.Value
{
if (v_8_11 == nil) {
__t12 = b_6
goto end_branch_12
} else {

}
}
{
if (v_8_11 != nil) {
b_6_loop = Call_Data_List_Lazy_Types_cons((v_8_11).V0, b_6)
xs_7_loop = (v_8_11).V1
continue go__go_5_10_4
__t12 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}
go__go_5_10_4 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_5_10_4(b_6_loop_val, xs_7_loop_val)
})
})
// TAST (Let): __local_var_5_9 shape=Other bindingType=Any
__local_var_5_9 := gopurs_runtime.RecordGet(Call_local_Data_List_Lazy_Types_go__go_4_2_3(func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	val gopurs_runtime.Value
}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "val", orig.acc, orig.val)
				}(), Call_local_Data_List_Lazy_Types_go__go_5_10_4(Get_Data_List_Lazy_Types_nil(), (v_2_0).V1)), "val")
_ = __local_var_5_9
__t13 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_5_9}))}
}))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
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
var Call_local_Data_List_Lazy_Types_go__go_4_1_5 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_1_5
var go__go_4_1_5 gopurs_runtime.Value
_ = go__go_4_1_5
Call_local_Data_List_Lazy_Types_go__go_4_1_5 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_1_5:
for {
if false { continue go__go_4_1_5 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_7_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_6))
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
// TAST (Let): __local_var_10_5 shape=App(Other) bindingType=(TypeVar b$scope179)
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
				return gopurs_runtime.RecordDict2("acc", "val", orig.acc, orig.val)
				}()
xs_6_loop = (v_7_2).V1
continue go__go_4_1_5
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
go__go_4_1_5 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_1_5(b_5_loop_val, xs_6_loop_val)
})
})
var Call_local_Data_List_Lazy_Types_go__go_5_7_6 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_5_7_6
var go__go_5_7_6 gopurs_runtime.Value
_ = go__go_5_7_6
Call_local_Data_List_Lazy_Types_go__go_5_7_6 = func(b_6_loop gopurs_runtime.Value, xs_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_5_7_6:
for {
if false { continue go__go_5_7_6 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_8 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_7))
_ = v_8_8
var __t9 gopurs_runtime.Value
{
if (v_8_8 == nil) {
__t9 = b_6
goto end_branch_9
} else {

}
}
{
if (v_8_8 != nil) {
b_6_loop = Call_Data_List_Lazy_Types_cons((v_8_8).V0, b_6)
xs_7_loop = (v_8_8).V1
continue go__go_5_7_6
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
go__go_5_7_6 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_5_7_6(b_6_loop_val, xs_7_loop_val)
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(Call_local_Data_List_Lazy_Types_go__go_4_1_5(func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	val gopurs_runtime.Value
}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "val", orig.acc, orig.val)
				}(), Call_local_Data_List_Lazy_Types_go__go_5_7_6(Get_Data_List_Lazy_Types_nil(), __local_var_2_0)), "val")}))}
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
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldl"), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldr"), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
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
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope229)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=App(Var) bindingType=(TypeVar m$scope229)
mempty_2_1 := Call_Data_Monoid_mempty(dictMonoid_0)
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())).V2, gopurs_runtime.Func2(func(i_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_5), gopurs_runtime.Apply(f_3, gopurs_runtime.Int(i_4.IntVal)))
}), mempty_2_1)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_2_2_8 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_2_2_8
var go__go_2_2_8 gopurs_runtime.Value
_ = go__go_2_2_8
Call_local_Data_List_Lazy_Types_go__go_2_2_8 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_2_8:
for {
if false { continue go__go_2_2_8 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (v_5_3 == nil) {
__t4 = b_3
goto end_branch_4
} else {

}
}
{
if (v_5_3 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V1, (v_5_3).V0)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
xs_4_loop = (v_5_3).V1
continue go__go_2_2_8
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
go__go_2_2_8 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_2_2_8(b_3_loop_val, xs_4_loop_val)
})
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_snd(), gopurs_runtime.Apply(go__go_2_2_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), acc_1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_3_6_9 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_6_9
var go__go_3_6_9 gopurs_runtime.Value
_ = go__go_3_6_9
Call_local_Data_List_Lazy_Types_go__go_3_6_9 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_6_9:
for {
if false { continue go__go_3_6_9 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_7 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_5))
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
// TAST (Let): __local_var_7_8 shape=Other bindingType=Any
__local_var_7_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1
_ = __local_var_7_8
// TAST (Let): __local_var_8_9 shape=Other bindingType=(TypeVar a$scope199)
__local_var_8_9 := (v_6_7).V0
_ = __local_var_8_9
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_8_9, __local_var_7_8}))}
}))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
xs_5_loop = (v_6_7).V1
continue go__go_3_6_9
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
go__go_3_6_9 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_6_9(b_4_loop_val, xs_5_loop_val)
})
})
// TAST (Let): v_3_5 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope211)])])])
v_3_5 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_3_6_9(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, xs_2)))
_ = v_3_5
var Call_local_Data_List_Lazy_Types_go__go_4_11_10 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_4_11_10
var go__go_4_11_10 gopurs_runtime.Value
_ = go__go_4_11_10
Call_local_Data_List_Lazy_Types_go__go_4_11_10 = func(b_5_loop gopurs_runtime.Value, xs_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_11_10:
for {
if false { continue go__go_4_11_10 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_6))
_ = v_7_12
var __t13 gopurs_runtime.Value
{
if (v_7_12 == nil) {
__t13 = b_5
goto end_branch_13
} else {

}
}
{
if (v_7_12 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), (v_7_12).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
xs_6_loop = (v_7_12).V1
continue go__go_4_11_10
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
go__go_4_11_10 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_4_11_10(b_5_loop_val, xs_6_loop_val)
})
})
return Call_Data_Tuple_snd(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_4_11_10(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_3_5).V0), b_1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, (v_3_5).V1)))))
})})))}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexList
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmpty sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2048287455_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2048287455(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))})))))}
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
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_1, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.Int(int64(0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_3
if (__t_tag_3 == nil) {
__t5 = gopurs_runtime.Int(int64(0))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_4
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_6
if (__t_tag_6 == nil) {
__t8 = gopurs_runtime.Int(int64(0))
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_7
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
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
var Call_local_Data_List_Lazy_Types_go__go_2_1_11 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_2_1_11
var go__go_2_1_11 gopurs_runtime.Value
_ = go__go_2_1_11
Call_local_Data_List_Lazy_Types_go__go_2_1_11 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_1_11:
for {
if false { continue go__go_2_1_11 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
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
// TAST (Let): __local_var_7_4 shape=Other bindingType=(TypeVar a$scope199)
__local_var_7_4 := (v_5_2).V0
_ = __local_var_7_4
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_4, __local_var_6_3}))}
}))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
xs_4_loop = (v_5_2).V1
continue go__go_2_1_11
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
go__go_2_1_11 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_2_1_11(b_3_loop_val, xs_4_loop_val)
})
})
// TAST (Let): v_2_0 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope211)])])])
v_2_0 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_2_1_11(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, xs_1)))
_ = v_2_0
var Call_local_Data_List_Lazy_Types_go__go_3_6_12 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_6_12
var go__go_3_6_12 gopurs_runtime.Value
_ = go__go_3_6_12
Call_local_Data_List_Lazy_Types_go__go_3_6_12 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_6_12:
for {
if false { continue go__go_3_6_12 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_7 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_5))
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
// TAST (Let): __local_var_7_8 shape=Other bindingType=(TypeVar b$scope212)
__local_var_7_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1
_ = __local_var_7_8
// TAST (Let): __local_var_8_9 shape=App(Other) bindingType=(TypeVar b$scope236)
__local_var_8_9 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), (v_6_7).V0)
_ = __local_var_8_9
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_8_9, __local_var_7_8}))}
}))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
xs_5_loop = (v_6_7).V1
continue go__go_3_6_12
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
go__go_3_6_12 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_6_12(b_4_loop_val, xs_5_loop_val)
})
})
return Call_Data_Tuple_snd(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_3_6_12(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_2_0).V0), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, (v_2_0).V1)))))
})})))}
	})
	return cache_Data_List_Lazy_Types_functorWithIndexList
}

var cache_Data_List_Lazy_Types_functorWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndex sync.Once
func Get_Data_List_Lazy_Types_functorWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndex = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_80275103_2412140840(Rebox_Data_List_Lazy_Types_2412140840_80275103(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_functorWithIndex(gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(Rebox_Data_List_Lazy_Types_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexList()))))})))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorWithIndex(gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(Rebox_Data_List_Lazy_Types_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexList()))))}), "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.Int(int64(0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1)))})))}
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
// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope241)])
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
_ = v1_2_0
// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
__local_var_3_1 := (v1_2_0).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
__local_var_4_2 := (v1_2_0).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope241)])])
__local_var_6_3 := Call_Data_List_Lazy_Types_toList(as_prime__1)
_ = __local_var_6_3
// TAST (Let): __local_var_7_4 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_7_4 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), __local_var_4_2)
_ = __local_var_7_4
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_5 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_9_5 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_4)
_ = __local_var_9_5
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_9_5.Type == 9 && __local_var_9_5.IntVal == 218341868 && __local_var_9_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_6_3))
goto end_branch_6
} else {

}
}
{
if (__local_var_9_5.Type == 9 && __local_var_9_5.IntVal == 218341868 && __local_var_9_5.UnsafePtr != nil) {
__t6 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_5.UnsafePtr).V1, __local_var_6_3)})
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t6)}
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
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope256)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope256)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope256) [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope255)])])])
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_4_2
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_6_3_13 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_6_3_13
var go__go_6_3_13 gopurs_runtime.Value
_ = go__go_6_3_13
Call_local_Data_List_Lazy_Types_go__go_6_3_13 = func(b_7_loop gopurs_runtime.Value, xs_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_6_3_13:
for {
if false { continue go__go_6_3_13 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_4 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_9_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_8))
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
b_7_loop = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, (v_9_4).V0)), b_7)
xs_8_loop = (v_9_4).V1
continue go__go_6_3_13
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
go__go_6_3_13 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_6_3_13(b_7_loop_val, xs_8_loop_val)
})
})
var Call_local_Data_List_Lazy_Types_go__go_7_6_14 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_7_6_14
var go__go_7_6_14 gopurs_runtime.Value
_ = go__go_7_6_14
Call_local_Data_List_Lazy_Types_go__go_7_6_14 = func(b_8_loop gopurs_runtime.Value, xs_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_7_6_14:
for {
if false { continue go__go_7_6_14 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_7 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_10_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_9))
_ = v_10_7
var __t8 gopurs_runtime.Value
{
if (v_10_7 == nil) {
__t8 = b_8
goto end_branch_8
} else {

}
}
{
if (v_10_7 != nil) {
b_8_loop = Call_Data_List_Lazy_Types_cons((v_10_7).V0, b_8)
xs_9_loop = (v_10_7).V1
continue go__go_7_6_14
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
go__go_7_6_14 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_7_6_14(b_8_loop_val, xs_9_loop_val)
})
})
return Call_local_Data_List_Lazy_Types_go__go_6_3_13(__local_var_4_2, Call_local_Data_List_Lazy_Types_go__go_7_6_14(Get_Data_List_Lazy_Types_nil(), xs_5))
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
		cache_Data_List_Lazy_Types_traversableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_4249989635_3043886126(Rebox_Data_List_Lazy_Types_3043886126_4249989635(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))})))))}
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
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope32)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_3))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope27)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
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
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(Rebox_Data_List_Lazy_Types_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope267)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope267)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope267) [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope266)])])])
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_4_2
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_Types_go__go_6_4_15 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_6_4_15
var go__go_6_4_15 gopurs_runtime.Value
_ = go__go_6_4_15
Call_local_Data_List_Lazy_Types_go__go_6_4_15 = func(b_7_loop gopurs_runtime.Value, xs_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_6_4_15:
for {
if false { continue go__go_6_4_15 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_5 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_9_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_8))
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
// TAST (Let): __local_var_11_7 shape=Other bindingType=(TypeVar a$scope199)
__local_var_11_7 := (v_9_5).V0
_ = __local_var_11_7
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_11_7, __local_var_10_6}))}
}))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
xs_8_loop = (v_9_5).V1
continue go__go_6_4_15
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
go__go_6_4_15 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_6_4_15(b_7_loop_val, xs_8_loop_val)
})
})
// TAST (Let): v_6_3 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope211)])])])
v_6_3 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_6_4_15(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, xs_5)))
_ = v_6_3
var Call_local_Data_List_Lazy_Types_go__go_7_9_16 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_7_9_16
var go__go_7_9_16 gopurs_runtime.Value
_ = go__go_7_9_16
Call_local_Data_List_Lazy_Types_go__go_7_9_16 = func(b_8_loop gopurs_runtime.Value, xs_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_7_9_16:
for {
if false { continue go__go_7_9_16 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_10 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_9))
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
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_8.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply2(f_3, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_8.UnsafePtr).V0.IntVal) - (int64(1))), (v_10_10).V0)), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_8.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
xs_9_loop = (v_10_10).V1
continue go__go_7_9_16
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
go__go_7_9_16 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_7_9_16(b_8_loop_val, xs_9_loop_val)
})
})
return Call_Data_Tuple_snd(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Lazy_Types_go__go_7_9_16(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_6_3).V0), __local_var_4_2}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, (v_6_3).V1)))))
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
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_525353375_1812164904(Rebox_Data_List_Lazy_Types_1812164904_525353375(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_traversableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904(Rebox_Data_List_Lazy_Types_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableWithIndexList()))))})))))}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty
}

var cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(Rebox_Data_List_Lazy_Types_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableNonEmptyList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope14)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904(Rebox_Data_List_Lazy_Types_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableWithIndexList()))))}), "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_2, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_4))
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.Int(int64(0))
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Lazy_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_4))
_ = __t_tag_2
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_4.UnsafePtr).V0.IntVal))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
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
var go__go_0_0_17 gopurs_runtime.Value
_ = go__go_0_0_17
var go__go_0_0_17_cell *gopurs_runtime.Value
_ = go__go_0_0_17_cell
// FALLBACK TCO: isLoop=false len=1
go__go_0_0_17 = gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope271), (ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope272)])])
v1_4_1 := Rebox_Data_List_Lazy_Types_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, b_2)))
_ = v1_4_1
var __t7 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v1_4_1).V1
_ = __t_tag_2
if (__t_tag_2 != nil) {
// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
__local_var_5_3 := (v1_4_1).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope271)])])
__local_var_6_4 := gopurs_runtime.Apply2((*go__go_0_0_17_cell), f_1, ((v1_4_1).V1).V0)
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
_ = __t_tag_5
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
return __t7
})))
})
go__go_0_0_17_cell = &go__go_0_0_17
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, go__go_0_0_17}))}
}()
	})
	return cache_Data_List_Lazy_Types_unfoldable1List
}

var cache_Data_List_Lazy_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldableList sync.Once
func Get_Data_List_Lazy_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldableList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldableList = func() gopurs_runtime.Value {
var go__go_0_0_18 gopurs_runtime.Value
_ = go__go_0_0_18
var go__go_0_0_18_cell *gopurs_runtime.Value
_ = go__go_0_0_18_cell
// FALLBACK TCO: isLoop=false len=1
go__go_0_0_18 = gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope276), (TypeVar b$scope277)])])
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
// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope276)])])
__local_var_6_3 := gopurs_runtime.Apply2((*go__go_0_0_18_cell), f_1, ((v1_4_1).V0).V1)
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
return __t4
})))
})
go__go_0_0_18_cell = &go__go_0_0_18
return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldable1List()))}
}), go__go_0_0_18}))}
}()
	})
	return cache_Data_List_Lazy_Types_unfoldableList
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmpty sync.Once
func Get_Data_List_Lazy_Types_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2751875267_2187088110(Rebox_Data_List_Lazy_Types_2187088110_2751875267(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldableList()))})))))}
	})
	return cache_Data_List_Lazy_Types_unfoldable1NonEmpty
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldableList()))}), "unfoldr1"), f_0, b_1)))}
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
return Call_Data_NonEmpty_head(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)))
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
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope297)])])
__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_0)
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
_ = __local_var_4_1
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope298)])])
__local_var_5_2 := gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, f_1)
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_6_3 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0))
_ = __local_var_6_3
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_4 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_8_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3)
_ = __local_var_8_4
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_5_2))
goto end_branch_5
} else {

}
}
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr != nil) {
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_4.UnsafePtr).V1, __local_var_5_2)})
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
}))))
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t6)}
}))
})}))}
	})
	return cache_Data_List_Lazy_Types_bindList
}

var cache_Data_List_Lazy_Types_applyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList sync.Once
func Get_Data_List_Lazy_Types_applyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList.Do(func() {
		cache_Data_List_Lazy_Types_applyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monadList()))}))}
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
// TAST (Let): v2_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope304)])
v2_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_1))
_ = v2_2_0
// TAST (Let): v3_3_1 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (Func [(TypeVar a$scope304)] (TypeVar b$scope305))])
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
// TAST (Let): __local_var_9_6 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope305)])])
__local_var_9_6 := gopurs_runtime.Apply2(Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monadList())), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_4, __local_var_7_5}))}
})), __local_var_5_3)
_ = __local_var_9_6
// TAST (Let): __local_var_10_7 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_10_7 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply2(Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monadList())), __local_var_7_5, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_2, Get_Data_List_Lazy_Types_nil()}))}
}))))
_ = __local_var_10_7
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_8 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_12_8 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_7)
_ = __local_var_12_8
var __t9 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_12_8.Type == 9 && __local_var_12_8.IntVal == 218341868 && __local_var_12_8.UnsafePtr == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_9_6))
goto end_branch_9
} else {

}
}
{
if (__local_var_12_8.Type == 9 && __local_var_12_8.IntVal == 218341868 && __local_var_12_8.UnsafePtr != nil) {
__t9 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_12_8.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_12_8.UnsafePtr).V1, __local_var_9_6)})
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t9)}
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
// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope285)])
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
_ = v1_2_0
// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
__local_var_3_1 := (v1_2_0).V1
_ = __local_var_3_1
// TAST (Let): v2_4_2 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar b$scope286)])
v2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply(f_1, (v1_2_0).V0))))
_ = v2_4_2
// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
__local_var_5_3 := (v2_4_2).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 shape=Other bindingType=Any
__local_var_6_4 := (v2_4_2).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 shape=App(Var) bindingType=(Func [(TypeVar a$scope285)] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope286)])]))
__local_var_8_6 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_toList(), f_1)
_ = __local_var_8_6
// TAST (Let): __local_var_9_7 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope297)])])
__local_var_9_7 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), __local_var_3_1)
_ = __local_var_9_7
// TAST (Let): __local_var_8_5 shape=Let(Let(App(Var))) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope286)])])
__local_var_8_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_8 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_11_8 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_7)
_ = __local_var_11_8
var __t16 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 218341868 && __local_var_11_8.UnsafePtr == nil) {
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_16
} else {

}
}
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 218341868 && __local_var_11_8.UnsafePtr != nil) {
// TAST (Let): __local_var_12_9 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope298)])])
__local_var_12_9 := gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_11_8.UnsafePtr).V1, __local_var_8_6)
_ = __local_var_12_9
// TAST (Let): __local_var_13_10 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_13_10 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply(__local_var_8_6, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_11_8.UnsafePtr).V0))
_ = __local_var_13_10
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_11 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_15_11 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_10)
_ = __local_var_15_11
var __t15 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_15_11.Type == 9 && __local_var_15_11.IntVal == 218341868 && __local_var_15_11.UnsafePtr == nil) {
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_12_9))
goto end_branch_15
} else {

}
}
{
if (__local_var_15_11.Type == 9 && __local_var_15_11.IntVal == 218341868 && __local_var_15_11.UnsafePtr != nil) {
// TAST (Let): __local_var_16_12 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_16_12 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_15_11.UnsafePtr).V1)
_ = __local_var_16_12
__t15 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_15_11.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_13 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_18_13 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_16_12)
_ = __local_var_18_13
var __t14 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_18_13.Type == 9 && __local_var_18_13.IntVal == 218341868 && __local_var_18_13.UnsafePtr == nil) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_12_9))
goto end_branch_14
} else {

}
}
{
if (__local_var_18_13.Type == 9 && __local_var_18_13.IntVal == 218341868 && __local_var_18_13.UnsafePtr != nil) {
__t14 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_18_13.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_18_13.UnsafePtr).V1, __local_var_12_9)})
goto end_branch_14
} else {

}
}
{
__t14 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t14)}
}))})
goto end_branch_15
} else {

}
}
{
__t15 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t15)}
}))))
goto end_branch_16
} else {

}
}
{
__t16 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t16)}
}))
_ = __local_var_8_5
// TAST (Let): __local_var_9_17 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_9_17 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), __local_var_6_4)
_ = __local_var_9_17
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_3, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_18 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_11_18 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_17)
_ = __local_var_11_18
var __t19 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_11_18.Type == 9 && __local_var_11_18.IntVal == 218341868 && __local_var_11_18.UnsafePtr == nil) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_8_5))
goto end_branch_19
} else {

}
}
{
if (__local_var_11_18.Type == 9 && __local_var_11_18.IntVal == 218341868 && __local_var_11_18.UnsafePtr != nil) {
__t19 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_11_18.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_11_18.UnsafePtr).V1, __local_var_8_5)})
goto end_branch_19
} else {

}
}
{
__t19 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t19)}
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
}), Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_semigroupNonEmptyList()))}))}
	})
	return cache_Data_List_Lazy_Types_altNonEmptyList
}

var cache_Data_List_Lazy_Types_altList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altList sync.Once
func Get_Data_List_Lazy_Types_altList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altList.Do(func() {
		cache_Data_List_Lazy_Types_altList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
}), Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_semigroupList()))}))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, Call_Control_Plus_empty(gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_plusList()))})}))}
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

func Call_Data_List_Lazy_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List())), dictEq_0)}))}
}

func Call_Data_List_Lazy_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Lazy_eqLazy(Call_Data_NonEmpty_eqNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))}, dictEq_0))))}
}

func Call_Data_List_Lazy_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqList1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope57)])])])
eqList1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_eqList(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
_ = eqList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_0)}
}), gopurs_runtime.Apply(Call_Data_Ord_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List())), dictOrd_0)}))}
}

func Call_Data_List_Lazy_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Lazy_ordLazy(gopurs_runtime.Apply(Call_Data_NonEmpty_ordNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()))}), dictOrd_0))))}
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
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope47)])
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_1))
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
var Call_local_Data_List_Lazy_Types_go__go_3_1_7 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_Types_go__go_3_1_7
var go__go_3_1_7 gopurs_runtime.Value
_ = go__go_3_1_7
Call_local_Data_List_Lazy_Types_go__go_3_1_7 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_1_7:
for {
if false { continue go__go_3_1_7 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_5))
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
continue go__go_3_1_7
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
go__go_3_1_7 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_Types_go__go_3_1_7(b_4_loop_val, xs_5_loop_val)
})
})
__t4 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_0).V0).StrVal())) + (Call_local_Data_List_Lazy_Types_go__go_3_1_7(gopurs_runtime.Str(""), (v_2_0).V1).StrVal())) + ("])")
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
// TAST (Let): showLazy_1_0 shape=App(Var) bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope44)])])])
showLazy_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Lazy_showLazy(Call_Data_NonEmpty_showNonEmpty(dictShow_0, Call_Data_List_Lazy_Types_showList(dictShow_0))))
_ = showLazy_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyList ") + (gopurs_runtime.Apply(showLazy_1_0.V0, v_2).StrVal())) + (")"))
})}))}
}

func Call_Data_List_Lazy_Types_showStep(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList1_1_0 shape=App(Var) bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope41)])])])
showList1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_showList(dictShow_0))
_ = showList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_928333203_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 string
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t3 = "Nil"
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2)
_ = __t_tag_2
if (__t_tag_2 != nil) {
__t3 = (((("(") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" : ")) + (gopurs_runtime.Apply(showList1_1_0.V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_3
} else {

}
}
{
__t3 = func() string { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Str(__t3)
})})))}
}

func Call_Data_List_Lazy_Types_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope238)])
v2_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
_ = v2_2_0
// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
__local_var_3_1 := (v2_2_0).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
__local_var_4_2 := (v2_2_0).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_4_2}))}
}))
})))
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

func Rebox_Data_List_Lazy_Types_167870147_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_1680800814_3071895939(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_List_Lazy_Types_1728839155_138441832(in *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_1812164904_2955889203(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
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

func Rebox_Data_List_Lazy_Types_2187088110_2751875267(in *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Lazy_Types_2412140840_2773701683(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_2412140840_80275103(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
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

func Rebox_Data_List_Lazy_Types_2812149806_1468200963(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
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

func Rebox_Data_List_Lazy_Types_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
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

func Rebox_Data_List_Lazy_Types_3725484264_2491554675(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Lazy_Types_3790796878_2176830691(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
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


