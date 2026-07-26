package Data_Time_Duration

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
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

var cache_Seconds gopurs_runtime.Value
var once_Seconds sync.Once
func Get_Seconds() gopurs_runtime.Value {
	once_Seconds.Do(func() {
		cache_Seconds = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Seconds(x_0_box)
})
	})
	return cache_Seconds
}

var cache_Minutes gopurs_runtime.Value
var once_Minutes sync.Once
func Get_Minutes() gopurs_runtime.Value {
	once_Minutes.Do(func() {
		cache_Minutes = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Minutes(x_0_box)
})
	})
	return cache_Minutes
}

var cache_Milliseconds gopurs_runtime.Value
var once_Milliseconds sync.Once
func Get_Milliseconds() gopurs_runtime.Value {
	once_Milliseconds.Do(func() {
		cache_Milliseconds = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Milliseconds(x_0_box)
})
	})
	return cache_Milliseconds
}

var cache_Hours gopurs_runtime.Value
var once_Hours sync.Once
func Get_Hours() gopurs_runtime.Value {
	once_Hours.Do(func() {
		cache_Hours = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Hours(x_0_box)
})
	})
	return cache_Hours
}

var cache_Days gopurs_runtime.Value
var once_Days sync.Once
func Get_Days() gopurs_runtime.Value {
	once_Days.Do(func() {
		cache_Days = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Days(x_0_box)
})
	})
	return cache_Days
}

var cache_toDuration gopurs_runtime.Value
var once_toDuration sync.Once
func Get_toDuration() gopurs_runtime.Value {
	once_toDuration.Do(func() {
		cache_toDuration = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toDuration((*Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_toDuration
}

var cache_showSeconds gopurs_runtime.Value
var once_showSeconds sync.Once
func Get_showSeconds() gopurs_runtime.Value {
	once_showSeconds.Do(func() {
		cache_showSeconds = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Seconds "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showNumber(), "show"), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showSeconds
}

var cache_showMinutes gopurs_runtime.Value
var once_showMinutes sync.Once
func Get_showMinutes() gopurs_runtime.Value {
	once_showMinutes.Do(func() {
		cache_showMinutes = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Minutes "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showNumber(), "show"), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showMinutes
}

var cache_showMilliseconds gopurs_runtime.Value
var once_showMilliseconds sync.Once
func Get_showMilliseconds() gopurs_runtime.Value {
	once_showMilliseconds.Do(func() {
		cache_showMilliseconds = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Milliseconds "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showNumber(), "show"), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showMilliseconds
}

var cache_showHours gopurs_runtime.Value
var once_showHours sync.Once
func Get_showHours() gopurs_runtime.Value {
	once_showHours.Do(func() {
		cache_showHours = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Hours "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showNumber(), "show"), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showHours
}

var cache_showDays gopurs_runtime.Value
var once_showDays sync.Once
func Get_showDays() gopurs_runtime.Value {
	once_showDays.Do(func() {
		cache_showDays = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Days "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showNumber(), "show"), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showDays
}

var cache_semigroupSeconds gopurs_runtime.Value
var once_semigroupSeconds sync.Once
func Get_semigroupSeconds() gopurs_runtime.Value {
	once_semigroupSeconds.Do(func() {
		cache_semigroupSeconds = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), v_0, v1_1)
}))
	})
	return cache_semigroupSeconds
}

var cache_semigroupMinutes gopurs_runtime.Value
var once_semigroupMinutes sync.Once
func Get_semigroupMinutes() gopurs_runtime.Value {
	once_semigroupMinutes.Do(func() {
		cache_semigroupMinutes = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), v_0, v1_1)
}))
	})
	return cache_semigroupMinutes
}

var cache_semigroupMilliseconds gopurs_runtime.Value
var once_semigroupMilliseconds sync.Once
func Get_semigroupMilliseconds() gopurs_runtime.Value {
	once_semigroupMilliseconds.Do(func() {
		cache_semigroupMilliseconds = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), v_0, v1_1)
}))
	})
	return cache_semigroupMilliseconds
}

var cache_semigroupHours gopurs_runtime.Value
var once_semigroupHours sync.Once
func Get_semigroupHours() gopurs_runtime.Value {
	once_semigroupHours.Do(func() {
		cache_semigroupHours = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), v_0, v1_1)
}))
	})
	return cache_semigroupHours
}

var cache_semigroupDays gopurs_runtime.Value
var once_semigroupDays sync.Once
func Get_semigroupDays() gopurs_runtime.Value {
	once_semigroupDays.Do(func() {
		cache_semigroupDays = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), v_0, v1_1)
}))
	})
	return cache_semigroupDays
}

var cache_ordSeconds gopurs_runtime.Value
var once_ordSeconds sync.Once
func Get_ordSeconds() gopurs_runtime.Value {
	once_ordSeconds.Do(func() {
		cache_ordSeconds = pkg_Data_Ord.Get_ordNumber()
	})
	return cache_ordSeconds
}

