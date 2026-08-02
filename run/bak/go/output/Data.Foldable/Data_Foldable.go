package Data_Foldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Maybe_First "gopurs/output/Data.Maybe.First"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
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

var cache_monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		cache_monoidEndo = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_0
semigroupEndo1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "compose"), v_1, v1_2)
}))
_ = semigroupEndo1_1_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_1_1
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}()
	})
	return cache_monoidEndo
}

var cache_Empty gopurs_runtime.Value
var once_Empty sync.Once
func Get_Empty() gopurs_runtime.Value {
	once_Empty.Do(func() {
		cache_Empty = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: nil}
	})
	return cache_Empty
}

var cache_Node gopurs_runtime.Value
var once_Node sync.Once
func Get_Node() gopurs_runtime.Value {
	once_Node.Do(func() {
		cache_Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2421944209, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Node
}

var cache_Append gopurs_runtime.Value
var once_Append sync.Once
func Get_Append() gopurs_runtime.Value {
	once_Append.Do(func() {
		cache_Append = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Append
}

var cache_semigroupFreeMonoidTree gopurs_runtime.Value
var once_semigroupFreeMonoidTree sync.Once
func Get_semigroupFreeMonoidTree() gopurs_runtime.Value {
	once_semigroupFreeMonoidTree.Do(func() {
		cache_semigroupFreeMonoidTree = gopurs_runtime.RecordDict1("append", Get_Append())
	})
	return cache_semigroupFreeMonoidTree
}

var cache_monoidFreeMonoidTree gopurs_runtime.Value
var once_monoidFreeMonoidTree sync.Once
func Get_monoidFreeMonoidTree() gopurs_runtime.Value {
	once_monoidFreeMonoidTree.Do(func() {
		cache_monoidFreeMonoidTree = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupFreeMonoidTree()
}), gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: nil})
	})
	return cache_monoidFreeMonoidTree
}

var cache_foldr gopurs_runtime.Value
var once_foldr sync.Once
func Get_foldr() gopurs_runtime.Value {
	once_foldr.Do(func() {
		cache_foldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr(dict_0_box)
})
	})
	return cache_foldr
}

var cache_foldr__gopurs_runtime_Value_473908296 gopurs_runtime.Value
var once_foldr__gopurs_runtime_Value_473908296 sync.Once
func Get_foldr__gopurs_runtime_Value_473908296() gopurs_runtime.Value {
	once_foldr__gopurs_runtime_Value_473908296.Do(func() {
		cache_foldr__gopurs_runtime_Value_473908296 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__gopurs_runtime_Value_473908296(dict_0_box)
})
	})
	return cache_foldr__gopurs_runtime_Value_473908296
}

var cache_foldr__gopurs_runtime_Value_3006709320 gopurs_runtime.Value
var once_foldr__gopurs_runtime_Value_3006709320 sync.Once
func Get_foldr__gopurs_runtime_Value_3006709320() gopurs_runtime.Value {
	once_foldr__gopurs_runtime_Value_3006709320.Do(func() {
		cache_foldr__gopurs_runtime_Value_3006709320 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__gopurs_runtime_Value_3006709320(dict_0_box)
})
	})
	return cache_foldr__gopurs_runtime_Value_3006709320
}

var cache_indexr gopurs_runtime.Value
var once_indexr sync.Once
func Get_indexr() gopurs_runtime.Value {
	once_indexr.Do(func() {
		cache_indexr = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, idx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_indexr(dictFoldable_0_box, idx_1_box.IntVal)
})
	})
	return cache_indexr
}

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_null(dictFoldable_0_box)
})
	})
	return cache_null
}

var cache_oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		cache_oneOf = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf(dictFoldable_0_box, dictPlus_1_box)
})
	})
	return cache_oneOf
}

var cache_oneOfMap gopurs_runtime.Value
var once_oneOfMap sync.Once
func Get_oneOfMap() gopurs_runtime.Value {
	once_oneOfMap.Do(func() {
		cache_oneOfMap = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOfMap(dictFoldable_0_box, dictPlus_1_box)
})
	})
	return cache_oneOfMap
}

var cache_oneOfMap__gopurs_runtime_Value_2786491193 gopurs_runtime.Value
var once_oneOfMap__gopurs_runtime_Value_2786491193 sync.Once
func Get_oneOfMap__gopurs_runtime_Value_2786491193() gopurs_runtime.Value {
	once_oneOfMap__gopurs_runtime_Value_2786491193.Do(func() {
		cache_oneOfMap__gopurs_runtime_Value_2786491193 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOfMap__gopurs_runtime_Value_2786491193(dictFoldable_0_box, dictPlus_1_box)
})
	})
	return cache_oneOfMap__gopurs_runtime_Value_2786491193
}

var cache_traverse_ gopurs_runtime.Value
var once_traverse_ sync.Once
func Get_traverse_() gopurs_runtime.Value {
	once_traverse_.Do(func() {
		cache_traverse_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse_(dictApplicative_0_box)
})
	})
	return cache_traverse_
}

