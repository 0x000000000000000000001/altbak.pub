package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Time_Duration_negate gopurs_runtime.Value
var once_Data_Time_Duration_negate sync.Once
func Get_Data_Time_Duration_negate() gopurs_runtime.Value {
	once_Data_Time_Duration_negate.Do(func() {
		cache_Data_Time_Duration_negate = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
}), Get_Data_Ring_numSub())
_ = __local_var_0_0
// TAST (Let): Semiring0_1_1 -> *Constructor_Data_Semiring_Semiring
Semiring0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_1_1
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "sub"), gopurs_runtime.Box(Semiring0_1_1.V3), a_2)
})
}()
	})
	return cache_Data_Time_Duration_negate
}

var cache_Data_Time_Duration_identity gopurs_runtime.Value
var once_Data_Time_Duration_identity sync.Once
func Get_Data_Time_Duration_identity() gopurs_runtime.Value {
	once_Data_Time_Duration_identity.Do(func() {
		cache_Data_Time_Duration_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Time_Duration_identity(x_0_box.FloatVal()))
})
	})
	return cache_Data_Time_Duration_identity
}

var cache_Data_Time_Duration_Seconds gopurs_runtime.Value
var once_Data_Time_Duration_Seconds sync.Once
func Get_Data_Time_Duration_Seconds() gopurs_runtime.Value {
	once_Data_Time_Duration_Seconds.Do(func() {
		cache_Data_Time_Duration_Seconds = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_Seconds(x_0_box)
})
	})
	return cache_Data_Time_Duration_Seconds
}

var cache_Data_Time_Duration_Minutes gopurs_runtime.Value
var once_Data_Time_Duration_Minutes sync.Once
func Get_Data_Time_Duration_Minutes() gopurs_runtime.Value {
	once_Data_Time_Duration_Minutes.Do(func() {
		cache_Data_Time_Duration_Minutes = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_Minutes(x_0_box)
})
	})
	return cache_Data_Time_Duration_Minutes
}

var cache_Data_Time_Duration_Milliseconds gopurs_runtime.Value
var once_Data_Time_Duration_Milliseconds sync.Once
func Get_Data_Time_Duration_Milliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_Milliseconds.Do(func() {
		cache_Data_Time_Duration_Milliseconds = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_Milliseconds(x_0_box)
})
	})
	return cache_Data_Time_Duration_Milliseconds
}

var cache_Data_Time_Duration_Hours gopurs_runtime.Value
var once_Data_Time_Duration_Hours sync.Once
func Get_Data_Time_Duration_Hours() gopurs_runtime.Value {
	once_Data_Time_Duration_Hours.Do(func() {
		cache_Data_Time_Duration_Hours = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_Hours(x_0_box)
})
	})
	return cache_Data_Time_Duration_Hours
}

var cache_Data_Time_Duration_Duration_dollarDict gopurs_runtime.Value
var once_Data_Time_Duration_Duration_dollarDict sync.Once
func Get_Data_Time_Duration_Duration_dollarDict() gopurs_runtime.Value {
	once_Data_Time_Duration_Duration_dollarDict.Do(func() {
		cache_Data_Time_Duration_Duration_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_Duration_dollarDict(x_0_box)
})
	})
	return cache_Data_Time_Duration_Duration_dollarDict
}

var cache_Data_Time_Duration_Days gopurs_runtime.Value
var once_Data_Time_Duration_Days sync.Once
func Get_Data_Time_Duration_Days() gopurs_runtime.Value {
	once_Data_Time_Duration_Days.Do(func() {
		cache_Data_Time_Duration_Days = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_Days(x_0_box)
})
	})
	return cache_Data_Time_Duration_Days
}

var cache_Data_Time_Duration_toDuration gopurs_runtime.Value
var once_Data_Time_Duration_toDuration sync.Once
func Get_Data_Time_Duration_toDuration() gopurs_runtime.Value {
	once_Data_Time_Duration_toDuration.Do(func() {
		cache_Data_Time_Duration_toDuration = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_toDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dict_0_box))
})
	})
	return cache_Data_Time_Duration_toDuration
}

