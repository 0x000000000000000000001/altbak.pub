package Effect_Aff

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Error_Class "gopurs/output/Control.Monad.Error.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Monad_ST_Class "gopurs/output/Control.Monad.ST.Class"
	pkg_Control_Parallel "gopurs/output/Control.Parallel"
	pkg_Control_Parallel_Class "gopurs/output/Control.Parallel.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Class "gopurs/output/Effect.Class"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
	pkg_Effect_Unsafe "gopurs/output/Effect.Unsafe"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_void1 gopurs_runtime.Value
var once_void1 sync.Once
func Get_void1() gopurs_runtime.Value {
	once_void1.Do(func() {
		cache_void1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void1
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

var cache_makeFiber gopurs_runtime.Value
var once_makeFiber sync.Once
func Get_makeFiber() gopurs_runtime.Value {
	once_makeFiber.Do(func() {
		cache_makeFiber = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeFiber(aff_0_box)
})
	})
	return cache_makeFiber
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
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadAff(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadAff(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAff()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(Get_pure__3514127574(), pkg_Data_Unit.Get_unit()), f_0)
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
		cache_semigroupCanceler = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(pkg_Control_Parallel.Get_parTraverse_(), Get_parallelAff(), Get_applicativeParAff(), pkg_Data_Foldable.Get_foldableArray(), pkg_Control_Parallel.Get_identity(), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(v_0, err_2), gopurs_runtime.Apply(v1_1, err_2)}))
})
})
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
		cache_functorFiber = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), Call_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorAff(), "map"), f_0, Call_joinFiber(t_1))))
})
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
}), gopurs_runtime.Func(func(t1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), Call_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyAff(), "apply"), Call_joinFiber(t1_0), Call_joinFiber(t2_1))))
})
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
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), Call_makeFiber(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), a_0)))
}))
	})
	return cache_applicativeFiber
}

var cache_forkAff gopurs_runtime.Value
var once_forkAff sync.Once
func Get_forkAff() gopurs_runtime.Value {
	once_forkAff.Do(func() {
		cache_forkAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_forkAff(aff_0_box)
})
	})
	return cache_forkAff
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

var cache_suspendAff gopurs_runtime.Value
var once_suspendAff sync.Once
func Get_suspendAff() gopurs_runtime.Value {
	once_suspendAff.Do(func() {
		cache_suspendAff = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_suspendAff(aff_0_box)
})
	})
	return cache_suspendAff
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

var cache_attempt gopurs_runtime.Value
var once_attempt sync.Once
func Get_attempt() gopurs_runtime.Value {
	once_attempt.Do(func() {
		cache_attempt = gopurs_runtime.Apply(pkg_Control_Monad_Error_Class.Get_try(), Get_monadErrorAff())
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
		cache_nonCanceler = gopurs_runtime.Apply(Get_const__3415939124(), gopurs_runtime.Apply(Get_pure__3514127574(), pkg_Data_Unit.Get_unit()))
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
return gopurs_runtime.Apply(Get_pure__2106705590(), gopurs_runtime.RecordGet(Get_monoidCanceler(), "mempty"))
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
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Apply(Get_attempt(), x_1))
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
}), gopurs_runtime.Func(func(a1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_monadErrorAff(), "catchError"), a1_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a2_1
}))
})
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

var cache_pure__3145599862 gopurs_runtime.Value
var once_pure__3145599862 sync.Once
func Get_pure__3145599862() gopurs_runtime.Value {
	once_pure__3145599862.Do(func() {
		cache_pure__3145599862 = gopurs_runtime.RecordGet(pkg_Data_Either.Get_applicativeEither(), "pure")
	})
	return cache_pure__3145599862
}

var cache_pure__3514127574 gopurs_runtime.Value
var once_pure__3514127574 sync.Once
func Get_pure__3514127574() gopurs_runtime.Value {
	once_pure__3514127574.Do(func() {
		cache_pure__3514127574 = gopurs_runtime.RecordGet(Get_applicativeAff(), "pure")
	})
	return cache_pure__3514127574
}

var cache_pure__2195681590 gopurs_runtime.Value
var once_pure__2195681590 sync.Once
func Get_pure__2195681590() gopurs_runtime.Value {
	once_pure__2195681590.Do(func() {
		cache_pure__2195681590 = gopurs_runtime.RecordGet(Get_applicativeAff(), "pure")
	})
	return cache_pure__2195681590
}

var cache_pure__3229300374 gopurs_runtime.Value
var once_pure__3229300374 sync.Once
func Get_pure__3229300374() gopurs_runtime.Value {
	once_pure__3229300374.Do(func() {
		cache_pure__3229300374 = gopurs_runtime.RecordGet(Get_applicativeAff(), "pure")
	})
	return cache_pure__3229300374
}

var cache_pure__3527452822 gopurs_runtime.Value
var once_pure__3527452822 sync.Once
func Get_pure__3527452822() gopurs_runtime.Value {
	once_pure__3527452822.Do(func() {
		cache_pure__3527452822 = gopurs_runtime.RecordGet(Get_applicativeParAff(), "pure")
	})
	return cache_pure__3527452822
}

var cache_pure__3540891798 gopurs_runtime.Value
var once_pure__3540891798 sync.Once
func Get_pure__3540891798() gopurs_runtime.Value {
	once_pure__3540891798.Do(func() {
		cache_pure__3540891798 = gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure")
	})
	return cache_pure__3540891798
}

var cache_pure__1641029622 gopurs_runtime.Value
var once_pure__1641029622 sync.Once
func Get_pure__1641029622() gopurs_runtime.Value {
	once_pure__1641029622.Do(func() {
		cache_pure__1641029622 = gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure")
	})
	return cache_pure__1641029622
}

var cache_pure__2106705590 gopurs_runtime.Value
var once_pure__2106705590 sync.Once
func Get_pure__2106705590() gopurs_runtime.Value {
	once_pure__2106705590.Do(func() {
		cache_pure__2106705590 = gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure")
	})
	return cache_pure__2106705590
}

var cache_pure__2644984438 gopurs_runtime.Value
var once_pure__2644984438 sync.Once
func Get_pure__2644984438() gopurs_runtime.Value {
	once_pure__2644984438.Do(func() {
		cache_pure__2644984438 = gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure")
	})
	return cache_pure__2644984438
}

var cache_pure__3453203222 gopurs_runtime.Value
var once_pure__3453203222 sync.Once
func Get_pure__3453203222() gopurs_runtime.Value {
	once_pure__3453203222.Do(func() {
		cache_pure__3453203222 = gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure")
	})
	return cache_pure__3453203222
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

var cache_apply__3993916842 gopurs_runtime.Value
var once_apply__3993916842 sync.Once
func Get_apply__3993916842() gopurs_runtime.Value {
	once_apply__3993916842.Do(func() {
		cache_apply__3993916842 = gopurs_runtime.RecordGet(Get_applyAff(), "apply")
	})
	return cache_apply__3993916842
}

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__1424073974 gopurs_runtime.Value
var once_lift2__1424073974 sync.Once
func Get_lift2__1424073974() gopurs_runtime.Value {
	once_lift2__1424073974.Do(func() {
		cache_lift2__1424073974 = func() gopurs_runtime.Value {
Functor0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyAff(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_0_0
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyAff(), "apply"), gopurs_runtime.Apply2(Functor0_0_0.V0, f_1, a_2), b_3)
})
})
})
}()
	})
	return cache_lift2__1424073974
}

