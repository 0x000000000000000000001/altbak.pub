package Data_Enum_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	unsafe "unsafe"
)

var cache_greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		cache_greaterThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420)) != (true))
})
}()
	})
	return cache_greaterThanOrEq
}

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
})
}()
	})
	return cache_lessThan
}

var cache_genericToEnum_prime gopurs_runtime.Value
var once_genericToEnum_prime sync.Once
func Get_genericToEnum_prime() gopurs_runtime.Value {
	once_genericToEnum_prime.Do(func() {
		cache_genericToEnum_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericToEnum_prime((*Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericToEnum_prime
}

var cache_genericToEnum gopurs_runtime.Value
var once_genericToEnum sync.Once
func Get_genericToEnum() gopurs_runtime.Value {
	once_genericToEnum.Do(func() {
		cache_genericToEnum = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericToEnum((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr))
})
	})
	return cache_genericToEnum
}

var cache_genericSucc_prime gopurs_runtime.Value
var once_genericSucc_prime sync.Once
func Get_genericSucc_prime() gopurs_runtime.Value {
	once_genericSucc_prime.Do(func() {
		cache_genericSucc_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSucc_prime((*Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericSucc_prime
}

var cache_genericSucc gopurs_runtime.Value
var once_genericSucc sync.Once
func Get_genericSucc() gopurs_runtime.Value {
	once_genericSucc.Do(func() {
		cache_genericSucc = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSucc((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr))
})
	})
	return cache_genericSucc
}

var cache_genericPred_prime gopurs_runtime.Value
var once_genericPred_prime sync.Once
func Get_genericPred_prime() gopurs_runtime.Value {
	once_genericPred_prime.Do(func() {
		cache_genericPred_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericPred_prime((*Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericPred_prime
}

var cache_genericPred gopurs_runtime.Value
var once_genericPred sync.Once
func Get_genericPred() gopurs_runtime.Value {
	once_genericPred.Do(func() {
		cache_genericPred = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericPred((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr))
})
	})
	return cache_genericPred
}

var cache_genericFromEnum_prime gopurs_runtime.Value
var once_genericFromEnum_prime sync.Once
func Get_genericFromEnum_prime() gopurs_runtime.Value {
	once_genericFromEnum_prime.Do(func() {
		cache_genericFromEnum_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFromEnum_prime((*Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericFromEnum_prime
}

var cache_genericFromEnum gopurs_runtime.Value
var once_genericFromEnum sync.Once
func Get_genericFromEnum() gopurs_runtime.Value {
	once_genericFromEnum.Do(func() {
		cache_genericFromEnum = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFromEnum((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value)(dictGenericBoundedEnum_1_box.UnsafePtr), x_2_box)
})
	})
	return cache_genericFromEnum
}

var cache_genericEnumSum gopurs_runtime.Value
var once_genericEnumSum sync.Once
func Get_genericEnumSum() gopurs_runtime.Value {
	once_genericEnumSum.Do(func() {
		cache_genericEnumSum = gopurs_runtime.Func2(func(dictGenericEnum_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumSum((*Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value)(dictGenericEnum_0_box.UnsafePtr), (*Record_genericTop_prime_gopurs_runtime_Value)(dictGenericTop_1_box.UnsafePtr))
})
	})
	return cache_genericEnumSum
}

var cache_genericEnumProduct gopurs_runtime.Value
var once_genericEnumProduct sync.Once
func Get_genericEnumProduct() gopurs_runtime.Value {
	once_genericEnumProduct.Do(func() {
		cache_genericEnumProduct = gopurs_runtime.Func5(func(dictGenericEnum_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value, dictGenericBottom_2_box gopurs_runtime.Value, dictGenericEnum1_3_box gopurs_runtime.Value, dictGenericTop1_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumProduct((*Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value)(dictGenericEnum_0_box.UnsafePtr), (*Record_genericTop_prime_gopurs_runtime_Value)(dictGenericTop_1_box.UnsafePtr), (*Record_genericBottom_prime_gopurs_runtime_Value)(dictGenericBottom_2_box.UnsafePtr), (*Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value)(dictGenericEnum1_3_box.UnsafePtr), (*Record_genericTop_prime_gopurs_runtime_Value)(dictGenericTop1_4_box.UnsafePtr))
})
	})
	return cache_genericEnumProduct
}

var cache_genericEnumNoArguments gopurs_runtime.Value
var once_genericEnumNoArguments sync.Once
func Get_genericEnumNoArguments() gopurs_runtime.Value {
	once_genericEnumNoArguments.Do(func() {
		cache_genericEnumNoArguments = gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}))
	})
	return cache_genericEnumNoArguments
}

var cache_genericEnumConstructor gopurs_runtime.Value
var once_genericEnumConstructor sync.Once
func Get_genericEnumConstructor() gopurs_runtime.Value {
	once_genericEnumConstructor.Do(func() {
		cache_genericEnumConstructor = gopurs_runtime.Func(func(dictGenericEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumConstructor((*Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value)(dictGenericEnum_0_box.UnsafePtr))
})
	})
	return cache_genericEnumConstructor
}

var cache_genericEnumArgument gopurs_runtime.Value
var once_genericEnumArgument sync.Once
func Get_genericEnumArgument() gopurs_runtime.Value {
	once_genericEnumArgument.Do(func() {
		cache_genericEnumArgument = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumArgument((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dictEnum_0_box.UnsafePtr))
})
	})
	return cache_genericEnumArgument
}

var cache_genericCardinality_prime gopurs_runtime.Value
var once_genericCardinality_prime sync.Once
func Get_genericCardinality_prime() gopurs_runtime.Value {
	once_genericCardinality_prime.Do(func() {
		cache_genericCardinality_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericCardinality_prime((*Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_genericCardinality_prime
}

var cache_genericCardinality gopurs_runtime.Value
var once_genericCardinality sync.Once
func Get_genericCardinality() gopurs_runtime.Value {
	once_genericCardinality.Do(func() {
		cache_genericCardinality = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericCardinality((*Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value)(dictGeneric_0_box.UnsafePtr), (*Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value)(dictGenericBoundedEnum_1_box.UnsafePtr))
})
	})
	return cache_genericCardinality
}

var cache_genericBoundedEnumSum gopurs_runtime.Value
var once_genericBoundedEnumSum sync.Once
func Get_genericBoundedEnumSum() gopurs_runtime.Value {
	once_genericBoundedEnumSum.Do(func() {
		cache_genericBoundedEnumSum = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBoundedEnumSum((*Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value)(dictGenericBoundedEnum_0_box.UnsafePtr))
})
	})
	return cache_genericBoundedEnumSum
}

var cache_genericBoundedEnumProduct gopurs_runtime.Value
var once_genericBoundedEnumProduct sync.Once
func Get_genericBoundedEnumProduct() gopurs_runtime.Value {
	once_genericBoundedEnumProduct.Do(func() {
		cache_genericBoundedEnumProduct = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBoundedEnumProduct((*Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value)(dictGenericBoundedEnum_0_box.UnsafePtr))
})
	})
	return cache_genericBoundedEnumProduct
}

var cache_genericBoundedEnumNoArguments gopurs_runtime.Value
var once_genericBoundedEnumNoArguments sync.Once
func Get_genericBoundedEnumNoArguments() gopurs_runtime.Value {
	once_genericBoundedEnumNoArguments.Do(func() {
		cache_genericBoundedEnumNoArguments = gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(1), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (i_0.IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}))
	})
	return cache_genericBoundedEnumNoArguments
}

