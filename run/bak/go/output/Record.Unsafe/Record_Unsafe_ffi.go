package Record_Unsafe

import "gopurs/output/gopurs_runtime"

func UnsafeHas(label string, rec map[string]any) bool {
	_, ok := rec[label]
	return ok
}
func UnsafeGet(label string, rec map[string]any) any {
	return rec[label]
}
func UnsafeSet(label string, value any, rec map[string]any) map[string]any {
	newMap := make(map[string]any, len(rec)+1)
	for k, v := range rec {
		newMap[k] = v
	}
	newMap[label] = value
	return newMap
}
func UnsafeDelete(label string, rec map[string]any) map[string]any {
	newMap := make(map[string]any)
	for k, v := range rec {
		if k != label {
			newMap[k] = v
		}
	}
	return newMap
}


// --- Auto-generated FFI wrappers ---
func Call_unsafeHas(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	arg1_map := gopurs_runtime.RecordToMap(arg1)
	go_arg1 := make(map[string]any)
	for k, v := range arg1_map { go_arg1[k] = v }
	go_res := UnsafeHas(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_UnsafeHas = gopurs_runtime.Func2(Call_unsafeHas)
func Call_unsafeGet(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	arg1_map := gopurs_runtime.RecordToMap(arg1)
	go_arg1 := make(map[string]any)
	for k, v := range arg1_map { go_arg1[k] = v }
	go_res := UnsafeGet(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_UnsafeGet = gopurs_runtime.Func2(Call_unsafeGet)
func Call_unsafeSet(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	arg2_map := gopurs_runtime.RecordToMap(arg2)
	go_arg2 := make(map[string]any)
	for k, v := range arg2_map { go_arg2[k] = v }
	go_res := UnsafeSet(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_map := make(map[string]gopurs_runtime.Value)
			for k, v := range go_res { res_map[k] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Record(res_map)
		}()
}
var _Gopurs_UnsafeSet = gopurs_runtime.Func3(Call_unsafeSet)
func Call_unsafeDelete(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	arg1_map := gopurs_runtime.RecordToMap(arg1)
	go_arg1 := make(map[string]any)
	for k, v := range arg1_map { go_arg1[k] = v }
	go_res := UnsafeDelete(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_map := make(map[string]gopurs_runtime.Value)
			for k, v := range go_res { res_map[k] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Record(res_map)
		}()
}
var _Gopurs_UnsafeDelete = gopurs_runtime.Func2(Call_unsafeDelete)
