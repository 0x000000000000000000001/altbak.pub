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
		cache_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void
}

var cache_new_ gopurs_runtime.Value
var once_new_ sync.Once
func Get_new_() gopurs_runtime.Value {
	once_new_.Do(func() {
		cache_new_ = Get__new()
	})
	return cache_new_
}

var cache_new__gopurs_runtime_Value_1731740265 gopurs_runtime.Value
var once_new__gopurs_runtime_Value_1731740265 sync.Once
func Get_new__gopurs_runtime_Value_1731740265() gopurs_runtime.Value {
	once_new__gopurs_runtime_Value_1731740265.Do(func() {
		cache_new__gopurs_runtime_Value_1731740265 = Get__new()
	})
	return cache_new__gopurs_runtime_Value_1731740265
}

var cache_new__gopurs_runtime_Value_22483300 gopurs_runtime.Value
var once_new__gopurs_runtime_Value_22483300 sync.Once
func Get_new__gopurs_runtime_Value_22483300() gopurs_runtime.Value {
	once_new__gopurs_runtime_Value_22483300.Do(func() {
		cache_new__gopurs_runtime_Value_22483300 = Get__new()
	})
	return cache_new__gopurs_runtime_Value_22483300
}

var cache_modify_prime gopurs_runtime.Value
var once_modify_prime sync.Once
func Get_modify_prime() gopurs_runtime.Value {
	once_modify_prime.Do(func() {
		cache_modify_prime = Get_modifyImpl()
	})
	return cache_modify_prime
}

var cache_modify_prime__gopurs_runtime_Value_2048033157 gopurs_runtime.Value
var once_modify_prime__gopurs_runtime_Value_2048033157 sync.Once
func Get_modify_prime__gopurs_runtime_Value_2048033157() gopurs_runtime.Value {
	once_modify_prime__gopurs_runtime_Value_2048033157.Do(func() {
		cache_modify_prime__gopurs_runtime_Value_2048033157 = Get_modifyImpl()
	})
	return cache_modify_prime__gopurs_runtime_Value_2048033157
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(f_0_box)
})
	})
	return cache_modify
}

var cache_modify__gopurs_runtime_Value_31931706 gopurs_runtime.Value
var once_modify__gopurs_runtime_Value_31931706 sync.Once
func Get_modify__gopurs_runtime_Value_31931706() gopurs_runtime.Value {
	once_modify__gopurs_runtime_Value_31931706.Do(func() {
		cache_modify__gopurs_runtime_Value_31931706 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__gopurs_runtime_Value_31931706(f_0_box)
})
	})
	return cache_modify__gopurs_runtime_Value_31931706
}

var cache_modify_ gopurs_runtime.Value
var once_modify_ sync.Once
func Get_modify_() gopurs_runtime.Value {
	once_modify_.Do(func() {
		cache_modify_ = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify_(f_0_box, s_1_box)
})
	})
	return cache_modify_
}

func Call_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_modify__gopurs_runtime_Value_31931706(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_modify_(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := *(s_1.PtrVal().(*gopurs_runtime.Value))
_ = __local_var_2_0
*(s_1.PtrVal().(*gopurs_runtime.Value)) = gopurs_runtime.Apply(f_0, __local_var_2_0)
return gopurs_runtime.Apply(f_0, __local_var_2_0)
}))
}

func Get__new() gopurs_runtime.Value {
	return _Gopurs__New
}

func Get_modifyImpl() gopurs_runtime.Value {
	return _Gopurs_ModifyImpl
}

func Get_newWithSelf() gopurs_runtime.Value {
	return _Gopurs_NewWithSelf
}

func Get_read() gopurs_runtime.Value {
	return _Gopurs_Read
}

func Get_write() gopurs_runtime.Value {
	return _Gopurs_Write
}
