package Data_Bounded

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	unsafe "unsafe"
)

var cache_topRecord gopurs_runtime.Value
var once_topRecord sync.Once
func Get_topRecord() gopurs_runtime.Value {
	once_topRecord.Do(func() {
		cache_topRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_topRecord((*Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_topRecord
}

var cache_top gopurs_runtime.Value
var once_top sync.Once
func Get_top() gopurs_runtime.Value {
	once_top.Do(func() {
		cache_top = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top((*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_top
}

var cache_boundedUnit gopurs_runtime.Value
var once_boundedUnit sync.Once
func Get_boundedUnit() gopurs_runtime.Value {
	once_boundedUnit.Do(func() {
		cache_boundedUnit = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordUnit()
}), pkg_Data_Unit.Get_unit(), pkg_Data_Unit.Get_unit())
	})
	return cache_boundedUnit
}

var cache_boundedRecordNil gopurs_runtime.Value
var once_boundedRecordNil sync.Once
func Get_boundedRecordNil() gopurs_runtime.Value {
	once_boundedRecordNil.Do(func() {
		cache_boundedRecordNil = gopurs_runtime.RecordDict3("OrdRecord0", "bottomRecord", "topRecord", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordRecordNil()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}))
	})
	return cache_boundedRecordNil
}

var cache_boundedProxy gopurs_runtime.Value
var once_boundedProxy sync.Once
func Get_boundedProxy() gopurs_runtime.Value {
	once_boundedProxy.Do(func() {
		cache_boundedProxy = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordProxy()
}), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})
	})
	return cache_boundedProxy
}

var cache_boundedOrdering gopurs_runtime.Value
var once_boundedOrdering sync.Once
func Get_boundedOrdering() gopurs_runtime.Value {
	once_boundedOrdering.Do(func() {
		cache_boundedOrdering = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordOrdering()
}), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
	})
	return cache_boundedOrdering
}

var cache_boundedNumber gopurs_runtime.Value
var once_boundedNumber sync.Once
func Get_boundedNumber() gopurs_runtime.Value {
	once_boundedNumber.Do(func() {
		cache_boundedNumber = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordNumber()
}), Get_bottomNumber(), Get_topNumber())
	})
	return cache_boundedNumber
}

var cache_boundedInt gopurs_runtime.Value
var once_boundedInt sync.Once
func Get_boundedInt() gopurs_runtime.Value {
	once_boundedInt.Do(func() {
		cache_boundedInt = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}))
}), Get_bottomInt(), Get_topInt())
	})
	return cache_boundedInt
}

var cache_boundedChar gopurs_runtime.Value
var once_boundedChar sync.Once
func Get_boundedChar() gopurs_runtime.Value {
	once_boundedChar.Do(func() {
		cache_boundedChar = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordChar()
}), Get_bottomChar(), Get_topChar())
	})
	return cache_boundedChar
}

var cache_boundedBoolean gopurs_runtime.Value
var once_boundedBoolean sync.Once
func Get_boundedBoolean() gopurs_runtime.Value {
	once_boundedBoolean.Do(func() {
		cache_boundedBoolean = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordBoolean()
}), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true))
	})
	return cache_boundedBoolean
}

