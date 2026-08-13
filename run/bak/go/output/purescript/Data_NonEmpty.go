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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
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
return Call_Data_NonEmpty_tail(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_NonEmpty_tail
}

var cache_Data_NonEmpty_singleton gopurs_runtime.Value
var once_Data_NonEmpty_singleton sync.Once
func Get_Data_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_NonEmpty_singleton.Do(func() {
		cache_Data_NonEmpty_singleton = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](dictPlus_0_box))
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
return Call_Data_NonEmpty_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_Data_NonEmpty_oneOf
}

var cache_Data_NonEmpty_head gopurs_runtime.Value
var once_Data_NonEmpty_head sync.Once
func Get_Data_NonEmpty_head() gopurs_runtime.Value {
	once_Data_NonEmpty_head.Do(func() {
		cache_Data_NonEmpty_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_head(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
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
return Call_Data_NonEmpty_fromNonEmpty(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
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
		cache_Data_NonEmpty_foldl1 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_foldl1(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
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
return Call_Data_NonEmpty_head__4279565926(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_NonEmpty_head__4279565926
}

var cache_Data_NonEmpty_singleton__3741573463 gopurs_runtime.Value
var once_Data_NonEmpty_singleton__3741573463 sync.Once
func Get_Data_NonEmpty_singleton__3741573463() gopurs_runtime.Value {
	once_Data_NonEmpty_singleton__3741573463.Do(func() {
		cache_Data_NonEmpty_singleton__3741573463 = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_NonEmpty_singleton__3741573463(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](dictPlus_0_box))
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

type Constructor_Data_NonEmpty_NonEmpty[T_f any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}


func Call_Data_NonEmpty_unfoldable1NonEmpty(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_1)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2)))}))
_ = __local_var_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr).V1})}
})
}))
}

func Call_Data_NonEmpty_tail(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_Data_NonEmpty_singleton(dictPlus_0_loop *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictPlus_0 *Constructor_Control_Plus_Plus[gopurs_runtime.Value] = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 -> gopurs_runtime.Value
empty_1_0 := gopurs_runtime.Box(dictPlus_0.V1)
_ = empty_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, empty_1_0})}
})
}

func Call_Data_NonEmpty_showNonEmpty(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(NonEmpty ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))
}

func Call_Data_NonEmpty_semigroupNonEmpty(dictApplicative_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))})}
})
}))
}

func Call_Data_NonEmpty_oneOf(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 -> *Constructor_Control_Alt_Alt[gopurs_runtime.Value]
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}

func Call_Data_NonEmpty_head(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_Data_NonEmpty_functorNonEmpty(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)})}
})
}))
}

func Call_Data_NonEmpty_functorWithIndex(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorNonEmpty1_1_0 -> gopurs_runtime.Value
functorNonEmpty1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_1_0
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_4})})
}), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
})
}))
}

func Call_Data_NonEmpty_fromNonEmpty(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}

func Call_Data_NonEmpty_foldableNonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})
}))
}

func Call_Data_NonEmpty_foldableWithIndexNonEmpty(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableNonEmpty1_1_0 -> gopurs_runtime.Value
foldableNonEmpty1_1_0 := Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableNonEmpty1_1_0
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_1_0
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_1 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_6})})
}), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_5})})
}), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, b_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_5})})
}), b_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
})
}))
}

func Call_Data_NonEmpty_traversableNonEmpty(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorNonEmpty1_1_0 -> gopurs_runtime.Value
functorNonEmpty1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_1_0
// TAST (Let): foldableNonEmpty1_2_2 -> gopurs_runtime.Value
foldableNonEmpty1_2_2 := Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableNonEmpty1_2_2
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_3 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_3
// TAST (Let): Functor0_5_4 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_4
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_3.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_4.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_5 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_5
// TAST (Let): Functor0_5_6 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_6
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_5.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_6.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
})
})
}))
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
// TAST (Let): functorNonEmpty1_2_2 -> gopurs_runtime.Value
functorNonEmpty1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_2_2
// TAST (Let): functorWithIndex1_1_0 -> gopurs_runtime.Value
functorWithIndex1_1_0 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_2_2
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_5})})
}), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
}))
_ = functorWithIndex1_1_0
// TAST (Let): foldableWithIndexNonEmpty1_2_4 -> gopurs_runtime.Value
foldableWithIndexNonEmpty1_2_4 := Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexNonEmpty1_2_4
// TAST (Let): traversableNonEmpty1_3_5 -> gopurs_runtime.Value
traversableNonEmpty1_3_5 := Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableNonEmpty1_3_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexNonEmpty1_2_4
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndex1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableNonEmpty1_3_5
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_6 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_6
// TAST (Let): Functor0_6_7 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_7
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_6.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_7.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply2(f_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_4))}, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_9})})
}), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
})
})
}))
}

func Call_Data_NonEmpty_foldable1NonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): foldableNonEmpty1_1_0 -> gopurs_runtime.Value
foldableNonEmpty1_1_0 := Call_Data_NonEmpty_foldableNonEmpty(dictFoldable_0)
_ = foldableNonEmpty1_1_0
return gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_1_0
}), gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), s_5, gopurs_runtime.Apply(f_3, a1_6))
})
}), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
__local_var_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(f_2, a1_5)
_ = __local_var_6_4
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr == nil) {
__t5 = a1_5
goto end_branch_5
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(__local_var_6_4, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
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
_ = __local_var_6_3
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_3, x_7)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
_ = __local_var_5_2
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.UnsafePtr == nil) {
__t6 = (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.UnsafePtr != nil) {
__t6 = gopurs_runtime.Apply(__local_var_4_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.UnsafePtr).V0)
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
}))
}

func Call_Data_NonEmpty_foldl1(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}), "foldl1")
}

func Call_Data_NonEmpty_eqNonEmpty(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
})
}))
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
// TAST (Let): eqNonEmpty2_3_1 -> gopurs_runtime.Value
eqNonEmpty2_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](__local_var_3_2))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal) != (0)))
})
}))
_ = eqNonEmpty2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqNonEmpty2_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_3 -> gopurs_runtime.Value
v_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_2, "compare"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
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
__t4 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_2))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
}))
})
}

func Call_Data_NonEmpty_eq1NonEmpty(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
})
})
}))
}

func Call_Data_NonEmpty_ord1NonEmpty(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): ordNonEmpty1_1_0 -> gopurs_runtime.Value
ordNonEmpty1_1_0 := Call_Data_NonEmpty_ordNonEmpty(dictOrd1_0)
_ = ordNonEmpty1_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): eq1NonEmpty1_2_1 -> gopurs_runtime.Value
eq1NonEmpty1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_3, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_3))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal) != (0)))
})
})
}))
_ = eq1NonEmpty1_2_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1NonEmpty1_2_1
}), gopurs_runtime.Func(func(dictOrd_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordNonEmpty1_1_0, dictOrd_3), "compare")
}))
}

func Call_Data_NonEmpty_head__4279565926(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_Data_NonEmpty_singleton__3741573463(dictPlus_0_loop *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictPlus_0 *Constructor_Control_Plus_Plus[gopurs_runtime.Value] = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 -> gopurs_runtime.Value
empty_1_0 := gopurs_runtime.Box(dictPlus_0.V1)
_ = empty_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, empty_1_0})}
})
}

func Call_Data_NonEmpty_singleton__532815287(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta0_0, gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_plusList(), "empty")})}
}


