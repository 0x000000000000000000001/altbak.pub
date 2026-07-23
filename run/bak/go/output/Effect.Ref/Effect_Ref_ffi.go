package Effect_Ref

import "gopurs/output/gopurs_runtime"

func _New(val any) func() any {
	return func() any {
		return map[string]any{"value": val}
	}
}
func NewWithSelf(f func(any) any) func() any {
	return func() any {
		ref := map[string]any{}
		ref["value"] = f(ref)
		return ref
	}
}
func Read(ref map[string]any) func() any {
	return func() any {
		return ref["value"]
	}
}
func ModifyImpl(f func(any) map[string]any, ref map[string]any) func() any {
	return func() any {
		t := f(ref["value"])
		ref["value"] = t["state"]
		return t["value"]
	}
}
func Write(val any, ref map[string]any) func() {
	return func() {
		ref["value"] = val
	}
}


// --- Auto-generated FFI wrappers ---
var _Gopurs__New = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := _New(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_NewWithSelf = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0 any) any {
		res := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0))
		return res.PtrVal
	}
	go_res := NewWithSelf(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_Read = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_map := arg0.PtrVal.(map[string]gopurs_runtime.Value)
	go_arg0 := make(map[string]any)
	for k, v := range arg0_map { go_arg0[k] = v.PtrVal }
	go_res := Read(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_ModifyImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0 any) map[string]any {
		res := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0))
		return gopurs_runtime.Unbox[map[string]any](res)
	}
	arg1_map := arg1.PtrVal.(map[string]gopurs_runtime.Value)
	go_arg1 := make(map[string]any)
	for k, v := range arg1_map { go_arg1[k] = v.PtrVal }
	go_res := ModifyImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_Write = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	arg1_map := arg1.PtrVal.(map[string]gopurs_runtime.Value)
	go_arg1 := make(map[string]any)
	for k, v := range arg1_map { go_arg1[k] = v.PtrVal }
	go_res := Write(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Value{}
		})
})
