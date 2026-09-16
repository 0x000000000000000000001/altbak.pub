package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Error_Class_MonadThrow_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Error_Class_MonadThrow_dollar_Dict sync.Once
func Get_Control_Monad_Error_Class_MonadThrow_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_MonadThrow_dollar_Dict.Do(func() {
		cache_Control_Monad_Error_Class_MonadThrow_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Error_Class_MonadThrow_dollar_Dict(func() struct{
	Monad0 gopurs_runtime.Value
	throwError gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Monad0 gopurs_runtime.Value
	throwError gopurs_runtime.Value
}{}
					clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
					clone.throwError = gopurs_runtime.RecordGet(orig, "throwError")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Error_Class_MonadThrow_dollar_Dict
}

var cache_Control_Monad_Error_Class_MonadError_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Error_Class_MonadError_dollar_Dict sync.Once
func Get_Control_Monad_Error_Class_MonadError_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_MonadError_dollar_Dict.Do(func() {
		cache_Control_Monad_Error_Class_MonadError_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Error_Class_MonadError_dollar_Dict(func() struct{
	MonadThrow0 gopurs_runtime.Value
	catchError gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	MonadThrow0 gopurs_runtime.Value
	catchError gopurs_runtime.Value
}{}
					clone.MonadThrow0 = gopurs_runtime.RecordGet(orig, "MonadThrow0")
					clone.catchError = gopurs_runtime.RecordGet(orig, "catchError")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Error_Class_MonadError_dollar_Dict
}

var cache_Control_Monad_Error_Class_throwError gopurs_runtime.Value
var once_Control_Monad_Error_Class_throwError sync.Once
func Get_Control_Monad_Error_Class_throwError() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_throwError.Do(func() {
		cache_Control_Monad_Error_Class_throwError = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_throwError(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_throwError
}

var cache_Control_Monad_Error_Class_monadThrowMaybe gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowMaybe sync.Once
func Get_Control_Monad_Error_Class_monadThrowMaybe() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowMaybe.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowMaybe = gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Error_Class_622782558_3188179519((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Error_Class_1642601656_2568689657(Rebox_Control_Monad_Error_Class_2568689657_1642601656(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Maybe_monadMaybe()))))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}
})})))}
	})
	return cache_Control_Monad_Error_Class_monadThrowMaybe
}

var cache_Control_Monad_Error_Class_monadThrowEither gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowEither sync.Once
func Get_Control_Monad_Error_Class_monadThrowEither() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowEither.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowEither = gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Either_monadEither()))}
}), Get_Data_Either_Left()}))}
	})
	return cache_Control_Monad_Error_Class_monadThrowEither
}

var cache_Control_Monad_Error_Class_monadThrowEffect gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowEffect sync.Once
func Get_Control_Monad_Error_Class_monadThrowEffect() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowEffect.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowEffect = gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}
}), Get_Effect_Exception_throwException()}))}
	})
	return cache_Control_Monad_Error_Class_monadThrowEffect
}

var cache_Control_Monad_Error_Class_monadErrorMaybe gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadErrorMaybe sync.Once
func Get_Control_Monad_Error_Class_monadErrorMaybe() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadErrorMaybe.Do(func() {
		cache_Control_Monad_Error_Class_monadErrorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Error_Class_4076617616_988098289((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Error_Class_622782558_3188179519(Rebox_Control_Monad_Error_Class_3188179519_622782558(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Control_Monad_Error_Class_monadThrowMaybe()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
})})))}
	})
	return cache_Control_Monad_Error_Class_monadErrorMaybe
}

var cache_Control_Monad_Error_Class_monadErrorEither gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadErrorEither sync.Once
func Get_Control_Monad_Error_Class_monadErrorEither() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadErrorEither.Do(func() {
		cache_Control_Monad_Error_Class_monadErrorEither = gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Control_Monad_Error_Class_monadThrowEither()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
	})
	return cache_Control_Monad_Error_Class_monadErrorEither
}

var cache_Control_Monad_Error_Class_monadErrorEffect gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadErrorEffect sync.Once
func Get_Control_Monad_Error_Class_monadErrorEffect() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadErrorEffect.Do(func() {
		cache_Control_Monad_Error_Class_monadErrorEffect = gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Control_Monad_Error_Class_monadThrowEffect()))}
}), gopurs_runtime.Func2(func(b_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Exception_catchException(), a_1, b_0)
})}))}
	})
	return cache_Control_Monad_Error_Class_monadErrorEffect
}

