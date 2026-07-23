package Effect_Aff

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Partial "gopurs/output/Partial"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Control_Parallel "gopurs/output/Control.Parallel"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Effect_Unsafe "gopurs/output/Effect.Unsafe"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
	pkg_Control_Monad_Error_Class "gopurs/output/Control.Monad.Error.Class"
)

var Canceler gopurs_runtime.Value
var once_Canceler sync.Once
func Get_Canceler() gopurs_runtime.Value {
	once_Canceler.Do(func() {
		Canceler = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Canceler
}

var suspendAff gopurs_runtime.Value
var once_suspendAff sync.Once
func Get_suspendAff() gopurs_runtime.Value {
	once_suspendAff.Do(func() {
		suspendAff = gopurs_runtime.Apply(Get__fork(), gopurs_runtime.Bool(false))
	})
	return suspendAff
}

var newtypeCanceler gopurs_runtime.Value
var once_newtypeCanceler sync.Once
func Get_newtypeCanceler() gopurs_runtime.Value {
	once_newtypeCanceler.Do(func() {
		newtypeCanceler = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeCanceler
}

var functorParAff gopurs_runtime.Value
var once_functorParAff sync.Once
func Get_functorParAff() gopurs_runtime.Value {
	once_functorParAff.Do(func() {
		functorParAff = gopurs_runtime.RecordDict1("map", Get__parAffMap())
	})
	return functorParAff
}

var functorAff gopurs_runtime.Value
var once_functorAff sync.Once
func Get_functorAff() gopurs_runtime.Value {
	once_functorAff.Do(func() {
		functorAff = gopurs_runtime.RecordDict1("map", Get__map())
	})
	return functorAff
}

var forkAff gopurs_runtime.Value
var once_forkAff sync.Once
func Get_forkAff() gopurs_runtime.Value {
	once_forkAff.Do(func() {
		forkAff = gopurs_runtime.Apply(Get__fork(), gopurs_runtime.Bool(true))
	})
	return forkAff
}

var ffiUtil gopurs_runtime.Value
var once_ffiUtil sync.Once
func Get_ffiUtil() gopurs_runtime.Value {
	once_ffiUtil.Do(func() {
		ffiUtil = gopurs_runtime.RecordDict5("isLeft", "fromLeft", "fromRight", "left", "right", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(v_0, "value0")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("unsafeFromLeft: Right"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.RecordGet(v_0, "value0")
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("unsafeFromRight: Left"))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), pkg_Data_Either.Get_Left(), pkg_Data_Either.Get_Right())
	})
	return ffiUtil
}

var makeFiber gopurs_runtime.Value
var once_makeFiber sync.Once
func Get_makeFiber() gopurs_runtime.Value {
	once_makeFiber.Do(func() {
		makeFiber = gopurs_runtime.Func(func(aff_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get__makeFiber(), Get_ffiUtil(), aff_0)
})
	})
	return makeFiber
}

var launchAff gopurs_runtime.Value
var once_launchAff sync.Once
func Get_launchAff() gopurs_runtime.Value {
	once_launchAff.Do(func() {
		launchAff = gopurs_runtime.Func(func(aff_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.UncurriedApp2(Get__makeFiber(), Get_ffiUtil(), aff_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
fiber_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
_ = fiber_2_1
_dollar__unused_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(fiber_2_1, "run"), gopurs_runtime.Value{})
_ = _dollar__unused_3_2
return fiber_2_1
})
})
	})
	return launchAff
}

var launchAff_ gopurs_runtime.Value
var once_launchAff_ sync.Once
func Get_launchAff_() gopurs_runtime.Value {
	once_launchAff_.Do(func() {
		launchAff_ = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.UncurriedApp2(Get__makeFiber(), Get_ffiUtil(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
fiber_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
_ = fiber_2_1
_dollar__unused_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(fiber_2_1, "run"), gopurs_runtime.Value{})
_ = _dollar__unused_3_2
return pkg_Data_Unit.Get_unit()
})
})
	})
	return launchAff_
}

