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
return Call_identity(x_0_box)
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

var cache_discard__gopurs_runtime_Value_1429772317 gopurs_runtime.Value
var once_discard__gopurs_runtime_Value_1429772317 sync.Once
func Get_discard__gopurs_runtime_Value_1429772317() gopurs_runtime.Value {
	once_discard__gopurs_runtime_Value_1429772317.Do(func() {
		cache_discard__gopurs_runtime_Value_1429772317 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__gopurs_runtime_Value_1429772317(dict_0_box)
})
	})
	return cache_discard__gopurs_runtime_Value_1429772317
}

var cache_bindProxy gopurs_runtime.Value
var once_bindProxy sync.Once
func Get_bindProxy() gopurs_runtime.Value {
	once_bindProxy.Do(func() {
		cache_bindProxy = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}
})
}))
	})
	return cache_bindProxy
}

var cache_bindFn gopurs_runtime.Value
var once_bindFn sync.Once
func Get_bindFn() gopurs_runtime.Value {
	once_bindFn.Do(func() {
		cache_bindFn = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
}), gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(m_0, x_2), x_2)
})
})
}))
	})
	return cache_bindFn
}

var cache_bindArray gopurs_runtime.Value
var once_bindArray sync.Once
func Get_bindArray() gopurs_runtime.Value {
	once_bindArray.Do(func() {
		cache_bindArray = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), Get_arrayBind())
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

var cache_bind__gopurs_runtime_Value_152598582 gopurs_runtime.Value
var once_bind__gopurs_runtime_Value_152598582 sync.Once
func Get_bind__gopurs_runtime_Value_152598582() gopurs_runtime.Value {
	once_bind__gopurs_runtime_Value_152598582.Do(func() {
		cache_bind__gopurs_runtime_Value_152598582 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__gopurs_runtime_Value_152598582(dict_0_box)
})
	})
	return cache_bind__gopurs_runtime_Value_152598582
}

var cache_bind__gopurs_runtime_Value_1745568086 gopurs_runtime.Value
var once_bind__gopurs_runtime_Value_1745568086 sync.Once
func Get_bind__gopurs_runtime_Value_1745568086() gopurs_runtime.Value {
	once_bind__gopurs_runtime_Value_1745568086.Do(func() {
		cache_bind__gopurs_runtime_Value_1745568086 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__gopurs_runtime_Value_1745568086(dict_0_box)
})
	})
	return cache_bind__gopurs_runtime_Value_1745568086
}

var cache_bindFlipped gopurs_runtime.Value
var once_bindFlipped sync.Once
func Get_bindFlipped() gopurs_runtime.Value {
	once_bindFlipped.Do(func() {
		cache_bindFlipped = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped(dictBind_0_box, b_1_box, a_2_box)
})
	})
	return cache_bindFlipped
}

var cache_bindFlipped__gopurs_runtime_Value_3556144182 gopurs_runtime.Value
var once_bindFlipped__gopurs_runtime_Value_3556144182 sync.Once
func Get_bindFlipped__gopurs_runtime_Value_3556144182() gopurs_runtime.Value {
	once_bindFlipped__gopurs_runtime_Value_3556144182.Do(func() {
		cache_bindFlipped__gopurs_runtime_Value_3556144182 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__gopurs_runtime_Value_3556144182(dictBind_0_box, b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__gopurs_runtime_Value_3556144182
}

var cache_composeKleisliFlipped gopurs_runtime.Value
var once_composeKleisliFlipped sync.Once
func Get_composeKleisliFlipped() gopurs_runtime.Value {
	once_composeKleisliFlipped.Do(func() {
		cache_composeKleisliFlipped = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeKleisliFlipped(dictBind_0_box, f_1_box, g_2_box, a_3_box)
})
	})
	return cache_composeKleisliFlipped
}

var cache_composeKleisliFlipped__gopurs_runtime_Value_1891628823 gopurs_runtime.Value
var once_composeKleisliFlipped__gopurs_runtime_Value_1891628823 sync.Once
func Get_composeKleisliFlipped__gopurs_runtime_Value_1891628823() gopurs_runtime.Value {
	once_composeKleisliFlipped__gopurs_runtime_Value_1891628823.Do(func() {
		cache_composeKleisliFlipped__gopurs_runtime_Value_1891628823 = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeKleisliFlipped__gopurs_runtime_Value_1891628823(dictBind_0_box, f_1_box, g_2_box, a_3_box)
})
	})
	return cache_composeKleisliFlipped__gopurs_runtime_Value_1891628823
}

