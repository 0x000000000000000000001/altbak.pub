package Data_Either

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_Left gopurs_runtime.Value
var once_Left sync.Once
func Get_Left() gopurs_runtime.Value {
	once_Left.Do(func() {
		cache_Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{value0})}
})
	})
	return cache_Left
}

var cache_Right gopurs_runtime.Value
var once_Right sync.Once
func Get_Right() gopurs_runtime.Value {
	once_Right.Do(func() {
		cache_Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Right{value0})}
})
	})
	return cache_Right
}

var cache_showEither gopurs_runtime.Value
var once_showEither sync.Once
func Get_showEither() gopurs_runtime.Value {
	once_showEither.Do(func() {
		cache_showEither = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showEither((*Record_show_gopurs_runtime_Value)(dictShow_0_box.UnsafePtr), (*Record_show_gopurs_runtime_Value)(dictShow1_1_box.UnsafePtr))
})
	})
	return cache_showEither
}

var cache_note_prime gopurs_runtime.Value
var once_note_prime sync.Once
func Get_note_prime() gopurs_runtime.Value {
	once_note_prime.Do(func() {
		cache_note_prime = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_note_prime(f_0_box, v2_1_box)
})
	})
	return cache_note_prime
}

var cache_note gopurs_runtime.Value
var once_note sync.Once
func Get_note() gopurs_runtime.Value {
	once_note.Do(func() {
		cache_note = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_note(a_0_box, v2_1_box)
})
	})
	return cache_note
}

var cache_genericEither gopurs_runtime.Value
var once_genericEither sync.Once
func Get_genericEither() gopurs_runtime.Value {
	once_genericEither.Do(func() {
		cache_genericEither = gopurs_runtime.RecordDict2("from", "to", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl{(*Data_Data_Either_Left)(x_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr{(*Data_Data_Either_Right)(x_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3478632216) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{(*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(x_0.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 492034566) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Right{(*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(x_0.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
	})
	return cache_genericEither
}

var cache_functorEither gopurs_runtime.Value
var once_functorEither sync.Once
func Get_functorEither() gopurs_runtime.Value {
	once_functorEither.Do(func() {
		cache_functorEither = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{(*Data_Data_Either_Left)(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Right{gopurs_runtime.Apply(f_0, (*Data_Data_Either_Right)(m_1.UnsafePtr).V0)})}
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
	})
	return cache_functorEither
}

var cache_invariantEither gopurs_runtime.Value
var once_invariantEither sync.Once
func Get_invariantEither() gopurs_runtime.Value {
	once_invariantEither.Do(func() {
		cache_invariantEither = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorEither(), "map"), f_0)
}))
	})
	return cache_invariantEither
}

var cache_fromRight_prime gopurs_runtime.Value
var once_fromRight_prime sync.Once
func Get_fromRight_prime() gopurs_runtime.Value {
	once_fromRight_prime.Do(func() {
		cache_fromRight_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromRight_prime(v_0_box, v1_1_box)
})
	})
	return cache_fromRight_prime
}

var cache_fromRight gopurs_runtime.Value
var once_fromRight sync.Once
func Get_fromRight() gopurs_runtime.Value {
	once_fromRight.Do(func() {
		cache_fromRight = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromRight(v_0_box, v1_1_box)
})
	})
	return cache_fromRight
}

var cache_fromLeft_prime gopurs_runtime.Value
var once_fromLeft_prime sync.Once
func Get_fromLeft_prime() gopurs_runtime.Value {
	once_fromLeft_prime.Do(func() {
		cache_fromLeft_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromLeft_prime(v_0_box, v1_1_box)
})
	})
	return cache_fromLeft_prime
}

var cache_fromLeft gopurs_runtime.Value
var once_fromLeft sync.Once
func Get_fromLeft() gopurs_runtime.Value {
	once_fromLeft.Do(func() {
		cache_fromLeft = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromLeft(v_0_box, v1_1_box)
})
	})
	return cache_fromLeft
}

