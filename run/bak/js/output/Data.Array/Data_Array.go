package Data_Array

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Array_ST_Iterator "gopurs/output/Data.Array.ST.Iterator"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Control_Bind "gopurs/output/Control.Bind"
)

var traverse_ gopurs_runtime.Value
var once_traverse_ sync.Once
func Get_traverse_() gopurs_runtime.Value {
	once_traverse_.Do(func() {
		traverse_ = gopurs_runtime.Apply(pkg_Data_Foldable.Get_traverse_(), pkg_Control_Monad_ST_Internal.Get_applicativeST())
	})
	return traverse_
}

var intercalate1 gopurs_runtime.Value
var once_intercalate1 sync.Once
func Get_intercalate1() gopurs_runtime.Value {
	once_intercalate1.Do(func() {
		intercalate1 = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func2(func(sep_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(pkg_Data_Foldable.Get_foldlArray(), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet(v_5, "init").IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Boolean(false), v1_6)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Boolean(false), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), sep_3, v1_6)))
}
end_branch_2:
return __t2
}), gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Boolean(true), mempty_2_1), xs_4), "acc")
})
})
	})
	return intercalate1
}

var zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		zipWith = gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_zipWithImpl(__local_var_0, __local_var_1, __local_var_2)
})
	})
	return zipWith
}

var zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		zipWithA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "traverse"), dictApplicative_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
_ = sequence1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_1_0, pkg_Data_Array.Call_zipWithImpl(f_2, xs_3, ys_4))
})
})
	})
	return zipWithA
}

var zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		zip = gopurs_runtime.Apply(Get_zipWith(), pkg_Data_Tuple.Get_Tuple())
	})
	return zip
}

var updateAtIndices gopurs_runtime.Value
var once_updateAtIndices sync.Once
func Get_updateAtIndices() gopurs_runtime.Value {
	once_updateAtIndices.Do(func() {
		updateAtIndices = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictFoldable_0)
_ = traverse_1_1_0
return gopurs_runtime.Func2(func(us_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Func(func(res_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(traverse_1_1_0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Data_Array_ST.Get_poke(), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], res_4)
}), us_2)
}), xs_3))
})
})
	})
	return updateAtIndices
}

var updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		updateAt = gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call__updateAt(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), __local_var_0, __local_var_1, __local_var_2)
})
	})
	return updateAt
}

var unsafeIndex gopurs_runtime.Value
var once_unsafeIndex sync.Once
func Get_unsafeIndex() gopurs_runtime.Value {
	once_unsafeIndex.Do(func() {
		unsafeIndex = gopurs_runtime.Func3(func(_dollar__unused_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ArrayAccess(__local_var_1, int(__local_var_2.IntVal))
})
	})
	return unsafeIndex
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(__local_var_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_unconsImpl(gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("head", "tail", x_1, xs_2))
}), __local_var_0)
})
	})
	return uncons
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func2(func(dictUnfoldable_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
len_2_0 := gopurs_runtime.Int(int64(len(xs_1.PtrVal.([]gopurs_runtime.Value))))
_ = len_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if i_3.IntVal < len_2_0.IntVal {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.ArrayAccess(xs_1, int(i_3.IntVal)), gopurs_runtime.Int(i_3.IntVal + 1)))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Int(0))
})
	})
	return toUnfoldable
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(__local_var_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_unconsImpl(gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Just", xs_2)
}), __local_var_0)
})
	})
	return tail
}

var sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		sortBy = gopurs_runtime.Func2(func(comp_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_sortByImpl(comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "GT").IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "EQ").IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "LT").IntVal != 0 {
__t0 = gopurs_runtime.Int(-1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), __local_var_1)
})
	})
	return sortBy
}

var sortWith gopurs_runtime.Value
var once_sortWith sync.Once
func Get_sortWith() gopurs_runtime.Value {
	once_sortWith.Do(func() {
		sortWith = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3))
}))
})
	})
	return sortWith
}

var sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		sort = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_sortBy(), compare_1_0, xs_2)
})
})
	})
	return sort
}

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), x_1), xs_0))
})
	})
	return snoc
}

var slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		slice = gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_sliceImpl(__local_var_0, __local_var_1, __local_var_2)
})
	})
	return slice
}

var splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		splitAt = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if v_0.IntVal <= 0 {
__t0 = gopurs_runtime.RecordDict2("before", "after", gopurs_runtime.Array([]gopurs_runtime.Value{}), v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("before", "after", pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(0), v_0, v1_1), pkg_Data_Array.Call_sliceImpl(v_0, gopurs_runtime.Int(int64(len(v1_1.PtrVal.([]gopurs_runtime.Value)))), v1_1))
}
end_branch_0:
return __t0
})
	})
	return splitAt
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if n_0.IntVal < 1 {
__t0 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_0
} else {

}
}
{
__t0 = pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(0), n_0, xs_1)
}
end_branch_0:
return __t0
})
	})
	return take
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{a_0})
})
	})
	return singleton
}

var scanr gopurs_runtime.Value
var once_scanr sync.Once
func Get_scanr() gopurs_runtime.Value {
	once_scanr.Do(func() {
		scanr = gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_scanrImpl(__local_var_0, __local_var_1, __local_var_2)
})
	})
	return scanr
}

var scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		scanl = gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_scanlImpl(__local_var_0, __local_var_1, __local_var_2)
})
	})
	return scanl
}

var replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		replicate = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_replicateImpl(__local_var_0, __local_var_1)
})
	})
	return replicate
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_rangeImpl(__local_var_0, __local_var_1)
})
	})
	return range_
}

var partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		partition = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_partitionImpl(__local_var_0, __local_var_1)
})
	})
	return partition
}

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(gopurs_runtime.Int(int64(len(xs_0.PtrVal.([]gopurs_runtime.Value)))).IntVal == 0)
})
	})
	return null
}

var modifyAtIndices gopurs_runtime.Value
var once_modifyAtIndices sync.Once
func Get_modifyAtIndices() gopurs_runtime.Value {
	once_modifyAtIndices.Do(func() {
		modifyAtIndices = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictFoldable_0)
_ = traverse_1_1_0
return gopurs_runtime.Func3(func(is_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Func(func(res_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(traverse_1_1_0, gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Data_Array_ST.Get_modify(), i_6, f_3, res_5)
}), is_2)
}), xs_4))
})
})
	})
	return modifyAtIndices
}

var mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		mapWithIndex = pkg_Data_FunctorWithIndex.Get_mapWithIndexArray()
	})
	return mapWithIndex
}

var intersperse gopurs_runtime.Value
var once_intersperse sync.Once
func Get_intersperse() gopurs_runtime.Value {
	once_intersperse.Do(func() {
		intersperse = gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, arr_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Int(int64(len(arr_1.PtrVal.([]gopurs_runtime.Value))))
_ = v_2_0
var __t7 gopurs_runtime.Value
{
if v_2_0.IntVal < 2 {
__t7 = arr_1
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
out_3_1 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Value{})
_ = out_3_1
_dollar__unused_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), gopurs_runtime.ArrayAccess(arr_1, 0), out_3_1), gopurs_runtime.Value{})
_ = _dollar__unused_4_2
_dollar__unused_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply3(pkg_Control_Monad_ST_Internal.Get_for_(), gopurs_runtime.Int(1), v_2_0, gopurs_runtime.Func(func(idx_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), a_0, out_3_1)
_ = __local_var_6_4
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_7_5 := gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Value{})
_ = _dollar__unused_7_5
__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), gopurs_runtime.ArrayAccess(arr_1, int(idx_5.IntVal)), out_3_1), gopurs_runtime.Value{})
_ = __local_var_8_6
return pkg_Data_Unit.Get_unit()
})
})), gopurs_runtime.Value{})
_ = _dollar__unused_5_3
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), out_3_1), gopurs_runtime.Value{})
}))
}
end_branch_7:
return __t7
})
	})
	return intersperse
}

var intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		intercalate = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_intercalate1(), dictMonoid_0)
})
	})
	return intercalate
}

var insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		insertAt = gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call__insertAt(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), __local_var_0, __local_var_1, __local_var_2)
})
	})
	return insertAt
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Int(int64(len(xs_0.PtrVal.([]gopurs_runtime.Value)))).IntVal == 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0.PtrVal.([]gopurs_runtime.Value)))).IntVal - 1), xs_0))
}
end_branch_0:
return __t0
})
	})
	return init_
}

var index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		index = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), __local_var_0, __local_var_1)
})
	})
	return index
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), xs_0, gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0.PtrVal.([]gopurs_runtime.Value)))).IntVal - 1))
})
	})
	return last
}

var unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		unsnoc = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Int(int64(len(xs_0.PtrVal.([]gopurs_runtime.Value)))).IntVal == 0 {
__t2 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_2
} else {

}
}
{
__local_var_1_0 := pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), xs_0, gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0.PtrVal.([]gopurs_runtime.Value)))).IntVal - 1))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_1_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("init", "last", pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0.PtrVal.([]gopurs_runtime.Value)))).IntVal - 1), xs_0), (*[1024]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return __t2
})
	})
	return unsnoc
}

var modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		modifyAt = gopurs_runtime.Func3(func(i_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), xs_2, i_0)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = pkg_Data_Array.Call__updateAt(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), i_0, gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]), xs_2)
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
	})
	return modifyAt
}

var span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		span = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, arr_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(i_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var i_3 gopurs_runtime.Value = i_3_loop
_ = i_3
v_4_1 := pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), arr_1, i_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4_1.StrVal == "Just").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Apply(p_0, (*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0]).IntVal != 0 {
i_3_loop = i_3.IntVal + 1
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor1("Just", i_3)
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_4_1.StrVal == "Nothing").IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Nothing")
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
breakIndex_3_4 := gopurs_runtime.Apply(go__2_0, gopurs_runtime.Int(0))
_ = breakIndex_3_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(breakIndex_3_4.StrVal == "Just").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (*[1024]gopurs_runtime.Value)(breakIndex_3_4.UnsafePtr)[0].IntVal == 0 {
__t6 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array([]gopurs_runtime.Value{}), arr_1)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.RecordDict2("init", "rest", pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(0), (*[1024]gopurs_runtime.Value)(breakIndex_3_4.UnsafePtr)[0], arr_1), pkg_Data_Array.Call_sliceImpl((*[1024]gopurs_runtime.Value)(breakIndex_3_4.UnsafePtr)[0], gopurs_runtime.Int(int64(len(arr_1.PtrVal.([]gopurs_runtime.Value)))), arr_1))
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(breakIndex_3_4.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.RecordDict2("init", "rest", arr_1, gopurs_runtime.Array([]gopurs_runtime.Value{}))
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
	return span
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_span(), p_0, xs_1), "init")
})
	})
	return takeWhile
}

var unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		unzip = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
fsts_1_0 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Value{})
_ = fsts_1_0
snds_2_1 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Value{})
_ = snds_2_1
__local_ref_3 := 0
_ = __local_ref_3
__local_var_3_2 := gopurs_runtime.Value{PtrVal: &__local_ref_3}
_ = __local_var_3_2
_dollar__unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST_Iterator.Get_iterate(), gopurs_runtime.Constructor2("Iterator", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), xs_0, v_4)
}), __local_var_3_2), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_5 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]
_ = __local_var_5_5
__local_var_6_6 := gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], fsts_1_0)
_ = __local_var_6_6
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_7 := gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Value{})
_ = __local_var_7_7
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), __local_var_5_5, snds_2_1), gopurs_runtime.Value{})
_ = __local_var_8_8
return pkg_Data_Unit.Get_unit()
})
})), gopurs_runtime.Value{})
_ = _dollar__unused_4_4
fsts_prime_5_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), fsts_1_0), gopurs_runtime.Value{})
_ = fsts_prime_5_9
snds_prime_6_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), snds_2_1), gopurs_runtime.Value{})
_ = snds_prime_6_10
return gopurs_runtime.Constructor2("Tuple", fsts_prime_5_9, snds_prime_6_10)
}))
})
	})
	return unzip
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), xs_0, gopurs_runtime.Int(0))
})
	})
	return head
}

var nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		nubBy = gopurs_runtime.Func2(func(comp_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
indexedAndSorted_2_0 := gopurs_runtime.Apply2(Get_sortBy(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(comp_0, (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[1])
}), gopurs_runtime.Apply2(pkg_Data_FunctorWithIndex.Get_mapWithIndexArray(), pkg_Data_Tuple.Get_Tuple(), xs_1))
_ = indexedAndSorted_2_0
v_3_1 := pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), indexedAndSorted_2_0, gopurs_runtime.Int(0))
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3_1.StrVal == "Nothing").IntVal != 0 {
__t2 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_3_1.StrVal == "Just").IntVal != 0 {
__local_var_4_3 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeThaw(), gopurs_runtime.Array([]gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(v_3_1.UnsafePtr)[0]}))
_ = __local_var_4_3
__t2 = gopurs_runtime.Apply2(pkg_Data_Functor.Get_arrayMap(), pkg_Data_Tuple.Get_snd(), gopurs_runtime.Apply3(Get_sortWith(), pkg_Data_Ord.Get_ordInt(), pkg_Data_Tuple.Get_fst(), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
result_5_4 := gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Value{})
_ = result_5_4
_dollar__unused_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_foreach(), indexedAndSorted_2_0, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_6 := (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]
_ = __local_var_7_6
__local_var_8_7 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), result_5_4)
_ = __local_var_8_7
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_8 := gopurs_runtime.Apply(__local_var_8_7, gopurs_runtime.Value{})
_ = __local_var_9_8
__local_var_10_10 := pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), __local_var_9_8, gopurs_runtime.Int(gopurs_runtime.Int(int64(len(__local_var_9_8.PtrVal.([]gopurs_runtime.Value)))).IntVal - 1))
_ = __local_var_10_10
var __t11 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_10_10.StrVal == "Just").IntVal != 0 {
__t11 = (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(__local_var_10_10.UnsafePtr)[0].UnsafePtr)[1]
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__local_var_10_9 := gopurs_runtime.Apply2(comp_0, __t11, __local_var_7_6)
_ = __local_var_10_9
__local_var_11_12 := gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), v1_6, result_5_4)
_ = __local_var_11_12
var __t13 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_10_9.StrVal == "LT").IntVal != 0 || gopurs_runtime.Bool(__local_var_10_9.StrVal == "GT").IntVal != 0 || gopurs_runtime.Bool(__local_var_10_9.StrVal == "EQ").IntVal != 0 != true {
__t13 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_12_14 := gopurs_runtime.Apply(__local_var_11_12, gopurs_runtime.Value{})
_ = __local_var_12_14
return pkg_Data_Unit.Get_unit()
})
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}
end_branch_13:
return gopurs_runtime.Apply(__t13, gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
_ = _dollar__unused_6_5
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), result_5_4), gopurs_runtime.Value{})
}))))
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
	})
	return nubBy
}

var nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		nub = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_nubBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
})
	})
	return nub
}

var groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		groupBy = gopurs_runtime.Func2(func(op_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
result_2_0 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Value{})
_ = result_2_0
__local_ref_2 := 0
_ = __local_ref_2
__local_var_3_1 := gopurs_runtime.Value{PtrVal: &__local_ref_2}
_ = __local_var_3_1
iter_4_3 := gopurs_runtime.Constructor2("Iterator", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), xs_1, v_4)
}), __local_var_3_1)
_ = iter_4_3
return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST_Iterator.Get_iterate(), iter_4_3, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
sub1_6_5 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Value{})
_ = sub1_6_5
_dollar__unused_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), x_5, sub1_6_5), gopurs_runtime.Value{})
_ = _dollar__unused_7_6
_dollar__unused_8_7 := gopurs_runtime.Apply(gopurs_runtime.Apply3(pkg_Data_Array_ST_Iterator.Get_pushWhile(), gopurs_runtime.Apply(op_0, x_5), iter_4_3, sub1_6_5), gopurs_runtime.Value{})
_ = _dollar__unused_8_7
grp_9_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), sub1_6_5), gopurs_runtime.Value{})
_ = grp_9_8
__local_var_10_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), grp_9_8, result_2_0), gopurs_runtime.Value{})
_ = __local_var_10_9
return pkg_Data_Unit.Get_unit()
})
})), gopurs_runtime.Value{})
_ = _dollar__unused_5_4
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), result_2_0), gopurs_runtime.Value{})
}), gopurs_runtime.Value{})
}))
})
	})
	return groupBy
}

var groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		groupAllBy = gopurs_runtime.Func(func(cmp_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(cmp_0, x_1, y_2).StrVal == "EQ")
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply2(Get_sortBy(), cmp_0, x_2))
})
})
	})
	return groupAllBy
}

var groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		groupAll = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_groupAllBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
})
	})
	return groupAll
}

var group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		group = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
eq2_1_0 := gopurs_runtime.RecordGet(dictEq_0, "eq")
_ = eq2_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_groupBy(), eq2_1_0, xs_2)
})
})
	})
	return group
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.RecordGet(dictFoldable_0, "foldr")
_ = __local_var_1_0
return gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_fromFoldableImpl(__local_var_1_0, __local_var_2)
})
})
	})
	return fromFoldable
}

var foldr gopurs_runtime.Value
var once_foldr sync.Once
func Get_foldr() gopurs_runtime.Value {
	once_foldr.Do(func() {
		foldr = pkg_Data_Foldable.Get_foldrArray()
	})
	return foldr
}

var foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		foldl = pkg_Data_Foldable.Get_foldlArray()
	})
	return foldl
}

var transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		transpose = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(idx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(allArrays_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var idx_2 gopurs_runtime.Value = idx_2_loop
_ = idx_2
var allArrays_3 gopurs_runtime.Value = allArrays_3_loop
_ = allArrays_3
v_4_1 := gopurs_runtime.Apply3(pkg_Data_Foldable.Get_foldlArray(), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, nextArr_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), nextArr_5, idx_2)
_ = __local_var_6_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_2.StrVal == "Nothing").IntVal != 0 {
__t3 = acc_4
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_6_2.StrVal == "Just").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(acc_4.StrVal == "Nothing").IntVal != 0 {
__t4 = gopurs_runtime.Array([]gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(__local_var_6_2.UnsafePtr)[0]})
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(acc_4.StrVal == "Just").IntVal != 0 {
__t4 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), (*[1024]gopurs_runtime.Value)(__local_var_6_2.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(acc_4.UnsafePtr)[0]))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = gopurs_runtime.Constructor1("Just", __t4)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.Constructor0("Nothing"), xs_0)
_ = v_4_1
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4_1.StrVal == "Nothing").IntVal != 0 {
__t5 = allArrays_3
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_4_1.StrVal == "Just").IntVal != 0 {
idx_2_loop = idx_2.IntVal + 1
allArrays_3_loop = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), (*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0]), allArrays_3))
continue go__1_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Int(0), gopurs_runtime.Array([]gopurs_runtime.Value{}))
})
	})
	return transpose
}

var foldRecM gopurs_runtime.Value
var once_foldRecM sync.Once
func Get_foldRecM() gopurs_runtime.Value {
	once_foldRecM.Do(func() {
		foldRecM = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, array_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet(o_6, "b").IntVal >= gopurs_runtime.Int(int64(len(array_5.PtrVal.([]gopurs_runtime.Value)))).IntVal {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Constructor1("Done", gopurs_runtime.RecordGet(o_6, "a")))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet(o_6, "a"), gopurs_runtime.ArrayAccess(array_5, int(gopurs_runtime.RecordGet(o_6, "b").IntVal))), gopurs_runtime.Func(func(res_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Constructor1("Loop", gopurs_runtime.RecordDict2("a", "b", res_prime_7, gopurs_runtime.Int(gopurs_runtime.RecordGet(o_6, "b").IntVal + 1))))
}))
}
end_branch_2:
return __t2
}), gopurs_runtime.RecordDict2("a", "b", b_4, gopurs_runtime.Int(0)))
})
})
	})
	return foldRecM
}

var foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		foldMap = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), dictMonoid_0)
})
	})
	return foldMap
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func4(Call_foldM)
	})
	return foldM
}

var fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		fold = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), dictMonoid_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return fold
}

var findMap gopurs_runtime.Value
var once_findMap sync.Once
func Get_findMap() gopurs_runtime.Value {
	once_findMap.Do(func() {
		findMap = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_findMapImpl(gopurs_runtime.Constructor0("Nothing"), pkg_Data_Maybe.Get_isJust(), __local_var_0, __local_var_1)
})
	})
	return findMap
}

var findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		findLastIndex = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_findLastIndexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), __local_var_0, __local_var_1)
})
	})
	return findLastIndex
}

var insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		insertBy = gopurs_runtime.Func3(func(cmp_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := pkg_Data_Array.Call_findLastIndexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(cmp_0, x_1, y_3).StrVal == "GT")
}), ys_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_1.StrVal == "Nothing").IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_1.StrVal == "Just").IntVal != 0 {
__t2 = gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(__local_var_3_1.UnsafePtr)[0].IntVal + 1)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__local_var_3_0 := pkg_Data_Array.Call__insertAt(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), __t2, x_1, ys_2)
_ = __local_var_3_0
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t3 = (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]
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
	return insertBy
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_insertBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
})
	})
	return insert
}

var findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		findIndex = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_findIndexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), __local_var_0, __local_var_1)
})
	})
	return findIndex
}

var find gopurs_runtime.Value
var once_find sync.Once
func Get_find() gopurs_runtime.Value {
	once_find.Do(func() {
		find = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := pkg_Data_Array.Call_findIndexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), f_0, xs_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.ArrayAccess(xs_1, int((*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0].IntVal)))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
})
	})
	return find
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_filterImpl(__local_var_0, __local_var_1)
})
	})
	return filter
}

var intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		intersectBy = gopurs_runtime.Func3(func(eq2_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_filterImpl(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := pkg_Data_Array.Call_findIndexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Apply(eq2_0, x_3), ys_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_4_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_4_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), xs_1)
})
	})
	return intersectBy
}

var intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		intersect = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_intersectBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return intersect
}

var elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		elemLastIndex = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1)
}))
})
	})
	return elemLastIndex
}

var elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		elemIndex = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1)
}))
})
	})
	return elemIndex
}

var notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		notElem = gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := pkg_Data_Array.Call_findIndexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_3, a_1)
}), arr_2)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(true)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(false)
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
	})
	return notElem
}

var elem gopurs_runtime.Value
var once_elem sync.Once
func Get_elem() gopurs_runtime.Value {
	once_elem.Do(func() {
		elem = gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := pkg_Data_Array.Call_findIndexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_3, a_1)
}), arr_2)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(true)
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
	})
	return elem
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_span(), p_0, xs_1), "rest")
})
	})
	return dropWhile
}

var dropEnd gopurs_runtime.Value
var once_dropEnd sync.Once
func Get_dropEnd() gopurs_runtime.Value {
	once_dropEnd.Do(func() {
		dropEnd = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Int(int64(len(xs_1.PtrVal.([]gopurs_runtime.Value)))).IntVal - n_0.IntVal
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if __local_var_2_0 < 1 {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(0), gopurs_runtime.Int(__local_var_2_0), xs_1)
}
end_branch_1:
return __t1
})
	})
	return dropEnd
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if n_0.IntVal < 1 {
__t0 = xs_1
goto end_branch_0
} else {

}
}
{
__t0 = pkg_Data_Array.Call_sliceImpl(n_0, gopurs_runtime.Int(int64(len(xs_1.PtrVal.([]gopurs_runtime.Value)))), xs_1)
}
end_branch_0:
return __t0
})
	})
	return drop
}