var cache_lift2__2401097718 gopurs_runtime.Value
var once_lift2__2401097718 sync.Once
func Get_lift2__2401097718() gopurs_runtime.Value {
	once_lift2__2401097718.Do(func() {
		cache_lift2__2401097718 = func() gopurs_runtime.Value {
Functor0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyParAff(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_0_0
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyParAff(), "apply"), gopurs_runtime.Apply2(Functor0_0_0.V0, f_1, a_2), b_3)
})
})
})
}()
	})
	return cache_lift2__2401097718
}

var cache_bind__3043330631 gopurs_runtime.Value
var once_bind__3043330631 sync.Once
func Get_bind__3043330631() gopurs_runtime.Value {
	once_bind__3043330631.Do(func() {
		cache_bind__3043330631 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3043330631(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3043330631
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__1182478273 gopurs_runtime.Value
var once_bind__1182478273 sync.Once
func Get_bind__1182478273() gopurs_runtime.Value {
	once_bind__1182478273.Do(func() {
		cache_bind__1182478273 = gopurs_runtime.RecordGet(Get_bindAff(), "bind")
	})
	return cache_bind__1182478273
}

var cache_bind__1451555105 gopurs_runtime.Value
var once_bind__1451555105 sync.Once
func Get_bind__1451555105() gopurs_runtime.Value {
	once_bind__1451555105.Do(func() {
		cache_bind__1451555105 = gopurs_runtime.RecordGet(Get_bindAff(), "bind")
	})
	return cache_bind__1451555105
}

var cache_bind__3390533889 gopurs_runtime.Value
var once_bind__3390533889 sync.Once
func Get_bind__3390533889() gopurs_runtime.Value {
	once_bind__3390533889.Do(func() {
		cache_bind__3390533889 = gopurs_runtime.RecordGet(Get_bindAff(), "bind")
	})
	return cache_bind__3390533889
}

var cache_bind__882999777 gopurs_runtime.Value
var once_bind__882999777 sync.Once
func Get_bind__882999777() gopurs_runtime.Value {
	once_bind__882999777.Do(func() {
		cache_bind__882999777 = gopurs_runtime.RecordGet(Get_bindAff(), "bind")
	})
	return cache_bind__882999777
}

var cache_bind__3831761345 gopurs_runtime.Value
var once_bind__3831761345 sync.Once
func Get_bind__3831761345() gopurs_runtime.Value {
	once_bind__3831761345.Do(func() {
		cache_bind__3831761345 = gopurs_runtime.RecordGet(Get_bindAff(), "bind")
	})
	return cache_bind__3831761345
}

var cache_bind__3103164513 gopurs_runtime.Value
var once_bind__3103164513 sync.Once
func Get_bind__3103164513() gopurs_runtime.Value {
	once_bind__3103164513.Do(func() {
		cache_bind__3103164513 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__3103164513
}

var cache_bind__1281531809 gopurs_runtime.Value
var once_bind__1281531809 sync.Once
func Get_bind__1281531809() gopurs_runtime.Value {
	once_bind__1281531809.Do(func() {
		cache_bind__1281531809 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__1281531809
}

var cache_bind__4047544097 gopurs_runtime.Value
var once_bind__4047544097 sync.Once
func Get_bind__4047544097() gopurs_runtime.Value {
	once_bind__4047544097.Do(func() {
		cache_bind__4047544097 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__4047544097
}

var cache_bind__3674668417 gopurs_runtime.Value
var once_bind__3674668417 sync.Once
func Get_bind__3674668417() gopurs_runtime.Value {
	once_bind__3674668417.Do(func() {
		cache_bind__3674668417 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__3674668417
}

var cache_bindFlipped__1485397639 gopurs_runtime.Value
var once_bindFlipped__1485397639 sync.Once
func Get_bindFlipped__1485397639() gopurs_runtime.Value {
	once_bindFlipped__1485397639.Do(func() {
		cache_bindFlipped__1485397639 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1485397639(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__1485397639
}

var cache_bindFlipped__1432323457 gopurs_runtime.Value
var once_bindFlipped__1432323457 sync.Once
func Get_bindFlipped__1432323457() gopurs_runtime.Value {
	once_bindFlipped__1432323457.Do(func() {
		cache_bindFlipped__1432323457 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1432323457(b_0_box, a_1_box)
})
	})
	return cache_bindFlipped__1432323457
}

var cache_discard__2561459590 gopurs_runtime.Value
var once_discard__2561459590 sync.Once
func Get_discard__2561459590() gopurs_runtime.Value {
	once_discard__2561459590.Do(func() {
		cache_discard__2561459590 = gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard")
	})
	return cache_discard__2561459590
}

var cache_discard__3153643456 gopurs_runtime.Value
var once_discard__3153643456 sync.Once
func Get_discard__3153643456() gopurs_runtime.Value {
	once_discard__3153643456.Do(func() {
		cache_discard__3153643456 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Get_bindAff())
	})
	return cache_discard__3153643456
}

var cache_discard__2110164512 gopurs_runtime.Value
var once_discard__2110164512 sync.Once
func Get_discard__2110164512() gopurs_runtime.Value {
	once_discard__2110164512.Do(func() {
		cache_discard__2110164512 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard__2110164512
}

var cache_discard__2399711136 gopurs_runtime.Value
var once_discard__2399711136 sync.Once
func Get_discard__2399711136() gopurs_runtime.Value {
	once_discard__2399711136.Do(func() {
		cache_discard__2399711136 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard__2399711136
}

var cache_discard__2520179008 gopurs_runtime.Value
var once_discard__2520179008 sync.Once
func Get_discard__2520179008() gopurs_runtime.Value {
	once_discard__2520179008.Do(func() {
		cache_discard__2520179008 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard__2520179008
}

var cache_discard__317162198 gopurs_runtime.Value
var once_discard__317162198 sync.Once
func Get_discard__317162198() gopurs_runtime.Value {
	once_discard__317162198.Do(func() {
		cache_discard__317162198 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__317162198(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_discard__317162198
}

var cache_discardUnit__2687062302 gopurs_runtime.Value
var once_discardUnit__2687062302 sync.Once
func Get_discardUnit__2687062302() gopurs_runtime.Value {
	once_discardUnit__2687062302.Do(func() {
		cache_discardUnit__2687062302 = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_discardUnit__2687062302
}

var cache_catchError__2657403463 gopurs_runtime.Value
var once_catchError__2657403463 sync.Once
func Get_catchError__2657403463() gopurs_runtime.Value {
	once_catchError__2657403463.Do(func() {
		cache_catchError__2657403463 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catchError__2657403463(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_catchError__2657403463
}

var cache_catchError__1612922415 gopurs_runtime.Value
var once_catchError__1612922415 sync.Once
func Get_catchError__1612922415() gopurs_runtime.Value {
	once_catchError__1612922415.Do(func() {
		cache_catchError__1612922415 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catchError__1612922415(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_catchError__1612922415
}

var cache_catchError__3892322529 gopurs_runtime.Value
var once_catchError__3892322529 sync.Once
func Get_catchError__3892322529() gopurs_runtime.Value {
	once_catchError__3892322529.Do(func() {
		cache_catchError__3892322529 = gopurs_runtime.RecordGet(Get_monadErrorAff(), "catchError")
	})
	return cache_catchError__3892322529
}

var cache_throwError__237885032 gopurs_runtime.Value
var once_throwError__237885032 sync.Once
func Get_throwError__237885032() gopurs_runtime.Value {
	once_throwError__237885032.Do(func() {
		cache_throwError__237885032 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_throwError__237885032(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Error_Class.Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_throwError__237885032
}

var cache_throwError__1668092494 gopurs_runtime.Value
var once_throwError__1668092494 sync.Once
func Get_throwError__1668092494() gopurs_runtime.Value {
	once_throwError__1668092494.Do(func() {
		cache_throwError__1668092494 = gopurs_runtime.RecordGet(Get_monadThrowAff(), "throwError")
	})
	return cache_throwError__1668092494
}

var cache_try__2648905537 gopurs_runtime.Value
var once_try__2648905537 sync.Once
func Get_try__2648905537() gopurs_runtime.Value {
	once_try__2648905537.Do(func() {
		cache_try__2648905537 = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_try__2648905537(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0_box))
})
	})
	return cache_try__2648905537
}

var cache_try__214520782 gopurs_runtime.Value
var once_try__214520782 sync.Once
func Get_try__214520782() gopurs_runtime.Value {
	once_try__214520782.Do(func() {
		cache_try__214520782 = func() gopurs_runtime.Value {
Monad0_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadErrorAff(), "MonadThrow0"), gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{})
_ = Monad0_0_0
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_0_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
pure_2_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_0_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_2
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_monadErrorAff(), "catchError"), gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Either.Get_Right(), a_3), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_2, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_4})})
}))
})
}()
	})
	return cache_try__214520782
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

var cache_parallel__1734610934 gopurs_runtime.Value
var once_parallel__1734610934 sync.Once
func Get_parallel__1734610934() gopurs_runtime.Value {
	once_parallel__1734610934.Do(func() {
		cache_parallel__1734610934 = gopurs_runtime.RecordGet(Get_parallelAff(), "parallel")
	})
	return cache_parallel__1734610934
}

var cache_parSequence___1071252918 gopurs_runtime.Value
var once_parSequence___1071252918 sync.Once
func Get_parSequence___1071252918() gopurs_runtime.Value {
	once_parSequence___1071252918.Do(func() {
		cache_parSequence___1071252918 = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, dictFoldable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parSequence___1071252918(gopurs_runtime.CoerceToStruct[pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2_box))
})
	})
	return cache_parSequence___1071252918
}

var cache_parSequence___3793531865 gopurs_runtime.Value
var once_parSequence___3793531865 sync.Once
func Get_parSequence___3793531865() gopurs_runtime.Value {
	once_parSequence___3793531865.Do(func() {
		cache_parSequence___3793531865 = gopurs_runtime.Apply4(pkg_Control_Parallel.Get_parTraverse_(), Get_parallelAff(), Get_applicativeParAff(), pkg_Data_Foldable.Get_foldableArray(), pkg_Control_Parallel.Get_identity())
	})
	return cache_parSequence___3793531865
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

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
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

var cache_applicativeEither__4081990212 gopurs_runtime.Value
var once_applicativeEither__4081990212 sync.Once
func Get_applicativeEither__4081990212() gopurs_runtime.Value {
	once_applicativeEither__4081990212.Do(func() {
		cache_applicativeEither__4081990212 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_applyEither()
}), pkg_Data_Either.Get_Right())
	})
	return cache_applicativeEither__4081990212
}

var cache_applyEither__3806012498 gopurs_runtime.Value
var once_applyEither__3806012498 sync.Once
func Get_applyEither__3806012498() gopurs_runtime.Value {
	once_applyEither__3806012498.Do(func() {
		cache_applyEither__3806012498 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Either.Get_functorEither(), "map"), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_applyEither__3806012498
}

var cache_functorEither__13820179 gopurs_runtime.Value
var once_functorEither__13820179 sync.Once
func Get_functorEither__13820179() gopurs_runtime.Value {
	once_functorEither__13820179.Do(func() {
		cache_functorEither__13820179 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_functorEither__13820179
}

var cache_functorEither__1771778897 gopurs_runtime.Value
var once_functorEither__1771778897 sync.Once
func Get_functorEither__1771778897() gopurs_runtime.Value {
	once_functorEither__1771778897.Do(func() {
		cache_functorEither__1771778897 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_functorEither__1771778897
}

var cache_foldableArray__2950015754 gopurs_runtime.Value
var once_foldableArray__2950015754 sync.Once
func Get_foldableArray__2950015754() gopurs_runtime.Value {
	once_foldableArray__2950015754.Do(func() {
		cache_foldableArray__2950015754 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), pkg_Data_Foldable.Get_foldlArray(), pkg_Data_Foldable.Get_foldrArray())
	})
	return cache_foldableArray__2950015754
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

var cache_const__1426827922 gopurs_runtime.Value
var once_const__1426827922 sync.Once
func Get_const__1426827922() gopurs_runtime.Value {
	once_const__1426827922.Do(func() {
		cache_const__1426827922 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1426827922(a_0_box, v_1_box)
})
	})
	return cache_const__1426827922
}

var cache_const__2857921436 gopurs_runtime.Value
var once_const__2857921436 sync.Once
func Get_const__2857921436() gopurs_runtime.Value {
	once_const__2857921436.Do(func() {
		cache_const__2857921436 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__2857921436(a_0_box, v_1_box)
})
	})
	return cache_const__2857921436
}

var cache_const__2050378404 gopurs_runtime.Value
var once_const__2050378404 sync.Once
func Get_const__2050378404() gopurs_runtime.Value {
	once_const__2050378404.Do(func() {
		cache_const__2050378404 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__2050378404(a_0_box, v_1_box)
})
	})
	return cache_const__2050378404
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_const__3848686068 gopurs_runtime.Value
var once_const__3848686068 sync.Once
func Get_const__3848686068() gopurs_runtime.Value {
	once_const__3848686068.Do(func() {
		cache_const__3848686068 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__3848686068(a_0_box, v_1_box)
})
	})
	return cache_const__3848686068
}

var cache_const__1155968100 gopurs_runtime.Value
var once_const__1155968100 sync.Once
func Get_const__1155968100() gopurs_runtime.Value {
	once_const__1155968100.Do(func() {
		cache_const__1155968100 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1155968100(a_0_box, v_1_box)
})
	})
	return cache_const__1155968100
}

var cache_const__73052052 gopurs_runtime.Value
var once_const__73052052 sync.Once
func Get_const__73052052() gopurs_runtime.Value {
	once_const__73052052.Do(func() {
		cache_const__73052052 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__73052052(a_0_box, v_1_box)
})
	})
	return cache_const__73052052
}

var cache_const__3415939124 gopurs_runtime.Value
var once_const__3415939124 sync.Once
func Get_const__3415939124() gopurs_runtime.Value {
	once_const__3415939124.Do(func() {
		cache_const__3415939124 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__3415939124(a_0_box, v_1_box)
})
	})
	return cache_const__3415939124
}

var cache_const__2189647754 gopurs_runtime.Value
var once_const__2189647754 sync.Once
func Get_const__2189647754() gopurs_runtime.Value {
	once_const__2189647754.Do(func() {
		cache_const__2189647754 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__2189647754(a_0_box, v_1_box)
})
	})
	return cache_const__2189647754
}

var cache_const__4270360676 gopurs_runtime.Value
var once_const__4270360676 sync.Once
func Get_const__4270360676() gopurs_runtime.Value {
	once_const__4270360676.Do(func() {
		cache_const__4270360676 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4270360676(a_0_box, v_1_box)
})
	})
	return cache_const__4270360676
}

