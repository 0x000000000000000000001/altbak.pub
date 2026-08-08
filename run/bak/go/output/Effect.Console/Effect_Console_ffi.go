package Effect_Console

import "gopurs/output/gopurs_runtime"


import (
	"fmt"
	"strings"
	"sync/atomic"
)

var indentation int32 = 0

func getIndent() string {
	ind := atomic.LoadInt32(&indentation)
	if ind < 0 {
		ind = 0
	}
	return strings.Repeat("  ", int(ind))
}

func Log(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Warn(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Error(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Info(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Debug(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
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
	Log(s, nil)
	atomic.AddInt32(&indentation, 1)
	return nil
}

func GroupCollapsed(s string, _ interface{}) interface{} {
	Log(s, nil)
	atomic.AddInt32(&indentation, 1)
	return nil
}

func GroupEnd(_ interface{}) interface{} {
	ind := atomic.LoadInt32(&indentation)
	if ind > 0 {
		atomic.AddInt32(&indentation, -1)
	}
	return nil
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Clear = // TAST: (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Clear(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Debug = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Debug(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Error = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Error(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Group = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Group(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_GroupCollapsed = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := GroupCollapsed(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_GroupEnd = // TAST: (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := GroupEnd(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Info = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Info(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Log = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Log(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Time = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Time(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_TimeEnd = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := TimeEnd(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_TimeLog = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := TimeLog(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Warn = // TAST: (Func [String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Warn(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})