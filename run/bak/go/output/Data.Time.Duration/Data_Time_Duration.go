package Data_Time_Duration

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var cache_Seconds gopurs_runtime.Value
var once_Seconds sync.Once
func Get_Seconds() gopurs_runtime.Value {
	once_Seconds.Do(func() {
		cache_Seconds = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Seconds
}

var cache_Minutes gopurs_runtime.Value
var once_Minutes sync.Once
func Get_Minutes() gopurs_runtime.Value {
	once_Minutes.Do(func() {
		cache_Minutes = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Minutes
}

var cache_Milliseconds gopurs_runtime.Value
var once_Milliseconds sync.Once
func Get_Milliseconds() gopurs_runtime.Value {
	once_Milliseconds.Do(func() {
		cache_Milliseconds = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Milliseconds
}

var cache_Hours gopurs_runtime.Value
var once_Hours sync.Once
func Get_Hours() gopurs_runtime.Value {
	once_Hours.Do(func() {
		cache_Hours = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Hours
}

var cache_Days gopurs_runtime.Value
var once_Days sync.Once
func Get_Days() gopurs_runtime.Value {
	once_Days.Do(func() {
		cache_Days = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Days
}

var cache_toDuration gopurs_runtime.Value
var once_toDuration sync.Once
func Get_toDuration() gopurs_runtime.Value {
	once_toDuration.Do(func() {
		cache_toDuration = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "toDuration")
}()
})
	})
	return cache_toDuration
}

var cache_showSeconds gopurs_runtime.Value
var once_showSeconds sync.Once
func Get_showSeconds() gopurs_runtime.Value {
	once_showSeconds.Do(func() {
		cache_showSeconds = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Seconds ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showSeconds
}

var cache_showMinutes gopurs_runtime.Value
var once_showMinutes sync.Once
func Get_showMinutes() gopurs_runtime.Value {
	once_showMinutes.Do(func() {
		cache_showMinutes = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Minutes ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showMinutes
}

var cache_showMilliseconds gopurs_runtime.Value
var once_showMilliseconds sync.Once
func Get_showMilliseconds() gopurs_runtime.Value {
	once_showMilliseconds.Do(func() {
		cache_showMilliseconds = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Milliseconds ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showMilliseconds
}

var cache_showHours gopurs_runtime.Value
var once_showHours sync.Once
func Get_showHours() gopurs_runtime.Value {
	once_showHours.Do(func() {
		cache_showHours = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Hours ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showHours
}

var cache_showDays gopurs_runtime.Value
var once_showDays sync.Once
func Get_showDays() gopurs_runtime.Value {
	once_showDays.Do(func() {
		cache_showDays = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Days ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showDays
}

var cache_semigroupSeconds gopurs_runtime.Value
var once_semigroupSeconds sync.Once
func Get_semigroupSeconds() gopurs_runtime.Value {
	once_semigroupSeconds.Do(func() {
		cache_semigroupSeconds = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
}))
	})
	return cache_semigroupSeconds
}

var cache_semigroupMinutes gopurs_runtime.Value
var once_semigroupMinutes sync.Once
func Get_semigroupMinutes() gopurs_runtime.Value {
	once_semigroupMinutes.Do(func() {
		cache_semigroupMinutes = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
}))
	})
	return cache_semigroupMinutes
}

var cache_semigroupMilliseconds gopurs_runtime.Value
var once_semigroupMilliseconds sync.Once
func Get_semigroupMilliseconds() gopurs_runtime.Value {
	once_semigroupMilliseconds.Do(func() {
		cache_semigroupMilliseconds = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
}))
	})
	return cache_semigroupMilliseconds
}

var cache_semigroupHours gopurs_runtime.Value
var once_semigroupHours sync.Once
func Get_semigroupHours() gopurs_runtime.Value {
	once_semigroupHours.Do(func() {
		cache_semigroupHours = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
}))
	})
	return cache_semigroupHours
}

var cache_semigroupDays gopurs_runtime.Value
var once_semigroupDays sync.Once
func Get_semigroupDays() gopurs_runtime.Value {
	once_semigroupDays.Do(func() {
		cache_semigroupDays = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
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
		cache_monoidSeconds = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupSeconds()
}))
	})
	return cache_monoidSeconds
}

var cache_monoidMinutes gopurs_runtime.Value
var once_monoidMinutes sync.Once
func Get_monoidMinutes() gopurs_runtime.Value {
	once_monoidMinutes.Do(func() {
		cache_monoidMinutes = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupMinutes()
}))
	})
	return cache_monoidMinutes
}

