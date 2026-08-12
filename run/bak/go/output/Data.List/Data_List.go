package Data_List

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Lazy "gopurs/output/Control.Lazy"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unit "gopurs/output/Data.Unit"
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_identity(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_identity
}

var cache_Pattern gopurs_runtime.Value
var once_Pattern sync.Once
func Get_Pattern() gopurs_runtime.Value {
	once_Pattern.Do(func() {
		cache_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Pattern(x_0_box)
})
	})
	return cache_Pattern
}

var cache_updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		cache_updateAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_updateAt
}

var cache_unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		cache_unzip = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
_ = __local_var_1_0
__local_var_2_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_1_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_1, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}})})
	})
	return cache_unzip
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_uncons
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_toUnfoldable
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_tail(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_tail
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_stripPrefix(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](s_2_box)))}
})
	})
	return cache_stripPrefix
}

var cache_span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		cache_span = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box))
})
	})
	return cache_span
}

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_snoc(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_0_box), x_1_box))}
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_singleton(a_0_box))}
})
	})
	return cache_singleton
}

var cache_sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		cache_sortBy = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy(cmp_0_box)
})
	})
	return cache_sortBy
}

var cache_sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		cache_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sort(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_sort
}

var cache_tails gopurs_runtime.Value
var once_tails sync.Once
func Get_tails() gopurs_runtime.Value {
	once_tails.Do(func() {
		cache_tails = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_tails(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_tails
}

var cache_showPattern gopurs_runtime.Value
var once_showPattern sync.Once
func Get_showPattern() gopurs_runtime.Value {
	once_showPattern.Do(func() {
		cache_showPattern = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showPattern(dictShow_0_box)
})
	})
	return cache_showPattern
}

var cache_reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		cache_reverse = func() gopurs_runtime.Value {
var go__go_0_0_6 gopurs_runtime.Value
go__go_0_0_6 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_6:
for {
if false { continue go__go_0_0_6 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
continue go__go_0_0_6
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_reverse
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = func() gopurs_runtime.Value {
var go__go_0_0_7 gopurs_runtime.Value
go__go_0_0_7 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop int64 = v1_2_loop_val.IntVal
var v2_3_loop gopurs_runtime.Value = v2_3_loop_val
go__go_0_0_7:
for {
if false { continue go__go_0_0_7 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 int64 = v1_2_loop
_ = v1_2
var v2_3 gopurs_runtime.Value = v2_3_loop
_ = v2_3
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v1_2), gopurs_runtime.Int(1))).IntVal) != (0) {
var go__go_4_1_8 gopurs_runtime.Value
go__go_4_1_8 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_8:
for {
if false { continue go__go_4_1_8 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_8
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})))}
goto end_branch_5
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr == nil) {
var go__go_4_3_9 gopurs_runtime.Value
go__go_4_3_9 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_3_9:
for {
if false { continue go__go_4_3_9 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_3_9
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_3_9, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})))}
goto end_branch_5
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_3.UnsafePtr).V0, v_1})})
v1_2_loop = Call_sub__1043827704(gopurs_runtime.Int(v1_2), gopurs_runtime.Int(1)).IntVal
v2_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_3.UnsafePtr).V1)}
continue go__go_0_0_7
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
})
return gopurs_runtime.Apply(go__go_0_0_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_take
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(p_0_box)
})
	})
	return cache_takeWhile
}

var cache_unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		cache_unsnoc = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_unsnoc(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](lst_0_box)))}
})
	})
	return cache_unsnoc
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_zipWith(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_zipWith
}

var cache_zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		cache_zip = gopurs_runtime.Apply(Get_zipWith(), pkg_Data_Tuple.Get_Tuple())
	})
	return cache_zip
}

var cache_zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		cache_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_2_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_3_box))
})
	})
	return cache_zipWithA
}

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_go__range(start_0_box.IntVal, end_1_box.IntVal))}
})
	})
	return cache_go__range
}

var cache_partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		cache_partition = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_partition(p_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box))
})
	})
	return cache_partition
}

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_null
}

var cache_nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		cache_nubBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubBy(p_0_box)
})
	})
	return cache_nubBy
}

var cache_nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		cache_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nub(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_nub
}

var cache_newtypePattern gopurs_runtime.Value
var once_newtypePattern sync.Once
func Get_newtypePattern() gopurs_runtime.Value {
	once_newtypePattern.Do(func() {
		cache_newtypePattern = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypePattern
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(f_0_box)
})
	})
	return cache_mapMaybe
}

var cache_manyRec gopurs_runtime.Value
var once_manyRec sync.Once
func Get_manyRec() gopurs_runtime.Value {
	once_manyRec.Do(func() {
		cache_manyRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_manyRec(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_manyRec
}

var cache_someRec gopurs_runtime.Value
var once_someRec sync.Once
func Get_someRec() gopurs_runtime.Value {
	once_someRec.Do(func() {
		cache_someRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_someRec(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_1_box))
})
	})
	return cache_someRec
}

var cache_some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		cache_some = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_some
}

var cache_many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		cache_many = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_many
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(acc_0.IntVal), gopurs_runtime.Int(1)).IntVal)
})
}), gopurs_runtime.Int(0))
	})
	return cache_length
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_last(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_last
}

var cache_insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		cache_insertBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_insertBy(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_insertBy
}

var cache_insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		cache_insertAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_insertAt
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_insert
}

var cache_init gopurs_runtime.Value
var once_init sync.Once
func Get_init() gopurs_runtime.Value {
	once_init.Do(func() {
		cache_init = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_init(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](lst_0_box)))}
})
	})
	return cache_init
}

var cache_index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		cache_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box.IntVal))}
})
	})
	return cache_index
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_head(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_head
}

var cache_transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		cache_transpose = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_transpose(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_0_box)))}
})
	})
	return cache_transpose
}

var cache_groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		cache_groupBy = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_groupBy(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_groupBy
}

var cache_groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		cache_groupAllBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAllBy(p_0_box)
})
	})
	return cache_groupAllBy
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_group
}

var cache_groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		cache_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAll(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_groupAll
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_fromFoldable
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldM
}

var cache_findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		cache_findIndex = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findIndex(fn_0_box)
})
	})
	return cache_findIndex
}

var cache_findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		cache_findLastIndex = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex(fn_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box)))}
})
	})
	return cache_findLastIndex
}

var cache_filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		cache_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterM(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_filterM
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(p_0_box)
})
	})
	return cache_filter
}

var cache_intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		cache_intersectBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_intersectBy(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_intersectBy
}

var cache_intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		cache_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersect(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_intersect
}

var cache_nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		cache_nubByEq = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_nubByEq(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_nubByEq
}

var cache_nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		cache_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubEq(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_nubEq
}

var cache_eqPattern gopurs_runtime.Value
var once_eqPattern sync.Once
func Get_eqPattern() gopurs_runtime.Value {
	once_eqPattern.Do(func() {
		cache_eqPattern = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqPattern(dictEq_0_box)
})
	})
	return cache_eqPattern
}

var cache_ordPattern gopurs_runtime.Value
var once_ordPattern sync.Once
func Get_ordPattern() gopurs_runtime.Value {
	once_ordPattern.Do(func() {
		cache_ordPattern = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordPattern(dictOrd_0_box)
})
	})
	return cache_ordPattern
}

var cache_elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		cache_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemLastIndex(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
})
	})
	return cache_elemLastIndex
}

var cache_elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		cache_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemIndex(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
})
	})
	return cache_elemIndex
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(p_0_box)
})
	})
	return cache_dropWhile
}

var cache_dropEnd gopurs_runtime.Value
var once_dropEnd sync.Once
func Get_dropEnd() gopurs_runtime.Value {
	once_dropEnd.Do(func() {
		cache_dropEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_dropEnd(n_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box)))}
})
	})
	return cache_dropEnd
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_drop(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_drop
}

var cache_slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		cache_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_slice(start_0_box.IntVal, end_1_box.IntVal, xs_2_box)
})
	})
	return cache_slice
}

var cache_takeEnd gopurs_runtime.Value
var once_takeEnd sync.Once
func Get_takeEnd() gopurs_runtime.Value {
	once_takeEnd.Do(func() {
		cache_takeEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_takeEnd(n_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box)))}
})
	})
	return cache_takeEnd
}

var cache_deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		cache_deleteBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_deleteBy(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_deleteBy
}

var cache_unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		cache_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_unionBy(eq_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_unionBy
}

var cache_union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		cache_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_union(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_union
}

var cache_deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		cache_deleteAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_deleteAt(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_deleteAt
}

var cache_delete gopurs_runtime.Value
var once_delete sync.Once
func Get_delete() gopurs_runtime.Value {
	once_delete.Do(func() {
		cache_delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_delete
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_difference
}

var cache_concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		cache_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_concatMap(b_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](a_1_box)))}
})
	})
	return cache_concatMap
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_concat(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_0_box)))}
})
	})
	return cache_concat
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = Call_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_catMaybes
}

var cache_alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		cache_alterAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_alterAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_alterAt
}

var cache_modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		cache_modifyAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyAt(n_0_box.IntVal, f_1_box)
})
	})
	return cache_modifyAt
}