var cache_genericBoundedEnumConstructor gopurs_runtime.Value
var once_genericBoundedEnumConstructor sync.Once
func Get_genericBoundedEnumConstructor() gopurs_runtime.Value {
	once_genericBoundedEnumConstructor.Do(func() {
		cache_genericBoundedEnumConstructor = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBoundedEnumConstructor((*Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value)(dictGenericBoundedEnum_0_box.UnsafePtr))
})
	})
	return cache_genericBoundedEnumConstructor
}

var cache_genericBoundedEnumArgument gopurs_runtime.Value
var once_genericBoundedEnumArgument sync.Once
func Get_genericBoundedEnumArgument() gopurs_runtime.Value {
	once_genericBoundedEnumArgument.Do(func() {
		cache_genericBoundedEnumArgument = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBoundedEnumArgument((*Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value)(dictBoundedEnum_0_box.UnsafePtr))
})
	})
	return cache_genericBoundedEnumArgument
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

func Call_genericToEnum_prime(dict_0_loop *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericToEnum_prime
}

func Call_genericToEnum(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := dictGeneric_0.to
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_2, "genericToEnum'"), x_4))
})
})
}

func Call_genericSucc_prime(dict_0_loop *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericSucc_prime
}

func Call_genericSucc(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := dictGeneric_0.to
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_2, "genericSucc'"), gopurs_runtime.Apply(dictGeneric_0.from, x_4)))
})
})
}

