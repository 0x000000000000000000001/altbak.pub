package Data_List_Lazy

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
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
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

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith
}

var cache_zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		cache_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, xs_2_box, ys_3_box)
})
	})
	return cache_zipWithA
}

var cache_zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		cache_zip = gopurs_runtime.Apply(Get_zipWith(), pkg_Data_Tuple.Get_Tuple())
	})
	return cache_zip
}

var cache_updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		cache_updateAt = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAt(n_0_box.IntVal, x_1_box, xs_2_box)
})
	})
	return cache_updateAt
}

var cache_unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		cache_unzip = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
_ = __local_var_1_0
__local_var_2_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_4_2
__local_var_5_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_1_0, __local_var_4_2})}
})), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_1, __local_var_5_3})}
}))})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_List_Lazy_Types.Get_nil(), pkg_Data_List_Lazy_Types.Get_nil()})})
	})
	return cache_unzip
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0_box))}
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

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take(n_0_box.IntVal)
})
	})
	return cache_take
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_tail(xs_0_box))}
})
	})
	return cache_tail
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_stripPrefix(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), v_1_box, s_2_box))}
})
	})
	return cache_stripPrefix
}

var cache_span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		cache_span = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span(p_0_box, xs_1_box)
})
	})
	return cache_span
}

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snoc(xs_0_box, x_1_box)
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(a_0_box)
})
	})
	return cache_singleton
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

var cache_scanlLazy gopurs_runtime.Value
var once_scanlLazy sync.Once
func Get_scanlLazy() gopurs_runtime.Value {
	once_scanlLazy.Do(func() {
		cache_scanlLazy = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanlLazy(f_0_box, acc_1_box, xs_2_box)
})
	})
	return cache_scanlLazy
}

var cache_reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		cache_reverse = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reverse(xs_0_box)
})
	})
	return cache_reverse
}

var cache_replicateM gopurs_runtime.Value
var once_replicateM sync.Once
func Get_replicateM() gopurs_runtime.Value {
	once_replicateM.Do(func() {
		cache_replicateM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicateM(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_replicateM
}

var cache_repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		cache_repeat = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat(x_0_box)
})
	})
	return cache_repeat
}

var cache_replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		cache_replicate = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate(i_0_box.IntVal, xs_1_box)
})
	})
	return cache_replicate
}

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__range(start_0_box.IntVal, end_1_box.IntVal)
})
	})
	return cache_go__range
}

var cache_partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		cache_partition = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_partition(f_0_box)
})
	})
	return cache_partition
}

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null(x_0_box))
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
		cache_length = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), l_0, gopurs_runtime.Int(1)).IntVal)
})
}), gopurs_runtime.Int(0))
	})
	return cache_length
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = func() gopurs_runtime.Value {
var go__go_0_0_5 gopurs_runtime.Value
go__go_0_0_5 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
go__go_0_0_5:
for {
if false { continue go__go_0_0_5 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t4 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
__local_var_2_1 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))}
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
if (__t2.IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0})}
goto end_branch_3
} else {

}
}
{
v_1_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
continue go__go_0_0_5
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
}
}()
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return cache_last
}

var cache_iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		cache_iterate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate(f_0_box, x_1_box)
})
	})
	return cache_iterate
}

var cache_insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		cache_insertAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAt(v_0_box.IntVal, v1_1_box, v2_2_box)
})
	})
	return cache_insertAt
}

var cache_init gopurs_runtime.Value
var once_init sync.Once
func Get_init() gopurs_runtime.Value {
	once_init.Do(func() {
		cache_init = func() gopurs_runtime.Value {
var go__go_0_0_7 gopurs_runtime.Value
_ = go__go_0_0_7
go__go_0_0_7 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
__local_var_2_1 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))}
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
if (__t2.IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, pkg_Data_List_Lazy_Types.Get_nil()})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_cons(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0), gopurs_runtime.Apply(go__go_0_0_7, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_7, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return cache_init
}

var cache_index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		cache_index = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_index(xs_0_box)
})
	})
	return cache_index
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_head(xs_0_box))}
})
	})
	return cache_head
}

var cache_transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		cache_transpose = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_transpose(xs_0_box)
})
	})
	return cache_transpose
}

var cache_groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		cache_groupBy = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupBy(eq_0_box)
})
	})
	return cache_groupBy
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

var cache_insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		cache_insertBy = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertBy(cmp_0_box, x_1_box, xs_2_box)
})
	})
	return cache_insertBy
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

var cache_foldrLazy gopurs_runtime.Value
var once_foldrLazy sync.Once
func Get_foldrLazy() gopurs_runtime.Value {
	once_foldrLazy.Do(func() {
		cache_foldrLazy = gopurs_runtime.Func3(func(dictLazy_0_box gopurs_runtime.Value, op_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrLazy(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_0_box), op_1_box, z_2_box)
})
	})
	return cache_foldrLazy
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex(fn_0_box, xs_1_box))}
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
		cache_intersectBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectBy(eq_0_box, xs_1_box, ys_2_box)
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
		cache_nubByEq = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubByEq(eq_0_box)
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

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop(n_0_box.IntVal)
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

var cache_deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		cache_deleteBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy(eq_0_box, x_1_box, xs_2_box)
})
	})
	return cache_deleteBy
}

var cache_unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		cache_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionBy(eq_0_box, xs_1_box, ys_2_box)
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
		cache_deleteAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteAt(n_0_box.IntVal, xs_1_box)
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

var cache_cycle gopurs_runtime.Value
var once_cycle sync.Once
func Get_cycle() gopurs_runtime.Value {
	once_cycle.Do(func() {
		cache_cycle = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cycle(xs_0_box)
})
	})
	return cache_cycle
}

var cache_concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		cache_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concatMap(b_0_box, a_1_box)
})
	})
	return cache_concatMap
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concat(v_0_box)
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
		cache_alterAt = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alterAt(n_0_box.IntVal, f_1_box, xs_2_box)
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

