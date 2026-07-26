package Data_Bifoldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Unit "gopurs/output/Data.Unit"
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

var cache_monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		cache_monoidEndo = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_0
semigroupEndo1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "compose"), v_1, v1_2)
}))
_ = semigroupEndo1_1_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_1_1
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}()
	})
	return cache_monoidEndo
}

var cache_monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		cache_monoidDual = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monoidEndo(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_0_0
semigroupDual1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "append"), v1_2, v_1)
}))
_ = semigroupDual1_1_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_1_1
}), gopurs_runtime.RecordGet(Get_monoidEndo(), "mempty"))
}()
	})
	return cache_monoidDual
}

var cache_bifoldr gopurs_runtime.Value
var once_bifoldr sync.Once
func Get_bifoldr() gopurs_runtime.Value {
	once_bifoldr.Do(func() {
		cache_bifoldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldr((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_bifoldr
}

var cache_bitraverse_ gopurs_runtime.Value
var once_bitraverse_ sync.Once
func Get_bitraverse_() gopurs_runtime.Value {
	once_bitraverse_.Do(func() {
		cache_bitraverse_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse_((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_bitraverse_
}

var cache_bifor_ gopurs_runtime.Value
var once_bifor_ sync.Once
func Get_bifor_() gopurs_runtime.Value {
	once_bifor_.Do(func() {
		cache_bifor_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifor_((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_bifor_
}

var cache_bisequence_ gopurs_runtime.Value
var once_bisequence_ sync.Once
func Get_bisequence_() gopurs_runtime.Value {
	once_bisequence_.Do(func() {
		cache_bisequence_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequence_((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_pure_gopurs_runtime_Value)(dictApplicative_1_box.UnsafePtr))
})
	})
	return cache_bisequence_
}

var cache_bifoldl gopurs_runtime.Value
var once_bifoldl sync.Once
func Get_bifoldl() gopurs_runtime.Value {
	once_bifoldl.Do(func() {
		cache_bifoldl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldl((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_bifoldl
}

var cache_bifoldableTuple gopurs_runtime.Value
var once_bifoldableTuple sync.Once
func Get_bifoldableTuple() gopurs_runtime.Value {
	once_bifoldableTuple.Do(func() {
		cache_bifoldableTuple = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func4(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), gopurs_runtime.Apply(g_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_3.UnsafePtr).V1))
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(g_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_3.UnsafePtr).V1, z_2))
}))
	})
	return cache_bifoldableTuple
}

var cache_bifoldableJoker gopurs_runtime.Value
var once_bifoldableJoker sync.Once
func Get_bifoldableJoker() gopurs_runtime.Value {
	once_bifoldableJoker.Do(func() {
		cache_bifoldableJoker = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableJoker((*Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value)(dictFoldable_0_box.UnsafePtr))
})
	})
	return cache_bifoldableJoker
}

var cache_bifoldableEither gopurs_runtime.Value
var once_bifoldableEither sync.Once
func Get_bifoldableEither() gopurs_runtime.Value {
	once_bifoldableEither.Do(func() {
		cache_bifoldableEither = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func4(func(dictMonoid_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_1, (*pkg_Data_Either.Data_Data_Either_Left)(v2_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_2, (*pkg_Data_Either.Data_Data_Either_Right)(v2_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value, v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply2(v_0, v2_2, (*pkg_Data_Either.Data_Data_Either_Left)(v3_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(v1_1, v2_2, (*pkg_Data_Either.Data_Data_Either_Right)(v3_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value, v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Either.Data_Data_Either_Left)(v3_3.UnsafePtr).V0, v2_2)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply2(v1_1, (*pkg_Data_Either.Data_Data_Either_Right)(v3_3.UnsafePtr).V0, v2_2)
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
	return cache_bifoldableEither
}

var cache_bifoldableConst gopurs_runtime.Value
var once_bifoldableConst sync.Once
func Get_bifoldableConst() gopurs_runtime.Value {
	once_bifoldableConst.Do(func() {
		cache_bifoldableConst = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func4(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v1_3)
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_2, v1_3)
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v1_3, z_2)
}))
	})
	return cache_bifoldableConst
}

var cache_bifoldableClown gopurs_runtime.Value
var once_bifoldableClown sync.Once
func Get_bifoldableClown() gopurs_runtime.Value {
	once_bifoldableClown.Do(func() {
		cache_bifoldableClown = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableClown((*Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value)(dictFoldable_0_box.UnsafePtr))
})
	})
	return cache_bifoldableClown
}

var cache_bifoldMapDefaultR gopurs_runtime.Value
var once_bifoldMapDefaultR sync.Once
func Get_bifoldMapDefaultR() gopurs_runtime.Value {
	once_bifoldMapDefaultR.Do(func() {
		cache_bifoldMapDefaultR = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMapDefaultR((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_mempty_gopurs_runtime_Value)(dictMonoid_1_box.UnsafePtr))
})
	})
	return cache_bifoldMapDefaultR
}

var cache_bifoldMapDefaultL gopurs_runtime.Value
var once_bifoldMapDefaultL sync.Once
func Get_bifoldMapDefaultL() gopurs_runtime.Value {
	once_bifoldMapDefaultL.Do(func() {
		cache_bifoldMapDefaultL = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMapDefaultL((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_mempty_gopurs_runtime_Value)(dictMonoid_1_box.UnsafePtr))
})
	})
	return cache_bifoldMapDefaultL
}

var cache_bifoldMap gopurs_runtime.Value
var once_bifoldMap sync.Once
func Get_bifoldMap() gopurs_runtime.Value {
	once_bifoldMap.Do(func() {
		cache_bifoldMap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMap((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_bifoldMap
}

var cache_bifoldableFlip gopurs_runtime.Value
var once_bifoldableFlip sync.Once
func Get_bifoldableFlip() gopurs_runtime.Value {
	once_bifoldableFlip.Do(func() {
		cache_bifoldableFlip = gopurs_runtime.Func(func(dictBifoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableFlip((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr))
})
	})
	return cache_bifoldableFlip
}

var cache_bifoldlDefault gopurs_runtime.Value
var once_bifoldlDefault sync.Once
func Get_bifoldlDefault() gopurs_runtime.Value {
	once_bifoldlDefault.Do(func() {
		cache_bifoldlDefault = gopurs_runtime.Func(func(dictBifoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldlDefault((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr))
})
	})
	return cache_bifoldlDefault
}

var cache_bifoldrDefault gopurs_runtime.Value
var once_bifoldrDefault sync.Once
func Get_bifoldrDefault() gopurs_runtime.Value {
	once_bifoldrDefault.Do(func() {
		cache_bifoldrDefault = gopurs_runtime.Func(func(dictBifoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldrDefault((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr))
})
	})
	return cache_bifoldrDefault
}

var cache_bifoldableProduct2 gopurs_runtime.Value
var once_bifoldableProduct2 sync.Once
func Get_bifoldableProduct2() gopurs_runtime.Value {
	once_bifoldableProduct2.Do(func() {
		cache_bifoldableProduct2 = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBifoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableProduct2((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable1_1_box.UnsafePtr))
})
	})
	return cache_bifoldableProduct2
}

var cache_bifold gopurs_runtime.Value
var once_bifold sync.Once
func Get_bifold() gopurs_runtime.Value {
	once_bifold.Do(func() {
		cache_bifold = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifold((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_mempty_gopurs_runtime_Value)(dictMonoid_1_box.UnsafePtr))
})
	})
	return cache_bifold
}

var cache_biany gopurs_runtime.Value
var once_biany sync.Once
func Get_biany() gopurs_runtime.Value {
	once_biany.Do(func() {
		cache_biany = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBooleanAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biany((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_)(dictBooleanAlgebra_1_box.UnsafePtr))
})
	})
	return cache_biany
}

