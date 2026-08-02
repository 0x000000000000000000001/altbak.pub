package Effect_Aff

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Effect "gopurs/output/Effect"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Control_Parallel "gopurs/output/Control.Parallel"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Effect_Unsafe "gopurs/output/Effect.Unsafe"
	pkg_Control_Monad_ST_Class "gopurs/output/Control.Monad.ST.Class"
	pkg_Control_Monad_Error_Class "gopurs/output/Control.Monad.Error.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
	pkg_Data_Either "gopurs/output/Data.Either"
	unsafe "unsafe"
)

var cache_discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		cache_discard = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard
}

var cache_void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		cache_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void
}

var cache_Canceler gopurs_runtime.Value
var once_Canceler sync.Once
func Get_Canceler() gopurs_runtime.Value {
	once_Canceler.Do(func() {
		cache_Canceler = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Canceler(x_0_box)
})
	})
	return cache_Canceler
}

var cache_suspendAff gopurs_runtime.Value
var once_suspendAff sync.Once
func Get_suspendAff() gopurs_runtime.Value {
	once_suspendAff.Do(func() {
		cache_suspendAff = gopurs_runtime.Apply(Get__fork(), gopurs_runtime.Bool(false))
	})
	return cache_suspendAff
}

var cache_newtypeCanceler gopurs_runtime.Value
var once_newtypeCanceler sync.Once
func Get_newtypeCanceler() gopurs_runtime.Value {
	once_newtypeCanceler.Do(func() {
		cache_newtypeCanceler = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCanceler
}

var cache_makeAff gopurs_runtime.Value
var once_makeAff sync.Once
func Get_makeAff() gopurs_runtime.Value {
	once_makeAff.Do(func() {
		cache_makeAff = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff(build_0_box)
})
	})
	return cache_makeAff
}

var cache_launchSuspendedAff gopurs_runtime.Value
var once_launchSuspendedAff sync.Once
func Get_launchSuspendedAff() gopurs_runtime.Value {
	once_launchSuspendedAff.Do(func() {
		cache_launchSuspendedAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_launchSuspendedAff(aff_0_box)
})
	})
	return cache_launchSuspendedAff
}

var cache_launchAff gopurs_runtime.Value
var once_launchAff sync.Once
func Get_launchAff() gopurs_runtime.Value {
	once_launchAff.Do(func() {
		cache_launchAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_launchAff(aff_0_box)
})
	})
	return cache_launchAff
}

var cache_launchAff_ gopurs_runtime.Value
var once_launchAff_ sync.Once
func Get_launchAff_() gopurs_runtime.Value {
	once_launchAff_.Do(func() {
		cache_launchAff_ = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_launchAff_(x_0_box)
})
	})
	return cache_launchAff_
}

var cache_functorParAff gopurs_runtime.Value
var once_functorParAff sync.Once
func Get_functorParAff() gopurs_runtime.Value {
	once_functorParAff.Do(func() {
		cache_functorParAff = gopurs_runtime.RecordDict1("map", Get__parAffMap())
	})
	return cache_functorParAff
}

var cache_functorAff gopurs_runtime.Value
var once_functorAff sync.Once
func Get_functorAff() gopurs_runtime.Value {
	once_functorAff.Do(func() {
		cache_functorAff = gopurs_runtime.RecordDict1("map", Get__map())
	})
	return cache_functorAff
}

var cache_forkAff gopurs_runtime.Value
var once_forkAff sync.Once
func Get_forkAff() gopurs_runtime.Value {
	once_forkAff.Do(func() {
		cache_forkAff = gopurs_runtime.Apply(Get__fork(), gopurs_runtime.Bool(true))
	})
	return cache_forkAff
}

var cache_delay gopurs_runtime.Value
var once_delay sync.Once
func Get_delay() gopurs_runtime.Value {
	once_delay.Do(func() {
		cache_delay = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delay(v_0_box.FloatVal())
})
	})
	return cache_delay
}

var cache_bracket gopurs_runtime.Value
var once_bracket sync.Once
func Get_bracket() gopurs_runtime.Value {
	once_bracket.Do(func() {
		cache_bracket = gopurs_runtime.Func2(func(acquire_0_box gopurs_runtime.Value, completed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bracket(acquire_0_box, completed_1_box)
})
	})
	return cache_bracket
}