var cache_const__4189285076 gopurs_runtime.Value
var once_const__4189285076 sync.Once
func Get_const__4189285076() gopurs_runtime.Value {
	once_const__4189285076.Do(func() {
		cache_const__4189285076 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4189285076(a_0_box, v_1_box)
})
	})
	return cache_const__4189285076
}

var cache_const__3953240484 gopurs_runtime.Value
var once_const__3953240484 sync.Once
func Get_const__3953240484() gopurs_runtime.Value {
	once_const__3953240484.Do(func() {
		cache_const__3953240484 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__3953240484(a_0_box, v_1_box)
})
	})
	return cache_const__3953240484
}

var cache_const__2557237620 gopurs_runtime.Value
var once_const__2557237620 sync.Once
func Get_const__2557237620() gopurs_runtime.Value {
	once_const__2557237620.Do(func() {
		cache_const__2557237620 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__2557237620(a_0_box, v_1_box)
})
	})
	return cache_const__2557237620
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__3261866592 gopurs_runtime.Value
var once_flip__3261866592 sync.Once
func Get_flip__3261866592() gopurs_runtime.Value {
	once_flip__3261866592.Do(func() {
		cache_flip__3261866592 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3261866592(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3261866592
}

var cache_flip__2253242624 gopurs_runtime.Value
var once_flip__2253242624 sync.Once
func Get_flip__2253242624() gopurs_runtime.Value {
	once_flip__2253242624.Do(func() {
		cache_flip__2253242624 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2253242624(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2253242624
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__328307316 gopurs_runtime.Value
var once_map__328307316 sync.Once
func Get_map__328307316() gopurs_runtime.Value {
	once_map__328307316.Do(func() {
		cache_map__328307316 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__328307316(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__328307316
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__3699108444 gopurs_runtime.Value
var once_map__3699108444 sync.Once
func Get_map__3699108444() gopurs_runtime.Value {
	once_map__3699108444.Do(func() {
		cache_map__3699108444 = gopurs_runtime.RecordGet(pkg_Data_Either.Get_functorEither(), "map")
	})
	return cache_map__3699108444
}

var cache_map__339096027 gopurs_runtime.Value
var once_map__339096027 sync.Once
func Get_map__339096027() gopurs_runtime.Value {
	once_map__339096027.Do(func() {
		cache_map__339096027 = gopurs_runtime.RecordGet(Get_functorAff(), "map")
	})
	return cache_map__339096027
}

var cache_map__2177087003 gopurs_runtime.Value
var once_map__2177087003 sync.Once
func Get_map__2177087003() gopurs_runtime.Value {
	once_map__2177087003.Do(func() {
		cache_map__2177087003 = gopurs_runtime.RecordGet(Get_functorAff(), "map")
	})
	return cache_map__2177087003
}

var cache_map__3065908595 gopurs_runtime.Value
var once_map__3065908595 sync.Once
func Get_map__3065908595() gopurs_runtime.Value {
	once_map__3065908595.Do(func() {
		cache_map__3065908595 = gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map")
	})
	return cache_map__3065908595
}

var cache_map__173660595 gopurs_runtime.Value
var once_map__173660595 sync.Once
func Get_map__173660595() gopurs_runtime.Value {
	once_map__173660595.Do(func() {
		cache_map__173660595 = gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map")
	})
	return cache_map__173660595
}

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
}

var cache_altAff__154760964 gopurs_runtime.Value
var once_altAff__154760964 sync.Once
func Get_altAff__154760964() gopurs_runtime.Value {
	once_altAff__154760964.Do(func() {
		cache_altAff__154760964 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAff()
}), gopurs_runtime.Func(func(a1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_monadErrorAff(), "catchError"), a1_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a2_1
}))
})
}))
	})
	return cache_altAff__154760964
}

var cache_altParAff__2031255559 gopurs_runtime.Value
var once_altParAff__2031255559 sync.Once
func Get_altParAff__2031255559() gopurs_runtime.Value {
	once_altParAff__2031255559.Do(func() {
		cache_altParAff__2031255559 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorParAff()
}), Get__parAffAlt())
	})
	return cache_altParAff__2031255559
}

