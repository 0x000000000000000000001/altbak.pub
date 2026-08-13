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
		cache_Control_Monad_List_Trans_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_identity(x_0_box)
})
	})
	return cache_Control_Monad_List_Trans_identity
}

var cache_Control_Monad_List_Trans_identity1 gopurs_runtime.Value
var once_Control_Monad_List_Trans_identity1 sync.Once
func Get_Control_Monad_List_Trans_identity1() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_identity1.Do(func() {
		cache_Control_Monad_List_Trans_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Monad_List_Trans_identity1(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](x_0_box)))}
})
	})
	return cache_Control_Monad_List_Trans_identity1
}

var cache_Control_Monad_List_Trans_Yield gopurs_runtime.Value
var once_Control_Monad_List_Trans_Yield sync.Once
func Get_Control_Monad_List_Trans_Yield() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_Yield.Do(func() {
		cache_Control_Monad_List_Trans_Yield = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, value0})}
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
return Call_Control_Monad_List_Trans_wrapLazy(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), v_1_box)
})
	})
	return cache_Control_Monad_List_Trans_wrapLazy
}

var cache_Control_Monad_List_Trans_wrapEffect gopurs_runtime.Value
var once_Control_Monad_List_Trans_wrapEffect sync.Once
func Get_Control_Monad_List_Trans_wrapEffect() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_wrapEffect.Do(func() {
		cache_Control_Monad_List_Trans_wrapEffect = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_wrapEffect(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), v_1_box)
})
	})
	return cache_Control_Monad_List_Trans_wrapEffect
}

