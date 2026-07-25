package Effect_Ref

import "gopurs/output/gopurs_runtime"

func _New(val interface{}) func() interface{} {
	return func() interface{} {
		return map[string]interface{}{"value": val}
	}
}
func NewWithSelf(f func(interface{}) interface{}) func() interface{} {
	return func() interface{} {
		ref := map[string]interface{}{}
		ref["value"] = f(ref)
		return ref
	}
}
func Read(ref map[string]interface{}) func() interface{} {
	return func() interface{} {
		return ref["value"]
	}
}
func ModifyImpl(f func(interface{}) map[string]interface{}, ref map[string]interface{}) func() interface{} {
	return func() interface{} {
		t := f(ref["value"])
		ref["value"] = t["state"]
		return t["value"]
	}
}
func Write(val interface{}, ref map[string]interface{}) func() {
	return func() {
		ref["value"] = val
	}
}


// --- Auto-generated FFI wrappers ---
func Call__New(arg0 interface{}) func() interface{} {
	return _New(arg0)
}
var _Gopurs__New = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _New(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_newWithSelf(arg0 func(interface{}) interface{}) func() interface{} {
	return NewWithSelf(arg0)
}
var _Gopurs_NewWithSelf = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := NewWithSelf(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_read(arg0 map[string]interface{}) func() interface{} {
	return Read(arg0)
}
var _Gopurs_Read = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_map := gopurs_runtime.RecordToMap(arg0)
	go_arg0 := make(map[string]interface{})
	for k, v := range arg0_map { go_arg0[k] = v }
	go_res := Read(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_modifyImpl(arg0 func(interface{}) map[string]interface{}, arg1 map[string]interface{}) func() interface{} {
	return ModifyImpl(arg0, arg1)
}
var _Gopurs_ModifyImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) map[string]interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[map[string]interface{}](inner_res0)
		}
	arg1_map := gopurs_runtime.RecordToMap(arg1)
	go_arg1 := make(map[string]interface{})
	for k, v := range arg1_map { go_arg1[k] = v }
	go_res := ModifyImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_write(arg0 interface{}, arg1 map[string]interface{}) func() {
	return Write(arg0, arg1)
}
var _Gopurs_Write = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	arg1_map := gopurs_runtime.RecordToMap(arg1)
	go_arg1 := make(map[string]interface{})
	for k, v := range arg1_map { go_arg1[k] = v }
	go_res := Write(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			go_res()
			return gopurs_runtime.Value{}
		})
})