var cache_composeKleisli gopurs_runtime.Value
var once_composeKleisli sync.Once
func Get_composeKleisli() gopurs_runtime.Value {
	once_composeKleisli.Do(func() {
		cache_composeKleisli = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeKleisli(dictBind_0_box, f_1_box, g_2_box, a_3_box)
})
	})
	return cache_composeKleisli
}

var cache_discardProxy gopurs_runtime.Value
var once_discardProxy sync.Once
func Get_discardProxy() gopurs_runtime.Value {
	once_discardProxy.Do(func() {
		cache_discardProxy = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_discardProxy
}

var cache_discardUnit gopurs_runtime.Value
var once_discardUnit sync.Once
func Get_discardUnit() gopurs_runtime.Value {
	once_discardUnit.Do(func() {
		cache_discardUnit = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_discardUnit
}

var cache_ifM gopurs_runtime.Value
var once_ifM sync.Once
func Get_ifM() gopurs_runtime.Value {
	once_ifM.Do(func() {
		cache_ifM = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, cond_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ifM(dictBind_0_box, cond_1_box, t_2_box, f_3_box)
})
	})
	return cache_ifM
}

var cache_join gopurs_runtime.Value
var once_join sync.Once
func Get_join() gopurs_runtime.Value {
	once_join.Do(func() {
		cache_join = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_join(dictBind_0_box, m_1_box)
})
	})
	return cache_join
}

var cache_join__gopurs_runtime_Value_3225547319 gopurs_runtime.Value
var once_join__gopurs_runtime_Value_3225547319 sync.Once
func Get_join__gopurs_runtime_Value_3225547319() gopurs_runtime.Value {
	once_join__gopurs_runtime_Value_3225547319.Do(func() {
		cache_join__gopurs_runtime_Value_3225547319 = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_join__gopurs_runtime_Value_3225547319(dictBind_0_box, m_1_box)
})
	})
	return cache_join__gopurs_runtime_Value_3225547319
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_discard(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "discard")
}

func Call_discard__gopurs_runtime_Value_1429772317(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "discard")
}

func Call_bind(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bind")
}

func Call_bind__gopurs_runtime_Value_152598582(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bind")
}

func Call_bind__gopurs_runtime_Value_1745568086(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bind")
}

func Call_bindFlipped(dictBind_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), a_2, b_1)
}

func Call_bindFlipped__gopurs_runtime_Value_3556144182(dictBind_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), a_2, b_1)
}

func Call_composeKleisliFlipped(dictBind_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(g_2, a_3), f_1)
}

func Call_composeKleisliFlipped__gopurs_runtime_Value_1891628823(dictBind_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(g_2, a_3), f_1)
}

func Call_composeKleisli(dictBind_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(f_1, a_3), g_2)
}

func Call_ifM(dictBind_0_loop gopurs_runtime.Value, cond_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var cond_1 gopurs_runtime.Value = cond_1_loop
_ = cond_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), cond_1, gopurs_runtime.Func(func(cond_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (cond_prime_4.IntVal) != (0) {
__t0 = t_2
goto end_branch_0
} else {

}
}
{
__t0 = f_3
}
end_branch_0:
return __t0
}))
}

func Call_join(dictBind_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), m_1, Get_identity())
}

func Call_join__gopurs_runtime_Value_3225547319(dictBind_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), m_1, Get_identity())
}

func Get_arrayBind() gopurs_runtime.Value {
	return _Gopurs_ArrayBind
}
