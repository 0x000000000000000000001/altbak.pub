package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Aff_pure gopurs_runtime.Value
var once_Effect_Aff_pure sync.Once
func Get_Effect_Aff_pure() gopurs_runtime.Value {
	once_Effect_Aff_pure.Do(func() {
		cache_Effect_Aff_pure = gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure")
	})
	return cache_Effect_Aff_pure
}

var cache_Effect_Aff_void gopurs_runtime.Value
var once_Effect_Aff_void sync.Once
func Get_Effect_Aff_void() gopurs_runtime.Value {
	once_Effect_Aff_void.Do(func() {
		cache_Effect_Aff_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_functorEffect(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
	})
	return cache_Effect_Aff_void
}

var cache_Effect_Aff_void1 gopurs_runtime.Value
var once_Effect_Aff_void1 sync.Once
func Get_Effect_Aff_void1() gopurs_runtime.Value {
	once_Effect_Aff_void1.Do(func() {
		cache_Effect_Aff_void1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_functorEffect(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
	})
	return cache_Effect_Aff_void1
}

var cache_Effect_Aff_Fiber gopurs_runtime.Value
var once_Effect_Aff_Fiber sync.Once
func Get_Effect_Aff_Fiber() gopurs_runtime.Value {
	once_Effect_Aff_Fiber.Do(func() {
		cache_Effect_Aff_Fiber = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Fiber(x_0_box)
})
	})
	return cache_Effect_Aff_Fiber
}

var cache_Effect_Aff_Canceler gopurs_runtime.Value
var once_Effect_Aff_Canceler sync.Once
func Get_Effect_Aff_Canceler() gopurs_runtime.Value {
	once_Effect_Aff_Canceler.Do(func() {
		cache_Effect_Aff_Canceler = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Canceler(x_0_box)
})
	})
	return cache_Effect_Aff_Canceler
}

var cache_Effect_Aff_newtypeCanceler gopurs_runtime.Value
var once_Effect_Aff_newtypeCanceler sync.Once
func Get_Effect_Aff_newtypeCanceler() gopurs_runtime.Value {
	once_Effect_Aff_newtypeCanceler.Do(func() {
		cache_Effect_Aff_newtypeCanceler = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Effect_Aff_newtypeCanceler
}

var cache_Effect_Aff_makeFiber gopurs_runtime.Value
var once_Effect_Aff_makeFiber sync.Once
func Get_Effect_Aff_makeFiber() gopurs_runtime.Value {
	once_Effect_Aff_makeFiber.Do(func() {
		cache_Effect_Aff_makeFiber = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeFiber(aff_0_box)
})
	})
	return cache_Effect_Aff_makeFiber
}

var cache_Effect_Aff_makeAff gopurs_runtime.Value
var once_Effect_Aff_makeAff sync.Once
func Get_Effect_Aff_makeAff() gopurs_runtime.Value {
	once_Effect_Aff_makeAff.Do(func() {
		cache_Effect_Aff_makeAff = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeAff(build_0_box)
})
	})
	return cache_Effect_Aff_makeAff
}

var cache_Effect_Aff_launchSuspendedAff gopurs_runtime.Value
var once_Effect_Aff_launchSuspendedAff sync.Once
func Get_Effect_Aff_launchSuspendedAff() gopurs_runtime.Value {
	once_Effect_Aff_launchSuspendedAff.Do(func() {
		cache_Effect_Aff_launchSuspendedAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_launchSuspendedAff(aff_0_box)
})
	})
	return cache_Effect_Aff_launchSuspendedAff
}

var cache_Effect_Aff_launchAff gopurs_runtime.Value
var once_Effect_Aff_launchAff sync.Once
func Get_Effect_Aff_launchAff() gopurs_runtime.Value {
	once_Effect_Aff_launchAff.Do(func() {
		cache_Effect_Aff_launchAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_launchAff(aff_0_box)
})
	})
	return cache_Effect_Aff_launchAff
}

var cache_Effect_Aff_launchAff_ gopurs_runtime.Value
var once_Effect_Aff_launchAff_ sync.Once
func Get_Effect_Aff_launchAff_() gopurs_runtime.Value {
	once_Effect_Aff_launchAff_.Do(func() {
		cache_Effect_Aff_launchAff_ = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_launchAff_(x_0_box)
})
	})
	return cache_Effect_Aff_launchAff_
}

var cache_Effect_Aff_functorParAff gopurs_runtime.Value
var once_Effect_Aff_functorParAff sync.Once
func Get_Effect_Aff_functorParAff() gopurs_runtime.Value {
	once_Effect_Aff_functorParAff.Do(func() {
		cache_Effect_Aff_functorParAff = gopurs_runtime.RecordDict1("map", Get_Effect_Aff__parAffMap())
	})
	return cache_Effect_Aff_functorParAff
}

var cache_Effect_Aff_functorAff gopurs_runtime.Value
var once_Effect_Aff_functorAff sync.Once
func Get_Effect_Aff_functorAff() gopurs_runtime.Value {
	once_Effect_Aff_functorAff.Do(func() {
		cache_Effect_Aff_functorAff = gopurs_runtime.RecordDict1("map", Get_Effect_Aff__map())
	})
	return cache_Effect_Aff_functorAff
}

var cache_Effect_Aff_delay gopurs_runtime.Value
var once_Effect_Aff_delay sync.Once
func Get_Effect_Aff_delay() gopurs_runtime.Value {
	once_Effect_Aff_delay.Do(func() {
		cache_Effect_Aff_delay = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_delay(v_0_box.FloatVal())
})
	})
	return cache_Effect_Aff_delay
}

var cache_Effect_Aff_bracket gopurs_runtime.Value
var once_Effect_Aff_bracket sync.Once
func Get_Effect_Aff_bracket() gopurs_runtime.Value {
	once_Effect_Aff_bracket.Do(func() {
		cache_Effect_Aff_bracket = gopurs_runtime.Func2(func(acquire_0_box gopurs_runtime.Value, completed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_bracket(acquire_0_box, completed_1_box)
})
	})
	return cache_Effect_Aff_bracket
}

var cache_Effect_Aff_applyParAff gopurs_runtime.Value
var once_Effect_Aff_applyParAff sync.Once
func Get_Effect_Aff_applyParAff() gopurs_runtime.Value {
	once_Effect_Aff_applyParAff.Do(func() {
		cache_Effect_Aff_applyParAff = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorParAff()
}), Get_Effect_Aff__parAffApply())
	})
	return cache_Effect_Aff_applyParAff
}

var cache_Effect_Aff_semigroupParAff gopurs_runtime.Value
var once_Effect_Aff_semigroupParAff sync.Once
func Get_Effect_Aff_semigroupParAff() gopurs_runtime.Value {
	once_Effect_Aff_semigroupParAff.Do(func() {
		cache_Effect_Aff_semigroupParAff = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_semigroupParAff(dictSemigroup_0_box)
})
	})
	return cache_Effect_Aff_semigroupParAff
}

var cache_Effect_Aff_monadAff gopurs_runtime.Value
var once_Effect_Aff_monadAff sync.Once
func Get_Effect_Aff_monadAff() gopurs_runtime.Value {
	once_Effect_Aff_monadAff.Do(func() {
		cache_Effect_Aff_monadAff = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applicativeAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_bindAff()
}))
	})
	return cache_Effect_Aff_monadAff
}

var cache_Effect_Aff_bindAff gopurs_runtime.Value
var once_Effect_Aff_bindAff sync.Once
func Get_Effect_Aff_bindAff() gopurs_runtime.Value {
	once_Effect_Aff_bindAff.Do(func() {
		cache_Effect_Aff_bindAff = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), Get_Effect_Aff__bind())
	})
	return cache_Effect_Aff_bindAff
}

