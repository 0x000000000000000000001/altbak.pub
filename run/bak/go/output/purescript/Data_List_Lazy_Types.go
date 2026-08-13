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
		cache_Data_List_Lazy_Types_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_identity(x_0_box)
})
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
		cache_Data_List_Lazy_Types_Nil = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
	})
	return cache_Data_List_Lazy_Types_Nil
}

var cache_Data_List_Lazy_Types_Cons gopurs_runtime.Value
var once_Data_List_Lazy_Types_Cons sync.Once
func Get_Data_List_Lazy_Types_Cons() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_Cons.Do(func() {
		cache_Data_List_Lazy_Types_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil
}

var cache_Data_List_Lazy_Types_newtypeNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_newtypeNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_newtypeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_newtypeNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_newtypeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_List_Lazy_Types_newtypeNonEmptyList
}

var cache_Data_List_Lazy_Types_newtypeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_newtypeList sync.Once
func Get_Data_List_Lazy_Types_newtypeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_newtypeList.Do(func() {
		cache_Data_List_Lazy_Types_newtypeList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
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
		cache_Data_List_Lazy_Types_semigroupList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1, ys_1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_semigroupList
}

var cache_Data_List_Lazy_Types_monoidList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monoidList sync.Once
func Get_Data_List_Lazy_Types_monoidList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monoidList.Do(func() {
		cache_Data_List_Lazy_Types_monoidList = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()))}
}), Get_Data_List_Lazy_Types_nil()})}
	})
	return cache_Data_List_Lazy_Types_monoidList
}

var cache_Data_List_Lazy_Types_lazyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_lazyList sync.Once
func Get_Data_List_Lazy_Types_lazyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_lazyList.Do(func() {
		cache_Data_List_Lazy_Types_lazyList = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(&Constructor_Control_Lazy_Lazy{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, x_1))
}))
})})}
	})
	return cache_Data_List_Lazy_Types_lazyList
}

var cache_Data_List_Lazy_Types_functorList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList sync.Once
func Get_Data_List_Lazy_Types_functorList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList.Do(func() {
		cache_Data_List_Lazy_Types_functorList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorList
}

var cache_Data_List_Lazy_Types_functorNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmpty sync.Once
func Get_Data_List_Lazy_Types_functorNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
_ = __local_var_4_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))})}
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorNonEmpty
}

var cache_Data_List_Lazy_Types_functorNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_functorNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmpty()).V0), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorNonEmptyList
}

var cache_Data_List_Lazy_Types_eq1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1List sync.Once
func Get_Data_List_Lazy_Types_eq1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1List.Do(func() {
		cache_Data_List_Lazy_Types_eq1List = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_0 gopurs_runtime.Value
go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
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
continue go__go_3_0_0
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})
})
})})}
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
return gopurs_runtime.Bool(Call_Data_List_Lazy_Types_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), xs_1_box, ys_2_box))
})
	})
	return cache_Data_List_Lazy_Types_eq1
}

var cache_Data_List_Lazy_Types_eq1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_eq1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_eq1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqNonEmpty1_1_0 -> *Constructor_Data_Eq_Eq
eqNonEmpty1_1_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_4 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_1_3 gopurs_runtime.Value
go__go_3_1_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_1_3:
for {
if false { continue go__go_3_1_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 bool
{
if (v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr == nil) {
var __t2 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)))}
continue go__go_3_1_3
__t3 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
}
}()
})
})
__t_and_4 = (gopurs_runtime.Apply2(go__go_3_1_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_4)
})
})}
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqNonEmpty1_1_0.V0), gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_3)).IntVal) != (0))
})
})
})})}
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
		cache_Data_List_Lazy_Types_ord1List = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_List_Lazy_Types_eq1List()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_6 gopurs_runtime.Value
go__go_3_0_6 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_6:
for {
if false { continue go__go_3_0_6 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t4 uint32
{
if (v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr == nil) {
var __t1 uint32
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
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
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) {
// TAST (Let): v2_6_2 -> gopurs_runtime.Value
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0)
_ = v2_6_2
var __t3 uint32
{
if (uint32(v2_6_2.IntVal) == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)))}
continue go__go_3_0_6
__t3 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(v2_6_2.IntVal)
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_3_0_6, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal)), UnsafePtr: nil}
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_ord1List
}

var cache_Data_List_Lazy_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordNonEmpty sync.Once
func Get_Data_List_Lazy_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_ordNonEmpty = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_ordNonEmpty(dictOrd_0_box)
})
	})
	return cache_Data_List_Lazy_Types_ordNonEmpty
}

var cache_Data_List_Lazy_Types_compare1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compare1 sync.Once
func Get_Data_List_Lazy_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compare1.Do(func() {
		cache_Data_List_Lazy_Types_compare1 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_List_Lazy_Types_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), xs_1_box, ys_2_box)), UnsafePtr: nil}
})
	})
	return cache_Data_List_Lazy_Types_compare1
}

var cache_Data_List_Lazy_Types_ord1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_ord1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_ord1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ord1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_ord1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_List_Lazy_Types_eq1NonEmptyList()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqNonEmpty2_1_1 -> *Constructor_Data_Eq_Eq
eqNonEmpty2_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_6 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0).IntVal) != (0) {

var go__go_4_3_10 gopurs_runtime.Value
go__go_4_3_10 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_3_10:
for {
if false { continue go__go_4_3_10 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 bool
{
if (v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr == nil) {
var __t4 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr == nil) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr != nil)) && (((v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V0).IntVal) != (0))) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V1)))}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V1)))}
continue go__go_4_3_10
__t5 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
}
}()
})
})
__t_and_6 = (gopurs_runtime.Apply2(go__go_4_3_10, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_6)
})
})))
_ = eqNonEmpty2_1_1
// TAST (Let): ordNonEmpty1_1_0 -> *Constructor_Data_Ord_Ord
ordNonEmpty1_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty2_1_1)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_7 -> gopurs_runtime.Value
v_4_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0)
_ = v_4_7
var __t13 uint32
{
if (uint32(v_4_7.IntVal) == 1527465420) {
__t13 = 1527465420
goto end_branch_13
} else {

}
}
{
if (uint32(v_4_7.IntVal) == 380165415) {
__t13 = 380165415
goto end_branch_13
} else {

}
}
{
var go__go_5_8_11 gopurs_runtime.Value
go__go_5_8_11 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_8_11:
for {
if false { continue go__go_5_8_11 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t12 uint32
{
if (v_6.Type == 9 && v_6.IntVal == 218341868 && v_6.UnsafePtr == nil) {
var __t9 uint32
{
if (v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr == nil) {
__t9 = 902936544
goto end_branch_9
} else {

}
}
{
__t9 = 1527465420
}
end_branch_9:
__t12 = __t9
goto end_branch_12
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr == nil) {
__t12 = 380165415
goto end_branch_12
} else {

}
}
{
if ((v_6.Type == 9 && v_6.IntVal == 218341868 && v_6.UnsafePtr != nil)) && ((v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr != nil)) {
// TAST (Let): v2_8_10 -> gopurs_runtime.Value
v2_8_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Lazy_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_7.UnsafePtr).V0)
_ = v2_8_10
var __t11 uint32
{
if (uint32(v2_8_10.IntVal) == 902936544) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_6.UnsafePtr).V1)))}
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_7.UnsafePtr).V1)))}
continue go__go_5_8_11
__t11 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_11
} else {

}
}
{
__t11 = uint32(v2_8_10.IntVal)
}
end_branch_11:
__t12 = __t11
goto end_branch_12
} else {

}
}
{
__t12 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t12), UnsafePtr: nil}
}
}()
})
})
__t13 = uint32(gopurs_runtime.Apply2(go__go_5_8_11, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1)))}).IntVal)
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t13), UnsafePtr: nil}
})
})}
_ = ordNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordNonEmpty1_1_0.V1), gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_3)).IntVal)), UnsafePtr: nil}
})
})
})})}
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
		cache_Data_List_Lazy_Types_foldableList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()).V1), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_16 gopurs_runtime.Value
go__go_1_2_16 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_16:
for {
if false { continue go__go_1_2_16 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
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
continue go__go_1_2_16
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
return go__go_1_2_16
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()).V1), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()).V1), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, a_4, b_3})}
}))
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableList
}

var cache_Data_List_Lazy_Types_foldableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmpty sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_1
var go__go_5_2_17 gopurs_runtime.Value
go__go_5_2_17 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_2_17:
for {
if false { continue go__go_5_2_17 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
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
continue go__go_5_2_17
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_5_2_17, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_18 gopurs_runtime.Value
go__go_3_5_18 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_5_18:
for {
if false { continue go__go_3_5_18 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
continue go__go_3_5_18
__t7 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_5_18, gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_19 gopurs_runtime.Value
go__go_3_8_19 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_8_19:
for {
if false { continue go__go_3_8_19 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_9 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
continue go__go_3_8_19
__t10 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_4_11_20 gopurs_runtime.Value
go__go_4_11_20 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_11_20:
for {
if false { continue go__go_4_11_20 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
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
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := (v_7_12).V0
_ = __local_var_8_13
b_5_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_13, b_5})}
}))
xs_6_loop = (v_7_12).V1
continue go__go_4_11_20
__t14 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_3_8_19, b_1, gopurs_runtime.Apply2(go__go_4_11_20, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableNonEmpty
}

var cache_Data_List_Lazy_Types_extendList gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendList sync.Once
func Get_Data_List_Lazy_Types_extendList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendList.Do(func() {
		cache_Data_List_Lazy_Types_extendList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
_ = v_2_0
var __t12 gopurs_runtime.Value
{
if (v_2_0 == nil) {
__t12 = Get_Data_List_Lazy_Types_nil()
goto end_branch_12
} else {

}
}
{
if (v_2_0 != nil) {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(f_0, l_1)
_ = __local_var_3_1
var go__go_4_3_21 gopurs_runtime.Value
go__go_4_3_21 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_3_21:
for {
if false { continue go__go_4_3_21 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_4 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_4
var __t7 gopurs_runtime.Value
{
if (v_7_4 == nil) {
__t7 = b_5
goto end_branch_7
} else {

}
}
{
if (v_7_4 != nil) {
// TAST (Let): acc_prime_8_5 -> gopurs_runtime.Value
acc_prime_8_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v_7_4).V0, gopurs_runtime.RecordGet(b_5, "acc")})}
}))
_ = acc_prime_8_5
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply(f_0, acc_prime_8_5)
_ = __local_var_9_6
b_5_loop = gopurs_runtime.RecordDict2("acc", "val", acc_prime_8_5, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_6, gopurs_runtime.RecordGet(b_5, "val")})}
})))
xs_6_loop = (v_7_4).V1
continue go__go_4_3_21
__t7 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_5_8_22 gopurs_runtime.Value
go__go_5_8_22 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_8_22:
for {
if false { continue go__go_5_8_22 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_9 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_9
var __t11 gopurs_runtime.Value
{
if (v_8_9 == nil) {
__t11 = b_6
goto end_branch_11
} else {

}
}
{
if (v_8_9 != nil) {
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := (v_8_9).V0
_ = __local_var_9_10
b_6_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_10, b_6})}
}))
xs_7_loop = (v_8_9).V1
continue go__go_5_8_22
__t11 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply2(go__go_4_3_21, gopurs_runtime.RecordDict2("acc", "val", Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_5_8_22, Get_Data_List_Lazy_Types_nil(), (v_2_0).V1)), "val")
_ = __local_var_4_2
__t12 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_3_1, __local_var_4_2})}
}))
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
})})}
	})
	return cache_Data_List_Lazy_Types_extendList
}

var cache_Data_List_Lazy_Types_extendNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_extendNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_extendNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1).UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_23 gopurs_runtime.Value
go__go_4_1_23 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_1_23:
for {
if false { continue go__go_4_1_23 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_2 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_2
var __t5 gopurs_runtime.Value
{
if (v_7_2 == nil) {
__t5 = b_5
goto end_branch_5
} else {

}
}
{
if (v_7_2 != nil) {
// TAST (Let): __local_var_8_3 -> gopurs_runtime.Value
__local_var_8_3 := gopurs_runtime.RecordGet(b_5, "acc")
_ = __local_var_8_3
// TAST (Let): __local_var_9_4 -> gopurs_runtime.Value
__local_var_9_4 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v_7_2).V0, __local_var_8_3})}
})))
_ = __local_var_9_4
b_5_loop = gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v_7_2).V0, __local_var_8_3})}
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_4, gopurs_runtime.RecordGet(b_5, "val")})}
})))
xs_6_loop = (v_7_2).V1
continue go__go_4_1_23
__t5 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_5_6_24 gopurs_runtime.Value
go__go_5_6_24 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_6_24:
for {
if false { continue go__go_5_6_24 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_7
var __t9 gopurs_runtime.Value
{
if (v_8_7 == nil) {
__t9 = b_6
goto end_branch_9
} else {

}
}
{
if (v_8_7 != nil) {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := (v_8_7).V0
_ = __local_var_9_8
b_6_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_8, b_6})}
}))
xs_7_loop = (v_8_7).V1
continue go__go_5_6_24
__t9 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(go__go_4_1_23, gopurs_runtime.RecordDict2("acc", "val", Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_5_6_24, Get_Data_List_Lazy_Types_nil(), __local_var_2_0)), "val")})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_extendNonEmptyList
}

var cache_Data_List_Lazy_Types_foldableNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V1), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V2), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
})})}
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
		cache_Data_List_Lazy_Types_foldableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](Get_Data_List_Lazy_Types_foldableWithIndexList()).V2), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_5)
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(f_3, gopurs_runtime.Int(i_4.IntVal))
_ = __local_var_7_3
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_7_3, x_8))
})
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_28 gopurs_runtime.Value
go__go_2_5_28 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_5_28:
for {
if false { continue go__go_2_5_28 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V1, (v_5_6).V0)})}
xs_4_loop = (v_5_6).V1
continue go__go_2_5_28
__t7 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply(go__go_2_5_28, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_9_29 gopurs_runtime.Value
go__go_3_9_29 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_9_29:
for {
if false { continue go__go_3_9_29 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_10 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
// TAST (Let): __local_var_7_11 -> gopurs_runtime.Value
__local_var_7_11 := (*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1
_ = __local_var_7_11
// TAST (Let): __local_var_8_12 -> gopurs_runtime.Value
__local_var_8_12 := (v_6_10).V0
_ = __local_var_8_12
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_12, __local_var_7_11})}
}))})}
xs_5_loop = (v_6_10).V1
continue go__go_3_9_29
__t13 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_3_8 -> *Constructor_Data_Tuple_Tuple
v_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_9_29, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_2))
_ = v_3_8
var go__go_4_14_30 gopurs_runtime.Value
go__go_4_14_30 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_14_30:
for {
if false { continue go__go_4_14_30 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_15 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
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
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), (v_7_15).V0, (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1)})}
xs_6_loop = (v_7_15).V1
continue go__go_4_14_30
__t16 = gopurs_runtime.Value{}
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
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_14_30, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_8).V0.IntVal), b_1})}, (v_3_8).V1).UnsafePtr).V1
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexList
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmpty sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty = func() gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_0_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_0_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
var go__go_5_3_31 gopurs_runtime.Value
go__go_5_3_31 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_3_31:
for {
if false { continue go__go_5_3_31 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_4 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_4
var __t5 gopurs_runtime.Value
{
if (v_8_4 == nil) {
__t5 = b_6
goto end_branch_5
} else {

}
}
{
if (v_8_4 != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_2.V0), b_6, gopurs_runtime.Apply(f_2, (v_8_4).V0))
xs_7_loop = (v_8_4).V1
continue go__go_5_3_31
__t5 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_1.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_5_3_31, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_6_32 gopurs_runtime.Value
go__go_3_6_32 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_6_32:
for {
if false { continue go__go_3_6_32 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_7
var __t8 gopurs_runtime.Value
{
if (v_6_7 == nil) {
__t8 = b_4
goto end_branch_8
} else {

}
}
{
if (v_6_7 != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (v_6_7).V0)
xs_5_loop = (v_6_7).V1
continue go__go_3_6_32
__t8 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_6_32, gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_9_33 gopurs_runtime.Value
go__go_3_9_33 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_9_33:
for {
if false { continue go__go_3_9_33 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_10 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_10
var __t11 gopurs_runtime.Value
{
if (v_6_10 == nil) {
__t11 = b_4
goto end_branch_11
} else {

}
}
{
if (v_6_10 != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (v_6_10).V0, b_4)
xs_5_loop = (v_6_10).V1
continue go__go_3_9_33
__t11 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_4_12_34 gopurs_runtime.Value
go__go_4_12_34 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_12_34:
for {
if false { continue go__go_4_12_34 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_13 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_13
var __t15 gopurs_runtime.Value
{
if (v_7_13 == nil) {
__t15 = b_5
goto end_branch_15
} else {

}
}
{
if (v_7_13 != nil) {
// TAST (Let): __local_var_8_14 -> gopurs_runtime.Value
__local_var_8_14 := (v_7_13).V0
_ = __local_var_8_14
b_5_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_14, b_5})}
}))
xs_6_loop = (v_7_13).V1
continue go__go_4_12_34
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_3_9_33, b_1, gopurs_runtime.Apply2(go__go_4_12_34, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_16 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_16
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_17 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_17
var go__go_6_18_35 gopurs_runtime.Value
go__go_6_18_35 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_18_35:
for {
if false { continue go__go_6_18_35 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_19 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_19
var __t20 gopurs_runtime.Value
{
if (v_9_19 == nil) {
__t20 = b_7
goto end_branch_20
} else {

}
}
{
if (v_9_19 != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_17.V0), (*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V1, gopurs_runtime.Apply2(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal)})}, (v_9_19).V0))})}
xs_8_loop = (v_9_19).V1
continue go__go_6_18_35
__t20 = gopurs_runtime.Value{}
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_16.V0), gopurs_runtime.Apply2(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_6_18_35, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.RecordGet(dictMonoid_1, "mempty")})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1).UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_21_36 gopurs_runtime.Value
go__go_4_21_36 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_21_36:
for {
if false { continue go__go_4_21_36 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_22 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_22
var __t23 gopurs_runtime.Value
{
if (v_7_22 == nil) {
__t23 = b_5
goto end_branch_23
} else {

}
}
{
if (v_7_22 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1, (v_7_22).V0)})}
xs_6_loop = (v_7_22).V1
continue go__go_4_21_36
__t23 = gopurs_runtime.Value{}
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
}
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_21_36, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply3(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_25_37 gopurs_runtime.Value
go__go_4_25_37 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_25_37:
for {
if false { continue go__go_4_25_37 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_26 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_26
var __t29 gopurs_runtime.Value
{
if (v_7_26 == nil) {
__t29 = b_5
goto end_branch_29
} else {

}
}
{
if (v_7_26 != nil) {
// TAST (Let): __local_var_8_27 -> gopurs_runtime.Value
__local_var_8_27 := (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1
_ = __local_var_8_27
// TAST (Let): __local_var_9_28 -> gopurs_runtime.Value
__local_var_9_28 := (v_7_26).V0
_ = __local_var_9_28
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_28, __local_var_8_27})}
}))})}
xs_6_loop = (v_7_26).V1
continue go__go_4_25_37
__t29 = gopurs_runtime.Value{}
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return __t29
}
}()
})
})
// TAST (Let): v_4_24 -> *Constructor_Data_Tuple_Tuple
v_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_4_25_37, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
_ = v_4_24
var go__go_5_30_38 gopurs_runtime.Value
go__go_5_30_38 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_30_38:
for {
if false { continue go__go_5_30_38 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_31 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_31
var __t32 gopurs_runtime.Value
{
if (v_8_31 == nil) {
__t32 = b_6
goto end_branch_32
} else {

}
}
{
if (v_8_31 != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1))})}, (v_8_31).V0, (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1)})}
xs_7_loop = (v_8_31).V1
continue go__go_5_30_38
__t32 = gopurs_runtime.Value{}
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
return __t32
}
}()
})
})
return gopurs_runtime.Apply3(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_30_38, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_4_24).V0.IntVal), b_2})}, (v_4_24).V1).UnsafePtr).V1)
})
})
})})}
}()
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmptyList()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2))
_ = __local_var_3_0
// TAST (Let): Semigroup0_4_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_1
var go__go_5_2_39 gopurs_runtime.Value
go__go_5_2_39 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_2_39:
for {
if false { continue go__go_5_2_39 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
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
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_1.V0), (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal)), (v_8_3).V0))})}
xs_7_loop = (v_8_3).V1
continue go__go_5_2_39
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_1, gopurs_runtime.Int(0), (__local_var_3_0).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_2_39, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})}, (__local_var_3_0).V1).UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2))
_ = __local_var_3_5
var go__go_4_6_40 gopurs_runtime.Value
go__go_4_6_40 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_6_40:
for {
if false { continue go__go_4_6_40 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_7
var __t8 gopurs_runtime.Value
{
if (v_7_7 == nil) {
__t8 = b_5
goto end_branch_8
} else {

}
}
{
if (v_7_7 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal)), (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1, (v_7_7).V0)})}
xs_6_loop = (v_7_7).V1
continue go__go_4_6_40
__t8 = gopurs_runtime.Value{}
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
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_6_40, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(0), b_1, (__local_var_3_5).V0)})}, (__local_var_3_5).V1).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_9 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2))
_ = __local_var_3_9
var go__go_4_11_41 gopurs_runtime.Value
go__go_4_11_41 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_11_41:
for {
if false { continue go__go_4_11_41 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_12
var __t15 gopurs_runtime.Value
{
if (v_7_12 == nil) {
__t15 = b_5
goto end_branch_15
} else {

}
}
{
if (v_7_12 != nil) {
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1
_ = __local_var_8_13
// TAST (Let): __local_var_9_14 -> gopurs_runtime.Value
__local_var_9_14 := (v_7_12).V0
_ = __local_var_9_14
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_14, __local_var_8_13})}
}))})}
xs_6_loop = (v_7_12).V1
continue go__go_4_11_41
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
// TAST (Let): v_4_10 -> *Constructor_Data_Tuple_Tuple
v_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_4_11_41, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (__local_var_3_9).V1))
_ = v_4_10
var go__go_5_16_42 gopurs_runtime.Value
go__go_5_16_42 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_16_42:
for {
if false { continue go__go_5_16_42 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_17 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_17
var __t18 gopurs_runtime.Value
{
if (v_8_17 == nil) {
__t18 = b_6
goto end_branch_18
} else {

}
}
{
if (v_8_17 != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1))), (v_8_17).V0, (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1)})}
xs_7_loop = (v_8_17).V1
continue go__go_5_16_42
__t18 = gopurs_runtime.Value{}
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}
}()
})
})
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(0), (__local_var_3_9).V0, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_16_42, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_4_10).V0.IntVal), b_1})}, (v_4_10).V1).UnsafePtr).V1)
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_functorWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexList sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_43 gopurs_runtime.Value
go__go_2_1_43 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_1_43:
for {
if false { continue go__go_2_1_43 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_2 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
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
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := (*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V1
_ = __local_var_6_3
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (v_5_2).V0
_ = __local_var_7_4
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_7_4, __local_var_6_3})}
}))})}
xs_4_loop = (v_5_2).V1
continue go__go_2_1_43
__t5 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_2_0 -> *Constructor_Data_Tuple_Tuple
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_2_1_43, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_1))
_ = v_2_0
var go__go_3_6_44 gopurs_runtime.Value
go__go_3_6_44 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_6_44:
for {
if false { continue go__go_3_6_44 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := (*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1
_ = __local_var_7_8
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1)), (v_6_7).V0)
_ = __local_var_8_9
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_9, __local_var_7_8})}
}))})}
xs_5_loop = (v_6_7).V1
continue go__go_3_6_44
__t10 = gopurs_runtime.Value{}
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
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_3_6_44, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_2_0).V0.IntVal), Get_Data_List_Lazy_Types_nil()})}, (v_2_0).V1).UnsafePtr).V1
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorWithIndexList
}

var cache_Data_List_Lazy_Types_functorWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndex sync.Once
func Get_Data_List_Lazy_Types_functorWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndex = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_0 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_0_0 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_1)
_ = __local_var_4_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 218341868 && __local_var_4_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 218341868 && __local_var_4_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_2.UnsafePtr).V1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))})}
})
})}
_ = functorNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_45 gopurs_runtime.Value
go__go_3_5_45 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_5_45:
for {
if false { continue go__go_3_5_45 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := (*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1
_ = __local_var_7_7
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := (v_6_6).V0
_ = __local_var_8_8
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_8, __local_var_7_7})}
}))})}
xs_5_loop = (v_6_6).V1
continue go__go_3_5_45
__t9 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_3_4 -> *Constructor_Data_Tuple_Tuple
v_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_5_45, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
_ = v_3_4
var go__go_4_10_46 gopurs_runtime.Value
go__go_4_10_46 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_10_46:
for {
if false { continue go__go_4_10_46 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_11 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
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
// TAST (Let): __local_var_8_12 -> gopurs_runtime.Value
__local_var_8_12 := (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1
_ = __local_var_8_12
// TAST (Let): __local_var_9_13 -> gopurs_runtime.Value
__local_var_9_13 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1))})}, (v_7_11).V0)
_ = __local_var_9_13
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_13, __local_var_8_12})}
}))})}
xs_6_loop = (v_7_11).V1
continue go__go_4_10_46
__t14 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_10_46, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_4).V0.IntVal), Get_Data_List_Lazy_Types_nil()})}, (v_3_4).V1).UnsafePtr).V1})}
})
})})}
}()
	})
	return cache_Data_List_Lazy_Types_functorWithIndex
}

var cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
_ = __local_var_3_0
var go__go_4_2_47 gopurs_runtime.Value
go__go_4_2_47 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_2_47:
for {
if false { continue go__go_4_2_47 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_3
var __t6 gopurs_runtime.Value
{
if (v_7_3 == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_7_3 != nil) {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1
_ = __local_var_8_4
// TAST (Let): __local_var_9_5 -> gopurs_runtime.Value
__local_var_9_5 := (v_7_3).V0
_ = __local_var_9_5
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_5, __local_var_8_4})}
}))})}
xs_6_loop = (v_7_3).V1
continue go__go_4_2_47
__t6 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_4_1 -> *Constructor_Data_Tuple_Tuple
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_4_2_47, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (__local_var_3_0).V1))
_ = v_4_1
var go__go_5_7_48 gopurs_runtime.Value
go__go_5_7_48 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_7_48:
for {
if false { continue go__go_5_7_48 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_8 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_8
var __t11 gopurs_runtime.Value
{
if (v_8_8 == nil) {
__t11 = b_6
goto end_branch_11
} else {

}
}
{
if (v_8_8 != nil) {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1
_ = __local_var_9_9
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1))), (v_8_8).V0)
_ = __local_var_10_10
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_10_10, __local_var_9_9})}
}))})}
xs_7_loop = (v_8_8).V1
continue go__go_5_7_48
__t11 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(0), (__local_var_3_0).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_7_48, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_4_1).V0.IntVal), Get_Data_List_Lazy_Types_nil()})}, (v_4_1).V1).UnsafePtr).V1})}
}))
})
})})}
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
		cache_Data_List_Lazy_Types_semigroupNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := Call_Data_List_Lazy_Types_toList(as_prime_1)
_ = __local_var_6_3
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_3_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2)
_ = __local_var_8_4
var __t5 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
goto end_branch_5
} else {

}
}
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr != nil) {
__t5 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_4.UnsafePtr).V1, __local_var_6_3)}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
}))})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_semigroupNonEmptyList
}

var cache_Data_List_Lazy_Types_traversableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableList sync.Once
func Get_Data_List_Lazy_Types_traversableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableList.Do(func() {
		cache_Data_List_Lazy_Types_traversableList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Lazy_Types_traversableList()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_4_2
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_3_49 gopurs_runtime.Value
go__go_6_3_49 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_3_49:
for {
if false { continue go__go_6_3_49 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_4 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
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
continue go__go_6_3_49
__t5 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_7_6_50 gopurs_runtime.Value
go__go_7_6_50 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_6_50:
for {
if false { continue go__go_7_6_50 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
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
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := (v_10_7).V0
_ = __local_var_11_8
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_11_8, b_8})}
}))
xs_9_loop = (v_10_7).V1
continue go__go_7_6_50
__t9 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_6_3_49, __local_var_4_2, gopurs_runtime.Apply2(go__go_7_6_50, Get_Data_List_Lazy_Types_nil(), xs_5))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_traversableList
}

var cache_Data_List_Lazy_Types_traversableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmpty sync.Once
func Get_Data_List_Lazy_Types_traversableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmpty = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_0 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_0_0 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_1)
_ = __local_var_4_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 218341868 && __local_var_4_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 218341868 && __local_var_4_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_2.UnsafePtr).V1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))})}
})
})}
_ = functorNonEmpty1_0_0
// TAST (Let): foldableNonEmpty1_1_4 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_4 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_5 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_5
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_6 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_6
var go__go_6_7_51 gopurs_runtime.Value
go__go_6_7_51 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_7_51:
for {
if false { continue go__go_6_7_51 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_8 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_8
var __t9 gopurs_runtime.Value
{
if (v_9_8 == nil) {
__t9 = b_7
goto end_branch_9
} else {

}
}
{
if (v_9_8 != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), b_7, gopurs_runtime.Apply(f_3, (v_9_8).V0))
xs_8_loop = (v_9_8).V1
continue go__go_6_7_51
__t9 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_5.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_6_7_51, gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_10_52 gopurs_runtime.Value
go__go_4_10_52 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_10_52:
for {
if false { continue go__go_4_10_52 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_11 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_11
var __t12 gopurs_runtime.Value
{
if (v_7_11 == nil) {
__t12 = b_5
goto end_branch_12
} else {

}
}
{
if (v_7_11 != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, b_5, (v_7_11).V0)
xs_6_loop = (v_7_11).V1
continue go__go_4_10_52
__t12 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_4_10_52, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_13_53 gopurs_runtime.Value
go__go_4_13_53 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_13_53:
for {
if false { continue go__go_4_13_53 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_14 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_14
var __t15 gopurs_runtime.Value
{
if (v_7_14 == nil) {
__t15 = b_5
goto end_branch_15
} else {

}
}
{
if (v_7_14 != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, (v_7_14).V0, b_5)
xs_6_loop = (v_7_14).V1
continue go__go_4_13_53
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
var go__go_5_16_54 gopurs_runtime.Value
go__go_5_16_54 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_16_54:
for {
if false { continue go__go_5_16_54 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_17 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_17
var __t19 gopurs_runtime.Value
{
if (v_8_17 == nil) {
__t19 = b_6
goto end_branch_19
} else {

}
}
{
if (v_8_17 != nil) {
// TAST (Let): __local_var_9_18 -> gopurs_runtime.Value
__local_var_9_18 := (v_8_17).V0
_ = __local_var_9_18
b_6_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_18, b_6})}
}))
xs_7_loop = (v_8_17).V1
continue go__go_5_16_54
__t19 = gopurs_runtime.Value{}
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}
}()
})
})
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_4_13_53, b_2, gopurs_runtime.Apply2(go__go_5_16_54, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_1_4
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_4)}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_3_20 -> *Constructor_Control_Apply_Apply
Apply0_3_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_20
// TAST (Let): Functor0_4_21 -> *Constructor_Data_Functor_Functor
Functor0_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_21
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_6_22 -> *Constructor_Control_Apply_Apply
Apply0_6_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_22
// TAST (Let): Functor0_7_23 -> *Constructor_Data_Functor_Functor
Functor0_7_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_23
var go__go_8_24_55 gopurs_runtime.Value
go__go_8_24_55 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var xs_10_loop gopurs_runtime.Value = xs_10_loop_val
go__go_8_24_55:
for {
if false { continue go__go_8_24_55 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_25 -> *Constructor_Data_List_Lazy_Types_Cons
v_11_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_10))
_ = v_11_25
var __t26 gopurs_runtime.Value
{
if (v_11_25 == nil) {
__t26 = b_9
goto end_branch_26
} else {

}
}
{
if (v_11_25 != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_22.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_23.V0), Get_Data_List_Lazy_Types_cons(), (v_11_25).V0), b_9)
xs_10_loop = (v_11_25).V1
continue go__go_8_24_55
__t26 = gopurs_runtime.Value{}
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
return __t26
}
}()
})
})
var go__go_9_27_56 gopurs_runtime.Value
go__go_9_27_56 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var xs_11_loop gopurs_runtime.Value = xs_11_loop_val
go__go_9_27_56:
for {
if false { continue go__go_9_27_56 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var xs_11 gopurs_runtime.Value = xs_11_loop
_ = xs_11
// TAST (Let): v_12_28 -> *Constructor_Data_List_Lazy_Types_Cons
v_12_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_11))
_ = v_12_28
var __t30 gopurs_runtime.Value
{
if (v_12_28 == nil) {
__t30 = b_10
goto end_branch_30
} else {

}
}
{
if (v_12_28 != nil) {
// TAST (Let): __local_var_13_29 -> gopurs_runtime.Value
__local_var_13_29 := (v_12_28).V0
_ = __local_var_13_29
b_10_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_13_29, b_10})}
}))
xs_11_loop = (v_12_28).V1
continue go__go_9_27_56
__t30 = gopurs_runtime.Value{}
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_20.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_21.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_8_24_55, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_9_27_56, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1)))
})
}), gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_3_31 -> *Constructor_Control_Apply_Apply
Apply0_3_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_31
// TAST (Let): Functor0_4_32 -> *Constructor_Data_Functor_Functor
Functor0_4_32 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_32
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_33 -> *Constructor_Control_Apply_Apply
Apply0_7_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_33
// TAST (Let): Functor0_8_34 -> *Constructor_Data_Functor_Functor
Functor0_8_34 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_34
var go__go_9_35_57 gopurs_runtime.Value
go__go_9_35_57 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var xs_11_loop gopurs_runtime.Value = xs_11_loop_val
go__go_9_35_57:
for {
if false { continue go__go_9_35_57 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var xs_11 gopurs_runtime.Value = xs_11_loop
_ = xs_11
// TAST (Let): v_12_36 -> *Constructor_Data_List_Lazy_Types_Cons
v_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_11))
_ = v_12_36
var __t37 gopurs_runtime.Value
{
if (v_12_36 == nil) {
__t37 = b_10
goto end_branch_37
} else {

}
}
{
if (v_12_36 != nil) {
b_10_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_33.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_34.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_5, (v_12_36).V0)), b_10)
xs_11_loop = (v_12_36).V1
continue go__go_9_35_57
__t37 = gopurs_runtime.Value{}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return __t37
}
}()
})
})
var go__go_10_38_58 gopurs_runtime.Value
go__go_10_38_58 = gopurs_runtime.Func(func(b_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_11_loop gopurs_runtime.Value = b_11_loop_val
var xs_12_loop gopurs_runtime.Value = xs_12_loop_val
go__go_10_38_58:
for {
if false { continue go__go_10_38_58 }
var b_11 gopurs_runtime.Value = b_11_loop
_ = b_11
var xs_12 gopurs_runtime.Value = xs_12_loop
_ = xs_12
// TAST (Let): v_13_39 -> *Constructor_Data_List_Lazy_Types_Cons
v_13_39 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_12))
_ = v_13_39
var __t41 gopurs_runtime.Value
{
if (v_13_39 == nil) {
__t41 = b_11
goto end_branch_41
} else {

}
}
{
if (v_13_39 != nil) {
// TAST (Let): __local_var_14_40 -> gopurs_runtime.Value
__local_var_14_40 := (v_13_39).V0
_ = __local_var_14_40
b_11_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_14_40, b_11})}
}))
xs_12_loop = (v_13_39).V1
continue go__go_10_38_58
__t41 = gopurs_runtime.Value{}
goto end_branch_41
} else {

}
}
{
__t41 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_41:
return __t41
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_31.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_32.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0)), gopurs_runtime.Apply2(go__go_9_35_57, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_10_38_58, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1)))
})
})
})})}
}()
	})
	return cache_Data_List_Lazy_Types_traversableNonEmpty
}

var cache_Data_List_Lazy_Types_traversableNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_traversableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2))
_ = __local_var_3_1
// TAST (Let): Apply0_4_2 -> *Constructor_Control_Apply_Apply
Apply0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
var go__go_6_4_59 gopurs_runtime.Value
go__go_6_4_59 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_4_59:
for {
if false { continue go__go_6_4_59 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_5 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_5
var __t6 gopurs_runtime.Value
{
if (v_9_5 == nil) {
__t6 = b_7
goto end_branch_6
} else {

}
}
{
if (v_9_5 != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), Get_Data_List_Lazy_Types_cons(), (v_9_5).V0), b_7)
xs_8_loop = (v_9_5).V1
continue go__go_6_4_59
__t6 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_7_7_60 gopurs_runtime.Value
go__go_7_7_60 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_7_60:
for {
if false { continue go__go_7_7_60 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_8 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_8
var __t10 gopurs_runtime.Value
{
if (v_10_8 == nil) {
__t10 = b_8
goto end_branch_10
} else {

}
}
{
if (v_10_8 != nil) {
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := (v_10_8).V0
_ = __local_var_11_9
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_11_9, b_8})}
}))
xs_9_loop = (v_10_8).V1
continue go__go_7_7_60
__t10 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xxs_3))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), (__local_var_3_1).V0), gopurs_runtime.Apply2(go__go_6_4_59, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_7_7_60, Get_Data_List_Lazy_Types_nil(), (__local_var_3_1).V1))))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_11 -> *Constructor_Data_Functor_Functor
Functor0_1_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_11
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_12 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3))
_ = __local_var_4_12
// TAST (Let): Apply0_5_13 -> *Constructor_Control_Apply_Apply
Apply0_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_13
// TAST (Let): Functor0_6_14 -> *Constructor_Data_Functor_Functor
Functor0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_14
var go__go_7_15_61 gopurs_runtime.Value
go__go_7_15_61 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_15_61:
for {
if false { continue go__go_7_15_61 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_16 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_16
var __t17 gopurs_runtime.Value
{
if (v_10_16 == nil) {
__t17 = b_8
goto end_branch_17
} else {

}
}
{
if (v_10_16 != nil) {
b_8_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_13.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_14.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_2, (v_10_16).V0)), b_8)
xs_9_loop = (v_10_16).V1
continue go__go_7_15_61
__t17 = gopurs_runtime.Value{}
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}
}()
})
})
var go__go_8_18_62 gopurs_runtime.Value
go__go_8_18_62 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var xs_10_loop gopurs_runtime.Value = xs_10_loop_val
go__go_8_18_62:
for {
if false { continue go__go_8_18_62 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_19 -> *Constructor_Data_List_Lazy_Types_Cons
v_11_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_10))
_ = v_11_19
var __t21 gopurs_runtime.Value
{
if (v_11_19 == nil) {
__t21 = b_9
goto end_branch_21
} else {

}
}
{
if (v_11_19 != nil) {
// TAST (Let): __local_var_12_20 -> gopurs_runtime.Value
__local_var_12_20 := (v_11_19).V0
_ = __local_var_12_20
b_9_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_12_20, b_9})}
}))
xs_10_loop = (v_11_19).V1
continue go__go_8_18_62
__t21 = gopurs_runtime.Value{}
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_11.V0), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xxs_4))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_2, (__local_var_4_12).V0)), gopurs_runtime.Apply2(go__go_7_15_61, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_8_18_62, Get_Data_List_Lazy_Types_nil(), (__local_var_4_12).V1))))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_traversableNonEmptyList
}

var cache_Data_List_Lazy_Types_traversableWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexList sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(&Constructor_Data_TraversableWithIndex_TraversableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](Get_Data_List_Lazy_Types_foldableWithIndexList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_List_Lazy_Types_functorWithIndexList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Lazy_Types_traversableList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_4_2
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_4_63 gopurs_runtime.Value
go__go_6_4_63 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_4_63:
for {
if false { continue go__go_6_4_63 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_5 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
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
// TAST (Let): __local_var_10_6 -> gopurs_runtime.Value
__local_var_10_6 := (*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V1
_ = __local_var_10_6
// TAST (Let): __local_var_11_7 -> gopurs_runtime.Value
__local_var_11_7 := (v_9_5).V0
_ = __local_var_11_7
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_11_7, __local_var_10_6})}
}))})}
xs_8_loop = (v_9_5).V1
continue go__go_6_4_63
__t8 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_6_3 -> *Constructor_Data_Tuple_Tuple
v_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_6_4_63, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_5))
_ = v_6_3
var go__go_7_9_64 gopurs_runtime.Value
go__go_7_9_64 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_9_64:
for {
if false { continue go__go_7_9_64 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_10 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
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
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply2(f_3, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V0.IntVal) - (1)), (v_10_10).V0)), (*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V1)})}
xs_9_loop = (v_10_10).V1
continue go__go_7_9_64
__t11 = gopurs_runtime.Value{}
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
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_7_9_64, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_6_3).V0.IntVal), __local_var_4_2})}, (v_6_3).V1).UnsafePtr).V1
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexList
}

var cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexNonEmpty sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_1 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_0_1 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1
_ = __local_var_2_2
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_2)
_ = __local_var_4_3
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_3.Type == 9 && __local_var_4_3.IntVal == 218341868 && __local_var_4_3.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (__local_var_4_3.Type == 9 && __local_var_4_3.IntVal == 218341868 && __local_var_4_3.UnsafePtr != nil) {
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_3.UnsafePtr).V1)}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))})}
})
})}
_ = functorNonEmpty1_0_1
// TAST (Let): functorWithIndex1_0_0 -> *Constructor_Data_FunctorWithIndex_FunctorWithIndex
functorWithIndex1_0_0 := &Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_0_1)}
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_6_65 gopurs_runtime.Value
go__go_3_6_65 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_6_65:
for {
if false { continue go__go_3_6_65 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := (*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1
_ = __local_var_7_8
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := (v_6_7).V0
_ = __local_var_8_9
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_9, __local_var_7_8})}
}))})}
xs_5_loop = (v_6_7).V1
continue go__go_3_6_65
__t10 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_3_5 -> *Constructor_Data_Tuple_Tuple
v_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_6_65, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
_ = v_3_5
var go__go_4_11_66 gopurs_runtime.Value
go__go_4_11_66 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_11_66:
for {
if false { continue go__go_4_11_66 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_12
var __t15 gopurs_runtime.Value
{
if (v_7_12 == nil) {
__t15 = b_5
goto end_branch_15
} else {

}
}
{
if (v_7_12 != nil) {
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1
_ = __local_var_8_13
// TAST (Let): __local_var_9_14 -> gopurs_runtime.Value
__local_var_9_14 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1))})}, (v_7_12).V0)
_ = __local_var_9_14
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_14, __local_var_8_13})}
}))})}
xs_6_loop = (v_7_12).V1
continue go__go_4_11_66
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_11_66, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_5).V0.IntVal), Get_Data_List_Lazy_Types_nil()})}, (v_3_5).V1).UnsafePtr).V1})}
})
})}
_ = functorWithIndex1_0_0
// TAST (Let): foldableNonEmpty1_1_17 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_17 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_18 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_18
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_19
var go__go_6_20_67 gopurs_runtime.Value
go__go_6_20_67 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_20_67:
for {
if false { continue go__go_6_20_67 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_21 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_21
var __t22 gopurs_runtime.Value
{
if (v_9_21 == nil) {
__t22 = b_7
goto end_branch_22
} else {

}
}
{
if (v_9_21 != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_19.V0), b_7, gopurs_runtime.Apply(f_3, (v_9_21).V0))
xs_8_loop = (v_9_21).V1
continue go__go_6_20_67
__t22 = gopurs_runtime.Value{}
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_18.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_6_20_67, gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_23_68 gopurs_runtime.Value
go__go_4_23_68 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_23_68:
for {
if false { continue go__go_4_23_68 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_24 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_24
var __t25 gopurs_runtime.Value
{
if (v_7_24 == nil) {
__t25 = b_5
goto end_branch_25
} else {

}
}
{
if (v_7_24 != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, b_5, (v_7_24).V0)
xs_6_loop = (v_7_24).V1
continue go__go_4_23_68
__t25 = gopurs_runtime.Value{}
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_23_68, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_26_69 gopurs_runtime.Value
go__go_4_26_69 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_26_69:
for {
if false { continue go__go_4_26_69 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_27 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_27
var __t28 gopurs_runtime.Value
{
if (v_7_27 == nil) {
__t28 = b_5
goto end_branch_28
} else {

}
}
{
if (v_7_27 != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, (v_7_27).V0, b_5)
xs_6_loop = (v_7_27).V1
continue go__go_4_26_69
__t28 = gopurs_runtime.Value{}
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
}
}()
})
})
var go__go_5_29_70 gopurs_runtime.Value
go__go_5_29_70 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_29_70:
for {
if false { continue go__go_5_29_70 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_30 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_30 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_30
var __t32 gopurs_runtime.Value
{
if (v_8_30 == nil) {
__t32 = b_6
goto end_branch_32
} else {

}
}
{
if (v_8_30 != nil) {
// TAST (Let): __local_var_9_31 -> gopurs_runtime.Value
__local_var_9_31 := (v_8_30).V0
_ = __local_var_9_31
b_6_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_31, b_6})}
}))
xs_7_loop = (v_8_30).V1
continue go__go_5_29_70
__t32 = gopurs_runtime.Value{}
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
return __t32
}
}()
})
})
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_4_26_69, b_2, gopurs_runtime.Apply2(go__go_5_29_70, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_1_17
// TAST (Let): foldableWithIndexNonEmpty1_1_16 -> *Constructor_Data_FoldableWithIndex_FoldableWithIndex
foldableWithIndexNonEmpty1_1_16 := &Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_17)}
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_33 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_33 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_33
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_34 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_34 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_34
var go__go_7_35_71 gopurs_runtime.Value
go__go_7_35_71 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_35_71:
for {
if false { continue go__go_7_35_71 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_36 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_36 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_36
var __t37 gopurs_runtime.Value
{
if (v_10_36 == nil) {
__t37 = b_8
goto end_branch_37
} else {

}
}
{
if (v_10_36 != nil) {
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_34.V0), (*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V1, gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V0.IntVal)})}, (v_10_36).V0))})}
xs_9_loop = (v_10_36).V1
continue go__go_7_35_71
__t37 = gopurs_runtime.Value{}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return __t37
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_33.V0), gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_7_35_71, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.RecordGet(dictMonoid_2, "mempty")})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1).UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_38_72 gopurs_runtime.Value
go__go_5_38_72 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_38_72:
for {
if false { continue go__go_5_38_72 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_39 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_39
var __t40 gopurs_runtime.Value
{
if (v_8_39 == nil) {
__t40 = b_6
goto end_branch_40
} else {

}
}
{
if (v_8_39 != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1, (v_8_39).V0)})}
xs_7_loop = (v_8_39).V1
continue go__go_5_38_72
__t40 = gopurs_runtime.Value{}
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
}
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_38_72, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0)})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_42_73 gopurs_runtime.Value
go__go_5_42_73 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_42_73:
for {
if false { continue go__go_5_42_73 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_43 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_43 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_43
var __t46 gopurs_runtime.Value
{
if (v_8_43 == nil) {
__t46 = b_6
goto end_branch_46
} else {

}
}
{
if (v_8_43 != nil) {
// TAST (Let): __local_var_9_44 -> gopurs_runtime.Value
__local_var_9_44 := (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1
_ = __local_var_9_44
// TAST (Let): __local_var_10_45 -> gopurs_runtime.Value
__local_var_10_45 := (v_8_43).V0
_ = __local_var_10_45
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_10_45, __local_var_9_44})}
}))})}
xs_7_loop = (v_8_43).V1
continue go__go_5_42_73
__t46 = gopurs_runtime.Value{}
goto end_branch_46
} else {

}
}
{
__t46 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_46:
return __t46
}
}()
})
})
// TAST (Let): v_5_41 -> *Constructor_Data_Tuple_Tuple
v_5_41 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_5_42_73, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
_ = v_5_41
var go__go_6_47_74 gopurs_runtime.Value
go__go_6_47_74 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_47_74:
for {
if false { continue go__go_6_47_74 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_48 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_48 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_48
var __t49 gopurs_runtime.Value
{
if (v_9_48 == nil) {
__t49 = b_7
goto end_branch_49
} else {

}
}
{
if (v_9_48 != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal) - (1))})}, (v_9_48).V0, (*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V1)})}
xs_8_loop = (v_9_48).V1
continue go__go_6_47_74
__t49 = gopurs_runtime.Value{}
goto end_branch_49
} else {

}
}
{
__t49 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_49:
return __t49
}
}()
})
})
return gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_6_47_74, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_5_41).V0.IntVal), b_3})}, (v_5_41).V1).UnsafePtr).V1)
})
})
})}
_ = foldableWithIndexNonEmpty1_1_16
// TAST (Let): functorNonEmpty1_2_51 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_2_51 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_52 -> gopurs_runtime.Value
__local_var_4_52 := (*Constructor_Data_NonEmpty_NonEmpty)(m_3.UnsafePtr).V1
_ = __local_var_4_52
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(m_3.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_53 -> gopurs_runtime.Value
__local_var_6_53 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_52)
_ = __local_var_6_53
var __t54 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_6_53.Type == 9 && __local_var_6_53.IntVal == 218341868 && __local_var_6_53.UnsafePtr == nil) {
__t54 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_54
} else {

}
}
{
if (__local_var_6_53.Type == 9 && __local_var_6_53.IntVal == 218341868 && __local_var_6_53.UnsafePtr != nil) {
__t54 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_53.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_2, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_53.UnsafePtr).V1)}
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_54:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t54)}
}))})}
})
})}
_ = functorNonEmpty1_2_51
// TAST (Let): foldableNonEmpty1_3_55 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_3_55 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_56 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_56 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_56
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_57 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_57 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_57
var go__go_8_58_75 gopurs_runtime.Value
go__go_8_58_75 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var xs_10_loop gopurs_runtime.Value = xs_10_loop_val
go__go_8_58_75:
for {
if false { continue go__go_8_58_75 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_59 -> *Constructor_Data_List_Lazy_Types_Cons
v_11_59 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_10))
_ = v_11_59
var __t60 gopurs_runtime.Value
{
if (v_11_59 == nil) {
__t60 = b_9
goto end_branch_60
} else {

}
}
{
if (v_11_59 != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_57.V0), b_9, gopurs_runtime.Apply(f_5, (v_11_59).V0))
xs_10_loop = (v_11_59).V1
continue go__go_8_58_75
__t60 = gopurs_runtime.Value{}
goto end_branch_60
} else {

}
}
{
__t60 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_60:
return __t60
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_56.V0), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_8_58_75, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_61_76 gopurs_runtime.Value
go__go_6_61_76 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_61_76:
for {
if false { continue go__go_6_61_76 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_62 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_62 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_62
var __t63 gopurs_runtime.Value
{
if (v_9_62 == nil) {
__t63 = b_7
goto end_branch_63
} else {

}
}
{
if (v_9_62 != nil) {
b_7_loop = gopurs_runtime.Apply2(f_3, b_7, (v_9_62).V0)
xs_8_loop = (v_9_62).V1
continue go__go_6_61_76
__t63 = gopurs_runtime.Value{}
goto end_branch_63
} else {

}
}
{
__t63 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_63:
return __t63
}
}()
})
})
return gopurs_runtime.Apply2(go__go_6_61_76, gopurs_runtime.Apply2(f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_64_77 gopurs_runtime.Value
go__go_6_64_77 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_64_77:
for {
if false { continue go__go_6_64_77 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_65 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_65 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_65
var __t66 gopurs_runtime.Value
{
if (v_9_65 == nil) {
__t66 = b_7
goto end_branch_66
} else {

}
}
{
if (v_9_65 != nil) {
b_7_loop = gopurs_runtime.Apply2(f_3, (v_9_65).V0, b_7)
xs_8_loop = (v_9_65).V1
continue go__go_6_64_77
__t66 = gopurs_runtime.Value{}
goto end_branch_66
} else {

}
}
{
__t66 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_66:
return __t66
}
}()
})
})
var go__go_7_67_78 gopurs_runtime.Value
go__go_7_67_78 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_67_78:
for {
if false { continue go__go_7_67_78 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_68 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_68 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_68
var __t70 gopurs_runtime.Value
{
if (v_10_68 == nil) {
__t70 = b_8
goto end_branch_70
} else {

}
}
{
if (v_10_68 != nil) {
// TAST (Let): __local_var_11_69 -> gopurs_runtime.Value
__local_var_11_69 := (v_10_68).V0
_ = __local_var_11_69
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_11_69, b_8})}
}))
xs_9_loop = (v_10_68).V1
continue go__go_7_67_78
__t70 = gopurs_runtime.Value{}
goto end_branch_70
} else {

}
}
{
__t70 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_70:
return __t70
}
}()
})
})
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_6_64_77, b_4, gopurs_runtime.Apply2(go__go_7_67_78, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_3_55
// TAST (Let): traversableNonEmpty1_2_50 -> *Constructor_Data_Traversable_Traversable
traversableNonEmpty1_2_50 := &Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_3_55)}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_2_51)}
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_71 -> *Constructor_Control_Apply_Apply
Apply0_5_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_71
// TAST (Let): Functor0_6_72 -> *Constructor_Data_Functor_Functor
Functor0_6_72 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_72
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_8_73 -> *Constructor_Control_Apply_Apply
Apply0_8_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_8_73
// TAST (Let): Functor0_9_74 -> *Constructor_Data_Functor_Functor
Functor0_9_74 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_74
var go__go_10_75_79 gopurs_runtime.Value
go__go_10_75_79 = gopurs_runtime.Func(func(b_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_11_loop gopurs_runtime.Value = b_11_loop_val
var xs_12_loop gopurs_runtime.Value = xs_12_loop_val
go__go_10_75_79:
for {
if false { continue go__go_10_75_79 }
var b_11 gopurs_runtime.Value = b_11_loop
_ = b_11
var xs_12 gopurs_runtime.Value = xs_12_loop
_ = xs_12
// TAST (Let): v_13_76 -> *Constructor_Data_List_Lazy_Types_Cons
v_13_76 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_12))
_ = v_13_76
var __t77 gopurs_runtime.Value
{
if (v_13_76 == nil) {
__t77 = b_11
goto end_branch_77
} else {

}
}
{
if (v_13_76 != nil) {
b_11_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_8_73.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_74.V0), Get_Data_List_Lazy_Types_cons(), (v_13_76).V0), b_11)
xs_12_loop = (v_13_76).V1
continue go__go_10_75_79
__t77 = gopurs_runtime.Value{}
goto end_branch_77
} else {

}
}
{
__t77 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_77:
return __t77
}
}()
})
})
var go__go_11_78_80 gopurs_runtime.Value
go__go_11_78_80 = gopurs_runtime.Func(func(b_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_12_loop gopurs_runtime.Value = b_12_loop_val
var xs_13_loop gopurs_runtime.Value = xs_13_loop_val
go__go_11_78_80:
for {
if false { continue go__go_11_78_80 }
var b_12 gopurs_runtime.Value = b_12_loop
_ = b_12
var xs_13 gopurs_runtime.Value = xs_13_loop
_ = xs_13
// TAST (Let): v_14_79 -> *Constructor_Data_List_Lazy_Types_Cons
v_14_79 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_13))
_ = v_14_79
var __t81 gopurs_runtime.Value
{
if (v_14_79 == nil) {
__t81 = b_12
goto end_branch_81
} else {

}
}
{
if (v_14_79 != nil) {
// TAST (Let): __local_var_15_80 -> gopurs_runtime.Value
__local_var_15_80 := (v_14_79).V0
_ = __local_var_15_80
b_12_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_15_80, b_12})}
}))
xs_13_loop = (v_14_79).V1
continue go__go_11_78_80
__t81 = gopurs_runtime.Value{}
goto end_branch_81
} else {

}
}
{
__t81 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_81:
return __t81
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_71.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_72.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_10_75_79, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "pure"), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_11_78_80, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1)))
})
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_82 -> *Constructor_Control_Apply_Apply
Apply0_5_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_82
// TAST (Let): Functor0_6_83 -> *Constructor_Data_Functor_Functor
Functor0_6_83 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_83
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_9_84 -> *Constructor_Control_Apply_Apply
Apply0_9_84 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_84
// TAST (Let): Functor0_10_85 -> *Constructor_Data_Functor_Functor
Functor0_10_85 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_85
var go__go_11_86_81 gopurs_runtime.Value
go__go_11_86_81 = gopurs_runtime.Func(func(b_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_12_loop gopurs_runtime.Value = b_12_loop_val
var xs_13_loop gopurs_runtime.Value = xs_13_loop_val
go__go_11_86_81:
for {
if false { continue go__go_11_86_81 }
var b_12 gopurs_runtime.Value = b_12_loop
_ = b_12
var xs_13 gopurs_runtime.Value = xs_13_loop
_ = xs_13
// TAST (Let): v_14_87 -> *Constructor_Data_List_Lazy_Types_Cons
v_14_87 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_13))
_ = v_14_87
var __t88 gopurs_runtime.Value
{
if (v_14_87 == nil) {
__t88 = b_12
goto end_branch_88
} else {

}
}
{
if (v_14_87 != nil) {
b_12_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_84.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_85.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_7, (v_14_87).V0)), b_12)
xs_13_loop = (v_14_87).V1
continue go__go_11_86_81
__t88 = gopurs_runtime.Value{}
goto end_branch_88
} else {

}
}
{
__t88 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_88:
return __t88
}
}()
})
})
var go__go_12_89_82 gopurs_runtime.Value
go__go_12_89_82 = gopurs_runtime.Func(func(b_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_13_loop gopurs_runtime.Value = b_13_loop_val
var xs_14_loop gopurs_runtime.Value = xs_14_loop_val
go__go_12_89_82:
for {
if false { continue go__go_12_89_82 }
var b_13 gopurs_runtime.Value = b_13_loop
_ = b_13
var xs_14 gopurs_runtime.Value = xs_14_loop
_ = xs_14
// TAST (Let): v_15_90 -> *Constructor_Data_List_Lazy_Types_Cons
v_15_90 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_14))
_ = v_15_90
var __t92 gopurs_runtime.Value
{
if (v_15_90 == nil) {
__t92 = b_13
goto end_branch_92
} else {

}
}
{
if (v_15_90 != nil) {
// TAST (Let): __local_var_16_91 -> gopurs_runtime.Value
__local_var_16_91 := (v_15_90).V0
_ = __local_var_16_91
b_13_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_16_91, b_13})}
}))
xs_14_loop = (v_15_90).V1
continue go__go_12_89_82
__t92 = gopurs_runtime.Value{}
goto end_branch_92
} else {

}
}
{
__t92 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_92:
return __t92
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_82.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_83.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)), gopurs_runtime.Apply2(go__go_11_86_81, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "pure"), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_12_89_82, Get_Data_List_Lazy_Types_nil(), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1)))
})
})
})}
_ = traversableNonEmpty1_2_50
return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(&Constructor_Data_TraversableWithIndex_TraversableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(foldableWithIndexNonEmpty1_1_16)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(functorWithIndex1_0_0)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(traversableNonEmpty1_2_50)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_93 -> *Constructor_Control_Apply_Apply
Apply0_4_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_93
// TAST (Let): Functor0_5_94 -> *Constructor_Data_Functor_Functor
Functor0_5_94 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_94
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_8_95 -> *Constructor_Control_Apply_Apply
Apply0_8_95 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_8_95
// TAST (Let): Functor0_9_96 -> *Constructor_Data_Functor_Functor
Functor0_9_96 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_96
var go__go_10_98_83 gopurs_runtime.Value
go__go_10_98_83 = gopurs_runtime.Func(func(b_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_11_loop gopurs_runtime.Value = b_11_loop_val
var xs_12_loop gopurs_runtime.Value = xs_12_loop_val
go__go_10_98_83:
for {
if false { continue go__go_10_98_83 }
var b_11 gopurs_runtime.Value = b_11_loop
_ = b_11
var xs_12 gopurs_runtime.Value = xs_12_loop
_ = xs_12
// TAST (Let): v_13_99 -> *Constructor_Data_List_Lazy_Types_Cons
v_13_99 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_12))
_ = v_13_99
var __t102 gopurs_runtime.Value
{
if (v_13_99 == nil) {
__t102 = b_11
goto end_branch_102
} else {

}
}
{
if (v_13_99 != nil) {
// TAST (Let): __local_var_14_100 -> gopurs_runtime.Value
__local_var_14_100 := (*Constructor_Data_Tuple_Tuple)(b_11.UnsafePtr).V1
_ = __local_var_14_100
// TAST (Let): __local_var_15_101 -> gopurs_runtime.Value
__local_var_15_101 := (v_13_99).V0
_ = __local_var_15_101
b_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_11.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_15_101, __local_var_14_100})}
}))})}
xs_12_loop = (v_13_99).V1
continue go__go_10_98_83
__t102 = gopurs_runtime.Value{}
goto end_branch_102
} else {

}
}
{
__t102 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_102:
return __t102
}
}()
})
})
// TAST (Let): v_10_97 -> *Constructor_Data_Tuple_Tuple
v_10_97 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_10_98_83, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
_ = v_10_97
var go__go_11_103_84 gopurs_runtime.Value
go__go_11_103_84 = gopurs_runtime.Func(func(b_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_12_loop gopurs_runtime.Value = b_12_loop_val
var xs_13_loop gopurs_runtime.Value = xs_13_loop_val
go__go_11_103_84:
for {
if false { continue go__go_11_103_84 }
var b_12 gopurs_runtime.Value = b_12_loop
_ = b_12
var xs_13 gopurs_runtime.Value = xs_13_loop
_ = xs_13
// TAST (Let): v_14_104 -> *Constructor_Data_List_Lazy_Types_Cons
v_14_104 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_13))
_ = v_14_104
var __t105 gopurs_runtime.Value
{
if (v_14_104 == nil) {
__t105 = b_12
goto end_branch_105
} else {

}
}
{
if (v_14_104 != nil) {
b_12_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_12.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_8_95.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_96.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply2(f_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_12.UnsafePtr).V0.IntVal) - (1))})}, (v_14_104).V0)), (*Constructor_Data_Tuple_Tuple)(b_12.UnsafePtr).V1)})}
xs_13_loop = (v_14_104).V1
continue go__go_11_103_84
__t105 = gopurs_runtime.Value{}
goto end_branch_105
} else {

}
}
{
__t105 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_105:
return __t105
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_93.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_94.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply2(f_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_11_103_84, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_10_97).V0.IntVal), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_List_Lazy_Types_nil())})}, (v_10_97).V1).UnsafePtr).V1)
})
})
})})}
}()
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty
}

var cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(&Constructor_Data_TraversableWithIndex_TraversableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Lazy_Types_traversableNonEmptyList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3))
_ = __local_var_4_1
// TAST (Let): Apply0_5_2 -> *Constructor_Control_Apply_Apply
Apply0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_2
// TAST (Let): Functor0_6_3 -> *Constructor_Data_Functor_Functor
Functor0_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_3
var go__go_7_5_85 gopurs_runtime.Value
go__go_7_5_85 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_5_85:
for {
if false { continue go__go_7_5_85 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_6
var __t9 gopurs_runtime.Value
{
if (v_10_6 == nil) {
__t9 = b_8
goto end_branch_9
} else {

}
}
{
if (v_10_6 != nil) {
// TAST (Let): __local_var_11_7 -> gopurs_runtime.Value
__local_var_11_7 := (*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V1
_ = __local_var_11_7
// TAST (Let): __local_var_12_8 -> gopurs_runtime.Value
__local_var_12_8 := (v_10_6).V0
_ = __local_var_12_8
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_12_8, __local_var_11_7})}
}))})}
xs_9_loop = (v_10_6).V1
continue go__go_7_5_85
__t9 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_7_4 -> *Constructor_Data_Tuple_Tuple
v_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_7_5_85, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (__local_var_4_1).V1))
_ = v_7_4
var go__go_8_10_86 gopurs_runtime.Value
go__go_8_10_86 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var xs_10_loop gopurs_runtime.Value = xs_10_loop_val
go__go_8_10_86:
for {
if false { continue go__go_8_10_86 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_11 -> *Constructor_Data_List_Lazy_Types_Cons
v_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_10))
_ = v_11_11
var __t12 gopurs_runtime.Value
{
if (v_11_11 == nil) {
__t12 = b_9
goto end_branch_12
} else {

}
}
{
if (v_11_11 != nil) {
b_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_9.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_3.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply2(f_2, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_9.UnsafePtr).V0.IntVal) - (1))), (v_11_11).V0)), (*Constructor_Data_Tuple_Tuple)(b_9.UnsafePtr).V1)})}
xs_10_loop = (v_11_11).V1
continue go__go_8_10_86
__t12 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xxs_4))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply2(f_2, gopurs_runtime.Int(0), (__local_var_4_1).V0)), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_8_10_86, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_7_4).V0.IntVal), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())})}, (v_7_4).V1).UnsafePtr).V1))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_unfoldable1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1List sync.Once
func Get_Data_List_Lazy_Types_unfoldable1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1List.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1List = func() gopurs_runtime.Value {
var go__go_0_0_87 gopurs_runtime.Value
_ = go__go_0_0_87
go__go_0_0_87 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Tuple_Tuple
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (v1_4_1).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(go__go_0_0_87, f_1, (*Constructor_Data_Maybe_Just)((v1_4_1).V1.UnsafePtr).V0)
_ = __local_var_5_3
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v1_4_1).V0, __local_var_5_3})}
}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (v1_4_1).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v1_4_1).V0, Get_Data_List_Lazy_Types_nil()})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t5)
}))
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, go__go_0_0_87})}
}()
	})
	return cache_Data_List_Lazy_Types_unfoldable1List
}

var cache_Data_List_Lazy_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldableList sync.Once
func Get_Data_List_Lazy_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldableList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldableList = func() gopurs_runtime.Value {
var go__go_0_0_88 gopurs_runtime.Value
_ = go__go_0_0_88
go__go_0_0_88 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Maybe_Just
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t3 gopurs_runtime.Value
{
if (v1_4_1 == nil) {
__t3 = Get_Data_List_Lazy_Types_nil()
goto end_branch_3
} else {

}
}
{
if (v1_4_1 != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply2(go__go_0_0_88, f_1, (*Constructor_Data_Tuple_Tuple)((v1_4_1).V0.UnsafePtr).V1)
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_Tuple_Tuple)((v1_4_1).V0.UnsafePtr).V0, __local_var_5_2})}
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t3)
}))
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable_Unfoldable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Lazy_Types_unfoldable1List()))}
}), go__go_0_0_88})}
}()
	})
	return cache_Data_List_Lazy_Types_unfoldableList
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmpty sync.Once
func Get_Data_List_Lazy_Types_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Tuple_Tuple
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, b_1))
_ = __local_var_2_1
var go__go_3_2_89 gopurs_runtime.Value
_ = go__go_3_2_89
go__go_3_2_89 = gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_7_3 -> *Constructor_Data_Maybe_Just
v1_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_4, b_5))
_ = v1_7_3
var __t5 gopurs_runtime.Value
{
if (v1_7_3 == nil) {
__t5 = Get_Data_List_Lazy_Types_nil()
goto end_branch_5
} else {

}
}
{
if (v1_7_3 != nil) {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := gopurs_runtime.Apply2(go__go_3_2_89, f_4, (*Constructor_Data_Tuple_Tuple)((v1_7_3).V0.UnsafePtr).V1)
_ = __local_var_8_4
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_Tuple_Tuple)((v1_7_3).V0.UnsafePtr).V0, __local_var_8_4})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t5)
}))
})
})
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Tuple_Tuple
var __local_var_2_0 *Constructor_Data_Tuple_Tuple = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (__local_var_2_1).V0, gopurs_runtime.Apply2(go__go_3_2_89, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Maybe_Just)(v1_4.UnsafePtr).V0)}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
}), (__local_var_2_1).V1)})})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (__local_var_2_0).V0, (__local_var_2_0).V1})}
})
})})}
	})
	return cache_Data_List_Lazy_Types_unfoldable1NonEmpty
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Lazy_Types_unfoldable1NonEmpty()).V0), f_0, b_1)))}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_unfoldable1NonEmptyList
}

var cache_Data_List_Lazy_Types_comonadNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_comonadNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_comonadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_comonadNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_comonadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](Get_Data_List_Lazy_Types_extendNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V0
})})}
	})
	return cache_Data_List_Lazy_Types_comonadNonEmptyList
}

var cache_Data_List_Lazy_Types_monadList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadList sync.Once
func Get_Data_List_Lazy_Types_monadList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadList.Do(func() {
		cache_Data_List_Lazy_Types_monadList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Lazy_Types_applicativeList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()))}
})})}
	})
	return cache_Data_List_Lazy_Types_monadList
}

var cache_Data_List_Lazy_Types_bindList gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList sync.Once
func Get_Data_List_Lazy_Types_bindList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList.Do(func() {
		cache_Data_List_Lazy_Types_bindList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyList()))}
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t5 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t5 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_5
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0)
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1, f_1)
_ = __local_var_5_2
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_7_3
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_7_3.Type == 9 && __local_var_7_3.IntVal == 218341868 && __local_var_7_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2))
goto end_branch_4
} else {

}
}
{
if (__local_var_7_3.Type == 9 && __local_var_7_3.IntVal == 218341868 && __local_var_7_3.UnsafePtr != nil) {
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_3.UnsafePtr).V1, __local_var_5_2)}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_bindList
}

var cache_Data_List_Lazy_Types_applyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList sync.Once
func Get_Data_List_Lazy_Types_applyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList.Do(func() {
		cache_Data_List_Lazy_Types_applyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(f_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1)
_ = __local_var_4_1
var __t9 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t9 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_9
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Lazy_Types_applicativeList()).V1), gopurs_runtime.Apply(f_prime_2, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0))
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Lazy_Types_applicativeList()).V1), gopurs_runtime.Apply(f_prime_2, a_prime_6))
}))
_ = __local_var_6_3
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2)
_ = __local_var_8_4
var __t8 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
goto end_branch_8
} else {

}
}
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr != nil) {
// TAST (Let): __local_var_9_5 -> gopurs_runtime.Value
__local_var_9_5 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_4.UnsafePtr).V1
_ = __local_var_9_5
__t8 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_4.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_6 -> gopurs_runtime.Value
__local_var_11_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_5)
_ = __local_var_11_6
var __t7 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_6.Type == 9 && __local_var_11_6.IntVal == 218341868 && __local_var_11_6.UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
goto end_branch_7
} else {

}
}
{
if (__local_var_11_6.Type == 9 && __local_var_11_6.IntVal == 218341868 && __local_var_11_6.UnsafePtr != nil) {
__t7 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_6.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_6.UnsafePtr).V1, __local_var_6_3)}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t7)}
}))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t8)}
}))))
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t9)}
}))
})
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_10 -> gopurs_runtime.Value
__local_var_4_10 := gopurs_runtime.Apply(Get_Data_Lazy_force(), f_0)
_ = __local_var_4_10
var __t18 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_10.Type == 9 && __local_var_4_10.IntVal == 218341868 && __local_var_4_10.UnsafePtr == nil) {
__t18 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_18
} else {

}
}
{
if (__local_var_4_10.Type == 9 && __local_var_4_10.IntVal == 218341868 && __local_var_4_10.UnsafePtr != nil) {
// TAST (Let): __local_var_5_11 -> gopurs_runtime.Value
__local_var_5_11 := gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_10.UnsafePtr).V0)
_ = __local_var_5_11
// TAST (Let): __local_var_6_12 -> gopurs_runtime.Value
__local_var_6_12 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_10.UnsafePtr).V1, __local_var_2_0)
_ = __local_var_6_12
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_11)
_ = __local_var_8_13
var __t17 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_8_13.Type == 9 && __local_var_8_13.IntVal == 218341868 && __local_var_8_13.UnsafePtr == nil) {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_12))
goto end_branch_17
} else {

}
}
{
if (__local_var_8_13.Type == 9 && __local_var_8_13.IntVal == 218341868 && __local_var_8_13.UnsafePtr != nil) {
// TAST (Let): __local_var_9_14 -> gopurs_runtime.Value
__local_var_9_14 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_13.UnsafePtr).V1
_ = __local_var_9_14
__t17 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_13.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_15 -> gopurs_runtime.Value
__local_var_11_15 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_14)
_ = __local_var_11_15
var __t16 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_15.Type == 9 && __local_var_11_15.IntVal == 218341868 && __local_var_11_15.UnsafePtr == nil) {
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_12))
goto end_branch_16
} else {

}
}
{
if (__local_var_11_15.Type == 9 && __local_var_11_15.IntVal == 218341868 && __local_var_11_15.UnsafePtr != nil) {
__t16 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_15.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_15.UnsafePtr).V1, __local_var_6_12)}
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t16)}
}))}
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t17)}
}))))
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t18)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_applyList
}

var cache_Data_List_Lazy_Types_applicativeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeList sync.Once
func Get_Data_List_Lazy_Types_applicativeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeList.Do(func() {
		cache_Data_List_Lazy_Types_applicativeList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyList()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, a_0, Get_Data_List_Lazy_Types_nil()})}
}))
})})}
	})
	return cache_Data_List_Lazy_Types_applicativeList
}