var cache_Control_Monad_List_Trans_unfold gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfold sync.Once
func Get_Control_Monad_List_Trans_unfold() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfold.Do(func() {
		cache_Control_Monad_List_Trans_unfold = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_unfold(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_unfold
}

var cache_Control_Monad_List_Trans_uncons gopurs_runtime.Value
var once_Control_Monad_List_Trans_uncons sync.Once
func Get_Control_Monad_List_Trans_uncons() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_uncons.Do(func() {
		cache_Control_Monad_List_Trans_uncons = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_uncons(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_uncons
}

var cache_Control_Monad_List_Trans_tail gopurs_runtime.Value
var once_Control_Monad_List_Trans_tail sync.Once
func Get_Control_Monad_List_Trans_tail() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_tail.Do(func() {
		cache_Control_Monad_List_Trans_tail = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_tail(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_tail
}

var cache_Control_Monad_List_Trans_stepMap gopurs_runtime.Value
var once_Control_Monad_List_Trans_stepMap sync.Once
func Get_Control_Monad_List_Trans_stepMap() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_stepMap.Do(func() {
		cache_Control_Monad_List_Trans_stepMap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_stepMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_stepMap
}

var cache_Control_Monad_List_Trans_takeWhile gopurs_runtime.Value
var once_Control_Monad_List_Trans_takeWhile sync.Once
func Get_Control_Monad_List_Trans_takeWhile() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_takeWhile.Do(func() {
		cache_Control_Monad_List_Trans_takeWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_takeWhile(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_takeWhile
}

var cache_Control_Monad_List_Trans_scanl gopurs_runtime.Value
var once_Control_Monad_List_Trans_scanl sync.Once
func Get_Control_Monad_List_Trans_scanl() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_scanl.Do(func() {
		cache_Control_Monad_List_Trans_scanl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_scanl(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_scanl
}

var cache_Control_Monad_List_Trans_prepend_prime gopurs_runtime.Value
var once_Control_Monad_List_Trans_prepend_prime sync.Once
func Get_Control_Monad_List_Trans_prepend_prime() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_prepend_prime.Do(func() {
		cache_Control_Monad_List_Trans_prepend_prime = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_prepend_prime(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_Control_Monad_List_Trans_prepend_prime
}

var cache_Control_Monad_List_Trans_prepend gopurs_runtime.Value
var once_Control_Monad_List_Trans_prepend sync.Once
func Get_Control_Monad_List_Trans_prepend() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_prepend.Do(func() {
		cache_Control_Monad_List_Trans_prepend = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_prepend(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), h_1_box, t_2_box)
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
return Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_singleton
}

var cache_Control_Monad_List_Trans_take gopurs_runtime.Value
var once_Control_Monad_List_Trans_take sync.Once
func Get_Control_Monad_List_Trans_take() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_take.Do(func() {
		cache_Control_Monad_List_Trans_take = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_take(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_take
}

var cache_Control_Monad_List_Trans_zipWith_prime gopurs_runtime.Value
var once_Control_Monad_List_Trans_zipWith_prime sync.Once
func Get_Control_Monad_List_Trans_zipWith_prime() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_zipWith_prime.Do(func() {
		cache_Control_Monad_List_Trans_zipWith_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_zipWith_prime(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_zipWith_prime
}

var cache_Control_Monad_List_Trans_zipWith gopurs_runtime.Value
var once_Control_Monad_List_Trans_zipWith sync.Once
func Get_Control_Monad_List_Trans_zipWith() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_zipWith.Do(func() {
		cache_Control_Monad_List_Trans_zipWith = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_zipWith(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_zipWith
}

var cache_Control_Monad_List_Trans_newtypeListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_newtypeListT sync.Once
func Get_Control_Monad_List_Trans_newtypeListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_newtypeListT.Do(func() {
		cache_Control_Monad_List_Trans_newtypeListT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Monad_List_Trans_newtypeListT
}

var cache_Control_Monad_List_Trans_mapMaybe gopurs_runtime.Value
var once_Control_Monad_List_Trans_mapMaybe sync.Once
func Get_Control_Monad_List_Trans_mapMaybe() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_mapMaybe.Do(func() {
		cache_Control_Monad_List_Trans_mapMaybe = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_mapMaybe
}

var cache_Control_Monad_List_Trans_iterate gopurs_runtime.Value
var once_Control_Monad_List_Trans_iterate sync.Once
func Get_Control_Monad_List_Trans_iterate() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_iterate.Do(func() {
		cache_Control_Monad_List_Trans_iterate = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_iterate(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_iterate
}

var cache_Control_Monad_List_Trans_repeat gopurs_runtime.Value
var once_Control_Monad_List_Trans_repeat sync.Once
func Get_Control_Monad_List_Trans_repeat() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_repeat.Do(func() {
		cache_Control_Monad_List_Trans_repeat = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_repeat(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_repeat
}

var cache_Control_Monad_List_Trans_head gopurs_runtime.Value
var once_Control_Monad_List_Trans_head sync.Once
func Get_Control_Monad_List_Trans_head() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_head.Do(func() {
		cache_Control_Monad_List_Trans_head = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_head(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
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
return Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_fromEffect
}

var cache_Control_Monad_List_Trans_monadTransListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_monadTransListT sync.Once
func Get_Control_Monad_List_Trans_monadTransListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_monadTransListT.Do(func() {
		cache_Control_Monad_List_Trans_monadTransListT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
}))
	})
	return cache_Control_Monad_List_Trans_monadTransListT
}

var cache_Control_Monad_List_Trans_lift gopurs_runtime.Value
var once_Control_Monad_List_Trans_lift sync.Once
func Get_Control_Monad_List_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_lift.Do(func() {
		cache_Control_Monad_List_Trans_lift = gopurs_runtime.RecordGet(Get_Control_Monad_List_Trans_monadTransListT(), "lift")
	})
	return cache_Control_Monad_List_Trans_lift
}

var cache_Control_Monad_List_Trans_foldlRec_prime gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldlRec_prime sync.Once
func Get_Control_Monad_List_Trans_foldlRec_prime() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldlRec_prime.Do(func() {
		cache_Control_Monad_List_Trans_foldlRec_prime = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldlRec_prime(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldlRec_prime
}

var cache_Control_Monad_List_Trans_runListTRec gopurs_runtime.Value
var once_Control_Monad_List_Trans_runListTRec sync.Once
func Get_Control_Monad_List_Trans_runListTRec() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_runListTRec.Do(func() {
		cache_Control_Monad_List_Trans_runListTRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_runListTRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_List_Trans_runListTRec
}

var cache_Control_Monad_List_Trans_foldlRec gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldlRec sync.Once
func Get_Control_Monad_List_Trans_foldlRec() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldlRec.Do(func() {
		cache_Control_Monad_List_Trans_foldlRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldlRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldlRec
}

var cache_Control_Monad_List_Trans_foldl_prime gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldl_prime sync.Once
func Get_Control_Monad_List_Trans_foldl_prime() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldl_prime.Do(func() {
		cache_Control_Monad_List_Trans_foldl_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldl_prime(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldl_prime
}

var cache_Control_Monad_List_Trans_runListT gopurs_runtime.Value
var once_Control_Monad_List_Trans_runListT sync.Once
func Get_Control_Monad_List_Trans_runListT() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_runListT.Do(func() {
		cache_Control_Monad_List_Trans_runListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_runListT(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_runListT
}

var cache_Control_Monad_List_Trans_foldl gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldl sync.Once
func Get_Control_Monad_List_Trans_foldl() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldl.Do(func() {
		cache_Control_Monad_List_Trans_foldl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldl(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldl
}

var cache_Control_Monad_List_Trans_filter gopurs_runtime.Value
var once_Control_Monad_List_Trans_filter sync.Once
func Get_Control_Monad_List_Trans_filter() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_filter.Do(func() {
		cache_Control_Monad_List_Trans_filter = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_filter
}

var cache_Control_Monad_List_Trans_dropWhile gopurs_runtime.Value
var once_Control_Monad_List_Trans_dropWhile sync.Once
func Get_Control_Monad_List_Trans_dropWhile() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_dropWhile.Do(func() {
		cache_Control_Monad_List_Trans_dropWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_dropWhile(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_dropWhile
}

var cache_Control_Monad_List_Trans_drop gopurs_runtime.Value
var once_Control_Monad_List_Trans_drop sync.Once
func Get_Control_Monad_List_Trans_drop() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_drop.Do(func() {
		cache_Control_Monad_List_Trans_drop = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_drop(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_drop
}

var cache_Control_Monad_List_Trans_cons gopurs_runtime.Value
var once_Control_Monad_List_Trans_cons sync.Once
func Get_Control_Monad_List_Trans_cons() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_cons.Do(func() {
		cache_Control_Monad_List_Trans_cons = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_cons(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), lh_1_box, t_2_box)
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
return Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
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
return Call_Control_Monad_List_Trans_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box))
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

var cache_Control_Monad_List_Trans_cons__808523158 gopurs_runtime.Value
var once_Control_Monad_List_Trans_cons__808523158 sync.Once
func Get_Control_Monad_List_Trans_cons__808523158() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_cons__808523158.Do(func() {
		cache_Control_Monad_List_Trans_cons__808523158 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_cons__808523158(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), lh_1_box, t_2_box)
})
	})
	return cache_Control_Monad_List_Trans_cons__808523158
}

var cache_Control_Monad_List_Trans_drop__1964165395 gopurs_runtime.Value
var once_Control_Monad_List_Trans_drop__1964165395 sync.Once
func Get_Control_Monad_List_Trans_drop__1964165395() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_drop__1964165395.Do(func() {
		cache_Control_Monad_List_Trans_drop__1964165395 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_drop__1964165395(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_drop__1964165395
}

var cache_Control_Monad_List_Trans_dropWhile__504781836 gopurs_runtime.Value
var once_Control_Monad_List_Trans_dropWhile__504781836 sync.Once
func Get_Control_Monad_List_Trans_dropWhile__504781836() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_dropWhile__504781836.Do(func() {
		cache_Control_Monad_List_Trans_dropWhile__504781836 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_dropWhile__504781836(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_dropWhile__504781836
}

var cache_Control_Monad_List_Trans_filter__1345510683 gopurs_runtime.Value
var once_Control_Monad_List_Trans_filter__1345510683 sync.Once
func Get_Control_Monad_List_Trans_filter__1345510683() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_filter__1345510683.Do(func() {
		cache_Control_Monad_List_Trans_filter__1345510683 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_filter__1345510683(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_filter__1345510683
}

var cache_Control_Monad_List_Trans_foldl_prime__3412851976 gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldl_prime__3412851976 sync.Once
func Get_Control_Monad_List_Trans_foldl_prime__3412851976() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldl_prime__3412851976.Do(func() {
		cache_Control_Monad_List_Trans_foldl_prime__3412851976 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldl_prime__3412851976(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldl_prime__3412851976
}

var cache_Control_Monad_List_Trans_foldl_prime__2387145256 gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldl_prime__2387145256 sync.Once
func Get_Control_Monad_List_Trans_foldl_prime__2387145256() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldl_prime__2387145256.Do(func() {
		cache_Control_Monad_List_Trans_foldl_prime__2387145256 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldl_prime__2387145256(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldl_prime__2387145256
}

var cache_Control_Monad_List_Trans_foldlRec_prime__4148996870 gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldlRec_prime__4148996870 sync.Once
func Get_Control_Monad_List_Trans_foldlRec_prime__4148996870() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldlRec_prime__4148996870.Do(func() {
		cache_Control_Monad_List_Trans_foldlRec_prime__4148996870 = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldlRec_prime__4148996870(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldlRec_prime__4148996870
}

var cache_Control_Monad_List_Trans_foldlRec_prime__1739794342 gopurs_runtime.Value
var once_Control_Monad_List_Trans_foldlRec_prime__1739794342 sync.Once
func Get_Control_Monad_List_Trans_foldlRec_prime__1739794342() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_foldlRec_prime__1739794342.Do(func() {
		cache_Control_Monad_List_Trans_foldlRec_prime__1739794342 = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_foldlRec_prime__1739794342(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_List_Trans_foldlRec_prime__1739794342
}

var cache_Control_Monad_List_Trans_iterate__4162284821 gopurs_runtime.Value
var once_Control_Monad_List_Trans_iterate__4162284821 sync.Once
func Get_Control_Monad_List_Trans_iterate__4162284821() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_iterate__4162284821.Do(func() {
		cache_Control_Monad_List_Trans_iterate__4162284821 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_iterate__4162284821(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_iterate__4162284821
}

var cache_Control_Monad_List_Trans_mapMaybe__3319479893 gopurs_runtime.Value
var once_Control_Monad_List_Trans_mapMaybe__3319479893 sync.Once
func Get_Control_Monad_List_Trans_mapMaybe__3319479893() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_mapMaybe__3319479893.Do(func() {
		cache_Control_Monad_List_Trans_mapMaybe__3319479893 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_mapMaybe__3319479893(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_mapMaybe__3319479893
}

var cache_Control_Monad_List_Trans_mapMaybe__3325666580 gopurs_runtime.Value
var once_Control_Monad_List_Trans_mapMaybe__3325666580 sync.Once
func Get_Control_Monad_List_Trans_mapMaybe__3325666580() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_mapMaybe__3325666580.Do(func() {
		cache_Control_Monad_List_Trans_mapMaybe__3325666580 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_mapMaybe__3325666580(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_mapMaybe__3325666580
}

var cache_Control_Monad_List_Trans_nil__1472516796 gopurs_runtime.Value
var once_Control_Monad_List_Trans_nil__1472516796 sync.Once
func Get_Control_Monad_List_Trans_nil__1472516796() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_nil__1472516796.Do(func() {
		cache_Control_Monad_List_Trans_nil__1472516796 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_nil__1472516796(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_List_Trans_nil__1472516796
}

var cache_Control_Monad_List_Trans_prepend__2860458454 gopurs_runtime.Value
var once_Control_Monad_List_Trans_prepend__2860458454 sync.Once
func Get_Control_Monad_List_Trans_prepend__2860458454() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_prepend__2860458454.Do(func() {
		cache_Control_Monad_List_Trans_prepend__2860458454 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_prepend__2860458454(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_Control_Monad_List_Trans_prepend__2860458454
}

var cache_Control_Monad_List_Trans_prepend_prime__1901723831 gopurs_runtime.Value
var once_Control_Monad_List_Trans_prepend_prime__1901723831 sync.Once
func Get_Control_Monad_List_Trans_prepend_prime__1901723831() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_prepend_prime__1901723831.Do(func() {
		cache_Control_Monad_List_Trans_prepend_prime__1901723831 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_prepend_prime__1901723831(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_Control_Monad_List_Trans_prepend_prime__1901723831
}

var cache_Control_Monad_List_Trans_singleton__2427543124 gopurs_runtime.Value
var once_Control_Monad_List_Trans_singleton__2427543124 sync.Once
func Get_Control_Monad_List_Trans_singleton__2427543124() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_singleton__2427543124.Do(func() {
		cache_Control_Monad_List_Trans_singleton__2427543124 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_singleton__2427543124(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_singleton__2427543124
}

var cache_Control_Monad_List_Trans_stepMap__3249590196 gopurs_runtime.Value
var once_Control_Monad_List_Trans_stepMap__3249590196 sync.Once
func Get_Control_Monad_List_Trans_stepMap__3249590196() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_stepMap__3249590196.Do(func() {
		cache_Control_Monad_List_Trans_stepMap__3249590196 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_stepMap__3249590196(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_stepMap__3249590196
}

var cache_Control_Monad_List_Trans_stepMap__167039253 gopurs_runtime.Value
var once_Control_Monad_List_Trans_stepMap__167039253 sync.Once
func Get_Control_Monad_List_Trans_stepMap__167039253() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_stepMap__167039253.Do(func() {
		cache_Control_Monad_List_Trans_stepMap__167039253 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_stepMap__167039253(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_List_Trans_stepMap__167039253
}

var cache_Control_Monad_List_Trans_take__1964165395 gopurs_runtime.Value
var once_Control_Monad_List_Trans_take__1964165395 sync.Once
func Get_Control_Monad_List_Trans_take__1964165395() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_take__1964165395.Do(func() {
		cache_Control_Monad_List_Trans_take__1964165395 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_take__1964165395(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_take__1964165395
}

var cache_Control_Monad_List_Trans_takeWhile__504781836 gopurs_runtime.Value
var once_Control_Monad_List_Trans_takeWhile__504781836 sync.Once
func Get_Control_Monad_List_Trans_takeWhile__504781836() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_takeWhile__504781836.Do(func() {
		cache_Control_Monad_List_Trans_takeWhile__504781836 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_takeWhile__504781836(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Monad_List_Trans_takeWhile__504781836
}

var cache_Control_Monad_List_Trans_uncons__1307401241 gopurs_runtime.Value
var once_Control_Monad_List_Trans_uncons__1307401241 sync.Once
func Get_Control_Monad_List_Trans_uncons__1307401241() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_uncons__1307401241.Do(func() {
		cache_Control_Monad_List_Trans_uncons__1307401241 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_uncons__1307401241(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_uncons__1307401241
}

var cache_Control_Monad_List_Trans_unfold__3487137686 gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfold__3487137686 sync.Once
func Get_Control_Monad_List_Trans_unfold__3487137686() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfold__3487137686.Do(func() {
		cache_Control_Monad_List_Trans_unfold__3487137686 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_unfold__3487137686(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_unfold__3487137686
}

var cache_Control_Monad_List_Trans_unfold__2471180757 gopurs_runtime.Value
var once_Control_Monad_List_Trans_unfold__2471180757 sync.Once
func Get_Control_Monad_List_Trans_unfold__2471180757() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_unfold__2471180757.Do(func() {
		cache_Control_Monad_List_Trans_unfold__2471180757 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_unfold__2471180757(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_unfold__2471180757
}

var cache_Control_Monad_List_Trans_wrapEffect__3965193927 gopurs_runtime.Value
var once_Control_Monad_List_Trans_wrapEffect__3965193927 sync.Once
func Get_Control_Monad_List_Trans_wrapEffect__3965193927() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_wrapEffect__3965193927.Do(func() {
		cache_Control_Monad_List_Trans_wrapEffect__3965193927 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_wrapEffect__3965193927(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), v_1_box)
})
	})
	return cache_Control_Monad_List_Trans_wrapEffect__3965193927
}

var cache_Control_Monad_List_Trans_zipWith_prime__376166203 gopurs_runtime.Value
var once_Control_Monad_List_Trans_zipWith_prime__376166203 sync.Once
func Get_Control_Monad_List_Trans_zipWith_prime__376166203() gopurs_runtime.Value {
	once_Control_Monad_List_Trans_zipWith_prime__376166203.Do(func() {
		cache_Control_Monad_List_Trans_zipWith_prime__376166203 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_zipWith_prime__376166203(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_List_Trans_zipWith_prime__376166203
}

type Constructor_Control_Monad_List_Trans_Yield struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Control_Monad_List_Trans_Skip struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Control_Monad_List_Trans_Done struct {
	Rc uint32
}


func Call_Control_Monad_List_Trans_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_List_Trans_identity1(x_0_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var x_0 *Constructor_Data_Maybe_Just = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_List_Trans_ListT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_List_Trans_wrapLazy(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, v_1})})
}

func Call_Control_Monad_List_Trans_wrapEffect(dictFunctor_0_loop *Constructor_Data_Functor_Functor, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))})}
}), v_1)
}

func Call_Control_Monad_List_Trans_unfold(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
unfold:
for {
if false { continue unfold }
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), f_2, __local_var_5_1)
}))})}
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
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
}), gopurs_runtime.Apply(f_2, z_3))
})
})
}
}

func Call_Control_Monad_List_Trans_uncons(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
uncons:
for {
if false { continue uncons }
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_4, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1320412129) {
__t3 = gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_List_Trans_Yield)(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield)(v1_5.UnsafePtr).V1)})}})})
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 813447293) {
__t3 = gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Skip)(v1_5.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 489128924) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
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

func Call_Control_Monad_List_Trans_tail(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Tuple_snd()), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_2))
})
}

func Call_Control_Monad_List_Trans_stepMap(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_1, v_2)
}

func Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0), f_2), (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0), f_2), (*Constructor_Control_Monad_List_Trans_Skip)(v_4.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
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
}), v_3)
})
})
}
}

func Call_Control_Monad_List_Trans_scanl(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0
_ = __local_var_6_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_7.Type == 9 && v1_7.IntVal == 1320412129) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(f_2, __local_var_6_1, (*Constructor_Control_Monad_List_Trans_Yield)(v1_7.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield)(v1_7.UnsafePtr).V1)})}, __local_var_6_1})}}
goto end_branch_2
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 813447293) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_6_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Skip)(v1_7.UnsafePtr).V0)})}, __local_var_6_1})}}
goto end_branch_2
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 489128924) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, b_3, l_4})})
})
})
})
}

