package Record_Unsafe

func UnsafeHas(label string, rec interface{}) bool {
	m := rec.(map[string]interface{})
	_, ok := m[label]
	return ok
}
func UnsafeGet(label string, rec interface{}) interface{} {
	m := rec.(map[string]interface{})
	return m[label]
}
func UnsafeSet(label string, val interface{}, rec interface{}) interface{} {
	m := rec.(map[string]interface{})
	res := make(map[string]interface{})
	for k, v := range m {
		res[k] = v
	}
	res[label] = val
	return res
}
func UnsafeDelete(label string, rec interface{}) interface{} {
	m := rec.(map[string]interface{})
	res := make(map[string]interface{})
	for k, v := range m {
		if k != label {
			res[k] = v
		}
	}
	return res
}
