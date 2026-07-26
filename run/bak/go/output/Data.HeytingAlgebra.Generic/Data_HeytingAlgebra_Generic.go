package Data_HeytingAlgebra_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	unsafe "unsafe"
)

var cache_genericTT_prime gopurs_runtime.Value
var once_genericTT_prime sync.Once
func Get_genericTT_prime() gopurs_runtime.Value {
	once_genericTT_prime.Do(func() {
		cache_genericTT_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTT_prime((*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericTT_prime
}

var cache_genericTT gopurs_runtime.Value
var once_genericTT sync.Once
func Get_genericTT() gopurs_runtime.Value {
	once_genericTT.Do(func() {
		cache_genericTT = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTT((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dictGenericHeytingAlgebra_1_box.UnsafePtr))
})
	})
	return cache_genericTT
}

var cache_genericNot_prime gopurs_runtime.Value
var once_genericNot_prime sync.Once
func Get_genericNot_prime() gopurs_runtime.Value {
	once_genericNot_prime.Do(func() {
		cache_genericNot_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericNot_prime((*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericNot_prime
}

var cache_genericNot gopurs_runtime.Value
var once_genericNot sync.Once
func Get_genericNot() gopurs_runtime.Value {
	once_genericNot.Do(func() {
		cache_genericNot = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericNot((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dictGenericHeytingAlgebra_1_box.UnsafePtr), x_2_box)
})
	})
	return cache_genericNot
}

var cache_genericImplies_prime gopurs_runtime.Value
var once_genericImplies_prime sync.Once
func Get_genericImplies_prime() gopurs_runtime.Value {
	once_genericImplies_prime.Do(func() {
		cache_genericImplies_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericImplies_prime((*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericImplies_prime
}

var cache_genericImplies gopurs_runtime.Value
var once_genericImplies sync.Once
func Get_genericImplies() gopurs_runtime.Value {
	once_genericImplies.Do(func() {
		cache_genericImplies = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericImplies((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dictGenericHeytingAlgebra_1_box.UnsafePtr), x_2_box, y_3_box)
})
	})
	return cache_genericImplies
}

var cache_genericHeytingAlgebraNoArguments gopurs_runtime.Value
var once_genericHeytingAlgebraNoArguments sync.Once
func Get_genericHeytingAlgebraNoArguments() gopurs_runtime.Value {
	once_genericHeytingAlgebraNoArguments.Do(func() {
		cache_genericHeytingAlgebraNoArguments = gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
}), gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
}), gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}})
	})
	return cache_genericHeytingAlgebraNoArguments
}

var cache_genericHeytingAlgebraArgument gopurs_runtime.Value
var once_genericHeytingAlgebraArgument sync.Once
func Get_genericHeytingAlgebraArgument() gopurs_runtime.Value {
	once_genericHeytingAlgebraArgument.Do(func() {
		cache_genericHeytingAlgebraArgument = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericHeytingAlgebraArgument((*Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value)(dictHeytingAlgebra_0_box.UnsafePtr))
})
	})
	return cache_genericHeytingAlgebraArgument
}

