package Effect_Console

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var warnShow gopurs_runtime.Value
var once_warnShow sync.Once
func Get_warnShow() gopurs_runtime.Value {
	once_warnShow.Do(func() {
		warnShow = gopurs_runtime.Func2(Call_warnShow)
	})
	return warnShow
}

var logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		logShow = gopurs_runtime.Func2(Call_logShow)
	})
	return logShow
}

var infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		infoShow = gopurs_runtime.Func2(Call_infoShow)
	})
	return infoShow
}

var errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		errorShow = gopurs_runtime.Func2(Call_errorShow)
	})
	return errorShow
}

var debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		debugShow = gopurs_runtime.Func2(Call_debugShow)
	})
	return debugShow
}

func Call_warnShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_warn(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), a_1_loop))
}

func Call_logShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), a_1_loop))
}

func Call_infoShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_info(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), a_1_loop))
}

func Call_errorShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_error(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), a_1_loop))
}

func Call_debugShow(dictShow_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_debug(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), a_1_loop))
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
