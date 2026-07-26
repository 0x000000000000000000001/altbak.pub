package Data_Functor_Coproduct

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	unsafe "unsafe"
)

var cache_Coproduct gopurs_runtime.Value
var once_Coproduct sync.Once
func Get_Coproduct() gopurs_runtime.Value {
	once_Coproduct.Do(func() {
		cache_Coproduct = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Coproduct(x_0_box)
})
	})
	return cache_Coproduct
}

var cache_showCoproduct gopurs_runtime.Value
var once_showCoproduct sync.Once
func Get_showCoproduct() gopurs_runtime.Value {
	once_showCoproduct.Do(func() {
		cache_showCoproduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showCoproduct((*Record_show_gopurs_runtime_Value)(dictShow_0_box.UnsafePtr), (*Record_show_gopurs_runtime_Value)(dictShow1_1_box.UnsafePtr))
})
	})
	return cache_showCoproduct
}

var cache_right gopurs_runtime.Value
var once_right sync.Once
func Get_right() gopurs_runtime.Value {
	once_right.Do(func() {
		cache_right = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right(ga_0_box)
})
	})
	return cache_right
}

var cache_newtypeCoproduct gopurs_runtime.Value
var once_newtypeCoproduct sync.Once
func Get_newtypeCoproduct() gopurs_runtime.Value {
	once_newtypeCoproduct.Do(func() {
		cache_newtypeCoproduct = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCoproduct
}

var cache_left gopurs_runtime.Value
var once_left sync.Once
func Get_left() gopurs_runtime.Value {
	once_left.Do(func() {
		cache_left = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_left(fa_0_box)
})
	})
	return cache_left
}

var cache_functorCoproduct gopurs_runtime.Value
var once_functorCoproduct sync.Once
func Get_functorCoproduct() gopurs_runtime.Value {
	once_functorCoproduct.Do(func() {
		cache_functorCoproduct = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorCoproduct((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr), (*Record_map__gopurs_runtime_Value)(dictFunctor1_1_box.UnsafePtr))
})
	})
	return cache_functorCoproduct
}

var cache_eq1Coproduct gopurs_runtime.Value
var once_eq1Coproduct sync.Once
func Get_eq1Coproduct() gopurs_runtime.Value {
	once_eq1Coproduct.Do(func() {
		cache_eq1Coproduct = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Coproduct((*Record_eq1_gopurs_runtime_Value)(dictEq1_0_box.UnsafePtr), (*Record_eq1_gopurs_runtime_Value)(dictEq11_1_box.UnsafePtr))
})
	})
	return cache_eq1Coproduct
}

var cache_eqCoproduct gopurs_runtime.Value
var once_eqCoproduct sync.Once
func Get_eqCoproduct() gopurs_runtime.Value {
	once_eqCoproduct.Do(func() {
		cache_eqCoproduct = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqCoproduct((*Record_eq1_gopurs_runtime_Value)(dictEq1_0_box.UnsafePtr), (*Record_eq1_gopurs_runtime_Value)(dictEq11_1_box.UnsafePtr), (*Record_eq_gopurs_runtime_Value)(dictEq_2_box.UnsafePtr))
})
	})
	return cache_eqCoproduct
}

var cache_ord1Coproduct gopurs_runtime.Value
var once_ord1Coproduct sync.Once
func Get_ord1Coproduct() gopurs_runtime.Value {
	once_ord1Coproduct.Do(func() {
		cache_ord1Coproduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Coproduct((*Record_compare1_gopurs_runtime_Value)(dictOrd1_0_box.UnsafePtr))
})
	})
	return cache_ord1Coproduct
}

var cache_ordCoproduct gopurs_runtime.Value
var once_ordCoproduct sync.Once
func Get_ordCoproduct() gopurs_runtime.Value {
	once_ordCoproduct.Do(func() {
		cache_ordCoproduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordCoproduct((*Record_compare1_gopurs_runtime_Value)(dictOrd1_0_box.UnsafePtr))
})
	})
	return cache_ordCoproduct
}

var cache_coproduct gopurs_runtime.Value
var once_coproduct sync.Once
func Get_coproduct() gopurs_runtime.Value {
	once_coproduct.Do(func() {
		cache_coproduct = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct
}

var cache_extendCoproduct gopurs_runtime.Value
var once_extendCoproduct sync.Once
func Get_extendCoproduct() gopurs_runtime.Value {
	once_extendCoproduct.Do(func() {
		cache_extendCoproduct = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendCoproduct((*Record_extend_gopurs_runtime_Value)(dictExtend_0_box.UnsafePtr))
})
	})
	return cache_extendCoproduct
}

