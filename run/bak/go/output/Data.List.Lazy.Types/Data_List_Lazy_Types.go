package Data_List_Lazy_Types

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Lazy "gopurs/output/Control.Lazy"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_TraversableWithIndex "gopurs/output/Data.TraversableWithIndex"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_List gopurs_runtime.Value
var once_List sync.Once
func Get_List() gopurs_runtime.Value {
	once_List.Do(func() {
		cache_List = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_List(x_0_box)
})
	})
	return cache_List
}

var cache_Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Nil
}

var cache_Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		cache_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Cons
}

var cache_NonEmptyList gopurs_runtime.Value
var once_NonEmptyList sync.Once
func Get_NonEmptyList() gopurs_runtime.Value {
	once_NonEmptyList.Do(func() {
		cache_NonEmptyList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_NonEmptyList(x_0_box)
})
	})
	return cache_NonEmptyList
}

var cache_nil gopurs_runtime.Value
var once_nil sync.Once
func Get_nil() gopurs_runtime.Value {
	once_nil.Do(func() {
		cache_nil = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil
}

var cache_newtypeNonEmptyList gopurs_runtime.Value
var once_newtypeNonEmptyList sync.Once
func Get_newtypeNonEmptyList() gopurs_runtime.Value {
	once_newtypeNonEmptyList.Do(func() {
		cache_newtypeNonEmptyList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeNonEmptyList
}

var cache_newtypeList gopurs_runtime.Value
var once_newtypeList sync.Once
func Get_newtypeList() gopurs_runtime.Value {
	once_newtypeList.Do(func() {
		cache_newtypeList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeList
}

var cache_step gopurs_runtime.Value
var once_step sync.Once
func Get_step() gopurs_runtime.Value {
	once_step.Do(func() {
		cache_step = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step(x_0_box))}
})
	})
	return cache_step
}

var cache_semigroupList gopurs_runtime.Value
var once_semigroupList sync.Once
func Get_semigroupList() gopurs_runtime.Value {
	once_semigroupList.Do(func() {
		cache_semigroupList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList
}

var cache_monoidList gopurs_runtime.Value
var once_monoidList sync.Once
func Get_monoidList() gopurs_runtime.Value {
	once_monoidList.Do(func() {
		cache_monoidList = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupList()
}), Get_nil())
	})
	return cache_monoidList
}

var cache_lazyList gopurs_runtime.Value
var once_lazyList sync.Once
func Get_lazyList() gopurs_runtime.Value {
	once_lazyList.Do(func() {
		cache_lazyList = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(f_0, x_1))
}))
}))
	})
	return cache_lazyList
}

var cache_functorList gopurs_runtime.Value
var once_functorList sync.Once
func Get_functorList() gopurs_runtime.Value {
	once_functorList.Do(func() {
		cache_functorList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_functorList
}

var cache_functorNonEmpty gopurs_runtime.Value
var once_functorNonEmpty sync.Once
func Get_functorNonEmpty() gopurs_runtime.Value {
	once_functorNonEmpty.Do(func() {
		cache_functorNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&pkg_Data_Functor.Constructor_Functor[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
})})}
	})
	return cache_functorNonEmpty
}

var cache_functorNonEmptyList gopurs_runtime.Value
var once_functorNonEmptyList sync.Once
func Get_functorNonEmptyList() gopurs_runtime.Value {
	once_functorNonEmptyList.Do(func() {
		cache_functorNonEmptyList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_functorNonEmpty()).V0, f_0), v_1)
})
}))
	})
	return cache_functorNonEmptyList
}

var cache_eq1List gopurs_runtime.Value
var once_eq1List sync.Once
func Get_eq1List() gopurs_runtime.Value {
	once_eq1List.Do(func() {
		cache_eq1List = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_2)))}).IntVal) != (0))
})
})
}))
	})
	return cache_eq1List
}

var cache_eq1NonEmptyList gopurs_runtime.Value
var once_eq1NonEmptyList sync.Once
func Get_eq1NonEmptyList() gopurs_runtime.Value {
	once_eq1NonEmptyList.Do(func() {
		cache_eq1NonEmptyList = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(pkg_Data_NonEmpty.Get_eqNonEmpty(), Get_eq1List(), dictEq_0))
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_eq1Lazy(), "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty1_1_0)}, v_2, v1_3).IntVal) != (0))
})
})
}))
	})
	return cache_eq1NonEmptyList
}

var cache_eqList gopurs_runtime.Value
var once_eqList sync.Once
func Get_eqList() gopurs_runtime.Value {
	once_eqList.Do(func() {
		cache_eqList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqList(dictEq_0_box)
})
	})
	return cache_eqList
}

var cache_eqNonEmptyList gopurs_runtime.Value
var once_eqNonEmptyList sync.Once
func Get_eqNonEmptyList() gopurs_runtime.Value {
	once_eqNonEmptyList.Do(func() {
		cache_eqNonEmptyList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqNonEmptyList(dictEq_0_box)
})
	})
	return cache_eqNonEmptyList
}

var cache_ord1List gopurs_runtime.Value
var once_ord1List sync.Once
func Get_ord1List() gopurs_runtime.Value {
	once_ord1List.Do(func() {
		cache_ord1List = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1List()
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
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0)
_ = v2_6_2
var __t3 uint32
{
if (uint32(v2_6_2.IntVal) == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_2)))}).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_ord1List
}

var cache_ordNonEmpty gopurs_runtime.Value
var once_ordNonEmpty sync.Once
func Get_ordNonEmpty() gopurs_runtime.Value {
	once_ordNonEmpty.Do(func() {
		cache_ordNonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_ordNonEmpty(), Get_ord1List())
	})
	return cache_ordNonEmpty
}

var cache_ord1NonEmptyList gopurs_runtime.Value
var once_ord1NonEmptyList sync.Once
func Get_ord1NonEmptyList() gopurs_runtime.Value {
	once_ord1NonEmptyList.Do(func() {
		cache_ord1NonEmptyList = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1NonEmptyList()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd_0))
_ = ordNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_ord1Lazy(), "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordNonEmpty1_1_0)}, v_2, v1_3).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_ord1NonEmptyList
}

var cache_ordList gopurs_runtime.Value
var once_ordList sync.Once
func Get_ordList() gopurs_runtime.Value {
	once_ordList.Do(func() {
		cache_ordList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordList(dictOrd_0_box)
})
	})
	return cache_ordList
}

var cache_ordNonEmptyList gopurs_runtime.Value
var once_ordNonEmptyList sync.Once
func Get_ordNonEmptyList() gopurs_runtime.Value {
	once_ordNonEmptyList.Do(func() {
		cache_ordNonEmptyList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordNonEmptyList(dictOrd_0_box)
})
	})
	return cache_ordNonEmptyList
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons(x_0_box, xs_1_box)
})
	})
	return cache_cons
}

var cache_foldableList gopurs_runtime.Value
var once_foldableList sync.Once
func Get_foldableList() gopurs_runtime.Value {
	once_foldableList.Do(func() {
		cache_foldableList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
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
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
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
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList
}

var cache_foldableNonEmpty gopurs_runtime.Value
var once_foldableNonEmpty sync.Once
func Get_foldableNonEmpty() gopurs_runtime.Value {
	once_foldableNonEmpty.Do(func() {
		cache_foldableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldableNonEmpty(), Get_foldableList())))}
	})
	return cache_foldableNonEmpty
}

var cache_extendList gopurs_runtime.Value
var once_extendList sync.Once
func Get_extendList() gopurs_runtime.Value {
	once_extendList.Do(func() {
		cache_extendList = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
_ = v_2_0
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr == nil) {
__t5 = Get_nil()
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr != nil) {
__local_var_3_1 := gopurs_runtime.Apply(f_0, l_1)
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
acc_prime_6_3 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, gopurs_runtime.RecordGet(v_5, "acc")})}
}))
_ = acc_prime_6_3
__local_var_7_4 := gopurs_runtime.Apply(f_0, acc_prime_6_3)
_ = __local_var_7_4
return gopurs_runtime.RecordDict2("acc", "val", acc_prime_6_3, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_7_4, gopurs_runtime.RecordGet(v_5, "val")})}
})))
})
}), gopurs_runtime.RecordDict2("acc", "val", Get_nil(), Get_nil()), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V1), "val")
_ = __local_var_4_2
__t5 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_4_2})}
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
	return cache_extendList
}

var cache_extendNonEmptyList gopurs_runtime.Value
var once_extendNonEmptyList sync.Once
func Get_extendNonEmptyList() gopurs_runtime.Value {
	once_extendNonEmptyList.Do(func() {
		cache_extendNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1).UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.RecordGet(v1_5, "acc")
_ = __local_var_6_1
__local_var_7_2 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, __local_var_6_1})}
})))
_ = __local_var_7_2
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, __local_var_6_1})}
})), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_7_2, gopurs_runtime.RecordGet(v1_5, "val")})}
})))
})
}), gopurs_runtime.RecordDict2("acc", "val", Get_nil(), Get_nil()), __local_var_2_0), "val")})}
}))
})
}))
	})
	return cache_extendNonEmptyList
}

var cache_foldableNonEmptyList gopurs_runtime.Value
var once_foldableNonEmptyList sync.Once
func Get_foldableNonEmptyList() gopurs_runtime.Value {
	once_foldableNonEmptyList.Do(func() {
		cache_foldableNonEmptyList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V1, f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V2, f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}))
	})
	return cache_foldableNonEmptyList
}

var cache_showList gopurs_runtime.Value
var once_showList sync.Once
func Get_showList() gopurs_runtime.Value {
	once_showList.Do(func() {
		cache_showList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showList(dictShow_0_box)
})
	})
	return cache_showList
}

var cache_showNonEmptyList gopurs_runtime.Value
var once_showNonEmptyList sync.Once
func Get_showNonEmptyList() gopurs_runtime.Value {
	once_showNonEmptyList.Do(func() {
		cache_showNonEmptyList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showNonEmptyList(dictShow_0_box)
})
	})
	return cache_showNonEmptyList
}

var cache_showStep gopurs_runtime.Value
var once_showStep sync.Once
func Get_showStep() gopurs_runtime.Value {
	once_showStep.Do(func() {
		cache_showStep = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showStep(dictShow_0_box)
})
	})
	return cache_showStep
}

var cache_foldableWithIndexList gopurs_runtime.Value
var once_foldableWithIndexList sync.Once
func Get_foldableWithIndexList() gopurs_runtime.Value {
	once_foldableWithIndexList.Do(func() {
		cache_foldableWithIndexList = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_5)
_ = __local_var_6_2
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
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
__local_var_4_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_4_6.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_8 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
__local_var_5_9 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_5_9.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, __local_var_4_8})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_nil()})}, xs_2))
_ = v_3_7
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
__local_var_6_11 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1).UnsafePtr).V1
})
})
}))
	})
	return cache_foldableWithIndexList
}

var cache_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_foldableWithIndexNonEmpty sync.Once
func Get_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_foldableWithIndexNonEmpty.Do(func() {
		cache_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldableWithIndexNonEmpty(), Get_foldableWithIndexList())))}
	})
	return cache_foldableWithIndexNonEmpty
}

var cache_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_foldableWithIndexNonEmptyList sync.Once
func Get_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyList.Do(func() {
		cache_foldableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_0
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(__local_var_3_0, x_4))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_1
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V2, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V3, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_2, x_4))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}))
	})
	return cache_foldableWithIndexNonEmptyList
}

