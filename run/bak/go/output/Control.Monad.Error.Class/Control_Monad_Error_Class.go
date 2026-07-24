package Control_Monad_Error_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
	pkg_Effect "gopurs/output/Effect"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var throwError gopurs_runtime.Value
var once_throwError sync.Once
func Get_throwError() gopurs_runtime.Value {
	once_throwError.Do(func() {
		throwError = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "throwError")
}()
})
	})
	return throwError
}

var monadThrowMaybe gopurs_runtime.Value
var once_monadThrowMaybe sync.Once
func Get_monadThrowMaybe() gopurs_runtime.Value {
	once_monadThrowMaybe.Do(func() {
		monadThrowMaybe = gopurs_runtime.RecordDict2("throwError", "Monad0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_monadMaybe()
}))
	})
	return monadThrowMaybe
}

var monadThrowEither gopurs_runtime.Value
var once_monadThrowEither sync.Once
func Get_monadThrowEither() gopurs_runtime.Value {
	once_monadThrowEither.Do(func() {
		monadThrowEither = gopurs_runtime.RecordDict2("throwError", "Monad0", pkg_Data_Either.Get_Left(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_monadEither()
}))
	})
	return monadThrowEither
}

var monadThrowEffect gopurs_runtime.Value
var once_monadThrowEffect sync.Once
func Get_monadThrowEffect() gopurs_runtime.Value {
	once_monadThrowEffect.Do(func() {
		monadThrowEffect = gopurs_runtime.RecordDict2("throwError", "Monad0", pkg_Effect_Exception.Get_throwException(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}))
	})
	return monadThrowEffect
}

var monadErrorMaybe gopurs_runtime.Value
var once_monadErrorMaybe sync.Once
func Get_monadErrorMaybe() gopurs_runtime.Value {
	once_monadErrorMaybe.Do(func() {
		monadErrorMaybe = gopurs_runtime.RecordDict2("catchError", "MonadThrow0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, pkg_Data_Unit.Get_unit())
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowMaybe()
}))
	})
	return monadErrorMaybe
}

var monadErrorEither gopurs_runtime.Value
var once_monadErrorEither sync.Once
func Get_monadErrorEither() gopurs_runtime.Value {
	once_monadErrorEither.Do(func() {
		monadErrorEither = gopurs_runtime.RecordDict2("catchError", "MonadThrow0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowEither()
}))
	})
	return monadErrorEither
}

var monadErrorEffect gopurs_runtime.Value
var once_monadErrorEffect sync.Once
func Get_monadErrorEffect() gopurs_runtime.Value {
	once_monadErrorEffect.Do(func() {
		monadErrorEffect = gopurs_runtime.RecordDict2("catchError", "MonadThrow0", gopurs_runtime.Func2(func(b_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Effect_Exception.Get_catchException(), a_1, b_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowEffect()
}))
	})
	return monadErrorEffect
}

var liftMaybe gopurs_runtime.Value
var once_liftMaybe sync.Once
func Get_liftMaybe() gopurs_runtime.Value {
	once_liftMaybe.Do(func() {
		liftMaybe = gopurs_runtime.Func(func(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0_loop, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(error_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0_loop, "throwError"), error_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_4.StrVal == "Nothing").IntVal != 0 {
__t2 = __local_var_3_1
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v2_4.StrVal == "Just").IntVal != 0 {
__t2 = gopurs_runtime.Apply(pure_1_0, (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
})
}()
})
	})
	return liftMaybe
}

var liftEither gopurs_runtime.Value
var once_liftEither sync.Once
func Get_liftEither() gopurs_runtime.Value {
	once_liftEither.Do(func() {
		liftEither = gopurs_runtime.Func(func(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
__local_var_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0_loop, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_0
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_2.StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0_loop, "throwError"), (*[1024]gopurs_runtime.Value)(v2_2.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v2_2.StrVal == "Right").IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_1_0, (*[1024]gopurs_runtime.Value)(v2_2.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
}()
})
	})
	return liftEither
}

var catchError gopurs_runtime.Value
var once_catchError sync.Once
func Get_catchError() gopurs_runtime.Value {
	once_catchError.Do(func() {
		catchError = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "catchError")
}()
})
	})
	return catchError
}

var catchJust gopurs_runtime.Value
var once_catchJust sync.Once
func Get_catchJust() gopurs_runtime.Value {
	once_catchJust.Do(func() {
		catchJust = gopurs_runtime.Func4(Call_catchJust)
	})
	return catchJust
}

var try gopurs_runtime.Value
var once_try sync.Once
func Get_try() gopurs_runtime.Value {
	once_try.Do(func() {
		try = gopurs_runtime.Func(func(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0_loop, "MonadThrow0"), gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0_loop, "catchError"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Either.Get_Right(), a_2), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Left", x_3))
}))
})
}()
})
	})
	return try
}

var withResource gopurs_runtime.Value
var once_withResource sync.Once
func Get_withResource() gopurs_runtime.Value {
	once_withResource.Do(func() {
		withResource = gopurs_runtime.Func(func(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
MonadThrow0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0_loop, "MonadThrow0"), gopurs_runtime.Value{})
_ = MonadThrow0_1_0
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadThrow0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
Bind1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_3_2
try1_4_3 := gopurs_runtime.Apply(Get_try(), dictMonadError_0_loop)
_ = try1_4_3
return gopurs_runtime.Func3(func(acquire_5 gopurs_runtime.Value, release_6 gopurs_runtime.Value, kleisli_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), acquire_5, gopurs_runtime.Func(func(resource_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(try1_4_3, gopurs_runtime.Apply(kleisli_7, resource_8)), gopurs_runtime.Func(func(result_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(release_6, resource_8), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(result_9.StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadThrow0_1_0, "throwError"), (*[1024]gopurs_runtime.Value)(result_9.UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(result_9.StrVal == "Right").IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), (*[1024]gopurs_runtime.Value)(result_9.UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
}))
}))
})
}()
})
	})
	return withResource
}

func Call_catchJust(dictMonadError_0_loop gopurs_runtime.Value, p_1_loop gopurs_runtime.Value, act_2_loop gopurs_runtime.Value, handler_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
var act_2 gopurs_runtime.Value = act_2_loop
_ = act_2
var handler_3 gopurs_runtime.Value = handler_3_loop
_ = handler_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0_loop, "catchError"), act_2_loop, gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_0 := gopurs_runtime.Apply(p_1_loop, e_4)
_ = v_5_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0_loop, "MonadThrow0"), gopurs_runtime.Value{}), "throwError"), e_4)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_5_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Apply(handler_3_loop, (*[1024]gopurs_runtime.Value)(v_5_0.UnsafePtr)[0])
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
}


