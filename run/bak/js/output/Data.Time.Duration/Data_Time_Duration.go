package Data_Time_Duration

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var Seconds gopurs_runtime.Value
var once_Seconds sync.Once
func Get_Seconds() gopurs_runtime.Value {
	once_Seconds.Do(func() {
		Seconds = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Seconds
}

var Minutes gopurs_runtime.Value
var once_Minutes sync.Once
func Get_Minutes() gopurs_runtime.Value {
	once_Minutes.Do(func() {
		Minutes = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Minutes
}

var Milliseconds gopurs_runtime.Value
var once_Milliseconds sync.Once
func Get_Milliseconds() gopurs_runtime.Value {
	once_Milliseconds.Do(func() {
		Milliseconds = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Milliseconds
}

var Hours gopurs_runtime.Value
var once_Hours sync.Once
func Get_Hours() gopurs_runtime.Value {
	once_Hours.Do(func() {
		Hours = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Hours
}

var Days gopurs_runtime.Value
var once_Days sync.Once
func Get_Days() gopurs_runtime.Value {
	once_Days.Do(func() {
		Days = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Days
}

var toDuration gopurs_runtime.Value
var once_toDuration sync.Once
func Get_toDuration() gopurs_runtime.Value {
	once_toDuration.Do(func() {
		toDuration = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "toDuration")
})
	})
	return toDuration
}

var showSeconds gopurs_runtime.Value
var once_showSeconds sync.Once
func Get_showSeconds() gopurs_runtime.Value {
	once_showSeconds.Do(func() {
		showSeconds = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Seconds " + gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal + ")")
}))
	})
	return showSeconds
}

var showMinutes gopurs_runtime.Value
var once_showMinutes sync.Once
func Get_showMinutes() gopurs_runtime.Value {
	once_showMinutes.Do(func() {
		showMinutes = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Minutes " + gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal + ")")
}))
	})
	return showMinutes
}

var showMilliseconds gopurs_runtime.Value
var once_showMilliseconds sync.Once
func Get_showMilliseconds() gopurs_runtime.Value {
	once_showMilliseconds.Do(func() {
		showMilliseconds = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Milliseconds " + gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal + ")")
}))
	})
	return showMilliseconds
}

var showHours gopurs_runtime.Value
var once_showHours sync.Once
func Get_showHours() gopurs_runtime.Value {
	once_showHours.Do(func() {
		showHours = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Hours " + gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal + ")")
}))
	})
	return showHours
}

var showDays gopurs_runtime.Value
var once_showDays sync.Once
func Get_showDays() gopurs_runtime.Value {
	once_showDays.Do(func() {
		showDays = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Days " + gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal + ")")
}))
	})
	return showDays
}

var semigroupSeconds gopurs_runtime.Value
var once_semigroupSeconds sync.Once
func Get_semigroupSeconds() gopurs_runtime.Value {
	once_semigroupSeconds.Do(func() {
		semigroupSeconds = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() + v1_1.FloatVal())
}))
	})
	return semigroupSeconds
}

var semigroupMinutes gopurs_runtime.Value
var once_semigroupMinutes sync.Once
func Get_semigroupMinutes() gopurs_runtime.Value {
	once_semigroupMinutes.Do(func() {
		semigroupMinutes = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() + v1_1.FloatVal())
}))
	})
	return semigroupMinutes
}

var semigroupMilliseconds gopurs_runtime.Value
var once_semigroupMilliseconds sync.Once
func Get_semigroupMilliseconds() gopurs_runtime.Value {
	once_semigroupMilliseconds.Do(func() {
		semigroupMilliseconds = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() + v1_1.FloatVal())
}))
	})
	return semigroupMilliseconds
}

var semigroupHours gopurs_runtime.Value
var once_semigroupHours sync.Once
func Get_semigroupHours() gopurs_runtime.Value {
	once_semigroupHours.Do(func() {
		semigroupHours = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() + v1_1.FloatVal())
}))
	})
	return semigroupHours
}

