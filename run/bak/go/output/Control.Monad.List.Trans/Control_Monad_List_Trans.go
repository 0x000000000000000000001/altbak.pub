package Control_Monad_List_Trans

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_identity1(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_identity1
}

var cache_Yield gopurs_runtime.Value
var once_Yield sync.Once
func Get_Yield() gopurs_runtime.Value {
	once_Yield.Do(func() {
		cache_Yield = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Yield
}

var cache_Skip gopurs_runtime.Value
var once_Skip sync.Once
func Get_Skip() gopurs_runtime.Value {
	once_Skip.Do(func() {
		cache_Skip = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Skip
}

var cache_Done gopurs_runtime.Value
var once_Done sync.Once
func Get_Done() gopurs_runtime.Value {
	once_Done.Do(func() {
		cache_Done = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Done
}

var cache_ListT gopurs_runtime.Value
var once_ListT sync.Once
func Get_ListT() gopurs_runtime.Value {
	once_ListT.Do(func() {
		cache_ListT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ListT(x_0_box)
})
	})
	return cache_ListT
}

var cache_wrapLazy gopurs_runtime.Value
var once_wrapLazy sync.Once
func Get_wrapLazy() gopurs_runtime.Value {
	once_wrapLazy.Do(func() {
		cache_wrapLazy = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrapLazy(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), v_1_box)
})
	})
	return cache_wrapLazy
}

var cache_wrapEffect gopurs_runtime.Value
var once_wrapEffect sync.Once
func Get_wrapEffect() gopurs_runtime.Value {
	once_wrapEffect.Do(func() {
		cache_wrapEffect = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrapEffect(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
})
	})
	return cache_wrapEffect
}

var cache_unfold gopurs_runtime.Value
var once_unfold sync.Once
func Get_unfold() gopurs_runtime.Value {
	once_unfold.Do(func() {
		cache_unfold = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfold(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_unfold
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_uncons
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_tail
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_takeWhile
}

var cache_scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		cache_scanl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanl(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_scanl
}

var cache_prepend_prime gopurs_runtime.Value
var once_prepend_prime sync.Once
func Get_prepend_prime() gopurs_runtime.Value {
	once_prepend_prime.Do(func() {
		cache_prepend_prime = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prepend_prime(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_prepend_prime
}

var cache_prepend gopurs_runtime.Value
var once_prepend sync.Once
func Get_prepend() gopurs_runtime.Value {
	once_prepend.Do(func() {
		cache_prepend = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prepend(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_prepend
}

var cache_nil gopurs_runtime.Value
var once_nil sync.Once
func Get_nil() gopurs_runtime.Value {
	once_nil.Do(func() {
		cache_nil = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nil(dictApplicative_0_box)
})
	})
	return cache_nil
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_singleton
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_take
}

var cache_zipWith_prime gopurs_runtime.Value
var once_zipWith_prime sync.Once
func Get_zipWith_prime() gopurs_runtime.Value {
	once_zipWith_prime.Do(func() {
		cache_zipWith_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith_prime(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_zipWith_prime
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_zipWith
}

var cache_newtypeListT gopurs_runtime.Value
var once_newtypeListT sync.Once
func Get_newtypeListT() gopurs_runtime.Value {
	once_newtypeListT.Do(func() {
		cache_newtypeListT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeListT
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_mapMaybe
}

var cache_iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		cache_iterate = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_iterate
}

var cache_repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		cache_repeat = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_repeat
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_head
}

var cache_functorListT gopurs_runtime.Value
var once_functorListT sync.Once
func Get_functorListT() gopurs_runtime.Value {
	once_functorListT.Do(func() {
		cache_functorListT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorListT(dictFunctor_0_box)
})
	})
	return cache_functorListT
}

var cache_fromEffect gopurs_runtime.Value
var once_fromEffect sync.Once
func Get_fromEffect() gopurs_runtime.Value {
	once_fromEffect.Do(func() {
		cache_fromEffect = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEffect(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_fromEffect
}

var cache_monadTransListT gopurs_runtime.Value
var once_monadTransListT sync.Once
func Get_monadTransListT() gopurs_runtime.Value {
	once_monadTransListT.Do(func() {
		cache_monadTransListT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEffect(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
}))
	})
	return cache_monadTransListT
}

var cache_foldlRec_prime gopurs_runtime.Value
var once_foldlRec_prime sync.Once
func Get_foldlRec_prime() gopurs_runtime.Value {
	once_foldlRec_prime.Do(func() {
		cache_foldlRec_prime = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec_prime(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_foldlRec_prime
}

var cache_runListTRec gopurs_runtime.Value
var once_runListTRec sync.Once
func Get_runListTRec() gopurs_runtime.Value {
	once_runListTRec.Do(func() {
		cache_runListTRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runListTRec(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_runListTRec
}

var cache_foldlRec gopurs_runtime.Value
var once_foldlRec sync.Once
func Get_foldlRec() gopurs_runtime.Value {
	once_foldlRec.Do(func() {
		cache_foldlRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_foldlRec
}

var cache_foldl_prime gopurs_runtime.Value
var once_foldl_prime sync.Once
func Get_foldl_prime() gopurs_runtime.Value {
	once_foldl_prime.Do(func() {
		cache_foldl_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl_prime(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldl_prime
}

var cache_runListT gopurs_runtime.Value
var once_runListT sync.Once
func Get_runListT() gopurs_runtime.Value {
	once_runListT.Do(func() {
		cache_runListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runListT(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_runListT
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldl
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_filter
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_dropWhile
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_drop
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), lh_1_box, t_2_box)
})
	})
	return cache_cons
}

var cache_unfoldable1ListT gopurs_runtime.Value
var once_unfoldable1ListT sync.Once
func Get_unfoldable1ListT() gopurs_runtime.Value {
	once_unfoldable1ListT.Do(func() {
		cache_unfoldable1ListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable1ListT(dictMonad_0_box)
})
	})
	return cache_unfoldable1ListT
}

var cache_unfoldableListT gopurs_runtime.Value
var once_unfoldableListT sync.Once
func Get_unfoldableListT() gopurs_runtime.Value {
	once_unfoldableListT.Do(func() {
		cache_unfoldableListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldableListT(dictMonad_0_box)
})
	})
	return cache_unfoldableListT
}

var cache_semigroupListT gopurs_runtime.Value
var once_semigroupListT sync.Once
func Get_semigroupListT() gopurs_runtime.Value {
	once_semigroupListT.Do(func() {
		cache_semigroupListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupListT(dictApplicative_0_box)
})
	})
	return cache_semigroupListT
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concat(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_concat
}

var cache_monoidListT gopurs_runtime.Value
var once_monoidListT sync.Once
func Get_monoidListT() gopurs_runtime.Value {
	once_monoidListT.Do(func() {
		cache_monoidListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidListT(dictApplicative_0_box)
})
	})
	return cache_monoidListT
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catMaybes(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
})
	})
	return cache_catMaybes
}

var cache_monadListT gopurs_runtime.Value
var once_monadListT sync.Once
func Get_monadListT() gopurs_runtime.Value {
	once_monadListT.Do(func() {
		cache_monadListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadListT(dictMonad_0_box)
})
	})
	return cache_monadListT
}

var cache_bindListT gopurs_runtime.Value
var once_bindListT sync.Once
func Get_bindListT() gopurs_runtime.Value {
	once_bindListT.Do(func() {
		cache_bindListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindListT(dictMonad_0_box)
})
	})
	return cache_bindListT
}

var cache_applyListT gopurs_runtime.Value
var once_applyListT sync.Once
func Get_applyListT() gopurs_runtime.Value {
	once_applyListT.Do(func() {
		cache_applyListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyListT(dictMonad_0_box)
})
	})
	return cache_applyListT
}

var cache_applicativeListT gopurs_runtime.Value
var once_applicativeListT sync.Once
func Get_applicativeListT() gopurs_runtime.Value {
	once_applicativeListT.Do(func() {
		cache_applicativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeListT(dictMonad_0_box)
})
	})
	return cache_applicativeListT
}

