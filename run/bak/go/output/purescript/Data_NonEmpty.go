package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_NonEmpty_NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_NonEmpty sync.Once
func Get_Data_NonEmpty_NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_NonEmpty.Do(func() {
		cache_Data_NonEmpty_NonEmpty = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, value0, value1})}
})
})
	})
	return cache_Data_NonEmpty_NonEmpty
}

var cache_Data_NonEmpty_unfoldable1NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_unfoldable1NonEmpty sync.Once
func Get_Data_NonEmpty_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_unfoldable1NonEmpty.Do(func() {
		cache_Data_NonEmpty_unfoldable1NonEmpty = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_unfoldable1NonEmpty(dictUnfoldable_0_box)
})
	})
	return cache_Data_NonEmpty_unfoldable1NonEmpty
}

var cache_Data_NonEmpty_tail gopurs_runtime.Value
var once_Data_NonEmpty_tail sync.Once
func Get_Data_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_NonEmpty_tail.Do(func() {
		cache_Data_NonEmpty_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_tail(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box))
})
	})
	return cache_Data_NonEmpty_tail
}

var cache_Data_NonEmpty_singleton gopurs_runtime.Value
var once_Data_NonEmpty_singleton sync.Once
func Get_Data_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_NonEmpty_singleton.Do(func() {
		cache_Data_NonEmpty_singleton = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](dictPlus_0_box))
})
	})
	return cache_Data_NonEmpty_singleton
}

var cache_Data_NonEmpty_showNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_showNonEmpty sync.Once
func Get_Data_NonEmpty_showNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_showNonEmpty.Do(func() {
		cache_Data_NonEmpty_showNonEmpty = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_showNonEmpty(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_NonEmpty_showNonEmpty
}

var cache_Data_NonEmpty_semigroupNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_semigroupNonEmpty sync.Once
func Get_Data_NonEmpty_semigroupNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_semigroupNonEmpty.Do(func() {
		cache_Data_NonEmpty_semigroupNonEmpty = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_semigroupNonEmpty(dictApplicative_0_box, dictSemigroup_1_box)
})
	})
	return cache_Data_NonEmpty_semigroupNonEmpty
}

var cache_Data_NonEmpty_oneOf gopurs_runtime.Value
var once_Data_NonEmpty_oneOf sync.Once
func Get_Data_NonEmpty_oneOf() gopurs_runtime.Value {
	once_Data_NonEmpty_oneOf.Do(func() {
		cache_Data_NonEmpty_oneOf = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_NonEmpty_oneOf
}

var cache_Data_NonEmpty_head gopurs_runtime.Value
var once_Data_NonEmpty_head sync.Once
func Get_Data_NonEmpty_head() gopurs_runtime.Value {
	once_Data_NonEmpty_head.Do(func() {
		cache_Data_NonEmpty_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_head(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box))
})
	})
	return cache_Data_NonEmpty_head
}

var cache_Data_NonEmpty_functorNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_functorNonEmpty sync.Once
func Get_Data_NonEmpty_functorNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_functorNonEmpty.Do(func() {
		cache_Data_NonEmpty_functorNonEmpty = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_functorNonEmpty(dictFunctor_0_box)
})
	})
	return cache_Data_NonEmpty_functorNonEmpty
}

var cache_Data_NonEmpty_functorWithIndex gopurs_runtime.Value
var once_Data_NonEmpty_functorWithIndex sync.Once
func Get_Data_NonEmpty_functorWithIndex() gopurs_runtime.Value {
	once_Data_NonEmpty_functorWithIndex.Do(func() {
		cache_Data_NonEmpty_functorWithIndex = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_functorWithIndex(dictFunctorWithIndex_0_box)
})
	})
	return cache_Data_NonEmpty_functorWithIndex
}

var cache_Data_NonEmpty_fromNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_fromNonEmpty sync.Once
func Get_Data_NonEmpty_fromNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_fromNonEmpty.Do(func() {
		cache_Data_NonEmpty_fromNonEmpty = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_fromNonEmpty(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box))
})
	})
	return cache_Data_NonEmpty_fromNonEmpty
}

