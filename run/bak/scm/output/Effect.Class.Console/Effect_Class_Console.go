package Effect_Class_Console

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
)

var warnShow gopurs_runtime.Value
var once_warnShow sync.Once
func Get_warnShow() gopurs_runtime.Value {
	once_warnShow.Do(func() {
		warnShow = gopurs_runtime.Func3(Call_warnShow)
	})
	return warnShow
}

var warn gopurs_runtime.Value
var once_warn sync.Once
func Get_warn() gopurs_runtime.Value {
	once_warn.Do(func() {
		warn = gopurs_runtime.Func2(Call_warn)
	})
	return warn
}

var timeLog gopurs_runtime.Value
var once_timeLog sync.Once
func Get_timeLog() gopurs_runtime.Value {
	once_timeLog.Do(func() {
		timeLog = gopurs_runtime.Func2(Call_timeLog)
	})
	return timeLog
}

var timeEnd gopurs_runtime.Value
var once_timeEnd sync.Once
func Get_timeEnd() gopurs_runtime.Value {
	once_timeEnd.Do(func() {
		timeEnd = gopurs_runtime.Func2(Call_timeEnd)
	})
	return timeEnd
}

var time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		time = gopurs_runtime.Func2(Call_time)
	})
	return time
}

var logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		logShow = gopurs_runtime.Func3(Call_logShow)
	})
	return logShow
}

var log gopurs_runtime.Value
var once_log sync.Once
func Get_log() gopurs_runtime.Value {
	once_log.Do(func() {
		log = gopurs_runtime.Func2(Call_log)
	})
	return log
}

var infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		infoShow = gopurs_runtime.Func3(Call_infoShow)
	})
	return infoShow
}

var info gopurs_runtime.Value
var once_info sync.Once
func Get_info() gopurs_runtime.Value {
	once_info.Do(func() {
		info = gopurs_runtime.Func2(Call_info)
	})
	return info
}

var errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		errorShow = gopurs_runtime.Func3(Call_errorShow)
	})
	return errorShow
}

var error gopurs_runtime.Value
var once_error sync.Once
func Get_error() gopurs_runtime.Value {
	once_error.Do(func() {
		error = gopurs_runtime.Func2(Call_error)
	})
	return error
}

var debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		debugShow = gopurs_runtime.Func3(Call_debugShow)
	})
	return debugShow
}

var debug gopurs_runtime.Value
var once_debug sync.Once
func Get_debug() gopurs_runtime.Value {
	once_debug.Do(func() {
		debug = gopurs_runtime.Func2(Call_debug)
	})
	return debug
}

var clear gopurs_runtime.Value
var once_clear sync.Once
func Get_clear() gopurs_runtime.Value {
	once_clear.Do(func() {
		clear = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), pkg_Effect_Console.Get_clear())
}()
})
	})
	return clear
}

func Call_warnShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_warn(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1_loop, "show"), x_2_loop)))
}

func Call_warn(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_warn(), x_1_loop))
}

func Call_timeLog(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_timeLog(), x_1_loop))
}

func Call_timeEnd(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_timeEnd(), x_1_loop))
}

func Call_time(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_time(), x_1_loop))
}

func Call_logShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1_loop, "show"), x_2_loop)))
}

func Call_log(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), x_1_loop))
}

func Call_infoShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_info(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1_loop, "show"), x_2_loop)))
}

func Call_info(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_info(), x_1_loop))
}

func Call_errorShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1_loop, "show"), x_2_loop)))
}

func Call_error(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), x_1_loop))
}

func Call_debugShow(dictMonadEffect_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_debug(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1_loop, "show"), x_2_loop)))
}

func Call_debug(dictMonadEffect_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_debug(), x_1_loop))
}