var cache_applicativeAff__3333162410 gopurs_runtime.Value
var once_applicativeAff__3333162410 sync.Once
func Get_applicativeAff__3333162410() gopurs_runtime.Value {
	once_applicativeAff__3333162410.Do(func() {
		cache_applicativeAff__3333162410 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), Get__pure())
	})
	return cache_applicativeAff__3333162410
}

var cache_applicativeAff__156155496 gopurs_runtime.Value
var once_applicativeAff__156155496 sync.Once
func Get_applicativeAff__156155496() gopurs_runtime.Value {
	once_applicativeAff__156155496.Do(func() {
		cache_applicativeAff__156155496 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), Get__pure())
	})
	return cache_applicativeAff__156155496
}

var cache_applicativeParAff__995286821 gopurs_runtime.Value
var once_applicativeParAff__995286821 sync.Once
func Get_applicativeParAff__995286821() gopurs_runtime.Value {
	once_applicativeParAff__995286821.Do(func() {
		cache_applicativeParAff__995286821 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_parallelAff(), "parallel"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), x_0))
}))
	})
	return cache_applicativeParAff__995286821
}

var cache_applicativeParAff__2568423465 gopurs_runtime.Value
var once_applicativeParAff__2568423465 sync.Once
func Get_applicativeParAff__2568423465() gopurs_runtime.Value {
	once_applicativeParAff__2568423465.Do(func() {
		cache_applicativeParAff__2568423465 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_parallelAff(), "parallel"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), x_0))
}))
	})
	return cache_applicativeParAff__2568423465
}

var cache_applicativeParAff__2496133224 gopurs_runtime.Value
var once_applicativeParAff__2496133224 sync.Once
func Get_applicativeParAff__2496133224() gopurs_runtime.Value {
	once_applicativeParAff__2496133224.Do(func() {
		cache_applicativeParAff__2496133224 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_parallelAff(), "parallel"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), x_0))
}))
	})
	return cache_applicativeParAff__2496133224
}

