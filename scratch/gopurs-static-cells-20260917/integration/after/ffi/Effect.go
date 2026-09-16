package Effect

func PureE(a any, _ interface{}) any {
	return a
}
func BindE(a func(interface{}) any, f func(any) func(interface{}) any, _ interface{}) any {
	resA := a(nil)
	return f(resA)(nil)
}
func getBool(v any) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	return v.(gopurs_runtime.Value).BoolVal()
}

func UntilE(f func(interface{}) any, _ interface{}) any {
	for {
		if getBool(f(nil)) {
			break
		}
	}
	return nil
}
func WhileE(f func(interface{}) any, a func(interface{}) any, _ interface{}) any {
	for {
		if !getBool(f(nil)) {
			break
		}
		a(nil)
	}
	return nil
}
func ForE(lo int64, hi int64, f func(any) func(interface{}) any, _ interface{}) any {
	for i := lo; i < hi; i++ {
		f(i)(nil)
	}
	return nil
}
func ForeachE(as []any, f func(any) func(interface{}) any, _ interface{}) any {
	for _, a := range as {
		f(a)(nil)
	}
	return nil
}