var cache_Data_NonEmpty_foldableNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_foldableNonEmpty sync.Once
func Get_Data_NonEmpty_foldableNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_foldableNonEmpty.Do(func() {
		cache_Data_NonEmpty_foldableNonEmpty = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_foldableNonEmpty(dictFoldable_0_box)
})
	})
	return cache_Data_NonEmpty_foldableNonEmpty
}

var cache_Data_NonEmpty_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_foldableWithIndexNonEmpty sync.Once
func Get_Data_NonEmpty_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_NonEmpty_foldableWithIndexNonEmpty = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_foldableWithIndexNonEmpty(dictFoldableWithIndex_0_box)
})
	})
	return cache_Data_NonEmpty_foldableWithIndexNonEmpty
}

var cache_Data_NonEmpty_traversableNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_traversableNonEmpty sync.Once
func Get_Data_NonEmpty_traversableNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_traversableNonEmpty.Do(func() {
		cache_Data_NonEmpty_traversableNonEmpty = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_traversableNonEmpty(dictTraversable_0_box)
})
	})
	return cache_Data_NonEmpty_traversableNonEmpty
}

var cache_Data_NonEmpty_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_traversableWithIndexNonEmpty sync.Once
func Get_Data_NonEmpty_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_NonEmpty_traversableWithIndexNonEmpty = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_traversableWithIndexNonEmpty(dictTraversableWithIndex_0_box)
})
	})
	return cache_Data_NonEmpty_traversableWithIndexNonEmpty
}

var cache_Data_NonEmpty_foldable1NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_foldable1NonEmpty sync.Once
func Get_Data_NonEmpty_foldable1NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_foldable1NonEmpty.Do(func() {
		cache_Data_NonEmpty_foldable1NonEmpty = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_foldable1NonEmpty(dictFoldable_0_box)
})
	})
	return cache_Data_NonEmpty_foldable1NonEmpty
}

var cache_Data_NonEmpty_foldl1 gopurs_runtime.Value
var once_Data_NonEmpty_foldl1 sync.Once
func Get_Data_NonEmpty_foldl1() gopurs_runtime.Value {
	once_Data_NonEmpty_foldl1.Do(func() {
		cache_Data_NonEmpty_foldl1 = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_foldl1(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box))
})
	})
	return cache_Data_NonEmpty_foldl1
}

var cache_Data_NonEmpty_eqNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_eqNonEmpty sync.Once
func Get_Data_NonEmpty_eqNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_eqNonEmpty.Do(func() {
		cache_Data_NonEmpty_eqNonEmpty = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_eqNonEmpty(dictEq1_0_box, dictEq_1_box)
})
	})
	return cache_Data_NonEmpty_eqNonEmpty
}

var cache_Data_NonEmpty_ordNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_ordNonEmpty sync.Once
func Get_Data_NonEmpty_ordNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_ordNonEmpty.Do(func() {
		cache_Data_NonEmpty_ordNonEmpty = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_ordNonEmpty(dictOrd1_0_box)
})
	})
	return cache_Data_NonEmpty_ordNonEmpty
}

var cache_Data_NonEmpty_eq1NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_eq1NonEmpty sync.Once
func Get_Data_NonEmpty_eq1NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_eq1NonEmpty.Do(func() {
		cache_Data_NonEmpty_eq1NonEmpty = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_eq1NonEmpty(dictEq1_0_box)
})
	})
	return cache_Data_NonEmpty_eq1NonEmpty
}

var cache_Data_NonEmpty_ord1NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_ord1NonEmpty sync.Once
func Get_Data_NonEmpty_ord1NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_ord1NonEmpty.Do(func() {
		cache_Data_NonEmpty_ord1NonEmpty = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_ord1NonEmpty(dictOrd1_0_box)
})
	})
	return cache_Data_NonEmpty_ord1NonEmpty
}

