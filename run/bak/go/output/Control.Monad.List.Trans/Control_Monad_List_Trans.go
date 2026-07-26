package Control_Monad_List_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
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

var cache_Yield gopurs_runtime.Value
var once_Yield sync.Once
func Get_Yield() gopurs_runtime.Value {
	once_Yield.Do(func() {
		cache_Yield = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{value0})}
})
	})
	return cache_Skip
}

var cache_Done gopurs_runtime.Value
var once_Done sync.Once
func Get_Done() gopurs_runtime.Value {
	once_Done.Do(func() {
		cache_Done = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
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
return Call_wrapLazy((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr), v_1_box)
})
	})
	return cache_wrapLazy
}

var cache_wrapEffect gopurs_runtime.Value
var once_wrapEffect sync.Once
func Get_wrapEffect() gopurs_runtime.Value {
	once_wrapEffect.Do(func() {
		cache_wrapEffect = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrapEffect((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr), v_1_box)
})
	})
	return cache_wrapEffect
}

var cache_unfold gopurs_runtime.Value
var once_unfold sync.Once
func Get_unfold() gopurs_runtime.Value {
	once_unfold.Do(func() {
		cache_unfold = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfold((*Record_)(dictMonad_0_box.UnsafePtr), f_1_box, z_2_box)
})
	})
	return cache_unfold
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_uncons
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_tail
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_takeWhile
}

var cache_scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		cache_scanl = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanl((*Record_)(dictMonad_0_box.UnsafePtr), f_1_box, b_2_box, l_3_box)
})
	})
	return cache_scanl
}

var cache_prepend_prime gopurs_runtime.Value
var once_prepend_prime sync.Once
func Get_prepend_prime() gopurs_runtime.Value {
	once_prepend_prime.Do(func() {
		cache_prepend_prime = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prepend_prime((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr), h_1_box, t_2_box)
})
	})
	return cache_prepend_prime
}

var cache_prepend gopurs_runtime.Value
var once_prepend sync.Once
func Get_prepend() gopurs_runtime.Value {
	once_prepend.Do(func() {
		cache_prepend = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prepend((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr), h_1_box, t_2_box)
})
	})
	return cache_prepend
}

var cache_nil gopurs_runtime.Value
var once_nil sync.Once
func Get_nil() gopurs_runtime.Value {
	once_nil.Do(func() {
		cache_nil = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nil((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_nil
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_singleton
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_take
}

var cache_zipWith_prime gopurs_runtime.Value
var once_zipWith_prime sync.Once
func Get_zipWith_prime() gopurs_runtime.Value {
	once_zipWith_prime.Do(func() {
		cache_zipWith_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith_prime((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_zipWith_prime
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith((*Record_)(dictMonad_0_box.UnsafePtr))
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
return Call_mapMaybe((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr), f_1_box, v_2_box)
})
	})
	return cache_mapMaybe
}

var cache_iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		cache_iterate = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate((*Record_)(dictMonad_0_box.UnsafePtr), f_1_box, a_2_box)
})
	})
	return cache_iterate
}

var cache_repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		cache_repeat = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_repeat
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_head
}

var cache_functorListT gopurs_runtime.Value
var once_functorListT sync.Once
func Get_functorListT() gopurs_runtime.Value {
	once_functorListT.Do(func() {
		cache_functorListT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorListT((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr))
})
	})
	return cache_functorListT
}

var cache_fromEffect gopurs_runtime.Value
var once_fromEffect sync.Once
func Get_fromEffect() gopurs_runtime.Value {
	once_fromEffect.Do(func() {
		cache_fromEffect = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEffect((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_fromEffect
}

var cache_monadTransListT gopurs_runtime.Value
var once_monadTransListT sync.Once
func Get_monadTransListT() gopurs_runtime.Value {
	once_monadTransListT.Do(func() {
		cache_monadTransListT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
return gopurs_runtime.Func(func(fa_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{a_4, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_2_1
})})}
}), fa_3)
})
}))
	})
	return cache_monadTransListT
}

var cache_foldlRec_prime gopurs_runtime.Value
var once_foldlRec_prime sync.Once
func Get_foldlRec_prime() gopurs_runtime.Value {
	once_foldlRec_prime.Do(func() {
		cache_foldlRec_prime = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec_prime((*Record_tailRecM_gopurs_runtime_Value)(dictMonadRec_0_box.UnsafePtr))
})
	})
	return cache_foldlRec_prime
}

var cache_runListTRec gopurs_runtime.Value
var once_runListTRec sync.Once
func Get_runListTRec() gopurs_runtime.Value {
	once_runListTRec.Do(func() {
		cache_runListTRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runListTRec((*Record_tailRecM_gopurs_runtime_Value)(dictMonadRec_0_box.UnsafePtr))
})
	})
	return cache_runListTRec
}

