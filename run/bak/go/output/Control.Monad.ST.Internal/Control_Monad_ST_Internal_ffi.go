package Control_Monad_ST_Internal

import "gopurs/output/gopurs_runtime"

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


// --- Auto-generated FFI wrappers ---
func Call_map_(arg0 interface{}) interface{} {
	return Map_(arg0)
}
var _Gopurs_Map_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Map_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_pure_(arg0 interface{}) interface{} {
	return Pure_(arg0)
}
var _Gopurs_Pure_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Pure_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_bind_(arg0 interface{}) interface{} {
	return Bind_(arg0)
}
var _Gopurs_Bind_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Bind_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_run(arg0 interface{}) interface{} {
	return Run(arg0)
}
var _Gopurs_Run = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Run(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_while(arg0 interface{}) interface{} {
	return While(arg0)
}
var _Gopurs_While = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := While(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_for_(arg0 interface{}) interface{} {
	return For_(arg0)
}
var _Gopurs_For_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := For_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_foreach(arg0 interface{}) interface{} {
	return Foreach(arg0)
}
var _Gopurs_Foreach = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreach(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_new_(arg0 interface{}) interface{} {
	return New_(arg0)
}
var _Gopurs_New_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := New_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_read(arg0 interface{}) interface{} {
	return Read(arg0)
}
var _Gopurs_Read = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Read(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_modifyImpl(arg0 interface{}) interface{} {
	return ModifyImpl(arg0)
}
var _Gopurs_ModifyImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := ModifyImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_write(arg0 interface{}) interface{} {
	return Write(arg0)
}
var _Gopurs_Write = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Write(go_arg0)
	return gopurs_runtime.Box(go_res)
})
