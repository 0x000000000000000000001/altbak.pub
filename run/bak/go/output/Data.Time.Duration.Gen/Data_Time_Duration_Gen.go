package Data_Time_Duration_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
)

var cache_genSeconds gopurs_runtime.Value
var once_genSeconds sync.Once
func Get_genSeconds() gopurs_runtime.Value {
	once_genSeconds.Do(func() {
		cache_genSeconds = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genSeconds(dictMonadGen_0_box)
})
	})
	return cache_genSeconds
}

var cache_genMinutes gopurs_runtime.Value
var once_genMinutes sync.Once
func Get_genMinutes() gopurs_runtime.Value {
	once_genMinutes.Do(func() {
		cache_genMinutes = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMinutes(dictMonadGen_0_box)
})
	})
	return cache_genMinutes
}

var cache_genMilliseconds gopurs_runtime.Value
var once_genMilliseconds sync.Once
func Get_genMilliseconds() gopurs_runtime.Value {
	once_genMilliseconds.Do(func() {
		cache_genMilliseconds = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMilliseconds(dictMonadGen_0_box)
})
	})
	return cache_genMilliseconds
}

var cache_genHours gopurs_runtime.Value
var once_genHours sync.Once
func Get_genHours() gopurs_runtime.Value {
	once_genHours.Do(func() {
		cache_genHours = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genHours(dictMonadGen_0_box)
})
	})
	return cache_genHours
}

var cache_genDays gopurs_runtime.Value
var once_genDays sync.Once
func Get_genDays() gopurs_runtime.Value {
	once_genDays.Do(func() {
		cache_genDays = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDays(dictMonadGen_0_box)
})
	})
	return cache_genDays
}

func Call_genSeconds(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Time_Duration.Get_Seconds(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(600.0)))
}

func Call_genMinutes(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Time_Duration.Get_Minutes(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(600.0)))
}

func Call_genMilliseconds(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Time_Duration.Get_Milliseconds(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(600000.0)))
}

func Call_genHours(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Time_Duration.Get_Hours(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(240.0)))
}

func Call_genDays(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Time_Duration.Get_Days(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(42.0)))
}


