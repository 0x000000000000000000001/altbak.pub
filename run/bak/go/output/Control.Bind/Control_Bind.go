package Control_Bind

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Apply "gopurs/output/Control.Apply"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_identity(gopurs_runtime.UnboxAny(x_0_box)))
})
	})
	return cache_identity
}

var cache_discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		cache_discard = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard(dict_0_box)
})
	})
	return cache_discard
}

var cache_discard__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____func_interface____interface____interface___3950241405 gopurs_runtime.Value
var once_discard__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____func_interface____interface____interface___3950241405 sync.Once
func Get_discard__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____func_interface____interface____interface___3950241405() gopurs_runtime.Value {
	once_discard__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____func_interface____interface____interface___3950241405.Do(func() {
		cache_discard__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____func_interface____interface____interface___3950241405 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____func_interface____interface____interface___3950241405(dict_0_box)
})
	})
	return cache_discard__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____func_interface____interface____interface___3950241405
}

var cache_bindProxy gopurs_runtime.Value
var once_bindProxy sync.Once
func Get_bindProxy() gopurs_runtime.Value {
	once_bindProxy.Do(func() {
		cache_bindProxy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})
}))))
	})
	return cache_bindProxy
}

var cache_bindFn gopurs_runtime.Value
var once_bindFn sync.Once
func Get_bindFn() gopurs_runtime.Value {
	once_bindFn.Do(func() {
		cache_bindFn = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
}), gopurs_runtime.Func3(func(m_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(m_0, x_2), x_2)
}))))
	})
	return cache_bindFn
}

var cache_bindArray gopurs_runtime.Value
var once_bindArray sync.Once
func Get_bindArray() gopurs_runtime.Value {
	once_bindArray.Do(func() {
		cache_bindArray = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), Get_arrayBind())))
	})
	return cache_bindArray
}

var cache_bind gopurs_runtime.Value
var once_bind sync.Once
func Get_bind() gopurs_runtime.Value {
	once_bind.Do(func() {
		cache_bind = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind(dict_0_box)
})
	})
	return cache_bind
}

var cache_bind__func_gopurs_runtime_Value__func_interface____interface____func_interface____func_interface____interface____func_interface____interface___470494006 gopurs_runtime.Value
var once_bind__func_gopurs_runtime_Value__func_interface____interface____func_interface____func_interface____interface____func_interface____interface___470494006 sync.Once
func Get_bind__func_gopurs_runtime_Value__func_interface____interface____func_interface____func_interface____interface____func_interface____interface___470494006() gopurs_runtime.Value {
	once_bind__func_gopurs_runtime_Value__func_interface____interface____func_interface____func_interface____interface____func_interface____interface___470494006.Do(func() {
		cache_bind__func_gopurs_runtime_Value__func_interface____interface____func_interface____func_interface____interface____func_interface____interface___470494006 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__func_gopurs_runtime_Value__func_interface____interface____func_interface____func_interface____interface____func_interface____interface___470494006(dict_0_box)
})
	})
	return cache_bind__func_gopurs_runtime_Value__func_interface____interface____func_interface____func_interface____interface____func_interface____interface___470494006
}

var cache_bind__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2164513878 gopurs_runtime.Value
var once_bind__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2164513878 sync.Once
func Get_bind__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2164513878() gopurs_runtime.Value {
	once_bind__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2164513878.Do(func() {
		cache_bind__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2164513878 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2164513878(dict_0_box)
})
	})
	return cache_bind__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2164513878
}

var cache_bindFlipped gopurs_runtime.Value
var once_bindFlipped sync.Once
func Get_bindFlipped() gopurs_runtime.Value {
	once_bindFlipped.Do(func() {
		cache_bindFlipped = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindFlipped(dictBind_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(b_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_bindFlipped
}

var cache_bindFlipped__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2292006102 gopurs_runtime.Value
var once_bindFlipped__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2292006102 sync.Once
func Get_bindFlipped__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2292006102() gopurs_runtime.Value {
	once_bindFlipped__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2292006102.Do(func() {
		cache_bindFlipped__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2292006102 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindFlipped__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2292006102(dictBind_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(b_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_bindFlipped__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2292006102
}

var cache_composeKleisliFlipped gopurs_runtime.Value
var once_composeKleisliFlipped sync.Once
func Get_composeKleisliFlipped() gopurs_runtime.Value {
	once_composeKleisliFlipped.Do(func() {
		cache_composeKleisliFlipped = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_composeKleisliFlipped(dictBind_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(a_3_box)))
})
	})
	return cache_composeKleisliFlipped
}

var cache_composeKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___477960372 gopurs_runtime.Value
var once_composeKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___477960372 sync.Once
func Get_composeKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___477960372() gopurs_runtime.Value {
	once_composeKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___477960372.Do(func() {
		cache_composeKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___477960372 = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_composeKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___477960372(dictBind_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(a_3_box)))
})
	})
	return cache_composeKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___477960372
}

var cache_composeKleisli gopurs_runtime.Value
var once_composeKleisli sync.Once
func Get_composeKleisli() gopurs_runtime.Value {
	once_composeKleisli.Do(func() {
		cache_composeKleisli = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_composeKleisli(dictBind_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(a_3_box)))
})
	})
	return cache_composeKleisli
}

var cache_discardProxy gopurs_runtime.Value
var once_discardProxy sync.Once
func Get_discardProxy() gopurs_runtime.Value {
	once_discardProxy.Do(func() {
		cache_discardProxy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))))
	})
	return cache_discardProxy
}

