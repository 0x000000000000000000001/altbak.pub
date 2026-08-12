package Data_Semigroup_Foldable

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
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
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity1(x_0_box)
})
	})
	return cache_identity1
}

var cache_FoldRight1 gopurs_runtime.Value
var once_FoldRight1 sync.Once
func Get_FoldRight1() gopurs_runtime.Value {
	once_FoldRight1.Do(func() {
		cache_FoldRight1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_FoldRight1[gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_FoldRight1
}

var cache_mkFoldRight1 gopurs_runtime.Value
var once_mkFoldRight1 sync.Once
func Get_mkFoldRight1() gopurs_runtime.Value {
	once_mkFoldRight1.Do(func() {
		cache_mkFoldRight1 = gopurs_runtime.Apply(Get_FoldRight1(), pkg_Data_Function.Get_go__const())
	})
	return cache_mkFoldRight1
}

var cache_mkFoldRight1__gopurs_runtime_Value_364767315 gopurs_runtime.Value
var once_mkFoldRight1__gopurs_runtime_Value_364767315 sync.Once
func Get_mkFoldRight1__gopurs_runtime_Value_364767315() gopurs_runtime.Value {
	once_mkFoldRight1__gopurs_runtime_Value_364767315.Do(func() {
		cache_mkFoldRight1__gopurs_runtime_Value_364767315 = gopurs_runtime.Apply(Get_FoldRight1(), pkg_Data_Function.Get_go__const())
	})
	return cache_mkFoldRight1__gopurs_runtime_Value_364767315
}

var cache_foldr1 gopurs_runtime.Value
var once_foldr1 sync.Once
func Get_foldr1() gopurs_runtime.Value {
	once_foldr1.Do(func() {
		cache_foldr1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr1
}

var cache_foldr1__gopurs_runtime_Value_3059734942 gopurs_runtime.Value
var once_foldr1__gopurs_runtime_Value_3059734942 sync.Once
func Get_foldr1__gopurs_runtime_Value_3059734942() gopurs_runtime.Value {
	once_foldr1__gopurs_runtime_Value_3059734942.Do(func() {
		cache_foldr1__gopurs_runtime_Value_3059734942 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1__gopurs_runtime_Value_3059734942(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr1__gopurs_runtime_Value_3059734942
}

var cache_foldl1 gopurs_runtime.Value
var once_foldl1 sync.Once
func Get_foldl1() gopurs_runtime.Value {
	once_foldl1.Do(func() {
		cache_foldl1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl1
}

var cache_foldl1__gopurs_runtime_Value_3059734942 gopurs_runtime.Value
var once_foldl1__gopurs_runtime_Value_3059734942 sync.Once
func Get_foldl1__gopurs_runtime_Value_3059734942() gopurs_runtime.Value {
	once_foldl1__gopurs_runtime_Value_3059734942.Do(func() {
		cache_foldl1__gopurs_runtime_Value_3059734942 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1__gopurs_runtime_Value_3059734942(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl1__gopurs_runtime_Value_3059734942
}

var cache_maximumBy gopurs_runtime.Value
var once_maximumBy sync.Once
func Get_maximumBy() gopurs_runtime.Value {
	once_maximumBy.Do(func() {
		cache_maximumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), cmp_1_box)
})
	})
	return cache_maximumBy
}

var cache_minimumBy gopurs_runtime.Value
var once_minimumBy sync.Once
func Get_minimumBy() gopurs_runtime.Value {
	once_minimumBy.Do(func() {
		cache_minimumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), cmp_1_box)
})
	})
	return cache_minimumBy
}

var cache_foldableTuple gopurs_runtime.Value
var once_foldableTuple sync.Once
func Get_foldableTuple() gopurs_runtime.Value {
	once_foldableTuple.Do(func() {
		cache_foldableTuple = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
})
}))
	})
	return cache_foldableTuple
}

var cache_foldableTuple__gopurs_runtime_Value_3696373503 gopurs_runtime.Value
var once_foldableTuple__gopurs_runtime_Value_3696373503 sync.Once
func Get_foldableTuple__gopurs_runtime_Value_3696373503() gopurs_runtime.Value {
	once_foldableTuple__gopurs_runtime_Value_3696373503.Do(func() {
		cache_foldableTuple__gopurs_runtime_Value_3696373503 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
})
}))
	})
	return cache_foldableTuple__gopurs_runtime_Value_3696373503
}