var cache_Data_Time_Duration_showSeconds gopurs_runtime.Value
var once_Data_Time_Duration_showSeconds sync.Once
func Get_Data_Time_Duration_showSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_showSeconds.Do(func() {
		cache_Data_Time_Duration_showSeconds = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Seconds ") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(v_0.FloatVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_Time_Duration_showSeconds
}

var cache_Data_Time_Duration_showMinutes gopurs_runtime.Value
var once_Data_Time_Duration_showMinutes sync.Once
func Get_Data_Time_Duration_showMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_showMinutes.Do(func() {
		cache_Data_Time_Duration_showMinutes = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Minutes ") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(v_0.FloatVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_Time_Duration_showMinutes
}

var cache_Data_Time_Duration_showMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_showMilliseconds sync.Once
func Get_Data_Time_Duration_showMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_showMilliseconds.Do(func() {
		cache_Data_Time_Duration_showMilliseconds = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Milliseconds ") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(v_0.FloatVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_Time_Duration_showMilliseconds
}

var cache_Data_Time_Duration_showHours gopurs_runtime.Value
var once_Data_Time_Duration_showHours sync.Once
func Get_Data_Time_Duration_showHours() gopurs_runtime.Value {
	once_Data_Time_Duration_showHours.Do(func() {
		cache_Data_Time_Duration_showHours = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Hours ") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(v_0.FloatVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_Time_Duration_showHours
}

var cache_Data_Time_Duration_showDays gopurs_runtime.Value
var once_Data_Time_Duration_showDays sync.Once
func Get_Data_Time_Duration_showDays() gopurs_runtime.Value {
	once_Data_Time_Duration_showDays.Do(func() {
		cache_Data_Time_Duration_showDays = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Days ") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(v_0.FloatVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_Time_Duration_showDays
}

var cache_Data_Time_Duration_semigroupSeconds gopurs_runtime.Value
var once_Data_Time_Duration_semigroupSeconds sync.Once
func Get_Data_Time_Duration_semigroupSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupSeconds.Do(func() {
		cache_Data_Time_Duration_semigroupSeconds = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})
}))
	})
	return cache_Data_Time_Duration_semigroupSeconds
}

var cache_Data_Time_Duration_semigroupMinutes gopurs_runtime.Value
var once_Data_Time_Duration_semigroupMinutes sync.Once
func Get_Data_Time_Duration_semigroupMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupMinutes.Do(func() {
		cache_Data_Time_Duration_semigroupMinutes = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})
}))
	})
	return cache_Data_Time_Duration_semigroupMinutes
}

var cache_Data_Time_Duration_semigroupMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_semigroupMilliseconds sync.Once
func Get_Data_Time_Duration_semigroupMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupMilliseconds.Do(func() {
		cache_Data_Time_Duration_semigroupMilliseconds = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})
}))
	})
	return cache_Data_Time_Duration_semigroupMilliseconds
}

var cache_Data_Time_Duration_semigroupHours gopurs_runtime.Value
var once_Data_Time_Duration_semigroupHours sync.Once
func Get_Data_Time_Duration_semigroupHours() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupHours.Do(func() {
		cache_Data_Time_Duration_semigroupHours = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})
}))
	})
	return cache_Data_Time_Duration_semigroupHours
}

var cache_Data_Time_Duration_semigroupDays gopurs_runtime.Value
var once_Data_Time_Duration_semigroupDays sync.Once
func Get_Data_Time_Duration_semigroupDays() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupDays.Do(func() {
		cache_Data_Time_Duration_semigroupDays = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})
}))
	})
	return cache_Data_Time_Duration_semigroupDays
}

var cache_Data_Time_Duration_ordSeconds gopurs_runtime.Value
var once_Data_Time_Duration_ordSeconds sync.Once
func Get_Data_Time_Duration_ordSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_ordSeconds.Do(func() {
		cache_Data_Time_Duration_ordSeconds = Get_Data_Ord_ordNumber()
	})
	return cache_Data_Time_Duration_ordSeconds
}

