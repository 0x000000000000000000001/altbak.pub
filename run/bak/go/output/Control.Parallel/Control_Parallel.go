package Control_Parallel

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var parTraverse_ gopurs_runtime.Value
var once_parTraverse_ sync.Once
func Get_parTraverse_() gopurs_runtime.Value {
	once_parTraverse_.Do(func() {
		parTraverse_ = gopurs_runtime.Func2(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse_(dictParallel_0_box, dictApplicative_1_box)
})
	})
	return parTraverse_
}

var parTraverse gopurs_runtime.Value
var once_parTraverse sync.Once
func Get_parTraverse() gopurs_runtime.Value {
	once_parTraverse.Do(func() {
		parTraverse = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse(dictParallel_0_box, dictApplicative_1_box, dictTraversable_2_box)
})
	})
	return parTraverse
}

var parSequence_ gopurs_runtime.Value
var once_parSequence_ sync.Once
func Get_parSequence_() gopurs_runtime.Value {
	once_parSequence_.Do(func() {
		parSequence_ = gopurs_runtime.Func2(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parSequence_(dictParallel_0_box, dictApplicative_1_box)
})
	})
	return parSequence_
}

var parSequence gopurs_runtime.Value
var once_parSequence sync.Once
func Get_parSequence() gopurs_runtime.Value {
	once_parSequence.Do(func() {
		parSequence = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parSequence(dictParallel_0_box, dictApplicative_1_box, dictTraversable_2_box)
})
	})
	return parSequence
}

var parOneOfMap gopurs_runtime.Value
var once_parOneOfMap sync.Once
func Get_parOneOfMap() gopurs_runtime.Value {
	once_parOneOfMap.Do(func() {
		parOneOfMap = gopurs_runtime.Func2(func(dictParallel_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parOneOfMap(dictParallel_0_box, dictAlternative_1_box)
})
	})
	return parOneOfMap
}

var parOneOf gopurs_runtime.Value
var once_parOneOf sync.Once
func Get_parOneOf() gopurs_runtime.Value {
	once_parOneOf.Do(func() {
		parOneOf = gopurs_runtime.Func2(func(dictParallel_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parOneOf(dictParallel_0_box, dictAlternative_1_box)
})
	})
	return parOneOf
}

var parApply gopurs_runtime.Value
var once_parApply sync.Once
func Get_parApply() gopurs_runtime.Value {
	once_parApply.Do(func() {
		parApply = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, mf_1_box gopurs_runtime.Value, ma_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parApply(dictParallel_0_box, mf_1_box, ma_2_box)
})
	})
	return parApply
}

func Call_parTraverse_(dictParallel_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
traverse__2_0 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_traverse_(), dictApplicative_1)
_ = traverse__2_0
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_4_1 := gopurs_runtime.Apply(traverse__2_0, dictFoldable_3)
_ = traverse_1_4_1
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(traverse_1_4_1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(f_5, x_6))
}))
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(__local_var_6_2, x_7))
})
})
})
}

func Call_parTraverse(dictParallel_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value, dictTraversable_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
var dictTraversable_2 gopurs_runtime.Value = dictTraversable_2_loop
_ = dictTraversable_2
traverse_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_2, "traverse"), dictApplicative_1)
_ = traverse_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(traverse_3_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(f_4, x_5))
}))
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(__local_var_5_1, x_6))
})
})
}

func Call_parSequence_(dictParallel_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
parTraverse_2_2_0 := Call_parTraverse_(dictParallel_0, dictApplicative_1)
_ = parTraverse_2_2_0
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(parTraverse_2_2_0, dictFoldable_3, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
}

func Call_parSequence(dictParallel_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value, dictTraversable_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
var dictTraversable_2 gopurs_runtime.Value = dictTraversable_2_loop
_ = dictTraversable_2
__local_var_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_2, "traverse"), dictApplicative_1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), x_3)
}))
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(__local_var_3_0, x_4))
})
}

func Call_parOneOfMap(dictParallel_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
Plus1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_2_0
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
empty_4_1 := gopurs_runtime.RecordGet(Plus1_2_0, "empty")
_ = empty_4_1
return gopurs_runtime.Func2(func(dictFunctor_5 gopurs_runtime.Value, f_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_3, "foldr"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_2_0, "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(f_6, x_7)))
}), empty_4_1)
_ = __local_var_7_2
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(__local_var_7_2, x_8))
})
})
})
}

func Call_parOneOf(dictParallel_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
Plus1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_2_0
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
empty_4_1 := gopurs_runtime.RecordGet(Plus1_2_0, "empty")
_ = empty_4_1
return gopurs_runtime.Func(func(dictFunctor_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_3, "foldr"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_2_0, "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), x_6))
}), empty_4_1)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(__local_var_6_2, x_7))
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
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), mf_1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), ma_2)))
}


