package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Parallel_identity gopurs_runtime.Value
var once_Control_Parallel_identity sync.Once
func Get_Control_Parallel_identity() gopurs_runtime.Value {
	once_Control_Parallel_identity.Do(func() {
		cache_Control_Parallel_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Control_Parallel_identity
}

var cache_Control_Parallel_identity1 gopurs_runtime.Value
var once_Control_Parallel_identity1 sync.Once
func Get_Control_Parallel_identity1() gopurs_runtime.Value {
	once_Control_Parallel_identity1.Do(func() {
		cache_Control_Parallel_identity1 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Control_Parallel_identity1
}

var cache_Control_Parallel_parTraverse_ gopurs_runtime.Value
var once_Control_Parallel_parTraverse_ sync.Once
func Get_Control_Parallel_parTraverse_() gopurs_runtime.Value {
	once_Control_Parallel_parTraverse_.Do(func() {
		cache_Control_Parallel_parTraverse_ = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parTraverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
})
	})
	return cache_Control_Parallel_parTraverse_
}

var cache_Control_Parallel_parTraverse gopurs_runtime.Value
var once_Control_Parallel_parTraverse sync.Once
func Get_Control_Parallel_parTraverse() gopurs_runtime.Value {
	once_Control_Parallel_parTraverse.Do(func() {
		cache_Control_Parallel_parTraverse = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parTraverse(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
})
	})
	return cache_Control_Parallel_parTraverse
}

var cache_Control_Parallel_parSequence_ gopurs_runtime.Value
var once_Control_Parallel_parSequence_ sync.Once
func Get_Control_Parallel_parSequence_() gopurs_runtime.Value {
	once_Control_Parallel_parSequence_.Do(func() {
		cache_Control_Parallel_parSequence_ = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parSequence_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_2_box))
})
	})
	return cache_Control_Parallel_parSequence_
}

var cache_Control_Parallel_parSequence___959287507 gopurs_runtime.Value
var once_Control_Parallel_parSequence___959287507 sync.Once
func Get_Control_Parallel_parSequence___959287507() gopurs_runtime.Value {
	once_Control_Parallel_parSequence___959287507.Do(func() {
		cache_Control_Parallel_parSequence___959287507 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parSequence___959287507(__eta_norm_0_0_box)
})
	})
	return cache_Control_Parallel_parSequence___959287507
}

var cache_Control_Parallel_parSequence gopurs_runtime.Value
var once_Control_Parallel_parSequence sync.Once
func Get_Control_Parallel_parSequence() gopurs_runtime.Value {
	once_Control_Parallel_parSequence.Do(func() {
		cache_Control_Parallel_parSequence = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parSequence(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_2_box))
})
	})
	return cache_Control_Parallel_parSequence
}

var cache_Control_Parallel_parOneOfMap gopurs_runtime.Value
var once_Control_Parallel_parOneOfMap sync.Once
func Get_Control_Parallel_parOneOfMap() gopurs_runtime.Value {
	once_Control_Parallel_parOneOfMap.Do(func() {
		cache_Control_Parallel_parOneOfMap = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parOneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
})
	})
	return cache_Control_Parallel_parOneOfMap
}

var cache_Control_Parallel_parOneOf gopurs_runtime.Value
var once_Control_Parallel_parOneOf sync.Once
func Get_Control_Parallel_parOneOf() gopurs_runtime.Value {
	once_Control_Parallel_parOneOf.Do(func() {
		cache_Control_Parallel_parOneOf = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parOneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
})
	})
	return cache_Control_Parallel_parOneOf
}

var cache_Control_Parallel_parApply gopurs_runtime.Value
var once_Control_Parallel_parApply sync.Once
func Get_Control_Parallel_parApply() gopurs_runtime.Value {
	once_Control_Parallel_parApply.Do(func() {
		cache_Control_Parallel_parApply = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parApply(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
})
	})
	return cache_Control_Parallel_parApply
}

func Call_Control_Parallel_parTraverse_(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): sequential_1_0 shape=App(Var) bindingType=(Func [(TypeApp Any [Unit])] (TypeApp (TypeVar m$scope6) [Unit]))
sequential_1_0 := Call_Control_Parallel_Class_sequential(dictParallel_0)
_ = sequential_1_0
// TAST (Let): parallel_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope6) [(TypeVar b$scope9)])] (TypeApp Any [(TypeVar b$scope9)]))
parallel_2_1 := Call_Control_Parallel_Class_parallel(dictParallel_0)
_ = parallel_2_1
return gopurs_runtime.Func3(func(dictApplicative_3 gopurs_runtime.Value, dictFoldable_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential_1_0, gopurs_runtime.Apply2(Call_Data_Foldable_traverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_4))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), parallel_2_1, f_5)))
})
}