func Call_Control_Monad_List_Trans_prepend_prime(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, h_1, t_2})})
}

func Call_Control_Monad_List_Trans_prepend(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, h_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return t_2
}))})})
}

func Call_Control_Monad_List_Trans_nil(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_Control_Monad_List_Trans_singleton(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): nil1_1_0 -> gopurs_runtime.Value
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, a_2, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}))})})
})
}

func Call_Control_Monad_List_Trans_take(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
take:
for {
if false { continue take }
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): nil1_1_0 -> gopurs_runtime.Value
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.IntVal) == (0) {
__t3 = nil1_1_0
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1320412129) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v2_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_take(dictApplicative_0), gopurs_runtime.Int((v_3.IntVal) - (1))), (*Constructor_Control_Monad_List_Trans_Yield)(v2_5.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_take(dictApplicative_0), gopurs_runtime.Int(v_3.IntVal)), (*Constructor_Control_Monad_List_Trans_Skip)(v2_5.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 489128924) {
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
}), v1_4)
}
end_branch_3:
return __t3
})
})
}
}

func Call_Control_Monad_List_Trans_zipWith_prime(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
zipWith_prime:
for {
if false { continue zipWith_prime }
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): prepend_prime1_3_2 -> gopurs_runtime.Value
prepend_prime1_3_2 := gopurs_runtime.Func(func(h_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, h_4, t_5})})
})
})
_ = prepend_prime1_3_2
// TAST (Let): Bind1_4_4 -> *Constructor_Control_Bind_Bind
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fa_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fb_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
}))})}
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), fa_6), gopurs_runtime.Func(func(ua_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), fb_7), gopurs_runtime.Func(func(ub_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](ub_9)
if (__t_tag_5 == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](ua_8)
if (__t_tag_6 == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](ua_8)
var __t_and_9 bool = false
if (__t_tag_7 != nil) {

var __t_tag_8 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](ub_9)
__t_and_9 = (__t_tag_8 != nil)
}
if __t_and_9 {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(ua_8.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_10_10
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(ub_9.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_11
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Control_Monad_List_Trans_zipWith_prime(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_5, __local_var_10_10, __local_var_11_11)
}))
_ = __local_var_12_12
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(prepend_prime1_3_2, a_13, __local_var_12_12)
}), gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(ua_8.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(ub_9.UnsafePtr).V0.UnsafePtr).V0))
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
})
})
}
}

