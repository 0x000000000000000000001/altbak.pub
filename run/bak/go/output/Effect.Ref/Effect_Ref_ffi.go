package Effect_Ref

import "gopurs/output/gopurs_runtime"

func _New(val interface{}, _ interface{}) interface{} {
	return map[string]interface{}{"value": val}
}
func NewWithSelf(f func(interface{}) interface{}, _ interface{}) interface{} {
	ref := map[string]interface{}{}
	ref["value"] = f(ref)
	return ref
}
func Read(ref interface{}, _ interface{}) interface{} {
	return ref.(map[string]interface{})["value"]
}
func ModifyImpl(f func(interface{}) interface{}, ref interface{}, _ interface{}) interface{} {
	t := f(ref.(map[string]interface{})["value"]).(map[string]interface{})
	ref.(map[string]interface{})["value"] = t["state"]
	return t["value"]
}
func Write(val interface{}, ref interface{}, _ interface{}) interface{} {
	ref.(map[string]interface{})["value"] = val
	return nil
}


// --- Auto-generated FFI wrappers ---
func Call__New(arg0 interface{}, arg1 interface{}) interface{} {
	return _New(arg0, arg1)
}
var _Gopurs__New = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _New(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_newWithSelf(arg0 func(interface{}) interface{}, arg1 interface{}) interface{} {
	return NewWithSelf(arg0, arg1)
}
var _Gopurs_NewWithSelf = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_res := NewWithSelf(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_read(arg0 interface{}, arg1 interface{}) interface{} {
	return Read(arg0, arg1)
}
var _Gopurs_Read = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Read(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_modifyImpl(arg0 func(interface{}) interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return ModifyImpl(arg0, arg1, arg2)
}
var _Gopurs_ModifyImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := ModifyImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_write(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return Write(arg0, arg1, arg2)
}
var _Gopurs_Write = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Write(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
