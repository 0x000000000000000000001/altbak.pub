package Data_List_Lazy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	unsafe "unsafe"
)

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_lessThanOrEq
}

var cache_tailRecM2 gopurs_runtime.Value
var once_tailRecM2 sync.Once
func Get_tailRecM2() gopurs_runtime.Value {
	once_tailRecM2.Do(func() {
		cache_tailRecM2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_tailRecM2
}

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_lessThan
}

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_greaterThan
}

var cache_greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		cache_greaterThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_greaterThanOrEq
}

var cache_any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		cache_any = func() gopurs_runtime.Value {
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
	return cache_any
}

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

var cache_zipWith__gopurs_runtime_Value_887591047 gopurs_runtime.Value
var once_zipWith__gopurs_runtime_Value_887591047 sync.Once
func Get_zipWith__gopurs_runtime_Value_887591047() gopurs_runtime.Value {
	once_zipWith__gopurs_runtime_Value_887591047.Do(func() {
		cache_zipWith__gopurs_runtime_Value_887591047 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith__gopurs_runtime_Value_887591047(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith__gopurs_runtime_Value_887591047
}

var cache_zipWith__gopurs_runtime_Value_177791335 gopurs_runtime.Value
var once_zipWith__gopurs_runtime_Value_177791335 sync.Once
func Get_zipWith__gopurs_runtime_Value_177791335() gopurs_runtime.Value {
	once_zipWith__gopurs_runtime_Value_177791335.Do(func() {
		cache_zipWith__gopurs_runtime_Value_177791335 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith__gopurs_runtime_Value_177791335(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith__gopurs_runtime_Value_177791335
}

var cache_zipWith__gopurs_runtime_Value_2990889615 gopurs_runtime.Value
var once_zipWith__gopurs_runtime_Value_2990889615 sync.Once
func Get_zipWith__gopurs_runtime_Value_2990889615() gopurs_runtime.Value {
	once_zipWith__gopurs_runtime_Value_2990889615.Do(func() {
		cache_zipWith__gopurs_runtime_Value_2990889615 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith__gopurs_runtime_Value_2990889615(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith__gopurs_runtime_Value_2990889615
}

var cache_zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		cache_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA(dictApplicative_0_box)
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

var cache_updateAt__gopurs_runtime_Value_3172091723 gopurs_runtime.Value
var once_updateAt__gopurs_runtime_Value_3172091723 sync.Once
func Get_updateAt__gopurs_runtime_Value_3172091723() gopurs_runtime.Value {
	once_updateAt__gopurs_runtime_Value_3172091723.Do(func() {
		cache_updateAt__gopurs_runtime_Value_3172091723 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAt__gopurs_runtime_Value_3172091723(n_0_box.IntVal, x_1_box, xs_2_box)
})
	})
	return cache_updateAt__gopurs_runtime_Value_3172091723
}

var cache_updateAt__gopurs_runtime_Value_911144478 gopurs_runtime.Value
var once_updateAt__gopurs_runtime_Value_911144478 sync.Once
func Get_updateAt__gopurs_runtime_Value_911144478() gopurs_runtime.Value {
	once_updateAt__gopurs_runtime_Value_911144478.Do(func() {
		cache_updateAt__gopurs_runtime_Value_911144478 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAt__gopurs_runtime_Value_911144478(n_0_box.IntVal, x_1_box, xs_2_box)
})
	})
	return cache_updateAt__gopurs_runtime_Value_911144478
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_1_0, __local_var_4_2})}.UnsafePtr))}
})), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_1, __local_var_5_3})}.UnsafePtr))}
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

var cache_uncons__gopurs_runtime_Value_2155426095 gopurs_runtime.Value
var once_uncons__gopurs_runtime_Value_2155426095 sync.Once
func Get_uncons__gopurs_runtime_Value_2155426095() gopurs_runtime.Value {
	once_uncons__gopurs_runtime_Value_2155426095.Do(func() {
		cache_uncons__gopurs_runtime_Value_2155426095 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__gopurs_runtime_Value_2155426095(xs_0_box))}
})
	})
	return cache_uncons__gopurs_runtime_Value_2155426095
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable(dictUnfoldable_0_box)
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

var cache_takeWhile__gopurs_runtime_Value_1526402856 gopurs_runtime.Value
var once_takeWhile__gopurs_runtime_Value_1526402856 sync.Once
func Get_takeWhile__gopurs_runtime_Value_1526402856() gopurs_runtime.Value {
	once_takeWhile__gopurs_runtime_Value_1526402856.Do(func() {
		cache_takeWhile__gopurs_runtime_Value_1526402856 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile__gopurs_runtime_Value_1526402856(p_0_box)
})
	})
	return cache_takeWhile__gopurs_runtime_Value_1526402856
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

var cache_take__gopurs_runtime_Value_3690802007 gopurs_runtime.Value
var once_take__gopurs_runtime_Value_3690802007 sync.Once
func Get_take__gopurs_runtime_Value_3690802007() gopurs_runtime.Value {
	once_take__gopurs_runtime_Value_3690802007.Do(func() {
		cache_take__gopurs_runtime_Value_3690802007 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take__gopurs_runtime_Value_3690802007(n_0_box.IntVal)
})
	})
	return cache_take__gopurs_runtime_Value_3690802007
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

var cache_tail__gopurs_runtime_Value_1935051898 gopurs_runtime.Value
var once_tail__gopurs_runtime_Value_1935051898 sync.Once
func Get_tail__gopurs_runtime_Value_1935051898() gopurs_runtime.Value {
	once_tail__gopurs_runtime_Value_1935051898.Do(func() {
		cache_tail__gopurs_runtime_Value_1935051898 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_tail__gopurs_runtime_Value_1935051898(xs_0_box))}
})
	})
	return cache_tail__gopurs_runtime_Value_1935051898
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_stripPrefix(dictEq_0_box, v_1_box, s_2_box))}
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

var cache_span__gopurs_runtime_Value_1778192253 gopurs_runtime.Value
var once_span__gopurs_runtime_Value_1778192253 sync.Once
func Get_span__gopurs_runtime_Value_1778192253() gopurs_runtime.Value {
	once_span__gopurs_runtime_Value_1778192253.Do(func() {
		cache_span__gopurs_runtime_Value_1778192253 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span__gopurs_runtime_Value_1778192253(p_0_box, xs_1_box)
})
	})
	return cache_span__gopurs_runtime_Value_1778192253
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

var cache_scanlLazy__gopurs_runtime_Value_4138780922 gopurs_runtime.Value
var once_scanlLazy__gopurs_runtime_Value_4138780922 sync.Once
func Get_scanlLazy__gopurs_runtime_Value_4138780922() gopurs_runtime.Value {
	once_scanlLazy__gopurs_runtime_Value_4138780922.Do(func() {
		cache_scanlLazy__gopurs_runtime_Value_4138780922 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanlLazy__gopurs_runtime_Value_4138780922(f_0_box, acc_1_box, xs_2_box)
})
	})
	return cache_scanlLazy__gopurs_runtime_Value_4138780922
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