func Call_Control_Monad_List_Trans_zipWith(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_zipWith_prime(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Apply2(f_2, a_3, b_4))
})
}))
})
}

func Call_Control_Monad_List_Trans_mapMaybe(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
// TAST (Let): __local_var_4_0 -> *Constructor_Data_Maybe_Just
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Control_Monad_List_Trans_Yield(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0)))}))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0 == nil) {
__t1 = Get_Control_Monad_List_Trans_Skip()
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 != nil) {
__t1 = (__local_var_4_0).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Apply(__t1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
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

func Call_Control_Monad_List_Trans_iterate(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, x_4), x_4})}})})
}), a_3)
})
})
}

func Call_Control_Monad_List_Trans_repeat(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply(Call_Control_Monad_List_Trans_iterate(dictMonad_0), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Control_Monad_List_Trans_head(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Tuple_fst()), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_2))
})
}

func Call_Control_Monad_List_Trans_functorListT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
functorListT:
for {
if false { continue functorListT }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(dictFunctor_0), "map"), f_1), (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_functorListT(dictFunctor_0), "map"), f_1), (*Constructor_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), v_2)
})
}))
}
}

func Call_Control_Monad_List_Trans_fromEffect(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}))
_ = __local_var_3_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, a_4, __local_var_3_1})}
}), fa_2)
})
}