var cache_foldableMultiplicative gopurs_runtime.Value
var once_foldableMultiplicative sync.Once
func Get_foldableMultiplicative() gopurs_runtime.Value {
	once_foldableMultiplicative.Do(func() {
		cache_foldableMultiplicative = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableMultiplicative
}

var cache_foldableMultiplicative__gopurs_runtime_Value_189846079 gopurs_runtime.Value
var once_foldableMultiplicative__gopurs_runtime_Value_189846079 sync.Once
func Get_foldableMultiplicative__gopurs_runtime_Value_189846079() gopurs_runtime.Value {
	once_foldableMultiplicative__gopurs_runtime_Value_189846079.Do(func() {
		cache_foldableMultiplicative__gopurs_runtime_Value_189846079 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative__gopurs_runtime_Value_1841171440()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableMultiplicative__gopurs_runtime_Value_189846079
}

var cache_foldableIdentity gopurs_runtime.Value
var once_foldableIdentity sync.Once
func Get_foldableIdentity() gopurs_runtime.Value {
	once_foldableIdentity.Do(func() {
		cache_foldableIdentity = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableIdentity
}

var cache_foldableIdentity__gopurs_runtime_Value_189846079 gopurs_runtime.Value
var once_foldableIdentity__gopurs_runtime_Value_189846079 sync.Once
func Get_foldableIdentity__gopurs_runtime_Value_189846079() gopurs_runtime.Value {
	once_foldableIdentity__gopurs_runtime_Value_189846079.Do(func() {
		cache_foldableIdentity__gopurs_runtime_Value_189846079 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity__gopurs_runtime_Value_1841171440()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableIdentity__gopurs_runtime_Value_189846079
}

var cache_foldableDual gopurs_runtime.Value
var once_foldableDual sync.Once
func Get_foldableDual() gopurs_runtime.Value {
	once_foldableDual.Do(func() {
		cache_foldableDual = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableDual
}

var cache_foldableDual__gopurs_runtime_Value_189846079 gopurs_runtime.Value
var once_foldableDual__gopurs_runtime_Value_189846079 sync.Once
func Get_foldableDual__gopurs_runtime_Value_189846079() gopurs_runtime.Value {
	once_foldableDual__gopurs_runtime_Value_189846079.Do(func() {
		cache_foldableDual__gopurs_runtime_Value_189846079 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual__gopurs_runtime_Value_1841171440()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableDual__gopurs_runtime_Value_189846079
}

var cache_foldRight1Semigroup gopurs_runtime.Value
var once_foldRight1Semigroup sync.Once
func Get_foldRight1Semigroup() gopurs_runtime.Value {
	once_foldRight1Semigroup.Do(func() {
		cache_foldRight1Semigroup = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := (*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_FoldRight1[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_FoldRight1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, a_3, f_4)), f_4)
})
}), (*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V1})}))}
})
}))
	})
	return cache_foldRight1Semigroup
}

var cache_foldRight1Semigroup__ptrData_Semigroup_Constructor_Semigroup_ptrConstructor_FoldRight1_gopurs_runtime_Value___1201419834 gopurs_runtime.Value
var once_foldRight1Semigroup__ptrData_Semigroup_Constructor_Semigroup_ptrConstructor_FoldRight1_gopurs_runtime_Value___1201419834 sync.Once
func Get_foldRight1Semigroup__ptrData_Semigroup_Constructor_Semigroup_ptrConstructor_FoldRight1_gopurs_runtime_Value___1201419834() gopurs_runtime.Value {
	once_foldRight1Semigroup__ptrData_Semigroup_Constructor_Semigroup_ptrConstructor_FoldRight1_gopurs_runtime_Value___1201419834.Do(func() {
		cache_foldRight1Semigroup__ptrData_Semigroup_Constructor_Semigroup_ptrConstructor_FoldRight1_gopurs_runtime_Value___1201419834 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&pkg_Data_Semigroup.Constructor_Semigroup[*Constructor_FoldRight1[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := (*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_FoldRight1[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_FoldRight1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, a_3, f_4)), f_4)
})
}), (*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V1})}))}
})
})})}
	})
	return cache_foldRight1Semigroup__ptrData_Semigroup_Constructor_Semigroup_ptrConstructor_FoldRight1_gopurs_runtime_Value___1201419834
}