var cache_Data_List_Lazy_Types_applyNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_applyNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_applyNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_1)
_ = v2_2_0
// TAST (Let): v3_3_1 -> gopurs_runtime.Value
v3_3_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v3_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V0
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_NonEmpty_NonEmpty)(v3_3_1.UnsafePtr).V0
_ = __local_var_6_4
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := (*Constructor_Data_NonEmpty_NonEmpty)(v3_3_1.UnsafePtr).V1
_ = __local_var_7_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_7 -> gopurs_runtime.Value
__local_var_9_7 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_4_2, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_9_7
// TAST (Let): __local_var_10_8 -> gopurs_runtime.Value
__local_var_10_8 := gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_9 -> gopurs_runtime.Value
__local_var_12_9 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_7)
_ = __local_var_12_9
var __t33 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_12_9.Type == 9 && __local_var_12_9.IntVal == 218341868 && __local_var_12_9.UnsafePtr == nil) {
__t33 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_33
} else {

}
}
{
if (__local_var_12_9.Type == 9 && __local_var_12_9.IntVal == 218341868 && __local_var_12_9.UnsafePtr != nil) {
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := gopurs_runtime.Apply(f_prime_10, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_9.UnsafePtr).V0)
_ = __local_var_13_11
// TAST (Let): __local_var_13_10 -> gopurs_runtime.Value
__local_var_13_10 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_13_11, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_13_10
// TAST (Let): __local_var_14_13 -> gopurs_runtime.Value
__local_var_14_13 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_9.UnsafePtr).V1
_ = __local_var_14_13
// TAST (Let): __local_var_14_12 -> gopurs_runtime.Value
__local_var_14_12 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_14 -> gopurs_runtime.Value
__local_var_16_14 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_13)
_ = __local_var_16_14
var __t24 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_14.Type == 9 && __local_var_16_14.IntVal == 218341868 && __local_var_16_14.UnsafePtr == nil) {
__t24 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_24
} else {

}
}
{
if (__local_var_16_14.Type == 9 && __local_var_16_14.IntVal == 218341868 && __local_var_16_14.UnsafePtr != nil) {
// TAST (Let): __local_var_17_16 -> gopurs_runtime.Value
__local_var_17_16 := gopurs_runtime.Apply(f_prime_10, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_14.UnsafePtr).V0)
_ = __local_var_17_16
// TAST (Let): __local_var_17_15 -> gopurs_runtime.Value
__local_var_17_15 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_17_16, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_17_15
// TAST (Let): __local_var_18_17 -> gopurs_runtime.Value
__local_var_18_17 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_14.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_18 -> gopurs_runtime.Value
__local_var_19_18 := gopurs_runtime.Apply(f_prime_10, a_prime_18)
_ = __local_var_19_18
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_19_18, Get_Data_List_Lazy_Types_nil()})}
}))
}))
_ = __local_var_18_17
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_19 -> gopurs_runtime.Value
__local_var_20_19 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_15)
_ = __local_var_20_19
var __t23 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_19.Type == 9 && __local_var_20_19.IntVal == 218341868 && __local_var_20_19.UnsafePtr == nil) {
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_17))
goto end_branch_23
} else {

}
}
{
if (__local_var_20_19.Type == 9 && __local_var_20_19.IntVal == 218341868 && __local_var_20_19.UnsafePtr != nil) {
// TAST (Let): __local_var_21_20 -> gopurs_runtime.Value
__local_var_21_20 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_19.UnsafePtr).V1
_ = __local_var_21_20
__t23 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_19.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_21 -> gopurs_runtime.Value
__local_var_23_21 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_21_20)
_ = __local_var_23_21
var __t22 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_23_21.Type == 9 && __local_var_23_21.IntVal == 218341868 && __local_var_23_21.UnsafePtr == nil) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_17))
goto end_branch_22
} else {

}
}
{
if (__local_var_23_21.Type == 9 && __local_var_23_21.IntVal == 218341868 && __local_var_23_21.UnsafePtr != nil) {
__t22 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_21.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_21.UnsafePtr).V1, __local_var_18_17)}
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t22)}
}))}
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t23)}
}))))
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_24:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t24)}
}))
_ = __local_var_14_12
__t33 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_25 -> gopurs_runtime.Value
__local_var_16_25 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_10)
_ = __local_var_16_25
var __t32 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_25.Type == 9 && __local_var_16_25.IntVal == 218341868 && __local_var_16_25.UnsafePtr == nil) {
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_12))
goto end_branch_32
} else {

}
}
{
if (__local_var_16_25.Type == 9 && __local_var_16_25.IntVal == 218341868 && __local_var_16_25.UnsafePtr != nil) {
// TAST (Let): __local_var_17_26 -> gopurs_runtime.Value
__local_var_17_26 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_25.UnsafePtr).V1
_ = __local_var_17_26
__t32 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_25.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_27 -> gopurs_runtime.Value
__local_var_19_27 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_26)
_ = __local_var_19_27
var __t31 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_19_27.Type == 9 && __local_var_19_27.IntVal == 218341868 && __local_var_19_27.UnsafePtr == nil) {
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_12))
goto end_branch_31
} else {

}
}
{
if (__local_var_19_27.Type == 9 && __local_var_19_27.IntVal == 218341868 && __local_var_19_27.UnsafePtr != nil) {
// TAST (Let): __local_var_20_28 -> gopurs_runtime.Value
__local_var_20_28 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_27.UnsafePtr).V1
_ = __local_var_20_28
__t31 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_27.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_22_29 -> gopurs_runtime.Value
__local_var_22_29 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_20_28)
_ = __local_var_22_29
var __t30 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_22_29.Type == 9 && __local_var_22_29.IntVal == 218341868 && __local_var_22_29.UnsafePtr == nil) {
__t30 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_12))
goto end_branch_30
} else {

}
}
{
if (__local_var_22_29.Type == 9 && __local_var_22_29.IntVal == 218341868 && __local_var_22_29.UnsafePtr != nil) {
__t30 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_22_29.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_22_29.UnsafePtr).V1, __local_var_14_12)}
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_30:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t30)}
}))}
goto end_branch_31
} else {

}
}
{
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_31:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t31)}
}))}
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_32:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t32)}
}))))
goto end_branch_33
} else {

}
}
{
__t33 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_33:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t33)}
}))
})
_ = __local_var_10_8
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_34 -> gopurs_runtime.Value
__local_var_12_34 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_5)
_ = __local_var_12_34
var __t55 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_12_34.Type == 9 && __local_var_12_34.IntVal == 218341868 && __local_var_12_34.UnsafePtr == nil) {
__t55 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_55
} else {

}
}
{
if (__local_var_12_34.Type == 9 && __local_var_12_34.IntVal == 218341868 && __local_var_12_34.UnsafePtr != nil) {
// TAST (Let): __local_var_13_35 -> gopurs_runtime.Value
__local_var_13_35 := gopurs_runtime.Apply(__local_var_10_8, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_34.UnsafePtr).V0)
_ = __local_var_13_35
// TAST (Let): __local_var_14_37 -> gopurs_runtime.Value
__local_var_14_37 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_34.UnsafePtr).V1
_ = __local_var_14_37
// TAST (Let): __local_var_14_36 -> gopurs_runtime.Value
__local_var_14_36 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_38 -> gopurs_runtime.Value
__local_var_16_38 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_37)
_ = __local_var_16_38
var __t46 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_38.Type == 9 && __local_var_16_38.IntVal == 218341868 && __local_var_16_38.UnsafePtr == nil) {
__t46 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_46
} else {

}
}
{
if (__local_var_16_38.Type == 9 && __local_var_16_38.IntVal == 218341868 && __local_var_16_38.UnsafePtr != nil) {
// TAST (Let): __local_var_17_39 -> gopurs_runtime.Value
__local_var_17_39 := gopurs_runtime.Apply(__local_var_10_8, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_38.UnsafePtr).V0)
_ = __local_var_17_39
// TAST (Let): __local_var_18_40 -> gopurs_runtime.Value
__local_var_18_40 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_38.UnsafePtr).V1, __local_var_10_8)
_ = __local_var_18_40
__t46 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_41 -> gopurs_runtime.Value
__local_var_20_41 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_39)
_ = __local_var_20_41
var __t45 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_41.Type == 9 && __local_var_20_41.IntVal == 218341868 && __local_var_20_41.UnsafePtr == nil) {
__t45 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_40))
goto end_branch_45
} else {

}
}
{
if (__local_var_20_41.Type == 9 && __local_var_20_41.IntVal == 218341868 && __local_var_20_41.UnsafePtr != nil) {
// TAST (Let): __local_var_21_42 -> gopurs_runtime.Value
__local_var_21_42 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_41.UnsafePtr).V1
_ = __local_var_21_42
__t45 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_41.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_43 -> gopurs_runtime.Value
__local_var_23_43 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_21_42)
_ = __local_var_23_43
var __t44 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_23_43.Type == 9 && __local_var_23_43.IntVal == 218341868 && __local_var_23_43.UnsafePtr == nil) {
__t44 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_40))
goto end_branch_44
} else {

}
}
{
if (__local_var_23_43.Type == 9 && __local_var_23_43.IntVal == 218341868 && __local_var_23_43.UnsafePtr != nil) {
__t44 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_43.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_43.UnsafePtr).V1, __local_var_18_40)}
goto end_branch_44
} else {

}
}
{
__t44 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_44:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t44)}
}))}
goto end_branch_45
} else {

}
}
{
__t45 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_45:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t45)}
}))))
goto end_branch_46
} else {

}
}
{
__t46 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_46:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t46)}
}))
_ = __local_var_14_36
__t55 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_47 -> gopurs_runtime.Value
__local_var_16_47 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_35)
_ = __local_var_16_47
var __t54 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_47.Type == 9 && __local_var_16_47.IntVal == 218341868 && __local_var_16_47.UnsafePtr == nil) {
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_36))
goto end_branch_54
} else {

}
}
{
if (__local_var_16_47.Type == 9 && __local_var_16_47.IntVal == 218341868 && __local_var_16_47.UnsafePtr != nil) {
// TAST (Let): __local_var_17_48 -> gopurs_runtime.Value
__local_var_17_48 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_47.UnsafePtr).V1
_ = __local_var_17_48
__t54 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_47.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_49 -> gopurs_runtime.Value
__local_var_19_49 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_48)
_ = __local_var_19_49
var __t53 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_19_49.Type == 9 && __local_var_19_49.IntVal == 218341868 && __local_var_19_49.UnsafePtr == nil) {
__t53 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_36))
goto end_branch_53
} else {

}
}
{
if (__local_var_19_49.Type == 9 && __local_var_19_49.IntVal == 218341868 && __local_var_19_49.UnsafePtr != nil) {
// TAST (Let): __local_var_20_50 -> gopurs_runtime.Value
__local_var_20_50 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_49.UnsafePtr).V1
_ = __local_var_20_50
__t53 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_49.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_22_51 -> gopurs_runtime.Value
__local_var_22_51 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_20_50)
_ = __local_var_22_51
var __t52 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_22_51.Type == 9 && __local_var_22_51.IntVal == 218341868 && __local_var_22_51.UnsafePtr == nil) {
__t52 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_36))
goto end_branch_52
} else {

}
}
{
if (__local_var_22_51.Type == 9 && __local_var_22_51.IntVal == 218341868 && __local_var_22_51.UnsafePtr != nil) {
__t52 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_22_51.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_22_51.UnsafePtr).V1, __local_var_14_36)}
goto end_branch_52
} else {

}
}
{
__t52 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_52:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t52)}
}))}
goto end_branch_53
} else {

}
}
{
__t53 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_53:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t53)}
}))}
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_54:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t54)}
}))))
goto end_branch_55
} else {

}
}
{
__t55 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_55:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t55)}
}))
_ = __local_var_9_6
// TAST (Let): __local_var_10_57 -> gopurs_runtime.Value
__local_var_10_57 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_6_4, __local_var_7_5})}
}))
_ = __local_var_10_57
// TAST (Let): __local_var_11_58 -> gopurs_runtime.Value
__local_var_11_58 := gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_59 -> gopurs_runtime.Value
__local_var_13_59 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3)
_ = __local_var_13_59
var __t83 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_59.Type == 9 && __local_var_13_59.IntVal == 218341868 && __local_var_13_59.UnsafePtr == nil) {
__t83 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_83
} else {

}
}
{
if (__local_var_13_59.Type == 9 && __local_var_13_59.IntVal == 218341868 && __local_var_13_59.UnsafePtr != nil) {
// TAST (Let): __local_var_14_61 -> gopurs_runtime.Value
__local_var_14_61 := gopurs_runtime.Apply(f_prime_11, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_59.UnsafePtr).V0)
_ = __local_var_14_61
// TAST (Let): __local_var_14_60 -> gopurs_runtime.Value
__local_var_14_60 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_14_61, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_14_60
// TAST (Let): __local_var_15_63 -> gopurs_runtime.Value
__local_var_15_63 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_59.UnsafePtr).V1
_ = __local_var_15_63
// TAST (Let): __local_var_15_62 -> gopurs_runtime.Value
__local_var_15_62 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_64 -> gopurs_runtime.Value
__local_var_17_64 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_63)
_ = __local_var_17_64
var __t74 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_64.Type == 9 && __local_var_17_64.IntVal == 218341868 && __local_var_17_64.UnsafePtr == nil) {
__t74 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_74
} else {

}
}
{
if (__local_var_17_64.Type == 9 && __local_var_17_64.IntVal == 218341868 && __local_var_17_64.UnsafePtr != nil) {
// TAST (Let): __local_var_18_66 -> gopurs_runtime.Value
__local_var_18_66 := gopurs_runtime.Apply(f_prime_11, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_64.UnsafePtr).V0)
_ = __local_var_18_66
// TAST (Let): __local_var_18_65 -> gopurs_runtime.Value
__local_var_18_65 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_18_66, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_18_65
// TAST (Let): __local_var_19_67 -> gopurs_runtime.Value
__local_var_19_67 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_64.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_68 -> gopurs_runtime.Value
__local_var_20_68 := gopurs_runtime.Apply(f_prime_11, a_prime_19)
_ = __local_var_20_68
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_20_68, Get_Data_List_Lazy_Types_nil()})}
}))
}))
_ = __local_var_19_67
__t74 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_69 -> gopurs_runtime.Value
__local_var_21_69 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_65)
_ = __local_var_21_69
var __t73 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_21_69.Type == 9 && __local_var_21_69.IntVal == 218341868 && __local_var_21_69.UnsafePtr == nil) {
__t73 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_67))
goto end_branch_73
} else {

}
}
{
if (__local_var_21_69.Type == 9 && __local_var_21_69.IntVal == 218341868 && __local_var_21_69.UnsafePtr != nil) {
// TAST (Let): __local_var_22_70 -> gopurs_runtime.Value
__local_var_22_70 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_21_69.UnsafePtr).V1
_ = __local_var_22_70
__t73 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_21_69.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_71 -> gopurs_runtime.Value
__local_var_24_71 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_22_70)
_ = __local_var_24_71
var __t72 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_24_71.Type == 9 && __local_var_24_71.IntVal == 218341868 && __local_var_24_71.UnsafePtr == nil) {
__t72 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_67))
goto end_branch_72
} else {

}
}
{
if (__local_var_24_71.Type == 9 && __local_var_24_71.IntVal == 218341868 && __local_var_24_71.UnsafePtr != nil) {
__t72 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_24_71.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_24_71.UnsafePtr).V1, __local_var_19_67)}
goto end_branch_72
} else {

}
}
{
__t72 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_72:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t72)}
}))}
goto end_branch_73
} else {

}
}
{
__t73 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_73:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t73)}
}))))
goto end_branch_74
} else {

}
}
{
__t74 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_74:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t74)}
}))
_ = __local_var_15_62
__t83 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_75 -> gopurs_runtime.Value
__local_var_17_75 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_60)
_ = __local_var_17_75
var __t82 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_75.Type == 9 && __local_var_17_75.IntVal == 218341868 && __local_var_17_75.UnsafePtr == nil) {
__t82 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_62))
goto end_branch_82
} else {

}
}
{
if (__local_var_17_75.Type == 9 && __local_var_17_75.IntVal == 218341868 && __local_var_17_75.UnsafePtr != nil) {
// TAST (Let): __local_var_18_76 -> gopurs_runtime.Value
__local_var_18_76 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_75.UnsafePtr).V1
_ = __local_var_18_76
__t82 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_75.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_77 -> gopurs_runtime.Value
__local_var_20_77 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_76)
_ = __local_var_20_77
var __t81 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_77.Type == 9 && __local_var_20_77.IntVal == 218341868 && __local_var_20_77.UnsafePtr == nil) {
__t81 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_62))
goto end_branch_81
} else {

}
}
{
if (__local_var_20_77.Type == 9 && __local_var_20_77.IntVal == 218341868 && __local_var_20_77.UnsafePtr != nil) {
// TAST (Let): __local_var_21_78 -> gopurs_runtime.Value
__local_var_21_78 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_77.UnsafePtr).V1
_ = __local_var_21_78
__t81 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_77.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_79 -> gopurs_runtime.Value
__local_var_23_79 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_21_78)
_ = __local_var_23_79
var __t80 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_23_79.Type == 9 && __local_var_23_79.IntVal == 218341868 && __local_var_23_79.UnsafePtr == nil) {
__t80 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_62))
goto end_branch_80
} else {

}
}
{
if (__local_var_23_79.Type == 9 && __local_var_23_79.IntVal == 218341868 && __local_var_23_79.UnsafePtr != nil) {
__t80 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_79.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_79.UnsafePtr).V1, __local_var_15_62)}
goto end_branch_80
} else {

}
}
{
__t80 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_80:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t80)}
}))}
goto end_branch_81
} else {

}
}
{
__t81 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_81:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t81)}
}))}
goto end_branch_82
} else {

}
}
{
__t82 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_82:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t82)}
}))))
goto end_branch_83
} else {

}
}
{
__t83 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_83:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t83)}
}))
})
_ = __local_var_11_58
// TAST (Let): __local_var_10_56 -> gopurs_runtime.Value
__local_var_10_56 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_84 -> gopurs_runtime.Value
__local_var_13_84 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_57)
_ = __local_var_13_84
var __t105 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_84.Type == 9 && __local_var_13_84.IntVal == 218341868 && __local_var_13_84.UnsafePtr == nil) {
__t105 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_105
} else {

}
}
{
if (__local_var_13_84.Type == 9 && __local_var_13_84.IntVal == 218341868 && __local_var_13_84.UnsafePtr != nil) {
// TAST (Let): __local_var_14_85 -> gopurs_runtime.Value
__local_var_14_85 := gopurs_runtime.Apply(__local_var_11_58, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_84.UnsafePtr).V0)
_ = __local_var_14_85
// TAST (Let): __local_var_15_87 -> gopurs_runtime.Value
__local_var_15_87 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_84.UnsafePtr).V1
_ = __local_var_15_87
// TAST (Let): __local_var_15_86 -> gopurs_runtime.Value
__local_var_15_86 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_88 -> gopurs_runtime.Value
__local_var_17_88 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_87)
_ = __local_var_17_88
var __t96 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_88.Type == 9 && __local_var_17_88.IntVal == 218341868 && __local_var_17_88.UnsafePtr == nil) {
__t96 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_96
} else {

}
}
{
if (__local_var_17_88.Type == 9 && __local_var_17_88.IntVal == 218341868 && __local_var_17_88.UnsafePtr != nil) {
// TAST (Let): __local_var_18_89 -> gopurs_runtime.Value
__local_var_18_89 := gopurs_runtime.Apply(__local_var_11_58, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_88.UnsafePtr).V0)
_ = __local_var_18_89
// TAST (Let): __local_var_19_90 -> gopurs_runtime.Value
__local_var_19_90 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_88.UnsafePtr).V1, __local_var_11_58)
_ = __local_var_19_90
__t96 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_91 -> gopurs_runtime.Value
__local_var_21_91 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_89)
_ = __local_var_21_91
var __t95 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_21_91.Type == 9 && __local_var_21_91.IntVal == 218341868 && __local_var_21_91.UnsafePtr == nil) {
__t95 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_90))
goto end_branch_95
} else {

}
}
{
if (__local_var_21_91.Type == 9 && __local_var_21_91.IntVal == 218341868 && __local_var_21_91.UnsafePtr != nil) {
// TAST (Let): __local_var_22_92 -> gopurs_runtime.Value
__local_var_22_92 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_21_91.UnsafePtr).V1
_ = __local_var_22_92
__t95 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_21_91.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_93 -> gopurs_runtime.Value
__local_var_24_93 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_22_92)
_ = __local_var_24_93
var __t94 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_24_93.Type == 9 && __local_var_24_93.IntVal == 218341868 && __local_var_24_93.UnsafePtr == nil) {
__t94 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_90))
goto end_branch_94
} else {

}
}
{
if (__local_var_24_93.Type == 9 && __local_var_24_93.IntVal == 218341868 && __local_var_24_93.UnsafePtr != nil) {
__t94 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_24_93.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_24_93.UnsafePtr).V1, __local_var_19_90)}
goto end_branch_94
} else {

}
}
{
__t94 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_94:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t94)}
}))}
goto end_branch_95
} else {

}
}
{
__t95 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_95:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t95)}
}))))
goto end_branch_96
} else {

}
}
{
__t96 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_96:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t96)}
}))
_ = __local_var_15_86
__t105 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_97 -> gopurs_runtime.Value
__local_var_17_97 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_85)
_ = __local_var_17_97
var __t104 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_97.Type == 9 && __local_var_17_97.IntVal == 218341868 && __local_var_17_97.UnsafePtr == nil) {
__t104 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_86))
goto end_branch_104
} else {

}
}
{
if (__local_var_17_97.Type == 9 && __local_var_17_97.IntVal == 218341868 && __local_var_17_97.UnsafePtr != nil) {
// TAST (Let): __local_var_18_98 -> gopurs_runtime.Value
__local_var_18_98 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_97.UnsafePtr).V1
_ = __local_var_18_98
__t104 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_97.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_99 -> gopurs_runtime.Value
__local_var_20_99 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_98)
_ = __local_var_20_99
var __t103 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_99.Type == 9 && __local_var_20_99.IntVal == 218341868 && __local_var_20_99.UnsafePtr == nil) {
__t103 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_86))
goto end_branch_103
} else {

}
}
{
if (__local_var_20_99.Type == 9 && __local_var_20_99.IntVal == 218341868 && __local_var_20_99.UnsafePtr != nil) {
// TAST (Let): __local_var_21_100 -> gopurs_runtime.Value
__local_var_21_100 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_99.UnsafePtr).V1
_ = __local_var_21_100
__t103 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_99.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_101 -> gopurs_runtime.Value
__local_var_23_101 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_21_100)
_ = __local_var_23_101
var __t102 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_23_101.Type == 9 && __local_var_23_101.IntVal == 218341868 && __local_var_23_101.UnsafePtr == nil) {
__t102 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_86))
goto end_branch_102
} else {

}
}
{
if (__local_var_23_101.Type == 9 && __local_var_23_101.IntVal == 218341868 && __local_var_23_101.UnsafePtr != nil) {
__t102 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_101.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_101.UnsafePtr).V1, __local_var_15_86)}
goto end_branch_102
} else {

}
}
{
__t102 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_102:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t102)}
}))}
goto end_branch_103
} else {

}
}
{
__t103 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_103:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t103)}
}))}
goto end_branch_104
} else {

}
}
{
__t104 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_104:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t104)}
}))))
goto end_branch_105
} else {

}
}
{
__t105 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_105:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t105)}
}))
_ = __local_var_10_56
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_106 -> gopurs_runtime.Value
__local_var_12_106 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_6)
_ = __local_var_12_106
var __t107 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_12_106.Type == 9 && __local_var_12_106.IntVal == 218341868 && __local_var_12_106.UnsafePtr == nil) {
__t107 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_56))
goto end_branch_107
} else {

}
}
{
if (__local_var_12_106.Type == 9 && __local_var_12_106.IntVal == 218341868 && __local_var_12_106.UnsafePtr != nil) {
__t107 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_106.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_106.UnsafePtr).V1, __local_var_10_56)}
goto end_branch_107
} else {

}
}
{
__t107 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_107:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t107)}
}))})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_applyNonEmptyList
}

var cache_Data_List_Lazy_Types_bindNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_bindNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_bindNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): v2_4_2 -> *Constructor_Data_NonEmpty_NonEmpty
v2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V0)))
_ = v2_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (v2_4_2).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (v2_4_2).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1)
_ = __local_var_9_6
var __t14 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_9_6.Type == 9 && __local_var_9_6.IntVal == 218341868 && __local_var_9_6.UnsafePtr == nil) {
__t14 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_14
} else {

}
}
{
if (__local_var_9_6.Type == 9 && __local_var_9_6.IntVal == 218341868 && __local_var_9_6.UnsafePtr != nil) {
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := Call_Data_List_Lazy_Types_toList(gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_9_6.UnsafePtr).V0))
_ = __local_var_10_7
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_9_6.UnsafePtr).V1, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_toList(gopurs_runtime.Apply(f_1, x_11))
}))
_ = __local_var_11_8
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_9 -> gopurs_runtime.Value
__local_var_13_9 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_7)
_ = __local_var_13_9
var __t13 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_9.Type == 9 && __local_var_13_9.IntVal == 218341868 && __local_var_13_9.UnsafePtr == nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_13
} else {

}
}
{
if (__local_var_13_9.Type == 9 && __local_var_13_9.IntVal == 218341868 && __local_var_13_9.UnsafePtr != nil) {
// TAST (Let): __local_var_14_10 -> gopurs_runtime.Value
__local_var_14_10 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_9.UnsafePtr).V1
_ = __local_var_14_10
__t13 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_9.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_11 -> gopurs_runtime.Value
__local_var_16_11 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_10)
_ = __local_var_16_11
var __t12 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_11.Type == 9 && __local_var_16_11.IntVal == 218341868 && __local_var_16_11.UnsafePtr == nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_12
} else {

}
}
{
if (__local_var_16_11.Type == 9 && __local_var_16_11.IntVal == 218341868 && __local_var_16_11.UnsafePtr != nil) {
__t12 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_11.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_11.UnsafePtr).V1, __local_var_11_8)}
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t12)}
}))}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t13)}
}))))
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t14)}
}))
_ = __local_var_8_5
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_5_3, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_15 -> gopurs_runtime.Value
__local_var_10_15 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4)
_ = __local_var_10_15
var __t16 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_10_15.Type == 9 && __local_var_10_15.IntVal == 218341868 && __local_var_10_15.UnsafePtr == nil) {
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_5))
goto end_branch_16
} else {

}
}
{
if (__local_var_10_15.Type == 9 && __local_var_10_15.IntVal == 218341868 && __local_var_10_15.UnsafePtr != nil) {
__t16 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_10_15.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_10_15.UnsafePtr).V1, __local_var_8_5)}
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t16)}
}))})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_bindNonEmptyList
}

var cache_Data_List_Lazy_Types_altNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_altNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_altNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := Call_Data_List_Lazy_Types_toList(as_prime_1)
_ = __local_var_6_3
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_3_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2)
_ = __local_var_8_4
var __t8 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
goto end_branch_8
} else {

}
}
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 218341868 && __local_var_8_4.UnsafePtr != nil) {
// TAST (Let): __local_var_9_5 -> gopurs_runtime.Value
__local_var_9_5 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_4.UnsafePtr).V1
_ = __local_var_9_5
__t8 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_4.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_6 -> gopurs_runtime.Value
__local_var_11_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_5)
_ = __local_var_11_6
var __t7 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_6.Type == 9 && __local_var_11_6.IntVal == 218341868 && __local_var_11_6.UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
goto end_branch_7
} else {

}
}
{
if (__local_var_11_6.Type == 9 && __local_var_11_6.IntVal == 218341868 && __local_var_11_6.UnsafePtr != nil) {
__t7 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_6.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_6.UnsafePtr).V1, __local_var_6_3)}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t7)}
}))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t8)}
}))})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_altNonEmptyList
}