var cache_functorWithIndexList gopurs_runtime.Value
var once_functorWithIndexList sync.Once
func Get_functorWithIndexList() gopurs_runtime.Value {
	once_functorWithIndexList.Do(func() {
		cache_functorWithIndexList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2)
_ = __local_var_4_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_0, acc_3})}
}))
})
})
}), Get_nil())
}))
	})
	return cache_functorWithIndexList
}

var cache_functorWithIndex gopurs_runtime.Value
var once_functorWithIndex sync.Once
func Get_functorWithIndex() gopurs_runtime.Value {
	once_functorWithIndex.Do(func() {
		cache_functorWithIndex = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorWithIndexList(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_0_1
functorNonEmpty1_0_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "map"), f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_0_0
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorWithIndexList(), "mapWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_3})})
}), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
}))))}
}()
	})
	return cache_functorWithIndex
}

var cache_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_functorWithIndexNonEmptyList sync.Once
func Get_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyList.Do(func() {
		cache_functorWithIndexNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_functorWithIndex()).V1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_0, x_4))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1)))})))}
}))
})
}))
	})
	return cache_functorWithIndexNonEmptyList
}

var cache_toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		cache_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toList(v_0_box)
})
	})
	return cache_toList
}

var cache_semigroupNonEmptyList gopurs_runtime.Value
var once_semigroupNonEmptyList sync.Once
func Get_semigroupNonEmptyList() gopurs_runtime.Value {
	once_semigroupNonEmptyList.Do(func() {
		cache_semigroupNonEmptyList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_2_0
__local_var_3_1 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0
_ = __local_var_3_1
__local_var_4_2 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V1
_ = __local_var_4_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), __local_var_4_2, Call_toList(as_prime_1))})}
}))
})
}))
	})
	return cache_semigroupNonEmptyList
}

var cache_traversableList gopurs_runtime.Value
var once_traversableList sync.Once
func Get_traversableList() gopurs_runtime.Value {
	once_traversableList.Do(func() {
		cache_traversableList = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_nil()))
})
}))
	})
	return cache_traversableList
}

var cache_traversableNonEmpty gopurs_runtime.Value
var once_traversableNonEmpty sync.Once
func Get_traversableNonEmpty() gopurs_runtime.Value {
	once_traversableNonEmpty.Do(func() {
		cache_traversableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableNonEmpty(), Get_traversableList())))}
	})
	return cache_traversableNonEmpty
}

var cache_traversableNonEmptyList gopurs_runtime.Value
var once_traversableNonEmptyList sync.Once
func Get_traversableNonEmptyList() gopurs_runtime.Value {
	once_traversableNonEmptyList.Do(func() {
		cache_traversableNonEmptyList = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_3))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableNonEmpty()).V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableNonEmpty()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_3)))}))
})
})
}))
	})
	return cache_traversableNonEmptyList
}

var cache_traversableWithIndexList gopurs_runtime.Value
var once_traversableWithIndexList sync.Once
func Get_traversableWithIndexList() gopurs_runtime.Value {
	once_traversableWithIndexList.Do(func() {
		cache_traversableWithIndexList = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_cons(), gopurs_runtime.Apply2(f_3, gopurs_runtime.Int(i_4.IntVal), a_5)), b_6)
})
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_nil()))
})
}))
	})
	return cache_traversableWithIndexList
}

var cache_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_traversableWithIndexNonEmpty sync.Once
func Get_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_traversableWithIndexNonEmpty.Do(func() {
		cache_traversableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableWithIndexNonEmpty(), Get_traversableWithIndexList())))}
	})
	return cache_traversableWithIndexNonEmpty
}

var cache_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_traversableWithIndexNonEmptyList sync.Once
func Get_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_traversableWithIndexNonEmptyList.Do(func() {
		cache_traversableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_4_1
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableWithIndexNonEmpty()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(__local_var_4_1, x_5))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_3)))}))
})
})
}))
	})
	return cache_traversableWithIndexNonEmptyList
}

var cache_unfoldable1List gopurs_runtime.Value
var once_unfoldable1List sync.Once
func Get_unfoldable1List() gopurs_runtime.Value {
	once_unfoldable1List.Do(func() {
		cache_unfoldable1List = func() gopurs_runtime.Value {
var go__go_0_0_3 gopurs_runtime.Value
_ = go__go_0_0_3
go__go_0_0_3 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
__local_var_5_3 := gopurs_runtime.Apply2(go__go_0_0_3, f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1.UnsafePtr).V0)
_ = __local_var_5_3
__t5 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0, __local_var_5_3})}
}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0, Get_nil()})}
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
	return cache_unfoldable1List
}

var cache_unfoldableList gopurs_runtime.Value
var once_unfoldableList sync.Once
func Get_unfoldableList() gopurs_runtime.Value {
	once_unfoldableList.Do(func() {
		cache_unfoldableList = func() gopurs_runtime.Value {
var go__go_0_0_4 gopurs_runtime.Value
_ = go__go_0_0_4
go__go_0_0_4 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr == nil) {
__t3 = Get_nil()
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr != nil) {
__local_var_5_2 := gopurs_runtime.Apply2(go__go_0_0_4, f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V1)
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V0, __local_var_5_2})}
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
return Get_unfoldable1List()
}), go__go_0_0_4)
}()
	})
	return cache_unfoldableList
}

var cache_unfoldable1NonEmpty gopurs_runtime.Value
var once_unfoldable1NonEmpty sync.Once
func Get_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_unfoldable1NonEmpty.Do(func() {
		cache_unfoldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_unfoldable1NonEmpty(), Get_unfoldableList())))}
	})
	return cache_unfoldable1NonEmpty
}

var cache_unfoldable1NonEmptyList gopurs_runtime.Value
var once_unfoldable1NonEmptyList sync.Once
func Get_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_unfoldable1NonEmptyList.Do(func() {
		cache_unfoldable1NonEmptyList = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_unfoldable1NonEmpty()).V0, f_0, b_1)))}
}))
})
}))
	})
	return cache_unfoldable1NonEmptyList
}

var cache_comonadNonEmptyList gopurs_runtime.Value
var once_comonadNonEmptyList sync.Once
func Get_comonadNonEmptyList() gopurs_runtime.Value {
	once_comonadNonEmptyList.Do(func() {
		cache_comonadNonEmptyList = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V0
}))
	})
	return cache_comonadNonEmptyList
}

var cache_monadList gopurs_runtime.Value
var once_monadList sync.Once
func Get_monadList() gopurs_runtime.Value {
	once_monadList.Do(func() {
		cache_monadList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindList()
}))
	})
	return cache_monadList
}

var cache_bindList gopurs_runtime.Value
var once_bindList sync.Once
func Get_bindList() gopurs_runtime.Value {
	once_bindList.Do(func() {
		cache_bindList = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_bindList
}

var cache_applyList gopurs_runtime.Value
var once_applyList sync.Once
func Get_applyList() gopurs_runtime.Value {
	once_applyList.Do(func() {
		cache_applyList = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyList
}

var cache_applicativeList gopurs_runtime.Value
var once_applicativeList sync.Once
func Get_applicativeList() gopurs_runtime.Value {
	once_applicativeList.Do(func() {
		cache_applicativeList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_0, Get_nil()})}
}))
}))
	})
	return cache_applicativeList
}

var cache_applyNonEmptyList gopurs_runtime.Value
var once_applyNonEmptyList sync.Once
func Get_applyNonEmptyList() gopurs_runtime.Value {
	once_applyNonEmptyList.Do(func() {
		cache_applyNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
v2_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v1_1)
_ = v2_2_0
v3_3_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v3_3_1
__local_var_4_2 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0
_ = __local_var_4_2
__local_var_5_3 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3_1.UnsafePtr).V0
_ = __local_var_6_4
__local_var_7_5 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3_1.UnsafePtr).V1
_ = __local_var_7_5
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), __local_var_7_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_2, Get_nil()})}
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_4, __local_var_7_5})}
})), __local_var_5_3))})}
}))
})
}))
	})
	return cache_applyNonEmptyList
}

var cache_bindNonEmptyList gopurs_runtime.Value
var once_bindNonEmptyList sync.Once
func Get_bindNonEmptyList() gopurs_runtime.Value {
	once_bindNonEmptyList.Do(func() {
		cache_bindNonEmptyList = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_2_0
__local_var_3_1 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V1
_ = __local_var_3_1
v2_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0)))
_ = v2_4_2
__local_var_5_3 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v2_4_2)}.UnsafePtr).V0
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v2_4_2)}.UnsafePtr).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), __local_var_6_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), __local_var_3_1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toList(gopurs_runtime.Apply(f_1, x_8))
})))})}
}))
})
}))
	})
	return cache_bindNonEmptyList
}

var cache_altNonEmptyList gopurs_runtime.Value
var once_altNonEmptyList sync.Once
func Get_altNonEmptyList() gopurs_runtime.Value {
	once_altNonEmptyList.Do(func() {
		cache_altNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.RecordGet(Get_semigroupNonEmptyList(), "append"))
	})
	return cache_altNonEmptyList
}

var cache_altList gopurs_runtime.Value
var once_altList sync.Once
func Get_altList() gopurs_runtime.Value {
	once_altList.Do(func() {
		cache_altList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.RecordGet(Get_semigroupList(), "append"))
	})
	return cache_altList
}

var cache_plusList gopurs_runtime.Value
var once_plusList sync.Once
func Get_plusList() gopurs_runtime.Value {
	once_plusList.Do(func() {
		cache_plusList = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altList()
}), Get_nil())
	})
	return cache_plusList
}

var cache_alternativeList gopurs_runtime.Value
var once_alternativeList sync.Once
func Get_alternativeList() gopurs_runtime.Value {
	once_alternativeList.Do(func() {
		cache_alternativeList = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusList()
}))
	})
	return cache_alternativeList
}

var cache_monadPlusList gopurs_runtime.Value
var once_monadPlusList sync.Once
func Get_monadPlusList() gopurs_runtime.Value {
	once_monadPlusList.Do(func() {
		cache_monadPlusList = gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_alternativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadList()
}))
	})
	return cache_monadPlusList
}

var cache_applicativeNonEmptyList gopurs_runtime.Value
var once_applicativeNonEmptyList sync.Once
func Get_applicativeNonEmptyList() gopurs_runtime.Value {
	once_applicativeNonEmptyList.Do(func() {
		cache_applicativeNonEmptyList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
}))
}))
	})
	return cache_applicativeNonEmptyList
}

var cache_monadNonEmptyList gopurs_runtime.Value
var once_monadNonEmptyList sync.Once
func Get_monadNonEmptyList() gopurs_runtime.Value {
	once_monadNonEmptyList.Do(func() {
		cache_monadNonEmptyList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindNonEmptyList()
}))
	})
	return cache_monadNonEmptyList
}

