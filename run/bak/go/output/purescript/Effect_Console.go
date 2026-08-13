package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Console_warnShow gopurs_runtime.Value
var once_Effect_Console_warnShow sync.Once
func Get_Effect_Console_warnShow() gopurs_runtime.Value {
	once_Effect_Console_warnShow.Do(func() {
		cache_Effect_Console_warnShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Console_warnShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), a_1_box)
})
	})
	return cache_Effect_Console_warnShow
}

var cache_Effect_Console_logShow gopurs_runtime.Value
var once_Effect_Console_logShow sync.Once
func Get_Effect_Console_logShow() gopurs_runtime.Value {
	once_Effect_Console_logShow.Do(func() {
		cache_Effect_Console_logShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Console_logShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), a_1_box)
})
	})
	return cache_Effect_Console_logShow
}

var cache_Effect_Console_infoShow gopurs_runtime.Value
var once_Effect_Console_infoShow sync.Once
func Get_Effect_Console_infoShow() gopurs_runtime.Value {
	once_Effect_Console_infoShow.Do(func() {
		cache_Effect_Console_infoShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Console_infoShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), a_1_box)
})
	})
	return cache_Effect_Console_infoShow
}

var cache_Effect_Console_grouped gopurs_runtime.Value
var once_Effect_Console_grouped sync.Once
func Get_Effect_Console_grouped() gopurs_runtime.Value {
	once_Effect_Console_grouped.Do(func() {
		cache_Effect_Console_grouped = gopurs_runtime.Func2(func(name_0_box gopurs_runtime.Value, inner_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Console_grouped(name_0_box.StrVal(), inner_1_box)
})
	})
	return cache_Effect_Console_grouped
}

var cache_Effect_Console_errorShow gopurs_runtime.Value
var once_Effect_Console_errorShow sync.Once
func Get_Effect_Console_errorShow() gopurs_runtime.Value {
	once_Effect_Console_errorShow.Do(func() {
		cache_Effect_Console_errorShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Console_errorShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), a_1_box)
})
	})
	return cache_Effect_Console_errorShow
}

var cache_Effect_Console_debugShow gopurs_runtime.Value
var once_Effect_Console_debugShow sync.Once
func Get_Effect_Console_debugShow() gopurs_runtime.Value {
	once_Effect_Console_debugShow.Do(func() {
		cache_Effect_Console_debugShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Console_debugShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), a_1_box)
})
	})
	return cache_Effect_Console_debugShow
}

var cache_Effect_Console_logShow__2885109999 gopurs_runtime.Value
var once_Effect_Console_logShow__2885109999 sync.Once
func Get_Effect_Console_logShow__2885109999() gopurs_runtime.Value {
	once_Effect_Console_logShow__2885109999.Do(func() {
		cache_Effect_Console_logShow__2885109999 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Console_logShow__2885109999(a_0_box.IntVal)
})
	})
	return cache_Effect_Console_logShow__2885109999
}

var cache_Effect_Console_logShow__339054415 gopurs_runtime.Value
var once_Effect_Console_logShow__339054415 sync.Once
func Get_Effect_Console_logShow__339054415() gopurs_runtime.Value {
	once_Effect_Console_logShow__339054415.Do(func() {
		cache_Effect_Console_logShow__339054415 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Console_logShow__339054415(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), a_1_box)
})
	})
	return cache_Effect_Console_logShow__339054415
}

func Call_Effect_Console_warnShow(dictShow_0_loop *Constructor_Data_Show_Show, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_Effect_Console_warn(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), a_1).StrVal()))
}

func Call_Effect_Console_logShow(dictShow_0_loop *Constructor_Data_Show_Show, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), a_1).StrVal()))
}

func Call_Effect_Console_infoShow(dictShow_0_loop *Constructor_Data_Show_Show, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_Effect_Console_info(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), a_1).StrVal()))
}

func Call_Effect_Console_grouped(name_0_loop string, inner_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var name_0 string = name_0_loop
_ = name_0
var inner_1 gopurs_runtime.Value = inner_1_loop
_ = inner_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(Get_Effect_Console_group(), gopurs_runtime.Str(name_0))
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(inner_1, gopurs_runtime.Value{})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(Get_Effect_Console_groupEnd(), gopurs_runtime.Value{})
_ = __local_var_5_3
return __local_var_4_2
})
}

func Call_Effect_Console_errorShow(dictShow_0_loop *Constructor_Data_Show_Show, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), a_1).StrVal()))
}

func Call_Effect_Console_debugShow(dictShow_0_loop *Constructor_Data_Show_Show, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_Effect_Console_debug(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), a_1).StrVal()))
}

func Call_Effect_Console_logShow__2885109999(a_0_loop int64) gopurs_runtime.Value {
var a_0 int64 = a_0_loop
_ = a_0
return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0)).StrVal()))
}

func Call_Effect_Console_logShow__339054415(dictShow_0_loop *Constructor_Data_Show_Show, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), a_1).StrVal()))
}

func Get_Effect_Console_clear() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Clear
}

func Get_Effect_Console_debug() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Debug
}

func Get_Effect_Console_error() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Error
}

func Get_Effect_Console_group() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Group
}

func Get_Effect_Console_groupCollapsed() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_GroupCollapsed
}

func Get_Effect_Console_groupEnd() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_GroupEnd
}

func Get_Effect_Console_info() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Info
}

func Get_Effect_Console_log() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Log
}

func Get_Effect_Console_time() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Time
}

func Get_Effect_Console_timeEnd() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_TimeEnd
}

func Get_Effect_Console_timeLog() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_TimeLog
}

func Get_Effect_Console_warn() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Warn
}
