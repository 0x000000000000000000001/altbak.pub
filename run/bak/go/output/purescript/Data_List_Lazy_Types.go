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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil
}

var cache_Data_List_Lazy_Types_newtypeNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_newtypeNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_newtypeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_newtypeNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_newtypeNonEmptyList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_List_Lazy_Types_newtypeNonEmptyList
}

var cache_Data_List_Lazy_Types_newtypeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_newtypeList sync.Once
func Get_Data_List_Lazy_Types_newtypeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_newtypeList.Do(func() {
		cache_Data_List_Lazy_Types_newtypeList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
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
		cache_Data_List_Lazy_Types_semigroupList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_semigroupList
}

var cache_Data_List_Lazy_Types_monoidList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monoidList sync.Once
func Get_Data_List_Lazy_Types_monoidList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monoidList.Do(func() {
		cache_Data_List_Lazy_Types_monoidList = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_semigroupList()
}), Get_Data_List_Lazy_Types_nil())
	})
	return cache_Data_List_Lazy_Types_monoidList
}

var cache_Data_List_Lazy_Types_lazyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_lazyList sync.Once
func Get_Data_List_Lazy_Types_lazyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_lazyList.Do(func() {
		cache_Data_List_Lazy_Types_lazyList = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, x_1))
}))
}))
	})
	return cache_Data_List_Lazy_Types_lazyList
}

var cache_Data_List_Lazy_Types_functorList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList sync.Once
func Get_Data_List_Lazy_Types_functorList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList.Do(func() {
		cache_Data_List_Lazy_Types_functorList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_functorList(), "map"), f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_Data_List_Lazy_Types_functorList
}

var cache_Data_List_Lazy_Types_functorNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmpty sync.Once
func Get_Data_List_Lazy_Types_functorNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_functorList(), "map"), f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
})})}
	})
	return cache_Data_List_Lazy_Types_functorNonEmpty
}

var cache_Data_List_Lazy_Types_functorNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_functorNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmptyList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_functorNonEmpty()).V0), f_0), v_1)
})
}))
	})
	return cache_Data_List_Lazy_Types_functorNonEmptyList
}

var cache_Data_List_Lazy_Types_eq1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1List sync.Once
func Get_Data_List_Lazy_Types_eq1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1List.Do(func() {
		cache_Data_List_Lazy_Types_eq1List = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})
})
}))
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
		cache_Data_List_Lazy_Types_eq1 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_eq1List(), "eq1")
	})
	return cache_Data_List_Lazy_Types_eq1
}

var cache_Data_List_Lazy_Types_eq1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_eq1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_eq1NonEmptyList = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqNonEmpty1_1_0 -> *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]
eqNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Call_Data_List_Lazy_Types_eqNonEmpty(dictEq_0))
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Lazy_eq1Lazy(), "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty1_1_0)}, v_2, v1_3).IntVal) != (0))
})
})
}))
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
		cache_Data_List_Lazy_Types_ord1List = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_eq1List()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
var __t4 gopurs_runtime.Value
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
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) {
// TAST (Let): v2_6_2 -> gopurs_runtime.Value
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0)
_ = v2_6_2
var __t3 uint32
{
if (uint32(v2_6_2.IntVal) == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
continue go__go_3_0_1
__t3 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(v2_6_2.IntVal)
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t4.IntVal)), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_Data_List_Lazy_Types_ord1List
}

var cache_Data_List_Lazy_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordNonEmpty sync.Once
func Get_Data_List_Lazy_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_ordNonEmpty = gopurs_runtime.Apply(Get_Data_NonEmpty_ordNonEmpty(), Get_Data_List_Lazy_Types_ord1List())
	})
	return cache_Data_List_Lazy_Types_ordNonEmpty
}

var cache_Data_List_Lazy_Types_compare1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compare1 sync.Once
func Get_Data_List_Lazy_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compare1.Do(func() {
		cache_Data_List_Lazy_Types_compare1 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_ord1List(), "compare1")
	})
	return cache_Data_List_Lazy_Types_compare1
}

var cache_Data_List_Lazy_Types_ord1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_ord1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_ord1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ord1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_ord1NonEmptyList = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_eq1NonEmptyList()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordNonEmpty1_1_0 -> *Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]
ordNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_ordNonEmpty(), dictOrd_0))
_ = ordNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Lazy_ord1Lazy(), "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordNonEmpty1_1_0)}, v_2, v1_3).IntVal)), UnsafePtr: nil}
})
})
}))
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
		cache_Data_List_Lazy_Types_foldableList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_2 gopurs_runtime.Value
go__go_1_2_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_2:
for {
if false { continue go__go_1_2_2 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 -> *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_2
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
return go__go_1_2_2
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableList
}

var cache_Data_List_Lazy_Types_foldableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmpty sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldableNonEmpty(), Get_Data_List_Lazy_Types_foldableList())))}
	})
	return cache_Data_List_Lazy_Types_foldableNonEmpty
}

