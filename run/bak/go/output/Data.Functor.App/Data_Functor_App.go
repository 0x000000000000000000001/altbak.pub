package Data_Functor_App

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	unsafe "unsafe"
)

var cache_App gopurs_runtime.Value
var once_App sync.Once
func Get_App() gopurs_runtime.Value {
	once_App.Do(func() {
		cache_App = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_App(x_0_box)
})
	})
	return cache_App
}

var cache_showApp gopurs_runtime.Value
var once_showApp sync.Once
func Get_showApp() gopurs_runtime.Value {
	once_showApp.Do(func() {
		cache_showApp = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showApp((*Record_show_gopurs_runtime_Value)(dictShow_0_box.UnsafePtr))
})
	})
	return cache_showApp
}

var cache_semigroupApp gopurs_runtime.Value
var once_semigroupApp sync.Once
func Get_semigroupApp() gopurs_runtime.Value {
	once_semigroupApp.Do(func() {
		cache_semigroupApp = gopurs_runtime.Func2(func(dictApply_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupApp((*Record_apply_gopurs_runtime_Value)(dictApply_0_box.UnsafePtr), (*Record_append__gopurs_runtime_Value)(dictSemigroup_1_box.UnsafePtr))
})
	})
	return cache_semigroupApp
}

var cache_plusApp gopurs_runtime.Value
var once_plusApp sync.Once
func Get_plusApp() gopurs_runtime.Value {
	once_plusApp.Do(func() {
		cache_plusApp = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_plusApp((*Record_empty_gopurs_runtime_Value)(dictPlus_0_box.UnsafePtr)))}
})
	})
	return cache_plusApp
}

var cache_newtypeApp gopurs_runtime.Value
var once_newtypeApp sync.Once
func Get_newtypeApp() gopurs_runtime.Value {
	once_newtypeApp.Do(func() {
		cache_newtypeApp = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeApp
}

var cache_monoidApp gopurs_runtime.Value
var once_monoidApp sync.Once
func Get_monoidApp() gopurs_runtime.Value {
	once_monoidApp.Do(func() {
		cache_monoidApp = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidApp((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_monoidApp
}

var cache_monadPlusApp gopurs_runtime.Value
var once_monadPlusApp sync.Once
func Get_monadPlusApp() gopurs_runtime.Value {
	once_monadPlusApp.Do(func() {
		cache_monadPlusApp = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadPlusApp((*Record_)(dictMonadPlus_0_box.UnsafePtr)))}
})
	})
	return cache_monadPlusApp
}

var cache_monadApp gopurs_runtime.Value
var once_monadApp sync.Once
func Get_monadApp() gopurs_runtime.Value {
	once_monadApp.Do(func() {
		cache_monadApp = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_monadApp((*Record_)(dictMonad_0_box.UnsafePtr)))}
})
	})
	return cache_monadApp
}

var cache_lazyApp gopurs_runtime.Value
var once_lazyApp sync.Once
func Get_lazyApp() gopurs_runtime.Value {
	once_lazyApp.Do(func() {
		cache_lazyApp = gopurs_runtime.Func(func(dictLazy_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_lazyApp((*Record_defer__gopurs_runtime_Value)(dictLazy_0_box.UnsafePtr)))}
})
	})
	return cache_lazyApp
}

var cache_hoistLowerApp gopurs_runtime.Value
var once_hoistLowerApp sync.Once
func Get_hoistLowerApp() gopurs_runtime.Value {
	once_hoistLowerApp.Do(func() {
		cache_hoistLowerApp = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_hoistLowerApp
}

var cache_hoistLiftApp gopurs_runtime.Value
var once_hoistLiftApp sync.Once
func Get_hoistLiftApp() gopurs_runtime.Value {
	once_hoistLiftApp.Do(func() {
		cache_hoistLiftApp = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_hoistLiftApp
}

var cache_hoistApp gopurs_runtime.Value
var once_hoistApp sync.Once
func Get_hoistApp() gopurs_runtime.Value {
	once_hoistApp.Do(func() {
		cache_hoistApp = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistApp(f_0_box, v_1_box)
})
	})
	return cache_hoistApp
}

var cache_functorApp gopurs_runtime.Value
var once_functorApp sync.Once
func Get_functorApp() gopurs_runtime.Value {
	once_functorApp.Do(func() {
		cache_functorApp = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_functorApp((*Record_map__gopurs_runtime_Value)(dictFunctor_0_box.UnsafePtr)))}
})
	})
	return cache_functorApp
}

var cache_extendApp gopurs_runtime.Value
var once_extendApp sync.Once
func Get_extendApp() gopurs_runtime.Value {
	once_extendApp.Do(func() {
		cache_extendApp = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_extendApp((*Record_extend_gopurs_runtime_Value)(dictExtend_0_box.UnsafePtr)))}
})
	})
	return cache_extendApp
}

