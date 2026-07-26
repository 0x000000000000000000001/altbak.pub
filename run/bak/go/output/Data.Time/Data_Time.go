package Data_Time

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_negate gopurs_runtime.Value
var once_negate sync.Once
func Get_negate() gopurs_runtime.Value {
	once_negate.Do(func() {
		cache_negate = func() gopurs_runtime.Value {
zero_0_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_0_0
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), zero_0_0, a_1)
})
}()
	})
	return cache_negate
}

var cache_Time gopurs_runtime.Value
var once_Time sync.Once
func Get_Time() gopurs_runtime.Value {
	once_Time.Do(func() {
		cache_Time = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Data_Data_Time_Time{value0, value1, value2, value3})}
})
})
})
})
	})
	return cache_Time
}

var cache_showTime gopurs_runtime.Value
var once_showTime sync.Once
func Get_showTime() gopurs_runtime.Value {
	once_showTime.Do(func() {
		cache_showTime = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Time "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_showHour(), "show"), (*Data_Data_Time_Time)(v_0.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_showMinute(), "show"), (*Data_Data_Time_Time)(v_0.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_showSecond(), "show"), (*Data_Data_Time_Time)(v_0.UnsafePtr).V2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_showMillisecond(), "show"), (*Data_Data_Time_Time)(v_0.UnsafePtr).V3), gopurs_runtime.Str(")")))))))))
}))
	})
	return cache_showTime
}

var cache_setSecond gopurs_runtime.Value
var once_setSecond sync.Once
func Get_setSecond() gopurs_runtime.Value {
	once_setSecond.Do(func() {
		cache_setSecond = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_setSecond(s_0_box, v_1_box)
})
	})
	return cache_setSecond
}

var cache_setMinute gopurs_runtime.Value
var once_setMinute sync.Once
func Get_setMinute() gopurs_runtime.Value {
	once_setMinute.Do(func() {
		cache_setMinute = gopurs_runtime.Func2(func(m_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_setMinute(m_0_box, v_1_box)
})
	})
	return cache_setMinute
}

var cache_setMillisecond gopurs_runtime.Value
var once_setMillisecond sync.Once
func Get_setMillisecond() gopurs_runtime.Value {
	once_setMillisecond.Do(func() {
		cache_setMillisecond = gopurs_runtime.Func2(func(ms_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_setMillisecond(ms_0_box, v_1_box)
})
	})
	return cache_setMillisecond
}

var cache_setHour gopurs_runtime.Value
var once_setHour sync.Once
func Get_setHour() gopurs_runtime.Value {
	once_setHour.Do(func() {
		cache_setHour = gopurs_runtime.Func2(func(h_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_setHour(h_0_box, v_1_box)
})
	})
	return cache_setHour
}

var cache_second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		cache_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_second(v_0_box)
})
	})
	return cache_second
}

var cache_minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		cache_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minute(v_0_box)
})
	})
	return cache_minute
}

var cache_millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		cache_millisecond = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_millisecond(v_0_box)
})
	})
	return cache_millisecond
}

var cache_millisToTime gopurs_runtime.Value
var once_millisToTime sync.Once
func Get_millisToTime() gopurs_runtime.Value {
	once_millisToTime.Do(func() {
		cache_millisToTime = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_millisToTime(v_0_box)
})
	})
	return cache_millisToTime
}

var cache_hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		cache_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hour(v_0_box)
})
	})
	return cache_hour
}

var cache_timeToMillis gopurs_runtime.Value
var once_timeToMillis sync.Once
func Get_timeToMillis() gopurs_runtime.Value {
	once_timeToMillis.Do(func() {
		cache_timeToMillis = gopurs_runtime.Func(func(t_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_timeToMillis(t_0_box)
})
	})
	return cache_timeToMillis
}

var cache_eqTime gopurs_runtime.Value
var once_eqTime sync.Once
func Get_eqTime() gopurs_runtime.Value {
	once_eqTime.Do(func() {
		cache_eqTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_eqHour(), "eq"), (*Data_Data_Time_Time)(x_0.UnsafePtr).V0, (*Data_Data_Time_Time)(y_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_eqMinute(), "eq"), (*Data_Data_Time_Time)(x_0.UnsafePtr).V1, (*Data_Data_Time_Time)(y_1.UnsafePtr).V1)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_eqSecond(), "eq"), (*Data_Data_Time_Time)(x_0.UnsafePtr).V2, (*Data_Data_Time_Time)(y_1.UnsafePtr).V2)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_eqMillisecond(), "eq"), (*Data_Data_Time_Time)(x_0.UnsafePtr).V3, (*Data_Data_Time_Time)(y_1.UnsafePtr).V3))
}))
	})
	return cache_eqTime
}