var cache_ordMinutes gopurs_runtime.Value
var once_ordMinutes sync.Once
func Get_ordMinutes() gopurs_runtime.Value {
	once_ordMinutes.Do(func() {
		cache_ordMinutes = pkg_Data_Ord.Get_ordNumber()
	})
	return cache_ordMinutes
}

var cache_ordMilliseconds gopurs_runtime.Value
var once_ordMilliseconds sync.Once
func Get_ordMilliseconds() gopurs_runtime.Value {
	once_ordMilliseconds.Do(func() {
		cache_ordMilliseconds = pkg_Data_Ord.Get_ordNumber()
	})
	return cache_ordMilliseconds
}

var cache_ordHours gopurs_runtime.Value
var once_ordHours sync.Once
func Get_ordHours() gopurs_runtime.Value {
	once_ordHours.Do(func() {
		cache_ordHours = pkg_Data_Ord.Get_ordNumber()
	})
	return cache_ordHours
}

var cache_ordDays gopurs_runtime.Value
var once_ordDays sync.Once
func Get_ordDays() gopurs_runtime.Value {
	once_ordDays.Do(func() {
		cache_ordDays = pkg_Data_Ord.Get_ordNumber()
	})
	return cache_ordDays
}

var cache_newtypeSeconds gopurs_runtime.Value
var once_newtypeSeconds sync.Once
func Get_newtypeSeconds() gopurs_runtime.Value {
	once_newtypeSeconds.Do(func() {
		cache_newtypeSeconds = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeSeconds
}

var cache_newtypeMinutes gopurs_runtime.Value
var once_newtypeMinutes sync.Once
func Get_newtypeMinutes() gopurs_runtime.Value {
	once_newtypeMinutes.Do(func() {
		cache_newtypeMinutes = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeMinutes
}

var cache_newtypeMilliseconds gopurs_runtime.Value
var once_newtypeMilliseconds sync.Once
func Get_newtypeMilliseconds() gopurs_runtime.Value {
	once_newtypeMilliseconds.Do(func() {
		cache_newtypeMilliseconds = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeMilliseconds
}

var cache_newtypeHours gopurs_runtime.Value
var once_newtypeHours sync.Once
func Get_newtypeHours() gopurs_runtime.Value {
	once_newtypeHours.Do(func() {
		cache_newtypeHours = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeHours
}

var cache_newtypeDays gopurs_runtime.Value
var once_newtypeDays sync.Once
func Get_newtypeDays() gopurs_runtime.Value {
	once_newtypeDays.Do(func() {
		cache_newtypeDays = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeDays
}

var cache_monoidSeconds gopurs_runtime.Value
var once_monoidSeconds sync.Once
func Get_monoidSeconds() gopurs_runtime.Value {
	once_monoidSeconds.Do(func() {
		cache_monoidSeconds = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupSeconds()
}), gopurs_runtime.Float(0.0))
	})
	return cache_monoidSeconds
}

var cache_monoidMinutes gopurs_runtime.Value
var once_monoidMinutes sync.Once
func Get_monoidMinutes() gopurs_runtime.Value {
	once_monoidMinutes.Do(func() {
		cache_monoidMinutes = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupMinutes()
}), gopurs_runtime.Float(0.0))
	})
	return cache_monoidMinutes
}

var cache_monoidMilliseconds gopurs_runtime.Value
var once_monoidMilliseconds sync.Once
func Get_monoidMilliseconds() gopurs_runtime.Value {
	once_monoidMilliseconds.Do(func() {
		cache_monoidMilliseconds = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupMilliseconds()
}), gopurs_runtime.Float(0.0))
	})
	return cache_monoidMilliseconds
}

var cache_monoidHours gopurs_runtime.Value
var once_monoidHours sync.Once
func Get_monoidHours() gopurs_runtime.Value {
	once_monoidHours.Do(func() {
		cache_monoidHours = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupHours()
}), gopurs_runtime.Float(0.0))
	})
	return cache_monoidHours
}

var cache_monoidDays gopurs_runtime.Value
var once_monoidDays sync.Once
func Get_monoidDays() gopurs_runtime.Value {
	once_monoidDays.Do(func() {
		cache_monoidDays = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupDays()
}), gopurs_runtime.Float(0.0))
	})
	return cache_monoidDays
}

