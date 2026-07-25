package Control_Monad_ST_Internal

import "gopurs/output/gopurs_runtime"

func Map_(f any) any {
	return func(a any) any {
		return func(_ any) any {
			return f.(func(any) any)(a.(func(any) any)(nil))
		}
	}
}

func Pure_(a any) any {
	return func(_ any) any {
		return a
	}
}

func Bind_(a any) any {
	return func(f any) any {
		return func(_ any) any {
			return f.(func(any) any)(a.(func(any) any)(nil)).(func(any) any)(nil)
		}
	}
}

func Run(f any) any {
	return f
}

func While(f any) any {
	return func(a any) any {
		return func(_ any) any {
			for f.(func(any) any)(nil).(int) != 0 {
				a.(func(any) any)(nil)
			}
			return nil
		}
	}
}

func For_(lo any) any {
	return func(hi any) any {
		return func(f any) any {
			return func(_ any) any {
				start := lo.(int)
				end := hi.(int)
				for i := start; i < end; i++ {
					f.(func(any) any)(i).(func(any) any)(nil)
				}
				return nil
			}
		}
	}
}

func Foreach(as any) any {
	return func(f any) any {
		return func(_ any) any {
			arr := as.([]any)
			for _, item := range arr {
				f.(func(any) any)(item).(func(any) any)(nil)
			}
			return nil
		}
	}
}

func New_(val any) any {
	return func(_ any) any {
		ref := &val
		return ref
	}
}

func Read(ref any) any {
	return func(_ any) any {
		ptr := ref.(*any)
		return *ptr
	}
}

func ModifyImpl(f any) any {
	return func(ref any) any {
		return func(_ any) any {
			ptr := ref.(*any)
			t := f.(func(any) any)(*ptr)

			// t is { state: s, value: v }
			dict := t.(map[string]any)
			*ptr = dict["state"]
			return dict["value"]
		}
	}
}

func Write(a any) any {
	return func(ref any) any {
		return func(_ any) any {
			ptr := ref.(*any)
			*ptr = a
			return a
		}
	}
}


// --- Auto-generated FFI wrappers ---
func Call_map_(arg0 any) any {
	return Map_(arg0)
}
var _Gopurs_Map_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Map_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_pure_(arg0 any) any {
	return Pure_(arg0)
}
var _Gopurs_Pure_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Pure_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_bind_(arg0 any) any {
	return Bind_(arg0)
}
var _Gopurs_Bind_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Bind_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_run(arg0 any) any {
	return Run(arg0)
}
var _Gopurs_Run = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Run(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_while(arg0 any) any {
	return While(arg0)
}
var _Gopurs_While = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := While(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_for_(arg0 any) any {
	return For_(arg0)
}
var _Gopurs_For_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := For_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_foreach(arg0 any) any {
	return Foreach(arg0)
}
var _Gopurs_Foreach = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreach(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_new_(arg0 any) any {
	return New_(arg0)
}
var _Gopurs_New_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := New_(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_read(arg0 any) any {
	return Read(arg0)
}
var _Gopurs_Read = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Read(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_modifyImpl(arg0 any) any {
	return ModifyImpl(arg0)
}
var _Gopurs_ModifyImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := ModifyImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_write(arg0 any) any {
	return Write(arg0)
}
var _Gopurs_Write = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Write(go_arg0)
	return gopurs_runtime.Box(go_res)
})
