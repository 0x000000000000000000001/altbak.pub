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