var launchSuspendedAff gopurs_runtime.Value
var once_launchSuspendedAff sync.Once
func Get_launchSuspendedAff() gopurs_runtime.Value {
	once_launchSuspendedAff.Do(func() {
		launchSuspendedAff = Get_makeFiber()
	})
	return launchSuspendedAff
}

var delay gopurs_runtime.Value
var once_delay sync.Once
func Get_delay() gopurs_runtime.Value {
	once_delay.Do(func() {
		delay = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get__delay(), pkg_Data_Either.Get_Right(), v_0)
})
	})
	return delay
}

var bracket gopurs_runtime.Value
var once_bracket sync.Once
func Get_bracket() gopurs_runtime.Value {
	once_bracket.Do(func() {
		bracket = gopurs_runtime.Func2(func(acquire_0 gopurs_runtime.Value, completed_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_generalBracket(), acquire_0, gopurs_runtime.RecordDict3("killed", "failed", "completed", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return completed_1
})))
})
	})
	return bracket
}

var applyParAff gopurs_runtime.Value
var once_applyParAff sync.Once
func Get_applyParAff() gopurs_runtime.Value {
	once_applyParAff.Do(func() {
		applyParAff = gopurs_runtime.RecordDict2("apply", "Functor0", Get__parAffApply(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorParAff()
}))
	})
	return applyParAff
}

var lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		lift2 = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get__parAffApply(), gopurs_runtime.Apply2(Get__parAffMap(), f_0, a_1), b_2)
})
	})
	return lift2
}

var semigroupParAff gopurs_runtime.Value
var once_semigroupParAff sync.Once
func Get_semigroupParAff() gopurs_runtime.Value {
	once_semigroupParAff.Do(func() {
		semigroupParAff = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.RecordGet(dictSemigroup_0, "append")))
})
	})
	return semigroupParAff
}

var monadAff gopurs_runtime.Value
var once_monadAff sync.Once
func Get_monadAff() gopurs_runtime.Value {
	once_monadAff.Do(func() {
		monadAff = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindAff()
}))
	})
	return monadAff
}

var bindAff gopurs_runtime.Value
var once_bindAff sync.Once
func Get_bindAff() gopurs_runtime.Value {
	once_bindAff.Do(func() {
		bindAff = gopurs_runtime.RecordDict2("bind", "Apply0", Get__bind(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}))
	})
	return bindAff
}

var applyAff gopurs_runtime.Value
var once_applyAff sync.Once
func Get_applyAff() gopurs_runtime.Value {
	once_applyAff.Do(func() {
		applyAff = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get__bind(), f_0, gopurs_runtime.Func(func(f_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get__bind(), a_1, gopurs_runtime.Func(func(a_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeAff(), "pure"), gopurs_runtime.Apply(f_prime_2, a_prime_3))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAff()
}))
	})
	return applyAff
}

var applicativeAff gopurs_runtime.Value
var once_applicativeAff sync.Once
func Get_applicativeAff() gopurs_runtime.Value {
	once_applicativeAff.Do(func() {
		applicativeAff = gopurs_runtime.RecordDict2("pure", "Apply0", Get__pure(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}))
	})
	return applicativeAff
}

var lift21 gopurs_runtime.Value
var once_lift21 sync.Once
func Get_lift21() gopurs_runtime.Value {
	once_lift21.Do(func() {
		lift21 = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyAff(), "apply"), gopurs_runtime.Apply2(Get__map(), f_0, a_1), b_2)
})
	})
	return lift21
}

var cancelWith gopurs_runtime.Value
var once_cancelWith sync.Once
func Get_cancelWith() gopurs_runtime.Value {
	once_cancelWith.Do(func() {
		cancelWith = gopurs_runtime.Func2(func(aff_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_generalBracket(), gopurs_runtime.Apply(Get__pure(), pkg_Data_Unit.Get_unit()), gopurs_runtime.RecordDict3("killed", "failed", "completed", gopurs_runtime.Func2(func(e_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, e_2)
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get__pure()
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get__pure()
})), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return aff_0
}))
})
	})
	return cancelWith
}