var cache_pure__189931222 gopurs_runtime.Value
var once_pure__189931222 sync.Once
func Get_pure__189931222() gopurs_runtime.Value {
	once_pure__189931222.Do(func() {
		cache_pure__189931222 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__189931222(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__189931222
}

var cache_pure__3236307030 gopurs_runtime.Value
var once_pure__3236307030 sync.Once
func Get_pure__3236307030() gopurs_runtime.Value {
	once_pure__3236307030.Do(func() {
		cache_pure__3236307030 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3236307030(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3236307030
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__355615152 gopurs_runtime.Value
var once_pure__355615152 sync.Once
func Get_pure__355615152() gopurs_runtime.Value {
	once_pure__355615152.Do(func() {
		cache_pure__355615152 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__355615152(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__355615152
}

var cache_apply__4203183626 gopurs_runtime.Value
var once_apply__4203183626 sync.Once
func Get_apply__4203183626() gopurs_runtime.Value {
	once_apply__4203183626.Do(func() {
		cache_apply__4203183626 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__4203183626(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__4203183626
}

var cache_apply__2962221386 gopurs_runtime.Value
var once_apply__2962221386 sync.Once
func Get_apply__2962221386() gopurs_runtime.Value {
	once_apply__2962221386.Do(func() {
		cache_apply__2962221386 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2962221386(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2962221386
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__1851858028 gopurs_runtime.Value
var once_apply__1851858028 sync.Once
func Get_apply__1851858028() gopurs_runtime.Value {
	once_apply__1851858028.Do(func() {
		cache_apply__1851858028 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__1851858028(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__1851858028
}

var cache_apply__2140510474 gopurs_runtime.Value
var once_apply__2140510474 sync.Once
func Get_apply__2140510474() gopurs_runtime.Value {
	once_apply__2140510474.Do(func() {
		cache_apply__2140510474 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2140510474(f_0_box, x_1_box)
})
	})
	return cache_apply__2140510474
}

var cache_apply__3620326986 gopurs_runtime.Value
var once_apply__3620326986 sync.Once
func Get_apply__3620326986() gopurs_runtime.Value {
	once_apply__3620326986.Do(func() {
		cache_apply__3620326986 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__3620326986(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_apply__3620326986
}

var cache_bind__3512795567 gopurs_runtime.Value
var once_bind__3512795567 sync.Once
func Get_bind__3512795567() gopurs_runtime.Value {
	once_bind__3512795567.Do(func() {
		cache_bind__3512795567 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3512795567(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3512795567
}

var cache_bind__3781138863 gopurs_runtime.Value
var once_bind__3781138863 sync.Once
func Get_bind__3781138863() gopurs_runtime.Value {
	once_bind__3781138863.Do(func() {
		cache_bind__3781138863 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3781138863(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3781138863
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__684765761 gopurs_runtime.Value
var once_bind__684765761 sync.Once
func Get_bind__684765761() gopurs_runtime.Value {
	once_bind__684765761.Do(func() {
		cache_bind__684765761 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__684765761(xs_0_box, f_1_box)
})
	})
	return cache_bind__684765761
}

var cache_bind__4082241 gopurs_runtime.Value
var once_bind__4082241 sync.Once
func Get_bind__4082241() gopurs_runtime.Value {
	once_bind__4082241.Do(func() {
		cache_bind__4082241 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__4082241(xs_0_box, f_1_box)
})
	})
	return cache_bind__4082241
}

var cache_defer__3967925939 gopurs_runtime.Value
var once_defer__3967925939 sync.Once
func Get_defer__3967925939() gopurs_runtime.Value {
	once_defer__3967925939.Do(func() {
		cache_defer__3967925939 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__3967925939(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_defer__3967925939
}

var cache_defer__2590380358 gopurs_runtime.Value
var once_defer__2590380358 sync.Once
func Get_defer__2590380358() gopurs_runtime.Value {
	once_defer__2590380358.Do(func() {
		cache_defer__2590380358 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__2590380358(f_0_box)
})
	})
	return cache_defer__2590380358
}

var cache_compose__1987728071 gopurs_runtime.Value
var once_compose__1987728071 sync.Once
func Get_compose__1987728071() gopurs_runtime.Value {
	once_compose__1987728071.Do(func() {
		cache_compose__1987728071 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1987728071(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__1987728071
}

var cache_compose__346034828 gopurs_runtime.Value
var once_compose__346034828 sync.Once
func Get_compose__346034828() gopurs_runtime.Value {
	once_compose__346034828.Do(func() {
		cache_compose__346034828 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__346034828(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__346034828
}

var cache_compose__1555187646 gopurs_runtime.Value
var once_compose__1555187646 sync.Once
func Get_compose__1555187646() gopurs_runtime.Value {
	once_compose__1555187646.Do(func() {
		cache_compose__1555187646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1555187646(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__1555187646
}

var cache_compose__4254807102 gopurs_runtime.Value
var once_compose__4254807102 sync.Once
func Get_compose__4254807102() gopurs_runtime.Value {
	once_compose__4254807102.Do(func() {
		cache_compose__4254807102 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__4254807102(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__4254807102
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq1__1773593252 gopurs_runtime.Value
var once_eq1__1773593252 sync.Once
func Get_eq1__1773593252() gopurs_runtime.Value {
	once_eq1__1773593252.Do(func() {
		cache_eq1__1773593252 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1__1773593252(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq1__1773593252
}

var cache_eq1__2184765036 gopurs_runtime.Value
var once_eq1__2184765036 sync.Once
func Get_eq1__2184765036() gopurs_runtime.Value {
	once_eq1__2184765036.Do(func() {
		cache_eq1__2184765036 = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1__2184765036(dictEq_0_box)
})
	})
	return cache_eq1__2184765036
}

var cache_eq1__3587165073 gopurs_runtime.Value
var once_eq1__3587165073 sync.Once
func Get_eq1__3587165073() gopurs_runtime.Value {
	once_eq1__3587165073.Do(func() {
		cache_eq1__3587165073 = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1__3587165073(dictEq_0_box)
})
	})
	return cache_eq1__3587165073
}

var cache_foldMap__4098395794 gopurs_runtime.Value
var once_foldMap__4098395794 sync.Once
func Get_foldMap__4098395794() gopurs_runtime.Value {
	once_foldMap__4098395794.Do(func() {
		cache_foldMap__4098395794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__4098395794(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__4098395794
}

var cache_foldMap__3562626100 gopurs_runtime.Value
var once_foldMap__3562626100 sync.Once
func Get_foldMap__3562626100() gopurs_runtime.Value {
	once_foldMap__3562626100.Do(func() {
		cache_foldMap__3562626100 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__3562626100(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_foldMap__3562626100
}

var cache_foldMap__2350611220 gopurs_runtime.Value
var once_foldMap__2350611220 sync.Once
func Get_foldMap__2350611220() gopurs_runtime.Value {
	once_foldMap__2350611220.Do(func() {
		cache_foldMap__2350611220 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__2350611220(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_foldMap__2350611220
}

var cache_foldl__1422885860 gopurs_runtime.Value
var once_foldl__1422885860 sync.Once
func Get_foldl__1422885860() gopurs_runtime.Value {
	once_foldl__1422885860.Do(func() {
		cache_foldl__1422885860 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1422885860(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__1422885860
}

var cache_foldl__94807652 gopurs_runtime.Value
var once_foldl__94807652 sync.Once
func Get_foldl__94807652() gopurs_runtime.Value {
	once_foldl__94807652.Do(func() {
		cache_foldl__94807652 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__94807652(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__94807652
}

var cache_foldl__506543652 gopurs_runtime.Value
var once_foldl__506543652 sync.Once
func Get_foldl__506543652() gopurs_runtime.Value {
	once_foldl__506543652.Do(func() {
		cache_foldl__506543652 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__506543652(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__506543652
}

var cache_foldl__267332164 gopurs_runtime.Value
var once_foldl__267332164 sync.Once
func Get_foldl__267332164() gopurs_runtime.Value {
	once_foldl__267332164.Do(func() {
		cache_foldl__267332164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__267332164(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__267332164
}

var cache_foldl__3131354468 gopurs_runtime.Value
var once_foldl__3131354468 sync.Once
func Get_foldl__3131354468() gopurs_runtime.Value {
	once_foldl__3131354468.Do(func() {
		cache_foldl__3131354468 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3131354468(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__3131354468
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldl__524683195 gopurs_runtime.Value
var once_foldl__524683195 sync.Once
func Get_foldl__524683195() gopurs_runtime.Value {
	once_foldl__524683195.Do(func() {
		cache_foldl__524683195 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__524683195(op_0_box)
})
	})
	return cache_foldl__524683195
}

var cache_foldl__3306117403 gopurs_runtime.Value
var once_foldl__3306117403 sync.Once
func Get_foldl__3306117403() gopurs_runtime.Value {
	once_foldl__3306117403.Do(func() {
		cache_foldl__3306117403 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3306117403(op_0_box)
})
	})
	return cache_foldl__3306117403
}

var cache_foldl__3737487037 gopurs_runtime.Value
var once_foldl__3737487037 sync.Once
func Get_foldl__3737487037() gopurs_runtime.Value {
	once_foldl__3737487037.Do(func() {
		cache_foldl__3737487037 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3737487037(op_0_box)
})
	})
	return cache_foldl__3737487037
}

var cache_foldl__1985071933 gopurs_runtime.Value
var once_foldl__1985071933 sync.Once
func Get_foldl__1985071933() gopurs_runtime.Value {
	once_foldl__1985071933.Do(func() {
		cache_foldl__1985071933 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1985071933(op_0_box)
})
	})
	return cache_foldl__1985071933
}

var cache_foldl__536153533 gopurs_runtime.Value
var once_foldl__536153533 sync.Once
func Get_foldl__536153533() gopurs_runtime.Value {
	once_foldl__536153533.Do(func() {
		cache_foldl__536153533 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__536153533(op_0_box)
})
	})
	return cache_foldl__536153533
}

var cache_foldl__170252797 gopurs_runtime.Value
var once_foldl__170252797 sync.Once
func Get_foldl__170252797() gopurs_runtime.Value {
	once_foldl__170252797.Do(func() {
		cache_foldl__170252797 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__170252797(op_0_box)
})
	})
	return cache_foldl__170252797
}

var cache_foldl__2188030845 gopurs_runtime.Value
var once_foldl__2188030845 sync.Once
func Get_foldl__2188030845() gopurs_runtime.Value {
	once_foldl__2188030845.Do(func() {
		cache_foldl__2188030845 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2188030845(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldl__2188030845
}

var cache_foldl__1444272061 gopurs_runtime.Value
var once_foldl__1444272061 sync.Once
func Get_foldl__1444272061() gopurs_runtime.Value {
	once_foldl__1444272061.Do(func() {
		cache_foldl__1444272061 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1444272061(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldl__1444272061
}

var cache_foldr__2111289130 gopurs_runtime.Value
var once_foldr__2111289130 sync.Once
func Get_foldr__2111289130() gopurs_runtime.Value {
	once_foldr__2111289130.Do(func() {
		cache_foldr__2111289130 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2111289130(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2111289130
}

var cache_foldr__926146538 gopurs_runtime.Value
var once_foldr__926146538 sync.Once
func Get_foldr__926146538() gopurs_runtime.Value {
	once_foldr__926146538.Do(func() {
		cache_foldr__926146538 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__926146538(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__926146538
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__1985071933 gopurs_runtime.Value
var once_foldr__1985071933 sync.Once
func Get_foldr__1985071933() gopurs_runtime.Value {
	once_foldr__1985071933.Do(func() {
		cache_foldr__1985071933 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__1985071933(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_foldr__1985071933
}

var cache_foldr__2389967549 gopurs_runtime.Value
var once_foldr__2389967549 sync.Once
func Get_foldr__2389967549() gopurs_runtime.Value {
	once_foldr__2389967549.Do(func() {
		cache_foldr__2389967549 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2389967549(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_foldr__2389967549
}

var cache_foldr__1278383325 gopurs_runtime.Value
var once_foldr__1278383325 sync.Once
func Get_foldr__1278383325() gopurs_runtime.Value {
	once_foldr__1278383325.Do(func() {
		cache_foldr__1278383325 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__1278383325(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_foldr__1278383325
}

var cache_foldr__2188030845 gopurs_runtime.Value
var once_foldr__2188030845 sync.Once
func Get_foldr__2188030845() gopurs_runtime.Value {
	once_foldr__2188030845.Do(func() {
		cache_foldr__2188030845 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2188030845(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldr__2188030845
}

var cache_foldr__3749276701 gopurs_runtime.Value
var once_foldr__3749276701 sync.Once
func Get_foldr__3749276701() gopurs_runtime.Value {
	once_foldr__3749276701.Do(func() {
		cache_foldr__3749276701 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3749276701(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldr__3749276701
}

var cache_foldMapWithIndex__2292551140 gopurs_runtime.Value
var once_foldMapWithIndex__2292551140 sync.Once
func Get_foldMapWithIndex__2292551140() gopurs_runtime.Value {
	once_foldMapWithIndex__2292551140.Do(func() {
		cache_foldMapWithIndex__2292551140 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndex__2292551140(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMapWithIndex__2292551140
}

var cache_foldMapWithIndex__1722031522 gopurs_runtime.Value
var once_foldMapWithIndex__1722031522 sync.Once
func Get_foldMapWithIndex__1722031522() gopurs_runtime.Value {
	once_foldMapWithIndex__1722031522.Do(func() {
		cache_foldMapWithIndex__1722031522 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndex__1722031522(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_foldMapWithIndex__1722031522
}

var cache_foldMapWithIndex__852526914 gopurs_runtime.Value
var once_foldMapWithIndex__852526914 sync.Once
func Get_foldMapWithIndex__852526914() gopurs_runtime.Value {
	once_foldMapWithIndex__852526914.Do(func() {
		cache_foldMapWithIndex__852526914 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndex__852526914(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_foldMapWithIndex__852526914
}

var cache_foldlWithIndex__2972270123 gopurs_runtime.Value
var once_foldlWithIndex__2972270123 sync.Once
func Get_foldlWithIndex__2972270123() gopurs_runtime.Value {
	once_foldlWithIndex__2972270123.Do(func() {
		cache_foldlWithIndex__2972270123 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__2972270123(f_0_box, acc_1_box)
})
	})
	return cache_foldlWithIndex__2972270123
}

var cache_foldlWithIndex__234438827 gopurs_runtime.Value
var once_foldlWithIndex__234438827 sync.Once
func Get_foldlWithIndex__234438827() gopurs_runtime.Value {
	once_foldlWithIndex__234438827.Do(func() {
		cache_foldlWithIndex__234438827 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__234438827(f_0_box, acc_1_box)
})
	})
	return cache_foldlWithIndex__234438827
}

var cache_foldlWithIndex__2986161357 gopurs_runtime.Value
var once_foldlWithIndex__2986161357 sync.Once
func Get_foldlWithIndex__2986161357() gopurs_runtime.Value {
	once_foldlWithIndex__2986161357.Do(func() {
		cache_foldlWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldlWithIndex__2986161357
}

var cache_foldlWithIndex__3610348555 gopurs_runtime.Value
var once_foldlWithIndex__3610348555 sync.Once
func Get_foldlWithIndex__3610348555() gopurs_runtime.Value {
	once_foldlWithIndex__3610348555.Do(func() {
		cache_foldlWithIndex__3610348555 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__3610348555(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldlWithIndex__3610348555
}

var cache_foldlWithIndex__446277963 gopurs_runtime.Value
var once_foldlWithIndex__446277963 sync.Once
func Get_foldlWithIndex__446277963() gopurs_runtime.Value {
	once_foldlWithIndex__446277963.Do(func() {
		cache_foldlWithIndex__446277963 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__446277963(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldlWithIndex__446277963
}

var cache_foldrWithIndex__2972270123 gopurs_runtime.Value
var once_foldrWithIndex__2972270123 sync.Once
func Get_foldrWithIndex__2972270123() gopurs_runtime.Value {
	once_foldrWithIndex__2972270123.Do(func() {
		cache_foldrWithIndex__2972270123 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__2972270123(f_0_box, b_1_box, xs_2_box)
})
	})
	return cache_foldrWithIndex__2972270123
}

var cache_foldrWithIndex__3735894283 gopurs_runtime.Value
var once_foldrWithIndex__3735894283 sync.Once
func Get_foldrWithIndex__3735894283() gopurs_runtime.Value {
	once_foldrWithIndex__3735894283.Do(func() {
		cache_foldrWithIndex__3735894283 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__3735894283(f_0_box, b_1_box, xs_2_box)
})
	})
	return cache_foldrWithIndex__3735894283
}

var cache_foldrWithIndex__500807083 gopurs_runtime.Value
var once_foldrWithIndex__500807083 sync.Once
func Get_foldrWithIndex__500807083() gopurs_runtime.Value {
	once_foldrWithIndex__500807083.Do(func() {
		cache_foldrWithIndex__500807083 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__500807083(f_0_box, b_1_box, xs_2_box)
})
	})
	return cache_foldrWithIndex__500807083
}

var cache_foldrWithIndex__2986161357 gopurs_runtime.Value
var once_foldrWithIndex__2986161357 sync.Once
func Get_foldrWithIndex__2986161357() gopurs_runtime.Value {
	once_foldrWithIndex__2986161357.Do(func() {
		cache_foldrWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldrWithIndex__2986161357
}

var cache_foldrWithIndex__3610348555 gopurs_runtime.Value
var once_foldrWithIndex__3610348555 sync.Once
func Get_foldrWithIndex__3610348555() gopurs_runtime.Value {
	once_foldrWithIndex__3610348555.Do(func() {
		cache_foldrWithIndex__3610348555 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__3610348555(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldrWithIndex__3610348555
}

var cache_foldrWithIndex__63302635 gopurs_runtime.Value
var once_foldrWithIndex__63302635 sync.Once
func Get_foldrWithIndex__63302635() gopurs_runtime.Value {
	once_foldrWithIndex__63302635.Do(func() {
		cache_foldrWithIndex__63302635 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__63302635(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldrWithIndex__63302635
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__3658931456 gopurs_runtime.Value
var once_flip__3658931456 sync.Once
func Get_flip__3658931456() gopurs_runtime.Value {
	once_flip__3658931456.Do(func() {
		cache_flip__3658931456 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3658931456(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3658931456
}

var cache_flip__3019832928 gopurs_runtime.Value
var once_flip__3019832928 sync.Once
func Get_flip__3019832928() gopurs_runtime.Value {
	once_flip__3019832928.Do(func() {
		cache_flip__3019832928 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3019832928(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3019832928
}

var cache_flip__2175652032 gopurs_runtime.Value
var once_flip__2175652032 sync.Once
func Get_flip__2175652032() gopurs_runtime.Value {
	once_flip__2175652032.Do(func() {
		cache_flip__2175652032 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2175652032(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2175652032
}

var cache_flip__848188896 gopurs_runtime.Value
var once_flip__848188896 sync.Once
func Get_flip__848188896() gopurs_runtime.Value {
	once_flip__848188896.Do(func() {
		cache_flip__848188896 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__848188896(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__848188896
}

var cache_flip__913470112 gopurs_runtime.Value
var once_flip__913470112 sync.Once
func Get_flip__913470112() gopurs_runtime.Value {
	once_flip__913470112.Do(func() {
		cache_flip__913470112 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__913470112(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__913470112
}

var cache_flip__4017289888 gopurs_runtime.Value
var once_flip__4017289888 sync.Once
func Get_flip__4017289888() gopurs_runtime.Value {
	once_flip__4017289888.Do(func() {
		cache_flip__4017289888 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__4017289888(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__4017289888
}

var cache_flip__3036739744 gopurs_runtime.Value
var once_flip__3036739744 sync.Once
func Get_flip__3036739744() gopurs_runtime.Value {
	once_flip__3036739744.Do(func() {
		cache_flip__3036739744 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3036739744(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3036739744
}

var cache_flip__3539834656 gopurs_runtime.Value
var once_flip__3539834656 sync.Once
func Get_flip__3539834656() gopurs_runtime.Value {
	once_flip__3539834656.Do(func() {
		cache_flip__3539834656 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3539834656(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3539834656
}

var cache_map__4040535013 gopurs_runtime.Value
var once_map__4040535013 sync.Once
func Get_map__4040535013() gopurs_runtime.Value {
	once_map__4040535013.Do(func() {
		cache_map__4040535013 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__4040535013(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__4040535013
}

var cache_map__2665381605 gopurs_runtime.Value
var once_map__2665381605 sync.Once
func Get_map__2665381605() gopurs_runtime.Value {
	once_map__2665381605.Do(func() {
		cache_map__2665381605 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2665381605(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2665381605
}

var cache_map__4285761829 gopurs_runtime.Value
var once_map__4285761829 sync.Once
func Get_map__4285761829() gopurs_runtime.Value {
	once_map__4285761829.Do(func() {
		cache_map__4285761829 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__4285761829(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__4285761829
}

var cache_map__3149795237 gopurs_runtime.Value
var once_map__3149795237 sync.Once
func Get_map__3149795237() gopurs_runtime.Value {
	once_map__3149795237.Do(func() {
		cache_map__3149795237 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3149795237(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3149795237
}

var cache_map__1542634789 gopurs_runtime.Value
var once_map__1542634789 sync.Once
func Get_map__1542634789() gopurs_runtime.Value {
	once_map__1542634789.Do(func() {
		cache_map__1542634789 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1542634789(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1542634789
}

var cache_map__2675323109 gopurs_runtime.Value
var once_map__2675323109 sync.Once
func Get_map__2675323109() gopurs_runtime.Value {
	once_map__2675323109.Do(func() {
		cache_map__2675323109 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2675323109(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2675323109
}

var cache_map__3871729957 gopurs_runtime.Value
var once_map__3871729957 sync.Once
func Get_map__3871729957() gopurs_runtime.Value {
	once_map__3871729957.Do(func() {
		cache_map__3871729957 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3871729957(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3871729957
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__831829748 gopurs_runtime.Value
var once_map__831829748 sync.Once
func Get_map__831829748() gopurs_runtime.Value {
	once_map__831829748.Do(func() {
		cache_map__831829748 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__831829748(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__831829748
}

var cache_map__2745625428 gopurs_runtime.Value
var once_map__2745625428 sync.Once
func Get_map__2745625428() gopurs_runtime.Value {
	once_map__2745625428.Do(func() {
		cache_map__2745625428 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2745625428(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2745625428
}

var cache_map__1510739772 gopurs_runtime.Value
var once_map__1510739772 sync.Once
func Get_map__1510739772() gopurs_runtime.Value {
	once_map__1510739772.Do(func() {
		cache_map__1510739772 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1510739772(f_0_box, l_1_box)
})
	})
	return cache_map__1510739772
}

var cache_map__3565923196 gopurs_runtime.Value
var once_map__3565923196 sync.Once
func Get_map__3565923196() gopurs_runtime.Value {
	once_map__3565923196.Do(func() {
		cache_map__3565923196 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3565923196(f_0_box, l_1_box)
})
	})
	return cache_map__3565923196
}

var cache_map__109003388 gopurs_runtime.Value
var once_map__109003388 sync.Once
func Get_map__109003388() gopurs_runtime.Value {
	once_map__109003388.Do(func() {
		cache_map__109003388 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__109003388(f_0_box, l_1_box)
})
	})
	return cache_map__109003388
}

var cache_map__1806510684 gopurs_runtime.Value
var once_map__1806510684 sync.Once
func Get_map__1806510684() gopurs_runtime.Value {
	once_map__1806510684.Do(func() {
		cache_map__1806510684 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1806510684(f_0_box, l_1_box)
})
	})
	return cache_map__1806510684
}

var cache_map__3269387708 gopurs_runtime.Value
var once_map__3269387708 sync.Once
func Get_map__3269387708() gopurs_runtime.Value {
	once_map__3269387708.Do(func() {
		cache_map__3269387708 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3269387708(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__3269387708
}

var cache_map__271334204 gopurs_runtime.Value
var once_map__271334204 sync.Once
func Get_map__271334204() gopurs_runtime.Value {
	once_map__271334204.Do(func() {
		cache_map__271334204 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__271334204(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__271334204
}

var cache_mapWithIndex__55256674 gopurs_runtime.Value
var once_mapWithIndex__55256674 sync.Once
func Get_mapWithIndex__55256674() gopurs_runtime.Value {
	once_mapWithIndex__55256674.Do(func() {
		cache_mapWithIndex__55256674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__55256674(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mapWithIndex__55256674
}

var cache_mapWithIndex__574674314 gopurs_runtime.Value
var once_mapWithIndex__574674314 sync.Once
func Get_mapWithIndex__574674314() gopurs_runtime.Value {
	once_mapWithIndex__574674314.Do(func() {
		cache_mapWithIndex__574674314 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__574674314(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mapWithIndex__574674314
}

var cache_mapWithIndex__3380890378 gopurs_runtime.Value
var once_mapWithIndex__3380890378 sync.Once
func Get_mapWithIndex__3380890378() gopurs_runtime.Value {
	once_mapWithIndex__3380890378.Do(func() {
		cache_mapWithIndex__3380890378 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__3380890378(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mapWithIndex__3380890378
}

var cache_eq1Lazy__3789574347 gopurs_runtime.Value
var once_eq1Lazy__3789574347 sync.Once
func Get_eq1Lazy__3789574347() gopurs_runtime.Value {
	once_eq1Lazy__3789574347.Do(func() {
		cache_eq1Lazy__3789574347 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_eqLazy(), dictEq_0), "eq")
}))
	})
	return cache_eq1Lazy__3789574347
}

var cache_eq1Lazy__251633054 gopurs_runtime.Value
var once_eq1Lazy__251633054 sync.Once
func Get_eq1Lazy__251633054() gopurs_runtime.Value {
	once_eq1Lazy__251633054.Do(func() {
		cache_eq1Lazy__251633054 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_eqLazy(), dictEq_0), "eq")
}))
	})
	return cache_eq1Lazy__251633054
}

var cache_functorLazy__491347738 gopurs_runtime.Value
var once_functorLazy__491347738 sync.Once
func Get_functorLazy__491347738() gopurs_runtime.Value {
	once_functorLazy__491347738.Do(func() {
		cache_functorLazy__491347738 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy__491347738
}

var cache_ord1Lazy__2079329387 gopurs_runtime.Value
var once_ord1Lazy__2079329387 sync.Once
func Get_ord1Lazy__2079329387() gopurs_runtime.Value {
	once_ord1Lazy__2079329387.Do(func() {
		cache_ord1Lazy__2079329387 = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Lazy.Get_eq1Lazy()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_ordLazy(), dictOrd_0), "compare")
}))
	})
	return cache_ord1Lazy__2079329387
}

var cache_altList__3296309911 gopurs_runtime.Value
var once_altList__3296309911 sync.Once
func Get_altList__3296309911() gopurs_runtime.Value {
	once_altList__3296309911.Do(func() {
		cache_altList__3296309911 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.RecordGet(Get_semigroupList(), "append"))
	})
	return cache_altList__3296309911
}

var cache_alternativeList__16377757 gopurs_runtime.Value
var once_alternativeList__16377757 sync.Once
func Get_alternativeList__16377757() gopurs_runtime.Value {
	once_alternativeList__16377757.Do(func() {
		cache_alternativeList__16377757 = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusList()
}))
	})
	return cache_alternativeList__16377757
}

var cache_applicativeList__37190504 gopurs_runtime.Value
var once_applicativeList__37190504 sync.Once
func Get_applicativeList__37190504() gopurs_runtime.Value {
	once_applicativeList__37190504.Do(func() {
		cache_applicativeList__37190504 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_0, Get_nil()})}
}))
}))
	})
	return cache_applicativeList__37190504
}

var cache_applicativeNonEmptyList__1064240936 gopurs_runtime.Value
var once_applicativeNonEmptyList__1064240936 sync.Once
func Get_applicativeNonEmptyList__1064240936() gopurs_runtime.Value {
	once_applicativeNonEmptyList__1064240936.Do(func() {
		cache_applicativeNonEmptyList__1064240936 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
}))
}))
	})
	return cache_applicativeNonEmptyList__1064240936
}

var cache_applyList__1470982352 gopurs_runtime.Value
var once_applyList__1470982352 sync.Once
func Get_applyList__1470982352() gopurs_runtime.Value {
	once_applyList__1470982352.Do(func() {
		cache_applyList__1470982352 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyList__1470982352
}

var cache_applyList__1358886895 gopurs_runtime.Value
var once_applyList__1358886895 sync.Once
func Get_applyList__1358886895() gopurs_runtime.Value {
	once_applyList__1358886895.Do(func() {
		cache_applyList__1358886895 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyList__1358886895
}

var cache_applyNonEmptyList__1016064038 gopurs_runtime.Value
var once_applyNonEmptyList__1016064038 sync.Once
func Get_applyNonEmptyList__1016064038() gopurs_runtime.Value {
	once_applyNonEmptyList__1016064038.Do(func() {
		cache_applyNonEmptyList__1016064038 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
v2_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v1_1)
_ = v2_2_0
v3_3_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v3_3_1
__local_var_4_2 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0
_ = __local_var_4_2
__local_var_5_3 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3_1.UnsafePtr).V0
_ = __local_var_6_4
__local_var_7_5 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3_1.UnsafePtr).V1
_ = __local_var_7_5
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), __local_var_7_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_2, Get_nil()})}
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_4, __local_var_7_5})}
})), __local_var_5_3))})}
}))
})
}))
	})
	return cache_applyNonEmptyList__1016064038
}

var cache_bindList__469219920 gopurs_runtime.Value
var once_bindList__469219920 sync.Once
func Get_bindList__469219920() gopurs_runtime.Value {
	once_bindList__469219920.Do(func() {
		cache_bindList__469219920 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_bindList__469219920
}

var cache_bindList__1503691431 gopurs_runtime.Value
var once_bindList__1503691431 sync.Once
func Get_bindList__1503691431() gopurs_runtime.Value {
	once_bindList__1503691431.Do(func() {
		cache_bindList__1503691431 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_bindList__1503691431
}

var cache_bindNonEmptyList__3512977895 gopurs_runtime.Value
var once_bindNonEmptyList__3512977895 sync.Once
func Get_bindNonEmptyList__3512977895() gopurs_runtime.Value {
	once_bindNonEmptyList__3512977895.Do(func() {
		cache_bindNonEmptyList__3512977895 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_2_0
__local_var_3_1 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V1
_ = __local_var_3_1
v2_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0)))
_ = v2_4_2
__local_var_5_3 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v2_4_2)}.UnsafePtr).V0
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v2_4_2)}.UnsafePtr).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), __local_var_6_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), __local_var_3_1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toList(gopurs_runtime.Apply(f_1, x_8))
})))})}
}))
})
}))
	})
	return cache_bindNonEmptyList__3512977895
}

var cache_cons__716923058 gopurs_runtime.Value
var once_cons__716923058 sync.Once
func Get_cons__716923058() gopurs_runtime.Value {
	once_cons__716923058.Do(func() {
		cache_cons__716923058 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__716923058(x_0_box, xs_1_box)
})
	})
	return cache_cons__716923058
}

var cache_cons__720046150 gopurs_runtime.Value
var once_cons__720046150 sync.Once
func Get_cons__720046150() gopurs_runtime.Value {
	once_cons__720046150.Do(func() {
		cache_cons__720046150 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__720046150(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](x_0_box), xs_1_box)
})
	})
	return cache_cons__720046150
}

var cache_cons__2305074921 gopurs_runtime.Value
var once_cons__2305074921 sync.Once
func Get_cons__2305074921() gopurs_runtime.Value {
	once_cons__2305074921.Do(func() {
		cache_cons__2305074921 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__2305074921(x_0_box, xs_1_box)
})
	})
	return cache_cons__2305074921
}

var cache_cons__3391588829 gopurs_runtime.Value
var once_cons__3391588829 sync.Once
func Get_cons__3391588829() gopurs_runtime.Value {
	once_cons__3391588829.Do(func() {
		cache_cons__3391588829 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__3391588829(x_0_box, xs_1_box)
})
	})
	return cache_cons__3391588829
}

var cache_cons__2134285409 gopurs_runtime.Value
var once_cons__2134285409 sync.Once
func Get_cons__2134285409() gopurs_runtime.Value {
	once_cons__2134285409.Do(func() {
		cache_cons__2134285409 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__2134285409(x_0_box, xs_1_box)
})
	})
	return cache_cons__2134285409
}

var cache_eq1List__2902948510 gopurs_runtime.Value
var once_eq1List__2902948510 sync.Once
func Get_eq1List__2902948510() gopurs_runtime.Value {
	once_eq1List__2902948510.Do(func() {
		cache_eq1List__2902948510 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_11 gopurs_runtime.Value
go__go_3_0_11 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_11:
for {
if false { continue go__go_3_0_11 }
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
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
continue go__go_3_0_11
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_0_11, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_2)))}).IntVal) != (0))
})
})
}))
	})
	return cache_eq1List__2902948510
}

var cache_eq1NonEmptyList__2973916286 gopurs_runtime.Value
var once_eq1NonEmptyList__2973916286 sync.Once
func Get_eq1NonEmptyList__2973916286() gopurs_runtime.Value {
	once_eq1NonEmptyList__2973916286.Do(func() {
		cache_eq1NonEmptyList__2973916286 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(pkg_Data_NonEmpty.Get_eqNonEmpty(), Get_eq1List(), dictEq_0))
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_eq1Lazy(), "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty1_1_0)}, v_2, v1_3).IntVal) != (0))
})
})
}))
	})
	return cache_eq1NonEmptyList__2973916286
}

var cache_extendNonEmptyList__2847971828 gopurs_runtime.Value
var once_extendNonEmptyList__2847971828 sync.Once
func Get_extendNonEmptyList__2847971828() gopurs_runtime.Value {
	once_extendNonEmptyList__2847971828.Do(func() {
		cache_extendNonEmptyList__2847971828 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1).UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.RecordGet(v1_5, "acc")
_ = __local_var_6_1
__local_var_7_2 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, __local_var_6_1})}
})))
_ = __local_var_7_2
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, __local_var_6_1})}
})), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_7_2, gopurs_runtime.RecordGet(v1_5, "val")})}
})))
})
}), gopurs_runtime.RecordDict2("acc", "val", Get_nil(), Get_nil()), __local_var_2_0), "val")})}
}))
})
}))
	})
	return cache_extendNonEmptyList__2847971828
}

var cache_foldableList__4097915271 gopurs_runtime.Value
var once_foldableList__4097915271 sync.Once
func Get_foldableList__4097915271() gopurs_runtime.Value {
	once_foldableList__4097915271.Do(func() {
		cache_foldableList__4097915271 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_12 gopurs_runtime.Value
go__go_1_2_12 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_12:
for {
if false { continue go__go_1_2_12 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_12
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
return go__go_1_2_12
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__4097915271
}

var cache_foldableList__331628915 gopurs_runtime.Value
var once_foldableList__331628915 sync.Once
func Get_foldableList__331628915() gopurs_runtime.Value {
	once_foldableList__331628915.Do(func() {
		cache_foldableList__331628915 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_13 gopurs_runtime.Value
go__go_1_2_13 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_13:
for {
if false { continue go__go_1_2_13 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_13
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
return go__go_1_2_13
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__331628915
}

var cache_foldableList__3094856796 gopurs_runtime.Value
var once_foldableList__3094856796 sync.Once
func Get_foldableList__3094856796() gopurs_runtime.Value {
	once_foldableList__3094856796.Do(func() {
		cache_foldableList__3094856796 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_14 gopurs_runtime.Value
go__go_1_2_14 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_14:
for {
if false { continue go__go_1_2_14 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_14
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
return go__go_1_2_14
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__3094856796
}

var cache_foldableList__1218280485 gopurs_runtime.Value
var once_foldableList__1218280485 sync.Once
func Get_foldableList__1218280485() gopurs_runtime.Value {
	once_foldableList__1218280485.Do(func() {
		cache_foldableList__1218280485 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_15 gopurs_runtime.Value
go__go_1_2_15 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_15:
for {
if false { continue go__go_1_2_15 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_15
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
return go__go_1_2_15
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__1218280485
}

var cache_foldableNonEmptyList__2027644716 gopurs_runtime.Value
var once_foldableNonEmptyList__2027644716 sync.Once
func Get_foldableNonEmptyList__2027644716() gopurs_runtime.Value {
	once_foldableNonEmptyList__2027644716.Do(func() {
		cache_foldableNonEmptyList__2027644716 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V1, f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V2, f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}))
	})
	return cache_foldableNonEmptyList__2027644716
}

var cache_foldableWithIndexList__662860203 gopurs_runtime.Value
var once_foldableWithIndexList__662860203 sync.Once
func Get_foldableWithIndexList__662860203() gopurs_runtime.Value {
	once_foldableWithIndexList__662860203.Do(func() {
		cache_foldableWithIndexList__662860203 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_5)
_ = __local_var_6_2
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
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
__local_var_4_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_4_6.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_8 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
__local_var_5_9 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_5_9.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, __local_var_4_8})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_nil()})}, xs_2))
_ = v_3_7
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
__local_var_6_11 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1).UnsafePtr).V1
})
})
}))
	})
	return cache_foldableWithIndexList__662860203
}

var cache_foldableWithIndexList__3899545502 gopurs_runtime.Value
var once_foldableWithIndexList__3899545502 sync.Once
func Get_foldableWithIndexList__3899545502() gopurs_runtime.Value {
	once_foldableWithIndexList__3899545502.Do(func() {
		cache_foldableWithIndexList__3899545502 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_5)
_ = __local_var_6_2
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
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
__local_var_4_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_4_6.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_8 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
__local_var_5_9 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_5_9.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, __local_var_4_8})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_nil()})}, xs_2))
_ = v_3_7
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
__local_var_6_11 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1).UnsafePtr).V1
})
})
}))
	})
	return cache_foldableWithIndexList__3899545502
}

