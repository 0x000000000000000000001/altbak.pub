package Data_Semigroup_Foldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Unit "gopurs/output/Data.Unit"
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

var cache_FoldRight1 gopurs_runtime.Value
var once_FoldRight1 sync.Once
func Get_FoldRight1() gopurs_runtime.Value {
	once_FoldRight1.Do(func() {
		cache_FoldRight1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_FoldRight1[gopurs_runtime.Value]{value0, value1})}
})
})
	})
	return cache_FoldRight1
}

var cache_semigroupAct gopurs_runtime.Value
var once_semigroupAct sync.Once
func Get_semigroupAct() gopurs_runtime.Value {
	once_semigroupAct.Do(func() {
		cache_semigroupAct = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupAct(dictApply_0_box)
})
	})
	return cache_semigroupAct
}

var cache_mkFoldRight1 gopurs_runtime.Value
var once_mkFoldRight1 sync.Once
func Get_mkFoldRight1() gopurs_runtime.Value {
	once_mkFoldRight1.Do(func() {
		cache_mkFoldRight1 = gopurs_runtime.Apply(Get_FoldRight1(), pkg_Data_Function.Get_const_())
	})
	return cache_mkFoldRight1
}

var cache_foldr1 gopurs_runtime.Value
var once_foldr1 sync.Once
func Get_foldr1() gopurs_runtime.Value {
	once_foldr1.Do(func() {
		cache_foldr1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1(dict_0_box)
})
	})
	return cache_foldr1
}

var cache_foldl1 gopurs_runtime.Value
var once_foldl1 sync.Once
func Get_foldl1() gopurs_runtime.Value {
	once_foldl1.Do(func() {
		cache_foldl1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1(dict_0_box)
})
	})
	return cache_foldl1
}

var cache_maximumBy gopurs_runtime.Value
var once_maximumBy sync.Once
func Get_maximumBy() gopurs_runtime.Value {
	once_maximumBy.Do(func() {
		cache_maximumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy(dictFoldable1_0_box, cmp_1_box)
})
	})
	return cache_maximumBy
}

var cache_minimumBy gopurs_runtime.Value
var once_minimumBy sync.Once
func Get_minimumBy() gopurs_runtime.Value {
	once_minimumBy.Do(func() {
		cache_minimumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy(dictFoldable1_0_box, cmp_1_box)
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
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
}))
	})
	return cache_foldableTuple
}

var cache_foldableMultiplicative gopurs_runtime.Value
var once_foldableMultiplicative sync.Once
func Get_foldableMultiplicative() gopurs_runtime.Value {
	once_foldableMultiplicative.Do(func() {
		cache_foldableMultiplicative = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return cache_foldableMultiplicative
}

var cache_foldableIdentity gopurs_runtime.Value
var once_foldableIdentity sync.Once
func Get_foldableIdentity() gopurs_runtime.Value {
	once_foldableIdentity.Do(func() {
		cache_foldableIdentity = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return cache_foldableIdentity
}

var cache_foldableDual gopurs_runtime.Value
var once_foldableDual sync.Once
func Get_foldableDual() gopurs_runtime.Value {
	once_foldableDual.Do(func() {
		cache_foldableDual = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return cache_foldableDual
}

var cache_foldRight1Semigroup gopurs_runtime.Value
var once_foldRight1Semigroup sync.Once
func Get_foldRight1Semigroup() gopurs_runtime.Value {
	once_foldRight1Semigroup.Do(func() {
		cache_foldRight1Semigroup = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := (*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_FoldRight1[gopurs_runtime.Value]{gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, a_3, f_4)), f_4)
}), (*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V1})}
}))
	})
	return cache_foldRight1Semigroup
}

var cache_semigroupDual gopurs_runtime.Value
var once_semigroupDual sync.Once
func Get_semigroupDual() gopurs_runtime.Value {
	once_semigroupDual.Do(func() {
		cache_semigroupDual = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldRight1Semigroup(), "append"), v1_1, v_0)
}))
	})
	return cache_semigroupDual
}

var cache_foldMap1DefaultR gopurs_runtime.Value
var once_foldMap1DefaultR sync.Once
func Get_foldMap1DefaultR() gopurs_runtime.Value {
	once_foldMap1DefaultR.Do(func() {
		cache_foldMap1DefaultR = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1DefaultR(dictFoldable1_0_box, dictFunctor_1_box, dictSemigroup_2_box)
})
	})
	return cache_foldMap1DefaultR
}

var cache_foldMap1DefaultL gopurs_runtime.Value
var once_foldMap1DefaultL sync.Once
func Get_foldMap1DefaultL() gopurs_runtime.Value {
	once_foldMap1DefaultL.Do(func() {
		cache_foldMap1DefaultL = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1DefaultL(dictFoldable1_0_box, dictFunctor_1_box, dictSemigroup_2_box)
})
	})
	return cache_foldMap1DefaultL
}

var cache_foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		cache_foldMap1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1(dict_0_box)
})
	})
	return cache_foldMap1
}