var cache_ordTime gopurs_runtime.Value
var once_ordTime sync.Once
func Get_ordTime() gopurs_runtime.Value {
	once_ordTime.Do(func() {
		cache_ordTime = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqTime()
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_ordHour(), "compare"), (*Data_Data_Time_Time)(x_0.UnsafePtr).V0, (*Data_Data_Time_Time)(y_1.UnsafePtr).V0)
_ = v_2_0
var __t5 gopurs_runtime.Value
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 1527465420) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 380165415) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_ordMinute(), "compare"), (*Data_Data_Time_Time)(x_0.UnsafePtr).V1, (*Data_Data_Time_Time)(y_1.UnsafePtr).V1)
_ = v1_3_1
var __t4 gopurs_runtime.Value
{
if (v1_3_1.Type == 9 && v1_3_1.IntVal == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (v1_3_1.Type == 9 && v1_3_1.IntVal == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
v2_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_ordSecond(), "compare"), (*Data_Data_Time_Time)(x_0.UnsafePtr).V2, (*Data_Data_Time_Time)(y_1.UnsafePtr).V2)
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 380165415) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_ordMillisecond(), "compare"), (*Data_Data_Time_Time)(x_0.UnsafePtr).V3, (*Data_Data_Time_Time)(y_1.UnsafePtr).V3)
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return __t5
}))
	})
	return cache_ordTime
}

var cache_diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		cache_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, t1_1_box gopurs_runtime.Value, t2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff((*Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value)(dictDuration_0_box.UnsafePtr), t1_1_box, t2_2_box)
})
	})
	return cache_diff
}

var cache_boundedTime gopurs_runtime.Value
var once_boundedTime sync.Once
func Get_boundedTime() gopurs_runtime.Value {
	once_boundedTime.Do(func() {
		cache_boundedTime = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordTime()
}), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Data_Data_Time_Time{gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedHour(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMinute(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedSecond(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMillisecond(), "bottom")})}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Data_Data_Time_Time{gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedHour(), "top"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMinute(), "top"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedSecond(), "top"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMillisecond(), "top")})})
	})
	return cache_boundedTime
}

var cache_maxTime gopurs_runtime.Value
var once_maxTime sync.Once
func Get_maxTime() gopurs_runtime.Value {
	once_maxTime.Do(func() {
		cache_maxTime = gopurs_runtime.Apply(Get_timeToMillis(), gopurs_runtime.RecordGet(Get_boundedTime(), "top"))
	})
	return cache_maxTime
}

var cache_minTime gopurs_runtime.Value
var once_minTime sync.Once
func Get_minTime() gopurs_runtime.Value {
	once_minTime.Do(func() {
		cache_minTime = gopurs_runtime.Apply(Get_timeToMillis(), gopurs_runtime.RecordGet(Get_boundedTime(), "bottom"))
	})
	return cache_minTime
}

var cache_adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		cache_adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_adjust((*Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value)(dictDuration_0_box.UnsafePtr), d_1_box, t_2_box)
})
	})
	return cache_adjust
}

type Data_Data_Time_Time struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}
func Is_Data_Data_Time_Time(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 922918650
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

func Call_setSecond(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Data_Data_Time_Time{(*Data_Data_Time_Time)(v_1.UnsafePtr).V0, (*Data_Data_Time_Time)(v_1.UnsafePtr).V1, s_0, (*Data_Data_Time_Time)(v_1.UnsafePtr).V3})}
}

func Call_setMinute(m_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Data_Data_Time_Time{(*Data_Data_Time_Time)(v_1.UnsafePtr).V0, m_0, (*Data_Data_Time_Time)(v_1.UnsafePtr).V2, (*Data_Data_Time_Time)(v_1.UnsafePtr).V3})}
}

func Call_setMillisecond(ms_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ms_0 gopurs_runtime.Value = ms_0_loop
_ = ms_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Data_Data_Time_Time{(*Data_Data_Time_Time)(v_1.UnsafePtr).V0, (*Data_Data_Time_Time)(v_1.UnsafePtr).V1, (*Data_Data_Time_Time)(v_1.UnsafePtr).V2, ms_0})}
}