var finally gopurs_runtime.Value
var once_finally sync.Once
func Get_finally() gopurs_runtime.Value {
	once_finally.Do(func() {
		finally = gopurs_runtime.Func2(func(fin_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_generalBracket(), gopurs_runtime.Apply(Get__pure(), pkg_Data_Unit.Get_unit()), gopurs_runtime.RecordDict3("killed", "failed", "completed", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return fin_0
})), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
})
	})
	return finally
}

var invincible gopurs_runtime.Value
var once_invincible sync.Once
func Get_invincible() gopurs_runtime.Value {
	once_invincible.Do(func() {
		invincible = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get__pure(), pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Apply3(Get_generalBracket(), a_0, gopurs_runtime.RecordDict3("killed", "failed", "completed", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})), Get__pure())
})
	})
	return invincible
}

var lazyAff gopurs_runtime.Value
var once_lazyAff sync.Once
func Get_lazyAff() gopurs_runtime.Value {
	once_lazyAff.Do(func() {
		lazyAff = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get__bind(), gopurs_runtime.Apply(Get__pure(), pkg_Data_Unit.Get_unit()), f_0)
}))
	})
	return lazyAff
}

var parallelAff gopurs_runtime.Value
var once_parallelAff sync.Once
func Get_parallelAff() gopurs_runtime.Value {
	once_parallelAff.Do(func() {
		parallelAff = gopurs_runtime.RecordDict4("parallel", "sequential", "Apply0", "Apply1", pkg_Unsafe_Coerce.Get_unsafeCoerce(), Get__sequential(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}))
	})
	return parallelAff
}

var applicativeParAff gopurs_runtime.Value
var once_applicativeParAff sync.Once
func Get_applicativeParAff() gopurs_runtime.Value {
	once_applicativeParAff.Do(func() {
		applicativeParAff = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get__pure(), x_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyParAff()
}))
	})
	return applicativeParAff
}

var parSequence_ gopurs_runtime.Value
var once_parSequence_ sync.Once
func Get_parSequence_() gopurs_runtime.Value {
	once_parSequence_.Do(func() {
		parSequence_ = gopurs_runtime.Apply4(pkg_Control_Parallel.Get_parTraverse_(), Get_parallelAff(), Get_applicativeParAff(), pkg_Data_Foldable.Get_foldableArray(), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
	})
	return parSequence_
}

var monoidParAff gopurs_runtime.Value
var once_monoidParAff sync.Once
func Get_monoidParAff() gopurs_runtime.Value {
	once_monoidParAff.Do(func() {
		monoidParAff = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(Get__pure(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")))
}))
})
	})
	return monoidParAff
}

var semigroupCanceler gopurs_runtime.Value
var once_semigroupCanceler sync.Once
func Get_semigroupCanceler() gopurs_runtime.Value {
	once_semigroupCanceler.Do(func() {
		semigroupCanceler = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, err_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_parSequence_(), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(v_0, err_2), gopurs_runtime.Apply(v1_1, err_2)}))
}))
	})
	return semigroupCanceler
}

var semigroupAff gopurs_runtime.Value
var once_semigroupAff sync.Once
func Get_semigroupAff() gopurs_runtime.Value {
	once_semigroupAff.Do(func() {
		semigroupAff = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift21(), gopurs_runtime.RecordGet(dictSemigroup_0, "append")))
})
	})
	return semigroupAff
}

var monadEffectAff gopurs_runtime.Value
var once_monadEffectAff sync.Once
func Get_monadEffectAff() gopurs_runtime.Value {
	once_monadEffectAff.Do(func() {
		monadEffectAff = gopurs_runtime.RecordDict2("liftEffect", "Monad0", Get__liftEffect(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}))
	})
	return monadEffectAff
}

var effectCanceler gopurs_runtime.Value
var once_effectCanceler sync.Once
func Get_effectCanceler() gopurs_runtime.Value {
	once_effectCanceler.Do(func() {
		effectCanceler = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get__liftEffect(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
})
	})
	return effectCanceler
}

var joinFiber gopurs_runtime.Value
var once_joinFiber sync.Once
func Get_joinFiber() gopurs_runtime.Value {
	once_joinFiber.Do(func() {
		joinFiber = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_makeAff(), gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_0, "join"), k_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = a_prime_3_1
__local_var_4_2 := gopurs_runtime.Apply(Get__liftEffect(), a_prime_3_1)
_ = __local_var_4_2
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_2
})
})
}))
})
	})
	return joinFiber
}