var cache_applyAff__4077982506 gopurs_runtime.Value
var once_applyAff__4077982506 sync.Once
func Get_applyAff__4077982506() gopurs_runtime.Value {
	once_applyAff__4077982506.Do(func() {
		cache_applyAff__4077982506 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadAff(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadAff(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAff()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyAff__4077982506
}

var cache_applyAff__2964533948 gopurs_runtime.Value
var once_applyAff__2964533948 sync.Once
func Get_applyAff__2964533948() gopurs_runtime.Value {
	once_applyAff__2964533948.Do(func() {
		cache_applyAff__2964533948 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadAff(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadAff(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAff()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyAff__2964533948
}

var cache_applyFiber__166674623 gopurs_runtime.Value
var once_applyFiber__166674623 sync.Once
func Get_applyFiber__166674623() gopurs_runtime.Value {
	once_applyFiber__166674623.Do(func() {
		cache_applyFiber__166674623 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorFiber()
}), gopurs_runtime.Func(func(t1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), Call_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyAff(), "apply"), Call_joinFiber(t1_0), Call_joinFiber(t2_1))))
})
}))
	})
	return cache_applyFiber__166674623
}

var cache_applyParAff__2385036585 gopurs_runtime.Value
var once_applyParAff__2385036585 sync.Once
func Get_applyParAff__2385036585() gopurs_runtime.Value {
	once_applyParAff__2385036585.Do(func() {
		cache_applyParAff__2385036585 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorParAff()
}), Get__parAffApply())
	})
	return cache_applyParAff__2385036585
}

var cache_applyParAff__3038657279 gopurs_runtime.Value
var once_applyParAff__3038657279 sync.Once
func Get_applyParAff__3038657279() gopurs_runtime.Value {
	once_applyParAff__3038657279.Do(func() {
		cache_applyParAff__3038657279 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorParAff()
}), Get__parAffApply())
	})
	return cache_applyParAff__3038657279
}

var cache_attempt__1549600275 gopurs_runtime.Value
var once_attempt__1549600275 sync.Once
func Get_attempt__1549600275() gopurs_runtime.Value {
	once_attempt__1549600275.Do(func() {
		cache_attempt__1549600275 = gopurs_runtime.Apply(pkg_Control_Monad_Error_Class.Get_try(), Get_monadErrorAff())
	})
	return cache_attempt__1549600275
}

var cache_bindAff__1273005738 gopurs_runtime.Value
var once_bindAff__1273005738 sync.Once
func Get_bindAff__1273005738() gopurs_runtime.Value {
	once_bindAff__1273005738.Do(func() {
		cache_bindAff__1273005738 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), Get__bind())
	})
	return cache_bindAff__1273005738
}

var cache_bindAff__1025486311 gopurs_runtime.Value
var once_bindAff__1025486311 sync.Once
func Get_bindAff__1025486311() gopurs_runtime.Value {
	once_bindAff__1025486311.Do(func() {
		cache_bindAff__1025486311 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), Get__bind())
	})
	return cache_bindAff__1025486311
}

var cache_bracket__3747730269 gopurs_runtime.Value
var once_bracket__3747730269 sync.Once
func Get_bracket__3747730269() gopurs_runtime.Value {
	once_bracket__3747730269.Do(func() {
		cache_bracket__3747730269 = gopurs_runtime.Func2(func(acquire_0_box gopurs_runtime.Value, completed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bracket__3747730269(acquire_0_box, completed_1_box)
})
	})
	return cache_bracket__3747730269
}

var cache_bracket__967388557 gopurs_runtime.Value
var once_bracket__967388557 sync.Once
func Get_bracket__967388557() gopurs_runtime.Value {
	once_bracket__967388557.Do(func() {
		cache_bracket__967388557 = gopurs_runtime.Func2(func(acquire_0_box gopurs_runtime.Value, completed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bracket__967388557(acquire_0_box, completed_1_box)
})
	})
	return cache_bracket__967388557
}

var cache_functorAff__1039414525 gopurs_runtime.Value
var once_functorAff__1039414525 sync.Once
func Get_functorAff__1039414525() gopurs_runtime.Value {
	once_functorAff__1039414525.Do(func() {
		cache_functorAff__1039414525 = gopurs_runtime.RecordDict1("map", Get__map())
	})
	return cache_functorAff__1039414525
}

var cache_functorAff__2378915857 gopurs_runtime.Value
var once_functorAff__2378915857 sync.Once
func Get_functorAff__2378915857() gopurs_runtime.Value {
	once_functorAff__2378915857.Do(func() {
		cache_functorAff__2378915857 = gopurs_runtime.RecordDict1("map", Get__map())
	})
	return cache_functorAff__2378915857
}

var cache_functorFiber__1732109553 gopurs_runtime.Value
var once_functorFiber__1732109553 sync.Once
func Get_functorFiber__1732109553() gopurs_runtime.Value {
	once_functorFiber__1732109553.Do(func() {
		cache_functorFiber__1732109553 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), Call_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorAff(), "map"), f_0, Call_joinFiber(t_1))))
})
}))
	})
	return cache_functorFiber__1732109553
}

var cache_functorParAff__4103318257 gopurs_runtime.Value
var once_functorParAff__4103318257 sync.Once
func Get_functorParAff__4103318257() gopurs_runtime.Value {
	once_functorParAff__4103318257.Do(func() {
		cache_functorParAff__4103318257 = gopurs_runtime.RecordDict1("map", Get__parAffMap())
	})
	return cache_functorParAff__4103318257
}

var cache_joinFiber__1248077776 gopurs_runtime.Value
var once_joinFiber__1248077776 sync.Once
func Get_joinFiber__1248077776() gopurs_runtime.Value {
	once_joinFiber__1248077776.Do(func() {
		cache_joinFiber__1248077776 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinFiber__1248077776(v_0_box)
})
	})
	return cache_joinFiber__1248077776
}

var cache_joinFiber__244086667 gopurs_runtime.Value
var once_joinFiber__244086667 sync.Once
func Get_joinFiber__244086667() gopurs_runtime.Value {
	once_joinFiber__244086667.Do(func() {
		cache_joinFiber__244086667 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinFiber__244086667(v_0_box)
})
	})
	return cache_joinFiber__244086667
}

var cache_joinFiber__1440991555 gopurs_runtime.Value
var once_joinFiber__1440991555 sync.Once
func Get_joinFiber__1440991555() gopurs_runtime.Value {
	once_joinFiber__1440991555.Do(func() {
		cache_joinFiber__1440991555 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinFiber__1440991555(v_0_box)
})
	})
	return cache_joinFiber__1440991555
}

var cache_killFiber__2435668841 gopurs_runtime.Value
var once_killFiber__2435668841 sync.Once
func Get_killFiber__2435668841() gopurs_runtime.Value {
	once_killFiber__2435668841.Do(func() {
		cache_killFiber__2435668841 = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_killFiber__2435668841(e_0_box, v_1_box)
})
	})
	return cache_killFiber__2435668841
}

var cache_killFiber__991707090 gopurs_runtime.Value
var once_killFiber__991707090 sync.Once
func Get_killFiber__991707090() gopurs_runtime.Value {
	once_killFiber__991707090.Do(func() {
		cache_killFiber__991707090 = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_killFiber__991707090(e_0_box, v_1_box)
})
	})
	return cache_killFiber__991707090
}