var cache_foldlRec gopurs_runtime.Value
var once_foldlRec sync.Once
func Get_foldlRec() gopurs_runtime.Value {
	once_foldlRec.Do(func() {
		cache_foldlRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec((*Record_tailRecM_gopurs_runtime_Value)(dictMonadRec_0_box.UnsafePtr))
})
	})
	return cache_foldlRec
}

var cache_foldl_prime gopurs_runtime.Value
var once_foldl_prime sync.Once
func Get_foldl_prime() gopurs_runtime.Value {
	once_foldl_prime.Do(func() {
		cache_foldl_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl_prime((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_foldl_prime
}

var cache_runListT gopurs_runtime.Value
var once_runListT sync.Once
func Get_runListT() gopurs_runtime.Value {
	once_runListT.Do(func() {
		cache_runListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_runListT
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_foldl
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr), f_1_box, v_2_box)
})
	})
	return cache_filter
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_dropWhile
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_drop
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr), lh_1_box, t_2_box)
})
	})
	return cache_cons
}

var cache_unfoldable1ListT gopurs_runtime.Value
var once_unfoldable1ListT sync.Once
func Get_unfoldable1ListT() gopurs_runtime.Value {
	once_unfoldable1ListT.Do(func() {
		cache_unfoldable1ListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable1ListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_unfoldable1ListT
}

var cache_unfoldableListT gopurs_runtime.Value
var once_unfoldableListT sync.Once
func Get_unfoldableListT() gopurs_runtime.Value {
	once_unfoldableListT.Do(func() {
		cache_unfoldableListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldableListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_unfoldableListT
}

var cache_semigroupListT gopurs_runtime.Value
var once_semigroupListT sync.Once
func Get_semigroupListT() gopurs_runtime.Value {
	once_semigroupListT.Do(func() {
		cache_semigroupListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupListT((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_semigroupListT
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concat((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_concat
}

var cache_monoidListT gopurs_runtime.Value
var once_monoidListT sync.Once
func Get_monoidListT() gopurs_runtime.Value {
	once_monoidListT.Do(func() {
		cache_monoidListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidListT((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_monoidListT
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catMaybes((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr))
})
	})
	return cache_catMaybes
}

var cache_monadListT gopurs_runtime.Value
var once_monadListT sync.Once
func Get_monadListT() gopurs_runtime.Value {
	once_monadListT.Do(func() {
		cache_monadListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_monadListT
}

var cache_bindListT gopurs_runtime.Value
var once_bindListT sync.Once
func Get_bindListT() gopurs_runtime.Value {
	once_bindListT.Do(func() {
		cache_bindListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_bindListT
}

var cache_applyListT gopurs_runtime.Value
var once_applyListT sync.Once
func Get_applyListT() gopurs_runtime.Value {
	once_applyListT.Do(func() {
		cache_applyListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_applyListT
}

var cache_applicativeListT gopurs_runtime.Value
var once_applicativeListT sync.Once
func Get_applicativeListT() gopurs_runtime.Value {
	once_applicativeListT.Do(func() {
		cache_applicativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_applicativeListT
}

var cache_monadEffectListT gopurs_runtime.Value
var once_monadEffectListT sync.Once
func Get_monadEffectListT() gopurs_runtime.Value {
	once_monadEffectListT.Do(func() {
		cache_monadEffectListT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectListT((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_monadEffectListT
}

var cache_monadSTListT gopurs_runtime.Value
var once_monadSTListT sync.Once
func Get_monadSTListT() gopurs_runtime.Value {
	once_monadSTListT.Do(func() {
		cache_monadSTListT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTListT((*Record_liftST_gopurs_runtime_Value)(dictMonadST_0_box.UnsafePtr))
})
	})
	return cache_monadSTListT
}

var cache_altListT gopurs_runtime.Value
var once_altListT sync.Once
func Get_altListT() gopurs_runtime.Value {
	once_altListT.Do(func() {
		cache_altListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altListT((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_altListT
}

var cache_plusListT gopurs_runtime.Value
var once_plusListT sync.Once
func Get_plusListT() gopurs_runtime.Value {
	once_plusListT.Do(func() {
		cache_plusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_plusListT
}

var cache_alternativeListT gopurs_runtime.Value
var once_alternativeListT sync.Once
func Get_alternativeListT() gopurs_runtime.Value {
	once_alternativeListT.Do(func() {
		cache_alternativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_alternativeListT
}

var cache_monadPlusListT gopurs_runtime.Value
var once_monadPlusListT sync.Once
func Get_monadPlusListT() gopurs_runtime.Value {
	once_monadPlusListT.Do(func() {
		cache_monadPlusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusListT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_monadPlusListT
}

type Data_Control_Monad_List_Trans_Yield struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Control_Monad_List_Trans_Yield(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1320412129
}

type Data_Control_Monad_List_Trans_Skip struct {
	V0 gopurs_runtime.Value
}
func Is_Data_Control_Monad_List_Trans_Skip(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 813447293
}

type Data_Control_Monad_List_Trans_Done struct {
	
}
func Is_Data_Control_Monad_List_Trans_Done(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 489128924
}

type Record_alt_gopurs_runtime_Value struct {
	alt gopurs_runtime.Value
}

type Record_ struct {
	
}

type Record_pure_gopurs_runtime_Value struct {
	pure gopurs_runtime.Value
}

type Record_apply_gopurs_runtime_Value struct {
	apply gopurs_runtime.Value
}

type Record_bipure_gopurs_runtime_Value struct {
	bipure gopurs_runtime.Value
}

type Record_biapply_gopurs_runtime_Value struct {
	biapply gopurs_runtime.Value
}

type Record_bind_gopurs_runtime_Value struct {
	bind gopurs_runtime.Value
}

type Record_discard_gopurs_runtime_Value struct {
	discard gopurs_runtime.Value
}

type Record_identity_gopurs_runtime_Value struct {
	identity gopurs_runtime.Value
}

type Record_ask_gopurs_runtime_Value struct {
	ask gopurs_runtime.Value
}

type Record_local_gopurs_runtime_Value struct {
	local gopurs_runtime.Value
}

type Record_peek_gopurs_runtime_Value_pos_gopurs_runtime_Value struct {
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
}

type Record_track_gopurs_runtime_Value struct {
	track gopurs_runtime.Value
}

type Record_extract_gopurs_runtime_Value struct {
	extract gopurs_runtime.Value
}

type Record_extend_gopurs_runtime_Value struct {
	extend gopurs_runtime.Value
}

type Record_defer__gopurs_runtime_Value struct {
	defer_ gopurs_runtime.Value
}

type Record_callCC_gopurs_runtime_Value struct {
	callCC gopurs_runtime.Value
}

type Record_catchError_gopurs_runtime_Value struct {
	catchError gopurs_runtime.Value
}

type Record_throwError_gopurs_runtime_Value struct {
	throwError gopurs_runtime.Value
}

type Record_chooseBool_gopurs_runtime_Value_chooseFloat_gopurs_runtime_Value_chooseInt_gopurs_runtime_Value_resize_gopurs_runtime_Value_sized_gopurs_runtime_Value struct {
	chooseBool gopurs_runtime.Value
	chooseFloat gopurs_runtime.Value
	chooseInt gopurs_runtime.Value
	resize gopurs_runtime.Value
	sized gopurs_runtime.Value
}

type Record_foldMap1_gopurs_runtime_Value_foldl1_gopurs_runtime_Value_foldr1_gopurs_runtime_Value struct {
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
}

type Record_append__gopurs_runtime_Value struct {
	append_ gopurs_runtime.Value
}

type Record_tailRecM_gopurs_runtime_Value struct {
	tailRecM gopurs_runtime.Value
}

type Record_unfoldr_gopurs_runtime_Value struct {
	unfoldr gopurs_runtime.Value
}

type Record_map__gopurs_runtime_Value struct {
	map_ gopurs_runtime.Value
}

type Record_state_gopurs_runtime_Value struct {
	state gopurs_runtime.Value
}

type Record_lift_gopurs_runtime_Value struct {
	lift gopurs_runtime.Value
}

type Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value struct {
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
}

type Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value struct {
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
}

type Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value struct {
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
}

type Record_mempty_gopurs_runtime_Value struct {
	mempty gopurs_runtime.Value
}

type Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value struct {
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
}

type Record_empty_gopurs_runtime_Value struct {
	empty gopurs_runtime.Value
}

type Record_compose_gopurs_runtime_Value struct {
	compose gopurs_runtime.Value
}

type Record_eq_gopurs_runtime_Value struct {
	eq gopurs_runtime.Value
}

type Record_compare_gopurs_runtime_Value struct {
	compare gopurs_runtime.Value
}

type Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value struct {
	bifoldMap gopurs_runtime.Value
	bifoldl gopurs_runtime.Value
	bifoldr gopurs_runtime.Value
}

type Record_bimap_gopurs_runtime_Value struct {
	bimap gopurs_runtime.Value
}

type Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value struct {
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
}

type Record_genericBottom_prime_gopurs_runtime_Value struct {
	genericBottom_prime gopurs_runtime.Value
}

type Record_genericTop_prime_gopurs_runtime_Value struct {
	genericTop_prime gopurs_runtime.Value
}

type Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value struct {
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
}

type Record_lose_gopurs_runtime_Value struct {
	lose gopurs_runtime.Value
}

type Record_choose_gopurs_runtime_Value struct {
	choose gopurs_runtime.Value
}

type Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value struct {
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
}

type Record_divide_gopurs_runtime_Value struct {
	divide gopurs_runtime.Value
}

type Record_recip_gopurs_runtime_Value struct {
	recip gopurs_runtime.Value
}

type Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value struct {
	genericCardinality_prime gopurs_runtime.Value
	genericFromEnum_prime gopurs_runtime.Value
	genericToEnum_prime gopurs_runtime.Value
}

type Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value struct {
	genericPred_prime gopurs_runtime.Value
	genericSucc_prime gopurs_runtime.Value
}

type Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value struct {
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}

type Record_unfoldr1_gopurs_runtime_Value struct {
	unfoldr1 gopurs_runtime.Value
}

type Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value struct {
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}

type Record_genericEq_prime_gopurs_runtime_Value struct {
	genericEq_prime gopurs_runtime.Value
}

type Record_eq1_gopurs_runtime_Value struct {
	eq1 gopurs_runtime.Value
}

type Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value struct {
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff gopurs_runtime.Value
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt gopurs_runtime.Value
}

type Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value struct {
	add gopurs_runtime.Value
	mul gopurs_runtime.Value
	one gopurs_runtime.Value
	zero gopurs_runtime.Value
}

type Record_foldMapWithIndex_gopurs_runtime_Value_foldlWithIndex_gopurs_runtime_Value_foldrWithIndex_gopurs_runtime_Value struct {
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
}

type Record_cmap_gopurs_runtime_Value struct {
	cmap gopurs_runtime.Value
}

type Record_imap_gopurs_runtime_Value struct {
	imap gopurs_runtime.Value
}

type Record_mapWithIndex_gopurs_runtime_Value struct {
	mapWithIndex gopurs_runtime.Value
}

type Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value struct {
	from gopurs_runtime.Value
	to gopurs_runtime.Value
}

type Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value struct {
	genericConj_prime gopurs_runtime.Value
	genericDisj_prime gopurs_runtime.Value
	genericFF_prime gopurs_runtime.Value
	genericImplies_prime gopurs_runtime.Value
	genericNot_prime gopurs_runtime.Value
	genericTT_prime gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_bool_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_bool struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff bool
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt bool
}

type Record_genericMempty_prime_gopurs_runtime_Value struct {
	genericMempty_prime gopurs_runtime.Value
}

type Record_genericCompare_prime_gopurs_runtime_Value struct {
	genericCompare_prime gopurs_runtime.Value
}

type Record_sub_gopurs_runtime_Value struct {
	sub gopurs_runtime.Value
}

type Record_compare1_gopurs_runtime_Value struct {
	compare1 gopurs_runtime.Value
}

type Record_left_gopurs_runtime_Value_right_gopurs_runtime_Value struct {
	left gopurs_runtime.Value
	right gopurs_runtime.Value
}

type Record_first_gopurs_runtime_Value_second_gopurs_runtime_Value struct {
	first gopurs_runtime.Value
	second gopurs_runtime.Value
}

type Record_dimap_gopurs_runtime_Value struct {
	dimap gopurs_runtime.Value
}

type Record_genericSub_prime_gopurs_runtime_Value struct {
	genericSub_prime gopurs_runtime.Value
}

type Record_genericAppend_prime_gopurs_runtime_Value struct {
	genericAppend_prime gopurs_runtime.Value
}

type Record_sequence1_gopurs_runtime_Value_traverse1_gopurs_runtime_Value struct {
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
}

type Record_genericAdd_prime_gopurs_runtime_Value_genericMul_prime_gopurs_runtime_Value_genericOne_prime_gopurs_runtime_Value_genericZero_prime_gopurs_runtime_Value struct {
	genericAdd_prime gopurs_runtime.Value
	genericMul_prime gopurs_runtime.Value
	genericOne_prime gopurs_runtime.Value
	genericZero_prime gopurs_runtime.Value
}

type Record_genericShow_prime_gopurs_runtime_Value struct {
	genericShow_prime gopurs_runtime.Value
}

type Record_genericShowArgs_gopurs_runtime_Value struct {
	genericShowArgs gopurs_runtime.Value
}

type Record_show_gopurs_runtime_Value struct {
	show gopurs_runtime.Value
}

type Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value struct {
	fromDuration gopurs_runtime.Value
	toDuration gopurs_runtime.Value
}

type Record_traverseWithIndex_gopurs_runtime_Value struct {
	traverseWithIndex gopurs_runtime.Value
}

type Record_liftEffect_gopurs_runtime_Value struct {
	liftEffect gopurs_runtime.Value
}

type Record_mappend__gopurs_runtime_Value_mempty__gopurs_runtime_Value struct {
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
}

type Record_proof_gopurs_runtime_Value struct {
	proof gopurs_runtime.Value
}

type Record_lower_gopurs_runtime_Value struct {
	lower gopurs_runtime.Value
}

type Record_liftST_gopurs_runtime_Value struct {
	liftST gopurs_runtime.Value
}

type Record_tell_gopurs_runtime_Value struct {
	tell gopurs_runtime.Value
}

type Record_reflectSymbol_gopurs_runtime_Value struct {
	reflectSymbol gopurs_runtime.Value
}

type Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value struct {
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
}

type Record_conquer_gopurs_runtime_Value struct {
	conquer gopurs_runtime.Value
}

type Record_inj_gopurs_runtime_Value_prj_gopurs_runtime_Value struct {
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
}

type Record_eqRecord_gopurs_runtime_Value struct {
	eqRecord gopurs_runtime.Value
}

type Record_conjRecord_gopurs_runtime_Value_disjRecord_gopurs_runtime_Value_ffRecord_gopurs_runtime_Value_impliesRecord_gopurs_runtime_Value_notRecord_gopurs_runtime_Value_ttRecord_gopurs_runtime_Value struct {
	conjRecord gopurs_runtime.Value
	disjRecord gopurs_runtime.Value
	ffRecord gopurs_runtime.Value
	impliesRecord gopurs_runtime.Value
	notRecord gopurs_runtime.Value
	ttRecord gopurs_runtime.Value
}

type Record_memptyRecord_gopurs_runtime_Value struct {
	memptyRecord gopurs_runtime.Value
}

type Record_compareRecord_gopurs_runtime_Value struct {
	compareRecord gopurs_runtime.Value
}

type Record_closed_gopurs_runtime_Value struct {
	closed gopurs_runtime.Value
}

type Record_unleft_gopurs_runtime_Value_unright_gopurs_runtime_Value struct {
	unleft gopurs_runtime.Value
	unright gopurs_runtime.Value
}

type Record_unfirst_gopurs_runtime_Value_unsecond_gopurs_runtime_Value struct {
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
}

type Record_reflectType_gopurs_runtime_Value struct {
	reflectType gopurs_runtime.Value
}

type Record_subRecord_gopurs_runtime_Value struct {
	subRecord gopurs_runtime.Value
}

type Record_appendRecord_gopurs_runtime_Value struct {
	appendRecord gopurs_runtime.Value
}

type Record_addRecord_gopurs_runtime_Value_mulRecord_gopurs_runtime_Value_oneRecord_gopurs_runtime_Value_zeroRecord_gopurs_runtime_Value struct {
	addRecord gopurs_runtime.Value
	mulRecord gopurs_runtime.Value
	oneRecord gopurs_runtime.Value
	zeroRecord gopurs_runtime.Value
}

type Record_showRecordFields_gopurs_runtime_Value struct {
	showRecordFields gopurs_runtime.Value
}

type Record_nes_gopurs_runtime_Value struct {
	nes gopurs_runtime.Value
}

type Record_liftAff_gopurs_runtime_Value struct {
	liftAff gopurs_runtime.Value
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_ListT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_wrapLazy(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{v_1})})
}

func Call_wrapEffect(dictFunctor_0_loop *Record_map__gopurs_runtime_Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(dictFunctor_0.map_, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
})})}
}), v_1)
}

func Call_unfold(dictMonad_0_loop *Record_, f_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unfold:
for {
if false { continue unfold }
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 930809136) {
__local_var_4_1 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_3.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_4_1
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_3.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfold(dictMonad_0, f_1, __local_var_4_1)
})})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 3589588149) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Apply(f_1, z_2))
}
}

func Call_uncons(dictMonad_0_loop *Record_) gopurs_runtime.Value {
uncons:
for {
if false { continue uncons }
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "bind"), v_2, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1320412129) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*Data_Control_Monad_List_Trans_Yield)(v1_3.UnsafePtr).V0, gopurs_runtime.Apply((*Data_Control_Monad_List_Trans_Yield)(v1_3.UnsafePtr).V1, pkg_Data_Unit.Get_unit())})}})})
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 813447293) {
__t1 = gopurs_runtime.Apply2(Get_uncons(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Apply((*Data_Control_Monad_List_Trans_Skip)(v1_3.UnsafePtr).V0, pkg_Data_Unit.Get_unit()))
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 489128924) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil})
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
}
}

func Call_tail(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = uncons1_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_snd()), gopurs_runtime.Apply(uncons1_1_0, l_2))
})
}

