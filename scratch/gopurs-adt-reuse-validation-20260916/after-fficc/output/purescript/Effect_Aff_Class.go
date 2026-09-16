package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Aff_Class_lift gopurs_runtime.Value
var once_Effect_Aff_Class_lift sync.Once
func Get_Effect_Aff_Class_lift() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift.Do(func() {
		cache_Effect_Aff_Class_lift = Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Cont_Trans_monadTransContT()))
	})
	return cache_Effect_Aff_Class_lift
}

var cache_Effect_Aff_Class_lift1 gopurs_runtime.Value
var once_Effect_Aff_Class_lift1 sync.Once
func Get_Effect_Aff_Class_lift1() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift1.Do(func() {
		cache_Effect_Aff_Class_lift1 = Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Except_Trans_monadTransExceptT()))
	})
	return cache_Effect_Aff_Class_lift1
}

var cache_Effect_Aff_Class_lift2 gopurs_runtime.Value
var once_Effect_Aff_Class_lift2 sync.Once
func Get_Effect_Aff_Class_lift2() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift2.Do(func() {
		cache_Effect_Aff_Class_lift2 = Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_List_Trans_monadTransListT()))
	})
	return cache_Effect_Aff_Class_lift2
}

var cache_Effect_Aff_Class_lift3 gopurs_runtime.Value
var once_Effect_Aff_Class_lift3 sync.Once
func Get_Effect_Aff_Class_lift3() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift3.Do(func() {
		cache_Effect_Aff_Class_lift3 = Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Maybe_Trans_monadTransMaybeT()))
	})
	return cache_Effect_Aff_Class_lift3
}

var cache_Effect_Aff_Class_lift4 gopurs_runtime.Value
var once_Effect_Aff_Class_lift4 sync.Once
func Get_Effect_Aff_Class_lift4() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift4.Do(func() {
		cache_Effect_Aff_Class_lift4 = Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Reader_Trans_monadTransReaderT()))
	})
	return cache_Effect_Aff_Class_lift4
}

var cache_Effect_Aff_Class_lift5 gopurs_runtime.Value
var once_Effect_Aff_Class_lift5 sync.Once
func Get_Effect_Aff_Class_lift5() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift5.Do(func() {
		cache_Effect_Aff_Class_lift5 = Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_State_Trans_monadTransStateT()))
	})
	return cache_Effect_Aff_Class_lift5
}

var cache_Effect_Aff_Class_MonadAff_dollar_Dict gopurs_runtime.Value
var once_Effect_Aff_Class_MonadAff_dollar_Dict sync.Once
func Get_Effect_Aff_Class_MonadAff_dollar_Dict() gopurs_runtime.Value {
	once_Effect_Aff_Class_MonadAff_dollar_Dict.Do(func() {
		cache_Effect_Aff_Class_MonadAff_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer(Call_Effect_Aff_Class_MonadAff_dollar_Dict(func() struct{
	MonadEffect0 gopurs_runtime.Value
	liftAff gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	MonadEffect0 gopurs_runtime.Value
	liftAff gopurs_runtime.Value
}{}
					clone.MonadEffect0 = gopurs_runtime.RecordGet(orig, "MonadEffect0")
					clone.liftAff = gopurs_runtime.RecordGet(orig, "liftAff")
					return clone
				}()))}
})
	})
	return cache_Effect_Aff_Class_MonadAff_dollar_Dict
}

var cache_Effect_Aff_Class_monadAffAff gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffAff sync.Once
func Get_Effect_Aff_Class_monadAffAff() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffAff.Do(func() {
		cache_Effect_Aff_Class_monadAffAff = gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff()))}
}), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})}))}
	})
	return cache_Effect_Aff_Class_monadAffAff
}