var cache_comonadCoproduct gopurs_runtime.Value
var once_comonadCoproduct sync.Once
func Get_comonadCoproduct() gopurs_runtime.Value {
	once_comonadCoproduct.Do(func() {
		cache_comonadCoproduct = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadCoproduct((*Record_extract_gopurs_runtime_Value)(dictComonad_0_box.UnsafePtr))
})
	})
	return cache_comonadCoproduct
}

var cache_bihoistCoproduct gopurs_runtime.Value
var once_bihoistCoproduct sync.Once
func Get_bihoistCoproduct() gopurs_runtime.Value {
	once_bihoistCoproduct.Do(func() {
		cache_bihoistCoproduct = gopurs_runtime.Func3(func(natF_0_box gopurs_runtime.Value, natG_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bihoistCoproduct(natF_0_box, natG_1_box, v_2_box)
})
	})
	return cache_bihoistCoproduct
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

func Call_Coproduct(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showCoproduct(dictShow_0_loop *Record_show_gopurs_runtime_Value, dictShow1_1_loop *Record_show_gopurs_runtime_Value) gopurs_runtime.Value {
var dictShow_0 *Record_show_gopurs_runtime_Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 *Record_show_gopurs_runtime_Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(left "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(dictShow_0.show, (*pkg_Data_Either.Data_Data_Either_Left)(v_2.UnsafePtr).V0), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(right "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(dictShow1_1.show, (*pkg_Data_Either.Data_Data_Either_Right)(v_2.UnsafePtr).V0), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}

func Call_right(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{ga_0})}
}

func Call_left(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{fa_0})}
}

func Call_functorCoproduct(dictFunctor_0_loop *Record_map__gopurs_runtime_Value, dictFunctor1_1_loop *Record_map__gopurs_runtime_Value) gopurs_runtime.Value {
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 *Record_map__gopurs_runtime_Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorEither(), "bimap"), gopurs_runtime.Apply(dictFunctor_0.map_, f_2), gopurs_runtime.Apply(dictFunctor1_1.map_, f_2), v_3)
}))
}

func Call_eq1Coproduct(dictEq1_0_loop *Record_eq1_gopurs_runtime_Value, dictEq11_1_loop *Record_eq1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq1_0 *Record_eq1_gopurs_runtime_Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 *Record_eq1_gopurs_runtime_Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_3_0 := gopurs_runtime.Apply(dictEq1_0.eq1, dictEq_2)
_ = eq12_3_0
eq13_4_1 := gopurs_runtime.Apply(dictEq11_1.eq1, dictEq_2)
_ = eq13_4_1
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
__t2 = gopurs_runtime.Bool(((v1_6.Type == 9 && v1_6.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(eq12_3_0, (*pkg_Data_Either.Data_Data_Either_Left)(v_5.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Left)(v1_6.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(((v_5.Type == 9 && v_5.IntVal == 2465973597)) && (((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(eq13_4_1, (*pkg_Data_Either.Data_Data_Either_Right)(v_5.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Right)(v1_6.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_2:
return __t2
})
}))
}

func Call_eqCoproduct(dictEq1_0_loop *Record_eq1_gopurs_runtime_Value, dictEq11_1_loop *Record_eq1_gopurs_runtime_Value, dictEq_2_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq1_0 *Record_eq1_gopurs_runtime_Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 *Record_eq1_gopurs_runtime_Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 *Record_eq_gopurs_runtime_Value = dictEq_2_loop
_ = dictEq_2
eq12_3_0 := gopurs_runtime.Apply(dictEq1_0.eq1, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_2)})
_ = eq12_3_0
eq13_4_1 := gopurs_runtime.Apply(dictEq11_1.eq1, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_2)})
_ = eq13_4_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
__t2 = gopurs_runtime.Bool(((v1_6.Type == 9 && v1_6.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(eq12_3_0, (*pkg_Data_Either.Data_Data_Either_Left)(v_5.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Left)(v1_6.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(((v_5.Type == 9 && v_5.IntVal == 2465973597)) && (((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(eq13_4_1, (*pkg_Data_Either.Data_Data_Either_Right)(v_5.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Right)(v1_6.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_2:
return __t2
}))
}

func Call_ord1Coproduct(dictOrd1_0_loop *Record_compare1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd1_0 *Record_compare1_gopurs_runtime_Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_1
eq1Coproduct2_4_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), dictEq_4)
_ = eq12_5_3
eq13_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "eq1"), dictEq_4)
_ = eq13_6_4
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 3711209382) {
__t5 = gopurs_runtime.Bool(((v1_8.Type == 9 && v1_8.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(eq12_5_3, (*pkg_Data_Either.Data_Data_Either_Left)(v_7.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Left)(v1_8.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Bool(((v_7.Type == 9 && v_7.IntVal == 2465973597)) && (((v1_8.Type == 9 && v1_8.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(eq13_6_4, (*pkg_Data_Either.Data_Data_Either_Right)(v_7.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Right)(v1_8.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_5:
return __t5
})
}))
_ = eq1Coproduct2_4_2
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Coproduct2_4_2
}), gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
compare12_6_6 := gopurs_runtime.Apply(dictOrd1_0.compare1, dictOrd_5)
_ = compare12_6_6
compare13_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), dictOrd_5)
_ = compare13_7_7
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
var __t9 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 3711209382) {
__t9 = gopurs_runtime.Apply2(compare12_6_6, (*pkg_Data_Either.Data_Data_Either_Left)(v_8.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Left)(v1_9.UnsafePtr).V0)
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_9:
__t8 = __t9
goto end_branch_8
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 3711209382) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_8
} else {

}
}
{
if ((v_8.Type == 9 && v_8.IntVal == 2465973597)) && ((v1_9.Type == 9 && v1_9.IntVal == 2465973597)) {
__t8 = gopurs_runtime.Apply2(compare13_7_7, (*pkg_Data_Either.Data_Data_Either_Right)(v_8.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Right)(v1_9.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
}))
})
}

func Call_ordCoproduct(dictOrd1_0_loop *Record_compare1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd1_0 *Record_compare1_gopurs_runtime_Value = dictOrd1_0_loop
_ = dictOrd1_0
ord1Coproduct1_1_0 := gopurs_runtime.Apply(Get_ord1Coproduct(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)})
_ = ord1Coproduct1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_6_3
eq12_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), __local_var_6_3)
_ = eq12_7_5
eq13_8_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "eq1"), __local_var_6_3)
_ = eq13_8_6
eqCoproduct3_7_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 3711209382) {
__t7 = gopurs_runtime.Bool(((v1_10.Type == 9 && v1_10.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(eq12_7_5, (*pkg_Data_Either.Data_Data_Either_Left)(v_9.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Left)(v1_10.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Bool(((v_9.Type == 9 && v_9.IntVal == 2465973597)) && (((v1_10.Type == 9 && v1_10.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(eq13_8_6, (*pkg_Data_Either.Data_Data_Either_Right)(v_9.UnsafePtr).V0, (*pkg_Data_Either.Data_Data_Either_Right)(v1_10.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_7:
return __t7
}))
_ = eqCoproduct3_7_4
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCoproduct3_7_4
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ord1Coproduct1_1_0, dictOrd11_3), "compare1"), dictOrd_5))
})
})
}