func Call_genericPred_prime(dict_0_loop *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericPred_prime
}

func Call_genericPred(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := dictGeneric_0.to
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_2, "genericPred'"), gopurs_runtime.Apply(dictGeneric_0.from, x_4)))
})
})
}

func Call_genericFromEnum_prime(dict_0_loop *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericFromEnum_prime
}

func Call_genericFromEnum(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericBoundedEnum_1_loop *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictGenericBoundedEnum_1.genericFromEnum_prime, gopurs_runtime.Apply(dictGeneric_0.from, x_2))
}

func Call_genericEnumSum(dictGenericEnum_0_loop *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value, dictGenericTop_1_loop *Record_genericTop_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericEnum_0 *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
var dictGenericTop_1 *Record_genericTop_prime_gopurs_runtime_Value = dictGenericTop_1_loop
_ = dictGenericTop_1
genericTop_prime_2_0 := dictGenericTop_1.genericTop_prime
_ = genericTop_prime_2_0
return gopurs_runtime.Func2(func(dictGenericEnum1_3 gopurs_runtime.Value, dictGenericBottom_4 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime_5_1 := gopurs_runtime.RecordGet(dictGenericBottom_4, "genericBottom'")
_ = genericBottom_prime_5_1
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3478632216) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inl(), gopurs_runtime.Apply(dictGenericEnum_0.genericPred_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(v_6.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 492034566) {
v1_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericPred'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(v_6.UnsafePtr).V0)
_ = v1_7_3
var __t4 gopurs_runtime.Value
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 3589588149) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl{genericTop_prime_2_0})}})}
goto end_branch_4
} else {

}
}
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_7_3.UnsafePtr).V0})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t2 = __t4
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3478632216) {
v1_7_6 := gopurs_runtime.Apply(dictGenericEnum_0.genericSucc_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(v_6.UnsafePtr).V0)
_ = v1_7_6
var __t7 gopurs_runtime.Value
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 3589588149) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr{genericBottom_prime_5_1})}})}
goto end_branch_7
} else {

}
}
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_7_6.UnsafePtr).V0})}})}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t5 = __t7
goto end_branch_5
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 492034566) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inr(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(v_6.UnsafePtr).V0))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
})
}

func Call_genericEnumProduct(dictGenericEnum_0_loop *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value, dictGenericTop_1_loop *Record_genericTop_prime_gopurs_runtime_Value, dictGenericBottom_2_loop *Record_genericBottom_prime_gopurs_runtime_Value, dictGenericEnum1_3_loop *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value, dictGenericTop1_4_loop *Record_genericTop_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericEnum_0 *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
var dictGenericTop_1 *Record_genericTop_prime_gopurs_runtime_Value = dictGenericTop_1_loop
_ = dictGenericTop_1
var dictGenericBottom_2 *Record_genericBottom_prime_gopurs_runtime_Value = dictGenericBottom_2_loop
_ = dictGenericBottom_2
var dictGenericEnum1_3 *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value = dictGenericEnum1_3_loop
_ = dictGenericEnum1_3
var dictGenericTop1_4 *Record_genericTop_prime_gopurs_runtime_Value = dictGenericTop1_4_loop
_ = dictGenericTop1_4
genericTop_prime_5_0 := dictGenericTop1_4.genericTop_prime
_ = genericTop_prime_5_0
return gopurs_runtime.Func(func(dictGenericBottom1_6 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime_7_1 := gopurs_runtime.RecordGet(dictGenericBottom1_6, "genericBottom'")
_ = genericBottom_prime_7_1
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
v1_9_2 := gopurs_runtime.Apply(dictGenericEnum1_3.genericPred_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_8.UnsafePtr).V1)
_ = v1_9_2
var __t3 gopurs_runtime.Value
{
if (v1_9_2.Type == 9 && v1_9_2.IntVal == 930809136) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{(*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_8.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_9_2.UnsafePtr).V0})}})}
goto end_branch_3
} else {

}
}
{
if (v1_9_2.Type == 9 && v1_9_2.IntVal == 3589588149) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{a_10, genericTop_prime_5_0})}
}), gopurs_runtime.Apply(dictGenericEnum_0.genericPred_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_8.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
v1_9_4 := gopurs_runtime.Apply(dictGenericEnum1_3.genericSucc_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_8.UnsafePtr).V1)
_ = v1_9_4
var __t5 gopurs_runtime.Value
{
if (v1_9_4.Type == 9 && v1_9_4.IntVal == 930809136) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{(*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_8.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_9_4.UnsafePtr).V0})}})}
goto end_branch_5
} else {

}
}
{
if (v1_9_4.Type == 9 && v1_9_4.IntVal == 3589588149) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{a_10, genericBottom_prime_7_1})}
}), gopurs_runtime.Apply(dictGenericEnum_0.genericSucc_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_8.UnsafePtr).V0))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
})
}