var cache_traverse___gopurs_runtime_Value_999280278 gopurs_runtime.Value
var once_traverse___gopurs_runtime_Value_999280278 sync.Once
func Get_traverse___gopurs_runtime_Value_999280278() gopurs_runtime.Value {
	once_traverse___gopurs_runtime_Value_999280278.Do(func() {
		cache_traverse___gopurs_runtime_Value_999280278 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse___gopurs_runtime_Value_999280278(dictApplicative_0_box)
})
	})
	return cache_traverse___gopurs_runtime_Value_999280278
}

var cache_for_ gopurs_runtime.Value
var once_for_ sync.Once
func Get_for_() gopurs_runtime.Value {
	once_for_.Do(func() {
		cache_for_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_for_(dictApplicative_0_box)
})
	})
	return cache_for_
}

var cache_sequence_ gopurs_runtime.Value
var once_sequence_ sync.Once
func Get_sequence_() gopurs_runtime.Value {
	once_sequence_.Do(func() {
		cache_sequence_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence_(dictApplicative_0_box)
})
	})
	return cache_sequence_
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(dict_0_box)
})
	})
	return cache_foldl
}

var cache_foldl__gopurs_runtime_Value_3006709320 gopurs_runtime.Value
var once_foldl__gopurs_runtime_Value_3006709320 sync.Once
func Get_foldl__gopurs_runtime_Value_3006709320() gopurs_runtime.Value {
	once_foldl__gopurs_runtime_Value_3006709320.Do(func() {
		cache_foldl__gopurs_runtime_Value_3006709320 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__gopurs_runtime_Value_3006709320(dict_0_box)
})
	})
	return cache_foldl__gopurs_runtime_Value_3006709320
}

var cache_indexl gopurs_runtime.Value
var once_indexl sync.Once
func Get_indexl() gopurs_runtime.Value {
	once_indexl.Do(func() {
		cache_indexl = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, idx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_indexl(dictFoldable_0_box, idx_1_box.IntVal)
})
	})
	return cache_indexl
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate(dictFoldable_0_box, dictMonoid_1_box)
})
	})
	return cache_intercalate
}

var cache_intercalate__gopurs_runtime_Value_363090812 gopurs_runtime.Value
var once_intercalate__gopurs_runtime_Value_363090812 sync.Once
func Get_intercalate__gopurs_runtime_Value_363090812() gopurs_runtime.Value {
	once_intercalate__gopurs_runtime_Value_363090812.Do(func() {
		cache_intercalate__gopurs_runtime_Value_363090812 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate__gopurs_runtime_Value_363090812(dictFoldable_0_box, dictMonoid_1_box)
})
	})
	return cache_intercalate__gopurs_runtime_Value_363090812
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length(dictFoldable_0_box, dictSemiring_1_box)
})
	})
	return cache_length
}

var cache_length__gopurs_runtime_Value_2374138440 gopurs_runtime.Value
var once_length__gopurs_runtime_Value_2374138440 sync.Once
func Get_length__gopurs_runtime_Value_2374138440() gopurs_runtime.Value {
	once_length__gopurs_runtime_Value_2374138440.Do(func() {
		cache_length__gopurs_runtime_Value_2374138440 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__gopurs_runtime_Value_2374138440(dictFoldable_0_box, dictSemiring_1_box)
})
	})
	return cache_length__gopurs_runtime_Value_2374138440
}

var cache_maximumBy gopurs_runtime.Value
var once_maximumBy sync.Once
func Get_maximumBy() gopurs_runtime.Value {
	once_maximumBy.Do(func() {
		cache_maximumBy = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy(dictFoldable_0_box, cmp_1_box)
})
	})
	return cache_maximumBy
}

var cache_maximumBy__gopurs_runtime_Value_562089896 gopurs_runtime.Value
var once_maximumBy__gopurs_runtime_Value_562089896 sync.Once
func Get_maximumBy__gopurs_runtime_Value_562089896() gopurs_runtime.Value {
	once_maximumBy__gopurs_runtime_Value_562089896.Do(func() {
		cache_maximumBy__gopurs_runtime_Value_562089896 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy__gopurs_runtime_Value_562089896(dictFoldable_0_box, cmp_1_box)
})
	})
	return cache_maximumBy__gopurs_runtime_Value_562089896
}

var cache_maximum gopurs_runtime.Value
var once_maximum sync.Once
func Get_maximum() gopurs_runtime.Value {
	once_maximum.Do(func() {
		cache_maximum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximum(dictOrd_0_box)
})
	})
	return cache_maximum
}

var cache_minimumBy gopurs_runtime.Value
var once_minimumBy sync.Once
func Get_minimumBy() gopurs_runtime.Value {
	once_minimumBy.Do(func() {
		cache_minimumBy = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy(dictFoldable_0_box, cmp_1_box)
})
	})
	return cache_minimumBy
}

var cache_minimumBy__gopurs_runtime_Value_562089896 gopurs_runtime.Value
var once_minimumBy__gopurs_runtime_Value_562089896 sync.Once
func Get_minimumBy__gopurs_runtime_Value_562089896() gopurs_runtime.Value {
	once_minimumBy__gopurs_runtime_Value_562089896.Do(func() {
		cache_minimumBy__gopurs_runtime_Value_562089896 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy__gopurs_runtime_Value_562089896(dictFoldable_0_box, cmp_1_box)
})
	})
	return cache_minimumBy__gopurs_runtime_Value_562089896
}