var cache_Effect_Aff_applyAff gopurs_runtime.Value
var once_Effect_Aff_applyAff sync.Once
func Get_Effect_Aff_applyAff() gopurs_runtime.Value {
	once_Effect_Aff_applyAff.Do(func() {
		cache_Effect_Aff_applyAff = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadAff(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadAff(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorAff()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Effect_Aff_applyAff
}

var cache_Effect_Aff_applicativeAff gopurs_runtime.Value
var once_Effect_Aff_applicativeAff sync.Once
func Get_Effect_Aff_applicativeAff() gopurs_runtime.Value {
	once_Effect_Aff_applicativeAff.Do(func() {
		cache_Effect_Aff_applicativeAff = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), Get_Effect_Aff__pure())
	})
	return cache_Effect_Aff_applicativeAff
}

var cache_Effect_Aff_pure1 gopurs_runtime.Value
var once_Effect_Aff_pure1 sync.Once
func Get_Effect_Aff_pure1() gopurs_runtime.Value {
	once_Effect_Aff_pure1.Do(func() {
		cache_Effect_Aff_pure1 = gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure")
	})
	return cache_Effect_Aff_pure1
}

var cache_Effect_Aff_cancelWith gopurs_runtime.Value
var once_Effect_Aff_cancelWith sync.Once
func Get_Effect_Aff_cancelWith() gopurs_runtime.Value {
	once_Effect_Aff_cancelWith.Do(func() {
		cache_Effect_Aff_cancelWith = gopurs_runtime.Func2(func(aff_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_cancelWith(aff_0_box, v_1_box)
})
	})
	return cache_Effect_Aff_cancelWith
}

var cache_Effect_Aff_finally gopurs_runtime.Value
var once_Effect_Aff_finally sync.Once
func Get_Effect_Aff_finally() gopurs_runtime.Value {
	once_Effect_Aff_finally.Do(func() {
		cache_Effect_Aff_finally = gopurs_runtime.Func2(func(fin_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_finally(fin_0_box, a_1_box)
})
	})
	return cache_Effect_Aff_finally
}

var cache_Effect_Aff_invincible gopurs_runtime.Value
var once_Effect_Aff_invincible sync.Once
func Get_Effect_Aff_invincible() gopurs_runtime.Value {
	once_Effect_Aff_invincible.Do(func() {
		cache_Effect_Aff_invincible = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_invincible(a_0_box)
})
	})
	return cache_Effect_Aff_invincible
}

var cache_Effect_Aff_lazyAff gopurs_runtime.Value
var once_Effect_Aff_lazyAff sync.Once
func Get_Effect_Aff_lazyAff() gopurs_runtime.Value {
	once_Effect_Aff_lazyAff.Do(func() {
		cache_Effect_Aff_lazyAff = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()), f_0)
}))
	})
	return cache_Effect_Aff_lazyAff
}

var cache_Effect_Aff_parallelAff gopurs_runtime.Value
var once_Effect_Aff_parallelAff sync.Once
func Get_Effect_Aff_parallelAff() gopurs_runtime.Value {
	once_Effect_Aff_parallelAff.Do(func() {
		cache_Effect_Aff_parallelAff = gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyParAff()
}), Get_Unsafe_Coerce_unsafeCoerce(), Get_Effect_Aff__sequential())
	})
	return cache_Effect_Aff_parallelAff
}

var cache_Effect_Aff_applicativeParAff gopurs_runtime.Value
var once_Effect_Aff_applicativeParAff sync.Once
func Get_Effect_Aff_applicativeParAff() gopurs_runtime.Value {
	once_Effect_Aff_applicativeParAff.Do(func() {
		cache_Effect_Aff_applicativeParAff = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyParAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_parallelAff(), "parallel"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"), x_0))
}))
	})
	return cache_Effect_Aff_applicativeParAff
}

var cache_Effect_Aff_monoidParAff gopurs_runtime.Value
var once_Effect_Aff_monoidParAff sync.Once
func Get_Effect_Aff_monoidParAff() gopurs_runtime.Value {
	once_Effect_Aff_monoidParAff.Do(func() {
		cache_Effect_Aff_monoidParAff = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_monoidParAff(dictMonoid_0_box)
})
	})
	return cache_Effect_Aff_monoidParAff
}

var cache_Effect_Aff_semigroupCanceler gopurs_runtime.Value
var once_Effect_Aff_semigroupCanceler sync.Once
func Get_Effect_Aff_semigroupCanceler() gopurs_runtime.Value {
	once_Effect_Aff_semigroupCanceler.Do(func() {
		cache_Effect_Aff_semigroupCanceler = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(Get_Control_Parallel_parTraverse_(), gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_parallelAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}, Get_Control_Parallel_identity(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(v_0, err_2), gopurs_runtime.Apply(v1_1, err_2)}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
})
}))
	})
	return cache_Effect_Aff_semigroupCanceler
}

var cache_Effect_Aff_semigroupAff gopurs_runtime.Value
var once_Effect_Aff_semigroupAff sync.Once
func Get_Effect_Aff_semigroupAff() gopurs_runtime.Value {
	once_Effect_Aff_semigroupAff.Do(func() {
		cache_Effect_Aff_semigroupAff = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_semigroupAff(dictSemigroup_0_box)
})
	})
	return cache_Effect_Aff_semigroupAff
}

var cache_Effect_Aff_monadEffectAff gopurs_runtime.Value
var once_Effect_Aff_monadEffectAff sync.Once
func Get_Effect_Aff_monadEffectAff() gopurs_runtime.Value {
	once_Effect_Aff_monadEffectAff.Do(func() {
		cache_Effect_Aff_monadEffectAff = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadAff()
}), Get_Effect_Aff__liftEffect())
	})
	return cache_Effect_Aff_monadEffectAff
}

var cache_Effect_Aff_liftEffect gopurs_runtime.Value
var once_Effect_Aff_liftEffect sync.Once
func Get_Effect_Aff_liftEffect() gopurs_runtime.Value {
	once_Effect_Aff_liftEffect.Do(func() {
		cache_Effect_Aff_liftEffect = gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect")
	})
	return cache_Effect_Aff_liftEffect
}

var cache_Effect_Aff_effectCanceler gopurs_runtime.Value
var once_Effect_Aff_effectCanceler sync.Once
func Get_Effect_Aff_effectCanceler() gopurs_runtime.Value {
	once_Effect_Aff_effectCanceler.Do(func() {
		cache_Effect_Aff_effectCanceler = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_effectCanceler(x_0_box)
})
	})
	return cache_Effect_Aff_effectCanceler
}

var cache_Effect_Aff_joinFiber gopurs_runtime.Value
var once_Effect_Aff_joinFiber sync.Once
func Get_Effect_Aff_joinFiber() gopurs_runtime.Value {
	once_Effect_Aff_joinFiber.Do(func() {
		cache_Effect_Aff_joinFiber = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_joinFiber(v_0_box)
})
	})
	return cache_Effect_Aff_joinFiber
}

var cache_Effect_Aff_functorFiber gopurs_runtime.Value
var once_Effect_Aff_functorFiber sync.Once
func Get_Effect_Aff_functorFiber() gopurs_runtime.Value {
	once_Effect_Aff_functorFiber.Do(func() {
		cache_Effect_Aff_functorFiber = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), Call_Effect_Aff_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_functorAff(), "map"), f_0, Call_Effect_Aff_joinFiber(t_1))))
})
}))
	})
	return cache_Effect_Aff_functorFiber
}

