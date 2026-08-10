package Effect_Class_Console

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Control_Bind "gopurs/output/Control.Bind"
)

var cache_warnShow gopurs_runtime.Value
var once_warnShow sync.Once
func Get_warnShow() gopurs_runtime.Value {
	once_warnShow.Do(func() {
		cache_warnShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_warnShow(dictMonadEffect_0_box, dictShow_1_box, x_2_box)
})
	})
	return cache_warnShow
}

var cache_warn gopurs_runtime.Value
var once_warn sync.Once
func Get_warn() gopurs_runtime.Value {
	once_warn.Do(func() {
		cache_warn = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_warn(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_warn
}

var cache_timeLog gopurs_runtime.Value
var once_timeLog sync.Once
func Get_timeLog() gopurs_runtime.Value {
	once_timeLog.Do(func() {
		cache_timeLog = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_timeLog(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_timeLog
}

var cache_timeEnd gopurs_runtime.Value
var once_timeEnd sync.Once
func Get_timeEnd() gopurs_runtime.Value {
	once_timeEnd.Do(func() {
		cache_timeEnd = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_timeEnd(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_timeEnd
}

var cache_time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		cache_time = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_time(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_time
}

var cache_logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		cache_logShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow(dictMonadEffect_0_box, dictShow_1_box, x_2_box)
})
	})
	return cache_logShow
}

var cache_log gopurs_runtime.Value
var once_log sync.Once
func Get_log() gopurs_runtime.Value {
	once_log.Do(func() {
		cache_log = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_log(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_log
}

var cache_infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		cache_infoShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_infoShow(dictMonadEffect_0_box, dictShow_1_box, x_2_box)
})
	})
	return cache_infoShow
}

var cache_info gopurs_runtime.Value
var once_info sync.Once
func Get_info() gopurs_runtime.Value {
	once_info.Do(func() {
		cache_info = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_info(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_info
}

var cache_groupEnd gopurs_runtime.Value
var once_groupEnd sync.Once
func Get_groupEnd() gopurs_runtime.Value {
	once_groupEnd.Do(func() {
		cache_groupEnd = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupEnd(dictMonadEffect_0_box)
})
	})
	return cache_groupEnd
}

var cache_groupCollapsed gopurs_runtime.Value
var once_groupCollapsed sync.Once
func Get_groupCollapsed() gopurs_runtime.Value {
	once_groupCollapsed.Do(func() {
		cache_groupCollapsed = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupCollapsed(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_groupCollapsed
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_group
}

var cache_group__gopurs_runtime_Value_227222742 gopurs_runtime.Value
var once_group__gopurs_runtime_Value_227222742 sync.Once
func Get_group__gopurs_runtime_Value_227222742() gopurs_runtime.Value {
	once_group__gopurs_runtime_Value_227222742.Do(func() {
		cache_group__gopurs_runtime_Value_227222742 = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group__gopurs_runtime_Value_227222742(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_group__gopurs_runtime_Value_227222742
}

var cache_grouped gopurs_runtime.Value
var once_grouped sync.Once
func Get_grouped() gopurs_runtime.Value {
	once_grouped.Do(func() {
		cache_grouped = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_grouped(dictMonadEffect_0_box)
})
	})
	return cache_grouped
}

var cache_errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		cache_errorShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_errorShow(dictMonadEffect_0_box, dictShow_1_box, x_2_box)
})
	})
	return cache_errorShow
}

var cache_error gopurs_runtime.Value
var once_error sync.Once
func Get_error() gopurs_runtime.Value {
	once_error.Do(func() {
		cache_error = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_error(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_error
}

var cache_debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		cache_debugShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_debugShow(dictMonadEffect_0_box, dictShow_1_box, x_2_box)
})
	})
	return cache_debugShow
}

var cache_debug gopurs_runtime.Value
var once_debug sync.Once
func Get_debug() gopurs_runtime.Value {
	once_debug.Do(func() {
		cache_debug = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_debug(dictMonadEffect_0_box, x_1_box)
})
	})
	return cache_debug
}

var cache_clear gopurs_runtime.Value
var once_clear sync.Once
func Get_clear() gopurs_runtime.Value {
	once_clear.Do(func() {
		cache_clear = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_clear(dictMonadEffect_0_box)
})
	})
	return cache_clear
}

func Call_warnShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_warn(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), x_2)))
}

func Call_warn(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_warn(), x_1))
}

func Call_timeLog(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_timeLog(), x_1))
}

func Call_timeEnd(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_timeEnd(), x_1))
}

func Call_time(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_time(), x_1))
}

func Call_logShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), x_2)))
}

func Call_log(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), x_1))
}

func Call_infoShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_info(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), x_2)))
}

func Call_info(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_info(), x_1))
}

func Call_groupEnd(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), pkg_Effect_Console.Get_groupEnd())
}

func Call_groupCollapsed(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_groupCollapsed(), x_1))
}

func Call_group(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_group(), x_1))
}

func Call_group__gopurs_runtime_Value_227222742(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_group(), x_1))
}

func Call_grouped(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
discard1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Bind1_2_1)
_ = discard1_3_2
groupEnd1_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), pkg_Effect_Console.Get_groupEnd())
_ = groupEnd1_4_3
return gopurs_runtime.Func(func(name_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(inner_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(discard1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_group(), name_5)), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), inner_6, gopurs_runtime.Func(func(result_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(discard1_3_2, groupEnd1_4_3, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), result_8)
}))
}))
}))
})
})
}

func Call_errorShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), x_2)))
}

func Call_error(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), x_1))
}

func Call_debugShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_debug(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), x_2)))
}

func Call_debug(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_debug(), x_1))
}

func Call_clear(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), pkg_Effect_Console.Get_clear())
}


