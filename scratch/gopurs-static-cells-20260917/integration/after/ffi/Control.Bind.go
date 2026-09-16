func ArrayBind(arr []interface{}, f func(interface{}) []interface{}) []interface{} {
	var result []interface{}
	for _, v := range arr {
		result = append(result, f(v)...)
	}
	return result
}
