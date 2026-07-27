package Effect



func PureE(a interface{}) interface{} {
	return func() interface{} {
		return a
	}
}

func BindE(a interface{}, f func(interface{}) interface{}) interface{} {
	return func() interface{} {
		resA := a.(func() interface{})()
		return f(resA).(func() interface{})()
	}
}

func UntilE(f interface{}) interface{} {
	return func() interface{} {
		for {
			if f.(func() interface{})().(bool) {
				break
			}
		}
		return nil
	}
}

func WhileE(f interface{}, a interface{}) interface{} {
	return func() interface{} {
		for {
			if !f.(func() interface{})().(bool) {
				break
			}
			a.(func() interface{})()
		}
		return nil
	}
}

func ForE(lo int64, hi int64, f func(int64) interface{}) interface{} {
	return func() interface{} {
		for i := lo; i < hi; i++ {
			f(i).(func() interface{})()
		}
		return nil
	}
}

func ForeachE(as []interface{}, f func(interface{}) interface{}) interface{} {
	return func() interface{} {
		for _, v := range as {
			f(v).(func() interface{})()
		}
		return nil
	}
}