var cache_Effect_Aff_Class_liftAff gopurs_runtime.Value
var once_Effect_Aff_Class_liftAff sync.Once
func Get_Effect_Aff_Class_liftAff() gopurs_runtime.Value {
	once_Effect_Aff_Class_liftAff.Do(func() {
		cache_Effect_Aff_Class_liftAff = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Effect_Aff_Class_liftAff
}

var cache_Effect_Aff_Class_monadAffContT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffContT sync.Once
func Get_Effect_Aff_Class_monadAffContT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffContT.Do(func() {
		cache_Effect_Aff_Class_monadAffContT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffContT(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffContT
}

var cache_Effect_Aff_Class_monadAffExceptT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffExceptT sync.Once
func Get_Effect_Aff_Class_monadAffExceptT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffExceptT.Do(func() {
		cache_Effect_Aff_Class_monadAffExceptT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffExceptT(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffExceptT
}

var cache_Effect_Aff_Class_monadAffListT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffListT sync.Once
func Get_Effect_Aff_Class_monadAffListT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffListT.Do(func() {
		cache_Effect_Aff_Class_monadAffListT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffListT(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffListT
}

var cache_Effect_Aff_Class_monadAffMaybe gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffMaybe sync.Once
func Get_Effect_Aff_Class_monadAffMaybe() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffMaybe.Do(func() {
		cache_Effect_Aff_Class_monadAffMaybe = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffMaybe(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffMaybe
}

var cache_Effect_Aff_Class_monadAffRWS gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffRWS sync.Once
func Get_Effect_Aff_Class_monadAffRWS() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffRWS.Do(func() {
		cache_Effect_Aff_Class_monadAffRWS = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffRWS(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffRWS
}

var cache_Effect_Aff_Class_monadAffReader gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffReader sync.Once
func Get_Effect_Aff_Class_monadAffReader() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffReader.Do(func() {
		cache_Effect_Aff_Class_monadAffReader = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffReader(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffReader
}

var cache_Effect_Aff_Class_monadAffState gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffState sync.Once
func Get_Effect_Aff_Class_monadAffState() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffState.Do(func() {
		cache_Effect_Aff_Class_monadAffState = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffState(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffState
}

var cache_Effect_Aff_Class_monadAffWriter gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffWriter sync.Once
func Get_Effect_Aff_Class_monadAffWriter() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffWriter.Do(func() {
		cache_Effect_Aff_Class_monadAffWriter = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffWriter(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffWriter
}

type Constructor_Effect_Aff_Class_MonadAff[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3183257445] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "MonadEffect0": return gopurs_runtime.Box(c.V0)
		case "liftAff": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Effect_Aff_Class_MonadAff: " + key)
		}
	}
}


func Call_Effect_Aff_Class_MonadAff_dollar_Dict(x_0_loop struct{
	MonadEffect0 gopurs_runtime.Value
	liftAff gopurs_runtime.Value
}) *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] {
var x_0 struct{
	MonadEffect0 gopurs_runtime.Value
	liftAff gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", orig.MonadEffect0, orig.liftAff)
				}())
}

func Call_Effect_Aff_Class_liftAff(dict_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Effect_Aff_Class_monadAffContT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectContT_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope9), (TypeVar m$scope8)])])
monadEffectContT_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_monadEffectContT(MonadEffect0_1_0))
_ = monadEffectContT_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectContT_2_1)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Cont_Trans_monadTransContT())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffExceptT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectExceptT_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (TypeVar m$scope16) [(ADT ["Data","Either","Either"] [(TypeVar e$scope17), (TypeVar a)])])])
monadEffectExceptT_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_monadEffectExceptT(MonadEffect0_1_0))
_ = monadEffectExceptT_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectExceptT_2_1)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Except_Trans_monadTransExceptT())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffListT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectListT_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (TypeVar m$scope23) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar m$scope23), (TypeVar a)])])])])
monadEffectListT_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_monadEffectListT(MonadEffect0_1_0))
_ = monadEffectListT_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectListT_2_1)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_List_Trans_monadTransListT())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffMaybe(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectMaybe_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (TypeVar m$scope29) [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])])
monadEffectMaybe_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_Maybe_Trans_monadEffectMaybe(MonadEffect0_1_0))
_ = monadEffectMaybe_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectMaybe_2_1)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Maybe_Trans_monadTransMaybeT())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffRWS(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): Monad0_2_1 shape=App(Other) bindingType=Any
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
// TAST (Let): liftAff1_3_2 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a$scope44)])] (TypeApp (TypeVar m$scope38) [(TypeVar a$scope44)]))
liftAff1_3_2 := Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0))
_ = liftAff1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadEffectRWS_5_3 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(Func [(TypeVar r$scope40), (TypeVar s$scope41)] (TypeApp (TypeVar m$scope38) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope41), (TypeVar a), (TypeVar w$scope39)])]))])
monadEffectRWS_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_RWS_Trans_monadEffectRWS(dictMonoid_4), MonadEffect0_1_0))
_ = monadEffectRWS_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectRWS_5_3)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_4))), Monad0_2_1), liftAff1_3_2)}))}
})
}

func Call_Effect_Aff_Class_monadAffReader(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectReader_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r$scope49), (TypeVar m$scope48)])])
monadEffectReader_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_Reader_Trans_monadEffectReader(MonadEffect0_1_0))
_ = monadEffectReader_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectReader_2_1)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Reader_Trans_monadTransReaderT())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffState(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectState_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(Func [(TypeVar s$scope57)] (TypeApp (TypeVar m$scope56) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope57)])]))])
monadEffectState_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadEffectState(MonadEffect0_1_0))
_ = monadEffectState_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectState_2_1)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_State_Trans_monadTransStateT())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffWriter(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): Monad0_2_1 shape=App(Other) bindingType=Any
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
// TAST (Let): liftAff1_3_2 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a$scope68)])] (TypeApp (TypeVar m$scope64) [(TypeVar a$scope68)]))
liftAff1_3_2 := Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0))
_ = liftAff1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadEffectWriter_5_3 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (TypeVar m$scope64) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope65)])])])
monadEffectWriter_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_4), MonadEffect0_1_0))
_ = monadEffectWriter_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectWriter_5_3)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_4))), Monad0_2_1), liftAff1_3_2)}))}
})
}


