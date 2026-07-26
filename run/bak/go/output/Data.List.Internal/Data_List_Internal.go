package Data_List_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	unsafe "unsafe"
)

var cache_Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		cache_Leaf = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: nil}
	})
	return cache_Leaf
}

var cache_Two gopurs_runtime.Value
var once_Two sync.Once
func Get_Two() gopurs_runtime.Value {
	once_Two.Do(func() {
		cache_Two = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{value0, value1, value2})}
})
})
})
	})
	return cache_Two
}

var cache_Three gopurs_runtime.Value
var once_Three sync.Once
func Get_Three() gopurs_runtime.Value {
	once_Three.Do(func() {
		cache_Three = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{value0, value1, value2, value3, value4})}
})
})
})
})
})
	})
	return cache_Three
}

var cache_TwoLeft gopurs_runtime.Value
var once_TwoLeft sync.Once
func Get_TwoLeft() gopurs_runtime.Value {
	once_TwoLeft.Do(func() {
		cache_TwoLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_TwoLeft{value0, value1})}
})
})
	})
	return cache_TwoLeft
}

var cache_TwoRight gopurs_runtime.Value
var once_TwoRight sync.Once
func Get_TwoRight() gopurs_runtime.Value {
	once_TwoRight.Do(func() {
		cache_TwoRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_TwoRight{value0, value1})}
})
})
	})
	return cache_TwoRight
}

var cache_ThreeLeft gopurs_runtime.Value
var once_ThreeLeft sync.Once
func Get_ThreeLeft() gopurs_runtime.Value {
	once_ThreeLeft.Do(func() {
		cache_ThreeLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeLeft{value0, value1, value2, value3})}
})
})
})
})
	})
	return cache_ThreeLeft
}

var cache_ThreeMiddle gopurs_runtime.Value
var once_ThreeMiddle sync.Once
func Get_ThreeMiddle() gopurs_runtime.Value {
	once_ThreeMiddle.Do(func() {
		cache_ThreeMiddle = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeMiddle{value0, value1, value2, value3})}
})
})
})
})
	})
	return cache_ThreeMiddle
}

var cache_ThreeRight gopurs_runtime.Value
var once_ThreeRight sync.Once
func Get_ThreeRight() gopurs_runtime.Value {
	once_ThreeRight.Do(func() {
		cache_ThreeRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeRight{value0, value1, value2, value3})}
})
})
})
})
	})
	return cache_ThreeRight
}

var cache_KickUp gopurs_runtime.Value
var once_KickUp sync.Once
func Get_KickUp() gopurs_runtime.Value {
	once_KickUp.Do(func() {
		cache_KickUp = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{value0, value1, value2})}
})
})
})
	})
	return cache_KickUp
}

var cache_fromZipper gopurs_runtime.Value
var once_fromZipper sync.Once
func Get_fromZipper() gopurs_runtime.Value {
	once_fromZipper.Do(func() {
		cache_fromZipper = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromZipper(v_0_box, v1_1_box)
})
	})
	return cache_fromZipper
}

var cache_insertAndLookupBy gopurs_runtime.Value
var once_insertAndLookupBy sync.Once
func Get_insertAndLookupBy() gopurs_runtime.Value {
	once_insertAndLookupBy.Do(func() {
		cache_insertAndLookupBy = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, orig_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAndLookupBy(comp_0_box, k_1_box, orig_2_box)
})
	})
	return cache_insertAndLookupBy
}

var cache_emptySet gopurs_runtime.Value
var once_emptySet sync.Once
func Get_emptySet() gopurs_runtime.Value {
	once_emptySet.Do(func() {
		cache_emptySet = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: nil}
	})
	return cache_emptySet
}

type Data_Data_List_Internal_Leaf struct {
	
}
func Is_Data_Data_List_Internal_Leaf(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2764020654
}

type Data_Data_List_Internal_Two struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_Two(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1177901036
}