var semigroupDays gopurs_runtime.Value
var once_semigroupDays sync.Once
func Get_semigroupDays() gopurs_runtime.Value {
	once_semigroupDays.Do(func() {
		semigroupDays = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() + v1_1.FloatVal())
}))
	})
	return semigroupDays
}

var ordSeconds gopurs_runtime.Value
var once_ordSeconds sync.Once
func Get_ordSeconds() gopurs_runtime.Value {
	once_ordSeconds.Do(func() {
		ordSeconds = pkg_Data_Ord.Get_ordNumber()
	})
	return ordSeconds
}

var ordMinutes gopurs_runtime.Value
var once_ordMinutes sync.Once
func Get_ordMinutes() gopurs_runtime.Value {
	once_ordMinutes.Do(func() {
		ordMinutes = pkg_Data_Ord.Get_ordNumber()
	})
	return ordMinutes
}

var ordMilliseconds gopurs_runtime.Value
var once_ordMilliseconds sync.Once
func Get_ordMilliseconds() gopurs_runtime.Value {
	once_ordMilliseconds.Do(func() {
		ordMilliseconds = pkg_Data_Ord.Get_ordNumber()
	})
	return ordMilliseconds
}

var ordHours gopurs_runtime.Value
var once_ordHours sync.Once
func Get_ordHours() gopurs_runtime.Value {
	once_ordHours.Do(func() {
		ordHours = pkg_Data_Ord.Get_ordNumber()
	})
	return ordHours
}

var ordDays gopurs_runtime.Value
var once_ordDays sync.Once
func Get_ordDays() gopurs_runtime.Value {
	once_ordDays.Do(func() {
		ordDays = pkg_Data_Ord.Get_ordNumber()
	})
	return ordDays
}

var newtypeSeconds gopurs_runtime.Value
var once_newtypeSeconds sync.Once
func Get_newtypeSeconds() gopurs_runtime.Value {
	once_newtypeSeconds.Do(func() {
		newtypeSeconds = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeSeconds
}

var newtypeMinutes gopurs_runtime.Value
var once_newtypeMinutes sync.Once
func Get_newtypeMinutes() gopurs_runtime.Value {
	once_newtypeMinutes.Do(func() {
		newtypeMinutes = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeMinutes
}

var newtypeMilliseconds gopurs_runtime.Value
var once_newtypeMilliseconds sync.Once
func Get_newtypeMilliseconds() gopurs_runtime.Value {
	once_newtypeMilliseconds.Do(func() {
		newtypeMilliseconds = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeMilliseconds
}

var newtypeHours gopurs_runtime.Value
var once_newtypeHours sync.Once
func Get_newtypeHours() gopurs_runtime.Value {
	once_newtypeHours.Do(func() {
		newtypeHours = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeHours
}

var newtypeDays gopurs_runtime.Value
var once_newtypeDays sync.Once
func Get_newtypeDays() gopurs_runtime.Value {
	once_newtypeDays.Do(func() {
		newtypeDays = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeDays
}

var monoidSeconds gopurs_runtime.Value
var once_monoidSeconds sync.Once
func Get_monoidSeconds() gopurs_runtime.Value {
	once_monoidSeconds.Do(func() {
		monoidSeconds = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupSeconds()
}))
	})
	return monoidSeconds
}

var monoidMinutes gopurs_runtime.Value
var once_monoidMinutes sync.Once
func Get_monoidMinutes() gopurs_runtime.Value {
	once_monoidMinutes.Do(func() {
		monoidMinutes = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupMinutes()
}))
	})
	return monoidMinutes
}

var monoidMilliseconds gopurs_runtime.Value
var once_monoidMilliseconds sync.Once
func Get_monoidMilliseconds() gopurs_runtime.Value {
	once_monoidMilliseconds.Do(func() {
		monoidMilliseconds = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupMilliseconds()
}))
	})
	return monoidMilliseconds
}

var monoidHours gopurs_runtime.Value
var once_monoidHours sync.Once
func Get_monoidHours() gopurs_runtime.Value {
	once_monoidHours.Do(func() {
		monoidHours = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupHours()
}))
	})
	return monoidHours
}

