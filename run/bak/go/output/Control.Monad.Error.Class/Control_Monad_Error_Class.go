package Control_Monad_Error_Class

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_throwError gopurs_runtime.Value
var once_throwError sync.Once
func Get_throwError() gopurs_runtime.Value {
	once_throwError.Do(func() {
		cache_throwError = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_throwError(gopurs_runtime.CoerceToStruct[Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_throwError
}

var cache_throwError__gopurs_runtime_Value_237885032 gopurs_runtime.Value
var once_throwError__gopurs_runtime_Value_237885032 sync.Once
func Get_throwError__gopurs_runtime_Value_237885032() gopurs_runtime.Value {
	once_throwError__gopurs_runtime_Value_237885032.Do(func() {
		cache_throwError__gopurs_runtime_Value_237885032 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_throwError__gopurs_runtime_Value_237885032(gopurs_runtime.CoerceToStruct[Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_throwError__gopurs_runtime_Value_237885032
}

var cache_monadThrowMaybe gopurs_runtime.Value
var once_monadThrowMaybe sync.Once
func Get_monadThrowMaybe() gopurs_runtime.Value {
	once_monadThrowMaybe.Do(func() {
		cache_monadThrowMaybe = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_monadMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}))
	})
	return cache_monadThrowMaybe
}

var cache_monadThrowMaybe__gopurs_runtime_Value_2229618385 gopurs_runtime.Value
var once_monadThrowMaybe__gopurs_runtime_Value_2229618385 sync.Once
func Get_monadThrowMaybe__gopurs_runtime_Value_2229618385() gopurs_runtime.Value {
	once_monadThrowMaybe__gopurs_runtime_Value_2229618385.Do(func() {
		cache_monadThrowMaybe__gopurs_runtime_Value_2229618385 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_monadMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}))
	})
	return cache_monadThrowMaybe__gopurs_runtime_Value_2229618385
}

var cache_monadThrowEither gopurs_runtime.Value
var once_monadThrowEither sync.Once
func Get_monadThrowEither() gopurs_runtime.Value {
	once_monadThrowEither.Do(func() {
		cache_monadThrowEither = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_monadEither()
}), pkg_Data_Either.Get_Left())
	})
	return cache_monadThrowEither
}

var cache_monadThrowEither__gopurs_runtime_Value_103604168 gopurs_runtime.Value
var once_monadThrowEither__gopurs_runtime_Value_103604168 sync.Once
func Get_monadThrowEither__gopurs_runtime_Value_103604168() gopurs_runtime.Value {
	once_monadThrowEither__gopurs_runtime_Value_103604168.Do(func() {
		cache_monadThrowEither__gopurs_runtime_Value_103604168 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_monadEither()
}), pkg_Data_Either.Get_Left())
	})
	return cache_monadThrowEither__gopurs_runtime_Value_103604168
}

var cache_monadThrowEffect gopurs_runtime.Value
var once_monadThrowEffect sync.Once
func Get_monadThrowEffect() gopurs_runtime.Value {
	once_monadThrowEffect.Do(func() {
		cache_monadThrowEffect = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}), pkg_Effect_Exception.Get_throwException())
	})
	return cache_monadThrowEffect
}

var cache_monadThrowEffect__gopurs_runtime_Value_18811790 gopurs_runtime.Value
var once_monadThrowEffect__gopurs_runtime_Value_18811790 sync.Once
func Get_monadThrowEffect__gopurs_runtime_Value_18811790() gopurs_runtime.Value {
	once_monadThrowEffect__gopurs_runtime_Value_18811790.Do(func() {
		cache_monadThrowEffect__gopurs_runtime_Value_18811790 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}), pkg_Effect_Exception.Get_throwException())
	})
	return cache_monadThrowEffect__gopurs_runtime_Value_18811790
}

var cache_monadErrorMaybe gopurs_runtime.Value
var once_monadErrorMaybe sync.Once
func Get_monadErrorMaybe() gopurs_runtime.Value {
	once_monadErrorMaybe.Do(func() {
		cache_monadErrorMaybe = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, pkg_Data_Unit.Get_unit())))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0})}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_monadErrorMaybe
}

var cache_monadErrorEither gopurs_runtime.Value
var once_monadErrorEither sync.Once
func Get_monadErrorEither() gopurs_runtime.Value {
	once_monadErrorEither.Do(func() {
		cache_monadErrorEither = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
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
	return cache_monadErrorEither
}

var cache_monadErrorEffect gopurs_runtime.Value
var once_monadErrorEffect sync.Once
func Get_monadErrorEffect() gopurs_runtime.Value {
	once_monadErrorEffect.Do(func() {
		cache_monadErrorEffect = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadThrowEffect()
}), gopurs_runtime.Func(func(b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Effect_Exception.Get_catchException(), a_1, b_0)
})
}))
	})
	return cache_monadErrorEffect
}

var cache_liftMaybe gopurs_runtime.Value
var once_liftMaybe sync.Once
func Get_liftMaybe() gopurs_runtime.Value {
	once_liftMaybe.Do(func() {
		cache_liftMaybe = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftMaybe(gopurs_runtime.CoerceToStruct[Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadThrow_0_box))
})
	})
	return cache_liftMaybe
}

var cache_liftEither gopurs_runtime.Value
var once_liftEither sync.Once
func Get_liftEither() gopurs_runtime.Value {
	once_liftEither.Do(func() {
		cache_liftEither = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEither(gopurs_runtime.CoerceToStruct[Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadThrow_0_box))
})
	})
	return cache_liftEither
}