var cache_alt__4218059372 gopurs_runtime.Value
var once_alt__4218059372 sync.Once
func Get_alt__4218059372() gopurs_runtime.Value {
	once_alt__4218059372.Do(func() {
		cache_alt__4218059372 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__4218059372(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__4218059372
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

var cache_pure__577281046 gopurs_runtime.Value
var once_pure__577281046 sync.Once
func Get_pure__577281046() gopurs_runtime.Value {
	once_pure__577281046.Do(func() {
		cache_pure__577281046 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure")
	})
	return cache_pure__577281046
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

var cache_apply__3620326986 gopurs_runtime.Value
var once_apply__3620326986 sync.Once
func Get_apply__3620326986() gopurs_runtime.Value {
	once_apply__3620326986.Do(func() {
		cache_apply__3620326986 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_applyList(), "apply")
	})
	return cache_apply__3620326986
}

var cache_bind__2073074151 gopurs_runtime.Value
var once_bind__2073074151 sync.Once
func Get_bind__2073074151() gopurs_runtime.Value {
	once_bind__2073074151.Do(func() {
		cache_bind__2073074151 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2073074151(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2073074151
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

var cache_bind__777862183 gopurs_runtime.Value
var once_bind__777862183 sync.Once
func Get_bind__777862183() gopurs_runtime.Value {
	once_bind__777862183.Do(func() {
		cache_bind__777862183 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__777862183(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__777862183
}

var cache_bind__523053991 gopurs_runtime.Value
var once_bind__523053991 sync.Once
func Get_bind__523053991() gopurs_runtime.Value {
	once_bind__523053991.Do(func() {
		cache_bind__523053991 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__523053991(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__523053991
}

var cache_bind__4082241 gopurs_runtime.Value
var once_bind__4082241 sync.Once
func Get_bind__4082241() gopurs_runtime.Value {
	once_bind__4082241.Do(func() {
		cache_bind__4082241 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindList(), "bind")
	})
	return cache_bind__4082241
}

var cache_bind__2775489217 gopurs_runtime.Value
var once_bind__2775489217 sync.Once
func Get_bind__2775489217() gopurs_runtime.Value {
	once_bind__2775489217.Do(func() {
		cache_bind__2775489217 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindList(), "bind")
	})
	return cache_bind__2775489217
}

var cache_bind__3799579873 gopurs_runtime.Value
var once_bind__3799579873 sync.Once
func Get_bind__3799579873() gopurs_runtime.Value {
	once_bind__3799579873.Do(func() {
		cache_bind__3799579873 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind")
	})
	return cache_bind__3799579873
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
		cache_defer__2590380358 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer")
	})
	return cache_defer__2590380358
}

var cache_defer__774977193 gopurs_runtime.Value
var once_defer__774977193 sync.Once
func Get_defer__774977193() gopurs_runtime.Value {
	once_defer__774977193.Do(func() {
		cache_defer__774977193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__774977193(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_defer__774977193
}

var cache_fix__1475205859 gopurs_runtime.Value
var once_fix__1475205859 sync.Once
func Get_fix__1475205859() gopurs_runtime.Value {
	once_fix__1475205859.Do(func() {
		cache_fix__1475205859 = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fix__1475205859(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_0_box), f_1_box)
})
	})
	return cache_fix__1475205859
}

var cache_fix__3570066147 gopurs_runtime.Value
var once_fix__3570066147 sync.Once
func Get_fix__3570066147() gopurs_runtime.Value {
	once_fix__3570066147.Do(func() {
		cache_fix__3570066147 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fix__3570066147(f_0_box)
})
	})
	return cache_fix__3570066147
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
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
var go__go_3_5_17 gopurs_runtime.Value
go__go_3_5_17 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_5_17:
for {
if false { continue go__go_3_5_17 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.UncurriedApp(__local_var_2_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_5_17
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_5_17, gopurs_runtime.UncurriedApp(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))))}
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

var cache_tailRecM2__3864450856 gopurs_runtime.Value
var once_tailRecM2__3864450856 sync.Once
func Get_tailRecM2__3864450856() gopurs_runtime.Value {
	once_tailRecM2__3864450856.Do(func() {
		cache_tailRecM2__3864450856 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2__3864450856(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_tailRecM2__3864450856
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

var cache_eq__2484408063 gopurs_runtime.Value
var once_eq__2484408063 sync.Once
func Get_eq__2484408063() gopurs_runtime.Value {
	once_eq__2484408063.Do(func() {
		cache_eq__2484408063 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2484408063(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2484408063
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

var cache_any__1385259145 gopurs_runtime.Value
var once_any__1385259145 sync.Once
func Get_any__1385259145() gopurs_runtime.Value {
	once_any__1385259145.Do(func() {
		cache_any__1385259145 = func() gopurs_runtime.Value {
semigroupDisj1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_0, v1_1)
})
}))
_ = semigroupDisj1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_0_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")))
}()
	})
	return cache_any__1385259145
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
		cache_foldl__524683195 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__524683195
}

var cache_foldl__3379885725 gopurs_runtime.Value
var once_foldl__3379885725 sync.Once
func Get_foldl__3379885725() gopurs_runtime.Value {
	once_foldl__3379885725.Do(func() {
		cache_foldl__3379885725 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__3379885725
}

var cache_foldl__1985071933 gopurs_runtime.Value
var once_foldl__1985071933 sync.Once
func Get_foldl__1985071933() gopurs_runtime.Value {
	once_foldl__1985071933.Do(func() {
		cache_foldl__1985071933 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__1985071933
}

var cache_foldl__536153533 gopurs_runtime.Value
var once_foldl__536153533 sync.Once
func Get_foldl__536153533() gopurs_runtime.Value {
	once_foldl__536153533.Do(func() {
		cache_foldl__536153533 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__536153533
}

var cache_foldl__176907901 gopurs_runtime.Value
var once_foldl__176907901 sync.Once
func Get_foldl__176907901() gopurs_runtime.Value {
	once_foldl__176907901.Do(func() {
		cache_foldl__176907901 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__176907901
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

var cache_foldr__2403185435 gopurs_runtime.Value
var once_foldr__2403185435 sync.Once
func Get_foldr__2403185435() gopurs_runtime.Value {
	once_foldr__2403185435.Do(func() {
		cache_foldr__2403185435 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2403185435(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2403185435
}

var cache_foldr__3192890333 gopurs_runtime.Value
var once_foldr__3192890333 sync.Once
func Get_foldr__3192890333() gopurs_runtime.Value {
	once_foldr__3192890333.Do(func() {
		cache_foldr__3192890333 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr")
	})
	return cache_foldr__3192890333
}

var cache_foldr__2389967549 gopurs_runtime.Value
var once_foldr__2389967549 sync.Once
func Get_foldr__2389967549() gopurs_runtime.Value {
	once_foldr__2389967549.Do(func() {
		cache_foldr__2389967549 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr")
	})
	return cache_foldr__2389967549
}

var cache_foldr__2492628765 gopurs_runtime.Value
var once_foldr__2492628765 sync.Once
func Get_foldr__2492628765() gopurs_runtime.Value {
	once_foldr__2492628765.Do(func() {
		cache_foldr__2492628765 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr")
	})
	return cache_foldr__2492628765
}

var cache_foldr__3433277981 gopurs_runtime.Value
var once_foldr__3433277981 sync.Once
func Get_foldr__3433277981() gopurs_runtime.Value {
	once_foldr__3433277981.Do(func() {
		cache_foldr__3433277981 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr")
	})
	return cache_foldr__3433277981
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_const__1245418657 gopurs_runtime.Value
var once_const__1245418657 sync.Once
func Get_const__1245418657() gopurs_runtime.Value {
	once_const__1245418657.Do(func() {
		cache_const__1245418657 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1245418657(a_0_box, v_1_box)
})
	})
	return cache_const__1245418657
}

var cache_const__754826900 gopurs_runtime.Value
var once_const__754826900 sync.Once
func Get_const__754826900() gopurs_runtime.Value {
	once_const__754826900.Do(func() {
		cache_const__754826900 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__754826900(a_0_box, v_1_box)
})
	})
	return cache_const__754826900
}

var cache_const__4115900574 gopurs_runtime.Value
var once_const__4115900574 sync.Once
func Get_const__4115900574() gopurs_runtime.Value {
	once_const__4115900574.Do(func() {
		cache_const__4115900574 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4115900574(a_0_box, v_1_box)
})
	})
	return cache_const__4115900574
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

var cache_flip__2422603136 gopurs_runtime.Value
var once_flip__2422603136 sync.Once
func Get_flip__2422603136() gopurs_runtime.Value {
	once_flip__2422603136.Do(func() {
		cache_flip__2422603136 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2422603136(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2422603136
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

var cache_map__1510739772 gopurs_runtime.Value
var once_map__1510739772 sync.Once
func Get_map__1510739772() gopurs_runtime.Value {
	once_map__1510739772.Do(func() {
		cache_map__1510739772 = gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map")
	})
	return cache_map__1510739772
}

var cache_map__109003388 gopurs_runtime.Value
var once_map__109003388 sync.Once
func Get_map__109003388() gopurs_runtime.Value {
	once_map__109003388.Do(func() {
		cache_map__109003388 = gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map")
	})
	return cache_map__109003388
}

var cache_map__829570556 gopurs_runtime.Value
var once_map__829570556 sync.Once
func Get_map__829570556() gopurs_runtime.Value {
	once_map__829570556.Do(func() {
		cache_map__829570556 = gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map")
	})
	return cache_map__829570556
}

var cache_map__2156385148 gopurs_runtime.Value
var once_map__2156385148 sync.Once
func Get_map__2156385148() gopurs_runtime.Value {
	once_map__2156385148.Do(func() {
		cache_map__2156385148 = gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map")
	})
	return cache_map__2156385148
}

var cache_map__291265340 gopurs_runtime.Value
var once_map__291265340 sync.Once
func Get_map__291265340() gopurs_runtime.Value {
	once_map__291265340.Do(func() {
		cache_map__291265340 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__291265340
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__901270812
}

var cache_map__1887399228 gopurs_runtime.Value
var once_map__1887399228 sync.Once
func Get_map__1887399228() gopurs_runtime.Value {
	once_map__1887399228.Do(func() {
		cache_map__1887399228 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__1887399228
}

var cache_map__1808515292 gopurs_runtime.Value
var once_map__1808515292 sync.Once
func Get_map__1808515292() gopurs_runtime.Value {
	once_map__1808515292.Do(func() {
		cache_map__1808515292 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__1808515292
}

var cache_map__539092636 gopurs_runtime.Value
var once_map__539092636 sync.Once
func Get_map__539092636() gopurs_runtime.Value {
	once_map__539092636.Do(func() {
		cache_map__539092636 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__539092636
}

var cache_map__1980149980 gopurs_runtime.Value
var once_map__1980149980 sync.Once
func Get_map__1980149980() gopurs_runtime.Value {
	once_map__1980149980.Do(func() {
		cache_map__1980149980 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__1980149980
}

var cache_map__316107900 gopurs_runtime.Value
var once_map__316107900 sync.Once
func Get_map__316107900() gopurs_runtime.Value {
	once_map__316107900.Do(func() {
		cache_map__316107900 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__316107900
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
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
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
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

var cache_applyLazy__879424557 gopurs_runtime.Value
var once_applyLazy__879424557 sync.Once
func Get_applyLazy__879424557() gopurs_runtime.Value {
	once_applyLazy__879424557.Do(func() {
		cache_applyLazy__879424557 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Lazy.Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Lazy.Get_force(), f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
}))
})
}))
	})
	return cache_applyLazy__879424557
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

var cache_functorLazy__3988504945 gopurs_runtime.Value
var once_functorLazy__3988504945 sync.Once
func Get_functorLazy__3988504945() gopurs_runtime.Value {
	once_functorLazy__3988504945.Do(func() {
		cache_functorLazy__3988504945 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy__3988504945
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

var cache_applyList__2495961956 gopurs_runtime.Value
var once_applyList__2495961956 sync.Once
func Get_applyList__2495961956() gopurs_runtime.Value {
	once_applyList__2495961956.Do(func() {
		cache_applyList__2495961956 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
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
	return cache_applyList__2495961956
}

var cache_applyList__1358886895 gopurs_runtime.Value
var once_applyList__1358886895 sync.Once
func Get_applyList__1358886895() gopurs_runtime.Value {
	once_applyList__1358886895.Do(func() {
		cache_applyList__1358886895 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
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

var cache_bindList__469219920 gopurs_runtime.Value
var once_bindList__469219920 sync.Once
func Get_bindList__469219920() gopurs_runtime.Value {
	once_bindList__469219920.Do(func() {
		cache_bindList__469219920 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindList(), "bind"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_bindList__469219920
}

var cache_bindList__1050117088 gopurs_runtime.Value
var once_bindList__1050117088 sync.Once
func Get_bindList__1050117088() gopurs_runtime.Value {
	once_bindList__1050117088.Do(func() {
		cache_bindList__1050117088 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_applyList()
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindList(), "bind"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, f_1)))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_bindList__1050117088
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

var cache_cons__891310957 gopurs_runtime.Value
var once_cons__891310957 sync.Once
func Get_cons__891310957() gopurs_runtime.Value {
	once_cons__891310957.Do(func() {
		cache_cons__891310957 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__891310957(x_0_box, xs_1_box)
})
	})
	return cache_cons__891310957
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_22 gopurs_runtime.Value
go__go_1_2_22 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_22:
for {
if false { continue go__go_1_2_22 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_22
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
return go__go_1_2_22
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_23 gopurs_runtime.Value
go__go_1_2_23 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_23:
for {
if false { continue go__go_1_2_23 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_23
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
return go__go_1_2_23
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__331628915
}

var cache_foldableList__21955931 gopurs_runtime.Value
var once_foldableList__21955931 sync.Once
func Get_foldableList__21955931() gopurs_runtime.Value {
	once_foldableList__21955931.Do(func() {
		cache_foldableList__21955931 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_24 gopurs_runtime.Value
go__go_1_2_24 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_24:
for {
if false { continue go__go_1_2_24 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_24
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
return go__go_1_2_24
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__21955931
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_25 gopurs_runtime.Value
go__go_1_2_25 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_25:
for {
if false { continue go__go_1_2_25 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_25
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
return go__go_1_2_25
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__1218280485
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_functorList__3996674161
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

var cache_nil__1478684294 gopurs_runtime.Value
var once_nil__1478684294 sync.Once
func Get_nil__1478684294() gopurs_runtime.Value {
	once_nil__1478684294.Do(func() {
		cache_nil__1478684294 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__1478684294
}

var cache_nil__2012296605 gopurs_runtime.Value
var once_nil__2012296605 sync.Once
func Get_nil__2012296605() gopurs_runtime.Value {
	once_nil__2012296605.Do(func() {
		cache_nil__2012296605 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__2012296605
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList__3612943602
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

var cache_traversableList__3068288903 gopurs_runtime.Value
var once_traversableList__3068288903 sync.Once
func Get_traversableList__3068288903() gopurs_runtime.Value {
	once_traversableList__3068288903.Do(func() {
		cache_traversableList__3068288903 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "traverse"), dictApplicative_0, pkg_Data_List_Lazy_Types.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
})
}))
	})
	return cache_traversableList__3068288903
}

var cache_traversableList__2371870579 gopurs_runtime.Value
var once_traversableList__2371870579 sync.Once
func Get_traversableList__2371870579() gopurs_runtime.Value {
	once_traversableList__2371870579.Do(func() {
		cache_traversableList__2371870579 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "traverse"), dictApplicative_0, pkg_Data_List_Lazy_Types.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
})
}))
	})
	return cache_traversableList__2371870579
}

var cache_traversableList__4246548956 gopurs_runtime.Value
var once_traversableList__4246548956 sync.Once
func Get_traversableList__4246548956() gopurs_runtime.Value {
	once_traversableList__4246548956.Do(func() {
		cache_traversableList__4246548956 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "traverse"), dictApplicative_0, pkg_Data_List_Lazy_Types.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
})
}))
	})
	return cache_traversableList__4246548956
}

var cache_unfoldable1List__4025223016 gopurs_runtime.Value
var once_unfoldable1List__4025223016 sync.Once
func Get_unfoldable1List__4025223016() gopurs_runtime.Value {
	once_unfoldable1List__4025223016.Do(func() {
		cache_unfoldable1List__4025223016 = func() gopurs_runtime.Value {
var go__go_0_0_26 gopurs_runtime.Value
_ = go__go_0_0_26
go__go_0_0_26 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t7 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
__local_var_5_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.Apply2(go__go_0_0_26, f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1.UnsafePtr).V0)
_ = __local_var_6_4
__t7 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_5_3, __local_var_6_4})}
}))
goto end_branch_7
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136 && __t_tag_5.UnsafePtr == nil) {
__local_var_5_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0
_ = __local_var_5_6
__t7 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_5_6, pkg_Data_List_Lazy_Types.Get_nil()})}
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
}))
})
})
return gopurs_runtime.RecordDict1("unfoldr1", go__go_0_0_26)
}()
	})
	return cache_unfoldable1List__4025223016
}

var cache_unfoldableList__825189991 gopurs_runtime.Value
var once_unfoldableList__825189991 sync.Once
func Get_unfoldableList__825189991() gopurs_runtime.Value {
	once_unfoldableList__825189991.Do(func() {
		cache_unfoldableList__825189991 = func() gopurs_runtime.Value {
var go__go_0_0_27 gopurs_runtime.Value
_ = go__go_0_0_27
go__go_0_0_27 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr == nil) {
__t4 = pkg_Data_List_Lazy_Types.Get_nil()
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr != nil) {
__local_var_5_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_2
__local_var_6_3 := gopurs_runtime.Apply2(go__go_0_0_27, f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V1)
_ = __local_var_6_3
__t4 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_5_2, __local_var_6_3})}
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
}))
})
})
return gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_unfoldable1List()
}), go__go_0_0_27)
}()
	})
	return cache_unfoldableList__825189991
}

var cache_unfoldableList__873755030 gopurs_runtime.Value
var once_unfoldableList__873755030 sync.Once
func Get_unfoldableList__873755030() gopurs_runtime.Value {
	once_unfoldableList__873755030.Do(func() {
		cache_unfoldableList__873755030 = func() gopurs_runtime.Value {
var go__go_0_0_28 gopurs_runtime.Value
_ = go__go_0_0_28
go__go_0_0_28 = gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2))
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr == nil) {
__t4 = pkg_Data_List_Lazy_Types.Get_nil()
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr != nil) {
__local_var_5_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_2
__local_var_6_3 := gopurs_runtime.Apply2(go__go_0_0_28, f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0.UnsafePtr).V1)
_ = __local_var_6_3
__t4 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_5_2, __local_var_6_3})}
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
}))
})
})
return gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_unfoldable1List()
}), go__go_0_0_28)
}()
	})
	return cache_unfoldableList__873755030
}