func Call_takeWhile(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{(*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_takeWhile(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_2), (*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_takeWhile(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_2), (*Data_Control_Monad_List_Trans_Skip)(v_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v_3)
})
}
}

func Call_scanl(dictMonad_0_loop *Record_, f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, l_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var l_3 gopurs_runtime.Value = l_3_loop
_ = l_3
return Call_unfold(dictMonad_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0
_ = __local_var_5_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(f_1, __local_var_5_0, (*Data_Control_Monad_List_Trans_Yield)(v1_6.UnsafePtr).V0), gopurs_runtime.Apply((*Data_Control_Monad_List_Trans_Yield)(v1_6.UnsafePtr).V1, pkg_Data_Unit.Get_unit())})}, __local_var_5_0})}})}
goto end_branch_1
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{__local_var_5_0, gopurs_runtime.Apply((*Data_Control_Monad_List_Trans_Skip)(v1_6.UnsafePtr).V0, pkg_Data_Unit.Get_unit())})}, __local_var_5_0})}})}
goto end_branch_1
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_2, l_3})})
}

func Call_prepend_prime(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value, h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{h_1, t_2})})
}

func Call_prepend(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value, h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{h_1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return t_2
})})})
}

func Call_nil(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
}

func Call_singleton(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{a_2, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
})})})
})
}