var cache_eqApp gopurs_runtime.Value
var once_eqApp sync.Once
func Get_eqApp() gopurs_runtime.Value {
	once_eqApp.Do(func() {
		cache_eqApp = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqApp((*Record_eq1_gopurs_runtime_Value)(dictEq1_0_box.UnsafePtr), (*Record_eq_gopurs_runtime_Value)(dictEq_1_box.UnsafePtr))
})
	})
	return cache_eqApp
}

var cache_ordApp gopurs_runtime.Value
var once_ordApp sync.Once
func Get_ordApp() gopurs_runtime.Value {
	once_ordApp.Do(func() {
		cache_ordApp = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordApp((*Record_compare1_gopurs_runtime_Value)(dictOrd1_0_box.UnsafePtr))
})
	})
	return cache_ordApp
}

var cache_eq1App gopurs_runtime.Value
var once_eq1App sync.Once
func Get_eq1App() gopurs_runtime.Value {
	once_eq1App.Do(func() {
		cache_eq1App = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1App((*Record_eq1_gopurs_runtime_Value)(dictEq1_0_box.UnsafePtr))
})
	})
	return cache_eq1App
}

var cache_ord1App gopurs_runtime.Value
var once_ord1App sync.Once
func Get_ord1App() gopurs_runtime.Value {
	once_ord1App.Do(func() {
		cache_ord1App = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1App((*Record_compare1_gopurs_runtime_Value)(dictOrd1_0_box.UnsafePtr))
})
	})
	return cache_ord1App
}

var cache_comonadApp gopurs_runtime.Value
var once_comonadApp sync.Once
func Get_comonadApp() gopurs_runtime.Value {
	once_comonadApp.Do(func() {
		cache_comonadApp = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_comonadApp((*Record_extract_gopurs_runtime_Value)(dictComonad_0_box.UnsafePtr)))}
})
	})
	return cache_comonadApp
}

var cache_bindApp gopurs_runtime.Value
var once_bindApp sync.Once
func Get_bindApp() gopurs_runtime.Value {
	once_bindApp.Do(func() {
		cache_bindApp = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_bindApp((*Record_bind_gopurs_runtime_Value)(dictBind_0_box.UnsafePtr)))}
})
	})
	return cache_bindApp
}

var cache_applyApp gopurs_runtime.Value
var once_applyApp sync.Once
func Get_applyApp() gopurs_runtime.Value {
	once_applyApp.Do(func() {
		cache_applyApp = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_applyApp((*Record_apply_gopurs_runtime_Value)(dictApply_0_box.UnsafePtr)))}
})
	})
	return cache_applyApp
}

var cache_applicativeApp gopurs_runtime.Value
var once_applicativeApp sync.Once
func Get_applicativeApp() gopurs_runtime.Value {
	once_applicativeApp.Do(func() {
		cache_applicativeApp = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_applicativeApp((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr)))}
})
	})
	return cache_applicativeApp
}

var cache_alternativeApp gopurs_runtime.Value
var once_alternativeApp sync.Once
func Get_alternativeApp() gopurs_runtime.Value {
	once_alternativeApp.Do(func() {
		cache_alternativeApp = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_alternativeApp((*Record_)(dictAlternative_0_box.UnsafePtr)))}
})
	})
	return cache_alternativeApp
}

var cache_altApp gopurs_runtime.Value
var once_altApp sync.Once
func Get_altApp() gopurs_runtime.Value {
	once_altApp.Do(func() {
		cache_altApp = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(Call_altApp((*Record_alt_gopurs_runtime_Value)(dictAlt_0_box.UnsafePtr)))}
})
	})
	return cache_altApp
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

func Call_App(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showApp(dictShow_0_loop *Record_show_gopurs_runtime_Value) gopurs_runtime.Value {
var dictShow_0 *Record_show_gopurs_runtime_Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(App "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(dictShow_0.show, v_1), gopurs_runtime.Str(")")))
}))
}

func Call_semigroupApp(dictApply_0_loop *Record_apply_gopurs_runtime_Value, dictSemigroup_1_loop *Record_append__gopurs_runtime_Value) gopurs_runtime.Value {
var dictApply_0 *Record_apply_gopurs_runtime_Value = dictApply_0_loop
_ = dictApply_0
var dictSemigroup_1 *Record_append__gopurs_runtime_Value = dictSemigroup_1_loop
_ = dictSemigroup_1
append1_2_0 := dictSemigroup_1.append_
_ = append1_2_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.apply, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApply_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}), "map"), append1_2_0, v_3), v1_4)
}))
}

func Call_plusApp(dictPlus_0_loop *Record_empty_gopurs_runtime_Value) *Record_empty_gopurs_runtime_Value {
var dictPlus_0 *Record_empty_gopurs_runtime_Value = dictPlus_0_loop
_ = dictPlus_0
return dictPlus_0
}