var cache_Data_List_Lazy_Types_extendList gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendList sync.Once
func Get_Data_List_Lazy_Types_extendList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendList.Do(func() {
		cache_Data_List_Lazy_Types_extendList = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 -> *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
_ = v_2_0
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr == nil) {
__t5 = Get_Data_List_Lazy_Types_nil()
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr != nil) {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(f_0, l_1)
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): acc_prime_6_3 -> gopurs_runtime.Value
acc_prime_6_3 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, gopurs_runtime.RecordGet(v_5, "acc")})}
}))
_ = acc_prime_6_3
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.Apply(f_0, acc_prime_6_3)
_ = __local_var_7_4
return gopurs_runtime.RecordDict2("acc", "val", acc_prime_6_3, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_4, gopurs_runtime.RecordGet(v_5, "val")})}
})))
})
}), gopurs_runtime.RecordDict2("acc", "val", Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V1), "val")
_ = __local_var_4_2
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_4_2})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})
}))
	})
	return cache_Data_List_Lazy_Types_extendList
}

var cache_Data_List_Lazy_Types_extendNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_extendNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_extendNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1).UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := gopurs_runtime.RecordGet(v1_5, "acc")
_ = __local_var_6_1
// TAST (Let): __local_var_7_2 -> gopurs_runtime.Value
__local_var_7_2 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, __local_var_6_1})}
})))
_ = __local_var_7_2
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, __local_var_6_1})}
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_2, gopurs_runtime.RecordGet(v1_5, "val")})}
})))
})
}), gopurs_runtime.RecordDict2("acc", "val", Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()), __local_var_2_0), "val")})}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_extendNonEmptyList
}

var cache_Data_List_Lazy_Types_foldableNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmptyList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V1), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V2), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}))
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
		cache_Data_List_Lazy_Types_foldableWithIndexList = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_6.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_7 -> *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]
v_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_9.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_6, __local_var_4_8})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_2))
_ = v_3_7
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
// TAST (Let): __local_var_6_11 -> gopurs_runtime.Value
__local_var_6_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1).UnsafePtr).V1
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexList
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmpty sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), Get_Data_List_Lazy_Types_foldableWithIndexList())))}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableNonEmptyList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_1, gopurs_runtime.Int(__t0.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V2), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t1.IntVal))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V3), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t2 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t2.IntVal))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_functorWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexList sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2)
_ = __local_var_4_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_0, acc_3})}
}))
})
})
}), Get_Data_List_Lazy_Types_nil())
}))
	})
	return cache_Data_List_Lazy_Types_functorWithIndexList
}

var cache_Data_List_Lazy_Types_functorWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndex sync.Once
func Get_Data_List_Lazy_Types_functorWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndex = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_functorWithIndexList(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): functorNonEmpty1_0_0 -> gopurs_runtime.Value
functorNonEmpty1_0_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "map"), f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_0_0
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_functorWithIndexList(), "mapWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_3})})
}), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
}))))}
}()
	})
	return cache_Data_List_Lazy_Types_functorWithIndex
}

var cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_functorWithIndex()).V1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t0.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1)))})))}
}))
})
}))
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
		cache_Data_List_Lazy_Types_semigroupNonEmptyList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), __local_var_4_2, Call_Data_List_Lazy_Types_toList(as_prime_1))})}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_semigroupNonEmptyList
}

var cache_Data_List_Lazy_Types_traversableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableList sync.Once
func Get_Data_List_Lazy_Types_traversableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableList.Do(func() {
		cache_Data_List_Lazy_Types_traversableList = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_List_Lazy_Types_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
})
}))
	})
	return cache_Data_List_Lazy_Types_traversableList
}

var cache_Data_List_Lazy_Types_traversableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmpty sync.Once
func Get_Data_List_Lazy_Types_traversableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), Get_Data_List_Lazy_Types_traversableList())))}
	})
	return cache_Data_List_Lazy_Types_traversableNonEmpty
}

var cache_Data_List_Lazy_Types_traversableNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_traversableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmptyList = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_3))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_traversableNonEmpty()).V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_traversableNonEmpty()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_traversableNonEmptyList
}

var cache_Data_List_Lazy_Types_traversableWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexList sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexList = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_traversableList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply2(f_3, gopurs_runtime.Int(i_4.IntVal), a_5)), b_6)
})
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
})
}))
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexList
}

var cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexNonEmpty sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableWithIndexNonEmpty(), Get_Data_List_Lazy_Types_traversableWithIndexList())))}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexNonEmpty
}

var cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_traversableWithIndexNonEmpty()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 930809136 && x_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 930809136 && x_4.UnsafePtr != nil) {
__t1 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_4.UnsafePtr).V0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, gopurs_runtime.Int(__t1.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_unfoldable1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1List sync.Once
func Get_Data_List_Lazy_Types_unfoldable1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1List.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1List = func() gopurs_runtime.Value {
var go__go_0_0_3 gopurs_runtime.Value
_ = go__go_0_0_3
go__go_0_0_3 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(go__go_0_0_3, f_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1.UnsafePtr).V0)
_ = __local_var_5_3
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0, __local_var_5_3})}
}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0, Get_Data_List_Lazy_Types_nil()})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
})
})
return gopurs_runtime.RecordDict1("unfoldr1", go__go_0_0_3)
}()
	})
	return cache_Data_List_Lazy_Types_unfoldable1List
}