func Call_take(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
take:
for {
if false { continue take }
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.IntVal) == (0) {
__t3 = nil1_1_0
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1320412129) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{(*Data_Control_Monad_List_Trans_Yield)(v2_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Int((v_3.IntVal) - (1))), (*Data_Control_Monad_List_Trans_Yield)(v2_5.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, v_3), (*Data_Control_Monad_List_Trans_Skip)(v2_5.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 489128924) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
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
}
}

func Call_zipWith_prime(dictMonad_0_loop *Record_) gopurs_runtime.Value {
zipWith_prime:
for {
if false { continue zipWith_prime }
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
Bind1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{})
_ = Bind1_3_2
Functor0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_4_3
uncons1_5_4 := gopurs_runtime.Apply(Get_uncons(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = uncons1_5_4
return gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, fa_7 gopurs_runtime.Value, fb_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
})})}
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(uncons1_5_4, fa_7), gopurs_runtime.Func(func(ua_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(uncons1_5_4, fb_8), gopurs_runtime.Func(func(ub_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (ub_10.Type == 9 && ub_10.IntVal == 3589588149) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), nil1_2_1)
goto end_branch_5
} else {

}
}
{
if (ua_9.Type == 9 && ua_9.IntVal == 3589588149) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), nil1_2_1)
goto end_branch_5
} else {

}
}
{
if ((ua_9.Type == 9 && ua_9.IntVal == 930809136)) && ((ub_10.Type == 9 && ub_10.IntVal == 930809136)) {
__local_var_11_6 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(ua_9.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_6
__local_var_12_7 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(ub_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_12_7
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{a_13, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_zipWith_prime(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_6, __local_var_11_6, __local_var_12_7)
})})})
}), gopurs_runtime.Apply2(f_6, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(ua_9.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(ub_10.UnsafePtr).V0.UnsafePtr).V0))
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
})))
})
}
}