var cache_monoidMilliseconds gopurs_runtime.Value
var once_monoidMilliseconds sync.Once
func Get_monoidMilliseconds() gopurs_runtime.Value {
	once_monoidMilliseconds.Do(func() {
		cache_monoidMilliseconds = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupMilliseconds()
}))
	})
	return cache_monoidMilliseconds
}

var cache_monoidHours gopurs_runtime.Value
var once_monoidHours sync.Once
func Get_monoidHours() gopurs_runtime.Value {
	once_monoidHours.Do(func() {
		cache_monoidHours = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupHours()
}))
	})
	return cache_monoidHours
}

var cache_monoidDays gopurs_runtime.Value
var once_monoidDays sync.Once
func Get_monoidDays() gopurs_runtime.Value {
	once_monoidDays.Do(func() {
		cache_monoidDays = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupDays()
}))
	})
	return cache_monoidDays
}

var cache_fromDuration gopurs_runtime.Value
var once_fromDuration sync.Once
func Get_fromDuration() gopurs_runtime.Value {
	once_fromDuration.Do(func() {
		cache_fromDuration = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "fromDuration")
}()
})
	})
	return cache_fromDuration
}

var cache_negateDuration gopurs_runtime.Value
var once_negateDuration sync.Once
func Get_negateDuration() gopurs_runtime.Value {
	once_negateDuration.Do(func() {
		cache_negateDuration = gopurs_runtime.Func2(func(dictDuration_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negateDuration(dictDuration_0_box, x_1_box)
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
return gopurs_runtime.Float((v_0.FloatVal()) * (1000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) / (1000.0))
}))
	})
	return cache_durationSeconds
}

var cache_durationMinutes gopurs_runtime.Value
var once_durationMinutes sync.Once
func Get_durationMinutes() gopurs_runtime.Value {
	once_durationMinutes.Do(func() {
		cache_durationMinutes = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) * (60000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) / (60000.0))
}))
	})
	return cache_durationMinutes
}

var cache_durationMilliseconds gopurs_runtime.Value
var once_durationMilliseconds sync.Once
func Get_durationMilliseconds() gopurs_runtime.Value {
	once_durationMilliseconds.Do(func() {
		cache_durationMilliseconds = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
	})
	return cache_durationMilliseconds
}

var cache_durationHours gopurs_runtime.Value
var once_durationHours sync.Once
func Get_durationHours() gopurs_runtime.Value {
	once_durationHours.Do(func() {
		cache_durationHours = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) * (3600000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) / (3600000.0))
}))
	})
	return cache_durationHours
}

var cache_durationDays gopurs_runtime.Value
var once_durationDays sync.Once
func Get_durationDays() gopurs_runtime.Value {
	once_durationDays.Do(func() {
		cache_durationDays = gopurs_runtime.RecordDict2("fromDuration", "toDuration", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) * (86400000.0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) / (86400000.0))
}))
	})
	return cache_durationDays
}

var cache_convertDuration gopurs_runtime.Value
var once_convertDuration sync.Once
func Get_convertDuration() gopurs_runtime.Value {
	once_convertDuration.Do(func() {
		cache_convertDuration = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dictDuration1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_convertDuration(dictDuration_0_box, dictDuration1_1_box, x_2_box)
})
	})
	return cache_convertDuration
}

func Call_negateDuration(dictDuration_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 gopurs_runtime.Value = dictDuration_0_loop
_ = dictDuration_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.FloatNeg(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "fromDuration"), x_1)))
}

func Call_convertDuration(dictDuration_0_loop gopurs_runtime.Value, dictDuration1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 gopurs_runtime.Value = dictDuration_0_loop
_ = dictDuration_0
var dictDuration1_1 gopurs_runtime.Value = dictDuration1_1_loop
_ = dictDuration1_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration1_1, "toDuration"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "fromDuration"), x_2))
}


