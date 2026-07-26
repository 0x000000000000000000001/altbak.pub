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
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{value0})}
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
return Call_wrapLazy(dictApplicative_0_box, v_1_box)
})
	})
	return cache_wrapLazy
}

var cache_wrapEffect gopurs_runtime.Value
var once_wrapEffect sync.Once
func Get_wrapEffect() gopurs_runtime.Value {
	once_wrapEffect.Do(func() {
		cache_wrapEffect = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrapEffect(dictFunctor_0_box, v_1_box)
})
	})
	return cache_wrapEffect
}

var cache_unfold gopurs_runtime.Value
var once_unfold sync.Once
func Get_unfold() gopurs_runtime.Value {
	once_unfold.Do(func() {
		cache_unfold = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfold(dictMonad_0_box, f_1_box, z_2_box)
})
	})
	return cache_unfold
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(dictMonad_0_box)
})
	})
	return cache_uncons
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail(dictMonad_0_box)
})
	})
	return cache_tail
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(dictApplicative_0_box)
})
	})
	return cache_takeWhile
}

var cache_scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		cache_scanl = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanl(dictMonad_0_box, f_1_box, b_2_box, l_3_box)
})
	})
	return cache_scanl
}

var cache_prepend_prime gopurs_runtime.Value
var once_prepend_prime sync.Once
func Get_prepend_prime() gopurs_runtime.Value {
	once_prepend_prime.Do(func() {
		cache_prepend_prime = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prepend_prime(dictApplicative_0_box, h_1_box, t_2_box)
})
	})
	return cache_prepend_prime
}

var cache_prepend gopurs_runtime.Value
var once_prepend sync.Once
func Get_prepend() gopurs_runtime.Value {
	once_prepend.Do(func() {
		cache_prepend = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prepend(dictApplicative_0_box, h_1_box, t_2_box)
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
return Call_singleton(dictApplicative_0_box)
})
	})
	return cache_singleton
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take(dictApplicative_0_box)
})
	})
	return cache_take
}

var cache_zipWith_prime gopurs_runtime.Value
var once_zipWith_prime sync.Once
func Get_zipWith_prime() gopurs_runtime.Value {
	once_zipWith_prime.Do(func() {
		cache_zipWith_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith_prime(dictMonad_0_box)
})
	})
	return cache_zipWith_prime
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith(dictMonad_0_box)
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
return Call_mapMaybe(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return cache_mapMaybe
}

var cache_iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		cache_iterate = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate(dictMonad_0_box, f_1_box, a_2_box)
})
	})
	return cache_iterate
}

var cache_repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		cache_repeat = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat(dictMonad_0_box)
})
	})
	return cache_repeat
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(dictMonad_0_box)
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
return Call_fromEffect(dictApplicative_0_box)
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
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{a_4, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
return Call_foldlRec_prime(dictMonadRec_0_box)
})
	})
	return cache_foldlRec_prime
}

var cache_runListTRec gopurs_runtime.Value
var once_runListTRec sync.Once
func Get_runListTRec() gopurs_runtime.Value {
	once_runListTRec.Do(func() {
		cache_runListTRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runListTRec(dictMonadRec_0_box)
})
	})
	return cache_runListTRec
}

var cache_foldlRec gopurs_runtime.Value
var once_foldlRec sync.Once
func Get_foldlRec() gopurs_runtime.Value {
	once_foldlRec.Do(func() {
		cache_foldlRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec(dictMonadRec_0_box)
})
	})
	return cache_foldlRec
}

var cache_foldl_prime gopurs_runtime.Value
var once_foldl_prime sync.Once
func Get_foldl_prime() gopurs_runtime.Value {
	once_foldl_prime.Do(func() {
		cache_foldl_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl_prime(dictMonad_0_box)
})
	})
	return cache_foldl_prime
}

var cache_runListT gopurs_runtime.Value
var once_runListT sync.Once
func Get_runListT() gopurs_runtime.Value {
	once_runListT.Do(func() {
		cache_runListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runListT(dictMonad_0_box)
})
	})
	return cache_runListT
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(dictMonad_0_box)
})
	})
	return cache_foldl
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return cache_filter
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(dictApplicative_0_box)
})
	})
	return cache_dropWhile
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop(dictApplicative_0_box)
})
	})
	return cache_drop
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons(dictApplicative_0_box, lh_1_box, t_2_box)
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
return Call_concat(dictApplicative_0_box)
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
return Call_catMaybes(dictFunctor_0_box)
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

type Constructor_Yield[T_a any, T_s any] struct {
	V0 T_a
	V1 gopurs_runtime.Value
}


type Constructor_Skip[T_a any, T_s any] struct {
	V0 gopurs_runtime.Value
}


type Constructor_Done[T_a any, T_s any] struct {
	
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

func Call_wrapLazy(dictApplicative_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{v_1})})
}

func Call_wrapEffect(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
})})}
}), v_1)
}

