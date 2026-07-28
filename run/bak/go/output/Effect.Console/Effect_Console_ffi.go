package Effect_Console

import "gopurs/output/gopurs_runtime"



import (
	"fmt"
)

func Log(s string) func() interface{} {
	return func() interface{} {
		fmt.Println(s)
		return nil
	}
}

func Warn(s string) func() interface{} {
	return func() interface{} {
		fmt.Println("[WARN]", s)
		return nil
	}
}

func Error(s string) func() interface{} {
	return func() interface{} {
		fmt.Println("[ERROR]", s)
		return nil
	}
}

func Info(s string) func() interface{} {
	return func() interface{} {
		fmt.Println("[INFO]", s)
		return nil
	}
}

func Debug(s string) func() interface{} {
	return Log(s)
}

func Time(s string) func() interface{} {
	return Log(s)
}

func TimeLog(s string) func() interface{} {
	return Log(s)
}

func TimeEnd(s string) func() interface{} {
	return Log(s)
}

func Clear() func() interface{} {
	return func() interface{} {
		return nil
	}
}

func Group(s string) func() interface{} {
	return Log(s)
}

func GroupCollapsed(s string) func() interface{} {
	return Log(s)
}

func GroupEnd() func() interface{} {
	return func() interface{} {
		return nil
	}
}


// --- Auto-generated FFI wrappers ---
func Call_log(arg0 string) func() interface{} {
	return Log(arg0)
}
var _Gopurs_Log = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Log(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_warn(arg0 string) func() interface{} {
	return Warn(arg0)
}
var _Gopurs_Warn = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Warn(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_error(arg0 string) func() interface{} {
	return Error(arg0)
}
var _Gopurs_Error = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Error(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_info(arg0 string) func() interface{} {
	return Info(arg0)
}
var _Gopurs_Info = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Info(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_debug(arg0 string) func() interface{} {
	return Debug(arg0)
}
var _Gopurs_Debug = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Debug(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_time(arg0 string) func() interface{} {
	return Time(arg0)
}
var _Gopurs_Time = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Time(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_timeLog(arg0 string) func() interface{} {
	return TimeLog(arg0)
}
var _Gopurs_TimeLog = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := TimeLog(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_timeEnd(arg0 string) func() interface{} {
	return TimeEnd(arg0)
}
var _Gopurs_TimeEnd = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := TimeEnd(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_clear() func() interface{} {
	return Clear()
}
var _Gopurs_Clear = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	go_res := Clear()
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_group(arg0 string) func() interface{} {
	return Group(arg0)
}
var _Gopurs_Group = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Group(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_groupCollapsed(arg0 string) func() interface{} {
	return GroupCollapsed(arg0)
}
var _Gopurs_GroupCollapsed = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := GroupCollapsed(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_groupEnd() func() interface{} {
	return GroupEnd()
}
var _Gopurs_GroupEnd = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	go_res := GroupEnd()
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
