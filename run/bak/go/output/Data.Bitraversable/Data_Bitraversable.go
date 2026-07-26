package Data_Bitraversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Bifoldable "gopurs/output/Data.Bifoldable"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Functor_Joker "gopurs/output/Data.Functor.Joker"
	pkg_Data_Functor_Clown "gopurs/output/Data.Functor.Clown"
	pkg_Data_Functor_Flip "gopurs/output/Data.Functor.Flip"
	pkg_Data_Functor_Product2 "gopurs/output/Data.Functor.Product2"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_bitraverse gopurs_runtime.Value
var once_bitraverse sync.Once
func Get_bitraverse() gopurs_runtime.Value {
	once_bitraverse.Do(func() {
		cache_bitraverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_bitraverse
}

var cache_lfor gopurs_runtime.Value
var once_lfor sync.Once
func Get_lfor() gopurs_runtime.Value {
	once_lfor.Do(func() {
		cache_lfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lfor((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_lfor
}

var cache_ltraverse gopurs_runtime.Value
var once_ltraverse sync.Once
func Get_ltraverse() gopurs_runtime.Value {
	once_ltraverse.Do(func() {
		cache_ltraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ltraverse((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_ltraverse
}

var cache_rfor gopurs_runtime.Value
var once_rfor sync.Once
func Get_rfor() gopurs_runtime.Value {
	once_rfor.Do(func() {
		cache_rfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rfor((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_rfor
}

var cache_rtraverse gopurs_runtime.Value
var once_rtraverse sync.Once
func Get_rtraverse() gopurs_runtime.Value {
	once_rtraverse.Do(func() {
		cache_rtraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rtraverse((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_rtraverse
}

var cache_bitraversableTuple gopurs_runtime.Value
var once_bitraversableTuple sync.Once
func Get_bitraversableTuple() gopurs_runtime.Value {
	once_bitraversableTuple.Do(func() {
		cache_bitraversableTuple = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_1
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_1, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0)), gopurs_runtime.Apply(g_3, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1))
})
}))
	})
	return cache_bitraversableTuple
}

var cache_bitraversableJoker gopurs_runtime.Value
var once_bitraversableJoker sync.Once
func Get_bitraversableJoker() gopurs_runtime.Value {
	once_bitraversableJoker.Do(func() {
		cache_bitraversableJoker = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraversableJoker((*Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value)(dictTraversable_0_box.UnsafePtr))
})
	})
	return cache_bitraversableJoker
}

var cache_bitraversableEither gopurs_runtime.Value
var once_bitraversableEither sync.Once
func Get_bitraversableEither() gopurs_runtime.Value {
	once_bitraversableEither.Do(func() {
		cache_bitraversableEither = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), pkg_Data_Either.Get_Left(), (*pkg_Data_Either.Data_Data_Either_Left)(v_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), pkg_Data_Either.Get_Right(), (*pkg_Data_Either.Data_Data_Either_Right)(v_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
return gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), pkg_Data_Either.Get_Left(), gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Data_Data_Either_Left)(v2_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.Apply(v1_3, (*pkg_Data_Either.Data_Data_Either_Right)(v2_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}))
	})
	return cache_bitraversableEither
}

var cache_bitraversableConst gopurs_runtime.Value
var once_bitraversableConst sync.Once
func Get_bitraversableConst() gopurs_runtime.Value {
	once_bitraversableConst.Do(func() {
		cache_bitraversableConst = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorConst()
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Const.Get_Const(), v_1)
}), gopurs_runtime.Func4(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Const.Get_Const(), gopurs_runtime.Apply(f_1, v1_3))
}))
	})
	return cache_bitraversableConst
}

var cache_bitraversableClown gopurs_runtime.Value
var once_bitraversableClown sync.Once
func Get_bitraversableClown() gopurs_runtime.Value {
	once_bitraversableClown.Do(func() {
		cache_bitraversableClown = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraversableClown((*Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value)(dictTraversable_0_box.UnsafePtr))
})
	})
	return cache_bitraversableClown
}