var cache_bracket__gopurs_runtime_Value_3112645395 gopurs_runtime.Value
var once_bracket__gopurs_runtime_Value_3112645395 sync.Once
func Get_bracket__gopurs_runtime_Value_3112645395() gopurs_runtime.Value {
	once_bracket__gopurs_runtime_Value_3112645395.Do(func() {
		cache_bracket__gopurs_runtime_Value_3112645395 = gopurs_runtime.Func2(func(acquire_0_box gopurs_runtime.Value, completed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bracket__gopurs_runtime_Value_3112645395(acquire_0_box, completed_1_box)
})
	})
	return cache_bracket__gopurs_runtime_Value_3112645395
}

var cache_applyParAff gopurs_runtime.Value
var once_applyParAff sync.Once
func Get_applyParAff() gopurs_runtime.Value {
	once_applyParAff.Do(func() {
		cache_applyParAff = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorParAff()
}), Get__parAffApply())
	})
	return cache_applyParAff
}

var cache_lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		cache_lift2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_lift2
}

var cache_semigroupParAff gopurs_runtime.Value
var once_semigroupParAff sync.Once
func Get_semigroupParAff() gopurs_runtime.Value {
	once_semigroupParAff.Do(func() {
		cache_semigroupParAff = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupParAff(dictSemigroup_0_box)
})
	})
	return cache_semigroupParAff
}

var cache_monadAff gopurs_runtime.Value
var once_monadAff sync.Once
func Get_monadAff() gopurs_runtime.Value {
	once_monadAff.Do(func() {
		cache_monadAff = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindAff()
}))
	})
	return cache_monadAff
}

var cache_bindAff gopurs_runtime.Value
var once_bindAff sync.Once
func Get_bindAff() gopurs_runtime.Value {
	once_bindAff.Do(func() {
		cache_bindAff = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), Get__bind())
	})
	return cache_bindAff
}

var cache_applyAff gopurs_runtime.Value
var once_applyAff sync.Once
func Get_applyAff() gopurs_runtime.Value {
	once_applyAff.Do(func() {
		cache_applyAff = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadAff(), "Bind1"), gopurs_runtime.Value{})
_ = __local_var_0_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAff()
}), gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "bind"), f_1, gopurs_runtime.Func(func(f_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "bind"), a_2, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadAff(), "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_prime_3, a_prime_4))
}))
}))
}))
}()
	})
	return cache_applyAff
}

var cache_applicativeAff gopurs_runtime.Value
var once_applicativeAff sync.Once
func Get_applicativeAff() gopurs_runtime.Value {
	once_applicativeAff.Do(func() {
		cache_applicativeAff = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), Get__pure())
	})
	return cache_applicativeAff
}

var cache_lift21 gopurs_runtime.Value
var once_lift21 sync.Once
func Get_lift21() gopurs_runtime.Value {
	once_lift21.Do(func() {
		cache_lift21 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift21(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_lift21
}

var cache_cancelWith gopurs_runtime.Value
var once_cancelWith sync.Once
func Get_cancelWith() gopurs_runtime.Value {
	once_cancelWith.Do(func() {
		cache_cancelWith = gopurs_runtime.Func2(func(aff_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cancelWith(aff_0_box, v_1_box)
})
	})
	return cache_cancelWith
}

var cache_finally gopurs_runtime.Value
var once_finally sync.Once
func Get_finally() gopurs_runtime.Value {
	once_finally.Do(func() {
		cache_finally = gopurs_runtime.Func2(func(fin_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_finally(fin_0_box, a_1_box)
})
	})
	return cache_finally
}

var cache_invincible gopurs_runtime.Value
var once_invincible sync.Once
func Get_invincible() gopurs_runtime.Value {
	once_invincible.Do(func() {
		cache_invincible = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_invincible(a_0_box)
})
	})
	return cache_invincible
}

var cache_lazyAff gopurs_runtime.Value
var once_lazyAff sync.Once
func Get_lazyAff() gopurs_runtime.Value {
	once_lazyAff.Do(func() {
		cache_lazyAff = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), pkg_Data_Unit.Get_unit()), f_0)
}))
	})
	return cache_lazyAff
}

var cache_parallelAff gopurs_runtime.Value
var once_parallelAff sync.Once
func Get_parallelAff() gopurs_runtime.Value {
	once_parallelAff.Do(func() {
		cache_parallelAff = gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}), pkg_Unsafe_Coerce.Get_unsafeCoerce(), Get__sequential())
	})
	return cache_parallelAff
}