var cache_biall gopurs_runtime.Value
var once_biall sync.Once
func Get_biall() gopurs_runtime.Value {
	once_biall.Do(func() {
		cache_biall = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBooleanAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biall((*Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value)(dictBifoldable_0_box.UnsafePtr), (*Record_)(dictBooleanAlgebra_1_box.UnsafePtr))
})
	})
	return cache_biall
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

func Call_bifoldr(dict_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.bifoldr
}

func Call_bitraverse_(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_0
applySecond_3_1 := gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_identity()
}), a_3), b_4)
})
_ = applySecond_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBifoldable_0.bifoldr, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_3_1, gopurs_runtime.Apply(f_4, x_6))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_3_1, gopurs_runtime.Apply(g_5, x_6))
}), gopurs_runtime.Apply(dictApplicative_1.pure, pkg_Data_Unit.Get_unit()))
})
}

func Call_bifor_(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse_2_2_0 := Call_bitraverse_(dictBifoldable_0, dictApplicative_1)
_ = bitraverse_2_2_0
return gopurs_runtime.Func3(func(t_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse_2_2_0, f_4, g_5, t_3)
})
}

func Call_bisequence_(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictApplicative_1_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *Record_pure_gopurs_runtime_Value = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(Call_bitraverse_(dictBifoldable_0, dictApplicative_1), Get_identity(), Get_identity())
}

