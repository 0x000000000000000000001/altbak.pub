package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_List_Trans_identity gopurs_runtime.Value
var once_Control_Monad_List_Trans_identity sync.Once
func Get_Control_Monad_List_Trans_identity() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_identity.Do(func() {
		cache_Control_Monad_List_Trans_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Control_Monad_List_Trans_identity
}

var cache_Control_Monad_List_Trans_identity1 gopurs_runtime.Value
var once_Control_Monad_List_Trans_identity1 sync.Once
func Get_Control_Monad_List_Trans_identity1() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_identity1.Do(func() {
		cache_Control_Monad_List_Trans_identity1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Control_Monad_List_Trans_identity1
}

var cache_Control_Monad_List_Trans_Yield gopurs_runtime.Value
var once_Control_Monad_List_Trans_Yield sync.Once
func Get_Control_Monad_List_Trans_Yield() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_Yield.Do(func() {
		cache_Control_Monad_List_Trans_Yield = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
})
})
	})
	return cache_Control_Monad_List_Trans_Yield
}

var cache_Control_Monad_List_Trans_Skip gopurs_runtime.Value
var once_Control_Monad_List_Trans_Skip sync.Once
func Get_Control_Monad_List_Trans_Skip() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_Skip.Do(func() {
		cache_Control_Monad_List_Trans_Skip = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Control_Monad_List_Trans_Skip
}

var cache_Control_Monad_List_Trans_Done gopurs_runtime.Value
var once_Control_Monad_List_Trans_Done sync.Once
func Get_Control_Monad_List_Trans_Done() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_Done.Do(func() {
		cache_Control_Monad_List_Trans_Done = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Control_Monad_List_Trans_Done
}

var cache_Control_Monad_List_Trans_ListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_ListT sync.Once
func Get_Control_Monad_List_Trans_ListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_ListT.Do(func() {
		cache_Control_Monad_List_Trans_ListT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_ListT(x_0_box)
})
	})
	return cache_Control_Monad_List_Trans_ListT
}

var cache_Control_Monad_List_Trans_wrapLazy gopurs_runtime.Value
var once_Control_Monad_List_Trans_wrapLazy sync.Once
func Get_Control_Monad_List_Trans_wrapLazy() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_wrapLazy.Do(func() {
		cache_Control_Monad_List_Trans_wrapLazy = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_wrapLazy(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), v_1_box)
})
	})
	return cache_Control_Monad_List_Trans_wrapLazy
}

var cache_Control_Monad_List_Trans_wrapEffect gopurs_runtime.Value
var once_Control_Monad_List_Trans_wrapEffect sync.Once
func Get_Control_Monad_List_Trans_wrapEffect() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_wrapEffect.Do(func() {
		cache_Control_Monad_List_Trans_wrapEffect = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_wrapEffect(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
})
	})
	return cache_Control_Monad_List_Trans_wrapEffect
}

var cache_Control_Monad_List_Trans_unfold gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfold sync.Once
func Get_Control_Monad_List_Trans_unfold() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfold.Do(func() {
		cache_Control_Monad_List_Trans_unfold = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_unfold(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_unfold
}

var cache_Control_Monad_List_Trans_uncons gopurs_runtime.Value
var once_Control_Monad_List_Trans_uncons sync.Once
func Get_Control_Monad_List_Trans_uncons() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_uncons.Do(func() {
		cache_Control_Monad_List_Trans_uncons = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_uncons(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_uncons
}

var cache_Control_Monad_List_Trans_tail gopurs_runtime.Value
var once_Control_Monad_List_Trans_tail sync.Once
func Get_Control_Monad_List_Trans_tail() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_tail.Do(func() {
		cache_Control_Monad_List_Trans_tail = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_tail(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_tail
}

var cache_Control_Monad_List_Trans_stepMap gopurs_runtime.Value
var once_Control_Monad_List_Trans_stepMap sync.Once
func Get_Control_Monad_List_Trans_stepMap() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_stepMap.Do(func() {
		cache_Control_Monad_List_Trans_stepMap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_stepMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_stepMap
}

var cache_Control_Monad_List_Trans_takeWhile gopurs_runtime.Value
var once_Control_Monad_List_Trans_takeWhile sync.Once
func Get_Control_Monad_List_Trans_takeWhile() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_takeWhile.Do(func() {
		cache_Control_Monad_List_Trans_takeWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_takeWhile(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_takeWhile
}

var cache_Control_Monad_List_Trans_scanl gopurs_runtime.Value
var once_Control_Monad_List_Trans_scanl sync.Once
func Get_Control_Monad_List_Trans_scanl() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_scanl.Do(func() {
		cache_Control_Monad_List_Trans_scanl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_scanl(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_scanl
}

var cache_Control_Monad_List_Trans_prepend_prime_ gopurs_runtime.Value
var once_Control_Monad_List_Trans_prepend_prime_ sync.Once
func Get_Control_Monad_List_Trans_prepend_prime_() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_prepend_prime_.Do(func() {
		cache_Control_Monad_List_Trans_prepend_prime_ = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_prepend_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_Control_Monad_List_Trans_prepend_prime_
}

var cache_Control_Monad_List_Trans_prepend gopurs_runtime.Value
var once_Control_Monad_List_Trans_prepend sync.Once
func Get_Control_Monad_List_Trans_prepend() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_prepend.Do(func() {
		cache_Control_Monad_List_Trans_prepend = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_prepend(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_Control_Monad_List_Trans_prepend
}

var cache_Control_Monad_List_Trans_nil gopurs_runtime.Value
var once_Control_Monad_List_Trans_nil sync.Once
func Get_Control_Monad_List_Trans_nil() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_nil.Do(func() {
		cache_Control_Monad_List_Trans_nil = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_nil(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_List_Trans_nil
}

var cache_Control_Monad_List_Trans_singleton gopurs_runtime.Value
var once_Control_Monad_List_Trans_singleton sync.Once
func Get_Control_Monad_List_Trans_singleton() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_singleton.Do(func() {
		cache_Control_Monad_List_Trans_singleton = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_singleton
}

var cache_Control_Monad_List_Trans_take gopurs_runtime.Value
var once_Control_Monad_List_Trans_take sync.Once
func Get_Control_Monad_List_Trans_take() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_take.Do(func() {
		cache_Control_Monad_List_Trans_take = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_take(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_take
}

var cache_Control_Monad_List_Trans_zipWith_prime_ gopurs_runtime.Value
var once_Control_Monad_List_Trans_zipWith_prime_ sync.Once
func Get_Control_Monad_List_Trans_zipWith_prime_() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_zipWith_prime_.Do(func() {
		cache_Control_Monad_List_Trans_zipWith_prime_ = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_zipWith_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_zipWith_prime_
}

var cache_Control_Monad_List_Trans_zipWith gopurs_runtime.Value
var once_Control_Monad_List_Trans_zipWith sync.Once
func Get_Control_Monad_List_Trans_zipWith() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_zipWith.Do(func() {
		cache_Control_Monad_List_Trans_zipWith = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_zipWith(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_zipWith
}

var cache_Control_Monad_List_Trans_newtypeListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_newtypeListT sync.Once
func Get_Control_Monad_List_Trans_newtypeListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_newtypeListT.Do(func() {
		cache_Control_Monad_List_Trans_newtypeListT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Control_Monad_List_Trans_newtypeListT
}

var cache_Control_Monad_List_Trans_mapMaybe gopurs_runtime.Value
var once_Control_Monad_List_Trans_mapMaybe sync.Once
func Get_Control_Monad_List_Trans_mapMaybe() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_mapMaybe.Do(func() {
		cache_Control_Monad_List_Trans_mapMaybe = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_mapMaybe
}

var cache_Control_Monad_List_Trans_iterate gopurs_runtime.Value
var once_Control_Monad_List_Trans_iterate sync.Once
func Get_Control_Monad_List_Trans_iterate() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_iterate.Do(func() {
		cache_Control_Monad_List_Trans_iterate = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_iterate(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_iterate
}

var cache_Control_Monad_List_Trans_repeat gopurs_runtime.Value
var once_Control_Monad_List_Trans_repeat sync.Once
func Get_Control_Monad_List_Trans_repeat() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_repeat.Do(func() {
		cache_Control_Monad_List_Trans_repeat = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_repeat(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_repeat
}

var cache_Control_Monad_List_Trans_head gopurs_runtime.Value
var once_Control_Monad_List_Trans_head sync.Once
func Get_Control_Monad_List_Trans_head() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_head.Do(func() {
		cache_Control_Monad_List_Trans_head = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_head(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_head
}

var cache_Control_Monad_List_Trans_functorListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_functorListT sync.Once
func Get_Control_Monad_List_Trans_functorListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_functorListT.Do(func() {
		cache_Control_Monad_List_Trans_functorListT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_functorListT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_List_Trans_functorListT
}

var cache_Control_Monad_List_Trans_fromEffect gopurs_runtime.Value
var once_Control_Monad_List_Trans_fromEffect sync.Once
func Get_Control_Monad_List_Trans_fromEffect() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_fromEffect.Do(func() {
		cache_Control_Monad_List_Trans_fromEffect = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_fromEffect
}

var cache_Control_Monad_List_Trans_monadTransListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadTransListT sync.Once
func Get_Control_Monad_List_Trans_monadTransListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadTransListT.Do(func() {
		cache_Control_Monad_List_Trans_monadTransListT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
})}))}
	})
	return cache_Control_Monad_List_Trans_monadTransListT
}

var cache_Control_Monad_List_Trans_lift gopurs_runtime.Value
var once_Control_Monad_List_Trans_lift sync.Once
func Get_Control_Monad_List_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_lift.Do(func() {
		cache_Control_Monad_List_Trans_lift = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_lift
}

var cache_Control_Monad_List_Trans_foldlRec_prime_ gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldlRec_prime_ sync.Once
func Get_Control_Monad_List_Trans_foldlRec_prime_() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldlRec_prime_.Do(func() {
		cache_Control_Monad_List_Trans_foldlRec_prime_ = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldlRec_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldlRec_prime_
}

var cache_Control_Monad_List_Trans_runListTRec gopurs_runtime.Value
var once_Control_Monad_List_Trans_runListTRec sync.Once
func Get_Control_Monad_List_Trans_runListTRec() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_runListTRec.Do(func() {
		cache_Control_Monad_List_Trans_runListTRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_runListTRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_List_Trans_runListTRec
}

var cache_Control_Monad_List_Trans_foldlRec gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldlRec sync.Once
func Get_Control_Monad_List_Trans_foldlRec() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldlRec.Do(func() {
		cache_Control_Monad_List_Trans_foldlRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldlRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldlRec
}

var cache_Control_Monad_List_Trans_foldl_prime_ gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldl_prime_ sync.Once
func Get_Control_Monad_List_Trans_foldl_prime_() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldl_prime_.Do(func() {
		cache_Control_Monad_List_Trans_foldl_prime_ = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldl_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldl_prime_
}

var cache_Control_Monad_List_Trans_runListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_runListT sync.Once
func Get_Control_Monad_List_Trans_runListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_runListT.Do(func() {
		cache_Control_Monad_List_Trans_runListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_runListT(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_runListT
}

var cache_Control_Monad_List_Trans_foldl gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldl sync.Once
func Get_Control_Monad_List_Trans_foldl() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldl.Do(func() {
		cache_Control_Monad_List_Trans_foldl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldl(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldl
}

var cache_Control_Monad_List_Trans_filter gopurs_runtime.Value
var once_Control_Monad_List_Trans_filter sync.Once
func Get_Control_Monad_List_Trans_filter() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_filter.Do(func() {
		cache_Control_Monad_List_Trans_filter = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_filter
}

var cache_Control_Monad_List_Trans_dropWhile gopurs_runtime.Value
var once_Control_Monad_List_Trans_dropWhile sync.Once
func Get_Control_Monad_List_Trans_dropWhile() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_dropWhile.Do(func() {
		cache_Control_Monad_List_Trans_dropWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_dropWhile(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_dropWhile
}

var cache_Control_Monad_List_Trans_drop gopurs_runtime.Value
var once_Control_Monad_List_Trans_drop sync.Once
func Get_Control_Monad_List_Trans_drop() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_drop.Do(func() {
		cache_Control_Monad_List_Trans_drop = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_drop(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_drop
}

var cache_Control_Monad_List_Trans_cons gopurs_runtime.Value
var once_Control_Monad_List_Trans_cons sync.Once
func Get_Control_Monad_List_Trans_cons() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_cons.Do(func() {
		cache_Control_Monad_List_Trans_cons = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_cons(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), lh_1_box, t_2_box)
})
	})
	return cache_Control_Monad_List_Trans_cons
}

var cache_Control_Monad_List_Trans_unfoldable1ListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfoldable1ListT sync.Once
func Get_Control_Monad_List_Trans_unfoldable1ListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfoldable1ListT.Do(func() {
		cache_Control_Monad_List_Trans_unfoldable1ListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_unfoldable1ListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_unfoldable1ListT
}

var cache_Control_Monad_List_Trans_unfoldableListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfoldableListT sync.Once
func Get_Control_Monad_List_Trans_unfoldableListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfoldableListT.Do(func() {
		cache_Control_Monad_List_Trans_unfoldableListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_unfoldableListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_unfoldableListT
}

var cache_Control_Monad_List_Trans_semigroupListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_semigroupListT sync.Once
func Get_Control_Monad_List_Trans_semigroupListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_semigroupListT.Do(func() {
		cache_Control_Monad_List_Trans_semigroupListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_semigroupListT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_List_Trans_semigroupListT
}

var cache_Control_Monad_List_Trans_concat gopurs_runtime.Value
var once_Control_Monad_List_Trans_concat sync.Once
func Get_Control_Monad_List_Trans_concat() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_concat.Do(func() {
		cache_Control_Monad_List_Trans_concat = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_concat
}

var cache_Control_Monad_List_Trans_monoidListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monoidListT sync.Once
func Get_Control_Monad_List_Trans_monoidListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monoidListT.Do(func() {
		cache_Control_Monad_List_Trans_monoidListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_monoidListT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_List_Trans_monoidListT
}

var cache_Control_Monad_List_Trans_catMaybes gopurs_runtime.Value
var once_Control_Monad_List_Trans_catMaybes sync.Once
func Get_Control_Monad_List_Trans_catMaybes() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_catMaybes.Do(func() {
		cache_Control_Monad_List_Trans_catMaybes = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
})
	})
	return cache_Control_Monad_List_Trans_catMaybes
}

var cache_Control_Monad_List_Trans_monadListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadListT sync.Once
func Get_Control_Monad_List_Trans_monadListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadListT.Do(func() {
		cache_Control_Monad_List_Trans_monadListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_monadListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_monadListT
}

var cache_Control_Monad_List_Trans_bindListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_bindListT sync.Once
func Get_Control_Monad_List_Trans_bindListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_bindListT.Do(func() {
		cache_Control_Monad_List_Trans_bindListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_bindListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_bindListT
}

var cache_Control_Monad_List_Trans_applyListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_applyListT sync.Once
func Get_Control_Monad_List_Trans_applyListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_applyListT.Do(func() {
		cache_Control_Monad_List_Trans_applyListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applyListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_applyListT
}

var cache_Control_Monad_List_Trans_applicativeListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_applicativeListT sync.Once
func Get_Control_Monad_List_Trans_applicativeListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_applicativeListT.Do(func() {
		cache_Control_Monad_List_Trans_applicativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_applicativeListT
}

var cache_Control_Monad_List_Trans_monadEffectListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadEffectListT sync.Once
func Get_Control_Monad_List_Trans_monadEffectListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadEffectListT.Do(func() {
		cache_Control_Monad_List_Trans_monadEffectListT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_monadEffectListT(dictMonadEffect_0_box)
})
	})
	return cache_Control_Monad_List_Trans_monadEffectListT
}

var cache_Control_Monad_List_Trans_monadSTListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadSTListT sync.Once
func Get_Control_Monad_List_Trans_monadSTListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadSTListT.Do(func() {
		cache_Control_Monad_List_Trans_monadSTListT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_monadSTListT(dictMonadST_0_box)
})
	})
	return cache_Control_Monad_List_Trans_monadSTListT
}

var cache_Control_Monad_List_Trans_altListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_altListT sync.Once
func Get_Control_Monad_List_Trans_altListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_altListT.Do(func() {
		cache_Control_Monad_List_Trans_altListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_altListT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_List_Trans_altListT
}

var cache_Control_Monad_List_Trans_plusListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_plusListT sync.Once
func Get_Control_Monad_List_Trans_plusListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_plusListT.Do(func() {
		cache_Control_Monad_List_Trans_plusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_plusListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_plusListT
}

var cache_Control_Monad_List_Trans_alternativeListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_alternativeListT sync.Once
func Get_Control_Monad_List_Trans_alternativeListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_alternativeListT.Do(func() {
		cache_Control_Monad_List_Trans_alternativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_alternativeListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_alternativeListT
}

var cache_Control_Monad_List_Trans_monadPlusListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadPlusListT sync.Once
func Get_Control_Monad_List_Trans_monadPlusListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadPlusListT.Do(func() {
		cache_Control_Monad_List_Trans_monadPlusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_monadPlusListT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_List_Trans_monadPlusListT
}

type Constructor_Control_Monad_List_Trans_Yield[T_a any, T_s any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}


type Constructor_Control_Monad_List_Trans_Skip[T_a any, T_s any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Control_Monad_List_Trans_Done[T_a any, T_s any] struct {
	Rc uint32
}


func Call_Control_Monad_List_Trans_ListT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_List_Trans_wrapLazy(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_1}))})
}

func Call_Control_Monad_List_Trans_wrapEffect(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))}))}
}), v_1)
}

func Call_Control_Monad_List_Trans_unfold(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
unfold:
for {
if false { continue unfold }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4))
if (__t_tag_1 != nil) {
// TAST (Let): __local_var_5_2 shape=Other bindingType=Any
__local_var_5_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_2
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), f_2, __local_var_5_2)
}))}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4))
if (__t_tag_3 == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}), gopurs_runtime.Apply(f_2, z_3))
})
}
}

func Call_Control_Monad_List_Trans_uncons(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
uncons:
for {
if false { continue uncons }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 shape=Other bindingType=(Func [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])])] (TypeApp (TypeVar f) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])])]))
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f)])
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_4, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1320412129) {
__t3 = gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}}))})
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 813447293) {
__t3 = gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 489128924) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
})
}
}

