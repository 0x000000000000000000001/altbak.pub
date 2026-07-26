package Control_Parallel_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_Maybe_Trans "gopurs/output/Control.Monad.Maybe.Trans"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Control_Monad_Except_Trans "gopurs/output/Control.Monad.Except.Trans"
	pkg_Data_Functor_Costar "gopurs/output/Data.Functor.Costar"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	unsafe "unsafe"
)

var cache_ParCont gopurs_runtime.Value
var once_ParCont sync.Once
func Get_ParCont() gopurs_runtime.Value {
	once_ParCont.Do(func() {
		cache_ParCont = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ParCont(x_0_box)
})
	})
	return cache_ParCont
}

var cache_sequential gopurs_runtime.Value
var once_sequential sync.Once
func Get_sequential() gopurs_runtime.Value {
	once_sequential.Do(func() {
		cache_sequential = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequential((*Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_sequential
}

var cache_parallel gopurs_runtime.Value
var once_parallel sync.Once
func Get_parallel() gopurs_runtime.Value {
	once_parallel.Do(func() {
		cache_parallel = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel((*Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_parallel
}

var cache_newtypeParCont gopurs_runtime.Value
var once_newtypeParCont sync.Once
func Get_newtypeParCont() gopurs_runtime.Value {
	once_newtypeParCont.Do(func() {
		cache_newtypeParCont = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeParCont
}

var cache_monadParWriterT gopurs_runtime.Value
var once_monadParWriterT sync.Once
func Get_monadParWriterT() gopurs_runtime.Value {
	once_monadParWriterT.Do(func() {
		cache_monadParWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParWriterT((*Record_mempty_gopurs_runtime_Value)(dictMonoid_0_box.UnsafePtr))
})
	})
	return cache_monadParWriterT
}

var cache_monadParStar gopurs_runtime.Value
var once_monadParStar sync.Once
func Get_monadParStar() gopurs_runtime.Value {
	once_monadParStar.Do(func() {
		cache_monadParStar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParStar((*Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value)(dictParallel_0_box.UnsafePtr))
})
	})
	return cache_monadParStar
}

var cache_monadParReaderT gopurs_runtime.Value
var once_monadParReaderT sync.Once
func Get_monadParReaderT() gopurs_runtime.Value {
	once_monadParReaderT.Do(func() {
		cache_monadParReaderT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParReaderT((*Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value)(dictParallel_0_box.UnsafePtr))
})
	})
	return cache_monadParReaderT
}

var cache_monadParMaybeT gopurs_runtime.Value
var once_monadParMaybeT sync.Once
func Get_monadParMaybeT() gopurs_runtime.Value {
	once_monadParMaybeT.Do(func() {
		cache_monadParMaybeT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParMaybeT((*Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value)(dictParallel_0_box.UnsafePtr))
})
	})
	return cache_monadParMaybeT
}

var cache_monadParExceptT gopurs_runtime.Value
var once_monadParExceptT sync.Once
func Get_monadParExceptT() gopurs_runtime.Value {
	once_monadParExceptT.Do(func() {
		cache_monadParExceptT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParExceptT((*Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value)(dictParallel_0_box.UnsafePtr))
})
	})
	return cache_monadParExceptT
}

var cache_monadParCostar gopurs_runtime.Value
var once_monadParCostar sync.Once
func Get_monadParCostar() gopurs_runtime.Value {
	once_monadParCostar.Do(func() {
		cache_monadParCostar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParCostar((*Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value)(dictParallel_0_box.UnsafePtr))
})
	})
	return cache_monadParCostar
}