var cache_monadEffectListT gopurs_runtime.Value
var once_monadEffectListT sync.Once
func Get_monadEffectListT() gopurs_runtime.Value {
	once_monadEffectListT.Do(func() {
		cache_monadEffectListT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectListT(dictMonadEffect_0_box)
})
	})
	return cache_monadEffectListT
}

var cache_monadSTListT gopurs_runtime.Value
var once_monadSTListT sync.Once
func Get_monadSTListT() gopurs_runtime.Value {
	once_monadSTListT.Do(func() {
		cache_monadSTListT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTListT(dictMonadST_0_box)
})
	})
	return cache_monadSTListT
}

var cache_altListT gopurs_runtime.Value
var once_altListT sync.Once
func Get_altListT() gopurs_runtime.Value {
	once_altListT.Do(func() {
		cache_altListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altListT(dictApplicative_0_box)
})
	})
	return cache_altListT
}

var cache_plusListT gopurs_runtime.Value
var once_plusListT sync.Once
func Get_plusListT() gopurs_runtime.Value {
	once_plusListT.Do(func() {
		cache_plusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusListT(dictMonad_0_box)
})
	})
	return cache_plusListT
}

var cache_alternativeListT gopurs_runtime.Value
var once_alternativeListT sync.Once
func Get_alternativeListT() gopurs_runtime.Value {
	once_alternativeListT.Do(func() {
		cache_alternativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeListT(dictMonad_0_box)
})
	})
	return cache_alternativeListT
}

var cache_monadPlusListT gopurs_runtime.Value
var once_monadPlusListT sync.Once
func Get_monadPlusListT() gopurs_runtime.Value {
	once_monadPlusListT.Do(func() {
		cache_monadPlusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusListT(dictMonad_0_box)
})
	})
	return cache_monadPlusListT
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