var cache_extendEither gopurs_runtime.Value
var once_extendEither sync.Once
func Get_extendEither() gopurs_runtime.Value {
	once_extendEither.Do(func() {
		cache_extendEither = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{(*Data_Data_Either_Left)(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Right{gopurs_runtime.Apply(v_0, v1_1)})}
}
end_branch_0:
return __t0
}))
	})
	return cache_extendEither
}

var cache_eqEither gopurs_runtime.Value
var once_eqEither sync.Once
func Get_eqEither() gopurs_runtime.Value {
	once_eqEither.Do(func() {
		cache_eqEither = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqEither((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr), (*Record_eq_gopurs_runtime_Value)(dictEq1_1_box.UnsafePtr))
})
	})
	return cache_eqEither
}

var cache_ordEither gopurs_runtime.Value
var once_ordEither sync.Once
func Get_ordEither() gopurs_runtime.Value {
	once_ordEither.Do(func() {
		cache_ordEither = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordEither((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_ordEither
}

var cache_eq1Either gopurs_runtime.Value
var once_eq1Either sync.Once
func Get_eq1Either() gopurs_runtime.Value {
	once_eq1Either.Do(func() {
		cache_eq1Either = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Either((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_eq1Either
}

var cache_ord1Either gopurs_runtime.Value
var once_ord1Either sync.Once
func Get_ord1Either() gopurs_runtime.Value {
	once_ord1Either.Do(func() {
		cache_ord1Either = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Either((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_ord1Either
}

var cache_either gopurs_runtime.Value
var once_either sync.Once
func Get_either() gopurs_runtime.Value {
	once_either.Do(func() {
		cache_either = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either
}

var cache_hush gopurs_runtime.Value
var once_hush sync.Once
func Get_hush() gopurs_runtime.Value {
	once_hush.Do(func() {
		cache_hush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hush(v2_0_box)
})
	})
	return cache_hush
}

var cache_isLeft gopurs_runtime.Value
var once_isLeft sync.Once
func Get_isLeft() gopurs_runtime.Value {
	once_isLeft.Do(func() {
		cache_isLeft = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_isLeft(v2_0_box)
})
	})
	return cache_isLeft
}

var cache_isRight gopurs_runtime.Value
var once_isRight sync.Once
func Get_isRight() gopurs_runtime.Value {
	once_isRight.Do(func() {
		cache_isRight = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_isRight(v2_0_box)
})
	})
	return cache_isRight
}

var cache_choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		cache_choose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose((*Record_alt_gopurs_runtime_Value)(dictAlt_0_box.UnsafePtr))
})
	})
	return cache_choose
}

var cache_boundedEither gopurs_runtime.Value
var once_boundedEither sync.Once
func Get_boundedEither() gopurs_runtime.Value {
	once_boundedEither.Do(func() {
		cache_boundedEither = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedEither((*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dictBounded_0_box.UnsafePtr))
})
	})
	return cache_boundedEither
}

var cache_blush gopurs_runtime.Value
var once_blush sync.Once
func Get_blush() gopurs_runtime.Value {
	once_blush.Do(func() {
		cache_blush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_blush(v2_0_box)
})
	})
	return cache_blush
}

var cache_applyEither gopurs_runtime.Value
var once_applyEither sync.Once
func Get_applyEither() gopurs_runtime.Value {
	once_applyEither.Do(func() {
		cache_applyEither = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{(*Data_Data_Either_Left)(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorEither(), "map"), (*Data_Data_Either_Right)(v_0.UnsafePtr).V0, v1_1)
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
	})
	return cache_applyEither
}

var cache_bindEither gopurs_runtime.Value
var once_bindEither sync.Once
func Get_bindEither() gopurs_runtime.Value {
	once_bindEither.Do(func() {
		cache_bindEither = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__local_var_1_1 := (*Data_Data_Either_Left)(v2_0.UnsafePtr).V0
_ = __local_var_1_1
__t0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{__local_var_1_1})}
})
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__local_var_1_2 := (*Data_Data_Either_Right)(v2_0.UnsafePtr).V0
_ = __local_var_1_2
__t0 = gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, __local_var_1_2)
})
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
	})
	return cache_bindEither
}