var cache_foldableWithIndexNonEmptyList__3565687582 gopurs_runtime.Value
var once_foldableWithIndexNonEmptyList__3565687582 sync.Once
func Get_foldableWithIndexNonEmptyList__3565687582() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyList__3565687582.Do(func() {
		cache_foldableWithIndexNonEmptyList__3565687582 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_0
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(__local_var_3_0, x_4))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_1
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V2, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V3, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_2, x_4))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))})
})
})
}))
	})
	return cache_foldableWithIndexNonEmptyList__3565687582
}

var cache_functorList__699353223 gopurs_runtime.Value
var once_functorList__699353223 sync.Once
func Get_functorList__699353223() gopurs_runtime.Value {
	once_functorList__699353223.Do(func() {
		cache_functorList__699353223 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_functorList__699353223
}

var cache_functorList__3996674161 gopurs_runtime.Value
var once_functorList__3996674161 sync.Once
func Get_functorList__3996674161() gopurs_runtime.Value {
	once_functorList__3996674161.Do(func() {
		cache_functorList__3996674161 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_functorList__3996674161
}

var cache_functorNonEmptyList__1934212625 gopurs_runtime.Value
var once_functorNonEmptyList__1934212625 sync.Once
func Get_functorNonEmptyList__1934212625() gopurs_runtime.Value {
	once_functorNonEmptyList__1934212625.Do(func() {
		cache_functorNonEmptyList__1934212625 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_functorNonEmpty()).V0, f_0), v_1)
})
}))
	})
	return cache_functorNonEmptyList__1934212625
}

