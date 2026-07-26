package Control_Monad_Identity_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_IdentityT gopurs_runtime.Value
var once_IdentityT sync.Once
func Get_IdentityT() gopurs_runtime.Value {
	once_IdentityT.Do(func() {
		cache_IdentityT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_IdentityT(x_0_box)
})
	})
	return cache_IdentityT
}

var cache_monadSTIdentityT gopurs_runtime.Value
var once_monadSTIdentityT sync.Once
func Get_monadSTIdentityT() gopurs_runtime.Value {
	once_monadSTIdentityT.Do(func() {
		cache_monadSTIdentityT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadSTIdentityT((*Record_liftST_gopurs_runtime_Value)(dictMonadST_0_box.UnsafePtr)))}
})
	})
	return cache_monadSTIdentityT
}

var cache_traversableIdentityT gopurs_runtime.Value
var once_traversableIdentityT sync.Once
func Get_traversableIdentityT() gopurs_runtime.Value {
	once_traversableIdentityT.Do(func() {
		cache_traversableIdentityT = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_traversableIdentityT((*Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value)(dictTraversable_0_box.UnsafePtr)))}
})
	})
	return cache_traversableIdentityT
}

var cache_runIdentityT gopurs_runtime.Value
var once_runIdentityT sync.Once
func Get_runIdentityT() gopurs_runtime.Value {
	once_runIdentityT.Do(func() {
		cache_runIdentityT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runIdentityT(v_0_box)
})
	})
	return cache_runIdentityT
}

var cache_plusIdentityT gopurs_runtime.Value
var once_plusIdentityT sync.Once
func Get_plusIdentityT() gopurs_runtime.Value {
	once_plusIdentityT.Do(func() {
		cache_plusIdentityT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_plusIdentityT((*Record_empty_gopurs_runtime_Value)(dictPlus_0_box.UnsafePtr)))}
})
	})
	return cache_plusIdentityT
}

var cache_newtypeIdentityT gopurs_runtime.Value
var once_newtypeIdentityT sync.Once
func Get_newtypeIdentityT() gopurs_runtime.Value {
	once_newtypeIdentityT.Do(func() {
		cache_newtypeIdentityT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeIdentityT
}

var cache_monadWriterIdentityT gopurs_runtime.Value
var once_monadWriterIdentityT sync.Once
func Get_monadWriterIdentityT() gopurs_runtime.Value {
	once_monadWriterIdentityT.Do(func() {
		cache_monadWriterIdentityT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadWriterIdentityT((*Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value)(dictMonadWriter_0_box.UnsafePtr)))}
})
	})
	return cache_monadWriterIdentityT
}

var cache_monadTransIdentityT gopurs_runtime.Value
var once_monadTransIdentityT sync.Once
func Get_monadTransIdentityT() gopurs_runtime.Value {
	once_monadTransIdentityT.Do(func() {
		cache_monadTransIdentityT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_IdentityT()
}))
	})
	return cache_monadTransIdentityT
}

var cache_monadThrowIdentityT gopurs_runtime.Value
var once_monadThrowIdentityT sync.Once
func Get_monadThrowIdentityT() gopurs_runtime.Value {
	once_monadThrowIdentityT.Do(func() {
		cache_monadThrowIdentityT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadThrowIdentityT((*Record_throwError_gopurs_runtime_Value)(dictMonadThrow_0_box.UnsafePtr)))}
})
	})
	return cache_monadThrowIdentityT
}

var cache_monadTellIdentityT gopurs_runtime.Value
var once_monadTellIdentityT sync.Once
func Get_monadTellIdentityT() gopurs_runtime.Value {
	once_monadTellIdentityT.Do(func() {
		cache_monadTellIdentityT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadTellIdentityT((*Record_tell_gopurs_runtime_Value)(dictMonadTell_0_box.UnsafePtr)))}
})
	})
	return cache_monadTellIdentityT
}

var cache_monadStateIdentityT gopurs_runtime.Value
var once_monadStateIdentityT sync.Once
func Get_monadStateIdentityT() gopurs_runtime.Value {
	once_monadStateIdentityT.Do(func() {
		cache_monadStateIdentityT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadStateIdentityT((*Record_state_gopurs_runtime_Value)(dictMonadState_0_box.UnsafePtr)))}
})
	})
	return cache_monadStateIdentityT
}