func Call_zipWith(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
zipWith_prime1_1_0 := gopurs_runtime.Apply(Get_zipWith_prime(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = zipWith_prime1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(zipWith_prime1_1_0, gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply2(f_2, a_3, b_4))
}))
})
}

func Call_mapMaybe(dictFunctor_0_loop *Record_map__gopurs_runtime_Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.map_, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Yield(), gopurs_runtime.Apply(f_1, (*Data_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0))
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 3589588149) {
__t2 = Get_Skip()
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 930809136) {
__t2 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_4_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = gopurs_runtime.Apply(__t2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Data_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1))
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, f_1), (*Data_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
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
}
}

func Call_iterate(dictMonad_0_loop *Record_, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return Call_unfold(dictMonad_0, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply(f_1, x_3), x_3})}})})
}), a_2)
}

func Call_repeat(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply2(Get_iterate(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, Get_identity())
}

func Call_head(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = uncons1_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_fst()), gopurs_runtime.Apply(uncons1_1_0, l_2))
})
}

func Call_functorListT(dictFunctor_0_loop *Record_map__gopurs_runtime_Value) gopurs_runtime.Value {
functorListT:
for {
if false { continue functorListT }
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFunctor_0.map_, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{gopurs_runtime.Apply(f_1, (*Data_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_functorListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictFunctor_0)}), "map"), f_1), (*Data_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_functorListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictFunctor_0)}), "map"), f_1), (*Data_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
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
}))
}
}