func Call_Control_Monad_List_Trans_tail(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0.UnsafePtr).V1, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_2))
})
}

func Call_Control_Monad_List_Trans_stepMap(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_1, v_2)
}

func Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])]))
__local_var_5_1 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0), f_2)
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_2 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_6_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_2))
}))}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
// TAST (Let): __local_var_5_4 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])]))
__local_var_5_4 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0), f_2)
_ = __local_var_5_4
// TAST (Let): __local_var_6_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_5 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
_ = __local_var_6_5
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_5))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), v_3)
})
}
}

func Call_Control_Monad_List_Trans_scanl(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_1 shape=Other bindingType=Any
__local_var_6_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 1320412129) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_1785332133_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_2, __local_var_6_1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1)}))}, __local_var_6_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}}))}
goto end_branch_2
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_1785332133_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0)}))}, __local_var_6_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}}))}
goto end_branch_2
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 489128924) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_2549197956_3094389156(Rebox_Control_Monad_List_Trans_3094389156_2549197956(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_2549197956_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]] { panic("Failed pattern match") }()))}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_2549197956_3094389156(Rebox_Control_Monad_List_Trans_3094389156_2549197956(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))))}
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{b_3, l_4}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})
}

func Call_Control_Monad_List_Trans_prepend_prime_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_1, t_2}))})
}

func Call_Control_Monad_List_Trans_prepend(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return t_2
}))}))})
}

func Call_Control_Monad_List_Trans_nil(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_Control_Monad_List_Trans_singleton(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): nil1_1_0 shape=App(Other) bindingType=(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}))}))})
})
}

func Call_Control_Monad_List_Trans_take(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
take:
for {
if false { continue take }
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): nil1_1_0 shape=App(Other) bindingType=(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_3.IntVal) == (int64(0)) {
__t7 = nil1_1_0
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1320412129) {
// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])]))
__local_var_6_2 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_take(dictApplicative_0), gopurs_runtime.Int((v_3.IntVal) - (int64(1))))
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_3 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1
_ = __local_var_7_3
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_3))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
// TAST (Let): __local_var_6_4 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])]))
__local_var_6_4 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_take(dictApplicative_0), gopurs_runtime.Int(v_3.IntVal))
_ = __local_var_6_4
// TAST (Let): __local_var_7_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_5 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0
_ = __local_var_7_5
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_5))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 489128924) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), v1_4)
}
end_branch_7:
return __t7
})
}
}

func Call_Control_Monad_List_Trans_zipWith_prime_(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
zipWith_prime_:
for {
if false { continue zipWith_prime_ }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=Any
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): prepend_prime_1_3_2 shape=Let(Abs(Abs(App(Other)))) bindingType=(Func [(TypeVar c), (ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar c), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar c)])])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar c), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar c)])])]))
prepend_prime_1_3_2 := gopurs_runtime.Func2(func(h_4 gopurs_runtime.Value, t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_4, t_5}))})
})
_ = prepend_prime_1_3_2
// TAST (Let): Bind1_4_4 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f)])
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_4_4
return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, fa_6 gopurs_runtime.Value, fb_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
}))}))}
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), fa_6), gopurs_runtime.Func(func(ua_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), fb_7), gopurs_runtime.Func(func(ub_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ub_9))
if (__t_tag_5 == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ua_8))
if (__t_tag_6 == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ua_8))
var __t_and_9 bool = false
if (__t_tag_7 != nil) {

var __t_tag_8 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ub_9))
__t_and_9 = (__t_tag_8 != nil)
}
if __t_and_9 {
// TAST (Let): __local_var_10_10 shape=Other bindingType=Any
__local_var_10_10 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ua_8.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_10_10
// TAST (Let): __local_var_11_11 shape=Other bindingType=Any
__local_var_11_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ub_9.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_11
// TAST (Let): __local_var_12_12 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar c), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar c)])])])])
__local_var_12_12 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Control_Monad_List_Trans_zipWith_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_5, __local_var_10_10, __local_var_11_11)
}))
_ = __local_var_12_12
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(prepend_prime_1_3_2, a_13, __local_var_12_12)
}), gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ua_8.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ub_9.UnsafePtr).V0.UnsafePtr).V0))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}))
})))
})
}
}

func Call_Control_Monad_List_Trans_zipWith(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 shape=Other bindingType=(Func [(TypeVar c)] (TypeApp (TypeVar f) [(TypeVar c)]))
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_zipWith_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Apply2(f_2, a_3, b_4))
}))
})
}