var cache_Effect_Aff_applyFiber gopurs_runtime.Value
var once_Effect_Aff_applyFiber sync.Once
func Get_Effect_Aff_applyFiber() gopurs_runtime.Value {
	once_Effect_Aff_applyFiber.Do(func() {
		cache_Effect_Aff_applyFiber = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorFiber()
}), gopurs_runtime.Func(func(t1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), Call_Effect_Aff_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_applyAff(), "apply"), Call_Effect_Aff_joinFiber(t1_0), Call_Effect_Aff_joinFiber(t2_1))))
})
}))
	})
	return cache_Effect_Aff_applyFiber
}

var cache_Effect_Aff_applicativeFiber gopurs_runtime.Value
var once_Effect_Aff_applicativeFiber sync.Once
func Get_Effect_Aff_applicativeFiber() gopurs_runtime.Value {
	once_Effect_Aff_applicativeFiber.Do(func() {
		cache_Effect_Aff_applicativeFiber = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyFiber()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), Call_Effect_Aff_makeFiber(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"), a_0)))
}))
	})
	return cache_Effect_Aff_applicativeFiber
}

var cache_Effect_Aff_forkAff gopurs_runtime.Value
var once_Effect_Aff_forkAff sync.Once
func Get_Effect_Aff_forkAff() gopurs_runtime.Value {
	once_Effect_Aff_forkAff.Do(func() {
		cache_Effect_Aff_forkAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_forkAff(aff_0_box)
})
	})
	return cache_Effect_Aff_forkAff
}

var cache_Effect_Aff_killFiber gopurs_runtime.Value
var once_Effect_Aff_killFiber sync.Once
func Get_Effect_Aff_killFiber() gopurs_runtime.Value {
	once_Effect_Aff_killFiber.Do(func() {
		cache_Effect_Aff_killFiber = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_killFiber(e_0_box, v_1_box)
})
	})
	return cache_Effect_Aff_killFiber
}

var cache_Effect_Aff_fiberCanceler gopurs_runtime.Value
var once_Effect_Aff_fiberCanceler sync.Once
func Get_Effect_Aff_fiberCanceler() gopurs_runtime.Value {
	once_Effect_Aff_fiberCanceler.Do(func() {
		cache_Effect_Aff_fiberCanceler = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_fiberCanceler(x_0_box, a_1_box)
})
	})
	return cache_Effect_Aff_fiberCanceler
}

var cache_Effect_Aff_supervise gopurs_runtime.Value
var once_Effect_Aff_supervise sync.Once
func Get_Effect_Aff_supervise() gopurs_runtime.Value {
	once_Effect_Aff_supervise.Do(func() {
		cache_Effect_Aff_supervise = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_supervise(aff_0_box)
})
	})
	return cache_Effect_Aff_supervise
}

var cache_Effect_Aff_suspendAff gopurs_runtime.Value
var once_Effect_Aff_suspendAff sync.Once
func Get_Effect_Aff_suspendAff() gopurs_runtime.Value {
	once_Effect_Aff_suspendAff.Do(func() {
		cache_Effect_Aff_suspendAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_suspendAff(aff_0_box)
})
	})
	return cache_Effect_Aff_suspendAff
}

var cache_Effect_Aff_monadSTAff gopurs_runtime.Value
var once_Effect_Aff_monadSTAff sync.Once
func Get_Effect_Aff_monadSTAff() gopurs_runtime.Value {
	once_Effect_Aff_monadSTAff.Do(func() {
		cache_Effect_Aff_monadSTAff = gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Class_monadSTEffect(), "liftST"), x_0))
}))
	})
	return cache_Effect_Aff_monadSTAff
}

var cache_Effect_Aff_monadThrowAff gopurs_runtime.Value
var once_Effect_Aff_monadThrowAff sync.Once
func Get_Effect_Aff_monadThrowAff() gopurs_runtime.Value {
	once_Effect_Aff_monadThrowAff.Do(func() {
		cache_Effect_Aff_monadThrowAff = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadAff()
}), Get_Effect_Aff__throwError())
	})
	return cache_Effect_Aff_monadThrowAff
}

var cache_Effect_Aff_monadErrorAff gopurs_runtime.Value
var once_Effect_Aff_monadErrorAff sync.Once
func Get_Effect_Aff_monadErrorAff() gopurs_runtime.Value {
	once_Effect_Aff_monadErrorAff.Do(func() {
		cache_Effect_Aff_monadErrorAff = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadThrowAff()
}), Get_Effect_Aff__catchError())
	})
	return cache_Effect_Aff_monadErrorAff
}

var cache_Effect_Aff_attempt gopurs_runtime.Value
var once_Effect_Aff_attempt sync.Once
func Get_Effect_Aff_attempt() gopurs_runtime.Value {
	once_Effect_Aff_attempt.Do(func() {
		cache_Effect_Aff_attempt = gopurs_runtime.Apply(Get_Control_Monad_Error_Class_try(), Get_Effect_Aff_monadErrorAff())
	})
	return cache_Effect_Aff_attempt
}

var cache_Effect_Aff_runAff gopurs_runtime.Value
var once_Effect_Aff_runAff sync.Once
func Get_Effect_Aff_runAff() gopurs_runtime.Value {
	once_Effect_Aff_runAff.Do(func() {
		cache_Effect_Aff_runAff = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_runAff(k_0_box, aff_1_box)
})
	})
	return cache_Effect_Aff_runAff
}

var cache_Effect_Aff_runAff_ gopurs_runtime.Value
var once_Effect_Aff_runAff_ sync.Once
func Get_Effect_Aff_runAff_() gopurs_runtime.Value {
	once_Effect_Aff_runAff_.Do(func() {
		cache_Effect_Aff_runAff_ = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_runAff_(k_0_box, aff_1_box)
})
	})
	return cache_Effect_Aff_runAff_
}

var cache_Effect_Aff_runSuspendedAff gopurs_runtime.Value
var once_Effect_Aff_runSuspendedAff sync.Once
func Get_Effect_Aff_runSuspendedAff() gopurs_runtime.Value {
	once_Effect_Aff_runSuspendedAff.Do(func() {
		cache_Effect_Aff_runSuspendedAff = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_runSuspendedAff(k_0_box, aff_1_box)
})
	})
	return cache_Effect_Aff_runSuspendedAff
}

var cache_Effect_Aff_monadRecAff gopurs_runtime.Value
var once_Effect_Aff_monadRecAff sync.Once
func Get_Effect_Aff_monadRecAff() gopurs_runtime.Value {
	once_Effect_Aff_monadRecAff.Do(func() {
		cache_Effect_Aff_monadRecAff = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadAff()
}), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
go__go_1_0_0 = gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), gopurs_runtime.Apply(k_0, a_2), gopurs_runtime.Func(func(res_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (res_3.Type == 9 && res_3.IntVal == 60402430) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"), (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(res_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (res_3.Type == 9 && res_3.IntVal == 525585346) {
__t1 = gopurs_runtime.Apply(go__go_1_0_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(res_3.UnsafePtr).V0)
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
return go__go_1_0_0
}))
	})
	return cache_Effect_Aff_monadRecAff
}

var cache_Effect_Aff_monoidAff gopurs_runtime.Value
var once_Effect_Aff_monoidAff sync.Once
func Get_Effect_Aff_monoidAff() gopurs_runtime.Value {
	once_Effect_Aff_monoidAff.Do(func() {
		cache_Effect_Aff_monoidAff = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_monoidAff(dictMonoid_0_box)
})
	})
	return cache_Effect_Aff_monoidAff
}

var cache_Effect_Aff_nonCanceler gopurs_runtime.Value
var once_Effect_Aff_nonCanceler sync.Once
func Get_Effect_Aff_nonCanceler() gopurs_runtime.Value {
	once_Effect_Aff_nonCanceler.Do(func() {
		cache_Effect_Aff_nonCanceler = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit())
_ = __local_var_0_0
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_0_0
})
}()
	})
	return cache_Effect_Aff_nonCanceler
}

