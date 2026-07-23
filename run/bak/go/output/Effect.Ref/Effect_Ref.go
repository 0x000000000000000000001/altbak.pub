package Effect_Ref

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var new_ gopurs_runtime.Value
var once_new_ sync.Once
func Get_new_() gopurs_runtime.Value {
	once_new_.Do(func() {
		new_ = Get__new()
	})
	return new_
}

var modify_prime gopurs_runtime.Value
var once_modify_prime sync.Once
func Get_modify_prime() gopurs_runtime.Value {
	once_modify_prime.Do(func() {
		modify_prime = Get_modifyImpl()
	})
	return modify_prime
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
})
	})
	return modify
}

var modify_ gopurs_runtime.Value
var once_modify_ sync.Once
func Get_modify_() gopurs_runtime.Value {
	once_modify_.Do(func() {
		modify_ = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(Get_modifyImpl(), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_3_1 := gopurs_runtime.Apply(f_0, s_2)
_ = s_prime_3_1
return gopurs_runtime.RecordDict2("state", "value", s_prime_3_1, s_prime_3_1)
}), s_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_3_2 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = a_prime_3_2
return pkg_Data_Unit.Get_unit()
})
})
	})
	return modify_
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