var cache_functorWithIndexList__353314402 gopurs_runtime.Value
var once_functorWithIndexList__353314402 sync.Once
func Get_functorWithIndexList__353314402() gopurs_runtime.Value {
	once_functorWithIndexList__353314402.Do(func() {
		cache_functorWithIndexList__353314402 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2)
_ = __local_var_4_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_0, acc_3})}
}))
})
})
}), Get_nil())
}))
	})
	return cache_functorWithIndexList__353314402
}

var cache_functorWithIndexNonEmptyList__812352994 gopurs_runtime.Value
var once_functorWithIndexNonEmptyList__812352994 sync.Once
func Get_functorWithIndexNonEmptyList__812352994() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyList__812352994.Do(func() {
		cache_functorWithIndexNonEmptyList__812352994 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_functorWithIndex()).V1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_0, x_4))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1)))})))}
}))
})
}))
	})
	return cache_functorWithIndexNonEmptyList__812352994
}

var cache_lazyList__601034736 gopurs_runtime.Value
var once_lazyList__601034736 sync.Once
func Get_lazyList__601034736() gopurs_runtime.Value {
	once_lazyList__601034736.Do(func() {
		cache_lazyList__601034736 = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(f_0, x_1))
}))
}))
	})
	return cache_lazyList__601034736
}