var cache_Effect_Aff_monoidCanceler gopurs_runtime.Value
var once_Effect_Aff_monoidCanceler sync.Once
func Get_Effect_Aff_monoidCanceler() gopurs_runtime.Value {
	once_Effect_Aff_monoidCanceler.Do(func() {
		cache_Effect_Aff_monoidCanceler = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_semigroupCanceler()
}), Get_Effect_Aff_nonCanceler())
	})
	return cache_Effect_Aff_monoidCanceler
}

var cache_Effect_Aff_never gopurs_runtime.Value
var once_Effect_Aff_never sync.Once
func Get_Effect_Aff_never() gopurs_runtime.Value {
	once_Effect_Aff_never.Do(func() {
		cache_Effect_Aff_never = Call_Effect_Aff_makeAff(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Get_Effect_Aff_monoidCanceler(), "mempty")
})
}))
	})
	return cache_Effect_Aff_never
}

var cache_Effect_Aff_apathize gopurs_runtime.Value
var once_Effect_Aff_apathize sync.Once
func Get_Effect_Aff_apathize() gopurs_runtime.Value {
	once_Effect_Aff_apathize.Do(func() {
		cache_Effect_Aff_apathize = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_functorAff(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Apply(Get_Effect_Aff_attempt(), x_1))
})
}()
	})
	return cache_Effect_Aff_apathize
}

var cache_Effect_Aff_altParAff gopurs_runtime.Value
var once_Effect_Aff_altParAff sync.Once
func Get_Effect_Aff_altParAff() gopurs_runtime.Value {
	once_Effect_Aff_altParAff.Do(func() {
		cache_Effect_Aff_altParAff = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorParAff()
}), Get_Effect_Aff__parAffAlt())
	})
	return cache_Effect_Aff_altParAff
}

var cache_Effect_Aff_altAff gopurs_runtime.Value
var once_Effect_Aff_altAff sync.Once
func Get_Effect_Aff_altAff() gopurs_runtime.Value {
	once_Effect_Aff_altAff.Do(func() {
		cache_Effect_Aff_altAff = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorAff()
}), gopurs_runtime.Func(func(a1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_monadErrorAff(), "catchError"), a1_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a2_1
}))
})
}))
	})
	return cache_Effect_Aff_altAff
}

var cache_Effect_Aff_plusAff gopurs_runtime.Value
var once_Effect_Aff_plusAff sync.Once
func Get_Effect_Aff_plusAff() gopurs_runtime.Value {
	once_Effect_Aff_plusAff.Do(func() {
		cache_Effect_Aff_plusAff = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_altAff()
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadThrowAff(), "throwError"), gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.Str("Always fails"))))
	})
	return cache_Effect_Aff_plusAff
}

var cache_Effect_Aff_plusParAff gopurs_runtime.Value
var once_Effect_Aff_plusParAff sync.Once
func Get_Effect_Aff_plusParAff() gopurs_runtime.Value {
	once_Effect_Aff_plusParAff.Do(func() {
		cache_Effect_Aff_plusParAff = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_altParAff()
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_parallelAff(), "parallel"), gopurs_runtime.RecordGet(Get_Effect_Aff_plusAff(), "empty")))
	})
	return cache_Effect_Aff_plusParAff
}

var cache_Effect_Aff_alternativeParAff gopurs_runtime.Value
var once_Effect_Aff_alternativeParAff sync.Once
func Get_Effect_Aff_alternativeParAff() gopurs_runtime.Value {
	once_Effect_Aff_alternativeParAff.Do(func() {
		cache_Effect_Aff_alternativeParAff = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applicativeParAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_plusParAff()
}))
	})
	return cache_Effect_Aff_alternativeParAff
}

var cache_Effect_Aff_altAff__154760964 gopurs_runtime.Value
var once_Effect_Aff_altAff__154760964 sync.Once
func Get_Effect_Aff_altAff__154760964() gopurs_runtime.Value {
	once_Effect_Aff_altAff__154760964.Do(func() {
		cache_Effect_Aff_altAff__154760964 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorAff()
}), gopurs_runtime.Func(func(a1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_monadErrorAff(), "catchError"), a1_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a2_1
}))
})
}))
	})
	return cache_Effect_Aff_altAff__154760964
}

var cache_Effect_Aff_altParAff__2031255559 gopurs_runtime.Value
var once_Effect_Aff_altParAff__2031255559 sync.Once
func Get_Effect_Aff_altParAff__2031255559() gopurs_runtime.Value {
	once_Effect_Aff_altParAff__2031255559.Do(func() {
		cache_Effect_Aff_altParAff__2031255559 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorParAff()
}), Get_Effect_Aff__parAffAlt())
	})
	return cache_Effect_Aff_altParAff__2031255559
}

var cache_Effect_Aff_applicativeAff__3333162410 gopurs_runtime.Value
var once_Effect_Aff_applicativeAff__3333162410 sync.Once
func Get_Effect_Aff_applicativeAff__3333162410() gopurs_runtime.Value {
	once_Effect_Aff_applicativeAff__3333162410.Do(func() {
		cache_Effect_Aff_applicativeAff__3333162410 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), Get_Effect_Aff__pure())
	})
	return cache_Effect_Aff_applicativeAff__3333162410
}

var cache_Effect_Aff_applicativeAff__156155496 gopurs_runtime.Value
var once_Effect_Aff_applicativeAff__156155496 sync.Once
func Get_Effect_Aff_applicativeAff__156155496() gopurs_runtime.Value {
	once_Effect_Aff_applicativeAff__156155496.Do(func() {
		cache_Effect_Aff_applicativeAff__156155496 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), Get_Effect_Aff__pure())
	})
	return cache_Effect_Aff_applicativeAff__156155496
}

var cache_Effect_Aff_applicativeParAff__995286821 gopurs_runtime.Value
var once_Effect_Aff_applicativeParAff__995286821 sync.Once
func Get_Effect_Aff_applicativeParAff__995286821() gopurs_runtime.Value {
	once_Effect_Aff_applicativeParAff__995286821.Do(func() {
		cache_Effect_Aff_applicativeParAff__995286821 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyParAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_parallelAff(), "parallel"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"), x_0))
}))
	})
	return cache_Effect_Aff_applicativeParAff__995286821
}

var cache_Effect_Aff_applicativeParAff__2568423465 gopurs_runtime.Value
var once_Effect_Aff_applicativeParAff__2568423465 sync.Once
func Get_Effect_Aff_applicativeParAff__2568423465() gopurs_runtime.Value {
	once_Effect_Aff_applicativeParAff__2568423465.Do(func() {
		cache_Effect_Aff_applicativeParAff__2568423465 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyParAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_parallelAff(), "parallel"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"), x_0))
}))
	})
	return cache_Effect_Aff_applicativeParAff__2568423465
}

var cache_Effect_Aff_applicativeParAff__2496133224 gopurs_runtime.Value
var once_Effect_Aff_applicativeParAff__2496133224 sync.Once
func Get_Effect_Aff_applicativeParAff__2496133224() gopurs_runtime.Value {
	once_Effect_Aff_applicativeParAff__2496133224.Do(func() {
		cache_Effect_Aff_applicativeParAff__2496133224 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyParAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_parallelAff(), "parallel"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"), x_0))
}))
	})
	return cache_Effect_Aff_applicativeParAff__2496133224
}