var cache_applicativeParAff gopurs_runtime.Value
var once_applicativeParAff sync.Once
func Get_applicativeParAff() gopurs_runtime.Value {
	once_applicativeParAff.Do(func() {
		cache_applicativeParAff = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_parallelAff(), "parallel"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), x_0))
}))
	})
	return cache_applicativeParAff
}

var cache_parSequence_ gopurs_runtime.Value
var once_parSequence_ sync.Once
func Get_parSequence_() gopurs_runtime.Value {
	once_parSequence_.Do(func() {
		cache_parSequence_ = gopurs_runtime.Apply4(pkg_Control_Parallel.Get_parTraverse_(), Get_parallelAff(), Get_applicativeParAff(), pkg_Data_Foldable.Get_foldableArray(), pkg_Control_Parallel.Get_identity())
	})
	return cache_parSequence_
}

var cache_monoidParAff gopurs_runtime.Value
var once_monoidParAff sync.Once
func Get_monoidParAff() gopurs_runtime.Value {
	once_monoidParAff.Do(func() {
		cache_monoidParAff = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidParAff(dictMonoid_0_box)
})
	})
	return cache_monoidParAff
}

var cache_semigroupCanceler gopurs_runtime.Value
var once_semigroupCanceler sync.Once
func Get_semigroupCanceler() gopurs_runtime.Value {
	once_semigroupCanceler.Do(func() {
		cache_semigroupCanceler = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_parSequence_(), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(v_0, err_2), gopurs_runtime.Apply(v1_1, err_2)}))
}))
	})
	return cache_semigroupCanceler
}

var cache_semigroupAff gopurs_runtime.Value
var once_semigroupAff sync.Once
func Get_semigroupAff() gopurs_runtime.Value {
	once_semigroupAff.Do(func() {
		cache_semigroupAff = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupAff(dictSemigroup_0_box)
})
	})
	return cache_semigroupAff
}

var cache_monadEffectAff gopurs_runtime.Value
var once_monadEffectAff sync.Once
func Get_monadEffectAff() gopurs_runtime.Value {
	once_monadEffectAff.Do(func() {
		cache_monadEffectAff = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}), Get__liftEffect())
	})
	return cache_monadEffectAff
}

var cache_effectCanceler gopurs_runtime.Value
var once_effectCanceler sync.Once
func Get_effectCanceler() gopurs_runtime.Value {
	once_effectCanceler.Do(func() {
		cache_effectCanceler = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_effectCanceler(x_0_box)
})
	})
	return cache_effectCanceler
}

var cache_joinFiber gopurs_runtime.Value
var once_joinFiber sync.Once
func Get_joinFiber() gopurs_runtime.Value {
	once_joinFiber.Do(func() {
		cache_joinFiber = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinFiber(v_0_box)
})
	})
	return cache_joinFiber
}

var cache_functorFiber gopurs_runtime.Value
var once_functorFiber sync.Once
func Get_functorFiber() gopurs_runtime.Value {
	once_functorFiber.Do(func() {
		cache_functorFiber = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, t_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(Get__makeFiber(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorAff(), "map"), f_0, Call_joinFiber(t_1))))
}))
	})
	return cache_functorFiber
}

var cache_applyFiber gopurs_runtime.Value
var once_applyFiber sync.Once
func Get_applyFiber() gopurs_runtime.Value {
	once_applyFiber.Do(func() {
		cache_applyFiber = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorFiber()
}), gopurs_runtime.Func2(func(t1_0 gopurs_runtime.Value, t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(Get__makeFiber(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyAff(), "apply"), Call_joinFiber(t1_0), Call_joinFiber(t2_1))))
}))
	})
	return cache_applyFiber
}

var cache_applicativeFiber gopurs_runtime.Value
var once_applicativeFiber sync.Once
func Get_applicativeFiber() gopurs_runtime.Value {
	once_applicativeFiber.Do(func() {
		cache_applicativeFiber = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyFiber()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(Get__makeFiber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), a_0)))
}))
	})
	return cache_applicativeFiber
}

var cache_killFiber gopurs_runtime.Value
var once_killFiber sync.Once
func Get_killFiber() gopurs_runtime.Value {
	once_killFiber.Do(func() {
		cache_killFiber = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_killFiber(e_0_box, v_1_box)
})
	})
	return cache_killFiber
}

var cache_fiberCanceler gopurs_runtime.Value
var once_fiberCanceler sync.Once
func Get_fiberCanceler() gopurs_runtime.Value {
	once_fiberCanceler.Do(func() {
		cache_fiberCanceler = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fiberCanceler(x_0_box, a_1_box)
})
	})
	return cache_fiberCanceler
}

