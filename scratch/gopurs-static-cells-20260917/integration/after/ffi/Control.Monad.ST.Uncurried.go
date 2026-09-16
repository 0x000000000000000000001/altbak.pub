package Control_Monad_ST_Uncurried

func MkSTFn1(fn func(any) any) any {
	return func(a any) any { return fn(a) }
}
func MkSTFn2(fn func(any) func(any) any) any {
	return func(a any, b any) any { return fn(a)(b) }
}
func MkSTFn3(fn func(any) func(any) func(any) any) any {
	return func(a any, b any, c any) any { return fn(a)(b)(c) }
}
func MkSTFn4(fn func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any) any { return fn(a)(b)(c)(d) }
}
func MkSTFn5(fn func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any) any { return fn(a)(b)(c)(d)(e) }
}
func MkSTFn6(fn func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any) any { return fn(a)(b)(c)(d)(e)(f) }
}
func MkSTFn7(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any) any { return fn(a)(b)(c)(d)(e)(f)(g) }
}
func MkSTFn8(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any, h any) any { return fn(a)(b)(c)(d)(e)(f)(g)(h) }
}
func MkSTFn9(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any, h any, i any) any { return fn(a)(b)(c)(d)(e)(f)(g)(h)(i) }
}
// func MkSTFn10(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	// return func any { return fn(a)(b)(c)(d)(e)(f)(g)(h)(i)(j) }
// }
func RunSTFn1(fn func(any) any, a any, _ interface{}) any { return fn(a) }
func RunSTFn2(fn func(any, any) any, a any, b any, _ interface{}) any { return fn(a, b) }
func RunSTFn3(fn func(any, any, any) any, a any, b any, c any, _ interface{}) any { return fn(a, b, c) }
func RunSTFn4(fn func(any, any, any, any) any, a any, b any, c any, d any, _ interface{}) any { return fn(a, b, c, d) }
func RunSTFn5(fn func(any, any, any, any, any) any, a any, b any, c any, d any, e any, _ interface{}) any { return fn(a, b, c, d, e) }
func RunSTFn6(fn func(any, any, any, any, any, any) any, a any, b any, c any, d any, e any, f any, _ interface{}) any { return fn(a, b, c, d, e, f) }
func RunSTFn7(fn func(any, any, any, any, any, any, any) any, a any, b any, c any, d any, e any, f any, g any, _ interface{}) any { return fn(a, b, c, d, e, f, g) }
func RunSTFn8(fn func(any, any, any, any, any, any, any, any) any, a any, b any, c any, d any, e any, f any, g any, h any, _ interface{}) any { return fn(a, b, c, d, e, f, g, h) }
func RunSTFn9(fn func(any, any, any, any, any, any, any, any, any) any, a any, b any, c any, d any, e any, f any, g any, h any, i any, _ interface{}) any { return fn(a, b, c, d, e, f, g, h, i) }
// func RunSTFn10(fn any, a any, b any, c any, d any, e any, f any, g any, h any, i any, j any, _ interface{}) any { return fn.(func(any, any, any, any, any, any, any, any, any, any) any)(a, b, c, d, e, f, g, h, i, j) }
