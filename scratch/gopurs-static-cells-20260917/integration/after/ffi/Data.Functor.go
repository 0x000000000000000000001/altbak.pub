func ArrayMap(f func(interface{}) interface{}, arr []interface{}) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}
