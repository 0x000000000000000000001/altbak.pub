package Control_Parallel

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
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

var cache_parTraverse_ gopurs_runtime.Value
var once_parTraverse_ sync.Once
func Get_parTraverse_() gopurs_runtime.Value {
	once_parTraverse_.Do(func() {
		cache_parTraverse_ = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse_(dictParallel_0_box)
})
	})
	return cache_parTraverse_
}

var cache_parTraverse gopurs_runtime.Value
var once_parTraverse sync.Once
func Get_parTraverse() gopurs_runtime.Value {
	once_parTraverse.Do(func() {
		cache_parTraverse = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse(dictParallel_0_box)
})
	})
	return cache_parTraverse
}

var cache_parSequence_ gopurs_runtime.Value
var once_parSequence_ sync.Once
func Get_parSequence_() gopurs_runtime.Value {
	once_parSequence_.Do(func() {
		cache_parSequence_ = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parSequence_(dictParallel_0_box)
})
	})
	return cache_parSequence_
}

var cache_parSequence gopurs_runtime.Value
var once_parSequence sync.Once
func Get_parSequence() gopurs_runtime.Value {
	once_parSequence.Do(func() {
		cache_parSequence = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parSequence(dictParallel_0_box)
})
	})
	return cache_parSequence
}

var cache_parOneOfMap gopurs_runtime.Value
var once_parOneOfMap sync.Once
func Get_parOneOfMap() gopurs_runtime.Value {
	once_parOneOfMap.Do(func() {
		cache_parOneOfMap = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parOneOfMap(dictParallel_0_box)
})
	})
	return cache_parOneOfMap
}

var cache_parOneOf gopurs_runtime.Value
var once_parOneOf sync.Once
func Get_parOneOf() gopurs_runtime.Value {
	once_parOneOf.Do(func() {
		cache_parOneOf = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parOneOf(dictParallel_0_box)
})
	})
	return cache_parOneOf
}

var cache_parApply gopurs_runtime.Value
var once_parApply sync.Once
func Get_parApply() gopurs_runtime.Value {
	once_parApply.Do(func() {
		cache_parApply = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, mf_1_box gopurs_runtime.Value, ma_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parApply(dictParallel_0_box, mf_1_box, ma_2_box)
})
	})
	return cache_parApply
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_parTraverse_(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
sequential_1_0 := ((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V1
_ = sequential_1_0
parallel_2_1 := ((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V0
_ = parallel_2_1
return gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
traverse__4_2 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_traverse_(), dictApplicative_3)
_ = traverse__4_2
return gopurs_runtime.Func(func(dictFoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_6_3 := gopurs_runtime.Apply(traverse__4_2, dictFoldable_5)
_ = traverse_1_6_3
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), sequential_1_0, gopurs_runtime.Apply(traverse_1_6_3, gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), parallel_2_1, f_7)))
})
})
})
}

func Call_parTraverse(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
sequential_1_0 := ((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V1
_ = sequential_1_0
parallel_2_1 := ((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V0
_ = parallel_2_1
return gopurs_runtime.Func2(func(dictApplicative_3 gopurs_runtime.Value, dictTraversable_4 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_4, "traverse"), dictApplicative_3)
_ = traverse_5_2
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), sequential_1_0, gopurs_runtime.Apply(traverse_5_2, gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), parallel_2_1, f_6)))
})
})
}

func Call_parSequence_(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
parTraverse_1_1_0 := gopurs_runtime.Apply(Get_parTraverse_(), dictParallel_0)
_ = parTraverse_1_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
parTraverse_2_3_1 := gopurs_runtime.Apply(parTraverse_1_1_0, dictApplicative_2)
_ = parTraverse_2_3_1
return gopurs_runtime.Func(func(dictFoldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(parTraverse_2_3_1, dictFoldable_4, Get_identity())
})
})
}

func Call_parSequence(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
parTraverse1_1_0 := gopurs_runtime.Apply(Get_parTraverse(), dictParallel_0)
_ = parTraverse1_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
parTraverse2_3_1 := gopurs_runtime.Apply(parTraverse1_1_0, dictApplicative_2)
_ = parTraverse2_3_1
return gopurs_runtime.Func(func(dictTraversable_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(parTraverse2_3_1, dictTraversable_4, Get_identity())
})
})
}

func Call_parOneOfMap(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
sequential_1_0 := ((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V1
_ = sequential_1_0
parallel_2_1 := ((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V0
_ = parallel_2_1
return gopurs_runtime.Func(func(dictAlternative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Plus1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_3, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_4_2
return gopurs_runtime.Func(func(dictFoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
oneOfMap_6_3 := gopurs_runtime.Apply2(pkg_Data_Foldable.Get_oneOfMap(), dictFoldable_5, Plus1_4_2)
_ = oneOfMap_6_3
return gopurs_runtime.Func2(func(dictFunctor_7 gopurs_runtime.Value, f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), sequential_1_0, gopurs_runtime.Apply(oneOfMap_6_3, gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), parallel_2_1, f_8)))
})
})
})
}

func Call_parOneOf(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
sequential_1_0 := ((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V1
_ = sequential_1_0
parallel_2_1 := ((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V0
_ = parallel_2_1
return gopurs_runtime.Func(func(dictAlternative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Plus1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_3, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_4_2
return gopurs_runtime.Func(func(dictFoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
oneOfMap_6_3 := gopurs_runtime.Apply2(pkg_Data_Foldable.Get_oneOfMap(), dictFoldable_5, Plus1_4_2)
_ = oneOfMap_6_3
return gopurs_runtime.Func(func(dictFunctor_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), sequential_1_0, gopurs_runtime.Apply(oneOfMap_6_3, parallel_2_1))
})
})
})
}

func Call_parApply(dictParallel_0_loop gopurs_runtime.Value, mf_1_loop gopurs_runtime.Value, ma_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
var mf_1 gopurs_runtime.Value = mf_1_loop
_ = mf_1
var ma_2 gopurs_runtime.Value = ma_2_loop
_ = ma_2
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1_NOT_FOUND"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V0, mf_1), gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictParallel_0.UnsafePtr)).V0, ma_2)))
}