var cache_alt__267341625 gopurs_runtime.Value
var once_alt__267341625 sync.Once
func Get_alt__267341625() gopurs_runtime.Value {
	once_alt__267341625.Do(func() {
		cache_alt__267341625 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__267341625(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__267341625
}

var cache_alt__1433662167 gopurs_runtime.Value
var once_alt__1433662167 sync.Once
func Get_alt__1433662167() gopurs_runtime.Value {
	once_alt__1433662167.Do(func() {
		cache_alt__1433662167 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__1433662167(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__1433662167
}

var cache_alt__1643437292 gopurs_runtime.Value
var once_alt__1643437292 sync.Once
func Get_alt__1643437292() gopurs_runtime.Value {
	once_alt__1643437292.Do(func() {
		cache_alt__1643437292 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__1643437292(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__1643437292
}

var cache_pure__993904534 gopurs_runtime.Value
var once_pure__993904534 sync.Once
func Get_pure__993904534() gopurs_runtime.Value {
	once_pure__993904534.Do(func() {
		cache_pure__993904534 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__993904534(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__993904534
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

var cache_pure__3286817552 gopurs_runtime.Value
var once_pure__3286817552 sync.Once
func Get_pure__3286817552() gopurs_runtime.Value {
	once_pure__3286817552.Do(func() {
		cache_pure__3286817552 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3286817552(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3286817552
}

var cache_pure__243192752 gopurs_runtime.Value
var once_pure__243192752 sync.Once
func Get_pure__243192752() gopurs_runtime.Value {
	once_pure__243192752.Do(func() {
		cache_pure__243192752 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__243192752(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__243192752
}

var cache_apply__1030762512 gopurs_runtime.Value
var once_apply__1030762512 sync.Once
func Get_apply__1030762512() gopurs_runtime.Value {
	once_apply__1030762512.Do(func() {
		cache_apply__1030762512 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__1030762512(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_apply__1030762512
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

var cache_apply__2007181164 gopurs_runtime.Value
var once_apply__2007181164 sync.Once
func Get_apply__2007181164() gopurs_runtime.Value {
	once_apply__2007181164.Do(func() {
		cache_apply__2007181164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2007181164(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2007181164
}

var cache_apply__2169384906 gopurs_runtime.Value
var once_apply__2169384906 sync.Once
func Get_apply__2169384906() gopurs_runtime.Value {
	once_apply__2169384906.Do(func() {
		cache_apply__2169384906 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_apply__2169384906(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_apply__2169384906
}

var cache_lift2__2286084406 gopurs_runtime.Value
var once_lift2__2286084406 sync.Once
func Get_lift2__2286084406() gopurs_runtime.Value {
	once_lift2__2286084406.Do(func() {
		cache_lift2__2286084406 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2286084406(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2286084406
}

var cache_lift2__3007543670 gopurs_runtime.Value
var once_lift2__3007543670 sync.Once
func Get_lift2__3007543670() gopurs_runtime.Value {
	once_lift2__3007543670.Do(func() {
		cache_lift2__3007543670 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__3007543670(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__3007543670
}

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__2114580400 gopurs_runtime.Value
var once_lift2__2114580400 sync.Once
func Get_lift2__2114580400() gopurs_runtime.Value {
	once_lift2__2114580400.Do(func() {
		cache_lift2__2114580400 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2114580400(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2114580400
}

var cache_bind__3842572251 gopurs_runtime.Value
var once_bind__3842572251 sync.Once
func Get_bind__3842572251() gopurs_runtime.Value {
	once_bind__3842572251.Do(func() {
		cache_bind__3842572251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3842572251(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_bind__3842572251
}

var cache_bind__556822235 gopurs_runtime.Value
var once_bind__556822235 sync.Once
func Get_bind__556822235() gopurs_runtime.Value {
	once_bind__556822235.Do(func() {
		cache_bind__556822235 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__556822235(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_bind__556822235
}

var cache_bind__1429695463 gopurs_runtime.Value
var once_bind__1429695463 sync.Once
func Get_bind__1429695463() gopurs_runtime.Value {
	once_bind__1429695463.Do(func() {
		cache_bind__1429695463 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1429695463(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1429695463
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

var cache_bind__226726151 gopurs_runtime.Value
var once_bind__226726151 sync.Once
func Get_bind__226726151() gopurs_runtime.Value {
	once_bind__226726151.Do(func() {
		cache_bind__226726151 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__226726151(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__226726151
}

var cache_bind__2057121831 gopurs_runtime.Value
var once_bind__2057121831 sync.Once
func Get_bind__2057121831() gopurs_runtime.Value {
	once_bind__2057121831.Do(func() {
		cache_bind__2057121831 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2057121831(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2057121831
}

var cache_bind__1872090113 gopurs_runtime.Value
var once_bind__1872090113 sync.Once
func Get_bind__1872090113() gopurs_runtime.Value {
	once_bind__1872090113.Do(func() {
		cache_bind__1872090113 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_bind__1872090113(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box))}
})
	})
	return cache_bind__1872090113
}

var cache_bind__2207507201 gopurs_runtime.Value
var once_bind__2207507201 sync.Once
func Get_bind__2207507201() gopurs_runtime.Value {
	once_bind__2207507201.Do(func() {
		cache_bind__2207507201 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_bind__2207507201(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_0_box), v1_1_box))}
})
	})
	return cache_bind__2207507201
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

var cache_defer__449752361 gopurs_runtime.Value
var once_defer__449752361 sync.Once
func Get_defer__449752361() gopurs_runtime.Value {
	once_defer__449752361.Do(func() {
		cache_defer__449752361 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__449752361(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_defer__449752361
}

var cache_bifunctorStep__1527141292 gopurs_runtime.Value
var once_bifunctorStep__1527141292 sync.Once
func Get_bifunctorStep__1527141292() gopurs_runtime.Value {
	once_bifunctorStep__1527141292.Do(func() {
		cache_bifunctorStep__1527141292 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
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
}))
	})
	return cache_bifunctorStep__1527141292
}

var cache_monadRecMaybe__796215523 gopurs_runtime.Value
var once_monadRecMaybe__796215523 sync.Once
func Get_monadRecMaybe__796215523() gopurs_runtime.Value {
	once_monadRecMaybe__796215523.Do(func() {
		cache_monadRecMaybe__796215523 = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_monadMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}})}
goto end_branch_4
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 525585346) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)))}})}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 60402430) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
_ = __local_var_2_0
var go__go_3_5_27 gopurs_runtime.Value
go__go_3_5_27 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_5_27:
for {
if false { continue go__go_3_5_27 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply(__local_var_2_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_5_27
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t6 = (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_5_27, gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a0_1)))}))))}
})
}))
	})
	return cache_monadRecMaybe__796215523
}

var cache_tailRec__2110844386 gopurs_runtime.Value
var once_tailRec__2110844386 sync.Once
func Get_tailRec__2110844386() gopurs_runtime.Value {
	once_tailRec__2110844386.Do(func() {
		cache_tailRec__2110844386 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec__2110844386(f_0_box)
})
	})
	return cache_tailRec__2110844386
}

var cache_tailRec__2666749533 gopurs_runtime.Value
var once_tailRec__2666749533 sync.Once
func Get_tailRec__2666749533() gopurs_runtime.Value {
	once_tailRec__2666749533.Do(func() {
		cache_tailRec__2666749533 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec__2666749533(f_0_box)
})
	})
	return cache_tailRec__2666749533
}

var cache_tailRecM__1824976720 gopurs_runtime.Value
var once_tailRecM__1824976720 sync.Once
func Get_tailRecM__1824976720() gopurs_runtime.Value {
	once_tailRecM__1824976720.Do(func() {
		cache_tailRecM__1824976720 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__1824976720(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__1824976720
}

var cache_tailRecM__3615056745 gopurs_runtime.Value
var once_tailRecM__3615056745 sync.Once
func Get_tailRecM__3615056745() gopurs_runtime.Value {
	once_tailRecM__3615056745.Do(func() {
		cache_tailRecM__3615056745 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3615056745(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3615056745
}

var cache_tailRecM__3865988408 gopurs_runtime.Value
var once_tailRecM__3865988408 sync.Once
func Get_tailRecM__3865988408() gopurs_runtime.Value {
	once_tailRecM__3865988408.Do(func() {
		cache_tailRecM__3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3865988408(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3865988408
}

var cache_tailRecM__3515372237 gopurs_runtime.Value
var once_tailRecM__3515372237 sync.Once
func Get_tailRecM__3515372237() gopurs_runtime.Value {
	once_tailRecM__3515372237.Do(func() {
		cache_tailRecM__3515372237 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3515372237(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3515372237
}

var cache_tailRecM__1444729948 gopurs_runtime.Value
var once_tailRecM__1444729948 sync.Once
func Get_tailRecM__1444729948() gopurs_runtime.Value {
	once_tailRecM__1444729948.Do(func() {
		cache_tailRecM__1444729948 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__1444729948(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__1444729948
}

var cache_tailRecM2__1943630176 gopurs_runtime.Value
var once_tailRecM2__1943630176 sync.Once
func Get_tailRecM2__1943630176() gopurs_runtime.Value {
	once_tailRecM2__1943630176.Do(func() {
		cache_tailRecM2__1943630176 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2__1943630176(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_tailRecM2__1943630176
}

var cache_tailRecM2__1136195496 gopurs_runtime.Value
var once_tailRecM2__1136195496 sync.Once
func Get_tailRecM2__1136195496() gopurs_runtime.Value {
	once_tailRecM2__1136195496.Do(func() {
		cache_tailRecM2__1136195496 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2__1136195496(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_tailRecM2__1136195496
}

var cache_compose__706970832 gopurs_runtime.Value
var once_compose__706970832 sync.Once
func Get_compose__706970832() gopurs_runtime.Value {
	once_compose__706970832.Do(func() {
		cache_compose__706970832 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__706970832(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__706970832
}

var cache_compose__2527254334 gopurs_runtime.Value
var once_compose__2527254334 sync.Once
func Get_compose__2527254334() gopurs_runtime.Value {
	once_compose__2527254334.Do(func() {
		cache_compose__2527254334 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2527254334(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2527254334
}

var cache_compose__1604328382 gopurs_runtime.Value
var once_compose__1604328382 sync.Once
func Get_compose__1604328382() gopurs_runtime.Value {
	once_compose__1604328382.Do(func() {
		cache_compose__1604328382 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1604328382(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__1604328382
}

var cache_compose__2995688990 gopurs_runtime.Value
var once_compose__2995688990 sync.Once
func Get_compose__2995688990() gopurs_runtime.Value {
	once_compose__2995688990.Do(func() {
		cache_compose__2995688990 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2995688990(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2995688990
}

var cache_compose__4020612094 gopurs_runtime.Value
var once_compose__4020612094 sync.Once
func Get_compose__4020612094() gopurs_runtime.Value {
	once_compose__4020612094.Do(func() {
		cache_compose__4020612094 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__4020612094(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__4020612094
}

var cache_compose__2710321297 gopurs_runtime.Value
var once_compose__2710321297 sync.Once
func Get_compose__2710321297() gopurs_runtime.Value {
	once_compose__2710321297.Do(func() {
		cache_compose__2710321297 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2710321297(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2710321297
}

var cache_compose__1933206353 gopurs_runtime.Value
var once_compose__1933206353 sync.Once
func Get_compose__1933206353() gopurs_runtime.Value {
	once_compose__1933206353.Do(func() {
		cache_compose__1933206353 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1933206353(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__1933206353
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

var cache_bimap__132457202 gopurs_runtime.Value
var once_bimap__132457202 sync.Once
func Get_bimap__132457202() gopurs_runtime.Value {
	once_bimap__132457202.Do(func() {
		cache_bimap__132457202 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__132457202(gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__132457202
}

var cache_bimap__2801350668 gopurs_runtime.Value
var once_bimap__2801350668 sync.Once
func Get_bimap__2801350668() gopurs_runtime.Value {
	once_bimap__2801350668.Do(func() {
		cache_bimap__2801350668 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__2801350668(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_bimap__2801350668
}

var cache_eq__789642299 gopurs_runtime.Value
var once_eq__789642299 sync.Once
func Get_eq__789642299() gopurs_runtime.Value {
	once_eq__789642299.Do(func() {
		cache_eq__789642299 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__789642299(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[bool]](dict_0_box))
})
	})
	return cache_eq__789642299
}

var cache_eq__1697837627 gopurs_runtime.Value
var once_eq__1697837627 sync.Once
func Get_eq__1697837627() gopurs_runtime.Value {
	once_eq__1697837627.Do(func() {
		cache_eq__1697837627 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__1697837627(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__1697837627
}

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2843686287(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2843686287
}

var cache_eq__2276491096 gopurs_runtime.Value
var once_eq__2276491096 sync.Once
func Get_eq__2276491096() gopurs_runtime.Value {
	once_eq__2276491096.Do(func() {
		cache_eq__2276491096 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2276491096(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2276491096
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

var cache_eq__1425708671 gopurs_runtime.Value
var once_eq__1425708671 sync.Once
func Get_eq__1425708671() gopurs_runtime.Value {
	once_eq__1425708671.Do(func() {
		cache_eq__1425708671 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__1425708671(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_eq__1425708671
}

var cache_eq__1272715810 gopurs_runtime.Value
var once_eq__1272715810 sync.Once
func Get_eq__1272715810() gopurs_runtime.Value {
	once_eq__1272715810.Do(func() {
		cache_eq__1272715810 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_eq__1272715810(uint32(v_0_box.IntVal), uint32(v1_1_box.IntVal)))
})
	})
	return cache_eq__1272715810
}

var cache_notEq__2384498378 gopurs_runtime.Value
var once_notEq__2384498378 sync.Once
func Get_notEq__2384498378() gopurs_runtime.Value {
	once_notEq__2384498378.Do(func() {
		cache_notEq__2384498378 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_notEq__2384498378
}

var cache_notEq__1272715810 gopurs_runtime.Value
var once_notEq__1272715810 sync.Once
func Get_notEq__1272715810() gopurs_runtime.Value {
	once_notEq__1272715810.Do(func() {
		cache_notEq__1272715810 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__1272715810(x_0_box, y_1_box))
})
	})
	return cache_notEq__1272715810
}

var cache_any__4179648253 gopurs_runtime.Value
var once_any__4179648253 sync.Once
func Get_any__4179648253() gopurs_runtime.Value {
	once_any__4179648253.Do(func() {
		cache_any__4179648253 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_any__4179648253(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_any__4179648253
}

var cache_any__842931401 gopurs_runtime.Value
var once_any__842931401 sync.Once
func Get_any__842931401() gopurs_runtime.Value {
	once_any__842931401.Do(func() {
		cache_any__842931401 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_any__842931401(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_any__842931401
}

var cache_foldl__3041692656 gopurs_runtime.Value
var once_foldl__3041692656 sync.Once
func Get_foldl__3041692656() gopurs_runtime.Value {
	once_foldl__3041692656.Do(func() {
		cache_foldl__3041692656 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3041692656(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__3041692656
}

var cache_foldl__66388714 gopurs_runtime.Value
var once_foldl__66388714 sync.Once
func Get_foldl__66388714() gopurs_runtime.Value {
	once_foldl__66388714.Do(func() {
		cache_foldl__66388714 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__66388714(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__66388714
}

var cache_foldl__2602334544 gopurs_runtime.Value
var once_foldl__2602334544 sync.Once
func Get_foldl__2602334544() gopurs_runtime.Value {
	once_foldl__2602334544.Do(func() {
		cache_foldl__2602334544 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2602334544(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__2602334544
}

var cache_foldl__165683952 gopurs_runtime.Value
var once_foldl__165683952 sync.Once
func Get_foldl__165683952() gopurs_runtime.Value {
	once_foldl__165683952.Do(func() {
		cache_foldl__165683952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__165683952(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__165683952
}

var cache_foldl__1148906672 gopurs_runtime.Value
var once_foldl__1148906672 sync.Once
func Get_foldl__1148906672() gopurs_runtime.Value {
	once_foldl__1148906672.Do(func() {
		cache_foldl__1148906672 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1148906672(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__1148906672
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

var cache_foldl__1754241693 gopurs_runtime.Value
var once_foldl__1754241693 sync.Once
func Get_foldl__1754241693() gopurs_runtime.Value {
	once_foldl__1754241693.Do(func() {
		cache_foldl__1754241693 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1754241693(f_0_box)
})
	})
	return cache_foldl__1754241693
}

var cache_foldl__3943124669 gopurs_runtime.Value
var once_foldl__3943124669 sync.Once
func Get_foldl__3943124669() gopurs_runtime.Value {
	once_foldl__3943124669.Do(func() {
		cache_foldl__3943124669 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3943124669(f_0_box)
})
	})
	return cache_foldl__3943124669
}

var cache_foldl__396932925 gopurs_runtime.Value
var once_foldl__396932925 sync.Once
func Get_foldl__396932925() gopurs_runtime.Value {
	once_foldl__396932925.Do(func() {
		cache_foldl__396932925 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__396932925(f_0_box)
})
	})
	return cache_foldl__396932925
}

var cache_foldl__3915700701 gopurs_runtime.Value
var once_foldl__3915700701 sync.Once
func Get_foldl__3915700701() gopurs_runtime.Value {
	once_foldl__3915700701.Do(func() {
		cache_foldl__3915700701 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3915700701(f_0_box)
})
	})
	return cache_foldl__3915700701
}

var cache_foldr__2512763050 gopurs_runtime.Value
var once_foldr__2512763050 sync.Once
func Get_foldr__2512763050() gopurs_runtime.Value {
	once_foldr__2512763050.Do(func() {
		cache_foldr__2512763050 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2512763050(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2512763050
}

var cache_foldr__3673994608 gopurs_runtime.Value
var once_foldr__3673994608 sync.Once
func Get_foldr__3673994608() gopurs_runtime.Value {
	once_foldr__3673994608.Do(func() {
		cache_foldr__3673994608 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3673994608(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldr__3673994608
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

var cache_foldr__2829803163 gopurs_runtime.Value
var once_foldr__2829803163 sync.Once
func Get_foldr__2829803163() gopurs_runtime.Value {
	once_foldr__2829803163.Do(func() {
		cache_foldr__2829803163 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2829803163(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2829803163
}

var cache_foldr__2979608669 gopurs_runtime.Value
var once_foldr__2979608669 sync.Once
func Get_foldr__2979608669() gopurs_runtime.Value {
	once_foldr__2979608669.Do(func() {
		cache_foldr__2979608669 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2979608669(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_1_box))
})
	})
	return cache_foldr__2979608669
}

var cache_foldr__3489910557 gopurs_runtime.Value
var once_foldr__3489910557 sync.Once
func Get_foldr__3489910557() gopurs_runtime.Value {
	once_foldr__3489910557.Do(func() {
		cache_foldr__3489910557 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3489910557(f_0_box, b_1_box)
})
	})
	return cache_foldr__3489910557
}

var cache_foldr__3234921885 gopurs_runtime.Value
var once_foldr__3234921885 sync.Once
func Get_foldr__3234921885() gopurs_runtime.Value {
	once_foldr__3234921885.Do(func() {
		cache_foldr__3234921885 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3234921885(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](b_1_box))
})
	})
	return cache_foldr__3234921885
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

var cache_flip__3709724320 gopurs_runtime.Value
var once_flip__3709724320 sync.Once
func Get_flip__3709724320() gopurs_runtime.Value {
	once_flip__3709724320.Do(func() {
		cache_flip__3709724320 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3709724320(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3709724320
}

var cache_flip__3563101792 gopurs_runtime.Value
var once_flip__3563101792 sync.Once
func Get_flip__3563101792() gopurs_runtime.Value {
	once_flip__3563101792.Do(func() {
		cache_flip__3563101792 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3563101792(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3563101792
}

var cache_flip__2270521984 gopurs_runtime.Value
var once_flip__2270521984 sync.Once
func Get_flip__2270521984() gopurs_runtime.Value {
	once_flip__2270521984.Do(func() {
		cache_flip__2270521984 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2270521984(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2270521984
}

var cache_flip__2106496000 gopurs_runtime.Value
var once_flip__2106496000 sync.Once
func Get_flip__2106496000() gopurs_runtime.Value {
	once_flip__2106496000.Do(func() {
		cache_flip__2106496000 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2106496000(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2106496000
}

var cache_flip__2373571712 gopurs_runtime.Value
var once_flip__2373571712 sync.Once
func Get_flip__2373571712() gopurs_runtime.Value {
	once_flip__2373571712.Do(func() {
		cache_flip__2373571712 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2373571712(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2373571712
}

var cache_flip__3192966848 gopurs_runtime.Value
var once_flip__3192966848 sync.Once
func Get_flip__3192966848() gopurs_runtime.Value {
	once_flip__3192966848.Do(func() {
		cache_flip__3192966848 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3192966848(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3192966848
}

var cache_flip__2681340064 gopurs_runtime.Value
var once_flip__2681340064 sync.Once
func Get_flip__2681340064() gopurs_runtime.Value {
	once_flip__2681340064.Do(func() {
		cache_flip__2681340064 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2681340064(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2681340064
}

var cache_flip__872296768 gopurs_runtime.Value
var once_flip__872296768 sync.Once
func Get_flip__872296768() gopurs_runtime.Value {
	once_flip__872296768.Do(func() {
		cache_flip__872296768 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__872296768(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__872296768
}

var cache_flip__4091748192 gopurs_runtime.Value
var once_flip__4091748192 sync.Once
func Get_flip__4091748192() gopurs_runtime.Value {
	once_flip__4091748192.Do(func() {
		cache_flip__4091748192 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__4091748192(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__4091748192
}

var cache_map__3116241637 gopurs_runtime.Value
var once_map__3116241637 sync.Once
func Get_map__3116241637() gopurs_runtime.Value {
	once_map__3116241637.Do(func() {
		cache_map__3116241637 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3116241637(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_map__3116241637
}

var cache_map__2876470885 gopurs_runtime.Value
var once_map__2876470885 sync.Once
func Get_map__2876470885() gopurs_runtime.Value {
	once_map__2876470885.Do(func() {
		cache_map__2876470885 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2876470885(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2876470885
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

var cache_map__3778170420 gopurs_runtime.Value
var once_map__3778170420 sync.Once
func Get_map__3778170420() gopurs_runtime.Value {
	once_map__3778170420.Do(func() {
		cache_map__3778170420 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3778170420(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3778170420
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__4242765044 gopurs_runtime.Value
var once_map__4242765044 sync.Once
func Get_map__4242765044() gopurs_runtime.Value {
	once_map__4242765044.Do(func() {
		cache_map__4242765044 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__4242765044(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__4242765044
}

var cache_map__3058795348 gopurs_runtime.Value
var once_map__3058795348 sync.Once
func Get_map__3058795348() gopurs_runtime.Value {
	once_map__3058795348.Do(func() {
		cache_map__3058795348 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3058795348(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3058795348
}

var cache_map__438443400 gopurs_runtime.Value
var once_map__438443400 sync.Once
func Get_map__438443400() gopurs_runtime.Value {
	once_map__438443400.Do(func() {
		cache_map__438443400 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__438443400(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__438443400
}

var cache_map__291265340 gopurs_runtime.Value
var once_map__291265340 sync.Once
func Get_map__291265340() gopurs_runtime.Value {
	once_map__291265340.Do(func() {
		cache_map__291265340 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__291265340(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v1_1_box)))}
})
	})
	return cache_map__291265340
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__901270812(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__901270812
}

var cache_map__140514012 gopurs_runtime.Value
var once_map__140514012 sync.Once
func Get_map__140514012() gopurs_runtime.Value {
	once_map__140514012.Do(func() {
		cache_map__140514012 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__140514012(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v1_1_box)))}
})
	})
	return cache_map__140514012
}

var cache_map__2202537180 gopurs_runtime.Value
var once_map__2202537180 sync.Once
func Get_map__2202537180() gopurs_runtime.Value {
	once_map__2202537180.Do(func() {
		cache_map__2202537180 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__2202537180(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__2202537180
}

var cache_map__1681779388 gopurs_runtime.Value
var once_map__1681779388 sync.Once
func Get_map__1681779388() gopurs_runtime.Value {
	once_map__1681779388.Do(func() {
		cache_map__1681779388 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__1681779388(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__1681779388
}

var cache_map__3486165692 gopurs_runtime.Value
var once_map__3486165692 sync.Once
func Get_map__3486165692() gopurs_runtime.Value {
	once_map__3486165692.Do(func() {
		cache_map__3486165692 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__3486165692(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__3486165692
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_emptySet__2398681994 gopurs_runtime.Value
var once_emptySet__2398681994 sync.Once
func Get_emptySet__2398681994() gopurs_runtime.Value {
	once_emptySet__2398681994.Do(func() {
		cache_emptySet__2398681994 = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_emptySet__2398681994
}

var cache_fromZipper__1019554324 gopurs_runtime.Value
var once_fromZipper__1019554324 sync.Once
func Get_fromZipper__1019554324() gopurs_runtime.Value {
	once_fromZipper__1019554324.Do(func() {
		cache_fromZipper__1019554324 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromZipper__1019554324(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box)
})
	})
	return cache_fromZipper__1019554324
}

var cache_insertAndLookupBy__3244745033 gopurs_runtime.Value
var once_insertAndLookupBy__3244745033 sync.Once
func Get_insertAndLookupBy__3244745033() gopurs_runtime.Value {
	once_insertAndLookupBy__3244745033.Do(func() {
		cache_insertAndLookupBy__3244745033 = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, orig_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAndLookupBy__3244745033(comp_0_box, k_1_box, orig_2_box)
})
	})
	return cache_insertAndLookupBy__3244745033
}

var cache_applyList__3072763993 gopurs_runtime.Value
var once_applyList__3072763993 sync.Once
func Get_applyList__3072763993() gopurs_runtime.Value {
	once_applyList__3072763993.Do(func() {
		cache_applyList__3072763993 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList__3072763993
}

var cache_applyList__1109325167 gopurs_runtime.Value
var once_applyList__1109325167 sync.Once
func Get_applyList__1109325167() gopurs_runtime.Value {
	once_applyList__1109325167.Do(func() {
		cache_applyList__1109325167 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList__1109325167
}

var cache_bindList__241263065 gopurs_runtime.Value
var once_bindList__241263065 sync.Once
func Get_bindList__241263065() gopurs_runtime.Value {
	once_bindList__241263065.Do(func() {
		cache_bindList__241263065 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_applyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindList__241263065
}

var cache_foldableList__1753400174 gopurs_runtime.Value
var once_foldableList__1753400174 sync.Once
func Get_foldableList__1753400174() gopurs_runtime.Value {
	once_foldableList__1753400174.Do(func() {
		cache_foldableList__1753400174 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_39 gopurs_runtime.Value
go__go_1_3_39 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_39:
for {
if false { continue go__go_1_3_39 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t4 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_39
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
return go__go_1_3_39
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_40 gopurs_runtime.Value
go__go_3_7_40 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_40:
for {
if false { continue go__go_3_7_40 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_40
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_40, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_foldableList__1753400174
}

var cache_foldableList__46095397 gopurs_runtime.Value
var once_foldableList__46095397 sync.Once
func Get_foldableList__46095397() gopurs_runtime.Value {
	once_foldableList__46095397.Do(func() {
		cache_foldableList__46095397 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_41 gopurs_runtime.Value
go__go_1_3_41 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_41:
for {
if false { continue go__go_1_3_41 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t4 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_41
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
return go__go_1_3_41
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_42 gopurs_runtime.Value
go__go_3_7_42 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_42:
for {
if false { continue go__go_3_7_42 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_42
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_42, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_foldableList__46095397
}

var cache_functorList__4121998062 gopurs_runtime.Value
var once_functorList__4121998062 sync.Once
func Get_functorList__4121998062() gopurs_runtime.Value {
	once_functorList__4121998062.Do(func() {
		cache_functorList__4121998062 = gopurs_runtime.RecordDict1("map", pkg_Data_List_Types.Get_listMap())
	})
	return cache_functorList__4121998062
}

var cache_functorList__1783129585 gopurs_runtime.Value
var once_functorList__1783129585 sync.Once
func Get_functorList__1783129585() gopurs_runtime.Value {
	once_functorList__1783129585.Do(func() {
		cache_functorList__1783129585 = gopurs_runtime.RecordDict1("map", pkg_Data_List_Types.Get_listMap())
	})
	return cache_functorList__1783129585
}

var cache_listMap__4135416762 gopurs_runtime.Value
var once_listMap__4135416762 sync.Once
func Get_listMap__4135416762() gopurs_runtime.Value {
	once_listMap__4135416762.Do(func() {
		cache_listMap__4135416762 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listMap__4135416762(f_0_box)
})
	})
	return cache_listMap__4135416762
}

var cache_semigroupList__2766094215 gopurs_runtime.Value
var once_semigroupList__2766094215 sync.Once
func Get_semigroupList__2766094215() gopurs_runtime.Value {
	once_semigroupList__2766094215.Do(func() {
		cache_semigroupList__2766094215 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_semigroupList__2766094215
}

var cache_traversableList__3361437934 gopurs_runtime.Value
var once_traversableList__3361437934 sync.Once
func Get_traversableList__3361437934() gopurs_runtime.Value {
	once_traversableList__3361437934.Do(func() {
		cache_traversableList__3361437934 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_List_Types.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_6_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_6_5
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_6_5.V0, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_traversableList__3361437934
}

var cache_alterAt__3453373293 gopurs_runtime.Value
var once_alterAt__3453373293 sync.Once
func Get_alterAt__3453373293() gopurs_runtime.Value {
	once_alterAt__3453373293.Do(func() {
		cache_alterAt__3453373293 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_alterAt__3453373293(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_alterAt__3453373293
}

var cache_deleteAt__2845095501 gopurs_runtime.Value
var once_deleteAt__2845095501 sync.Once
func Get_deleteAt__2845095501() gopurs_runtime.Value {
	once_deleteAt__2845095501.Do(func() {
		cache_deleteAt__2845095501 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_deleteAt__2845095501(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_deleteAt__2845095501
}

var cache_deleteBy__697302515 gopurs_runtime.Value
var once_deleteBy__697302515 sync.Once
func Get_deleteBy__697302515() gopurs_runtime.Value {
	once_deleteBy__697302515.Do(func() {
		cache_deleteBy__697302515 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_deleteBy__697302515(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_deleteBy__697302515
}

var cache_drop__1836090668 gopurs_runtime.Value
var once_drop__1836090668 sync.Once
func Get_drop__1836090668() gopurs_runtime.Value {
	once_drop__1836090668.Do(func() {
		cache_drop__1836090668 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_drop__1836090668(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_drop__1836090668
}

var cache_filter__1617261107 gopurs_runtime.Value
var once_filter__1617261107 sync.Once
func Get_filter__1617261107() gopurs_runtime.Value {
	once_filter__1617261107.Do(func() {
		cache_filter__1617261107 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter__1617261107(p_0_box)
})
	})
	return cache_filter__1617261107
}

var cache_filterM__14771079 gopurs_runtime.Value
var once_filterM__14771079 sync.Once
func Get_filterM__14771079() gopurs_runtime.Value {
	once_filterM__14771079.Do(func() {
		cache_filterM__14771079 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterM__14771079(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_filterM__14771079
}

var cache_findIndex__2366045378 gopurs_runtime.Value
var once_findIndex__2366045378 sync.Once
func Get_findIndex__2366045378() gopurs_runtime.Value {
	once_findIndex__2366045378.Do(func() {
		cache_findIndex__2366045378 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findIndex__2366045378(fn_0_box)
})
	})
	return cache_findIndex__2366045378
}

var cache_findLastIndex__2366045378 gopurs_runtime.Value
var once_findLastIndex__2366045378 sync.Once
func Get_findLastIndex__2366045378() gopurs_runtime.Value {
	once_findLastIndex__2366045378.Do(func() {
		cache_findLastIndex__2366045378 = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex__2366045378(fn_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box)))}
})
	})
	return cache_findLastIndex__2366045378
}

var cache_foldM__3577257629 gopurs_runtime.Value
var once_foldM__3577257629 sync.Once
func Get_foldM__3577257629() gopurs_runtime.Value {
	once_foldM__3577257629.Do(func() {
		cache_foldM__3577257629 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM__3577257629(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldM__3577257629
}

var cache_groupBy__1039549870 gopurs_runtime.Value
var once_groupBy__1039549870 sync.Once
func Get_groupBy__1039549870() gopurs_runtime.Value {
	once_groupBy__1039549870.Do(func() {
		cache_groupBy__1039549870 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_groupBy__1039549870(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_groupBy__1039549870
}

var cache_head__3729839663 gopurs_runtime.Value
var once_head__3729839663 sync.Once
func Get_head__3729839663() gopurs_runtime.Value {
	once_head__3729839663.Do(func() {
		cache_head__3729839663 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_head__3729839663(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_head__3729839663
}

var cache_index__304299960 gopurs_runtime.Value
var once_index__304299960 sync.Once
func Get_index__304299960() gopurs_runtime.Value {
	once_index__304299960.Do(func() {
		cache_index__304299960 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index__304299960(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box.IntVal))}
})
	})
	return cache_index__304299960
}

var cache_insertAt__2634211748 gopurs_runtime.Value
var once_insertAt__2634211748 sync.Once
func Get_insertAt__2634211748() gopurs_runtime.Value {
	once_insertAt__2634211748.Do(func() {
		cache_insertAt__2634211748 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt__2634211748(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_insertAt__2634211748
}

var cache_insertBy__1738998985 gopurs_runtime.Value
var once_insertBy__1738998985 sync.Once
func Get_insertBy__1738998985() gopurs_runtime.Value {
	once_insertBy__1738998985.Do(func() {
		cache_insertBy__1738998985 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_insertBy__1738998985(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_insertBy__1738998985
}

var cache_intersectBy__1190504998 gopurs_runtime.Value
var once_intersectBy__1190504998 sync.Once
func Get_intersectBy__1190504998() gopurs_runtime.Value {
	once_intersectBy__1190504998.Do(func() {
		cache_intersectBy__1190504998 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_intersectBy__1190504998(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_intersectBy__1190504998
}

var cache_last__4043133652 gopurs_runtime.Value
var once_last__4043133652 sync.Once
func Get_last__4043133652() gopurs_runtime.Value {
	once_last__4043133652.Do(func() {
		cache_last__4043133652 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_last__4043133652(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_last__4043133652
}

var cache_length__3003998832 gopurs_runtime.Value
var once_length__3003998832 sync.Once
func Get_length__3003998832() gopurs_runtime.Value {
	once_length__3003998832.Do(func() {
		cache_length__3003998832 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(acc_0.IntVal), gopurs_runtime.Int(1)).IntVal)
})
}), gopurs_runtime.Int(0))
	})
	return cache_length__3003998832
}

var cache_many__542682753 gopurs_runtime.Value
var once_many__542682753 sync.Once
func Get_many__542682753() gopurs_runtime.Value {
	once_many__542682753.Do(func() {
		cache_many__542682753 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many__542682753(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_many__542682753
}

var cache_manyRec__4046352885 gopurs_runtime.Value
var once_manyRec__4046352885 sync.Once
func Get_manyRec__4046352885() gopurs_runtime.Value {
	once_manyRec__4046352885.Do(func() {
		cache_manyRec__4046352885 = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_manyRec__4046352885(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_manyRec__4046352885
}

var cache_mapMaybe__1486753757 gopurs_runtime.Value
var once_mapMaybe__1486753757 sync.Once
func Get_mapMaybe__1486753757() gopurs_runtime.Value {
	once_mapMaybe__1486753757.Do(func() {
		cache_mapMaybe__1486753757 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__1486753757(f_0_box)
})
	})
	return cache_mapMaybe__1486753757
}

var cache_mapMaybe__748617661 gopurs_runtime.Value
var once_mapMaybe__748617661 sync.Once
func Get_mapMaybe__748617661() gopurs_runtime.Value {
	once_mapMaybe__748617661.Do(func() {
		cache_mapMaybe__748617661 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__748617661(f_0_box)
})
	})
	return cache_mapMaybe__748617661
}

var cache_mapMaybe__4251473821 gopurs_runtime.Value
var once_mapMaybe__4251473821 sync.Once
func Get_mapMaybe__4251473821() gopurs_runtime.Value {
	once_mapMaybe__4251473821.Do(func() {
		cache_mapMaybe__4251473821 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__4251473821(f_0_box)
})
	})
	return cache_mapMaybe__4251473821
}

var cache_mapMaybe__2491277821 gopurs_runtime.Value
var once_mapMaybe__2491277821 sync.Once
func Get_mapMaybe__2491277821() gopurs_runtime.Value {
	once_mapMaybe__2491277821.Do(func() {
		cache_mapMaybe__2491277821 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__2491277821(f_0_box)
})
	})
	return cache_mapMaybe__2491277821
}

var cache_nubBy__1502591776 gopurs_runtime.Value
var once_nubBy__1502591776 sync.Once
func Get_nubBy__1502591776() gopurs_runtime.Value {
	once_nubBy__1502591776.Do(func() {
		cache_nubBy__1502591776 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubBy__1502591776(p_0_box)
})
	})
	return cache_nubBy__1502591776
}

var cache_nubByEq__3655321914 gopurs_runtime.Value
var once_nubByEq__3655321914 sync.Once
func Get_nubByEq__3655321914() gopurs_runtime.Value {
	once_nubByEq__3655321914.Do(func() {
		cache_nubByEq__3655321914 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_nubByEq__3655321914(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_nubByEq__3655321914
}

var cache_reverse__1174136571 gopurs_runtime.Value
var once_reverse__1174136571 sync.Once
func Get_reverse__1174136571() gopurs_runtime.Value {
	once_reverse__1174136571.Do(func() {
		cache_reverse__1174136571 = func() gopurs_runtime.Value {
var go__go_0_0_60 gopurs_runtime.Value
go__go_0_0_60 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_60:
for {
if false { continue go__go_0_0_60 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
continue go__go_0_0_60
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_60, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_reverse__1174136571
}

var cache_reverse__4230102656 gopurs_runtime.Value
var once_reverse__4230102656 sync.Once
func Get_reverse__4230102656() gopurs_runtime.Value {
	once_reverse__4230102656.Do(func() {
		cache_reverse__4230102656 = func() gopurs_runtime.Value {
var go__go_0_0_61 gopurs_runtime.Value
go__go_0_0_61 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_61:
for {
if false { continue go__go_0_0_61 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
continue go__go_0_0_61
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_61, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_reverse__4230102656
}

var cache_singleton__2450819477 gopurs_runtime.Value
var once_singleton__2450819477 sync.Once
func Get_singleton__2450819477() gopurs_runtime.Value {
	once_singleton__2450819477.Do(func() {
		cache_singleton__2450819477 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_singleton__2450819477(a_0_box))}
})
	})
	return cache_singleton__2450819477
}

var cache_singleton__707062261 gopurs_runtime.Value
var once_singleton__707062261 sync.Once
func Get_singleton__707062261() gopurs_runtime.Value {
	once_singleton__707062261.Do(func() {
		cache_singleton__707062261 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_singleton__707062261(a_0_box))}
})
	})
	return cache_singleton__707062261
}

var cache_singleton__3932757557 gopurs_runtime.Value
var once_singleton__3932757557 sync.Once
func Get_singleton__3932757557() gopurs_runtime.Value {
	once_singleton__3932757557.Do(func() {
		cache_singleton__3932757557 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_singleton__3932757557(a_0_box))}
})
	})
	return cache_singleton__3932757557
}

var cache_some__542682753 gopurs_runtime.Value
var once_some__542682753 sync.Once
func Get_some__542682753() gopurs_runtime.Value {
	once_some__542682753.Do(func() {
		cache_some__542682753 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some__542682753(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_some__542682753
}

var cache_sortBy__1502591776 gopurs_runtime.Value
var once_sortBy__1502591776 sync.Once
func Get_sortBy__1502591776() gopurs_runtime.Value {
	once_sortBy__1502591776.Do(func() {
		cache_sortBy__1502591776 = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy__1502591776(cmp_0_box)
})
	})
	return cache_sortBy__1502591776
}

var cache_span__799093643 gopurs_runtime.Value
var once_span__799093643 sync.Once
func Get_span__799093643() gopurs_runtime.Value {
	once_span__799093643.Do(func() {
		cache_span__799093643 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span__799093643(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box))
})
	})
	return cache_span__799093643
}

var cache_tail__1771843450 gopurs_runtime.Value
var once_tail__1771843450 sync.Once
func Get_tail__1771843450() gopurs_runtime.Value {
	once_tail__1771843450.Do(func() {
		cache_tail__1771843450 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_tail__1771843450(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_tail__1771843450
}

var cache_tails__3932757557 gopurs_runtime.Value
var once_tails__3932757557 sync.Once
func Get_tails__3932757557() gopurs_runtime.Value {
	once_tails__3932757557.Do(func() {
		cache_tails__3932757557 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_tails__3932757557(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_tails__3932757557
}

var cache_take__1836090668 gopurs_runtime.Value
var once_take__1836090668 sync.Once
func Get_take__1836090668() gopurs_runtime.Value {
	once_take__1836090668.Do(func() {
		cache_take__1836090668 = func() gopurs_runtime.Value {
var go__go_0_0_68 gopurs_runtime.Value
go__go_0_0_68 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop int64 = v1_2_loop_val.IntVal
var v2_3_loop gopurs_runtime.Value = v2_3_loop_val
go__go_0_0_68:
for {
if false { continue go__go_0_0_68 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 int64 = v1_2_loop
_ = v1_2
var v2_3 gopurs_runtime.Value = v2_3_loop
_ = v2_3
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v1_2), gopurs_runtime.Int(1))).IntVal) != (0) {
var go__go_4_1_69 gopurs_runtime.Value
go__go_4_1_69 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_69:
for {
if false { continue go__go_4_1_69 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_69
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_69, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})))}
goto end_branch_5
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr == nil) {
var go__go_4_3_70 gopurs_runtime.Value
go__go_4_3_70 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_3_70:
for {
if false { continue go__go_4_3_70 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_3_70
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_3_70, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})))}
goto end_branch_5
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_3.UnsafePtr).V0, v_1})})
v1_2_loop = Call_sub__1043827704(gopurs_runtime.Int(v1_2), gopurs_runtime.Int(1)).IntVal
v2_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_3.UnsafePtr).V1)}
continue go__go_0_0_68
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
})
return gopurs_runtime.Apply(go__go_0_0_68, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_take__1836090668
}

var cache_transpose__682228544 gopurs_runtime.Value
var once_transpose__682228544 sync.Once
func Get_transpose__682228544() gopurs_runtime.Value {
	once_transpose__682228544.Do(func() {
		cache_transpose__682228544 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_transpose__682228544(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_0_box)))}
})
	})
	return cache_transpose__682228544
}

var cache_uncons__3009258782 gopurs_runtime.Value
var once_uncons__3009258782 sync.Once
func Get_uncons__3009258782() gopurs_runtime.Value {
	once_uncons__3009258782.Do(func() {
		cache_uncons__3009258782 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__3009258782(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_uncons__3009258782
}

var cache_unionBy__1190504998 gopurs_runtime.Value
var once_unionBy__1190504998 sync.Once
func Get_unionBy__1190504998() gopurs_runtime.Value {
	once_unionBy__1190504998.Do(func() {
		cache_unionBy__1190504998 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_unionBy__1190504998(eq_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_unionBy__1190504998
}

var cache_unsnoc__2942606998 gopurs_runtime.Value
var once_unsnoc__2942606998 sync.Once
func Get_unsnoc__2942606998() gopurs_runtime.Value {
	once_unsnoc__2942606998.Do(func() {
		cache_unsnoc__2942606998 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_unsnoc__2942606998(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](lst_0_box)))}
})
	})
	return cache_unsnoc__2942606998
}

var cache_updateAt__2634211748 gopurs_runtime.Value
var once_updateAt__2634211748 sync.Once
func Get_updateAt__2634211748() gopurs_runtime.Value {
	once_updateAt__2634211748.Do(func() {
		cache_updateAt__2634211748 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt__2634211748(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_updateAt__2634211748
}

var cache_zipWith__884793877 gopurs_runtime.Value
var once_zipWith__884793877 sync.Once
func Get_zipWith__884793877() gopurs_runtime.Value {
	once_zipWith__884793877.Do(func() {
		cache_zipWith__884793877 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_zipWith__884793877(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_zipWith__884793877
}

var cache_zipWith__4203240021 gopurs_runtime.Value
var once_zipWith__4203240021 sync.Once
func Get_zipWith__4203240021() gopurs_runtime.Value {
	once_zipWith__4203240021.Do(func() {
		cache_zipWith__4203240021 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_zipWith__4203240021(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_zipWith__4203240021
}

var cache_zipWith__3856182069 gopurs_runtime.Value
var once_zipWith__3856182069 sync.Once
func Get_zipWith__3856182069() gopurs_runtime.Value {
	once_zipWith__3856182069.Do(func() {
		cache_zipWith__3856182069 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_zipWith__3856182069(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_zipWith__3856182069
}

var cache_applicativeMaybe__500933224 gopurs_runtime.Value
var once_applicativeMaybe__500933224 sync.Once
func Get_applicativeMaybe__500933224() gopurs_runtime.Value {
	once_applicativeMaybe__500933224.Do(func() {
		cache_applicativeMaybe__500933224 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), pkg_Data_Maybe.Get_Just())
	})
	return cache_applicativeMaybe__500933224
}

var cache_applyMaybe__3698865467 gopurs_runtime.Value
var once_applyMaybe__3698865467 sync.Once
func Get_applyMaybe__3698865467() gopurs_runtime.Value {
	once_applyMaybe__3698865467.Do(func() {
		cache_applyMaybe__3698865467 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))})))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3698865467
}

var cache_bindMaybe__3591110311 gopurs_runtime.Value
var once_bindMaybe__3591110311 sync.Once
func Get_bindMaybe__3591110311() gopurs_runtime.Value {
	once_bindMaybe__3591110311.Do(func() {
		cache_bindMaybe__3591110311 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindMaybe__3591110311
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_monadMaybe__3072900051 gopurs_runtime.Value
var once_monadMaybe__3072900051 sync.Once
func Get_monadMaybe__3072900051() gopurs_runtime.Value {
	once_monadMaybe__3072900051.Do(func() {
		cache_monadMaybe__3072900051 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_bindMaybe()
}))
	})
	return cache_monadMaybe__3072900051
}

var cache_alaF__2824451980 gopurs_runtime.Value
var once_alaF__2824451980 sync.Once
func Get_alaF__2824451980() gopurs_runtime.Value {
	once_alaF__2824451980.Do(func() {
		cache_alaF__2824451980 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF__2824451980(_dollar__unused_0_box, _dollar__unused_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_2_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_3_box), v_4_box)
})
	})
	return cache_alaF__2824451980
}

var cache_alaF__4085337484 gopurs_runtime.Value
var once_alaF__4085337484 sync.Once
func Get_alaF__4085337484() gopurs_runtime.Value {
	once_alaF__4085337484.Do(func() {
		cache_alaF__4085337484 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF__4085337484(_dollar__unused_0_box, _dollar__unused_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_2_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_3_box), v_4_box)
})
	})
	return cache_alaF__4085337484
}

var cache_compare__669572705 gopurs_runtime.Value
var once_compare__669572705 sync.Once
func Get_compare__669572705() gopurs_runtime.Value {
	once_compare__669572705.Do(func() {
		cache_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__669572705(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__669572705
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

var cache_compare__3029065925 gopurs_runtime.Value
var once_compare__3029065925 sync.Once
func Get_compare__3029065925() gopurs_runtime.Value {
	once_compare__3029065925.Do(func() {
		cache_compare__3029065925 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__3029065925(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compare__3029065925
}

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_lessThan__4087042607 gopurs_runtime.Value
var once_lessThan__4087042607 sync.Once
func Get_lessThan__4087042607() gopurs_runtime.Value {
	once_lessThan__4087042607.Do(func() {
		cache_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__4087042607
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_negate__2635823316 gopurs_runtime.Value
var once_negate__2635823316 sync.Once
func Get_negate__2635823316() gopurs_runtime.Value {
	once_negate__2635823316.Do(func() {
		cache_negate__2635823316 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2635823316(__eta0_0_box)
})
	})
	return cache_negate__2635823316
}

var cache_negate__1364373265 gopurs_runtime.Value
var once_negate__1364373265 sync.Once
func Get_negate__1364373265() gopurs_runtime.Value {
	once_negate__1364373265.Do(func() {
		cache_negate__1364373265 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__1364373265(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate__1364373265
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

var cache_append__2832914972 gopurs_runtime.Value
var once_append__2832914972 sync.Once
func Get_append__2832914972() gopurs_runtime.Value {
	once_append__2832914972.Do(func() {
		cache_append__2832914972 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__2832914972(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_append__2832914972
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

var cache_append__2013893496 gopurs_runtime.Value
var once_append__2013893496 sync.Once
func Get_append__2013893496() gopurs_runtime.Value {
	once_append__2013893496.Do(func() {
		cache_append__2013893496 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_append__2013893496(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](xs_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](ys_1_box)))}
})
	})
	return cache_append__2013893496
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

var cache_zero__1556010056 gopurs_runtime.Value
var once_zero__1556010056 sync.Once
func Get_zero__1556010056() gopurs_runtime.Value {
	once_zero__1556010056.Do(func() {
		cache_zero__1556010056 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1556010056(dict_0_box)
})
	})
	return cache_zero__1556010056
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
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

var cache_show__2130238610 gopurs_runtime.Value
var once_show__2130238610 sync.Once
func Get_show__2130238610() gopurs_runtime.Value {
	once_show__2130238610.Do(func() {
		cache_show__2130238610 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2130238610(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_show__2130238610
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

var cache_sequence__2616145509 gopurs_runtime.Value
var once_sequence__2616145509 sync.Once
func Get_sequence__2616145509() gopurs_runtime.Value {
	once_sequence__2616145509.Do(func() {
		cache_sequence__2616145509 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__2616145509(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_sequence__2616145509
}

var cache_traverse__3956862083 gopurs_runtime.Value
var once_traverse__3956862083 sync.Once
func Get_traverse__3956862083() gopurs_runtime.Value {
	once_traverse__3956862083.Do(func() {
		cache_traverse__3956862083 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__3956862083(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_traverse__3956862083
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

var cache_traverse__2839486329 gopurs_runtime.Value
var once_traverse__2839486329 sync.Once
func Get_traverse__2839486329() gopurs_runtime.Value {
	once_traverse__2839486329.Do(func() {
		cache_traverse__2839486329 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__2839486329(dictApplicative_0_box)
})
	})
	return cache_traverse__2839486329
}

var cache_unfoldr__1128708256 gopurs_runtime.Value
var once_unfoldr__1128708256 sync.Once
func Get_unfoldr__1128708256() gopurs_runtime.Value {
	once_unfoldr__1128708256.Do(func() {
		cache_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1128708256
}

var cache_unfoldr__3827943605 gopurs_runtime.Value
var once_unfoldr__3827943605 sync.Once
func Get_unfoldr__3827943605() gopurs_runtime.Value {
	once_unfoldr__3827943605.Do(func() {
		cache_unfoldr__3827943605 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__3827943605(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__3827943605
}

func Call_identity(x_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var x_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Pattern(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_updateAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
updateAt:
for {
if false { continue updateAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))}))})))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t2)
}
}

func Call_uncons(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_toUnfoldable(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1))}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1))}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V1)})})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(rec_2, "tail")))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))})))}
}))
}

func Call_tail(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t0)
}

func Call_stripPrefix(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], s_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var s_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = s_2_loop
_ = s_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_Rec_Class.Get_monadRecMaybe(), "tailRecM"), gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "b")
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(dictEq_0.V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(o_3, "a").UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(o_3, "b").UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(o_3, "a").UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(o_3, "b").UnsafePtr).V1)})})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
__t4 = __t2
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(o_3, "b")))}})}})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1358893437 && __t_tag_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(o_3, "b")))}})}})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t6))}
}), gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(s_2)})))
}

func Call_span(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
span:
for {
if false { continue span }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0).IntVal) != (0)) {
v2_2_0 := Call_span(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)
_ = v2_2_0
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}
}

func Call_snoc(xs_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var xs_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_singleton(a_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})
}

func Call_sortBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0_0 gopurs_runtime.Value
_ = merge_1_0_0
merge_1_0_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_eq__1272715810(uint32(gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal), 380165415)).IntVal) != (0) {
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
goto end_branch_1
} else {

}
}
{
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3))}
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
})
})
var mergePairs_2_4_1 gopurs_runtime.Value
_ = mergePairs_2_4_1
mergePairs_2_4_1 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_and_6 bool = false
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {

var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
__t_and_6 = (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1358893437 && __t_tag_5.UnsafePtr != nil)
}
if __t_and_6 {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}.UnsafePtr).V0))})))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(mergePairs_2_4_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}.UnsafePtr).V1)})))})})}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_3))}
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t7))}
})
var mergeAll_3_8_2 gopurs_runtime.Value
mergeAll_3_8_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_4_loop_val)
mergeAll_3_8_2:
for {
if false { continue mergeAll_3_8_2 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_4_loop
_ = v_4
var __t11 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_10 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr != nil) {

var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr).V1)}
__t_and_10 = (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr == nil)
}
if __t_and_10 {
__t11 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(mergePairs_2_4_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
continue mergeAll_3_8_2
__t11 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}
}
}()
})
var sequences_4_12_3 gopurs_runtime.Value
_ = sequences_4_12_3
var descending_4_13_4 gopurs_runtime.Value
_ = descending_4_13_4
var ascending_4_14_5 gopurs_runtime.Value
_ = ascending_4_14_5
sequences_4_12_3 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
var __t_and_16 bool = false
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}
__t_and_16 = (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil)
}
if __t_and_16 {
var __t18 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]
{
if (gopurs_runtime.Bool(Call_eq__1272715810(uint32(gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0).IntVal), 380165415)).IntVal) != (0) {
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(descending_4_13_4, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V1)}))
goto end_branch_18
} else {

}
}
{
__local_var_6_17 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_17
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(ascending_4_14_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_17, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_7)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V1)}))
}
end_branch_18:
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t18)}
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t19))}
})
descending_4_13_4 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && ((gopurs_runtime.Bool(Call_eq__1272715810(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0).IntVal), 380165415)).IntVal) != (0)) {
__t20 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(descending_4_13_4, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_6)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))}
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_6)})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(sequences_4_12_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_7))})))})})}
}
end_branch_20:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t20))}
})
})
})
ascending_4_14_5 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && ((gopurs_runtime.Bool(Call_notEq__1272715810(gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})).IntVal) != (0)) {
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(ascending_4_14_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_8)})})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))}
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(sequences_4_12_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_7))})))})})}
}
end_branch_21:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t21))}
})
})
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_8_2, gopurs_runtime.Apply(sequences_4_12_3, x_5))
})
}