var cache_Effect_Aff_applyAff__4077982506 gopurs_runtime.Value
var once_Effect_Aff_applyAff__4077982506 sync.Once
func Get_Effect_Aff_applyAff__4077982506() gopurs_runtime.Value {
	once_Effect_Aff_applyAff__4077982506.Do(func() {
		cache_Effect_Aff_applyAff__4077982506 = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadAff(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadAff(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorAff()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Effect_Aff_applyAff__4077982506
}

var cache_Effect_Aff_applyAff__2964533948 gopurs_runtime.Value
var once_Effect_Aff_applyAff__2964533948 sync.Once
func Get_Effect_Aff_applyAff__2964533948() gopurs_runtime.Value {
	once_Effect_Aff_applyAff__2964533948.Do(func() {
		cache_Effect_Aff_applyAff__2964533948 = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadAff(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadAff(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorAff()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Effect_Aff_applyAff__2964533948
}

var cache_Effect_Aff_applyFiber__166674623 gopurs_runtime.Value
var once_Effect_Aff_applyFiber__166674623 sync.Once
func Get_Effect_Aff_applyFiber__166674623() gopurs_runtime.Value {
	once_Effect_Aff_applyFiber__166674623.Do(func() {
		cache_Effect_Aff_applyFiber__166674623 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorFiber()
}), gopurs_runtime.Func(func(t1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), Call_Effect_Aff_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_applyAff(), "apply"), Call_Effect_Aff_joinFiber(t1_0), Call_Effect_Aff_joinFiber(t2_1))))
})
}))
	})
	return cache_Effect_Aff_applyFiber__166674623
}

var cache_Effect_Aff_applyParAff__2385036585 gopurs_runtime.Value
var once_Effect_Aff_applyParAff__2385036585 sync.Once
func Get_Effect_Aff_applyParAff__2385036585() gopurs_runtime.Value {
	once_Effect_Aff_applyParAff__2385036585.Do(func() {
		cache_Effect_Aff_applyParAff__2385036585 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorParAff()
}), Get_Effect_Aff__parAffApply())
	})
	return cache_Effect_Aff_applyParAff__2385036585
}

var cache_Effect_Aff_applyParAff__3038657279 gopurs_runtime.Value
var once_Effect_Aff_applyParAff__3038657279 sync.Once
func Get_Effect_Aff_applyParAff__3038657279() gopurs_runtime.Value {
	once_Effect_Aff_applyParAff__3038657279.Do(func() {
		cache_Effect_Aff_applyParAff__3038657279 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_functorParAff()
}), Get_Effect_Aff__parAffApply())
	})
	return cache_Effect_Aff_applyParAff__3038657279
}

var cache_Effect_Aff_attempt__1549600275 gopurs_runtime.Value
var once_Effect_Aff_attempt__1549600275 sync.Once
func Get_Effect_Aff_attempt__1549600275() gopurs_runtime.Value {
	once_Effect_Aff_attempt__1549600275.Do(func() {
		cache_Effect_Aff_attempt__1549600275 = gopurs_runtime.Apply(Get_Control_Monad_Error_Class_try(), Get_Effect_Aff_monadErrorAff())
	})
	return cache_Effect_Aff_attempt__1549600275
}

var cache_Effect_Aff_bindAff__1273005738 gopurs_runtime.Value
var once_Effect_Aff_bindAff__1273005738 sync.Once
func Get_Effect_Aff_bindAff__1273005738() gopurs_runtime.Value {
	once_Effect_Aff_bindAff__1273005738.Do(func() {
		cache_Effect_Aff_bindAff__1273005738 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), Get_Effect_Aff__bind())
	})
	return cache_Effect_Aff_bindAff__1273005738
}

var cache_Effect_Aff_bindAff__1025486311 gopurs_runtime.Value
var once_Effect_Aff_bindAff__1025486311 sync.Once
func Get_Effect_Aff_bindAff__1025486311() gopurs_runtime.Value {
	once_Effect_Aff_bindAff__1025486311.Do(func() {
		cache_Effect_Aff_bindAff__1025486311 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), Get_Effect_Aff__bind())
	})
	return cache_Effect_Aff_bindAff__1025486311
}

var cache_Effect_Aff_bracket__3747730269 gopurs_runtime.Value
var once_Effect_Aff_bracket__3747730269 sync.Once
func Get_Effect_Aff_bracket__3747730269() gopurs_runtime.Value {
	once_Effect_Aff_bracket__3747730269.Do(func() {
		cache_Effect_Aff_bracket__3747730269 = gopurs_runtime.Func2(func(acquire_0_box gopurs_runtime.Value, completed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_bracket__3747730269(acquire_0_box, completed_1_box)
})
	})
	return cache_Effect_Aff_bracket__3747730269
}

var cache_Effect_Aff_bracket__967388557 gopurs_runtime.Value
var once_Effect_Aff_bracket__967388557 sync.Once
func Get_Effect_Aff_bracket__967388557() gopurs_runtime.Value {
	once_Effect_Aff_bracket__967388557.Do(func() {
		cache_Effect_Aff_bracket__967388557 = gopurs_runtime.Func2(func(acquire_0_box gopurs_runtime.Value, completed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_bracket__967388557(acquire_0_box, completed_1_box)
})
	})
	return cache_Effect_Aff_bracket__967388557
}

var cache_Effect_Aff_functorAff__1039414525 gopurs_runtime.Value
var once_Effect_Aff_functorAff__1039414525 sync.Once
func Get_Effect_Aff_functorAff__1039414525() gopurs_runtime.Value {
	once_Effect_Aff_functorAff__1039414525.Do(func() {
		cache_Effect_Aff_functorAff__1039414525 = gopurs_runtime.RecordDict1("map", Get_Effect_Aff__map())
	})
	return cache_Effect_Aff_functorAff__1039414525
}

var cache_Effect_Aff_functorAff__2378915857 gopurs_runtime.Value
var once_Effect_Aff_functorAff__2378915857 sync.Once
func Get_Effect_Aff_functorAff__2378915857() gopurs_runtime.Value {
	once_Effect_Aff_functorAff__2378915857.Do(func() {
		cache_Effect_Aff_functorAff__2378915857 = gopurs_runtime.RecordDict1("map", Get_Effect_Aff__map())
	})
	return cache_Effect_Aff_functorAff__2378915857
}

var cache_Effect_Aff_functorFiber__1732109553 gopurs_runtime.Value
var once_Effect_Aff_functorFiber__1732109553 sync.Once
func Get_Effect_Aff_functorFiber__1732109553() gopurs_runtime.Value {
	once_Effect_Aff_functorFiber__1732109553.Do(func() {
		cache_Effect_Aff_functorFiber__1732109553 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), Call_Effect_Aff_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_functorAff(), "map"), f_0, Call_Effect_Aff_joinFiber(t_1))))
})
}))
	})
	return cache_Effect_Aff_functorFiber__1732109553
}

var cache_Effect_Aff_functorParAff__4103318257 gopurs_runtime.Value
var once_Effect_Aff_functorParAff__4103318257 sync.Once
func Get_Effect_Aff_functorParAff__4103318257() gopurs_runtime.Value {
	once_Effect_Aff_functorParAff__4103318257.Do(func() {
		cache_Effect_Aff_functorParAff__4103318257 = gopurs_runtime.RecordDict1("map", Get_Effect_Aff__parAffMap())
	})
	return cache_Effect_Aff_functorParAff__4103318257
}

var cache_Effect_Aff_joinFiber__1248077776 gopurs_runtime.Value
var once_Effect_Aff_joinFiber__1248077776 sync.Once
func Get_Effect_Aff_joinFiber__1248077776() gopurs_runtime.Value {
	once_Effect_Aff_joinFiber__1248077776.Do(func() {
		cache_Effect_Aff_joinFiber__1248077776 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_joinFiber__1248077776(v_0_box)
})
	})
	return cache_Effect_Aff_joinFiber__1248077776
}

