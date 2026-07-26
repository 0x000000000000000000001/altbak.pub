package Data_Array_NonEmpty_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_TraversableWithIndex "gopurs/output/Data.TraversableWithIndex"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Semigroup_Traversable "gopurs/output/Data.Semigroup.Traversable"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Data_Show "gopurs/output/Data.Show"
	unsafe "unsafe"
)

var cache_NonEmptyArray gopurs_runtime.Value
var once_NonEmptyArray sync.Once
func Get_NonEmptyArray() gopurs_runtime.Value {
	once_NonEmptyArray.Do(func() {
		cache_NonEmptyArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_NonEmptyArray(x_0_box)
})
	})
	return cache_NonEmptyArray
}

var cache_unfoldable1NonEmptyArray gopurs_runtime.Value
var once_unfoldable1NonEmptyArray sync.Once
func Get_unfoldable1NonEmptyArray() gopurs_runtime.Value {
	once_unfoldable1NonEmptyArray.Do(func() {
		cache_unfoldable1NonEmptyArray = pkg_Data_Unfoldable1.Get_unfoldable1Array()
	})
	return cache_unfoldable1NonEmptyArray
}

var cache_traversableWithIndexNonEmptyArray gopurs_runtime.Value
var once_traversableWithIndexNonEmptyArray sync.Once
func Get_traversableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_traversableWithIndexNonEmptyArray.Do(func() {
		cache_traversableWithIndexNonEmptyArray = pkg_Data_TraversableWithIndex.Get_traversableWithIndexArray()
	})
	return cache_traversableWithIndexNonEmptyArray
}

var cache_traversableNonEmptyArray gopurs_runtime.Value
var once_traversableNonEmptyArray sync.Once
func Get_traversableNonEmptyArray() gopurs_runtime.Value {
	once_traversableNonEmptyArray.Do(func() {
		cache_traversableNonEmptyArray = pkg_Data_Traversable.Get_traversableArray()
	})
	return cache_traversableNonEmptyArray
}

var cache_showNonEmptyArray gopurs_runtime.Value
var once_showNonEmptyArray sync.Once
func Get_showNonEmptyArray() gopurs_runtime.Value {
	once_showNonEmptyArray.Do(func() {
		cache_showNonEmptyArray = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showNonEmptyArray((*Record_show_gopurs_runtime_Value)(dictShow_0_box.UnsafePtr))
})
	})
	return cache_showNonEmptyArray
}

var cache_semigroupNonEmptyArray gopurs_runtime.Value
var once_semigroupNonEmptyArray sync.Once
func Get_semigroupNonEmptyArray() gopurs_runtime.Value {
	once_semigroupNonEmptyArray.Do(func() {
		cache_semigroupNonEmptyArray = pkg_Data_Semigroup.Get_semigroupArray()
	})
	return cache_semigroupNonEmptyArray
}

var cache_ordNonEmptyArray gopurs_runtime.Value
var once_ordNonEmptyArray sync.Once
func Get_ordNonEmptyArray() gopurs_runtime.Value {
	once_ordNonEmptyArray.Do(func() {
		cache_ordNonEmptyArray = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordNonEmptyArray((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_ordNonEmptyArray
}

var cache_ord1NonEmptyArray gopurs_runtime.Value
var once_ord1NonEmptyArray sync.Once
func Get_ord1NonEmptyArray() gopurs_runtime.Value {
	once_ord1NonEmptyArray.Do(func() {
		cache_ord1NonEmptyArray = pkg_Data_Ord.Get_ord1Array()
	})
	return cache_ord1NonEmptyArray
}

var cache_monadNonEmptyArray gopurs_runtime.Value
var once_monadNonEmptyArray sync.Once
func Get_monadNonEmptyArray() gopurs_runtime.Value {
	once_monadNonEmptyArray.Do(func() {
		cache_monadNonEmptyArray = pkg_Control_Monad.Get_monadArray()
	})
	return cache_monadNonEmptyArray
}

var cache_functorWithIndexNonEmptyArray gopurs_runtime.Value
var once_functorWithIndexNonEmptyArray sync.Once
func Get_functorWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyArray.Do(func() {
		cache_functorWithIndexNonEmptyArray = pkg_Data_FunctorWithIndex.Get_functorWithIndexArray()
	})
	return cache_functorWithIndexNonEmptyArray
}

var cache_functorNonEmptyArray gopurs_runtime.Value
var once_functorNonEmptyArray sync.Once
func Get_functorNonEmptyArray() gopurs_runtime.Value {
	once_functorNonEmptyArray.Do(func() {
		cache_functorNonEmptyArray = pkg_Data_Functor.Get_functorArray()
	})
	return cache_functorNonEmptyArray
}

var cache_foldableWithIndexNonEmptyArray gopurs_runtime.Value
var once_foldableWithIndexNonEmptyArray sync.Once
func Get_foldableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyArray.Do(func() {
		cache_foldableWithIndexNonEmptyArray = pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray()
	})
	return cache_foldableWithIndexNonEmptyArray
}

var cache_foldableNonEmptyArray gopurs_runtime.Value
var once_foldableNonEmptyArray sync.Once
func Get_foldableNonEmptyArray() gopurs_runtime.Value {
	once_foldableNonEmptyArray.Do(func() {
		cache_foldableNonEmptyArray = pkg_Data_Foldable.Get_foldableArray()
	})
	return cache_foldableNonEmptyArray
}

var cache_foldable1NonEmptyArray gopurs_runtime.Value
var once_foldable1NonEmptyArray sync.Once
func Get_foldable1NonEmptyArray() gopurs_runtime.Value {
	once_foldable1NonEmptyArray.Do(func() {
		cache_foldable1NonEmptyArray = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), f_2)
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldable1NonEmptyArray(), "foldl1"), append_1_0)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_foldl1Impl(), __local_var_0, __local_var_1)
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_foldr1Impl(), __local_var_0, __local_var_1)
}))
	})
	return cache_foldable1NonEmptyArray
}