var cache_monadParParCont gopurs_runtime.Value
var once_monadParParCont sync.Once
func Get_monadParParCont() gopurs_runtime.Value {
	once_monadParParCont.Do(func() {
		cache_monadParParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParParCont((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_monadParParCont
}

var cache_functorParCont gopurs_runtime.Value
var once_functorParCont sync.Once
func Get_functorParCont() gopurs_runtime.Value {
	once_functorParCont.Do(func() {
		cache_functorParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorParCont((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_functorParCont
}

var cache_applyParCont gopurs_runtime.Value
var once_applyParCont sync.Once
func Get_applyParCont() gopurs_runtime.Value {
	once_applyParCont.Do(func() {
		cache_applyParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyParCont((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_applyParCont
}

var cache_applicativeParCont gopurs_runtime.Value
var once_applicativeParCont sync.Once
func Get_applicativeParCont() gopurs_runtime.Value {
	once_applicativeParCont.Do(func() {
		cache_applicativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeParCont((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_applicativeParCont
}

var cache_altParCont gopurs_runtime.Value
var once_altParCont sync.Once
func Get_altParCont() gopurs_runtime.Value {
	once_altParCont.Do(func() {
		cache_altParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altParCont((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_altParCont
}

var cache_plusParCont gopurs_runtime.Value
var once_plusParCont sync.Once
func Get_plusParCont() gopurs_runtime.Value {
	once_plusParCont.Do(func() {
		cache_plusParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusParCont((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_plusParCont
}

var cache_alternativeParCont gopurs_runtime.Value
var once_alternativeParCont sync.Once
func Get_alternativeParCont() gopurs_runtime.Value {
	once_alternativeParCont.Do(func() {
		cache_alternativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeParCont((*Record_liftEffect_gopurs_runtime_Value)(dictMonadEffect_0_box.UnsafePtr))
})
	})
	return cache_alternativeParCont
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

func Call_ParCont(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_sequential(dict_0_loop *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.sequential
}

func Call_parallel(dict_0_loop *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.parallel
}

func Call_monadParWriterT(dictMonoid_0_loop *Record_mempty_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonoid_0 *Record_mempty_gopurs_runtime_Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonoid_0)}, "Semigroup0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
applyWriterT_2_1 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_3_2
functorWriterT1_4_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_3_2, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply(f_4, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V1})}
}))
}))
_ = functorWriterT1_4_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_4_3
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_3_2, "map"), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v3_7.UnsafePtr).V0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v3_7.UnsafePtr).V1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v4_8.UnsafePtr).V1)})}
}), v_5), v1_6)
}))
})
_ = applyWriterT_2_1
return gopurs_runtime.Func(func(dictParallel_3 gopurs_runtime.Value) gopurs_runtime.Value {
applyWriterT1_4_4 := gopurs_runtime.UncurriedApp(applyWriterT_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "Apply0"), gopurs_runtime.Value{}))
_ = applyWriterT1_4_4
applyWriterT2_5_5 := gopurs_runtime.UncurriedApp(applyWriterT_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "Apply1"), gopurs_runtime.Value{}))
_ = applyWriterT2_5_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT1_4_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_5_5
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "parallel"), v_6)
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "sequential"), v_6)
}))
})
}

func Call_monadParStar(dictParallel_0_loop *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value) gopurs_runtime.Value {
var dictParallel_0 *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictParallel_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorStar1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
}))
_ = functorStar1_3_3
applyStar_3_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_3
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
}))
_ = applyStar_3_2
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictParallel_0)}, "Apply1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
functorStar1_6_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "map"), f_6)
_ = __local_var_8_9
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_9, gopurs_runtime.Apply(v_7, x_9))
})
}))
_ = functorStar1_6_8
applyStar1_6_7 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_6_8
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "apply"), gopurs_runtime.Apply(v_7, a_9), gopurs_runtime.Apply(v1_8, a_9))
}))
_ = applyStar1_6_7
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar_3_2
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_6_7
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.parallel, gopurs_runtime.Apply(v_7, x_8))
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.sequential, gopurs_runtime.Apply(v_7, x_8))
}))
}

func Call_monadParReaderT(dictParallel_0_loop *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value) gopurs_runtime.Value {
var dictParallel_0 *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictParallel_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
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
applyReaderT_3_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
}))
_ = applyReaderT_3_2
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictParallel_0)}, "Apply1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
functorReaderT1_6_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "map"), x_6)
_ = __local_var_7_9
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_9, gopurs_runtime.Apply(v_8, x_9))
})
}))
_ = functorReaderT1_6_8
applyReaderT1_6_7 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_6_8
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "apply"), gopurs_runtime.Apply(v_7, r_9), gopurs_runtime.Apply(v1_8, r_9))
}))
_ = applyReaderT1_6_7
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT_3_2
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_6_7
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.parallel, gopurs_runtime.Apply(v_7, x_8))
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.sequential, gopurs_runtime.Apply(v_7, x_8))
}))
}

func Call_monadParMaybeT(dictParallel_0_loop *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value) gopurs_runtime.Value {
var dictParallel_0 *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictParallel_0)}, "Apply1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
functorCompose2_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), f_4), v_5)
}))
_ = functorCompose2_4_4
applyCompose_4_3 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_4_4
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), v_5), v1_6)
}))
_ = applyCompose_4_3
return gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
applyMaybeT_6_5 := gopurs_runtime.Apply(pkg_Control_Monad_Maybe_Trans.Get_applyMaybeT(), dictMonad_5)
_ = applyMaybeT_6_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMaybeT_6_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_4_3
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.parallel, v_7)
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.sequential, v_7)
}))
})
}

