package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Ref_new gopurs_runtime.Value
var once_Effect_Ref_new sync.Once
func Get_Effect_Ref_new() gopurs_runtime.Value {
	once_Effect_Ref_new.Do(func() {
		cache_Effect_Ref_new = Get_Effect_Ref__new()
	})
	return cache_Effect_Ref_new
}

var cache_Effect_Ref_modify_prime gopurs_runtime.Value
var once_Effect_Ref_modify_prime sync.Once
func Get_Effect_Ref_modify_prime() gopurs_runtime.Value {
	once_Effect_Ref_modify_prime.Do(func() {
		cache_Effect_Ref_modify_prime = Get_Effect_Ref_modifyImpl()
	})
	return cache_Effect_Ref_modify_prime
}

var cache_Effect_Ref_modify gopurs_runtime.Value
var once_Effect_Ref_modify sync.Once
func Get_Effect_Ref_modify() gopurs_runtime.Value {
	once_Effect_Ref_modify.Do(func() {
		cache_Effect_Ref_modify = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Ref_modify(f_0_box)
})
	})
	return cache_Effect_Ref_modify
}

var cache_Effect_Ref_modify_prime__3296699741 gopurs_runtime.Value
var once_Effect_Ref_modify_prime__3296699741 sync.Once
func Get_Effect_Ref_modify_prime__3296699741() gopurs_runtime.Value {
	once_Effect_Ref_modify_prime__3296699741.Do(func() {
		cache_Effect_Ref_modify_prime__3296699741 = Get_Effect_Ref_modifyImpl()
	})
	return cache_Effect_Ref_modify_prime__3296699741
}

var cache_Effect_Ref_new__2045356851 gopurs_runtime.Value
var once_Effect_Ref_new__2045356851 sync.Once
func Get_Effect_Ref_new__2045356851() gopurs_runtime.Value {
	once_Effect_Ref_new__2045356851.Do(func() {
		cache_Effect_Ref_new__2045356851 = Get_Effect_Ref__new()
	})
	return cache_Effect_Ref_new__2045356851
}

var cache_Effect_Ref_new__3544820218 gopurs_runtime.Value
var once_Effect_Ref_new__3544820218 sync.Once
func Get_Effect_Ref_new__3544820218() gopurs_runtime.Value {
	once_Effect_Ref_new__3544820218.Do(func() {
		cache_Effect_Ref_new__3544820218 = Get_Effect_Ref__new()
	})
	return cache_Effect_Ref_new__3544820218
}

var cache_Effect_Ref_new__1693026106 gopurs_runtime.Value
var once_Effect_Ref_new__1693026106 sync.Once
func Get_Effect_Ref_new__1693026106() gopurs_runtime.Value {
	once_Effect_Ref_new__1693026106.Do(func() {
		cache_Effect_Ref_new__1693026106 = Get_Effect_Ref__new()
	})
	return cache_Effect_Ref_new__1693026106
}

var cache_Effect_Ref_new__1017896474 gopurs_runtime.Value
var once_Effect_Ref_new__1017896474 sync.Once
func Get_Effect_Ref_new__1017896474() gopurs_runtime.Value {
	once_Effect_Ref_new__1017896474.Do(func() {
		cache_Effect_Ref_new__1017896474 = Get_Effect_Ref__new()
	})
	return cache_Effect_Ref_new__1017896474
}

var cache_Effect_Ref_new__337624346 gopurs_runtime.Value
var once_Effect_Ref_new__337624346 sync.Once
func Get_Effect_Ref_new__337624346() gopurs_runtime.Value {
	once_Effect_Ref_new__337624346.Do(func() {
		cache_Effect_Ref_new__337624346 = Get_Effect_Ref__new()
	})
	return cache_Effect_Ref_new__337624346
}

func Call_Effect_Ref_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Effect_Ref_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): s_prime_2_0 -> gopurs_runtime.Value
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Get_Effect_Ref__new() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref__New
}

func Get_Effect_Ref_modifyImpl() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_ModifyImpl
}

func Get_Effect_Ref_modify_() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_Modify_
}

func Get_Effect_Ref_newWithSelf() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_NewWithSelf
}

func Get_Effect_Ref_read() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_Read
}

func Get_Effect_Ref_write() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_Write
}
