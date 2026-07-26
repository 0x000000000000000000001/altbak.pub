package Test_RBTree

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	unsafe "unsafe"
)

var cache_R gopurs_runtime.Value
var once_R sync.Once
func Get_R() gopurs_runtime.Value {
	once_R.Do(func() {
		cache_R = gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}
	})
	return cache_R
}

var cache_B gopurs_runtime.Value
var once_B sync.Once
func Get_B() gopurs_runtime.Value {
	once_B.Do(func() {
		cache_B = gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}
	})
	return cache_B
}

var cache_E gopurs_runtime.Value
var once_E sync.Once
func Get_E() gopurs_runtime.Value {
	once_E.Do(func() {
		cache_E = gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: nil}
	})
	return cache_E
}

var cache_T gopurs_runtime.Value
var once_T sync.Once
func Get_T() gopurs_runtime.Value {
	once_T.Do(func() {
		cache_T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{value0, value1, value2.IntVal, value3})}
})
})
})
})
	})
	return cache_T
}

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max(x_0_box.IntVal, y_1_box.IntVal)
})
	})
	return cache_max
}

var cache_makeBlack gopurs_runtime.Value
var once_makeBlack sync.Once
func Get_makeBlack() gopurs_runtime.Value {
	once_makeBlack.Do(func() {
		cache_makeBlack = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeBlack(v_0_box)
})
	})
	return cache_makeBlack
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return cache_describe
}

var cache_depth gopurs_runtime.Value
var once_depth sync.Once
func Get_depth() gopurs_runtime.Value {
	once_depth.Do(func() {
		cache_depth = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_depth(v_0_box)
})
	})
	return cache_depth
}

var cache_balance gopurs_runtime.Value
var once_balance sync.Once
func Get_balance() gopurs_runtime.Value {
	once_balance.Do(func() {
		cache_balance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_balance(v_0_box, v1_1_box, v2_2_box.IntVal, v3_3_box)
})
	})
	return cache_balance
}

var cache_ins gopurs_runtime.Value
var once_ins sync.Once
func Get_ins() gopurs_runtime.Value {
	once_ins.Do(func() {
		cache_ins = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ins(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_ins
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert(x_0_box.IntVal, s_1_box)
})
	})
	return cache_insert
}

var cache_buildTree gopurs_runtime.Value
var once_buildTree sync.Once
func Get_buildTree() gopurs_runtime.Value {
	once_buildTree.Do(func() {
		cache_buildTree = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_buildTree(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_buildTree
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(100000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Apply(Get_depth(), Call_buildTree(dummy_0.IntVal, gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: nil}))))
}))
	})
	return cache_act
}

type Data_Test_RBTree_R struct {
	
}
func Is_Data_Test_RBTree_R(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3668501016
}

type Data_Test_RBTree_B struct {
	
}
func Is_Data_Test_RBTree_B(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1583507464
}

type Data_Test_RBTree_E struct {
	
}
func Is_Data_Test_RBTree_E(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1548554223
}

type Data_Test_RBTree_T struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 int64
	V3 gopurs_runtime.Value
}
func Is_Data_Test_RBTree_T(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3983586014
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

func Call_max(x_0_loop int64, y_1_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
var __t0 gopurs_runtime.Value
{
if (x_0) > (y_1) {
__t0 = gopurs_runtime.Int(x_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Int(y_1)
}
end_branch_0:
return __t0
}

func Call_makeBlack(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3983586014) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)(v_0.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)(v_0.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v_0.UnsafePtr).V3})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1548554223) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: nil}
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

