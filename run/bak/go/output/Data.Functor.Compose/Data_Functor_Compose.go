package Data_Functor_Compose

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	unsafe "unsafe"
)

var cache_Compose gopurs_runtime.Value
var once_Compose sync.Once
func Get_Compose() gopurs_runtime.Value {
	once_Compose.Do(func() {
		cache_Compose = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Compose(x_0_box)
})
	})
	return cache_Compose
}

var cache_showCompose gopurs_runtime.Value
var once_showCompose sync.Once
func Get_showCompose() gopurs_runtime.Value {
	once_showCompose.Do(func() {
		cache_showCompose = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showCompose((*Record_show_gopurs_runtime_Value)(dictShow_0_box.UnsafePtr))
})
	})
	return cache_showCompose
}

var cache_newtypeCompose gopurs_runtime.Value
var once_newtypeCompose sync.Once
func Get_newtypeCompose() gopurs_runtime.Value {
	once_newtypeCompose.Do(func() {
		cache_newtypeCompose = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCompose
}

var cache_functorCompose gopurs_runtime.Value
var once_functorCompose sync.Once
func Get_functorCompose() gopurs_runtime.Value {
	once_functorCompose.Do(func() {
		cache_functorCompose = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorCompose((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr), (*Record_map__gopurs_runtime_Value)(dictFunctor1_1_box.UnsafePtr))
})
	})
	return cache_functorCompose
}

var cache_eqCompose gopurs_runtime.Value
var once_eqCompose sync.Once
func Get_eqCompose() gopurs_runtime.Value {
	once_eqCompose.Do(func() {
		cache_eqCompose = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqCompose((*Record_eq1_gopurs_runtime_Value)(dictEq1_0_box.UnsafePtr), (*Record_eq1_gopurs_runtime_Value)(dictEq11_1_box.UnsafePtr), (*Record_eq_gopurs_runtime_Value)(dictEq_2_box.UnsafePtr))
})
	})
	return cache_eqCompose
}

var cache_ordCompose gopurs_runtime.Value
var once_ordCompose sync.Once
func Get_ordCompose() gopurs_runtime.Value {
	once_ordCompose.Do(func() {
		cache_ordCompose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordCompose((*Record_compare1_gopurs_runtime_Value)(dictOrd1_0_box.UnsafePtr))
})
	})
	return cache_ordCompose
}

var cache_eq1Compose gopurs_runtime.Value
var once_eq1Compose sync.Once
func Get_eq1Compose() gopurs_runtime.Value {
	once_eq1Compose.Do(func() {
		cache_eq1Compose = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Compose((*Record_eq1_gopurs_runtime_Value)(dictEq1_0_box.UnsafePtr), (*Record_eq1_gopurs_runtime_Value)(dictEq11_1_box.UnsafePtr))
})
	})
	return cache_eq1Compose
}

var cache_ord1Compose gopurs_runtime.Value
var once_ord1Compose sync.Once
func Get_ord1Compose() gopurs_runtime.Value {
	once_ord1Compose.Do(func() {
		cache_ord1Compose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Compose((*Record_compare1_gopurs_runtime_Value)(dictOrd1_0_box.UnsafePtr))
})
	})
	return cache_ord1Compose
}

var cache_bihoistCompose gopurs_runtime.Value
var once_bihoistCompose sync.Once
func Get_bihoistCompose() gopurs_runtime.Value {
	once_bihoistCompose.Do(func() {
		cache_bihoistCompose = gopurs_runtime.Func4(func(dictFunctor_0_box gopurs_runtime.Value, natF_1_box gopurs_runtime.Value, natG_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bihoistCompose((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr), natF_1_box, natG_2_box, v_3_box)
})
	})
	return cache_bihoistCompose
}

var cache_applyCompose gopurs_runtime.Value
var once_applyCompose sync.Once
func Get_applyCompose() gopurs_runtime.Value {
	once_applyCompose.Do(func() {
		cache_applyCompose = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyCompose((*Record_apply_gopurs_runtime_Value)(dictApply_0_box.UnsafePtr))
})
	})
	return cache_applyCompose
}