var cache_minimum gopurs_runtime.Value
var once_minimum sync.Once
func Get_minimum() gopurs_runtime.Value {
	once_minimum.Do(func() {
		cache_minimum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimum(dictOrd_0_box)
})
	})
	return cache_minimum
}

var cache_product gopurs_runtime.Value
var once_product sync.Once
func Get_product() gopurs_runtime.Value {
	once_product.Do(func() {
		cache_product = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product(dictFoldable_0_box, dictSemiring_1_box)
})
	})
	return cache_product
}

var cache_sum gopurs_runtime.Value
var once_sum sync.Once
func Get_sum() gopurs_runtime.Value {
	once_sum.Do(func() {
		cache_sum = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sum(dictFoldable_0_box, dictSemiring_1_box)
})
	})
	return cache_sum
}

var cache_sum__gopurs_runtime_Value_2374138440 gopurs_runtime.Value
var once_sum__gopurs_runtime_Value_2374138440 sync.Once
func Get_sum__gopurs_runtime_Value_2374138440() gopurs_runtime.Value {
	once_sum__gopurs_runtime_Value_2374138440.Do(func() {
		cache_sum__gopurs_runtime_Value_2374138440 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sum__gopurs_runtime_Value_2374138440(dictFoldable_0_box, dictSemiring_1_box)
})
	})
	return cache_sum__gopurs_runtime_Value_2374138440
}

var cache_foldableTuple gopurs_runtime.Value
var once_foldableTuple sync.Once
func Get_foldableTuple() gopurs_runtime.Value {
	once_foldableTuple.Do(func() {
		cache_foldableTuple = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
}))
	})
	return cache_foldableTuple
}

var cache_foldableMultiplicative gopurs_runtime.Value
var once_foldableMultiplicative sync.Once
func Get_foldableMultiplicative() gopurs_runtime.Value {
	once_foldableMultiplicative.Do(func() {
		cache_foldableMultiplicative = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}))
	})
	return cache_foldableMultiplicative
}

var cache_foldableMaybe gopurs_runtime.Value
var once_foldableMaybe sync.Once
func Get_foldableMaybe() gopurs_runtime.Value {
	once_foldableMaybe.Do(func() {
		cache_foldableMaybe = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_foldableMaybe
}

var cache_foldableIdentity gopurs_runtime.Value
var once_foldableIdentity sync.Once
func Get_foldableIdentity() gopurs_runtime.Value {
	once_foldableIdentity.Do(func() {
		cache_foldableIdentity = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}))
	})
	return cache_foldableIdentity
}

var cache_foldableEither gopurs_runtime.Value
var once_foldableEither sync.Once
func Get_foldableEither() gopurs_runtime.Value {
	once_foldableEither.Do(func() {
		cache_foldableEither = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_foldableEither
}

var cache_foldableDual gopurs_runtime.Value
var once_foldableDual sync.Once
func Get_foldableDual() gopurs_runtime.Value {
	once_foldableDual.Do(func() {
		cache_foldableDual = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}))
	})
	return cache_foldableDual
}

var cache_foldableDisj gopurs_runtime.Value
var once_foldableDisj sync.Once
func Get_foldableDisj() gopurs_runtime.Value {
	once_foldableDisj.Do(func() {
		cache_foldableDisj = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}))
	})
	return cache_foldableDisj
}

var cache_foldableConst gopurs_runtime.Value
var once_foldableConst sync.Once
func Get_foldableConst() gopurs_runtime.Value {
	once_foldableConst.Do(func() {
		cache_foldableConst = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
}))
	})
	return cache_foldableConst
}

var cache_foldableConj gopurs_runtime.Value
var once_foldableConj sync.Once
func Get_foldableConj() gopurs_runtime.Value {
	once_foldableConj.Do(func() {
		cache_foldableConj = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}))
	})
	return cache_foldableConj
}

var cache_foldableAdditive gopurs_runtime.Value
var once_foldableAdditive sync.Once
func Get_foldableAdditive() gopurs_runtime.Value {
	once_foldableAdditive.Do(func() {
		cache_foldableAdditive = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}))
	})
	return cache_foldableAdditive
}

var cache_foldMapDefaultR gopurs_runtime.Value
var once_foldMapDefaultR sync.Once
func Get_foldMapDefaultR() gopurs_runtime.Value {
	once_foldMapDefaultR.Do(func() {
		cache_foldMapDefaultR = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapDefaultR(dictFoldable_0_box, dictMonoid_1_box)
})
	})
	return cache_foldMapDefaultR
}

var cache_foldableArray gopurs_runtime.Value
var once_foldableArray sync.Once
func Get_foldableArray() gopurs_runtime.Value {
	once_foldableArray.Do(func() {
		cache_foldableArray = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableArray(), "foldr"), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_2, x_3), acc_4)
}), mempty_1_0)
})
}), Get_foldlArray(), Get_foldrArray())
	})
	return cache_foldableArray
}

