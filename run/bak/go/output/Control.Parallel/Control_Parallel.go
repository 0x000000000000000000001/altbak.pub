package Control_Parallel

import (
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Parallel_Class "gopurs/output/Control.Parallel.Class"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
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

var cache_parTraverse_ gopurs_runtime.Value
var once_parTraverse_ sync.Once
func Get_parTraverse_() gopurs_runtime.Value {
	once_parTraverse_.Do(func() {
		cache_parTraverse_ = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse_(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2_box), f_3_box)
})
	})
	return cache_parTraverse_
}

var cache_parTraverse___gopurs_runtime_Value_1426351978 gopurs_runtime.Value
var once_parTraverse___gopurs_runtime_Value_1426351978 sync.Once
func Get_parTraverse___gopurs_runtime_Value_1426351978() gopurs_runtime.Value {
	once_parTraverse___gopurs_runtime_Value_1426351978.Do(func() {
		cache_parTraverse___gopurs_runtime_Value_1426351978 = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse___gopurs_runtime_Value_1426351978(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2_box), f_3_box)
})
	})
	return cache_parTraverse___gopurs_runtime_Value_1426351978
}

var cache_parTraverse gopurs_runtime.Value
var once_parTraverse sync.Once
func Get_parTraverse() gopurs_runtime.Value {
	once_parTraverse.Do(func() {
		cache_parTraverse = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_2_box), f_3_box)
})
	})
	return cache_parTraverse
}

var cache_parTraverse__gopurs_runtime_Value_1055730709 gopurs_runtime.Value
var once_parTraverse__gopurs_runtime_Value_1055730709 sync.Once
func Get_parTraverse__gopurs_runtime_Value_1055730709() gopurs_runtime.Value {
	once_parTraverse__gopurs_runtime_Value_1055730709.Do(func() {
		cache_parTraverse__gopurs_runtime_Value_1055730709 = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse__gopurs_runtime_Value_1055730709(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_2_box), f_3_box)
})
	})
	return cache_parTraverse__gopurs_runtime_Value_1055730709
}

var cache_parSequence_ gopurs_runtime.Value
var once_parSequence_ sync.Once
func Get_parSequence_() gopurs_runtime.Value {
	once_parSequence_.Do(func() {
		cache_parSequence_ = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parSequence_(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2_box))
})
	})
	return cache_parSequence_
}

var cache_parSequence___gopurs_runtime_Value_1071252918 gopurs_runtime.Value
var once_parSequence___gopurs_runtime_Value_1071252918 sync.Once
func Get_parSequence___gopurs_runtime_Value_1071252918() gopurs_runtime.Value {
	once_parSequence___gopurs_runtime_Value_1071252918.Do(func() {
		cache_parSequence___gopurs_runtime_Value_1071252918 = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parSequence___gopurs_runtime_Value_1071252918(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2_box))
})
	})
	return cache_parSequence___gopurs_runtime_Value_1071252918
}

var cache_parSequence gopurs_runtime.Value
var once_parSequence sync.Once
func Get_parSequence() gopurs_runtime.Value {
	once_parSequence.Do(func() {
		cache_parSequence = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parSequence(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_2_box))
})
	})
	return cache_parSequence
}

var cache_parOneOfMap gopurs_runtime.Value
var once_parOneOfMap sync.Once
func Get_parOneOfMap() gopurs_runtime.Value {
	once_parOneOfMap.Do(func() {
		cache_parOneOfMap = gopurs_runtime.Func2(func(dictParallel_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parOneOfMap(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_1_box))
})
	})
	return cache_parOneOfMap
}

var cache_parOneOf gopurs_runtime.Value
var once_parOneOf sync.Once
func Get_parOneOf() gopurs_runtime.Value {
	once_parOneOf.Do(func() {
		cache_parOneOf = gopurs_runtime.Func2(func(dictParallel_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parOneOf(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_1_box))
})
	})
	return cache_parOneOf
}