var cache_monadList__2596899283 gopurs_runtime.Value
var once_monadList__2596899283 sync.Once
func Get_monadList__2596899283() gopurs_runtime.Value {
	once_monadList__2596899283.Do(func() {
		cache_monadList__2596899283 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindList()
}))
	})
	return cache_monadList__2596899283
}

var cache_nil__1478684294 gopurs_runtime.Value
var once_nil__1478684294 sync.Once
func Get_nil__1478684294() gopurs_runtime.Value {
	once_nil__1478684294.Do(func() {
		cache_nil__1478684294 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__1478684294
}

var cache_nil__3988504114 gopurs_runtime.Value
var once_nil__3988504114 sync.Once
func Get_nil__3988504114() gopurs_runtime.Value {
	once_nil__3988504114.Do(func() {
		cache_nil__3988504114 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__3988504114
}

var cache_nil__2012296605 gopurs_runtime.Value
var once_nil__2012296605 sync.Once
func Get_nil__2012296605() gopurs_runtime.Value {
	once_nil__2012296605.Do(func() {
		cache_nil__2012296605 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__2012296605
}

var cache_nil__2288399465 gopurs_runtime.Value
var once_nil__2288399465 sync.Once
func Get_nil__2288399465() gopurs_runtime.Value {
	once_nil__2288399465.Do(func() {
		cache_nil__2288399465 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__2288399465
}

var cache_nil__4122162182 gopurs_runtime.Value
var once_nil__4122162182 sync.Once
func Get_nil__4122162182() gopurs_runtime.Value {
	once_nil__4122162182.Do(func() {
		cache_nil__4122162182 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__4122162182
}

var cache_plusList__2873873584 gopurs_runtime.Value
var once_plusList__2873873584 sync.Once
func Get_plusList__2873873584() gopurs_runtime.Value {
	once_plusList__2873873584.Do(func() {
		cache_plusList__2873873584 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altList()
}), Get_nil())
	})
	return cache_plusList__2873873584
}

var cache_plusList__3460472018 gopurs_runtime.Value
var once_plusList__3460472018 sync.Once
func Get_plusList__3460472018() gopurs_runtime.Value {
	once_plusList__3460472018.Do(func() {
		cache_plusList__3460472018 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altList()
}), Get_nil())
	})
	return cache_plusList__3460472018
}

var cache_semigroupList__1199693447 gopurs_runtime.Value
var once_semigroupList__1199693447 sync.Once
func Get_semigroupList__1199693447() gopurs_runtime.Value {
	once_semigroupList__1199693447.Do(func() {
		cache_semigroupList__1199693447 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList__1199693447
}

var cache_semigroupList__3612943602 gopurs_runtime.Value
var once_semigroupList__3612943602 sync.Once
func Get_semigroupList__3612943602() gopurs_runtime.Value {
	once_semigroupList__3612943602.Do(func() {
		cache_semigroupList__3612943602 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList__3612943602
}

var cache_semigroupList__2598308723 gopurs_runtime.Value
var once_semigroupList__2598308723 sync.Once
func Get_semigroupList__2598308723() gopurs_runtime.Value {
	once_semigroupList__2598308723.Do(func() {
		cache_semigroupList__2598308723 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList__2598308723
}

var cache_semigroupList__4136327256 gopurs_runtime.Value
var once_semigroupList__4136327256 sync.Once
func Get_semigroupList__4136327256() gopurs_runtime.Value {
	once_semigroupList__4136327256.Do(func() {
		cache_semigroupList__4136327256 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList__4136327256
}

var cache_step__3545407802 gopurs_runtime.Value
var once_step__3545407802 sync.Once
func Get_step__3545407802() gopurs_runtime.Value {
	once_step__3545407802.Do(func() {
		cache_step__3545407802 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step__3545407802(x_0_box))}
})
	})
	return cache_step__3545407802
}

var cache_step__4184651873 gopurs_runtime.Value
var once_step__4184651873 sync.Once
func Get_step__4184651873() gopurs_runtime.Value {
	once_step__4184651873.Do(func() {
		cache_step__4184651873 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step__4184651873(x_0_box))}
})
	})
	return cache_step__4184651873
}

var cache_step__4057057377 gopurs_runtime.Value
var once_step__4057057377 sync.Once
func Get_step__4057057377() gopurs_runtime.Value {
	once_step__4057057377.Do(func() {
		cache_step__4057057377 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step__4057057377(x_0_box))}
})
	})
	return cache_step__4057057377
}

var cache_step__2999566881 gopurs_runtime.Value
var once_step__2999566881 sync.Once
func Get_step__2999566881() gopurs_runtime.Value {
	once_step__2999566881.Do(func() {
		cache_step__2999566881 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step__2999566881(x_0_box))}
})
	})
	return cache_step__2999566881
}

var cache_toList__1017592434 gopurs_runtime.Value
var once_toList__1017592434 sync.Once
func Get_toList__1017592434() gopurs_runtime.Value {
	once_toList__1017592434.Do(func() {
		cache_toList__1017592434 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toList__1017592434(v_0_box)
})
	})
	return cache_toList__1017592434
}

