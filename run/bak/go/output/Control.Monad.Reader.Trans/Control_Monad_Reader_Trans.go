package Control_Monad_Reader_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_ReaderT gopurs_runtime.Value
var once_ReaderT sync.Once
func Get_ReaderT() gopurs_runtime.Value {
	once_ReaderT.Do(func() {
		cache_ReaderT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ReaderT(x_0_box)
})
	})
	return cache_ReaderT
}

var cache_withReaderT gopurs_runtime.Value
var once_withReaderT sync.Once
func Get_withReaderT() gopurs_runtime.Value {
	once_withReaderT.Do(func() {
		cache_withReaderT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withReaderT(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_withReaderT
}

var cache_runReaderT gopurs_runtime.Value
var once_runReaderT sync.Once
func Get_runReaderT() gopurs_runtime.Value {
	once_runReaderT.Do(func() {
		cache_runReaderT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runReaderT(v_0_box)
})
	})
	return cache_runReaderT
}

var cache_newtypeReaderT gopurs_runtime.Value
var once_newtypeReaderT sync.Once
func Get_newtypeReaderT() gopurs_runtime.Value {
	once_newtypeReaderT.Do(func() {
		cache_newtypeReaderT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeReaderT
}

var cache_monadTransReaderT gopurs_runtime.Value
var once_monadTransReaderT sync.Once
func Get_monadTransReaderT() gopurs_runtime.Value {
	once_monadTransReaderT.Do(func() {
		cache_monadTransReaderT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func3(func(dictMonad_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
	})
	return cache_monadTransReaderT
}

var cache_mapReaderT gopurs_runtime.Value
var once_mapReaderT sync.Once
func Get_mapReaderT() gopurs_runtime.Value {
	once_mapReaderT.Do(func() {
		cache_mapReaderT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapReaderT(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_mapReaderT
}

var cache_functorReaderT gopurs_runtime.Value
var once_functorReaderT sync.Once
func Get_functorReaderT() gopurs_runtime.Value {
	once_functorReaderT.Do(func() {
		cache_functorReaderT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorReaderT((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr))
})
	})
	return cache_functorReaderT
}

var cache_distributiveReaderT gopurs_runtime.Value
var once_distributiveReaderT sync.Once
func Get_distributiveReaderT() gopurs_runtime.Value {
	once_distributiveReaderT.Do(func() {
		cache_distributiveReaderT = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distributiveReaderT((*Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value)(dictDistributive_0_box.UnsafePtr))
})
	})
	return cache_distributiveReaderT
}

var cache_applyReaderT gopurs_runtime.Value
var once_applyReaderT sync.Once
func Get_applyReaderT() gopurs_runtime.Value {
	once_applyReaderT.Do(func() {
		cache_applyReaderT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyReaderT((*Record_apply_gopurs_runtime_Value)(dictApply_0_box.UnsafePtr))
})
	})
	return cache_applyReaderT
}

var cache_bindReaderT gopurs_runtime.Value
var once_bindReaderT sync.Once
func Get_bindReaderT() gopurs_runtime.Value {
	once_bindReaderT.Do(func() {
		cache_bindReaderT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindReaderT((*Record_bind_gopurs_runtime_Value)(dictBind_0_box.UnsafePtr))
})
	})
	return cache_bindReaderT
}

var cache_semigroupReaderT gopurs_runtime.Value
var once_semigroupReaderT sync.Once
func Get_semigroupReaderT() gopurs_runtime.Value {
	once_semigroupReaderT.Do(func() {
		cache_semigroupReaderT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupReaderT((*Record_apply_gopurs_runtime_Value)(dictApply_0_box.UnsafePtr))
})
	})
	return cache_semigroupReaderT
}