var cache_Data_List_Lazy_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldableList sync.Once
func Get_Data_List_Lazy_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldableList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldableList = func() gopurs_runtime.Value {
var go__go_0_0_4 gopurs_runtime.Value
_ = go__go_0_0_4
go__go_0_0_4 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr == nil) {
__t3 = Get_Data_List_Lazy_Types_nil()
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply2(go__go_0_0_4, f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V1)
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V0, __local_var_5_2})}
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
})
})
return gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_unfoldable1List()
}), go__go_0_0_4)
}()
	})
	return cache_Data_List_Lazy_Types_unfoldableList
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmpty sync.Once
func Get_Data_List_Lazy_Types_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_NonEmpty_unfoldable1NonEmpty(), Get_Data_List_Lazy_Types_unfoldableList())))}
	})
	return cache_Data_List_Lazy_Types_unfoldable1NonEmpty
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmptyList sync.Once
func Get_Data_List_Lazy_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmptyList = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_unfoldable1NonEmpty()).V0), f_0, b_1)))}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_unfoldable1NonEmptyList
}

var cache_Data_List_Lazy_Types_comonadNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_comonadNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_comonadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_comonadNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_comonadNonEmptyList = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_extendNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V0
}))
	})
	return cache_Data_List_Lazy_Types_comonadNonEmptyList
}

var cache_Data_List_Lazy_Types_monadList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadList sync.Once
func Get_Data_List_Lazy_Types_monadList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadList.Do(func() {
		cache_Data_List_Lazy_Types_monadList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_bindList()
}))
	})
	return cache_Data_List_Lazy_Types_monadList
}

var cache_Data_List_Lazy_Types_bindList gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList sync.Once
func Get_Data_List_Lazy_Types_bindList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList.Do(func() {
		cache_Data_List_Lazy_Types_bindList = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_bindList
}

var cache_Data_List_Lazy_Types_applyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList sync.Once
func Get_Data_List_Lazy_Types_applyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList.Do(func() {
		cache_Data_List_Lazy_Types_applyList = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Data_List_Lazy_Types_applyList
}

var cache_Data_List_Lazy_Types_applicativeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeList sync.Once
func Get_Data_List_Lazy_Types_applicativeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeList.Do(func() {
		cache_Data_List_Lazy_Types_applicativeList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()})}
}))
}))
	})
	return cache_Data_List_Lazy_Types_applicativeList
}

var cache_Data_List_Lazy_Types_applyNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_applyNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_applyNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_1)
_ = v2_2_0
// TAST (Let): v3_3_1 -> gopurs_runtime.Value
v3_3_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v3_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3_1.UnsafePtr).V0
_ = __local_var_6_4
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3_1.UnsafePtr).V1
_ = __local_var_7_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_applyList(), "apply"), __local_var_7_5, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_2, Get_Data_List_Lazy_Types_nil()})}
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_applyList(), "apply"), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_4, __local_var_7_5})}
})), __local_var_5_3))})}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_applyNonEmptyList
}

var cache_Data_List_Lazy_Types_bindNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_bindNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_bindNonEmptyList = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): v2_4_2 -> *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]
v2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0)))
_ = v2_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v2_4_2)}.UnsafePtr).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v2_4_2)}.UnsafePtr).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), __local_var_6_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), __local_var_3_1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_toList(gopurs_runtime.Apply(f_1, x_8))
})))})}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_bindNonEmptyList
}

var cache_Data_List_Lazy_Types_altNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_altNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_altNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupNonEmptyList(), "append"))
	})
	return cache_Data_List_Lazy_Types_altNonEmptyList
}

var cache_Data_List_Lazy_Types_altList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altList sync.Once
func Get_Data_List_Lazy_Types_altList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altList.Do(func() {
		cache_Data_List_Lazy_Types_altList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"))
	})
	return cache_Data_List_Lazy_Types_altList
}

var cache_Data_List_Lazy_Types_plusList gopurs_runtime.Value
var once_Data_List_Lazy_Types_plusList sync.Once
func Get_Data_List_Lazy_Types_plusList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_plusList.Do(func() {
		cache_Data_List_Lazy_Types_plusList = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_altList()
}), Get_Data_List_Lazy_Types_nil())
	})
	return cache_Data_List_Lazy_Types_plusList
}

var cache_Data_List_Lazy_Types_alternativeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_alternativeList sync.Once
func Get_Data_List_Lazy_Types_alternativeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_alternativeList.Do(func() {
		cache_Data_List_Lazy_Types_alternativeList = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_plusList()
}))
	})
	return cache_Data_List_Lazy_Types_alternativeList
}

var cache_Data_List_Lazy_Types_monadPlusList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadPlusList sync.Once
func Get_Data_List_Lazy_Types_monadPlusList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadPlusList.Do(func() {
		cache_Data_List_Lazy_Types_monadPlusList = gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_alternativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_monadList()
}))
	})
	return cache_Data_List_Lazy_Types_monadPlusList
}

var cache_Data_List_Lazy_Types_applicativeNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_applicativeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_applicativeNonEmptyList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_plusList(), "empty")})}
}))
}))
	})
	return cache_Data_List_Lazy_Types_applicativeNonEmptyList
}