var cache_parApply gopurs_runtime.Value
var once_parApply sync.Once
func Get_parApply() gopurs_runtime.Value {
	once_parApply.Do(func() {
		cache_parApply = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parApply(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
})
	})
	return cache_parApply
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_parTraverse_(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictFoldable_2 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_2_loop
_ = dictFoldable_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
__local_var_4_0 := gopurs_runtime.Apply3(pkg_Data_Foldable.Get_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_2)}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V2, gopurs_runtime.Apply(f_3, x_4))
}))
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_4_0, x_5))
})
}

func Call_parTraverse___gopurs_runtime_Value_1426351978(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictFoldable_2 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_2_loop
_ = dictFoldable_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
__local_var_4_0 := gopurs_runtime.Apply3(pkg_Data_Foldable.Get_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_2)}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V2, gopurs_runtime.Apply(f_3, x_4))
}))
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_4_0, x_5))
})
}

func Call_parTraverse(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictTraversable_2_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictTraversable_2 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
_ = dictTraversable_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
__local_var_4_0 := gopurs_runtime.Apply2(dictTraversable_2.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V2, gopurs_runtime.Apply(f_3, x_4))
}))
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_4_0, x_5))
})
}

func Call_parTraverse__gopurs_runtime_Value_1055730709(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictTraversable_2_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictTraversable_2 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
_ = dictTraversable_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
__local_var_4_0 := gopurs_runtime.Apply2(dictTraversable_2.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V2, gopurs_runtime.Apply(f_3, x_4))
}))
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_4_0, x_5))
})
}

func Call_parSequence_(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictFoldable_2 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_2_loop
_ = dictFoldable_2
return Call_parTraverse_(dictParallel_0, dictApplicative_1, dictFoldable_2, Get_identity())
}

func Call_parSequence___gopurs_runtime_Value_1071252918(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictFoldable_2 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_2_loop
_ = dictFoldable_2
return Call_parTraverse_(dictParallel_0, dictApplicative_1, dictFoldable_2, Get_identity())
}

func Call_parSequence(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictTraversable_2_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictTraversable_2 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
_ = dictTraversable_2
__local_var_3_0 := gopurs_runtime.Apply2(dictTraversable_2.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V2, x_3)
}))
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_3_0, x_4))
})
}

func Call_parOneOfMap(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictAlternative_1_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictAlternative_1 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_1_loop
_ = dictAlternative_1
Plus1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_1.V1, gopurs_runtime.Value{}))
_ = Plus1_2_0
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
alt_6_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(Plus1_2_0.V0, gopurs_runtime.Value{}), "alt")
_ = alt_6_1
__local_var_7_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_3, "foldr"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_6_1, gopurs_runtime.Apply(dictParallel_0.V2, gopurs_runtime.Apply(f_5, x_7)))
}), Plus1_2_0.V1)
_ = __local_var_7_2
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_7_2, x_8))
})
})
})
})
}

func Call_parOneOf(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictAlternative_1_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictAlternative_1 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_1_loop
_ = dictAlternative_1
Plus1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_1.V1, gopurs_runtime.Value{}))
_ = Plus1_2_0
return gopurs_runtime.Func(func(dictFoldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
alt_5_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(Plus1_2_0.V0, gopurs_runtime.Value{}), "alt")
_ = alt_5_1
__local_var_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_3, "foldr"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_5_1, gopurs_runtime.Apply(dictParallel_0.V2, x_6))
}), Plus1_2_0.V1)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_6_2, x_7))
})
})
})
}

func Call_parApply(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
Apply1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictParallel_0.V1, gopurs_runtime.Value{}))
_ = Apply1_1_0
return gopurs_runtime.Func(func(mf_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ma_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply2(Apply1_1_0.V1, gopurs_runtime.Apply(dictParallel_0.V2, mf_2), gopurs_runtime.Apply(dictParallel_0.V2, ma_3)))
})
})
}