var cache_launchAff__227652174 gopurs_runtime.Value
var once_launchAff__227652174 sync.Once
func Get_launchAff__227652174() gopurs_runtime.Value {
	once_launchAff__227652174.Do(func() {
		cache_launchAff__227652174 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_launchAff__227652174(aff_0_box)
})
	})
	return cache_launchAff__227652174
}

var cache_launchSuspendedAff__227652174 gopurs_runtime.Value
var once_launchSuspendedAff__227652174 sync.Once
func Get_launchSuspendedAff__227652174() gopurs_runtime.Value {
	once_launchSuspendedAff__227652174.Do(func() {
		cache_launchSuspendedAff__227652174 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_launchSuspendedAff__227652174(aff_0_box)
})
	})
	return cache_launchSuspendedAff__227652174
}

var cache_makeAff__3447620704 gopurs_runtime.Value
var once_makeAff__3447620704 sync.Once
func Get_makeAff__3447620704() gopurs_runtime.Value {
	once_makeAff__3447620704.Do(func() {
		cache_makeAff__3447620704 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff__3447620704(build_0_box)
})
	})
	return cache_makeAff__3447620704
}

var cache_makeAff__3958971776 gopurs_runtime.Value
var once_makeAff__3958971776 sync.Once
func Get_makeAff__3958971776() gopurs_runtime.Value {
	once_makeAff__3958971776.Do(func() {
		cache_makeAff__3958971776 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff__3958971776(build_0_box)
})
	})
	return cache_makeAff__3958971776
}

var cache_makeFiber__2414720213 gopurs_runtime.Value
var once_makeFiber__2414720213 sync.Once
func Get_makeFiber__2414720213() gopurs_runtime.Value {
	once_makeFiber__2414720213.Do(func() {
		cache_makeFiber__2414720213 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeFiber__2414720213(aff_0_box)
})
	})
	return cache_makeFiber__2414720213
}

var cache_makeFiber__4185835653 gopurs_runtime.Value
var once_makeFiber__4185835653 sync.Once
func Get_makeFiber__4185835653() gopurs_runtime.Value {
	once_makeFiber__4185835653.Do(func() {
		cache_makeFiber__4185835653 = gopurs_runtime.Func(func(aff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeFiber__4185835653(aff_0_box)
})
	})
	return cache_makeFiber__4185835653
}

var cache_monadAff__2914113427 gopurs_runtime.Value
var once_monadAff__2914113427 sync.Once
func Get_monadAff__2914113427() gopurs_runtime.Value {
	once_monadAff__2914113427.Do(func() {
		cache_monadAff__2914113427 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindAff()
}))
	})
	return cache_monadAff__2914113427
}

var cache_monadEffectAff__2194637066 gopurs_runtime.Value
var once_monadEffectAff__2194637066 sync.Once
func Get_monadEffectAff__2194637066() gopurs_runtime.Value {
	once_monadEffectAff__2194637066.Do(func() {
		cache_monadEffectAff__2194637066 = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}), Get__liftEffect())
	})
	return cache_monadEffectAff__2194637066
}

var cache_monadErrorAff__3346684269 gopurs_runtime.Value
var once_monadErrorAff__3346684269 sync.Once
func Get_monadErrorAff__3346684269() gopurs_runtime.Value {
	once_monadErrorAff__3346684269.Do(func() {
		cache_monadErrorAff__3346684269 = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowAff()
}), Get__catchError())
	})
	return cache_monadErrorAff__3346684269
}

var cache_monadErrorAff__2250703981 gopurs_runtime.Value
var once_monadErrorAff__2250703981 sync.Once
func Get_monadErrorAff__2250703981() gopurs_runtime.Value {
	once_monadErrorAff__2250703981.Do(func() {
		cache_monadErrorAff__2250703981 = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowAff()
}), Get__catchError())
	})
	return cache_monadErrorAff__2250703981
}

var cache_monadThrowAff__1033845923 gopurs_runtime.Value
var once_monadThrowAff__1033845923 sync.Once
func Get_monadThrowAff__1033845923() gopurs_runtime.Value {
	once_monadThrowAff__1033845923.Do(func() {
		cache_monadThrowAff__1033845923 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}), Get__throwError())
	})
	return cache_monadThrowAff__1033845923
}

var cache_monadThrowAff__799187270 gopurs_runtime.Value
var once_monadThrowAff__799187270 sync.Once
func Get_monadThrowAff__799187270() gopurs_runtime.Value {
	once_monadThrowAff__799187270.Do(func() {
		cache_monadThrowAff__799187270 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}), Get__throwError())
	})
	return cache_monadThrowAff__799187270
}

var cache_parallelAff__3386337330 gopurs_runtime.Value
var once_parallelAff__3386337330 sync.Once
func Get_parallelAff__3386337330() gopurs_runtime.Value {
	once_parallelAff__3386337330.Do(func() {
		cache_parallelAff__3386337330 = gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}), pkg_Unsafe_Coerce.Get_unsafeCoerce(), Get__sequential())
	})
	return cache_parallelAff__3386337330
}

var cache_parallelAff__959558577 gopurs_runtime.Value
var once_parallelAff__959558577 sync.Once
func Get_parallelAff__959558577() gopurs_runtime.Value {
	once_parallelAff__959558577.Do(func() {
		cache_parallelAff__959558577 = gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}), pkg_Unsafe_Coerce.Get_unsafeCoerce(), Get__sequential())
	})
	return cache_parallelAff__959558577
}

var cache_plusParAff__4391090 gopurs_runtime.Value
var once_plusParAff__4391090 sync.Once
func Get_plusParAff__4391090() gopurs_runtime.Value {
	once_plusParAff__4391090.Do(func() {
		cache_plusParAff__4391090 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altParAff()
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_parallelAff(), "parallel"), gopurs_runtime.RecordGet(Get_plusAff(), "empty")))
	})
	return cache_plusParAff__4391090
}

var cache_runAff__2713492946 gopurs_runtime.Value
var once_runAff__2713492946 sync.Once
func Get_runAff__2713492946() gopurs_runtime.Value {
	once_runAff__2713492946.Do(func() {
		cache_runAff__2713492946 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, aff_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runAff__2713492946(k_0_box, aff_1_box)
})
	})
	return cache_runAff__2713492946
}

