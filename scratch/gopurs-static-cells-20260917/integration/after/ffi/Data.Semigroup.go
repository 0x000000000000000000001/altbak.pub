func ConcatString(s1 string, s2 string) string {
	return s1 + s2
}
func ConcatArray(xs []interface{}, ys []interface{}) []interface{} {
	if len(xs) == 0 {
		return ys
	}
	if len(ys) == 0 {
		return xs
	}
	res := make([]interface{}, 0, len(xs)+len(ys))
	res = append(res, xs...)
	res = append(res, ys...)
	return res
}