var cache_Data_List_Lazy_Types_monadNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadNonEmptyList sync.Once
func Get_Data_List_Lazy_Types_monadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_monadNonEmptyList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applicativeNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_bindNonEmptyList()
}))
	})
	return cache_Data_List_Lazy_Types_monadNonEmptyList
}

var cache_Data_List_Lazy_Types_altList__3296309911 gopurs_runtime.Value
var once_Data_List_Lazy_Types_altList__3296309911 sync.Once
func Get_Data_List_Lazy_Types_altList__3296309911() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altList__3296309911.Do(func() {
		cache_Data_List_Lazy_Types_altList__3296309911 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"))
	})
	return cache_Data_List_Lazy_Types_altList__3296309911
}

var cache_Data_List_Lazy_Types_alternativeList__16377757 gopurs_runtime.Value
var once_Data_List_Lazy_Types_alternativeList__16377757 sync.Once
func Get_Data_List_Lazy_Types_alternativeList__16377757() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_alternativeList__16377757.Do(func() {
		cache_Data_List_Lazy_Types_alternativeList__16377757 = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_plusList()
}))
	})
	return cache_Data_List_Lazy_Types_alternativeList__16377757
}

var cache_Data_List_Lazy_Types_applicativeList__37190504 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeList__37190504 sync.Once
func Get_Data_List_Lazy_Types_applicativeList__37190504() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeList__37190504.Do(func() {
		cache_Data_List_Lazy_Types_applicativeList__37190504 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()})}
}))
}))
	})
	return cache_Data_List_Lazy_Types_applicativeList__37190504
}

var cache_Data_List_Lazy_Types_applicativeNonEmptyList__1064240936 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeNonEmptyList__1064240936 sync.Once
func Get_Data_List_Lazy_Types_applicativeNonEmptyList__1064240936() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeNonEmptyList__1064240936.Do(func() {
		cache_Data_List_Lazy_Types_applicativeNonEmptyList__1064240936 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_plusList(), "empty")})}
}))
}))
	})
	return cache_Data_List_Lazy_Types_applicativeNonEmptyList__1064240936
}

var cache_Data_List_Lazy_Types_applyList__1470982352 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList__1470982352 sync.Once
func Get_Data_List_Lazy_Types_applyList__1470982352() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList__1470982352.Do(func() {
		cache_Data_List_Lazy_Types_applyList__1470982352 = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Data_List_Lazy_Types_applyList__1470982352
}

var cache_Data_List_Lazy_Types_applyList__1358886895 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList__1358886895 sync.Once
func Get_Data_List_Lazy_Types_applyList__1358886895() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList__1358886895.Do(func() {
		cache_Data_List_Lazy_Types_applyList__1358886895 = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Data_List_Lazy_Types_applyList__1358886895
}

var cache_Data_List_Lazy_Types_applyNonEmptyList__1016064038 gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyNonEmptyList__1016064038 sync.Once
func Get_Data_List_Lazy_Types_applyNonEmptyList__1016064038() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyNonEmptyList__1016064038.Do(func() {
		cache_Data_List_Lazy_Types_applyNonEmptyList__1016064038 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_1)
_ = v2_2_0
// TAST (Let): v3_3_1 -> gopurs_runtime.Value
v3_3_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v3_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3_1.UnsafePtr).V0
_ = __local_var_6_4
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3_1.UnsafePtr).V1
_ = __local_var_7_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_applyList(), "apply"), __local_var_7_5, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_2, Get_Data_List_Lazy_Types_nil()})}
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_applyList(), "apply"), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_4, __local_var_7_5})}
})), __local_var_5_3))})}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_applyNonEmptyList__1016064038
}

var cache_Data_List_Lazy_Types_bindList__469219920 gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList__469219920 sync.Once
func Get_Data_List_Lazy_Types_bindList__469219920() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList__469219920.Do(func() {
		cache_Data_List_Lazy_Types_bindList__469219920 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_bindList__469219920
}

var cache_Data_List_Lazy_Types_bindList__1050117088 gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList__1050117088 sync.Once
func Get_Data_List_Lazy_Types_bindList__1050117088() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList__1050117088.Do(func() {
		cache_Data_List_Lazy_Types_bindList__1050117088 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_bindList__1050117088
}

var cache_Data_List_Lazy_Types_bindList__1503691431 gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList__1503691431 sync.Once
func Get_Data_List_Lazy_Types_bindList__1503691431() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList__1503691431.Do(func() {
		cache_Data_List_Lazy_Types_bindList__1503691431 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_bindList__1503691431
}

var cache_Data_List_Lazy_Types_bindNonEmptyList__3512977895 gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindNonEmptyList__3512977895 sync.Once
func Get_Data_List_Lazy_Types_bindNonEmptyList__3512977895() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindNonEmptyList__3512977895.Do(func() {
		cache_Data_List_Lazy_Types_bindNonEmptyList__3512977895 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): v2_4_2 -> *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]
v2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0)))
_ = v2_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v2_4_2)}.UnsafePtr).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v2_4_2)}.UnsafePtr).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), __local_var_6_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), __local_var_3_1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_toList(gopurs_runtime.Apply(f_1, x_8))
})))})}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_bindNonEmptyList__3512977895
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
return Call_Data_List_Lazy_Types_cons__720046150(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__720046150
}

