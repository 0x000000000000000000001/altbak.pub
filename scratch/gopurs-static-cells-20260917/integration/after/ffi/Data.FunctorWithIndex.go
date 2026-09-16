func MapWithIndexArray(f func(int64, interface{}) interface{}, xs []interface{}) []interface{} {
	result := make([]interface{}, len(xs))
	for i, x := range xs {
		result[i] = f(int64(i), x)
	}
	return result
}
