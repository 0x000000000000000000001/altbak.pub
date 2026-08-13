package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Class_Console_warnShow gopurs_runtime.Value
var once_Effect_Class_Console_warnShow sync.Once
func Get_Effect_Class_Console_warnShow() gopurs_runtime.Value {
	once_Effect_Class_Console_warnShow.Do(func() {
		cache_Effect_Class_Console_warnShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_warnShow(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_1_box), x_2_box)
})
	})
	return cache_Effect_Class_Console_warnShow
}

var cache_Effect_Class_Console_warn gopurs_runtime.Value
var once_Effect_Class_Console_warn sync.Once
func Get_Effect_Class_Console_warn() gopurs_runtime.Value {
	once_Effect_Class_Console_warn.Do(func() {
		cache_Effect_Class_Console_warn = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_warn(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_warn
}

var cache_Effect_Class_Console_timeLog gopurs_runtime.Value
var once_Effect_Class_Console_timeLog sync.Once
func Get_Effect_Class_Console_timeLog() gopurs_runtime.Value {
	once_Effect_Class_Console_timeLog.Do(func() {
		cache_Effect_Class_Console_timeLog = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_timeLog(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_timeLog
}

var cache_Effect_Class_Console_timeEnd gopurs_runtime.Value
var once_Effect_Class_Console_timeEnd sync.Once
func Get_Effect_Class_Console_timeEnd() gopurs_runtime.Value {
	once_Effect_Class_Console_timeEnd.Do(func() {
		cache_Effect_Class_Console_timeEnd = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_timeEnd(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_timeEnd
}

var cache_Effect_Class_Console_time gopurs_runtime.Value
var once_Effect_Class_Console_time sync.Once
func Get_Effect_Class_Console_time() gopurs_runtime.Value {
	once_Effect_Class_Console_time.Do(func() {
		cache_Effect_Class_Console_time = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_time(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_time
}

var cache_Effect_Class_Console_logShow gopurs_runtime.Value
var once_Effect_Class_Console_logShow sync.Once
func Get_Effect_Class_Console_logShow() gopurs_runtime.Value {
	once_Effect_Class_Console_logShow.Do(func() {
		cache_Effect_Class_Console_logShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_logShow(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_1_box), x_2_box)
})
	})
	return cache_Effect_Class_Console_logShow
}

var cache_Effect_Class_Console_log gopurs_runtime.Value
var once_Effect_Class_Console_log sync.Once
func Get_Effect_Class_Console_log() gopurs_runtime.Value {
	once_Effect_Class_Console_log.Do(func() {
		cache_Effect_Class_Console_log = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_log(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_log
}

var cache_Effect_Class_Console_infoShow gopurs_runtime.Value
var once_Effect_Class_Console_infoShow sync.Once
func Get_Effect_Class_Console_infoShow() gopurs_runtime.Value {
	once_Effect_Class_Console_infoShow.Do(func() {
		cache_Effect_Class_Console_infoShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_infoShow(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_1_box), x_2_box)
})
	})
	return cache_Effect_Class_Console_infoShow
}

var cache_Effect_Class_Console_info gopurs_runtime.Value
var once_Effect_Class_Console_info sync.Once
func Get_Effect_Class_Console_info() gopurs_runtime.Value {
	once_Effect_Class_Console_info.Do(func() {
		cache_Effect_Class_Console_info = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_info(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_info
}

var cache_Effect_Class_Console_groupEnd gopurs_runtime.Value
var once_Effect_Class_Console_groupEnd sync.Once
func Get_Effect_Class_Console_groupEnd() gopurs_runtime.Value {
	once_Effect_Class_Console_groupEnd.Do(func() {
		cache_Effect_Class_Console_groupEnd = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_groupEnd(dictMonadEffect_0_box)
})
	})
	return cache_Effect_Class_Console_groupEnd
}

var cache_Effect_Class_Console_groupCollapsed gopurs_runtime.Value
var once_Effect_Class_Console_groupCollapsed sync.Once
func Get_Effect_Class_Console_groupCollapsed() gopurs_runtime.Value {
	once_Effect_Class_Console_groupCollapsed.Do(func() {
		cache_Effect_Class_Console_groupCollapsed = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_groupCollapsed(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_groupCollapsed
}

var cache_Effect_Class_Console_group gopurs_runtime.Value
var once_Effect_Class_Console_group sync.Once
func Get_Effect_Class_Console_group() gopurs_runtime.Value {
	once_Effect_Class_Console_group.Do(func() {
		cache_Effect_Class_Console_group = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_group(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_group
}

var cache_Effect_Class_Console_grouped gopurs_runtime.Value
var once_Effect_Class_Console_grouped sync.Once
func Get_Effect_Class_Console_grouped() gopurs_runtime.Value {
	once_Effect_Class_Console_grouped.Do(func() {
		cache_Effect_Class_Console_grouped = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_grouped(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box))
})
	})
	return cache_Effect_Class_Console_grouped
}

var cache_Effect_Class_Console_errorShow gopurs_runtime.Value
var once_Effect_Class_Console_errorShow sync.Once
func Get_Effect_Class_Console_errorShow() gopurs_runtime.Value {
	once_Effect_Class_Console_errorShow.Do(func() {
		cache_Effect_Class_Console_errorShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_errorShow(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_1_box), x_2_box)
})
	})
	return cache_Effect_Class_Console_errorShow
}

var cache_Effect_Class_Console_error gopurs_runtime.Value
var once_Effect_Class_Console_error sync.Once
func Get_Effect_Class_Console_error() gopurs_runtime.Value {
	once_Effect_Class_Console_error.Do(func() {
		cache_Effect_Class_Console_error = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_error(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_error
}

var cache_Effect_Class_Console_debugShow gopurs_runtime.Value
var once_Effect_Class_Console_debugShow sync.Once
func Get_Effect_Class_Console_debugShow() gopurs_runtime.Value {
	once_Effect_Class_Console_debugShow.Do(func() {
		cache_Effect_Class_Console_debugShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_debugShow(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_1_box), x_2_box)
})
	})
	return cache_Effect_Class_Console_debugShow
}

var cache_Effect_Class_Console_debug gopurs_runtime.Value
var once_Effect_Class_Console_debug sync.Once
func Get_Effect_Class_Console_debug() gopurs_runtime.Value {
	once_Effect_Class_Console_debug.Do(func() {
		cache_Effect_Class_Console_debug = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_debug(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_debug
}

var cache_Effect_Class_Console_clear gopurs_runtime.Value
var once_Effect_Class_Console_clear sync.Once
func Get_Effect_Class_Console_clear() gopurs_runtime.Value {
	once_Effect_Class_Console_clear.Do(func() {
		cache_Effect_Class_Console_clear = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_clear(dictMonadEffect_0_box)
})
	})
	return cache_Effect_Class_Console_clear
}

var cache_Effect_Class_Console_group__227222742 gopurs_runtime.Value
var once_Effect_Class_Console_group__227222742 sync.Once
func Get_Effect_Class_Console_group__227222742() gopurs_runtime.Value {
	once_Effect_Class_Console_group__227222742.Do(func() {
		cache_Effect_Class_Console_group__227222742 = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_Console_group__227222742(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_Effect_Class_Console_group__227222742
}

func Call_Effect_Class_Console_warnShow(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, dictShow_1_loop *Constructor_Data_Show_Show, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *Constructor_Data_Show_Show = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_warn(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_1.V0), x_2).StrVal())))
}

func Call_Effect_Class_Console_warn(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_warn(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_timeLog(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_timeLog(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_timeEnd(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_timeEnd(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_time(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_time(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_logShow(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, dictShow_1_loop *Constructor_Data_Show_Show, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *Constructor_Data_Show_Show = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_1.V0), x_2).StrVal())))
}

func Call_Effect_Class_Console_log(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_infoShow(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, dictShow_1_loop *Constructor_Data_Show_Show, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *Constructor_Data_Show_Show = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_info(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_1.V0), x_2).StrVal())))
}

func Call_Effect_Class_Console_info(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_info(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_groupEnd(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), Get_Effect_Console_groupEnd())
}

func Call_Effect_Class_Console_groupCollapsed(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_groupCollapsed(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_group(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_group(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_grouped(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): groupEnd1_3_2 -> gopurs_runtime.Value
groupEnd1_3_2 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), Get_Effect_Console_groupEnd())
_ = groupEnd1_3_2
// TAST (Let): Applicative0_4_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.Func(func(name_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(inner_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_group(), gopurs_runtime.Str(name_5.StrVal()))), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), inner_6, gopurs_runtime.Func(func(result_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), groupEnd1_3_2, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_3.V1), result_8)
}))
}))
}))
})
})
}

func Call_Effect_Class_Console_errorShow(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, dictShow_1_loop *Constructor_Data_Show_Show, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *Constructor_Data_Show_Show = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_1.V0), x_2).StrVal())))
}

func Call_Effect_Class_Console_error(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_debugShow(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, dictShow_1_loop *Constructor_Data_Show_Show, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *Constructor_Data_Show_Show = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_debug(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_1.V0), x_2).StrVal())))
}

func Call_Effect_Class_Console_debug(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_debug(), gopurs_runtime.Str(x_1)))
}

func Call_Effect_Class_Console_clear(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), Get_Effect_Console_clear())
}

func Call_Effect_Class_Console_group__227222742(dictMonadEffect_0_loop *Constructor_Effect_Class_MonadEffect, x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *Constructor_Effect_Class_MonadEffect = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadEffect_0.V1), gopurs_runtime.Apply(Get_Effect_Console_group(), gopurs_runtime.Str(x_1)))
}