var cache_Data_Time_Duration_ordMinutes gopurs_runtime.Value
var once_Data_Time_Duration_ordMinutes sync.Once
func Get_Data_Time_Duration_ordMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_ordMinutes.Do(func() {
		cache_Data_Time_Duration_ordMinutes = Get_Data_Ord_ordNumber()
	})
	return cache_Data_Time_Duration_ordMinutes
}

var cache_Data_Time_Duration_ordMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_ordMilliseconds sync.Once
func Get_Data_Time_Duration_ordMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_ordMilliseconds.Do(func() {
		cache_Data_Time_Duration_ordMilliseconds = Get_Data_Ord_ordNumber()
	})
	return cache_Data_Time_Duration_ordMilliseconds
}

var cache_Data_Time_Duration_ordHours gopurs_runtime.Value
var once_Data_Time_Duration_ordHours sync.Once
func Get_Data_Time_Duration_ordHours() gopurs_runtime.Value {
	once_Data_Time_Duration_ordHours.Do(func() {
		cache_Data_Time_Duration_ordHours = Get_Data_Ord_ordNumber()
	})
	return cache_Data_Time_Duration_ordHours
}

var cache_Data_Time_Duration_ordDays gopurs_runtime.Value
var once_Data_Time_Duration_ordDays sync.Once
func Get_Data_Time_Duration_ordDays() gopurs_runtime.Value {
	once_Data_Time_Duration_ordDays.Do(func() {
		cache_Data_Time_Duration_ordDays = Get_Data_Ord_ordNumber()
	})
	return cache_Data_Time_Duration_ordDays
}

var cache_Data_Time_Duration_newtypeSeconds gopurs_runtime.Value
var once_Data_Time_Duration_newtypeSeconds sync.Once
func Get_Data_Time_Duration_newtypeSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeSeconds.Do(func() {
		cache_Data_Time_Duration_newtypeSeconds = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Time_Duration_newtypeSeconds
}

var cache_Data_Time_Duration_newtypeMinutes gopurs_runtime.Value
var once_Data_Time_Duration_newtypeMinutes sync.Once
func Get_Data_Time_Duration_newtypeMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeMinutes.Do(func() {
		cache_Data_Time_Duration_newtypeMinutes = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Time_Duration_newtypeMinutes
}

var cache_Data_Time_Duration_newtypeMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_newtypeMilliseconds sync.Once
func Get_Data_Time_Duration_newtypeMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeMilliseconds.Do(func() {
		cache_Data_Time_Duration_newtypeMilliseconds = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Time_Duration_newtypeMilliseconds
}

var cache_Data_Time_Duration_newtypeHours gopurs_runtime.Value
var once_Data_Time_Duration_newtypeHours sync.Once
func Get_Data_Time_Duration_newtypeHours() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeHours.Do(func() {
		cache_Data_Time_Duration_newtypeHours = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Time_Duration_newtypeHours
}

var cache_Data_Time_Duration_newtypeDays gopurs_runtime.Value
var once_Data_Time_Duration_newtypeDays sync.Once
func Get_Data_Time_Duration_newtypeDays() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeDays.Do(func() {
		cache_Data_Time_Duration_newtypeDays = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Time_Duration_newtypeDays
}

var cache_Data_Time_Duration_monoidSeconds gopurs_runtime.Value
var once_Data_Time_Duration_monoidSeconds sync.Once
func Get_Data_Time_Duration_monoidSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidSeconds.Do(func() {
		cache_Data_Time_Duration_monoidSeconds = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Time_Duration_semigroupSeconds()
}), gopurs_runtime.Float(0.0))
	})
	return cache_Data_Time_Duration_monoidSeconds
}

var cache_Data_Time_Duration_monoidMinutes gopurs_runtime.Value
var once_Data_Time_Duration_monoidMinutes sync.Once
func Get_Data_Time_Duration_monoidMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidMinutes.Do(func() {
		cache_Data_Time_Duration_monoidMinutes = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Time_Duration_semigroupMinutes()
}), gopurs_runtime.Float(0.0))
	})
	return cache_Data_Time_Duration_monoidMinutes
}