var cache_Data_List_Lazy_Types_cons__376540526 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__376540526 sync.Once
func Get_Data_List_Lazy_Types_cons__376540526() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__376540526.Do(func() {
		cache_Data_List_Lazy_Types_cons__376540526 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__376540526(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__376540526
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
return Call_Data_List_Lazy_Types_cons__1901546616(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[int64]](x_0_box), xs_1_box)
})
	})
	return cache_Data_List_Lazy_Types_cons__1901546616
}

var cache_Data_List_Lazy_Types_cons__3391588829 gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons__3391588829 sync.Once
func Get_Data_List_Lazy_Types_cons__3391588829() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons__3391588829.Do(func() {
		cache_Data_List_Lazy_Types_cons__3391588829 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Types_cons__3391588829(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](x_0_box), xs_1_box)
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

var cache_Data_List_Lazy_Types_eq1List__2902948510 gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1List__2902948510 sync.Once
func Get_Data_List_Lazy_Types_eq1List__2902948510() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1List__2902948510.Do(func() {
		cache_Data_List_Lazy_Types_eq1List__2902948510 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_5 gopurs_runtime.Value
go__go_3_0_5 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_5:
for {
if false { continue go__go_3_0_5 }
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
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
continue go__go_3_0_5
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_5, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_eq1List__2902948510
}

var cache_Data_List_Lazy_Types_eq1NonEmptyList__2973916286 gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1NonEmptyList__2973916286 sync.Once
func Get_Data_List_Lazy_Types_eq1NonEmptyList__2973916286() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1NonEmptyList__2973916286.Do(func() {
		cache_Data_List_Lazy_Types_eq1NonEmptyList__2973916286 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqNonEmpty1_1_0 -> *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]
eqNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Call_Data_List_Lazy_Types_eqNonEmpty(dictEq_0))
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Lazy_eq1Lazy(), "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty1_1_0)}, v_2, v1_3).IntVal) != (0))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_eq1NonEmptyList__2973916286
}

var cache_Data_List_Lazy_Types_extendNonEmptyList__2847971828 gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendNonEmptyList__2847971828 sync.Once
func Get_Data_List_Lazy_Types_extendNonEmptyList__2847971828() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendNonEmptyList__2847971828.Do(func() {
		cache_Data_List_Lazy_Types_extendNonEmptyList__2847971828 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1).UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := gopurs_runtime.RecordGet(v1_5, "acc")
_ = __local_var_6_1
// TAST (Let): __local_var_7_2 -> gopurs_runtime.Value
__local_var_7_2 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, __local_var_6_1})}
})))
_ = __local_var_7_2
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, __local_var_6_1})}
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_2, gopurs_runtime.RecordGet(v1_5, "val")})}
})))
})
}), gopurs_runtime.RecordDict2("acc", "val", Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()), __local_var_2_0), "val")})}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_extendNonEmptyList__2847971828
}

var cache_Data_List_Lazy_Types_foldableList__4097915271 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__4097915271 sync.Once
func Get_Data_List_Lazy_Types_foldableList__4097915271() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__4097915271.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__4097915271 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_6 gopurs_runtime.Value
go__go_1_2_6 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_6:
for {
if false { continue go__go_1_2_6 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 -> *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_6
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
return go__go_1_2_6
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableList__4097915271
}

var cache_Data_List_Lazy_Types_foldableList__331628915 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__331628915 sync.Once
func Get_Data_List_Lazy_Types_foldableList__331628915() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__331628915.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__331628915 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_7 gopurs_runtime.Value
go__go_1_2_7 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_7:
for {
if false { continue go__go_1_2_7 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 -> *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_7
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
return go__go_1_2_7
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableList__331628915
}

var cache_Data_List_Lazy_Types_foldableList__21955931 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__21955931 sync.Once
func Get_Data_List_Lazy_Types_foldableList__21955931() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__21955931.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__21955931 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_8 gopurs_runtime.Value
go__go_1_2_8 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_8:
for {
if false { continue go__go_1_2_8 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 -> *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_8
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
return go__go_1_2_8
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableList__21955931
}

var cache_Data_List_Lazy_Types_foldableList__3094856796 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__3094856796 sync.Once
func Get_Data_List_Lazy_Types_foldableList__3094856796() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__3094856796.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__3094856796 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_9 gopurs_runtime.Value
go__go_1_2_9 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_9:
for {
if false { continue go__go_1_2_9 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 -> *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_9
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
return go__go_1_2_9
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableList__3094856796
}

var cache_Data_List_Lazy_Types_foldableList__1218280485 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList__1218280485 sync.Once
func Get_Data_List_Lazy_Types_foldableList__1218280485() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList__1218280485.Do(func() {
		cache_Data_List_Lazy_Types_foldableList__1218280485 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_10 gopurs_runtime.Value
go__go_1_2_10 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_10:
for {
if false { continue go__go_1_2_10 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_3 -> *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_10
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
return go__go_1_2_10
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableList__1218280485
}

var cache_Data_List_Lazy_Types_foldableNonEmptyList__2027644716 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmptyList__2027644716 sync.Once
func Get_Data_List_Lazy_Types_foldableNonEmptyList__2027644716() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmptyList__2027644716.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmptyList__2027644716 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V1), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableNonEmpty()).V2), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableNonEmptyList__2027644716
}

var cache_Data_List_Lazy_Types_foldableWithIndexList__662860203 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexList__662860203 sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexList__662860203() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexList__662860203.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexList__662860203 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_6.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_7 -> *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]
v_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_9.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_6, __local_var_4_8})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_2))
_ = v_3_7
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
// TAST (Let): __local_var_6_11 -> gopurs_runtime.Value
__local_var_6_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1).UnsafePtr).V1
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexList__662860203
}

