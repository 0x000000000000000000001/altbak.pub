package Control_Monad_ST_Internal

import "gopurs/output/gopurs_runtime"

func Map_(f func(interface{}) interface{}, a func(interface{}) interface{}, _ interface{}) interface{} {
	return f(a(nil))
}

func Pure_(a interface{}, _ interface{}) interface{} {
	return a
}

func Bind_(a func(interface{}) interface{}, f func(interface{}) func(interface{}) interface{}, _ interface{}) interface{} {
	return f(a(nil))(nil)
}

func Run(f func(interface{}) interface{}) interface{} {
	return f(nil)
}

func While(f func() bool, a func(interface{}) interface{}, _ interface{}) interface{} {
	for f() {
		a(nil)
	}
	return nil
}

func ForImpl(lo int64, hi int64, f func(int64) func(interface{}) interface{}, _ interface{}) interface{} {
	for i := lo; i < hi; i++ {
		f(i)(nil)
	}
	return nil
}

func Foreach(as []interface{}, f func(interface{}) func(interface{}) interface{}, _ interface{}) interface{} {
	for _, item := range as {
		f(item)(nil)
	}
	return nil
}

func NewImpl(val interface{}, _ interface{}) interface{} {
	v := val
	return &v
}

func Read(ref interface{}, _ interface{}) interface{} {
	ptr := ref.(*interface{})
	return *ptr
}

func ModifyImpl(f func(interface{}) interface{}, ref interface{}, _ interface{}) interface{} {
	ptr := ref.(*interface{})
	t := f(*ptr)

	dict := t.(map[string]interface{})
	*ptr = dict["state"]
	return dict["value"]
}

func Write(a interface{}, ref interface{}, _ interface{}) interface{} {
	ptr := ref.(*interface{})
	*ptr = a
	return a
}


// --- Auto-generated FFI wrappers ---
func Call_map_(arg0 func(interface{}) interface{}, arg1 func(interface{}) interface{}, arg2 interface{}) interface{} {
	return Map_(arg0, arg1, arg2)
}
var _Gopurs_Map_ = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_res := Map_(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_pure_(arg0 interface{}, arg1 interface{}) interface{} {
	return Pure_(arg0, arg1)
}
var _Gopurs_Pure_ = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Pure_(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_bind_(arg0 func(interface{}) interface{}, arg1 func(interface{}) func(interface{}) interface{}, arg2 interface{}) interface{} {
	return Bind_(arg0, arg1, arg2)
}
var _Gopurs_Bind_ = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := arg2
	go_res := Bind_(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_run(arg0 func(interface{}) interface{}) interface{} {
	return Run(arg0)
}
var _Gopurs_Run = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := Run(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_while(arg0 func() bool, arg1 func(interface{}) interface{}, arg2 interface{}) interface{} {
	return While(arg0, arg1, arg2)
}
var _Gopurs_While = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg1 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_res := While(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_forImpl(arg0 int64, arg1 int64, arg2 func(int64) func(interface{}) interface{}, arg3 interface{}) interface{} {
	return ForImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_ForImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := func(p0_0 int64) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg3 := arg3
	go_res := ForImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_foreach(arg0 []interface{}, arg1 func(interface{}) func(interface{}) interface{}, arg2 interface{}) interface{} {
	return Foreach(arg0, arg1, arg2)
}
var _Gopurs_Foreach = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := arg2
	go_res := Foreach(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_newImpl(arg0 interface{}, arg1 interface{}) interface{} {
	return NewImpl(arg0, arg1)
}
var _Gopurs_NewImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := NewImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_read(arg0 interface{}, arg1 interface{}) interface{} {
	return Read(arg0, arg1)
}
var _Gopurs_Read = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Read(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_modifyImpl(arg0 func(interface{}) interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return ModifyImpl(arg0, arg1, arg2)
}
var _Gopurs_ModifyImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := ModifyImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_write(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return Write(arg0, arg1, arg2)
}
var _Gopurs_Write = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Write(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