var cache_Effect_Aff_joinFiber__244086667 gopurs_runtime.Value
var once_Effect_Aff_joinFiber__244086667 sync.Once
func Get_Effect_Aff_joinFiber__244086667() gopurs_runtime.Value {
	once_Effect_Aff_joinFiber__244086667.Do(func() {
		cache_Effect_Aff_joinFiber__244086667 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_joinFiber__244086667(v_0_box)
})
	})
	return cache_Effect_Aff_joinFiber__244086667
}

var cache_Effect_Aff_joinFiber__1440991555 gopurs_runtime.Value
var once_Effect_Aff_joinFiber__1440991555 sync.Once
func Get_Effect_Aff_joinFiber__1440991555() gopurs_runtime.Value {
	once_Effect_Aff_joinFiber__1440991555.Do(func() {
		cache_Effect_Aff_joinFiber__1440991555 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_joinFiber__1440991555(v_0_box)
})
	})
	return cache_Effect_Aff_joinFiber__1440991555
}

var cache_Effect_Aff_killFiber__2435668841 gopurs_runtime.Value
var once_Effect_Aff_killFiber__2435668841 sync.Once
func Get_Effect_Aff_killFiber__2435668841() gopurs_runtime.Value {
	once_Effect_Aff_killFiber__2435668841.Do(func() {
		cache_Effect_Aff_killFiber__2435668841 = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_killFiber__2435668841(e_0_box, v_1_box)
})
	})
	return cache_Effect_Aff_killFiber__2435668841
}

var cache_Effect_Aff_killFiber__991707090 gopurs_runtime.Value
var once_Effect_Aff_killFiber__991707090 sync.Once
func Get_Effect_Aff_killFiber__991707090() gopurs_runtime.Value {
	once_Effect_Aff_killFiber__991707090.Do(func() {
		cache_Effect_Aff_killFiber__991707090 = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_killFiber__991707090(e_0_box, v_1_box)
})
	})
	return cache_Effect_Aff_killFiber__991707090
}

var cache_Effect_Aff_launchAff__227652174 gopurs_runtime.Value
var once_Effect_Aff_launchAff__227652174 sync.Once
func Get_Effect_Aff_launchAff__227652174() gopurs_runtime.Value {
	once_Effect_Aff_launchAff__227652174.Do(func() {
		cache_Effect_Aff_launchAff__227652174 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_launchAff__227652174(aff_0_box)
})
	})
	return cache_Effect_Aff_launchAff__227652174
}

var cache_Effect_Aff_launchSuspendedAff__227652174 gopurs_runtime.Value
var once_Effect_Aff_launchSuspendedAff__227652174 sync.Once
func Get_Effect_Aff_launchSuspendedAff__227652174() gopurs_runtime.Value {
	once_Effect_Aff_launchSuspendedAff__227652174.Do(func() {
		cache_Effect_Aff_launchSuspendedAff__227652174 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_launchSuspendedAff__227652174(aff_0_box)
})
	})
	return cache_Effect_Aff_launchSuspendedAff__227652174
}

var cache_Effect_Aff_makeAff__3447620704 gopurs_runtime.Value
var once_Effect_Aff_makeAff__3447620704 sync.Once
func Get_Effect_Aff_makeAff__3447620704() gopurs_runtime.Value {
	once_Effect_Aff_makeAff__3447620704.Do(func() {
		cache_Effect_Aff_makeAff__3447620704 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeAff__3447620704(build_0_box)
})
	})
	return cache_Effect_Aff_makeAff__3447620704
}

var cache_Effect_Aff_makeAff__3958971776 gopurs_runtime.Value
var once_Effect_Aff_makeAff__3958971776 sync.Once
func Get_Effect_Aff_makeAff__3958971776() gopurs_runtime.Value {
	once_Effect_Aff_makeAff__3958971776.Do(func() {
		cache_Effect_Aff_makeAff__3958971776 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeAff__3958971776(build_0_box)
})
	})
	return cache_Effect_Aff_makeAff__3958971776
}

var cache_Effect_Aff_makeAff__829681120 gopurs_runtime.Value
var once_Effect_Aff_makeAff__829681120 sync.Once
func Get_Effect_Aff_makeAff__829681120() gopurs_runtime.Value {
	once_Effect_Aff_makeAff__829681120.Do(func() {
		cache_Effect_Aff_makeAff__829681120 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeAff__829681120(build_0_box)
})
	})
	return cache_Effect_Aff_makeAff__829681120
}

var cache_Effect_Aff_makeFiber__2414720213 gopurs_runtime.Value
var once_Effect_Aff_makeFiber__2414720213 sync.Once
func Get_Effect_Aff_makeFiber__2414720213() gopurs_runtime.Value {
	once_Effect_Aff_makeFiber__2414720213.Do(func() {
		cache_Effect_Aff_makeFiber__2414720213 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeFiber__2414720213(aff_0_box)
})
	})
	return cache_Effect_Aff_makeFiber__2414720213
}

var cache_Effect_Aff_makeFiber__4185835653 gopurs_runtime.Value
var once_Effect_Aff_makeFiber__4185835653 sync.Once
func Get_Effect_Aff_makeFiber__4185835653() gopurs_runtime.Value {
	once_Effect_Aff_makeFiber__4185835653.Do(func() {
		cache_Effect_Aff_makeFiber__4185835653 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeFiber__4185835653(aff_0_box)
})
	})
	return cache_Effect_Aff_makeFiber__4185835653
}

var cache_Effect_Aff_monadAff__2914113427 gopurs_runtime.Value
var once_Effect_Aff_monadAff__2914113427 sync.Once
func Get_Effect_Aff_monadAff__2914113427() gopurs_runtime.Value {
	once_Effect_Aff_monadAff__2914113427.Do(func() {
		cache_Effect_Aff_monadAff__2914113427 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applicativeAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_bindAff()
}))
	})
	return cache_Effect_Aff_monadAff__2914113427
}

var cache_Effect_Aff_monadEffectAff__2194637066 gopurs_runtime.Value
var once_Effect_Aff_monadEffectAff__2194637066 sync.Once
func Get_Effect_Aff_monadEffectAff__2194637066() gopurs_runtime.Value {
	once_Effect_Aff_monadEffectAff__2194637066.Do(func() {
		cache_Effect_Aff_monadEffectAff__2194637066 = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadAff()
}), Get_Effect_Aff__liftEffect())
	})
	return cache_Effect_Aff_monadEffectAff__2194637066
}

var cache_Effect_Aff_monadEffectAff__1856968838 gopurs_runtime.Value
var once_Effect_Aff_monadEffectAff__1856968838 sync.Once
func Get_Effect_Aff_monadEffectAff__1856968838() gopurs_runtime.Value {
	once_Effect_Aff_monadEffectAff__1856968838.Do(func() {
		cache_Effect_Aff_monadEffectAff__1856968838 = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadAff()
}), Get_Effect_Aff__liftEffect())
	})
	return cache_Effect_Aff_monadEffectAff__1856968838
}

var cache_Effect_Aff_monadErrorAff__3346684269 gopurs_runtime.Value
var once_Effect_Aff_monadErrorAff__3346684269 sync.Once
func Get_Effect_Aff_monadErrorAff__3346684269() gopurs_runtime.Value {
	once_Effect_Aff_monadErrorAff__3346684269.Do(func() {
		cache_Effect_Aff_monadErrorAff__3346684269 = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadThrowAff()
}), Get_Effect_Aff__catchError())
	})
	return cache_Effect_Aff_monadErrorAff__3346684269
}

var cache_Effect_Aff_monadErrorAff__2250703981 gopurs_runtime.Value
var once_Effect_Aff_monadErrorAff__2250703981 sync.Once
func Get_Effect_Aff_monadErrorAff__2250703981() gopurs_runtime.Value {
	once_Effect_Aff_monadErrorAff__2250703981.Do(func() {
		cache_Effect_Aff_monadErrorAff__2250703981 = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadThrowAff()
}), Get_Effect_Aff__catchError())
	})
	return cache_Effect_Aff_monadErrorAff__2250703981
}

