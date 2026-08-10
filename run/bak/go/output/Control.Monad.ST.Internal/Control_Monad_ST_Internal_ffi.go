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
	var ptr *interface{}
	if val, ok := ref.(gopurs_runtime.Value); ok {
		ptr = val.PtrVal().(*interface{})
	} else {
		ptr = ref.(*interface{})
	}
	return *ptr
}

func ModifyImpl(f func(interface{}) interface{}, ref interface{}, _ interface{}) interface{} {
	var ptr *interface{}
	if val, ok := ref.(gopurs_runtime.Value); ok {
		ptr = val.PtrVal().(*interface{})
	} else {
		ptr = ref.(*interface{})
	}
	
	t := f(*ptr)

	switch val := t.(type) {
	case map[string]interface{}:
		*ptr = val["state"]
		return val["value"]
	case gopurs_runtime.Value:
		*ptr = gopurs_runtime.RecordGet(val, "state")
		return gopurs_runtime.RecordGet(val, "value")
	default:
		panic("ModifyImpl: expected map[string]interface{} or gopurs_runtime.Value")
	}
}

func Write(a interface{}, ref interface{}, _ interface{}) interface{} {
	var ptr *interface{}
	if val, ok := ref.(gopurs_runtime.Value); ok {
		ptr = val.PtrVal().(*interface{})
	} else {
		ptr = ref.(*interface{})
	}
	*ptr = a
	return a
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Bind_ = // TAST: (ForAll [r, a, b] (Func [(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)]), (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)]))] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := arg2
	go_res := Bind_(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ForImpl = // TAST: (ForAll [r, a] (Func [Int, Int, (Func [Int] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)]))] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Unit])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := func(p0_0 int64) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg3 := arg3
	go_res := ForImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreach = // TAST: (ForAll [r, a] (Func [(Array (TypeVar a)), (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Unit]))] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Unit])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := arg2
	go_res := Foreach(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Map_ = // TAST: (ForAll [r, a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_res := Map_(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ModifyImpl = // TAST: (ForAll [r, a, b] (Func [(Func [(TypeVar a)] (Record (Row [state: (TypeVar a), value: (TypeVar b)] Empty))), (ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := ModifyImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_NewImpl = // TAST: (ForAll [a, r] (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), (TypeVar a)])])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := NewImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Pure_ = // TAST: (ForAll [r, a] (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Pure_(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Read = // TAST: (ForAll [a, r] (Func [(ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Read(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Run = // TAST: (ForAll [a] (Func [(ForAll [r] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)]))] (TypeVar a)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := Run(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_While = // TAST: (ForAll [r, a] (Func [(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Boolean]), (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Unit])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg1 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_res := While(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Write = // TAST: (ForAll [a, r] (Func [(TypeVar a), (ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Write(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})