var cache_genericFF_prime gopurs_runtime.Value
var once_genericFF_prime sync.Once
func Get_genericFF_prime() gopurs_runtime.Value {
	once_genericFF_prime.Do(func() {
		cache_genericFF_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFF_prime((*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericFF_prime
}

var cache_genericFF gopurs_runtime.Value
var once_genericFF sync.Once
func Get_genericFF() gopurs_runtime.Value {
	once_genericFF.Do(func() {
		cache_genericFF = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFF((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dictGenericHeytingAlgebra_1_box.UnsafePtr))
})
	})
	return cache_genericFF
}

var cache_genericDisj_prime gopurs_runtime.Value
var once_genericDisj_prime sync.Once
func Get_genericDisj_prime() gopurs_runtime.Value {
	once_genericDisj_prime.Do(func() {
		cache_genericDisj_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericDisj_prime((*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericDisj_prime
}

var cache_genericDisj gopurs_runtime.Value
var once_genericDisj sync.Once
func Get_genericDisj() gopurs_runtime.Value {
	once_genericDisj.Do(func() {
		cache_genericDisj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericDisj((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dictGenericHeytingAlgebra_1_box.UnsafePtr), x_2_box, y_3_box)
})
	})
	return cache_genericDisj
}

var cache_genericConj_prime gopurs_runtime.Value
var once_genericConj_prime sync.Once
func Get_genericConj_prime() gopurs_runtime.Value {
	once_genericConj_prime.Do(func() {
		cache_genericConj_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericConj_prime((*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericConj_prime
}

var cache_genericHeytingAlgebraConstructor gopurs_runtime.Value
var once_genericHeytingAlgebraConstructor sync.Once
func Get_genericHeytingAlgebraConstructor() gopurs_runtime.Value {
	once_genericHeytingAlgebraConstructor.Do(func() {
		cache_genericHeytingAlgebraConstructor = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericHeytingAlgebraConstructor((*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dictGenericHeytingAlgebra_0_box.UnsafePtr))
})
	})
	return cache_genericHeytingAlgebraConstructor
}

var cache_genericHeytingAlgebraProduct gopurs_runtime.Value
var once_genericHeytingAlgebraProduct sync.Once
func Get_genericHeytingAlgebraProduct() gopurs_runtime.Value {
	once_genericHeytingAlgebraProduct.Do(func() {
		cache_genericHeytingAlgebraProduct = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericHeytingAlgebraProduct((*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dictGenericHeytingAlgebra_0_box.UnsafePtr))
})
	})
	return cache_genericHeytingAlgebraProduct
}

var cache_genericConj gopurs_runtime.Value
var once_genericConj sync.Once
func Get_genericConj() gopurs_runtime.Value {
	once_genericConj.Do(func() {
		cache_genericConj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericConj((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value)(dictGenericHeytingAlgebra_1_box.UnsafePtr), x_2_box, y_3_box)
})
	})
	return cache_genericConj
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

func Call_genericTT_prime(dict_0_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericTT_prime
}

func Call_genericTT(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericHeytingAlgebra_1_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
return gopurs_runtime.Apply(dictGeneric_0.to, dictGenericHeytingAlgebra_1.genericTT_prime)
}

func Call_genericNot_prime(dict_0_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericNot_prime
}

func Call_genericNot(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericHeytingAlgebra_1_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictGeneric_0.to, gopurs_runtime.Apply(dictGenericHeytingAlgebra_1.genericNot_prime, gopurs_runtime.Apply(dictGeneric_0.from, x_2)))
}

func Call_genericImplies_prime(dict_0_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericImplies_prime
}

func Call_genericImplies(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericHeytingAlgebra_1_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.to, gopurs_runtime.Apply2(dictGenericHeytingAlgebra_1.genericImplies_prime, gopurs_runtime.Apply(dictGeneric_0.from, x_2), gopurs_runtime.Apply(dictGeneric_0.from, y_3)))
}

func Call_genericHeytingAlgebraArgument(dictHeytingAlgebra_0_loop *Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 *Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_0.conj, v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_0.disj, v_1, v1_2)
}), dictHeytingAlgebra_0.ff, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_0.implies, v_1, v1_2)
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictHeytingAlgebra_0.not, v_1)
}), dictHeytingAlgebra_0.tt})
}

func Call_genericFF_prime(dict_0_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericFF_prime
}

func Call_genericFF(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericHeytingAlgebra_1_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
return gopurs_runtime.Apply(dictGeneric_0.to, dictGenericHeytingAlgebra_1.genericFF_prime)
}

func Call_genericDisj_prime(dict_0_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericDisj_prime
}

func Call_genericDisj(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericHeytingAlgebra_1_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.to, gopurs_runtime.Apply2(dictGenericHeytingAlgebra_1.genericDisj_prime, gopurs_runtime.Apply(dictGeneric_0.from, x_2), gopurs_runtime.Apply(dictGeneric_0.from, y_3)))
}

func Call_genericConj_prime(dict_0_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericConj_prime
}

func Call_genericHeytingAlgebraConstructor(dictGenericHeytingAlgebra_0_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericHeytingAlgebra_0 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dictGenericHeytingAlgebra_0_loop
_ = dictGenericHeytingAlgebra_0
return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictGenericHeytingAlgebra_0.genericConj_prime, v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictGenericHeytingAlgebra_0.genericDisj_prime, v_1, v1_2)
}), dictGenericHeytingAlgebra_0.genericFF_prime, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictGenericHeytingAlgebra_0.genericImplies_prime, v_1, v1_2)
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.genericNot_prime, v_1)
}), dictGenericHeytingAlgebra_0.genericTT_prime})
}

func Call_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericHeytingAlgebra_0 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dictGenericHeytingAlgebra_0_loop
_ = dictGenericHeytingAlgebra_0
genericFF_prime1_1_0 := dictGenericHeytingAlgebra_0.genericFF_prime
_ = genericFF_prime1_1_0
genericTT_prime1_2_1 := dictGenericHeytingAlgebra_0.genericTT_prime
_ = genericTT_prime1_2_1
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{gopurs_runtime.Apply2(dictGenericHeytingAlgebra_0.genericConj_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericConj'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_5.UnsafePtr).V1)})}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{gopurs_runtime.Apply2(dictGenericHeytingAlgebra_0.genericDisj_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericDisj'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_5.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{genericFF_prime1_1_0, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericFF'")})}, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{gopurs_runtime.Apply2(dictGenericHeytingAlgebra_0.genericImplies_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericImplies'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_5.UnsafePtr).V1)})}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.genericNot_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_4.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericNot'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_4.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{genericTT_prime1_2_1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericTT'")})}})
})
}

func Call_genericConj(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericHeytingAlgebra_1_loop *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.to, gopurs_runtime.Apply2(dictGenericHeytingAlgebra_1.genericConj_prime, gopurs_runtime.Apply(dictGeneric_0.from, x_2), gopurs_runtime.Apply(dictGeneric_0.from, y_3)))
}