var cache_Effect_Aff_monadThrowAff__1033845923 gopurs_runtime.Value
var once_Effect_Aff_monadThrowAff__1033845923 sync.Once
func Get_Effect_Aff_monadThrowAff__1033845923() gopurs_runtime.Value {
	once_Effect_Aff_monadThrowAff__1033845923.Do(func() {
		cache_Effect_Aff_monadThrowAff__1033845923 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadAff()
}), Get_Effect_Aff__throwError())
	})
	return cache_Effect_Aff_monadThrowAff__1033845923
}

var cache_Effect_Aff_monadThrowAff__799187270 gopurs_runtime.Value
var once_Effect_Aff_monadThrowAff__799187270 sync.Once
func Get_Effect_Aff_monadThrowAff__799187270() gopurs_runtime.Value {
	once_Effect_Aff_monadThrowAff__799187270.Do(func() {
		cache_Effect_Aff_monadThrowAff__799187270 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadAff()
}), Get_Effect_Aff__throwError())
	})
	return cache_Effect_Aff_monadThrowAff__799187270
}

var cache_Effect_Aff_parallelAff__3386337330 gopurs_runtime.Value
var once_Effect_Aff_parallelAff__3386337330 sync.Once
func Get_Effect_Aff_parallelAff__3386337330() gopurs_runtime.Value {
	once_Effect_Aff_parallelAff__3386337330.Do(func() {
		cache_Effect_Aff_parallelAff__3386337330 = gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyParAff()
}), Get_Unsafe_Coerce_unsafeCoerce(), Get_Effect_Aff__sequential())
	})
	return cache_Effect_Aff_parallelAff__3386337330
}

var cache_Effect_Aff_parallelAff__959558577 gopurs_runtime.Value
var once_Effect_Aff_parallelAff__959558577 sync.Once
func Get_Effect_Aff_parallelAff__959558577() gopurs_runtime.Value {
	once_Effect_Aff_parallelAff__959558577.Do(func() {
		cache_Effect_Aff_parallelAff__959558577 = gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_applyParAff()
}), Get_Unsafe_Coerce_unsafeCoerce(), Get_Effect_Aff__sequential())
	})
	return cache_Effect_Aff_parallelAff__959558577
}

var cache_Effect_Aff_plusParAff__4391090 gopurs_runtime.Value
var once_Effect_Aff_plusParAff__4391090 sync.Once
func Get_Effect_Aff_plusParAff__4391090() gopurs_runtime.Value {
	once_Effect_Aff_plusParAff__4391090.Do(func() {
		cache_Effect_Aff_plusParAff__4391090 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_altParAff()
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_parallelAff(), "parallel"), gopurs_runtime.RecordGet(Get_Effect_Aff_plusAff(), "empty")))
	})
	return cache_Effect_Aff_plusParAff__4391090
}

var cache_Effect_Aff_runAff__2713492946 gopurs_runtime.Value
var once_Effect_Aff_runAff__2713492946 sync.Once
func Get_Effect_Aff_runAff__2713492946() gopurs_runtime.Value {
	once_Effect_Aff_runAff__2713492946.Do(func() {
		cache_Effect_Aff_runAff__2713492946 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_runAff__2713492946(k_0_box, aff_1_box)
})
	})
	return cache_Effect_Aff_runAff__2713492946
}

func Call_Effect_Aff_Fiber(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Effect_Aff_Canceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Effect_Aff_makeFiber(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), gopurs_runtime.Apply(Get_Effect_Aff__makeFiberNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), nf_1)))
}))
}

func Call_Effect_Aff_makeAff(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (either_3.Type == 9 && either_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (either_3.Type == 9 && either_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
}))
}

func Call_Effect_Aff_launchSuspendedAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return Call_Effect_Aff_makeFiber(aff_0)
}

func Call_Effect_Aff_launchAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), Call_Effect_Aff_makeFiber(aff_0), gopurs_runtime.Func(func(fiber_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()))}, gopurs_runtime.RecordGet(fiber_1, "run"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), fiber_1)
}))
}))
}

func Call_Effect_Aff_launchAff_(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Effect_Aff_void(), Call_Effect_Aff_launchAff(x_0))
}

func Call_Effect_Aff_delay(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.UncurriedApp2(Get_Effect_Aff__delay(), Get_Data_Either_Right(), gopurs_runtime.Float(v_0))
}

func Call_Effect_Aff_bracket(acquire_0_loop gopurs_runtime.Value, completed_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var acquire_0 gopurs_runtime.Value = acquire_0_loop
_ = acquire_0
var completed_1 gopurs_runtime.Value = completed_1_loop
_ = completed_1
return gopurs_runtime.Apply2(Get_Effect_Aff_generalBracket(), acquire_0, gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
})))
}

func Call_Effect_Aff_semigroupParAff(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applyParAff(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_applyParAff(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), __local_var_2_1, a_3), b_4)
})
}))
}

func Call_Effect_Aff_cancelWith(aff_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Applicative_pure__3514127574()
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Applicative_pure__3514127574()
}), gopurs_runtime.Func(func(e_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, e_2)
})
})), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return aff_0
}))
}

func Call_Effect_Aff_finally(fin_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fin_0 gopurs_runtime.Value = fin_0_loop
_ = fin_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
})
})), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_Effect_Aff_invincible(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit())
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_1
})
_ = __local_var_1_0
return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), a_0, gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})), gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"))
}

func Call_Effect_Aff_monoidParAff(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): semigroupParAff1_1_0 -> gopurs_runtime.Value
semigroupParAff1_1_0 := Call_Effect_Aff_semigroupParAff(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupParAff1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupParAff1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeParAff(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Call_Effect_Aff_semigroupAff(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applyAff(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_applyAff(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), __local_var_2_1, a_3), b_4)
})
}))
}

func Call_Effect_Aff_effectCanceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}

func Call_Effect_Aff_joinFiber(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return Call_Effect_Aff_makeAff(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), Get_Effect_Aff_effectCanceler()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(v_0, "join"), gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_2})})
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2})})
})))
}))
}

func Call_Effect_Aff_forkAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), gopurs_runtime.Apply(Get_Effect_Aff__forkAffNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()))}, gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), nf_1)), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"), gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), nf_1)))
}))
}))
}

func Call_Effect_Aff_killFiber(e_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.RecordGet(v_1, "isSuspended")), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (suspended_2.IntVal) != (0) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(Get_Effect_Aff_void1(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}))))
goto end_branch_0
} else {

}
}
{
__t0 = Call_Effect_Aff_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), Get_Effect_Aff_effectCanceler()), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_4})})
}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()})})
})))
}))
}
end_branch_0:
return __t0
}))
}

func Call_Effect_Aff_fiberCanceler(x_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return Call_Effect_Aff_killFiber(a_1, x_0)
}