func Call_Control_Parallel_parTraverse(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): sequential_1_0 shape=App(Var) bindingType=(Func [(TypeApp Any [(TypeApp (TypeVar t$scope19) [(TypeVar b$scope21)])])] (TypeApp (TypeVar m$scope18) [(TypeApp (TypeVar t$scope19) [(TypeVar b$scope21)])]))
sequential_1_0 := Call_Control_Parallel_Class_sequential(dictParallel_0)
_ = sequential_1_0
// TAST (Let): parallel_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope18) [(TypeVar b$scope21)])] (TypeApp Any [(TypeVar b$scope21)]))
parallel_2_1 := Call_Control_Parallel_Class_parallel(dictParallel_0)
_ = parallel_2_1
return gopurs_runtime.Func3(func(dictApplicative_3 gopurs_runtime.Value, dictTraversable_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential_1_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_4, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), parallel_2_1, f_5)))
})
}

func Call_Control_Parallel_parSequence_(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictFoldable_2 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_2_loop
_ = dictFoldable_2
return gopurs_runtime.Apply3(Call_Control_Parallel_parTraverse_(dictParallel_0), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_2)}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Control_Parallel_parSequence___959287507(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
parSequence___959287507:
for {
if false { continue parSequence___959287507 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply4(Call_Control_Parallel_parTraverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_parallelAff())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), __eta_norm_0_0)
}
}

func Call_Control_Parallel_parSequence(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictTraversable_2_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictTraversable_2 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
_ = dictTraversable_2
return gopurs_runtime.Apply3(Call_Control_Parallel_parTraverse(dictParallel_0), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(dictTraversable_2)}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Control_Parallel_parOneOfMap(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): sequential_1_0 shape=App(Var) bindingType=(Func [(TypeApp Any [(TypeVar b$scope48)])] (TypeApp (TypeVar m$scope50) [(TypeVar b$scope48)]))
sequential_1_0 := Call_Control_Parallel_Class_sequential(dictParallel_0)
_ = sequential_1_0
// TAST (Let): parallel_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope50) [(TypeVar b$scope48)])] (TypeApp Any [(TypeVar b$scope48)]))
parallel_2_1 := Call_Control_Parallel_Class_parallel(dictParallel_0)
_ = parallel_2_1
return gopurs_runtime.Func(func(dictAlternative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Plus1_4_2 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [Any])
Plus1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_3, "Plus1"), gopurs_runtime.Value{}))
_ = Plus1_4_2
return gopurs_runtime.Func3(func(dictFoldable_5 gopurs_runtime.Value, dictFunctor_6 gopurs_runtime.Value, f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential_1_0, gopurs_runtime.Apply(Call_Data_Foldable_oneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_5), Plus1_4_2), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), parallel_2_1, f_7)))
})
})
}

func Call_Control_Parallel_parOneOf(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): sequential_1_0 shape=App(Var) bindingType=(Func [(TypeApp Any [(TypeVar a$scope58)])] (TypeApp (TypeVar m$scope60) [(TypeVar a$scope58)]))
sequential_1_0 := Call_Control_Parallel_Class_sequential(dictParallel_0)
_ = sequential_1_0
// TAST (Let): parallel_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope60) [(TypeVar a$scope58)])] (TypeApp Any [(TypeVar a$scope58)]))
parallel_2_1 := Call_Control_Parallel_Class_parallel(dictParallel_0)
_ = parallel_2_1
return gopurs_runtime.Func(func(dictAlternative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Plus1_4_2 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [Any])
Plus1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_3, "Plus1"), gopurs_runtime.Value{}))
_ = Plus1_4_2
return gopurs_runtime.Func2(func(dictFoldable_5 gopurs_runtime.Value, dictFunctor_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential_1_0, gopurs_runtime.Apply(Call_Data_Foldable_oneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_5), Plus1_4_2), parallel_2_1))
})
})
}

func Call_Control_Parallel_parApply(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): Apply1_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [Any])
Apply1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictParallel_0.V1, gopurs_runtime.Value{}))
_ = Apply1_1_0
return gopurs_runtime.Func2(func(mf_2 gopurs_runtime.Value, ma_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply2(Apply1_1_0.V1, gopurs_runtime.Apply(dictParallel_0.V2, mf_2), gopurs_runtime.Apply(dictParallel_0.V2, ma_3)))
})
}


