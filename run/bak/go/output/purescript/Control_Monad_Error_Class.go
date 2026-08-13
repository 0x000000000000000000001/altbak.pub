package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Error_Class_MonadThrow_dollarDict gopurs_runtime.Value
var once_Control_Monad_Error_Class_MonadThrow_dollarDict sync.Once
func Get_Control_Monad_Error_Class_MonadThrow_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_MonadThrow_dollarDict.Do(func() {
		cache_Control_Monad_Error_Class_MonadThrow_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_MonadThrow_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Error_Class_MonadThrow_dollarDict
}

var cache_Control_Monad_Error_Class_MonadError_dollarDict gopurs_runtime.Value
var once_Control_Monad_Error_Class_MonadError_dollarDict sync.Once
func Get_Control_Monad_Error_Class_MonadError_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_MonadError_dollarDict.Do(func() {
		cache_Control_Monad_Error_Class_MonadError_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_MonadError_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Error_Class_MonadError_dollarDict
}

var cache_Control_Monad_Error_Class_throwError gopurs_runtime.Value
var once_Control_Monad_Error_Class_throwError sync.Once
func Get_Control_Monad_Error_Class_throwError() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_throwError.Do(func() {
		cache_Control_Monad_Error_Class_throwError = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_throwError(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_throwError
}

var cache_Control_Monad_Error_Class_monadThrowMaybe gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowMaybe sync.Once
func Get_Control_Monad_Error_Class_monadThrowMaybe() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowMaybe.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowMaybe = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_monadMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Control_Monad_Error_Class_monadThrowMaybe
}

var cache_Control_Monad_Error_Class_monadThrowEither gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowEither sync.Once
func Get_Control_Monad_Error_Class_monadThrowEither() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowEither.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowEither = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_monadEither()
}), Get_Data_Either_Left())
	})
	return cache_Control_Monad_Error_Class_monadThrowEither
}

var cache_Control_Monad_Error_Class_monadThrowEffect gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowEffect sync.Once
func Get_Control_Monad_Error_Class_monadThrowEffect() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowEffect.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowEffect = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_monadEffect()
}), Get_Effect_Exception_throwException())
	})
	return cache_Control_Monad_Error_Class_monadThrowEffect
}

var cache_Control_Monad_Error_Class_monadErrorMaybe gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadErrorMaybe sync.Once
func Get_Control_Monad_Error_Class_monadErrorMaybe() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadErrorMaybe.Do(func() {
		cache_Control_Monad_Error_Class_monadErrorMaybe = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_monadMaybe()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Maybe_Just)(v_0.UnsafePtr).V0}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
}))
	})
	return cache_Control_Monad_Error_Class_monadErrorMaybe
}

var cache_Control_Monad_Error_Class_monadErrorEither gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadErrorEither sync.Once
func Get_Control_Monad_Error_Class_monadErrorEither() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadErrorEither.Do(func() {
		cache_Control_Monad_Error_Class_monadErrorEither = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_monadEither()
}), Get_Data_Either_Left())
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Either_Right
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Either_Right](gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Left)(v_0.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = &Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(v_0.UnsafePtr).V0}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Either_Right](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(__t0)}
})
}))
	})
	return cache_Control_Monad_Error_Class_monadErrorEither
}

var cache_Control_Monad_Error_Class_monadErrorEffect gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadErrorEffect sync.Once
func Get_Control_Monad_Error_Class_monadErrorEffect() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadErrorEffect.Do(func() {
		cache_Control_Monad_Error_Class_monadErrorEffect = gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_monadEffect()
}), Get_Effect_Exception_throwException())
}), gopurs_runtime.Func(func(b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Exception_catchException(), a_1, b_0)
})
}))
	})
	return cache_Control_Monad_Error_Class_monadErrorEffect
}

var cache_Control_Monad_Error_Class_liftMaybe gopurs_runtime.Value
var once_Control_Monad_Error_Class_liftMaybe sync.Once
func Get_Control_Monad_Error_Class_liftMaybe() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_liftMaybe.Do(func() {
		cache_Control_Monad_Error_Class_liftMaybe = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_liftMaybe(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow](dictMonadThrow_0_box))
})
	})
	return cache_Control_Monad_Error_Class_liftMaybe
}