var cache_monadRecIdentityT gopurs_runtime.Value
var once_monadRecIdentityT sync.Once
func Get_monadRecIdentityT() gopurs_runtime.Value {
	once_monadRecIdentityT.Do(func() {
		cache_monadRecIdentityT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadRecIdentityT((*Record_tailRecM_gopurs_runtime_Value)(dictMonadRec_0_box.UnsafePtr)))}
})
	})
	return cache_monadRecIdentityT
}

var cache_monadReaderIdentityT gopurs_runtime.Value
var once_monadReaderIdentityT sync.Once
func Get_monadReaderIdentityT() gopurs_runtime.Value {
	once_monadReaderIdentityT.Do(func() {
		cache_monadReaderIdentityT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadReaderIdentityT((*Record_local_gopurs_runtime_Value)(dictMonadReader_0_box.UnsafePtr)))}
})
	})
	return cache_monadReaderIdentityT
}

var cache_monadPlusIdentityT gopurs_runtime.Value
var once_monadPlusIdentityT sync.Once
func Get_monadPlusIdentityT() gopurs_runtime.Value {
	once_monadPlusIdentityT.Do(func() {
		cache_monadPlusIdentityT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadPlusIdentityT((*Record_)(dictMonadPlus_0_box.UnsafePtr)))}
})
	})
	return cache_monadPlusIdentityT
}

var cache_monadIdentityT gopurs_runtime.Value
var once_monadIdentityT sync.Once
func Get_monadIdentityT() gopurs_runtime.Value {
	once_monadIdentityT.Do(func() {
		cache_monadIdentityT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadIdentityT((*Record_)(dictMonad_0_box.UnsafePtr)))}
})
	})
	return cache_monadIdentityT
}

var cache_monadErrorIdentityT gopurs_runtime.Value
var once_monadErrorIdentityT sync.Once
func Get_monadErrorIdentityT() gopurs_runtime.Value {
	once_monadErrorIdentityT.Do(func() {
		cache_monadErrorIdentityT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadErrorIdentityT((*Record_catchError_gopurs_runtime_Value)(dictMonadError_0_box.UnsafePtr)))}
})
	})
	return cache_monadErrorIdentityT
}

var cache_monadEffectIdentityT gopurs_runtime.Value
var once_monadEffectIdentityT sync.Once
func Get_monadEffectIdentityT() gopurs_runtime.Value {
	once_monadEffectIdentityT.Do(func() {
		cache_monadEffectIdentityT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadEffectIdentityT((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr)))}
})
	})
	return cache_monadEffectIdentityT
}

var cache_monadContIdentityT gopurs_runtime.Value
var once_monadContIdentityT sync.Once
func Get_monadContIdentityT() gopurs_runtime.Value {
	once_monadContIdentityT.Do(func() {
		cache_monadContIdentityT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadContIdentityT((*Record_callCC_gopurs_runtime_Value)(dictMonadCont_0_box.UnsafePtr)))}
})
	})
	return cache_monadContIdentityT
}

var cache_monadAskIdentityT gopurs_runtime.Value
var once_monadAskIdentityT sync.Once
func Get_monadAskIdentityT() gopurs_runtime.Value {
	once_monadAskIdentityT.Do(func() {
		cache_monadAskIdentityT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadAskIdentityT((*Record_ask_gopurs_runtime_Value)(dictMonadAsk_0_box.UnsafePtr)))}
})
	})
	return cache_monadAskIdentityT
}

var cache_mapIdentityT gopurs_runtime.Value
var once_mapIdentityT sync.Once
func Get_mapIdentityT() gopurs_runtime.Value {
	once_mapIdentityT.Do(func() {
		cache_mapIdentityT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapIdentityT(f_0_box, v_1_box)
})
	})
	return cache_mapIdentityT
}

var cache_functorIdentityT gopurs_runtime.Value
var once_functorIdentityT sync.Once
func Get_functorIdentityT() gopurs_runtime.Value {
	once_functorIdentityT.Do(func() {
		cache_functorIdentityT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_functorIdentityT((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr)))}
})
	})
	return cache_functorIdentityT
}

var cache_foldableIdentityT gopurs_runtime.Value
var once_foldableIdentityT sync.Once
func Get_foldableIdentityT() gopurs_runtime.Value {
	once_foldableIdentityT.Do(func() {
		cache_foldableIdentityT = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_foldableIdentityT((*Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value)(dictFoldable_0_box.UnsafePtr)))}
})
	})
	return cache_foldableIdentityT
}