var cache_Data_NonEmpty_head__4279565926 gopurs_runtime.Value
var once_Data_NonEmpty_head__4279565926 sync.Once
func Get_Data_NonEmpty_head__4279565926() gopurs_runtime.Value {
	once_Data_NonEmpty_head__4279565926.Do(func() {
		cache_Data_NonEmpty_head__4279565926 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_head__4279565926(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box))
})
	})
	return cache_Data_NonEmpty_head__4279565926
}

var cache_Data_NonEmpty_singleton__3741573463 gopurs_runtime.Value
var once_Data_NonEmpty_singleton__3741573463 sync.Once
func Get_Data_NonEmpty_singleton__3741573463() gopurs_runtime.Value {
	once_Data_NonEmpty_singleton__3741573463.Do(func() {
		cache_Data_NonEmpty_singleton__3741573463 = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_singleton__3741573463(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](dictPlus_0_box))
})
	})
	return cache_Data_NonEmpty_singleton__3741573463
}

var cache_Data_NonEmpty_singleton__532815287 gopurs_runtime.Value
var once_Data_NonEmpty_singleton__532815287 sync.Once
func Get_Data_NonEmpty_singleton__532815287() gopurs_runtime.Value {
	once_Data_NonEmpty_singleton__532815287.Do(func() {
		cache_Data_NonEmpty_singleton__532815287 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_singleton__532815287(__eta0_0_box)
})
	})
	return cache_Data_NonEmpty_singleton__532815287
}

type Constructor_Data_NonEmpty_NonEmpty struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_Data_NonEmpty_unfoldable1NonEmpty(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> *Constructor_Data_Tuple_Tuple
__local_var_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_1, b_2))
_ = __local_var_3_1
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Tuple_Tuple
var __local_var_3_0 *Constructor_Data_Tuple_Tuple = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (__local_var_3_1).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(v1_4.UnsafePtr).V0)}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), (__local_var_3_1).V1)})})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (__local_var_3_0).V0, (__local_var_3_0).V1})}
})
})})}
}

func Call_Data_NonEmpty_tail(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_NonEmpty_singleton(dictPlus_0_loop *Constructor_Control_Plus_Plus) gopurs_runtime.Value {
var dictPlus_0 *Constructor_Control_Plus_Plus = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 -> gopurs_runtime.Value
empty_1_0 := gopurs_runtime.Box(dictPlus_0.V1)
_ = empty_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_2, empty_1_0})}
})
}

func Call_Data_NonEmpty_showNonEmpty(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(NonEmpty ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1).StrVal())) + (")"))
})})}
}

func Call_Data_NonEmpty_semigroupNonEmpty(dictApplicative_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), (*Constructor_Data_NonEmpty_NonEmpty)(v1_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v1_3.UnsafePtr).V1))})}
})
})})}
}

func Call_Data_NonEmpty_oneOf(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 -> *Constructor_Control_Alt_Alt
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
}

func Call_Data_NonEmpty_head(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_NonEmpty_functorNonEmpty(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(m_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Data_NonEmpty_NonEmpty)(m_2.UnsafePtr).V1)})}
})
})})}
}

func Call_Data_NonEmpty_functorWithIndex(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorNonEmpty1_1_0 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(m_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_Data_NonEmpty_NonEmpty)(m_3.UnsafePtr).V1)})}
})
})))
_ = functorNonEmpty1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
}), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)})}
})
})})}
}

func Call_Data_NonEmpty_fromNonEmpty(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (v_1).V0, (v_1).V1)
}

func Call_Data_NonEmpty_foldableNonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
})})}
}

func Call_Data_NonEmpty_foldableWithIndexNonEmpty(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): foldableNonEmpty1_1_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_2.V0), gopurs_runtime.Apply(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldl"), f_2, gopurs_runtime.Apply2(f_2, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldr"), f_2, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
})))
_ = foldableNonEmpty1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_0)}
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_3 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_3
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_3.V0), gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_6})})
}), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
}), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
}), b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
})})}
}

