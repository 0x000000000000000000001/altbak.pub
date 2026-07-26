package Data_Const

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	unsafe "unsafe"
)

var cache_Const gopurs_runtime.Value
var once_Const sync.Once
func Get_Const() gopurs_runtime.Value {
	once_Const.Do(func() {
		cache_Const = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Const(x_0_box)
})
	})
	return cache_Const
}

var cache_showConst gopurs_runtime.Value
var once_showConst sync.Once
func Get_showConst() gopurs_runtime.Value {
	once_showConst.Do(func() {
		cache_showConst = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showConst((*Record_show_gopurs_runtime_Value)(dictShow_0_box.UnsafePtr))
})
	})
	return cache_showConst
}

var cache_semiringConst gopurs_runtime.Value
var once_semiringConst sync.Once
func Get_semiringConst() gopurs_runtime.Value {
	once_semiringConst.Do(func() {
		cache_semiringConst = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_semiringConst((*Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value)(dictSemiring_0_box.UnsafePtr)))}
})
	})
	return cache_semiringConst
}

var cache_semigroupoidConst gopurs_runtime.Value
var once_semigroupoidConst sync.Once
func Get_semigroupoidConst() gopurs_runtime.Value {
	once_semigroupoidConst.Do(func() {
		cache_semigroupoidConst = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return cache_semigroupoidConst
}

var cache_semigroupConst gopurs_runtime.Value
var once_semigroupConst sync.Once
func Get_semigroupConst() gopurs_runtime.Value {
	once_semigroupConst.Do(func() {
		cache_semigroupConst = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_semigroupConst((*Record_append__gopurs_runtime_Value)(dictSemigroup_0_box.UnsafePtr)))}
})
	})
	return cache_semigroupConst
}

var cache_ringConst gopurs_runtime.Value
var once_ringConst sync.Once
func Get_ringConst() gopurs_runtime.Value {
	once_ringConst.Do(func() {
		cache_ringConst = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_ringConst((*Record_sub_gopurs_runtime_Value)(dictRing_0_box.UnsafePtr)))}
})
	})
	return cache_ringConst
}

var cache_ordConst gopurs_runtime.Value
var once_ordConst sync.Once
func Get_ordConst() gopurs_runtime.Value {
	once_ordConst.Do(func() {
		cache_ordConst = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_ordConst((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr)))}
})
	})
	return cache_ordConst
}

var cache_newtypeConst gopurs_runtime.Value
var once_newtypeConst sync.Once
func Get_newtypeConst() gopurs_runtime.Value {
	once_newtypeConst.Do(func() {
		cache_newtypeConst = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeConst
}

var cache_monoidConst gopurs_runtime.Value
var once_monoidConst sync.Once
func Get_monoidConst() gopurs_runtime.Value {
	once_monoidConst.Do(func() {
		cache_monoidConst = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monoidConst((*Record_mempty_gopurs_runtime_Value)(dictMonoid_0_box.UnsafePtr)))}
})
	})
	return cache_monoidConst
}

var cache_heytingAlgebraConst gopurs_runtime.Value
var once_heytingAlgebraConst sync.Once
func Get_heytingAlgebraConst() gopurs_runtime.Value {
	once_heytingAlgebraConst.Do(func() {
		cache_heytingAlgebraConst = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_heytingAlgebraConst((*Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value)(dictHeytingAlgebra_0_box.UnsafePtr)))}
})
	})
	return cache_heytingAlgebraConst
}

var cache_functorConst gopurs_runtime.Value
var once_functorConst sync.Once
func Get_functorConst() gopurs_runtime.Value {
	once_functorConst.Do(func() {
		cache_functorConst = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return m_1
}))
	})
	return cache_functorConst
}

var cache_invariantConst gopurs_runtime.Value
var once_invariantConst sync.Once
func Get_invariantConst() gopurs_runtime.Value {
	once_invariantConst.Do(func() {
		cache_invariantConst = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorConst(), "map"), f_0)
}))
	})
	return cache_invariantConst
}

var cache_euclideanRingConst gopurs_runtime.Value
var once_euclideanRingConst sync.Once
func Get_euclideanRingConst() gopurs_runtime.Value {
	once_euclideanRingConst.Do(func() {
		cache_euclideanRingConst = gopurs_runtime.Func(func(dictEuclideanRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_euclideanRingConst((*Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value)(dictEuclideanRing_0_box.UnsafePtr)))}
})
	})
	return cache_euclideanRingConst
}

var cache_eqConst gopurs_runtime.Value
var once_eqConst sync.Once
func Get_eqConst() gopurs_runtime.Value {
	once_eqConst.Do(func() {
		cache_eqConst = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_eqConst((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr)))}
})
	})
	return cache_eqConst
}