var cache_traversable1NonEmptyArray gopurs_runtime.Value
var once_traversable1NonEmptyArray sync.Once
func Get_traversable1NonEmptyArray() gopurs_runtime.Value {
	once_traversable1NonEmptyArray.Do(func() {
		cache_traversable1NonEmptyArray = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1NonEmptyArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableArray()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversable1NonEmptyArray(), "traverse1"), dictApply_0, pkg_Data_Semigroup_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
apply_1_0 := gopurs_runtime.RecordGet(dictApply_0, "apply")
_ = apply_1_0
map__2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map")
_ = map__2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_traverse1Impl(), apply_1_0, map__2_1, f_3)
})
}))
	})
	return cache_traversable1NonEmptyArray
}

var cache_eqNonEmptyArray gopurs_runtime.Value
var once_eqNonEmptyArray sync.Once
func Get_eqNonEmptyArray() gopurs_runtime.Value {
	once_eqNonEmptyArray.Do(func() {
		cache_eqNonEmptyArray = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqNonEmptyArray((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_eqNonEmptyArray
}

var cache_eq1NonEmptyArray gopurs_runtime.Value
var once_eq1NonEmptyArray sync.Once
func Get_eq1NonEmptyArray() gopurs_runtime.Value {
	once_eq1NonEmptyArray.Do(func() {
		cache_eq1NonEmptyArray = pkg_Data_Eq.Get_eq1Array()
	})
	return cache_eq1NonEmptyArray
}

var cache_bindNonEmptyArray gopurs_runtime.Value
var once_bindNonEmptyArray sync.Once
func Get_bindNonEmptyArray() gopurs_runtime.Value {
	once_bindNonEmptyArray.Do(func() {
		cache_bindNonEmptyArray = pkg_Control_Bind.Get_bindArray()
	})
	return cache_bindNonEmptyArray
}

var cache_applyNonEmptyArray gopurs_runtime.Value
var once_applyNonEmptyArray sync.Once
func Get_applyNonEmptyArray() gopurs_runtime.Value {
	once_applyNonEmptyArray.Do(func() {
		cache_applyNonEmptyArray = pkg_Control_Apply.Get_applyArray()
	})
	return cache_applyNonEmptyArray
}

var cache_applicativeNonEmptyArray gopurs_runtime.Value
var once_applicativeNonEmptyArray sync.Once
func Get_applicativeNonEmptyArray() gopurs_runtime.Value {
	once_applicativeNonEmptyArray.Do(func() {
		cache_applicativeNonEmptyArray = pkg_Control_Applicative.Get_applicativeArray()
	})
	return cache_applicativeNonEmptyArray
}

var cache_altNonEmptyArray gopurs_runtime.Value
var once_altNonEmptyArray sync.Once
func Get_altNonEmptyArray() gopurs_runtime.Value {
	once_altNonEmptyArray.Do(func() {
		cache_altNonEmptyArray = pkg_Control_Alt.Get_altArray()
	})
	return cache_altNonEmptyArray
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

func Call_NonEmptyArray(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showNonEmptyArray(dictShow_0_loop *Record_show_gopurs_runtime_Value) gopurs_runtime.Value {
var dictShow_0 *Record_show_gopurs_runtime_Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(NonEmptyArray "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(pkg_Data_Show.Get_showArrayImpl(), dictShow_0.show, v_1), gopurs_runtime.Str(")")))
}))
}

func Call_ordNonEmptyArray(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Ord.Get_ordArray(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_eqNonEmptyArray(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), dictEq_0.eq))
}

func Get_foldl1Impl() gopurs_runtime.Value {
	return _Gopurs_Foldl1Impl
}

func Get_foldr1Impl() gopurs_runtime.Value {
	return _Gopurs_Foldr1Impl
}

func Get_traverse1Impl() gopurs_runtime.Value {
	return _Gopurs_Traverse1Impl
}