var cache_applicativeReaderT gopurs_runtime.Value
var once_applicativeReaderT sync.Once
func Get_applicativeReaderT() gopurs_runtime.Value {
	once_applicativeReaderT.Do(func() {
		cache_applicativeReaderT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeReaderT((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_applicativeReaderT
}

var cache_monadReaderT gopurs_runtime.Value
var once_monadReaderT sync.Once
func Get_monadReaderT() gopurs_runtime.Value {
	once_monadReaderT.Do(func() {
		cache_monadReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_monadReaderT
}

var cache_monadAskReaderT gopurs_runtime.Value
var once_monadAskReaderT sync.Once
func Get_monadAskReaderT() gopurs_runtime.Value {
	once_monadAskReaderT.Do(func() {
		cache_monadAskReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskReaderT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_monadAskReaderT
}

var cache_monadReaderReaderT gopurs_runtime.Value
var once_monadReaderReaderT sync.Once
func Get_monadReaderReaderT() gopurs_runtime.Value {
	once_monadReaderReaderT.Do(func() {
		cache_monadReaderReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderReaderT((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_monadReaderReaderT
}

var cache_monadContReaderT gopurs_runtime.Value
var once_monadContReaderT sync.Once
func Get_monadContReaderT() gopurs_runtime.Value {
	once_monadContReaderT.Do(func() {
		cache_monadContReaderT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContReaderT((*Record_callCC_gopurs_runtime_Value)(dictMonadCont_0_box.UnsafePtr))
})
	})
	return cache_monadContReaderT
}

var cache_monadEffectReader gopurs_runtime.Value
var once_monadEffectReader sync.Once
func Get_monadEffectReader() gopurs_runtime.Value {
	once_monadEffectReader.Do(func() {
		cache_monadEffectReader = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectReader((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_monadEffectReader
}

var cache_monadRecReaderT gopurs_runtime.Value
var once_monadRecReaderT sync.Once
func Get_monadRecReaderT() gopurs_runtime.Value {
	once_monadRecReaderT.Do(func() {
		cache_monadRecReaderT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadRecReaderT((*Record_tailRecM_gopurs_runtime_Value)(dictMonadRec_0_box.UnsafePtr))
})
	})
	return cache_monadRecReaderT
}

var cache_monadStateReaderT gopurs_runtime.Value
var once_monadStateReaderT sync.Once
func Get_monadStateReaderT() gopurs_runtime.Value {
	once_monadStateReaderT.Do(func() {
		cache_monadStateReaderT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateReaderT((*Record_state_gopurs_runtime_Value)(dictMonadState_0_box.UnsafePtr))
})
	})
	return cache_monadStateReaderT
}

var cache_monadTellReaderT gopurs_runtime.Value
var once_monadTellReaderT sync.Once
func Get_monadTellReaderT() gopurs_runtime.Value {
	once_monadTellReaderT.Do(func() {
		cache_monadTellReaderT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTellReaderT((*Record_tell_gopurs_runtime_Value)(dictMonadTell_0_box.UnsafePtr))
})
	})
	return cache_monadTellReaderT
}

var cache_monadWriterReaderT gopurs_runtime.Value
var once_monadWriterReaderT sync.Once
func Get_monadWriterReaderT() gopurs_runtime.Value {
	once_monadWriterReaderT.Do(func() {
		cache_monadWriterReaderT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadWriterReaderT((*Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value)(dictMonadWriter_0_box.UnsafePtr))
})
	})
	return cache_monadWriterReaderT
}

var cache_monadThrowReaderT gopurs_runtime.Value
var once_monadThrowReaderT sync.Once
func Get_monadThrowReaderT() gopurs_runtime.Value {
	once_monadThrowReaderT.Do(func() {
		cache_monadThrowReaderT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadThrowReaderT((*Record_throwError_gopurs_runtime_Value)(dictMonadThrow_0_box.UnsafePtr))
})
	})
	return cache_monadThrowReaderT
}

var cache_monadErrorReaderT gopurs_runtime.Value
var once_monadErrorReaderT sync.Once
func Get_monadErrorReaderT() gopurs_runtime.Value {
	once_monadErrorReaderT.Do(func() {
		cache_monadErrorReaderT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadErrorReaderT((*Record_catchError_gopurs_runtime_Value)(dictMonadError_0_box.UnsafePtr))
})
	})
	return cache_monadErrorReaderT
}

var cache_monadSTReaderT gopurs_runtime.Value
var once_monadSTReaderT sync.Once
func Get_monadSTReaderT() gopurs_runtime.Value {
	once_monadSTReaderT.Do(func() {
		cache_monadSTReaderT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTReaderT((*Record_liftST_gopurs_runtime_Value)(dictMonadST_0_box.UnsafePtr))
})
	})
	return cache_monadSTReaderT
}