type Data_Data_List_Internal_Three struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_Three(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1064476974
}

type Data_Data_List_Internal_TwoLeft struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_TwoLeft(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1304506903
}

type Data_Data_List_Internal_TwoRight struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_TwoRight(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2884341868
}

type Data_Data_List_Internal_ThreeLeft struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_ThreeLeft(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2195694037
}

type Data_Data_List_Internal_ThreeMiddle struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_ThreeMiddle(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1584522659
}

type Data_Data_List_Internal_ThreeRight struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_ThreeRight(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3952671150
}

type Data_Data_List_Internal_KickUp struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_KickUp(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2023586927
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

func Call_fromZipper(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromZipper:
for {
if false { continue fromZipper }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1304506903) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{v1_1, (*Data_Data_List_Internal_TwoLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_TwoLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2884341868) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_TwoRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_TwoRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, v1_1})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2195694037) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{v1_1, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V3})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1584522659) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{(*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, v1_1, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V3})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3952671150) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{(*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V3, v1_1})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_insertAndLookupBy(comp_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, orig_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var orig_2 gopurs_runtime.Value = orig_2_loop
_ = orig_2
var up_3_0 gopurs_runtime.Value
up_3_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
up_3_0:
for {
if false { continue up_3_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1304506903) {
__t2 = Call_fromZipper((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2, (*Data_Data_List_Internal_TwoLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_TwoLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2884341868) {
__t2 = Call_fromZipper((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{(*Data_Data_List_Internal_TwoRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_TwoRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2})})
goto end_branch_2
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2195694037) {
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2})}, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 1584522659) {
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0})}, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 3952671150) {
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V2})}, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2})}})}
continue up_3_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
})
var down_4_8 gopurs_runtime.Value
down_4_8 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
down_4_8:
for {
if false { continue down_4_8 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t9 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 2764020654) {
__t9 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(up_3_0, v_5, gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: nil}, k_1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: nil}})}))
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1177901036) {
v2_7_10 := gopurs_runtime.Apply2(comp_0, k_1, (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V1)
_ = v2_7_10
var __t11 gopurs_runtime.Value
{
if (v2_7_10.Type == 9 && v2_7_10.IntVal == 902936544) {
__t11 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_11
} else {

}
}
{
if (v2_7_10.Type == 9 && v2_7_10.IntVal == 1527465420) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_TwoLeft{(*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V1, (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V2})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V0
continue down_4_8
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_TwoRight{(*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V0, (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V1})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V2
continue down_4_8
__t11 = gopurs_runtime.Value{}
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1064476974) {
v2_7_12 := gopurs_runtime.Apply2(comp_0, k_1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V1)
_ = v2_7_12
var __t15 gopurs_runtime.Value
{
if (v2_7_12.Type == 9 && v2_7_12.IntVal == 902936544) {
__t15 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_15
} else {

}
}
{
v3_8_13 := gopurs_runtime.Apply2(comp_0, k_1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V3)
_ = v3_8_13
var __t14 gopurs_runtime.Value
{
if (v3_8_13.Type == 9 && v3_8_13.IntVal == 902936544) {
__t14 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_14
} else {

}
}
{
if (v2_7_12.Type == 9 && v2_7_12.IntVal == 1527465420) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeLeft{(*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V2, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V3, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V4})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V0
continue down_4_8
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
if ((v2_7_12.Type == 9 && v2_7_12.IntVal == 380165415)) && ((v3_8_13.Type == 9 && v3_8_13.IntVal == 1527465420)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeMiddle{(*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V0, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V3, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V4})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V2
continue down_4_8
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeRight{(*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V0, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V2, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V3})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V4
continue down_4_8
__t14 = gopurs_runtime.Value{}
}
end_branch_14:
__t15 = __t14
}
end_branch_15:
__t9 = __t15
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
return gopurs_runtime.Apply2(down_4_8, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, orig_2)
}


