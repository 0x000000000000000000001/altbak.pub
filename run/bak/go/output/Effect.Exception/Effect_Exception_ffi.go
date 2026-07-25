package Effect_Exception

import "gopurs/output/gopurs_runtime"

import "errors"
func Error(msg string) error { return errors.New(msg) }
func Message(e error) string { return e.Error() }
func Name(e error) string { return "Error" }
func ShowErrorImpl(e error) string { return e.Error() }
func StackImpl(just func(string) interface{}, nothing interface{}, e error) interface{} { return nothing }
func ThrowException(e error) func() interface{} { return func() interface{} { panic(e) } }
func CatchException(c func(error) func() interface{}, t func() interface{}) func() interface{} {
	return func() (res interface{}) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					res = c(err)()
				} else {
					// Fallback if panic is not an error type
					res = c(errors.New("panic"))()
				}
			}
		}()
		return t()
	}
}
func ErrorWithCause(msg string, cause error) error { return errors.New(msg) }
func ErrorWithName(name string, msg string) error { return errors.New(msg) }


// --- Auto-generated FFI wrappers ---
func Call_error(arg0 string) error { return errors.New(msg) } {
	return Error(arg0)
}
var _Gopurs_Error = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Error(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_message(arg0 error) string { return e.Error() } {
	return Message(arg0)
}
var _Gopurs_Message = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := Message(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_name(arg0 error) string { return "Error" } {
	return Name(arg0)
}
var _Gopurs_Name = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := Name(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_showErrorImpl(arg0 error) string { return e.Error() } {
	return ShowErrorImpl(arg0)
}
var _Gopurs_ShowErrorImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := ShowErrorImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_stackImpl(arg0 func(string) interface{}, arg1 interface{}, arg2 error) interface{} { return nothing } {
	return StackImpl(arg0, arg1, arg2)
}
var _Gopurs_StackImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[error](arg2)
	go_res := StackImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_throwException(arg0 error) func() interface{} { return func() interface{} { panic(e) } } {
	return ThrowException(arg0)
}
var _Gopurs_ThrowException = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := ThrowException(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_catchException(arg0 func(error) func() interface{}, arg1 func() interface{}) func() interface{} {
	return CatchException(arg0, arg1)
}
var _Gopurs_CatchException = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 error) func() interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func() interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_arg1 := func() interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Value{})
		}
	go_res := CatchException(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_errorWithCause(arg0 string, arg1 error) error { return errors.New(msg) } {
	return ErrorWithCause(arg0, arg1)
}
var _Gopurs_ErrorWithCause = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[error](arg1)
	go_res := ErrorWithCause(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_errorWithName(arg0 string, arg1 string) error { return errors.New(msg) } {
	return ErrorWithName(arg0, arg1)
}
var _Gopurs_ErrorWithName = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := ErrorWithName(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