var cache_monoidReaderT gopurs_runtime.Value
var once_monoidReaderT sync.Once
func Get_monoidReaderT() gopurs_runtime.Value {
	once_monoidReaderT.Do(func() {
		cache_monoidReaderT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidReaderT((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_monoidReaderT
}

var cache_altReaderT gopurs_runtime.Value
var once_altReaderT sync.Once
func Get_altReaderT() gopurs_runtime.Value {
	once_altReaderT.Do(func() {
		cache_altReaderT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altReaderT((*Record_alt_gopurs_runtime_Value)(dictAlt_0_box.UnsafePtr))
})
	})
	return cache_altReaderT
}

var cache_plusReaderT gopurs_runtime.Value
var once_plusReaderT sync.Once
func Get_plusReaderT() gopurs_runtime.Value {
	once_plusReaderT.Do(func() {
		cache_plusReaderT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusReaderT((*Record_empty_gopurs_runtime_Value)(dictPlus_0_box.UnsafePtr))
})
	})
	return cache_plusReaderT
}

var cache_alternativeReaderT gopurs_runtime.Value
var once_alternativeReaderT sync.Once
func Get_alternativeReaderT() gopurs_runtime.Value {
	once_alternativeReaderT.Do(func() {
		cache_alternativeReaderT = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeReaderT((*Record_)(dictAlternative_0_box.UnsafePtr))
})
	})
	return cache_alternativeReaderT
}

var cache_monadPlusReaderT gopurs_runtime.Value
var once_monadPlusReaderT sync.Once
func Get_monadPlusReaderT() gopurs_runtime.Value {
	once_monadPlusReaderT.Do(func() {
		cache_monadPlusReaderT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusReaderT((*Record_)(dictMonadPlus_0_box.UnsafePtr))
})
	})
	return cache_monadPlusReaderT
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

func Call_ReaderT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_withReaderT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}

func Call_runReaderT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_mapReaderT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_functorReaderT(dictFunctor_0_loop *Record_map__gopurs_runtime_Value) gopurs_runtime.Value {
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictFunctor_0.map_, x_1)
_ = __local_var_2_0
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(v_3, x_4))
})
}))
}

func Call_distributiveReaderT(dictDistributive_0_loop *Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value) gopurs_runtime.Value {
distributiveReaderT:
for {
if false { continue distributiveReaderT }
var dictDistributive_0 *Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value = dictDistributive_0_loop
_ = dictDistributive_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictDistributive_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorReaderT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
}))
_ = functorReaderT1_2_1
return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
}), gopurs_runtime.Func2(func(dictFunctor_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_distributiveReaderT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictDistributive_0)}), "distribute"), dictFunctor_3)
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_3, "map"), f_4)
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(__local_var_6_4, x_7))
})
}), gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
collect1_4_5 := gopurs_runtime.Apply(dictDistributive_0.collect, dictFunctor_3)
_ = collect1_4_5
return gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(collect1_4_5, gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(r_7, e_6)
}), a_5)
})
}))
}
}

func Call_applyReaderT(dictApply_0_loop *Record_apply_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApply_0 *Record_apply_gopurs_runtime_Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApply_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorReaderT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
}))
_ = functorReaderT1_2_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.apply, gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
}))
}

func Call_bindReaderT(dictBind_0_loop *Record_bind_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBind_0 *Record_bind_gopurs_runtime_Value = dictBind_0_loop
_ = dictBind_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBind_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorReaderT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
}))
_ = functorReaderT1_3_3
applyReaderT1_3_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
}))
_ = applyReaderT1_3_2
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_3_2
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictBind_0.bind, gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_5, a_7, r_6)
}))
}))
}

