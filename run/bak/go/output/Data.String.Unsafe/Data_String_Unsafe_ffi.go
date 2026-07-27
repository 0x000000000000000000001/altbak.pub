package Data_String_Unsafe

func CharAt(i interface{}) interface{} {
	return func(s interface{}) interface{} {
		str := s.(string)
		idx := int(i.(int))
		if idx >= 0 && idx < len(str) {
			return (string(str[idx]))
		}
		panic("Data.String.Unsafe.charAt: Invalid index.")
	}
}

func Char(s interface{}) interface{} {
	str := s.(string)
	if len(str) == 1 {
		return (string(str[0]))
	}
	panic("Data.String.Unsafe.char: Expected string of length 1.")
}