var cache_pure__160425008 gopurs_runtime.Value
var once_pure__160425008 sync.Once
func Get_pure__160425008() gopurs_runtime.Value {
	once_pure__160425008.Do(func() {
		cache_pure__160425008 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__160425008(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__160425008
}

var cache_pure__3197665392 gopurs_runtime.Value
var once_pure__3197665392 sync.Once
func Get_pure__3197665392() gopurs_runtime.Value {
	once_pure__3197665392.Do(func() {
		cache_pure__3197665392 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3197665392(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3197665392
}

var cache_pure__3820067664 gopurs_runtime.Value
var once_pure__3820067664 sync.Once
func Get_pure__3820067664() gopurs_runtime.Value {
	once_pure__3820067664.Do(func() {
		cache_pure__3820067664 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3820067664(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3820067664
}

var cache_pure__1670386480 gopurs_runtime.Value
var once_pure__1670386480 sync.Once
func Get_pure__1670386480() gopurs_runtime.Value {
	once_pure__1670386480.Do(func() {
		cache_pure__1670386480 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1670386480(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1670386480
}

var cache_pure__1715998582 gopurs_runtime.Value
var once_pure__1715998582 sync.Once
func Get_pure__1715998582() gopurs_runtime.Value {
	once_pure__1715998582.Do(func() {
		cache_pure__1715998582 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1715998582(a_0_box)
})
	})
	return cache_pure__1715998582
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

var cache_bind__3227627207 gopurs_runtime.Value
var once_bind__3227627207 sync.Once
func Get_bind__3227627207() gopurs_runtime.Value {
	once_bind__3227627207.Do(func() {
		cache_bind__3227627207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3227627207(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3227627207
}

var cache_bind__2892370023 gopurs_runtime.Value
var once_bind__2892370023 sync.Once
func Get_bind__2892370023() gopurs_runtime.Value {
	once_bind__2892370023.Do(func() {
		cache_bind__2892370023 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2892370023(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2892370023
}

var cache_bind__2931166727 gopurs_runtime.Value
var once_bind__2931166727 sync.Once
func Get_bind__2931166727() gopurs_runtime.Value {
	once_bind__2931166727.Do(func() {
		cache_bind__2931166727 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2931166727(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2931166727
}

var cache_bind__991288455 gopurs_runtime.Value
var once_bind__991288455 sync.Once
func Get_bind__991288455() gopurs_runtime.Value {
	once_bind__991288455.Do(func() {
		cache_bind__991288455 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__991288455(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__991288455
}

var cache_bind__3078669415 gopurs_runtime.Value
var once_bind__3078669415 sync.Once
func Get_bind__3078669415() gopurs_runtime.Value {
	once_bind__3078669415.Do(func() {
		cache_bind__3078669415 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3078669415(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3078669415
}

var cache_bind__961482919 gopurs_runtime.Value
var once_bind__961482919 sync.Once
func Get_bind__961482919() gopurs_runtime.Value {
	once_bind__961482919.Do(func() {
		cache_bind__961482919 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__961482919(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__961482919
}

var cache_cons__808523158 gopurs_runtime.Value
var once_cons__808523158 sync.Once
func Get_cons__808523158() gopurs_runtime.Value {
	once_cons__808523158.Do(func() {
		cache_cons__808523158 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__808523158(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), lh_1_box, t_2_box)
})
	})
	return cache_cons__808523158
}

var cache_drop__1964165395 gopurs_runtime.Value
var once_drop__1964165395 sync.Once
func Get_drop__1964165395() gopurs_runtime.Value {
	once_drop__1964165395.Do(func() {
		cache_drop__1964165395 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop__1964165395(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_drop__1964165395
}

var cache_dropWhile__504781836 gopurs_runtime.Value
var once_dropWhile__504781836 sync.Once
func Get_dropWhile__504781836() gopurs_runtime.Value {
	once_dropWhile__504781836.Do(func() {
		cache_dropWhile__504781836 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile__504781836(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_dropWhile__504781836
}

var cache_filter__1345510683 gopurs_runtime.Value
var once_filter__1345510683 sync.Once
func Get_filter__1345510683() gopurs_runtime.Value {
	once_filter__1345510683.Do(func() {
		cache_filter__1345510683 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter__1345510683(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_filter__1345510683
}

var cache_foldl_prime__3412851976 gopurs_runtime.Value
var once_foldl_prime__3412851976 sync.Once
func Get_foldl_prime__3412851976() gopurs_runtime.Value {
	once_foldl_prime__3412851976.Do(func() {
		cache_foldl_prime__3412851976 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl_prime__3412851976(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldl_prime__3412851976
}

var cache_foldl_prime__2387145256 gopurs_runtime.Value
var once_foldl_prime__2387145256 sync.Once
func Get_foldl_prime__2387145256() gopurs_runtime.Value {
	once_foldl_prime__2387145256.Do(func() {
		cache_foldl_prime__2387145256 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl_prime__2387145256(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldl_prime__2387145256
}

var cache_foldlRec_prime__4148996870 gopurs_runtime.Value
var once_foldlRec_prime__4148996870 sync.Once
func Get_foldlRec_prime__4148996870() gopurs_runtime.Value {
	once_foldlRec_prime__4148996870.Do(func() {
		cache_foldlRec_prime__4148996870 = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec_prime__4148996870(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_foldlRec_prime__4148996870
}

var cache_foldlRec_prime__1739794342 gopurs_runtime.Value
var once_foldlRec_prime__1739794342 sync.Once
func Get_foldlRec_prime__1739794342() gopurs_runtime.Value {
	once_foldlRec_prime__1739794342.Do(func() {
		cache_foldlRec_prime__1739794342 = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec_prime__1739794342(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_foldlRec_prime__1739794342
}

var cache_iterate__4162284821 gopurs_runtime.Value
var once_iterate__4162284821 sync.Once
func Get_iterate__4162284821() gopurs_runtime.Value {
	once_iterate__4162284821.Do(func() {
		cache_iterate__4162284821 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate__4162284821(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_iterate__4162284821
}

var cache_mapMaybe__3319479893 gopurs_runtime.Value
var once_mapMaybe__3319479893 sync.Once
func Get_mapMaybe__3319479893() gopurs_runtime.Value {
	once_mapMaybe__3319479893.Do(func() {
		cache_mapMaybe__3319479893 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__3319479893(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_mapMaybe__3319479893
}

var cache_mapMaybe__3325666580 gopurs_runtime.Value
var once_mapMaybe__3325666580 sync.Once
func Get_mapMaybe__3325666580() gopurs_runtime.Value {
	once_mapMaybe__3325666580.Do(func() {
		cache_mapMaybe__3325666580 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__3325666580(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_mapMaybe__3325666580
}

var cache_nil__1472516796 gopurs_runtime.Value
var once_nil__1472516796 sync.Once
func Get_nil__1472516796() gopurs_runtime.Value {
	once_nil__1472516796.Do(func() {
		cache_nil__1472516796 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nil__1472516796(dictApplicative_0_box)
})
	})
	return cache_nil__1472516796
}

var cache_prepend__2860458454 gopurs_runtime.Value
var once_prepend__2860458454 sync.Once
func Get_prepend__2860458454() gopurs_runtime.Value {
	once_prepend__2860458454.Do(func() {
		cache_prepend__2860458454 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prepend__2860458454(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_prepend__2860458454
}

var cache_prepend_prime__1901723831 gopurs_runtime.Value
var once_prepend_prime__1901723831 sync.Once
func Get_prepend_prime__1901723831() gopurs_runtime.Value {
	once_prepend_prime__1901723831.Do(func() {
		cache_prepend_prime__1901723831 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prepend_prime__1901723831(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), h_1_box, t_2_box)
})
	})
	return cache_prepend_prime__1901723831
}

var cache_singleton__2427543124 gopurs_runtime.Value
var once_singleton__2427543124 sync.Once
func Get_singleton__2427543124() gopurs_runtime.Value {
	once_singleton__2427543124.Do(func() {
		cache_singleton__2427543124 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton__2427543124(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_singleton__2427543124
}

var cache_stepMap__3249590196 gopurs_runtime.Value
var once_stepMap__3249590196 sync.Once
func Get_stepMap__3249590196() gopurs_runtime.Value {
	once_stepMap__3249590196.Do(func() {
		cache_stepMap__3249590196 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stepMap__3249590196(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_stepMap__3249590196
}

var cache_take__1964165395 gopurs_runtime.Value
var once_take__1964165395 sync.Once
func Get_take__1964165395() gopurs_runtime.Value {
	once_take__1964165395.Do(func() {
		cache_take__1964165395 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take__1964165395(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_take__1964165395
}

var cache_takeWhile__504781836 gopurs_runtime.Value
var once_takeWhile__504781836 sync.Once
func Get_takeWhile__504781836() gopurs_runtime.Value {
	once_takeWhile__504781836.Do(func() {
		cache_takeWhile__504781836 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile__504781836(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_takeWhile__504781836
}

var cache_uncons__1307401241 gopurs_runtime.Value
var once_uncons__1307401241 sync.Once
func Get_uncons__1307401241() gopurs_runtime.Value {
	once_uncons__1307401241.Do(func() {
		cache_uncons__1307401241 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons__1307401241(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_uncons__1307401241
}

var cache_unfold__3487137686 gopurs_runtime.Value
var once_unfold__3487137686 sync.Once
func Get_unfold__3487137686() gopurs_runtime.Value {
	once_unfold__3487137686.Do(func() {
		cache_unfold__3487137686 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfold__3487137686(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_unfold__3487137686
}

var cache_unfold__2471180757 gopurs_runtime.Value
var once_unfold__2471180757 sync.Once
func Get_unfold__2471180757() gopurs_runtime.Value {
	once_unfold__2471180757.Do(func() {
		cache_unfold__2471180757 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfold__2471180757(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_unfold__2471180757
}

var cache_wrapEffect__3965193927 gopurs_runtime.Value
var once_wrapEffect__3965193927 sync.Once
func Get_wrapEffect__3965193927() gopurs_runtime.Value {
	once_wrapEffect__3965193927.Do(func() {
		cache_wrapEffect__3965193927 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrapEffect__3965193927(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
})
	})
	return cache_wrapEffect__3965193927
}

var cache_zipWith_prime__376166203 gopurs_runtime.Value
var once_zipWith_prime__376166203 sync.Once
func Get_zipWith_prime__376166203() gopurs_runtime.Value {
	once_zipWith_prime__376166203.Do(func() {
		cache_zipWith_prime__376166203 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith_prime__376166203(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_zipWith_prime__376166203
}

var cache_tailRecM__3865988408 gopurs_runtime.Value
var once_tailRecM__3865988408 sync.Once
func Get_tailRecM__3865988408() gopurs_runtime.Value {
	once_tailRecM__3865988408.Do(func() {
		cache_tailRecM__3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3865988408(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3865988408
}

var cache_tailRecM__1444729948 gopurs_runtime.Value
var once_tailRecM__1444729948 sync.Once
func Get_tailRecM__1444729948() gopurs_runtime.Value {
	once_tailRecM__1444729948.Do(func() {
		cache_tailRecM__1444729948 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__1444729948(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__1444729948
}

var cache_tailRecM2__1943630176 gopurs_runtime.Value
var once_tailRecM2__1943630176 sync.Once
func Get_tailRecM2__1943630176() gopurs_runtime.Value {
	once_tailRecM2__1943630176.Do(func() {
		cache_tailRecM2__1943630176 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2__1943630176(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_tailRecM2__1943630176
}

var cache_tailRecM2__2551820843 gopurs_runtime.Value
var once_tailRecM2__2551820843 sync.Once
func Get_tailRecM2__2551820843() gopurs_runtime.Value {
	once_tailRecM2__2551820843.Do(func() {
		cache_tailRecM2__2551820843 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2__2551820843(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_tailRecM2__2551820843
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

var cache_const__4181451586 gopurs_runtime.Value
var once_const__4181451586 sync.Once
func Get_const__4181451586() gopurs_runtime.Value {
	once_const__4181451586.Do(func() {
		cache_const__4181451586 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4181451586(a_0_box, v_1_box)
})
	})
	return cache_const__4181451586
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

var cache_const__3952683620 gopurs_runtime.Value
var once_const__3952683620 sync.Once
func Get_const__3952683620() gopurs_runtime.Value {
	once_const__3952683620.Do(func() {
		cache_const__3952683620 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__3952683620(a_0_box, v_1_box)
})
	})
	return cache_const__3952683620
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

var cache_flip__1117087808 gopurs_runtime.Value
var once_flip__1117087808 sync.Once
func Get_flip__1117087808() gopurs_runtime.Value {
	once_flip__1117087808.Do(func() {
		cache_flip__1117087808 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1117087808(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1117087808
}

var cache_flip__3675729664 gopurs_runtime.Value
var once_flip__3675729664 sync.Once
func Get_flip__3675729664() gopurs_runtime.Value {
	once_flip__3675729664.Do(func() {
		cache_flip__3675729664 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3675729664(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3675729664
}

var cache_flip__3858636736 gopurs_runtime.Value
var once_flip__3858636736 sync.Once
func Get_flip__3858636736() gopurs_runtime.Value {
	once_flip__3858636736.Do(func() {
		cache_flip__3858636736 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3858636736(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3858636736
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

var cache_map__1668665428 gopurs_runtime.Value
var once_map__1668665428 sync.Once
func Get_map__1668665428() gopurs_runtime.Value {
	once_map__1668665428.Do(func() {
		cache_map__1668665428 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1668665428(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1668665428
}

var cache_map__3674493396 gopurs_runtime.Value
var once_map__3674493396 sync.Once
func Get_map__3674493396() gopurs_runtime.Value {
	once_map__3674493396.Do(func() {
		cache_map__3674493396 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3674493396(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3674493396
}

var cache_map__2322598548 gopurs_runtime.Value
var once_map__2322598548 sync.Once
func Get_map__2322598548() gopurs_runtime.Value {
	once_map__2322598548.Do(func() {
		cache_map__2322598548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2322598548(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2322598548
}

var cache_map__2753776532 gopurs_runtime.Value
var once_map__2753776532 sync.Once
func Get_map__2753776532() gopurs_runtime.Value {
	once_map__2753776532.Do(func() {
		cache_map__2753776532 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2753776532(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2753776532
}

var cache_map__657998836 gopurs_runtime.Value
var once_map__657998836 sync.Once
func Get_map__657998836() gopurs_runtime.Value {
	once_map__657998836.Do(func() {
		cache_map__657998836 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__657998836(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__657998836
}

var cache_map__528096244 gopurs_runtime.Value
var once_map__528096244 sync.Once
func Get_map__528096244() gopurs_runtime.Value {
	once_map__528096244.Do(func() {
		cache_map__528096244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__528096244(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__528096244
}

var cache_map__3228596244 gopurs_runtime.Value
var once_map__3228596244 sync.Once
func Get_map__3228596244() gopurs_runtime.Value {
	once_map__3228596244.Do(func() {
		cache_map__3228596244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3228596244(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3228596244
}

var cache_map__1729183892 gopurs_runtime.Value
var once_map__1729183892 sync.Once
func Get_map__1729183892() gopurs_runtime.Value {
	once_map__1729183892.Do(func() {
		cache_map__1729183892 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1729183892(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1729183892
}

var cache_map__3124798356 gopurs_runtime.Value
var once_map__3124798356 sync.Once
func Get_map__3124798356() gopurs_runtime.Value {
	once_map__3124798356.Do(func() {
		cache_map__3124798356 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3124798356(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3124798356
}

var cache_map__596534652 gopurs_runtime.Value
var once_map__596534652 sync.Once
func Get_map__596534652() gopurs_runtime.Value {
	once_map__596534652.Do(func() {
		cache_map__596534652 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__596534652(f_0_box, l_1_box)
})
	})
	return cache_map__596534652
}

var cache_map__2275717084 gopurs_runtime.Value
var once_map__2275717084 sync.Once
func Get_map__2275717084() gopurs_runtime.Value {
	once_map__2275717084.Do(func() {
		cache_map__2275717084 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__2275717084(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__2275717084
}

var cache_map__2615158204 gopurs_runtime.Value
var once_map__2615158204 sync.Once
func Get_map__2615158204() gopurs_runtime.Value {
	once_map__2615158204.Do(func() {
		cache_map__2615158204 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__2615158204(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](v1_1_box)))}
})
	})
	return cache_map__2615158204
}

var cache_map__125648636 gopurs_runtime.Value
var once_map__125648636 sync.Once
func Get_map__125648636() gopurs_runtime.Value {
	once_map__125648636.Do(func() {
		cache_map__125648636 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__125648636(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](v1_1_box)))}
})
	})
	return cache_map__125648636
}

var cache_applicativeLazy__690919725 gopurs_runtime.Value
var once_applicativeLazy__690919725 sync.Once
func Get_applicativeLazy__690919725() gopurs_runtime.Value {
	once_applicativeLazy__690919725.Do(func() {
		cache_applicativeLazy__690919725 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Lazy.Get_applyLazy()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}))
	})
	return cache_applicativeLazy__690919725
}

var cache_applyLazy__225241115 gopurs_runtime.Value
var once_applyLazy__225241115 sync.Once
func Get_applyLazy__225241115() gopurs_runtime.Value {
	once_applyLazy__225241115.Do(func() {
		cache_applyLazy__225241115 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Lazy.Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Lazy.Get_force(), f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
}))
})
}))
	})
	return cache_applyLazy__225241115
}

var cache_functorLazy__491347738 gopurs_runtime.Value
var once_functorLazy__491347738 sync.Once
func Get_functorLazy__491347738() gopurs_runtime.Value {
	once_functorLazy__491347738.Do(func() {
		cache_functorLazy__491347738 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy__491347738
}

var cache_functorLazy__3988504945 gopurs_runtime.Value
var once_functorLazy__3988504945 sync.Once
func Get_functorLazy__3988504945() gopurs_runtime.Value {
	once_functorLazy__3988504945.Do(func() {
		cache_functorLazy__3988504945 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy__3988504945
}

var cache_fromMaybe__430429096 gopurs_runtime.Value
var once_fromMaybe__430429096 sync.Once
func Get_fromMaybe__430429096() gopurs_runtime.Value {
	once_fromMaybe__430429096.Do(func() {
		cache_fromMaybe__430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__430429096(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__430429096
}

var cache_fromMaybe__656947263 gopurs_runtime.Value
var once_fromMaybe__656947263 sync.Once
func Get_fromMaybe__656947263() gopurs_runtime.Value {
	once_fromMaybe__656947263.Do(func() {
		cache_fromMaybe__656947263 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__656947263(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__656947263
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_append__988370296 gopurs_runtime.Value
var once_append__988370296 sync.Once
func Get_append__988370296() gopurs_runtime.Value {
	once_append__988370296.Do(func() {
		cache_append__988370296 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__988370296(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__988370296
}

var cache_fst__2554656696 gopurs_runtime.Value
var once_fst__2554656696 sync.Once
func Get_fst__2554656696() gopurs_runtime.Value {
	once_fst__2554656696.Do(func() {
		cache_fst__2554656696 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__2554656696(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__2554656696
}

var cache_snd__4038973427 gopurs_runtime.Value
var once_snd__4038973427 sync.Once
func Get_snd__4038973427() gopurs_runtime.Value {
	once_snd__4038973427.Do(func() {
		cache_snd__4038973427 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__4038973427(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__4038973427
}

type Constructor_Yield[T_a any, T_s any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}


type Constructor_Skip[T_a any, T_s any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Done[T_a any, T_s any] struct {
	Rc uint32
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var x_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_ListT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_wrapLazy(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_1})})
}

func Call_wrapEffect(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))})}
}), v_1)
}

func Call_unfold(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
unfold:
for {
if false { continue unfold }
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
__local_var_5_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_unfold(dictMonad_0), f_2, __local_var_5_1)
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

func Call_uncons(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
uncons:
for {
if false { continue uncons }
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure")
_ = pure_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_2.V1, v_4, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1320412129) {
__t3 = gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}})})
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 813447293) {
__t3 = gopurs_runtime.Apply(Call_uncons(dictMonad_0), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 489128924) {
__t3 = gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
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

func Call_tail(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_snd()), gopurs_runtime.Apply(Call_uncons(dictMonad_0), l_2))
})
}

func Call_takeWhile(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_takeWhile(dictApplicative_0), f_2), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_takeWhile(dictApplicative_0), f_2), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}
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

func Call_scanl(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_unfold(dictMonad_0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_1
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 1320412129) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_2, __local_var_6_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1)})}, __local_var_6_1})}})}
goto end_branch_2
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0)})}, __local_var_6_1})}})}
goto end_branch_2
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 489128924) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](__t2))}
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_3, l_4})})
})
})
})
}

func Call_prepend_prime(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_1, t_2})})
}

func Call_prepend(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return t_2
}))})})
}

func Call_nil(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_singleton(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}))})})
})
}

func Call_take(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
take:
for {
if false { continue take }
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
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
__t3 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1320412129) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_take(dictApplicative_0), gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(v_3.IntVal), gopurs_runtime.Int(1)).IntVal)), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_take(dictApplicative_0), gopurs_runtime.Int(v_3.IntVal)), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)})}
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

func Call_zipWith_prime(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
zipWith_prime:
for {
if false { continue zipWith_prime }
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
__local_var_3_3 := gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{})
_ = __local_var_3_3
prepend_prime1_3_2 := gopurs_runtime.Func(func(h_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_4, t_5})})
})
})
_ = prepend_prime1_3_2
Bind1_4_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fa_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fb_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
}))})}
}), gopurs_runtime.Apply2(Bind1_4_4.V1, gopurs_runtime.Apply(Call_uncons(dictMonad_0), fa_6), gopurs_runtime.Func(func(ua_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_4_4.V1, gopurs_runtime.Apply(Call_uncons(dictMonad_0), fb_7), gopurs_runtime.Func(func(ub_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ub_9))}
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136 && __t_tag_5.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ua_8))}
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 930809136 && __t_tag_6.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ua_8))}
var __t_and_9 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 930809136 && __t_tag_7.UnsafePtr != nil) {

var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ub_9))}
__t_and_9 = (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 930809136 && __t_tag_8.UnsafePtr != nil)
}
if __t_and_9 {
__local_var_10_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_8.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_10_10
__local_var_11_11 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_9.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_11
__local_var_12_12 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_zipWith_prime(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_5, __local_var_10_10, __local_var_11_11)
}))
_ = __local_var_12_12
__t13 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(prepend_prime1_3_2, a_13, __local_var_12_12)
}), gopurs_runtime.Apply2(f_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_8.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_9.UnsafePtr).V0.UnsafePtr).V0))
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

func Call_zipWith(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_zipWith_prime(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Apply2(f_2, a_3, b_4))
})
}))
})
}

func Call_mapMaybe(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__local_var_4_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Yield(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)))}))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr == nil) {
__t1 = Get_Skip()
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Apply(__t1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_iterate(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_unfold(dictMonad_0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, x_4), x_4})}})})
}), a_3)
})
})
}

