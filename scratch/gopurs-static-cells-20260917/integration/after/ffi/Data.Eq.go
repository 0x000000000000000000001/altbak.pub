func refEq(r1 interface{}, r2 interface{}) bool {
	return r1 == r2
}
func EqBooleanImpl(r1 bool, r2 bool) bool {
	return r1 == r2
}
func EqIntImpl(r1 int64, r2 int64) bool {
	return r1 == r2
}
func EqNumberImpl(r1 float64, r2 float64) bool {
	return r1 == r2
}
func EqCharImpl(r1 string, r2 string) bool {
	return r1 == r2
}
func EqStringImpl(r1 string, r2 string) bool {
	return r1 == r2
}
func EqArrayImpl(f func(interface{}, interface{}) bool, xs []interface{}, ys []interface{}) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if !f(xs[i], ys[i]) {
			return false
		}
	}
	return true
}