var cache_supervise gopurs_runtime.Value
var once_supervise sync.Once
func Get_supervise() gopurs_runtime.Value {
	once_supervise.Do(func() {
		cache_supervise = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_supervise(aff_0_box)
})
	})
	return cache_supervise
}

var cache_monadSTAff gopurs_runtime.Value
var once_monadSTAff sync.Once
func Get_monadSTAff() gopurs_runtime.Value {
	once_monadSTAff.Do(func() {
		cache_monadSTAff = gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Class.Get_monadSTEffect(), "liftST"), x_0))
}))
	})
	return cache_monadSTAff
}

var cache_monadThrowAff gopurs_runtime.Value
var once_monadThrowAff sync.Once
func Get_monadThrowAff() gopurs_runtime.Value {
	once_monadThrowAff.Do(func() {
		cache_monadThrowAff = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}), Get__throwError())
	})
	return cache_monadThrowAff
}

var cache_monadErrorAff gopurs_runtime.Value
var once_monadErrorAff sync.Once
func Get_monadErrorAff() gopurs_runtime.Value {
	once_monadErrorAff.Do(func() {
		cache_monadErrorAff = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowAff()
}), Get__catchError())
	})
	return cache_monadErrorAff
}

var cache_try gopurs_runtime.Value
var once_try sync.Once
func Get_try() gopurs_runtime.Value {
	once_try.Do(func() {
		cache_try = gopurs_runtime.Apply(pkg_Control_Monad_Error_Class.Get_try(), Get_monadErrorAff())
	})
	return cache_try
}

var cache_attempt gopurs_runtime.Value
var once_attempt sync.Once
func Get_attempt() gopurs_runtime.Value {
	once_attempt.Do(func() {
		cache_attempt = Get_try()
	})
	return cache_attempt
}

var cache_runAff gopurs_runtime.Value
var once_runAff sync.Once
func Get_runAff() gopurs_runtime.Value {
	once_runAff.Do(func() {
		cache_runAff = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runAff(k_0_box, aff_1_box)
})
	})
	return cache_runAff
}

var cache_runAff_ gopurs_runtime.Value
var once_runAff_ sync.Once
func Get_runAff_() gopurs_runtime.Value {
	once_runAff_.Do(func() {
		cache_runAff_ = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runAff_(k_0_box, aff_1_box)
})
	})
	return cache_runAff_
}

var cache_runSuspendedAff gopurs_runtime.Value
var once_runSuspendedAff sync.Once
func Get_runSuspendedAff() gopurs_runtime.Value {
	once_runSuspendedAff.Do(func() {
		cache_runSuspendedAff = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runSuspendedAff(k_0_box, aff_1_box)
})
	})
	return cache_runSuspendedAff
}

var cache_monadRecAff gopurs_runtime.Value
var once_monadRecAff sync.Once
func Get_monadRecAff() gopurs_runtime.Value {
	once_monadRecAff.Do(func() {
		cache_monadRecAff = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
go__go_1_0_0 = gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(k_0, a_2), gopurs_runtime.Func(func(res_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (res_3.Type == 9 && res_3.IntVal == 60402430) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(res_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (res_3.Type == 9 && res_3.IntVal == 525585346) {
__t1 = gopurs_runtime.Apply(go__go_1_0_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(res_3.UnsafePtr).V0)
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
	return cache_monadRecAff
}

var cache_monoidAff gopurs_runtime.Value
var once_monoidAff sync.Once
func Get_monoidAff() gopurs_runtime.Value {
	once_monoidAff.Do(func() {
		cache_monoidAff = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidAff(dictMonoid_0_box)
})
	})
	return cache_monoidAff
}

var cache_nonCanceler gopurs_runtime.Value
var once_nonCanceler sync.Once
func Get_nonCanceler() gopurs_runtime.Value {
	once_nonCanceler.Do(func() {
		cache_nonCanceler = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), pkg_Data_Unit.Get_unit())
_ = __local_var_0_0
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_0_0
})
}()
	})
	return cache_nonCanceler
}

var cache_monoidCanceler gopurs_runtime.Value
var once_monoidCanceler sync.Once
func Get_monoidCanceler() gopurs_runtime.Value {
	once_monoidCanceler.Do(func() {
		cache_monoidCanceler = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupCanceler()
}), Get_nonCanceler())
	})
	return cache_monoidCanceler
}

var cache_never gopurs_runtime.Value
var once_never sync.Once
func Get_never() gopurs_runtime.Value {
	once_never.Do(func() {
		cache_never = Call_makeAff(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.RecordGet(Get_monoidCanceler(), "mempty"))
}))
	})
	return cache_never
}