func Call_repeat(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply(Call_iterate(dictMonad_0), Get_identity())
}

func Call_head(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_fst()), gopurs_runtime.Apply(Call_uncons(dictMonad_0), l_2))
})
}

func Call_functorListT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_functorListT(dictFunctor_0), "map"), f_1), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_functorListT(dictFunctor_0), "map"), f_1), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_fromEffect(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}))
_ = __local_var_3_1
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, __local_var_3_1})}
}), fa_2)
})
}

func Call_foldlRec_prime(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
Monad01_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(Call_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_9_4})})
goto end_branch_6
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
__local_var_11_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_5
__t6 = gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply2(f_5, __local_var_9_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_5)})})
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

func Call_runListTRec(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Apply3(Get_foldlRec_prime(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_Unit.Get_unit())
})
}), pkg_Data_Unit.Get_unit())
}

func Call_foldlRec(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
Monad01_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(Call_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_9_4})})
goto end_branch_5
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Apply2(f_5, __local_var_9_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V1)})})
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

func Call_foldl_prime(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_0 gopurs_runtime.Value
_ = loop_4_2_0
loop_4_2_0 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(Call_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, b_5)
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__local_var_8_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_3
__t4 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(f_3, b_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_runListT(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Apply3(Get_foldl_prime(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_Unit.Get_unit())
})
}), pkg_Data_Unit.Get_unit())
}