func Call_sort(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_sortBy(compare_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_2))})))}
})
}

func Call_tails(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
tails:
for {
if false { continue tails }
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_tails((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))}))})})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t0)
}
}

func Call_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showList_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_List_Types.Get_showList(), dictShow_0))
_ = showList_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Pattern "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(showList_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_10 gopurs_runtime.Value
go__go_1_0_10 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_10:
for {
if false { continue go__go_1_0_10 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if ((v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_10
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
var go__go_4_1_11 gopurs_runtime.Value
go__go_4_1_11 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_11:
for {
if false { continue go__go_4_1_11 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_11
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_11, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_unsnoc(lst_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var lst_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
var go__go_1_0_12 gopurs_runtime.Value
go__go_1_0_12 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
go__go_1_0_12:
for {
if false { continue go__go_1_0_12 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("last", "revInit", (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)})})}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3})})
continue go__go_1_0_12
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(h_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_13 gopurs_runtime.Value
go__go_3_4_13 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_4_13:
for {
if false { continue go__go_3_4_13 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t5 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_5
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_4_13
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_3_4_13, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(h_2, "revInit")))})))}, gopurs_runtime.RecordGet(h_2, "last"))
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_1_0_12, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}))
}

func Call_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var go__go_3_0_14 gopurs_runtime.Value
go__go_3_0_14 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_6_loop_val)
go__go_3_0_14:
for {
if false { continue go__go_3_0_14 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_6_loop
_ = v2_6
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0), v2_6})})
continue go__go_3_0_14
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
})
var go__go_4_2_15 gopurs_runtime.Value
go__go_4_2_15 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_15:
for {
if false { continue go__go_4_2_15 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_2_15
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_2_15, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(go__go_3_0_14, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}))
}