var cache_foldableFreeMonoidTree gopurs_runtime.Value
var once_foldableFreeMonoidTree sync.Once
func Get_foldableFreeMonoidTree() gopurs_runtime.Value {
	once_foldableFreeMonoidTree.Do(func() {
		cache_foldableFreeMonoidTree = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_2, x_3), acc_4)
}), mempty_1_0)
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_1_0 gopurs_runtime.Value
go__go_1_1_0 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_1_0:
for {
if false { continue go__go_1_1_0 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t2 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*Constructor_Node[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: nil}
continue go__go_1_1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
continue go__go_1_1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1, rhs_4})}
continue go__go_1_1_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
var __t5 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
__t5 = acc_2
goto end_branch_5
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: nil}
continue go__go_1_1_0
__t5 = gopurs_runtime.Value{}
}
end_branch_5:
__t2 = __t5
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
})
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_1_0, a_2, b_3, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: nil})
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_6_1 gopurs_runtime.Value
go__go_1_6_1 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_6_1:
for {
if false { continue go__go_1_6_1 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t7 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*Constructor_Node[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: nil}
rhs_4_loop = lhs_3
continue go__go_1_6_1
__t7 = gopurs_runtime.Value{}
goto end_branch_7
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t8 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_6_1
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_6_1
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, lhs_3, (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_6_1
__t8 = gopurs_runtime.Value{}
}
end_branch_8:
__t7 = __t8
goto end_branch_7
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t10 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t10 = acc_2
goto end_branch_10
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: nil}
rhs_4_loop = lhs_3
continue go__go_1_6_1
__t10 = gopurs_runtime.Value{}
}
end_branch_10:
__t7 = __t10
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
})
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_6_1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: nil}, b_3)
})
}))
	})
	return cache_foldableFreeMonoidTree
}

var cache_foldMapDefaultL gopurs_runtime.Value
var once_foldMapDefaultL sync.Once
func Get_foldMapDefaultL() gopurs_runtime.Value {
	once_foldMapDefaultL.Do(func() {
		cache_foldMapDefaultL = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapDefaultL(dictFoldable_0_box, dictMonoid_1_box)
})
	})
	return cache_foldMapDefaultL
}

var cache_foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		cache_foldMap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap(dict_0_box)
})
	})
	return cache_foldMap
}

var cache_foldMap__gopurs_runtime_Value_1916751092 gopurs_runtime.Value
var once_foldMap__gopurs_runtime_Value_1916751092 sync.Once
func Get_foldMap__gopurs_runtime_Value_1916751092() gopurs_runtime.Value {
	once_foldMap__gopurs_runtime_Value_1916751092.Do(func() {
		cache_foldMap__gopurs_runtime_Value_1916751092 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__gopurs_runtime_Value_1916751092(dict_0_box)
})
	})
	return cache_foldMap__gopurs_runtime_Value_1916751092
}

var cache_foldableApp gopurs_runtime.Value
var once_foldableApp sync.Once
func Get_foldableApp() gopurs_runtime.Value {
	once_foldableApp.Do(func() {
		cache_foldableApp = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableApp(dictFoldable_0_box)
})
	})
	return cache_foldableApp
}

var cache_foldableCompose gopurs_runtime.Value
var once_foldableCompose sync.Once
func Get_foldableCompose() gopurs_runtime.Value {
	once_foldableCompose.Do(func() {
		cache_foldableCompose = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableCompose(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_foldableCompose
}

var cache_foldableCoproduct gopurs_runtime.Value
var once_foldableCoproduct sync.Once
func Get_foldableCoproduct() gopurs_runtime.Value {
	once_foldableCoproduct.Do(func() {
		cache_foldableCoproduct = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableCoproduct(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_foldableCoproduct
}

var cache_foldableFirst gopurs_runtime.Value
var once_foldableFirst sync.Once
func Get_foldableFirst() gopurs_runtime.Value {
	once_foldableFirst.Do(func() {
		cache_foldableFirst = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldMap"), dictMonoid_0)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldl"), f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldr"), f_0, z_1, v_2)
}))
	})
	return cache_foldableFirst
}

var cache_foldableLast gopurs_runtime.Value
var once_foldableLast sync.Once
func Get_foldableLast() gopurs_runtime.Value {
	once_foldableLast.Do(func() {
		cache_foldableLast = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldMap"), dictMonoid_0)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldl"), f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldr"), f_0, z_1, v_2)
}))
	})
	return cache_foldableLast
}

var cache_foldableProduct gopurs_runtime.Value
var once_foldableProduct sync.Once
func Get_foldableProduct() gopurs_runtime.Value {
	once_foldableProduct.Do(func() {
		cache_foldableProduct = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableProduct(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_foldableProduct
}

var cache_foldlDefault gopurs_runtime.Value
var once_foldlDefault sync.Once
func Get_foldlDefault() gopurs_runtime.Value {
	once_foldlDefault.Do(func() {
		cache_foldlDefault = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlDefault(dictFoldable_0_box)
})
	})
	return cache_foldlDefault
}

var cache_foldrDefault gopurs_runtime.Value
var once_foldrDefault sync.Once
func Get_foldrDefault() gopurs_runtime.Value {
	once_foldrDefault.Do(func() {
		cache_foldrDefault = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault(dictFoldable_0_box)
})
	})
	return cache_foldrDefault
}