var cache_alterAt__950766476 gopurs_runtime.Value
var once_alterAt__950766476 sync.Once
func Get_alterAt__950766476() gopurs_runtime.Value {
	once_alterAt__950766476.Do(func() {
		cache_alterAt__950766476 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alterAt__950766476(n_0_box.IntVal, f_1_box, xs_2_box)
})
	})
	return cache_alterAt__950766476
}

var cache_deleteAt__4024047148 gopurs_runtime.Value
var once_deleteAt__4024047148 sync.Once
func Get_deleteAt__4024047148() gopurs_runtime.Value {
	once_deleteAt__4024047148.Do(func() {
		cache_deleteAt__4024047148 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteAt__4024047148(n_0_box.IntVal, xs_1_box)
})
	})
	return cache_deleteAt__4024047148
}

var cache_deleteBy__501100275 gopurs_runtime.Value
var once_deleteBy__501100275 sync.Once
func Get_deleteBy__501100275() gopurs_runtime.Value {
	once_deleteBy__501100275.Do(func() {
		cache_deleteBy__501100275 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy__501100275(eq_0_box, x_1_box, xs_2_box)
})
	})
	return cache_deleteBy__501100275
}

var cache_drop__4024047148 gopurs_runtime.Value
var once_drop__4024047148 sync.Once
func Get_drop__4024047148() gopurs_runtime.Value {
	once_drop__4024047148.Do(func() {
		cache_drop__4024047148 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop__4024047148(n_0_box.IntVal)
})
	})
	return cache_drop__4024047148
}

var cache_filter__638755635 gopurs_runtime.Value
var once_filter__638755635 sync.Once
func Get_filter__638755635() gopurs_runtime.Value {
	once_filter__638755635.Do(func() {
		cache_filter__638755635 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter__638755635(p_0_box)
})
	})
	return cache_filter__638755635
}