var cache_Data_Time_Duration_monoidMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_monoidMilliseconds sync.Once
func Get_Data_Time_Duration_monoidMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidMilliseconds.Do(func() {
		cache_Data_Time_Duration_monoidMilliseconds = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Time_Duration_semigroupMilliseconds()
}), gopurs_runtime.Float(0.0))
	})
	return cache_Data_Time_Duration_monoidMilliseconds
}

var cache_Data_Time_Duration_monoidHours gopurs_runtime.Value
var once_Data_Time_Duration_monoidHours sync.Once
func Get_Data_Time_Duration_monoidHours() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidHours.Do(func() {
		cache_Data_Time_Duration_monoidHours = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Time_Duration_semigroupHours()
}), gopurs_runtime.Float(0.0))
	})
	return cache_Data_Time_Duration_monoidHours
}

var cache_Data_Time_Duration_monoidDays gopurs_runtime.Value
var once_Data_Time_Duration_monoidDays sync.Once
func Get_Data_Time_Duration_monoidDays() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidDays.Do(func() {
		cache_Data_Time_Duration_monoidDays = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Time_Duration_semigroupDays()
}), gopurs_runtime.Float(0.0))
	})
	return cache_Data_Time_Duration_monoidDays
}

var cache_Data_Time_Duration_fromDuration gopurs_runtime.Value
var once_Data_Time_Duration_fromDuration sync.Once
func Get_Data_Time_Duration_fromDuration() gopurs_runtime.Value {
	once_Data_Time_Duration_fromDuration.Do(func() {
		cache_Data_Time_Duration_fromDuration = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_fromDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dict_0_box))
})
	})
	return cache_Data_Time_Duration_fromDuration
}

var cache_Data_Time_Duration_negateDuration gopurs_runtime.Value
var once_Data_Time_Duration_negateDuration sync.Once
func Get_Data_Time_Duration_negateDuration() gopurs_runtime.Value {
	once_Data_Time_Duration_negateDuration.Do(func() {
		cache_Data_Time_Duration_negateDuration = gopurs_runtime.Func(func(dictDuration_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_negateDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box))
})
	})
	return cache_Data_Time_Duration_negateDuration
}

var cache_Data_Time_Duration_eqSeconds gopurs_runtime.Value
var once_Data_Time_Duration_eqSeconds sync.Once
func Get_Data_Time_Duration_eqSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_eqSeconds.Do(func() {
		cache_Data_Time_Duration_eqSeconds = Get_Data_Eq_eqNumber()
	})
	return cache_Data_Time_Duration_eqSeconds
}

var cache_Data_Time_Duration_eqMinutes gopurs_runtime.Value
var once_Data_Time_Duration_eqMinutes sync.Once
func Get_Data_Time_Duration_eqMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_eqMinutes.Do(func() {
		cache_Data_Time_Duration_eqMinutes = Get_Data_Eq_eqNumber()
	})
	return cache_Data_Time_Duration_eqMinutes
}

var cache_Data_Time_Duration_eqMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_eqMilliseconds sync.Once
func Get_Data_Time_Duration_eqMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_eqMilliseconds.Do(func() {
		cache_Data_Time_Duration_eqMilliseconds = Get_Data_Eq_eqNumber()
	})
	return cache_Data_Time_Duration_eqMilliseconds
}

var cache_Data_Time_Duration_eqHours gopurs_runtime.Value
var once_Data_Time_Duration_eqHours sync.Once
func Get_Data_Time_Duration_eqHours() gopurs_runtime.Value {
	once_Data_Time_Duration_eqHours.Do(func() {
		cache_Data_Time_Duration_eqHours = Get_Data_Eq_eqNumber()
	})
	return cache_Data_Time_Duration_eqHours
}

var cache_Data_Time_Duration_eqDays gopurs_runtime.Value
var once_Data_Time_Duration_eqDays sync.Once
func Get_Data_Time_Duration_eqDays() gopurs_runtime.Value {
	once_Data_Time_Duration_eqDays.Do(func() {
		cache_Data_Time_Duration_eqDays = Get_Data_Eq_eqNumber()
	})
	return cache_Data_Time_Duration_eqDays
}