var cache_foldrDefault__gopurs_runtime_Value_2409537192 gopurs_runtime.Value
var once_foldrDefault__gopurs_runtime_Value_2409537192 sync.Once
func Get_foldrDefault__gopurs_runtime_Value_2409537192() gopurs_runtime.Value {
	once_foldrDefault__gopurs_runtime_Value_2409537192.Do(func() {
		cache_foldrDefault__gopurs_runtime_Value_2409537192 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault__gopurs_runtime_Value_2409537192(dictFoldable_0_box)
})
	})
	return cache_foldrDefault__gopurs_runtime_Value_2409537192
}

var cache_foldrDefault__gopurs_runtime_Value_3006709320 gopurs_runtime.Value
var once_foldrDefault__gopurs_runtime_Value_3006709320 sync.Once
func Get_foldrDefault__gopurs_runtime_Value_3006709320() gopurs_runtime.Value {
	once_foldrDefault__gopurs_runtime_Value_3006709320.Do(func() {
		cache_foldrDefault__gopurs_runtime_Value_3006709320 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault__gopurs_runtime_Value_3006709320(dictFoldable_0_box)
})
	})
	return cache_foldrDefault__gopurs_runtime_Value_3006709320
}

var cache_lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		cache_lookup = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup(dictFoldable_0_box)
})
	})
	return cache_lookup
}

var cache_surroundMap gopurs_runtime.Value
var once_surroundMap sync.Once
func Get_surroundMap() gopurs_runtime.Value {
	once_surroundMap.Do(func() {
		cache_surroundMap = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_surroundMap(dictFoldable_0_box)
})
	})
	return cache_surroundMap
}

var cache_surroundMap__gopurs_runtime_Value_75699306 gopurs_runtime.Value
var once_surroundMap__gopurs_runtime_Value_75699306 sync.Once
func Get_surroundMap__gopurs_runtime_Value_75699306() gopurs_runtime.Value {
	once_surroundMap__gopurs_runtime_Value_75699306.Do(func() {
		cache_surroundMap__gopurs_runtime_Value_75699306 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_surroundMap__gopurs_runtime_Value_75699306(dictFoldable_0_box)
})
	})
	return cache_surroundMap__gopurs_runtime_Value_75699306
}

var cache_surround gopurs_runtime.Value
var once_surround sync.Once
func Get_surround() gopurs_runtime.Value {
	once_surround.Do(func() {
		cache_surround = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_surround(dictFoldable_0_box)
})
	})
	return cache_surround
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, b0_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictFoldable_0_box, dictMonad_1_box, f_2_box, b0_3_box)
})
	})
	return cache_foldM
}

var cache_fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		cache_fold = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold(dictFoldable_0_box, dictMonoid_1_box)
})
	})
	return cache_fold
}

var cache_fold__gopurs_runtime_Value_3625874517 gopurs_runtime.Value
var once_fold__gopurs_runtime_Value_3625874517 sync.Once
func Get_fold__gopurs_runtime_Value_3625874517() gopurs_runtime.Value {
	once_fold__gopurs_runtime_Value_3625874517.Do(func() {
		cache_fold__gopurs_runtime_Value_3625874517 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold__gopurs_runtime_Value_3625874517(dictFoldable_0_box, dictMonoid_1_box)
})
	})
	return cache_fold__gopurs_runtime_Value_3625874517
}

var cache_findMap gopurs_runtime.Value
var once_findMap sync.Once
func Get_findMap() gopurs_runtime.Value {
	once_findMap.Do(func() {
		cache_findMap = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findMap(dictFoldable_0_box, p_1_box)
})
	})
	return cache_findMap
}

var cache_find gopurs_runtime.Value
var once_find sync.Once
func Get_find() gopurs_runtime.Value {
	once_find.Do(func() {
		cache_find = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_find(dictFoldable_0_box, p_1_box)
})
	})
	return cache_find
}

var cache_any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		cache_any = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_any(dictFoldable_0_box, dictHeytingAlgebra_1_box)
})
	})
	return cache_any
}

var cache_any__gopurs_runtime_Value_968068125 gopurs_runtime.Value
var once_any__gopurs_runtime_Value_968068125 sync.Once
func Get_any__gopurs_runtime_Value_968068125() gopurs_runtime.Value {
	once_any__gopurs_runtime_Value_968068125.Do(func() {
		cache_any__gopurs_runtime_Value_968068125 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_any__gopurs_runtime_Value_968068125(dictFoldable_0_box, dictHeytingAlgebra_1_box)
})
	})
	return cache_any__gopurs_runtime_Value_968068125
}

var cache_elem gopurs_runtime.Value
var once_elem sync.Once
func Get_elem() gopurs_runtime.Value {
	once_elem.Do(func() {
		cache_elem = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elem(dictFoldable_0_box)
})
	})
	return cache_elem
}

var cache_elem__gopurs_runtime_Value_2313606065 gopurs_runtime.Value
var once_elem__gopurs_runtime_Value_2313606065 sync.Once
func Get_elem__gopurs_runtime_Value_2313606065() gopurs_runtime.Value {
	once_elem__gopurs_runtime_Value_2313606065.Do(func() {
		cache_elem__gopurs_runtime_Value_2313606065 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elem__gopurs_runtime_Value_2313606065(dictFoldable_0_box)
})
	})
	return cache_elem__gopurs_runtime_Value_2313606065
}