var cache_Control_Monad_Error_Class_liftEither gopurs_runtime.Value
var once_Control_Monad_Error_Class_liftEither sync.Once
func Get_Control_Monad_Error_Class_liftEither() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_liftEither.Do(func() {
		cache_Control_Monad_Error_Class_liftEither = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_liftEither(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow](dictMonadThrow_0_box))
})
	})
	return cache_Control_Monad_Error_Class_liftEither
}

var cache_Control_Monad_Error_Class_catchError gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError sync.Once
func Get_Control_Monad_Error_Class_catchError() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError.Do(func() {
		cache_Control_Monad_Error_Class_catchError = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchError
}

var cache_Control_Monad_Error_Class_catchJust gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchJust sync.Once
func Get_Control_Monad_Error_Class_catchJust() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchJust.Do(func() {
		cache_Control_Monad_Error_Class_catchJust = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchJust(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dictMonadError_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchJust
}

var cache_Control_Monad_Error_Class_try gopurs_runtime.Value
var once_Control_Monad_Error_Class_try sync.Once
func Get_Control_Monad_Error_Class_try() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_try.Do(func() {
		cache_Control_Monad_Error_Class_try = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_try(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dictMonadError_0_box))
})
	})
	return cache_Control_Monad_Error_Class_try
}

var cache_Control_Monad_Error_Class_withResource gopurs_runtime.Value
var once_Control_Monad_Error_Class_withResource sync.Once
func Get_Control_Monad_Error_Class_withResource() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_withResource.Do(func() {
		cache_Control_Monad_Error_Class_withResource = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_withResource(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dictMonadError_0_box))
})
	})
	return cache_Control_Monad_Error_Class_withResource
}

var cache_Control_Monad_Error_Class_catchError__3620969455 gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError__3620969455 sync.Once
func Get_Control_Monad_Error_Class_catchError__3620969455() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError__3620969455.Do(func() {
		cache_Control_Monad_Error_Class_catchError__3620969455 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError__3620969455(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchError__3620969455
}

var cache_Control_Monad_Error_Class_catchError__2657403463 gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError__2657403463 sync.Once
func Get_Control_Monad_Error_Class_catchError__2657403463() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError__2657403463.Do(func() {
		cache_Control_Monad_Error_Class_catchError__2657403463 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError__2657403463(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchError__2657403463
}

var cache_Control_Monad_Error_Class_catchError__1102377099 gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError__1102377099 sync.Once
func Get_Control_Monad_Error_Class_catchError__1102377099() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError__1102377099.Do(func() {
		cache_Control_Monad_Error_Class_catchError__1102377099 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError__1102377099(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchError__1102377099
}

var cache_Control_Monad_Error_Class_catchError__1612922415 gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError__1612922415 sync.Once
func Get_Control_Monad_Error_Class_catchError__1612922415() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError__1612922415.Do(func() {
		cache_Control_Monad_Error_Class_catchError__1612922415 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError__1612922415(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchError__1612922415
}

var cache_Control_Monad_Error_Class_catchError__4177389606 gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError__4177389606 sync.Once
func Get_Control_Monad_Error_Class_catchError__4177389606() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError__4177389606.Do(func() {
		cache_Control_Monad_Error_Class_catchError__4177389606 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError__4177389606(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchError__4177389606
}

var cache_Control_Monad_Error_Class_catchError__3649261295 gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError__3649261295 sync.Once
func Get_Control_Monad_Error_Class_catchError__3649261295() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError__3649261295.Do(func() {
		cache_Control_Monad_Error_Class_catchError__3649261295 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError__3649261295(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_catchError__3649261295
}

var cache_Control_Monad_Error_Class_catchError__3892322529 gopurs_runtime.Value
var once_Control_Monad_Error_Class_catchError__3892322529 sync.Once
func Get_Control_Monad_Error_Class_catchError__3892322529() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_catchError__3892322529.Do(func() {
		cache_Control_Monad_Error_Class_catchError__3892322529 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_catchError__3892322529(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Monad_Error_Class_catchError__3892322529
}

var cache_Control_Monad_Error_Class_monadThrowEffect__18811790 gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowEffect__18811790 sync.Once
func Get_Control_Monad_Error_Class_monadThrowEffect__18811790() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowEffect__18811790.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowEffect__18811790 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_monadEffect()
}), Get_Effect_Exception_throwException())
	})
	return cache_Control_Monad_Error_Class_monadThrowEffect__18811790
}

var cache_Control_Monad_Error_Class_monadThrowEither__103604168 gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowEither__103604168 sync.Once
func Get_Control_Monad_Error_Class_monadThrowEither__103604168() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowEither__103604168.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowEither__103604168 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_monadEither()
}), Get_Data_Either_Left())
	})
	return cache_Control_Monad_Error_Class_monadThrowEither__103604168
}

var cache_Control_Monad_Error_Class_monadThrowMaybe__2229618385 gopurs_runtime.Value
var once_Control_Monad_Error_Class_monadThrowMaybe__2229618385 sync.Once
func Get_Control_Monad_Error_Class_monadThrowMaybe__2229618385() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_monadThrowMaybe__2229618385.Do(func() {
		cache_Control_Monad_Error_Class_monadThrowMaybe__2229618385 = gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_monadMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Control_Monad_Error_Class_monadThrowMaybe__2229618385
}

var cache_Control_Monad_Error_Class_throwError__237885032 gopurs_runtime.Value
var once_Control_Monad_Error_Class_throwError__237885032 sync.Once
func Get_Control_Monad_Error_Class_throwError__237885032() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_throwError__237885032.Do(func() {
		cache_Control_Monad_Error_Class_throwError__237885032 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_throwError__237885032(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_throwError__237885032
}

var cache_Control_Monad_Error_Class_throwError__1338676736 gopurs_runtime.Value
var once_Control_Monad_Error_Class_throwError__1338676736 sync.Once
func Get_Control_Monad_Error_Class_throwError__1338676736() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_throwError__1338676736.Do(func() {
		cache_Control_Monad_Error_Class_throwError__1338676736 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_throwError__1338676736(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow](dict_0_box))
})
	})
	return cache_Control_Monad_Error_Class_throwError__1338676736
}

var cache_Control_Monad_Error_Class_throwError__1668092494 gopurs_runtime.Value
var once_Control_Monad_Error_Class_throwError__1668092494 sync.Once
func Get_Control_Monad_Error_Class_throwError__1668092494() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_throwError__1668092494.Do(func() {
		cache_Control_Monad_Error_Class_throwError__1668092494 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_throwError__1668092494(__eta0_0_box)
})
	})
	return cache_Control_Monad_Error_Class_throwError__1668092494
}

