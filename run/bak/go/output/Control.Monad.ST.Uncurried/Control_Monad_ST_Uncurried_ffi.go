package Control_Monad_ST_Uncurried

func MkSTFn1(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn2(fn interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn3(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn4(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn5(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn6(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn7(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn8(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn9(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(nil)
	}
}
func MkSTFn10(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
		return fn.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(f).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(j).(func(interface{}) interface{})(nil)
	}
}
func RunSTFn1(fn interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}) interface{})(a) }
	}
}
func RunSTFn2(fn interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}) interface{})(a, b) }
	}
}
func RunSTFn3(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}, interface{}) interface{})(a, b, c) }
	}
}
func RunSTFn4(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d) }
	}
}
func RunSTFn5(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e) }
	}
}
func RunSTFn6(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, f) }
	}
}
func RunSTFn7(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, f, g) }
	}
}
func RunSTFn8(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, f, g, h) }
	}
}
func RunSTFn9(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, f, g, h, i) }
	}
}
func RunSTFn10(fn interface{}) interface{} {
	return func(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
		return func() interface{} { return fn.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, f, g, h, i, j) }
	}
}
