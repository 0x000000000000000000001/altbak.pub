package Effect_Ref

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect "gopurs/output/Effect"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		cache_void = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func(inner_arg0 func() interface{}) func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0())
})), nil)
}
}(func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, nil))
})()
})
})
	})
	return cache_void
}

var cache_new_ gopurs_runtime.Value
var once_new_ sync.Once
func Get_new_() gopurs_runtime.Value {
	once_new_.Do(func() {
		cache_new_ = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func(inner_arg0 interface{}) func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get__new(), gopurs_runtime.Any(inner_arg0)), nil)
}
}(gopurs_runtime.UnboxAny(arg0))()
})
})
	})
	return cache_new_
}

var cache_new__func_bool__func___gopurs_runtime_Value_3385984446 gopurs_runtime.Value
var once_new__func_bool__func___gopurs_runtime_Value_3385984446 sync.Once
func Get_new__func_bool__func___gopurs_runtime_Value_3385984446() gopurs_runtime.Value {
	once_new__func_bool__func___gopurs_runtime_Value_3385984446.Do(func() {
		cache_new__func_bool__func___gopurs_runtime_Value_3385984446 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func(inner_arg0 bool) func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get__new(), gopurs_runtime.Bool(inner_arg0)), nil)
}
}((arg0.IntVal) != (0))()
})
})
	})
	return cache_new__func_bool__func___gopurs_runtime_Value_3385984446
}

var cache_new__func_interface____func___gopurs_runtime_Value_2656710515 gopurs_runtime.Value
var once_new__func_interface____func___gopurs_runtime_Value_2656710515 sync.Once
func Get_new__func_interface____func___gopurs_runtime_Value_2656710515() gopurs_runtime.Value {
	once_new__func_interface____func___gopurs_runtime_Value_2656710515.Do(func() {
		cache_new__func_interface____func___gopurs_runtime_Value_2656710515 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func(inner_arg0 interface{}) func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get__new(), gopurs_runtime.Any(inner_arg0)), nil)
}
}(gopurs_runtime.UnboxAny(arg0))()
})
})
	})
	return cache_new__func_interface____func___gopurs_runtime_Value_2656710515
}

var cache_modify_prime gopurs_runtime.Value
var once_modify_prime sync.Once
func Get_modify_prime() gopurs_runtime.Value {
	once_modify_prime.Do(func() {
		cache_modify_prime = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(interface{}) interface{}, inner_arg1 gopurs_runtime.Value) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_modifyImpl(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), inner_arg1), nil))
}
}(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, arg1)())
})
})
	})
	return cache_modify_prime
}

var cache_modify_prime__func_func_interface____interface____gopurs_runtime_Value__func___interface___3101787068 gopurs_runtime.Value
var once_modify_prime__func_func_interface____interface____gopurs_runtime_Value__func___interface___3101787068 sync.Once
func Get_modify_prime__func_func_interface____interface____gopurs_runtime_Value__func___interface___3101787068() gopurs_runtime.Value {
	once_modify_prime__func_func_interface____interface____gopurs_runtime_Value__func___interface___3101787068.Do(func() {
		cache_modify_prime__func_func_interface____interface____gopurs_runtime_Value__func___interface___3101787068 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(interface{}) interface{}, inner_arg1 gopurs_runtime.Value) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_modifyImpl(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), inner_arg1), nil))
}
}(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, arg1)())
})
})
	})
	return cache_modify_prime__func_func_interface____interface____gopurs_runtime_Value__func___interface___3101787068
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_modify
}

var cache_modify__func_func_interface____interface____gopurs_runtime_Value__func___interface___2373286386 gopurs_runtime.Value
var once_modify__func_func_interface____interface____gopurs_runtime_Value__func___interface___2373286386 sync.Once
func Get_modify__func_func_interface____interface____gopurs_runtime_Value__func___interface___2373286386() gopurs_runtime.Value {
	once_modify__func_func_interface____interface____gopurs_runtime_Value__func___interface___2373286386.Do(func() {
		cache_modify__func_func_interface____interface____gopurs_runtime_Value__func___interface___2373286386 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__func_func_interface____interface____gopurs_runtime_Value__func___interface___2373286386(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_modify__func_func_interface____interface____gopurs_runtime_Value__func___interface___2373286386
}

var cache_modify_ gopurs_runtime.Value
var once_modify_ sync.Once
func Get_modify_() gopurs_runtime.Value {
	once_modify_.Do(func() {
		cache_modify_ = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Call_modify_(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, s_1_box)()
})
})
	})
	return cache_modify_
}

var cache__new gopurs_runtime.Value
var once__new sync.Once
func Get__new() gopurs_runtime.Value {
	once__new.Do(func() {
		cache__new = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return _New(gopurs_runtime.UnboxAny(arg0))()
})
})
	})
	return cache__new
}

var cache_modifyImpl gopurs_runtime.Value
var once_modifyImpl sync.Once
func Get_modifyImpl() gopurs_runtime.Value {
	once_modifyImpl.Do(func() {
		cache_modifyImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(ModifyImpl(func(inner_arg0 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0))
}, arg1)())
})
})
	})
	return cache_modifyImpl
}

var cache_newWithSelf gopurs_runtime.Value
var once_newWithSelf sync.Once
func Get_newWithSelf() gopurs_runtime.Value {
	once_newWithSelf.Do(func() {
		cache_newWithSelf = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return NewWithSelf(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
})()
})
})
	})
	return cache_newWithSelf
}

var cache_read gopurs_runtime.Value
var once_read sync.Once
func Get_read() gopurs_runtime.Value {
	once_read.Do(func() {
		cache_read = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(Read(arg0)())
})
})
	})
	return cache_read
}

var cache_write gopurs_runtime.Value
var once_write sync.Once
func Get_write() gopurs_runtime.Value {
	once_write.Do(func() {
		cache_write = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Write(gopurs_runtime.UnboxAny(arg0), arg1)()
})
})
	})
	return cache_write
}

func Call_modify(f_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(s_1)))
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_modify__func_func_interface____interface____gopurs_runtime_Value__func___interface___2373286386(f_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(s_1)))
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_modify_(f_0_loop func(interface{}) interface{}, s_1_loop gopurs_runtime.Value) func() gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := *(s_1.PtrVal().(*gopurs_runtime.Value))
_ = __local_var_2_0
*(s_1.PtrVal().(*gopurs_runtime.Value)) = gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(__local_var_2_0)))
return gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(__local_var_2_0)))
})), nil)
}
}