func Call_coproduct(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Data_Data_Either_Left)(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Data_Data_Either_Right)(v2_2.UnsafePtr).V0)
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

func Call_extendCoproduct(dictExtend_0_loop *Record_extend_gopurs_runtime_Value) gopurs_runtime.Value {
var dictExtend_0 *Record_extend_gopurs_runtime_Value = dictExtend_0_loop
_ = dictExtend_0
functorCoproduct1_1_0 := gopurs_runtime.Apply(Get_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictExtend_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = functorCoproduct1_1_0
return gopurs_runtime.Func(func(dictExtend1_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorCoproduct2_3_1 := gopurs_runtime.Apply(functorCoproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct2_3_1
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(dictExtend_0.extend, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{x_5})})
}))
_ = __local_var_5_2
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{x_6})})
}))
_ = __local_var_6_3
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{gopurs_runtime.Apply(__local_var_5_2, (*pkg_Data_Either.Data_Data_Either_Left)(v2_7.UnsafePtr).V0)})}
goto end_branch_4
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Apply(__local_var_6_3, (*pkg_Data_Either.Data_Data_Either_Right)(v2_7.UnsafePtr).V0)})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
}))
})
}

func Call_comonadCoproduct(dictComonad_0_loop *Record_extract_gopurs_runtime_Value) gopurs_runtime.Value {
var dictComonad_0 *Record_extract_gopurs_runtime_Value = dictComonad_0_loop
_ = dictComonad_0
extendCoproduct1_1_0 := gopurs_runtime.Apply(Get_extendCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictComonad_0)}, "Extend0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = extendCoproduct1_1_0
return gopurs_runtime.Func(func(dictComonad1_2 gopurs_runtime.Value) gopurs_runtime.Value {
extendCoproduct2_3_1 := gopurs_runtime.Apply(extendCoproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "Extend0"), gopurs_runtime.Value{}))
_ = extendCoproduct2_3_1
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendCoproduct2_3_1
}), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(dictComonad_0.extract, (*pkg_Data_Either.Data_Data_Either_Left)(v2_4.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "extract"), (*pkg_Data_Either.Data_Data_Either_Right)(v2_4.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
}

func Call_bihoistCoproduct(natF_0_loop gopurs_runtime.Value, natG_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var natF_0 gopurs_runtime.Value = natF_0_loop
_ = natF_0
var natG_1 gopurs_runtime.Value = natG_1_loop
_ = natG_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorEither(), "bimap"), natF_0, natG_1, v_2)
}