var cache_semigroupEither gopurs_runtime.Value
var once_semigroupEither sync.Once
func Get_semigroupEither() gopurs_runtime.Value {
	once_semigroupEither.Do(func() {
		cache_semigroupEither = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupEither((*Record_append__gopurs_runtime_Value)(dictSemigroup_0_box.UnsafePtr))
})
	})
	return cache_semigroupEither
}

var cache_applicativeEither gopurs_runtime.Value
var once_applicativeEither sync.Once
func Get_applicativeEither() gopurs_runtime.Value {
	once_applicativeEither.Do(func() {
		cache_applicativeEither = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}), Get_Right())
	})
	return cache_applicativeEither
}

var cache_monadEither gopurs_runtime.Value
var once_monadEither sync.Once
func Get_monadEither() gopurs_runtime.Value {
	once_monadEither.Do(func() {
		cache_monadEither = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEither()
}))
	})
	return cache_monadEither
}

var cache_altEither gopurs_runtime.Value
var once_altEither sync.Once
func Get_altEither() gopurs_runtime.Value {
	once_altEither.Do(func() {
		cache_altEither = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = v_0
}
end_branch_0:
return __t0
}))
	})
	return cache_altEither
}

type Data_Data_Either_Left struct {
	V0 gopurs_runtime.Value
}
func Is_Data_Data_Either_Left(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3711209382
}

type Data_Data_Either_Right struct {
	V0 gopurs_runtime.Value
}
func Is_Data_Data_Either_Right(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2465973597
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

func Call_showEither(dictShow_0_loop *Record_show_gopurs_runtime_Value, dictShow1_1_loop *Record_show_gopurs_runtime_Value) gopurs_runtime.Value {
var dictShow_0 *Record_show_gopurs_runtime_Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 *Record_show_gopurs_runtime_Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Left "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(dictShow_0.show, (*Data_Data_Either_Left)(v_2.UnsafePtr).V0), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Right "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(dictShow1_1.show, (*Data_Data_Either_Right)(v_2.UnsafePtr).V0), gopurs_runtime.Str(")")))
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

func Call_note_prime(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v2_1 gopurs_runtime.Value = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 3589588149) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 930809136) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Right{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v2_1.UnsafePtr).V0})}
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

func Call_note(a_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 gopurs_runtime.Value = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 3589588149) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{a_0})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 930809136) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Right{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v2_1.UnsafePtr).V0})}
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