var cache_apathize gopurs_runtime.Value
var once_apathize sync.Once
func Get_apathize() gopurs_runtime.Value {
	once_apathize.Do(func() {
		cache_apathize = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorAff(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Apply(Get_try(), x_1))
})
}()
	})
	return cache_apathize
}

var cache_altParAff gopurs_runtime.Value
var once_altParAff sync.Once
func Get_altParAff() gopurs_runtime.Value {
	once_altParAff.Do(func() {
		cache_altParAff = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorParAff()
}), Get__parAffAlt())
	})
	return cache_altParAff
}

var cache_altAff gopurs_runtime.Value
var once_altAff sync.Once
func Get_altAff() gopurs_runtime.Value {
	once_altAff.Do(func() {
		cache_altAff = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAff()
}), gopurs_runtime.Func2(func(a1_0 gopurs_runtime.Value, a2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_monadErrorAff(), "catchError"), a1_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a2_1
}))
}))
	})
	return cache_altAff
}

var cache_plusAff gopurs_runtime.Value
var once_plusAff sync.Once
func Get_plusAff() gopurs_runtime.Value {
	once_plusAff.Do(func() {
		cache_plusAff = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altAff()
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadThrowAff(), "throwError"), gopurs_runtime.Apply(pkg_Effect_Exception.Get_error(), gopurs_runtime.Str("Always fails"))))
	})
	return cache_plusAff
}

var cache_plusParAff gopurs_runtime.Value
var once_plusParAff sync.Once
func Get_plusParAff() gopurs_runtime.Value {
	once_plusParAff.Do(func() {
		cache_plusParAff = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altParAff()
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_parallelAff(), "parallel"), gopurs_runtime.RecordGet(Get_plusAff(), "empty")))
	})
	return cache_plusParAff
}

var cache_alternativeParAff gopurs_runtime.Value
var once_alternativeParAff sync.Once
func Get_alternativeParAff() gopurs_runtime.Value {
	once_alternativeParAff.Do(func() {
		cache_alternativeParAff = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeParAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusParAff()
}))
	})
	return cache_alternativeParAff
}

func Call_Canceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_makeAff(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(Get__makeAffImpl(), gopurs_runtime.Func2(func(onError_1 gopurs_runtime.Value, onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (either_3.Type == 9 && either_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(onError_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (either_3.Type == 9 && either_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(onSuccess_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
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
}))
}

func Call_launchSuspendedAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply(Get__makeFiber(), aff_0)
}

func Call_launchAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(Get__makeFiber(), aff_0), gopurs_runtime.Func(func(fiber_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.RecordGet(fiber_1, "run"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), fiber_1)
}))
}))
}

func Call_launchAff_(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_void(), Call_launchAff(x_0))
}

func Call_delay(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.UncurriedApp2(Get__delay(), pkg_Data_Either.Get_Right(), gopurs_runtime.Float(v_0))
}

func Call_bracket(acquire_0_loop gopurs_runtime.Value, completed_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var acquire_0 gopurs_runtime.Value = acquire_0_loop
_ = acquire_0
var completed_1 gopurs_runtime.Value = completed_1_loop
_ = completed_1
return gopurs_runtime.Apply2(Get_generalBracket(), acquire_0, gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
})))
}

func Call_bracket__gopurs_runtime_Value_3112645395(acquire_0_loop gopurs_runtime.Value, completed_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var acquire_0 gopurs_runtime.Value = acquire_0_loop
_ = acquire_0
var completed_1 gopurs_runtime.Value = completed_1_loop
_ = completed_1
return gopurs_runtime.Apply2(Get_generalBracket(), acquire_0, gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
})))
}

func Call_lift2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyParAff(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyParAff(), "Functor0"), gopurs_runtime.Value{}), "map"), f_0, a_1), b_2)
}

func Call_semigroupParAff(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.RecordGet(dictSemigroup_0, "append")))
}

func Call_lift21(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyAff(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyAff(), "Functor0"), gopurs_runtime.Value{}), "map"), f_0, a_1), b_2)
}

func Call_cancelWith(aff_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(Get_generalBracket(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), pkg_Data_Unit.Get_unit()), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Get_applicativeAff(), "pure")
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Get_applicativeAff(), "pure")
}), gopurs_runtime.Func2(func(e_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, e_2)
})), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return aff_0
}))
}