var cache_bottomRecord gopurs_runtime.Value
var once_bottomRecord sync.Once
func Get_bottomRecord() gopurs_runtime.Value {
	once_bottomRecord.Do(func() {
		cache_bottomRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottomRecord((*Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_bottomRecord
}

var cache_boundedRecord gopurs_runtime.Value
var once_boundedRecord sync.Once
func Get_boundedRecord() gopurs_runtime.Value {
	once_boundedRecord.Do(func() {
		cache_boundedRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictBoundedRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedRecord((*Record_)(_dollar__unused_0_box.UnsafePtr), (*Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value)(dictBoundedRecord_1_box.UnsafePtr))
})
	})
	return cache_boundedRecord
}

var cache_bottom gopurs_runtime.Value
var once_bottom sync.Once
func Get_bottom() gopurs_runtime.Value {
	once_bottom.Do(func() {
		cache_bottom = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom((*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_bottom
}

var cache_boundedRecordCons gopurs_runtime.Value
var once_boundedRecordCons sync.Once
func Get_boundedRecordCons() gopurs_runtime.Value {
	once_boundedRecordCons.Do(func() {
		cache_boundedRecordCons = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictBounded_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedRecordCons((*Record_reflectSymbol_gopurs_runtime_Value)(dictIsSymbol_0_box.UnsafePtr), (*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dictBounded_1_box.UnsafePtr))
})
	})
	return cache_boundedRecordCons
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

func Call_topRecord(dict_0_loop *Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.topRecord
}

func Call_top(dict_0_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.top
}

func Call_bottomRecord(dict_0_loop *Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.bottomRecord
}

func Call_boundedRecord(_dollar__unused_0_loop *Record_, dictBoundedRecord_1_loop *Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value) gopurs_runtime.Value {
var _dollar__unused_0 *Record_ = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictBoundedRecord_1 *Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value = dictBoundedRecord_1_loop
_ = dictBoundedRecord_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBoundedRecord_1)}, "OrdRecord0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_0
eqRec1_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = eqRec1_3_1
ordRecord1_4_2 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_3_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = ordRecord1_4_2
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecord1_4_2
}), gopurs_runtime.Apply2(dictBoundedRecord_1.bottomRecord, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), gopurs_runtime.Apply2(dictBoundedRecord_1.topRecord, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
}

func Call_bottom(dict_0_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.bottom
}

func Call_boundedRecordCons(dictIsSymbol_0_loop *Record_reflectSymbol_gopurs_runtime_Value, dictBounded_1_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) gopurs_runtime.Value {
var dictIsSymbol_0 *Record_reflectSymbol_gopurs_runtime_Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictBounded_1 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dictBounded_1_loop
_ = dictBounded_1
top1_2_0 := dictBounded_1.top
_ = top1_2_0
bottom1_3_1 := dictBounded_1.bottom
_ = bottom1_3_1
Ord0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBounded_1)}, "Ord0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Ord0_4_2
return gopurs_runtime.Func3(func(_dollar__unused_5 gopurs_runtime.Value, _dollar__unused_6 gopurs_runtime.Value, dictBoundedRecord_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_7, "OrdRecord0"), gopurs_runtime.Value{})
_ = __local_var_8_3
__local_var_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_3, "EqRecord0"), gopurs_runtime.Value{})
_ = __local_var_9_4
__local_var_10_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_4_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_10_5
eqRowCons2_11_7 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func3(func(v_11 gopurs_runtime.Value, ra_12 gopurs_runtime.Value, rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
get_14_8 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Apply(dictIsSymbol_0.reflectSymbol, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = get_14_8
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_5, "eq"), gopurs_runtime.Apply(get_14_8, ra_12), gopurs_runtime.Apply(get_14_8, rb_13)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_4, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, ra_12, rb_13))
}))
_ = eqRowCons2_11_7
ordRecordCons_11_6 := gopurs_runtime.RecordDict2("EqRecord0", "compareRecord", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRowCons2_11_7
}), gopurs_runtime.Func3(func(v_12 gopurs_runtime.Value, ra_13 gopurs_runtime.Value, rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
key_15_9 := gopurs_runtime.Apply(dictIsSymbol_0.reflectSymbol, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})
_ = key_15_9
left_16_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_4_2, "compare"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_15_9, ra_13), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_15_9, rb_14))
_ = left_16_10
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), left_16_10, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Bool(false)).IntVal) != (0) {
__t11 = left_16_10
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_3, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, ra_13, rb_14)
}
end_branch_11:
return __t11
}))
_ = ordRecordCons_11_6
return gopurs_runtime.RecordDict3("OrdRecord0", "bottomRecord", "topRecord", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecordCons_11_6
}), gopurs_runtime.Func2(func(v_12 gopurs_runtime.Value, rowProxy_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.reflectSymbol, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), bottom1_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, rowProxy_13))
}), gopurs_runtime.Func2(func(v_12 gopurs_runtime.Value, rowProxy_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.reflectSymbol, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), top1_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, rowProxy_13))
}))
})
}

func Get_bottomChar() gopurs_runtime.Value {
	return _Gopurs_BottomChar
}

func Get_bottomInt() gopurs_runtime.Value {
	return _Gopurs_BottomInt
}

func Get_bottomNumber() gopurs_runtime.Value {
	return _Gopurs_BottomNumber
}

func Get_topChar() gopurs_runtime.Value {
	return _Gopurs_TopChar
}

func Get_topInt() gopurs_runtime.Value {
	return _Gopurs_TopInt
}

func Get_topNumber() gopurs_runtime.Value {
	return _Gopurs_TopNumber
}