func Call_setHour(h_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var h_0 gopurs_runtime.Value = h_0_loop
_ = h_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Data_Data_Time_Time{h_0, (*Data_Data_Time_Time)(v_1.UnsafePtr).V1, (*Data_Data_Time_Time)(v_1.UnsafePtr).V2, (*Data_Data_Time_Time)(v_1.UnsafePtr).V3})}
}

func Call_second(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Time_Time)(v_0.UnsafePtr).V2
}

func Call_minute(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Time_Time)(v_0.UnsafePtr).V1
}

func Call_millisecond(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Time_Time)(v_0.UnsafePtr).V3
}

func Call_millisToTime(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
hours_1_0 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), v_0, gopurs_runtime.Float(3600000.0)))
_ = hours_1_0
minutes_2_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), v_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), hours_1_0, gopurs_runtime.Float(3600000.0))), gopurs_runtime.Float(60000.0)))
_ = minutes_2_1
seconds_3_2 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), v_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), minutes_2_1, gopurs_runtime.Float(60000.0)))), gopurs_runtime.Float(1000.0)))
_ = seconds_3_2
__local_var_4_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Time(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumHour(), "toEnum"), gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), hours_1_0)))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMinute(), "toEnum"), gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), minutes_2_1)))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumSecond(), "toEnum"), gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), seconds_3_2)))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMillisecond(), "toEnum"), gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), v_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), seconds_3_2, gopurs_runtime.Float(1000.0))))))))
_ = __local_var_4_3
var __t4 gopurs_runtime.Value
{
if (__local_var_4_3.Type == 9 && __local_var_4_3.IntVal == 930809136) {
__t4 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_4_3.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}

func Call_hour(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Time_Time)(v_0.UnsafePtr).V0
}

func Call_timeToMillis(t_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var t_0 gopurs_runtime.Value = t_0_loop
_ = t_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), gopurs_runtime.Float(3600000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumHour(), "fromEnum"), (*Data_Data_Time_Time)(t_0.UnsafePtr).V0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), gopurs_runtime.Float(60000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMinute(), "fromEnum"), (*Data_Data_Time_Time)(t_0.UnsafePtr).V1)))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), gopurs_runtime.Float(1000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumSecond(), "fromEnum"), (*Data_Data_Time_Time)(t_0.UnsafePtr).V2)))), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMillisecond(), "fromEnum"), (*Data_Data_Time_Time)(t_0.UnsafePtr).V3)))
}

func Call_diff(dictDuration_0_loop *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value, t1_1_loop gopurs_runtime.Value, t2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value = dictDuration_0_loop
_ = dictDuration_0
var t1_1 gopurs_runtime.Value = t1_1_loop
_ = t1_1
var t2_2 gopurs_runtime.Value = t2_2_loop
_ = t2_2
return gopurs_runtime.Apply(dictDuration_0.toDuration, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), gopurs_runtime.Apply(Get_timeToMillis(), t1_1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "toDuration"), gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negate(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "fromDuration"), gopurs_runtime.Apply(Get_timeToMillis(), t2_2))))))
}

func Call_adjust(dictDuration_0_loop *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value, d_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
d_prime_3_0 := gopurs_runtime.Apply(dictDuration_0.fromDuration, d_1)
_ = d_prime_3_0
wholeDays_4_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), d_prime_3_0, gopurs_runtime.Float(86400000.0)))
_ = wholeDays_4_1
msAdjusted_5_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), gopurs_runtime.Apply(Get_timeToMillis(), t_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), d_prime_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "toDuration"), gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negate(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "fromDuration"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationDays(), "fromDuration"), wholeDays_4_1))))))
_ = msAdjusted_5_2
var __t4 gopurs_runtime.Value
{
if (msAdjusted_5_2.FloatVal()) > (Get_maxTime().FloatVal()) {
__t4 = gopurs_runtime.Float(1.0)
goto end_branch_4
} else {

}
}
{
if (msAdjusted_5_2.FloatVal()) < (Get_minTime().FloatVal()) {
__t4 = gopurs_runtime.Apply(Get_negate(), gopurs_runtime.Float(1.0))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Float(0.0)
}
end_branch_4:
wrap_6_3 := __t4
_ = wrap_6_3
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupDays(), "append"), wholeDays_4_1, wrap_6_3), gopurs_runtime.Apply(Get_millisToTime(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), msAdjusted_5_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), gopurs_runtime.Float(86400000.0), gopurs_runtime.Apply(Get_negate(), wrap_6_3))))})}
}