func Call_foldl(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_1 gopurs_runtime.Value
_ = loop_4_2_1
loop_4_2_1 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(Call_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(Applicative0_1_0.V1, b_5)
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(loop_4_2_1, gopurs_runtime.Apply2(f_3, b_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1)
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

func Call_filter(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
s_prime_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
_ = s_prime_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_prime_4_0})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_prime_4_0})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_dropWhile(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
dropWhile:
for {
if false { continue dropWhile }
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_dropWhile(dictApplicative_0), f_2), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_dropWhile(dictApplicative_0), f_2), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}
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

func Call_drop(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
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
__t2 = gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_drop(dictApplicative_0), gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(v_2.IntVal), gopurs_runtime.Int(1)).IntVal)), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_drop(dictApplicative_0), gopurs_runtime.Int(v_2.IntVal)), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)})}
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

func Call_cons(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], lh_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 gopurs_runtime.Value = lh_1_loop
_ = lh_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), lh_1), t_2})})
}

func Call_unfoldable1ListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_2 gopurs_runtime.Value
_ = go__go_4_1_2
go__go_4_1_2 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(Call_singleton(Applicative0_1_0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__local_var_6_4 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1.UnsafePtr).V0
_ = __local_var_6_4
__t5 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_4_1_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_2, __local_var_6_4)))})
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
return gopurs_runtime.Apply(go__go_4_1_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_2, b_3)))})
})
}))
}