func Call_Control_Monad_List_Trans_foldlRec_prime(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Monad01_4_3 -> *Constructor_Control_Monad_Monad
Monad01_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 -> gopurs_runtime.Value
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, __local_var_9_4})})
goto end_branch_6
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
// TAST (Let): __local_var_11_5 -> gopurs_runtime.Value
__local_var_11_5 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_5
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(f_5, __local_var_9_4, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_5)})})
}))
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
}), gopurs_runtime.RecordDict2("a", "b", a_6, b_7))
})
})
})
}

func Call_Control_Monad_List_Trans_runListTRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Apply3(Get_Control_Monad_List_Trans_foldlRec_prime(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
})
}), Get_Data_Unit_unit())
}

func Call_Control_Monad_List_Trans_foldlRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Monad01_4_3 -> *Constructor_Control_Monad_Monad
Monad01_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 -> gopurs_runtime.Value
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, __local_var_9_4})})
goto end_branch_5
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Apply2(f_5, __local_var_9_4, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V1)})})
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
}), gopurs_runtime.RecordDict2("a", "b", a_6, b_7))
})
})
})
}

func Call_Control_Monad_List_Trans_foldl_prime(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_0 gopurs_runtime.Value
_ = loop_4_2_0
loop_4_2_0 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_5)
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
// TAST (Let): __local_var_8_3 -> gopurs_runtime.Value
__local_var_8_3 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_3
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_5, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(loop_4_2_0, a_9, __local_var_8_3)
}))
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
})
})
return loop_4_2_0
})
}