func Call_fromEffect(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_1_0
return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{a_3, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
})})}
}), fa_2)
})
}

func Call_foldlRec_prime(dictMonadRec_0_loop *Record_tailRecM_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadRec_0 *Record_tailRecM_gopurs_runtime_Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
uncons1_4_3 := gopurs_runtime.Apply(Get_uncons(), Monad0_1_0)
_ = uncons1_4_3
return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.tailRecM, gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(uncons1_4_3, gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 3589588149) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{__local_var_9_4})})
goto end_branch_5
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136) {
__local_var_11_6 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_6
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply2(f_5, __local_var_9_4, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_6)})})
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
}), gopurs_runtime.RecordDict2("a", "b", a_6, b_7))
})
}

func Call_runListTRec(dictMonadRec_0_loop *Record_tailRecM_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadRec_0 *Record_tailRecM_gopurs_runtime_Value = dictMonadRec_0_loop
_ = dictMonadRec_0
return gopurs_runtime.Apply3(Get_foldlRec_prime(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), pkg_Data_Unit.Get_unit())
}

func Call_foldlRec(dictMonadRec_0_loop *Record_tailRecM_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadRec_0 *Record_tailRecM_gopurs_runtime_Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
uncons1_3_2 := gopurs_runtime.Apply(Get_uncons(), Monad0_1_0)
_ = uncons1_3_2
return gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.tailRecM, gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := gopurs_runtime.RecordGet(o_7, "a")
_ = __local_var_8_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_3_2, gopurs_runtime.RecordGet(o_7, "b")), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 3589588149) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{__local_var_8_3})})
goto end_branch_4
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 930809136) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Apply2(f_4, __local_var_8_3, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_9.UnsafePtr).V0.UnsafePtr).V0), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_9.UnsafePtr).V0.UnsafePtr).V1)})})
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
}), gopurs_runtime.RecordDict2("a", "b", a_5, b_6))
})
}