var functorFiber gopurs_runtime.Value
var once_functorFiber sync.Once
func Get_functorFiber() gopurs_runtime.Value {
	once_functorFiber.Do(func() {
		functorFiber = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, t_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.UncurriedApp2(Get__makeFiber(), Get_ffiUtil(), gopurs_runtime.Apply2(Get__map(), f_0, gopurs_runtime.Apply(Get_joinFiber(), t_1))))
}))
	})
	return functorFiber
}

var applyFiber gopurs_runtime.Value
var once_applyFiber sync.Once
func Get_applyFiber() gopurs_runtime.Value {
	once_applyFiber.Do(func() {
		applyFiber = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(t1_0 gopurs_runtime.Value, t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.UncurriedApp2(Get__makeFiber(), Get_ffiUtil(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyAff(), "apply"), gopurs_runtime.Apply(Get_joinFiber(), t1_0), gopurs_runtime.Apply(Get_joinFiber(), t2_1))))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorFiber()
}))
	})
	return applyFiber
}

var applicativeFiber gopurs_runtime.Value
var once_applicativeFiber sync.Once
func Get_applicativeFiber() gopurs_runtime.Value {
	once_applicativeFiber.Do(func() {
		applicativeFiber = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.UncurriedApp2(Get__makeFiber(), Get_ffiUtil(), gopurs_runtime.Apply(Get__pure(), a_0)))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyFiber()
}))
	})
	return applicativeFiber
}

var killFiber gopurs_runtime.Value
var once_killFiber sync.Once
func Get_killFiber() gopurs_runtime.Value {
	once_killFiber.Do(func() {
		killFiber = gopurs_runtime.Func2(func(e_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get__bind(), gopurs_runtime.Apply(Get__liftEffect(), gopurs_runtime.RecordGet(v_1, "isSuspended")), gopurs_runtime.Func(func(suspended_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (suspended_2).IntVal != 0 {
__local_var_3_4 := gopurs_runtime.UncurriedApp2(gopurs_runtime.RecordGet(v_1, "kill"), e_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}))
_ = __local_var_3_4
__t3 = gopurs_runtime.Apply(Get__liftEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_4_5 := gopurs_runtime.Apply(__local_var_3_4, gopurs_runtime.Value{})
_ = a_prime_4_5
return pkg_Data_Unit.Get_unit()
}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply(Get_makeAff(), gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.UncurriedApp2(gopurs_runtime.RecordGet(v_1, "kill"), e_0, k_3)
_ = __local_var_4_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
_ = a_prime_5_1
__local_var_6_2 := gopurs_runtime.Apply(Get__liftEffect(), a_prime_5_1)
_ = __local_var_6_2
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_2
})
})
}))
}
end_branch_3:
return __t3
}))
})
	})
	return killFiber
}

var fiberCanceler gopurs_runtime.Value
var once_fiberCanceler sync.Once
func Get_fiberCanceler() gopurs_runtime.Value {
	once_fiberCanceler.Do(func() {
		fiberCanceler = gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_killFiber(), a_1, x_0)
})
	})
	return fiberCanceler
}