var cache_fromDuration gopurs_runtime.Value
var once_fromDuration sync.Once
func Get_fromDuration() gopurs_runtime.Value {
	once_fromDuration.Do(func() {
		cache_fromDuration = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromDuration((*Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_fromDuration
}

var cache_negateDuration gopurs_runtime.Value
var once_negateDuration sync.Once
func Get_negateDuration() gopurs_runtime.Value {
	once_negateDuration.Do(func() {
		cache_negateDuration = gopurs_runtime.Func2(func(dictDuration_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negateDuration((*Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value)(dictDuration_0_box.UnsafePtr), x_1_box)
})
	})
	return cache_negateDuration
}

var cache_eqSeconds gopurs_runtime.Value
var once_eqSeconds sync.Once
func Get_eqSeconds() gopurs_runtime.Value {
	once_eqSeconds.Do(func() {
		cache_eqSeconds = pkg_Data_Eq.Get_eqNumber()
	})
	return cache_eqSeconds
}

var cache_eqMinutes gopurs_runtime.Value
var once_eqMinutes sync.Once
func Get_eqMinutes() gopurs_runtime.Value {
	once_eqMinutes.Do(func() {
		cache_eqMinutes = pkg_Data_Eq.Get_eqNumber()
	})
	return cache_eqMinutes
}

var cache_eqMilliseconds gopurs_runtime.Value
var once_eqMilliseconds sync.Once
func Get_eqMilliseconds() gopurs_runtime.Value {
	once_eqMilliseconds.Do(func() {
		cache_eqMilliseconds = pkg_Data_Eq.Get_eqNumber()
	})
	return cache_eqMilliseconds
}

var cache_eqHours gopurs_runtime.Value
var once_eqHours sync.Once
func Get_eqHours() gopurs_runtime.Value {
	once_eqHours.Do(func() {
		cache_eqHours = pkg_Data_Eq.Get_eqNumber()
	})
	return cache_eqHours
}

var cache_eqDays gopurs_runtime.Value
var once_eqDays sync.Once
func Get_eqDays() gopurs_runtime.Value {
	once_eqDays.Do(func() {
		cache_eqDays = pkg_Data_Eq.Get_eqNumber()
	})
	return cache_eqDays
}

var cache_durationSeconds gopurs_runtime.Value
var once_durationSeconds sync.Once
func Get_durationSeconds() gopurs_runtime.Value {
	once_durationSeconds.Do(func() {
		cache_durationSeconds = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), v_0, gopurs_runtime.Float(1000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), v_0, gopurs_runtime.Float(1000.0))
}))
	})
	return cache_durationSeconds
}

var cache_durationMinutes gopurs_runtime.Value
var once_durationMinutes sync.Once
func Get_durationMinutes() gopurs_runtime.Value {
	once_durationMinutes.Do(func() {
		cache_durationMinutes = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), v_0, gopurs_runtime.Float(60000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), v_0, gopurs_runtime.Float(60000.0))
}))
	})
	return cache_durationMinutes
}

var cache_durationMilliseconds gopurs_runtime.Value
var once_durationMilliseconds sync.Once
func Get_durationMilliseconds() gopurs_runtime.Value {
	once_durationMilliseconds.Do(func() {
		cache_durationMilliseconds = gopurs_runtime.RecordDict2("fromDuration", "toDuration", Get_identity(), Get_identity())
	})
	return cache_durationMilliseconds
}

var cache_durationHours gopurs_runtime.Value
var once_durationHours sync.Once
func Get_durationHours() gopurs_runtime.Value {
	once_durationHours.Do(func() {
		cache_durationHours = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), v_0, gopurs_runtime.Float(3600000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), v_0, gopurs_runtime.Float(3600000.0))
}))
	})
	return cache_durationHours
}

var cache_durationDays gopurs_runtime.Value
var once_durationDays sync.Once
func Get_durationDays() gopurs_runtime.Value {
	once_durationDays.Do(func() {
		cache_durationDays = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), v_0, gopurs_runtime.Float(86400000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), v_0, gopurs_runtime.Float(86400000.0))
}))
	})
	return cache_durationDays
}

var cache_convertDuration gopurs_runtime.Value
var once_convertDuration sync.Once
func Get_convertDuration() gopurs_runtime.Value {
	once_convertDuration.Do(func() {
		cache_convertDuration = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dictDuration1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_convertDuration((*Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value)(dictDuration_0_box.UnsafePtr), (*Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value)(dictDuration1_1_box.UnsafePtr), x_2_box)
})
	})
	return cache_convertDuration
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

func Call_Seconds(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Minutes(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Milliseconds(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Hours(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Days(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_toDuration(dict_0_loop *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.toDuration
}

func Call_fromDuration(dict_0_loop *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.fromDuration
}

func Call_negateDuration(dictDuration_0_loop *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value = dictDuration_0_loop
_ = dictDuration_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictDuration_0.toDuration, gopurs_runtime.Apply(Get_negate(), gopurs_runtime.Apply(dictDuration_0.fromDuration, x_1)))
}

func Call_convertDuration(dictDuration_0_loop *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value, dictDuration1_1_loop *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value = dictDuration_0_loop
_ = dictDuration_0
var dictDuration1_1 *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value = dictDuration1_1_loop
_ = dictDuration1_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictDuration1_1.toDuration, gopurs_runtime.Apply(dictDuration_0.fromDuration, x_2))
}