func Call_monoidApp(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
append1_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = append1_3_1
semigroupApp2_4_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), append1_3_1, v_4), v1_5)
}))
_ = semigroupApp2_4_2
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupApp2_4_2
}), gopurs_runtime.Apply(dictApplicative_0.pure, gopurs_runtime.RecordGet(dictMonoid_2, "mempty")))
})
}

func Call_monadPlusApp(dictMonadPlus_0_loop *Record_) *Record_ {
var dictMonadPlus_0 *Record_ = dictMonadPlus_0_loop
_ = dictMonadPlus_0
return dictMonadPlus_0
}

func Call_monadApp(dictMonad_0_loop *Record_) *Record_ {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
return dictMonad_0
}

func Call_lazyApp(dictLazy_0_loop *Record_defer__gopurs_runtime_Value) *Record_defer__gopurs_runtime_Value {
var dictLazy_0 *Record_defer__gopurs_runtime_Value = dictLazy_0_loop
_ = dictLazy_0
return dictLazy_0
}

func Call_hoistApp(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_functorApp(dictFunctor_0_loop *Record_map__gopurs_runtime_Value) *Record_map__gopurs_runtime_Value {
var dictFunctor_0 *Record_map__gopurs_runtime_Value = dictFunctor_0_loop
_ = dictFunctor_0
return dictFunctor_0
}

func Call_extendApp(dictExtend_0_loop *Record_extend_gopurs_runtime_Value) *Record_extend_gopurs_runtime_Value {
var dictExtend_0 *Record_extend_gopurs_runtime_Value = dictExtend_0_loop
_ = dictExtend_0
return dictExtend_0
}

func Call_eqApp(dictEq1_0_loop *Record_eq1_gopurs_runtime_Value, dictEq_1_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq1_0 *Record_eq1_gopurs_runtime_Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 *Record_eq_gopurs_runtime_Value = dictEq_1_loop
_ = dictEq_1
eq11_2_0 := gopurs_runtime.Apply(dictEq1_0.eq1, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_1)})
_ = eq11_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_2_0, x_3, y_4)
}))
}

func Call_ordApp(dictOrd1_0_loop *Record_compare1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd1_0 *Record_compare1_gopurs_runtime_Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(dictOrd1_0.compare1, dictOrd_2)
_ = compare11_3_1
eq11_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_4_2
eqApp2_5_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_4_2, x_5, y_6)
}))
_ = eqApp2_5_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_5_3
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_3_1, x_6, y_7)
}))
})
}

func Call_eq1App(dictEq1_0_loop *Record_eq1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq1_0 *Record_eq1_gopurs_runtime_Value = dictEq1_0_loop
_ = dictEq1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictEq1_0.eq1, dictEq_1)
}))
}

func Call_ord1App(dictOrd1_0_loop *Record_compare1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd1_0 *Record_compare1_gopurs_runtime_Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd1_0)}, "Eq10_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1App1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), dictEq_3)
}))
_ = eq1App1_3_2
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1App1_3_2
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_5_3 := gopurs_runtime.Apply(dictOrd1_0.compare1, dictOrd_4)
_ = compare11_5_3
eq11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_4, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_6_4
eqApp2_7_5 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_6_4, x_7, y_8)
}))
_ = eqApp2_7_5
return gopurs_runtime.RecordGet(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_7_5
}), gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_5_3, x_8, y_9)
})), "compare")
}))
}

func Call_comonadApp(dictComonad_0_loop *Record_extract_gopurs_runtime_Value) *Record_extract_gopurs_runtime_Value {
var dictComonad_0 *Record_extract_gopurs_runtime_Value = dictComonad_0_loop
_ = dictComonad_0
return dictComonad_0
}

func Call_bindApp(dictBind_0_loop *Record_bind_gopurs_runtime_Value) *Record_bind_gopurs_runtime_Value {
var dictBind_0 *Record_bind_gopurs_runtime_Value = dictBind_0_loop
_ = dictBind_0
return dictBind_0
}

func Call_applyApp(dictApply_0_loop *Record_apply_gopurs_runtime_Value) *Record_apply_gopurs_runtime_Value {
var dictApply_0 *Record_apply_gopurs_runtime_Value = dictApply_0_loop
_ = dictApply_0
return dictApply_0
}

func Call_applicativeApp(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) *Record_pure_gopurs_runtime_Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
return dictApplicative_0
}

func Call_alternativeApp(dictAlternative_0_loop *Record_) *Record_ {
var dictAlternative_0 *Record_ = dictAlternative_0_loop
_ = dictAlternative_0
return dictAlternative_0
}

func Call_altApp(dictAlt_0_loop *Record_alt_gopurs_runtime_Value) *Record_alt_gopurs_runtime_Value {
var dictAlt_0 *Record_alt_gopurs_runtime_Value = dictAlt_0_loop
_ = dictAlt_0
return dictAlt_0
}