var cache_notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		cache_notElem = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_notElem(dictFoldable_0_box)
})
	})
	return cache_notElem
}

var cache_or gopurs_runtime.Value
var once_or sync.Once
func Get_or() gopurs_runtime.Value {
	once_or.Do(func() {
		cache_or = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_or(dictFoldable_0_box, dictHeytingAlgebra_1_box)
})
	})
	return cache_or
}

var cache_all gopurs_runtime.Value
var once_all sync.Once
func Get_all() gopurs_runtime.Value {
	once_all.Do(func() {
		cache_all = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_all(dictFoldable_0_box, dictHeytingAlgebra_1_box)
})
	})
	return cache_all
}

var cache_all__gopurs_runtime_Value_968068125 gopurs_runtime.Value
var once_all__gopurs_runtime_Value_968068125 sync.Once
func Get_all__gopurs_runtime_Value_968068125() gopurs_runtime.Value {
	once_all__gopurs_runtime_Value_968068125.Do(func() {
		cache_all__gopurs_runtime_Value_968068125 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_all__gopurs_runtime_Value_968068125(dictFoldable_0_box, dictHeytingAlgebra_1_box)
})
	})
	return cache_all__gopurs_runtime_Value_968068125
}

var cache_and gopurs_runtime.Value
var once_and sync.Once
func Get_and() gopurs_runtime.Value {
	once_and.Do(func() {
		cache_and = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_and(dictFoldable_0_box, dictHeytingAlgebra_1_box)
})
	})
	return cache_and
}

type Constructor_Empty[T_a any] struct {
	Rc uint32
}


type Constructor_Node[T_a any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Append[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_foldr(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldr")
}

func Call_foldr__gopurs_runtime_Value_473908296(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldr")
}

func Call_foldr__gopurs_runtime_Value_3006709320(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldr")
}

func Call_indexr(dictFoldable_0_loop gopurs_runtime.Value, idx_1_loop int64) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var idx_1 int64 = idx_1_loop
_ = idx_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, cursor_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.RecordGet(cursor_3, "elem")
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
__t1 = cursor_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.RecordGet(cursor_3, "pos").IntVal) == (idx_1) {
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_2})}, gopurs_runtime.RecordGet(cursor_3, "pos"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.RecordGet(cursor_3, "elem"), gopurs_runtime.Int((gopurs_runtime.RecordGet(cursor_3, "pos").IntVal) + (1)))
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Int(0)))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_2_0, x_3), "elem")
})
}

func Call_null(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
}), gopurs_runtime.Bool(true))
}

func Call_oneOf(dictFoldable_0_loop gopurs_runtime.Value, dictPlus_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 gopurs_runtime.Value = dictPlus_1_loop
_ = dictPlus_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_1, "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.RecordGet(dictPlus_1, "empty"))
}

func Call_oneOfMap(dictFoldable_0_loop gopurs_runtime.Value, dictPlus_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 gopurs_runtime.Value = dictPlus_1_loop
_ = dictPlus_1
empty_2_0 := gopurs_runtime.RecordGet(dictPlus_1, "empty")
_ = empty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_1, "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply(f_3, x_4))
}), empty_2_0)
})
}

func Call_oneOfMap__gopurs_runtime_Value_2786491193(dictFoldable_0_loop gopurs_runtime.Value, dictPlus_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 gopurs_runtime.Value = dictPlus_1_loop
_ = dictPlus_1
empty_2_0 := gopurs_runtime.RecordGet(dictPlus_1, "empty")
_ = empty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_1, "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply(f_3, x_4))
}), empty_2_0)
})
}

func Call_traverse_(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(dictFoldable_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(f_3, x_4)
_ = __local_var_5_1
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_identity()
}), __local_var_5_1), b_6)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_Unit.Get_unit()))
})
}

func Call_traverse___gopurs_runtime_Value_999280278(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(dictFoldable_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(f_3, x_4)
_ = __local_var_5_1
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_identity()
}), __local_var_5_1), b_6)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_Unit.Get_unit()))
})
}

func Call_for_(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
traverse_1_1_0 := Call_traverse_(dictApplicative_0)
_ = traverse_1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(traverse_1_1_0, dictFoldable_2)
_ = __local_var_3_1
return gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_3_1, a_5, b_4)
})
})
}

func Call_sequence_(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
traverse_1_1_0 := Call_traverse_(dictApplicative_0)
_ = traverse_1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(traverse_1_1_0, dictFoldable_2, Get_identity())
})
}

func Call_foldl(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldl")
}

func Call_foldl__gopurs_runtime_Value_3006709320(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldl")
}

func Call_indexl(dictFoldable_0_loop gopurs_runtime.Value, idx_1_loop int64) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var idx_1 int64 = idx_1_loop
_ = idx_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(cursor_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.RecordGet(cursor_2, "elem")
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
__t1 = cursor_2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.RecordGet(cursor_2, "pos").IntVal) == (idx_1) {
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_3})}, gopurs_runtime.RecordGet(cursor_2, "pos"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.RecordGet(cursor_2, "elem"), gopurs_runtime.Int((gopurs_runtime.RecordGet(cursor_2, "pos").IntVal) + (1)))
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Int(0)))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_2_0, x_3), "elem")
})
}

