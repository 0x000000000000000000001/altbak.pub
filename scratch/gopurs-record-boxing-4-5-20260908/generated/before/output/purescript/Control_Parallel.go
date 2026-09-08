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
		cache_Control_Parallel_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Control_Parallel_identity
}

var cache_Control_Parallel_parTraverse_ gopurs_runtime.Value
var once_Control_Parallel_parTraverse_ sync.Once
func Get_Control_Parallel_parTraverse_() gopurs_runtime.Value {
	once_Control_Parallel_parTraverse_.Do(func() {
		cache_Control_Parallel_parTraverse_ = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parTraverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_2_box), f_3_box)
})
	})
	return cache_Control_Parallel_parTraverse_
}

var cache_Control_Parallel_parTraverse___1848589797 gopurs_runtime.Value
var once_Control_Parallel_parTraverse___1848589797 sync.Once
func Get_Control_Parallel_parTraverse___1848589797() gopurs_runtime.Value {
	once_Control_Parallel_parTraverse___1848589797.Do(func() {
		cache_Control_Parallel_parTraverse___1848589797 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parTraverse___1848589797(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Control_Parallel_parTraverse___1848589797
}

var cache_Control_Parallel_parTraverse gopurs_runtime.Value
var once_Control_Parallel_parTraverse sync.Once
func Get_Control_Parallel_parTraverse() gopurs_runtime.Value {
	once_Control_Parallel_parTraverse.Do(func() {
		cache_Control_Parallel_parTraverse = gopurs_runtime.Func4(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parTraverse(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_2_box), f_3_box)
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

var cache_Control_Parallel_parSequence___182504245 gopurs_runtime.Value
var once_Control_Parallel_parSequence___182504245 sync.Once
func Get_Control_Parallel_parSequence___182504245() gopurs_runtime.Value {
	once_Control_Parallel_parSequence___182504245.Do(func() {
		cache_Control_Parallel_parSequence___182504245 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parSequence___182504245(__eta_norm_0_0_box)
})
	})
	return cache_Control_Parallel_parSequence___182504245
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
		cache_Control_Parallel_parOneOfMap = gopurs_runtime.Func2(func(dictParallel_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parOneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_1_box))
})
	})
	return cache_Control_Parallel_parOneOfMap
}

var cache_Control_Parallel_parOneOf gopurs_runtime.Value
var once_Control_Parallel_parOneOf sync.Once
func Get_Control_Parallel_parOneOf() gopurs_runtime.Value {
	once_Control_Parallel_parOneOf.Do(func() {
		cache_Control_Parallel_parOneOf = gopurs_runtime.Func2(func(dictParallel_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_parOneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_1_box))
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

func Call_Control_Parallel_parTraverse_(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictFoldable_2 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_2_loop
_ = dictFoldable_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp Any [Unit]))
__local_var_4_0 := gopurs_runtime.Apply3(Get_Data_Foldable_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_2)}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V2), gopurs_runtime.Apply(f_3, x_4))
}))
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V3), gopurs_runtime.Apply(__local_var_4_0, x_5))
})
}

func Call_Control_Parallel_parTraverse___1848589797(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
parTraverse___1848589797:
for {
if false { continue parTraverse___1848589797 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply(Get_Effect_Aff__sequential(), gopurs_runtime.Apply4(Get_Data_Foldable_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), __eta_norm_0_1))
}
}

func Call_Control_Parallel_parTraverse(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictTraversable_2_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictTraversable_2 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
_ = dictTraversable_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp Any [(TypeApp (TypeVar t) [(TypeVar b)])]))
__local_var_4_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_2.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V2), gopurs_runtime.Apply(f_3, x_4))
}))
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V3), gopurs_runtime.Apply(__local_var_4_0, x_5))
})
}

func Call_Control_Parallel_parSequence_(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictFoldable_2 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_2_loop
_ = dictFoldable_2
return Call_Control_Parallel_parTraverse_(dictParallel_0, dictApplicative_1, dictFoldable_2, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Control_Parallel_parSequence___182504245(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
parSequence___182504245:
for {
if false { continue parSequence___182504245 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Call_Control_Parallel_parTraverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_parallelAff()), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)), __eta_norm_0_0)
}
}

func Call_Control_Parallel_parSequence(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictTraversable_2_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictTraversable_2 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
_ = dictTraversable_2
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp Any [(TypeApp (TypeVar t) [(TypeVar b)])]))
__local_var_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_2.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V2), x_3)
}))
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V3), gopurs_runtime.Apply(__local_var_3_0, x_4))
})
}

func Call_Control_Parallel_parOneOfMap(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictAlternative_1_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictAlternative_1 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_1_loop
_ = dictAlternative_1
// TAST (Let): Plus1_2_0 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [Any])
Plus1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_1.V1), gopurs_runtime.Value{}))
_ = Plus1_2_0
return gopurs_runtime.Func3(func(dictFoldable_3 gopurs_runtime.Value, dictFunctor_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): alt_6_2 shape=Other bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar b)]), (TypeApp (TypeVar g) [(TypeVar b)])] (TypeApp (TypeVar g) [(TypeVar b)]))
alt_6_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Plus1_2_0.V0), gopurs_runtime.Value{}), "alt")
_ = alt_6_2
// TAST (Let): __local_var_6_1 shape=Let(App(Other)) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp Any [(TypeVar b)]))
__local_var_6_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_3, "foldr"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_6_2, gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V2), gopurs_runtime.Apply(f_5, x_7)))
}), gopurs_runtime.Box(Plus1_2_0.V1))
_ = __local_var_6_1
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V3), gopurs_runtime.Apply(__local_var_6_1, x_7))
})
})
}

func Call_Control_Parallel_parOneOf(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictAlternative_1_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictAlternative_1 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_1_loop
_ = dictAlternative_1
// TAST (Let): Plus1_2_0 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [Any])
Plus1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_1.V1), gopurs_runtime.Value{}))
_ = Plus1_2_0
return gopurs_runtime.Func2(func(dictFoldable_3 gopurs_runtime.Value, dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): alt_5_2 shape=Other bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar b)]), (TypeApp (TypeVar g) [(TypeVar b)])] (TypeApp (TypeVar g) [(TypeVar b)]))
alt_5_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Plus1_2_0.V0), gopurs_runtime.Value{}), "alt")
_ = alt_5_2
// TAST (Let): __local_var_5_1 shape=Let(App(Other)) bindingType=(Func [(TypeApp (TypeVar t) [(TypeApp (TypeVar m) [(TypeVar a)])])] (TypeApp Any [(TypeVar a)]))
__local_var_5_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_3, "foldr"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_5_2, gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V2), x_6))
}), gopurs_runtime.Box(Plus1_2_0.V1))
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V3), gopurs_runtime.Apply(__local_var_5_1, x_6))
})
})
}

func Call_Control_Parallel_parApply(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): Apply1_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [Any])
Apply1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V1), gopurs_runtime.Value{}))
_ = Apply1_1_0
return gopurs_runtime.Func2(func(mf_2 gopurs_runtime.Value, ma_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V3), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V2), mf_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictParallel_0.V2), ma_3)))
})
}