var cache_Data_Time_Duration_durationSeconds gopurs_runtime.Value
var once_Data_Time_Duration_durationSeconds sync.Once
func Get_Data_Time_Duration_durationSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_durationSeconds.Do(func() {
		cache_Data_Time_Duration_durationSeconds = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) * (1000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) / (1000.0))
}))
	})
	return cache_Data_Time_Duration_durationSeconds
}

var cache_Data_Time_Duration_durationMinutes gopurs_runtime.Value
var once_Data_Time_Duration_durationMinutes sync.Once
func Get_Data_Time_Duration_durationMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_durationMinutes.Do(func() {
		cache_Data_Time_Duration_durationMinutes = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) * (60000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) / (60000.0))
}))
	})
	return cache_Data_Time_Duration_durationMinutes
}

var cache_Data_Time_Duration_durationMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_durationMilliseconds sync.Once
func Get_Data_Time_Duration_durationMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_durationMilliseconds.Do(func() {
		cache_Data_Time_Duration_durationMilliseconds = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Time_Duration_durationMilliseconds
}

var cache_Data_Time_Duration_durationHours gopurs_runtime.Value
var once_Data_Time_Duration_durationHours sync.Once
func Get_Data_Time_Duration_durationHours() gopurs_runtime.Value {
	once_Data_Time_Duration_durationHours.Do(func() {
		cache_Data_Time_Duration_durationHours = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) * (3600000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) / (3600000.0))
}))
	})
	return cache_Data_Time_Duration_durationHours
}

var cache_Data_Time_Duration_durationDays gopurs_runtime.Value
var once_Data_Time_Duration_durationDays sync.Once
func Get_Data_Time_Duration_durationDays() gopurs_runtime.Value {
	once_Data_Time_Duration_durationDays.Do(func() {
		cache_Data_Time_Duration_durationDays = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) * (86400000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) / (86400000.0))
}))
	})
	return cache_Data_Time_Duration_durationDays
}

var cache_Data_Time_Duration_convertDuration gopurs_runtime.Value
var once_Data_Time_Duration_convertDuration sync.Once
func Get_Data_Time_Duration_convertDuration() gopurs_runtime.Value {
	once_Data_Time_Duration_convertDuration.Do(func() {
		cache_Data_Time_Duration_convertDuration = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dictDuration1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_convertDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration1_1_box), x_2_box)
})
	})
	return cache_Data_Time_Duration_convertDuration
}

var cache_Data_Time_Duration_fromDuration__4195558286 gopurs_runtime.Value
var once_Data_Time_Duration_fromDuration__4195558286 sync.Once
func Get_Data_Time_Duration_fromDuration__4195558286() gopurs_runtime.Value {
	once_Data_Time_Duration_fromDuration__4195558286.Do(func() {
		cache_Data_Time_Duration_fromDuration__4195558286 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_fromDuration__4195558286(__eta0_0_box)
})
	})
	return cache_Data_Time_Duration_fromDuration__4195558286
}

var cache_Data_Time_Duration_fromDuration__1721614606 gopurs_runtime.Value
var once_Data_Time_Duration_fromDuration__1721614606 sync.Once
func Get_Data_Time_Duration_fromDuration__1721614606() gopurs_runtime.Value {
	once_Data_Time_Duration_fromDuration__1721614606.Do(func() {
		cache_Data_Time_Duration_fromDuration__1721614606 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_fromDuration__1721614606(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dict_0_box))
})
	})
	return cache_Data_Time_Duration_fromDuration__1721614606
}

var cache_Data_Time_Duration_negateDuration__4195558286 gopurs_runtime.Value
var once_Data_Time_Duration_negateDuration__4195558286 sync.Once
func Get_Data_Time_Duration_negateDuration__4195558286() gopurs_runtime.Value {
	once_Data_Time_Duration_negateDuration__4195558286.Do(func() {
		cache_Data_Time_Duration_negateDuration__4195558286 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_negateDuration__4195558286(__eta0_0_box)
})
	})
	return cache_Data_Time_Duration_negateDuration__4195558286
}