var cache_Data_List_Lazy_Types_foldableWithIndexList__3899545502 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexList__3899545502 sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexList__3899545502() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexList__3899545502.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexList__3899545502 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_6.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_7 -> *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]
v_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_9.IntVal) + (1)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_6, __local_var_4_8})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_2))
_ = v_3_7
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
// TAST (Let): __local_var_6_11 -> gopurs_runtime.Value
__local_var_6_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1).UnsafePtr).V1
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexList__3899545502
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__3565687582 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__3565687582 sync.Once
func Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__3565687582() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__3565687582.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__3565687582 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableNonEmptyList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_1, gopurs_runtime.Int(__t0.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V2), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t1.IntVal))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V3), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t2 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t2.IntVal))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
})
})
}))
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList__3565687582
}

var cache_Data_List_Lazy_Types_functorList__699353223 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList__699353223 sync.Once
func Get_Data_List_Lazy_Types_functorList__699353223() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList__699353223.Do(func() {
		cache_Data_List_Lazy_Types_functorList__699353223 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_functorList(), "map"), f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_Data_List_Lazy_Types_functorList__699353223
}

var cache_Data_List_Lazy_Types_functorList__3996674161 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList__3996674161 sync.Once
func Get_Data_List_Lazy_Types_functorList__3996674161() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList__3996674161.Do(func() {
		cache_Data_List_Lazy_Types_functorList__3996674161 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_functorList(), "map"), f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_Data_List_Lazy_Types_functorList__3996674161
}

var cache_Data_List_Lazy_Types_functorNonEmptyList__1934212625 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmptyList__1934212625 sync.Once
func Get_Data_List_Lazy_Types_functorNonEmptyList__1934212625() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmptyList__1934212625.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmptyList__1934212625 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_functorNonEmpty()).V0), f_0), v_1)
})
}))
	})
	return cache_Data_List_Lazy_Types_functorNonEmptyList__1934212625
}

var cache_Data_List_Lazy_Types_functorWithIndexList__353314402 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexList__353314402 sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexList__353314402() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexList__353314402.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexList__353314402 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2)
_ = __local_var_4_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_0, acc_3})}
}))
})
})
}), Get_Data_List_Lazy_Types_nil())
}))
	})
	return cache_Data_List_Lazy_Types_functorWithIndexList__353314402
}

var cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList__812352994 gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexNonEmptyList__812352994 sync.Once
func Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList__812352994() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexNonEmptyList__812352994.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList__812352994 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_functorWithIndex()).V1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t0.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1)))})))}
}))
})
}))
	})
	return cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList__812352994
}

var cache_Data_List_Lazy_Types_lazyList__601034736 gopurs_runtime.Value
var once_Data_List_Lazy_Types_lazyList__601034736 sync.Once
func Get_Data_List_Lazy_Types_lazyList__601034736() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_lazyList__601034736.Do(func() {
		cache_Data_List_Lazy_Types_lazyList__601034736 = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, x_1))
}))
}))
	})
	return cache_Data_List_Lazy_Types_lazyList__601034736
}

var cache_Data_List_Lazy_Types_lazyList__706698273 gopurs_runtime.Value
var once_Data_List_Lazy_Types_lazyList__706698273 sync.Once
func Get_Data_List_Lazy_Types_lazyList__706698273() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_lazyList__706698273.Do(func() {
		cache_Data_List_Lazy_Types_lazyList__706698273 = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, x_1))
}))
}))
	})
	return cache_Data_List_Lazy_Types_lazyList__706698273
}

var cache_Data_List_Lazy_Types_monadList__2596899283 gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadList__2596899283 sync.Once
func Get_Data_List_Lazy_Types_monadList__2596899283() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadList__2596899283.Do(func() {
		cache_Data_List_Lazy_Types_monadList__2596899283 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_bindList()
}))
	})
	return cache_Data_List_Lazy_Types_monadList__2596899283
}

var cache_Data_List_Lazy_Types_monoidList__245587391 gopurs_runtime.Value
var once_Data_List_Lazy_Types_monoidList__245587391 sync.Once
func Get_Data_List_Lazy_Types_monoidList__245587391() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monoidList__245587391.Do(func() {
		cache_Data_List_Lazy_Types_monoidList__245587391 = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_semigroupList()
}), Get_Data_List_Lazy_Types_nil())
	})
	return cache_Data_List_Lazy_Types_monoidList__245587391
}

