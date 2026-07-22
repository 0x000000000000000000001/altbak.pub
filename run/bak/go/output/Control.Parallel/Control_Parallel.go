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
		parTraverse_ = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
traverse__2_0 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_traverse_(), dictApplicative_1)
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_4_1 := gopurs_runtime.Apply(traverse__2_0, dictFoldable_3)
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(traverse_1_4_1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], gopurs_runtime.Apply(f_5, x_6))
}))
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], gopurs_runtime.Apply(__local_var_6_2, x_7))
})
})
})
})
})
	})
	return parTraverse_
}

var parTraverse gopurs_runtime.Value
var once_parTraverse sync.Once
func Get_parTraverse() gopurs_runtime.Value {
	once_parTraverse.Do(func() {
		parTraverse = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTraversable_2 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_3_0 := gopurs_runtime.Apply(dictTraversable_2.PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_1)
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(traverse_3_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], gopurs_runtime.Apply(f_4, x_5))
}))
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], gopurs_runtime.Apply(__local_var_5_1, x_6))
})
})
})
})
})
	})
	return parTraverse
}

var parSequence_ gopurs_runtime.Value
var once_parSequence_ sync.Once
func Get_parSequence_() gopurs_runtime.Value {
	once_parSequence_.Do(func() {
		parSequence_ = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
parTraverse_2_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_parTraverse_(), dictParallel_0), dictApplicative_1)
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(parTraverse_2_2_0, dictFoldable_3), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
})
	})
	return parSequence_
}

var parSequence gopurs_runtime.Value
var once_parSequence sync.Once
func Get_parSequence() gopurs_runtime.Value {
	once_parSequence.Do(func() {
		parSequence = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTraversable_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictTraversable_2.PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], x_3)
}))
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], gopurs_runtime.Apply(__local_var_3_0, x_4))
})
})
})
})
	})
	return parSequence
}

var parOneOfMap gopurs_runtime.Value
var once_parOneOfMap sync.Once
func Get_parOneOfMap() gopurs_runtime.Value {
	once_parOneOfMap.Do(func() {
		parOneOfMap = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictAlternative_1 gopurs_runtime.Value) gopurs_runtime.Value {
Plus1_2_0 := gopurs_runtime.Apply(dictAlternative_1.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
empty_4_1 := Plus1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
return gopurs_runtime.Func(func(dictFunctor_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_3.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Plus1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], gopurs_runtime.Apply(f_6, x_7)))
})), empty_4_1)
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], gopurs_runtime.Apply(__local_var_7_2, x_8))
})
})
})
})
})
})
	})
	return parOneOfMap
}

var parOneOf gopurs_runtime.Value
var once_parOneOf sync.Once
func Get_parOneOf() gopurs_runtime.Value {
	once_parOneOf.Do(func() {
		parOneOf = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictAlternative_1 gopurs_runtime.Value) gopurs_runtime.Value {
Plus1_2_0 := gopurs_runtime.Apply(dictAlternative_1.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
empty_4_1 := Plus1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
return gopurs_runtime.Func(func(dictFunctor_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_3.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Plus1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], x_6))
})), empty_4_1)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], gopurs_runtime.Apply(__local_var_6_2, x_7))
})
})
})
})
})
	})
	return parOneOf
}

var parApply gopurs_runtime.Value
var once_parApply sync.Once
func Get_parApply() gopurs_runtime.Value {
	once_parApply.Do(func() {
		parApply = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mf_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ma_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], mf_1)), gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], ma_2)))
})
})
})
	})
	return parApply
}