func Call_unfoldableListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
unfoldable1ListT1_2_1 := Call_unfoldable1ListT(dictMonad_0)
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
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
goto end_branch_4
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 930809136 && v_6.UnsafePtr != nil) {
__local_var_7_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0.UnsafePtr).V0)), gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_5_2_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_3, __local_var_7_3)))})
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
return gopurs_runtime.Apply(go__go_5_2_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_3, b_4)))})
})
}))
}

func Call_semigroupListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.RecordDict1("append", Call_concat(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0)))
}

func Call_concat(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_concat(dictApplicative_0), v1_5, y_3)
}), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_concat(dictApplicative_0), v1_5, y_3)
}), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_monoidListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
semigroupListT1_1_0 := gopurs_runtime.RecordDict1("append", Call_concat(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0)))
_ = semigroupListT1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupListT1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
}

func Call_catMaybes(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, Get_identity1())
}

func Call_monadListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeListT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindListT(dictMonad_0)
}))
}

func Call_bindListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
semigroupListT1_1_0 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, Call_concat(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))}
_ = semigroupListT1_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyListT(dictMonad_0)
}), gopurs_runtime.Func(func(fa_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1320412129) {
__local_var_6_2 := (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(semigroupListT1_1_0.V0, gopurs_runtime.Apply(f_4, __local_var_6_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_bindListT(dictMonad_0), "bind"), s_prime_7, f_4))
}), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)})}
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 813447293) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_bindListT(dictMonad_0), "bind"), v1_6, f_4)
}), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)})}
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