var cache_Control_Monad_Error_Class_liftMaybe gopurs_runtime.Value
var once_Control_Monad_Error_Class_liftMaybe sync.Once
func Get_Control_Monad_Error_Class_liftMaybe() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_liftMaybe.Do(func() {
		cache_Control_Monad_Error_Class_liftMaybe = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_liftMaybe(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadThrow_0_box))
})
	})
	return cache_Control_Monad_Error_Class_liftMaybe
}

var cache_Control_Monad_Error_Class_liftEither gopurs_runtime.Value
var once_Control_Monad_Error_Class_liftEither sync.Once
func Get_Control_Monad_Error_Class_liftEither() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_liftEither.Do(func() {
		cache_Control_Monad_Error_Class_liftEither = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_liftEither(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadThrow_0_box))
})
	})
	return cache_Control_Monad_Error_Class_liftEither
}

var cache_Control_Monad_Error_Class_catchError gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError sync.Once
func Get_Control_Monad_Error_Class_catchError() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError.Do(func() {
		cache_Control_Monad_Error_Class_catchError = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchError
}

var cache_Control_Monad_Error_Class_catchJust gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchJust sync.Once
func Get_Control_Monad_Error_Class_catchJust() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchJust.Do(func() {
		cache_Control_Monad_Error_Class_catchJust = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchJust(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchJust
}

var cache_Control_Monad_Error_Class_try gopurs_runtime.Value
var once_Control_Monad_Error_Class_try sync.Once
func Get_Control_Monad_Error_Class_try() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_try.Do(func() {
		cache_Control_Monad_Error_Class_try = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_try(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0_box))
})
	})
	return cache_Control_Monad_Error_Class_try
}

var cache_Control_Monad_Error_Class_withResource gopurs_runtime.Value
var once_Control_Monad_Error_Class_withResource sync.Once
func Get_Control_Monad_Error_Class_withResource() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_withResource.Do(func() {
		cache_Control_Monad_Error_Class_withResource = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_withResource(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0_box))
})
	})
	return cache_Control_Monad_Error_Class_withResource
}

type Constructor_Control_Monad_Error_Class_MonadThrow[T_e any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[23967309] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "throwError": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Error_Class_MonadThrow: " + key)
		}
	}
}


type Constructor_Control_Monad_Error_Class_MonadError[T_e any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1402181699] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "MonadThrow0": return gopurs_runtime.Box(c.V0)
		case "catchError": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Error_Class_MonadError: " + key)
		}
	}
}


func Call_Control_Monad_Error_Class_MonadThrow_dollar_Dict(x_0_loop struct{
	Monad0 gopurs_runtime.Value
	throwError gopurs_runtime.Value
}) *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	Monad0 gopurs_runtime.Value
	throwError gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Monad0", "throwError", orig.Monad0, orig.throwError)
				}())
}

func Call_Control_Monad_Error_Class_MonadError_dollar_Dict(x_0_loop struct{
	MonadThrow0 gopurs_runtime.Value
	catchError gopurs_runtime.Value
}) *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	MonadThrow0 gopurs_runtime.Value
	catchError gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", orig.MonadThrow0, orig.catchError)
				}())
}

func Call_Control_Monad_Error_Class_throwError(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Control_Monad_Error_Class_liftMaybe(dictMonadThrow_0_loop *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadThrow_0 *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope44)] (TypeApp (TypeVar m$scope42) [(TypeVar a$scope44)]))
pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadThrow_0.V0, gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{})))
_ = pure_1_0
return gopurs_runtime.Func(func(error_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=(TypeApp (TypeVar m) [(TypeVar a)])
__local_var_3_1 := gopurs_runtime.Apply(dictMonadThrow_0.V1, error_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_4)
_ = __t_tag_2
if (__t_tag_2 == nil) {
__t4 = __local_var_3_1
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_4)
_ = __t_tag_3
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.Apply(pure_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_4.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
})
}

