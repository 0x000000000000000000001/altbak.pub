package Control_Parallel

import (
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Parallel_Class "gopurs/output/Control.Parallel.Class"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Unit "gopurs/output/Data.Unit"
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

var cache_pure__2935994064 gopurs_runtime.Value
var once_pure__2935994064 sync.Once
func Get_pure__2935994064() gopurs_runtime.Value {
	once_pure__2935994064.Do(func() {
		cache_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2935994064(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2935994064
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_parallel__2242335472 gopurs_runtime.Value
var once_parallel__2242335472 sync.Once
func Get_parallel__2242335472() gopurs_runtime.Value {
	once_parallel__2242335472.Do(func() {
		cache_parallel__2242335472 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel__2242335472(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_parallel__2242335472
}

var cache_parallel__4223476656 gopurs_runtime.Value
var once_parallel__4223476656 sync.Once
func Get_parallel__4223476656() gopurs_runtime.Value {
	once_parallel__4223476656.Do(func() {
		cache_parallel__4223476656 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel__4223476656(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_parallel__4223476656
}

var cache_sequential__2242335472 gopurs_runtime.Value
var once_sequential__2242335472 sync.Once
func Get_sequential__2242335472() gopurs_runtime.Value {
	once_sequential__2242335472.Do(func() {
		cache_sequential__2242335472 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequential__2242335472(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequential__2242335472
}

var cache_parTraverse__1055730709 gopurs_runtime.Value
var once_parTraverse__1055730709 sync.Once
func Get_parTraverse__1055730709() gopurs_runtime.Value {
	once_parTraverse__1055730709.Do(func() {
		cache_parTraverse__1055730709 = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse__1055730709(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_2_box), f_3_box)
})
	})
	return cache_parTraverse__1055730709
}

var cache_parTraverse__4270952213 gopurs_runtime.Value
var once_parTraverse__4270952213 sync.Once
func Get_parTraverse__4270952213() gopurs_runtime.Value {
	once_parTraverse__4270952213.Do(func() {
		cache_parTraverse__4270952213 = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse__4270952213(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_2_box), f_3_box)
})
	})
	return cache_parTraverse__4270952213
}

var cache_parTraverse___1426351978 gopurs_runtime.Value
var once_parTraverse___1426351978 sync.Once
func Get_parTraverse___1426351978() gopurs_runtime.Value {
	once_parTraverse___1426351978.Do(func() {
		cache_parTraverse___1426351978 = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse___1426351978(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2_box), f_3_box)
})
	})
	return cache_parTraverse___1426351978
}

var cache_parTraverse___1113625962 gopurs_runtime.Value
var once_parTraverse___1113625962 sync.Once
func Get_parTraverse___1113625962() gopurs_runtime.Value {
	once_parTraverse___1113625962.Do(func() {
		cache_parTraverse___1113625962 = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parTraverse___1113625962(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2_box), f_3_box)
})
	})
	return cache_parTraverse___1113625962
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__3591001499 gopurs_runtime.Value
var once_foldr__3591001499 sync.Once
func Get_foldr__3591001499() gopurs_runtime.Value {
	once_foldr__3591001499.Do(func() {
		cache_foldr__3591001499 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3591001499(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__3591001499
}

var cache_foldr__2492367323 gopurs_runtime.Value
var once_foldr__2492367323 sync.Once
func Get_foldr__2492367323() gopurs_runtime.Value {
	once_foldr__2492367323.Do(func() {
		cache_foldr__2492367323 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2492367323(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2492367323
}

var cache_oneOfMap__3719016818 gopurs_runtime.Value
var once_oneOfMap__3719016818 sync.Once
func Get_oneOfMap__3719016818() gopurs_runtime.Value {
	once_oneOfMap__3719016818.Do(func() {
		cache_oneOfMap__3719016818 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOfMap__3719016818(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](dictPlus_1_box))
})
	})
	return cache_oneOfMap__3719016818
}

var cache_oneOfMap__1349369970 gopurs_runtime.Value
var once_oneOfMap__1349369970 sync.Once
func Get_oneOfMap__1349369970() gopurs_runtime.Value {
	once_oneOfMap__1349369970.Do(func() {
		cache_oneOfMap__1349369970 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOfMap__1349369970(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](dictPlus_1_box))
})
	})
	return cache_oneOfMap__1349369970
}

var cache_traverse___996968168 gopurs_runtime.Value
var once_traverse___996968168 sync.Once
func Get_traverse___996968168() gopurs_runtime.Value {
	once_traverse___996968168.Do(func() {
		cache_traverse___996968168 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse___996968168(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_traverse___996968168
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
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

func Call_parSequence_(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
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
alt_6_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(Plus1_2_0.V0, gopurs_runtime.Value{}), "alt")
_ = alt_6_2
__local_var_6_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_3, "foldr"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_6_2, gopurs_runtime.Apply(dictParallel_0.V2, gopurs_runtime.Apply(f_5, x_7)))
}), Plus1_2_0.V1)
_ = __local_var_6_1
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_6_1, x_7))
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
alt_5_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(Plus1_2_0.V0, gopurs_runtime.Value{}), "alt")
_ = alt_5_2
__local_var_5_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_3, "foldr"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_5_2, gopurs_runtime.Apply(dictParallel_0.V2, x_6))
}), Plus1_2_0.V1)
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply(__local_var_5_1, x_6))
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

func Call_pure__2935994064(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_parallel__2242335472(dict_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_parallel__4223476656(dict_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_sequential__2242335472(dict_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_parTraverse__1055730709(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictTraversable_2_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_parTraverse__4270952213(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictTraversable_2_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_parTraverse___1426351978(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_parTraverse___1113625962(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__3591001499(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2492367323(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_oneOfMap__3719016818(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictPlus_1_loop *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value] = dictPlus_1_loop
_ = dictPlus_1
alt_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictPlus_1.V0, gopurs_runtime.Value{}), "alt")
_ = alt_2_0
empty_3_1 := dictPlus_1.V1
_ = empty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_2_0, gopurs_runtime.Apply(f_4, x_5))
}), empty_3_1)
})
}

func Call_oneOfMap__1349369970(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictPlus_1_loop *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value] = dictPlus_1_loop
_ = dictPlus_1
alt_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictPlus_1.V0, gopurs_runtime.Value{}), "alt")
_ = alt_2_0
empty_3_1 := dictPlus_1.V1
_ = empty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_2_0, gopurs_runtime.Apply(f_4, x_5))
}), empty_3_1)
})
}

func Call_traverse___996968168(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_1 := gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{})
_ = __local_var_1_1
Functor0_2_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(Functor0_2_2.V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), a_3), b_4)
})
})
_ = applySecond_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(f_3, x_4))
}), gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit()))
})
})
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}