var cache_eq1Const gopurs_runtime.Value
var once_eq1Const sync.Once
func Get_eq1Const() gopurs_runtime.Value {
	once_eq1Const.Do(func() {
		cache_eq1Const = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Const((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_eq1Const
}

var cache_ord1Const gopurs_runtime.Value
var once_ord1Const sync.Once
func Get_ord1Const() gopurs_runtime.Value {
	once_ord1Const.Do(func() {
		cache_ord1Const = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Const((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_ord1Const
}

var cache_commutativeRingConst gopurs_runtime.Value
var once_commutativeRingConst sync.Once
func Get_commutativeRingConst() gopurs_runtime.Value {
	once_commutativeRingConst.Do(func() {
		cache_commutativeRingConst = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_commutativeRingConst((*Record_)(dictCommutativeRing_0_box.UnsafePtr)))}
})
	})
	return cache_commutativeRingConst
}

var cache_boundedConst gopurs_runtime.Value
var once_boundedConst sync.Once
func Get_boundedConst() gopurs_runtime.Value {
	once_boundedConst.Do(func() {
		cache_boundedConst = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_boundedConst((*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dictBounded_0_box.UnsafePtr)))}
})
	})
	return cache_boundedConst
}

var cache_booleanAlgebraConst gopurs_runtime.Value
var once_booleanAlgebraConst sync.Once
func Get_booleanAlgebraConst() gopurs_runtime.Value {
	once_booleanAlgebraConst.Do(func() {
		cache_booleanAlgebraConst = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_booleanAlgebraConst((*Record_)(dictBooleanAlgebra_0_box.UnsafePtr)))}
})
	})
	return cache_booleanAlgebraConst
}

var cache_applyConst gopurs_runtime.Value
var once_applyConst sync.Once
func Get_applyConst() gopurs_runtime.Value {
	once_applyConst.Do(func() {
		cache_applyConst = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyConst((*Record_append__gopurs_runtime_Value)(dictSemigroup_0_box.UnsafePtr))
})
	})
	return cache_applyConst
}

var cache_applicativeConst gopurs_runtime.Value
var once_applicativeConst sync.Once
func Get_applicativeConst() gopurs_runtime.Value {
	once_applicativeConst.Do(func() {
		cache_applicativeConst = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeConst((*Record_mempty_gopurs_runtime_Value)(dictMonoid_0_box.UnsafePtr))
})
	})
	return cache_applicativeConst
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

func Call_Const(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showConst(dictShow_0_loop *Record_show_gopurs_runtime_Value) gopurs_runtime.Value {
var dictShow_0 *Record_show_gopurs_runtime_Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Const "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(dictShow_0.show, v_1), gopurs_runtime.Str(")")))
}))
}

func Call_semiringConst(dictSemiring_0_loop *Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value) *Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value {
var dictSemiring_0 *Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value = dictSemiring_0_loop
_ = dictSemiring_0
return dictSemiring_0
}

func Call_semigroupConst(dictSemigroup_0_loop *Record_append__gopurs_runtime_Value) *Record_append__gopurs_runtime_Value {
var dictSemigroup_0 *Record_append__gopurs_runtime_Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return dictSemigroup_0
}

func Call_ringConst(dictRing_0_loop *Record_sub_gopurs_runtime_Value) *Record_sub_gopurs_runtime_Value {
var dictRing_0 *Record_sub_gopurs_runtime_Value = dictRing_0_loop
_ = dictRing_0
return dictRing_0
}

func Call_ordConst(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) *Record_compare_gopurs_runtime_Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_monoidConst(dictMonoid_0_loop *Record_mempty_gopurs_runtime_Value) *Record_mempty_gopurs_runtime_Value {
var dictMonoid_0 *Record_mempty_gopurs_runtime_Value = dictMonoid_0_loop
_ = dictMonoid_0
return dictMonoid_0
}

func Call_heytingAlgebraConst(dictHeytingAlgebra_0_loop *Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value) *Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value {
var dictHeytingAlgebra_0 *Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return dictHeytingAlgebra_0
}

func Call_euclideanRingConst(dictEuclideanRing_0_loop *Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value) *Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value {
var dictEuclideanRing_0 *Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
return dictEuclideanRing_0
}

func Call_eqConst(dictEq_0_loop *Record_eq_gopurs_runtime_Value) *Record_eq_gopurs_runtime_Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_eq1Const(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
eq_1_0 := dictEq_0.eq
_ = eq_1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eq_1_0
}))
}

func Call_ord1Const(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.compare
_ = compare_1_0
eq_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)}, "Eq0_NOT_FOUND"), gopurs_runtime.Value{}), "eq")
_ = eq_2_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq_2_1
}))
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return compare_1_0
}))
}

func Call_commutativeRingConst(dictCommutativeRing_0_loop *Record_) *Record_ {
var dictCommutativeRing_0 *Record_ = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
return dictCommutativeRing_0
}

func Call_boundedConst(dictBounded_0_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value {
var dictBounded_0 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}

func Call_booleanAlgebraConst(dictBooleanAlgebra_0_loop *Record_) *Record_ {
var dictBooleanAlgebra_0 *Record_ = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
return dictBooleanAlgebra_0
}

func Call_applyConst(dictSemigroup_0_loop *Record_append__gopurs_runtime_Value) gopurs_runtime.Value {
var dictSemigroup_0 *Record_append__gopurs_runtime_Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConst()
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_0.append_, v_1, v1_2)
}))
}

func Call_applicativeConst(dictMonoid_0_loop *Record_mempty_gopurs_runtime_Value) gopurs_runtime.Value {
var dictMonoid_0 *Record_mempty_gopurs_runtime_Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := dictMonoid_0.mempty
_ = mempty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonoid_0)}, "Semigroup0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
applyConst1_3_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConst()
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), v_3, v1_4)
}))
_ = applyConst1_3_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyConst1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
}))
}