func Call_Control_Monad_List_Trans_mapMaybe(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))
_ = __local_var_4_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_4_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(Get_Control_Monad_List_Trans_Yield(), (__local_var_4_1).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_2:
// TAST (Let): __local_var_4_0 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])] (ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])]))])
__local_var_4_0 := __t2
_ = __local_var_4_0
var __t3 gopurs_runtime.Value
{
if (__local_var_4_0 == nil) {
__t3 = Get_Control_Monad_List_Trans_Skip()
goto end_branch_3
} else {

}
}
{
if (__local_var_4_0 != nil) {
__t3 = (__local_var_4_0).V0
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
// TAST (Let): __local_var_5_4 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_5_4 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1
_ = __local_var_5_4
__t6 = gopurs_runtime.Apply(__t3, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_mapMaybe(dictFunctor_0, f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_4))
})))
goto end_branch_6
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
// TAST (Let): __local_var_4_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_4_5 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
_ = __local_var_4_5
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_mapMaybe(dictFunctor_0, f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_5))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), v_2)
}
}

func Call_Control_Monad_List_Trans_iterate(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 shape=Other bindingType=(Func [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar a)])])] (TypeApp (TypeVar f) [(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar a)])])]))
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, x_4), x_4}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))})
}), a_3)
})
}

func Call_Control_Monad_List_Trans_repeat(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply(Call_Control_Monad_List_Trans_iterate(dictMonad_0), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Control_Monad_List_Trans_head(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_2))
})
}

func Call_Control_Monad_List_Trans_functorListT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
functorListT:
for {
if false { continue functorListT }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(dictFunctor_0), "map"), f_1)
_ = __local_var_4_0
// TAST (Let): __local_var_5_1 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_5_1 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1
_ = __local_var_5_1
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_1))
}))}))}
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(dictFunctor_0), "map"), f_1)
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_5_3 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
_ = __local_var_5_3
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_3))
}))}))}
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}), v_2)
})}))}
}
}

func Call_Control_Monad_List_Trans_fromEffect(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}))
_ = __local_var_3_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, __local_var_3_1}))}
}), fa_2)
})
}

func Call_Control_Monad_List_Trans_lift(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
return Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{})))
}

func Call_Control_Monad_List_Trans_foldlRec_prime_(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f)])
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Monad01_4_3 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar f)])
Monad01_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 shape=Other bindingType=(TypeVar a)
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_10))
if (__t_tag_5 == nil) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_1239976062_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, __local_var_9_4})))})
goto end_branch_8
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_10))
if (__t_tag_6 != nil) {
// TAST (Let): __local_var_11_7 shape=Other bindingType=Any
__local_var_11_7 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_7
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(f_5, __local_var_9_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_784458114_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, func() struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
} {
					orig := gopurs_runtime.RecordDict2("a", "b", b_prime__12, __local_var_11_7)
					_ = orig
					clone := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a")
					clone.b = gopurs_runtime.RecordGet(orig, "b")
					return clone
				}()})))})
}))
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
}), func() gopurs_runtime.Value {
				orig := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{a_6, b_7}
				_ = orig
				return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
				}())
})
}

func Call_Control_Monad_List_Trans_runListTRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Apply3(Get_Control_Monad_List_Trans_foldlRec_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
}), Get_Data_Unit_unit())
}

func Call_Control_Monad_List_Trans_foldlRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f)])
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Monad01_4_3 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar f)])
Monad01_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 shape=Other bindingType=(TypeVar a)
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_10))
if (__t_tag_5 == nil) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_1239976062_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, __local_var_9_4})))})
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_10))
if (__t_tag_6 != nil) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_784458114_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, func() struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
} {
					orig := gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Apply2(f_5, __local_var_9_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V1)
					_ = orig
					clone := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a")
					clone.b = gopurs_runtime.RecordGet(orig, "b")
					return clone
				}()})))})
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
}), func() gopurs_runtime.Value {
				orig := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{a_6, b_7}
				_ = orig
				return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
				}())
})
}

func Call_Control_Monad_List_Trans_foldl_prime_(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_0 gopurs_runtime.Value
_ = loop_4_2_0
// FALLBACK TCO: isLoop=false len=1
loop_4_2_0 = gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_7))
if (__t_tag_3 == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_5)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_7))
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(loop_4_2_0, a_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1)
}))
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
})
return loop_4_2_0
})
}

func Call_Control_Monad_List_Trans_runListT(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Apply3(Get_Control_Monad_List_Trans_foldl_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
}), Get_Data_Unit_unit())
}

func Call_Control_Monad_List_Trans_foldl(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_1 gopurs_runtime.Value
_ = loop_4_2_1
// FALLBACK TCO: isLoop=false len=1
loop_4_2_1 = gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_7))
if (__t_tag_3 == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_5)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_7))
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Apply2(loop_4_2_1, gopurs_runtime.Apply2(f_3, b_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1)
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
})
return loop_4_2_1
})
}

func Call_Control_Monad_List_Trans_filter(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
// TAST (Let): s_prime__4_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeVar b)])
s_prime__4_0 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_filter(dictFunctor_0, f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
}))
_ = s_prime__4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_prime__4_0}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_prime__4_0}))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_filter(dictFunctor_0, f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))
}))}))}
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), v_2)
}
}

func Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
dropWhile:
for {
if false { continue dropWhile }
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])]))
__local_var_5_1 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0), f_2)
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_2 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_6_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_2))
}))}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1}))}
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
// TAST (Let): __local_var_5_4 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])]))
__local_var_5_4 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0), f_2)
_ = __local_var_5_4
// TAST (Let): __local_var_6_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_5 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
_ = __local_var_6_5
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_5))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), v_3)
})
}
}

func Call_Control_Monad_List_Trans_drop(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_2.IntVal) == (int64(0)) {
__t6 = v1_3
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1320412129) {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])]))
__local_var_5_1 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_drop(dictApplicative_0), gopurs_runtime.Int((v_2.IntVal) - (int64(1))))
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_2 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V1
_ = __local_var_6_2
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_2))
}))}))}
goto end_branch_5
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
// TAST (Let): __local_var_5_3 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])]))
__local_var_5_3 := gopurs_runtime.Apply(Call_Control_Monad_List_Trans_drop(dictApplicative_0), gopurs_runtime.Int(v_2.IntVal))
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_4 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0
_ = __local_var_6_4
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4))
}))}))}
goto end_branch_5
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 489128924) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}), v1_3)
}
end_branch_6:
return __t6
})
}
}

func Call_Control_Monad_List_Trans_cons(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], lh_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 gopurs_runtime.Value = lh_1_loop
_ = lh_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), lh_1), t_2}))})
}

func Call_Control_Monad_List_Trans_unfoldable1ListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_2 gopurs_runtime.Value
_ = go__go_4_1_2
// FALLBACK TCO: isLoop=false len=1
go__go_4_1_2 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Call_Control_Monad_List_Trans_singleton(Applicative0_1_0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
// TAST (Let): __local_var_6_4 shape=Other bindingType=Any
__local_var_6_4 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1.UnsafePtr).V0
_ = __local_var_6_4
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
}))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_4_1_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_3804580809_138441832(Rebox_Control_Monad_List_Trans_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_2, __local_var_6_4)))))})
}))}))})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})
return gopurs_runtime.Apply(go__go_4_1_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_3804580809_138441832(Rebox_Control_Monad_List_Trans_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_2, b_3)))))})
})}))}
}

func Call_Control_Monad_List_Trans_unfoldableListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Applicative0_2_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): unfoldable1ListT1_2_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Unfoldable1","Unfoldable1"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
unfoldable1ListT1_2_1 := (&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_3_3 gopurs_runtime.Value
_ = go__go_5_3_3
// FALLBACK TCO: isLoop=false len=1
go__go_5_3_3 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t7 = gopurs_runtime.Apply(Call_Control_Monad_List_Trans_singleton(Applicative0_2_2), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136 && __t_tag_5.UnsafePtr != nil) {
// TAST (Let): __local_var_7_6 shape=Other bindingType=Any
__local_var_7_6 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1.UnsafePtr).V0
_ = __local_var_7_6
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0
}))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_5_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_3804580809_138441832(Rebox_Control_Monad_List_Trans_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, __local_var_7_6)))))})
}))}))})
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
return gopurs_runtime.Apply(go__go_5_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_3804580809_138441832(Rebox_Control_Monad_List_Trans_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, b_4)))))})
})})
_ = unfoldable1ListT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(unfoldable1ListT1_2_1)}
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_8_4 gopurs_runtime.Value
_ = go__go_5_8_4
// FALLBACK TCO: isLoop=false len=1
go__go_5_8_4 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
var __t_tag_9 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_6))
if (__t_tag_9 == nil) {
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
goto end_branch_12
} else {

}
}
{
var __t_tag_10 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_6))
if (__t_tag_10 != nil) {
// TAST (Let): __local_var_7_11 shape=Other bindingType=Any
__local_var_7_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_7_11
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0.UnsafePtr).V0
}))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_5_8_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, __local_var_7_11)))))})
}))}))})
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
return gopurs_runtime.Apply(go__go_5_8_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_4010058633_3094389156(Rebox_Control_Monad_List_Trans_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, b_4)))))})
})}))}
}

func Call_Control_Monad_List_Trans_semigroupListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}))}
}

func Call_Control_Monad_List_Trans_concat(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
// TAST (Let): __local_var_5_1 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_5_1 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_5_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_semigroupListT(gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}), "append"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_1), y_3)
}))}))}
goto end_branch_3
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
// TAST (Let): __local_var_5_2 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_5_2 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
_ = __local_var_5_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_semigroupListT(gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}), "append"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2), y_3)
}))}))}
goto end_branch_3
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return y_3
}))}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), x_2)
})
}

func Call_Control_Monad_List_Trans_monoidListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): semigroupListT1_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
semigroupListT1_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))})
_ = semigroupListT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupListT1_1_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})}))}
}

func Call_Control_Monad_List_Trans_catMaybes(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Control_Monad_List_Trans_monadListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_bindListT(dictMonad_0)))}
})}))}
}

func Call_Control_Monad_List_Trans_bindListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): semigroupListT1_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applyListT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(fa_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1320412129) {
// TAST (Let): __local_var_6_2 shape=Other bindingType=Any
__local_var_6_2 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_3 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_7_3
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_1_0.V0), gopurs_runtime.Apply(f_4, __local_var_6_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_3), f_4))
}))}))}
goto end_branch_5
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 813447293) {
// TAST (Let): __local_var_6_4 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_4 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_4
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4), f_4)
}))}))}
goto end_branch_5
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 489128924) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}), fa_3)
})}))}
}

func Call_Control_Monad_List_Trans_applyListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorListT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_1_1), "map"), f_2)
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_3 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_6_3
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_1_1), "map"), f_2)
_ = __local_var_5_4
// TAST (Let): __local_var_6_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_5 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
_ = __local_var_6_5
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_5))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), v_3)
})})
_ = functorListT1_1_0
// TAST (Let): __local_var_2_7 shape=App(Var) bindingType=Any
__local_var_2_7 := Call_Control_Monad_List_Trans_monadListT(dictMonad_0)
_ = __local_var_2_7
// TAST (Let): Bind1_3_8 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_8
// TAST (Let): Applicative0_4_9 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_9
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_1_0)}
}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_8.V1), f_5, gopurs_runtime.Func(func(f_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_8.V1), a_6, gopurs_runtime.Func(func(a_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_9.V1), gopurs_runtime.Apply(f_prime__7, a_prime__8))
}))
}))
})}))}
}

func Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applyListT(dictMonad_0)))}
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))}))}
}