func Call_semigroupReaderT(dictApply_0_loop *Record_apply_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApply_0 *Record_apply_gopurs_runtime_Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApply_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), __local_var_3_1)
_ = __local_var_6_2
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.apply, gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(a_4, r_7)), gopurs_runtime.Apply(b_5, r_7))
})
}))
})
}

func Call_applicativeReaderT(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorReaderT1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), x_3)
_ = __local_var_4_3
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(v_5, x_6))
})
}))
_ = functorReaderT1_3_2
applyReaderT1_4_4 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_2
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
}))
_ = applyReaderT1_4_4
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_4_4
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(dictApplicative_0.pure, x_5)
_ = __local_var_6_5
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_5
})
}))
}

func Call_monadReaderT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
functorReaderT1_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), x_4)
_ = __local_var_5_5
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, gopurs_runtime.Apply(v_6, x_7))
})
}))
_ = functorReaderT1_4_4
applyReaderT1_5_6 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_4_4
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
}))
_ = applyReaderT1_5_6
applicativeReaderT1_3_2 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_5_6
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), x_6)
_ = __local_var_7_7
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_7
})
}))
_ = applicativeReaderT1_3_2
bindReaderT1_4_8 := gopurs_runtime.Apply(Get_bindReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = bindReaderT1_4_8
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bindReaderT1_4_8
}))
}

func Call_monadAskReaderT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = monadReaderT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"))
}

func Call_monadReaderReaderT(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = monadReaderT1_1_0
monadAskReaderT1_2_1 := gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"))
_ = monadAskReaderT1_2_1
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskReaderT1_2_1
}), Get_withReaderT())
}

func Call_monadContReaderT(dictMonadCont_0_loop *Record_callCC_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadCont_0 *Record_callCC_gopurs_runtime_Value = dictMonadCont_0_loop
_ = dictMonadCont_0
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadCont_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadReaderT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadCont_0.callCC, gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(c_4, x_5)
_ = __local_var_6_1
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_1
})
}), r_3)
}))
}))
}

func Call_monadEffectReader(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, x_4))
}))
}

func Call_monadRecReaderT(dictMonadRec_0_loop *Record_tailRecM_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadRec_0 *Record_tailRecM_gopurs_runtime_Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
monadReaderT1_4_3 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_4_3
}), gopurs_runtime.Func3(func(k_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.tailRecM, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(k_5, a_prime_8, r_7), pure_3_2)
}), a_6)
}))
}

func Call_monadStateReaderT(dictMonadState_0_loop *Record_state_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadState_0 *Record_state_gopurs_runtime_Value = dictMonadState_0_loop
_ = dictMonadState_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadState_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(dictMonadState_0.state, x_4))
}))
}

func Call_monadTellReaderT(dictMonadTell_0_loop *Record_tell_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadTell_0 *Record_tell_gopurs_runtime_Value = dictMonadTell_0_loop
_ = dictMonadTell_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadTell_0)}, "Monad1_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad1_1_0
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadTell_0)}, "Semigroup0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
monadReaderT1_3_2 := gopurs_runtime.Apply(Get_monadReaderT(), Monad1_1_0)
_ = monadReaderT1_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad1_1_0)
_ = __local_var_4_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(dictMonadTell_0.tell, x_5))
}))
}

func Call_monadWriterReaderT(dictMonadWriter_0_loop *Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadWriter_0 *Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monoid0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadWriter_0)}, "Monoid0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monoid0_1_0
monadTellReaderT1_2_1 := gopurs_runtime.Apply(Get_monadTellReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadWriter_0)}, "MonadTell1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadTellReaderT1_2_1
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellReaderT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_1_0
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadWriter_0.listen, gopurs_runtime.Apply(v_3, x_4))
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadWriter_0.pass, gopurs_runtime.Apply(v_3, x_4))
}))
}

func Call_monadThrowReaderT(dictMonadThrow_0_loop *Record_throwError_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadThrow_0 *Record_throwError_gopurs_runtime_Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadThrow_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(dictMonadThrow_0.throwError, x_4))
}))
}