func Call_depth(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
depth:
for {
if false { continue depth }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1548554223) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3983586014) {
__local_var_1_1 := gopurs_runtime.Apply(Get_depth(), (*Data_Test_RBTree_T)(v_0.UnsafePtr).V1)
_ = __local_var_1_1
__local_var_2_2 := gopurs_runtime.Apply(Get_depth(), (*Data_Test_RBTree_T)(v_0.UnsafePtr).V3)
_ = __local_var_2_2
var __t3 gopurs_runtime.Value
{
if (__local_var_1_1.IntVal) > (__local_var_2_2.IntVal) {
__t3 = gopurs_runtime.Int((1) + (__local_var_1_1.IntVal))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int((1) + (__local_var_2_2.IntVal))
}
end_branch_3:
__t0 = __t3
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

func Call_balance(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop int64, v3_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 int64 = v2_2_loop
_ = v2_2
var v3_3 gopurs_runtime.Value = v3_3_loop
_ = v3_3
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1583507464) {
var __t1 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3983586014) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 3668501016) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3983586014) {
var __t6 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 3668501016) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V3})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3, v2_2, v3_3})}})}
goto end_branch_6
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3983586014) {
var __t9 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 3668501016) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V3, v2_2, v3_3})}})}
goto end_branch_9
} else {

}
}
{
var __t_and_12 bool = false
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) {

var __t_tag_11 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0
__t_and_12 = (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 3668501016)
}
if __t_and_12 {
var __t13 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3983586014) {
var __t15 gopurs_runtime.Value
{
var __t_tag_16 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0
if (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 3668501016) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_15
} else {

}
}
{
var __t_tag_17 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_19 bool = false
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 3983586014) {

var __t_tag_18 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_19 = (__t_tag_18.Type == 9 && __t_tag_18.IntVal == 3668501016)
}
if __t_and_19 {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_15:
__t13 = __t15
goto end_branch_13
} else {

}
}
{
var __t_tag_20 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_22 bool = false
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 3983586014) {

var __t_tag_21 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_22 = (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 3668501016)
}
if __t_and_22 {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_13:
__t9 = __t13
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_9:
__t6 = __t9
goto end_branch_6
} else {

}
}
{
var __t_and_24 bool = false
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) {

var __t_tag_23 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0
__t_and_24 = (__t_tag_23.Type == 9 && __t_tag_23.IntVal == 3668501016)
}
if __t_and_24 {
var __t25 gopurs_runtime.Value
{
var __t_tag_26 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1
if (__t_tag_26.Type == 9 && __t_tag_26.IntVal == 3983586014) {
var __t27 gopurs_runtime.Value
{
var __t_tag_28 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0
if (__t_tag_28.Type == 9 && __t_tag_28.IntVal == 3668501016) {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_27
} else {

}
}
{
var __t_tag_29 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_31 bool = false
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 3983586014) {

var __t_tag_30 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_31 = (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 3668501016)
}
if __t_and_31 {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_27:
__t25 = __t27
goto end_branch_25
} else {

}
}
{
var __t_tag_32 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_34 bool = false
if (__t_tag_32.Type == 9 && __t_tag_32.IntVal == 3983586014) {

var __t_tag_33 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_34 = (__t_tag_33.Type == 9 && __t_tag_33.IntVal == 3668501016)
}
if __t_and_34 {
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_25:
__t6 = __t25
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
var __t_tag_35 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3
if (__t_tag_35.Type == 9 && __t_tag_35.IntVal == 3983586014) {
var __t36 gopurs_runtime.Value
{
var __t_tag_37 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V0
if (__t_tag_37.Type == 9 && __t_tag_37.IntVal == 3668501016) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V3, v2_2, v3_3})}})}
goto end_branch_36
} else {

}
}
{
var __t_and_39 bool = false
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) {

var __t_tag_38 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0
__t_and_39 = (__t_tag_38.Type == 9 && __t_tag_38.IntVal == 3668501016)
}
if __t_and_39 {
var __t40 gopurs_runtime.Value
{
var __t_tag_41 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1
if (__t_tag_41.Type == 9 && __t_tag_41.IntVal == 3983586014) {
var __t42 gopurs_runtime.Value
{
var __t_tag_43 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0
if (__t_tag_43.Type == 9 && __t_tag_43.IntVal == 3668501016) {
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_42
} else {

}
}
{
var __t_tag_44 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_46 bool = false
if (__t_tag_44.Type == 9 && __t_tag_44.IntVal == 3983586014) {

var __t_tag_45 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_46 = (__t_tag_45.Type == 9 && __t_tag_45.IntVal == 3668501016)
}
if __t_and_46 {
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_42
} else {

}
}
{
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_42:
__t40 = __t42
goto end_branch_40
} else {

}
}
{
var __t_tag_47 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_49 bool = false
if (__t_tag_47.Type == 9 && __t_tag_47.IntVal == 3983586014) {

var __t_tag_48 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_49 = (__t_tag_48.Type == 9 && __t_tag_48.IntVal == 3668501016)
}
if __t_and_49 {
__t40 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_40:
__t36 = __t40
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_36:
__t4 = __t36
goto end_branch_4
} else {

}
}
{
var __t_and_51 bool = false
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) {

var __t_tag_50 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0
__t_and_51 = (__t_tag_50.Type == 9 && __t_tag_50.IntVal == 3668501016)
}
if __t_and_51 {
var __t52 gopurs_runtime.Value
{
var __t_tag_53 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1
if (__t_tag_53.Type == 9 && __t_tag_53.IntVal == 3983586014) {
var __t54 gopurs_runtime.Value
{
var __t_tag_55 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0
if (__t_tag_55.Type == 9 && __t_tag_55.IntVal == 3668501016) {
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_54
} else {

}
}
{
var __t_tag_56 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_58 bool = false
if (__t_tag_56.Type == 9 && __t_tag_56.IntVal == 3983586014) {

var __t_tag_57 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_58 = (__t_tag_57.Type == 9 && __t_tag_57.IntVal == 3668501016)
}
if __t_and_58 {
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_54:
__t52 = __t54
goto end_branch_52
} else {

}
}
{
var __t_tag_59 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_61 bool = false
if (__t_tag_59.Type == 9 && __t_tag_59.IntVal == 3983586014) {

var __t_tag_60 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_61 = (__t_tag_60.Type == 9 && __t_tag_60.IntVal == 3668501016)
}
if __t_and_61 {
__t52 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_52
} else {

}
}
{
__t52 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_52:
__t4 = __t52
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_4:
__t2 = __t4
goto end_branch_2
} else {

}
}
{
var __t_and_63 bool = false
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) {

var __t_tag_62 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0
__t_and_63 = (__t_tag_62.Type == 9 && __t_tag_62.IntVal == 3668501016)
}
if __t_and_63 {
var __t64 gopurs_runtime.Value
{
var __t_tag_65 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1
if (__t_tag_65.Type == 9 && __t_tag_65.IntVal == 3983586014) {
var __t66 gopurs_runtime.Value
{
var __t_tag_67 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0
if (__t_tag_67.Type == 9 && __t_tag_67.IntVal == 3668501016) {
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_66
} else {

}
}
{
var __t_tag_68 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_70 bool = false
if (__t_tag_68.Type == 9 && __t_tag_68.IntVal == 3983586014) {

var __t_tag_69 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_70 = (__t_tag_69.Type == 9 && __t_tag_69.IntVal == 3668501016)
}
if __t_and_70 {
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_66
} else {

}
}
{
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_66:
__t64 = __t66
goto end_branch_64
} else {

}
}
{
var __t_tag_71 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_73 bool = false
if (__t_tag_71.Type == 9 && __t_tag_71.IntVal == 3983586014) {

var __t_tag_72 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_73 = (__t_tag_72.Type == 9 && __t_tag_72.IntVal == 3668501016)
}
if __t_and_73 {
__t64 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_64
} else {

}
}
{
__t64 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_64:
__t2 = __t64
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
var __t_and_75 bool = false
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) {

var __t_tag_74 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0
__t_and_75 = (__t_tag_74.Type == 9 && __t_tag_74.IntVal == 3668501016)
}
if __t_and_75 {
var __t76 gopurs_runtime.Value
{
var __t_tag_77 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1
if (__t_tag_77.Type == 9 && __t_tag_77.IntVal == 3983586014) {
var __t78 gopurs_runtime.Value
{
var __t_tag_79 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0
if (__t_tag_79.Type == 9 && __t_tag_79.IntVal == 3668501016) {
__t78 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_78
} else {

}
}
{
var __t_tag_80 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_82 bool = false
if (__t_tag_80.Type == 9 && __t_tag_80.IntVal == 3983586014) {

var __t_tag_81 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_82 = (__t_tag_81.Type == 9 && __t_tag_81.IntVal == 3668501016)
}
if __t_and_82 {
__t78 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_78
} else {

}
}
{
__t78 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_78:
__t76 = __t78
goto end_branch_76
} else {

}
}
{
var __t_tag_83 gopurs_runtime.Value = (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3
var __t_and_85 bool = false
if (__t_tag_83.Type == 9 && __t_tag_83.IntVal == 3983586014) {

var __t_tag_84 gopurs_runtime.Value = (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0
__t_and_85 = (__t_tag_84.Type == 9 && __t_tag_84.IntVal == 3668501016)
}
if __t_and_85 {
__t76 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, gopurs_runtime.Int((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2).IntVal, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_76
} else {

}
}
{
__t76 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_76:
__t1 = __t76
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_0:
return __t0
}

func Call_ins(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
ins:
for {
if false { continue ins }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 1548554223) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: nil}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 3983586014) {
var __t1 gopurs_runtime.Value
{
if (v_0) < (gopurs_runtime.Int((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2).IntVal) {
__t1 = Call_balance((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0, Call_ins(v_0, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1), gopurs_runtime.Int((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3)
goto end_branch_1
} else {

}
}
{
if (v_0) > (gopurs_runtime.Int((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2).IntVal) {
__t1 = Call_balance((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2).IntVal, Call_ins(v_0, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{(*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3})}
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

func Call_insert(x_0_loop int64, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
__local_var_2_0 := Call_ins(x_0, s_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3983586014) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Data_Test_RBTree_T)(__local_var_2_0.UnsafePtr).V1, gopurs_runtime.Int((*Data_Test_RBTree_T)(__local_var_2_0.UnsafePtr).V2).IntVal, (*Data_Test_RBTree_T)(__local_var_2_0.UnsafePtr).V3})}
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 1548554223) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: nil}
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

func Call_buildTree(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
buildTree:
for {
if false { continue buildTree }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = Call_insert(v_0, v1_1)
continue buildTree
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0
}
}


