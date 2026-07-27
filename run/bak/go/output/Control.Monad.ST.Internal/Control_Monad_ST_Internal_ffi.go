package Control_Monad_ST_Internal

func Map_(f interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(_ interface{}) interface{} {
			return f.(func(interface{}) interface{})(a.(func(interface{}) interface{})(nil))
		}
	}
}

func Pure_(a interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return a
	}
}

func Bind_(a interface{}) interface{} {
	return func(f interface{}) interface{} {
		return func(_ interface{}) interface{} {
			return f.(func(interface{}) interface{})(a.(func(interface{}) interface{})(nil)).(func(interface{}) interface{})(nil)
		}
	}
}

func Run(f interface{}) interface{} {
	return f
}

func While(f interface{}) interface{} {
	return func(a interface{}) interface{} {
		return func(_ interface{}) interface{} {
			for f.(func(interface{}) interface{})(nil).(int) != 0 {
				a.(func(interface{}) interface{})(nil)
			}
			return nil
		}
	}
}

func For_(lo interface{}) interface{} {
	return func(hi interface{}) interface{} {
		return func(f interface{}) interface{} {
			return func(_ interface{}) interface{} {
				start := lo.(int)
				end := hi.(int)
				for i := start; i < end; i++ {
					f.(func(interface{}) interface{})(i).(func(interface{}) interface{})(nil)
				}
				return nil
			}
		}
	}
}

func Foreach(as interface{}) interface{} {
	return func(f interface{}) interface{} {
		return func(_ interface{}) interface{} {
			arr := as.([]interface{})
			for _, item := range arr {
				f.(func(interface{}) interface{})(item).(func(interface{}) interface{})(nil)
			}
			return nil
		}
	}
}

func New_(val interface{}) interface{} {
	return func(_ interface{}) interface{} {
		ref := &val
		return ref
	}
}

func Read(ref interface{}) interface{} {
	return func(_ interface{}) interface{} {
		ptr := ref.(*interface{})
		return *ptr
	}
}

func ModifyImpl(f interface{}) interface{} {
	return func(ref interface{}) interface{} {
		return func(_ interface{}) interface{} {
			ptr := ref.(*interface{})
			t := f.(func(interface{}) interface{})(*ptr)

			// t is { state: s, value: v }
			dict := t.(map[string]interface{})
			*ptr = dict["state"]
			return dict["value"]
		}
	}
}

func Write(a interface{}) interface{} {
	return func(ref interface{}) interface{} {
		return func(_ interface{}) interface{} {
			ptr := ref.(*interface{})
			*ptr = a
			return a
		}
	}
}