func Call_Control_Monad_List_Trans_runListT(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Apply3(Get_Control_Monad_List_Trans_foldl_prime(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
})
}), Get_Data_Unit_unit())
}

func Call_Control_Monad_List_Trans_foldl(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_1 gopurs_runtime.Value
_ = loop_4_2_1
loop_4_2_1 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_5)
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(loop_4_2_1, gopurs_runtime.Apply2(f_3, b_5, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V1)
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
})
return loop_4_2_1
})
}

func Call_Control_Monad_List_Trans_filter(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
// TAST (Let): s_prime_4_0 -> gopurs_runtime.Value
s_prime_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1)
_ = s_prime_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0, s_prime_4_0})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, s_prime_4_0})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
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

func Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
dropWhile:
for {
if false { continue dropWhile }
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0), f_2), (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0), f_2), (*Constructor_Control_Monad_List_Trans_Skip)(v_4.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
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
}), v_3)
})
})
}
}

func Call_Control_Monad_List_Trans_drop(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_drop(dictApplicative_0), gopurs_runtime.Int((v_2.IntVal) - (1))), (*Constructor_Control_Monad_List_Trans_Yield)(v2_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_drop(dictApplicative_0), gopurs_runtime.Int(v_2.IntVal)), (*Constructor_Control_Monad_List_Trans_Skip)(v2_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v1_3)
}
end_branch_2:
return __t2
})
})
}
}

func Call_Control_Monad_List_Trans_cons(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, lh_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 gopurs_runtime.Value = lh_1_loop
_ = lh_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), lh_1), t_2})})
}

func Call_Control_Monad_List_Trans_unfoldable1ListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_2 gopurs_runtime.Value
_ = go__go_4_1_2
go__go_4_1_2 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Call_Control_Monad_List_Trans_singleton(Applicative0_1_0), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_Maybe_Just)((*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1.UnsafePtr).V0
_ = __local_var_6_4
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_applicativeLazy(), "pure"), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_4_1_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_2, __local_var_6_4)))})
}))})})
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
return gopurs_runtime.Apply(go__go_4_1_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_2, b_3)))})
})
}))
}

func Call_Control_Monad_List_Trans_unfoldableListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): unfoldable1ListT1_2_1 -> gopurs_runtime.Value
unfoldable1ListT1_2_1 := Call_Control_Monad_List_Trans_unfoldable1ListT(dictMonad_0)
_ = unfoldable1ListT1_2_1
return gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return unfoldable1ListT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_2_3 gopurs_runtime.Value
_ = go__go_5_2_3
go__go_5_2_3 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 930809136 && v_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
goto end_branch_4
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 930809136 && v_6.UnsafePtr != nil) {
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_6.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_applicativeLazy(), "pure"), (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_6.UnsafePtr).V0.UnsafePtr).V0)), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_5_2_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_3, __local_var_7_3)))})
}))})})
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
return gopurs_runtime.Apply(go__go_5_2_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_3, b_4)))})
})
}))
}

func Call_Control_Monad_List_Trans_semigroupListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.RecordDict1("append", Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0)))
}

func Call_Control_Monad_List_Trans_concat(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_concat(dictApplicative_0), v1_5, y_3)
}), (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_concat(dictApplicative_0), v1_5, y_3)
}), (*Constructor_Control_Monad_List_Trans_Skip)(v_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return y_3
}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), x_2)
})
})
}

func Call_Control_Monad_List_Trans_monoidListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): semigroupListT1_1_0 -> gopurs_runtime.Value
semigroupListT1_1_0 := gopurs_runtime.RecordDict1("append", Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0)))
_ = semigroupListT1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupListT1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
}

func Call_Control_Monad_List_Trans_catMaybes(dictFunctor_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Control_Monad_List_Trans_monadListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_bindListT(dictMonad_0)
}))
}

func Call_Control_Monad_List_Trans_bindListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): semigroupListT1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupListT1_1_0 := &Constructor_Data_Semigroup_Semigroup{1, Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))}
_ = semigroupListT1_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applyListT(dictMonad_0)
}), gopurs_runtime.Func(func(fa_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1320412129) {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := (*Constructor_Control_Monad_List_Trans_Yield)(v_5.UnsafePtr).V0
_ = __local_var_6_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(semigroupListT1_1_0.V0), gopurs_runtime.Apply(f_4, __local_var_6_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), s_prime_7, f_4))
}), (*Constructor_Control_Monad_List_Trans_Yield)(v_5.UnsafePtr).V1)})}
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 813447293) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_List_Trans_bindListT(dictMonad_0), "bind"), v1_6, f_4)
}), (*Constructor_Control_Monad_List_Trans_Skip)(v_5.UnsafePtr).V0)})}
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 489128924) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), fa_3)
})
}))
}