var cache_liftEffect__1892566677 gopurs_runtime.Value
var once_liftEffect__1892566677 sync.Once
func Get_liftEffect__1892566677() gopurs_runtime.Value {
	once_liftEffect__1892566677.Do(func() {
		cache_liftEffect__1892566677 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__1892566677(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__1892566677
}

var cache_liftEffect__574228595 gopurs_runtime.Value
var once_liftEffect__574228595 sync.Once
func Get_liftEffect__574228595() gopurs_runtime.Value {
	once_liftEffect__574228595.Do(func() {
		cache_liftEffect__574228595 = gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect")
	})
	return cache_liftEffect__574228595
}

var cache_liftEffect__3226494803 gopurs_runtime.Value
var once_liftEffect__3226494803 sync.Once
func Get_liftEffect__3226494803() gopurs_runtime.Value {
	once_liftEffect__3226494803.Do(func() {
		cache_liftEffect__3226494803 = gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect")
	})
	return cache_liftEffect__3226494803
}

var cache_liftEffect__1442411827 gopurs_runtime.Value
var once_liftEffect__1442411827 sync.Once
func Get_liftEffect__1442411827() gopurs_runtime.Value {
	once_liftEffect__1442411827.Do(func() {
		cache_liftEffect__1442411827 = gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect")
	})
	return cache_liftEffect__1442411827
}

var cache_liftEffect__88347923 gopurs_runtime.Value
var once_liftEffect__88347923 sync.Once
func Get_liftEffect__88347923() gopurs_runtime.Value {
	once_liftEffect__88347923.Do(func() {
		cache_liftEffect__88347923 = gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect")
	})
	return cache_liftEffect__88347923
}

var cache_liftEffect__2769380243 gopurs_runtime.Value
var once_liftEffect__2769380243 sync.Once
func Get_liftEffect__2769380243() gopurs_runtime.Value {
	once_liftEffect__2769380243.Do(func() {
		cache_liftEffect__2769380243 = gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect")
	})
	return cache_liftEffect__2769380243
}

var cache_applicativeEffect__284161122 gopurs_runtime.Value
var once_applicativeEffect__284161122 sync.Once
func Get_applicativeEffect__284161122() gopurs_runtime.Value {
	once_applicativeEffect__284161122.Do(func() {
		cache_applicativeEffect__284161122 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_pureE())
	})
	return cache_applicativeEffect__284161122
}

var cache_applyEffect__2014400020 gopurs_runtime.Value
var once_applyEffect__2014400020 sync.Once
func Get_applyEffect__2014400020() gopurs_runtime.Value {
	once_applyEffect__2014400020.Do(func() {
		cache_applyEffect__2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__2014400020
}

var cache_bindEffect__2113658466 gopurs_runtime.Value
var once_bindEffect__2113658466 sync.Once
func Get_bindEffect__2113658466() gopurs_runtime.Value {
	once_bindEffect__2113658466.Do(func() {
		cache_bindEffect__2113658466 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_bindE())
	})
	return cache_bindEffect__2113658466
}

var cache_functorEffect__347161653 gopurs_runtime.Value
var once_functorEffect__347161653 sync.Once
func Get_functorEffect__347161653() gopurs_runtime.Value {
	once_functorEffect__347161653.Do(func() {
		cache_functorEffect__347161653 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__347161653
}

var cache_functorEffect__3107547953 gopurs_runtime.Value
var once_functorEffect__3107547953 sync.Once
func Get_functorEffect__3107547953() gopurs_runtime.Value {
	once_functorEffect__3107547953.Do(func() {
		cache_functorEffect__3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__3107547953
}

func Call_Canceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_makeFiber(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(Get__makeFiberNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get__joinFiber(), nf_1), gopurs_runtime.Apply(Get__killFiber(), nf_1), gopurs_runtime.Apply(Get__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get__runFiber(), nf_1)))
}))
}

func Call_makeAff(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(Get__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
}))
}

func Call_launchSuspendedAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return Call_makeFiber(aff_0)
}

func Call_launchAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), Call_makeFiber(aff_0), gopurs_runtime.Func(func(fiber_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect(), gopurs_runtime.RecordGet(fiber_1, "run"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_semigroupParAff(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyParAff(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyParAff(), "apply"), gopurs_runtime.Apply2(Functor0_1_0.V0, __local_var_2_1, a_3), b_4)
})
}))
}

func Call_cancelWith(aff_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(Get_generalBracket(), gopurs_runtime.Apply(Get_pure__3514127574(), pkg_Data_Unit.Get_unit()), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Apply(Get_const__1155968100(), Get_pure__3514127574()), gopurs_runtime.Apply(Get_const__1155968100(), Get_pure__3514127574()), gopurs_runtime.Func(func(e_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, e_2)
})
})), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return aff_0
}))
}

func Call_finally(fin_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fin_0 gopurs_runtime.Value = fin_0_loop
_ = fin_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
__local_var_2_0 := gopurs_runtime.Apply(Get_const__1155968100(), fin_0)
_ = __local_var_2_0
return gopurs_runtime.Apply3(Get_generalBracket(), gopurs_runtime.Apply(Get_pure__3514127574(), pkg_Data_Unit.Get_unit()), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_0
})), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_invincible(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
__local_var_1_0 := gopurs_runtime.Apply(Get_pure__3514127574(), pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Apply3(Get_generalBracket(), a_0, gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
})), gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"))
}

func Call_monoidParAff(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
semigroupParAff1_1_0 := Call_semigroupParAff(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupParAff1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupParAff1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeParAff(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Call_semigroupAff(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyAff(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyAff(), "apply"), gopurs_runtime.Apply2(Functor0_1_0.V0, __local_var_2_1, a_3), b_4)
})
}))
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
return gopurs_runtime.Apply2(Get_map__173660595(), Get_effectCanceler(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(v_0, "join"), gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_2})})
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2})})
})))
}))
}

func Call_forkAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(Get__forkAffNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
fiber_2_0 := gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get__joinFiber(), nf_1), gopurs_runtime.Apply(Get__killFiber(), nf_1), gopurs_runtime.Apply(Get__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get__runFiber(), nf_1))
_ = fiber_2_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Get_bindAff(), gopurs_runtime.Apply(Get_liftEffect__3226494803(), gopurs_runtime.Apply(Get__runFiber(), nf_1)), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), fiber_2_0)
}))
}))
}

func Call_killFiber(e_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(Get_bind__1182478273(), gopurs_runtime.Apply(Get_liftEffect__574228595(), gopurs_runtime.RecordGet(v_1, "isSuspended")), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (suspended_2.IntVal) != (0) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(Get_void1(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Apply(Get_const__2557237620(), gopurs_runtime.Apply(Get_pure__3540891798(), pkg_Data_Unit.Get_unit())), gopurs_runtime.Apply(Get_const__3953240484(), gopurs_runtime.Apply(Get_pure__3540891798(), pkg_Data_Unit.Get_unit())))))
goto end_branch_0
} else {

}
}
{
__t0 = Call_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_map__173660595(), Get_effectCanceler(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
fiber_3_1 := gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get__isSuspendedFiber(), gopurs_runtime.RecordGet(sup_2, "fiber")), gopurs_runtime.Apply(Get__joinFiber(), gopurs_runtime.RecordGet(sup_2, "fiber")), gopurs_runtime.Apply(Get__killFiber(), gopurs_runtime.RecordGet(sup_2, "fiber")), gopurs_runtime.Apply(Get__onCompleteFiber(), gopurs_runtime.RecordGet(sup_2, "fiber")), gopurs_runtime.Apply(Get__runFiber(), gopurs_runtime.RecordGet(sup_2, "fiber")))
_ = fiber_3_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect(), gopurs_runtime.RecordGet(fiber_3_1, "run"), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.RecordDict2("fiber", "supervisor", fiber_3_1, gopurs_runtime.RecordGet(sup_2, "supervisor")))
}))
}))), gopurs_runtime.RecordDict3("completed", "failed", "killed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff__3447620704(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), killError_1_0, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(Get_pure__3145599862(), pkg_Data_Unit.Get_unit())))
}))
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff__3447620704(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), killError_1_0, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(Get_pure__3145599862(), pkg_Data_Unit.Get_unit())))
}))
})
}), gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(pkg_Control_Parallel.Get_parTraverse_(), Get_parallelAff(), Get_applicativeParAff(), pkg_Data_Foldable.Get_foldableArray(), pkg_Control_Parallel.Get_identity(), gopurs_runtime.Array([]gopurs_runtime.Value{Call_killFiber(err_2, gopurs_runtime.RecordGet(sup_3, "fiber")), Call_makeAff__3447620704(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), err_2, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(Get_pure__3145599862(), pkg_Data_Unit.Get_unit())))
}))}))
})
})), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinFiber(gopurs_runtime.RecordGet(x_2, "fiber"))
}))
}