func Call_Effect_Aff_supervise(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.Str("[Aff] Child fiber outlived parent"))
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(sup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Effect_Aff__killAll(), __local_var_1_1, gopurs_runtime.RecordGet(sup_2, "supervisor"), gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()})}))
}))
})
_ = __local_var_1_0
// TAST (Let): __local_var_1_3 -> gopurs_runtime.Value
__local_var_1_3 := gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.Str("[Aff] Child fiber outlived parent"))
_ = __local_var_1_3
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Func(func(sup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Effect_Aff__killAll(), __local_var_1_3, gopurs_runtime.RecordGet(sup_2, "supervisor"), gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()})}))
}))
})
_ = __local_var_1_2
return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), gopurs_runtime.Apply(Get_Effect_Aff__makeSupervisedFiber(), aff_0), gopurs_runtime.Func(func(sup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()))}, gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), gopurs_runtime.RecordGet(sup_1, "fiber")), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), gopurs_runtime.RecordDict2("fiber", "supervisor", gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), gopurs_runtime.RecordGet(sup_1, "fiber")), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), gopurs_runtime.RecordGet(sup_1, "fiber")), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), gopurs_runtime.RecordGet(sup_1, "fiber")), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), gopurs_runtime.RecordGet(sup_1, "fiber")), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), gopurs_runtime.RecordGet(sup_1, "fiber"))), gopurs_runtime.RecordGet(sup_1, "supervisor")))
}))
}))), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_2
}), gopurs_runtime.Func(func(err_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(sup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(Get_Control_Parallel_parTraverse_(), gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_parallelAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeParAff()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}, Get_Control_Parallel_identity(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Call_Effect_Aff_killFiber(err_1, gopurs_runtime.RecordGet(sup_2, "fiber")), Call_Effect_Aff_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Effect_Aff__killAll(), err_1, gopurs_runtime.RecordGet(sup_2, "supervisor"), gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()})}))
}))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
})), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_joinFiber(gopurs_runtime.RecordGet(x_1, "fiber"))
}))
}

func Call_Effect_Aff_suspendAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), Call_Effect_Aff_makeFiber(aff_0))
}

func Call_Effect_Aff_runAff(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return Call_Effect_Aff_launchAff(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), gopurs_runtime.Apply2(Get_Control_Monad_Error_Class_try(), gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff()))}, aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(k_0, x_2))
})))
}

func Call_Effect_Aff_runAff_(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return gopurs_runtime.Apply(Get_Effect_Aff_void(), Call_Effect_Aff_runAff(k_0, aff_1))
}

func Call_Effect_Aff_runSuspendedAff(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return Call_Effect_Aff_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), gopurs_runtime.Apply2(Get_Control_Monad_Error_Class_try(), gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff()))}, aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(k_0, x_2))
})))
}

func Call_Effect_Aff_monoidAff(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): semigroupAff1_1_0 -> gopurs_runtime.Value
semigroupAff1_1_0 := Call_Effect_Aff_semigroupAff(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupAff1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAff1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_applicativeAff(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Call_Effect_Aff_bracket__3747730269(acquire_0_loop gopurs_runtime.Value, completed_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var acquire_0 gopurs_runtime.Value = acquire_0_loop
_ = acquire_0
var completed_1 gopurs_runtime.Value = completed_1_loop
_ = completed_1
return gopurs_runtime.Apply2(Get_Effect_Aff_generalBracket(), acquire_0, gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
})))
}

func Call_Effect_Aff_bracket__967388557(acquire_0_loop gopurs_runtime.Value, completed_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var acquire_0 gopurs_runtime.Value = acquire_0_loop
_ = acquire_0
var completed_1 gopurs_runtime.Value = completed_1_loop
_ = completed_1
return gopurs_runtime.Apply2(Get_Effect_Aff_generalBracket(), acquire_0, gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
})))
}

func Call_Effect_Aff_joinFiber__1248077776(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return Call_Effect_Aff_makeAff(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), Get_Effect_Aff_effectCanceler()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(v_0, "join"), gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_2})})
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2})})
})))
}))
}

func Call_Effect_Aff_joinFiber__244086667(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return Call_Effect_Aff_makeAff(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), Get_Effect_Aff_effectCanceler()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(v_0, "join"), gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_2})})
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2})})
})))
}))
}

func Call_Effect_Aff_joinFiber__1440991555(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return Call_Effect_Aff_makeAff(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), Get_Effect_Aff_effectCanceler()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(v_0, "join"), gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_2})})
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2})})
})))
}))
}

func Call_Effect_Aff_killFiber__2435668841(e_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.RecordGet(v_1, "isSuspended")), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (suspended_2.IntVal) != (0) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(Get_Effect_Aff_void1(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}))))
goto end_branch_0
} else {

}
}
{
__t0 = Call_Effect_Aff_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), Get_Effect_Aff_effectCanceler()), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_4})})
}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()})})
})))
}))
}
end_branch_0:
return __t0
}))
}

func Call_Effect_Aff_killFiber__991707090(e_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.RecordGet(v_1, "isSuspended")), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (suspended_2.IntVal) != (0) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(Get_Effect_Aff_void1(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}))))
goto end_branch_0
} else {

}
}
{
__t0 = Call_Effect_Aff_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), Get_Effect_Aff_effectCanceler()), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_4})})
}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()})})
})))
}))
}
end_branch_0:
return __t0
}))
}

func Call_Effect_Aff_launchAff__227652174(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), Call_Effect_Aff_makeFiber(aff_0), gopurs_runtime.Func(func(fiber_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()))}, gopurs_runtime.RecordGet(fiber_1, "run"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), fiber_1)
}))
}))
}

func Call_Effect_Aff_launchSuspendedAff__227652174(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return Call_Effect_Aff_makeFiber(aff_0)
}

func Call_Effect_Aff_makeAff__3447620704(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (either_3.Type == 9 && either_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (either_3.Type == 9 && either_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
}))
}

func Call_Effect_Aff_makeAff__3958971776(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (either_3.Type == 9 && either_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (either_3.Type == 9 && either_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
}))
}

func Call_Effect_Aff_makeAff__829681120(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(Get_Effect_Aff__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (either_3.Type == 9 && either_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(onError_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (either_3.Type == 9 && either_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(onSuccess_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
}))
}

func Call_Effect_Aff_makeFiber__2414720213(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), gopurs_runtime.Apply(Get_Effect_Aff__makeFiberNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), nf_1)))
}))
}

func Call_Effect_Aff_makeFiber__4185835653(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), gopurs_runtime.Apply(Get_Effect_Aff__makeFiberNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get_Effect_Aff__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__joinFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__killFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get_Effect_Aff__runFiber(), nf_1)))
}))
}

func Call_Effect_Aff_runAff__2713492946(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return Call_Effect_Aff_launchAff(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), gopurs_runtime.Apply2(Get_Control_Monad_Error_Class_try(), gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff()))}, aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(k_0, x_2))
})))
}

func Get_Effect_Aff__bind() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Bind
}

func Get_Effect_Aff__catchError() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__CatchError
}

func Get_Effect_Aff__delay() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Delay
}

func Get_Effect_Aff__forkAffNative() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ForkAffNative
}

func Get_Effect_Aff__isSuspendedFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__IsSuspendedFiber
}

func Get_Effect_Aff__joinFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__JoinFiber
}

func Get_Effect_Aff__killAll() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__KillAll
}

func Get_Effect_Aff__killFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__KillFiber
}

func Get_Effect_Aff__liftEffect() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__LiftEffect
}

func Get_Effect_Aff__makeAffImpl() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__MakeAffImpl
}

func Get_Effect_Aff__makeFiberNative() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__MakeFiberNative
}

func Get_Effect_Aff__makeSupervisedFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__MakeSupervisedFiber
}

func Get_Effect_Aff__map() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Map
}

func Get_Effect_Aff__onCompleteFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__OnCompleteFiber
}

func Get_Effect_Aff__parAffAlt() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ParAffAlt
}

func Get_Effect_Aff__parAffApply() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ParAffApply
}

func Get_Effect_Aff__parAffMap() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ParAffMap
}

func Get_Effect_Aff__pure() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Pure
}

func Get_Effect_Aff__runFiber() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__RunFiber
}

func Get_Effect_Aff__sequential() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__Sequential
}

func Get_Effect_Aff__throwError() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff__ThrowError
}

func Get_Effect_Aff_generalBracket() gopurs_runtime.Value {
	return _Gopurs_Effect_Aff_GeneralBracket
}