func Call_unfold(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unfold:
for {
if false { continue unfold }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 930809136) {
__local_var_4_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_4_1
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_uncons(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
uncons:
for {
if false { continue uncons }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), v_2, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1320412129) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1, pkg_Data_Unit.Get_unit())})}})})
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 813447293) {
__t1 = gopurs_runtime.Apply2(Get_uncons(), dictMonad_0, gopurs_runtime.Apply((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0, pkg_Data_Unit.Get_unit()))
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

func Call_tail(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_snd()), gopurs_runtime.Apply(uncons1_1_0, l_2))
})
}

func Call_takeWhile(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_takeWhile(), dictApplicative_0, f_2), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_takeWhile(), dictApplicative_0, f_2), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}
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

func Call_scanl(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, l_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var l_3 gopurs_runtime.Value = l_3_loop
_ = l_3
return Call_unfold(dictMonad_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
_ = __local_var_5_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(f_1, __local_var_5_0, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1, pkg_Data_Unit.Get_unit())})}, __local_var_5_0})}})}
goto end_branch_1
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{__local_var_5_0, gopurs_runtime.Apply((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0, pkg_Data_Unit.Get_unit())})}, __local_var_5_0})}})}
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
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_2, l_3})})
}

func Call_prepend_prime(dictApplicative_0_loop gopurs_runtime.Value, h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{h_1, t_2})})
}

func Call_prepend(dictApplicative_0_loop gopurs_runtime.Value, h_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 gopurs_runtime.Value = h_1_loop
_ = h_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{h_1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return t_2
})})})
}

func Call_nil(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
}

func Call_singleton(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
})})})
})
}

func Call_take(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
take:
for {
if false { continue take }
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_take(), dictApplicative_0, gopurs_runtime.Int((v_3.IntVal) - (1))), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_take(), dictApplicative_0, v_3), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)})}
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

func Call_zipWith_prime(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
zipWith_prime:
for {
if false { continue zipWith_prime }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
Bind1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_3_2
Functor0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_4_3
uncons1_5_4 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_5_4
return gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, fa_7 gopurs_runtime.Value, fb_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
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
__local_var_11_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_9.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_6
__local_var_12_7 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_12_7
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{a_13, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_zipWith_prime(), dictMonad_0, f_6, __local_var_11_6, __local_var_12_7)
})})})
}), gopurs_runtime.Apply2(f_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_9.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_10.UnsafePtr).V0.UnsafePtr).V0))
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

func Call_zipWith(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
zipWith_prime1_1_0 := gopurs_runtime.Apply(Get_zipWith_prime(), dictMonad_0)
_ = zipWith_prime1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(zipWith_prime1_1_0, gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply2(f_2, a_3, b_4))
}))
})
}

func Call_mapMaybe(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Yield(), gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))
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
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = gopurs_runtime.Apply(__t2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, f_1), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, f_1), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_iterate(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return Call_unfold(dictMonad_0, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_1, x_3), x_3})}})})
}), a_2)
}

func Call_repeat(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply2(Get_iterate(), dictMonad_0, Get_identity())
}

func Call_head(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_fst()), gopurs_runtime.Apply(uncons1_1_0, l_2))
})
}

func Call_functorListT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
functorListT:
for {
if false { continue functorListT }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_functorListT(), dictFunctor_0), "map"), f_1), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_functorListT(), dictFunctor_0), "map"), f_1), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_fromEffect(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_1_0
return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{a_3, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
})})}
}), fa_2)
})
}

func Call_foldlRec_prime(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
uncons1_4_3 := gopurs_runtime.Apply(Get_uncons(), Monad0_1_0)
_ = uncons1_4_3
return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictMonadRec_0.UnsafePtr)).V0, gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(uncons1_4_3, gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 3589588149) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{__local_var_9_4})})
goto end_branch_5
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136) {
__local_var_11_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_11_6
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply2(f_5, __local_var_9_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_6)})})
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

func Call_runListTRec(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
return gopurs_runtime.Apply3(Get_foldlRec_prime(), dictMonadRec_0, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), pkg_Data_Unit.Get_unit())
}

func Call_foldlRec(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
uncons1_3_2 := gopurs_runtime.Apply(Get_uncons(), Monad0_1_0)
_ = uncons1_3_2
return gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictMonadRec_0.UnsafePtr)).V0, gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := gopurs_runtime.RecordGet(o_7, "a")
_ = __local_var_8_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_3_2, gopurs_runtime.RecordGet(o_7, "b")), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 3589588149) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{__local_var_8_3})})
goto end_branch_4
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 930809136) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Apply2(f_4, __local_var_8_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V1)})})
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

func Call_foldl_prime(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_0
uncons1_2_1 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2 gopurs_runtime.Value
_ = loop_4_2
loop_4_2 = gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply(uncons1_2_1, l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 3589588149) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_5)
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136) {
__local_var_8_4 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_4
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply2(f_3, b_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_runListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply3(Get_foldl_prime(), dictMonad_0, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), pkg_Data_Unit.Get_unit())
}