func Call_finally(fin_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fin_0 gopurs_runtime.Value = fin_0_loop
_ = fin_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply3(Get_generalBracket(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), pkg_Data_Unit.Get_unit()), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
})), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_invincible(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Apply3(Get_generalBracket(), a_0, gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})), gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"))
}

func Call_monoidParAff(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeParAff(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Call_semigroupAff(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift21(), gopurs_runtime.RecordGet(dictSemigroup_0, "append")))
}

func Call_effectCanceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}

func Call_joinFiber(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return Call_makeAff(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), Get_effectCanceler(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(v_0, "join"), gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_2})})
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2})})
})))
}))
}

func Call_killFiber(e_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.RecordGet(v_1, "isSuspended")), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (suspended_2.IntVal) != (0) {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), pkg_Data_Unit.Get_unit())
_ = __local_var_3_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), pkg_Data_Unit.Get_unit())
_ = __local_var_3_2
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_2
}))))
goto end_branch_0
} else {

}
}
{
__t0 = Call_makeAff(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), Get_effectCanceler(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_4})})
}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})})
})))
}))
}
end_branch_0:
return __t0
}))
}

func Call_fiberCanceler(x_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return Call_killFiber(a_1, x_0)
}

func Call_supervise(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
killError_1_0 := gopurs_runtime.Apply(pkg_Effect_Exception.Get_error(), gopurs_runtime.Str("[Aff] Child fiber outlived parent"))
_ = killError_1_0
return gopurs_runtime.Apply3(Get_generalBracket(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(Get__makeSupervisedFiber(), aff_0), gopurs_runtime.Func(func(sup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(sup_2, "fiber"), "run"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), sup_2)
}))
}))), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), killError_1_0, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Either.Get_applicativeEither(), "pure"), pkg_Data_Unit.Get_unit())))
}))
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), killError_1_0, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Either.Get_applicativeEither(), "pure"), pkg_Data_Unit.Get_unit())))
}))
}), gopurs_runtime.Func2(func(err_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_parSequence_(), gopurs_runtime.Array([]gopurs_runtime.Value{Call_killFiber(err_2, gopurs_runtime.RecordGet(sup_3, "fiber")), Call_makeAff(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), err_2, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Either.Get_applicativeEither(), "pure"), pkg_Data_Unit.Get_unit())))
}))}))
})), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinFiber(gopurs_runtime.RecordGet(x_2, "fiber"))
}))
}

func Call_runAff(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return Call_launchAff(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(Get_try(), aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(k_0, x_2))
})))
}

func Call_runAff_(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return gopurs_runtime.Apply(Get_void(), Call_runAff(k_0, aff_1))
}

func Call_runSuspendedAff(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return gopurs_runtime.Apply(Get__makeFiber(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(Get_try(), aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(k_0, x_2))
})))
}

func Call_monoidAff(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift21(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Get__bind() gopurs_runtime.Value {
	return _Gopurs__Bind
}

func Get__catchError() gopurs_runtime.Value {
	return _Gopurs__CatchError
}

func Get__delay() gopurs_runtime.Value {
	return _Gopurs__Delay
}

func Get__fork() gopurs_runtime.Value {
	return _Gopurs__Fork
}

func Get__killAll() gopurs_runtime.Value {
	return _Gopurs__KillAll
}

func Get__liftEffect() gopurs_runtime.Value {
	return _Gopurs__LiftEffect
}

func Get__makeAffImpl() gopurs_runtime.Value {
	return _Gopurs__MakeAffImpl
}

func Get__makeFiber() gopurs_runtime.Value {
	return _Gopurs__MakeFiber
}

func Get__makeSupervisedFiber() gopurs_runtime.Value {
	return _Gopurs__MakeSupervisedFiber
}

func Get__map() gopurs_runtime.Value {
	return _Gopurs__Map
}

func Get__parAffAlt() gopurs_runtime.Value {
	return _Gopurs__ParAffAlt
}

func Get__parAffApply() gopurs_runtime.Value {
	return _Gopurs__ParAffApply
}

func Get__parAffMap() gopurs_runtime.Value {
	return _Gopurs__ParAffMap
}

func Get__pure() gopurs_runtime.Value {
	return _Gopurs__Pure
}

func Get__sequential() gopurs_runtime.Value {
	return _Gopurs__Sequential
}

func Get__throwError() gopurs_runtime.Value {
	return _Gopurs__ThrowError
}

func Get_generalBracket() gopurs_runtime.Value {
	return _Gopurs_GeneralBracket
}
