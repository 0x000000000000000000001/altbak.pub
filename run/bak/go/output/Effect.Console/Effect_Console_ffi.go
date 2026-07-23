package Effect_Console

import "gopurs/output/gopurs_runtime"



import (
	"fmt"
)

func Log(s string) func() any {
	return func() any {
		fmt.Println(s)
		return nil
	}
}

func Warn(s string) func() any {
	return func() any {
		fmt.Println("[WARN]", s)
		return nil
	}
}

func Error(s string) func() any {
	return func() any {
		fmt.Println("[ERROR]", s)
		return nil
	}
}

func Info(s string) func() any {
	return func() any {
		fmt.Println("[INFO]", s)
		return nil
	}
}

func Debug(s string) func() any {
	return Log(s)
}

func Time(s string) func() any {
	return Log(s)
}

func TimeLog(s string) func() any {
	return Log(s)
}

func TimeEnd(s string) func() any {
	return Log(s)
}

func Clear() {
}

func Group(s string) func() any {
	return Log(s)
}

func GroupCollapsed(s string) func() any {
	return Log(s)
}

func GroupEnd() {
	Clear()
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Log = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Log(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_Warn = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Warn(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_Error = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Error(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_Info = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Info(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_Debug = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Debug(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_Time = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Time(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_TimeLog = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := TimeLog(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_TimeEnd = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := TimeEnd(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_Clear = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	Clear()
	return gopurs_runtime.Value{}
})
var _Gopurs_Group = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Group(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_GroupCollapsed = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := GroupCollapsed(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_GroupEnd = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	GroupEnd()
	return gopurs_runtime.Value{}
})