func Call_zipWithA(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, xs_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_2_loop
_ = xs_2
var ys_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversableList(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_zipWith(f_1, xs_2, ys_3))}))})
}

func Call_go__range(start_0_loop int64, end_1_loop int64) *pkg_Data_List_Types.Constructor_Cons[int64] {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var __t3 *pkg_Data_List_Types.Constructor_Cons[int64]
{
if (Call_eq__2843686287(gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1)).IntVal) != (0) {
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_singleton__2450819477(gopurs_runtime.Int(start_0)))})
goto end_branch_3
} else {

}
}
{
var go__go_2_0_16 gopurs_runtime.Value
go__go_2_0_16 = gopurs_runtime.Func(func(s_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(step_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rest_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var s_3_loop int64 = s_3_loop_val.IntVal
var e_4_loop int64 = e_4_loop_val.IntVal
var step_5_loop int64 = step_5_loop_val.IntVal
var rest_6_loop *pkg_Data_List_Types.Constructor_Cons[int64] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[int64]](rest_6_loop_val)
go__go_2_0_16:
for {
if false { continue go__go_2_0_16 }
var s_3 int64 = s_3_loop
_ = s_3
var e_4 int64 = e_4_loop
_ = e_4
var step_5 int64 = step_5_loop
_ = step_5
var rest_6 *pkg_Data_List_Types.Constructor_Cons[int64] = rest_6_loop
_ = rest_6
var __t1 gopurs_runtime.Value
{
if (Call_eq__2843686287(gopurs_runtime.Int(s_3), gopurs_runtime.Int(e_4)).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(s_3), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(rest_6)})})}
goto end_branch_1
} else {

}
}
{
s_3_loop = Call_add__560788792(gopurs_runtime.Int(s_3), gopurs_runtime.Int(step_5)).IntVal
e_4_loop = e_4
step_5_loop = step_5
rest_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(s_3), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(rest_6)})})})
continue go__go_2_0_16
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[int64]](gopurs_runtime.Value{}))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[int64]](__t1))}
}
}()
})
})
})
})
var __t2 int64
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1))).IntVal) != (0) {
__t2 = 1
goto end_branch_2
} else {

}
}
{
__t2 = Call_negate__2635823316(gopurs_runtime.Int(1)).IntVal
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[int64]](gopurs_runtime.Apply4(go__go_2_0_16, gopurs_runtime.Int(end_1), gopurs_runtime.Int(start_0), gopurs_runtime.Int(__t2), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}))
}
end_branch_3:
return __t3
}