var cache_bisequenceDefault gopurs_runtime.Value
var once_bisequenceDefault sync.Once
func Get_bisequenceDefault() gopurs_runtime.Value {
	once_bisequenceDefault.Do(func() {
		cache_bisequenceDefault = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequenceDefault((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_bisequenceDefault
}

var cache_bisequence gopurs_runtime.Value
var once_bisequence sync.Once
func Get_bisequence() gopurs_runtime.Value {
	once_bisequence.Do(func() {
		cache_bisequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequence((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_bisequence
}

var cache_bitraversableFlip gopurs_runtime.Value
var once_bitraversableFlip sync.Once
func Get_bitraversableFlip() gopurs_runtime.Value {
	once_bitraversableFlip.Do(func() {
		cache_bitraversableFlip = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraversableFlip((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr))
})
	})
	return cache_bitraversableFlip
}

var cache_bitraversableProduct2 gopurs_runtime.Value
var once_bitraversableProduct2 sync.Once
func Get_bitraversableProduct2() gopurs_runtime.Value {
	once_bitraversableProduct2.Do(func() {
		cache_bitraversableProduct2 = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraversableProduct2((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr))
})
	})
	return cache_bitraversableProduct2
}

var cache_bitraverseDefault gopurs_runtime.Value
var once_bitraverseDefault sync.Once
func Get_bitraverseDefault() gopurs_runtime.Value {
	once_bitraverseDefault.Do(func() {
		cache_bitraverseDefault = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverseDefault((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_bitraverseDefault
}

var cache_bifor gopurs_runtime.Value
var once_bifor sync.Once
func Get_bifor() gopurs_runtime.Value {
	once_bifor.Do(func() {
		cache_bifor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifor((*Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value)(dictBitraversable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_bifor
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

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_bitraverse(dict_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.bitraverse
}

func Call_lfor(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.bitraverse, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
_ = bitraverse2_2_0
pure_3_1 := dictApplicative_1.pure
_ = pure_3_1
return gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, f_5, pure_3_1, t_4)
})
}

func Call_ltraverse(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.bitraverse, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
_ = bitraverse2_2_0
pure_3_1 := dictApplicative_1.pure
_ = pure_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(bitraverse2_2_0, f_4, pure_3_1)
})
}

func Call_rfor(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.bitraverse, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
_ = bitraverse2_2_0
pure_3_1 := dictApplicative_1.pure
_ = pure_3_1
return gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, pure_3_1, f_5, t_4)
})
}

func Call_rtraverse(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(dictBitraversable_0.bitraverse, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, dictApplicative_1.pure)
}

func Call_bitraversableJoker(dictTraversable_0_loop *Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value) gopurs_runtime.Value {
var dictTraversable_0 *Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictTraversable_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorJoker_2_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), g_3, v1_4)
}))
_ = bifunctorJoker_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictTraversable_0)}, "Foldable1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_3_2
bifoldableJoker_4_3 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldMap"), dictMonoid_4)
_ = foldMap1_5_4
return gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_5_4, r_7, v1_8)
})
}), gopurs_runtime.Func4(func(v_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldl"), r_5, u_6, v1_7)
}), gopurs_runtime.Func4(func(v_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldr"), r_5, u_6, v1_7)
}))
_ = bifoldableJoker_4_3
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableJoker_4_3
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker_2_1
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_6_5 := gopurs_runtime.Apply(dictTraversable_0.sequence, dictApplicative_5)
_ = sequence1_6_5
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Joker.Get_Joker(), gopurs_runtime.Apply(sequence1_6_5, v_7))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_6_6 := gopurs_runtime.Apply(dictTraversable_0.traverse, dictApplicative_5)
_ = traverse1_6_6
return gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Joker.Get_Joker(), gopurs_runtime.Apply2(traverse1_6_6, r_8, v1_9))
})
}))
}

func Call_bitraversableClown(dictTraversable_0_loop *Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value) gopurs_runtime.Value {
var dictTraversable_0 *Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictTraversable_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorClown_2_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, v1_4)
}))
_ = bifunctorClown_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictTraversable_0)}, "Foldable1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_3_2
bifoldableClown_4_3 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldMap"), dictMonoid_4)
_ = foldMap1_5_4
return gopurs_runtime.Func3(func(l_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_5_4, l_6, v1_8)
})
}), gopurs_runtime.Func4(func(l_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldl"), l_4, u_6, v1_7)
}), gopurs_runtime.Func4(func(l_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldr"), l_4, u_6, v1_7)
}))
_ = bifoldableClown_4_3
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableClown_4_3
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown_2_1
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_6_5 := gopurs_runtime.Apply(dictTraversable_0.sequence, dictApplicative_5)
_ = sequence1_6_5
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Clown.Get_Clown(), gopurs_runtime.Apply(sequence1_6_5, v_7))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_6_6 := gopurs_runtime.Apply(dictTraversable_0.traverse, dictApplicative_5)
_ = traverse1_6_6
return gopurs_runtime.Func3(func(l_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Clown.Get_Clown(), gopurs_runtime.Apply2(traverse1_6_6, l_7, v1_9))
})
}))
}

