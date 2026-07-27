package Data_FunctorWithIndex

func MapWithIndexArray(f func(int) func(interface{}) interface{}, xs []interface{}) []interface{} {
	result := make([]interface{}, len(xs))
	for i, x := range xs {
		result[i] = f(i)(x)
	}
	return result
}