func Call_partition(p_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, x_2).IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "no")))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "yes"))})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "no"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "yes")))})
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})
}

func Call_null(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) bool {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_17 gopurs_runtime.Value
go__go_1_0_17 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
var v2_4_loop gopurs_runtime.Value = v2_4_loop_val
go__go_1_0_17:
for {
if false { continue go__go_1_0_17 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var v2_4 gopurs_runtime.Value = v2_4_loop
_ = v2_4
var __t3 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr != nil) {
v3_5_1 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V0, v_2)
_ = v3_5_1
var __t2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.RecordGet(v3_5_1, "found").IntVal) != (0) {
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = v1_3
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V1)}
continue go__go_1_0_17
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V0, v1_3})})
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V1)}
continue go__go_1_0_17
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
})
__local_var_2_4 := gopurs_runtime.Apply2(go__go_1_0_17, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_18 gopurs_runtime.Value
go__go_4_5_18 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_18:
for {
if false { continue go__go_4_5_18 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t6 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_5_18
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t6))}
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_5_18, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Apply(__local_var_2_4, x_3))
})
}

func Call_nub(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_nubBy(dictOrd_0.V1)
}

func Call_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_19 gopurs_runtime.Value
go__go_1_0_19 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_19:
for {
if false { continue go__go_1_0_19 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_20 gopurs_runtime.Value
go__go_4_1_20 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_20:
for {
if false { continue go__go_4_1_20 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_20
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_20, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_19
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_19
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_19, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_manyRec(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
Plus1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_3_1
Alt0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_4_2
Functor0_5_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
pure_7_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_5
return gopurs_runtime.Func(func(p_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(Alt0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, pkg_Control_Monad_Rec_Class.Get_Loop(), p_8), gopurs_runtime.Apply(Applicative0_6_4.V1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})})), gopurs_runtime.Func(func(aa_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_5, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Monad_Rec_Class.Get_bifunctorStep(), "bimap"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_11, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](acc_9)})}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_12_6_21 gopurs_runtime.Value
go__go_12_6_21 = gopurs_runtime.Func(func(v_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_13_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_13_loop_val)
var v1_14_loop gopurs_runtime.Value = v1_14_loop_val
go__go_12_6_21:
for {
if false { continue go__go_12_6_21 }
var v_13 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_13_loop
_ = v_13
var v1_14 gopurs_runtime.Value = v1_14_loop
_ = v1_14
var __t7 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 1358893437 && v1_14.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_13)}
goto end_branch_7
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 1358893437 && v1_14.UnsafePtr != nil) {
v_13_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_14.UnsafePtr).V0, v_13})})
v1_14_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_14.UnsafePtr).V1)}
continue go__go_12_6_21
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t7))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_12_6_21, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](acc_9))})))}
}), aa_10))
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
})
})
}