func Call_intercalate(dictFoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_0
mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_3_1
return gopurs_runtime.Func2(func(sep_4 gopurs_runtime.Value, xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_6, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_7, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), sep_4, v1_7)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
}), gopurs_runtime.RecordDict2("acc", "init", mempty_3_1, gopurs_runtime.Bool(true)), xs_5), "acc")
})
}

func Call_intercalate__gopurs_runtime_Value_363090812(dictFoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_0
mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_3_1
return gopurs_runtime.Func2(func(sep_4 gopurs_runtime.Value, xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_6, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_7, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), sep_4, v1_7)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
}), gopurs_runtime.RecordDict2("acc", "init", mempty_3_1, gopurs_runtime.Bool(true)), xs_5), "acc")
})
}

func Call_length(dictFoldable_0_loop gopurs_runtime.Value, dictSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 gopurs_runtime.Value = dictSemiring_1_loop
_ = dictSemiring_1
one_2_0 := gopurs_runtime.RecordGet(dictSemiring_1, "one")
_ = one_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(c_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_1, "add"), one_2_0, c_3)
}), gopurs_runtime.RecordGet(dictSemiring_1, "zero"))
}

func Call_length__gopurs_runtime_Value_2374138440(dictFoldable_0_loop gopurs_runtime.Value, dictSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 gopurs_runtime.Value = dictSemiring_1_loop
_ = dictSemiring_1
one_2_0 := gopurs_runtime.RecordGet(dictSemiring_1, "one")
_ = one_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(c_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_1, "add"), one_2_0, c_3)
}), gopurs_runtime.RecordGet(dictSemiring_1, "zero"))
}

func Call_maximumBy(dictFoldable_0_loop gopurs_runtime.Value, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}).IntVal) != (0) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __t1})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_maximumBy__gopurs_runtime_Value_562089896(dictFoldable_0_loop gopurs_runtime.Value, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}).IntVal) != (0) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __t1})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_maximum(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy(dictFoldable_2, compare_1_0)
})
}

func Call_minimumBy(dictFoldable_0_loop gopurs_runtime.Value, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}).IntVal) != (0) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __t1})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_minimumBy__gopurs_runtime_Value_562089896(dictFoldable_0_loop gopurs_runtime.Value, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}).IntVal) != (0) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __t1})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_minimum(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy(dictFoldable_2, compare_1_0)
})
}

func Call_product(dictFoldable_0_loop gopurs_runtime.Value, dictSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 gopurs_runtime.Value = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.RecordGet(dictSemiring_1, "mul"), gopurs_runtime.RecordGet(dictSemiring_1, "one"))
}

func Call_sum(dictFoldable_0_loop gopurs_runtime.Value, dictSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 gopurs_runtime.Value = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.RecordGet(dictSemiring_1, "add"), gopurs_runtime.RecordGet(dictSemiring_1, "zero"))
}

func Call_sum__gopurs_runtime_Value_2374138440(dictFoldable_0_loop gopurs_runtime.Value, dictSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 gopurs_runtime.Value = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.RecordGet(dictSemiring_1, "add"), gopurs_runtime.RecordGet(dictSemiring_1, "zero"))
}

func Call_foldMapDefaultR(dictFoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_3, x_4), acc_5)
}), mempty_2_0)
})
}

func Call_foldMapDefaultL(dictFoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_4, gopurs_runtime.Apply(f_3, x_5))
}), mempty_2_0)
})
}

func Call_foldMap(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldMap")
}

func Call_foldMap__gopurs_runtime_Value_1916751092(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldMap")
}

func Call_foldableApp(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, i_2, v_3)
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, i_2, v_3)
}))
}

func Call_foldableCompose(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2)
_ = foldMap4_3_0
foldMap5_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2)
_ = foldMap5_4_1
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap4_3_0, gopurs_runtime.Apply(foldMap5_4_1, f_5), v_6)
})
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2), i_3, v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2)
_ = __local_var_5_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_5_2, a_7, b_6)
}), i_3, v_4)
}))
}

func Call_foldableCoproduct(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2)
_ = foldMap4_3_0
foldMap5_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2)
_ = foldMap5_4_1
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(foldMap4_3_0, f_5)
_ = __local_var_6_2
__local_var_7_3 := gopurs_runtime.Apply(foldMap5_4_1, f_5)
_ = __local_var_7_3
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(__local_var_6_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t4 = gopurs_runtime.Apply(__local_var_7_3, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
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
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3)
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, z_3)
_ = __local_var_5_6
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(__local_var_4_5, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(__local_var_5_6, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
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
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, z_3)
_ = __local_var_4_8
__local_var_5_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(__local_var_4_8, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t10 = gopurs_runtime.Apply(__local_var_5_9, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
}))
}

func Call_foldableProduct(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2)
_ = foldMap4_3_0
foldMap5_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2)
_ = foldMap5_4_1
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(foldMap4_3_0, f_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(foldMap5_4_1, f_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
}))
}

func Call_foldlDefault(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidFreeMonoidTree())
_ = foldMap2_1_0
return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldl"), c_2, u_3, gopurs_runtime.Apply2(foldMap2_1_0, Get_Node(), xs_4))
})
}