var cache_filterM__647926151 gopurs_runtime.Value
var once_filterM__647926151 sync.Once
func Get_filterM__647926151() gopurs_runtime.Value {
	once_filterM__647926151.Do(func() {
		cache_filterM__647926151 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterM__647926151(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_filterM__647926151
}

var cache_findIndex__1594900290 gopurs_runtime.Value
var once_findIndex__1594900290 sync.Once
func Get_findIndex__1594900290() gopurs_runtime.Value {
	once_findIndex__1594900290.Do(func() {
		cache_findIndex__1594900290 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findIndex__1594900290(fn_0_box)
})
	})
	return cache_findIndex__1594900290
}

var cache_findLastIndex__1594900290 gopurs_runtime.Value
var once_findLastIndex__1594900290 sync.Once
func Get_findLastIndex__1594900290() gopurs_runtime.Value {
	once_findLastIndex__1594900290.Do(func() {
		cache_findLastIndex__1594900290 = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex__1594900290(fn_0_box, xs_1_box))}
})
	})
	return cache_findLastIndex__1594900290
}

var cache_foldM__3505933597 gopurs_runtime.Value
var once_foldM__3505933597 sync.Once
func Get_foldM__3505933597() gopurs_runtime.Value {
	once_foldM__3505933597.Do(func() {
		cache_foldM__3505933597 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM__3505933597(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldM__3505933597
}

var cache_fromStep__1398792641 gopurs_runtime.Value
var once_fromStep__1398792641 sync.Once
func Get_fromStep__1398792641() gopurs_runtime.Value {
	once_fromStep__1398792641.Do(func() {
		cache_fromStep__1398792641 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromStep__1398792641(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](x_0_box))
})
	})
	return cache_fromStep__1398792641
}

var cache_groupBy__1659362014 gopurs_runtime.Value
var once_groupBy__1659362014 sync.Once
func Get_groupBy__1659362014() gopurs_runtime.Value {
	once_groupBy__1659362014.Do(func() {
		cache_groupBy__1659362014 = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupBy__1659362014(eq_0_box)
})
	})
	return cache_groupBy__1659362014
}

var cache_head__2155426095 gopurs_runtime.Value
var once_head__2155426095 sync.Once
func Get_head__2155426095() gopurs_runtime.Value {
	once_head__2155426095.Do(func() {
		cache_head__2155426095 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_head__2155426095(xs_0_box))}
})
	})
	return cache_head__2155426095
}

var cache_insertAt__725610501 gopurs_runtime.Value
var once_insertAt__725610501 sync.Once
func Get_insertAt__725610501() gopurs_runtime.Value {
	once_insertAt__725610501.Do(func() {
		cache_insertAt__725610501 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAt__725610501(v_0_box.IntVal, v1_1_box, v2_2_box)
})
	})
	return cache_insertAt__725610501
}

var cache_insertBy__2098566601 gopurs_runtime.Value
var once_insertBy__2098566601 sync.Once
func Get_insertBy__2098566601() gopurs_runtime.Value {
	once_insertBy__2098566601.Do(func() {
		cache_insertBy__2098566601 = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertBy__2098566601(cmp_0_box, x_1_box, xs_2_box)
})
	})
	return cache_insertBy__2098566601
}

var cache_intersectBy__3844889126 gopurs_runtime.Value
var once_intersectBy__3844889126 sync.Once
func Get_intersectBy__3844889126() gopurs_runtime.Value {
	once_intersectBy__3844889126.Do(func() {
		cache_intersectBy__3844889126 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectBy__3844889126(eq_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_intersectBy__3844889126
}

var cache_length__162861552 gopurs_runtime.Value
var once_length__162861552 sync.Once
func Get_length__162861552() gopurs_runtime.Value {
	once_length__162861552.Do(func() {
		cache_length__162861552 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), l_0, gopurs_runtime.Int(1)).IntVal)
})
}), gopurs_runtime.Int(0))
	})
	return cache_length__162861552
}

var cache_many__956417025 gopurs_runtime.Value
var once_many__956417025 sync.Once
func Get_many__956417025() gopurs_runtime.Value {
	once_many__956417025.Do(func() {
		cache_many__956417025 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many__956417025(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_many__956417025
}

var cache_mapMaybe__3574309085 gopurs_runtime.Value
var once_mapMaybe__3574309085 sync.Once
func Get_mapMaybe__3574309085() gopurs_runtime.Value {
	once_mapMaybe__3574309085.Do(func() {
		cache_mapMaybe__3574309085 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__3574309085(f_0_box)
})
	})
	return cache_mapMaybe__3574309085
}

var cache_mapMaybe__1687744733 gopurs_runtime.Value
var once_mapMaybe__1687744733 sync.Once
func Get_mapMaybe__1687744733() gopurs_runtime.Value {
	once_mapMaybe__1687744733.Do(func() {
		cache_mapMaybe__1687744733 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__1687744733(f_0_box)
})
	})
	return cache_mapMaybe__1687744733
}

var cache_mapMaybe__899591645 gopurs_runtime.Value
var once_mapMaybe__899591645 sync.Once
func Get_mapMaybe__899591645() gopurs_runtime.Value {
	once_mapMaybe__899591645.Do(func() {
		cache_mapMaybe__899591645 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__899591645(f_0_box)
})
	})
	return cache_mapMaybe__899591645
}

var cache_mapMaybe__600226685 gopurs_runtime.Value
var once_mapMaybe__600226685 sync.Once
func Get_mapMaybe__600226685() gopurs_runtime.Value {
	once_mapMaybe__600226685.Do(func() {
		cache_mapMaybe__600226685 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__600226685(f_0_box)
})
	})
	return cache_mapMaybe__600226685
}

var cache_nubBy__2220739616 gopurs_runtime.Value
var once_nubBy__2220739616 sync.Once
func Get_nubBy__2220739616() gopurs_runtime.Value {
	once_nubBy__2220739616.Do(func() {
		cache_nubBy__2220739616 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubBy__2220739616(p_0_box)
})
	})
	return cache_nubBy__2220739616
}

var cache_nubByEq__616397370 gopurs_runtime.Value
var once_nubByEq__616397370 sync.Once
func Get_nubByEq__616397370() gopurs_runtime.Value {
	once_nubByEq__616397370.Do(func() {
		cache_nubByEq__616397370 = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubByEq__616397370(eq_0_box)
})
	})
	return cache_nubByEq__616397370
}

var cache_null__1674339719 gopurs_runtime.Value
var once_null__1674339719 sync.Once
func Get_null__1674339719() gopurs_runtime.Value {
	once_null__1674339719.Do(func() {
		cache_null__1674339719 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null__1674339719(x_0_box))
})
	})
	return cache_null__1674339719
}

var cache_repeat__2149902581 gopurs_runtime.Value
var once_repeat__2149902581 sync.Once
func Get_repeat__2149902581() gopurs_runtime.Value {
	once_repeat__2149902581.Do(func() {
		cache_repeat__2149902581 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat__2149902581(x_0_box)
})
	})
	return cache_repeat__2149902581
}

var cache_replicateM__3816548429 gopurs_runtime.Value
var once_replicateM__3816548429 sync.Once
func Get_replicateM__3816548429() gopurs_runtime.Value {
	once_replicateM__3816548429.Do(func() {
		cache_replicateM__3816548429 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicateM__3816548429(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_replicateM__3816548429
}

var cache_reverse__1315655552 gopurs_runtime.Value
var once_reverse__1315655552 sync.Once
func Get_reverse__1315655552() gopurs_runtime.Value {
	once_reverse__1315655552.Do(func() {
		cache_reverse__1315655552 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reverse__1315655552(xs_0_box)
})
	})
	return cache_reverse__1315655552
}

var cache_scanlLazy__3747838620 gopurs_runtime.Value
var once_scanlLazy__3747838620 sync.Once
func Get_scanlLazy__3747838620() gopurs_runtime.Value {
	once_scanlLazy__3747838620.Do(func() {
		cache_scanlLazy__3747838620 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanlLazy__3747838620(f_0_box, acc_1_box, xs_2_box)
})
	})
	return cache_scanlLazy__3747838620
}

var cache_some__956417025 gopurs_runtime.Value
var once_some__956417025 sync.Once
func Get_some__956417025() gopurs_runtime.Value {
	once_some__956417025.Do(func() {
		cache_some__956417025 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some__956417025(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_some__956417025
}

var cache_span__776304907 gopurs_runtime.Value
var once_span__776304907 sync.Once
func Get_span__776304907() gopurs_runtime.Value {
	once_span__776304907.Do(func() {
		cache_span__776304907 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span__776304907(p_0_box, xs_1_box)
})
	})
	return cache_span__776304907
}

var cache_tail__1935051898 gopurs_runtime.Value
var once_tail__1935051898 sync.Once
func Get_tail__1935051898() gopurs_runtime.Value {
	once_tail__1935051898.Do(func() {
		cache_tail__1935051898 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_tail__1935051898(xs_0_box))}
})
	})
	return cache_tail__1935051898
}

var cache_take__4024047148 gopurs_runtime.Value
var once_take__4024047148 sync.Once
func Get_take__4024047148() gopurs_runtime.Value {
	once_take__4024047148.Do(func() {
		cache_take__4024047148 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take__4024047148(n_0_box.IntVal)
})
	})
	return cache_take__4024047148
}

var cache_takeWhile__638755635 gopurs_runtime.Value
var once_takeWhile__638755635 sync.Once
func Get_takeWhile__638755635() gopurs_runtime.Value {
	once_takeWhile__638755635.Do(func() {
		cache_takeWhile__638755635 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile__638755635(p_0_box)
})
	})
	return cache_takeWhile__638755635
}