func Call_Data_NonEmpty_traversableNonEmpty(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorNonEmpty1_1_0 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(m_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_Data_NonEmpty_NonEmpty)(m_3.UnsafePtr).V1)})}
})
})))
_ = functorNonEmpty1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): foldableNonEmpty1_2_2 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_4 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_4.V0), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), f_3, gopurs_runtime.Apply2(f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
})))
_ = foldableNonEmpty1_2_2
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_2_2)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_1_0)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_5 -> *Constructor_Control_Apply_Apply
Apply0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_5
// TAST (Let): Functor0_5_6 -> *Constructor_Data_Functor_Functor
Functor0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_6
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_5.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_6.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_7 -> *Constructor_Control_Apply_Apply
Apply0_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_7
// TAST (Let): Functor0_5_8 -> *Constructor_Data_Functor_Functor
Functor0_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_8
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_7.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_8.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
})
})
})})}
}

func Call_Data_NonEmpty_traversableWithIndexNonEmpty(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorNonEmpty1_2_2 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(m_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, (*Constructor_Data_NonEmpty_NonEmpty)(m_4.UnsafePtr).V1)})}
})
})))
_ = functorNonEmpty1_2_2
// TAST (Let): functorWithIndex1_1_0 -> *Constructor_Data_FunctorWithIndex_FunctorWithIndex
functorWithIndex1_1_0 := &Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_2_2)}
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
}), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)})}
})
})}
_ = functorWithIndex1_1_0
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{})
_ = __local_var_2_5
// TAST (Let): __local_var_3_7 -> gopurs_runtime.Value
__local_var_3_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_5, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_3_7
// TAST (Let): foldableNonEmpty1_3_6 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_3_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_8 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_8
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_8.V0), gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_7, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_4))}, f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_7, "foldl"), f_4, gopurs_runtime.Apply2(f_4, b_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_7, "foldr"), f_4, b_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
})))
_ = foldableNonEmpty1_3_6
// TAST (Let): foldableWithIndexNonEmpty1_2_4 -> *Constructor_Data_FoldableWithIndex_FoldableWithIndex
foldableWithIndexNonEmpty1_2_4 := &Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_3_6)}
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_9 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_9
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_9.V0), gopurs_runtime.Apply2(f_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_5, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_4))}, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_8})})
}), (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_5, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_7})})
}), gopurs_runtime.Apply3(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, b_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_5, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_7})})
}), b_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
})}
_ = foldableWithIndexNonEmpty1_2_4
// TAST (Let): __local_var_3_11 -> gopurs_runtime.Value
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 -> gopurs_runtime.Value
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): functorNonEmpty1_4_12 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(m_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "map"), f_5, (*Constructor_Data_NonEmpty_NonEmpty)(m_6.UnsafePtr).V1)})}
})
})))
_ = functorNonEmpty1_4_12
// TAST (Let): __local_var_5_15 -> gopurs_runtime.Value
__local_var_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_5_15
// TAST (Let): foldableNonEmpty1_5_14 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_5_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_16 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_16
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_16.V0), gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_15, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_6))}, f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_15, "foldl"), f_6, gopurs_runtime.Apply2(f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_15, "foldr"), f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
})))
_ = foldableNonEmpty1_5_14
// TAST (Let): traversableNonEmpty1_3_10 -> *Constructor_Data_Traversable_Traversable
traversableNonEmpty1_3_10 := &Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_5_14)}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_4_12)}
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_17 -> *Constructor_Control_Apply_Apply
Apply0_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_17
// TAST (Let): Functor0_8_18 -> *Constructor_Data_Functor_Functor
Functor0_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_18
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_17.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_18.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_19 -> *Constructor_Control_Apply_Apply
Apply0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_19
// TAST (Let): Functor0_8_20 -> *Constructor_Data_Functor_Functor
Functor0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_20
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_20.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_9, (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_11, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, f_9, (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V1))
})
})
})}
_ = traversableNonEmpty1_3_10
return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(&Constructor_Data_TraversableWithIndex_TraversableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(foldableWithIndexNonEmpty1_2_4)}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(functorWithIndex1_1_0)}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(traversableNonEmpty1_3_10)}
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_21 -> *Constructor_Control_Apply_Apply
Apply0_5_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_21
// TAST (Let): Functor0_6_22 -> *Constructor_Data_Functor_Functor
Functor0_6_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_22
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_21.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_22.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply2(f_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_4))}, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
}), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
})})}
}