var cache_toList__4101396777 gopurs_runtime.Value
var once_toList__4101396777 sync.Once
func Get_toList__4101396777() gopurs_runtime.Value {
	once_toList__4101396777.Do(func() {
		cache_toList__4101396777 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toList__4101396777(v_0_box)
})
	})
	return cache_toList__4101396777
}

var cache_traversableList__3068288903 gopurs_runtime.Value
var once_traversableList__3068288903 sync.Once
func Get_traversableList__3068288903() gopurs_runtime.Value {
	once_traversableList__3068288903.Do(func() {
		cache_traversableList__3068288903 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_nil()))
})
}))
	})
	return cache_traversableList__3068288903
}

var cache_traversableList__589375054 gopurs_runtime.Value
var once_traversableList__589375054 sync.Once
func Get_traversableList__589375054() gopurs_runtime.Value {
	once_traversableList__589375054.Do(func() {
		cache_traversableList__589375054 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_nil()))
})
}))
	})
	return cache_traversableList__589375054
}

var cache_traversableNonEmptyList__3951740999 gopurs_runtime.Value
var once_traversableNonEmptyList__3951740999 sync.Once
func Get_traversableNonEmptyList__3951740999() gopurs_runtime.Value {
	once_traversableNonEmptyList__3951740999.Do(func() {
		cache_traversableNonEmptyList__3951740999 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_3))}
}))
}), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableNonEmpty()).V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xxs_4))}
}))
}), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableNonEmpty()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_3)))}))
})
})
}))
	})
	return cache_traversableNonEmptyList__3951740999
}

var cache_unfoldable1List__4025223016 gopurs_runtime.Value
var once_unfoldable1List__4025223016 sync.Once
func Get_unfoldable1List__4025223016() gopurs_runtime.Value {
	once_unfoldable1List__4025223016.Do(func() {
		cache_unfoldable1List__4025223016 = func() gopurs_runtime.Value {
var go__go_0_0_16 gopurs_runtime.Value
_ = go__go_0_0_16
go__go_0_0_16 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
__local_var_5_3 := gopurs_runtime.Apply2(go__go_0_0_16, f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1.UnsafePtr).V0)
_ = __local_var_5_3
__t5 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0, __local_var_5_3})}
}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0, Get_nil()})}
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
return gopurs_runtime.RecordDict1("unfoldr1", go__go_0_0_16)
}()
	})
	return cache_unfoldable1List__4025223016
}

var cache_maybe__919206801 gopurs_runtime.Value
var once_maybe__919206801 sync.Once
func Get_maybe__919206801() gopurs_runtime.Value {
	once_maybe__919206801.Do(func() {
		cache_maybe__919206801 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__919206801(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__919206801
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_unwrap__1971311275 gopurs_runtime.Value
var once_unwrap__1971311275 sync.Once
func Get_unwrap__1971311275() gopurs_runtime.Value {
	once_unwrap__1971311275.Do(func() {
		cache_unwrap__1971311275 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__1971311275(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__1971311275
}

var cache_unwrap__1997596318 gopurs_runtime.Value
var once_unwrap__1997596318 sync.Once
func Get_unwrap__1997596318() gopurs_runtime.Value {
	once_unwrap__1997596318.Do(func() {
		cache_unwrap__1997596318 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__1997596318(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__1997596318
}

var cache_unwrap__4073450430 gopurs_runtime.Value
var once_unwrap__4073450430 sync.Once
func Get_unwrap__4073450430() gopurs_runtime.Value {
	once_unwrap__4073450430.Do(func() {
		cache_unwrap__4073450430 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__4073450430(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__4073450430
}

var cache_unwrap__3159071614 gopurs_runtime.Value
var once_unwrap__3159071614 sync.Once
func Get_unwrap__3159071614() gopurs_runtime.Value {
	once_unwrap__3159071614.Do(func() {
		cache_unwrap__3159071614 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3159071614(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3159071614
}

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

var cache_unwrap__831442259 gopurs_runtime.Value
var once_unwrap__831442259 sync.Once
func Get_unwrap__831442259() gopurs_runtime.Value {
	once_unwrap__831442259.Do(func() {
		cache_unwrap__831442259 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__831442259(__eta0_0_box)
})
	})
	return cache_unwrap__831442259
}

var cache_unwrap__4291124211 gopurs_runtime.Value
var once_unwrap__4291124211 sync.Once
func Get_unwrap__4291124211() gopurs_runtime.Value {
	once_unwrap__4291124211.Do(func() {
		cache_unwrap__4291124211 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__4291124211(__eta0_0_box)
})
	})
	return cache_unwrap__4291124211
}

var cache_head__4279565926 gopurs_runtime.Value
var once_head__4279565926 sync.Once
func Get_head__4279565926() gopurs_runtime.Value {
	once_head__4279565926.Do(func() {
		cache_head__4279565926 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head__4279565926(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_head__4279565926
}

var cache_singleton__3741573463 gopurs_runtime.Value
var once_singleton__3741573463 sync.Once
func Get_singleton__3741573463() gopurs_runtime.Value {
	once_singleton__3741573463.Do(func() {
		cache_singleton__3741573463 = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton__3741573463(gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](dictPlus_0_box))
})
	})
	return cache_singleton__3741573463
}

var cache_singleton__532815287 gopurs_runtime.Value
var once_singleton__532815287 sync.Once
func Get_singleton__532815287() gopurs_runtime.Value {
	once_singleton__532815287.Do(func() {
		cache_singleton__532815287 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton__532815287(__eta0_0_box)
})
	})
	return cache_singleton__532815287
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_compare1__650153534 gopurs_runtime.Value
var once_compare1__650153534 sync.Once
func Get_compare1__650153534() gopurs_runtime.Value {
	once_compare1__650153534.Do(func() {
		cache_compare1__650153534 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare1__650153534(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare1__650153534
}

var cache_compare1__3282065035 gopurs_runtime.Value
var once_compare1__3282065035 sync.Once
func Get_compare1__3282065035() gopurs_runtime.Value {
	once_compare1__3282065035.Do(func() {
		cache_compare1__3282065035 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare1__3282065035(dictOrd_0_box)
})
	})
	return cache_compare1__3282065035
}

var cache_sub__1124926121 gopurs_runtime.Value
var once_sub__1124926121 sync.Once
func Get_sub__1124926121() gopurs_runtime.Value {
	once_sub__1124926121.Do(func() {
		cache_sub__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__1124926121
}

var cache_sub__190951261 gopurs_runtime.Value
var once_sub__190951261 sync.Once
func Get_sub__190951261() gopurs_runtime.Value {
	once_sub__190951261.Do(func() {
		cache_sub__190951261 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__190951261(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_sub__190951261
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__1124926121 gopurs_runtime.Value
var once_append__1124926121 sync.Once
func Get_append__1124926121() gopurs_runtime.Value {
	once_append__1124926121.Do(func() {
		cache_append__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1124926121
}

var cache_append__2462288412 gopurs_runtime.Value
var once_append__2462288412 sync.Once
func Get_append__2462288412() gopurs_runtime.Value {
	once_append__2462288412.Do(func() {
		cache_append__2462288412 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__2462288412(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__2462288412
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_append__2734706680 gopurs_runtime.Value
var once_append__2734706680 sync.Once
func Get_append__2734706680() gopurs_runtime.Value {
	once_append__2734706680.Do(func() {
		cache_append__2734706680 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__2734706680(xs_0_box, ys_1_box)
})
	})
	return cache_append__2734706680
}

var cache_add__1124926121 gopurs_runtime.Value
var once_add__1124926121 sync.Once
func Get_add__1124926121() gopurs_runtime.Value {
	once_add__1124926121.Do(func() {
		cache_add__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1124926121
}

var cache_add__190951261 gopurs_runtime.Value
var once_add__190951261 sync.Once
func Get_add__190951261() gopurs_runtime.Value {
	once_add__190951261.Do(func() {
		cache_add__190951261 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__190951261(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_add__190951261
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__3316320786 gopurs_runtime.Value
var once_show__3316320786 sync.Once
func Get_show__3316320786() gopurs_runtime.Value {
	once_show__3316320786.Do(func() {
		cache_show__3316320786 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3316320786(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__3316320786
}

var cache_show__1092279890 gopurs_runtime.Value
var once_show__1092279890 sync.Once
func Get_show__1092279890() gopurs_runtime.Value {
	once_show__1092279890.Do(func() {
		cache_show__1092279890 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__1092279890(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__1092279890
}

var cache_sequence__1886310617 gopurs_runtime.Value
var once_sequence__1886310617 sync.Once
func Get_sequence__1886310617() gopurs_runtime.Value {
	once_sequence__1886310617.Do(func() {
		cache_sequence__1886310617 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__1886310617(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence__1886310617
}

var cache_sequence__3720417425 gopurs_runtime.Value
var once_sequence__3720417425 sync.Once
func Get_sequence__3720417425() gopurs_runtime.Value {
	once_sequence__3720417425.Do(func() {
		cache_sequence__3720417425 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__3720417425(__eta0_0_box)
})
	})
	return cache_sequence__3720417425
}

var cache_sequence__3634000753 gopurs_runtime.Value
var once_sequence__3634000753 sync.Once
func Get_sequence__3634000753() gopurs_runtime.Value {
	once_sequence__3634000753.Do(func() {
		cache_sequence__3634000753 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__3634000753(__eta0_0_box)
})
	})
	return cache_sequence__3634000753
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
}

var cache_traverse__894989549 gopurs_runtime.Value
var once_traverse__894989549 sync.Once
func Get_traverse__894989549() gopurs_runtime.Value {
	once_traverse__894989549.Do(func() {
		cache_traverse__894989549 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__894989549(dictApplicative_0_box)
})
	})
	return cache_traverse__894989549
}

var cache_traverse__1157172365 gopurs_runtime.Value
var once_traverse__1157172365 sync.Once
func Get_traverse__1157172365() gopurs_runtime.Value {
	once_traverse__1157172365.Do(func() {
		cache_traverse__1157172365 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__1157172365(dictApplicative_0_box)
})
	})
	return cache_traverse__1157172365
}

var cache_traverse__3246764013 gopurs_runtime.Value
var once_traverse__3246764013 sync.Once
func Get_traverse__3246764013() gopurs_runtime.Value {
	once_traverse__3246764013.Do(func() {
		cache_traverse__3246764013 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__3246764013(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_traverse__3246764013
}

var cache_traverse__694301997 gopurs_runtime.Value
var once_traverse__694301997 sync.Once
func Get_traverse__694301997() gopurs_runtime.Value {
	once_traverse__694301997.Do(func() {
		cache_traverse__694301997 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__694301997(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_traverse__694301997
}

var cache_traverseWithIndex__2726076659 gopurs_runtime.Value
var once_traverseWithIndex__2726076659 sync.Once
func Get_traverseWithIndex__2726076659() gopurs_runtime.Value {
	once_traverseWithIndex__2726076659.Do(func() {
		cache_traverseWithIndex__2726076659 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex__2726076659(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverseWithIndex__2726076659
}

var cache_traverseWithIndex__2841069947 gopurs_runtime.Value
var once_traverseWithIndex__2841069947 sync.Once
func Get_traverseWithIndex__2841069947() gopurs_runtime.Value {
	once_traverseWithIndex__2841069947.Do(func() {
		cache_traverseWithIndex__2841069947 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex__2841069947(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_traverseWithIndex__2841069947
}

var cache_snd__1234761462 gopurs_runtime.Value
var once_snd__1234761462 sync.Once
func Get_snd__1234761462() gopurs_runtime.Value {
	once_snd__1234761462.Do(func() {
		cache_snd__1234761462 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__1234761462(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__1234761462
}

var cache_snd__4069939766 gopurs_runtime.Value
var once_snd__4069939766 sync.Once
func Get_snd__4069939766() gopurs_runtime.Value {
	once_snd__4069939766.Do(func() {
		cache_snd__4069939766 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__4069939766(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__4069939766
}

var cache_snd__2618926102 gopurs_runtime.Value
var once_snd__2618926102 sync.Once
func Get_snd__2618926102() gopurs_runtime.Value {
	once_snd__2618926102.Do(func() {
		cache_snd__2618926102 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__2618926102(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__2618926102
}

var cache_unfoldr1__2402610528 gopurs_runtime.Value
var once_unfoldr1__2402610528 sync.Once
func Get_unfoldr1__2402610528() gopurs_runtime.Value {
	once_unfoldr1__2402610528.Do(func() {
		cache_unfoldr1__2402610528 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr1__2402610528(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr1__2402610528
}

var cache_unfoldr1__2387656390 gopurs_runtime.Value
var once_unfoldr1__2387656390 sync.Once
func Get_unfoldr1__2387656390() gopurs_runtime.Value {
	once_unfoldr1__2387656390.Do(func() {
		cache_unfoldr1__2387656390 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr1__2387656390(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_unfoldr1__2387656390
}

type Constructor_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_List(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_NonEmptyList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_step(x_0_loop gopurs_runtime.Value) *Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_eq1List(), "eq1"), dictEq_0))
}

func Call_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_eqLazy(), gopurs_runtime.Apply2(pkg_Data_NonEmpty.Get_eqNonEmpty(), Get_eq1List(), dictEq_0))
}

func Call_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqList1_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_eq1List(), "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
_ = eqList1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqList1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_ord1List(), "compare1"), dictOrd_0))
}

func Call_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_ordLazy(), gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd_0))
}

func Call_cons(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1))
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Str("")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(shown_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(shown_3.StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(","), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), x_prime_4).StrVal())).StrVal())).StrVal())
})
}), gopurs_runtime.Str(""), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V1).StrVal())).StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(fromFoldable ["), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(__t1.StrVal()), gopurs_runtime.Str("])")).StrVal())).StrVal())
}))
}

func Call_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showLazy_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_showLazy(), gopurs_runtime.Apply2(pkg_Data_NonEmpty.Get_showNonEmpty(), dictShow_0, Call_showList(dictShow_0))))
_ = showLazy_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(NonEmptyList "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(showLazy_1_0.V0, v_2).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_showStep(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showList1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](Call_showList(dictShow_0))
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
__t1 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("("), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" : "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(showList1_1_0.V0, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())
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

func Call_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_lazyList(), "defer"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
v2_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1})}
}))
}))
}

func Call_pure__189931222(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3236307030(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__355615152(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__4203183626(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2962221386(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__1851858028(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2140510474(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Lazy.Get_force(), f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
}))
}

func Call_apply__3620326986(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
Bind1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_1
return gopurs_runtime.Apply2(Bind1_2_0.V1, __eta0_0, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_0.V1, __eta1_1, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_3_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
}

func Call_bind__3512795567(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3781138863(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__684765761(xs_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
}

func Call_bind__4082241(xs_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](__t0))}
}), xs_0)
}

func Call_defer__3967925939(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_defer__2590380358(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(f_0, x_1))
}))
}

func Call_compose__1987728071(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__346034828(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1555187646(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__4254807102(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq1__1773593252(dict_0_loop *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq1__2184765036(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_eqLazy(), dictEq_0), "eq")
}

func Call_eq1__3587165073(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_eqLazy(), dictEq_0), "eq")
}

func Call_foldMap__4098395794(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldMap__3562626100(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V0, __eta0_0, __eta1_1)
}

func Call_foldMap__2350611220(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V0, __eta0_0, __eta1_1)
}

func Call_foldl__1422885860(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__94807652(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__506543652(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__267332164(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__3131354468(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__524683195(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_5 gopurs_runtime.Value
go__go_1_0_5 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_5:
for {
if false { continue go__go_1_0_5 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_5
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
return go__go_1_0_5
}

func Call_foldl__3306117403(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_6 gopurs_runtime.Value
go__go_1_0_6 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](b_2_loop_val)
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_6:
for {
if false { continue go__go_1_0_6 }
var b_2 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply2(op_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0))
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_6
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](__t2))}
}
}()
})
})
return go__go_1_0_6
}

func Call_foldl__3737487037(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_7 gopurs_runtime.Value
go__go_1_0_7 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop string = b_2_loop_val.StrVal()
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_7:
for {
if false { continue go__go_1_0_7 }
var b_2 string = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Str(b_2)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, gopurs_runtime.Str(b_2), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0))}).StrVal()
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_7
__t2 = gopurs_runtime.Str(gopurs_runtime.Value{}.StrVal())
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Str(__t2.StrVal())
}
}()
})
})
return go__go_1_0_7
}

func Call_foldl__1985071933(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_8 gopurs_runtime.Value
go__go_1_0_8 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_8:
for {
if false { continue go__go_1_0_8 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0)
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_8
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
return go__go_1_0_8
}

func Call_foldl__536153533(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_9 gopurs_runtime.Value
go__go_1_0_9 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_9:
for {
if false { continue go__go_1_0_9 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0))})
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_9
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
return go__go_1_0_9
}