func Call_suspendAff(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), Call_makeFiber(aff_0))
}

func Call_runAff(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return Call_launchAff(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply2(pkg_Control_Monad_Error_Class.Get_try(), Get_monadErrorAff(), aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
return Call_makeFiber(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply2(pkg_Control_Monad_Error_Class.Get_try(), Get_monadErrorAff(), aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(k_0, x_2))
})))
}

func Call_monoidAff(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
semigroupAff1_1_0 := Call_semigroupAff(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupAff1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAff1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
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

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_bind__3043330631(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bindFlipped__1485397639(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_bindFlipped__1432323457(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), a_1, b_0)
}

func Call_discard__317162198(dict_0_loop *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_catchError__2657403463(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_catchError__1612922415(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_throwError__237885032(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_try__2648905537(dictMonadError_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadError_0 *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadError_0_loop
_ = dictMonadError_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadError_0.V0, gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadError_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_Either.Get_Right(), a_4), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_5})})
}))
})
}

func Call_parallel__2242335472(dict_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_parSequence___1071252918(dictParallel_0_loop *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_2_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictParallel_0 *pkg_Control_Parallel_Class.Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
_ = dictParallel_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var dictFoldable_2 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_2_loop
_ = dictFoldable_2
return gopurs_runtime.Apply4(pkg_Control_Parallel.Get_parTraverse_(), gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer(dictParallel_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_2)}, pkg_Control_Parallel.Get_identity())
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

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
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

func Call_const__1426827922(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2857921436(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2050378404(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__3848686068(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__1155968100(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__73052052(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__3415939124(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2189647754(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__4270360676(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__4189285076(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__3953240484(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2557237620(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3261866592(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2253242624(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__328307316(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_bracket__3747730269(acquire_0_loop gopurs_runtime.Value, completed_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_bracket__967388557(acquire_0_loop gopurs_runtime.Value, completed_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_joinFiber__1248077776(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_joinFiber__244086667(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_joinFiber__1440991555(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_killFiber__2435668841(e_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(Get_liftEffect__574228595(), gopurs_runtime.RecordGet(v_1, "isSuspended")), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (suspended_2.IntVal) != (0) {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), pkg_Data_Unit.Get_unit())
_ = __local_var_3_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), pkg_Data_Unit.Get_unit())
_ = __local_var_3_1
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(Get_void1(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_0
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_1
}))))
goto end_branch_2
} else {

}
}
{
__t2 = Call_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), Get_effectCanceler(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_4})})
}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})})
})))
}))
}
end_branch_2:
return __t2
}))
}

func Call_killFiber__991707090(e_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply(Get_liftEffect__574228595(), gopurs_runtime.RecordGet(v_1, "isSuspended")), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (suspended_2.IntVal) != (0) {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), pkg_Data_Unit.Get_unit())
_ = __local_var_3_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), pkg_Data_Unit.Get_unit())
_ = __local_var_3_1
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(Get_void1(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_0
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_1
}))))
goto end_branch_2
} else {

}
}
{
__t2 = Call_makeAff__3447620704(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), Get_effectCanceler(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, err_4})})
}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})})
})))
}))
}
end_branch_2:
return __t2
}))
}

func Call_launchAff__227652174(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), Call_makeFiber(aff_0), gopurs_runtime.Func(func(fiber_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect(), gopurs_runtime.RecordGet(fiber_1, "run"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), fiber_1)
}))
}))
}

func Call_launchSuspendedAff__227652174(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return Call_makeFiber(aff_0)
}

func Call_makeAff__3447620704(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(Get__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
}))
}

func Call_makeAff__3958971776(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(Get__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
}))
}

func Call_makeFiber__2414720213(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(Get__makeFiberNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get__joinFiber(), nf_1), gopurs_runtime.Apply(Get__killFiber(), nf_1), gopurs_runtime.Apply(Get__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get__runFiber(), nf_1)))
}))
}

func Call_makeFiber__4185835653(aff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var aff_0 gopurs_runtime.Value = aff_0_loop
_ = aff_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(Get__makeFiberNative(), aff_0), gopurs_runtime.Func(func(nf_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", gopurs_runtime.Apply(Get__isSuspendedFiber(), nf_1), gopurs_runtime.Apply(Get__joinFiber(), nf_1), gopurs_runtime.Apply(Get__killFiber(), nf_1), gopurs_runtime.Apply(Get__onCompleteFiber(), nf_1), gopurs_runtime.Apply(Get__runFiber(), nf_1)))
}))
}

func Call_runAff__2713492946(k_0_loop gopurs_runtime.Value, aff_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var aff_1 gopurs_runtime.Value = aff_1_loop
_ = aff_1
return Call_launchAff(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindAff(), "bind"), gopurs_runtime.Apply2(pkg_Control_Monad_Error_Class.Get_try(), Get_monadErrorAff(), aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(k_0, x_2))
})))
}

func Call_liftEffect__1892566677(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Get__forkAffNative() gopurs_runtime.Value {
	return _Gopurs__ForkAffNative
}

func Get__isSuspendedFiber() gopurs_runtime.Value {
	return _Gopurs__IsSuspendedFiber
}

func Get__joinFiber() gopurs_runtime.Value {
	return _Gopurs__JoinFiber
}

func Get__killAll() gopurs_runtime.Value {
	return _Gopurs__KillAll
}

func Get__killFiber() gopurs_runtime.Value {
	return _Gopurs__KillFiber
}

func Get__liftEffect() gopurs_runtime.Value {
	return _Gopurs__LiftEffect
}

func Get__makeAffImpl() gopurs_runtime.Value {
	return _Gopurs__MakeAffImpl
}

func Get__makeFiberNative() gopurs_runtime.Value {
	return _Gopurs__MakeFiberNative
}

func Get__makeSupervisedFiber() gopurs_runtime.Value {
	return _Gopurs__MakeSupervisedFiber
}

func Get__map() gopurs_runtime.Value {
	return _Gopurs__Map
}

func Get__onCompleteFiber() gopurs_runtime.Value {
	return _Gopurs__OnCompleteFiber
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

func Get__runFiber() gopurs_runtime.Value {
	return _Gopurs__RunFiber
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