func Call_foldl(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_3_1 gopurs_runtime.Value
_ = loop_3_1
loop_3_1 = gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_1_0, l_5), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3589588149) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_4)
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 930809136) {
__t2 = gopurs_runtime.Apply2(loop_3_1, gopurs_runtime.Apply2(f_2, b_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0.UnsafePtr).V1)
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

func Call_filter(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Apply(dictFunctor_0, v_filter5).BoolVal() {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}(), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
_ = s_prime_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_prime_4_1})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{s_prime_4_1})}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), func() gopurs_runtime.Value {
arr_val_filter5 := f_1
_ = arr_val_filter5
arr_go_filter5 := (*[]gopurs_runtime.Value)(arr_val_filter5.UnsafePtr)
_ = arr_go_filter5
res_go_filter5 := make([]gopurs_runtime.Value, 0)
_ = res_go_filter5
for _, v_filter5 := range *arr_go_filter5 {
if gopurs_runtime.Apply(dictFunctor_0, v_filter5).BoolVal() {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}(), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)})}
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

func Call_dropWhile(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
dropWhile:
for {
if false { continue dropWhile }
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_dropWhile(), dictApplicative_0, f_2), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_dropWhile(), dictApplicative_0, f_2), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}
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

func Call_drop(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_drop(), dictApplicative_0, gopurs_runtime.Int((v_2.IntVal) - (1))), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_drop(), dictApplicative_0, v_2), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)})}
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

func Call_cons(dictApplicative_0_loop gopurs_runtime.Value, lh_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 gopurs_runtime.Value = lh_1_loop
_ = lh_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(lh_1, pkg_Data_Unit.Get_unit()), t_2})})
}

func Call_unfoldable1ListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_2 gopurs_runtime.Value
_ = go__5_2
go__5_2 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3589588149) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_2_1
})})})
goto end_branch_3
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136) {
__local_var_7_6 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1.UnsafePtr).V0
_ = __local_var_7_6
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, pkg_Data_Unit.Get_unit()), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_unfoldableListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
unfoldable1ListT1_3_2 := gopurs_runtime.Apply(Get_unfoldable1ListT(), dictMonad_0)
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
__local_var_8_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_8_5
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0.UnsafePtr).V0, pkg_Data_Unit.Get_unit()), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_semigroupListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_concat(), dictApplicative_0))
}

func Call_concat(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_concat(), dictApplicative_0, v1_5, y_3)
}), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_concat(), dictApplicative_0, v1_5, y_3)
}), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_monoidListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
semigroupListT1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_concat(), dictApplicative_0))
_ = semigroupListT1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupListT1_1_0
}), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
}

func Call_catMaybes(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, Get_identity())
}

func Call_monadListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeListT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), dictMonad_0)
}))
}

func Call_bindListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
append_1_0 := gopurs_runtime.Apply(Get_concat(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = append_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), dictMonad_0)
}), gopurs_runtime.Func2(func(fa_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1320412129) {
__local_var_6_3 := (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(append_1_0, gopurs_runtime.Apply(f_4, __local_var_6_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_bindListT(), dictMonad_0), "bind"), s_prime_7, f_4))
}), (*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_bindListT(), dictMonad_0), "bind"), v1_6, f_4)
}), (*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)})}
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

func Call_applyListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
functorListT1_1_0 := gopurs_runtime.Apply(Get_functorListT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
__local_var_2_1 := gopurs_runtime.Apply(Get_bindListT(), dictMonad_0)
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeListT(), dictMonad_0), "pure"), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
}))
}

func Call_applicativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
_ = nil1_2_1
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), dictMonad_0)
}), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{a_3, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_2_1
})})})
}))
}

func Call_monadEffectListT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
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
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictMonadEffect_0.UnsafePtr)).V0, x_4))
}))
}

func Call_monadSTListT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
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
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictMonadST_0.UnsafePtr)).V0, x_4))
}))
}

func Call_altListT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
functorListT1_1_0 := gopurs_runtime.Apply(Get_functorListT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), gopurs_runtime.Apply(Get_concat(), dictApplicative_0))
}

func Call_plusListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
altListT1_2_1 := gopurs_runtime.Apply(Get_altListT(), Applicative0_1_0)
_ = altListT1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altListT1_2_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
}

func Call_alternativeListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeListT1_1_0 := gopurs_runtime.Apply(Get_applicativeListT(), dictMonad_0)
_ = applicativeListT1_1_0
plusListT1_2_1 := gopurs_runtime.Apply(Get_plusListT(), dictMonad_0)
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
return gopurs_runtime.Apply(Get_applicativeListT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), dictMonad_0)
}))
_ = monadListT1_1_0
alternativeListT1_2_1 := gopurs_runtime.Apply(Get_alternativeListT(), dictMonad_0)
_ = alternativeListT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeListT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_1_0
}))
}