func Call_Control_Monad_List_Trans_applyListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): functorListT1_1_0 -> gopurs_runtime.Value
functorListT1_1_0 := Call_Control_Monad_List_Trans_functorListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_bindListT(dictMonad_0)
}))
_ = __local_var_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Applicative0_4_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_3.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_Control_Monad_List_Trans_applicativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applyListT(dictMonad_0)
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))))
}

func Call_Control_Monad_List_Trans_monadEffectListT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadListT1_2_1 -> gopurs_runtime.Value
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applyListT(Monad0_1_0)
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_bindListT(Monad0_1_0)
}))
_ = monadListT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_List_Trans_monadTransListT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_Control_Monad_List_Trans_monadSTListT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadListT1_2_1 -> gopurs_runtime.Value
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applyListT(Monad0_1_0)
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_bindListT(Monad0_1_0)
}))
_ = monadListT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_List_Trans_monadTransListT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_Control_Monad_List_Trans_altListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): functorListT1_1_0 -> gopurs_runtime.Value
functorListT1_1_0 := Call_Control_Monad_List_Trans_functorListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), Call_Control_Monad_List_Trans_concat(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0)))
}

func Call_Control_Monad_List_Trans_plusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> gopurs_runtime.Value
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
// TAST (Let): altListT1_2_1 -> gopurs_runtime.Value
altListT1_2_1 := Call_Control_Monad_List_Trans_altListT(Applicative0_1_0)
_ = altListT1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altListT1_2_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
}

func Call_Control_Monad_List_Trans_alternativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeListT1_1_0 -> gopurs_runtime.Value
applicativeListT1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applyListT(dictMonad_0)
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))))
_ = applicativeListT1_1_0
// TAST (Let): plusListT1_2_1 -> gopurs_runtime.Value
plusListT1_2_1 := Call_Control_Monad_List_Trans_plusListT(dictMonad_0)
_ = plusListT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeListT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusListT1_2_1
}))
}

func Call_Control_Monad_List_Trans_monadPlusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadListT1_1_0 -> gopurs_runtime.Value
monadListT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_applyListT(dictMonad_0)
}), Call_Control_Monad_List_Trans_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_List_Trans_bindListT(dictMonad_0)
}))
_ = monadListT1_1_0
// TAST (Let): alternativeListT1_2_1 -> gopurs_runtime.Value
alternativeListT1_2_1 := Call_Control_Monad_List_Trans_alternativeListT(dictMonad_0)
_ = alternativeListT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeListT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_1_0
}))
}

func Call_Control_Monad_List_Trans_cons__808523158(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, lh_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 gopurs_runtime.Value = lh_1_loop
_ = lh_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, gopurs_runtime.Apply(Get_Data_Lazy_force(), lh_1), t_2})})
}

func Call_Control_Monad_List_Trans_drop__1964165395(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_drop(dictApplicative_0), gopurs_runtime.Int((v_2.IntVal) - (1))), (*Constructor_Control_Monad_List_Trans_Yield)(v2_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_drop(dictApplicative_0), gopurs_runtime.Int(v_2.IntVal)), (*Constructor_Control_Monad_List_Trans_Skip)(v2_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v1_3)
}
end_branch_2:
return __t2
})
})
}

func Call_Control_Monad_List_Trans_dropWhile__504781836(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0), f_2), (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_dropWhile(dictApplicative_0), f_2), (*Constructor_Control_Monad_List_Trans_Skip)(v_4.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
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
}), v_3)
})
})
}

func Call_Control_Monad_List_Trans_filter__1345510683(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
// TAST (Let): s_prime_4_0 -> gopurs_runtime.Value
s_prime_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1)
_ = s_prime_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0, s_prime_4_0})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, s_prime_4_0})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
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

func Call_Control_Monad_List_Trans_foldl_prime__3412851976(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_4 gopurs_runtime.Value
_ = loop_4_2_4
loop_4_2_4 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_5)
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
// TAST (Let): __local_var_8_3 -> gopurs_runtime.Value
__local_var_8_3 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_3
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_5, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(loop_4_2_4, a_9, __local_var_8_3)
}))
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
})
})
return loop_4_2_4
})
}

func Call_Control_Monad_List_Trans_foldl_prime__2387145256(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_5 gopurs_runtime.Value
_ = loop_4_2_5
loop_4_2_5 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_5)
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
// TAST (Let): __local_var_8_3 -> gopurs_runtime.Value
__local_var_8_3 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_3
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_5, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(loop_4_2_5, a_9, __local_var_8_3)
}))
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
})
})
return loop_4_2_5
})
}

func Call_Control_Monad_List_Trans_foldlRec_prime__4148996870(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Monad01_4_3 -> *Constructor_Control_Monad_Monad
Monad01_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 -> gopurs_runtime.Value
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, __local_var_9_4})})
goto end_branch_6
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
// TAST (Let): __local_var_11_5 -> gopurs_runtime.Value
__local_var_11_5 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_5
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(f_5, __local_var_9_4, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_5)})})
}))
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
}), gopurs_runtime.RecordDict2("a", "b", a_6, b_7))
})
})
})
}