var supervise gopurs_runtime.Value
var once_supervise sync.Once
func Get_supervise() gopurs_runtime.Value {
	once_supervise.Do(func() {
		supervise = gopurs_runtime.Func(func(aff_0 gopurs_runtime.Value) gopurs_runtime.Value {
killError_1_0 := gopurs_runtime.Apply(pkg_Effect_Exception.Get_error(), gopurs_runtime.Str("[Aff] Child fiber outlived parent"))
_ = killError_1_0
__local_var_2_1 := gopurs_runtime.UncurriedApp2(Get__makeSupervisedFiber(), Get_ffiUtil(), aff_0)
_ = __local_var_2_1
return gopurs_runtime.Apply3(Get_generalBracket(), gopurs_runtime.Apply(Get__liftEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
sup_3_2 := gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Value{})
_ = sup_3_2
_dollar__unused_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(sup_3_2, "fiber"), "run"), gopurs_runtime.Value{})
_ = _dollar__unused_4_3
return sup_3_2
})), gopurs_runtime.RecordDict3("killed", "failed", "completed", gopurs_runtime.Func2(func(err_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_parSequence_(), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply2(Get_killFiber(), err_2, gopurs_runtime.RecordGet(sup_3, "fiber")), gopurs_runtime.Apply(Get_makeAff(), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), err_2, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Right"), pkg_Data_Unit.Get_unit())))
}))}))
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_makeAff(), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), killError_1_0, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Right"), pkg_Data_Unit.Get_unit())))
}))
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, sup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_makeAff(), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get__killAll(), killError_1_0, gopurs_runtime.RecordGet(sup_3, "supervisor"), gopurs_runtime.Apply(k_4, gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Right"), pkg_Data_Unit.Get_unit())))
}))
})), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_joinFiber(), gopurs_runtime.RecordGet(x_2, "fiber"))
}))
})
	})
	return supervise
}

var monadSTAff gopurs_runtime.Value
var once_monadSTAff sync.Once
func Get_monadSTAff() gopurs_runtime.Value {
	once_monadSTAff.Do(func() {
		monadSTAff = gopurs_runtime.RecordDict2("liftST", "Monad0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get__liftEffect(), x_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}))
	})
	return monadSTAff
}

var monadThrowAff gopurs_runtime.Value
var once_monadThrowAff sync.Once
func Get_monadThrowAff() gopurs_runtime.Value {
	once_monadThrowAff.Do(func() {
		monadThrowAff = gopurs_runtime.RecordDict2("throwError", "Monad0", Get__throwError(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}))
	})
	return monadThrowAff
}

var monadErrorAff gopurs_runtime.Value
var once_monadErrorAff sync.Once
func Get_monadErrorAff() gopurs_runtime.Value {
	once_monadErrorAff.Do(func() {
		monadErrorAff = gopurs_runtime.RecordDict2("catchError", "MonadThrow0", Get__catchError(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowAff()
}))
	})
	return monadErrorAff
}

var try gopurs_runtime.Value
var once_try sync.Once
func Get_try() gopurs_runtime.Value {
	once_try.Do(func() {
		try = gopurs_runtime.Apply(pkg_Control_Monad_Error_Class.Get_try(), Get_monadErrorAff())
	})
	return try
}

var attempt gopurs_runtime.Value
var once_attempt sync.Once
func Get_attempt() gopurs_runtime.Value {
	once_attempt.Do(func() {
		attempt = Get_try()
	})
	return attempt
}

var runAff gopurs_runtime.Value
var once_runAff sync.Once
func Get_runAff() gopurs_runtime.Value {
	once_runAff.Do(func() {
		runAff = gopurs_runtime.Func2(func(k_0 gopurs_runtime.Value, aff_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.UncurriedApp2(Get__makeFiber(), Get_ffiUtil(), gopurs_runtime.Apply2(Get__bind(), gopurs_runtime.Apply(Get_try(), aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get__liftEffect(), gopurs_runtime.Apply(k_0, x_2))
})))
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
fiber_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = fiber_3_1
_dollar__unused_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(fiber_3_1, "run"), gopurs_runtime.Value{})
_ = _dollar__unused_4_2
return fiber_3_1
})
})
	})
	return runAff
}

var runAff_ gopurs_runtime.Value
var once_runAff_ sync.Once
func Get_runAff_() gopurs_runtime.Value {
	once_runAff_.Do(func() {
		runAff_ = gopurs_runtime.Func2(func(k_0 gopurs_runtime.Value, aff_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(Get_runAff(), k_0, aff_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = a_prime_3_1
return pkg_Data_Unit.Get_unit()
})
})
	})
	return runAff_
}

var runSuspendedAff gopurs_runtime.Value
var once_runSuspendedAff sync.Once
func Get_runSuspendedAff() gopurs_runtime.Value {
	once_runSuspendedAff.Do(func() {
		runSuspendedAff = gopurs_runtime.Func2(func(k_0 gopurs_runtime.Value, aff_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get__makeFiber(), Get_ffiUtil(), gopurs_runtime.Apply2(Get__bind(), gopurs_runtime.Apply(Get_try(), aff_1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get__liftEffect(), gopurs_runtime.Apply(k_0, x_2))
})))
})
	})
	return runSuspendedAff
}

