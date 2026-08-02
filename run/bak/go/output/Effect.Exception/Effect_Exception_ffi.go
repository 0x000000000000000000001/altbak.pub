package Effect_Exception

import "gopurs/output/gopurs_runtime"


import "errors"

func Error(msg string) error { return errors.New(msg) }
func Message(e error) string { return e.Error() }
func Name(e error) string { return "Error" }
func ShowErrorImpl(e error) string { return e.Error() }
func StackImpl(just func(string) interface{}, nothing interface{}, e error) interface{} { return nothing }

func ThrowException(e error, _ interface{}) interface{} {
	panic(e)
}

func CatchException(c func(error) func(interface{}) interface{}, t func(interface{}) interface{}, _ interface{}) (res interface{}) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				res = c(err)(nil)
			} else {
				res = c(errors.New("panic"))(nil)
			}
		}
	}()
	return t(nil)
}

func ErrorWithCause(msg string, cause error) error { return errors.New(msg) }
func ErrorWithName(name string, msg string) error { return errors.New(msg) }


// --- Auto-generated FFI wrappers ---
var _Gopurs_CatchException = // TAST: (Func [(Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [(TypeVar a)])), (ADT ["Effect","Effect"] [(TypeVar a)])] (ADT ["Effect","Effect"] [(TypeVar a)]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 error) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_res := CatchException(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Error = // TAST: (Func [String] (ADT ["Effect","Exception","Error"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Error(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ErrorWithCause = // TAST: (Func [String, (ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Exception","Error"] []))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[error](arg1)
	go_res := ErrorWithCause(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ErrorWithName = // TAST: (Func [String, String] (ADT ["Effect","Exception","Error"] []))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := ErrorWithName(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Message = // TAST: (Func [(ADT ["Effect","Exception","Error"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := Message(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Name = // TAST: (Func [(ADT ["Effect","Exception","Error"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := Name(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ShowErrorImpl = // TAST: (Func [(ADT ["Effect","Exception","Error"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := ShowErrorImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_StackImpl = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), (ADT ["Effect","Exception","Error"] [])] (ADT ["Data","Maybe","Maybe"] [String]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[error](arg2)
	go_res := StackImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ThrowException = // TAST: (Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [(TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_arg1 := arg1
	go_res := ThrowException(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})