func Call_Control_Monad_List_Trans_monadEffectListT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadListT1_2_1 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar m), (TypeVar a)])])])])
monadListT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorListT1_4_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_4_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1320412129) {
// TAST (Let): __local_var_8_4 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_4 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_8_4
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1320412129) {
// TAST (Let): __local_var_11_5 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_3), "map"), f_5)
_ = __local_var_11_5
// TAST (Let): __local_var_12_6 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_6 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_12_6
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_6))
}))}))}
goto end_branch_9
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 813447293) {
// TAST (Let): __local_var_11_7 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_3), "map"), f_5)
_ = __local_var_11_7
// TAST (Let): __local_var_12_8 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_8 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0
_ = __local_var_12_8
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_7, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_8))
}))}))}
goto end_branch_9
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 489128924) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_4))
}))}))}
goto end_branch_16
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 813447293) {
// TAST (Let): __local_var_8_10 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_10 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_10
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1320412129) {
// TAST (Let): __local_var_11_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_3), "map"), f_5)
_ = __local_var_11_11
// TAST (Let): __local_var_12_12 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_12 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_12_12
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_11, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_12))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 813447293) {
// TAST (Let): __local_var_11_13 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_3), "map"), f_5)
_ = __local_var_11_13
// TAST (Let): __local_var_12_14 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_14 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0
_ = __local_var_12_14
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_13, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_14))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 489128924) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_10))
}))}))}
goto end_branch_16
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 489128924) {
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}), v_6)
})})
_ = functorListT1_4_2
// TAST (Let): __local_var_5_17 shape=App(Var) bindingType=Any
__local_var_5_17 := Call_Control_Monad_List_Trans_monadListT(Monad0_1_0)
_ = __local_var_5_17
// TAST (Let): Bind1_6_18 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_4_2)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime__10, a_prime__11))
}))
}))
})}))}
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))}))}
}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_3_20 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_3_20 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_3_20
// TAST (Let): Functor0_4_21 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_21
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_23 shape=App(Other) bindingType=Any
__local_var_6_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_23
// TAST (Let): functorListT1_6_22 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_6_22 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1320412129) {
// TAST (Let): __local_var_10_24 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_24 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_10_24
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_25 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_23), "map"), f_7)
_ = __local_var_13_25
// TAST (Let): __local_var_14_26 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_26 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_26
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_25, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_26))
}))}))}
goto end_branch_29
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_27 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_23), "map"), f_7)
_ = __local_var_13_27
// TAST (Let): __local_var_14_28 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_28 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_14_28
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_27, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_28))
}))}))}
goto end_branch_29
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return __t29
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_24))
}))}))}
goto end_branch_36
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 813447293) {
// TAST (Let): __local_var_10_30 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_30 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_10_30
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_31 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_23), "map"), f_7)
_ = __local_var_13_31
// TAST (Let): __local_var_14_32 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_32 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_32
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_31, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_32))
}))}))}
goto end_branch_35
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_33 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_23), "map"), f_7)
_ = __local_var_13_33
// TAST (Let): __local_var_14_34 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_34 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_14_34
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_33, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_34))
}))}))}
goto end_branch_35
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
return __t35
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_30))
}))}))}
goto end_branch_36
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 489128924) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_36
} else {

}
}
{
__t36 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_36:
return __t36
}), v_8)
})})
_ = functorListT1_6_22
// TAST (Let): __local_var_7_37 shape=App(Var) bindingType=Any
__local_var_7_37 := Call_Control_Monad_List_Trans_monadListT(Monad0_1_0)
_ = __local_var_7_37
// TAST (Let): Bind1_8_38 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_8_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_38
// TAST (Let): Applicative0_9_39 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_39
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_6_22)}
}), gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), f_10, gopurs_runtime.Func(func(f_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), a_11, gopurs_runtime.Func(func(a_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_39.V1), gopurs_runtime.Apply(f_prime__12, a_prime__13))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(fa_5 gopurs_runtime.Value, f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_21.V0), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1320412129) {
// TAST (Let): __local_var_8_40 shape=Other bindingType=Any
__local_var_8_40 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_40
// TAST (Let): __local_var_9_41 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_9_41 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_9_41
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_3_20.V0), gopurs_runtime.Apply(f_6, __local_var_8_40), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(Monad0_1_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_41), f_6))
}))}))}
goto end_branch_43
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 813447293) {
// TAST (Let): __local_var_8_42 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_42 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_42
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(Monad0_1_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_42), f_6)
}))}))}
goto end_branch_43
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 489128924) {
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_43
} else {

}
}
{
__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_43:
return __t43
}), fa_5)
})}))}
})})
_ = monadListT1_2_1
// TAST (Let): __local_var_3_44 shape=App(Var) bindingType=Any
__local_var_3_44 := Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))
_ = __local_var_3_44
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadListT1_2_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_44, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
})}))}
}

func Call_Control_Monad_List_Trans_monadSTListT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadListT1_2_1 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar m), (TypeVar a)])])])])
monadListT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorListT1_4_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_4_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1320412129) {
// TAST (Let): __local_var_8_4 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_4 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_8_4
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1320412129) {
// TAST (Let): __local_var_11_5 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_3), "map"), f_5)
_ = __local_var_11_5
// TAST (Let): __local_var_12_6 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_6 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_12_6
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_6))
}))}))}
goto end_branch_9
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 813447293) {
// TAST (Let): __local_var_11_7 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_3), "map"), f_5)
_ = __local_var_11_7
// TAST (Let): __local_var_12_8 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_8 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0
_ = __local_var_12_8
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_7, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_8))
}))}))}
goto end_branch_9
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 489128924) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_4))
}))}))}
goto end_branch_16
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 813447293) {
// TAST (Let): __local_var_8_10 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_10 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_10
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1320412129) {
// TAST (Let): __local_var_11_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_3), "map"), f_5)
_ = __local_var_11_11
// TAST (Let): __local_var_12_12 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_12 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_12_12
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_11, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_12))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 813447293) {
// TAST (Let): __local_var_11_13 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_3), "map"), f_5)
_ = __local_var_11_13
// TAST (Let): __local_var_12_14 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_14 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0
_ = __local_var_12_14
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_13, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_14))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 489128924) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_10))
}))}))}
goto end_branch_16
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 489128924) {
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}), v_6)
})})
_ = functorListT1_4_2
// TAST (Let): __local_var_5_17 shape=App(Var) bindingType=Any
__local_var_5_17 := Call_Control_Monad_List_Trans_monadListT(Monad0_1_0)
_ = __local_var_5_17
// TAST (Let): Bind1_6_18 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_4_2)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime__10, a_prime__11))
}))
}))
})}))}
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))}))}
}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_3_20 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_3_20 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_3_20
// TAST (Let): Functor0_4_21 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_21
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_23 shape=App(Other) bindingType=Any
__local_var_6_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_23
// TAST (Let): functorListT1_6_22 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_6_22 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1320412129) {
// TAST (Let): __local_var_10_24 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_24 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_10_24
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_25 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_23), "map"), f_7)
_ = __local_var_13_25
// TAST (Let): __local_var_14_26 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_26 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_26
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_25, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_26))
}))}))}
goto end_branch_29
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_27 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_23), "map"), f_7)
_ = __local_var_13_27
// TAST (Let): __local_var_14_28 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_28 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_14_28
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_27, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_28))
}))}))}
goto end_branch_29
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return __t29
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_24))
}))}))}
goto end_branch_36
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 813447293) {
// TAST (Let): __local_var_10_30 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_30 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_10_30
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_31 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_23), "map"), f_7)
_ = __local_var_13_31
// TAST (Let): __local_var_14_32 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_32 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_32
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_31, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_32))
}))}))}
goto end_branch_35
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_33 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_23), "map"), f_7)
_ = __local_var_13_33
// TAST (Let): __local_var_14_34 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_34 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_14_34
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_33, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_34))
}))}))}
goto end_branch_35
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
return __t35
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_30))
}))}))}
goto end_branch_36
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 489128924) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_36
} else {

}
}
{
__t36 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_36:
return __t36
}), v_8)
})})
_ = functorListT1_6_22
// TAST (Let): __local_var_7_37 shape=App(Var) bindingType=Any
__local_var_7_37 := Call_Control_Monad_List_Trans_monadListT(Monad0_1_0)
_ = __local_var_7_37
// TAST (Let): Bind1_8_38 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_8_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_38
// TAST (Let): Applicative0_9_39 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_39
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_6_22)}
}), gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), f_10, gopurs_runtime.Func(func(f_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), a_11, gopurs_runtime.Func(func(a_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_39.V1), gopurs_runtime.Apply(f_prime__12, a_prime__13))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(fa_5 gopurs_runtime.Value, f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_21.V0), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1320412129) {
// TAST (Let): __local_var_8_40 shape=Other bindingType=Any
__local_var_8_40 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_40
// TAST (Let): __local_var_9_41 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_9_41 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_9_41
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_3_20.V0), gopurs_runtime.Apply(f_6, __local_var_8_40), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(Monad0_1_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_41), f_6))
}))}))}
goto end_branch_43
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 813447293) {
// TAST (Let): __local_var_8_42 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_42 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_42
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(Monad0_1_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_42), f_6)
}))}))}
goto end_branch_43
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 489128924) {
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_43
} else {

}
}
{
__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_43:
return __t43
}), fa_5)
})}))}
})})
_ = monadListT1_2_1
// TAST (Let): __local_var_3_44 shape=App(Var) bindingType=Any
__local_var_3_44 := Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))
_ = __local_var_3_44
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadListT1_2_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_44, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
})}))}
}

func Call_Control_Monad_List_Trans_altListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorListT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_1_1), "map"), f_2)
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_3 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_6_3
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_1_1), "map"), f_2)
_ = __local_var_5_4
// TAST (Let): __local_var_6_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_5 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
_ = __local_var_6_5
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_5))
}))}))}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), v_3)
})})
_ = functorListT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_1_0)}
}), Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}))}
}

func Call_Control_Monad_List_Trans_plusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=Any
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorListT1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1320412129) {
// TAST (Let): __local_var_6_4 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_4 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_6_4
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1320412129) {
// TAST (Let): __local_var_9_5 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_2_3), "map"), f_3)
_ = __local_var_9_5
// TAST (Let): __local_var_10_6 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_6 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_6
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_6))
}))}))}
goto end_branch_9
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 813447293) {
// TAST (Let): __local_var_9_7 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_2_3), "map"), f_3)
_ = __local_var_9_7
// TAST (Let): __local_var_10_8 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_8 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_10_8
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_7, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_8))
}))}))}
goto end_branch_9
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 489128924) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4))
}))}))}
goto end_branch_16
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 813447293) {
// TAST (Let): __local_var_6_10 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_10 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_10
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1320412129) {
// TAST (Let): __local_var_9_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_2_3), "map"), f_3)
_ = __local_var_9_11
// TAST (Let): __local_var_10_12 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_12 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_12
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_11, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_12))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 813447293) {
// TAST (Let): __local_var_9_13 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_2_3), "map"), f_3)
_ = __local_var_9_13
// TAST (Let): __local_var_10_14 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_14 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_10_14
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_13, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_14))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 489128924) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_10))
}))}))}
goto end_branch_16
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 489128924) {
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}), v_4)
})})
_ = functorListT1_2_2
// TAST (Let): altListT1_2_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
altListT1_2_1 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_2_2)}
}), Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_1_0))})
_ = altListT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altListT1_2_1)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})}))}
}