var cache_transpose__1534541312 gopurs_runtime.Value
var once_transpose__1534541312 sync.Once
func Get_transpose__1534541312() gopurs_runtime.Value {
	once_transpose__1534541312.Do(func() {
		cache_transpose__1534541312 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_transpose__1534541312(xs_0_box)
})
	})
	return cache_transpose__1534541312
}

var cache_uncons__3647012005 gopurs_runtime.Value
var once_uncons__3647012005 sync.Once
func Get_uncons__3647012005() gopurs_runtime.Value {
	once_uncons__3647012005.Do(func() {
		cache_uncons__3647012005 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__3647012005(xs_0_box))}
})
	})
	return cache_uncons__3647012005
}

var cache_uncons__1321764894 gopurs_runtime.Value
var once_uncons__1321764894 sync.Once
func Get_uncons__1321764894() gopurs_runtime.Value {
	once_uncons__1321764894.Do(func() {
		cache_uncons__1321764894 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__1321764894(xs_0_box))}
})
	})
	return cache_uncons__1321764894
}

var cache_uncons__974566859 gopurs_runtime.Value
var once_uncons__974566859 sync.Once
func Get_uncons__974566859() gopurs_runtime.Value {
	once_uncons__974566859.Do(func() {
		cache_uncons__974566859 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__974566859(xs_0_box))}
})
	})
	return cache_uncons__974566859
}

var cache_uncons__1420258522 gopurs_runtime.Value
var once_uncons__1420258522 sync.Once
func Get_uncons__1420258522() gopurs_runtime.Value {
	once_uncons__1420258522.Do(func() {
		cache_uncons__1420258522 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__1420258522(xs_0_box))}
})
	})
	return cache_uncons__1420258522
}

var cache_unionBy__3844889126 gopurs_runtime.Value
var once_unionBy__3844889126 sync.Once
func Get_unionBy__3844889126() gopurs_runtime.Value {
	once_unionBy__3844889126.Do(func() {
		cache_unionBy__3844889126 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionBy__3844889126(eq_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_unionBy__3844889126
}

var cache_updateAt__725610501 gopurs_runtime.Value
var once_updateAt__725610501 sync.Once
func Get_updateAt__725610501() gopurs_runtime.Value {
	once_updateAt__725610501.Do(func() {
		cache_updateAt__725610501 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAt__725610501(n_0_box.IntVal, x_1_box, xs_2_box)
})
	})
	return cache_updateAt__725610501
}

var cache_zipWith__3539178005 gopurs_runtime.Value
var once_zipWith__3539178005 sync.Once
func Get_zipWith__3539178005() gopurs_runtime.Value {
	once_zipWith__3539178005.Do(func() {
		cache_zipWith__3539178005 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith__3539178005(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith__3539178005
}

var cache_zipWith__3210333397 gopurs_runtime.Value
var once_zipWith__3210333397 sync.Once
func Get_zipWith__3210333397() gopurs_runtime.Value {
	once_zipWith__3210333397.Do(func() {
		cache_zipWith__3210333397 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith__3210333397(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith__3210333397
}

var cache_zipWith__3984071349 gopurs_runtime.Value
var once_zipWith__3984071349 sync.Once
func Get_zipWith__3984071349() gopurs_runtime.Value {
	once_zipWith__3984071349.Do(func() {
		cache_zipWith__3984071349 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith__3984071349(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith__3984071349
}

var cache_applicativeMaybe__3016118221 gopurs_runtime.Value
var once_applicativeMaybe__3016118221 sync.Once
func Get_applicativeMaybe__3016118221() gopurs_runtime.Value {
	once_applicativeMaybe__3016118221.Do(func() {
		cache_applicativeMaybe__3016118221 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), pkg_Data_Maybe.Get_Just())
	})
	return cache_applicativeMaybe__3016118221
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)))}
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

var cache_bindMaybe__1910292045 gopurs_runtime.Value
var once_bindMaybe__1910292045 sync.Once
func Get_bindMaybe__1910292045() gopurs_runtime.Value {
	once_bindMaybe__1910292045.Do(func() {
		cache_bindMaybe__1910292045 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_bindMaybe__1910292045
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

var cache_isNothing__1401305026 gopurs_runtime.Value
var once_isNothing__1401305026 sync.Once
func Get_isNothing__1401305026() gopurs_runtime.Value {
	once_isNothing__1401305026.Do(func() {
		cache_isNothing__1401305026 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__1401305026(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__1401305026
}

var cache_maybe__2641488518 gopurs_runtime.Value
var once_maybe__2641488518 sync.Once
func Get_maybe__2641488518() gopurs_runtime.Value {
	once_maybe__2641488518.Do(func() {
		cache_maybe__2641488518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__2641488518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__2641488518
}

var cache_maybe__3931678292 gopurs_runtime.Value
var once_maybe__3931678292 sync.Once
func Get_maybe__3931678292() gopurs_runtime.Value {
	once_maybe__3931678292.Do(func() {
		cache_maybe__3931678292 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3931678292(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3931678292
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
		cache_unwrap__831442259 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_unwrap__831442259
}

var cache_unwrap__4291124211 gopurs_runtime.Value
var once_unwrap__4291124211 sync.Once
func Get_unwrap__4291124211() gopurs_runtime.Value {
	once_unwrap__4291124211.Do(func() {
		cache_unwrap__4291124211 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_unwrap__4291124211
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

var cache_compare__2349537221 gopurs_runtime.Value
var once_compare__2349537221 sync.Once
func Get_compare__2349537221() gopurs_runtime.Value {
	once_compare__2349537221.Do(func() {
		cache_compare__2349537221 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__2349537221(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__2349537221
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

var cache_greaterThanOrEq__4087042607 gopurs_runtime.Value
var once_greaterThanOrEq__4087042607 sync.Once
func Get_greaterThanOrEq__4087042607() gopurs_runtime.Value {
	once_greaterThanOrEq__4087042607.Do(func() {
		cache_greaterThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__4087042607
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
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

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
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

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
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
		cache_append__2734706680 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append")
	})
	return cache_append__2734706680
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
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

var cache_sequence__2904194897 gopurs_runtime.Value
var once_sequence__2904194897 sync.Once
func Get_sequence__2904194897() gopurs_runtime.Value {
	once_sequence__2904194897.Do(func() {
		cache_sequence__2904194897 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "sequence")
	})
	return cache_sequence__2904194897
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

var cache_traverse__1157172365 gopurs_runtime.Value
var once_traverse__1157172365 sync.Once
func Get_traverse__1157172365() gopurs_runtime.Value {
	once_traverse__1157172365.Do(func() {
		cache_traverse__1157172365 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "traverse")
	})
	return cache_traverse__1157172365
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

var cache_unfoldr__1779806517 gopurs_runtime.Value
var once_unfoldr__1779806517 sync.Once
func Get_unfoldr__1779806517() gopurs_runtime.Value {
	once_unfoldr__1779806517.Do(func() {
		cache_unfoldr__1779806517 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1779806517(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1779806517
}

var cache_unfoldr__193332035 gopurs_runtime.Value
var once_unfoldr__193332035 sync.Once
func Get_unfoldr__193332035() gopurs_runtime.Value {
	once_unfoldr__193332035.Do(func() {
		cache_unfoldr__193332035 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_unfoldableList(), "unfoldr")
	})
	return cache_unfoldr__193332035
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Pattern(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
zipWith:
for {
if false { continue zipWith }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), Call_zipWith(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}), xs_1), ys_2)
}
}

func Call_zipWithA(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value, ys_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 gopurs_runtime.Value = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, Call_zipWith(f_1, xs_2, ys_3))
}

func Call_updateAt(n_0_loop int64, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
updateAt:
for {
if false { continue updateAt }
var n_0 int64 = n_0_loop
_ = n_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t0 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (n_0) == (0) {
__t0 = &pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1}
goto end_branch_0
} else {

}
}
{
__t0 = &pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_updateAt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
}
end_branch_0:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t0)}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}), xs_2)
}
}

func Call_uncons(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1)})}
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

func Call_toUnfoldable(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.RecordGet(rec_2, "tail")})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_1))})))}
}))
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0).IntVal) != (0)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_takeWhile(p_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
}

func Call_take(n_0_loop int64) gopurs_runtime.Value {
take:
for {
if false { continue take }
var n_0 int64 = n_0_loop
_ = n_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(n_0), gopurs_runtime.Int(0))).IntVal) != (0) {
__t2 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_nil()
})
goto end_branch_2
} else {

}
}
{
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 218341868 && v1_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 218341868 && v1_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_take(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}))
_ = __local_var_1_0
__t2 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
end_branch_2:
return __t2
}
}

func Call_tail(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "tail")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}))
}