var cache_Data_Time_Duration_negateDuration__3870190523 gopurs_runtime.Value
var once_Data_Time_Duration_negateDuration__3870190523 sync.Once
func Get_Data_Time_Duration_negateDuration__3870190523() gopurs_runtime.Value {
	once_Data_Time_Duration_negateDuration__3870190523.Do(func() {
		cache_Data_Time_Duration_negateDuration__3870190523 = gopurs_runtime.Func(func(dictDuration_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_negateDuration__3870190523(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box))
})
	})
	return cache_Data_Time_Duration_negateDuration__3870190523
}

var cache_Data_Time_Duration_toDuration__2440169646 gopurs_runtime.Value
var once_Data_Time_Duration_toDuration__2440169646 sync.Once
func Get_Data_Time_Duration_toDuration__2440169646() gopurs_runtime.Value {
	once_Data_Time_Duration_toDuration__2440169646.Do(func() {
		cache_Data_Time_Duration_toDuration__2440169646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Duration_toDuration__2440169646(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dict_0_box))
})
	})
	return cache_Data_Time_Duration_toDuration__2440169646
}

type Constructor_Data_Time_Duration_Duration struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[487663984] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Time_Duration_Duration)(ptr)
		_ = c
		switch key {
		case "fromDuration": return gopurs_runtime.Box(c.V0)
		case "toDuration": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Time_Duration_Duration: " + key)
		}
	}
}


func Call_Data_Time_Duration_identity(x_0_loop float64) float64 {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Float(x_0).FloatVal()
}

func Call_Data_Time_Duration_Seconds(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Duration_Minutes(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Duration_Milliseconds(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Duration_Hours(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Duration_Duration_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Duration_Days(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Duration_toDuration(dict_0_loop *Constructor_Data_Time_Duration_Duration) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Time_Duration_Duration = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Time_Duration_fromDuration(dict_0_loop *Constructor_Data_Time_Duration_Duration) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Time_Duration_Duration = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Time_Duration_negateDuration(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
}), Get_Data_Ring_numSub())
_ = __local_var_1_2
// TAST (Let): Semiring0_2_3 -> *Constructor_Data_Semiring_Semiring
Semiring0_2_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_2_3
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "sub"), gopurs_runtime.Box(Semiring0_2_3.V3), a_3)
})
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V0), x_2))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V1), gopurs_runtime.Apply(__local_var_1_0, x_2))
})
}

func Call_Data_Time_Duration_convertDuration(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration, dictDuration1_1_loop *Constructor_Data_Time_Duration_Duration, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
var dictDuration1_1 *Constructor_Data_Time_Duration_Duration = dictDuration1_1_loop
_ = dictDuration1_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration1_1.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V0), x_2))
}

func Call_Data_Time_Duration_fromDuration__4195558286(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Float((__eta0_0.FloatVal()) * (86400000.0))
}

func Call_Data_Time_Duration_fromDuration__1721614606(dict_0_loop *Constructor_Data_Time_Duration_Duration) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Time_Duration_Duration = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Time_Duration_negateDuration__4195558286(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
}), Get_Data_Ring_numSub())
_ = __local_var_1_0
return gopurs_runtime.Float(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Time_Duration_durationMilliseconds(), "toDuration"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "sub"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Semiring0"), gopurs_runtime.Value{}), "zero"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Time_Duration_durationMilliseconds(), "fromDuration"), __eta0_0))).FloatVal())
}

func Call_Data_Time_Duration_negateDuration__3870190523(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
}), Get_Data_Ring_numSub())
_ = __local_var_1_2
// TAST (Let): Semiring0_2_3 -> *Constructor_Data_Semiring_Semiring
Semiring0_2_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Semiring0"), gopurs_runtime.Value{}))
_ = Semiring0_2_3
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "sub"), gopurs_runtime.Box(Semiring0_2_3.V3), a_3)
})
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V0), x_2))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V1), gopurs_runtime.Apply(__local_var_1_0, x_2))
})
}

func Call_Data_Time_Duration_toDuration__2440169646(dict_0_loop *Constructor_Data_Time_Duration_Duration) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Time_Duration_Duration = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}