func Call_someRec(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictAlternative_1_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictAlternative_1 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_1_loop
_ = dictAlternative_1
Apply0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_1.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_0
Functor0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_1.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_0.V1, gopurs_runtime.Apply2(Functor0_3_1.V0, pkg_Data_List_Types.Get_Cons(), v_4), gopurs_runtime.Apply2(Call_manyRec(dictMonadRec_0), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(dictAlternative_1)}, v_4))
})
}

func Call_some(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Types.Get_Cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4)
})))
})
})
}

func Call_many(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Alt0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Alt0_1_0.V1, gopurs_runtime.Apply2(Call_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4), gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}))
})
})
}

func Call_last(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
last:
for {
if false { continue last }
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
continue last
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}
}

func Call_insertBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
insertBy:
for {
if false { continue insertBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_insertBy(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))})}
goto end_branch_1
} else {

}
}
{
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2)
}
}

func Call_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
insertAt:
for {
if false { continue insertAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})}})}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))}))})))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t1)
}
}

func Call_insert(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_insertBy(), dictOrd_0.V1)
}

func Call_init(lst_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var lst_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_1, "init")))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_unsnoc(lst_0))}))}))
}

func Call_index(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
index:
for {
if false { continue index }
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (v1_1) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
v_0_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
v1_1_loop = Call_sub__1043827704(gopurs_runtime.Int(v1_1), gopurs_runtime.Int(1)).IntVal
continue index
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1)
}
}

func Call_head(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_transpose(v_0_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
transpose:
for {
if false { continue transpose }
var v_0 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_0_loop
_ = v_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
v_0_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})
continue transpose
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_mapMaybe(Get_head()), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}))})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_transpose(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1)}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(Call_mapMaybe(Get_tail()), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})))})})})))}))})})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t3)
}
}

func Call_groupBy(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
groupBy:
for {
if false { continue groupBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
v2_2_0 := Call_span(gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init")))}})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_groupBy(v_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest"))))}))})})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t1)
}
}

func Call_groupAllBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Bool(Call_eq__1272715810(uint32(gopurs_runtime.Apply2(p_0, x_1, y_2).IntVal), 902936544)).IntVal) != (0))
})
}))
_ = __local_var_1_0
__local_var_2_1 := Call_sortBy(p_0)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
}

func Call_group(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_groupBy(), dictEq_0.V0)
}

func Call_groupAll(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := Call_group(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(dictOrd_0.V0, gopurs_runtime.Value{})))
_ = __local_var_1_0
compare_2_2 := dictOrd_0.V1
_ = compare_2_2
__local_var_2_1 := gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_sortBy(compare_2_2), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_3))})))}
})
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
}

func Call_fromFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V2, pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_foldM(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(Applicative0_1_0.V1, v1_4)
goto end_branch_3
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {
__local_var_6_2 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1
_ = __local_var_6_2
__t3 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(v_3, v1_4, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_foldM(dictMonad_0), v_3, b_prime_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__local_var_6_2)})
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
})
})
})
}
}

func Call_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_22 gopurs_runtime.Value
go__go_1_0_22 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop int64 = v_2_loop_val.IntVal
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_22:
for {
if false { continue go__go_1_0_22 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(v_2)})}
goto end_branch_1
} else {

}
}
{
v_2_loop = Call_add__560788792(gopurs_runtime.Int(v_2), gopurs_runtime.Int(1)).IntVal
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_22
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t2))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_22, gopurs_runtime.Int(0))
}

func Call_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[int64] {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var go__go_2_0_23 gopurs_runtime.Value
go__go_2_0_23 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_3_loop_val)
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_0_23:
for {
if false { continue go__go_2_0_23 }
var v_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_3)}
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
v_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0, v_3})})
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)}
continue go__go_2_0_23
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__291265340(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(gopurs_runtime.Apply(Get_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Int(v_2.IntVal)).IntVal)
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Call_findIndex(fn_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_23, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))))})
}

func Call_filterM(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
filterM:
for {
if false { continue filterM }
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
__local_var_5_2 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_5_2
__local_var_6_3 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_6_3
__t5 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(v_3, __local_var_5_2), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_filterM(dictMonad_0), v_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__local_var_6_3)}), gopurs_runtime.Func(func(xs_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (b_7.IntVal) != (0) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_5_2, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_prime_8)})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_prime_8))}
}
end_branch_4:
return gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))})
}))
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
})
}
}

func Call_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_24 gopurs_runtime.Value
go__go_1_0_24 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_24:
for {
if false { continue go__go_1_0_24 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_25 gopurs_runtime.Value
go__go_4_1_25 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_25:
for {
if false { continue go__go_4_1_25 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_25
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_25, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_24
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_24
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_24, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_intersectBy(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_4, v1_5)
})
}))
_ = semigroupDisj1_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_4_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")), gopurs_runtime.Apply(v_0, x_3), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}).IntVal) != (0))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_intersect(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy(), dictEq_0.V0)
}

func Call_nubByEq(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
nubByEq:
for {
if false { continue nubByEq }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__local_var_2_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0
_ = __local_var_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_nubByEq(v_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_not__3201284355(gopurs_runtime.Bool((gopurs_runtime.Apply2(v_0, __local_var_2_0, y_3).IntVal) != (0))).IntVal) != (0))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}))))})})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}
}

func Call_nubEq(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_nubByEq(), dictEq_0.V0)
}

func Call_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
eqList_1_0 := &pkg_Data_Eq.Constructor_Eq[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_eq1List(), "eq1"), dictEq_0)}
_ = eqList_1_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eqList_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](y_3))}).IntVal) != (0))
})
}))
}

func Call_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordList_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_List_Types.Get_ordList(), dictOrd_0))
_ = ordList_1_0
eqPattern1_2_1 := Call_eqPattern(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqPattern1_2_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqPattern1_2_1
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(ordList_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](y_4))}).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_elemLastIndex(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1).IntVal) != (0))
}))
}

func Call_elemIndex(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_findIndex(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1).IntVal) != (0))
}))
}