var cache_reverse__gopurs_runtime_Value_1535498235 gopurs_runtime.Value
var once_reverse__gopurs_runtime_Value_1535498235 sync.Once
func Get_reverse__gopurs_runtime_Value_1535498235() gopurs_runtime.Value {
	once_reverse__gopurs_runtime_Value_1535498235.Do(func() {
		cache_reverse__gopurs_runtime_Value_1535498235 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reverse__gopurs_runtime_Value_1535498235(xs_0_box)
})
	})
	return cache_reverse__gopurs_runtime_Value_1535498235
}

var cache_replicateM gopurs_runtime.Value
var once_replicateM sync.Once
func Get_replicateM() gopurs_runtime.Value {
	once_replicateM.Do(func() {
		cache_replicateM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicateM(dictMonad_0_box)
})
	})
	return cache_replicateM
}

var cache_replicateM__gopurs_runtime_Value_2313876346 gopurs_runtime.Value
var once_replicateM__gopurs_runtime_Value_2313876346 sync.Once
func Get_replicateM__gopurs_runtime_Value_2313876346() gopurs_runtime.Value {
	once_replicateM__gopurs_runtime_Value_2313876346.Do(func() {
		cache_replicateM__gopurs_runtime_Value_2313876346 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicateM__gopurs_runtime_Value_2313876346(dictMonad_0_box)
})
	})
	return cache_replicateM__gopurs_runtime_Value_2313876346
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

var cache_repeat__gopurs_runtime_Value_2462085934 gopurs_runtime.Value
var once_repeat__gopurs_runtime_Value_2462085934 sync.Once
func Get_repeat__gopurs_runtime_Value_2462085934() gopurs_runtime.Value {
	once_repeat__gopurs_runtime_Value_2462085934.Do(func() {
		cache_repeat__gopurs_runtime_Value_2462085934 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat__gopurs_runtime_Value_2462085934(x_0_box)
})
	})
	return cache_repeat__gopurs_runtime_Value_2462085934
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

var cache_null__gopurs_runtime_Value_3517949020 gopurs_runtime.Value
var once_null__gopurs_runtime_Value_3517949020 sync.Once
func Get_null__gopurs_runtime_Value_3517949020() gopurs_runtime.Value {
	once_null__gopurs_runtime_Value_3517949020.Do(func() {
		cache_null__gopurs_runtime_Value_3517949020 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null__gopurs_runtime_Value_3517949020(x_0_box))
})
	})
	return cache_null__gopurs_runtime_Value_3517949020
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

var cache_nubBy__gopurs_runtime_Value_3654717659 gopurs_runtime.Value
var once_nubBy__gopurs_runtime_Value_3654717659 sync.Once
func Get_nubBy__gopurs_runtime_Value_3654717659() gopurs_runtime.Value {
	once_nubBy__gopurs_runtime_Value_3654717659.Do(func() {
		cache_nubBy__gopurs_runtime_Value_3654717659 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubBy__gopurs_runtime_Value_3654717659(p_0_box)
})
	})
	return cache_nubBy__gopurs_runtime_Value_3654717659
}

var cache_nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		cache_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nub(dictOrd_0_box)
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

var cache_mapMaybe__gopurs_runtime_Value_1173728558 gopurs_runtime.Value
var once_mapMaybe__gopurs_runtime_Value_1173728558 sync.Once
func Get_mapMaybe__gopurs_runtime_Value_1173728558() gopurs_runtime.Value {
	once_mapMaybe__gopurs_runtime_Value_1173728558.Do(func() {
		cache_mapMaybe__gopurs_runtime_Value_1173728558 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__gopurs_runtime_Value_1173728558(f_0_box)
})
	})
	return cache_mapMaybe__gopurs_runtime_Value_1173728558
}

var cache_mapMaybe__gopurs_runtime_Value_3066361755 gopurs_runtime.Value
var once_mapMaybe__gopurs_runtime_Value_3066361755 sync.Once
func Get_mapMaybe__gopurs_runtime_Value_3066361755() gopurs_runtime.Value {
	once_mapMaybe__gopurs_runtime_Value_3066361755.Do(func() {
		cache_mapMaybe__gopurs_runtime_Value_3066361755 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__gopurs_runtime_Value_3066361755(f_0_box)
})
	})
	return cache_mapMaybe__gopurs_runtime_Value_3066361755
}

var cache_some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		cache_some = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some(dictAlternative_0_box, dictLazy_1_box, v_2_box)
})
	})
	return cache_some
}

var cache_some__gopurs_runtime_Value_2240457450 gopurs_runtime.Value
var once_some__gopurs_runtime_Value_2240457450 sync.Once
func Get_some__gopurs_runtime_Value_2240457450() gopurs_runtime.Value {
	once_some__gopurs_runtime_Value_2240457450.Do(func() {
		cache_some__gopurs_runtime_Value_2240457450 = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some__gopurs_runtime_Value_2240457450(dictAlternative_0_box, dictLazy_1_box, v_2_box)
})
	})
	return cache_some__gopurs_runtime_Value_2240457450
}

var cache_many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		cache_many = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(dictAlternative_0_box, dictLazy_1_box, v_2_box)
})
	})
	return cache_many
}

var cache_many__gopurs_runtime_Value_2240457450 gopurs_runtime.Value
var once_many__gopurs_runtime_Value_2240457450 sync.Once
func Get_many__gopurs_runtime_Value_2240457450() gopurs_runtime.Value {
	once_many__gopurs_runtime_Value_2240457450.Do(func() {
		cache_many__gopurs_runtime_Value_2240457450 = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many__gopurs_runtime_Value_2240457450(dictAlternative_0_box, dictLazy_1_box, v_2_box)
})
	})
	return cache_many__gopurs_runtime_Value_2240457450
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((l_0.IntVal) + (1))
})
}), gopurs_runtime.Int(0))
	})
	return cache_length
}

var cache_length__gopurs_runtime_Value_2692518155 gopurs_runtime.Value
var once_length__gopurs_runtime_Value_2692518155 sync.Once
func Get_length__gopurs_runtime_Value_2692518155() gopurs_runtime.Value {
	once_length__gopurs_runtime_Value_2692518155.Do(func() {
		cache_length__gopurs_runtime_Value_2692518155 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((l_0.IntVal) + (1))
})
}), gopurs_runtime.Int(0))
	})
	return cache_length__gopurs_runtime_Value_2692518155
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = func() gopurs_runtime.Value {
var go__go_0_0_10 gopurs_runtime.Value
go__go_0_0_10 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
go__go_0_0_10:
for {
if false { continue go__go_0_0_10 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
__local_var_2_3 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))}
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Bool(false)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
v_1_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
continue go__go_0_0_10
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_10, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return cache_last
}