func Call_bifoldl(dict_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.bifoldl
}

func Call_bifoldableJoker(dictFoldable_0_loop *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictFoldable_0 *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(dictFoldable_0.foldMap, dictMonoid_1)
_ = foldMap1_2_0
return gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_2_0, r_4, v1_5)
})
}), gopurs_runtime.Func4(func(v_1 gopurs_runtime.Value, r_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictFoldable_0.foldl, r_2, u_3, v1_4)
}), gopurs_runtime.Func4(func(v_1 gopurs_runtime.Value, r_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictFoldable_0.foldr, r_2, u_3, v1_4)
}))
}

func Call_bifoldableClown(dictFoldable_0_loop *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictFoldable_0 *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(dictFoldable_0.foldMap, dictMonoid_1)
_ = foldMap1_2_0
return gopurs_runtime.Func3(func(l_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_2_0, l_3, v1_5)
})
}), gopurs_runtime.Func4(func(l_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictFoldable_0.foldl, l_1, u_3, v1_4)
}), gopurs_runtime.Func4(func(l_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictFoldable_0.foldr, l_1, u_3, v1_4)
}))
}

func Call_bifoldMapDefaultR(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictMonoid_1_loop *Record_mempty_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *Record_mempty_gopurs_runtime_Value = dictMonoid_1_loop
_ = dictMonoid_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, "Semigroup0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_0
mempty_3_1 := dictMonoid_1.mempty
_ = mempty_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBifoldable_0.bifoldr, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.Apply(f_4, x_6))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.Apply(g_5, x_6))
}), mempty_3_1)
})
}

func Call_bifoldMapDefaultL(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictMonoid_1_loop *Record_mempty_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *Record_mempty_gopurs_runtime_Value = dictMonoid_1_loop
_ = dictMonoid_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, "Semigroup0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_0
mempty_3_1 := dictMonoid_1.mempty
_ = mempty_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBifoldable_0.bifoldl, gopurs_runtime.Func2(func(m_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), m_6, gopurs_runtime.Apply(f_4, a_7))
}), gopurs_runtime.Func2(func(m_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), m_6, gopurs_runtime.Apply(g_5, b_7))
}), mempty_3_1)
})
}

func Call_bifoldMap(dict_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.bifoldMap
}

func Call_bifoldableFlip(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap2_2_0 := gopurs_runtime.Apply(dictBifoldable_0.bifoldMap, dictMonoid_1)
_ = bifoldMap2_2_0
return gopurs_runtime.Func3(func(r_3 gopurs_runtime.Value, l_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bifoldMap2_2_0, l_4, r_3, v_5)
})
}), gopurs_runtime.Func4(func(r_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(dictBifoldable_0.bifoldl, l_2, r_1, u_3, v_4)
}), gopurs_runtime.Func4(func(r_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(dictBifoldable_0.bifoldr, l_2, r_1, u_3, v_4)
}))
}