func Call_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_26 gopurs_runtime.Value
go__go_1_0_26 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
go__go_1_0_26:
for {
if false { continue go__go_1_0_26 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1
continue go__go_1_0_26
__t1 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = v_2
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_1_0_26
}

func Call_dropEnd(n_0_loop int64, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(gopurs_runtime.Apply(Get_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal), gopurs_runtime.Int(n_0)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}))
}

func Call_drop(v_0_loop int64, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
drop:
for {
if false { continue drop }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
v_0_loop = Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal
v1_1_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1
continue drop
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}
}

func Call_slice(start_0_loop int64, end_1_loop int64, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(end_1), gopurs_runtime.Int(start_0)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_drop(start_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_2)))}))})))}
}

func Call_takeEnd(n_0_loop int64, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_drop(Call_sub__1043827704(gopurs_runtime.Int(gopurs_runtime.Apply(Get_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal), gopurs_runtime.Int(n_0)).IntVal, xs_1))})
}

func Call_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
deleteBy:
for {
if false { continue deleteBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_deleteBy(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))})})}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}
}

func Call_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_deleteBy(eq_0, a_4, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_3)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_nubByEq(eq_0, ys_2))}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
}

func Call_union(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy(), dictEq_0.V0)
}

func Call_deleteAt(v_0_loop int64, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
deleteAt:
for {
if false { continue deleteAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}})}
goto end_branch_1
} else {

}
}
{
__local_var_2_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0
_ = __local_var_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_3)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_deleteAt(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1))}))})))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t2)
}
}

func Call_delete(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_deleteBy(), dictEq_0.V0)
}

func Call_difference(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_deleteBy(dictEq_0.V0, a_2, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_1)))}
})
}))
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = a_1_loop
_ = a_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_1)}, b_0))
}

func Call_concat(v_0_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}, Get_identity()))
}

func Call_alterAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
alterAt:
for {
if false { continue alterAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (v_0) == (0) {
v3_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0))
_ = v3_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}})}
goto end_branch_3
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_alterAt(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))}))})))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_4:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t4)
}
}

func Call_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(Get_alterAt(), gopurs_runtime.Int(n_0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, x_2)})}
}))
}

func Call_alt__267341625(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_alt__1433662167(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_alt__1643437292(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__993904534(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3286817552(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__243192752(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__1030762512(dict_0_loop *pkg_Control_Apply.Constructor_Apply[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2007181164(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2169384906(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_lift2__2286084406(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__3007543670(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__2114580400(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_bind__3842572251(dict_0_loop *pkg_Control_Bind.Constructor_Bind[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__556822235(dict_0_loop *pkg_Control_Bind.Constructor_Bind[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1429695463(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__226726151(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2057121831(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1872090113(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_bind__2207507201(v_0_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]], v1_1_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_defer__3967925939(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_defer__449752361(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_tailRec__2110844386(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_28 gopurs_runtime.Value
go__go_1_0_28 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_28:
for {
if false { continue go__go_1_0_28 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_28
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
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
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_28, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRec__2666749533(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_29 gopurs_runtime.Value
go__go_1_0_29 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_29:
for {
if false { continue go__go_1_0_29 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_29
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
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
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_29, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRecM__1824976720(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__3615056745(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__3865988408(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__3515372237(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__1444729948(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM2__1943630176(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_tailRecM2__1136195496(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_Rec_Class.Get_monadRecMaybe(), "tailRecM"), gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(o_3, "a"), gopurs_runtime.RecordGet(o_3, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_1, b_2))
}

func Call_compose__706970832(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2527254334(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1604328382(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2995688990(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__4020612094(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2710321297(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1933206353(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bimap__132457202(dict_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bimap__2801350668(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)))}})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)))}})}
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

func Call_eq__789642299(dict_0_loop *pkg_Data_Eq.Constructor_Eq[bool]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[bool] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__1697837627(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2843686287(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.IntVal) == (__eta1_1.IntVal))
}

func Call_eq__2276491096(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) == ((__eta1_1.IntVal) != (0)))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__1425708671(dict_0_loop *pkg_Data_Eq.Constructor_Eq[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__1272715810(v_0_loop uint32, v1_1_loop uint32) bool {
var v_0 uint32 = v_0_loop
_ = v_0
var v1_1 uint32 = v1_1_loop
_ = v1_1
var __t2 bool
{
if (v_0 == 1527465420) {
var __t0 bool
{
if (v1_1 == 1527465420) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t2 = __t0
goto end_branch_2
} else {

}
}
{
if (v_0 == 380165415) {
var __t1 bool
{
if (v1_1 == 380165415) {
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
if ((v_0 == 902936544)) && ((v1_1 == 902936544)) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return __t2
}

func Call_notEq__2384498378(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return ((gopurs_runtime.Apply2(dictEq_0.V0, x_1, y_2).IntVal) != (0)) != (true)
}

func Call_notEq__1272715810(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), x_0, y_1).IntVal) != (0)) != (true)
}

func Call_any__4179648253(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_1.V1, v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), dictHeytingAlgebra_1.V2))
}

func Call_any__842931401(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")), __eta0_0, __eta1_1)
}

func Call_foldl__3041692656(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__66388714(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2602334544(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__165683952(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__1148906672(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__1754241693(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_30 gopurs_runtime.Value
go__go_1_0_30 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop int64 = b_2_loop_val.IntVal
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_30:
for {
if false { continue go__go_1_0_30 }
var b_2 int64 = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(b_2)
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(b_2), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_30
__t1 = gopurs_runtime.Int(gopurs_runtime.Value{}.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Int(__t1.IntVal)
}
}()
})
})
return go__go_1_0_30
}

func Call_foldl__3943124669(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_31 gopurs_runtime.Value
go__go_1_0_31 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_31:
for {
if false { continue go__go_1_0_31 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_31
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
return go__go_1_0_31
}

func Call_foldl__396932925(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_32 gopurs_runtime.Value
go__go_1_0_32 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_32:
for {
if false { continue go__go_1_0_32 }
var b_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_32
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return go__go_1_0_32
}

func Call_foldl__3915700701(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_33 gopurs_runtime.Value
go__go_1_0_33 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_33:
for {
if false { continue go__go_1_0_33 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_33
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
return go__go_1_0_33
}

func Call_foldr__2512763050(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__3673994608(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2829803163(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2979608669(f_0_loop gopurs_runtime.Value, b_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = b_1_loop
_ = b_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)})
_ = __local_var_2_0
var go__go_3_2_34 gopurs_runtime.Value
go__go_3_2_34 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_34:
for {
if false { continue go__go_3_2_34 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_2_34
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_34, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_foldr__3489910557(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_0
var go__go_3_2_35 gopurs_runtime.Value
go__go_3_2_35 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_35:
for {
if false { continue go__go_3_2_35 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_2_35
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_35, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_foldr__3234921885(f_0_loop gopurs_runtime.Value, b_1_loop *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = b_1_loop
_ = b_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_1)})
_ = __local_var_2_0
var go__go_3_2_36 gopurs_runtime.Value
go__go_3_2_36 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_36:
for {
if false { continue go__go_3_2_36 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)})})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_2_36
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](__t3))}
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_36, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
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

func Call_flip__3709724320(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3563101792(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2270521984(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2106496000(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2373571712(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3192966848(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2681340064(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__872296768(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__4091748192(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__3116241637(dict_0_loop *pkg_Data_Functor.Constructor_Functor[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2876470885(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3778170420(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__4242765044(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3058795348(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__438443400(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(pkg_Data_List_Types.Get_listMap(), __eta0_0, __eta1_1)
}

func Call_map__291265340(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[int64]) *pkg_Data_Maybe.Constructor_Just[int64] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[int64] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.Apply(v_0, gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0.IntVal)).IntVal)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0)
}

func Call_map__901270812(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_map__140514012(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0))})))}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t0)
}

func Call_map__2202537180(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)))}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]]](__t0)
}

func Call_map__1681779388(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)))}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t0)
}

func Call_map__3486165692(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_fromZipper__1019554324(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t6 = v1_1
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1304506903) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, v1_1, (*pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_5
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2884341868) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1})})
goto end_branch_5
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2195694037) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, v1_1, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1584522659) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3952671150) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3, v1_1})})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = __t5
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

func Call_insertAndLookupBy__3244745033(comp_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, orig_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var orig_2 gopurs_runtime.Value = orig_2_loop
_ = orig_2
var up_3_0_37 gopurs_runtime.Value
up_3_0_37 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
up_3_0_37:
for {
if false { continue up_3_0_37 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t7 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}
goto end_branch_7
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
var __t6 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1304506903) {
__t6 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_6
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2884341868) {
__t6 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})})
goto end_branch_6
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2195694037) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0_37
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1584522659) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0})}, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0_37
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3952671150) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2})}, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}})}
continue up_3_0_37
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
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
var down_4_8_38 gopurs_runtime.Value
down_4_8_38 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
down_4_8_38:
for {
if false { continue down_4_8_38 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t15 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 2764020654) {
__t15 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(up_3_0_37, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}, gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, k_1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}})}))
goto end_branch_15
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1177901036) {
v2_7_9 := gopurs_runtime.Apply2(comp_0, k_1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
_ = v2_7_9
var __t10 gopurs_runtime.Value
{
if (uint32(v2_7_9.IntVal) == 902936544) {
__t10 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_10
} else {

}
}
{
if (uint32(v2_7_9.IntVal) == 1527465420) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_38
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_38
__t10 = gopurs_runtime.Value{}
}
end_branch_10:
__t15 = __t10
goto end_branch_15
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1064476974) {
v2_7_11 := gopurs_runtime.Apply2(comp_0, k_1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
_ = v2_7_11
var __t14 gopurs_runtime.Value
{
if (uint32(v2_7_11.IntVal) == 902936544) {
__t14 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_14
} else {

}
}
{
v3_8_12 := gopurs_runtime.Apply2(comp_0, k_1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3)
_ = v3_8_12
var __t13 gopurs_runtime.Value
{
if (uint32(v3_8_12.IntVal) == 902936544) {
__t13 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_13
} else {

}
}
{
if (uint32(v2_7_11.IntVal) == 1527465420) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_38
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
if ((uint32(v2_7_11.IntVal) == 380165415)) && ((uint32(v3_8_12.IntVal) == 1527465420)) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_38
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4
continue down_4_8_38
__t13 = gopurs_runtime.Value{}
}
end_branch_13:
__t14 = __t13
}
end_branch_14:
__t15 = __t14
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
return gopurs_runtime.Apply2(down_4_8_38, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, orig_2)
}

func Call_listMap__4135416762(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_43 gopurs_runtime.Value
chunkedRevMap_1_0_43 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_43:
for {
if false { continue chunkedRevMap_1_0_43 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_43
__t19 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_44 gopurs_runtime.Value
reverseUnrolledMap_4_1_44 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_6_loop_val)
reverseUnrolledMap_4_1_44:
for {
if false { continue reverseUnrolledMap_4_1_44 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v3_6_loop
_ = v3_6
var __t8 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0), v3_6})})})})})})
continue reverseUnrolledMap_4_1_44
__t8 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = v3_6
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
var __t13 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_44, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_43, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_alterAt__3453373293(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (v_0) == (0) {
v3_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0))
_ = v3_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}})}
goto end_branch_3
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_alterAt(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))}))})))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_4:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t4)
}

func Call_deleteAt__2845095501(v_0_loop int64, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}})}
goto end_branch_1
} else {

}
}
{
__local_var_2_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0
_ = __local_var_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_3)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_deleteAt(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1))}))})))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t2)
}

func Call_deleteBy__697302515(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_deleteBy(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))})})}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_drop__1836090668(v_0_loop int64, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_drop(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1))}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_filter__1617261107(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_45 gopurs_runtime.Value
go__go_1_0_45 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_45:
for {
if false { continue go__go_1_0_45 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_46 gopurs_runtime.Value
go__go_4_1_46 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_46:
for {
if false { continue go__go_4_1_46 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_46
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_46, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_45
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_45
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_45, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_filterM__14771079(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
__local_var_5_2 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_5_2
__local_var_6_3 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_6_3
__t5 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(v_3, __local_var_5_2), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_filterM(dictMonad_0), v_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__local_var_6_3)}), gopurs_runtime.Func(func(xs_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (b_7.IntVal) != (0) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_5_2, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_prime_8)})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_prime_8))}
}
end_branch_4:
return gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))})
}))
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
})
}

func Call_findIndex__2366045378(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_47 gopurs_runtime.Value
go__go_1_0_47 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop int64 = v_2_loop_val.IntVal
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_47:
for {
if false { continue go__go_1_0_47 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(v_2)})}
goto end_branch_1
} else {

}
}
{
v_2_loop = Call_add__560788792(gopurs_runtime.Int(v_2), gopurs_runtime.Int(1)).IntVal
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_47
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t2))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_47, gopurs_runtime.Int(0))
}

