package Record_Unsafe

import "gopurs/output/gopurs_runtime"

func UnsafeHas(label string, recVal interface{}) bool {
	v := recVal.(gopurs_runtime.Value)
	m := gopurs_runtime.RecordToMap(v)
	_, ok := m[label]
	return ok
}

func UnsafeGet(label string, recVal interface{}) interface{} {
	v := recVal.(gopurs_runtime.Value)
	return gopurs_runtime.RecordGet(v, label)
}

func UnsafeSet(label string, value interface{}, recVal interface{}) interface{} {
	v := recVal.(gopurs_runtime.Value)
	val := value.(gopurs_runtime.Value)
	return gopurs_runtime.RecordUpdate1(v, label, val)
}

func UnsafeDelete(label string, recVal interface{}) interface{} {
	v := recVal.(gopurs_runtime.Value)
	m := gopurs_runtime.RecordToMap(v)
	newMap := make(map[string]gopurs_runtime.Value)
	for k, val := range m {
		if k != label {
			newMap[k] = val
		}
	}
	return gopurs_runtime.Record(newMap)
}


// --- Auto-generated FFI wrappers ---
func Call_unsafeHas(arg0 string, arg1 interface{}) bool {
	return UnsafeHas(arg0, arg1)
}
var _Gopurs_UnsafeHas = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := UnsafeHas(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_unsafeGet(arg0 string, arg1 interface{}) interface{} {
	return UnsafeGet(arg0, arg1)
}
var _Gopurs_UnsafeGet = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := UnsafeGet(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_unsafeSet(arg0 string, arg1 interface{}, arg2 interface{}) interface{} {
	return UnsafeSet(arg0, arg1, arg2)
}
var _Gopurs_UnsafeSet = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := UnsafeSet(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_unsafeDelete(arg0 string, arg1 interface{}) interface{} {
	return UnsafeDelete(arg0, arg1)
}
var _Gopurs_UnsafeDelete = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := UnsafeDelete(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