var cache_last__gopurs_runtime_Value_2155426095 gopurs_runtime.Value
var once_last__gopurs_runtime_Value_2155426095 sync.Once
func Get_last__gopurs_runtime_Value_2155426095() gopurs_runtime.Value {
	once_last__gopurs_runtime_Value_2155426095.Do(func() {
		cache_last__gopurs_runtime_Value_2155426095 = func() gopurs_runtime.Value {
var go__go_0_0_11 gopurs_runtime.Value
go__go_0_0_11 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
go__go_0_0_11:
for {
if false { continue go__go_0_0_11 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
__local_var_2_3 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))}
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Bool(false)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
v_1_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
continue go__go_0_0_11
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_11, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return cache_last__gopurs_runtime_Value_2155426095
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

var cache_iterate__gopurs_runtime_Value_4112245263 gopurs_runtime.Value
var once_iterate__gopurs_runtime_Value_4112245263 sync.Once
func Get_iterate__gopurs_runtime_Value_4112245263() gopurs_runtime.Value {
	once_iterate__gopurs_runtime_Value_4112245263.Do(func() {
		cache_iterate__gopurs_runtime_Value_4112245263 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate__gopurs_runtime_Value_4112245263(f_0_box, x_1_box)
})
	})
	return cache_iterate__gopurs_runtime_Value_4112245263
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

var cache_insertAt__gopurs_runtime_Value_3172091723 gopurs_runtime.Value
var once_insertAt__gopurs_runtime_Value_3172091723 sync.Once
func Get_insertAt__gopurs_runtime_Value_3172091723() gopurs_runtime.Value {
	once_insertAt__gopurs_runtime_Value_3172091723.Do(func() {
		cache_insertAt__gopurs_runtime_Value_3172091723 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAt__gopurs_runtime_Value_3172091723(v_0_box.IntVal, v1_1_box, v2_2_box)
})
	})
	return cache_insertAt__gopurs_runtime_Value_3172091723
}

var cache_insertAt__gopurs_runtime_Value_911144478 gopurs_runtime.Value
var once_insertAt__gopurs_runtime_Value_911144478 sync.Once
func Get_insertAt__gopurs_runtime_Value_911144478() gopurs_runtime.Value {
	once_insertAt__gopurs_runtime_Value_911144478.Do(func() {
		cache_insertAt__gopurs_runtime_Value_911144478 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAt__gopurs_runtime_Value_911144478(v_0_box.IntVal, v1_1_box, v2_2_box)
})
	})
	return cache_insertAt__gopurs_runtime_Value_911144478
}

var cache_init gopurs_runtime.Value
var once_init sync.Once
func Get_init() gopurs_runtime.Value {
	once_init.Do(func() {
		cache_init = func() gopurs_runtime.Value {
var go__go_0_0_14 gopurs_runtime.Value
_ = go__go_0_0_14
go__go_0_0_14 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
__local_var_2_3 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))}
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Bool(false)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_cons(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0), gopurs_runtime.Apply(go__go_0_0_14, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))).UnsafePtr))}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr))}
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_14, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return cache_init
}

var cache_init__gopurs_runtime_Value_1935051898 gopurs_runtime.Value
var once_init__gopurs_runtime_Value_1935051898 sync.Once
func Get_init__gopurs_runtime_Value_1935051898() gopurs_runtime.Value {
	once_init__gopurs_runtime_Value_1935051898.Do(func() {
		cache_init__gopurs_runtime_Value_1935051898 = func() gopurs_runtime.Value {
var go__go_0_0_15 gopurs_runtime.Value
_ = go__go_0_0_15
go__go_0_0_15 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
__local_var_2_3 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))}
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Bool(false)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_cons(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0), gopurs_runtime.Apply(go__go_0_0_15, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))).UnsafePtr))}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr))}
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_15, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return cache_init__gopurs_runtime_Value_1935051898
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

var cache_head__gopurs_runtime_Value_2155426095 gopurs_runtime.Value
var once_head__gopurs_runtime_Value_2155426095 sync.Once
func Get_head__gopurs_runtime_Value_2155426095() gopurs_runtime.Value {
	once_head__gopurs_runtime_Value_2155426095.Do(func() {
		cache_head__gopurs_runtime_Value_2155426095 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_head__gopurs_runtime_Value_2155426095(xs_0_box))}
})
	})
	return cache_head__gopurs_runtime_Value_2155426095
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

var cache_transpose__gopurs_runtime_Value_1519875707 gopurs_runtime.Value
var once_transpose__gopurs_runtime_Value_1519875707 sync.Once
func Get_transpose__gopurs_runtime_Value_1519875707() gopurs_runtime.Value {
	once_transpose__gopurs_runtime_Value_1519875707.Do(func() {
		cache_transpose__gopurs_runtime_Value_1519875707 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_transpose__gopurs_runtime_Value_1519875707(xs_0_box)
})
	})
	return cache_transpose__gopurs_runtime_Value_1519875707
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

var cache_groupBy__gopurs_runtime_Value_2430346981 gopurs_runtime.Value
var once_groupBy__gopurs_runtime_Value_2430346981 sync.Once
func Get_groupBy__gopurs_runtime_Value_2430346981() gopurs_runtime.Value {
	once_groupBy__gopurs_runtime_Value_2430346981.Do(func() {
		cache_groupBy__gopurs_runtime_Value_2430346981 = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupBy__gopurs_runtime_Value_2430346981(eq_0_box)
})
	})
	return cache_groupBy__gopurs_runtime_Value_2430346981
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group(dictEq_0_box)
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

var cache_insertBy__gopurs_runtime_Value_3848889042 gopurs_runtime.Value
var once_insertBy__gopurs_runtime_Value_3848889042 sync.Once
func Get_insertBy__gopurs_runtime_Value_3848889042() gopurs_runtime.Value {
	once_insertBy__gopurs_runtime_Value_3848889042.Do(func() {
		cache_insertBy__gopurs_runtime_Value_3848889042 = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertBy__gopurs_runtime_Value_3848889042(cmp_0_box, x_1_box, xs_2_box)
})
	})
	return cache_insertBy__gopurs_runtime_Value_3848889042
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert(dictOrd_0_box)
})
	})
	return cache_insert
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(dictFoldable_0_box)
})
	})
	return cache_fromFoldable
}

var cache_fromFoldable__gopurs_runtime_Value_3298799452 gopurs_runtime.Value
var once_fromFoldable__gopurs_runtime_Value_3298799452 sync.Once
func Get_fromFoldable__gopurs_runtime_Value_3298799452() gopurs_runtime.Value {
	once_fromFoldable__gopurs_runtime_Value_3298799452.Do(func() {
		cache_fromFoldable__gopurs_runtime_Value_3298799452 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable__gopurs_runtime_Value_3298799452(dictFoldable_0_box)
})
	})
	return cache_fromFoldable__gopurs_runtime_Value_3298799452
}

var cache_foldrLazy gopurs_runtime.Value
var once_foldrLazy sync.Once
func Get_foldrLazy() gopurs_runtime.Value {
	once_foldrLazy.Do(func() {
		cache_foldrLazy = gopurs_runtime.Func3(func(dictLazy_0_box gopurs_runtime.Value, op_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrLazy(dictLazy_0_box, op_1_box, z_2_box)
})
	})
	return cache_foldrLazy
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictMonad_0_box, f_1_box, b_2_box, xs_3_box)
})
	})
	return cache_foldM
}

var cache_foldM__gopurs_runtime_Value_3630384258 gopurs_runtime.Value
var once_foldM__gopurs_runtime_Value_3630384258 sync.Once
func Get_foldM__gopurs_runtime_Value_3630384258() gopurs_runtime.Value {
	once_foldM__gopurs_runtime_Value_3630384258.Do(func() {
		cache_foldM__gopurs_runtime_Value_3630384258 = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM__gopurs_runtime_Value_3630384258(dictMonad_0_box, f_1_box, b_2_box, xs_3_box)
})
	})
	return cache_foldM__gopurs_runtime_Value_3630384258
}