var cache_semigroupDual gopurs_runtime.Value
var once_semigroupDual sync.Once
func Get_semigroupDual() gopurs_runtime.Value {
	once_semigroupDual.Do(func() {
		cache_semigroupDual = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldRight1Semigroup(), "append"), v1_1, v_0)
})
}))
	})
	return cache_semigroupDual
}

var cache_foldMap1DefaultR gopurs_runtime.Value
var once_foldMap1DefaultR sync.Once
func Get_foldMap1DefaultR() gopurs_runtime.Value {
	once_foldMap1DefaultR.Do(func() {
		cache_foldMap1DefaultR = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1DefaultR(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_2_box))
})
	})
	return cache_foldMap1DefaultR
}

var cache_foldMap1DefaultL gopurs_runtime.Value
var once_foldMap1DefaultL sync.Once
func Get_foldMap1DefaultL() gopurs_runtime.Value {
	once_foldMap1DefaultL.Do(func() {
		cache_foldMap1DefaultL = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1DefaultL(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_2_box))
})
	})
	return cache_foldMap1DefaultL
}

var cache_foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		cache_foldMap1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1
}

var cache_foldMap1__gopurs_runtime_Value_3342855683 gopurs_runtime.Value
var once_foldMap1__gopurs_runtime_Value_3342855683 sync.Once
func Get_foldMap1__gopurs_runtime_Value_3342855683() gopurs_runtime.Value {
	once_foldMap1__gopurs_runtime_Value_3342855683.Do(func() {
		cache_foldMap1__gopurs_runtime_Value_3342855683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__gopurs_runtime_Value_3342855683(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__gopurs_runtime_Value_3342855683
}

var cache_foldl1Default gopurs_runtime.Value
var once_foldl1Default sync.Once
func Get_foldl1Default() gopurs_runtime.Value {
	once_foldl1Default.Do(func() {
		cache_foldl1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1Default(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_foldl1Default
}

var cache_foldr1Default gopurs_runtime.Value
var once_foldr1Default sync.Once
func Get_foldr1Default() gopurs_runtime.Value {
	once_foldr1Default.Do(func() {
		cache_foldr1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1Default(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_foldr1Default
}

var cache_intercalateMap gopurs_runtime.Value
var once_intercalateMap sync.Once
func Get_intercalateMap() gopurs_runtime.Value {
	once_intercalateMap.Do(func() {
		cache_intercalateMap = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalateMap(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_intercalateMap
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_intercalate
}

var cache_maximum gopurs_runtime.Value
var once_maximum sync.Once
func Get_maximum() gopurs_runtime.Value {
	once_maximum.Do(func() {
		cache_maximum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximum(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_maximum
}

var cache_minimum gopurs_runtime.Value
var once_minimum sync.Once
func Get_minimum() gopurs_runtime.Value {
	once_minimum.Do(func() {
		cache_minimum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimum(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_minimum
}

var cache_traverse1_ gopurs_runtime.Value
var once_traverse1_ sync.Once
func Get_traverse1_() gopurs_runtime.Value {
	once_traverse1_.Do(func() {
		cache_traverse1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1_(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_traverse1_
}

var cache_traverse1___gopurs_runtime_Value_3055398386 gopurs_runtime.Value
var once_traverse1___gopurs_runtime_Value_3055398386 sync.Once
func Get_traverse1___gopurs_runtime_Value_3055398386() gopurs_runtime.Value {
	once_traverse1___gopurs_runtime_Value_3055398386.Do(func() {
		cache_traverse1___gopurs_runtime_Value_3055398386 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1___gopurs_runtime_Value_3055398386(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_traverse1___gopurs_runtime_Value_3055398386
}

var cache_for1_ gopurs_runtime.Value
var once_for1_ sync.Once
func Get_for1_() gopurs_runtime.Value {
	once_for1_.Do(func() {
		cache_for1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_for1_(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_for1_
}

var cache_sequence1_ gopurs_runtime.Value
var once_sequence1_ sync.Once
func Get_sequence1_() gopurs_runtime.Value {
	once_sequence1_.Do(func() {
		cache_sequence1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1_(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_sequence1_
}

var cache_fold1 gopurs_runtime.Value
var once_fold1 sync.Once
func Get_fold1() gopurs_runtime.Value {
	once_fold1.Do(func() {
		cache_fold1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold1(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_fold1
}

type Constructor_FoldRight1[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}


type Constructor_Foldable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2465059545] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Foldable1[gopurs_runtime.Value])(ptr)
		switch key {
		case "Foldable0": return c.V0
		case "foldMap1": return c.V1
		case "foldl1": return c.V2
		case "foldr1": return c.V3
		default: panic("Key not found in dictionary Constructor_Foldable1: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_foldr1(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_foldr1__gopurs_runtime_Value_3059734942(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_foldl1(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldl1__gopurs_runtime_Value_3059734942(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_maximumBy(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(dictFoldable1_0.V2, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, x_2, y_3), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(nil)}).IntVal) != (0) {
__t0 = x_2
goto end_branch_0
} else {

}
}
{
__t0 = y_3
}
end_branch_0:
return __t0
})
}))
}

func Call_minimumBy(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(dictFoldable1_0.V2, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, x_2, y_3), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(nil)}).IntVal) != (0) {
__t0 = x_2
goto end_branch_0
} else {

}
}
{
__t0 = y_3
}
end_branch_0:
return __t0
})
}))
}

func Call_foldMap1DefaultR(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictFunctor_1_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], dictSemigroup_2_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_2_loop
_ = dictSemigroup_2
append_3_0 := dictSemigroup_2.V0
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(dictFunctor_1.V0, f_4)
_ = __local_var_5_1
__local_var_6_2 := gopurs_runtime.Apply(dictFoldable1_0.V3, append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_foldMap1DefaultL(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictFunctor_1_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], dictSemigroup_2_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_2_loop
_ = dictSemigroup_2
append_3_0 := dictSemigroup_2.V0
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(dictFunctor_1.V0, f_4)
_ = __local_var_5_1
__local_var_6_2 := gopurs_runtime.Apply(dictFoldable1_0.V2, append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_foldMap1(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__gopurs_runtime_Value_3342855683(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl1Default(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_0 := gopurs_runtime.Apply2(dictFoldable1_0.V1, Get_semigroupDual(), Get_mkFoldRight1())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_1_0, a_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, (*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(x_2, a_6, b_5)
})
}))
})
})
}

func Call_foldr1Default(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_0 := gopurs_runtime.Apply2(dictFoldable1_0.V1, Get_foldRight1Semigroup(), Get_mkFoldRight1())
_ = __local_var_1_0
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_1_0, a_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, (*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, b_2)
})
})
}

func Call_intercalateMap(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
semigroupJoinWith1_2_0 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_1.V0, gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(dictSemigroup_1.V0, j_4, gopurs_runtime.Apply(v1_3, j_4)))
})
})
})}
_ = semigroupJoinWith1_2_0
return gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_2_0)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_1 := gopurs_runtime.Apply(f_4, x_6)
_ = __local_var_7_1
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_1
})
}), foldable_5, j_3)
})
})
})
}

