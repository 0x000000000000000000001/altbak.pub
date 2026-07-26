package Data_Bounded_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	unsafe "unsafe"
)

var cache_genericTopNoArguments gopurs_runtime.Value
var once_genericTopNoArguments sync.Once
func Get_genericTopNoArguments() gopurs_runtime.Value {
	once_genericTopNoArguments.Do(func() {
		cache_genericTopNoArguments = gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil})
	})
	return cache_genericTopNoArguments
}

var cache_genericTopArgument gopurs_runtime.Value
var once_genericTopArgument sync.Once
func Get_genericTopArgument() gopurs_runtime.Value {
	once_genericTopArgument.Do(func() {
		cache_genericTopArgument = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTopArgument((*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dictBounded_0_box.UnsafePtr))
})
	})
	return cache_genericTopArgument
}

var cache_genericTop_prime gopurs_runtime.Value
var once_genericTop_prime sync.Once
func Get_genericTop_prime() gopurs_runtime.Value {
	once_genericTop_prime.Do(func() {
		cache_genericTop_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTop_prime((*Record_genericTop_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericTop_prime
}

var cache_genericTopConstructor gopurs_runtime.Value
var once_genericTopConstructor sync.Once
func Get_genericTopConstructor() gopurs_runtime.Value {
	once_genericTopConstructor.Do(func() {
		cache_genericTopConstructor = gopurs_runtime.Func(func(dictGenericTop_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTopConstructor((*Record_genericTop_prime_gopurs_runtime_Value)(dictGenericTop_0_box.UnsafePtr))
})
	})
	return cache_genericTopConstructor
}

var cache_genericTopProduct gopurs_runtime.Value
var once_genericTopProduct sync.Once
func Get_genericTopProduct() gopurs_runtime.Value {
	once_genericTopProduct.Do(func() {
		cache_genericTopProduct = gopurs_runtime.Func(func(dictGenericTop_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTopProduct((*Record_genericTop_prime_gopurs_runtime_Value)(dictGenericTop_0_box.UnsafePtr))
})
	})
	return cache_genericTopProduct
}

var cache_genericTopSum gopurs_runtime.Value
var once_genericTopSum sync.Once
func Get_genericTopSum() gopurs_runtime.Value {
	once_genericTopSum.Do(func() {
		cache_genericTopSum = gopurs_runtime.Func(func(dictGenericTop_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTopSum((*Record_genericTop_prime_gopurs_runtime_Value)(dictGenericTop_0_box.UnsafePtr))
})
	})
	return cache_genericTopSum
}

var cache_genericTop gopurs_runtime.Value
var once_genericTop sync.Once
func Get_genericTop() gopurs_runtime.Value {
	once_genericTop.Do(func() {
		cache_genericTop = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTop((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericTop_prime_gopurs_runtime_Value)(dictGenericTop_1_box.UnsafePtr))
})
	})
	return cache_genericTop
}

var cache_genericBottomNoArguments gopurs_runtime.Value
var once_genericBottomNoArguments sync.Once
func Get_genericBottomNoArguments() gopurs_runtime.Value {
	once_genericBottomNoArguments.Do(func() {
		cache_genericBottomNoArguments = gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil})
	})
	return cache_genericBottomNoArguments
}

var cache_genericBottomArgument gopurs_runtime.Value
var once_genericBottomArgument sync.Once
func Get_genericBottomArgument() gopurs_runtime.Value {
	once_genericBottomArgument.Do(func() {
		cache_genericBottomArgument = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottomArgument((*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dictBounded_0_box.UnsafePtr))
})
	})
	return cache_genericBottomArgument
}