var cache_Data_List_Lazy_Types_altList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altList sync.Once
func Get_Data_List_Lazy_Types_altList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altList.Do(func() {
		cache_Data_List_Lazy_Types_altList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1, ys_1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_altList
}

var cache_Data_List_Lazy_Types_plusList gopurs_runtime.Value
var once_Data_List_Lazy_Types_plusList sync.Once
func Get_Data_List_Lazy_Types_plusList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_plusList.Do(func() {
		cache_Data_List_Lazy_Types_plusList = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Data_List_Lazy_Types_altList()))}
}), Get_Data_List_Lazy_Types_nil()})}
	})
	return cache_Data_List_Lazy_Types_plusList
}

var cache_Data_List_Lazy_Types_alternativeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_alternativeList sync.Once
func Get_Data_List_Lazy_Types_alternativeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_alternativeList.Do(func() {
		cache_Data_List_Lazy_Types_alternativeList = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Lazy_Types_applicativeList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_List_Lazy_Types_plusList()))}
})})}
	})
	return cache_Data_List_Lazy_Types_alternativeList
}

var cache_Data_List_Lazy_Types_monadPlusList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadPlusList sync.Once
func Get_Data_List_Lazy_Types_monadPlusList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadPlusList.Do(func() {
		cache_Data_List_Lazy_Types_monadPlusList = gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(&Constructor_Control_MonadPlus_MonadPlus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](Get_Data_List_Lazy_Types_alternativeList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Data_List_Lazy_Types_monadList()))}
})})}
	})
	return cache_Data_List_Lazy_Types_monadPlusList
}

var cache_Data_List_Lazy_Types_applicativeNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_applicativeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_applicativeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_0, Get_Data_List_Lazy_Types_nil()})}
}))
})})}
	})
	return cache_Data_List_Lazy_Types_applicativeNonEmptyList
}

var cache_Data_List_Lazy_Types_monadNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_monadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_monadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Lazy_Types_applicativeNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindNonEmptyList()))}
})})}
	})
	return cache_Data_List_Lazy_Types_monadNonEmptyList
}

var cache_Data_List_Lazy_Types_altList__3343465296 gopurs_runtime.Value
var once_Data_List_Lazy_Types_altList__3343465296 sync.Once
func Get_Data_List_Lazy_Types_altList__3343465296() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altList__3343465296.Do(func() {
		cache_Data_List_Lazy_Types_altList__3343465296 = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1, ys_1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_altList__3343465296
}

var cache_Data_List_Lazy_Types_alternativeList__577105552 gopurs_runtime.Value
var once_Data_List_Lazy_Types_alternativeList__577105552 sync.Once
func Get_Data_List_Lazy_Types_alternativeList__577105552() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_alternativeList__577105552.Do(func() {
		cache_Data_List_Lazy_Types_alternativeList__577105552 = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Lazy_Types_applicativeList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_List_Lazy_Types_plusList()))}
})})}
	})
	return cache_Data_List_Lazy_Types_alternativeList__577105552
}

var cache_Data_List_Lazy_Types_applicativeList__4243212624 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeList__4243212624 sync.Once
func Get_Data_List_Lazy_Types_applicativeList__4243212624() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeList__4243212624.Do(func() {
		cache_Data_List_Lazy_Types_applicativeList__4243212624 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyList()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, a_0, Get_Data_List_Lazy_Types_nil()})}
}))
})})}
	})
	return cache_Data_List_Lazy_Types_applicativeList__4243212624
}

var cache_Data_List_Lazy_Types_applicativeNonEmptyList__3829787129 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeNonEmptyList__3829787129 sync.Once
func Get_Data_List_Lazy_Types_applicativeNonEmptyList__3829787129() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeNonEmptyList__3829787129.Do(func() {
		cache_Data_List_Lazy_Types_applicativeNonEmptyList__3829787129 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_0, Get_Data_List_Lazy_Types_nil()})}
}))
})})}
	})
	return cache_Data_List_Lazy_Types_applicativeNonEmptyList__3829787129
}

var cache_Data_List_Lazy_Types_applyList__1470982352 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList__1470982352 sync.Once
func Get_Data_List_Lazy_Types_applyList__1470982352() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList__1470982352.Do(func() {
		cache_Data_List_Lazy_Types_applyList__1470982352 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(f_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1)
_ = __local_var_4_1
var __t11 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t11 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_11
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(f_prime_2, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0)
_ = __local_var_5_3
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_5_3, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_5_2
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(f_prime_2, a_prime_6)
_ = __local_var_7_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_7_5, Get_Data_List_Lazy_Types_nil()})}
}))
}))
_ = __local_var_6_4
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2)
_ = __local_var_8_6
var __t10 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_8_6.Type == 9 && __local_var_8_6.IntVal == 218341868 && __local_var_8_6.UnsafePtr == nil) {
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4))
goto end_branch_10
} else {

}
}
{
if (__local_var_8_6.Type == 9 && __local_var_8_6.IntVal == 218341868 && __local_var_8_6.UnsafePtr != nil) {
// TAST (Let): __local_var_9_7 -> gopurs_runtime.Value
__local_var_9_7 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_6.UnsafePtr).V1
_ = __local_var_9_7
__t10 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_6.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_7)
_ = __local_var_11_8
var __t9 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 218341868 && __local_var_11_8.UnsafePtr == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4))
goto end_branch_9
} else {

}
}
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 218341868 && __local_var_11_8.UnsafePtr != nil) {
__t9 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_8.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_8.UnsafePtr).V1, __local_var_6_4)}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t9)}
}))}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t10)}
}))))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t11)}
}))
})
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_12 -> gopurs_runtime.Value
__local_var_4_12 := gopurs_runtime.Apply(Get_Data_Lazy_force(), f_0)
_ = __local_var_4_12
var __t20 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_12.Type == 9 && __local_var_4_12.IntVal == 218341868 && __local_var_4_12.UnsafePtr == nil) {
__t20 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_20
} else {

}
}
{
if (__local_var_4_12.Type == 9 && __local_var_4_12.IntVal == 218341868 && __local_var_4_12.UnsafePtr != nil) {
// TAST (Let): __local_var_5_13 -> gopurs_runtime.Value
__local_var_5_13 := gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_12.UnsafePtr).V0)
_ = __local_var_5_13
// TAST (Let): __local_var_6_14 -> gopurs_runtime.Value
__local_var_6_14 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_12.UnsafePtr).V1, __local_var_2_0)
_ = __local_var_6_14
__t20 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_15 -> gopurs_runtime.Value
__local_var_8_15 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_13)
_ = __local_var_8_15
var __t19 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_8_15.Type == 9 && __local_var_8_15.IntVal == 218341868 && __local_var_8_15.UnsafePtr == nil) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_14))
goto end_branch_19
} else {

}
}
{
if (__local_var_8_15.Type == 9 && __local_var_8_15.IntVal == 218341868 && __local_var_8_15.UnsafePtr != nil) {
// TAST (Let): __local_var_9_16 -> gopurs_runtime.Value
__local_var_9_16 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_15.UnsafePtr).V1
_ = __local_var_9_16
__t19 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_15.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_17 -> gopurs_runtime.Value
__local_var_11_17 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_16)
_ = __local_var_11_17
var __t18 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_17.Type == 9 && __local_var_11_17.IntVal == 218341868 && __local_var_11_17.UnsafePtr == nil) {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_14))
goto end_branch_18
} else {

}
}
{
if (__local_var_11_17.Type == 9 && __local_var_11_17.IntVal == 218341868 && __local_var_11_17.UnsafePtr != nil) {
__t18 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_17.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_17.UnsafePtr).V1, __local_var_6_14)}
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t18)}
}))}
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t19)}
}))))
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_20:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t20)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_applyList__1470982352
}

var cache_Data_List_Lazy_Types_applyList__421157088 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList__421157088 sync.Once
func Get_Data_List_Lazy_Types_applyList__421157088() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList__421157088.Do(func() {
		cache_Data_List_Lazy_Types_applyList__421157088 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(f_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1)
_ = __local_var_4_1
var __t11 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t11 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_11
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(f_prime_2, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0)
_ = __local_var_5_3
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_5_3, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_5_2
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(f_prime_2, a_prime_6)
_ = __local_var_7_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_7_5, Get_Data_List_Lazy_Types_nil()})}
}))
}))
_ = __local_var_6_4
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2)
_ = __local_var_8_6
var __t10 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_8_6.Type == 9 && __local_var_8_6.IntVal == 218341868 && __local_var_8_6.UnsafePtr == nil) {
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4))
goto end_branch_10
} else {

}
}
{
if (__local_var_8_6.Type == 9 && __local_var_8_6.IntVal == 218341868 && __local_var_8_6.UnsafePtr != nil) {
// TAST (Let): __local_var_9_7 -> gopurs_runtime.Value
__local_var_9_7 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_6.UnsafePtr).V1
_ = __local_var_9_7
__t10 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_6.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_7)
_ = __local_var_11_8
var __t9 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 218341868 && __local_var_11_8.UnsafePtr == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4))
goto end_branch_9
} else {

}
}
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 218341868 && __local_var_11_8.UnsafePtr != nil) {
__t9 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_8.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_8.UnsafePtr).V1, __local_var_6_4)}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t9)}
}))}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t10)}
}))))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t11)}
}))
})
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_12 -> gopurs_runtime.Value
__local_var_4_12 := gopurs_runtime.Apply(Get_Data_Lazy_force(), f_0)
_ = __local_var_4_12
var __t20 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_12.Type == 9 && __local_var_4_12.IntVal == 218341868 && __local_var_4_12.UnsafePtr == nil) {
__t20 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_20
} else {

}
}
{
if (__local_var_4_12.Type == 9 && __local_var_4_12.IntVal == 218341868 && __local_var_4_12.UnsafePtr != nil) {
// TAST (Let): __local_var_5_13 -> gopurs_runtime.Value
__local_var_5_13 := gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_12.UnsafePtr).V0)
_ = __local_var_5_13
// TAST (Let): __local_var_6_14 -> gopurs_runtime.Value
__local_var_6_14 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_12.UnsafePtr).V1, __local_var_2_0)
_ = __local_var_6_14
__t20 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_15 -> gopurs_runtime.Value
__local_var_8_15 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_13)
_ = __local_var_8_15
var __t19 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_8_15.Type == 9 && __local_var_8_15.IntVal == 218341868 && __local_var_8_15.UnsafePtr == nil) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_14))
goto end_branch_19
} else {

}
}
{
if (__local_var_8_15.Type == 9 && __local_var_8_15.IntVal == 218341868 && __local_var_8_15.UnsafePtr != nil) {
// TAST (Let): __local_var_9_16 -> gopurs_runtime.Value
__local_var_9_16 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_15.UnsafePtr).V1
_ = __local_var_9_16
__t19 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_8_15.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_17 -> gopurs_runtime.Value
__local_var_11_17 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_16)
_ = __local_var_11_17
var __t18 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_17.Type == 9 && __local_var_11_17.IntVal == 218341868 && __local_var_11_17.UnsafePtr == nil) {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_14))
goto end_branch_18
} else {

}
}
{
if (__local_var_11_17.Type == 9 && __local_var_11_17.IntVal == 218341868 && __local_var_11_17.UnsafePtr != nil) {
__t18 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_17.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_17.UnsafePtr).V1, __local_var_6_14)}
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t18)}
}))}
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t19)}
}))))
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_20:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t20)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_applyList__421157088
}

var cache_Data_List_Lazy_Types_applyNonEmptyList__4077822201 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyNonEmptyList__4077822201 sync.Once
func Get_Data_List_Lazy_Types_applyNonEmptyList__4077822201() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyNonEmptyList__4077822201.Do(func() {
		cache_Data_List_Lazy_Types_applyNonEmptyList__4077822201 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_1)
_ = v2_2_0
// TAST (Let): v3_3_1 -> gopurs_runtime.Value
v3_3_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v3_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V0
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_NonEmpty_NonEmpty)(v3_3_1.UnsafePtr).V0
_ = __local_var_6_4
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := (*Constructor_Data_NonEmpty_NonEmpty)(v3_3_1.UnsafePtr).V1
_ = __local_var_7_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_7 -> gopurs_runtime.Value
__local_var_9_7 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_4_2, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_9_7
// TAST (Let): __local_var_10_8 -> gopurs_runtime.Value
__local_var_10_8 := gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_9 -> gopurs_runtime.Value
__local_var_12_9 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_7)
_ = __local_var_12_9
var __t33 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_12_9.Type == 9 && __local_var_12_9.IntVal == 218341868 && __local_var_12_9.UnsafePtr == nil) {
__t33 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_33
} else {

}
}
{
if (__local_var_12_9.Type == 9 && __local_var_12_9.IntVal == 218341868 && __local_var_12_9.UnsafePtr != nil) {
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := gopurs_runtime.Apply(f_prime_10, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_9.UnsafePtr).V0)
_ = __local_var_13_11
// TAST (Let): __local_var_13_10 -> gopurs_runtime.Value
__local_var_13_10 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_13_11, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_13_10
// TAST (Let): __local_var_14_13 -> gopurs_runtime.Value
__local_var_14_13 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_9.UnsafePtr).V1
_ = __local_var_14_13
// TAST (Let): __local_var_14_12 -> gopurs_runtime.Value
__local_var_14_12 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_14 -> gopurs_runtime.Value
__local_var_16_14 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_13)
_ = __local_var_16_14
var __t24 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_14.Type == 9 && __local_var_16_14.IntVal == 218341868 && __local_var_16_14.UnsafePtr == nil) {
__t24 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_24
} else {

}
}
{
if (__local_var_16_14.Type == 9 && __local_var_16_14.IntVal == 218341868 && __local_var_16_14.UnsafePtr != nil) {
// TAST (Let): __local_var_17_16 -> gopurs_runtime.Value
__local_var_17_16 := gopurs_runtime.Apply(f_prime_10, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_14.UnsafePtr).V0)
_ = __local_var_17_16
// TAST (Let): __local_var_17_15 -> gopurs_runtime.Value
__local_var_17_15 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_17_16, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_17_15
// TAST (Let): __local_var_18_17 -> gopurs_runtime.Value
__local_var_18_17 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_14.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_18 -> gopurs_runtime.Value
__local_var_19_18 := gopurs_runtime.Apply(f_prime_10, a_prime_18)
_ = __local_var_19_18
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_19_18, Get_Data_List_Lazy_Types_nil()})}
}))
}))
_ = __local_var_18_17
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_19 -> gopurs_runtime.Value
__local_var_20_19 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_15)
_ = __local_var_20_19
var __t23 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_19.Type == 9 && __local_var_20_19.IntVal == 218341868 && __local_var_20_19.UnsafePtr == nil) {
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_17))
goto end_branch_23
} else {

}
}
{
if (__local_var_20_19.Type == 9 && __local_var_20_19.IntVal == 218341868 && __local_var_20_19.UnsafePtr != nil) {
// TAST (Let): __local_var_21_20 -> gopurs_runtime.Value
__local_var_21_20 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_19.UnsafePtr).V1
_ = __local_var_21_20
__t23 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_19.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_21 -> gopurs_runtime.Value
__local_var_23_21 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_21_20)
_ = __local_var_23_21
var __t22 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_23_21.Type == 9 && __local_var_23_21.IntVal == 218341868 && __local_var_23_21.UnsafePtr == nil) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_17))
goto end_branch_22
} else {

}
}
{
if (__local_var_23_21.Type == 9 && __local_var_23_21.IntVal == 218341868 && __local_var_23_21.UnsafePtr != nil) {
__t22 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_21.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_21.UnsafePtr).V1, __local_var_18_17)}
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t22)}
}))}
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t23)}
}))))
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_24:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t24)}
}))
_ = __local_var_14_12
__t33 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_25 -> gopurs_runtime.Value
__local_var_16_25 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_10)
_ = __local_var_16_25
var __t32 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_25.Type == 9 && __local_var_16_25.IntVal == 218341868 && __local_var_16_25.UnsafePtr == nil) {
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_12))
goto end_branch_32
} else {

}
}
{
if (__local_var_16_25.Type == 9 && __local_var_16_25.IntVal == 218341868 && __local_var_16_25.UnsafePtr != nil) {
// TAST (Let): __local_var_17_26 -> gopurs_runtime.Value
__local_var_17_26 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_25.UnsafePtr).V1
_ = __local_var_17_26
__t32 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_25.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_27 -> gopurs_runtime.Value
__local_var_19_27 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_26)
_ = __local_var_19_27
var __t31 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_19_27.Type == 9 && __local_var_19_27.IntVal == 218341868 && __local_var_19_27.UnsafePtr == nil) {
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_12))
goto end_branch_31
} else {

}
}
{
if (__local_var_19_27.Type == 9 && __local_var_19_27.IntVal == 218341868 && __local_var_19_27.UnsafePtr != nil) {
// TAST (Let): __local_var_20_28 -> gopurs_runtime.Value
__local_var_20_28 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_27.UnsafePtr).V1
_ = __local_var_20_28
__t31 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_27.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_22_29 -> gopurs_runtime.Value
__local_var_22_29 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_20_28)
_ = __local_var_22_29
var __t30 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_22_29.Type == 9 && __local_var_22_29.IntVal == 218341868 && __local_var_22_29.UnsafePtr == nil) {
__t30 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_12))
goto end_branch_30
} else {

}
}
{
if (__local_var_22_29.Type == 9 && __local_var_22_29.IntVal == 218341868 && __local_var_22_29.UnsafePtr != nil) {
__t30 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_22_29.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_22_29.UnsafePtr).V1, __local_var_14_12)}
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_30:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t30)}
}))}
goto end_branch_31
} else {

}
}
{
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_31:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t31)}
}))}
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_32:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t32)}
}))))
goto end_branch_33
} else {

}
}
{
__t33 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_33:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t33)}
}))
})
_ = __local_var_10_8
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_34 -> gopurs_runtime.Value
__local_var_12_34 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_5)
_ = __local_var_12_34
var __t55 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_12_34.Type == 9 && __local_var_12_34.IntVal == 218341868 && __local_var_12_34.UnsafePtr == nil) {
__t55 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_55
} else {

}
}
{
if (__local_var_12_34.Type == 9 && __local_var_12_34.IntVal == 218341868 && __local_var_12_34.UnsafePtr != nil) {
// TAST (Let): __local_var_13_35 -> gopurs_runtime.Value
__local_var_13_35 := gopurs_runtime.Apply(__local_var_10_8, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_34.UnsafePtr).V0)
_ = __local_var_13_35
// TAST (Let): __local_var_14_37 -> gopurs_runtime.Value
__local_var_14_37 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_34.UnsafePtr).V1
_ = __local_var_14_37
// TAST (Let): __local_var_14_36 -> gopurs_runtime.Value
__local_var_14_36 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_38 -> gopurs_runtime.Value
__local_var_16_38 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_37)
_ = __local_var_16_38
var __t46 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_38.Type == 9 && __local_var_16_38.IntVal == 218341868 && __local_var_16_38.UnsafePtr == nil) {
__t46 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_46
} else {

}
}
{
if (__local_var_16_38.Type == 9 && __local_var_16_38.IntVal == 218341868 && __local_var_16_38.UnsafePtr != nil) {
// TAST (Let): __local_var_17_39 -> gopurs_runtime.Value
__local_var_17_39 := gopurs_runtime.Apply(__local_var_10_8, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_38.UnsafePtr).V0)
_ = __local_var_17_39
// TAST (Let): __local_var_18_40 -> gopurs_runtime.Value
__local_var_18_40 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_38.UnsafePtr).V1, __local_var_10_8)
_ = __local_var_18_40
__t46 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_41 -> gopurs_runtime.Value
__local_var_20_41 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_39)
_ = __local_var_20_41
var __t45 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_41.Type == 9 && __local_var_20_41.IntVal == 218341868 && __local_var_20_41.UnsafePtr == nil) {
__t45 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_40))
goto end_branch_45
} else {

}
}
{
if (__local_var_20_41.Type == 9 && __local_var_20_41.IntVal == 218341868 && __local_var_20_41.UnsafePtr != nil) {
// TAST (Let): __local_var_21_42 -> gopurs_runtime.Value
__local_var_21_42 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_41.UnsafePtr).V1
_ = __local_var_21_42
__t45 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_41.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_43 -> gopurs_runtime.Value
__local_var_23_43 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_21_42)
_ = __local_var_23_43
var __t44 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_23_43.Type == 9 && __local_var_23_43.IntVal == 218341868 && __local_var_23_43.UnsafePtr == nil) {
__t44 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_40))
goto end_branch_44
} else {

}
}
{
if (__local_var_23_43.Type == 9 && __local_var_23_43.IntVal == 218341868 && __local_var_23_43.UnsafePtr != nil) {
__t44 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_43.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_43.UnsafePtr).V1, __local_var_18_40)}
goto end_branch_44
} else {

}
}
{
__t44 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_44:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t44)}
}))}
goto end_branch_45
} else {

}
}
{
__t45 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_45:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t45)}
}))))
goto end_branch_46
} else {

}
}
{
__t46 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_46:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t46)}
}))
_ = __local_var_14_36
__t55 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_47 -> gopurs_runtime.Value
__local_var_16_47 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_35)
_ = __local_var_16_47
var __t54 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_47.Type == 9 && __local_var_16_47.IntVal == 218341868 && __local_var_16_47.UnsafePtr == nil) {
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_36))
goto end_branch_54
} else {

}
}
{
if (__local_var_16_47.Type == 9 && __local_var_16_47.IntVal == 218341868 && __local_var_16_47.UnsafePtr != nil) {
// TAST (Let): __local_var_17_48 -> gopurs_runtime.Value
__local_var_17_48 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_47.UnsafePtr).V1
_ = __local_var_17_48
__t54 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_47.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_49 -> gopurs_runtime.Value
__local_var_19_49 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_48)
_ = __local_var_19_49
var __t53 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_19_49.Type == 9 && __local_var_19_49.IntVal == 218341868 && __local_var_19_49.UnsafePtr == nil) {
__t53 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_36))
goto end_branch_53
} else {

}
}
{
if (__local_var_19_49.Type == 9 && __local_var_19_49.IntVal == 218341868 && __local_var_19_49.UnsafePtr != nil) {
// TAST (Let): __local_var_20_50 -> gopurs_runtime.Value
__local_var_20_50 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_49.UnsafePtr).V1
_ = __local_var_20_50
__t53 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_49.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_22_51 -> gopurs_runtime.Value
__local_var_22_51 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_20_50)
_ = __local_var_22_51
var __t52 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_22_51.Type == 9 && __local_var_22_51.IntVal == 218341868 && __local_var_22_51.UnsafePtr == nil) {
__t52 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_36))
goto end_branch_52
} else {

}
}
{
if (__local_var_22_51.Type == 9 && __local_var_22_51.IntVal == 218341868 && __local_var_22_51.UnsafePtr != nil) {
__t52 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_22_51.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_22_51.UnsafePtr).V1, __local_var_14_36)}
goto end_branch_52
} else {

}
}
{
__t52 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_52:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t52)}
}))}
goto end_branch_53
} else {

}
}
{
__t53 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_53:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t53)}
}))}
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_54:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t54)}
}))))
goto end_branch_55
} else {

}
}
{
__t55 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_55:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t55)}
}))
_ = __local_var_9_6
// TAST (Let): __local_var_10_57 -> gopurs_runtime.Value
__local_var_10_57 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_6_4, __local_var_7_5})}
}))
_ = __local_var_10_57
// TAST (Let): __local_var_11_58 -> gopurs_runtime.Value
__local_var_11_58 := gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_59 -> gopurs_runtime.Value
__local_var_13_59 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3)
_ = __local_var_13_59
var __t83 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_59.Type == 9 && __local_var_13_59.IntVal == 218341868 && __local_var_13_59.UnsafePtr == nil) {
__t83 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_83
} else {

}
}
{
if (__local_var_13_59.Type == 9 && __local_var_13_59.IntVal == 218341868 && __local_var_13_59.UnsafePtr != nil) {
// TAST (Let): __local_var_14_61 -> gopurs_runtime.Value
__local_var_14_61 := gopurs_runtime.Apply(f_prime_11, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_59.UnsafePtr).V0)
_ = __local_var_14_61
// TAST (Let): __local_var_14_60 -> gopurs_runtime.Value
__local_var_14_60 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_14_61, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_14_60
// TAST (Let): __local_var_15_63 -> gopurs_runtime.Value
__local_var_15_63 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_59.UnsafePtr).V1
_ = __local_var_15_63
// TAST (Let): __local_var_15_62 -> gopurs_runtime.Value
__local_var_15_62 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_64 -> gopurs_runtime.Value
__local_var_17_64 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_63)
_ = __local_var_17_64
var __t74 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_64.Type == 9 && __local_var_17_64.IntVal == 218341868 && __local_var_17_64.UnsafePtr == nil) {
__t74 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_74
} else {

}
}
{
if (__local_var_17_64.Type == 9 && __local_var_17_64.IntVal == 218341868 && __local_var_17_64.UnsafePtr != nil) {
// TAST (Let): __local_var_18_66 -> gopurs_runtime.Value
__local_var_18_66 := gopurs_runtime.Apply(f_prime_11, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_64.UnsafePtr).V0)
_ = __local_var_18_66
// TAST (Let): __local_var_18_65 -> gopurs_runtime.Value
__local_var_18_65 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_18_66, Get_Data_List_Lazy_Types_nil()})}
}))
_ = __local_var_18_65
// TAST (Let): __local_var_19_67 -> gopurs_runtime.Value
__local_var_19_67 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_64.UnsafePtr).V1, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_68 -> gopurs_runtime.Value
__local_var_20_68 := gopurs_runtime.Apply(f_prime_11, a_prime_19)
_ = __local_var_20_68
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_20_68, Get_Data_List_Lazy_Types_nil()})}
}))
}))
_ = __local_var_19_67
__t74 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_69 -> gopurs_runtime.Value
__local_var_21_69 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_65)
_ = __local_var_21_69
var __t73 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_21_69.Type == 9 && __local_var_21_69.IntVal == 218341868 && __local_var_21_69.UnsafePtr == nil) {
__t73 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_67))
goto end_branch_73
} else {

}
}
{
if (__local_var_21_69.Type == 9 && __local_var_21_69.IntVal == 218341868 && __local_var_21_69.UnsafePtr != nil) {
// TAST (Let): __local_var_22_70 -> gopurs_runtime.Value
__local_var_22_70 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_21_69.UnsafePtr).V1
_ = __local_var_22_70
__t73 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_21_69.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_71 -> gopurs_runtime.Value
__local_var_24_71 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_22_70)
_ = __local_var_24_71
var __t72 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_24_71.Type == 9 && __local_var_24_71.IntVal == 218341868 && __local_var_24_71.UnsafePtr == nil) {
__t72 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_67))
goto end_branch_72
} else {

}
}
{
if (__local_var_24_71.Type == 9 && __local_var_24_71.IntVal == 218341868 && __local_var_24_71.UnsafePtr != nil) {
__t72 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_24_71.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_24_71.UnsafePtr).V1, __local_var_19_67)}
goto end_branch_72
} else {

}
}
{
__t72 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_72:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t72)}
}))}
goto end_branch_73
} else {

}
}
{
__t73 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_73:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t73)}
}))))
goto end_branch_74
} else {

}
}
{
__t74 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_74:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t74)}
}))
_ = __local_var_15_62
__t83 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_75 -> gopurs_runtime.Value
__local_var_17_75 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_60)
_ = __local_var_17_75
var __t82 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_75.Type == 9 && __local_var_17_75.IntVal == 218341868 && __local_var_17_75.UnsafePtr == nil) {
__t82 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_62))
goto end_branch_82
} else {

}
}
{
if (__local_var_17_75.Type == 9 && __local_var_17_75.IntVal == 218341868 && __local_var_17_75.UnsafePtr != nil) {
// TAST (Let): __local_var_18_76 -> gopurs_runtime.Value
__local_var_18_76 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_75.UnsafePtr).V1
_ = __local_var_18_76
__t82 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_75.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_77 -> gopurs_runtime.Value
__local_var_20_77 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_76)
_ = __local_var_20_77
var __t81 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_77.Type == 9 && __local_var_20_77.IntVal == 218341868 && __local_var_20_77.UnsafePtr == nil) {
__t81 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_62))
goto end_branch_81
} else {

}
}
{
if (__local_var_20_77.Type == 9 && __local_var_20_77.IntVal == 218341868 && __local_var_20_77.UnsafePtr != nil) {
// TAST (Let): __local_var_21_78 -> gopurs_runtime.Value
__local_var_21_78 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_77.UnsafePtr).V1
_ = __local_var_21_78
__t81 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_77.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_79 -> gopurs_runtime.Value
__local_var_23_79 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_21_78)
_ = __local_var_23_79
var __t80 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_23_79.Type == 9 && __local_var_23_79.IntVal == 218341868 && __local_var_23_79.UnsafePtr == nil) {
__t80 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_62))
goto end_branch_80
} else {

}
}
{
if (__local_var_23_79.Type == 9 && __local_var_23_79.IntVal == 218341868 && __local_var_23_79.UnsafePtr != nil) {
__t80 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_79.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_79.UnsafePtr).V1, __local_var_15_62)}
goto end_branch_80
} else {

}
}
{
__t80 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_80:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t80)}
}))}
goto end_branch_81
} else {

}
}
{
__t81 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_81:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t81)}
}))}
goto end_branch_82
} else {

}
}
{
__t82 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_82:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t82)}
}))))
goto end_branch_83
} else {

}
}
{
__t83 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_83:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t83)}
}))
})
_ = __local_var_11_58
// TAST (Let): __local_var_10_56 -> gopurs_runtime.Value
__local_var_10_56 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_84 -> gopurs_runtime.Value
__local_var_13_84 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_57)
_ = __local_var_13_84
var __t105 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_84.Type == 9 && __local_var_13_84.IntVal == 218341868 && __local_var_13_84.UnsafePtr == nil) {
__t105 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_105
} else {

}
}
{
if (__local_var_13_84.Type == 9 && __local_var_13_84.IntVal == 218341868 && __local_var_13_84.UnsafePtr != nil) {
// TAST (Let): __local_var_14_85 -> gopurs_runtime.Value
__local_var_14_85 := gopurs_runtime.Apply(__local_var_11_58, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_84.UnsafePtr).V0)
_ = __local_var_14_85
// TAST (Let): __local_var_15_87 -> gopurs_runtime.Value
__local_var_15_87 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_84.UnsafePtr).V1
_ = __local_var_15_87
// TAST (Let): __local_var_15_86 -> gopurs_runtime.Value
__local_var_15_86 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_88 -> gopurs_runtime.Value
__local_var_17_88 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_87)
_ = __local_var_17_88
var __t96 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_88.Type == 9 && __local_var_17_88.IntVal == 218341868 && __local_var_17_88.UnsafePtr == nil) {
__t96 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_96
} else {

}
}
{
if (__local_var_17_88.Type == 9 && __local_var_17_88.IntVal == 218341868 && __local_var_17_88.UnsafePtr != nil) {
// TAST (Let): __local_var_18_89 -> gopurs_runtime.Value
__local_var_18_89 := gopurs_runtime.Apply(__local_var_11_58, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_88.UnsafePtr).V0)
_ = __local_var_18_89
// TAST (Let): __local_var_19_90 -> gopurs_runtime.Value
__local_var_19_90 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_88.UnsafePtr).V1, __local_var_11_58)
_ = __local_var_19_90
__t96 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_91 -> gopurs_runtime.Value
__local_var_21_91 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_89)
_ = __local_var_21_91
var __t95 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_21_91.Type == 9 && __local_var_21_91.IntVal == 218341868 && __local_var_21_91.UnsafePtr == nil) {
__t95 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_90))
goto end_branch_95
} else {

}
}
{
if (__local_var_21_91.Type == 9 && __local_var_21_91.IntVal == 218341868 && __local_var_21_91.UnsafePtr != nil) {
// TAST (Let): __local_var_22_92 -> gopurs_runtime.Value
__local_var_22_92 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_21_91.UnsafePtr).V1
_ = __local_var_22_92
__t95 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_21_91.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_93 -> gopurs_runtime.Value
__local_var_24_93 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_22_92)
_ = __local_var_24_93
var __t94 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_24_93.Type == 9 && __local_var_24_93.IntVal == 218341868 && __local_var_24_93.UnsafePtr == nil) {
__t94 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_90))
goto end_branch_94
} else {

}
}
{
if (__local_var_24_93.Type == 9 && __local_var_24_93.IntVal == 218341868 && __local_var_24_93.UnsafePtr != nil) {
__t94 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_24_93.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_24_93.UnsafePtr).V1, __local_var_19_90)}
goto end_branch_94
} else {

}
}
{
__t94 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_94:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t94)}
}))}
goto end_branch_95
} else {

}
}
{
__t95 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_95:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t95)}
}))))
goto end_branch_96
} else {

}
}
{
__t96 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_96:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t96)}
}))
_ = __local_var_15_86
__t105 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_97 -> gopurs_runtime.Value
__local_var_17_97 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_85)
_ = __local_var_17_97
var __t104 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_97.Type == 9 && __local_var_17_97.IntVal == 218341868 && __local_var_17_97.UnsafePtr == nil) {
__t104 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_86))
goto end_branch_104
} else {

}
}
{
if (__local_var_17_97.Type == 9 && __local_var_17_97.IntVal == 218341868 && __local_var_17_97.UnsafePtr != nil) {
// TAST (Let): __local_var_18_98 -> gopurs_runtime.Value
__local_var_18_98 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_97.UnsafePtr).V1
_ = __local_var_18_98
__t104 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_97.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_99 -> gopurs_runtime.Value
__local_var_20_99 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_98)
_ = __local_var_20_99
var __t103 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_99.Type == 9 && __local_var_20_99.IntVal == 218341868 && __local_var_20_99.UnsafePtr == nil) {
__t103 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_86))
goto end_branch_103
} else {

}
}
{
if (__local_var_20_99.Type == 9 && __local_var_20_99.IntVal == 218341868 && __local_var_20_99.UnsafePtr != nil) {
// TAST (Let): __local_var_21_100 -> gopurs_runtime.Value
__local_var_21_100 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_99.UnsafePtr).V1
_ = __local_var_21_100
__t103 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_99.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_101 -> gopurs_runtime.Value
__local_var_23_101 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_21_100)
_ = __local_var_23_101
var __t102 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_23_101.Type == 9 && __local_var_23_101.IntVal == 218341868 && __local_var_23_101.UnsafePtr == nil) {
__t102 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_86))
goto end_branch_102
} else {

}
}
{
if (__local_var_23_101.Type == 9 && __local_var_23_101.IntVal == 218341868 && __local_var_23_101.UnsafePtr != nil) {
__t102 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_101.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_23_101.UnsafePtr).V1, __local_var_15_86)}
goto end_branch_102
} else {

}
}
{
__t102 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_102:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t102)}
}))}
goto end_branch_103
} else {

}
}
{
__t103 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_103:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t103)}
}))}
goto end_branch_104
} else {

}
}
{
__t104 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_104:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t104)}
}))))
goto end_branch_105
} else {

}
}
{
__t105 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_105:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t105)}
}))
_ = __local_var_10_56
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_106 -> gopurs_runtime.Value
__local_var_12_106 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_6)
_ = __local_var_12_106
var __t107 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_12_106.Type == 9 && __local_var_12_106.IntVal == 218341868 && __local_var_12_106.UnsafePtr == nil) {
__t107 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_56))
goto end_branch_107
} else {

}
}
{
if (__local_var_12_106.Type == 9 && __local_var_12_106.IntVal == 218341868 && __local_var_12_106.UnsafePtr != nil) {
__t107 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_106.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_12_106.UnsafePtr).V1, __local_var_10_56)}
goto end_branch_107
} else {

}
}
{
__t107 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_107:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t107)}
}))})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_applyNonEmptyList__4077822201
}

