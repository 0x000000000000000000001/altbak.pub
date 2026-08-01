package Effect_Console

import "gopurs/output/gopurs_runtime"

import (
	"fmt"
)

func Log(s string, _ interface{}) interface{} {
	fmt.Println(s)
	return nil
}

func Warn(s string, _ interface{}) interface{} {
	fmt.Println("[WARN]", s)
	return nil
}

func Error(s string, _ interface{}) interface{} {
	fmt.Println("[ERROR]", s)
	return nil
}

func Info(s string, _ interface{}) interface{} {
	fmt.Println("[INFO]", s)
	return nil
}

func Debug(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func Time(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func TimeLog(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func TimeEnd(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func Clear(_ interface{}) interface{} {
	return nil
}

func Group(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func GroupCollapsed(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func GroupEnd(_ interface{}) interface{} {
	return nil
}


// --- Auto-generated FFI wrappers ---
func Call_log(arg0 string, arg1 interface{}) interface{} {
	return Log(arg0, arg1)
}
var _Gopurs_Log = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Log(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_warn(arg0 string, arg1 interface{}) interface{} {
	return Warn(arg0, arg1)
}
var _Gopurs_Warn = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Warn(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_error(arg0 string, arg1 interface{}) interface{} {
	return Error(arg0, arg1)
}
var _Gopurs_Error = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Error(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_info(arg0 string, arg1 interface{}) interface{} {
	return Info(arg0, arg1)
}
var _Gopurs_Info = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Info(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_debug(arg0 string, arg1 interface{}) interface{} {
	return Debug(arg0, arg1)
}
var _Gopurs_Debug = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Debug(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_time(arg0 string, arg1 interface{}) interface{} {
	return Time(arg0, arg1)
}
var _Gopurs_Time = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Time(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_timeLog(arg0 string, arg1 interface{}) interface{} {
	return TimeLog(arg0, arg1)
}
var _Gopurs_TimeLog = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := TimeLog(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_timeEnd(arg0 string, arg1 interface{}) interface{} {
	return TimeEnd(arg0, arg1)
}
var _Gopurs_TimeEnd = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := TimeEnd(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_clear(arg0 interface{}) interface{} {
	return Clear(arg0)
}
var _Gopurs_Clear = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Clear(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_group(arg0 string, arg1 interface{}) interface{} {
	return Group(arg0, arg1)
}
var _Gopurs_Group = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Group(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_groupCollapsed(arg0 string, arg1 interface{}) interface{} {
	return GroupCollapsed(arg0, arg1)
}
var _Gopurs_GroupCollapsed = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := GroupCollapsed(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_groupEnd(arg0 interface{}) interface{} {
	return GroupEnd(arg0)
}
var _Gopurs_GroupEnd = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := GroupEnd(go_arg0)
	return gopurs_runtime.Box(go_res)
})