var cache_foldM__gopurs_runtime_Value_2269678914 gopurs_runtime.Value
var once_foldM__gopurs_runtime_Value_2269678914 sync.Once
func Get_foldM__gopurs_runtime_Value_2269678914() gopurs_runtime.Value {
	once_foldM__gopurs_runtime_Value_2269678914.Do(func() {
		cache_foldM__gopurs_runtime_Value_2269678914 = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM__gopurs_runtime_Value_2269678914(dictMonad_0_box, f_1_box, b_2_box, xs_3_box)
})
	})
	return cache_foldM__gopurs_runtime_Value_2269678914
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

var cache_findIndex__gopurs_runtime_Value_2097162489 gopurs_runtime.Value
var once_findIndex__gopurs_runtime_Value_2097162489 sync.Once
func Get_findIndex__gopurs_runtime_Value_2097162489() gopurs_runtime.Value {
	once_findIndex__gopurs_runtime_Value_2097162489.Do(func() {
		cache_findIndex__gopurs_runtime_Value_2097162489 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findIndex__gopurs_runtime_Value_2097162489(fn_0_box)
})
	})
	return cache_findIndex__gopurs_runtime_Value_2097162489
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

var cache_findLastIndex__gopurs_runtime_Value_2097162489 gopurs_runtime.Value
var once_findLastIndex__gopurs_runtime_Value_2097162489 sync.Once
func Get_findLastIndex__gopurs_runtime_Value_2097162489() gopurs_runtime.Value {
	once_findLastIndex__gopurs_runtime_Value_2097162489.Do(func() {
		cache_findLastIndex__gopurs_runtime_Value_2097162489 = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex__gopurs_runtime_Value_2097162489(fn_0_box, xs_1_box))}
})
	})
	return cache_findLastIndex__gopurs_runtime_Value_2097162489
}

var cache_filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		cache_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterM(dictMonad_0_box)
})
	})
	return cache_filterM
}

var cache_filterM__gopurs_runtime_Value_698007938 gopurs_runtime.Value
var once_filterM__gopurs_runtime_Value_698007938 sync.Once
func Get_filterM__gopurs_runtime_Value_698007938() gopurs_runtime.Value {
	once_filterM__gopurs_runtime_Value_698007938.Do(func() {
		cache_filterM__gopurs_runtime_Value_698007938 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterM__gopurs_runtime_Value_698007938(dictMonad_0_box)
})
	})
	return cache_filterM__gopurs_runtime_Value_698007938
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

var cache_filter__gopurs_runtime_Value_1526402856 gopurs_runtime.Value
var once_filter__gopurs_runtime_Value_1526402856 sync.Once
func Get_filter__gopurs_runtime_Value_1526402856() gopurs_runtime.Value {
	once_filter__gopurs_runtime_Value_1526402856.Do(func() {
		cache_filter__gopurs_runtime_Value_1526402856 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter__gopurs_runtime_Value_1526402856(p_0_box)
})
	})
	return cache_filter__gopurs_runtime_Value_1526402856
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

var cache_intersectBy__gopurs_runtime_Value_2139240221 gopurs_runtime.Value
var once_intersectBy__gopurs_runtime_Value_2139240221 sync.Once
func Get_intersectBy__gopurs_runtime_Value_2139240221() gopurs_runtime.Value {
	once_intersectBy__gopurs_runtime_Value_2139240221.Do(func() {
		cache_intersectBy__gopurs_runtime_Value_2139240221 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectBy__gopurs_runtime_Value_2139240221(eq_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_intersectBy__gopurs_runtime_Value_2139240221
}

var cache_intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		cache_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersect(dictEq_0_box)
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

var cache_nubByEq__gopurs_runtime_Value_1127729793 gopurs_runtime.Value
var once_nubByEq__gopurs_runtime_Value_1127729793 sync.Once
func Get_nubByEq__gopurs_runtime_Value_1127729793() gopurs_runtime.Value {
	once_nubByEq__gopurs_runtime_Value_1127729793.Do(func() {
		cache_nubByEq__gopurs_runtime_Value_1127729793 = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubByEq__gopurs_runtime_Value_1127729793(eq_0_box)
})
	})
	return cache_nubByEq__gopurs_runtime_Value_1127729793
}

var cache_nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		cache_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubEq(dictEq_0_box)
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
return Call_elemLastIndex(dictEq_0_box, x_1_box)
})
	})
	return cache_elemLastIndex
}

var cache_elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		cache_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemIndex(dictEq_0_box, x_1_box)
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

var cache_drop__gopurs_runtime_Value_3690802007 gopurs_runtime.Value
var once_drop__gopurs_runtime_Value_3690802007 sync.Once
func Get_drop__gopurs_runtime_Value_3690802007() gopurs_runtime.Value {
	once_drop__gopurs_runtime_Value_3690802007.Do(func() {
		cache_drop__gopurs_runtime_Value_3690802007 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop__gopurs_runtime_Value_3690802007(n_0_box.IntVal)
})
	})
	return cache_drop__gopurs_runtime_Value_3690802007
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

var cache_deleteBy__gopurs_runtime_Value_3550949736 gopurs_runtime.Value
var once_deleteBy__gopurs_runtime_Value_3550949736 sync.Once
func Get_deleteBy__gopurs_runtime_Value_3550949736() gopurs_runtime.Value {
	once_deleteBy__gopurs_runtime_Value_3550949736.Do(func() {
		cache_deleteBy__gopurs_runtime_Value_3550949736 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy__gopurs_runtime_Value_3550949736(eq_0_box, x_1_box, xs_2_box)
})
	})
	return cache_deleteBy__gopurs_runtime_Value_3550949736
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

var cache_unionBy__gopurs_runtime_Value_2139240221 gopurs_runtime.Value
var once_unionBy__gopurs_runtime_Value_2139240221 sync.Once
func Get_unionBy__gopurs_runtime_Value_2139240221() gopurs_runtime.Value {
	once_unionBy__gopurs_runtime_Value_2139240221.Do(func() {
		cache_unionBy__gopurs_runtime_Value_2139240221 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionBy__gopurs_runtime_Value_2139240221(eq_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_unionBy__gopurs_runtime_Value_2139240221
}

var cache_union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		cache_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_union(dictEq_0_box)
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

var cache_deleteAt__gopurs_runtime_Value_3690802007 gopurs_runtime.Value
var once_deleteAt__gopurs_runtime_Value_3690802007 sync.Once
func Get_deleteAt__gopurs_runtime_Value_3690802007() gopurs_runtime.Value {
	once_deleteAt__gopurs_runtime_Value_3690802007.Do(func() {
		cache_deleteAt__gopurs_runtime_Value_3690802007 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteAt__gopurs_runtime_Value_3690802007(n_0_box.IntVal, xs_1_box)
})
	})
	return cache_deleteAt__gopurs_runtime_Value_3690802007
}

var cache_delete gopurs_runtime.Value
var once_delete sync.Once
func Get_delete() gopurs_runtime.Value {
	once_delete.Do(func() {
		cache_delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete(dictEq_0_box)
})
	})
	return cache_delete
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference(dictEq_0_box)
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

