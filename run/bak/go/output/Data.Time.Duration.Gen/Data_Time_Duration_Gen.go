package Data_Time_Duration_Gen

import (
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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

var cache_chooseFloat__1964853975 gopurs_runtime.Value
var once_chooseFloat__1964853975 sync.Once
func Get_chooseFloat__1964853975() gopurs_runtime.Value {
	once_chooseFloat__1964853975.Do(func() {
		cache_chooseFloat__1964853975 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseFloat__1964853975(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseFloat__1964853975
}

var cache_map__2224957140 gopurs_runtime.Value
var once_map__2224957140 sync.Once
func Get_map__2224957140() gopurs_runtime.Value {
	once_map__2224957140.Do(func() {
		cache_map__2224957140 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2224957140(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2224957140
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
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

func Call_chooseFloat__1964853975(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_map__2224957140(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


