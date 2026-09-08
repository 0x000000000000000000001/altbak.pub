package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Ref_go__new gopurs_runtime.Value
var once_Effect_Ref_go__new sync.Once
func Get_Effect_Ref_go__new() gopurs_runtime.Value {
	once_Effect_Ref_go__new.Do(func() {
		cache_Effect_Ref_go__new = Get_Effect_Ref__new()
	})
	return cache_Effect_Ref_go__new
}

var cache_Effect_Ref_new__2291742540 gopurs_runtime.Value
var once_Effect_Ref_new__2291742540 sync.Once
func Get_Effect_Ref_new__2291742540() gopurs_runtime.Value {
	once_Effect_Ref_new__2291742540.Do(func() {
		cache_Effect_Ref_new__2291742540 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Ref_new__2291742540((__eta_norm_0_0_box.IntVal) != (0))
})
	})
	return cache_Effect_Ref_new__2291742540
}

var cache_Effect_Ref_modify_prime_ gopurs_runtime.Value
var once_Effect_Ref_modify_prime_ sync.Once
func Get_Effect_Ref_modify_prime_() gopurs_runtime.Value {
	once_Effect_Ref_modify_prime_.Do(func() {
		cache_Effect_Ref_modify_prime_ = Get_Effect_Ref_modifyImpl()
	})
	return cache_Effect_Ref_modify_prime_
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

func Call_Effect_Ref_new__2291742540(__eta_norm_0_0_loop bool) gopurs_runtime.Value {
new__2291742540:
for {
if false { continue new__2291742540 }
var __eta_norm_0_0 bool = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Bool(__eta_norm_0_0))
}
}

func Call_Effect_Ref_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Effect_Ref_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): s_prime__2_0 shape=App(Other) bindingType=(TypeVar s)
s_prime__2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime__2_0
return func() gopurs_runtime.Value {
				orig := struct{
	state gopurs_runtime.Value
	value gopurs_runtime.Value
}{s_prime__2_0, s_prime__2_0}
				_ = orig
				return gopurs_runtime.RecordDict2("state", "value", orig.state, orig.value)
				}()
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