var takeEnd gopurs_runtime.Value
var once_takeEnd sync.Once
func Get_takeEnd() gopurs_runtime.Value {
	once_takeEnd.Do(func() {
		takeEnd = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Int(int64(len(xs_1.PtrVal.([]gopurs_runtime.Value)))).IntVal - n_0.IntVal
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if __local_var_2_0 < 1 {
__t1 = xs_1
goto end_branch_1
} else {

}
}
{
__t1 = pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(int64(len(xs_1.PtrVal.([]gopurs_runtime.Value)))), xs_1)
}
end_branch_1:
return __t1
})
	})
	return takeEnd
}

var deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		deleteAt = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call__deleteAt(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), __local_var_0, __local_var_1)
})
	})
	return deleteAt
}

var deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		deleteBy = gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Int(int64(len(v2_2.PtrVal.([]gopurs_runtime.Value)))).IntVal == 0 {
__t4 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__local_var_3_0 := pkg_Data_Array.Call_findIndexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Apply(v_0, v1_1), v2_2)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nothing").IntVal != 0 {
__t1 = v2_2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__local_var_4_2 := pkg_Data_Array.Call__deleteAt(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], v2_2)
_ = __local_var_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_4_2.StrVal == "Just").IntVal != 0 {
__t3 = (*[1024]gopurs_runtime.Value)(__local_var_4_2.UnsafePtr)[0]
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
__t4 = __t1
}
end_branch_4:
return __t4
})
	})
	return deleteBy
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return delete_
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldrArray(), gopurs_runtime.Apply(Get_delete_(), dictEq_0))
})
	})
	return difference
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), xs_1)
})
	})
	return cons
}

var some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		some = gopurs_runtime.Func3(func(dictAlternative_0 gopurs_runtime.Value, dictLazy_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_cons(), v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_1, "defer"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_many(), dictAlternative_0, dictLazy_1, v_2)
})))
})
	})
	return some
}

var many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		many = gopurs_runtime.Func3(func(dictAlternative_0 gopurs_runtime.Value, dictLazy_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply3(Get_some(), dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Array([]gopurs_runtime.Value{})))
})
	})
	return many
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func2(func(b_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Bind.Get_arrayBind(), a_1, b_0)
})
	})
	return concatMap
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0]})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
})
	})
	return mapMaybe
}

var filterA gopurs_runtime.Value
var once_filterA sync.Once
func Get_filterA() gopurs_runtime.Value {
	once_filterA.Do(func() {
		filterA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "traverse"), dictApplicative_0)
_ = traverse1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(traverse1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), x_4), gopurs_runtime.Apply(p_3, x_4))
}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(Get_mapMaybe(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1].IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_4:
return __t4
})))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(__local_var_4_2, x_6))
})
})
})
	})
	return filterA
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Apply(Get_mapMaybe(), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
	})
	return catMaybes
}

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_anyImpl(__local_var_0, __local_var_1)
})
	})
	return any
}

var nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		nubByEq = gopurs_runtime.Func2(func(eq2_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
arr_2_0 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Value{})
_ = arr_2_0
_dollar__unused_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_foreach(), xs_1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Get_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq2_0, v_4, x_3)
}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), arr_2_0)
_ = __local_var_5_3
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Value{})
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), x_3, arr_2_0)
_ = __local_var_7_5
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Apply(__local_var_4_2, __local_var_6_4).IntVal != 0 != true {
__t6 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_7 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
_ = __local_var_8_7
return pkg_Data_Unit.Get_unit()
})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}
end_branch_6:
return gopurs_runtime.Apply(__t6, gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
_ = _dollar__unused_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), arr_2_0), gopurs_runtime.Value{})
}))
})
	})
	return nubByEq
}

var nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		nubEq = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_nubByEq(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return nubEq
}

var unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		unionBy = gopurs_runtime.Func3(func(eq2_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), xs_1, gopurs_runtime.Apply3(pkg_Data_Foldable.Get_foldlArray(), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_deleteBy(), eq2_0, a_4, b_3)
}), gopurs_runtime.Apply2(Get_nubByEq(), eq2_0, ys_2), xs_1))
})
	})
	return unionBy
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unionBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return union
}

var alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		alterAt = gopurs_runtime.Func3(func(i_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := pkg_Data_Array.Call_indexImpl(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), xs_2, i_0)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
v_4_2 := gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0])
_ = v_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4_2.StrVal == "Nothing").IntVal != 0 {
__t3 = pkg_Data_Array.Call__deleteAt(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), i_0, xs_2)
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_4_2.StrVal == "Just").IntVal != 0 {
__t3 = pkg_Data_Array.Call__updateAt(pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), i_0, (*[1024]gopurs_runtime.Value)(v_4_2.UnsafePtr)[0], xs_2)
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
})
	})
	return alterAt
}

var all gopurs_runtime.Value
var once_all sync.Once
func Get_all() gopurs_runtime.Value {
	once_all.Do(func() {
		all = gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_allImpl(__local_var_0, __local_var_1)
})
	})
	return all
}

func Call_foldM(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
return pkg_Data_Array.Call_unconsImpl(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_2_loop)
}), gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, as_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1_loop, b_2_loop, a_4), gopurs_runtime.Func(func(b_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_foldM(), dictMonad_0_loop, f_1_loop, b_prime_6, as_5)
}))
}), __local_var_3_loop)
}
}

func Get__deleteAt() gopurs_runtime.Value {
	return _Gopurs__DeleteAt
}

func Get__insertAt() gopurs_runtime.Value {
	return _Gopurs__InsertAt
}

func Get__updateAt() gopurs_runtime.Value {
	return _Gopurs__UpdateAt
}

func Get_allImpl() gopurs_runtime.Value {
	return _Gopurs_AllImpl
}

func Get_anyImpl() gopurs_runtime.Value {
	return _Gopurs_AnyImpl
}

func Get_concat() gopurs_runtime.Value {
	return _Gopurs_Concat
}

func Get_filterImpl() gopurs_runtime.Value {
	return _Gopurs_FilterImpl
}

func Get_findIndexImpl() gopurs_runtime.Value {
	return _Gopurs_FindIndexImpl
}

func Get_findLastIndexImpl() gopurs_runtime.Value {
	return _Gopurs_FindLastIndexImpl
}

func Get_findMapImpl() gopurs_runtime.Value {
	return _Gopurs_FindMapImpl
}

func Get_fromFoldableImpl() gopurs_runtime.Value {
	return _Gopurs_FromFoldableImpl
}

func Get_indexImpl() gopurs_runtime.Value {
	return _Gopurs_IndexImpl
}

func Get_length() gopurs_runtime.Value {
	return _Gopurs_Length
}

func Get_partitionImpl() gopurs_runtime.Value {
	return _Gopurs_PartitionImpl
}

func Get_rangeImpl() gopurs_runtime.Value {
	return _Gopurs_RangeImpl
}

func Get_replicateImpl() gopurs_runtime.Value {
	return _Gopurs_ReplicateImpl
}

func Get_reverse() gopurs_runtime.Value {
	return _Gopurs_Reverse
}

func Get_scanlImpl() gopurs_runtime.Value {
	return _Gopurs_ScanlImpl
}

func Get_scanrImpl() gopurs_runtime.Value {
	return _Gopurs_ScanrImpl
}

func Get_sliceImpl() gopurs_runtime.Value {
	return _Gopurs_SliceImpl
}

func Get_sortByImpl() gopurs_runtime.Value {
	return _Gopurs_SortByImpl
}

func Get_unconsImpl() gopurs_runtime.Value {
	return _Gopurs_UnconsImpl
}

func Get_unsafeIndexImpl() gopurs_runtime.Value {
	return _Gopurs_UnsafeIndexImpl
}

func Get_zipWithImpl() gopurs_runtime.Value {
	return _Gopurs_ZipWithImpl
}