var cache_alterAt__gopurs_runtime_Value_1764238455 gopurs_runtime.Value
var once_alterAt__gopurs_runtime_Value_1764238455 sync.Once
func Get_alterAt__gopurs_runtime_Value_1764238455() gopurs_runtime.Value {
	once_alterAt__gopurs_runtime_Value_1764238455.Do(func() {
		cache_alterAt__gopurs_runtime_Value_1764238455 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alterAt__gopurs_runtime_Value_1764238455(n_0_box.IntVal, f_1_box, xs_2_box)
})
	})
	return cache_alterAt__gopurs_runtime_Value_1764238455
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

func Call_tailRecM2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), Call_zipWith(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
})
}), xs_1), ys_2)
}

func Call_zipWith__gopurs_runtime_Value_887591047(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), Call_zipWith(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
})
}), xs_1), ys_2)
}

func Call_zipWith__gopurs_runtime_Value_177791335(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), Call_zipWith(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
})
}), xs_1), ys_2)
}

func Call_zipWith__gopurs_runtime_Value_2990889615(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), Call_zipWith(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
})
}), xs_1), ys_2)
}

func Call_zipWithA(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
sequence1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "sequence"), dictApplicative_0)
_ = sequence1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_1_0, Call_zipWith(f_2, xs_3, ys_4))
})
})
})
}

func Call_updateAt(n_0_loop int64, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (n_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_updateAt((n_0) - (1), x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_updateAt__gopurs_runtime_Value_3172091723(n_0_loop int64, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (n_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_updateAt((n_0) - (1), x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_updateAt__gopurs_runtime_Value_911144478(n_0_loop int64, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (n_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_updateAt((n_0) - (1), x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_uncons(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0)
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 218341868 && v_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 218341868 && v_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1_0.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr)
}

func Call_uncons__gopurs_runtime_Value_2155426095(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0)
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 218341868 && v_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 218341868 && v_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1_0.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr)
}

func Call_toUnfoldable(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.RecordGet(rec_2, "tail")})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_1))}).UnsafePtr))}
}))
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0).IntVal) != (0)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_takeWhile(p_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_takeWhile__gopurs_runtime_Value_1526402856(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0).IntVal) != (0)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_takeWhile(p_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_take(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(0)).IntVal) != (0) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 218341868 && v1_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_take((n_0) - (1)), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
_ = __local_var_1_0
__t2 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
end_branch_2:
return __t2
}

func Call_take__gopurs_runtime_Value_3690802007(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(0)).IntVal) != (0) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 218341868 && v1_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_take((n_0) - (1)), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
_ = __local_var_1_0
__t2 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
end_branch_2:
return __t2
}

func Call_tail(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "tail")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}).UnsafePtr)
}

func Call_tail__gopurs_runtime_Value_1935051898(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "tail")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}).UnsafePtr)
}

func Call_stripPrefix(dictEq_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(Call_tailRecM2(gopurs_runtime.Func(func(prefix_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(input_4 gopurs_runtime.Value) gopurs_runtime.Value {
v1_5_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), prefix_3)
_ = v1_5_0
var __t1 gopurs_runtime.Value
{
if (v1_5_0.Type == 9 && v1_5_0.IntVal == 218341868 && v1_5_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, input_4})}})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v1_5_0.Type == 9 && v1_5_0.IntVal == 218341868 && v1_5_0.UnsafePtr != nil) {
v2_6_2 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), input_4)
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if ((v2_6_2.Type == 9 && v2_6_2.IntVal == 218341868 && v2_6_2.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5_0.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v2_6_2.UnsafePtr).V0).IntVal) != (0)) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5_0.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v2_6_2.UnsafePtr).V1)})}})}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_3:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t3.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr))}
})
}), v_1, s_2).UnsafePtr)
}

func Call_span(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
v_2_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_1))}
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if ((v_2_0.Type == 9 && v_2_0.IntVal == 930809136 && v_2_0.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "head")).IntVal) != (0)) {
__local_var_3_2 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "head")
_ = __local_var_3_2
v1_4_3 := Call_span(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "tail"))
_ = v1_4_3
__local_var_5_4 := gopurs_runtime.RecordGet(v1_4_3, "init")
_ = __local_var_5_4
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_2, __local_var_5_4})}.UnsafePtr))}
})), gopurs_runtime.RecordGet(v1_4_3, "rest"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", pkg_Data_List_Lazy_Types.Get_nil(), xs_1)
}
end_branch_1:
return __t1
}

func Call_span__gopurs_runtime_Value_1778192253(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
v_2_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_1))}
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if ((v_2_0.Type == 9 && v_2_0.IntVal == 930809136 && v_2_0.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "head")).IntVal) != (0)) {
__local_var_3_2 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "head")
_ = __local_var_3_2
v1_4_3 := Call_span(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "tail"))
_ = v1_4_3
__local_var_5_4 := gopurs_runtime.RecordGet(v1_4_3, "init")
_ = __local_var_5_4
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_2, __local_var_5_4})}.UnsafePtr))}
})), gopurs_runtime.RecordGet(v1_4_3, "rest"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", pkg_Data_List_Lazy_Types.Get_nil(), xs_1)
}
end_branch_1:
return __t1
}

func Call_snoc(xs_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
})), xs_0)
}

func Call_singleton(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
}))
}

func Call_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Pattern "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_showList(), dictShow_0), "show"), v_1), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_scanlLazy(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
acc_prime_4_1 := gopurs_runtime.Apply2(f_0, acc_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
_ = acc_prime_4_1
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, acc_prime_4_1, Call_scanlLazy(f_0, acc_prime_4_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_scanlLazy__gopurs_runtime_Value_4138780922(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
acc_prime_4_1 := gopurs_runtime.Apply2(f_0, acc_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
_ = acc_prime_4_1
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, acc_prime_4_1, Call_scanlLazy(f_0, acc_prime_4_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_reverse(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_3, b_2})}.UnsafePtr))}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_0)
}))
}

func Call_reverse__gopurs_runtime_Value_1535498235(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_3, b_2})}.UnsafePtr))}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_0)
}))
}

func Call_replicateM(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), n_3, gopurs_runtime.Int(1)).IntVal) != (0) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(Call_replicateM(dictMonad_0), gopurs_runtime.Int((n_3.IntVal) - (1)), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_5, as_6})}.UnsafePtr))}
})))
}))
}))
}
end_branch_2:
return __t2
})
})
}

func Call_replicateM__gopurs_runtime_Value_2313876346(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), n_3, gopurs_runtime.Int(1)).IntVal) != (0) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(Call_replicateM(dictMonad_0), gopurs_runtime.Int((n_3.IntVal) - (1)), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_5, as_6})}.UnsafePtr))}
})))
}))
}))
}
end_branch_2:
return __t2
})
})
}

func Call_repeat(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
go__go_1_0_0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_0})}.UnsafePtr))}
}))
}))
return go__go_1_0_0
}

func Call_repeat__gopurs_runtime_Value_2462085934(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_1 gopurs_runtime.Value
_ = go__go_1_0_1
go__go_1_0_1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_1})}.UnsafePtr))}
}))
}))
return go__go_1_0_1
}