func Call_Control_Monad_List_Trans_foldlRec_prime__1739794342(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Monad01_4_3 -> *Constructor_Control_Monad_Monad
Monad01_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 -> gopurs_runtime.Value
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, __local_var_9_4})})
goto end_branch_6
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
// TAST (Let): __local_var_11_5 -> gopurs_runtime.Value
__local_var_11_5 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_5
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(f_5, __local_var_9_4, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_5)})})
}))
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
}), gopurs_runtime.RecordDict2("a", "b", a_6, b_7))
})
})
})
}

func Call_Control_Monad_List_Trans_iterate__4162284821(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, x_4), x_4})}})})
}), a_3)
})
})
}

func Call_Control_Monad_List_Trans_mapMaybe__3319479893(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
// TAST (Let): __local_var_4_0 -> *Constructor_Data_Maybe_Just
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Control_Monad_List_Trans_Yield(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0)))}))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0 == nil) {
__t1 = Get_Control_Monad_List_Trans_Skip()
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 != nil) {
__t1 = (__local_var_4_0).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Apply(__t1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
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

func Call_Control_Monad_List_Trans_mapMaybe__3325666580(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
// TAST (Let): __local_var_4_0 -> *Constructor_Data_Maybe_Just
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Control_Monad_List_Trans_Yield(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0))})))}))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0 == nil) {
__t1 = Get_Control_Monad_List_Trans_Skip()
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 != nil) {
__t1 = (__local_var_4_0).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Apply(__t1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply2(Get_Control_Monad_List_Trans_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
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

func Call_Control_Monad_List_Trans_nil__1472516796(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_Control_Monad_List_Trans_prepend__2860458454(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, h_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return t_2
}))})})
}

func Call_Control_Monad_List_Trans_prepend_prime__1901723831(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, h_1, t_2})})
}

func Call_Control_Monad_List_Trans_singleton__2427543124(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): nil1_1_0 -> gopurs_runtime.Value
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, a_2, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}))})})
})
}

func Call_Control_Monad_List_Trans_stepMap__3249590196(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_1, v_2)
}

func Call_Control_Monad_List_Trans_stepMap__167039253(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_1, v_2)
}

func Call_Control_Monad_List_Trans_take__1964165395(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): nil1_1_0 -> gopurs_runtime.Value
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.IntVal) == (0) {
__t3 = nil1_1_0
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1320412129) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v2_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_take(dictApplicative_0), gopurs_runtime.Int((v_3.IntVal) - (1))), (*Constructor_Control_Monad_List_Trans_Yield)(v2_5.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_take(dictApplicative_0), gopurs_runtime.Int(v_3.IntVal)), (*Constructor_Control_Monad_List_Trans_Skip)(v2_5.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 489128924) {
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
}), v1_4)
}
end_branch_3:
return __t3
})
})
}

func Call_Control_Monad_List_Trans_takeWhile__504781836(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0), f_2), (*Constructor_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_takeWhile(dictApplicative_0), f_2), (*Constructor_Control_Monad_List_Trans_Skip)(v_4.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
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
}), v_3)
})
})
}

func Call_Control_Monad_List_Trans_uncons__1307401241(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_4, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1320412129) {
__t3 = gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_List_Trans_Yield)(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Yield)(v1_5.UnsafePtr).V1)})}})})
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 813447293) {
__t3 = gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Control_Monad_List_Trans_Skip)(v1_5.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 489128924) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
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

func Call_Control_Monad_List_Trans_unfold__3487137686(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), f_2, __local_var_5_1)
}))})}
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
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
}), gopurs_runtime.Apply(f_2, z_3))
})
})
}

func Call_Control_Monad_List_Trans_unfold__2471180757(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Monad_List_Trans_unfold(dictMonad_0), f_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](__local_var_5_1))})
}))})}
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
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
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](z_3))}))
})
})
}

func Call_Control_Monad_List_Trans_wrapEffect__3965193927(dictFunctor_0_loop *Constructor_Data_Functor_Functor, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))})}
}), v_1)
}

func Call_Control_Monad_List_Trans_zipWith_prime__376166203(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): prepend_prime1_3_2 -> gopurs_runtime.Value
prepend_prime1_3_2 := gopurs_runtime.Func(func(h_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Yield{1, h_4, t_5})})
})
})
_ = prepend_prime1_3_2
// TAST (Let): Bind1_4_4 -> *Constructor_Control_Bind_Bind
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fa_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fb_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_List_Trans_Skip{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
}))})}
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), fa_6), gopurs_runtime.Func(func(ua_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), gopurs_runtime.Apply(Call_Control_Monad_List_Trans_uncons(dictMonad_0), fb_7), gopurs_runtime.Func(func(ub_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](ub_9)
if (__t_tag_5 == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](ua_8)
if (__t_tag_6 == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](ua_8)
var __t_and_9 bool = false
if (__t_tag_7 != nil) {

var __t_tag_8 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](ub_9)
__t_and_9 = (__t_tag_8 != nil)
}
if __t_and_9 {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(ua_8.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_10_10
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(ub_9.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_11
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Control_Monad_List_Trans_zipWith_prime(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_5, __local_var_10_10, __local_var_11_11)
}))
_ = __local_var_12_12
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(prepend_prime1_3_2, a_13, __local_var_12_12)
}), gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(ua_8.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(ub_9.UnsafePtr).V0.UnsafePtr).V0))
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
})
})
}