func Call_bifoldlDefault(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
bifoldMap1_1_0 := gopurs_runtime.Apply(dictBifoldable_0.bifoldMap, Get_monoidDual())
_ = bifoldMap1_1_0
return gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, z_4 gopurs_runtime.Value, p_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(bifoldMap1_1_0, gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_7, x_6)
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_3, a_7, x_6)
}), p_5, z_4)
})
}

func Call_bifoldrDefault(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
bifoldMap1_1_0 := gopurs_runtime.Apply(dictBifoldable_0.bifoldMap, Get_monoidEndo())
_ = bifoldMap1_1_0
return gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, z_4 gopurs_runtime.Value, p_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(bifoldMap1_1_0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, x_6)
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_3, x_6)
}), p_5, z_4)
})
}

func Call_bifoldableProduct2(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictBifoldable1_1_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value) gopurs_runtime.Value {
bifoldableProduct2:
for {
if false { continue bifoldableProduct2 }
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBifoldable1_1 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable1_1_loop
_ = dictBifoldable1_1
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap3_3_0 := gopurs_runtime.Apply(dictBifoldable_0.bifoldMap, dictMonoid_2)
_ = bifoldMap3_3_0
bifoldMap4_4_1 := gopurs_runtime.Apply(dictBifoldable1_1.bifoldMap, dictMonoid_2)
_ = bifoldMap4_4_1
return gopurs_runtime.Func3(func(l_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply3(bifoldMap3_3_0, l_5, r_6, (*pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2)(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(bifoldMap4_4_1, l_5, r_6, (*pkg_Data_Functor_Product2.Data_Data_Functor_Product2_Product2)(v_7.UnsafePtr).V1))
})
}), gopurs_runtime.Func4(func(l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, u_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(Get_bifoldlDefault(), Call_bifoldableProduct2(dictBifoldable_0, dictBifoldable1_1), l_2, r_3, u_4, m_5)
}), gopurs_runtime.Func4(func(l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, u_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(Get_bifoldrDefault(), Call_bifoldableProduct2(dictBifoldable_0, dictBifoldable1_1), l_2, r_3, u_4, m_5)
}))
}
}

func Call_bifold(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictMonoid_1_loop *Record_mempty_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *Record_mempty_gopurs_runtime_Value = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(dictBifoldable_0.bifoldMap, gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, Get_identity(), Get_identity())
}

func Call_biany(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictBooleanAlgebra_1_loop *Record_) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBooleanAlgebra_1 *Record_ = dictBooleanAlgebra_1_loop
_ = dictBooleanAlgebra_1
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBooleanAlgebra_1)}, "HeytingAlgebra0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupDisj1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "disj"), v_3, v1_4)
}))
_ = semigroupDisj1_3_2
bifoldMap2_2_0 := gopurs_runtime.Apply(dictBifoldable_0.bifoldMap, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_3_2
}), gopurs_runtime.RecordGet(__local_var_2_1, "ff")))
_ = bifoldMap2_2_0
return gopurs_runtime.Func2(func(p_3 gopurs_runtime.Value, q_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(bifoldMap2_2_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(p_3, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(q_4, x_5)
}))
})
}

func Call_biall(dictBifoldable_0_loop *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value, dictBooleanAlgebra_1_loop *Record_) gopurs_runtime.Value {
var dictBifoldable_0 *Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBooleanAlgebra_1 *Record_ = dictBooleanAlgebra_1_loop
_ = dictBooleanAlgebra_1
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBooleanAlgebra_1)}, "HeytingAlgebra0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupConj1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "conj"), v_3, v1_4)
}))
_ = semigroupConj1_3_2
bifoldMap2_2_0 := gopurs_runtime.Apply(dictBifoldable_0.bifoldMap, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_3_2
}), gopurs_runtime.RecordGet(__local_var_2_1, "tt")))
_ = bifoldMap2_2_0
return gopurs_runtime.Func2(func(p_3 gopurs_runtime.Value, q_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(bifoldMap2_2_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(p_3, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(q_4, x_5)
}))
})
}