var cache_Data_List_Lazy_Types_bindList__469219920 gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList__469219920 sync.Once
func Get_Data_List_Lazy_Types_bindList__469219920() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList__469219920.Do(func() {
		cache_Data_List_Lazy_Types_bindList__469219920 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyList()))}
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t15 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t15 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_15
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0)
_ = __local_var_4_1
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3)
_ = __local_var_7_4
var __t12 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_7_4.Type == 9 && __local_var_7_4.IntVal == 218341868 && __local_var_7_4.UnsafePtr == nil) {
__t12 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_12
} else {

}
}
{
if (__local_var_7_4.Type == 9 && __local_var_7_4.IntVal == 218341868 && __local_var_7_4.UnsafePtr != nil) {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_4.UnsafePtr).V0)
_ = __local_var_8_5
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_4.UnsafePtr).V1, f_1)
_ = __local_var_9_6
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_7 -> gopurs_runtime.Value
__local_var_11_7 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_5)
_ = __local_var_11_7
var __t11 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_7.Type == 9 && __local_var_11_7.IntVal == 218341868 && __local_var_11_7.UnsafePtr == nil) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_6))
goto end_branch_11
} else {

}
}
{
if (__local_var_11_7.Type == 9 && __local_var_11_7.IntVal == 218341868 && __local_var_11_7.UnsafePtr != nil) {
// TAST (Let): __local_var_12_8 -> gopurs_runtime.Value
__local_var_12_8 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_7.UnsafePtr).V1
_ = __local_var_12_8
__t11 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_7.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_9 -> gopurs_runtime.Value
__local_var_14_9 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_8)
_ = __local_var_14_9
var __t10 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_14_9.Type == 9 && __local_var_14_9.IntVal == 218341868 && __local_var_14_9.UnsafePtr == nil) {
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_6))
goto end_branch_10
} else {

}
}
{
if (__local_var_14_9.Type == 9 && __local_var_14_9.IntVal == 218341868 && __local_var_14_9.UnsafePtr != nil) {
__t10 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_14_9.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_14_9.UnsafePtr).V1, __local_var_9_6)}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t10)}
}))}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t11)}
}))))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t12)}
}))
_ = __local_var_5_2
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_13 -> gopurs_runtime.Value
__local_var_7_13 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_7_13
var __t14 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_7_13.Type == 9 && __local_var_7_13.IntVal == 218341868 && __local_var_7_13.UnsafePtr == nil) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2))
goto end_branch_14
} else {

}
}
{
if (__local_var_7_13.Type == 9 && __local_var_7_13.IntVal == 218341868 && __local_var_7_13.UnsafePtr != nil) {
__t14 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_13.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_13.UnsafePtr).V1, __local_var_5_2)}
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t14)}
}))))
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t15)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_bindList__469219920
}

var cache_Data_List_Lazy_Types_bindList__1050117088 gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList__1050117088 sync.Once
func Get_Data_List_Lazy_Types_bindList__1050117088() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList__1050117088.Do(func() {
		cache_Data_List_Lazy_Types_bindList__1050117088 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyList()))}
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t15 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t15 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_15
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0))})
_ = __local_var_4_1
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3)
_ = __local_var_7_4
var __t12 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_7_4.Type == 9 && __local_var_7_4.IntVal == 218341868 && __local_var_7_4.UnsafePtr == nil) {
__t12 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_12
} else {

}
}
{
if (__local_var_7_4.Type == 9 && __local_var_7_4.IntVal == 218341868 && __local_var_7_4.UnsafePtr != nil) {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_4.UnsafePtr).V0)
_ = __local_var_8_5
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_4.UnsafePtr).V1, f_1)
_ = __local_var_9_6
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_7 -> gopurs_runtime.Value
__local_var_11_7 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_5)
_ = __local_var_11_7
var __t11 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_11_7.Type == 9 && __local_var_11_7.IntVal == 218341868 && __local_var_11_7.UnsafePtr == nil) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_6))
goto end_branch_11
} else {

}
}
{
if (__local_var_11_7.Type == 9 && __local_var_11_7.IntVal == 218341868 && __local_var_11_7.UnsafePtr != nil) {
// TAST (Let): __local_var_12_8 -> gopurs_runtime.Value
__local_var_12_8 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_7.UnsafePtr).V1
_ = __local_var_12_8
__t11 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_11_7.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_9 -> gopurs_runtime.Value
__local_var_14_9 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_8)
_ = __local_var_14_9
var __t10 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_14_9.Type == 9 && __local_var_14_9.IntVal == 218341868 && __local_var_14_9.UnsafePtr == nil) {
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_6))
goto end_branch_10
} else {

}
}
{
if (__local_var_14_9.Type == 9 && __local_var_14_9.IntVal == 218341868 && __local_var_14_9.UnsafePtr != nil) {
__t10 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_14_9.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_14_9.UnsafePtr).V1, __local_var_9_6)}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t10)}
}))}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t11)}
}))))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t12)}
}))
_ = __local_var_5_2
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_13 -> gopurs_runtime.Value
__local_var_7_13 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_7_13
var __t14 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_7_13.Type == 9 && __local_var_7_13.IntVal == 218341868 && __local_var_7_13.UnsafePtr == nil) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2))
goto end_branch_14
} else {

}
}
{
if (__local_var_7_13.Type == 9 && __local_var_7_13.IntVal == 218341868 && __local_var_7_13.UnsafePtr != nil) {
__t14 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_13.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_13.UnsafePtr).V1, __local_var_5_2)}
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t14)}
}))))
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t15)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_bindList__1050117088
}

var cache_Data_List_Lazy_Types_bindNonEmptyList__3420378873 gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindNonEmptyList__3420378873 sync.Once
func Get_Data_List_Lazy_Types_bindNonEmptyList__3420378873() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindNonEmptyList__3420378873.Do(func() {
		cache_Data_List_Lazy_Types_bindNonEmptyList__3420378873 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Lazy_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): v2_4_2 -> *Constructor_Data_NonEmpty_NonEmpty
v2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V0)))
_ = v2_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (v2_4_2).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (v2_4_2).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1)
_ = __local_var_9_6
var __t14 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_9_6.Type == 9 && __local_var_9_6.IntVal == 218341868 && __local_var_9_6.UnsafePtr == nil) {
__t14 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_14
} else {

}
}
{
if (__local_var_9_6.Type == 9 && __local_var_9_6.IntVal == 218341868 && __local_var_9_6.UnsafePtr != nil) {
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := Call_Data_List_Lazy_Types_toList(gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_9_6.UnsafePtr).V0))
_ = __local_var_10_7
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_9_6.UnsafePtr).V1, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_toList(gopurs_runtime.Apply(f_1, x_11))
}))
_ = __local_var_11_8
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_9 -> gopurs_runtime.Value
__local_var_13_9 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_7)
_ = __local_var_13_9
var __t13 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_9.Type == 9 && __local_var_13_9.IntVal == 218341868 && __local_var_13_9.UnsafePtr == nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_13
} else {

}
}
{
if (__local_var_13_9.Type == 9 && __local_var_13_9.IntVal == 218341868 && __local_var_13_9.UnsafePtr != nil) {
// TAST (Let): __local_var_14_10 -> gopurs_runtime.Value
__local_var_14_10 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_9.UnsafePtr).V1
_ = __local_var_14_10
__t13 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_9.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_11 -> gopurs_runtime.Value
__local_var_16_11 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_10)
_ = __local_var_16_11
var __t12 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_11.Type == 9 && __local_var_16_11.IntVal == 218341868 && __local_var_16_11.UnsafePtr == nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_12
} else {

}
}
{
if (__local_var_16_11.Type == 9 && __local_var_16_11.IntVal == 218341868 && __local_var_16_11.UnsafePtr != nil) {
__t12 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_11.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_11.UnsafePtr).V1, __local_var_11_8)}
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t12)}
}))}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t13)}
}))))
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t14)}
}))
_ = __local_var_8_5
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_5_3, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_15 -> gopurs_runtime.Value
__local_var_10_15 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4)
_ = __local_var_10_15
var __t16 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_10_15.Type == 9 && __local_var_10_15.IntVal == 218341868 && __local_var_10_15.UnsafePtr == nil) {
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_5))
goto end_branch_16
} else {

}
}
{
if (__local_var_10_15.Type == 9 && __local_var_10_15.IntVal == 218341868 && __local_var_10_15.UnsafePtr != nil) {
__t16 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_10_15.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_10_15.UnsafePtr).V1, __local_var_8_5)}
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t16)}
}))})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_bindNonEmptyList__3420378873
}

var cache_Data_List_Lazy_Types_cons__716923058 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__716923058 sync.Once
func Get_Data_List_Lazy_Types_cons__716923058() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__716923058.Do(func() {
		cache_Data_List_Lazy_Types_cons__716923058 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__716923058(x_0_box, xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__716923058
}

var cache_Data_List_Lazy_Types_cons__720046150 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__720046150 sync.Once
func Get_Data_List_Lazy_Types_cons__720046150() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__720046150.Do(func() {
		cache_Data_List_Lazy_Types_cons__720046150 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__720046150(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__720046150
}

var cache_Data_List_Lazy_Types_cons__894912754 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__894912754 sync.Once
func Get_Data_List_Lazy_Types_cons__894912754() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__894912754.Do(func() {
		cache_Data_List_Lazy_Types_cons__894912754 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__894912754(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__894912754
}

var cache_Data_List_Lazy_Types_cons__376540526 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__376540526 sync.Once
func Get_Data_List_Lazy_Types_cons__376540526() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__376540526.Do(func() {
		cache_Data_List_Lazy_Types_cons__376540526 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__376540526(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__376540526
}

var cache_Data_List_Lazy_Types_cons__673400617 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__673400617 sync.Once
func Get_Data_List_Lazy_Types_cons__673400617() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__673400617.Do(func() {
		cache_Data_List_Lazy_Types_cons__673400617 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__673400617(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__673400617
}

var cache_Data_List_Lazy_Types_cons__2305074921 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__2305074921 sync.Once
func Get_Data_List_Lazy_Types_cons__2305074921() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__2305074921.Do(func() {
		cache_Data_List_Lazy_Types_cons__2305074921 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__2305074921(x_0_box, xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__2305074921
}

var cache_Data_List_Lazy_Types_cons__891310957 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__891310957 sync.Once
func Get_Data_List_Lazy_Types_cons__891310957() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__891310957.Do(func() {
		cache_Data_List_Lazy_Types_cons__891310957 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__891310957(x_0_box, xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__891310957
}

var cache_Data_List_Lazy_Types_cons__1901546616 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__1901546616 sync.Once
func Get_Data_List_Lazy_Types_cons__1901546616() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__1901546616.Do(func() {
		cache_Data_List_Lazy_Types_cons__1901546616 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__1901546616(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__1901546616
}

var cache_Data_List_Lazy_Types_cons__3391588829 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__3391588829 sync.Once
func Get_Data_List_Lazy_Types_cons__3391588829() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__3391588829.Do(func() {
		cache_Data_List_Lazy_Types_cons__3391588829 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__3391588829(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__3391588829
}

var cache_Data_List_Lazy_Types_cons__2134285409 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__2134285409 sync.Once
func Get_Data_List_Lazy_Types_cons__2134285409() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__2134285409.Do(func() {
		cache_Data_List_Lazy_Types_cons__2134285409 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__2134285409(x_0_box, xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__2134285409
}

var cache_Data_List_Lazy_Types_eq1List__394800310 gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1List__394800310 sync.Once
func Get_Data_List_Lazy_Types_eq1List__394800310() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1List__394800310.Do(func() {
		cache_Data_List_Lazy_Types_eq1List__394800310 = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_90 gopurs_runtime.Value
go__go_3_0_90 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_90:
for {
if false { continue go__go_3_0_90 }
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
continue go__go_3_0_90
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_90, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_eq1List__394800310
}

var cache_Data_List_Lazy_Types_eq1NonEmptyList__1055554591 gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1NonEmptyList__1055554591 sync.Once
func Get_Data_List_Lazy_Types_eq1NonEmptyList__1055554591() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1NonEmptyList__1055554591.Do(func() {
		cache_Data_List_Lazy_Types_eq1NonEmptyList__1055554591 = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqNonEmpty1_1_0 -> *Constructor_Data_Eq_Eq
eqNonEmpty1_1_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_4 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_1_91 gopurs_runtime.Value
go__go_3_1_91 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_1_91:
for {
if false { continue go__go_3_1_91 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 bool
{
if (v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr == nil) {
var __t2 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)))}
continue go__go_3_1_91
__t3 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
}
}()
})
})
__t_and_4 = (gopurs_runtime.Apply2(go__go_3_1_91, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_4)
})
})}
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqNonEmpty1_1_0.V0), gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_3)).IntVal) != (0))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_eq1NonEmptyList__1055554591
}

var cache_Data_List_Lazy_Types_extendNonEmptyList__848880921 gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendNonEmptyList__848880921 sync.Once
func Get_Data_List_Lazy_Types_extendNonEmptyList__848880921() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendNonEmptyList__848880921.Do(func() {
		cache_Data_List_Lazy_Types_extendNonEmptyList__848880921 = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1).UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_92 gopurs_runtime.Value
go__go_4_1_92 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_1_92:
for {
if false { continue go__go_4_1_92 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_2 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_2
var __t5 gopurs_runtime.Value
{
if (v_7_2 == nil) {
__t5 = b_5
goto end_branch_5
} else {

}
}
{
if (v_7_2 != nil) {
// TAST (Let): __local_var_8_3 -> gopurs_runtime.Value
__local_var_8_3 := gopurs_runtime.RecordGet(b_5, "acc")
_ = __local_var_8_3
// TAST (Let): __local_var_9_4 -> gopurs_runtime.Value
__local_var_9_4 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v_7_2).V0, __local_var_8_3})}
})))
_ = __local_var_9_4
b_5_loop = gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v_7_2).V0, __local_var_8_3})}
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_4, gopurs_runtime.RecordGet(b_5, "val")})}
})))
xs_6_loop = (v_7_2).V1
continue go__go_4_1_92
__t5 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_5_6_93 gopurs_runtime.Value
go__go_5_6_93 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_6_93:
for {
if false { continue go__go_5_6_93 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_7
var __t9 gopurs_runtime.Value
{
if (v_8_7 == nil) {
__t9 = b_6
goto end_branch_9
} else {

}
}
{
if (v_8_7 != nil) {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := (v_8_7).V0
_ = __local_var_9_8
b_6_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_8, b_6})}
}))
xs_7_loop = (v_8_7).V1
continue go__go_5_6_93
__t9 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(go__go_4_1_92, gopurs_runtime.RecordDict2("acc", "val", Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_5_6_93, Get_Data_List_Lazy_Types_nil(), __local_var_2_0)), "val")})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_extendNonEmptyList__848880921
}