func Call_Control_Monad_List_Trans_alternativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeListT1_1_0 shape=LitRecord bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
applicativeListT1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorListT1_2_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_2_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1320412129) {
// TAST (Let): __local_var_6_3 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_3 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_6_3
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1320412129) {
// TAST (Let): __local_var_9_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_2_2), "map"), f_3)
_ = __local_var_9_4
// TAST (Let): __local_var_10_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_5 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_5
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_5))
}))}))}
goto end_branch_8
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 813447293) {
// TAST (Let): __local_var_9_6 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_2_2), "map"), f_3)
_ = __local_var_9_6
// TAST (Let): __local_var_10_7 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_7 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_10_7
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_6, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_7))
}))}))}
goto end_branch_8
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 489128924) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_3))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 813447293) {
// TAST (Let): __local_var_6_9 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_6_9 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_9
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1320412129) {
// TAST (Let): __local_var_9_10 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_2_2), "map"), f_3)
_ = __local_var_9_10
// TAST (Let): __local_var_10_11 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_11 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_11
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_10, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_11))
}))}))}
goto end_branch_14
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 813447293) {
// TAST (Let): __local_var_9_12 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_2_2), "map"), f_3)
_ = __local_var_9_12
// TAST (Let): __local_var_10_13 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_13 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_10_13
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_12, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_13))
}))}))}
goto end_branch_14
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 489128924) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_9))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 489128924) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}), v_4)
})})
_ = functorListT1_2_1
// TAST (Let): __local_var_3_16 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_3_16 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_4_17 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_4_17 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_4_17
// TAST (Let): Functor0_5_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_5_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_18
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applyListT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(fa_6 gopurs_runtime.Value, f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_18.V0), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1320412129) {
// TAST (Let): __local_var_9_19 shape=Other bindingType=Any
__local_var_9_19 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_9_19
// TAST (Let): __local_var_10_20 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_20 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_20
__t22 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_4_17.V0), gopurs_runtime.Apply(f_7, __local_var_9_19), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_20), f_7))
}))}))}
goto end_branch_22
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 813447293) {
// TAST (Let): __local_var_9_21 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_9_21 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_9_21
__t22 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_21), f_7)
}))}))}
goto end_branch_22
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 489128924) {
__t22 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
}), fa_6)
})}))}
})})
_ = __local_var_3_16
// TAST (Let): Bind1_4_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_16.V1), gopurs_runtime.Value{}))
_ = Bind1_4_23
// TAST (Let): Applicative0_5_24 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_5_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_16.V0), gopurs_runtime.Value{}))
_ = Applicative0_5_24
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_2_1)}
}), gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_23.V1), f_6, gopurs_runtime.Func(func(f_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_23.V1), a_7, gopurs_runtime.Func(func(a_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_24.V1), gopurs_runtime.Apply(f_prime__8, a_prime__9))
}))
}))
})}))}
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = applicativeListT1_1_0
// TAST (Let): Applicative0_2_26 shape=App(Other) bindingType=Any
Applicative0_2_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_2_26
// TAST (Let): __local_var_3_29 shape=App(Other) bindingType=Any
__local_var_3_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_26, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_29
// TAST (Let): functorListT1_3_28 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_3_28 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_29, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t58 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1320412129) {
// TAST (Let): __local_var_7_30 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_30 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
_ = __local_var_7_30
__t58 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_29, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1320412129) {
// TAST (Let): __local_var_10_31 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_31 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_10_31
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_29, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_32 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_29), "map"), f_4)
_ = __local_var_13_32
// TAST (Let): __local_var_14_33 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_33 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_33
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_32, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_33))
}))}))}
goto end_branch_36
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_34 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_29), "map"), f_4)
_ = __local_var_13_34
// TAST (Let): __local_var_14_35 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_35 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_14_35
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_34, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_35))
}))}))}
goto end_branch_36
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_36
} else {

}
}
{
__t36 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_36:
return __t36
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_31))
}))}))}
goto end_branch_43
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 813447293) {
// TAST (Let): __local_var_10_37 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_37 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_10_37
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_29, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_38 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_38 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_29), "map"), f_4)
_ = __local_var_13_38
// TAST (Let): __local_var_14_39 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_39 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_39
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_38, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_39))
}))}))}
goto end_branch_42
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_40 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_40 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_29), "map"), f_4)
_ = __local_var_13_40
// TAST (Let): __local_var_14_41 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_41 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_14_41
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_40, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_41))
}))}))}
goto end_branch_42
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
return __t42
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_37))
}))}))}
goto end_branch_43
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 489128924) {
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_43
} else {

}
}
{
__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_43:
return __t43
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_30))
}))}))}
goto end_branch_58
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 813447293) {
// TAST (Let): __local_var_7_44 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_44 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0
_ = __local_var_7_44
__t58 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_29, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t57 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1320412129) {
// TAST (Let): __local_var_10_45 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_45 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_10_45
__t57 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_29, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t50 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_46 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_29), "map"), f_4)
_ = __local_var_13_46
// TAST (Let): __local_var_14_47 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_47 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_47
__t50 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_46, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_47))
}))}))}
goto end_branch_50
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_48 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_48 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_29), "map"), f_4)
_ = __local_var_13_48
// TAST (Let): __local_var_14_49 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_49 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_14_49
__t50 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_48, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_49))
}))}))}
goto end_branch_50
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t50 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_50
} else {

}
}
{
__t50 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_50:
return __t50
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_45))
}))}))}
goto end_branch_57
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 813447293) {
// TAST (Let): __local_var_10_51 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_51 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_10_51
__t57 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_29, "map"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t56 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_52 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_52 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_29), "map"), f_4)
_ = __local_var_13_52
// TAST (Let): __local_var_14_53 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_53 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_53
__t56 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_52, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_53))
}))}))}
goto end_branch_56
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_54 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_13_54 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_29), "map"), f_4)
_ = __local_var_13_54
// TAST (Let): __local_var_14_55 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_55 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_14_55
__t56 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_54, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_55))
}))}))}
goto end_branch_56
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t56 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_56
} else {

}
}
{
__t56 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_56:
return __t56
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_51))
}))}))}
goto end_branch_57
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 489128924) {
__t57 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_57
} else {

}
}
{
__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_57:
return __t57
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_44))
}))}))}
goto end_branch_58
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 489128924) {
__t58 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_58
} else {

}
}
{
__t58 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_58:
return __t58
}), v_5)
})})
_ = functorListT1_3_28
// TAST (Let): altListT1_3_27 shape=Let(LitRecord) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
altListT1_3_27 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_3_28)}
}), Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_2_26))})
_ = altListT1_3_27
// TAST (Let): plusListT1_2_25 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
plusListT1_2_25 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altListT1_3_27)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_26, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})})
_ = plusListT1_2_25
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeListT1_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusListT1_2_25)}
})}))}
}