var cache_Data_List_Lazy_Types_nil__1478684294 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__1478684294 sync.Once
func Get_Data_List_Lazy_Types_nil__1478684294() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__1478684294.Do(func() {
		cache_Data_List_Lazy_Types_nil__1478684294 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__1478684294
}

var cache_Data_List_Lazy_Types_nil__3988504114 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__3988504114 sync.Once
func Get_Data_List_Lazy_Types_nil__3988504114() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__3988504114.Do(func() {
		cache_Data_List_Lazy_Types_nil__3988504114 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__3988504114
}

var cache_Data_List_Lazy_Types_nil__1778182234 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__1778182234 sync.Once
func Get_Data_List_Lazy_Types_nil__1778182234() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__1778182234.Do(func() {
		cache_Data_List_Lazy_Types_nil__1778182234 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__1778182234
}

var cache_Data_List_Lazy_Types_nil__2012296605 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__2012296605 sync.Once
func Get_Data_List_Lazy_Types_nil__2012296605() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__2012296605.Do(func() {
		cache_Data_List_Lazy_Types_nil__2012296605 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__2012296605
}

var cache_Data_List_Lazy_Types_nil__2014033708 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__2014033708 sync.Once
func Get_Data_List_Lazy_Types_nil__2014033708() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__2014033708.Do(func() {
		cache_Data_List_Lazy_Types_nil__2014033708 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__2014033708
}

var cache_Data_List_Lazy_Types_nil__2288399465 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__2288399465 sync.Once
func Get_Data_List_Lazy_Types_nil__2288399465() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__2288399465.Do(func() {
		cache_Data_List_Lazy_Types_nil__2288399465 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__2288399465
}

var cache_Data_List_Lazy_Types_nil__4122162182 gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil__4122162182 sync.Once
func Get_Data_List_Lazy_Types_nil__4122162182() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil__4122162182.Do(func() {
		cache_Data_List_Lazy_Types_nil__4122162182 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_List_Lazy_Types_nil__4122162182
}

var cache_Data_List_Lazy_Types_plusList__2873873584 gopurs_runtime.Value
var once_Data_List_Lazy_Types_plusList__2873873584 sync.Once
func Get_Data_List_Lazy_Types_plusList__2873873584() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_plusList__2873873584.Do(func() {
		cache_Data_List_Lazy_Types_plusList__2873873584 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_altList()
}), Get_Data_List_Lazy_Types_nil())
	})
	return cache_Data_List_Lazy_Types_plusList__2873873584
}

var cache_Data_List_Lazy_Types_plusList__3460472018 gopurs_runtime.Value
var once_Data_List_Lazy_Types_plusList__3460472018 sync.Once
func Get_Data_List_Lazy_Types_plusList__3460472018() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_plusList__3460472018.Do(func() {
		cache_Data_List_Lazy_Types_plusList__3460472018 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_altList()
}), Get_Data_List_Lazy_Types_nil())
	})
	return cache_Data_List_Lazy_Types_plusList__3460472018
}

var cache_Data_List_Lazy_Types_semigroupList__1199693447 gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList__1199693447 sync.Once
func Get_Data_List_Lazy_Types_semigroupList__1199693447() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList__1199693447.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList__1199693447 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_semigroupList__1199693447
}

var cache_Data_List_Lazy_Types_semigroupList__3612943602 gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList__3612943602 sync.Once
func Get_Data_List_Lazy_Types_semigroupList__3612943602() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList__3612943602.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList__3612943602 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_semigroupList__3612943602
}

var cache_Data_List_Lazy_Types_semigroupList__2598308723 gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList__2598308723 sync.Once
func Get_Data_List_Lazy_Types_semigroupList__2598308723() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList__2598308723.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList__2598308723 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_semigroupList__2598308723
}

var cache_Data_List_Lazy_Types_semigroupList__4136327256 gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList__4136327256 sync.Once
func Get_Data_List_Lazy_Types_semigroupList__4136327256() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList__4136327256.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList__4136327256 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_Data_List_Lazy_Types_semigroupList__4136327256
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
		cache_Data_List_Lazy_Types_traversableList__3068288903 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_List_Lazy_Types_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
})
}))
	})
	return cache_Data_List_Lazy_Types_traversableList__3068288903
}

var cache_Data_List_Lazy_Types_traversableList__2371870579 gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableList__2371870579 sync.Once
func Get_Data_List_Lazy_Types_traversableList__2371870579() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableList__2371870579.Do(func() {
		cache_Data_List_Lazy_Types_traversableList__2371870579 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_List_Lazy_Types_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
})
}))
	})
	return cache_Data_List_Lazy_Types_traversableList__2371870579
}

var cache_Data_List_Lazy_Types_traversableList__589375054 gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableList__589375054 sync.Once
func Get_Data_List_Lazy_Types_traversableList__589375054() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableList__589375054.Do(func() {
		cache_Data_List_Lazy_Types_traversableList__589375054 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_List_Lazy_Types_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
})
}))
	})
	return cache_Data_List_Lazy_Types_traversableList__589375054
}