func Call_applyListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
functorListT1_1_0 := Call_functorListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeListT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindListT(dictMonad_0)
}))
_ = __local_var_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_2.V1, f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_2.V1, a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_4_3.V1, gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_applicativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyListT(dictMonad_0)
}), Call_singleton(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))))
}

func Call_monadEffectListT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyListT(Monad0_1_0)
}), Call_singleton(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindListT(Monad0_1_0)
}))
_ = monadListT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransListT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_monadSTListT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyListT(Monad0_1_0)
}), Call_singleton(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindListT(Monad0_1_0)
}))
_ = monadListT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransListT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_altListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
functorListT1_1_0 := Call_functorListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), Call_concat(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0)))
}

func Call_plusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
altListT1_2_1 := Call_altListT(Applicative0_1_0)
_ = altListT1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altListT1_2_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
}

func Call_alternativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeListT1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyListT(dictMonad_0)
}), Call_singleton(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))))
_ = applicativeListT1_1_0
plusListT1_2_1 := Call_plusListT(dictMonad_0)
_ = plusListT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeListT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusListT1_2_1
}))
}

func Call_monadPlusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadListT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyListT(dictMonad_0)
}), Call_singleton(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindListT(dictMonad_0)
}))
_ = monadListT1_1_0
alternativeListT1_2_1 := Call_alternativeListT(dictMonad_0)
_ = alternativeListT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeListT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_1_0
}))
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