func Call_Control_Monad_List_Trans_monadPlusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadListT1_1_0 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
monadListT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=Any
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorListT1_3_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_3_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_7_3 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_7_3
var __local_var_7_3 gopurs_runtime.Value
_ = __local_var_7_3
Call_local_Control_Monad_List_Trans___local_var_7_3 = func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1320412129) {
// TAST (Let): __local_var_9_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_2), "map"), f_4)
_ = __local_var_9_4
// TAST (Let): __local_var_10_5 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_5 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_5
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_5))
}))}))}
goto end_branch_8
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 813447293) {
// TAST (Let): __local_var_9_6 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_2), "map"), f_4)
_ = __local_var_9_6
// TAST (Let): __local_var_10_7 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_7 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_10_7
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_6, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_7))
}))}))}
goto end_branch_8
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 489128924) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}), v_7)
}
__local_var_7_3 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_7_3(v_7_loop_val)
})
// TAST (Let): __local_var_8_9 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_9 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
_ = __local_var_8_9
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_7_3(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_9))
}))}))}
goto end_branch_17
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_7_10 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_7_10
var __local_var_7_10 gopurs_runtime.Value
_ = __local_var_7_10
Call_local_Control_Monad_List_Trans___local_var_7_10 = func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1320412129) {
// TAST (Let): __local_var_9_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_2), "map"), f_4)
_ = __local_var_9_11
// TAST (Let): __local_var_10_12 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_12 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_12
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_11, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_12))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 813447293) {
// TAST (Let): __local_var_9_13 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_9_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_2), "map"), f_4)
_ = __local_var_9_13
// TAST (Let): __local_var_10_14 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_14 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_10_14
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_13, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_14))
}))}))}
goto end_branch_15
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 489128924) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}), v_7)
}
__local_var_7_10 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_7_10(v_7_loop_val)
})
// TAST (Let): __local_var_8_16 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_16 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0
_ = __local_var_8_16
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_7_10(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_16))
}))}))}
goto end_branch_17
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 489128924) {
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}), v_5)
})})
_ = functorListT1_3_1
// TAST (Let): __local_var_4_18 shape=App(Var) bindingType=Any
__local_var_4_18 := Call_Control_Monad_List_Trans_monadListT(dictMonad_0)
_ = __local_var_4_18
// TAST (Let): Bind1_5_19 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_19
// TAST (Let): Applicative0_6_20 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_20
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_3_1)}
}), gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_19.V1), f_7, gopurs_runtime.Func(func(f_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_19.V1), a_8, gopurs_runtime.Func(func(a_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_20.V1), gopurs_runtime.Apply(f_prime__9, a_prime__10))
}))
}))
})}))}
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))}))}
}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_2_21 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_2_21 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_2_21
// TAST (Let): Functor0_3_22 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_22
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_24 shape=App(Other) bindingType=Any
__local_var_5_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_24
// TAST (Let): functorListT1_5_23 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_5_23 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_24, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t39 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_9_25 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_9_25
var __local_var_9_25 gopurs_runtime.Value
_ = __local_var_9_25
Call_local_Control_Monad_List_Trans___local_var_9_25 = func(v_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_24, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1320412129) {
// TAST (Let): __local_var_11_26 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_5_24), "map"), f_6)
_ = __local_var_11_26
// TAST (Let): __local_var_12_27 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_27 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_12_27
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_26, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_27))
}))}))}
goto end_branch_30
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 813447293) {
// TAST (Let): __local_var_11_28 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_5_24), "map"), f_6)
_ = __local_var_11_28
// TAST (Let): __local_var_12_29 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_29 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0
_ = __local_var_12_29
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_28, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_29))
}))}))}
goto end_branch_30
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 489128924) {
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
}), v_9)
}
__local_var_9_25 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_9_25(v_9_loop_val)
})
// TAST (Let): __local_var_10_31 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_31 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_31
__t39 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_9_25(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_31))
}))}))}
goto end_branch_39
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_9_32 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_9_32
var __local_var_9_32 gopurs_runtime.Value
_ = __local_var_9_32
Call_local_Control_Monad_List_Trans___local_var_9_32 = func(v_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_24, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1320412129) {
// TAST (Let): __local_var_11_33 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_5_24), "map"), f_6)
_ = __local_var_11_33
// TAST (Let): __local_var_12_34 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_34 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_12_34
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_33, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_34))
}))}))}
goto end_branch_37
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 813447293) {
// TAST (Let): __local_var_11_35 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_11_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_5_24), "map"), f_6)
_ = __local_var_11_35
// TAST (Let): __local_var_12_36 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_12_36 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0
_ = __local_var_12_36
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_35, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_12_36))
}))}))}
goto end_branch_37
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 489128924) {
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return __t37
}), v_9)
}
__local_var_9_32 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_9_32(v_9_loop_val)
})
// TAST (Let): __local_var_10_38 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_38 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_10_38
__t39 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_9_32(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_38))
}))}))}
goto end_branch_39
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 489128924) {
__t39 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_39
} else {

}
}
{
__t39 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_39:
return __t39
}), v_7)
})})
_ = functorListT1_5_23
// TAST (Let): __local_var_6_40 shape=App(Var) bindingType=Any
__local_var_6_40 := Call_Control_Monad_List_Trans_monadListT(dictMonad_0)
_ = __local_var_6_40
// TAST (Let): Bind1_7_41 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_40, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_41
// TAST (Let): Applicative0_8_42 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_40, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_42
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_5_23)}
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_41.V1), f_9, gopurs_runtime.Func(func(f_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_41.V1), a_10, gopurs_runtime.Func(func(a_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_42.V1), gopurs_runtime.Apply(f_prime__11, a_prime__12))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(fa_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_22.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1320412129) {
// TAST (Let): __local_var_7_43 shape=Other bindingType=Any
__local_var_7_43 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0
_ = __local_var_7_43
// TAST (Let): __local_var_8_44 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_44 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
_ = __local_var_8_44
__t46 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_2_21.V0), gopurs_runtime.Apply(f_5, __local_var_7_43), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_44), f_5))
}))}))}
goto end_branch_46
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 813447293) {
// TAST (Let): __local_var_7_45 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_45 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0
_ = __local_var_7_45
__t46 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_45), f_5)
}))}))}
goto end_branch_46
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 489128924) {
__t46 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_46
} else {

}
}
{
__t46 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_46:
return __t46
}), fa_4)
})}))}
})})
_ = monadListT1_1_0
// TAST (Let): applicativeListT1_2_48 shape=LitRecord bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
applicativeListT1_2_48 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_50 shape=App(Other) bindingType=Any
__local_var_3_50 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_50
// TAST (Let): functorListT1_3_49 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_3_49 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_50, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t83 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1320412129) {
// TAST (Let): __local_var_7_51 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_51 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
_ = __local_var_7_51
__t83 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_50, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t66 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_10_52 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_10_52
var __local_var_10_52 gopurs_runtime.Value
_ = __local_var_10_52
Call_local_Control_Monad_List_Trans___local_var_10_52 = func(v_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_50, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t57 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1320412129) {
// TAST (Let): __local_var_12_53 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_53 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_50), "map"), f_4)
_ = __local_var_12_53
// TAST (Let): __local_var_13_54 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_54 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1
_ = __local_var_13_54
__t57 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_53, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_54))
}))}))}
goto end_branch_57
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 813447293) {
// TAST (Let): __local_var_12_55 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_55 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_50), "map"), f_4)
_ = __local_var_12_55
// TAST (Let): __local_var_13_56 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_56 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0
_ = __local_var_13_56
__t57 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_55, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_56))
}))}))}
goto end_branch_57
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 489128924) {
__t57 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_57
} else {

}
}
{
__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_57:
return __t57
}), v_10)
}
__local_var_10_52 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_52(v_10_loop_val)
})
// TAST (Let): __local_var_11_58 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_58 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_11_58
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_52(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_58))
}))}))}
goto end_branch_66
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_10_59 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_10_59
var __local_var_10_59 gopurs_runtime.Value
_ = __local_var_10_59
Call_local_Control_Monad_List_Trans___local_var_10_59 = func(v_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_50, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t64 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1320412129) {
// TAST (Let): __local_var_12_60 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_60 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_50), "map"), f_4)
_ = __local_var_12_60
// TAST (Let): __local_var_13_61 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_61 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1
_ = __local_var_13_61
__t64 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_60, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_61))
}))}))}
goto end_branch_64
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 813447293) {
// TAST (Let): __local_var_12_62 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_62 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_50), "map"), f_4)
_ = __local_var_12_62
// TAST (Let): __local_var_13_63 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_63 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0
_ = __local_var_13_63
__t64 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_62, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_63))
}))}))}
goto end_branch_64
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 489128924) {
__t64 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_64
} else {

}
}
{
__t64 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_64:
return __t64
}), v_10)
}
__local_var_10_59 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_59(v_10_loop_val)
})
// TAST (Let): __local_var_11_65 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_65 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_11_65
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_59(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_65))
}))}))}
goto end_branch_66
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 489128924) {
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_66
} else {

}
}
{
__t66 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_66:
return __t66
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_51))
}))}))}
goto end_branch_83
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 813447293) {
// TAST (Let): __local_var_7_67 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_67 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0
_ = __local_var_7_67
__t83 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_50, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t82 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_10_68 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_10_68
var __local_var_10_68 gopurs_runtime.Value
_ = __local_var_10_68
Call_local_Control_Monad_List_Trans___local_var_10_68 = func(v_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_50, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t73 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1320412129) {
// TAST (Let): __local_var_12_69 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_69 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_50), "map"), f_4)
_ = __local_var_12_69
// TAST (Let): __local_var_13_70 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_70 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1
_ = __local_var_13_70
__t73 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_69, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_70))
}))}))}
goto end_branch_73
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 813447293) {
// TAST (Let): __local_var_12_71 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_71 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_50), "map"), f_4)
_ = __local_var_12_71
// TAST (Let): __local_var_13_72 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_72 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0
_ = __local_var_13_72
__t73 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_71, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_72))
}))}))}
goto end_branch_73
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 489128924) {
__t73 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_73
} else {

}
}
{
__t73 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_73:
return __t73
}), v_10)
}
__local_var_10_68 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_68(v_10_loop_val)
})
// TAST (Let): __local_var_11_74 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_74 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_11_74
__t82 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_68(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_74))
}))}))}
goto end_branch_82
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_10_75 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_10_75
var __local_var_10_75 gopurs_runtime.Value
_ = __local_var_10_75
Call_local_Control_Monad_List_Trans___local_var_10_75 = func(v_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_50, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t80 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1320412129) {
// TAST (Let): __local_var_12_76 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_76 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_50), "map"), f_4)
_ = __local_var_12_76
// TAST (Let): __local_var_13_77 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_77 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1
_ = __local_var_13_77
__t80 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_76, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_77))
}))}))}
goto end_branch_80
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 813447293) {
// TAST (Let): __local_var_12_78 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_78 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_3_50), "map"), f_4)
_ = __local_var_12_78
// TAST (Let): __local_var_13_79 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_79 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0
_ = __local_var_13_79
__t80 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_78, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_79))
}))}))}
goto end_branch_80
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 489128924) {
__t80 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_80
} else {

}
}
{
__t80 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_80:
return __t80
}), v_10)
}
__local_var_10_75 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_75(v_10_loop_val)
})
// TAST (Let): __local_var_11_81 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_81 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_11_81
__t82 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_75(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_81))
}))}))}
goto end_branch_82
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 489128924) {
__t82 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_82
} else {

}
}
{
__t82 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_82:
return __t82
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_67))
}))}))}
goto end_branch_83
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 489128924) {
__t83 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_83
} else {

}
}
{
__t83 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_83:
return __t83
}), v_5)
})})
_ = functorListT1_3_49
// TAST (Let): __local_var_4_84 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_4_84 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_86 shape=App(Other) bindingType=Any
__local_var_6_86 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_86
// TAST (Let): functorListT1_6_85 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_6_85 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_86, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t101 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_10_87 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_10_87
var __local_var_10_87 gopurs_runtime.Value
_ = __local_var_10_87
Call_local_Control_Monad_List_Trans___local_var_10_87 = func(v_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_86, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t92 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1320412129) {
// TAST (Let): __local_var_12_88 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_88 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_86), "map"), f_7)
_ = __local_var_12_88
// TAST (Let): __local_var_13_89 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_89 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1
_ = __local_var_13_89
__t92 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_88, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_89))
}))}))}
goto end_branch_92
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 813447293) {
// TAST (Let): __local_var_12_90 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_90 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_86), "map"), f_7)
_ = __local_var_12_90
// TAST (Let): __local_var_13_91 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_91 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0
_ = __local_var_13_91
__t92 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_90, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_91))
}))}))}
goto end_branch_92
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 489128924) {
__t92 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_92
} else {

}
}
{
__t92 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_92:
return __t92
}), v_10)
}
__local_var_10_87 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_87(v_10_loop_val)
})
// TAST (Let): __local_var_11_93 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_93 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_11_93
__t101 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_87(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_93))
}))}))}
goto end_branch_101
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_10_94 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_10_94
var __local_var_10_94 gopurs_runtime.Value
_ = __local_var_10_94
Call_local_Control_Monad_List_Trans___local_var_10_94 = func(v_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_86, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t99 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1320412129) {
// TAST (Let): __local_var_12_95 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_95 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_86), "map"), f_7)
_ = __local_var_12_95
// TAST (Let): __local_var_13_96 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_96 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1
_ = __local_var_13_96
__t99 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_95, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_96))
}))}))}
goto end_branch_99
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 813447293) {
// TAST (Let): __local_var_12_97 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_12_97 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_6_86), "map"), f_7)
_ = __local_var_12_97
// TAST (Let): __local_var_13_98 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_98 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0
_ = __local_var_13_98
__t99 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_97, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_98))
}))}))}
goto end_branch_99
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 489128924) {
__t99 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_99
} else {

}
}
{
__t99 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_99:
return __t99
}), v_10)
}
__local_var_10_94 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_94(v_10_loop_val)
})
// TAST (Let): __local_var_11_100 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_100 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_11_100
__t101 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_10_94(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_100))
}))}))}
goto end_branch_101
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 489128924) {
__t101 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_101
} else {

}
}
{
__t101 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_101:
return __t101
}), v_8)
})})
_ = functorListT1_6_85
// TAST (Let): __local_var_7_102 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_7_102 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_8_103 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_8_103 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_8_103
// TAST (Let): Functor0_9_104 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_9_104 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_104
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applyListT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(fa_10 gopurs_runtime.Value, f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_104.V0), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t108 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1320412129) {
// TAST (Let): __local_var_13_105 shape=Other bindingType=Any
__local_var_13_105 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_13_105
// TAST (Let): __local_var_14_106 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_14_106 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1
_ = __local_var_14_106
__t108 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_8_103.V0), gopurs_runtime.Apply(f_11, __local_var_13_105), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_106), f_11))
}))}))}
goto end_branch_108
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 813447293) {
// TAST (Let): __local_var_13_107 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_107 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0
_ = __local_var_13_107
__t108 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_107), f_11)
}))}))}
goto end_branch_108
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 489128924) {
__t108 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_108
} else {

}
}
{
__t108 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_108:
return __t108
}), fa_10)
})}))}
})})
_ = __local_var_7_102
// TAST (Let): Bind1_8_109 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_8_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_7_102.V1), gopurs_runtime.Value{}))
_ = Bind1_8_109
// TAST (Let): Applicative0_9_110 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_9_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_7_102.V0), gopurs_runtime.Value{}))
_ = Applicative0_9_110
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_6_85)}
}), gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_109.V1), f_10, gopurs_runtime.Func(func(f_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_109.V1), a_11, gopurs_runtime.Func(func(a_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_110.V1), gopurs_runtime.Apply(f_prime__12, a_prime__13))
}))
}))
})}))}
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))}))}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_5_111 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_5_111 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_5_111
// TAST (Let): Functor0_6_112 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_6_112 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_112
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_114 shape=App(Other) bindingType=Any
__local_var_8_114 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_114
// TAST (Let): functorListT1_8_113 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_8_113 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_114, "map"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t129 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_12_115 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_12_115
var __local_var_12_115 gopurs_runtime.Value
_ = __local_var_12_115
Call_local_Control_Monad_List_Trans___local_var_12_115 = func(v_12_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_12 gopurs_runtime.Value = v_12_loop
_ = v_12
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_114, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t120 gopurs_runtime.Value
{
if (v_13.Type == 9 && v_13.IntVal == 1320412129) {
// TAST (Let): __local_var_14_116 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_14_116 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_8_114), "map"), f_9)
_ = __local_var_14_116
// TAST (Let): __local_var_15_117 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_117 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V1
_ = __local_var_15_117
__t120 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_116, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_117))
}))}))}
goto end_branch_120
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 813447293) {
// TAST (Let): __local_var_14_118 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_14_118 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_8_114), "map"), f_9)
_ = __local_var_14_118
// TAST (Let): __local_var_15_119 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_119 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0
_ = __local_var_15_119
__t120 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_118, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_119))
}))}))}
goto end_branch_120
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 489128924) {
__t120 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_120
} else {

}
}
{
__t120 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_120:
return __t120
}), v_12)
}
__local_var_12_115 = gopurs_runtime.Func(func(v_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_12_115(v_12_loop_val)
})
// TAST (Let): __local_var_13_121 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_121 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1
_ = __local_var_13_121
__t129 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_12_115(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_121))
}))}))}
goto end_branch_129
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_12_122 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_12_122
var __local_var_12_122 gopurs_runtime.Value
_ = __local_var_12_122
Call_local_Control_Monad_List_Trans___local_var_12_122 = func(v_12_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_12 gopurs_runtime.Value = v_12_loop
_ = v_12
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_114, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t127 gopurs_runtime.Value
{
if (v_13.Type == 9 && v_13.IntVal == 1320412129) {
// TAST (Let): __local_var_14_123 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_14_123 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_8_114), "map"), f_9)
_ = __local_var_14_123
// TAST (Let): __local_var_15_124 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_124 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V1
_ = __local_var_15_124
__t127 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_123, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_124))
}))}))}
goto end_branch_127
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 813447293) {
// TAST (Let): __local_var_14_125 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_14_125 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_8_114), "map"), f_9)
_ = __local_var_14_125
// TAST (Let): __local_var_15_126 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_126 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0
_ = __local_var_15_126
__t127 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_125, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_126))
}))}))}
goto end_branch_127
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 489128924) {
__t127 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_127
} else {

}
}
{
__t127 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_127:
return __t127
}), v_12)
}
__local_var_12_122 = gopurs_runtime.Func(func(v_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_12_122(v_12_loop_val)
})
// TAST (Let): __local_var_13_128 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_13_128 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0
_ = __local_var_13_128
__t129 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_12_122(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_13_128))
}))}))}
goto end_branch_129
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 489128924) {
__t129 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_129
} else {

}
}
{
__t129 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_129:
return __t129
}), v_10)
})})
_ = functorListT1_8_113
// TAST (Let): __local_var_9_130 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_9_130 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applyListT(dictMonad_0)))}
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))}))}
}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_10_131 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_10_131 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_10_131
// TAST (Let): Functor0_11_132 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_11_132 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_11_132
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_applyListT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(fa_12 gopurs_runtime.Value, f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_132.V0), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t136 gopurs_runtime.Value
{
if (v_14.Type == 9 && v_14.IntVal == 1320412129) {
// TAST (Let): __local_var_15_133 shape=Other bindingType=Any
__local_var_15_133 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V0
_ = __local_var_15_133
// TAST (Let): __local_var_16_134 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_16_134 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V1
_ = __local_var_16_134
__t136 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_10_131.V0), gopurs_runtime.Apply(f_13, __local_var_15_133), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_16_134), f_13))
}))}))}
goto end_branch_136
} else {

}
}
{
if (v_14.Type == 9 && v_14.IntVal == 813447293) {
// TAST (Let): __local_var_15_135 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_135 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V0
_ = __local_var_15_135
__t136 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_135), f_13)
}))}))}
goto end_branch_136
} else {

}
}
{
if (v_14.Type == 9 && v_14.IntVal == 489128924) {
__t136 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_136
} else {

}
}
{
__t136 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_136:
return __t136
}), fa_12)
})}))}
})})
_ = __local_var_9_130
// TAST (Let): Bind1_10_137 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_10_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_9_130.V1), gopurs_runtime.Value{}))
_ = Bind1_10_137
// TAST (Let): Applicative0_11_138 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_11_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_9_130.V0), gopurs_runtime.Value{}))
_ = Applicative0_11_138
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_8_113)}
}), gopurs_runtime.Func2(func(f_12 gopurs_runtime.Value, a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_137.V1), f_12, gopurs_runtime.Func(func(f_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_137.V1), a_13, gopurs_runtime.Func(func(a_prime__15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_138.V1), gopurs_runtime.Apply(f_prime__14, a_prime__15))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(fa_7 gopurs_runtime.Value, f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_112.V0), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t156 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1320412129) {
// TAST (Let): __local_var_10_139 shape=Other bindingType=Any
__local_var_10_139 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_10_139
// TAST (Let): __local_var_11_140 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_140 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_11_140
__t156 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_13_141 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_13_141 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_13_141
// TAST (Let): Functor0_14_142 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_14_142 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_14_142
var Call_local_Control_Monad_List_Trans___local_var_15_143 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_15_143
var __local_var_15_143 gopurs_runtime.Value
_ = __local_var_15_143
Call_local_Control_Monad_List_Trans___local_var_15_143 = func(fa_15_loop gopurs_runtime.Value, f_16_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_15 gopurs_runtime.Value = fa_15_loop
_ = fa_15
var f_16 gopurs_runtime.Value = f_16_loop
_ = f_16
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_14_142.V0), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t147 gopurs_runtime.Value
{
if (v_17.Type == 9 && v_17.IntVal == 1320412129) {
// TAST (Let): __local_var_18_144 shape=Other bindingType=Any
__local_var_18_144 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_17.UnsafePtr).V0
_ = __local_var_18_144
// TAST (Let): __local_var_19_145 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_19_145 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_17.UnsafePtr).V1
_ = __local_var_19_145
__t147 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_13_141.V0), gopurs_runtime.Apply(f_16, __local_var_18_144), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_19_145), f_16))
}))}))}
goto end_branch_147
} else {

}
}
{
if (v_17.Type == 9 && v_17.IntVal == 813447293) {
// TAST (Let): __local_var_18_146 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_18_146 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_17.UnsafePtr).V0
_ = __local_var_18_146
__t147 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_146), f_16)
}))}))}
goto end_branch_147
} else {

}
}
{
if (v_17.Type == 9 && v_17.IntVal == 489128924) {
__t147 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_147
} else {

}
}
{
__t147 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_147:
return __t147
}), fa_15)
}
__local_var_15_143 = gopurs_runtime.Func(func(fa_15_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_15_143(fa_15_loop_val, f_16_loop_val)
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_5_111.V0), gopurs_runtime.Apply(f_8, __local_var_10_139), Call_local_Control_Monad_List_Trans___local_var_15_143(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_140), f_8))
}))}))}
goto end_branch_156
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 813447293) {
// TAST (Let): __local_var_10_148 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_10_148 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
_ = __local_var_10_148
__t156 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupListT1_12_149 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])])])
semigroupListT1_12_149 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = semigroupListT1_12_149
// TAST (Let): Functor0_13_150 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_13_150 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_13_150
var Call_local_Control_Monad_List_Trans___local_var_14_151 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_151
var __local_var_14_151 gopurs_runtime.Value
_ = __local_var_14_151
Call_local_Control_Monad_List_Trans___local_var_14_151 = func(fa_14_loop gopurs_runtime.Value, f_15_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_14 gopurs_runtime.Value = fa_14_loop
_ = fa_14
var f_15 gopurs_runtime.Value = f_15_loop
_ = f_15
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_13_150.V0), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t155 gopurs_runtime.Value
{
if (v_16.Type == 9 && v_16.IntVal == 1320412129) {
// TAST (Let): __local_var_17_152 shape=Other bindingType=Any
__local_var_17_152 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_16.UnsafePtr).V0
_ = __local_var_17_152
// TAST (Let): __local_var_18_153 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_18_153 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_16.UnsafePtr).V1
_ = __local_var_18_153
__t155 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_12_149.V0), gopurs_runtime.Apply(f_15, __local_var_17_152), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_153), f_15))
}))}))}
goto end_branch_155
} else {

}
}
{
if (v_16.Type == 9 && v_16.IntVal == 813447293) {
// TAST (Let): __local_var_17_154 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_154 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_16.UnsafePtr).V0
_ = __local_var_17_154
__t155 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_154), f_15)
}))}))}
goto end_branch_155
} else {

}
}
{
if (v_16.Type == 9 && v_16.IntVal == 489128924) {
__t155 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_155
} else {

}
}
{
__t155 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_155:
return __t155
}), fa_14)
}
__local_var_14_151 = gopurs_runtime.Func(func(fa_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_151(fa_14_loop_val, f_15_loop_val)
})
})
return Call_local_Control_Monad_List_Trans___local_var_14_151(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_148), f_8)
}))}))}
goto end_branch_156
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 489128924) {
__t156 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_156
} else {

}
}
{
__t156 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_156:
return __t156
}), fa_7)
})}))}
})})
_ = __local_var_4_84
// TAST (Let): Bind1_5_157 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_157 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_4_84.V1), gopurs_runtime.Value{}))
_ = Bind1_5_157
// TAST (Let): Applicative0_6_158 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_6_158 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_4_84.V0), gopurs_runtime.Value{}))
_ = Applicative0_6_158
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_3_49)}
}), gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_157.V1), f_7, gopurs_runtime.Func(func(f_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_157.V1), a_8, gopurs_runtime.Func(func(a_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_158.V1), gopurs_runtime.Apply(f_prime__9, a_prime__10))
}))
}))
})}))}
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))})
_ = applicativeListT1_2_48
// TAST (Let): Applicative0_3_160 shape=App(Other) bindingType=Any
Applicative0_3_160 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_3_160
// TAST (Let): __local_var_4_163 shape=App(Other) bindingType=Any
__local_var_4_163 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_3_160, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_163
// TAST (Let): functorListT1_4_162 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
functorListT1_4_162 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t232 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1320412129) {
// TAST (Let): __local_var_8_164 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_164 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_8_164
__t232 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t197 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1320412129) {
// TAST (Let): __local_var_11_165 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_165 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_11_165
__t197 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t180 gopurs_runtime.Value
{
if (v_13.Type == 9 && v_13.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_14_166 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_166
var __local_var_14_166 gopurs_runtime.Value
_ = __local_var_14_166
Call_local_Control_Monad_List_Trans___local_var_14_166 = func(v_14_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_14 gopurs_runtime.Value = v_14_loop
_ = v_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t171 gopurs_runtime.Value
{
if (v_15.Type == 9 && v_15.IntVal == 1320412129) {
// TAST (Let): __local_var_16_167 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_167 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_167
// TAST (Let): __local_var_17_168 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_168 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V1
_ = __local_var_17_168
__t171 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_167, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_168))
}))}))}
goto end_branch_171
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 813447293) {
// TAST (Let): __local_var_16_169 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_169 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_169
// TAST (Let): __local_var_17_170 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_170 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0
_ = __local_var_17_170
__t171 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_169, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_170))
}))}))}
goto end_branch_171
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 489128924) {
__t171 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_171
} else {

}
}
{
__t171 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_171:
return __t171
}), v_14)
}
__local_var_14_166 = gopurs_runtime.Func(func(v_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_166(v_14_loop_val)
})
// TAST (Let): __local_var_15_172 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_172 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V1
_ = __local_var_15_172
__t180 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_166(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_172))
}))}))}
goto end_branch_180
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_14_173 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_173
var __local_var_14_173 gopurs_runtime.Value
_ = __local_var_14_173
Call_local_Control_Monad_List_Trans___local_var_14_173 = func(v_14_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_14 gopurs_runtime.Value = v_14_loop
_ = v_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t178 gopurs_runtime.Value
{
if (v_15.Type == 9 && v_15.IntVal == 1320412129) {
// TAST (Let): __local_var_16_174 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_174 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_174
// TAST (Let): __local_var_17_175 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_175 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V1
_ = __local_var_17_175
__t178 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_174, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_175))
}))}))}
goto end_branch_178
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 813447293) {
// TAST (Let): __local_var_16_176 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_176 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_176
// TAST (Let): __local_var_17_177 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_177 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0
_ = __local_var_17_177
__t178 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_176, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_177))
}))}))}
goto end_branch_178
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 489128924) {
__t178 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_178
} else {

}
}
{
__t178 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_178:
return __t178
}), v_14)
}
__local_var_14_173 = gopurs_runtime.Func(func(v_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_173(v_14_loop_val)
})
// TAST (Let): __local_var_15_179 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_179 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0
_ = __local_var_15_179
__t180 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_173(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_179))
}))}))}
goto end_branch_180
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 489128924) {
__t180 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_180
} else {

}
}
{
__t180 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_180:
return __t180
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_165))
}))}))}
goto end_branch_197
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 813447293) {
// TAST (Let): __local_var_11_181 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_181 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0
_ = __local_var_11_181
__t197 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t196 gopurs_runtime.Value
{
if (v_13.Type == 9 && v_13.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_14_182 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_182
var __local_var_14_182 gopurs_runtime.Value
_ = __local_var_14_182
Call_local_Control_Monad_List_Trans___local_var_14_182 = func(v_14_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_14 gopurs_runtime.Value = v_14_loop
_ = v_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t187 gopurs_runtime.Value
{
if (v_15.Type == 9 && v_15.IntVal == 1320412129) {
// TAST (Let): __local_var_16_183 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_183 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_183
// TAST (Let): __local_var_17_184 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_184 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V1
_ = __local_var_17_184
__t187 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_183, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_184))
}))}))}
goto end_branch_187
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 813447293) {
// TAST (Let): __local_var_16_185 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_185 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_185
// TAST (Let): __local_var_17_186 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_186 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0
_ = __local_var_17_186
__t187 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_185, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_186))
}))}))}
goto end_branch_187
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 489128924) {
__t187 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_187
} else {

}
}
{
__t187 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_187:
return __t187
}), v_14)
}
__local_var_14_182 = gopurs_runtime.Func(func(v_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_182(v_14_loop_val)
})
// TAST (Let): __local_var_15_188 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_188 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V1
_ = __local_var_15_188
__t196 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_182(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_188))
}))}))}
goto end_branch_196
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_14_189 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_189
var __local_var_14_189 gopurs_runtime.Value
_ = __local_var_14_189
Call_local_Control_Monad_List_Trans___local_var_14_189 = func(v_14_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_14 gopurs_runtime.Value = v_14_loop
_ = v_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t194 gopurs_runtime.Value
{
if (v_15.Type == 9 && v_15.IntVal == 1320412129) {
// TAST (Let): __local_var_16_190 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_190 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_190
// TAST (Let): __local_var_17_191 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_191 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V1
_ = __local_var_17_191
__t194 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_190, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_191))
}))}))}
goto end_branch_194
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 813447293) {
// TAST (Let): __local_var_16_192 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_192 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_192
// TAST (Let): __local_var_17_193 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_193 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0
_ = __local_var_17_193
__t194 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_192, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_193))
}))}))}
goto end_branch_194
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 489128924) {
__t194 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_194
} else {

}
}
{
__t194 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_194:
return __t194
}), v_14)
}
__local_var_14_189 = gopurs_runtime.Func(func(v_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_189(v_14_loop_val)
})
// TAST (Let): __local_var_15_195 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_195 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0
_ = __local_var_15_195
__t196 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_189(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_195))
}))}))}
goto end_branch_196
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 489128924) {
__t196 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_196
} else {

}
}
{
__t196 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_196:
return __t196
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_181))
}))}))}
goto end_branch_197
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 489128924) {
__t197 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_197
} else {

}
}
{
__t197 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_197:
return __t197
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_164))
}))}))}
goto end_branch_232
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 813447293) {
// TAST (Let): __local_var_8_198 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_8_198 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_198
__t232 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t231 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1320412129) {
// TAST (Let): __local_var_11_199 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_199 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_11_199
__t231 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t214 gopurs_runtime.Value
{
if (v_13.Type == 9 && v_13.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_14_200 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_200
var __local_var_14_200 gopurs_runtime.Value
_ = __local_var_14_200
Call_local_Control_Monad_List_Trans___local_var_14_200 = func(v_14_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_14 gopurs_runtime.Value = v_14_loop
_ = v_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t205 gopurs_runtime.Value
{
if (v_15.Type == 9 && v_15.IntVal == 1320412129) {
// TAST (Let): __local_var_16_201 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_201 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_201
// TAST (Let): __local_var_17_202 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_202 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V1
_ = __local_var_17_202
__t205 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_201, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_202))
}))}))}
goto end_branch_205
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 813447293) {
// TAST (Let): __local_var_16_203 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_203 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_203
// TAST (Let): __local_var_17_204 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_204 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0
_ = __local_var_17_204
__t205 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_203, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_204))
}))}))}
goto end_branch_205
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 489128924) {
__t205 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_205
} else {

}
}
{
__t205 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_205:
return __t205
}), v_14)
}
__local_var_14_200 = gopurs_runtime.Func(func(v_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_200(v_14_loop_val)
})
// TAST (Let): __local_var_15_206 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_206 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V1
_ = __local_var_15_206
__t214 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_200(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_206))
}))}))}
goto end_branch_214
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_14_207 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_207
var __local_var_14_207 gopurs_runtime.Value
_ = __local_var_14_207
Call_local_Control_Monad_List_Trans___local_var_14_207 = func(v_14_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_14 gopurs_runtime.Value = v_14_loop
_ = v_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t212 gopurs_runtime.Value
{
if (v_15.Type == 9 && v_15.IntVal == 1320412129) {
// TAST (Let): __local_var_16_208 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_208 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_208
// TAST (Let): __local_var_17_209 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_209 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V1
_ = __local_var_17_209
__t212 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_208, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_209))
}))}))}
goto end_branch_212
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 813447293) {
// TAST (Let): __local_var_16_210 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_210 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_210
// TAST (Let): __local_var_17_211 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_211 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0
_ = __local_var_17_211
__t212 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_210, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_211))
}))}))}
goto end_branch_212
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 489128924) {
__t212 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_212
} else {

}
}
{
__t212 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_212:
return __t212
}), v_14)
}
__local_var_14_207 = gopurs_runtime.Func(func(v_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_207(v_14_loop_val)
})
// TAST (Let): __local_var_15_213 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_213 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0
_ = __local_var_15_213
__t214 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_207(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_213))
}))}))}
goto end_branch_214
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 489128924) {
__t214 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_214
} else {

}
}
{
__t214 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_214:
return __t214
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_199))
}))}))}
goto end_branch_231
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 813447293) {
// TAST (Let): __local_var_11_215 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_11_215 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0
_ = __local_var_11_215
__t231 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t230 gopurs_runtime.Value
{
if (v_13.Type == 9 && v_13.IntVal == 1320412129) {
var Call_local_Control_Monad_List_Trans___local_var_14_216 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_216
var __local_var_14_216 gopurs_runtime.Value
_ = __local_var_14_216
Call_local_Control_Monad_List_Trans___local_var_14_216 = func(v_14_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_14 gopurs_runtime.Value = v_14_loop
_ = v_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t221 gopurs_runtime.Value
{
if (v_15.Type == 9 && v_15.IntVal == 1320412129) {
// TAST (Let): __local_var_16_217 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_217 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_217
// TAST (Let): __local_var_17_218 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_218 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V1
_ = __local_var_17_218
__t221 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_217, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_218))
}))}))}
goto end_branch_221
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 813447293) {
// TAST (Let): __local_var_16_219 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_219 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_219
// TAST (Let): __local_var_17_220 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_220 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0
_ = __local_var_17_220
__t221 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_219, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_220))
}))}))}
goto end_branch_221
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 489128924) {
__t221 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_221
} else {

}
}
{
__t221 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_221:
return __t221
}), v_14)
}
__local_var_14_216 = gopurs_runtime.Func(func(v_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_216(v_14_loop_val)
})
// TAST (Let): __local_var_15_222 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_222 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V1
_ = __local_var_15_222
__t230 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_216(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_222))
}))}))}
goto end_branch_230
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 813447293) {
var Call_local_Control_Monad_List_Trans___local_var_14_223 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_List_Trans___local_var_14_223
var __local_var_14_223 gopurs_runtime.Value
_ = __local_var_14_223
Call_local_Control_Monad_List_Trans___local_var_14_223 = func(v_14_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_14 gopurs_runtime.Value = v_14_loop
_ = v_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_163, "map"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t228 gopurs_runtime.Value
{
if (v_15.Type == 9 && v_15.IntVal == 1320412129) {
// TAST (Let): __local_var_16_224 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_224 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_224
// TAST (Let): __local_var_17_225 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_225 := (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V1
_ = __local_var_17_225
__t228 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_List_Trans_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_224, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_225))
}))}))}
goto end_branch_228
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 813447293) {
// TAST (Let): __local_var_16_226 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar b), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar b)])])]))
__local_var_16_226 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(__local_var_4_163), "map"), f_5)
_ = __local_var_16_226
// TAST (Let): __local_var_17_227 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_17_227 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_15.UnsafePtr).V0
_ = __local_var_17_227
__t228 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_226, gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_227))
}))}))}
goto end_branch_228
} else {

}
}
{
if (v_15.Type == 9 && v_15.IntVal == 489128924) {
__t228 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_228
} else {

}
}
{
__t228 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_228:
return __t228
}), v_14)
}
__local_var_14_223 = gopurs_runtime.Func(func(v_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_223(v_14_loop_val)
})
// TAST (Let): __local_var_15_229 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
__local_var_15_229 := (*Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0
_ = __local_var_15_229
__t230 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_List_Trans_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_List_Trans___local_var_14_223(gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_229))
}))}))}
goto end_branch_230
} else {

}
}
{
if (v_13.Type == 9 && v_13.IntVal == 489128924) {
__t230 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_230
} else {

}
}
{
__t230 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_230:
return __t230
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_215))
}))}))}
goto end_branch_231
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 489128924) {
__t231 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_231
} else {

}
}
{
__t231 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_231:
return __t231
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_198))
}))}))}
goto end_branch_232
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 489128924) {
__t232 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_232
} else {

}
}
{
__t232 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_232:
return __t232
}), v_6)
})})
_ = functorListT1_4_162
// TAST (Let): altListT1_4_161 shape=Let(LitRecord) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
altListT1_4_161 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorListT1_4_162)}
}), Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_3_160))})
_ = altListT1_4_161
// TAST (Let): plusListT1_3_159 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
plusListT1_3_159 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altListT1_4_161)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_3_160, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})})
_ = plusListT1_3_159
// TAST (Let): alternativeListT1_2_47 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alternative","Alternative"] [(TypeApp (TypeVar f) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar f), (TypeVar a)])])])])
alternativeListT1_2_47 := (&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeListT1_2_48)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusListT1_3_159)}
})})
_ = alternativeListT1_2_47
return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeListT1_2_47)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadListT1_1_0)}
})}))}
}

func Rebox_Control_Monad_List_Trans_1239976062_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Control_Monad_List_Trans_138441832_1785332133(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_List_Trans_138441832_3804580809(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Control_Monad_List_Trans_1785332133_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_List_Trans_2549197956_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_List_Trans_1785332133_138441832(in.V0))}
	return out
}

func Rebox_Control_Monad_List_Trans_3094389156_2549197956(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = Rebox_Control_Monad_List_Trans_138441832_1785332133(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Control_Monad_List_Trans_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Control_Monad_List_Trans_3804580809_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Control_Monad_List_Trans_4010058633_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_List_Trans_784458114_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
				}()
	return out
}