var cache_applicativeCompose gopurs_runtime.Value
var once_applicativeCompose sync.Once
func Get_applicativeCompose() gopurs_runtime.Value {
	once_applicativeCompose.Do(func() {
		cache_applicativeCompose = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeCompose((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_applicativeCompose
}

var cache_altCompose gopurs_runtime.Value
var once_altCompose sync.Once
func Get_altCompose() gopurs_runtime.Value {
	once_altCompose.Do(func() {
		cache_altCompose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altCompose((*Record_alt_gopurs_runtime_Value)(dictAlt_0_box.UnsafePtr))
})
	})
	return cache_altCompose
}

var cache_plusCompose gopurs_runtime.Value
var once_plusCompose sync.Once
func Get_plusCompose() gopurs_runtime.Value {
	once_plusCompose.Do(func() {
		cache_plusCompose = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusCompose((*Record_empty_gopurs_runtime_Value)(dictPlus_0_box.UnsafePtr))
})
	})
	return cache_plusCompose
}

var cache_alternativeCompose gopurs_runtime.Value
var once_alternativeCompose sync.Once
func Get_alternativeCompose() gopurs_runtime.Value {
	once_alternativeCompose.Do(func() {
		cache_alternativeCompose = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeCompose((*Record_)(dictAlternative_0_box.UnsafePtr))
})
	})
	return cache_alternativeCompose
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

func Call_Compose(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showCompose(dictShow_0_loop *Record_show_gopurs_runtime_Value) gopurs_runtime.Value {
var dictShow_0 *Record_show_gopurs_runtime_Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Compose "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(dictShow_0.show, v_1), gopurs_runtime.Str(")")))
}))
}

func Call_functorCompose(dictFunctor_0_loop *Record_map__gopurs_runtime_Value, dictFunctor1_1_loop *Record_map__gopurs_runtime_Value) gopurs_runtime.Value {
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 *Record_map__gopurs_runtime_Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFunctor_0.map_, gopurs_runtime.Apply(dictFunctor1_1.map_, f_2), v_3)
}))
}

func Call_eqCompose(dictEq1_0_loop *Record_eq1_gopurs_runtime_Value, dictEq11_1_loop *Record_eq1_gopurs_runtime_Value, dictEq_2_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq1_0 *Record_eq1_gopurs_runtime_Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 *Record_eq1_gopurs_runtime_Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 *Record_eq_gopurs_runtime_Value = dictEq_2_loop
_ = dictEq_2
eq11_3_1 := gopurs_runtime.Apply(dictEq11_1.eq1, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_2)})
_ = eq11_3_1
eq11_3_0 := gopurs_runtime.Apply(dictEq1_0.eq1, gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_1, x_4, y_5)
})))
_ = eq11_3_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_0, v_4, v1_5)
}))
}

func Call_ordCompose(dictOrd1_0_loop *Record_compare1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd1_0 *Record_compare1_gopurs_runtime_Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), dictOrd_5)
_ = compare11_6_4
eq11_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_7_5
eqApp2_8_6 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_7_5, x_8, y_9)
}))
_ = eqApp2_8_6
compare11_6_3 := gopurs_runtime.Apply(dictOrd1_0.compare1, gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_8_6
}), gopurs_runtime.Func2(func(x_9 gopurs_runtime.Value, y_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_6_4, x_9, y_10)
})))
_ = compare11_6_3
eq11_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_7_8
eq11_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_7_8, x_8, y_9)
})))
_ = eq11_7_7
eqCompose3_8_9 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_7_7, v_8, v1_9)
}))
_ = eqCompose3_8_9
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCompose3_8_9
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_6_3, v_9, v1_10)
}))
})
})
}

func Call_eq1Compose(dictEq1_0_loop *Record_eq1_gopurs_runtime_Value, dictEq11_1_loop *Record_eq1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq1_0 *Record_eq1_gopurs_runtime_Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 *Record_eq1_gopurs_runtime_Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_3_0 := gopurs_runtime.Apply(dictEq11_1.eq1, dictEq_2)
_ = eq11_3_0
return gopurs_runtime.Apply(dictEq1_0.eq1, gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_0, x_4, y_5)
})))
}))
}