func Call_replicate(i_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var go__go_2_0_2 gopurs_runtime.Value
_ = go__go_2_0_2
go__go_2_0_2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, xs_1, go__go_2_0_2})}.UnsafePtr))}
}))
}))
return gopurs_runtime.Apply(Call_take(i_0), go__go_2_0_2)
}

func Call_go__range(start_0_loop int64, end_1_loop int64) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1)).IntVal) != (0) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_unfoldableList(), "unfoldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThanOrEq(), x_2, gopurs_runtime.Int(end_1)).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2, gopurs_runtime.Int((x_2.IntVal) - (1))})}})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t2.UnsafePtr))}
}), gopurs_runtime.Int(start_0))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_unfoldableList(), "unfoldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), x_2, gopurs_runtime.Int(end_1)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2, gopurs_runtime.Int((x_2.IntVal) + (1))})}})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), gopurs_runtime.Int(start_0))
}
end_branch_1:
return __t1
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.RecordGet(v_2, "yes")})}.UnsafePtr))}
})))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.RecordGet(v_2, "no")})}.UnsafePtr))}
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

func Call_null__gopurs_runtime_Value_3517949020(x_0_loop gopurs_runtime.Value) bool {
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
var goStep_1_0_3 gopurs_runtime.Value
_ = goStep_1_0_3
var go__go_1_1_4 gopurs_runtime.Value
_ = go__go_1_1_4
goStep_1_0_3 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2)
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v2_4_3, "found").IntVal) != (0) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(go__go_1_1_4, gopurs_runtime.RecordGet(v2_4_3, "result"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)).UnsafePtr))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_1_1_4, gopurs_runtime.RecordGet(v2_4_3, "result"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_4:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t4.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t2.UnsafePtr))}
})
})
go__go_1_1_4 = gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(goStep_1_0_3, s_2), v_3)
})
})
return gopurs_runtime.Apply(go__go_1_1_4, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: nil})
}

func Call_nubBy__gopurs_runtime_Value_3654717659(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var goStep_1_0_5 gopurs_runtime.Value
_ = goStep_1_0_5
var go__go_1_1_6 gopurs_runtime.Value
_ = go__go_1_1_6
goStep_1_0_5 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2)
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v2_4_3, "found").IntVal) != (0) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(go__go_1_1_6, gopurs_runtime.RecordGet(v2_4_3, "result"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)).UnsafePtr))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_1_1_6, gopurs_runtime.RecordGet(v2_4_3, "result"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_4:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t4.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t2.UnsafePtr))}
})
})
go__go_1_1_6 = gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(goStep_1_0_5, s_2), v_3)
})
})
return gopurs_runtime.Apply(go__go_1_1_6, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: nil})
}

func Call_nub(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return Call_nubBy(gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}

func Call_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_7 gopurs_runtime.Value
go__go_1_0_7 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_7:
for {
if false { continue go__go_1_0_7 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
v1_3_2 := gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
_ = v1_3_2
var __t3 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 930809136 && v1_3_2.UnsafePtr == nil) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_7
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 930809136 && v1_3_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_mapMaybe(f_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t3.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_7)
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, x_3)
})
}

func Call_mapMaybe__gopurs_runtime_Value_1173728558(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_8 gopurs_runtime.Value
go__go_1_0_8 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_8:
for {
if false { continue go__go_1_0_8 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
v1_3_2 := gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
_ = v1_3_2
var __t3 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 930809136 && v1_3_2.UnsafePtr == nil) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_8
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 930809136 && v1_3_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_mapMaybe(f_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t3.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_8)
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, x_3)
})
}

func Call_mapMaybe__gopurs_runtime_Value_3066361755(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_9 gopurs_runtime.Value
go__go_1_0_9 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_9:
for {
if false { continue go__go_1_0_9 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
v1_3_2 := gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
_ = v1_3_2
var __t3 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 930809136 && v1_3_2.UnsafePtr == nil) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_9
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 930809136 && v1_3_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_mapMaybe(f_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t3.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_9)
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, x_3)
})
}

func Call_some(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Lazy_Types.Get_cons(), v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_1, "defer"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(dictAlternative_0, dictLazy_1, v_2)
})))
}

func Call_some__gopurs_runtime_Value_2240457450(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Lazy_Types.Get_cons(), v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_1, "defer"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(dictAlternative_0, dictLazy_1, v_2)
})))
}

func Call_many(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), Call_some(dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
}

func Call_many__gopurs_runtime_Value_2240457450(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), Call_some(dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
}

func Call_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_12 gopurs_runtime.Value
_ = go__go_2_0_12
go__go_2_0_12 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, go__go_2_0_12)
_ = __local_var_4_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, __local_var_4_1})}.UnsafePtr))}
}))
}))
return go__go_2_0_12
}

func Call_iterate__gopurs_runtime_Value_4112245263(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_13 gopurs_runtime.Value
_ = go__go_2_0_13
go__go_2_0_13 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, go__go_2_0_13)
_ = __local_var_4_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, __local_var_4_1})}.UnsafePtr))}
}))
}))
return go__go_2_0_13
}

func Call_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})}.UnsafePtr))}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 218341868 && v3_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V0, Call_insertAt((v_0) - (1), v1_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), v2_2)
}
end_branch_1:
return __t1
}

func Call_insertAt__gopurs_runtime_Value_3172091723(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})}.UnsafePtr))}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 218341868 && v3_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V0, Call_insertAt((v_0) - (1), v1_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), v2_2)
}
end_branch_1:
return __t1
}

func Call_insertAt__gopurs_runtime_Value_911144478(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})}.UnsafePtr))}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 218341868 && v3_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V0, Call_insertAt((v_0) - (1), v1_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v3_3.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), v2_2)
}
end_branch_1:
return __t1
}

func Call_index(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_16 gopurs_runtime.Value
go__go_1_0_16 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_16:
for {
if false { continue go__go_1_0_16 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (v1_3.IntVal) == (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
v1_3_loop = gopurs_runtime.Int((v1_3.IntVal) - (1))
continue go__go_1_0_16
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_16, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
}

func Call_head(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "head")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}).UnsafePtr)
}

func Call_head__gopurs_runtime_Value_2155426095(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "head")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}).UnsafePtr)
}

func Call_transpose(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr == nil) {
__t1 = xs_0
goto end_branch_1
} else {

}
}
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr != nil) {
v1_2_2 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, "head")))}
_ = v1_2_2
var __t3 gopurs_runtime.Value
{
if (v1_2_2.Type == 9 && v1_2_2.IntVal == 930809136 && v1_2_2.UnsafePtr == nil) {
__t3 = Call_transpose(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, "tail"))
goto end_branch_3
} else {

}
}
{
if (v1_2_2.Type == 9 && v1_2_2.IntVal == 930809136 && v1_2_2.UnsafePtr != nil) {
__local_var_3_4 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_2_2.UnsafePtr).V0, "head")
_ = __local_var_3_4
__local_var_4_5 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_2_2.UnsafePtr).V0, "tail")
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply(Call_mapMaybe(Get_head()), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, "tail"))
_ = __local_var_5_6
__local_var_6_7 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_4, __local_var_5_6})}.UnsafePtr))}
}))
_ = __local_var_6_7
__local_var_7_9 := gopurs_runtime.Apply(Call_mapMaybe(Get_tail()), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, "tail"))
_ = __local_var_7_9
__local_var_7_8 := Call_transpose(gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_5, __local_var_7_9})}.UnsafePtr))}
})))
_ = __local_var_7_8
__t3 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_7, __local_var_7_8})}.UnsafePtr))}
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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

