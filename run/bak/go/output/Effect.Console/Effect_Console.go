package Effect_Console

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Effect "gopurs/output/Effect"
)

var cache_discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		cache_discard = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard
}

var cache_warnShow gopurs_runtime.Value
var once_warnShow sync.Once
func Get_warnShow() gopurs_runtime.Value {
	once_warnShow.Do(func() {
		cache_warnShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Call_warnShow(dictShow_0_box, gopurs_runtime.UnboxAny(a_1_box))()
})
})
	})
	return cache_warnShow
}

var cache_logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		cache_logShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Call_logShow(dictShow_0_box, gopurs_runtime.UnboxAny(a_1_box))()
})
})
	})
	return cache_logShow
}

var cache_logShow__func_gopurs_runtime_Value__interface____func___gopurs_runtime_Value_2566372785 gopurs_runtime.Value
var once_logShow__func_gopurs_runtime_Value__interface____func___gopurs_runtime_Value_2566372785 sync.Once
func Get_logShow__func_gopurs_runtime_Value__interface____func___gopurs_runtime_Value_2566372785() gopurs_runtime.Value {
	once_logShow__func_gopurs_runtime_Value__interface____func___gopurs_runtime_Value_2566372785.Do(func() {
		cache_logShow__func_gopurs_runtime_Value__interface____func___gopurs_runtime_Value_2566372785 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Call_logShow__func_gopurs_runtime_Value__interface____func___gopurs_runtime_Value_2566372785(dictShow_0_box, gopurs_runtime.UnboxAny(a_1_box))()
})
})
	})
	return cache_logShow__func_gopurs_runtime_Value__interface____func___gopurs_runtime_Value_2566372785
}

var cache_infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		cache_infoShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Call_infoShow(dictShow_0_box, gopurs_runtime.UnboxAny(a_1_box))()
})
})
	})
	return cache_infoShow
}

var cache_grouped gopurs_runtime.Value
var once_grouped sync.Once
func Get_grouped() gopurs_runtime.Value {
	once_grouped.Do(func() {
		cache_grouped = gopurs_runtime.Func2(func(name_0_box gopurs_runtime.Value, inner_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(Call_grouped(name_0_box.StrVal(), func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(inner_1_box, nil))
})())
})
})
	})
	return cache_grouped
}

var cache_errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		cache_errorShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Call_errorShow(dictShow_0_box, gopurs_runtime.UnboxAny(a_1_box))()
})
})
	})
	return cache_errorShow
}

var cache_debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		cache_debugShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Call_debugShow(dictShow_0_box, gopurs_runtime.UnboxAny(a_1_box))()
})
})
	})
	return cache_debugShow
}

var cache_clear gopurs_runtime.Value
var once_clear sync.Once
func Get_clear() gopurs_runtime.Value {
	once_clear.Do(func() {
		cache_clear = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Clear()
})
	})
	return cache_clear
}

var cache_debug gopurs_runtime.Value
var once_debug sync.Once
func Get_debug() gopurs_runtime.Value {
	once_debug.Do(func() {
		cache_debug = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Debug(arg0.StrVal())()
})
})
	})
	return cache_debug
}

var cache_error gopurs_runtime.Value
var once_error sync.Once
func Get_error() gopurs_runtime.Value {
	once_error.Do(func() {
		cache_error = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Error(arg0.StrVal())()
})
})
	})
	return cache_error
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Group(arg0.StrVal())()
})
})
	})
	return cache_group
}

var cache_groupCollapsed gopurs_runtime.Value
var once_groupCollapsed sync.Once
func Get_groupCollapsed() gopurs_runtime.Value {
	once_groupCollapsed.Do(func() {
		cache_groupCollapsed = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return GroupCollapsed(arg0.StrVal())()
})
})
	})
	return cache_groupCollapsed
}

var cache_groupEnd gopurs_runtime.Value
var once_groupEnd sync.Once
func Get_groupEnd() gopurs_runtime.Value {
	once_groupEnd.Do(func() {
		cache_groupEnd = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return GroupEnd()
})
	})
	return cache_groupEnd
}

var cache_info gopurs_runtime.Value
var once_info sync.Once
func Get_info() gopurs_runtime.Value {
	once_info.Do(func() {
		cache_info = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Info(arg0.StrVal())()
})
})
	})
	return cache_info
}

var cache_log gopurs_runtime.Value
var once_log sync.Once
func Get_log() gopurs_runtime.Value {
	once_log.Do(func() {
		cache_log = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Log(arg0.StrVal())()
})
})
	})
	return cache_log
}

var cache_time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		cache_time = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Time(arg0.StrVal())()
})
})
	})
	return cache_time
}

var cache_timeEnd gopurs_runtime.Value
var once_timeEnd sync.Once
func Get_timeEnd() gopurs_runtime.Value {
	once_timeEnd.Do(func() {
		cache_timeEnd = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return TimeEnd(arg0.StrVal())()
})
})
	})
	return cache_timeEnd
}

var cache_timeLog gopurs_runtime.Value
var once_timeLog sync.Once
func Get_timeLog() gopurs_runtime.Value {
	once_timeLog.Do(func() {
		cache_timeLog = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return TimeLog(arg0.StrVal())()
})
})
	})
	return cache_timeLog
}

var cache_warn gopurs_runtime.Value
var once_warn sync.Once
func Get_warn() gopurs_runtime.Value {
	once_warn.Do(func() {
		cache_warn = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Warn(arg0.StrVal())()
})
})
	})
	return cache_warn
}

func Call_warnShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop interface{}) func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 interface{} = a_1_loop
_ = a_1
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_warn(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Any(a_1))), nil)
}
}

func Call_logShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop interface{}) func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 interface{} = a_1_loop
_ = a_1
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Any(a_1))), nil)
}
}

func Call_logShow__func_gopurs_runtime_Value__interface____func___gopurs_runtime_Value_2566372785(dictShow_0_loop gopurs_runtime.Value, a_1_loop interface{}) func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 interface{} = a_1_loop
_ = a_1
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Any(a_1))), nil)
}
}

func Call_infoShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop interface{}) func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 interface{} = a_1_loop
_ = a_1
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_info(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Any(a_1))), nil)
}
}

func Call_grouped(name_0_loop string, inner_1_loop func() interface{}) func() interface{} {
var name_0 string = name_0_loop
_ = name_0
var inner_1 func() interface{} = inner_1_loop
_ = inner_1
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply(Get_group(), gopurs_runtime.Str(name_0)), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(inner_1())
}), gopurs_runtime.Func(func(result_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), Get_groupEnd(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), result_3)
}))
}))
})), nil))
}
}

func Call_errorShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop interface{}) func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 interface{} = a_1_loop
_ = a_1
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_error(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Any(a_1))), nil)
}
}

func Call_debugShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop interface{}) func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 interface{} = a_1_loop
_ = a_1
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_debug(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Any(a_1))), nil)
}
}