func Call_stripPrefix(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_Rec_Class.Get_monadRecMaybe(), "tailRecM"), gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(o_3, "a")))
_ = v1_4_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_4_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_4_0)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(o_3, "b")})}})}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_4_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_4_0)}.UnsafePtr != nil) {
v2_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(o_3, "b")))
_ = v2_5_1
var __t2 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v2_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v2_5_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v2_5_1)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(dictEq_0.V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_4_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v2_5_1)}.UnsafePtr).V0).IntVal) != (0)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_4_0)}.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v2_5_1)}.UnsafePtr).V1)})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.RecordDict2("a", "b", v_1, s_2)))
}

func Call_span(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
span:
for {
if false { continue span }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
v_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_1))})
_ = v_2_0
var __t4 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0, "head")).IntVal) != (0)) {
__local_var_3_1 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0, "head")
_ = __local_var_3_1
v1_4_2 := Call_span(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0, "tail"))
_ = v1_4_2
__local_var_5_3 := gopurs_runtime.RecordGet(v1_4_2, "init")
_ = __local_var_5_3
__t4 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_5_3})}
})), gopurs_runtime.RecordGet(v1_4_2, "rest"))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.RecordDict2("init", "rest", pkg_Data_List_Lazy_Types.Get_nil(), xs_1)
}
end_branch_4:
return __t4
}
}

func Call_snoc(xs_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, pkg_Data_List_Lazy_Types.Get_nil()})}
})), xs_0)
}

func Call_singleton(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, pkg_Data_List_Lazy_Types.Get_nil()})}
}))
}

func Call_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showList_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_showList(), dictShow_0))
_ = showList_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Pattern "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(showList_1_0.V0, v_2), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_scanlLazy(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
scanlLazy:
for {
if false { continue scanlLazy }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
acc_prime_4_0 := gopurs_runtime.Apply2(f_0, acc_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
_ = acc_prime_4_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, acc_prime_4_0, Call_scanlLazy(f_0, acc_prime_4_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}), xs_2)
}
}

func Call_reverse(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_3, b_2})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_0)
}))
}

func Call_replicateM(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
replicateM:
for {
if false { continue replicateM }
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(n_3, gopurs_runtime.Int(1))).IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(Bind1_2_1.V1, m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_replicateM(dictMonad_0), gopurs_runtime.Apply2(Get_sub__1043827704(), n_3, gopurs_runtime.Int(1)), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_5, as_6})}
})))
}))
}))
}
end_branch_2:
return __t2
})
})
}
}

func Call_repeat(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
go__go_1_0_0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_0})}
}))
}))
return go__go_1_0_0
}

func Call_replicate(i_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var go__go_2_0_1 gopurs_runtime.Value
_ = go__go_2_0_1
go__go_2_0_1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, xs_1, go__go_2_0_1})}
}))
}))
return gopurs_runtime.Apply(Call_take(i_0), go__go_2_0_1)
}

func Call_go__range(start_0_loop int64, end_1_loop int64) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1))).IntVal) != (0) {
__t2 = gopurs_runtime.Apply2(Get_unfoldr__193332035(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(x_2, gopurs_runtime.Int(end_1))).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2, gopurs_runtime.Apply2(Get_sub__1043827704(), x_2, gopurs_runtime.Int(1))})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[int64, int64]]](__t1))}
}), gopurs_runtime.Int(start_0))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(Get_unfoldr__193332035(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(x_2, gopurs_runtime.Int(end_1))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2, gopurs_runtime.Apply2(Get_add__560788792(), x_2, gopurs_runtime.Int(1))})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[int64, int64]]](__t0))}
}), gopurs_runtime.Int(start_0))
}
end_branch_2:
return __t2
}

func Call_partition(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, x_1).IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.RecordGet(v_2, "no"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.RecordGet(v_2, "yes")})}
})))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.RecordGet(v_2, "no")})}
})), gopurs_runtime.RecordGet(v_2, "yes"))
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("no", "yes", pkg_Data_List_Lazy_Types.Get_nil(), pkg_Data_List_Lazy_Types.Get_nil()))
}

func Call_null(x_0_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(x_0))}
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var goStep_1_0_2 gopurs_runtime.Value
_ = goStep_1_0_2
var go__go_1_1_3 gopurs_runtime.Value
_ = go__go_1_1_3
goStep_1_0_2 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
v2_4_2 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2)
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v2_4_2, "found").IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(go__go_1_1_3, gopurs_runtime.RecordGet(v2_4_2, "result"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1))))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_1_1_3, gopurs_runtime.RecordGet(v2_4_2, "result"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
})
})
go__go_1_1_3 = gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(goStep_1_0_2, s_2), v_3)
})
})
return gopurs_runtime.Apply(go__go_1_1_3, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_nub(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_nubBy(dictOrd_0.V1)
}

func Call_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_4 gopurs_runtime.Value
go__go_1_0_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_4:
for {
if false { continue go__go_1_0_4 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
v1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr == nil) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_4
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr).V0, gopurs_runtime.Apply(Call_mapMaybe(f_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_4)
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, x_3)
})
}
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
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_many(dictAlternative_0), dictLazy_3, v_4)
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
return gopurs_runtime.Apply2(Alt0_1_0.V1, gopurs_runtime.Apply2(Call_some(dictAlternative_0), dictLazy_3, v_4), gopurs_runtime.Apply(Applicative0_2_1.V1, pkg_Data_List_Lazy_Types.Get_nil()))
})
})
}

func Call_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_6 gopurs_runtime.Value
_ = go__go_2_0_6
go__go_2_0_6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, go__go_2_0_6)
_ = __local_var_4_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, __local_var_4_1})}
}))
}))
return go__go_2_0_6
}

func Call_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertAt:
for {
if false { continue insertAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})}
}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 218341868 && v3_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, pkg_Data_List_Lazy_Types.Get_nil()})}
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 218341868 && v3_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V0, Call_insertAt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, v1_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), v2_2)
}
end_branch_1:
return __t1
}
}

func Call_index(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_8 gopurs_runtime.Value
go__go_1_0_8 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_8:
for {
if false { continue go__go_1_0_8 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t2 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (v1_3.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
v1_3_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), v1_3, gopurs_runtime.Int(1))
continue go__go_1_0_8
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_8, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
}

func Call_head(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "head")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}))
}

func Call_transpose(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
transpose:
for {
if false { continue transpose }
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))})
_ = v_1_0
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t9 = xs_0
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
v1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, "head")))})
_ = v1_2_1
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.UnsafePtr == nil) {
xs_0_loop = gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, "tail")
continue transpose
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.UnsafePtr != nil) {
__local_var_3_2 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.UnsafePtr).V0, "head")
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.UnsafePtr).V0, "tail")
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply(Call_mapMaybe(Get_head()), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, "tail"))
_ = __local_var_5_4
__local_var_6_5 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_2, __local_var_5_4})}
}))
_ = __local_var_6_5
__local_var_7_7 := gopurs_runtime.Apply(Call_mapMaybe(Get_tail()), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, "tail"))
_ = __local_var_7_7
__local_var_7_6 := Call_transpose(gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_3, __local_var_7_7})}
})))
_ = __local_var_7_6
__t8 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_5, __local_var_7_6})}
}))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = __t8
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

func Call_groupBy(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
groupBy:
for {
if false { continue groupBy }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_4
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
__local_var_2_1 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_1
v1_3_2 := Call_span(gopurs_runtime.Apply(eq_0, __local_var_2_1), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
_ = v1_3_2
__local_var_4_3 := gopurs_runtime.RecordGet(v1_3_2, "init")
_ = __local_var_4_3
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_1, __local_var_4_3})}
})), gopurs_runtime.Apply(Call_groupBy(eq_0), gopurs_runtime.RecordGet(v1_3_2, "rest"))})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
}

func Call_group(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return Call_groupBy(dictEq_0.V0)
}

func Call_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertBy:
for {
if false { continue insertBy }
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, pkg_Data_List_Lazy_Types.Get_nil()})}
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
var __t1 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(cmp_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = &pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, Call_insertBy(cmp_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = &pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), v_3)}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}), xs_2)
}
}

func Call_insert(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_insertBy(), dictOrd_0.V1)
}

func Call_fromFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V2, pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
}

func Call_foldrLazy(dictLazy_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value], op_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
_ = dictLazy_0
var op_1 gopurs_runtime.Value = op_1_loop
_ = op_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
var go__go_3_0_9 gopurs_runtime.Value
_ = go__go_3_0_9
go__go_3_0_9 = gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_4))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr != nil) {
__local_var_6_2 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0
_ = __local_var_6_2
__local_var_7_3 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply(dictLazy_0.V0, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_1, __local_var_6_2, gopurs_runtime.Apply(go__go_3_0_9, __local_var_7_3))
}))
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr == nil) {
__t4 = z_2
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
return go__go_3_0_9
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
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_5))})
_ = v_6_2
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, b_4)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.UnsafePtr != nil) {
__local_var_7_3 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.UnsafePtr).V0, "tail")
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(f_3, b_4, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.UnsafePtr).V0, "head")), gopurs_runtime.Func(func(b_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_foldM(dictMonad_0), f_3, b_prime_8, __local_var_7_3)
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
})
})
})
}
}