func Call_transpose__gopurs_runtime_Value_1519875707(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_0))}
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr == nil) {
__t1 = xs_0
goto end_branch_1
} else {

}
}
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr != nil) {
v1_2_2 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, "head")))}
_ = v1_2_2
var __t3 gopurs_runtime.Value
{
if (v1_2_2.Type == 9 && v1_2_2.IntVal == 930809136 && v1_2_2.UnsafePtr == nil) {
__t3 = Call_transpose(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, "tail"))
goto end_branch_3
} else {

}
}
{
if (v1_2_2.Type == 9 && v1_2_2.IntVal == 930809136 && v1_2_2.UnsafePtr != nil) {
__local_var_3_4 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_2_2.UnsafePtr).V0, "head")
_ = __local_var_3_4
__local_var_4_5 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_2_2.UnsafePtr).V0, "tail")
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply(Call_mapMaybe(Get_head()), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, "tail"))
_ = __local_var_5_6
__local_var_6_7 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_4, __local_var_5_6})}.UnsafePtr))}
}))
_ = __local_var_6_7
__local_var_7_9 := gopurs_runtime.Apply(Call_mapMaybe(Get_tail()), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0, "tail"))
_ = __local_var_7_9
__local_var_7_8 := Call_transpose(gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_5, __local_var_7_9})}.UnsafePtr))}
})))
_ = __local_var_7_8
__t3 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_7, __local_var_7_8})}.UnsafePtr))}
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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

func Call_groupBy(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
__local_var_2_2 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_2
v1_3_3 := Call_span(gopurs_runtime.Apply(eq_0, __local_var_2_2), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
_ = v1_3_3
__local_var_4_4 := gopurs_runtime.RecordGet(v1_3_3, "init")
_ = __local_var_4_4
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_2, __local_var_4_4})}
})), gopurs_runtime.Apply(Call_groupBy(eq_0), gopurs_runtime.RecordGet(v1_3_3, "rest"))})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_groupBy__gopurs_runtime_Value_2430346981(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
__local_var_2_2 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_2
v1_3_3 := Call_span(gopurs_runtime.Apply(eq_0, __local_var_2_2), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
_ = v1_3_3
__local_var_4_4 := gopurs_runtime.RecordGet(v1_3_3, "init")
_ = __local_var_4_4
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_2, __local_var_4_4})}
})), gopurs_runtime.Apply(Call_groupBy(eq_0), gopurs_runtime.RecordGet(v1_3_3, "rest"))})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_group(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return Call_groupBy(gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(cmp_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, Call_insertBy(cmp_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), v_3)})}.UnsafePtr))}
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_insertBy__gopurs_runtime_Value_3848889042(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, pkg_Data_List_Lazy_Types.Get_nil()})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(cmp_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, Call_insertBy(cmp_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), v_3)})}.UnsafePtr))}
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_insert(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_insertBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}

func Call_fromFoldable(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
}

func Call_fromFoldable__gopurs_runtime_Value_3298799452(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
}

func Call_foldrLazy(dictLazy_0_loop gopurs_runtime.Value, op_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
var op_1 gopurs_runtime.Value = op_1_loop
_ = op_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
var go__go_3_0_17 gopurs_runtime.Value
_ = go__go_3_0_17
go__go_3_0_17 = gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_4)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 218341868 && v_5_1.UnsafePtr != nil) {
__local_var_6_3 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_5_1.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_5_1.UnsafePtr).V1
_ = __local_var_7_4
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_0, "defer"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_1, __local_var_6_3, gopurs_runtime.Apply(go__go_3_0_17, __local_var_7_4))
}))
goto end_branch_2
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 218341868 && v_5_1.UnsafePtr == nil) {
__t2 = z_2
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
return go__go_3_0_17
}

func Call_foldM(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_3))}
_ = v_4_0
var __t1 gopurs_runtime.Value
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 930809136 && v_4_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_2)
goto end_branch_1
} else {

}
}
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 930809136 && v_4_0.UnsafePtr != nil) {
__local_var_5_2 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_0.UnsafePtr).V0, "tail")
_ = __local_var_5_2
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_0.UnsafePtr).V0, "head")), gopurs_runtime.Func(func(b_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictMonad_0, f_1, b_prime_6, __local_var_5_2)
}))
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

func Call_foldM__gopurs_runtime_Value_3630384258(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_3))}
_ = v_4_0
var __t1 gopurs_runtime.Value
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 930809136 && v_4_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_2)
goto end_branch_1
} else {

}
}
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 930809136 && v_4_0.UnsafePtr != nil) {
__local_var_5_2 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_0.UnsafePtr).V0, "tail")
_ = __local_var_5_2
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_0.UnsafePtr).V0, "head")), gopurs_runtime.Func(func(b_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictMonad_0, f_1, b_prime_6, __local_var_5_2)
}))
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

func Call_foldM__gopurs_runtime_Value_2269678914(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(xs_3))}
_ = v_4_0
var __t1 gopurs_runtime.Value
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 930809136 && v_4_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_2)
goto end_branch_1
} else {

}
}
{
if (v_4_0.Type == 9 && v_4_0.IntVal == 930809136 && v_4_0.UnsafePtr != nil) {
__local_var_5_2 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_0.UnsafePtr).V0, "tail")
_ = __local_var_5_2
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_0.UnsafePtr).V0, "head")), gopurs_runtime.Func(func(b_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictMonad_0, f_1, b_prime_6, __local_var_5_2)
}))
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

func Call_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_18 gopurs_runtime.Value
_ = go__go_1_0_18
go__go_1_0_18 = gopurs_runtime.Func(func(n_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(list_3))}, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet(o_4, "head")).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure"), n_2).UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(go__go_1_0_18, gopurs_runtime.Int((n_2.IntVal) + (1)), gopurs_runtime.RecordGet(o_4, "tail")).UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(__t1.UnsafePtr))}
})).UnsafePtr))}
})
})
return gopurs_runtime.Apply(go__go_1_0_18, gopurs_runtime.Int(0))
}

func Call_findIndex__gopurs_runtime_Value_2097162489(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_19 gopurs_runtime.Value
_ = go__go_1_0_19
go__go_1_0_19 = gopurs_runtime.Func(func(n_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(list_3))}, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet(o_4, "head")).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure"), n_2).UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(go__go_1_0_19, gopurs_runtime.Int((n_2.IntVal) + (1)), gopurs_runtime.RecordGet(o_4, "tail")).UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(__t1.UnsafePtr))}
})).UnsafePtr))}
})
})
return gopurs_runtime.Apply(go__go_1_0_19, gopurs_runtime.Int(0))
}

func Call_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(((gopurs_runtime.Apply(Get_length(), xs_1).IntVal) - (1)) - (v_2.IntVal))
}), gopurs_runtime.Apply(Call_findIndex(fn_0), Call_reverse(xs_1))).UnsafePtr)
}