var cache_Control_Monad_Error_Class_try__2648905537 gopurs_runtime.Value
var once_Control_Monad_Error_Class_try__2648905537 sync.Once
func Get_Control_Monad_Error_Class_try__2648905537() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_try__2648905537.Do(func() {
		cache_Control_Monad_Error_Class_try__2648905537 = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_try__2648905537(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError](dictMonadError_0_box))
})
	})
	return cache_Control_Monad_Error_Class_try__2648905537
}

var cache_Control_Monad_Error_Class_try__214520782 gopurs_runtime.Value
var once_Control_Monad_Error_Class_try__214520782 sync.Once
func Get_Control_Monad_Error_Class_try__214520782() gopurs_runtime.Value {
	once_Control_Monad_Error_Class_try__214520782.Do(func() {
		cache_Control_Monad_Error_Class_try__214520782 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Error_Class_try__214520782(__eta0_0_box)
})
	})
	return cache_Control_Monad_Error_Class_try__214520782
}

type Constructor_Control_Monad_Error_Class_MonadThrow struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[23967309] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Error_Class_MonadThrow)(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "throwError": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Error_Class_MonadThrow: " + key)
		}
	}
}


type Constructor_Control_Monad_Error_Class_MonadError struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1402181699] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Error_Class_MonadError)(ptr)
		_ = c
		switch key {
		case "MonadThrow0": return gopurs_runtime.Box(c.V0)
		case "catchError": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Error_Class_MonadError: " + key)
		}
	}
}


func Call_Control_Monad_Error_Class_MonadThrow_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Error_Class_MonadError_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Error_Class_throwError(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadThrow) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadThrow = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_liftMaybe(dictMonadThrow_0_loop *Constructor_Control_Monad_Error_Class_MonadThrow) gopurs_runtime.Value {
var dictMonadThrow_0 *Constructor_Control_Monad_Error_Class_MonadThrow = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadThrow_0.V0), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(error_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadThrow_0.V1), error_2)
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
__t2 = gopurs_runtime.Apply(pure_1_0, (*Constructor_Data_Maybe_Just)(v2_4.UnsafePtr).V0)
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