var cache_Data_List_Lazy_Types_foldableList__4097915271 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__4097915271 sync.Once
func Get_Data_List_Lazy_Types_foldableList__4097915271() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__4097915271.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__4097915271 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_94 gopurs_runtime.Value
go__go_4_2_94 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_2_94:
for {
if false { continue go__go_4_2_94 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3 == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_7_3 != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_5, gopurs_runtime.Apply(f_3, (v_7_3).V0))
xs_6_loop = (v_7_3).V1
continue go__go_4_2_94
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
return gopurs_runtime.Apply(go__go_4_2_94, mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_5_95 gopurs_runtime.Value
go__go_1_5_95 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_5_95:
for {
if false { continue go__go_1_5_95 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_6
var __t7 gopurs_runtime.Value
{
if (v_4_6 == nil) {
__t7 = b_2
goto end_branch_7
} else {

}
}
{
if (v_4_6 != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (v_4_6).V0)
xs_3_loop = (v_4_6).V1
continue go__go_1_5_95
__t7 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_1_5_95
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_96 gopurs_runtime.Value
go__go_3_8_96 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_8_96:
for {
if false { continue go__go_3_8_96 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_9 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
b_4_loop = gopurs_runtime.Apply2(op_0, (v_6_9).V0, b_4)
xs_5_loop = (v_6_9).V1
continue go__go_3_8_96
__t10 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_4_11_97 gopurs_runtime.Value
go__go_4_11_97 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_11_97:
for {
if false { continue go__go_4_11_97 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
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
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := (v_7_12).V0
_ = __local_var_8_13
b_5_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_13, b_5})}
}))
xs_6_loop = (v_7_12).V1
continue go__go_4_11_97
__t14 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_8_96, z_1, gopurs_runtime.Apply2(go__go_4_11_97, Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableList__4097915271
}

var cache_Data_List_Lazy_Types_foldableList__331628915 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__331628915 sync.Once
func Get_Data_List_Lazy_Types_foldableList__331628915() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__331628915.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__331628915 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_98 gopurs_runtime.Value
go__go_4_2_98 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_2_98:
for {
if false { continue go__go_4_2_98 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3 == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_7_3 != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_5, gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_7_3).V0))}))
xs_6_loop = (v_7_3).V1
continue go__go_4_2_98
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
return gopurs_runtime.Apply(go__go_4_2_98, mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_5_99 gopurs_runtime.Value
go__go_1_5_99 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_5_99:
for {
if false { continue go__go_1_5_99 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_6
var __t7 gopurs_runtime.Value
{
if (v_4_6 == nil) {
__t7 = b_2
goto end_branch_7
} else {

}
}
{
if (v_4_6 != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_4_6).V0))})
xs_3_loop = (v_4_6).V1
continue go__go_1_5_99
__t7 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_1_5_99
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_100 gopurs_runtime.Value
go__go_3_8_100 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_8_100:
for {
if false { continue go__go_3_8_100 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_9 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
b_4_loop = gopurs_runtime.Apply2(op_0, (v_6_9).V0, b_4)
xs_5_loop = (v_6_9).V1
continue go__go_3_8_100
__t10 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_4_11_101 gopurs_runtime.Value
go__go_4_11_101 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_11_101:
for {
if false { continue go__go_4_11_101 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
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
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := (v_7_12).V0
_ = __local_var_8_13
b_5_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_13, b_5})}
}))
xs_6_loop = (v_7_12).V1
continue go__go_4_11_101
__t14 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_8_100, z_1, gopurs_runtime.Apply2(go__go_4_11_101, Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableList__331628915
}

var cache_Data_List_Lazy_Types_foldableList__21955931 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__21955931 sync.Once
func Get_Data_List_Lazy_Types_foldableList__21955931() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__21955931.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__21955931 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_102 gopurs_runtime.Value
go__go_4_2_102 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_2_102:
for {
if false { continue go__go_4_2_102 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3 == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_7_3 != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_5, gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_7_3).V0))}))
xs_6_loop = (v_7_3).V1
continue go__go_4_2_102
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
return gopurs_runtime.Apply(go__go_4_2_102, mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_5_103 gopurs_runtime.Value
go__go_1_5_103 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_5_103:
for {
if false { continue go__go_1_5_103 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_6
var __t7 gopurs_runtime.Value
{
if (v_4_6 == nil) {
__t7 = b_2
goto end_branch_7
} else {

}
}
{
if (v_4_6 != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_4_6).V0))})
xs_3_loop = (v_4_6).V1
continue go__go_1_5_103
__t7 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_1_5_103
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_104 gopurs_runtime.Value
go__go_3_8_104 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_8_104:
for {
if false { continue go__go_3_8_104 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_9 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
b_4_loop = gopurs_runtime.Apply2(op_0, (v_6_9).V0, b_4)
xs_5_loop = (v_6_9).V1
continue go__go_3_8_104
__t10 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_4_11_105 gopurs_runtime.Value
go__go_4_11_105 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_11_105:
for {
if false { continue go__go_4_11_105 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
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
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := (v_7_12).V0
_ = __local_var_8_13
b_5_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_13, b_5})}
}))
xs_6_loop = (v_7_12).V1
continue go__go_4_11_105
__t14 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_8_104, z_1, gopurs_runtime.Apply2(go__go_4_11_105, Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableList__21955931
}

var cache_Data_List_Lazy_Types_foldableList__3094856796 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__3094856796 sync.Once
func Get_Data_List_Lazy_Types_foldableList__3094856796() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__3094856796.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__3094856796 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_106 gopurs_runtime.Value
go__go_4_2_106 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_2_106:
for {
if false { continue go__go_4_2_106 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3 == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_7_3 != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_5, gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_7_3).V0))}))
xs_6_loop = (v_7_3).V1
continue go__go_4_2_106
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
return gopurs_runtime.Apply(go__go_4_2_106, mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_5_107 gopurs_runtime.Value
go__go_1_5_107 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_5_107:
for {
if false { continue go__go_1_5_107 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_6
var __t7 gopurs_runtime.Value
{
if (v_4_6 == nil) {
__t7 = b_2
goto end_branch_7
} else {

}
}
{
if (v_4_6 != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_4_6).V0))})
xs_3_loop = (v_4_6).V1
continue go__go_1_5_107
__t7 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_1_5_107
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_108 gopurs_runtime.Value
go__go_3_8_108 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_8_108:
for {
if false { continue go__go_3_8_108 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_9 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
b_4_loop = gopurs_runtime.Apply2(op_0, (v_6_9).V0, b_4)
xs_5_loop = (v_6_9).V1
continue go__go_3_8_108
__t10 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_4_11_109 gopurs_runtime.Value
go__go_4_11_109 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_11_109:
for {
if false { continue go__go_4_11_109 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
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
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := (v_7_12).V0
_ = __local_var_8_13
b_5_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_13, b_5})}
}))
xs_6_loop = (v_7_12).V1
continue go__go_4_11_109
__t14 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_8_108, z_1, gopurs_runtime.Apply2(go__go_4_11_109, Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableList__3094856796
}

var cache_Data_List_Lazy_Types_foldableNonEmptyList__1614682446 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmptyList__1614682446 sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmptyList__1614682446() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmptyList__1614682446.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmptyList__1614682446 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V1), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V2), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableNonEmptyList__1614682446
}

var cache_Data_List_Lazy_Types_foldableWithIndexList__662860203 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexList__662860203 sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexList__662860203() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexList__662860203.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexList__662860203 = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_110 gopurs_runtime.Value
go__go_4_3_110 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_3_110:
for {
if false { continue go__go_4_3_110 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_4 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_4
var __t5 gopurs_runtime.Value
{
if (v_7_4 == nil) {
__t5 = b_5
goto end_branch_5
} else {

}
}
{
if (v_7_4 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1, gopurs_runtime.Apply2(f_3, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal), (v_7_4).V0))})}
xs_6_loop = (v_7_4).V1
continue go__go_4_3_110
__t5 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(go__go_4_3_110, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), mempty_2_1})})
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(__local_var_4_2, x_5).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_7_111 gopurs_runtime.Value
go__go_2_7_111 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_7_111:
for {
if false { continue go__go_2_7_111 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_8 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_8
var __t9 gopurs_runtime.Value
{
if (v_5_8 == nil) {
__t9 = b_3
goto end_branch_9
} else {

}
}
{
if (v_5_8 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V1, (v_5_8).V0)})}
xs_4_loop = (v_5_8).V1
continue go__go_2_7_111
__t9 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply(go__go_2_7_111, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_6
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(__local_var_2_6, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_11_112 gopurs_runtime.Value
go__go_3_11_112 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_11_112:
for {
if false { continue go__go_3_11_112 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_12
var __t15 gopurs_runtime.Value
{
if (v_6_12 == nil) {
__t15 = b_4
goto end_branch_15
} else {

}
}
{
if (v_6_12 != nil) {
// TAST (Let): __local_var_7_13 -> gopurs_runtime.Value
__local_var_7_13 := (*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1
_ = __local_var_7_13
// TAST (Let): __local_var_8_14 -> gopurs_runtime.Value
__local_var_8_14 := (v_6_12).V0
_ = __local_var_8_14
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_14, __local_var_7_13})}
}))})}
xs_5_loop = (v_6_12).V1
continue go__go_3_11_112
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
// TAST (Let): v_3_10 -> *Constructor_Data_Tuple_Tuple
v_3_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_11_112, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_2))
_ = v_3_10
var go__go_4_16_113 gopurs_runtime.Value
go__go_4_16_113 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_16_113:
for {
if false { continue go__go_4_16_113 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_17 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_17
var __t18 gopurs_runtime.Value
{
if (v_7_17 == nil) {
__t18 = b_5
goto end_branch_18
} else {

}
}
{
if (v_7_17 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), (v_7_17).V0, (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1)})}
xs_6_loop = (v_7_17).V1
continue go__go_4_16_113
__t18 = gopurs_runtime.Value{}
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_16_113, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_10).V0.IntVal), b_1})}, (v_3_10).V1).UnsafePtr).V1
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexList__662860203
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__2582116578 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__2582116578 sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__2582116578() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__2582116578.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__2582116578 = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmptyList()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2))
_ = __local_var_3_0
// TAST (Let): Semigroup0_4_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_1
var go__go_5_2_114 gopurs_runtime.Value
go__go_5_2_114 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_2_114:
for {
if false { continue go__go_5_2_114 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
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
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_1.V0), (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal)), (v_8_3).V0))})}
xs_7_loop = (v_8_3).V1
continue go__go_5_2_114
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_1, gopurs_runtime.Int(0), (__local_var_3_0).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_2_114, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})}, (__local_var_3_0).V1).UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2))
_ = __local_var_3_5
var go__go_4_6_115 gopurs_runtime.Value
go__go_4_6_115 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_6_115:
for {
if false { continue go__go_4_6_115 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_7
var __t8 gopurs_runtime.Value
{
if (v_7_7 == nil) {
__t8 = b_5
goto end_branch_8
} else {

}
}
{
if (v_7_7 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal)), (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1, (v_7_7).V0)})}
xs_6_loop = (v_7_7).V1
continue go__go_4_6_115
__t8 = gopurs_runtime.Value{}
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
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_6_115, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(0), b_1, (__local_var_3_5).V0)})}, (__local_var_3_5).V1).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_9 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2))
_ = __local_var_3_9
var go__go_4_11_116 gopurs_runtime.Value
go__go_4_11_116 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_11_116:
for {
if false { continue go__go_4_11_116 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_12 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_12
var __t15 gopurs_runtime.Value
{
if (v_7_12 == nil) {
__t15 = b_5
goto end_branch_15
} else {

}
}
{
if (v_7_12 != nil) {
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1
_ = __local_var_8_13
// TAST (Let): __local_var_9_14 -> gopurs_runtime.Value
__local_var_9_14 := (v_7_12).V0
_ = __local_var_9_14
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_14, __local_var_8_13})}
}))})}
xs_6_loop = (v_7_12).V1
continue go__go_4_11_116
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
// TAST (Let): v_4_10 -> *Constructor_Data_Tuple_Tuple
v_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_4_11_116, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (__local_var_3_9).V1))
_ = v_4_10
var go__go_5_16_117 gopurs_runtime.Value
go__go_5_16_117 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_16_117:
for {
if false { continue go__go_5_16_117 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_17 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_17
var __t18 gopurs_runtime.Value
{
if (v_8_17 == nil) {
__t18 = b_6
goto end_branch_18
} else {

}
}
{
if (v_8_17 != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1))), (v_8_17).V0, (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1)})}
xs_7_loop = (v_8_17).V1
continue go__go_5_16_117
__t18 = gopurs_runtime.Value{}
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}
}()
})
})
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(0), (__local_var_3_9).V0, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_16_117, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_4_10).V0.IntVal), b_1})}, (v_4_10).V1).UnsafePtr).V1)
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__2582116578
}

var cache_Data_List_Lazy_Types_functorList__699353223 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList__699353223 sync.Once
func Get_Data_List_Lazy_Types_functorList__699353223() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList__699353223.Do(func() {
		cache_Data_List_Lazy_Types_functorList__699353223 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_3_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1
_ = __local_var_4_1
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorList__699353223
}

var cache_Data_List_Lazy_Types_functorList__1388140403 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList__1388140403 sync.Once
func Get_Data_List_Lazy_Types_functorList__1388140403() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList__1388140403.Do(func() {
		cache_Data_List_Lazy_Types_functorList__1388140403 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_3_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1
_ = __local_var_4_1
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0))}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorList__1388140403
}

var cache_Data_List_Lazy_Types_functorList__1718231415 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList__1718231415 sync.Once
func Get_Data_List_Lazy_Types_functorList__1718231415() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList__1718231415.Do(func() {
		cache_Data_List_Lazy_Types_functorList__1718231415 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_3_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1
_ = __local_var_4_1
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0))}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorList__1718231415
}

var cache_Data_List_Lazy_Types_functorNonEmptyList__2698868942 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmptyList__2698868942 sync.Once
func Get_Data_List_Lazy_Types_functorNonEmptyList__2698868942() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmptyList__2698868942.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmptyList__2698868942 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmpty()).V0), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorNonEmptyList__2698868942
}

var cache_Data_List_Lazy_Types_functorWithIndexList__1530074091 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexList__1530074091 sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexList__1530074091() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexList__1530074091.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexList__1530074091 = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_118 gopurs_runtime.Value
go__go_2_1_118 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_1_118:
for {
if false { continue go__go_2_1_118 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_2 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
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
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := (*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V1
_ = __local_var_6_3
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (v_5_2).V0
_ = __local_var_7_4
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_7_4, __local_var_6_3})}
}))})}
xs_4_loop = (v_5_2).V1
continue go__go_2_1_118
__t5 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_2_0 -> *Constructor_Data_Tuple_Tuple
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_2_1_118, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_1))
_ = v_2_0
var go__go_3_6_119 gopurs_runtime.Value
go__go_3_6_119 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_6_119:
for {
if false { continue go__go_3_6_119 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := (*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1
_ = __local_var_7_8
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1)), (v_6_7).V0)
_ = __local_var_8_9
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_8_9, __local_var_7_8})}
}))})}
xs_5_loop = (v_6_7).V1
continue go__go_3_6_119
__t10 = gopurs_runtime.Value{}
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
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_3_6_119, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_2_0).V0.IntVal), Get_Data_List_Lazy_Types_nil()})}, (v_2_0).V1).UnsafePtr).V1
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorWithIndexList__1530074091
}

var cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList__1050619298 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexNonEmptyList__1050619298 sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList__1050619298() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexNonEmptyList__1050619298.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList__1050619298 = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
_ = __local_var_3_0
var go__go_4_2_120 gopurs_runtime.Value
go__go_4_2_120 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var xs_6_loop gopurs_runtime.Value = xs_6_loop_val
go__go_4_2_120:
for {
if false { continue go__go_4_2_120 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var xs_6 gopurs_runtime.Value = xs_6_loop
_ = xs_6
// TAST (Let): v_7_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_6))
_ = v_7_3
var __t6 gopurs_runtime.Value
{
if (v_7_3 == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_7_3 != nil) {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1
_ = __local_var_8_4
// TAST (Let): __local_var_9_5 -> gopurs_runtime.Value
__local_var_9_5 := (v_7_3).V0
_ = __local_var_9_5
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_9_5, __local_var_8_4})}
}))})}
xs_6_loop = (v_7_3).V1
continue go__go_4_2_120
__t6 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): v_4_1 -> *Constructor_Data_Tuple_Tuple
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_4_2_120, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, (__local_var_3_0).V1))
_ = v_4_1
var go__go_5_7_121 gopurs_runtime.Value
go__go_5_7_121 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_7_121:
for {
if false { continue go__go_5_7_121 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_8 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_8
var __t11 gopurs_runtime.Value
{
if (v_8_8 == nil) {
__t11 = b_6
goto end_branch_11
} else {

}
}
{
if (v_8_8 != nil) {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1
_ = __local_var_9_9
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1))), (v_8_8).V0)
_ = __local_var_10_10
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_10_10, __local_var_9_9})}
}))})}
xs_7_loop = (v_8_8).V1
continue go__go_5_7_121
__t11 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(0), (__local_var_3_0).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_7_121, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_4_1).V0.IntVal), Get_Data_List_Lazy_Types_nil()})}, (v_4_1).V1).UnsafePtr).V1})}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList__1050619298
}

var cache_Data_List_Lazy_Types_lazyList__601034736 gopurs_runtime.Value
var once_Data_List_Lazy_Types_lazyList__601034736 sync.Once
func Get_Data_List_Lazy_Types_lazyList__601034736() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_lazyList__601034736.Do(func() {
		cache_Data_List_Lazy_Types_lazyList__601034736 = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(&Constructor_Control_Lazy_Lazy{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, x_1))
}))
})})}
	})
	return cache_Data_List_Lazy_Types_lazyList__601034736
}

var cache_Data_List_Lazy_Types_lazyList__706698273 gopurs_runtime.Value
var once_Data_List_Lazy_Types_lazyList__706698273 sync.Once
func Get_Data_List_Lazy_Types_lazyList__706698273() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_lazyList__706698273.Do(func() {
		cache_Data_List_Lazy_Types_lazyList__706698273 = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(&Constructor_Control_Lazy_Lazy{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, x_1))
}))
})})}
	})
	return cache_Data_List_Lazy_Types_lazyList__706698273
}

var cache_Data_List_Lazy_Types_monadList__976767312 gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadList__976767312 sync.Once
func Get_Data_List_Lazy_Types_monadList__976767312() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadList__976767312.Do(func() {
		cache_Data_List_Lazy_Types_monadList__976767312 = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Lazy_Types_applicativeList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()))}
})})}
	})
	return cache_Data_List_Lazy_Types_monadList__976767312
}

var cache_Data_List_Lazy_Types_monoidList__1794469159 gopurs_runtime.Value
var once_Data_List_Lazy_Types_monoidList__1794469159 sync.Once
func Get_Data_List_Lazy_Types_monoidList__1794469159() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monoidList__1794469159.Do(func() {
		cache_Data_List_Lazy_Types_monoidList__1794469159 = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()))}
}), Get_Data_List_Lazy_Types_nil()})}
	})
	return cache_Data_List_Lazy_Types_monoidList__1794469159
}

var cache_Data_List_Lazy_Types_nil__1478684294 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__1478684294 sync.Once
func Get_Data_List_Lazy_Types_nil__1478684294() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__1478684294.Do(func() {
		cache_Data_List_Lazy_Types_nil__1478684294 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__1478684294
}

var cache_Data_List_Lazy_Types_nil__3988504114 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__3988504114 sync.Once
func Get_Data_List_Lazy_Types_nil__3988504114() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__3988504114.Do(func() {
		cache_Data_List_Lazy_Types_nil__3988504114 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__3988504114
}

var cache_Data_List_Lazy_Types_nil__1778182234 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__1778182234 sync.Once
func Get_Data_List_Lazy_Types_nil__1778182234() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__1778182234.Do(func() {
		cache_Data_List_Lazy_Types_nil__1778182234 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__1778182234
}

var cache_Data_List_Lazy_Types_nil__1439122877 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__1439122877 sync.Once
func Get_Data_List_Lazy_Types_nil__1439122877() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__1439122877.Do(func() {
		cache_Data_List_Lazy_Types_nil__1439122877 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__1439122877
}

var cache_Data_List_Lazy_Types_nil__2012296605 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__2012296605 sync.Once
func Get_Data_List_Lazy_Types_nil__2012296605() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__2012296605.Do(func() {
		cache_Data_List_Lazy_Types_nil__2012296605 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__2012296605
}

var cache_Data_List_Lazy_Types_nil__2014033708 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__2014033708 sync.Once
func Get_Data_List_Lazy_Types_nil__2014033708() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__2014033708.Do(func() {
		cache_Data_List_Lazy_Types_nil__2014033708 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__2014033708
}

var cache_Data_List_Lazy_Types_nil__2288399465 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__2288399465 sync.Once
func Get_Data_List_Lazy_Types_nil__2288399465() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__2288399465.Do(func() {
		cache_Data_List_Lazy_Types_nil__2288399465 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__2288399465
}

var cache_Data_List_Lazy_Types_nil__4122162182 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__4122162182 sync.Once
func Get_Data_List_Lazy_Types_nil__4122162182() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__4122162182.Do(func() {
		cache_Data_List_Lazy_Types_nil__4122162182 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons)(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__4122162182
}

var cache_Data_List_Lazy_Types_plusList__2873873584 gopurs_runtime.Value
var once_Data_List_Lazy_Types_plusList__2873873584 sync.Once
func Get_Data_List_Lazy_Types_plusList__2873873584() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_plusList__2873873584.Do(func() {
		cache_Data_List_Lazy_Types_plusList__2873873584 = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Data_List_Lazy_Types_altList()))}
}), Get_Data_List_Lazy_Types_nil()})}
	})
	return cache_Data_List_Lazy_Types_plusList__2873873584
}

var cache_Data_List_Lazy_Types_semigroupList__1199693447 gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList__1199693447 sync.Once
func Get_Data_List_Lazy_Types_semigroupList__1199693447() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList__1199693447.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList__1199693447 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1
_ = __local_var_4_1
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_3
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V1, ys_1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_semigroupList__1199693447
}

var cache_Data_List_Lazy_Types_semigroupList__3612943602 gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList__3612943602 sync.Once
func Get_Data_List_Lazy_Types_semigroupList__3612943602() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList__3612943602.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList__3612943602 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1
_ = __local_var_4_1
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_3
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V1, ys_1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_semigroupList__3612943602
}

var cache_Data_List_Lazy_Types_semigroupList__2598308723 gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList__2598308723 sync.Once
func Get_Data_List_Lazy_Types_semigroupList__2598308723() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList__2598308723.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList__2598308723 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1
_ = __local_var_4_1
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0))}, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
goto end_branch_3
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_2.UnsafePtr).V1, ys_1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))})})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
})
})})}
	})
	return cache_Data_List_Lazy_Types_semigroupList__2598308723
}