var cache_foldl1Default gopurs_runtime.Value
var once_foldl1Default sync.Once
func Get_foldl1Default() gopurs_runtime.Value {
	once_foldl1Default.Do(func() {
		cache_foldl1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1Default(dictFoldable1_0_box)
})
	})
	return cache_foldl1Default
}

var cache_foldr1Default gopurs_runtime.Value
var once_foldr1Default sync.Once
func Get_foldr1Default() gopurs_runtime.Value {
	once_foldr1Default.Do(func() {
		cache_foldr1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1Default(dictFoldable1_0_box)
})
	})
	return cache_foldr1Default
}

var cache_intercalateMap gopurs_runtime.Value
var once_intercalateMap sync.Once
func Get_intercalateMap() gopurs_runtime.Value {
	once_intercalateMap.Do(func() {
		cache_intercalateMap = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalateMap(dictFoldable1_0_box, dictSemigroup_1_box)
})
	})
	return cache_intercalateMap
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate(dictFoldable1_0_box, dictSemigroup_1_box)
})
	})
	return cache_intercalate
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

var cache_traverse1_ gopurs_runtime.Value
var once_traverse1_ sync.Once
func Get_traverse1_() gopurs_runtime.Value {
	once_traverse1_.Do(func() {
		cache_traverse1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1_(dictFoldable1_0_box, dictApply_1_box)
})
	})
	return cache_traverse1_
}

var cache_for1_ gopurs_runtime.Value
var once_for1_ sync.Once
func Get_for1_() gopurs_runtime.Value {
	once_for1_.Do(func() {
		cache_for1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_for1_(dictFoldable1_0_box, dictApply_1_box)
})
	})
	return cache_for1_
}

var cache_sequence1_ gopurs_runtime.Value
var once_sequence1_ sync.Once
func Get_sequence1_() gopurs_runtime.Value {
	once_sequence1_.Do(func() {
		cache_sequence1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1_(dictFoldable1_0_box, dictApply_1_box)
})
	})
	return cache_sequence1_
}

var cache_fold1 gopurs_runtime.Value
var once_fold1 sync.Once
func Get_fold1() gopurs_runtime.Value {
	once_fold1.Do(func() {
		cache_fold1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold1(dictFoldable1_0_box, dictSemigroup_1_box)
})
	})
	return cache_fold1
}

type Constructor_FoldRight1[T_a any] struct {
	V0 gopurs_runtime.Value
	V1 T_a
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_semigroupAct(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_identity()
}), v_1), v1_2)
}))
}

func Call_foldr1(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldr1")
}

func Call_foldl1(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldl1")
}

func Call_maximumBy(dictFoldable1_0_loop gopurs_runtime.Value, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldl1"), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, x_2, y_3), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}).IntVal) != (0) {
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
}))
}

func Call_minimumBy(dictFoldable1_0_loop gopurs_runtime.Value, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldl1"), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, x_2, y_3), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}).IntVal) != (0) {
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
}))
}

func Call_foldMap1DefaultR(dictFoldable1_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
append_3_0 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1, "map"), f_4)
_ = __local_var_5_1
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldr1"), append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_foldMap1DefaultL(dictFoldable1_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
append_3_0 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1, "map"), f_4)
_ = __local_var_5_1
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldl1"), append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_foldMap1(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldMap1")
}

func Call_foldl1Default(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), Get_semigroupDual(), Get_mkFoldRight1())
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_1_0, a_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, (*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(x_2, a_6, b_5)
}))
})
}

func Call_foldr1Default(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), Get_foldRight1Semigroup(), Get_mkFoldRight1())
_ = __local_var_1_0
return gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_1_0, a_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, (*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, b_2)
})
}

func Call_intercalateMap(dictFoldable1_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
foldMap12_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), j_4, gopurs_runtime.Apply(v1_3, j_4)))
})))
_ = foldMap12_2_0
return gopurs_runtime.Func3(func(j_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, foldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMap12_2_0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_1 := gopurs_runtime.Apply(f_4, x_6)
_ = __local_var_7_1
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_1
})
}), foldable_5, j_3)
})
}

func Call_intercalate(dictFoldable1_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
foldMap12_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), j_4, gopurs_runtime.Apply(v1_3, j_4)))
})))
_ = foldMap12_2_0
return gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, foldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMap12_2_0, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
}), foldable_4, a_3)
})
}

func Call_maximum(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
semigroupMax_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
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
}))
_ = semigroupMax_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), semigroupMax_1_0, pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
}

func Call_minimum(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
semigroupMin_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
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
}))
_ = semigroupMin_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), semigroupMin_1_0, pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
}

func Call_traverse1_(dictFoldable1_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_0
foldMap12_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), gopurs_runtime.Apply(Get_semigroupAct(), dictApply_1))
_ = foldMap12_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Apply2(foldMap12_3_1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
}

func Call_for1_(dictFoldable1_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
__local_var_2_0 := Call_traverse1_(dictFoldable1_0, dictApply_1)
_ = __local_var_2_0
return gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
}

func Call_sequence1_(dictFoldable1_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
return gopurs_runtime.Apply(Call_traverse1_(dictFoldable1_0, dictApply_1), Get_identity())
}

func Call_fold1(dictFoldable1_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), dictSemigroup_1, Get_identity())
}


