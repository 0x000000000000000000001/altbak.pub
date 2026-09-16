func CharAt(i interface{}) interface{} {
	return func(s interface{}) interface{} {
		str := gopurs_runtime.Unbox[string](s)
		idx := gopurs_runtime.Unbox[int](i)
		if idx >= 0 && idx < len(str) {
			return string(str[idx])
		}
		panic("Data.String.Unsafe.charAt: Invalid index.")
	}
}

func Char(s interface{}) interface{} {
	str := gopurs_runtime.Unbox[string](s)
	if len(str) == 1 {
		return string(str[0])
	}
	panic("Data.String.Unsafe.char: Expected string of length 1.")
}