var cache_Data_List_Lazy_Types_step__3545407802 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__3545407802 sync.Once
func Get_Data_List_Lazy_Types_step__3545407802() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__3545407802.Do(func() {
		cache_Data_List_Lazy_Types_step__3545407802 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__3545407802(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__3545407802
}

var cache_Data_List_Lazy_Types_step__3687829882 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__3687829882 sync.Once
func Get_Data_List_Lazy_Types_step__3687829882() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__3687829882.Do(func() {
		cache_Data_List_Lazy_Types_step__3687829882 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__3687829882(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__3687829882
}

var cache_Data_List_Lazy_Types_step__4184651873 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__4184651873 sync.Once
func Get_Data_List_Lazy_Types_step__4184651873() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__4184651873.Do(func() {
		cache_Data_List_Lazy_Types_step__4184651873 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__4184651873(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__4184651873
}

var cache_Data_List_Lazy_Types_step__4057057377 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__4057057377 sync.Once
func Get_Data_List_Lazy_Types_step__4057057377() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__4057057377.Do(func() {
		cache_Data_List_Lazy_Types_step__4057057377 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__4057057377(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__4057057377
}

var cache_Data_List_Lazy_Types_step__2986341153 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__2986341153 sync.Once
func Get_Data_List_Lazy_Types_step__2986341153() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__2986341153.Do(func() {
		cache_Data_List_Lazy_Types_step__2986341153 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__2986341153(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__2986341153
}

var cache_Data_List_Lazy_Types_step__2999566881 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__2999566881 sync.Once
func Get_Data_List_Lazy_Types_step__2999566881() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__2999566881.Do(func() {
		cache_Data_List_Lazy_Types_step__2999566881 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__2999566881(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__2999566881
}

var cache_Data_List_Lazy_Types_step__2528948705 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__2528948705 sync.Once
func Get_Data_List_Lazy_Types_step__2528948705() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__2528948705.Do(func() {
		cache_Data_List_Lazy_Types_step__2528948705 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__2528948705(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__2528948705
}

var cache_Data_List_Lazy_Types_step__2322903873 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__2322903873 sync.Once
func Get_Data_List_Lazy_Types_step__2322903873() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__2322903873.Do(func() {
		cache_Data_List_Lazy_Types_step__2322903873 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__2322903873(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__2322903873
}

var cache_Data_List_Lazy_Types_step__2597188449 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__2597188449 sync.Once
func Get_Data_List_Lazy_Types_step__2597188449() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__2597188449.Do(func() {
		cache_Data_List_Lazy_Types_step__2597188449 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__2597188449(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__2597188449
}

var cache_Data_List_Lazy_Types_step__1903922273 gopurs_runtime.Value
var once_Data_List_Lazy_Types_step__1903922273 sync.Once
func Get_Data_List_Lazy_Types_step__1903922273() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step__1903922273.Do(func() {
		cache_Data_List_Lazy_Types_step__1903922273 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_step__1903922273(x_0_box))}
})
	})
	return cache_Data_List_Lazy_Types_step__1903922273
}

var cache_Data_List_Lazy_Types_toList__1017592434 gopurs_runtime.Value
var once_Data_List_Lazy_Types_toList__1017592434 sync.Once
func Get_Data_List_Lazy_Types_toList__1017592434() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_toList__1017592434.Do(func() {
		cache_Data_List_Lazy_Types_toList__1017592434 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_toList__1017592434(v_0_box)
})
	})
	return cache_Data_List_Lazy_Types_toList__1017592434
}

var cache_Data_List_Lazy_Types_toList__4101396777 gopurs_runtime.Value
var once_Data_List_Lazy_Types_toList__4101396777 sync.Once
func Get_Data_List_Lazy_Types_toList__4101396777() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_toList__4101396777.Do(func() {
		cache_Data_List_Lazy_Types_toList__4101396777 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_toList__4101396777(v_0_box)
})
	})
	return cache_Data_List_Lazy_Types_toList__4101396777
}

var cache_Data_List_Lazy_Types_traversableList__3068288903 gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableList__3068288903 sync.Once
func Get_Data_List_Lazy_Types_traversableList__3068288903() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableList__3068288903.Do(func() {
		cache_Data_List_Lazy_Types_traversableList__3068288903 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_3_2
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_3_122 gopurs_runtime.Value
go__go_5_3_122 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_3_122:
for {
if false { continue go__go_5_3_122 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_4 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_4
var __t5 gopurs_runtime.Value
{
if (v_8_4 == nil) {
__t5 = b_6
goto end_branch_5
} else {

}
}
{
if (v_8_4 != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), (v_8_4).V0), b_6)
xs_7_loop = (v_8_4).V1
continue go__go_5_3_122
__t5 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_6_6_123 gopurs_runtime.Value
go__go_6_6_123 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_6_123:
for {
if false { continue go__go_6_6_123 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_7
var __t9 gopurs_runtime.Value
{
if (v_9_7 == nil) {
__t9 = b_7
goto end_branch_9
} else {

}
}
{
if (v_9_7 != nil) {
// TAST (Let): __local_var_10_8 -> gopurs_runtime.Value
__local_var_10_8 := (v_9_7).V0
_ = __local_var_10_8
b_7_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_10_8, b_7})}
}))
xs_8_loop = (v_9_7).V1
continue go__go_6_6_123
__t9 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_5_3_122, __local_var_3_2, gopurs_runtime.Apply2(go__go_6_6_123, Get_Data_List_Lazy_Types_nil(), xs_4))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_10 -> *Constructor_Control_Apply_Apply
Apply0_1_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_10
// TAST (Let): Functor0_2_11 -> *Constructor_Data_Functor_Functor
Functor0_2_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_11
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_12 -> gopurs_runtime.Value
__local_var_4_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_4_12
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_13_124 gopurs_runtime.Value
go__go_6_13_124 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_13_124:
for {
if false { continue go__go_6_13_124 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_14 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_14
var __t15 gopurs_runtime.Value
{
if (v_9_14 == nil) {
__t15 = b_7
goto end_branch_15
} else {

}
}
{
if (v_9_14 != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_10.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_11.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, (v_9_14).V0)), b_7)
xs_8_loop = (v_9_14).V1
continue go__go_6_13_124
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
var go__go_7_16_125 gopurs_runtime.Value
go__go_7_16_125 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_16_125:
for {
if false { continue go__go_7_16_125 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_17 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_17
var __t19 gopurs_runtime.Value
{
if (v_10_17 == nil) {
__t19 = b_8
goto end_branch_19
} else {

}
}
{
if (v_10_17 != nil) {
// TAST (Let): __local_var_11_18 -> gopurs_runtime.Value
__local_var_11_18 := (v_10_17).V0
_ = __local_var_11_18
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_11_18, b_8})}
}))
xs_9_loop = (v_10_17).V1
continue go__go_7_16_125
__t19 = gopurs_runtime.Value{}
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}
}()
})
})
return gopurs_runtime.Apply2(go__go_6_13_124, __local_var_4_12, gopurs_runtime.Apply2(go__go_7_16_125, Get_Data_List_Lazy_Types_nil(), xs_5))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_traversableList__3068288903
}

var cache_Data_List_Lazy_Types_traversableList__2371870579 gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableList__2371870579 sync.Once
func Get_Data_List_Lazy_Types_traversableList__2371870579() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableList__2371870579.Do(func() {
		cache_Data_List_Lazy_Types_traversableList__2371870579 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_3_2
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_3_126 gopurs_runtime.Value
go__go_5_3_126 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var xs_7_loop gopurs_runtime.Value = xs_7_loop_val
go__go_5_3_126:
for {
if false { continue go__go_5_3_126 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var xs_7 gopurs_runtime.Value = xs_7_loop
_ = xs_7
// TAST (Let): v_8_4 -> *Constructor_Data_List_Lazy_Types_Cons
v_8_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_7))
_ = v_8_4
var __t5 gopurs_runtime.Value
{
if (v_8_4 == nil) {
__t5 = b_6
goto end_branch_5
} else {

}
}
{
if (v_8_4 != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), (v_8_4).V0), b_6)
xs_7_loop = (v_8_4).V1
continue go__go_5_3_126
__t5 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_6_6_127 gopurs_runtime.Value
go__go_6_6_127 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_6_127:
for {
if false { continue go__go_6_6_127 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_7 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_7
var __t9 gopurs_runtime.Value
{
if (v_9_7 == nil) {
__t9 = b_7
goto end_branch_9
} else {

}
}
{
if (v_9_7 != nil) {
// TAST (Let): __local_var_10_8 -> gopurs_runtime.Value
__local_var_10_8 := (v_9_7).V0
_ = __local_var_10_8
b_7_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_10_8, b_7})}
}))
xs_8_loop = (v_9_7).V1
continue go__go_6_6_127
__t9 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_5_3_126, __local_var_3_2, gopurs_runtime.Apply2(go__go_6_6_127, Get_Data_List_Lazy_Types_nil(), xs_4))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_10 -> *Constructor_Control_Apply_Apply
Apply0_1_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_10
// TAST (Let): Functor0_2_11 -> *Constructor_Data_Functor_Functor
Functor0_2_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_11
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_12 -> gopurs_runtime.Value
__local_var_4_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil())
_ = __local_var_4_12
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_13_128 gopurs_runtime.Value
go__go_6_13_128 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_13_128:
for {
if false { continue go__go_6_13_128 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_14 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_14
var __t15 gopurs_runtime.Value
{
if (v_9_14 == nil) {
__t15 = b_7
goto end_branch_15
} else {

}
}
{
if (v_9_14 != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_10.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_11.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_9_14).V0))})), b_7)
xs_8_loop = (v_9_14).V1
continue go__go_6_13_128
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
var go__go_7_16_129 gopurs_runtime.Value
go__go_7_16_129 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_16_129:
for {
if false { continue go__go_7_16_129 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_17 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_17
var __t19 gopurs_runtime.Value
{
if (v_10_17 == nil) {
__t19 = b_8
goto end_branch_19
} else {

}
}
{
if (v_10_17 != nil) {
// TAST (Let): __local_var_11_18 -> gopurs_runtime.Value
__local_var_11_18 := (v_10_17).V0
_ = __local_var_11_18
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_11_18, b_8})}
}))
xs_9_loop = (v_10_17).V1
continue go__go_7_16_129
__t19 = gopurs_runtime.Value{}
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}
}()
})
})
return gopurs_runtime.Apply2(go__go_6_13_128, __local_var_4_12, gopurs_runtime.Apply2(go__go_7_16_129, Get_Data_List_Lazy_Types_nil(), xs_5))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_traversableList__2371870579
}

var cache_Data_List_Lazy_Types_traversableNonEmptyList__1985249486 gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmptyList__1985249486 sync.Once
func Get_Data_List_Lazy_Types_traversableNonEmptyList__1985249486() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmptyList__1985249486.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmptyList__1985249486 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2))
_ = __local_var_3_1
// TAST (Let): Apply0_4_2 -> *Constructor_Control_Apply_Apply
Apply0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
var go__go_6_4_130 gopurs_runtime.Value
go__go_6_4_130 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_4_130:
for {
if false { continue go__go_6_4_130 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_5 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_5
var __t6 gopurs_runtime.Value
{
if (v_9_5 == nil) {
__t6 = b_7
goto end_branch_6
} else {

}
}
{
if (v_9_5 != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), Get_Data_List_Lazy_Types_cons(), (v_9_5).V0), b_7)
xs_8_loop = (v_9_5).V1
continue go__go_6_4_130
__t6 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_7_7_131 gopurs_runtime.Value
go__go_7_7_131 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_7_131:
for {
if false { continue go__go_7_7_131 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_8 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_8
var __t10 gopurs_runtime.Value
{
if (v_10_8 == nil) {
__t10 = b_8
goto end_branch_10
} else {

}
}
{
if (v_10_8 != nil) {
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := (v_10_8).V0
_ = __local_var_11_9
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_11_9, b_8})}
}))
xs_9_loop = (v_10_8).V1
continue go__go_7_7_131
__t10 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xxs_3))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), (__local_var_3_1).V0), gopurs_runtime.Apply2(go__go_6_4_130, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_7_7_131, Get_Data_List_Lazy_Types_nil(), (__local_var_3_1).V1))))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_11 -> *Constructor_Data_Functor_Functor
Functor0_1_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_11
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_12 -> *Constructor_Data_NonEmpty_NonEmpty
__local_var_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3))
_ = __local_var_4_12
// TAST (Let): Apply0_5_13 -> *Constructor_Control_Apply_Apply
Apply0_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_13
// TAST (Let): Functor0_6_14 -> *Constructor_Data_Functor_Functor
Functor0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_14
var go__go_7_15_132 gopurs_runtime.Value
go__go_7_15_132 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_15_132:
for {
if false { continue go__go_7_15_132 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_16 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_16
var __t17 gopurs_runtime.Value
{
if (v_10_16 == nil) {
__t17 = b_8
goto end_branch_17
} else {

}
}
{
if (v_10_16 != nil) {
b_8_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_13.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_14.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_2, (v_10_16).V0)), b_8)
xs_9_loop = (v_10_16).V1
continue go__go_7_15_132
__t17 = gopurs_runtime.Value{}
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}
}()
})
})
var go__go_8_18_133 gopurs_runtime.Value
go__go_8_18_133 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var xs_10_loop gopurs_runtime.Value = xs_10_loop_val
go__go_8_18_133:
for {
if false { continue go__go_8_18_133 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_19 -> *Constructor_Data_List_Lazy_Types_Cons
v_11_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_10))
_ = v_11_19
var __t21 gopurs_runtime.Value
{
if (v_11_19 == nil) {
__t21 = b_9
goto end_branch_21
} else {

}
}
{
if (v_11_19 != nil) {
// TAST (Let): __local_var_12_20 -> gopurs_runtime.Value
__local_var_12_20 := (v_11_19).V0
_ = __local_var_12_20
b_9_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_12_20, b_9})}
}))
xs_10_loop = (v_11_19).V1
continue go__go_8_18_133
__t21 = gopurs_runtime.Value{}
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_11.V0), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xxs_4))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_2, (__local_var_4_12).V0)), gopurs_runtime.Apply2(go__go_7_15_132, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_8_18_133, Get_Data_List_Lazy_Types_nil(), (__local_var_4_12).V1))))
})
})
})})}
	})
	return cache_Data_List_Lazy_Types_traversableNonEmptyList__1985249486
}

var cache_Data_List_Lazy_Types_unfoldable1List__338845127 gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1List__338845127 sync.Once
func Get_Data_List_Lazy_Types_unfoldable1List__338845127() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1List__338845127.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1List__338845127 = func() gopurs_runtime.Value {
var go__go_0_0_134 gopurs_runtime.Value
_ = go__go_0_0_134
go__go_0_0_134 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Tuple_Tuple
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (v1_4_1).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(go__go_0_0_134, f_1, (*Constructor_Data_Maybe_Just)((v1_4_1).V1.UnsafePtr).V0)
_ = __local_var_5_3
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v1_4_1).V0, __local_var_5_3})}
}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (v1_4_1).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v1_4_1).V0, Get_Data_List_Lazy_Types_nil()})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t5)
}))
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, go__go_0_0_134})}
}()
	})
	return cache_Data_List_Lazy_Types_unfoldable1List__338845127
}

var cache_Data_List_Lazy_Types_unfoldableList__825189991 gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldableList__825189991 sync.Once
func Get_Data_List_Lazy_Types_unfoldableList__825189991() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldableList__825189991.Do(func() {
		cache_Data_List_Lazy_Types_unfoldableList__825189991 = func() gopurs_runtime.Value {
var go__go_0_0_135 gopurs_runtime.Value
_ = go__go_0_0_135
go__go_0_0_135 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Maybe_Just
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t3 gopurs_runtime.Value
{
if (v1_4_1 == nil) {
__t3 = Get_Data_List_Lazy_Types_nil()
goto end_branch_3
} else {

}
}
{
if (v1_4_1 != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply2(go__go_0_0_135, f_1, (*Constructor_Data_Tuple_Tuple)((v1_4_1).V0.UnsafePtr).V1)
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_Tuple_Tuple)((v1_4_1).V0.UnsafePtr).V0, __local_var_5_2})}
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t3)
}))
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable_Unfoldable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Lazy_Types_unfoldable1List()))}
}), go__go_0_0_135})}
}()
	})
	return cache_Data_List_Lazy_Types_unfoldableList__825189991
}

type Constructor_Data_List_Lazy_Types_Nil struct {
	Rc uint32
}


type Constructor_Data_List_Lazy_Types_Cons struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_Data_List_Lazy_Types_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
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

func Call_Data_List_Lazy_Types_step(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_eqNonEmpty(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_3 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_1 gopurs_runtime.Value
go__go_3_0_1 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_1:
for {
if false { continue go__go_3_0_1 }
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
continue go__go_3_0_1
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
__t_and_3 = (gopurs_runtime.Apply2(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_3)
})
})})}
}

func Call_Data_List_Lazy_Types_eq1(dictEq_0_loop *Constructor_Data_Eq_Eq, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var go__go_3_0_2 gopurs_runtime.Value
go__go_3_0_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_2:
for {
if false { continue go__go_3_0_2 }
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
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)))}
continue go__go_3_0_2
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
return (gopurs_runtime.Apply2(go__go_3_0_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0)
}

func Call_Data_List_Lazy_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_4 gopurs_runtime.Value
go__go_3_0_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_4:
for {
if false { continue go__go_3_0_4 }
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
continue go__go_3_0_4
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_4, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})
})})}
}

func Call_Data_List_Lazy_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_4 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_1_5 gopurs_runtime.Value
go__go_3_1_5 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_1_5:
for {
if false { continue go__go_3_1_5 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 bool
{
if (v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr == nil) {
var __t2 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)))}
continue go__go_3_1_5
__t3 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
}
}()
})
})
__t_and_4 = (gopurs_runtime.Apply2(go__go_3_1_5, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_4)
})
}))
_ = __local_var_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), y_3)).IntVal) != (0))
})
}))))}
}

func Call_Data_List_Lazy_Types_ordNonEmpty(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqNonEmpty2_1_0 -> *Constructor_Data_Eq_Eq
eqNonEmpty2_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_5 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0).IntVal) != (0) {

var go__go_4_2_7 gopurs_runtime.Value
go__go_4_2_7 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_7:
for {
if false { continue go__go_4_2_7 }
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
continue go__go_4_2_7
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
__t_and_5 = (gopurs_runtime.Apply2(go__go_4_2_7, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_5)
})
})))
_ = eqNonEmpty2_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty2_1_0)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_6 -> gopurs_runtime.Value
v_4_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0)
_ = v_4_6
var __t12 uint32
{
if (uint32(v_4_6.IntVal) == 1527465420) {
__t12 = 1527465420
goto end_branch_12
} else {

}
}
{
if (uint32(v_4_6.IntVal) == 380165415) {
__t12 = 380165415
goto end_branch_12
} else {

}
}
{
var go__go_5_7_8 gopurs_runtime.Value
go__go_5_7_8 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_7_8:
for {
if false { continue go__go_5_7_8 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t11 uint32
{
if (v_6.Type == 9 && v_6.IntVal == 218341868 && v_6.UnsafePtr == nil) {
var __t8 uint32
{
if (v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr == nil) {
__t8 = 902936544
goto end_branch_8
} else {

}
}
{
__t8 = 1527465420
}
end_branch_8:
__t11 = __t8
goto end_branch_11
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr == nil) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if ((v_6.Type == 9 && v_6.IntVal == 218341868 && v_6.UnsafePtr != nil)) && ((v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr != nil)) {
// TAST (Let): v2_8_9 -> gopurs_runtime.Value
v2_8_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Lazy_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_7.UnsafePtr).V0)
_ = v2_8_9
var __t10 uint32
{
if (uint32(v2_8_9.IntVal) == 902936544) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_6.UnsafePtr).V1)))}
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_7.UnsafePtr).V1)))}
continue go__go_5_7_8
__t10 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_10
} else {

}
}
{
__t10 = uint32(v2_8_9.IntVal)
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil}
}
}()
})
})
__t12 = uint32(gopurs_runtime.Apply2(go__go_5_7_8, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1)))}).IntVal)
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t12), UnsafePtr: nil}
})
})})}
}

func Call_Data_List_Lazy_Types_compare1(dictOrd_0_loop *Constructor_Data_Ord_Ord, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var go__go_3_0_9 gopurs_runtime.Value
go__go_3_0_9 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_9:
for {
if false { continue go__go_3_0_9 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t4 uint32
{
if (v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr == nil) {
var __t1 uint32
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
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
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) {
// TAST (Let): v2_6_2 -> gopurs_runtime.Value
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0)
_ = v2_6_2
var __t3 uint32
{
if (uint32(v2_6_2.IntVal) == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)))}
continue go__go_3_0_9
__t3 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(v2_6_2.IntVal)
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
}
}()
})
})
return uint32(gopurs_runtime.Apply2(go__go_3_0_9, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal)
}

func Call_Data_List_Lazy_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqList1_1_0 -> *Constructor_Data_Eq_Eq
eqList1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_12 gopurs_runtime.Value
go__go_4_2_12 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_12:
for {
if false { continue go__go_4_2_12 }
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
continue go__go_4_2_12
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_4_2_12, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))}).IntVal) != (0))
})
})))
_ = eqList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_0)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_13 gopurs_runtime.Value
go__go_4_5_13 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_13:
for {
if false { continue go__go_4_5_13 }
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
continue go__go_4_5_13
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_4_5_13, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))}).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_List_Lazy_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqNonEmpty2_1_0 -> *Constructor_Data_Eq_Eq
eqNonEmpty2_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_5 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0).IntVal) != (0) {

var go__go_4_2_14 gopurs_runtime.Value
go__go_4_2_14 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_14:
for {
if false { continue go__go_4_2_14 }
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
continue go__go_4_2_14
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
__t_and_5 = (gopurs_runtime.Apply2(go__go_4_2_14, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1)))}).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_5)
})
})))
_ = eqNonEmpty2_1_0
// TAST (Let): __local_var_2_6 -> *Constructor_Data_Ord_Ord
__local_var_2_6 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty2_1_0)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_7 -> gopurs_runtime.Value
v_4_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0)
_ = v_4_7
var __t13 uint32
{
if (uint32(v_4_7.IntVal) == 1527465420) {
__t13 = 1527465420
goto end_branch_13
} else {

}
}
{
if (uint32(v_4_7.IntVal) == 380165415) {
__t13 = 380165415
goto end_branch_13
} else {

}
}
{
var go__go_5_8_15 gopurs_runtime.Value
go__go_5_8_15 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_8_15:
for {
if false { continue go__go_5_8_15 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t12 uint32
{
if (v_6.Type == 9 && v_6.IntVal == 218341868 && v_6.UnsafePtr == nil) {
var __t9 uint32
{
if (v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr == nil) {
__t9 = 902936544
goto end_branch_9
} else {

}
}
{
__t9 = 1527465420
}
end_branch_9:
__t12 = __t9
goto end_branch_12
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr == nil) {
__t12 = 380165415
goto end_branch_12
} else {

}
}
{
if ((v_6.Type == 9 && v_6.IntVal == 218341868 && v_6.UnsafePtr != nil)) && ((v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr != nil)) {
// TAST (Let): v2_8_10 -> gopurs_runtime.Value
v2_8_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Lazy_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_7.UnsafePtr).V0)
_ = v2_8_10
var __t11 uint32
{
if (uint32(v2_8_10.IntVal) == 902936544) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_6.UnsafePtr).V1)))}
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_7.UnsafePtr).V1)))}
continue go__go_5_8_15
__t11 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_11
} else {

}
}
{
__t11 = uint32(v2_8_10.IntVal)
}
end_branch_11:
__t12 = __t11
goto end_branch_12
} else {

}
}
{
__t12 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t12), UnsafePtr: nil}
}
}()
})
})
__t13 = uint32(gopurs_runtime.Apply2(go__go_5_8_15, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1)))}).IntVal)
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t13), UnsafePtr: nil}
})
})}
_ = __local_var_2_6
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_6.V0), gopurs_runtime.Value{})
_ = __local_var_3_15
// TAST (Let): eqLazy1_3_14 -> *Constructor_Data_Eq_Eq
eqLazy1_3_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_15, "eq"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_4), gopurs_runtime.Apply(Get_Data_Lazy_force(), y_5)).IntVal) != (0))
})
})))
_ = eqLazy1_3_14
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqLazy1_3_14)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_2_6.V1), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_4), gopurs_runtime.Apply(Get_Data_Lazy_force(), y_5)).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_List_Lazy_Types_cons(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
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
var go__go_3_1_25 gopurs_runtime.Value
go__go_3_1_25 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_1_25:
for {
if false { continue go__go_3_1_25 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_2 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
continue go__go_3_1_25
__t3 = gopurs_runtime.Value{}
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
}()
})
})
__t4 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_0).V0).StrVal())) + (gopurs_runtime.Apply2(go__go_3_1_25, gopurs_runtime.Str(""), (v_2_0).V1).StrVal())) + ("])")
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_4:
return gopurs_runtime.Str(__t4)
})})}
}

func Call_Data_List_Lazy_Types_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_2 -> *Constructor_Data_List_Lazy_Types_Cons
v_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
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
var go__go_3_3_26 gopurs_runtime.Value
go__go_3_3_26 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_3_26:
for {
if false { continue go__go_3_3_26 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_4 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
continue go__go_3_3_26
__t5 = gopurs_runtime.Value{}
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
}()
})
})
__t6 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_2).V0).StrVal())) + (gopurs_runtime.Apply2(go__go_3_3_26, gopurs_runtime.Str(""), (v_2_2).V1).StrVal())) + ("])")
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_6:
return gopurs_runtime.Str(__t6)
}))
_ = __local_var_1_1
// TAST (Let): __local_var_2_7 -> gopurs_runtime.Value
__local_var_2_7 := gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(NonEmpty ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "show"), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))
_ = __local_var_2_7
// TAST (Let): showLazy_1_0 -> *Constructor_Data_Show_Show
showLazy_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(defer \\_ -> ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "show"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_3)).StrVal())) + (")"))
})))
_ = showLazy_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyList ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showLazy_1_0.V0), v_2).StrVal())) + (")"))
})})}
}

func Call_Data_List_Lazy_Types_showStep(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList1_1_0 -> *Constructor_Data_Show_Show
showList1_1_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
var go__go_3_2_27 gopurs_runtime.Value
go__go_3_2_27 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_2_27:
for {
if false { continue go__go_3_2_27 }
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
continue go__go_3_2_27
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
__t5 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_1).V0).StrVal())) + (gopurs_runtime.Apply2(go__go_3_2_27, gopurs_runtime.Str(""), (v_2_1).V1).StrVal())) + ("])")
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
_ = showList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 string
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t6 = "Nil"
goto end_branch_6
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t6 = (((("(") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0).StrVal())) + (" : ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(showList1_1_0.V0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_6:
return gopurs_runtime.Str(__t6)
})})}
}

func Call_Data_List_Lazy_Types_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V1})}
})))
}))
}

func Call_Data_List_Lazy_Types_cons__716923058(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__720046150(x_0_loop *Constructor_Data_List_Lazy_Types_Cons, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__894912754(x_0_loop *Constructor_Data_List_Lazy_Types_Cons, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__376540526(x_0_loop *Constructor_Data_List_Lazy_Types_Cons, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__673400617(x_0_loop *Constructor_Data_List_Lazy_Types_Cons, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__2305074921(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__891310957(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__1901546616(x_0_loop *Constructor_Data_List_Lazy_Types_Cons, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__3391588829(x_0_loop *Constructor_Data_List_Lazy_Types_Cons, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__2134285409(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_step__3545407802(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__3687829882(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__4184651873(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__4057057377(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__2986341153(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__2999566881(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__2528948705(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__2322903873(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__2597188449(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__1903922273(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_toList__1017592434(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V1})}
})))
}))
}

func Call_Data_List_Lazy_Types_toList__4101396777(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v2_2_0.UnsafePtr).V1})}
})))
}))
}