func Call_pure__160425008(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3197665392(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3820067664(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1670386480(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1715998582(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3227627207(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2892370023(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2931166727(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__991288455(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3078669415(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__961482919(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_cons__808523158(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], lh_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 gopurs_runtime.Value = lh_1_loop
_ = lh_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), lh_1), t_2})})
}

func Call_drop__1964165395(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
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
__t2 = gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_drop(dictApplicative_0), gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(v_2.IntVal), gopurs_runtime.Int(1)).IntVal)), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_drop(dictApplicative_0), gopurs_runtime.Int(v_2.IntVal)), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)})}
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

func Call_dropWhile__504781836(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_dropWhile(dictApplicative_0), f_2), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_dropWhile(dictApplicative_0), f_2), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}
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

func Call_filter__1345510683(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
s_prime_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
_ = s_prime_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_prime_4_0})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_prime_4_0})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_filter(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_foldl_prime__3412851976(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_4 gopurs_runtime.Value
_ = loop_4_2_4
loop_4_2_4 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(Call_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, b_5)
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__local_var_8_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_3
__t4 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(f_3, b_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_foldl_prime__2387145256(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2_5 gopurs_runtime.Value
_ = loop_4_2_5
loop_4_2_5 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(Call_uncons(dictMonad_0), l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, b_5)
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__local_var_8_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_3
__t4 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(f_3, b_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_foldlRec_prime__4148996870(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
Monad01_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(Call_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_9_4})})
goto end_branch_6
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
__local_var_11_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_5
__t6 = gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply2(f_5, __local_var_9_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_5)})})
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

func Call_foldlRec_prime__1739794342(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
Monad01_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}))
_ = Monad01_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(Call_uncons(Monad01_4_3), gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_9_4})})
goto end_branch_6
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
__local_var_11_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_5
__t6 = gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply2(f_5, __local_var_9_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_5)})})
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

func Call_iterate__4162284821(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_unfold(dictMonad_0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, x_4), x_4})}})})
}), a_3)
})
})
}

func Call_mapMaybe__3319479893(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__local_var_4_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Yield(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)))}))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr == nil) {
__t1 = Get_Skip()
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Apply(__t1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_mapMaybe__3325666580(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__local_var_4_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Yield(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)))}))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr == nil) {
__t1 = Get_Skip()
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Apply(__t1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_nil__1472516796(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_prepend__2860458454(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return t_2
}))})})
}

func Call_prepend_prime__1901723831(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_1, t_2})})
}

func Call_singleton__2427543124(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}))})})
})
}

func Call_stepMap__3249590196(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_1, v_2)
}

func Call_take__1964165395(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)})
_ = nil1_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
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
__t3 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1320412129) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_take(dictApplicative_0), gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(v_3.IntVal), gopurs_runtime.Int(1)).IntVal)), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_take(dictApplicative_0), gopurs_runtime.Int(v_3.IntVal)), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)})}
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

func Call_takeWhile__504781836(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_takeWhile(dictApplicative_0), f_2), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_takeWhile(dictApplicative_0), f_2), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}
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

func Call_uncons__1307401241(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure")
_ = pure_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_2.V1, v_4, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1320412129) {
__t3 = gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}})})
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 813447293) {
__t3 = gopurs_runtime.Apply(Call_uncons(dictMonad_0), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 489128924) {
__t3 = gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
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

func Call_unfold__3487137686(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
__local_var_5_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_unfold(dictMonad_0), f_2, __local_var_5_1)
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

func Call_unfold__2471180757(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
__local_var_5_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_5_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_unfold(dictMonad_0), f_2, __local_var_5_1)
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

func Call_wrapEffect__3965193927(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))})}
}), v_1)
}

func Call_zipWith_prime__376166203(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
__local_var_3_3 := gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{})
_ = __local_var_3_3
prepend_prime1_3_2 := gopurs_runtime.Func(func(h_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_4, t_5})})
})
})
_ = prepend_prime1_3_2
Bind1_4_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fa_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fb_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
}))})}
}), gopurs_runtime.Apply2(Bind1_4_4.V1, gopurs_runtime.Apply(Call_uncons(dictMonad_0), fa_6), gopurs_runtime.Func(func(ua_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_4_4.V1, gopurs_runtime.Apply(Call_uncons(dictMonad_0), fb_7), gopurs_runtime.Func(func(ub_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ub_9))}
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136 && __t_tag_5.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ua_8))}
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 930809136 && __t_tag_6.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: unsafe.Pointer(nil)}))
goto end_branch_13
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ua_8))}
var __t_and_9 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 930809136 && __t_tag_7.UnsafePtr != nil) {

var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ub_9))}
__t_and_9 = (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 930809136 && __t_tag_8.UnsafePtr != nil)
}
if __t_and_9 {
__local_var_10_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_8.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_10_10
__local_var_11_11 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_9.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_11
__local_var_12_12 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_zipWith_prime(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_5, __local_var_10_10, __local_var_11_11)
}))
_ = __local_var_12_12
__t13 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(prepend_prime1_3_2, a_13, __local_var_12_12)
}), gopurs_runtime.Apply2(f_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_8.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_9.UnsafePtr).V0.UnsafePtr).V0))
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

func Call_tailRecM__3865988408(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__1444729948(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM2__1943630176(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_tailRecM2__2551820843(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_const__4181451586(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_const__3952683620(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_flip__1117087808(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3675729664(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3858636736(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_map__1668665428(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3674493396(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2322598548(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2753776532(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__657998836(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__528096244(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3228596244(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1729183892(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3124798356(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__596534652(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
}

func Call_map__2275717084(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_map__2615158204(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0))})})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_map__125648636(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0))})})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_fromMaybe__430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromMaybe__656947263(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__988370296(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fst__2554656696(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_snd__4038973427(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}