var cache_Data_List_Lazy_Types_traversableNonEmptyList__3951740999 gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmptyList__3951740999 sync.Once
func Get_Data_List_Lazy_Types_traversableNonEmptyList__3951740999() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmptyList__3951740999.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmptyList__3951740999 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_foldableNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_3))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_traversableNonEmpty()).V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_traversableNonEmpty()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
})
})
}))
	})
	return cache_Data_List_Lazy_Types_traversableNonEmptyList__3951740999
}

var cache_Data_List_Lazy_Types_unfoldable1List__4025223016 gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1List__4025223016 sync.Once
func Get_Data_List_Lazy_Types_unfoldable1List__4025223016() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1List__4025223016.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1List__4025223016 = func() gopurs_runtime.Value {
var go__go_0_0_11 gopurs_runtime.Value
_ = go__go_0_0_11
go__go_0_0_11 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(go__go_0_0_11, f_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1.UnsafePtr).V0)
_ = __local_var_5_3
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0, __local_var_5_3})}
}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0, Get_Data_List_Lazy_Types_nil()})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
})
})
return gopurs_runtime.RecordDict1("unfoldr1", go__go_0_0_11)
}()
	})
	return cache_Data_List_Lazy_Types_unfoldable1List__4025223016
}

var cache_Data_List_Lazy_Types_unfoldableList__825189991 gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldableList__825189991 sync.Once
func Get_Data_List_Lazy_Types_unfoldableList__825189991() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldableList__825189991.Do(func() {
		cache_Data_List_Lazy_Types_unfoldableList__825189991 = func() gopurs_runtime.Value {
var go__go_0_0_12 gopurs_runtime.Value
_ = go__go_0_0_12
go__go_0_0_12 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr == nil) {
__t3 = Get_Data_List_Lazy_Types_nil()
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply2(go__go_0_0_12, f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V1)
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V0, __local_var_5_2})}
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
})
})
return gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_unfoldable1List()
}), go__go_0_0_12)
}()
	})
	return cache_Data_List_Lazy_Types_unfoldableList__825189991
}

type Constructor_Data_List_Lazy_Types_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Data_List_Lazy_Types_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
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

func Call_Data_List_Lazy_Types_step(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_eqNonEmpty(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_eq1List(), "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1).IntVal) != (0)))
})
}))
}

func Call_Data_List_Lazy_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_eq1List(), "eq1"), dictEq_0))
}

func Call_Data_List_Lazy_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Lazy_eqLazy(), Call_Data_List_Lazy_Types_eqNonEmpty(dictEq_0))
}

func Call_Data_List_Lazy_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqList1_1_0 -> gopurs_runtime.Value
eqList1_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_eq1List(), "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
_ = eqList1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqList1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_ord1List(), "compare1"), dictOrd_0))
}

func Call_Data_List_Lazy_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Lazy_ordLazy(), gopurs_runtime.Apply(Get_Data_List_Lazy_Types_ordNonEmpty(), dictOrd_0))
}

func Call_Data_List_Lazy_Types_cons(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 -> *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Str("(fromFoldable [])")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Str(((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0).StrVal())) + (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(shown_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(((shown_3.StrVal()) + (",")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), x_prime_4).StrVal()))
})
}), gopurs_runtime.Str(""), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V1).StrVal())) + ("])"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Str(__t1.StrVal())
}))
}

func Call_Data_List_Lazy_Types_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := Call_Data_List_Lazy_Types_showList(dictShow_0)
_ = __local_var_1_1
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(NonEmpty ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))
_ = __local_var_2_2
// TAST (Let): showLazy_1_0 -> *Constructor_Data_Show_Show[gopurs_runtime.Value]
showLazy_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(defer \\_ -> ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "show"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_3)).StrVal())) + (")"))
})))
_ = showLazy_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyList ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showLazy_1_0.V0), v_2).StrVal())) + (")"))
}))
}

func Call_Data_List_Lazy_Types_showStep(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList1_1_0 -> *Constructor_Data_Show_Show[gopurs_runtime.Value]
showList1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_showList(dictShow_0))
_ = showList1_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Str("Nil")
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Str((((("(") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" : ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(showList1_1_0.V0), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Str(__t1.StrVal())
}))
}

func Call_Data_List_Lazy_Types_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1})}
}))
}))
}

func Call_Data_List_Lazy_Types_cons__716923058(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__720046150(x_0_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__376540526(x_0_loop *Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]], xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__2305074921(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__891310957(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__1901546616(x_0_loop *Constructor_Data_List_Lazy_Types_Cons[int64], xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons[int64] = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__3391588829(x_0_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_cons__2134285409(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_Data_List_Lazy_Types_step__3545407802(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__4184651873(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__4057057377(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__2986341153(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__2999566881(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__2528948705(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_step__1903922273(x_0_loop gopurs_runtime.Value) *Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_0))
}

func Call_Data_List_Lazy_Types_toList__1017592434(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1})}
}))
}))
}

func Call_Data_List_Lazy_Types_toList__4101396777(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1})}
}))
}))
}