func Call_foldrDefault(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidFreeMonoidTree())
_ = foldMap2_1_0
return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), c_2, u_3, gopurs_runtime.Apply2(foldMap2_1_0, Get_Node(), xs_4))
})
}

func Call_foldrDefault__gopurs_runtime_Value_2409537192(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidFreeMonoidTree())
_ = foldMap2_1_0
return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), c_2, u_3, gopurs_runtime.Apply2(foldMap2_1_0, Get_Node(), xs_4))
})
}

func Call_foldrDefault__gopurs_runtime_Value_3006709320(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidFreeMonoidTree())
_ = foldMap2_1_0
return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), c_2, u_3, gopurs_runtime.Apply2(foldMap2_1_0, Get_Node(), xs_4))
})
}

func Call_lookup(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), pkg_Data_Maybe_First.Get_monoidFirst())
_ = foldMap2_1_0
return gopurs_runtime.Func2(func(dictEq_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap2_1_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_2, "eq"), a_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}
end_branch_1:
return __t1
}))
})
}

func Call_surroundMap(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidEndo())
_ = foldMap2_1_0
return gopurs_runtime.Func4(func(dictSemigroup_2 gopurs_runtime.Value, d_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMap2_1_0, gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), d_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), gopurs_runtime.Apply(t_4, a_6), m_7))
}), f_5, d_3)
})
}

func Call_surroundMap__gopurs_runtime_Value_75699306(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidEndo())
_ = foldMap2_1_0
return gopurs_runtime.Func4(func(dictSemigroup_2 gopurs_runtime.Value, d_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMap2_1_0, gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), d_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), gopurs_runtime.Apply(t_4, a_6), m_7))
}), f_5, d_3)
})
}

func Call_surround(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
surroundMap1_1_0 := Call_surroundMap(dictFoldable_0)
_ = surroundMap1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
surroundMap2_3_1 := gopurs_runtime.Apply(surroundMap1_1_0, dictSemigroup_2)
_ = surroundMap2_3_1
return gopurs_runtime.Func(func(d_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(surroundMap2_3_1, d_4, Get_identity())
})
})
}

func Call_foldM(dictFoldable_0_loop gopurs_runtime.Value, dictMonad_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, b0_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonad_1 gopurs_runtime.Value = dictMonad_1_loop
_ = dictMonad_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var b0_3 gopurs_runtime.Value = b0_3_loop
_ = b0_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}), "bind"), b_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_6, a_5)
}))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), b0_3))
}

func Call_fold(dictFoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1, Get_identity())
}

func Call_fold__gopurs_runtime_Value_3625874517(dictFoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1, Get_identity())
}

func Call_findMap(dictFoldable_0_loop gopurs_runtime.Value, p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(p_1, v1_3)
goto end_branch_0
} else {

}
}
{
__t0 = v_2
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_find(dictFoldable_0_loop gopurs_runtime.Value, p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil)) && ((gopurs_runtime.Apply(p_1, v1_3).IntVal) != (0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}
goto end_branch_0
} else {

}
}
{
__t0 = v_2
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_any(dictFoldable_0_loop gopurs_runtime.Value, dictHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 gopurs_runtime.Value = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "disj"), v_2, v1_3)
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "ff")))
}

func Call_any__gopurs_runtime_Value_968068125(dictFoldable_0_loop gopurs_runtime.Value, dictHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 gopurs_runtime.Value = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "disj"), v_2, v1_3)
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "ff")))
}

func Call_elem(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
semigroupDisj1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_1, v1_2)
}))
_ = semigroupDisj1_1_1
any1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_1
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")))
_ = any1_1_0
return gopurs_runtime.Func2(func(dictEq_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq_2, "eq"), x_3))
})
}

func Call_elem__gopurs_runtime_Value_2313606065(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
semigroupDisj1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_1, v1_2)
}))
_ = semigroupDisj1_1_1
any1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_1
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")))
_ = any1_1_0
return gopurs_runtime.Func2(func(dictEq_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq_2, "eq"), x_3))
})
}

func Call_notElem(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
elem1_1_0 := Call_elem(dictFoldable_0)
_ = elem1_1_0
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
elem2_3_1 := gopurs_runtime.Apply(elem1_1_0, dictEq_2)
_ = elem2_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(elem2_3_1, x_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(__local_var_5_2, x_6))
})
})
})
}

func Call_or(dictFoldable_0_loop gopurs_runtime.Value, dictHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 gopurs_runtime.Value = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "disj"), v_2, v1_3)
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "ff")), Get_identity())
}

func Call_all(dictFoldable_0_loop gopurs_runtime.Value, dictHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 gopurs_runtime.Value = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "conj"), v_2, v1_3)
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "tt")))
}

func Call_all__gopurs_runtime_Value_968068125(dictFoldable_0_loop gopurs_runtime.Value, dictHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 gopurs_runtime.Value = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "conj"), v_2, v1_3)
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "tt")))
}

func Call_and(dictFoldable_0_loop gopurs_runtime.Value, dictHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 gopurs_runtime.Value = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "conj"), v_2, v1_3)
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "tt")), Get_identity())
}

func Get_foldlArray() gopurs_runtime.Value {
	return _Gopurs_FoldlArray
}

func Get_foldrArray() gopurs_runtime.Value {
	return _Gopurs_FoldrArray
}