var cache_catchError gopurs_runtime.Value
var once_catchError sync.Once
func Get_catchError() gopurs_runtime.Value {
	once_catchError.Do(func() {
		cache_catchError = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catchError(gopurs_runtime.CoerceToStruct[Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_catchError
}

var cache_catchError__gopurs_runtime_Value_2657403463 gopurs_runtime.Value
var once_catchError__gopurs_runtime_Value_2657403463 sync.Once
func Get_catchError__gopurs_runtime_Value_2657403463() gopurs_runtime.Value {
	once_catchError__gopurs_runtime_Value_2657403463.Do(func() {
		cache_catchError__gopurs_runtime_Value_2657403463 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catchError__gopurs_runtime_Value_2657403463(gopurs_runtime.CoerceToStruct[Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_catchError__gopurs_runtime_Value_2657403463
}

var cache_catchJust gopurs_runtime.Value
var once_catchJust sync.Once
func Get_catchJust() gopurs_runtime.Value {
	once_catchJust.Do(func() {
		cache_catchJust = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catchJust(gopurs_runtime.CoerceToStruct[Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0_box))
})
	})
	return cache_catchJust
}

var cache_try gopurs_runtime.Value
var once_try sync.Once
func Get_try() gopurs_runtime.Value {
	once_try.Do(func() {
		cache_try = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_try(gopurs_runtime.CoerceToStruct[Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0_box))
})
	})
	return cache_try
}

var cache_try__gopurs_runtime_Value_2648905537 gopurs_runtime.Value
var once_try__gopurs_runtime_Value_2648905537 sync.Once
func Get_try__gopurs_runtime_Value_2648905537() gopurs_runtime.Value {
	once_try__gopurs_runtime_Value_2648905537.Do(func() {
		cache_try__gopurs_runtime_Value_2648905537 = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_try__gopurs_runtime_Value_2648905537(gopurs_runtime.CoerceToStruct[Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0_box))
})
	})
	return cache_try__gopurs_runtime_Value_2648905537
}

var cache_withResource gopurs_runtime.Value
var once_withResource sync.Once
func Get_withResource() gopurs_runtime.Value {
	once_withResource.Do(func() {
		cache_withResource = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withResource(gopurs_runtime.CoerceToStruct[Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0_box))
})
	})
	return cache_withResource
}

type Constructor_MonadThrow[T_e any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[23967309] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "throwError": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadThrow: " + key)
		}
	}
}


type Constructor_MonadError[T_e any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1402181699] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "MonadThrow0": return c.V0
		case "catchError": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadError: " + key)
		}
	}
}


func Call_throwError(dict_0_loop *Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_throwError__gopurs_runtime_Value_237885032(dict_0_loop *Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftMaybe(dictMonadThrow_0_loop *Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadThrow_0 *Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadThrow_0_loop
_ = dictMonadThrow_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadThrow_0.V0, gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(error_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonadThrow_0.V1, error_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 930809136 && v2_4.UnsafePtr == nil) {
__t2 = __local_var_3_1
goto end_branch_2
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 930809136 && v2_4.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply(pure_1_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_4.UnsafePtr).V0)
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
}

func Call_liftEither(dictMonadThrow_0_loop *Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadThrow_0 *Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadThrow_0_loop
_ = dictMonadThrow_0
__local_var_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadThrow_0.V0, gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_0
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(dictMonadThrow_0.V1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply(__local_var_1_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
}

func Call_catchError(dict_0_loop *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_catchError__gopurs_runtime_Value_2657403463(dict_0_loop *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_catchJust(dictMonadError_0_loop *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadError_0_loop
_ = dictMonadError_0
MonadThrow0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonadError_0.V0, gopurs_runtime.Value{}))
_ = MonadThrow0_1_0
return gopurs_runtime.Func(func(p_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(act_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(handler_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadError_0.V1, act_3, gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(p_2, e_5))
_ = v_6_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply(MonadThrow0_1_0.V1, e_5)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply(handler_4, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_6_1)}.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
})
})
}

func Call_try(dictMonadError_0_loop *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadError_0_loop
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

func Call_try__gopurs_runtime_Value_2648905537(dictMonadError_0_loop *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadError_0_loop
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

func Call_withResource(dictMonadError_0_loop *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadError_0_loop
_ = dictMonadError_0
MonadThrow0_1_0 := gopurs_runtime.Apply(dictMonadError_0.V0, gopurs_runtime.Value{})
_ = MonadThrow0_1_0
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadThrow0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
try1_4_3 := Call_try(dictMonadError_0)
_ = try1_4_3
pure_5_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_4
return gopurs_runtime.Func(func(acquire_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(release_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kleisli_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_2.V1, acquire_6, gopurs_runtime.Func(func(resource_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(try1_4_3, gopurs_runtime.Apply(kleisli_8, resource_9)), gopurs_runtime.Func(func(result_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Bind1_3_2)}, gopurs_runtime.Apply(release_7, resource_9), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (result_10.Type == 9 && result_10.IntVal == 3711209382) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadThrow0_1_0, "throwError"), (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(result_10.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
if (result_10.Type == 9 && result_10.IntVal == 2465973597) {
__t5 = gopurs_runtime.Apply(pure_5_4, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(result_10.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
}))
}))
})
})
})
}