func Call_Control_Monad_Error_Class_liftEither(dictMonadThrow_0_loop *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadThrow_0 *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := Call_Control_Monad_Error_Class_throwError(dictMonadThrow_0)
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=(Func [(TypeVar a$scope50)] (TypeApp (TypeVar m$scope48) [(TypeVar a$scope50)]))
__local_var_2_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadThrow_0.V0, gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{})))
_ = __local_var_2_1
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(__local_var_2_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
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
}

func Call_Control_Monad_Error_Class_catchError(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Control_Monad_Error_Class_catchJust(dictMonadError_0_loop *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): MonadThrow0_1_0 shape=App(Other) bindingType=(ADT ["Control","Monad","Error","Class","MonadThrow"] [(TypeVar e$scope58), (TypeVar m$scope59)])
MonadThrow0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonadError_0.V0, gopurs_runtime.Value{}))
_ = MonadThrow0_1_0
return gopurs_runtime.Func3(func(p_2 gopurs_runtime.Value, act_3 gopurs_runtime.Value, handler_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadError_0.V1, act_3, gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope61)])
v_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(p_2, e_5))
_ = v_6_1
var __t2 gopurs_runtime.Value
{
if (v_6_1 == nil) {
__t2 = gopurs_runtime.Apply(MonadThrow0_1_0.V1, e_5)
goto end_branch_2
} else {

}
}
{
if (v_6_1 != nil) {
__t2 = gopurs_runtime.Apply(handler_4, (v_6_1).V0)
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
}

func Call_Control_Monad_Error_Class_try(dictMonadError_0_loop *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadError_0.V0, gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope66)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): pure_3_2 shape=App(Var) bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e$scope65), (TypeVar a$scope67)])] (TypeApp (TypeVar m$scope66) [(ADT ["Data","Either","Either"] [(TypeVar e$scope65), (TypeVar a$scope67)])]))
pure_3_2 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))
_ = pure_3_2
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadError_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Either_Right(), a_4), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), pure_3_2, Get_Data_Either_Left()))
})
}

func Call_Control_Monad_Error_Class_withResource(dictMonadError_0_loop *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): MonadThrow0_1_0 shape=App(Other) bindingType=Any
MonadThrow0_1_0 := gopurs_runtime.Apply(dictMonadError_0.V0, gopurs_runtime.Value{})
_ = MonadThrow0_1_0
// TAST (Let): Monad0_2_1 shape=App(Other) bindingType=Any
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadThrow0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
// TAST (Let): Bind1_3_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope78)])
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): try1_4_3 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope78) [(TypeVar a$scope80)])] (TypeApp (TypeVar m$scope78) [(ADT ["Data","Either","Either"] [Any, (TypeVar a$scope80)])]))
try1_4_3 := Call_Control_Monad_Error_Class_try(dictMonadError_0)
_ = try1_4_3
// TAST (Let): throwError1_5_4 shape=App(Var) bindingType=(Func [Any] (TypeApp (TypeVar m$scope78) [(TypeVar a$scope80)]))
throwError1_5_4 := Call_Control_Monad_Error_Class_throwError(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](MonadThrow0_1_0))
_ = throwError1_5_4
// TAST (Let): pure_6_5 shape=App(Var) bindingType=(Func [(TypeVar a$scope80)] (TypeApp (TypeVar m$scope78) [(TypeVar a$scope80)]))
pure_6_5 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{})))
_ = pure_6_5
return gopurs_runtime.Func3(func(acquire_7 gopurs_runtime.Value, release_8 gopurs_runtime.Value, kleisli_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_2.V1, acquire_7, gopurs_runtime.Func(func(resource_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(try1_4_3, gopurs_runtime.Apply(kleisli_9, resource_10)), gopurs_runtime.Func(func(result_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(Bind1_3_2), gopurs_runtime.Apply(release_8, resource_10), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (result_11.Type == 9 && result_11.IntVal == 3711209382) {
__t6 = gopurs_runtime.Apply(throwError1_5_4, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(result_11.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
if (result_11.Type == 9 && result_11.IntVal == 2465973597) {
__t6 = gopurs_runtime.Apply(pure_6_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(result_11.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
}))
}))
})
}

func Rebox_Control_Monad_Error_Class_1642601656_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Error_Class_2568689657_1642601656(in *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Error_Class_3188179519_622782558(in *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Error_Class_4076617616_988098289(in *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Error_Class_622782558_3188179519(in *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