func Call_Control_Monad_Error_Class_liftEither(dictMonadThrow_0_loop *Constructor_Control_Monad_Error_Class_MonadThrow) gopurs_runtime.Value {
var dictMonadThrow_0 *Constructor_Control_Monad_Error_Class_MonadThrow = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadThrow_0.V0), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_0
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadThrow_0.V1), (*Constructor_Data_Either_Left)(v2_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0)
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

func Call_Control_Monad_Error_Class_catchError(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadError = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_catchJust(dictMonadError_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_Control_Monad_Error_Class_MonadError = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): MonadThrow0_1_0 -> *Constructor_Control_Monad_Error_Class_MonadThrow
MonadThrow0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadError_0.V0), gopurs_runtime.Value{}))
_ = MonadThrow0_1_0
return gopurs_runtime.Func(func(p_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(act_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(handler_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadError_0.V1), act_3, gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_1 -> *Constructor_Data_Maybe_Just
v_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(p_2, e_5))
_ = v_6_1
var __t2 gopurs_runtime.Value
{
if (v_6_1 == nil) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(MonadThrow0_1_0.V1), e_5)
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
})
})
}

func Call_Control_Monad_Error_Class_try(dictMonadError_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_Control_Monad_Error_Class_MonadError = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadError_0.V0), gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadError_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Either_Right(), a_4), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_5})})
}))
})
}

func Call_Control_Monad_Error_Class_withResource(dictMonadError_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_Control_Monad_Error_Class_MonadError = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): MonadThrow0_1_0 -> gopurs_runtime.Value
MonadThrow0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadError_0.V0), gopurs_runtime.Value{})
_ = MonadThrow0_1_0
// TAST (Let): Monad0_2_1 -> gopurs_runtime.Value
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadThrow0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Monad0_4_4 -> gopurs_runtime.Value
Monad0_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadError_0.V0), gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_4
// TAST (Let): Functor0_5_5 -> *Constructor_Data_Functor_Functor
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
// TAST (Let): pure_6_6 -> gopurs_runtime.Value
pure_6_6 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_6
// TAST (Let): try1_4_3 -> gopurs_runtime.Value
try1_4_3 := gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadError_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), Get_Data_Either_Right(), a_7), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_6, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_8})})
}))
})
_ = try1_4_3
// TAST (Let): pure_5_7 -> gopurs_runtime.Value
pure_5_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_7
return gopurs_runtime.Func(func(acquire_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(release_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kleisli_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), acquire_6, gopurs_runtime.Func(func(resource_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(try1_4_3, gopurs_runtime.Apply(kleisli_8, resource_9)), gopurs_runtime.Func(func(result_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Bind1_3_2)}, gopurs_runtime.Apply(release_7, resource_9), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (result_10.Type == 9 && result_10.IntVal == 3711209382) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadThrow0_1_0, "throwError"), (*Constructor_Data_Either_Left)(result_10.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
if (result_10.Type == 9 && result_10.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(pure_5_7, (*Constructor_Data_Either_Right)(result_10.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
}))
}))
})
})
})
}

func Call_Control_Monad_Error_Class_catchError__3620969455(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadError = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_catchError__2657403463(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadError = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_catchError__1102377099(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadError = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_catchError__1612922415(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadError = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_catchError__4177389606(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadError = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_catchError__3649261295(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadError = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_catchError__3892322529(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__catchError(), __eta0_0, __eta1_1)
}

func Call_Control_Monad_Error_Class_throwError__237885032(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadThrow) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadThrow = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_throwError__1338676736(dict_0_loop *Constructor_Control_Monad_Error_Class_MonadThrow) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Error_Class_MonadThrow = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Error_Class_throwError__1668092494(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__throwError(), __eta0_0)
}

func Call_Control_Monad_Error_Class_try__2648905537(dictMonadError_0_loop *Constructor_Control_Monad_Error_Class_MonadError) gopurs_runtime.Value {
var dictMonadError_0 *Constructor_Control_Monad_Error_Class_MonadError = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadError_0.V0), gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadError_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Either_Right(), a_4), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_5})})
}))
})
}

func Call_Control_Monad_Error_Class_try__214520782(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_Aff_monadThrowAff(), "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_monadErrorAff(), "catchError"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Either_Right(), __eta0_0), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_3})})
}))
}