var monoidDays gopurs_runtime.Value
var once_monoidDays sync.Once
func Get_monoidDays() gopurs_runtime.Value {
	once_monoidDays.Do(func() {
		monoidDays = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupDays()
}))
	})
	return monoidDays
}

var fromDuration gopurs_runtime.Value
var once_fromDuration sync.Once
func Get_fromDuration() gopurs_runtime.Value {
	once_fromDuration.Do(func() {
		fromDuration = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "fromDuration")
})
	})
	return fromDuration
}

var negateDuration gopurs_runtime.Value
var once_negateDuration sync.Once
func Get_negateDuration() gopurs_runtime.Value {
	once_negateDuration.Do(func() {
		negateDuration = gopurs_runtime.Func2(func(dictDuration_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.FloatNeg(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "fromDuration"), x_1)))
})
	})
	return negateDuration
}

var eqSeconds gopurs_runtime.Value
var once_eqSeconds sync.Once
func Get_eqSeconds() gopurs_runtime.Value {
	once_eqSeconds.Do(func() {
		eqSeconds = pkg_Data_Eq.Get_eqNumber()
	})
	return eqSeconds
}

var eqMinutes gopurs_runtime.Value
var once_eqMinutes sync.Once
func Get_eqMinutes() gopurs_runtime.Value {
	once_eqMinutes.Do(func() {
		eqMinutes = pkg_Data_Eq.Get_eqNumber()
	})
	return eqMinutes
}

var eqMilliseconds gopurs_runtime.Value
var once_eqMilliseconds sync.Once
func Get_eqMilliseconds() gopurs_runtime.Value {
	once_eqMilliseconds.Do(func() {
		eqMilliseconds = pkg_Data_Eq.Get_eqNumber()
	})
	return eqMilliseconds
}

var eqHours gopurs_runtime.Value
var once_eqHours sync.Once
func Get_eqHours() gopurs_runtime.Value {
	once_eqHours.Do(func() {
		eqHours = pkg_Data_Eq.Get_eqNumber()
	})
	return eqHours
}

var eqDays gopurs_runtime.Value
var once_eqDays sync.Once
func Get_eqDays() gopurs_runtime.Value {
	once_eqDays.Do(func() {
		eqDays = pkg_Data_Eq.Get_eqNumber()
	})
	return eqDays
}

var durationSeconds gopurs_runtime.Value
var once_durationSeconds sync.Once
func Get_durationSeconds() gopurs_runtime.Value {
	once_durationSeconds.Do(func() {
		durationSeconds = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() * 1000.0)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() / 1000.0)
}))
	})
	return durationSeconds
}

var durationMinutes gopurs_runtime.Value
var once_durationMinutes sync.Once
func Get_durationMinutes() gopurs_runtime.Value {
	once_durationMinutes.Do(func() {
		durationMinutes = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() * 60000.0)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() / 60000.0)
}))
	})
	return durationMinutes
}

var durationMilliseconds gopurs_runtime.Value
var once_durationMilliseconds sync.Once
func Get_durationMilliseconds() gopurs_runtime.Value {
	once_durationMilliseconds.Do(func() {
		durationMilliseconds = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
	})
	return durationMilliseconds
}

var durationHours gopurs_runtime.Value
var once_durationHours sync.Once
func Get_durationHours() gopurs_runtime.Value {
	once_durationHours.Do(func() {
		durationHours = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() * 3600000.0)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() / 3600000.0)
}))
	})
	return durationHours
}

var durationDays gopurs_runtime.Value
var once_durationDays sync.Once
func Get_durationDays() gopurs_runtime.Value {
	once_durationDays.Do(func() {
		durationDays = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() * 86400000.0)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(v_0.FloatVal() / 86400000.0)
}))
	})
	return durationDays
}

var convertDuration gopurs_runtime.Value
var once_convertDuration sync.Once
func Get_convertDuration() gopurs_runtime.Value {
	once_convertDuration.Do(func() {
		convertDuration = gopurs_runtime.Func3(func(dictDuration_0 gopurs_runtime.Value, dictDuration1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration1_1, "toDuration"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "fromDuration"), x_2))
})
	})
	return convertDuration
}