var cache_extendIdentityI gopurs_runtime.Value
var once_extendIdentityI sync.Once
func Get_extendIdentityI() gopurs_runtime.Value {
	once_extendIdentityI.Do(func() {
		cache_extendIdentityI = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendIdentityI((*Record_extend_gopurs_runtime_Value)(dictExtend_0_box.UnsafePtr))
})
	})
	return cache_extendIdentityI
}

var cache_eqIdentityT gopurs_runtime.Value
var once_eqIdentityT sync.Once
func Get_eqIdentityT() gopurs_runtime.Value {
	once_eqIdentityT.Do(func() {
		cache_eqIdentityT = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqIdentityT((*Record_eq1_gopurs_runtime_Value)(dictEq1_0_box.UnsafePtr), (*Record_eq_gopurs_runtime_Value)(dictEq_1_box.UnsafePtr))
})
	})
	return cache_eqIdentityT
}

var cache_ordIdentityT gopurs_runtime.Value
var once_ordIdentityT sync.Once
func Get_ordIdentityT() gopurs_runtime.Value {
	once_ordIdentityT.Do(func() {
		cache_ordIdentityT = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordIdentityT((*Record_compare1_gopurs_runtime_Value)(dictOrd1_0_box.UnsafePtr))
})
	})
	return cache_ordIdentityT
}

var cache_eq1IdentityT gopurs_runtime.Value
var once_eq1IdentityT sync.Once
func Get_eq1IdentityT() gopurs_runtime.Value {
	once_eq1IdentityT.Do(func() {
		cache_eq1IdentityT = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1IdentityT((*Record_eq1_gopurs_runtime_Value)(dictEq1_0_box.UnsafePtr))
})
	})
	return cache_eq1IdentityT
}

var cache_ord1IdentityT gopurs_runtime.Value
var once_ord1IdentityT sync.Once
func Get_ord1IdentityT() gopurs_runtime.Value {
	once_ord1IdentityT.Do(func() {
		cache_ord1IdentityT = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1IdentityT((*Record_compare1_gopurs_runtime_Value)(dictOrd1_0_box.UnsafePtr))
})
	})
	return cache_ord1IdentityT
}

var cache_comonadIdentityT gopurs_runtime.Value
var once_comonadIdentityT sync.Once
func Get_comonadIdentityT() gopurs_runtime.Value {
	once_comonadIdentityT.Do(func() {
		cache_comonadIdentityT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadIdentityT((*Record_extract_gopurs_runtime_Value)(dictComonad_0_box.UnsafePtr))
})
	})
	return cache_comonadIdentityT
}

var cache_bindIdentityT gopurs_runtime.Value
var once_bindIdentityT sync.Once
func Get_bindIdentityT() gopurs_runtime.Value {
	once_bindIdentityT.Do(func() {
		cache_bindIdentityT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_bindIdentityT((*Record_bind_gopurs_runtime_Value)(dictBind_0_box.UnsafePtr)))}
})
	})
	return cache_bindIdentityT
}

var cache_applyIdentityT gopurs_runtime.Value
var once_applyIdentityT sync.Once
func Get_applyIdentityT() gopurs_runtime.Value {
	once_applyIdentityT.Do(func() {
		cache_applyIdentityT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_applyIdentityT((*Record_apply_gopurs_runtime_Value)(dictApply_0_box.UnsafePtr)))}
})
	})
	return cache_applyIdentityT
}

var cache_applicativeIdentityT gopurs_runtime.Value
var once_applicativeIdentityT sync.Once
func Get_applicativeIdentityT() gopurs_runtime.Value {
	once_applicativeIdentityT.Do(func() {
		cache_applicativeIdentityT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_applicativeIdentityT((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr)))}
})
	})
	return cache_applicativeIdentityT
}

var cache_alternativeIdentityT gopurs_runtime.Value
var once_alternativeIdentityT sync.Once
func Get_alternativeIdentityT() gopurs_runtime.Value {
	once_alternativeIdentityT.Do(func() {
		cache_alternativeIdentityT = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_alternativeIdentityT((*Record_)(dictAlternative_0_box.UnsafePtr)))}
})
	})
	return cache_alternativeIdentityT
}

