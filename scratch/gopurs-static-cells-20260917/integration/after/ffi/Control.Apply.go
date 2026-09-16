func ArrayApply(fs []func(interface{}) interface{}, xs []interface{}) []interface{} {
	result := make([]interface{}, 0, len(fs)*len(xs))
	for _, f := range fs {
		for _, x := range xs {
			result = append(result, f(x))
		}
	}
	return result
}