func Call_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_10 gopurs_runtime.Value
_ = go__go_1_0_10
go__go_1_0_10 = gopurs_runtime.Func(func(n_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(list_3))}, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *pkg_Data_Maybe.Constructor_Just[int64]
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet(o_4, "head")).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Get_pure__577281046(), n_2))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(go__go_1_0_10, gopurs_runtime.Apply2(Get_add__560788792(), n_2, gopurs_runtime.Int(1)), gopurs_runtime.RecordGet(o_4, "tail")))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))))}
})
})
return gopurs_runtime.Apply(go__go_1_0_10, gopurs_runtime.Int(0))
}

func Call_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(Get_map__291265340(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply(Get_length(), xs_1), gopurs_runtime.Int(1)), v_2).IntVal)
}), gopurs_runtime.Apply(Call_findIndex(fn_0), Call_reverse(xs_1))))
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
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(list_4))})
_ = v_5_2
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.UnsafePtr != nil) {
__local_var_6_3 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.UnsafePtr).V0, "head")
_ = __local_var_6_3
__local_var_7_4 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.UnsafePtr).V0, "tail")
_ = __local_var_7_4
__t6 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(p_3, __local_var_6_3), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_filterM(dictMonad_0), p_3, __local_var_7_4), gopurs_runtime.Func(func(xs_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (b_8.IntVal) != (0) {
__t5 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_3, xs_prime_9})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = xs_prime_9
}
end_branch_5:
return gopurs_runtime.Apply(Applicative0_1_0.V1, __t5)
}))
}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
})
}
}

func Call_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_11 gopurs_runtime.Value
go__go_1_0_11 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_11:
for {
if false { continue go__go_1_0_11 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t2 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_filter(p_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_11
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_11)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, x_3)
})
}
}

func Call_intersectBy(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_4, v1_5)
})
}))
_ = semigroupDisj1_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_4_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")), gopurs_runtime.Apply(eq_0, x_3), ys_2).IntVal) != (0))
})), xs_1)
}

func Call_intersect(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy(), dictEq_0.V0)
}

func Call_nubByEq(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
nubByEq:
for {
if false { continue nubByEq }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
__local_var_2_1 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_1, gopurs_runtime.Apply(Call_nubByEq(eq_0), gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(Get_not__3201284355(), gopurs_runtime.Apply2(eq_0, __local_var_2_1, y_3)).IntVal) != (0))
})), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
}

func Call_nubEq(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return Call_nubByEq(dictEq_0.V0)
}

func Call_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
eqList_1_0 := &pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_eq1List(), "eq1"), dictEq_0)}
_ = eqList_1_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eqList_1_0.V0, x_2, y_3).IntVal) != (0))
})
}))
}

func Call_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordList_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_ordList(), dictOrd_0))
_ = ordList_1_0
eqPattern1_2_1 := Call_eqPattern(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqPattern1_2_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqPattern1_2_1
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(ordList_1_0.V1, x_3, y_4).IntVal)), UnsafePtr: nil}
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
var go__go_1_0_12 gopurs_runtime.Value
go__go_1_0_12 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_12:
for {
if false { continue go__go_1_0_12 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if ((v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_12
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), v_2)
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_12, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_2))
})
}

func Call_drop(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var go__go_1_0_13 gopurs_runtime.Value
go__go_1_0_13 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
go__go_1_0_13:
for {
if false { continue go__go_1_0_13 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), v_2, gopurs_runtime.Int(1))
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1))
continue go__go_1_0_13
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(go__go_1_0_13, gopurs_runtime.Int(n_0)))
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, x_3)
})
}

func Call_slice(start_0_loop int64, end_1_loop int64, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Call_take(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(end_1), gopurs_runtime.Int(start_0)).IntVal), gopurs_runtime.Apply(Call_drop(start_0), xs_2))
}

func Call_deleteBy(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteBy:
for {
if false { continue deleteBy }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(eq_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, Call_deleteBy(eq_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}), xs_2)
}
}

func Call_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), xs_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy(eq_0, a_4, b_3)
})
}), gopurs_runtime.Apply(Call_nubByEq(eq_0), ys_2), xs_1))
}

func Call_union(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy(), dictEq_0.V0)
}

func Call_deleteAt(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteAt:
for {
if false { continue deleteAt }
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 218341868 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 218341868 && v1_2.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (n_0) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, Call_deleteAt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}), xs_1)
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
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy(dictEq_0.V0, a_2, b_1)
})
}))
}

func Call_cycle(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_14 gopurs_runtime.Value
_ = go__go_1_0_14
go__go_1_0_14 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), xs_0, go__go_1_0_14)
}))
return go__go_1_0_14
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindList(), "bind"), a_1, b_0)
}

func Call_concat(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindList(), "bind"), v_0, Get_identity())
}

func Call_alterAt(n_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
alterAt:
for {
if false { continue alterAt }
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (n_0) == (0) {
v2_4_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_alterAt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal, f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}), xs_2)
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

func Call_alt__4218059372(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
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

func Call_bind__2073074151(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__777862183(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__523053991(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_defer__3967925939(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_defer__774977193(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fix__1475205859(dictLazy_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
_ = dictLazy_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_15 gopurs_runtime.Value
_ = go__go_2_0_15
go__go_2_0_15 = gopurs_runtime.Apply(dictLazy_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, go__go_2_0_15)
}))
return go__go_2_0_15
}

func Call_fix__3570066147(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_16 gopurs_runtime.Value
_ = go__go_1_0_16
go__go_1_0_16 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, go__go_1_0_16)
}))
return go__go_1_0_16
}

func Call_tailRec__2110844386(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_18 gopurs_runtime.Value
go__go_1_0_18 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_18:
for {
if false { continue go__go_1_0_18 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_18
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
return gopurs_runtime.Apply(go__go_1_0_18, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRec__2666749533(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_19 gopurs_runtime.Value
go__go_1_0_19 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_19:
for {
if false { continue go__go_1_0_19 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_19
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
return gopurs_runtime.Apply(go__go_1_0_19, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRecM__3865988408(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_tailRecM2__3864450856(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_eq__2484408063(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2403185435(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__1245418657(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__754826900(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__4115900574(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
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

func Call_flip__2422603136(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
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

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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
var up_3_0_20 gopurs_runtime.Value
up_3_0_20 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
up_3_0_20:
for {
if false { continue up_3_0_20 }
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
continue up_3_0_20
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
continue up_3_0_20
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
continue up_3_0_20
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
var down_4_8_21 gopurs_runtime.Value
down_4_8_21 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
down_4_8_21:
for {
if false { continue down_4_8_21 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t15 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 2764020654) {
__t15 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(up_3_0_20, v_5, gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, k_1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}})}))
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
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_21
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_21
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
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_21
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
if ((uint32(v2_7_11.IntVal) == 380165415)) && ((uint32(v3_8_12.IntVal) == 1527465420)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_21
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4
continue down_4_8_21
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
return gopurs_runtime.Apply2(down_4_8_21, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, orig_2)
}

func Call_cons__716923058(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_cons__2305074921(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_cons__891310957(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_step__3545407802(x_0_loop gopurs_runtime.Value) *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_step__4184651873(x_0_loop gopurs_runtime.Value) *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_alterAt__950766476(n_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (n_0) == (0) {
v2_4_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_alterAt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal, f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}), xs_2)
}

func Call_deleteAt__4024047148(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 218341868 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 218341868 && v1_2.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (n_0) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, Call_deleteAt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}), xs_1)
}

func Call_deleteBy__501100275(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(eq_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, Call_deleteBy(eq_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}), xs_2)
}

func Call_drop__4024047148(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var go__go_1_0_29 gopurs_runtime.Value
go__go_1_0_29 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
go__go_1_0_29:
for {
if false { continue go__go_1_0_29 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), v_2, gopurs_runtime.Int(1))
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1))
continue go__go_1_0_29
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(go__go_1_0_29, gopurs_runtime.Int(n_0)))
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, x_3)
})
}

func Call_filter__638755635(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_30 gopurs_runtime.Value
go__go_1_0_30 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_30:
for {
if false { continue go__go_1_0_30 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t2 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_filter(p_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_30
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_30)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, x_3)
})
}

func Call_filterM__647926151(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(list_4))})
_ = v_5_2
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.UnsafePtr != nil) {
__local_var_6_3 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.UnsafePtr).V0, "head")
_ = __local_var_6_3
__local_var_7_4 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_2)}.UnsafePtr).V0, "tail")
_ = __local_var_7_4
__t6 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(p_3, __local_var_6_3), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_filterM(dictMonad_0), p_3, __local_var_7_4), gopurs_runtime.Func(func(xs_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (b_8.IntVal) != (0) {
__t5 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_3, xs_prime_9})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = xs_prime_9
}
end_branch_5:
return gopurs_runtime.Apply(Applicative0_1_0.V1, __t5)
}))
}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
})
}

func Call_findIndex__1594900290(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_31 gopurs_runtime.Value
_ = go__go_1_0_31
go__go_1_0_31 = gopurs_runtime.Func(func(n_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(list_3))}, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *pkg_Data_Maybe.Constructor_Just[int64]
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet(o_4, "head")).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure"), n_2))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(go__go_1_0_31, gopurs_runtime.Apply2(Get_add__560788792(), n_2, gopurs_runtime.Int(1)), gopurs_runtime.RecordGet(o_4, "tail")))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))))}
})
})
return gopurs_runtime.Apply(go__go_1_0_31, gopurs_runtime.Int(0))
}