func Call_findLastIndex__2366045378(fn_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[int64] {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var go__go_2_1_48 gopurs_runtime.Value
go__go_2_1_48 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_3_loop_val)
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_1_48:
for {
if false { continue go__go_2_1_48 }
var v_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t2 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_3)}
goto end_branch_2
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
v_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0, v_3})})
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)}
continue go__go_2_1_48
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__local_var_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Call_findIndex(fn_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_1_48, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
_ = __local_var_2_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_2_0)}.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(gopurs_runtime.Apply(Get_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_2_0)}.UnsafePtr).V0.IntVal)).IntVal)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t3)
}

func Call_foldM__3577257629(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(Applicative0_1_0.V1, v1_4)
goto end_branch_3
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {
__local_var_6_2 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1
_ = __local_var_6_2
__t3 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(v_3, v1_4, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_foldM(dictMonad_0), v_3, b_prime_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__local_var_6_2)})
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
})
})
})
}

func Call_groupBy__1039549870(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
v2_2_0 := Call_span(gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init")))}})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_groupBy(v_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest"))))}))})})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t1)
}

func Call_head__3729839663(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_index__304299960(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (v1_1) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1, Call_sub__1043827704(gopurs_runtime.Int(v1_1), gopurs_runtime.Int(1)).IntVal))}))}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1)
}

func Call_insertAt__2634211748(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})}})}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))}))})))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t1)
}

func Call_insertBy__1738998985(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_insertBy(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))})}
goto end_branch_1
} else {

}
}
{
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2)
}

func Call_intersectBy__1190504998(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_4, v1_5)
})
}))
_ = semigroupDisj1_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_4_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")), gopurs_runtime.Apply(v_0, x_3), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}).IntVal) != (0))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_last__4043133652(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_5
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V0})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_last((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V1))}))}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_3:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_6:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t6)
}

func Call_many__542682753(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Alt0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Alt0_1_0.V1, gopurs_runtime.Apply2(Call_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4), gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}))
})
})
}

func Call_manyRec__4046352885(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
Plus1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_3_1
Alt0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_4_2
Functor0_5_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
pure_7_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_5
return gopurs_runtime.Func(func(p_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(Alt0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, pkg_Control_Monad_Rec_Class.Get_Loop(), p_8), gopurs_runtime.Apply(Applicative0_6_4.V1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})})), gopurs_runtime.Func(func(aa_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_5, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Monad_Rec_Class.Get_bifunctorStep(), "bimap"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_11, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](acc_9)})}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_12_6_49 gopurs_runtime.Value
go__go_12_6_49 = gopurs_runtime.Func(func(v_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_13_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_13_loop_val)
var v1_14_loop gopurs_runtime.Value = v1_14_loop_val
go__go_12_6_49:
for {
if false { continue go__go_12_6_49 }
var v_13 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_13_loop
_ = v_13
var v1_14 gopurs_runtime.Value = v1_14_loop
_ = v1_14
var __t7 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 1358893437 && v1_14.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_13)}
goto end_branch_7
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 1358893437 && v1_14.UnsafePtr != nil) {
v_13_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_14.UnsafePtr).V0, v_13})})
v1_14_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_14.UnsafePtr).V1)}
continue go__go_12_6_49
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t7))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_12_6_49, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](acc_9))})))}
}), aa_10))
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
})
})
}

func Call_mapMaybe__1486753757(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_50 gopurs_runtime.Value
go__go_1_0_50 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_50:
for {
if false { continue go__go_1_0_50 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_51 gopurs_runtime.Value
go__go_4_1_51 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_51:
for {
if false { continue go__go_4_1_51 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_51
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_51, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_50
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_50
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_50, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_mapMaybe__748617661(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_52 gopurs_runtime.Value
go__go_1_0_52 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_52:
for {
if false { continue go__go_1_0_52 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_53 gopurs_runtime.Value
go__go_4_1_53 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_53:
for {
if false { continue go__go_4_1_53 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_53
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_53, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_52
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_52
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_52, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_mapMaybe__4251473821(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_54 gopurs_runtime.Value
go__go_1_0_54 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_54:
for {
if false { continue go__go_1_0_54 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_55 gopurs_runtime.Value
go__go_4_1_55 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_55:
for {
if false { continue go__go_4_1_55 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_55
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_55, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_54
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_54
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_54, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_mapMaybe__2491277821(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_56 gopurs_runtime.Value
go__go_1_0_56 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_56:
for {
if false { continue go__go_1_0_56 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_57 gopurs_runtime.Value
go__go_4_1_57 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_57:
for {
if false { continue go__go_4_1_57 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_57
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_57, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_56
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_56
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_56, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_nubBy__1502591776(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_58 gopurs_runtime.Value
go__go_1_0_58 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
var v2_4_loop gopurs_runtime.Value = v2_4_loop_val
go__go_1_0_58:
for {
if false { continue go__go_1_0_58 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var v2_4 gopurs_runtime.Value = v2_4_loop
_ = v2_4
var __t3 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr != nil) {
v3_5_1 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V0, v_2)
_ = v3_5_1
var __t2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.RecordGet(v3_5_1, "found").IntVal) != (0) {
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = v1_3
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V1)}
continue go__go_1_0_58
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V0, v1_3})})
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V1)}
continue go__go_1_0_58
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
})
__local_var_2_4 := gopurs_runtime.Apply2(go__go_1_0_58, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_59 gopurs_runtime.Value
go__go_4_5_59 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_59:
for {
if false { continue go__go_4_5_59 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t6 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_5_59
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t6))}
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_5_59, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Apply(__local_var_2_4, x_3))
})
}

func Call_nubByEq__3655321914(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__local_var_2_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0
_ = __local_var_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_nubByEq(v_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_0, __local_var_2_0, y_3).IntVal) != (0)) != (true))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}))))})})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_singleton__2450819477(a_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})
}

func Call_singleton__707062261(a_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})
}

func Call_singleton__3932757557(a_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})
}

func Call_some__542682753(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Types.Get_Cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4)
})))
})
})
}

func Call_sortBy__1502591776(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0_62 gopurs_runtime.Value
_ = merge_1_0_62
merge_1_0_62 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
__local_var_4_1 := uint32(gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal)
_ = __local_var_4_1
var __t2 bool
{
if (__local_var_4_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 == 380165415) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_62, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
goto end_branch_3
} else {

}
}
{
__t3 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_62, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3))}))}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3))}
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
})
})
var mergePairs_2_6_63 gopurs_runtime.Value
_ = mergePairs_2_6_63
mergePairs_2_6_63 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_and_8 bool = false
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {

var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
__t_and_8 = (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 1358893437 && __t_tag_7.UnsafePtr != nil)
}
if __t_and_8 {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_62, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}.UnsafePtr).V0))})))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(mergePairs_2_6_63, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}.UnsafePtr).V1)})))})})}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_3))}
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t9))}
})
var mergeAll_3_10_64 gopurs_runtime.Value
mergeAll_3_10_64 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_4_loop_val)
mergeAll_3_10_64:
for {
if false { continue mergeAll_3_10_64 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_4_loop
_ = v_4
var __t13 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_12 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr != nil) {

var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr).V1)}
__t_and_12 = (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 1358893437 && __t_tag_11.UnsafePtr == nil)
}
if __t_and_12 {
__t13 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(mergePairs_2_6_63, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
continue mergeAll_3_10_64
__t13 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t13)}
}
}()
})
var sequences_4_14_65 gopurs_runtime.Value
_ = sequences_4_14_65
var descending_4_15_66 gopurs_runtime.Value
_ = descending_4_15_66
var ascending_4_16_67 gopurs_runtime.Value
_ = ascending_4_16_67
sequences_4_14_65 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
var __t_and_18 bool = false
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {

var __t_tag_17 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}
__t_and_18 = (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 1358893437 && __t_tag_17.UnsafePtr != nil)
}
if __t_and_18 {
var __t22 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]
{
__local_var_6_20 := uint32(gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0).IntVal)
_ = __local_var_6_20
var __t21 bool
{
if (__local_var_6_20 == 1527465420) {
__t21 = false
goto end_branch_21
} else {

}
}
{
if (__local_var_6_20 == 380165415) {
__t21 = true
goto end_branch_21
} else {

}
}
{
__t21 = false
}
end_branch_21:
if __t21 {
__t22 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(descending_4_15_66, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V1)}))
goto end_branch_22
} else {

}
}
{
__local_var_6_19 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_19
__t22 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(ascending_4_16_67, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_19, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_7)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V1)}))
}
end_branch_22:
__t23 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)}
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t23))}
})
descending_4_15_66 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 gopurs_runtime.Value
{
var __t_and_26 bool = false
if (v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil) {

__local_var_8_24 := uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0).IntVal)
_ = __local_var_8_24
var __t25 bool
{
if (__local_var_8_24 == 1527465420) {
__t25 = false
goto end_branch_25
} else {

}
}
{
if (__local_var_8_24 == 380165415) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
__t_and_26 = __t25
}
if __t_and_26 {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(descending_4_15_66, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_6)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))}
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_6)})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(sequences_4_14_65, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_7))})))})})}
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t27))}
})
})
})
ascending_4_16_67 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0)) != (true)) {
__t28 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(ascending_4_16_67, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_8)})})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))}
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(sequences_4_14_65, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_7))})))})})}
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t28))}
})
})
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_10_64, gopurs_runtime.Apply(sequences_4_14_65, x_5))
})
}

func Call_span__799093643(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0).IntVal) != (0)) {
v2_2_0 := Call_span(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)
_ = v2_2_0
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}

func Call_tail__1771843450(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t0)
}

func Call_tails__3932757557(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_tails((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))}))})})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t0)
}

func Call_transpose__682228544(v_0_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_0_loop
_ = v_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_transpose(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})))}))}
goto end_branch_2
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_mapMaybe(Get_head()), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}))})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_transpose(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1)}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(Call_mapMaybe(Get_tail()), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})))})})})))}))})})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t3)
}

func Call_uncons__3009258782(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_unionBy__1190504998(eq_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_deleteBy(eq_0, a_4, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_3)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_nubByEq(eq_0, ys_2))}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
}

func Call_unsnoc__2942606998(lst_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var lst_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
var go__go_1_0_71 gopurs_runtime.Value
go__go_1_0_71 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
go__go_1_0_71:
for {
if false { continue go__go_1_0_71 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("last", "revInit", (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)})})}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3})})
continue go__go_1_0_71
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(h_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_72 gopurs_runtime.Value
go__go_3_4_72 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_4_72:
for {
if false { continue go__go_3_4_72 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t5 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_5
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_4_72
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_3_4_72, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(h_2, "revInit")))})))}, gopurs_runtime.RecordGet(h_2, "last"))
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_1_0_71, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}))
}

func Call_updateAt__2634211748(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1))}))})))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t2)
}

func Call_zipWith__884793877(f_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var go__go_3_0_73 gopurs_runtime.Value
go__go_3_0_73 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_6_loop_val)
go__go_3_0_73:
for {
if false { continue go__go_3_0_73 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_6_loop
_ = v2_6
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0), v2_6})})
continue go__go_3_0_73
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
})
var go__go_4_2_74 gopurs_runtime.Value
go__go_4_2_74 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_74:
for {
if false { continue go__go_4_2_74 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_2_74
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_2_74, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(go__go_3_0_73, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}))
}

func Call_zipWith__4203240021(f_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var go__go_3_0_75 gopurs_runtime.Value
go__go_3_0_75 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_6_loop_val)
go__go_3_0_75:
for {
if false { continue go__go_3_0_75 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_6_loop
_ = v2_6
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0), v2_6})})
continue go__go_3_0_75
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
})
var go__go_4_2_76 gopurs_runtime.Value
go__go_4_2_76 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_76:
for {
if false { continue go__go_4_2_76 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_2_76
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_2_76, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(go__go_3_0_75, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}))
}

func Call_zipWith__3856182069(f_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var go__go_3_0_77 gopurs_runtime.Value
go__go_3_0_77 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_6_loop_val)
go__go_3_0_77:
for {
if false { continue go__go_3_0_77 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_6_loop
_ = v2_6
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0), v2_6})})
continue go__go_3_0_77
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
})
var go__go_4_2_78 gopurs_runtime.Value
go__go_4_2_78 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_78:
for {
if false { continue go__go_4_2_78 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_2_78
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_2_78, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(go__go_3_0_77, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}))
}

func Call_alaF__2824451980(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_3_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_alaF__4085337484(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_3_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__669572705(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__3029065925(dict_0_loop *pkg_Data_Ord.Constructor_Ord[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_negate__2635823316(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Int(-(__eta0_0.IntVal))
}

func Call_negate__1364373265(dictRing_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_sub__1124926121(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
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

func Call_append__2832914972(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
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

func Call_append__2013893496(xs_0_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]], ys_1_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var xs_0 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = xs_0_loop
_ = xs_0
var ys_1 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = ys_1_loop
_ = ys_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
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

func Call_zero__1556010056(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2130238610(dict_0_loop *pkg_Data_Show.Constructor_Show[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_sequence__1886310617(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_sequence__2616145509(dictApplicative_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, pkg_Data_List_Types.Get_identity())
}

func Call_traverse__3956862083(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__2839486329(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_6_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_6_5
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_6_5.V0, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__3827943605(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