func Call_monadParExceptT(dictParallel_0_loop *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value) gopurs_runtime.Value {
var dictParallel_0 *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictParallel_0)}, "Apply1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Either.Get_applyEither(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
functorCompose2_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), f_4), v_5)
}))
_ = functorCompose2_4_4
applyCompose_4_3 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_4_4
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.RecordGet(pkg_Data_Either.Get_applyEither(), "apply"), v_5), v1_6)
}))
_ = applyCompose_4_3
return gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
applyExceptT_6_5 := gopurs_runtime.Apply(pkg_Control_Monad_Except_Trans.Get_applyExceptT(), dictMonad_5)
_ = applyExceptT_6_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyExceptT_6_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_4_3
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.parallel, v_7)
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.sequential, v_7)
}))
})
}

func Call_monadParCostar(dictParallel_0_loop *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value) gopurs_runtime.Value {
var dictParallel_0 *Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value = dictParallel_0_loop
_ = dictParallel_0
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(dictParallel_0.sequential, x_2))
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(dictParallel_0.parallel, x_2))
}))
}

func Call_monadParParCont(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
}))
_ = functorContT1_1_0
applyContT_2_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
}))
_ = applyContT_2_1
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)})
}), Get_ParCont(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return v_3
}))
}

func Call_functorParCont(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_monadParParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)}), "sequential"), x_2)
_ = __local_var_3_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_monadParParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)}), "parallel"), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_1, a_5))
}))
}))
}))
}

func Call_applyParCont(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
discard1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Bind1_1_0)
_ = discard1_2_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_functorParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)})
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_2 := gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
_ = __local_ref_2
return gopurs_runtime.Any(&__local_ref_2)
})), gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_3 := gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
_ = __local_ref_3
return gopurs_runtime.Any(&__local_ref_3)
})), gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(discard1_2_1, gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(rb_7.PtrVal().(*gopurs_runtime.Value))
})), gopurs_runtime.Func(func(mb_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (mb_9.Type == 9 && mb_9.IntVal == 3589588149) {
__t4 = gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(ra_6.PtrVal().(*gopurs_runtime.Value)) = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{a_8})}
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{a_8})}
}))
goto end_branch_4
} else {

}
}
{
if (mb_9.Type == 9 && mb_9.IntVal == 930809136) {
__t4 = gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(a_8, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(mb_9.UnsafePtr).V0))
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
})), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(ra_6.PtrVal().(*gopurs_runtime.Value))
})), gopurs_runtime.Func(func(ma_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (ma_10.Type == 9 && ma_10.IntVal == 3589588149) {
__t5 = gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(rb_7.PtrVal().(*gopurs_runtime.Value)) = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{b_9})}
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{b_9})}
}))
goto end_branch_5
} else {

}
}
{
if (ma_10.Type == 9 && ma_10.IntVal == 930809136) {
__t5 = gopurs_runtime.Apply(k_5, gopurs_runtime.Apply((*pkg_Data_Maybe.Data_Data_Maybe_Just)(ma_10.UnsafePtr).V0, b_9))
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
}))
}))
}))
}))
}))
}

func Call_applicativeParCont(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
applyParCont1_1_0 := gopurs_runtime.Apply(Get_applyParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)})
_ = applyParCont1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyParCont1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_monadParParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)}), "parallel"), gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
}))
}))
}

func Call_altParCont(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
discard1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Bind1_2_1)
_ = discard1_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_3
functorParCont1_5_4 := gopurs_runtime.Apply(Get_functorParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)})
_ = functorParCont1_5_4
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorParCont1_5_4
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_5 := gopurs_runtime.Bool(false)
_ = __local_ref_5
return gopurs_runtime.Any(&__local_ref_5)
})), gopurs_runtime.Func(func(done_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(discard1_3_2, gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(done_9.PtrVal().(*gopurs_runtime.Value))
})), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (b_11.IntVal) != (0) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply2(discard1_3_2, gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(done_9.PtrVal().(*gopurs_runtime.Value)) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
})), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, a_10)
}))
}
end_branch_6:
return __t6
}))
})), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_7, gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(done_9.PtrVal().(*gopurs_runtime.Value))
})), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (b_12.IntVal) != (0) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply2(discard1_3_2, gopurs_runtime.Apply(dictMonadEffect_0.liftEffect, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(done_9.PtrVal().(*gopurs_runtime.Value)) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
})), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, a_11)
}))
}
end_branch_7:
return __t7
}))
}))
}))
}))
}))
}

func Call_plusParCont(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
altParCont1_1_0 := gopurs_runtime.Apply(Get_altParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)})
_ = altParCont1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altParCont1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}))
}

func Call_alternativeParCont(dictMonadEffect_0_loop *Record_liftEffect_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Record_liftEffect_gopurs_runtime_Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
applicativeParCont1_1_0 := gopurs_runtime.Apply(Get_applicativeParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)})
_ = applicativeParCont1_1_0
plusParCont1_2_1 := gopurs_runtime.Apply(Get_plusParCont(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadEffect_0)})
_ = plusParCont1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeParCont1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusParCont1_2_1
}))
}