func Call_monadErrorReaderT(dictMonadError_0_loop *Record_catchError_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadError_0 *Record_catchError_gopurs_runtime_Value = dictMonadError_0_loop
_ = dictMonadError_0
monadThrowReaderT1_1_0 := gopurs_runtime.Apply(Get_monadThrowReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadError_0)}, "MonadThrow0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadThrowReaderT1_1_0
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowReaderT1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadError_0.catchError, gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(h_3, e_5, r_4)
}))
}))
}

func Call_monadSTReaderT(dictMonadST_0_loop *Record_liftST_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadST_0 *Record_liftST_gopurs_runtime_Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadST_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(dictMonadST_0.liftST, x_4))
}))
}

func Call_monoidReaderT(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_4_3
semigroupReaderT2_4_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), __local_var_4_3)
_ = __local_var_7_4
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(a_5, r_8)), gopurs_runtime.Apply(b_6, r_8))
})
}))
_ = semigroupReaderT2_4_2
__local_var_5_5 := gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))
_ = __local_var_5_5
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupReaderT2_4_2
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_5
}))
})
}

func Call_altReaderT(dictAlt_0_loop *Record_alt_gopurs_runtime_Value) gopurs_runtime.Value {
var dictAlt_0 *Record_alt_gopurs_runtime_Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlt_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorReaderT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
}))
_ = functorReaderT1_2_1
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictAlt_0.alt, gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
}))
}

func Call_plusReaderT(dictPlus_0_loop *Record_empty_gopurs_runtime_Value) gopurs_runtime.Value {
var dictPlus_0 *Record_empty_gopurs_runtime_Value = dictPlus_0_loop
_ = dictPlus_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictPlus_0)}, "Alt0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorReaderT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
}))
_ = functorReaderT1_3_3
altReaderT1_3_2 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "alt"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
}))
_ = altReaderT1_3_2
__local_var_4_5 := dictPlus_0.empty
_ = __local_var_4_5
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_3_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_5
}))
}

func Call_alternativeReaderT(dictAlternative_0_loop *Record_) gopurs_runtime.Value {
var dictAlternative_0 *Record_ = dictAlternative_0_loop
_ = dictAlternative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
functorReaderT1_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), x_4)
_ = __local_var_5_5
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, gopurs_runtime.Apply(v_6, x_7))
})
}))
_ = functorReaderT1_4_4
applyReaderT1_5_6 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_4_4
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
}))
_ = applyReaderT1_5_6
applicativeReaderT1_3_2 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_5_6
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), x_6)
_ = __local_var_7_7
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_7
})
}))
_ = applicativeReaderT1_3_2
__local_var_4_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, "Plus1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_4_8
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_9
__local_var_6_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_9, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_11
functorReaderT1_7_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_11, "map"), x_7)
_ = __local_var_8_13
return gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_13, gopurs_runtime.Apply(v_9, x_10))
})
}))
_ = functorReaderT1_7_12
altReaderT1_8_14 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_7_12
}), gopurs_runtime.Func3(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_9, "alt"), gopurs_runtime.Apply(v_8, r_10), gopurs_runtime.Apply(v1_9, r_10))
}))
_ = altReaderT1_8_14
__local_var_9_15 := gopurs_runtime.RecordGet(__local_var_4_8, "empty")
_ = __local_var_9_15
plusReaderT1_6_10 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_8_14
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_9_15
}))
_ = plusReaderT1_6_10
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusReaderT1_6_10
}))
}

func Call_monadPlusReaderT(dictMonadPlus_0_loop *Record_) gopurs_runtime.Value {
var dictMonadPlus_0 *Record_ = dictMonadPlus_0_loop
_ = dictMonadPlus_0
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadPlus_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadReaderT1_1_0
alternativeReaderT1_2_1 := gopurs_runtime.Apply(Get_alternativeReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadPlus_0)}, "Alternative1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = alternativeReaderT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeReaderT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}))
}