func Call_foldl_prime(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
uncons1_2_1 := gopurs_runtime.Apply(Get_uncons(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = uncons1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2 gopurs_runtime.Value
_ = loop_4_2
loop_4_2 = gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply(uncons1_2_1, l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 3589588149) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"), b_5)
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136) {
__local_var_8_4 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_4
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply2(f_3, b_5, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(loop_4_2, a_9, __local_var_8_4)
}))
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
return loop_4_2
})
}

func Call_runListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply3(Get_foldl_prime(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), pkg_Data_Unit.Get_unit())
}

func Call_foldl(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = uncons1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_3_1 gopurs_runtime.Value
_ = loop_3_1
loop_3_1 = gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_1_0, l_5), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3589588149) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"), b_4)
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 930809136) {
__t2 = gopurs_runtime.Apply2(loop_3_1, gopurs_runtime.Apply2(f_2, b_4, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_6.UnsafePtr).V0.UnsafePtr).V0), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_6.UnsafePtr).V0.UnsafePtr).V1)
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
return loop_3_1
})
}

func Call_filter(dictFunctor_0_loop *Record_map__gopurs_runtime_Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.map_, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
s_prime_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), func() gopurs_runtime.Value {
arr_val_filter5 := f_1
_ = arr_val_filter5
arr_go_filter5 := (*[]gopurs_runtime.Value)(arr_val_filter5.UnsafePtr)
_ = arr_go_filter5
res_go_filter5 := make([]gopurs_runtime.Value, 0)
_ = res_go_filter5
for _, v_filter5 := range *arr_go_filter5 {
if gopurs_runtime.Apply(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, v_filter5).BoolVal() {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}(), (*Data_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V1)
_ = s_prime_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, (*Data_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{(*Data_Control_Monad_List_Trans_Yield)(v_3.UnsafePtr).V0, s_prime_4_1})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{s_prime_4_1})}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), func() gopurs_runtime.Value {
arr_val_filter5 := f_1
_ = arr_val_filter5
arr_go_filter5 := (*[]gopurs_runtime.Value)(arr_val_filter5.UnsafePtr)
_ = arr_go_filter5
res_go_filter5 := make([]gopurs_runtime.Value, 0)
_ = res_go_filter5
for _, v_filter5 := range *arr_go_filter5 {
if gopurs_runtime.Apply(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, v_filter5).BoolVal() {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}(), (*Data_Control_Monad_List_Trans_Skip)(v_3.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
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
}
}

func Call_dropWhile(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
dropWhile:
for {
if false { continue dropWhile }
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_dropWhile(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_2), (*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{(*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0, (*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_dropWhile(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_2), (*Data_Control_Monad_List_Trans_Skip)(v_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v_3)
})
}
}

func Call_drop(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Int((v_2.IntVal) - (1))), (*Data_Control_Monad_List_Trans_Yield)(v2_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, v_2), (*Data_Control_Monad_List_Trans_Skip)(v2_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
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
}
}

func Call_cons(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value, lh_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 gopurs_runtime.Value = lh_1_loop
_ = lh_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{gopurs_runtime.Apply(lh_1, pkg_Data_Unit.Get_unit()), t_2})})
}