func Call_findLastIndex__1594900290(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply(Get_length(), xs_1), gopurs_runtime.Int(1)), v_2).IntVal)
}), gopurs_runtime.Apply(Call_findIndex(fn_0), Call_reverse(xs_1))))
}

func Call_foldM__3505933597(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_5))})
_ = v_6_2
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, b_4)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.UnsafePtr != nil) {
__local_var_7_3 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.UnsafePtr).V0, "tail")
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(f_3, b_4, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_2)}.UnsafePtr).V0, "head")), gopurs_runtime.Func(func(b_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_foldM(dictMonad_0), f_3, b_prime_8, __local_var_7_3)
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
})
})
})
}

func Call_fromStep__1398792641(x_0_loop *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var x_0 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = x_0_loop
_ = x_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)})
}

func Call_groupBy__1659362014(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_4
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
__local_var_2_1 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_1
v1_3_2 := Call_span(gopurs_runtime.Apply(eq_0, __local_var_2_1), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
_ = v1_3_2
__local_var_4_3 := gopurs_runtime.RecordGet(v1_3_2, "init")
_ = __local_var_4_3
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_1, __local_var_4_3})}
})), gopurs_runtime.Apply(Call_groupBy(eq_0), gopurs_runtime.RecordGet(v1_3_2, "rest"))})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_head__2155426095(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "head")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}))
}

func Call_insertAt__725610501(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})}
}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 218341868 && v3_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, pkg_Data_List_Lazy_Types.Get_nil()})}
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 218341868 && v3_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V0, Call_insertAt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, v1_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), v2_2)
}
end_branch_1:
return __t1
}

func Call_insertBy__2098566601(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, pkg_Data_List_Lazy_Types.Get_nil()})}
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
var __t1 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(cmp_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = &pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, Call_insertBy(cmp_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = &pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), v_3)}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}), xs_2)
}

func Call_intersectBy__3844889126(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_4, v1_5)
})
}))
_ = semigroupDisj1_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_4_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")), gopurs_runtime.Apply(eq_0, x_3), ys_2).IntVal) != (0))
})), xs_1)
}

func Call_many__956417025(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Alt0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Alt0_1_0.V1, gopurs_runtime.Apply2(Call_some(dictAlternative_0), dictLazy_3, v_4), gopurs_runtime.Apply(Applicative0_2_1.V1, pkg_Data_List_Lazy_Types.Get_nil()))
})
})
}

func Call_mapMaybe__3574309085(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_32 gopurs_runtime.Value
go__go_1_0_32 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_32:
for {
if false { continue go__go_1_0_32 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
v1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr == nil) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_32
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr).V0, gopurs_runtime.Apply(Call_mapMaybe(f_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_32)
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, x_3)
})
}

func Call_mapMaybe__1687744733(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_33 gopurs_runtime.Value
go__go_1_0_33 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_33:
for {
if false { continue go__go_1_0_33 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
v1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr == nil) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_33
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr).V0, gopurs_runtime.Apply(Call_mapMaybe(f_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_33)
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, x_3)
})
}

func Call_mapMaybe__899591645(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_34 gopurs_runtime.Value
go__go_1_0_34 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_34:
for {
if false { continue go__go_1_0_34 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
v1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr == nil) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_34
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr).V0, gopurs_runtime.Apply(Call_mapMaybe(f_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_34)
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, x_3)
})
}

func Call_mapMaybe__600226685(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_35 gopurs_runtime.Value
go__go_1_0_35 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_35:
for {
if false { continue go__go_1_0_35 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
v1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr == nil) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_35
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_3_1)}.UnsafePtr).V0, gopurs_runtime.Apply(Call_mapMaybe(f_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_35)
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, x_3)
})
}

func Call_nubBy__2220739616(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var goStep_1_0_36 gopurs_runtime.Value
_ = goStep_1_0_36
var go__go_1_1_37 gopurs_runtime.Value
_ = go__go_1_1_37
goStep_1_0_36 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
v2_4_2 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2)
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v2_4_2, "found").IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(go__go_1_1_37, gopurs_runtime.RecordGet(v2_4_2, "result"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1))))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_1_1_37, gopurs_runtime.RecordGet(v2_4_2, "result"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
})
})
go__go_1_1_37 = gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(goStep_1_0_36, s_2), v_3)
})
})
return gopurs_runtime.Apply(go__go_1_1_37, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_nubByEq__616397370(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
__local_var_2_1 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_1, gopurs_runtime.Apply(Call_nubByEq(eq_0), gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply2(eq_0, __local_var_2_1, y_3)).IntVal) != (0))
})), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_null__1674339719(x_0_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(x_0))}
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_repeat__2149902581(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_38 gopurs_runtime.Value
_ = go__go_1_0_38
go__go_1_0_38 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_38})}
}))
}))
return go__go_1_0_38
}

func Call_replicateM__3816548429(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(n_3, gopurs_runtime.Int(1))).IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(Bind1_2_1.V1, m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_replicateM(dictMonad_0), gopurs_runtime.Apply2(Get_sub__1043827704(), n_3, gopurs_runtime.Int(1)), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_5, as_6})}
})))
}))
}))
}
end_branch_2:
return __t2
})
})
}

func Call_reverse__1315655552(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_3, b_2})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_0)
}))
}

func Call_scanlLazy__3747838620(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
acc_prime_4_0 := gopurs_runtime.Apply2(f_0, acc_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
_ = acc_prime_4_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, acc_prime_4_0, Call_scanlLazy(f_0, acc_prime_4_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}), xs_2)
}

func Call_some__956417025(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_many(dictAlternative_0), dictLazy_3, v_4)
})))
})
})
}

func Call_span__776304907(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
v_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_1))})
_ = v_2_0
var __t4 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0, "head")).IntVal) != (0)) {
__local_var_3_1 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0, "head")
_ = __local_var_3_1
v1_4_2 := Call_span(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0, "tail"))
_ = v1_4_2
__local_var_5_3 := gopurs_runtime.RecordGet(v1_4_2, "init")
_ = __local_var_5_3
__t4 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_5_3})}
})), gopurs_runtime.RecordGet(v1_4_2, "rest"))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.RecordDict2("init", "rest", pkg_Data_List_Lazy_Types.Get_nil(), xs_1)
}
end_branch_4:
return __t4
}

func Call_tail__1935051898(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "tail")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}))
}

func Call_take__4024047148(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(n_0), gopurs_runtime.Int(0))).IntVal) != (0) {
__t2 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_nil()
})
goto end_branch_2
} else {

}
}
{
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 218341868 && v1_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 218341868 && v1_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_take(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}))
_ = __local_var_1_0
__t2 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
end_branch_2:
return __t2
}

func Call_takeWhile__638755635(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0).IntVal) != (0)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_takeWhile(p_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_transpose__1534541312(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))})
_ = v_1_0
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t9 = xs_0
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
v1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, "head")))})
_ = v1_2_1
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.UnsafePtr == nil) {
__t8 = Call_transpose(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, "tail"))
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.UnsafePtr != nil) {
__local_var_3_2 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.UnsafePtr).V0, "head")
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}.UnsafePtr).V0, "tail")
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply(Call_mapMaybe(Get_head()), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, "tail"))
_ = __local_var_5_4
__local_var_6_5 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_2, __local_var_5_4})}
}))
_ = __local_var_6_5
__local_var_7_7 := gopurs_runtime.Apply(Call_mapMaybe(Get_tail()), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, "tail"))
_ = __local_var_7_7
__local_var_7_6 := Call_transpose(gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_3, __local_var_7_7})}
})))
_ = __local_var_7_6
__t8 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_5, __local_var_7_6})}
}))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = __t8
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

func Call_uncons__3647012005(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1)})}
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

func Call_uncons__1321764894(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1)})}
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

func Call_uncons__974566859(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1)})}
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

func Call_uncons__1420258522(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1)})}
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

func Call_unionBy__3844889126(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), xs_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy(eq_0, a_4, b_3)
})
}), gopurs_runtime.Apply(Call_nubByEq(eq_0), ys_2), xs_1))
}

func Call_updateAt__725610501(n_0_loop int64, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t0 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (n_0) == (0) {
__t0 = &pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1}
goto end_branch_0
} else {

}
}
{
__t0 = &pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_updateAt(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
}
end_branch_0:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t0)}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}), xs_2)
}

func Call_zipWith__3539178005(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), Call_zipWith(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}), xs_1), ys_2)
}

func Call_zipWith__3210333397(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), Call_zipWith(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}), xs_1), ys_2)
}

func Call_zipWith__3984071349(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), Call_zipWith(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}), xs_1), ys_2)
}

func Call_isNothing__1401305026(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__2641488518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__3931678292(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__2349537221(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
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

func Call_greaterThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
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

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_sequence__1886310617(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__1779806517(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