func Call_genericEnumConstructor(dictGenericEnum_0_loop *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericEnum_0 *Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(dictGenericEnum_0.genericPred_prime, v_1))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(dictGenericEnum_0.genericSucc_prime, v_1))
}))
}

func Call_genericEnumArgument(dictEnum_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEnum_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dictEnum_0_loop
_ = dictEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(dictEnum_0.pred, v_1))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(dictEnum_0.succ, v_1))
}))
}

func Call_genericCardinality_prime(dict_0_loop *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.genericCardinality_prime
}

func Call_genericCardinality(dictGeneric_0_loop *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value, dictGenericBoundedEnum_1_loop *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGeneric_0 *Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
return dictGenericBoundedEnum_1.genericCardinality_prime
}

func Call_genericBoundedEnumSum(dictGenericBoundedEnum_0_loop *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
genericCardinality_prime1_1_0 := dictGenericBoundedEnum_0.genericCardinality_prime
_ = genericCardinality_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int((genericCardinality_prime1_1_0.IntVal) + (gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'").IntVal)), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 3478632216) {
__t1 = gopurs_runtime.Apply(dictGenericBoundedEnum_0.genericFromEnum_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(v_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 492034566) {
__t1 = gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(v_3.UnsafePtr).V0).IntVal) + (genericCardinality_prime1_1_0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThanOrEq(), n_3, gopurs_runtime.Int(0)), gopurs_runtime.Apply2(Get_lessThan(), n_3, genericCardinality_prime1_1_0)).IntVal) != (0) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inl(), gopurs_runtime.Apply(dictGenericBoundedEnum_0.genericToEnum_prime, n_3))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inr(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int((n_3.IntVal) - (genericCardinality_prime1_1_0.IntVal))))
}
end_branch_2:
return __t2
}))
})
}

func Call_genericBoundedEnumProduct(dictGenericBoundedEnum_0_loop *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
genericCardinality_prime1_1_0 := dictGenericBoundedEnum_0.genericCardinality_prime
_ = genericCardinality_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
genericCardinality_prime2_3_1 := gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'")
_ = genericCardinality_prime2_3_1
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int((genericCardinality_prime1_1_0.IntVal) * (genericCardinality_prime2_3_1.IntVal)), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(((gopurs_runtime.Apply(dictGenericBoundedEnum_0.genericFromEnum_prime, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_4.UnsafePtr).V0).IntVal) * (genericCardinality_prime2_3_1.IntVal)) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v1_4.UnsafePtr).V1).IntVal))
}), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Product(), gopurs_runtime.Apply(dictGenericBoundedEnum_0.genericToEnum_prime, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "div"), n_4, genericCardinality_prime2_3_1))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), n_4, genericCardinality_prime2_3_1)))
}))
})
}

func Call_genericBoundedEnumConstructor(dictGenericBoundedEnum_0_loop *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 *Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", dictGenericBoundedEnum_0.genericCardinality_prime, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGenericBoundedEnum_0.genericFromEnum_prime, v_1)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(dictGenericBoundedEnum_0.genericToEnum_prime, i_1))
}))
}

func Call_genericBoundedEnumArgument(dictBoundedEnum_0_loop *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBoundedEnum_0 *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", dictBoundedEnum_0.cardinality, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictBoundedEnum_0.fromEnum, v_1)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(dictBoundedEnum_0.toEnum, i_1))
}))
}