func Call_foldl__170252797(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_10 gopurs_runtime.Value
go__go_1_0_10 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](b_2_loop_val)
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_10:
for {
if false { continue go__go_1_0_10 }
var b_2 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply2(op_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0))}))
xs_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_10
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](__t2))}
}
}()
})
})
return go__go_1_0_10
}

func Call_foldl__2188030845(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V1, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldl__1444272061(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V1, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldr__2111289130(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__926146538(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__1985071933(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_nil(), xs_2))
}

func Call_foldr__2389967549(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_nil(), xs_2))
}

func Call_foldr__1278383325(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), Get_nil(), xs_2))
}

func Call_foldr__2188030845(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V2, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldr__3749276701(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableNonEmpty()).V2, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldMapWithIndex__2292551140(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMapWithIndex__1722031522(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V1, __eta0_0, __eta1_1)
}

func Call_foldMapWithIndex__852526914(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V1, __eta0_0, __eta1_1)
}

func Call_foldlWithIndex__2972270123(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_1
__local_var_4_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_2.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_0, x_3).UnsafePtr).V1
})
}

func Call_foldlWithIndex__234438827(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_1
__local_var_4_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_2.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](a_5))})})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_0, x_3).UnsafePtr).V1
})
}

func Call_foldlWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldlWithIndex__3610348555(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V2, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldlWithIndex__446277963(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V2, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldrWithIndex__2972270123(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
v_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
__local_var_5_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, __local_var_4_1})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_nil()})}, xs_2))
_ = v_3_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), a_7, __local_var_5_3)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), b_1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1).UnsafePtr).V1
}

func Call_foldrWithIndex__3735894283(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
v_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
__local_var_5_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](a_6))}, __local_var_4_1})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_nil()})}, xs_2))
_ = v_3_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](a_7))}, __local_var_5_3)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), b_1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1).UnsafePtr).V1
}

func Call_foldrWithIndex__500807083(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
v_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
__local_var_5_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](a_6))}, __local_var_4_1})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_nil()})}, xs_2))
_ = v_3_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](a_7))}, __local_var_5_3)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), b_1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1).UnsafePtr).V1
}

func Call_foldrWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_foldrWithIndex__3610348555(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V3, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldrWithIndex__63302635(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V3, __eta0_0, __eta1_1, __eta2_2)
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3658931456(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3019832928(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2175652032(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__848188896(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__913470112(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__4017289888(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3036739744(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3539834656(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__4040535013(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2665381605(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__4285761829(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3149795237(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1542634789(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2675323109(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3871729957(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__831829748(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2745625428(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1510739772(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
}

func Call_map__3565923196(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))))}
}))
}

func Call_map__109003388(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1)))})))}
}))
}

func Call_map__1806510684(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1)))})))}
}))
}

func Call_map__3269387708(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_functorNonEmpty()).V0, __eta0_0, __eta1_1)
}

func Call_map__271334204(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_functorNonEmpty()).V0, __eta0_0, __eta1_1)
}

func Call_mapWithIndex__55256674(dict_0_loop *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_mapWithIndex__574674314(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_functorWithIndex()).V1, __eta0_0, __eta1_1)
}

func Call_mapWithIndex__3380890378(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_functorWithIndex()).V1, __eta0_0, __eta1_1)
}

func Call_cons__716923058(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_cons__720046150(x_0_loop *Constructor_Cons[gopurs_runtime.Value], xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *Constructor_Cons[gopurs_runtime.Value] = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_cons__2305074921(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_cons__3391588829(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_cons__2134285409(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_step__3545407802(x_0_loop gopurs_runtime.Value) *Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_step__4184651873(x_0_loop gopurs_runtime.Value) *Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_step__4057057377(x_0_loop gopurs_runtime.Value) *Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_step__2999566881(x_0_loop gopurs_runtime.Value) *Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_toList__1017592434(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_lazyList(), "defer"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
v2_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1})}
}))
}))
}

func Call_toList__4101396777(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_lazyList(), "defer"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
v2_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v2_2_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1})}
}))
}))
}

func Call_maybe__919206801(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_unwrap__1971311275(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__1997596318(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__4073450430(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__3159071614(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__831442259(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_unwrap__4291124211(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_head__4279565926(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_singleton__3741573463(dictPlus_0_loop *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictPlus_0 *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value] = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := dictPlus_0.V1
_ = empty_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, empty_1_0})}
})
}

func Call_singleton__532815287(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta0_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare1__650153534(dict_0_loop *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare1__3282065035(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_ordLazy(), dictOrd_0), "compare")
}

func Call_sub__1124926121(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__190951261(dict_0_loop *pkg_Data_Ring.Constructor_Ring[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1124926121(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__2462288412(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__2734706680(xs_0_loop gopurs_runtime.Value, ys_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 gopurs_runtime.Value = ys_1_loop
_ = ys_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
}

func Call_add__1124926121(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__190951261(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__3316320786(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__1092279890(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_sequence__1886310617(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_sequence__3720417425(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableNonEmpty()).V2, __eta0_0)
}

func Call_sequence__3634000753(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableNonEmpty()).V2, __eta0_0)
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__894989549(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_nil()))
})
}

func Call_traverse__1157172365(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_cons(), gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](a_4))})), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_nil()))
})
}

func Call_traverse__3246764013(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableNonEmpty()).V3, __eta0_0, __eta1_1)
}

func Call_traverse__694301997(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableNonEmpty()).V3, __eta0_0, __eta1_1)
}

func Call_traverseWithIndex__2726076659(dict_0_loop *pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverseWithIndex__2841069947(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_traversableWithIndexNonEmpty()).V3, __eta0_0, __eta1_1)
}

func Call_snd__1234761462(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_snd__4069939766(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_snd__2618926102(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_unfoldr1__2402610528(dict_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unfoldr1__2387656390(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_unfoldable1NonEmpty()).V0, __eta0_0, __eta1_1)
}