func Call_ord1Compose(dictOrd1_0_loop *Record_compare1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd1_0 *Record_compare1_gopurs_runtime_Value = dictOrd1_0_loop
_ = dictOrd1_0
ordCompose1_1_0 := gopurs_runtime.Apply(Get_ordCompose(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)})
_ = ordCompose1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordCompose2_4_2 := gopurs_runtime.Apply(ordCompose1_1_0, dictOrd11_3)
_ = ordCompose2_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_5_3
eq1Compose2_6_4 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_6 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "eq1"), dictEq_6)
_ = eq11_7_5
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_7_5, x_8, y_9)
})))
}))
_ = eq1Compose2_6_4
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Compose2_6_4
}), gopurs_runtime.Func(func(dictOrd_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordCompose2_4_2, dictOrd_7), "compare")
}))
})
}

func Call_bihoistCompose(dictFunctor_0_loop *Record_map__gopurs_runtime_Value, natF_1_loop gopurs_runtime.Value, natG_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
var natF_1 gopurs_runtime.Value = natF_1_loop
_ = natF_1
var natG_2 gopurs_runtime.Value = natG_2_loop
_ = natG_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
return gopurs_runtime.Apply(natF_1, gopurs_runtime.Apply2(dictFunctor_0.map_, natG_2, v_3))
}

func Call_applyCompose(dictApply_0_loop *Record_apply_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApply_0 *Record_apply_gopurs_runtime_Value = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApply_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
apply1_3_1 := gopurs_runtime.RecordGet(dictApply1_2, "apply")
_ = apply1_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
functorCompose2_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "map"), f_5), v_6)
}))
_ = functorCompose2_5_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.apply, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), apply1_3_1, v_6), v1_7)
}))
})
}

func Call_applicativeCompose(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictApplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_2
apply1_5_3 := gopurs_runtime.RecordGet(__local_var_4_2, "apply")
_ = apply1_5_3
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
functorCompose2_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "map"), f_7), v_8)
}))
_ = functorCompose2_7_6
applyCompose2_6_4 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), apply1_5_3, v_8), v1_9)
}))
_ = applyCompose2_6_4
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose2_6_4
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "pure"), x_7))
}))
})
}

func Call_altCompose(dictAlt_0_loop *Record_alt_gopurs_runtime_Value) gopurs_runtime.Value {
var dictAlt_0 *Record_alt_gopurs_runtime_Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlt_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3), v_4)
}))
_ = functorCompose2_3_1
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_1
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictAlt_0.alt, v_4, v1_5)
}))
})
}

func Call_plusCompose(dictPlus_0_loop *Record_empty_gopurs_runtime_Value) gopurs_runtime.Value {
var dictPlus_0 *Record_empty_gopurs_runtime_Value = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := dictPlus_0.empty
_ = empty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictPlus_0)}, "Alt0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_4, "map"), f_5), v_6)
}))
_ = functorCompose2_5_3
altCompose2_6_4 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "alt"), v_6, v1_7)
}))
_ = altCompose2_6_4
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_6_4
}), empty_1_0)
})
}

func Call_alternativeCompose(dictAlternative_0_loop *Record_) gopurs_runtime.Value {
var dictAlternative_0 *Record_ = dictAlternative_0_loop
_ = dictAlternative_0
applicativeCompose1_1_0 := gopurs_runtime.Apply(Get_applicativeCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applicativeCompose1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, "Plus1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
empty_3_2 := gopurs_runtime.RecordGet(__local_var_2_1, "empty")
_ = empty_3_2
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_4
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
plusCompose1_4_3 := gopurs_runtime.Func(func(dictFunctor_6 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_6, "map"), f_7), v_8)
}))
_ = functorCompose2_7_6
altCompose2_8_7 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "alt"), v_8, v1_9)
}))
_ = altCompose2_8_7
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_8_7
}), empty_3_2)
})
_ = plusCompose1_4_3
return gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeCompose2_6_8 := gopurs_runtime.Apply(applicativeCompose1_1_0, dictApplicative_5)
_ = applicativeCompose2_6_8
plusCompose2_7_9 := gopurs_runtime.Apply(plusCompose1_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = plusCompose2_7_9
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeCompose2_6_8
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return plusCompose2_7_9
}))
})
}


