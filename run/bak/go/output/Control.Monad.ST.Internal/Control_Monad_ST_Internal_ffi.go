package Control_Monad_ST_Internal

import "gopurs/output/gopurs_runtime"

func Map_(f func(interface{}) interface{}, a func() interface{}) func() interface{} {
	return func() interface{} {
		return f(a())
	}
}

func Pure_(a interface{}) func() interface{} {
	return func() interface{} {
		return a
	}
}

func Bind_(a func() interface{}, f func(interface{}) func() interface{}) func() interface{} {
	return func() interface{} {
		return f(a())()
	}
}

func Run(f func() interface{}) interface{} {
	return f()
}

func While(f func() bool, a func() interface{}) func() interface{} {
	return func() interface{} {
		for f() {
			a()
		}
		return nil
	}
}

func ForImpl(lo int64, hi int64, f func(int64) func() interface{}) func() interface{} {
	return func() interface{} {
		for i := lo; i < hi; i++ {
			f(i)()
		}
		return nil
	}
}

func Foreach(as []interface{}, f func(interface{}) func() interface{}) func() interface{} {
	return func() interface{} {
		for _, item := range as {
			f(item)()
		}
		return nil
	}
}

func NewImpl(val interface{}) func() interface{} {
	return func() interface{} {
		v := val
		return &v
	}
}

func Read(ref interface{}) func() interface{} {
	return func() interface{} {
		ptr := ref.(*interface{})
		return *ptr
	}
}

func ModifyImpl(f func(interface{}) interface{}, ref interface{}) func() interface{} {
	return func() interface{} {
		ptr := ref.(*interface{})
		t := f(*ptr)

		dict := t.(map[string]interface{})
		*ptr = dict["state"]
		return dict["value"]
	}
}

func Write(a interface{}, ref interface{}) func() interface{} {
	return func() interface{} {
		ptr := ref.(*interface{})
		*ptr = a
		return a
	}
}


// --- Auto-generated FFI wrappers ---
func Call_map_(arg0 func(interface{}) interface{}, arg1 func() interface{}) func() interface{} {
	return Map_(arg0, arg1)
}
var _Gopurs_Map_ = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func() interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Value{})
		}
	go_res := Map_(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_pure_(arg0 interface{}) func() interface{} {
	return Pure_(arg0)
}
var _Gopurs_Pure_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Pure_(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_bind_(arg0 func() interface{}, arg1 func(interface{}) func() interface{}) func() interface{} {
	return Bind_(arg0, arg1)
}
var _Gopurs_Bind_ = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
		}
	go_arg1 := func(p0_0 interface{}) func() interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func() interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := Bind_(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_run(arg0 func() interface{}) interface{} {
	return Run(arg0)
}
var _Gopurs_Run = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
		}
	go_res := Run(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_while(arg0 func() bool, arg1 func() interface{}) func() interface{} {
	return While(arg0, arg1)
}
var _Gopurs_While = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg1 := func() interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Value{})
		}
	go_res := While(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_forImpl(arg0 int64, arg1 int64, arg2 func(int64) func() interface{}) func() interface{} {
	return ForImpl(arg0, arg1, arg2)
}
var _Gopurs_ForImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := func(p0_0 int64) func() interface{} {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func() interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := ForImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_foreach(arg0 []interface{}, arg1 func(interface{}) func() interface{}) func() interface{} {
	return Foreach(arg0, arg1)
}
var _Gopurs_Foreach = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 interface{}) func() interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func() interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := Foreach(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_newImpl(arg0 interface{}) func() interface{} {
	return NewImpl(arg0)
}
var _Gopurs_NewImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := NewImpl(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_read(arg0 interface{}) func() interface{} {
	return Read(arg0)
}
var _Gopurs_Read = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Read(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_modifyImpl(arg0 func(interface{}) interface{}, arg1 interface{}) func() interface{} {
	return ModifyImpl(arg0, arg1)
}
var _Gopurs_ModifyImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_res := ModifyImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_write(arg0 interface{}, arg1 interface{}) func() interface{} {
	return Write(arg0, arg1)
}
var _Gopurs_Write = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Write(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