func Call_unfoldable1ListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_2 gopurs_runtime.Value
_ = go__5_2
go__5_2 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3589588149) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V0, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_2_1
})})})
goto end_branch_3
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V1
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136) {
__local_var_7_6 := (*pkg_Data_Maybe.Data_Data_Maybe_Just)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V1.UnsafePtr).V0
_ = __local_var_7_6
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V0, pkg_Data_Unit.Get_unit()), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__5_2, gopurs_runtime.Apply(f_3, __local_var_7_6))
})})})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
return gopurs_runtime.Apply(go__5_2, gopurs_runtime.Apply(f_3, b_4))
}))
}

func Call_unfoldableListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
unfoldable1ListT1_3_2 := gopurs_runtime.Apply(Get_unfoldable1ListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = unfoldable1ListT1_3_2
return gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return unfoldable1ListT1_3_2
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__6_3 gopurs_runtime.Value
_ = go__6_3
go__6_3 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 3589588149) {
__t4 = nil1_2_1
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136) {
__local_var_8_5 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_5
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_7.UnsafePtr).V0.UnsafePtr).V0, pkg_Data_Unit.Get_unit()), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__6_3, gopurs_runtime.Apply(f_4, __local_var_8_5))
})})})
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
return gopurs_runtime.Apply(go__6_3, gopurs_runtime.Apply(f_4, b_5))
}))
}

func Call_semigroupListT(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_concat(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}))
}

func Call_concat(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{(*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_concat(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, v1_5, y_3)
}), (*Data_Control_Monad_List_Trans_Yield)(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_concat(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, v1_5, y_3)
}), (*Data_Control_Monad_List_Trans_Skip)(v_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return y_3
})})}
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
}

func Call_monoidListT(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
semigroupListT1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_concat(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}))
_ = semigroupListT1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupListT1_1_0
}), gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
}

func Call_catMaybes(dictFunctor_0_loop *Record_map__gopurs_runtime_Value) gopurs_runtime.Value {
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply2(Get_mapMaybe(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, Get_identity())
}

func Call_monadListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
}))
}

func Call_bindListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
append_1_0 := gopurs_runtime.Apply(Get_concat(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = append_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
}), gopurs_runtime.Func2(func(fa_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1320412129) {
__local_var_6_3 := (*Data_Control_Monad_List_Trans_Yield)(v_5.UnsafePtr).V0
_ = __local_var_6_3
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(append_1_0, gopurs_runtime.Apply(f_4, __local_var_6_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_bindListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}), "bind"), s_prime_7, f_4))
}), (*Data_Control_Monad_List_Trans_Yield)(v_5.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Skip{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_bindListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}), "bind"), v1_6, f_4)
}), (*Data_Control_Monad_List_Trans_Skip)(v_5.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 489128924) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), fa_3)
}))
}

func Call_applyListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
functorListT1_1_0 := gopurs_runtime.Apply(Get_functorListT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
__local_var_2_1 := gopurs_runtime.Apply(Get_bindListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}), "pure"), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
}))
}

func Call_applicativeListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
}), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_List_Trans_Yield{a_3, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_2_1
})})})
}))
}

func Call_monadEffectListT(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeListT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), Monad0_1_0)
}))
_ = monadListT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransListT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, x_4))
}))
}

func Call_monadSTListT(dictMonadST_0_loop *Record_liftST_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadST_0 *Record_liftST_gopurs_runtime_Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadST_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeListT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), Monad0_1_0)
}))
_ = monadListT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransListT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(dictMonadST_0.liftST, x_4))
}))
}

func Call_altListT(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
functorListT1_1_0 := gopurs_runtime.Apply(Get_functorListT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), gopurs_runtime.Apply(Get_concat(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}))
}

func Call_plusListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Applicative0_1_0
altListT1_2_1 := gopurs_runtime.Apply(Get_altListT(), Applicative0_1_0)
_ = altListT1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altListT1_2_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
}

func Call_alternativeListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
applicativeListT1_1_0 := gopurs_runtime.Apply(Get_applicativeListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = applicativeListT1_1_0
plusListT1_2_1 := gopurs_runtime.Apply(Get_plusListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = plusListT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeListT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusListT1_2_1
}))
}

func Call_monadPlusListT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
monadListT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
}))
_ = monadListT1_1_0
alternativeListT1_2_1 := gopurs_runtime.Apply(Get_alternativeListT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = alternativeListT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeListT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_1_0
}))
}