var cache_altIdentityT gopurs_runtime.Value
var once_altIdentityT sync.Once
func Get_altIdentityT() gopurs_runtime.Value {
	once_altIdentityT.Do(func() {
		cache_altIdentityT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_altIdentityT((*Record_alt_gopurs_runtime_Value)(dictAlt_0_box.UnsafePtr)))}
})
	})
	return cache_altIdentityT
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

func Call_IdentityT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_monadSTIdentityT(dictMonadST_0_loop *Record_liftST_gopurs_runtime_Value) *Record_liftST_gopurs_runtime_Value {
var dictMonadST_0 *Record_liftST_gopurs_runtime_Value = dictMonadST_0_loop
_ = dictMonadST_0
return dictMonadST_0
}

func Call_traversableIdentityT(dictTraversable_0_loop *Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value) *Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value {
var dictTraversable_0 *Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value = dictTraversable_0_loop
_ = dictTraversable_0
return dictTraversable_0
}

func Call_runIdentityT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_plusIdentityT(dictPlus_0_loop *Record_empty_gopurs_runtime_Value) *Record_empty_gopurs_runtime_Value {
var dictPlus_0 *Record_empty_gopurs_runtime_Value = dictPlus_0_loop
_ = dictPlus_0
return dictPlus_0
}

func Call_monadWriterIdentityT(dictMonadWriter_0_loop *Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value) *Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value {
var dictMonadWriter_0 *Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
return dictMonadWriter_0
}

func Call_monadThrowIdentityT(dictMonadThrow_0_loop *Record_throwError_gopurs_runtime_Value) *Record_throwError_gopurs_runtime_Value {
var dictMonadThrow_0 *Record_throwError_gopurs_runtime_Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
return dictMonadThrow_0
}

func Call_monadTellIdentityT(dictMonadTell_0_loop *Record_tell_gopurs_runtime_Value) *Record_tell_gopurs_runtime_Value {
var dictMonadTell_0 *Record_tell_gopurs_runtime_Value = dictMonadTell_0_loop
_ = dictMonadTell_0
return dictMonadTell_0
}

func Call_monadStateIdentityT(dictMonadState_0_loop *Record_state_gopurs_runtime_Value) *Record_state_gopurs_runtime_Value {
var dictMonadState_0 *Record_state_gopurs_runtime_Value = dictMonadState_0_loop
_ = dictMonadState_0
return dictMonadState_0
}

func Call_monadRecIdentityT(dictMonadRec_0_loop *Record_tailRecM_gopurs_runtime_Value) *Record_tailRecM_gopurs_runtime_Value {
var dictMonadRec_0 *Record_tailRecM_gopurs_runtime_Value = dictMonadRec_0_loop
_ = dictMonadRec_0
return dictMonadRec_0
}

func Call_monadReaderIdentityT(dictMonadReader_0_loop *Record_local_gopurs_runtime_Value) *Record_local_gopurs_runtime_Value {
var dictMonadReader_0 *Record_local_gopurs_runtime_Value = dictMonadReader_0_loop
_ = dictMonadReader_0
return dictMonadReader_0
}

func Call_monadPlusIdentityT(dictMonadPlus_0_loop *Record_) *Record_ {
var dictMonadPlus_0 *Record_ = dictMonadPlus_0_loop
_ = dictMonadPlus_0
return dictMonadPlus_0
}

func Call_monadIdentityT(dictMonad_0_loop *Record_) *Record_ {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
return dictMonad_0
}

func Call_monadErrorIdentityT(dictMonadError_0_loop *Record_catchError_gopurs_runtime_Value) *Record_catchError_gopurs_runtime_Value {
var dictMonadError_0 *Record_catchError_gopurs_runtime_Value = dictMonadError_0_loop
_ = dictMonadError_0
return dictMonadError_0
}

func Call_monadEffectIdentityT(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) *Record_liftEffect_gopurs_runtime_Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return dictMonadEffect_0
}

func Call_monadContIdentityT(dictMonadCont_0_loop *Record_callCC_gopurs_runtime_Value) *Record_callCC_gopurs_runtime_Value {
var dictMonadCont_0 *Record_callCC_gopurs_runtime_Value = dictMonadCont_0_loop
_ = dictMonadCont_0
return dictMonadCont_0
}

func Call_monadAskIdentityT(dictMonadAsk_0_loop *Record_ask_gopurs_runtime_Value) *Record_ask_gopurs_runtime_Value {
var dictMonadAsk_0 *Record_ask_gopurs_runtime_Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
return dictMonadAsk_0
}

