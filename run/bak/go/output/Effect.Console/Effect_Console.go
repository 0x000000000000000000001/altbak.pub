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
return Call_warnShow(dictShow_0_box, a_1_box)
})
	})
	return cache_warnShow
}

var cache_logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		cache_logShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow(dictShow_0_box, a_1_box)
})
	})
	return cache_logShow
}

var cache_logShow__gopurs_runtime_Value_3442281361 gopurs_runtime.Value
var once_logShow__gopurs_runtime_Value_3442281361 sync.Once
func Get_logShow__gopurs_runtime_Value_3442281361() gopurs_runtime.Value {
	once_logShow__gopurs_runtime_Value_3442281361.Do(func() {
		cache_logShow__gopurs_runtime_Value_3442281361 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow__gopurs_runtime_Value_3442281361(dictShow_0_box, a_1_box)
})
	})
	return cache_logShow__gopurs_runtime_Value_3442281361
}

var cache_infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		cache_infoShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_infoShow(dictShow_0_box, a_1_box)
})
	})
	return cache_infoShow
}

var cache_grouped gopurs_runtime.Value
var once_grouped sync.Once
func Get_grouped() gopurs_runtime.Value {
	once_grouped.Do(func() {
		cache_grouped = gopurs_runtime.Func2(func(name_0_box gopurs_runtime.Value, inner_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_grouped(name_0_box.StrVal(), inner_1_box)
})
	})
	return cache_grouped
}

var cache_errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		cache_errorShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_errorShow(dictShow_0_box, a_1_box)
})
	})
	return cache_errorShow
}

var cache_debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		cache_debugShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_debugShow(dictShow_0_box, a_1_box)
})
	})
	return cache_debugShow
}

func Call_warnShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_warn(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), a_1))
}

func Call_logShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), a_1))
}

func Call_logShow__gopurs_runtime_Value_3442281361(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), a_1))
}

func Call_infoShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_info(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), a_1))
}

func Call_grouped(name_0_loop string, inner_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var name_0 string = name_0_loop
_ = name_0
var inner_1 gopurs_runtime.Value = inner_1_loop
_ = inner_1
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply(Get_group(), gopurs_runtime.Str(name_0)), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), inner_1, gopurs_runtime.Func(func(result_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), Get_groupEnd(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), result_3)
}))
}))
}))
}

func Call_errorShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_error(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), a_1))
}

func Call_debugShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_debug(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), a_1))
}

func Get_clear() gopurs_runtime.Value {
	return _Gopurs_Clear
}

func Get_debug() gopurs_runtime.Value {
	return _Gopurs_Debug
}

func Get_error() gopurs_runtime.Value {
	return _Gopurs_Error
}

func Get_group() gopurs_runtime.Value {
	return _Gopurs_Group
}

func Get_groupCollapsed() gopurs_runtime.Value {
	return _Gopurs_GroupCollapsed
}

func Get_groupEnd() gopurs_runtime.Value {
	return _Gopurs_GroupEnd
}

func Get_info() gopurs_runtime.Value {
	return _Gopurs_Info
}

func Get_log() gopurs_runtime.Value {
	return _Gopurs_Log
}

func Get_time() gopurs_runtime.Value {
	return _Gopurs_Time
}

func Get_timeEnd() gopurs_runtime.Value {
	return _Gopurs_TimeEnd
}

func Get_timeLog() gopurs_runtime.Value {
	return _Gopurs_TimeLog
}

func Get_warn() gopurs_runtime.Value {
	return _Gopurs_Warn
}