var cache_genericBottom_prime gopurs_runtime.Value
var once_genericBottom_prime sync.Once
func Get_genericBottom_prime() gopurs_runtime.Value {
	once_genericBottom_prime.Do(func() {
		cache_genericBottom_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottom_prime((*Record_genericBottom_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericBottom_prime
}

var cache_genericBottomConstructor gopurs_runtime.Value
var once_genericBottomConstructor sync.Once
func Get_genericBottomConstructor() gopurs_runtime.Value {
	once_genericBottomConstructor.Do(func() {
		cache_genericBottomConstructor = gopurs_runtime.Func(func(dictGenericBottom_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottomConstructor((*Record_genericBottom_prime_gopurs_runtime_Value)(dictGenericBottom_0_box.UnsafePtr))
})
	})
	return cache_genericBottomConstructor
}

var cache_genericBottomProduct gopurs_runtime.Value
var once_genericBottomProduct sync.Once
func Get_genericBottomProduct() gopurs_runtime.Value {
	once_genericBottomProduct.Do(func() {
		cache_genericBottomProduct = gopurs_runtime.Func(func(dictGenericBottom_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottomProduct((*Record_genericBottom_prime_gopurs_runtime_Value)(dictGenericBottom_0_box.UnsafePtr))
})
	})
	return cache_genericBottomProduct
}

var cache_genericBottomSum gopurs_runtime.Value
var once_genericBottomSum sync.Once
func Get_genericBottomSum() gopurs_runtime.Value {
	once_genericBottomSum.Do(func() {
		cache_genericBottomSum = gopurs_runtime.Func(func(dictGenericBottom_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottomSum((*Record_genericBottom_prime_gopurs_runtime_Value)(dictGenericBottom_0_box.UnsafePtr))
})
	})
	return cache_genericBottomSum
}

var cache_genericBottom gopurs_runtime.Value
var once_genericBottom sync.Once
func Get_genericBottom() gopurs_runtime.Value {
	once_genericBottom.Do(func() {
		cache_genericBottom = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBottom_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottom((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericBottom_prime_gopurs_runtime_Value)(dictGenericBottom_1_box.UnsafePtr))
})
	})
	return cache_genericBottom
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

func Call_genericTopArgument(dictBounded_0_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBounded_0 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.RecordDict1("genericTop'", dictBounded_0.top)
}

func Call_genericTop_prime(dict_0_loop *Record_genericTop_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericTop_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericTop_prime
}

func Call_genericTopConstructor(dictGenericTop_0_loop *Record_genericTop_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericTop_0 *Record_genericTop_prime_gopurs_runtime_Value = dictGenericTop_0_loop
_ = dictGenericTop_0
return gopurs_runtime.RecordDict1("genericTop'", dictGenericTop_0.genericTop_prime)
}

func Call_genericTopProduct(dictGenericTop_0_loop *Record_genericTop_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericTop_0 *Record_genericTop_prime_gopurs_runtime_Value = dictGenericTop_0_loop
_ = dictGenericTop_0
genericTop_prime1_1_0 := dictGenericTop_0.genericTop_prime
_ = genericTop_prime1_1_0
return gopurs_runtime.Func(func(dictGenericTop1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{genericTop_prime1_1_0, gopurs_runtime.RecordGet(dictGenericTop1_2, "genericTop'")})})
})
}

func Call_genericTopSum(dictGenericTop_0_loop *Record_genericTop_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericTop_0 *Record_genericTop_prime_gopurs_runtime_Value = dictGenericTop_0_loop
_ = dictGenericTop_0
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr{dictGenericTop_0.genericTop_prime})})
}

func Call_genericTop(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericTop_1_loop *Record_genericTop_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericTop_1 *Record_genericTop_prime_gopurs_runtime_Value = dictGenericTop_1_loop
_ = dictGenericTop_1
return gopurs_runtime.Apply(dictGeneric_0.to, dictGenericTop_1.genericTop_prime)
}

func Call_genericBottomArgument(dictBounded_0_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBounded_0 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.RecordDict1("genericBottom'", dictBounded_0.bottom)
}

func Call_genericBottom_prime(dict_0_loop *Record_genericBottom_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericBottom_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericBottom_prime
}

func Call_genericBottomConstructor(dictGenericBottom_0_loop *Record_genericBottom_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericBottom_0 *Record_genericBottom_prime_gopurs_runtime_Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
return gopurs_runtime.RecordDict1("genericBottom'", dictGenericBottom_0.genericBottom_prime)
}

func Call_genericBottomProduct(dictGenericBottom_0_loop *Record_genericBottom_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericBottom_0 *Record_genericBottom_prime_gopurs_runtime_Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
genericBottom_prime1_1_0 := dictGenericBottom_0.genericBottom_prime
_ = genericBottom_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBottom1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{genericBottom_prime1_1_0, gopurs_runtime.RecordGet(dictGenericBottom1_2, "genericBottom'")})})
})
}

func Call_genericBottomSum(dictGenericBottom_0_loop *Record_genericBottom_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericBottom_0 *Record_genericBottom_prime_gopurs_runtime_Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl{dictGenericBottom_0.genericBottom_prime})})
}

func Call_genericBottom(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericBottom_1_loop *Record_genericBottom_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBottom_1 *Record_genericBottom_prime_gopurs_runtime_Value = dictGenericBottom_1_loop
_ = dictGenericBottom_1
return gopurs_runtime.Apply(dictGeneric_0.to, dictGenericBottom_1.genericBottom_prime)
}