func Call_intercalate(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
semigroupJoinWith1_2_0 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_1.V0, gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(dictSemigroup_1.V0, j_4, gopurs_runtime.Apply(v1_3, j_4)))
})
})
})}
_ = semigroupJoinWith1_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_2_0)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), foldable_4, a_3)
})
})
}

func Call_maximum(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
semigroupMax_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply2(dictOrd_0.V1, v_1, v1_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 1527465420) {
__t2 = v1_2
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 902936544) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 380165415) {
__t2 = v_1
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
}))
_ = semigroupMax_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), semigroupMax_1_0, pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
}

func Call_minimum(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
semigroupMin_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply2(dictOrd_0.V1, v_1, v1_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 1527465420) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 902936544) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 380165415) {
__t2 = v1_2
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
}))
_ = semigroupMin_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), semigroupMin_1_0, pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
}

func Call_traverse1_(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
Functor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}))
_ = Functor0_2_0
semigroupAct1_3_1 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_1.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_0.V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Apply3(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
}

func Call_traverse1___gopurs_runtime_Value_3055398386(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
Functor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}))
_ = Functor0_2_0
semigroupAct1_3_1 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_1.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_0.V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Apply3(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
}

func Call_for1_(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
__local_var_2_0 := Call_traverse1_(dictFoldable1_0, dictApply_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
})
}

func Call_sequence1_(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
return gopurs_runtime.Apply(Call_traverse1_(dictFoldable1_0, dictApply_1), Get_identity1())
}

func Call_fold1(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.Apply2(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_1)}, Get_identity())
}