func Call_mapIdentityT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_functorIdentityT(dictFunctor_0_loop *Record_map__gopurs_runtime_Value) *Record_map__gopurs_runtime_Value {
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
return dictFunctor_0
}

func Call_foldableIdentityT(dictFoldable_0_loop *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value) *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value {
var dictFoldable_0 *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value = dictFoldable_0_loop
_ = dictFoldable_0
return dictFoldable_0
}

func Call_extendIdentityI(dictExtend_0_loop *Record_extend_gopurs_runtime_Value) gopurs_runtime.Value {
var dictExtend_0 *Record_extend_gopurs_runtime_Value = dictExtend_0_loop
_ = dictExtend_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictExtend_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictExtend_0.extend, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, x_4)
}), v_3)
}))
}

func Call_eqIdentityT(dictEq1_0_loop *Record_eq1_gopurs_runtime_Value, dictEq_1_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq1_0 *Record_eq1_gopurs_runtime_Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 *Record_eq_gopurs_runtime_Value = dictEq_1_loop
_ = dictEq_1
eq11_2_0 := gopurs_runtime.Apply(dictEq1_0.eq1, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_1)})
_ = eq11_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_2_0, x_3, y_4)
}))
}

func Call_ordIdentityT(dictOrd1_0_loop *Record_compare1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd1_0 *Record_compare1_gopurs_runtime_Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(dictOrd1_0.compare1, dictOrd_2)
_ = compare11_3_1
eq11_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_4_2
eqIdentityT2_5_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_4_2, x_5, y_6)
}))
_ = eqIdentityT2_5_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqIdentityT2_5_3
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_3_1, x_6, y_7)
}))
})
}

func Call_eq1IdentityT(dictEq1_0_loop *Record_eq1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq1_0 *Record_eq1_gopurs_runtime_Value = dictEq1_0_loop
_ = dictEq1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictEq1_0.eq1, dictEq_1)
}))
}

func Call_ord1IdentityT(dictOrd1_0_loop *Record_compare1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd1_0 *Record_compare1_gopurs_runtime_Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1IdentityT1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), dictEq_3)
}))
_ = eq1IdentityT1_3_2
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1IdentityT1_3_2
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_5_3 := gopurs_runtime.Apply(dictOrd1_0.compare1, dictOrd_4)
_ = compare11_5_3
eq11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_4, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_6_4
eqIdentityT2_7_5 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_6_4, x_7, y_8)
}))
_ = eqIdentityT2_7_5
return gopurs_runtime.RecordGet(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqIdentityT2_7_5
}), gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_5_3, x_8, y_9)
})), "compare")
}))
}

func Call_comonadIdentityT(dictComonad_0_loop *Record_extract_gopurs_runtime_Value) gopurs_runtime.Value {
var dictComonad_0 *Record_extract_gopurs_runtime_Value = dictComonad_0_loop
_ = dictComonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictComonad_0)}, "Extend0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
extendIdentityI1_3_2 := gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "extend"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, x_5)
}), v_4)
}))
_ = extendIdentityI1_3_2
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendIdentityI1_3_2
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictComonad_0.extract, x_4)
}))
}

func Call_bindIdentityT(dictBind_0_loop *Record_bind_gopurs_runtime_Value) *Record_bind_gopurs_runtime_Value {
var dictBind_0 *Record_bind_gopurs_runtime_Value = dictBind_0_loop
_ = dictBind_0
return dictBind_0
}

func Call_applyIdentityT(dictApply_0_loop *Record_apply_gopurs_runtime_Value) *Record_apply_gopurs_runtime_Value {
var dictApply_0 *Record_apply_gopurs_runtime_Value = dictApply_0_loop
_ = dictApply_0
return dictApply_0
}

func Call_applicativeIdentityT(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) *Record_pure_gopurs_runtime_Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
return dictApplicative_0
}

func Call_alternativeIdentityT(dictAlternative_0_loop *Record_) *Record_ {
var dictAlternative_0 *Record_ = dictAlternative_0_loop
_ = dictAlternative_0
return dictAlternative_0
}

func Call_altIdentityT(dictAlt_0_loop *Record_alt_gopurs_runtime_Value) *Record_alt_gopurs_runtime_Value {
var dictAlt_0 *Record_alt_gopurs_runtime_Value = dictAlt_0_loop
_ = dictAlt_0
return dictAlt_0
}


