package Test_FileOps

import "gopurs/output/gopurs_runtime"


import (
	"os"
)

func WriteFileSync(path string, content string, _ interface{}) interface{} {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

func ReadFileSync(path string, _ interface{}) interface{} {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func LoopE(n int, action func(interface{}) interface{}, _ interface{}) interface{} {
	for i := 0; i < n; i++ {
		action(nil)
	}
	return nil
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_LoopE = // TAST: (Func [Int, (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_res := LoopE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ReadFileSync = // TAST: (Func [String] (ADT ["Effect","Effect"] [String]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := ReadFileSync(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_WriteFileSync = // TAST: (Func [String, String] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := arg2
	go_res := WriteFileSync(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})