var monadRecAff gopurs_runtime.Value
var once_monadRecAff sync.Once
func Get_monadRecAff() gopurs_runtime.Value {
	once_monadRecAff.Do(func() {
		monadRecAff = gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get__bind(), gopurs_runtime.Apply(k_0, a_2), gopurs_runtime.Func(func(res_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(res_3, "_tag").StrVal == "Done")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(Get__pure(), gopurs_runtime.RecordGet(res_3, "value0"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(res_3, "_tag").StrVal == "Loop")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(go__1_0, gopurs_runtime.RecordGet(res_3, "value0"))
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
return go__1_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAff()
}))
	})
	return monadRecAff
}

var monoidAff gopurs_runtime.Value
var once_monoidAff sync.Once
func Get_monoidAff() gopurs_runtime.Value {
	once_monoidAff.Do(func() {
		monoidAff = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(Get__pure(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift21(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")))
}))
})
	})
	return monoidAff
}

var nonCanceler gopurs_runtime.Value
var once_nonCanceler sync.Once
func Get_nonCanceler() gopurs_runtime.Value {
	once_nonCanceler.Do(func() {
		nonCanceler = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(Get__pure(), pkg_Data_Unit.Get_unit())
_ = __local_var_0_0
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_0_0
})
}()
	})
	return nonCanceler
}

var monoidCanceler gopurs_runtime.Value
var once_monoidCanceler sync.Once
func Get_monoidCanceler() gopurs_runtime.Value {
	once_monoidCanceler.Do(func() {
		monoidCanceler = gopurs_runtime.RecordDict2("mempty", "Semigroup0", Get_nonCanceler(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupCanceler()
}))
	})
	return monoidCanceler
}

var never gopurs_runtime.Value
var once_never sync.Once
func Get_never() gopurs_runtime.Value {
	once_never.Do(func() {
		never = gopurs_runtime.Apply(Get_makeAff(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_nonCanceler()
})
}))
	})
	return never
}

var apathize gopurs_runtime.Value
var once_apathize sync.Once
func Get_apathize() gopurs_runtime.Value {
	once_apathize.Do(func() {
		apathize = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(Get__map(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Apply(Get_try(), x_1))
})
}()
	})
	return apathize
}

var altParAff gopurs_runtime.Value
var once_altParAff sync.Once
func Get_altParAff() gopurs_runtime.Value {
	once_altParAff.Do(func() {
		altParAff = gopurs_runtime.RecordDict2("alt", "Functor0", Get__parAffAlt(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorParAff()
}))
	})
	return altParAff
}

var altAff gopurs_runtime.Value
var once_altAff sync.Once
func Get_altAff() gopurs_runtime.Value {
	once_altAff.Do(func() {
		altAff = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(a1_0 gopurs_runtime.Value, a2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get__catchError(), a1_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a2_1
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAff()
}))
	})
	return altAff
}

var plusAff gopurs_runtime.Value
var once_plusAff sync.Once
func Get_plusAff() gopurs_runtime.Value {
	once_plusAff.Do(func() {
		plusAff = gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Apply(Get__throwError(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_error(), gopurs_runtime.Str("Always fails"))), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altAff()
}))
	})
	return plusAff
}

var plusParAff gopurs_runtime.Value
var once_plusParAff sync.Once
func Get_plusParAff() gopurs_runtime.Value {
	once_plusParAff.Do(func() {
		plusParAff = gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.RecordGet(Get_plusAff(), "empty"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altParAff()
}))
	})
	return plusParAff
}

var alternativeParAff gopurs_runtime.Value
var once_alternativeParAff sync.Once
func Get_alternativeParAff() gopurs_runtime.Value {
	once_alternativeParAff.Do(func() {
		alternativeParAff = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeParAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusParAff()
}))
	})
	return alternativeParAff
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

func Get_makeAff() gopurs_runtime.Value {
	return _Gopurs_MakeAff
}