func Call_bisequenceDefault(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply3(dictBitraversable_0.bitraverse, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, Get_identity(), Get_identity())
}

func Call_bisequence(dict_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.bisequence
}

func Call_bitraversableFlip(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBitraversable_0)}, "Bifunctor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorFlip_2_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), g_3, f_2, v_4)
}))
_ = bifunctorFlip_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBitraversable_0)}, "Bifoldable1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_3_2
bifoldableFlip_4_3 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap2_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "bifoldMap"), dictMonoid_4)
_ = bifoldMap2_5_4
return gopurs_runtime.Func3(func(r_6 gopurs_runtime.Value, l_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bifoldMap2_5_4, l_7, r_6, v_8)
})
}), gopurs_runtime.Func4(func(r_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_3_2, "bifoldl"), l_5, r_4, u_6, v_7)
}), gopurs_runtime.Func4(func(r_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_3_2, "bifoldr"), l_5, r_4, u_6, v_7)
}))
_ = bifoldableFlip_4_3
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableFlip_4_3
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip_2_1
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
bisequence2_6_5 := gopurs_runtime.Apply(dictBitraversable_0.bisequence, dictApplicative_5)
_ = bisequence2_6_5
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Flip.Get_Flip(), gopurs_runtime.Apply(bisequence2_6_5, v_7))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_6_6 := gopurs_runtime.Apply(dictBitraversable_0.bitraverse, dictApplicative_5)
_ = bitraverse2_6_6
return gopurs_runtime.Func3(func(r_7 gopurs_runtime.Value, l_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Flip.Get_Flip(), gopurs_runtime.Apply3(bitraverse2_6_6, l_8, r_7, v_9))
})
}))
}

func Call_bitraversableProduct2(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBitraversable_0)}, "Bifunctor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifoldableProduct2_2_1 := gopurs_runtime.Apply(pkg_Data_Bifoldable.Get_bifoldableProduct2(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBitraversable_0)}, "Bifoldable1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = bifoldableProduct2_2_1
return gopurs_runtime.Func(func(dictBitraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
bifunctorProduct21_5_3 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, g_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_5, g_6, (*pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2)(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_2, "bimap"), f_5, g_6, (*pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2)(v_7.UnsafePtr).V1)})}
}))
_ = bifunctorProduct21_5_3
bifoldableProduct21_6_4 := gopurs_runtime.Apply(bifoldableProduct2_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifoldable1"), gopurs_runtime.Value{}))
_ = bifoldableProduct21_6_4
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableProduct21_6_4
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct21_5_3
}), gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_7, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_8_5
bisequence3_9_6 := gopurs_runtime.Apply(dictBitraversable_0.bisequence, dictApplicative_7)
_ = bisequence3_9_6
bisequence4_10_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "bisequence"), dictApplicative_7)
_ = bisequence4_10_7
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_8_5, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_5, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product2.Get_Product2(), gopurs_runtime.Apply(bisequence3_9_6, (*pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2)(v_11.UnsafePtr).V0)), gopurs_runtime.Apply(bisequence4_10_7, (*pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2)(v_11.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_7, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_8_8
bitraverse3_9_9 := gopurs_runtime.Apply(dictBitraversable_0.bitraverse, dictApplicative_7)
_ = bitraverse3_9_9
bitraverse4_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "bitraverse"), dictApplicative_7)
_ = bitraverse4_10_10
return gopurs_runtime.Func3(func(l_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_8_8, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_8, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product2.Get_Product2(), gopurs_runtime.Apply3(bitraverse3_9_9, l_11, r_12, (*pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2)(v_13.UnsafePtr).V0)), gopurs_runtime.Apply3(bitraverse4_10_10, l_11, r_12, (*pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2)(v_13.UnsafePtr).V1))
})
}))
})
}

func Call_bitraverseDefault(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
bisequence2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.bisequence, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
_ = bisequence2_2_0
return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(bisequence2_2_0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBitraversable_0)}, "Bifunctor0_NOT_FOUND"), gopurs_runtime.Value{}), "bimap"), f_3, g_4, t_5))
})
}

func Call_bifor(dictBitraversable_0_loop *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBitraversable_0 *Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.bitraverse, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
_ = bitraverse2_2_0
return gopurs_runtime.Func3(func(t_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, f_4, g_5, t_3)
})
}