func Call_fromRight_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = (*Data_Data_Either_Right)(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_fromRight(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = (*Data_Data_Either_Right)(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = v_0
}
end_branch_0:
return __t0
}

func Call_fromLeft_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = (*Data_Data_Either_Left)(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_fromLeft(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = (*Data_Data_Either_Left)(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = v_0
}
end_branch_0:
return __t0
}

func Call_eqEither(dictEq_0_loop *Record_eq_gopurs_runtime_Value, dictEq1_1_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 *Record_eq_gopurs_runtime_Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Bool(((y_3.Type == 9 && y_3.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(dictEq_0.eq, (*Data_Data_Either_Left)(x_2.UnsafePtr).V0, (*Data_Data_Either_Left)(y_3.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((x_2.Type == 9 && x_2.IntVal == 2465973597)) && (((y_3.Type == 9 && y_3.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(dictEq1_1.eq, (*Data_Data_Either_Right)(x_2.UnsafePtr).V0, (*Data_Data_Either_Right)(y_3.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_0:
return __t0
}))
}

func Call_ordEither(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)}, "Eq0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_1
eqEither2_4_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
__t3 = gopurs_runtime.Bool(((y_5.Type == 9 && y_5.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Data_Data_Either_Left)(x_4.UnsafePtr).V0, (*Data_Data_Either_Left)(y_5.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(((x_4.Type == 9 && x_4.IntVal == 2465973597)) && (((y_5.Type == 9 && y_5.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "eq"), (*Data_Data_Either_Right)(x_4.UnsafePtr).V0, (*Data_Data_Either_Right)(y_5.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_3:
return __t3
}))
_ = eqEither2_4_2
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqEither2_4_2
}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (x_5.Type == 9 && x_5.IntVal == 3711209382) {
var __t5 gopurs_runtime.Value
{
if (y_6.Type == 9 && y_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Apply2(dictOrd_0.compare, (*Data_Data_Either_Left)(x_5.UnsafePtr).V0, (*Data_Data_Either_Left)(y_6.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 3711209382) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if ((x_5.Type == 9 && x_5.IntVal == 2465973597)) && ((y_6.Type == 9 && y_6.IntVal == 2465973597)) {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Data_Data_Either_Right)(x_5.UnsafePtr).V0, (*Data_Data_Either_Right)(y_6.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
})
}

func Call_eq1Either(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Bool(((y_3.Type == 9 && y_3.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(dictEq_0.eq, (*Data_Data_Either_Left)(x_2.UnsafePtr).V0, (*Data_Data_Either_Left)(y_3.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((x_2.Type == 9 && x_2.IntVal == 2465973597)) && (((y_3.Type == 9 && y_3.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Data_Data_Either_Right)(x_2.UnsafePtr).V0, (*Data_Data_Either_Right)(y_3.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_0:
return __t0
}))
}

func Call_ord1Either(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
ordEither1_1_0 := gopurs_runtime.Apply(Get_ordEither(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)})
_ = ordEither1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)}, "Eq0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1Either1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq1_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
__t3 = gopurs_runtime.Bool(((y_5.Type == 9 && y_5.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*Data_Data_Either_Left)(x_4.UnsafePtr).V0, (*Data_Data_Either_Left)(y_5.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(((x_4.Type == 9 && x_4.IntVal == 2465973597)) && (((y_5.Type == 9 && y_5.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Data_Data_Either_Right)(x_4.UnsafePtr).V0, (*Data_Data_Either_Right)(y_5.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_3:
return __t3
}))
_ = eq1Either1_3_2
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Either1_3_2
}), gopurs_runtime.Func(func(dictOrd1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordEither1_1_0, dictOrd1_4), "compare")
}))
}

func Call_either(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Data_Data_Either_Left)(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Data_Data_Either_Right)(v2_2.UnsafePtr).V0)
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

func Call_hush(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Data_Either_Right)(v2_0.UnsafePtr).V0})}
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

func Call_isLeft(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Bool(false)
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

func Call_isRight(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Bool(true)
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

func Call_choose(dictAlt_0_loop *Record_alt_gopurs_runtime_Value) gopurs_runtime.Value {
var dictAlt_0 *Record_alt_gopurs_runtime_Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlt_0)}, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictAlt_0.alt, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), Get_Left(), a_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), Get_Right(), b_3))
})
}

func Call_boundedEither(dictBounded_0_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBounded_0 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dictBounded_0_loop
_ = dictBounded_0
bottom_1_0 := dictBounded_0.bottom
_ = bottom_1_0
ordEither1_2_1 := gopurs_runtime.Apply(Get_ordEither(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBounded_0)}, "Ord0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = ordEither1_2_1
return gopurs_runtime.Func(func(dictBounded1_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordEither2_4_2 := gopurs_runtime.Apply(ordEither1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_3, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither2_4_2
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither2_4_2
}), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Left{bottom_1_0})}, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Data_Data_Either_Right{gopurs_runtime.RecordGet(dictBounded1_3, "top")})})
})
}

func Call_blush(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Data_Either_Left)(v2_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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

func Call_semigroupEither(dictSemigroup_0_loop *Record_append__gopurs_runtime_Value) gopurs_runtime.Value {
var dictSemigroup_0 *Record_append__gopurs_runtime_Value = dictSemigroup_0_loop
_ = dictSemigroup_0
append1_1_0 := dictSemigroup_0.append_
_ = append1_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyEither(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorEither(), "map"), append1_1_0, x_2), y_3)
}))
}