func Call_findLastIndex__gopurs_runtime_Value_2097162489(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(((gopurs_runtime.Apply(Get_length(), xs_1).IntVal) - (1)) - (v_2.IntVal))
}), gopurs_runtime.Apply(Call_findIndex(fn_0), Call_reverse(xs_1))).UnsafePtr)
}

func Call_filterM(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_2 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(list_4))}
_ = v_5_2
var __t3 gopurs_runtime.Value
{
if (v_5_2.Type == 9 && v_5_2.IntVal == 930809136 && v_5_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_3
} else {

}
}
{
if (v_5_2.Type == 9 && v_5_2.IntVal == 930809136 && v_5_2.UnsafePtr != nil) {
__local_var_6_4 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_2.UnsafePtr).V0, "head")
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_2.UnsafePtr).V0, "tail")
_ = __local_var_7_5
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply(p_3, __local_var_6_4), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(Call_filterM(dictMonad_0), p_3, __local_var_7_5), gopurs_runtime.Func(func(xs_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (b_8.IntVal) != (0) {
__t6 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_4, xs_prime_9})}.UnsafePtr))}
}))
goto end_branch_6
} else {

}
}
{
__t6 = xs_prime_9
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), __t6)
}))
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
}

func Call_filterM__gopurs_runtime_Value_698007938(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_2 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(list_4))}
_ = v_5_2
var __t3 gopurs_runtime.Value
{
if (v_5_2.Type == 9 && v_5_2.IntVal == 930809136 && v_5_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_3
} else {

}
}
{
if (v_5_2.Type == 9 && v_5_2.IntVal == 930809136 && v_5_2.UnsafePtr != nil) {
__local_var_6_4 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_2.UnsafePtr).V0, "head")
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_2.UnsafePtr).V0, "tail")
_ = __local_var_7_5
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply(p_3, __local_var_6_4), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(Call_filterM(dictMonad_0), p_3, __local_var_7_5), gopurs_runtime.Func(func(xs_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (b_8.IntVal) != (0) {
__t6 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_4, xs_prime_9})}.UnsafePtr))}
}))
goto end_branch_6
} else {

}
}
{
__t6 = xs_prime_9
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), __t6)
}))
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
}

func Call_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_20 gopurs_runtime.Value
go__go_1_0_20 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_20:
for {
if false { continue go__go_1_0_20 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_filter(p_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_20
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_20)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, x_3)
})
}

func Call_filter__gopurs_runtime_Value_1526402856(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_21 gopurs_runtime.Value
go__go_1_0_21 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_21:
for {
if false { continue go__go_1_0_21 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_filter(p_0), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_21
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), go__go_1_0_21)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, x_3)
})
}

func Call_intersectBy(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_any(), gopurs_runtime.Apply(eq_0, x_3), ys_2).IntVal) != (0))
})), xs_1)
}

func Call_intersectBy__gopurs_runtime_Value_2139240221(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_any(), gopurs_runtime.Apply(eq_0, x_3), ys_2).IntVal) != (0))
})), xs_1)
}

func Call_intersect(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_nubByEq(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
__local_var_2_2 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_2
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_2, gopurs_runtime.Apply(Call_nubByEq(eq_0), gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply2(eq_0, __local_var_2_2, y_3)).IntVal) != (0))
})), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_nubByEq__gopurs_runtime_Value_1127729793(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
__local_var_2_2 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_2
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_2, gopurs_runtime.Apply(Call_nubByEq(eq_0), gopurs_runtime.Apply(Call_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply2(eq_0, __local_var_2_2, y_3)).IntVal) != (0))
})), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_nubEq(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return Call_nubByEq(gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_eq1List(), "eq1"), dictEq_0, x_1, y_2).IntVal) != (0))
})
}))
}

func Call_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
eqPattern1_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_eq1List(), "eq1"), __local_var_1_1, x_2, y_3).IntVal) != (0))
})
}))
_ = eqPattern1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqPattern1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_ordList(), dictOrd_0), "compare"), x_2, y_3)
})
}))
}

func Call_elemLastIndex(dictEq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1).IntVal) != (0))
}))
}

func Call_elemIndex(dictEq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_findIndex(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1).IntVal) != (0))
}))
}

func Call_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_22 gopurs_runtime.Value
go__go_1_0_22 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_22:
for {
if false { continue go__go_1_0_22 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if ((v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)
continue go__go_1_0_22
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
return gopurs_runtime.Apply(go__go_1_0_22, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_2))
})
}

func Call_drop(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var go__go_1_0_23 gopurs_runtime.Value
go__go_1_0_23 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3_loop_val.UnsafePtr)
go__go_1_0_23:
for {
if false { continue go__go_1_0_23 }
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Int((v_2.IntVal) - (1))
v1_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1).UnsafePtr)
continue go__go_1_0_23
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
})
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(go__go_1_0_23, gopurs_runtime.Int(n_0)))
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, x_3)
})
}

func Call_drop__gopurs_runtime_Value_3690802007(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var go__go_1_0_24 gopurs_runtime.Value
go__go_1_0_24 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3_loop_val.UnsafePtr)
go__go_1_0_24:
for {
if false { continue go__go_1_0_24 }
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Int((v_2.IntVal) - (1))
v1_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1).UnsafePtr)
continue go__go_1_0_24
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr))}
}
}()
})
})
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(go__go_1_0_24, gopurs_runtime.Int(n_0)))
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
return gopurs_runtime.Apply(Call_take((end_1) - (start_0)), gopurs_runtime.Apply(Call_drop(start_0), xs_2))
}

func Call_deleteBy(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(eq_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1).UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, Call_deleteBy(eq_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_deleteBy__gopurs_runtime_Value_3550949736(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(eq_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1).UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, Call_deleteBy(eq_0, x_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
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

func Call_unionBy__gopurs_runtime_Value_2139240221(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_union(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_deleteAt(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 218341868 && v1_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 218341868 && v1_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (n_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1).UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, Call_deleteAt((n_0) - (1), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_1)
}

func Call_deleteAt__gopurs_runtime_Value_3690802007(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 218341868 && v1_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 218341868 && v1_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (n_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1).UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, Call_deleteAt((n_0) - (1), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_1)
}

func Call_delete(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_difference(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy(gopurs_runtime.RecordGet(dictEq_0, "eq"), a_2, b_1)
})
}))
}

func Call_cycle(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_25 gopurs_runtime.Value
_ = go__go_1_0_25
go__go_1_0_25 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), xs_0, go__go_1_0_25)
}))
return go__go_1_0_25
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
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (n_0) == (0) {
v2_4_2 := gopurs_runtime.Apply(f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 930809136 && v2_4_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1).UnsafePtr))}
goto end_branch_3
} else {

}
}
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 930809136 && v2_4_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_4_2.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1})}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t3.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_alterAt((n_0) - (1), f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
}

func Call_alterAt__gopurs_runtime_Value_1764238455(n_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (n_0) == (0) {
v2_4_2 := gopurs_runtime.Apply(f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 930809136 && v2_4_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1).UnsafePtr))}
goto end_branch_3
} else {

}
}
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 930809136 && v2_4_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_4_2.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1})}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t3.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, Call_alterAt((n_0) - (1), f_1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), xs_2)
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