var cache_discardUnit gopurs_runtime.Value
var once_discardUnit sync.Once
func Get_discardUnit() gopurs_runtime.Value {
	once_discardUnit.Do(func() {
		cache_discardUnit = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))))
	})
	return cache_discardUnit
}

var cache_ifM gopurs_runtime.Value
var once_ifM sync.Once
func Get_ifM() gopurs_runtime.Value {
	once_ifM.Do(func() {
		cache_ifM = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, cond_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_ifM(dictBind_0_box, gopurs_runtime.UnboxAny(cond_1_box), gopurs_runtime.UnboxAny(t_2_box), gopurs_runtime.UnboxAny(f_3_box)))
})
	})
	return cache_ifM
}

var cache_join gopurs_runtime.Value
var once_join sync.Once
func Get_join() gopurs_runtime.Value {
	once_join.Do(func() {
		cache_join = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_join(dictBind_0_box, gopurs_runtime.UnboxAny(m_1_box)))
})
	})
	return cache_join
}

var cache_join__func_gopurs_runtime_Value__interface____interface___803698843 gopurs_runtime.Value
var once_join__func_gopurs_runtime_Value__interface____interface___803698843 sync.Once
func Get_join__func_gopurs_runtime_Value__interface____interface___803698843() gopurs_runtime.Value {
	once_join__func_gopurs_runtime_Value__interface____interface___803698843.Do(func() {
		cache_join__func_gopurs_runtime_Value__interface____interface___803698843 = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_join__func_gopurs_runtime_Value__interface____interface___803698843(dictBind_0_box, gopurs_runtime.UnboxAny(m_1_box)))
})
	})
	return cache_join__func_gopurs_runtime_Value__interface____interface___803698843
}

var cache_arrayBind gopurs_runtime.Value
var once_arrayBind sync.Once
func Get_arrayBind() gopurs_runtime.Value {
	once_arrayBind.Do(func() {
		cache_arrayBind = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := ArrayBind(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func(inner_arg0 interface{}) []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(arg1, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
})
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_arrayBind
}

func Call_identity(x_0_loop interface{}) interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
return x_0
}

func Call_discard(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "discard")
}

func Call_discard__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____func_interface____interface____interface___3950241405(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "discard")
}

func Call_bind(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bind")
}

func Call_bind__func_gopurs_runtime_Value__func_interface____interface____func_interface____func_interface____interface____func_interface____interface___470494006(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bind")
}

func Call_bind__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2164513878(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bind")
}

func Call_bindFlipped(dictBind_0_loop gopurs_runtime.Value, b_1_loop func(interface{}) interface{}, a_2_loop interface{}) interface{} {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var b_1 func(interface{}) interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Any(a_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(b_1(gopurs_runtime.UnboxAny(arg0)))
})))
}

func Call_bindFlipped__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2292006102(dictBind_0_loop gopurs_runtime.Value, b_1_loop func(interface{}) interface{}, a_2_loop interface{}) interface{} {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var b_1 func(interface{}) interface{} = b_1_loop
_ = b_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Any(a_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(b_1(gopurs_runtime.UnboxAny(arg0)))
})))
}

func Call_composeKleisliFlipped(dictBind_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, g_2_loop func(interface{}) interface{}, a_3_loop interface{}) interface{} {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var g_2 func(interface{}) interface{} = g_2_loop
_ = g_2
var a_3 interface{} = a_3_loop
_ = a_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Any(g_2(a_3)), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0)))
})))
}

func Call_composeKleisliFlipped__func_gopurs_runtime_Value__func_interface____interface____func_interface____interface____interface____interface___477960372(dictBind_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, g_2_loop func(interface{}) interface{}, a_3_loop interface{}) interface{} {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var g_2 func(interface{}) interface{} = g_2_loop
_ = g_2
var a_3 interface{} = a_3_loop
_ = a_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Any(g_2(a_3)), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0)))
})))
}

func Call_composeKleisli(dictBind_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, g_2_loop func(interface{}) interface{}, a_3_loop interface{}) interface{} {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var g_2 func(interface{}) interface{} = g_2_loop
_ = g_2
var a_3 interface{} = a_3_loop
_ = a_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Any(f_1(a_3)), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(g_2(gopurs_runtime.UnboxAny(arg0)))
})))
}

func Call_ifM(dictBind_0_loop gopurs_runtime.Value, cond_1_loop interface{}, t_2_loop interface{}, f_3_loop interface{}) interface{} {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var cond_1 interface{} = cond_1_loop
_ = cond_1
var t_2 interface{} = t_2_loop
_ = t_2
var f_3 interface{} = f_3_loop
_ = f_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Any(cond_1), gopurs_runtime.Func(func(cond_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (cond_prime_4.IntVal) != (0) {
__t0 = gopurs_runtime.Any(t_2)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Any(f_3)
}
end_branch_0:
return __t0
})))
}

func Call_join(dictBind_0_loop gopurs_runtime.Value, m_1_loop interface{}) interface{} {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var m_1 interface{} = m_1_loop
_ = m_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Any(m_1), Get_identity()))
}

func Call_join__func_gopurs_runtime_Value__interface____interface___803698843(dictBind_0_loop gopurs_runtime.Value, m_1_loop interface{}) interface{} {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var m_1 interface{} = m_1_loop
_ = m_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Any(m_1), Get_identity()))
}
