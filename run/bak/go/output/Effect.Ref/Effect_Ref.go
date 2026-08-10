package Effect_Ref

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_new gopurs_runtime.Value
var once_new sync.Once
func Get_new() gopurs_runtime.Value {
	once_new.Do(func() {
		cache_new = Get__new()
	})
	return cache_new
}

var cache_new__gopurs_runtime_Value_2045356851 gopurs_runtime.Value
var once_new__gopurs_runtime_Value_2045356851 sync.Once
func Get_new__gopurs_runtime_Value_2045356851() gopurs_runtime.Value {
	once_new__gopurs_runtime_Value_2045356851.Do(func() {
		cache_new__gopurs_runtime_Value_2045356851 = Get__new()
	})
	return cache_new__gopurs_runtime_Value_2045356851
}

var cache_new__gopurs_runtime_Value_1693026106 gopurs_runtime.Value
var once_new__gopurs_runtime_Value_1693026106 sync.Once
func Get_new__gopurs_runtime_Value_1693026106() gopurs_runtime.Value {
	once_new__gopurs_runtime_Value_1693026106.Do(func() {
		cache_new__gopurs_runtime_Value_1693026106 = Get__new()
	})
	return cache_new__gopurs_runtime_Value_1693026106
}

var cache_modify_prime gopurs_runtime.Value
var once_modify_prime sync.Once
func Get_modify_prime() gopurs_runtime.Value {
	once_modify_prime.Do(func() {
		cache_modify_prime = Get_modifyImpl()
	})
	return cache_modify_prime
}

var cache_modify_prime__gopurs_runtime_Value_3296699741 gopurs_runtime.Value
var once_modify_prime__gopurs_runtime_Value_3296699741 sync.Once
func Get_modify_prime__gopurs_runtime_Value_3296699741() gopurs_runtime.Value {
	once_modify_prime__gopurs_runtime_Value_3296699741.Do(func() {
		cache_modify_prime__gopurs_runtime_Value_3296699741 = Get_modifyImpl()
	})
	return cache_modify_prime__gopurs_runtime_Value_3296699741
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

func Call_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Get__new() gopurs_runtime.Value {
	return _Gopurs__New
}

func Get_modifyImpl() gopurs_runtime.Value {
	return _Gopurs_ModifyImpl
}

func Get_modify_() gopurs_runtime.Value {
	return _Gopurs_Modify_
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