func Call_Data_NonEmpty_foldable1NonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): foldableNonEmpty1_1_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_0)}
}), gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), s_5, gopurs_runtime.Apply(f_3, a1_6))
})
}), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> *Constructor_Data_Maybe_Just
__local_var_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(f_2, a1_5)
_ = __local_var_6_5
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr == nil) {
__t6 = a1_5
goto end_branch_6
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr != nil) {
__t6 = gopurs_runtime.Apply(__local_var_6_5, (*Constructor_Data_Maybe_Just)(v2_7.UnsafePtr).V0)
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
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_6_4, x_7)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
_ = __local_var_5_3
var __t7 gopurs_runtime.Value
{
if (__local_var_5_3 == nil) {
__t7 = (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0
goto end_branch_7
} else {

}
}
{
if (__local_var_5_3 != nil) {
__t7 = gopurs_runtime.Apply(__local_var_4_2, (__local_var_5_3).V0)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
})})}
}

func Call_Data_NonEmpty_foldl1(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V1), f_1, (v_2).V0, (v_2).V1)
}

func Call_Data_NonEmpty_eqNonEmpty(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_1))}, (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1).IntVal) != (0)))
})
})})}
}

func Call_Data_NonEmpty_ordNonEmpty(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): eqNonEmpty2_3_1 -> *Constructor_Data_Eq_Eq
eqNonEmpty2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_4.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](__local_var_3_2))}, (*Constructor_Data_NonEmpty_NonEmpty)(x_4.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty)(y_5.UnsafePtr).V1).IntVal) != (0)))
})
})))
_ = eqNonEmpty2_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty2_3_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_3 -> gopurs_runtime.Value
v_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_2, "compare"), (*Constructor_Data_NonEmpty_NonEmpty)(x_4.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_5.UnsafePtr).V0)
_ = v_6_3
var __t4 uint32
{
if (uint32(v_6_3.IntVal) == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (uint32(v_6_3.IntVal) == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
__t4 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_2))}, (*Constructor_Data_NonEmpty_NonEmpty)(x_4.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty)(y_5.UnsafePtr).V1).IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
})})}
})
}

func Call_Data_NonEmpty_eq1NonEmpty(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_1))}, (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1).IntVal) != (0)))
})
})
})})}
}

func Call_Data_NonEmpty_ord1NonEmpty(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eq1NonEmpty1_1_0 -> *Constructor_Data_Eq_Eq1
eq1NonEmpty1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_2, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_2))}, (*Constructor_Data_NonEmpty_NonEmpty)(x_3.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty)(y_4.UnsafePtr).V1).IntVal) != (0)))
})
})
})))
_ = eq1NonEmpty1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1NonEmpty1_1_0)}
}), gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_2 -> gopurs_runtime.Value
v_5_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_2, "compare"), (*Constructor_Data_NonEmpty_NonEmpty)(x_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_4.UnsafePtr).V0)
_ = v_5_2
var __t3 uint32
{
if (uint32(v_5_2.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_5_2.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_2))}, (*Constructor_Data_NonEmpty_NonEmpty)(x_3.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty)(y_4.UnsafePtr).V1).IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
})
})})}
}

func Call_Data_NonEmpty_head__4279565926(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_NonEmpty_singleton__3741573463(dictPlus_0_loop *Constructor_Control_Plus_Plus) gopurs_runtime.Value {
var dictPlus_0 *Constructor_Control_Plus_Plus = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 -> gopurs_runtime.Value
empty_1_0 := gopurs_runtime.Box(dictPlus_0.V1)
_ = empty_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_2, empty_1_0})}
})
}

func Call_Data_NonEmpty_singleton__532815287(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __eta0_0, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_List_Lazy_Types_plusList()).V1)})}
}


