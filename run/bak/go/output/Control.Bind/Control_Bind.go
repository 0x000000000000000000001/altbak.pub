package Control_Bind

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
return Call_discard(gopurs_runtime.CoerceToStruct[Constructor_Discard[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_discard
}

var cache_bindProxy gopurs_runtime.Value
var once_bindProxy sync.Once
func Get_bindProxy() gopurs_runtime.Value {
	once_bindProxy.Do(func() {
		cache_bindProxy = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
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
return Call_bind(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind
}

var cache_bindFlipped gopurs_runtime.Value
var once_bindFlipped sync.Once
func Get_bindFlipped() gopurs_runtime.Value {
	once_bindFlipped.Do(func() {
		cache_bindFlipped = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped
}

var cache_composeKleisliFlipped gopurs_runtime.Value
var once_composeKleisliFlipped sync.Once
func Get_composeKleisliFlipped() gopurs_runtime.Value {
	once_composeKleisliFlipped.Do(func() {
		cache_composeKleisliFlipped = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeKleisliFlipped(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), f_1_box, g_2_box, a_3_box)
})
	})
	return cache_composeKleisliFlipped
}

var cache_composeKleisli gopurs_runtime.Value
var once_composeKleisli sync.Once
func Get_composeKleisli() gopurs_runtime.Value {
	once_composeKleisli.Do(func() {
		cache_composeKleisli = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeKleisli(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), f_1_box, g_2_box, a_3_box)
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
return Call_ifM(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), cond_1_box, t_2_box, f_3_box)
})
	})
	return cache_ifM
}

var cache_join gopurs_runtime.Value
var once_join sync.Once
func Get_join() gopurs_runtime.Value {
	once_join.Do(func() {
		cache_join = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_join(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), m_1_box)
})
	})
	return cache_join
}

var cache_applyArray__2998472828 gopurs_runtime.Value
var once_applyArray__2998472828 sync.Once
func Get_applyArray__2998472828() gopurs_runtime.Value {
	once_applyArray__2998472828.Do(func() {
		cache_applyArray__2998472828 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Control_Apply.Get_arrayApply())
	})
	return cache_applyArray__2998472828
}

var cache_applyFn__4042184691 gopurs_runtime.Value
var once_applyFn__4042184691 sync.Once
func Get_applyFn__4042184691() gopurs_runtime.Value {
	once_applyFn__4042184691.Do(func() {
		cache_applyFn__4042184691 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_applyFn__4042184691
}

var cache_applyProxy__315643445 gopurs_runtime.Value
var once_applyProxy__315643445 sync.Once
func Get_applyProxy__315643445() gopurs_runtime.Value {
	once_applyProxy__315643445.Do(func() {
		cache_applyProxy__315643445 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_applyProxy__315643445
}

var cache_bind__4146772295 gopurs_runtime.Value
var once_bind__4146772295 sync.Once
func Get_bind__4146772295() gopurs_runtime.Value {
	once_bind__4146772295.Do(func() {
		cache_bind__4146772295 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__4146772295(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__4146772295
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__737692327 gopurs_runtime.Value
var once_bind__737692327 sync.Once
func Get_bind__737692327() gopurs_runtime.Value {
	once_bind__737692327.Do(func() {
		cache_bind__737692327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__737692327(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__737692327
}

var cache_bindFlipped__1485397639 gopurs_runtime.Value
var once_bindFlipped__1485397639 sync.Once
func Get_bindFlipped__1485397639() gopurs_runtime.Value {
	once_bindFlipped__1485397639.Do(func() {
		cache_bindFlipped__1485397639 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1485397639(gopurs_runtime.CoerceToStruct[Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__1485397639
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__2253242624 gopurs_runtime.Value
var once_flip__2253242624 sync.Once
func Get_flip__2253242624() gopurs_runtime.Value {
	once_flip__2253242624.Do(func() {
		cache_flip__2253242624 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2253242624(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2253242624
}

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
}

var cache_functorFn__20325936 gopurs_runtime.Value
var once_functorFn__20325936 sync.Once
func Get_functorFn__20325936() gopurs_runtime.Value {
	once_functorFn__20325936.Do(func() {
		cache_functorFn__20325936 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn__20325936
}

var cache_functorProxy__1157108209 gopurs_runtime.Value
var once_functorProxy__1157108209 sync.Once
func Get_functorProxy__1157108209() gopurs_runtime.Value {
	once_functorProxy__1157108209.Do(func() {
		cache_functorProxy__1157108209 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_functorProxy__1157108209
}

type Constructor_Bind[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4032919565] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Bind[gopurs_runtime.Value])(ptr)
		switch key {
		case "Apply0": return c.V0
		case "bind": return c.V1
		default: panic("Key not found in dictionary Constructor_Bind: " + key)
		}
	}
}


type Constructor_Discard[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2260728934] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Discard[gopurs_runtime.Value])(ptr)
		switch key {
		case "discard": return c.V0
		default: panic("Key not found in dictionary Constructor_Discard: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_discard(dict_0_loop *Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bind(dict_0_loop *Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bindFlipped(dictBind_0_loop *Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_composeKleisliFlipped(dictBind_0_loop *Constructor_Bind[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(dictBind_0.V1, gopurs_runtime.Apply(g_2, a_3), f_1)
}

func Call_composeKleisli(dictBind_0_loop *Constructor_Bind[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(dictBind_0.V1, gopurs_runtime.Apply(f_1, a_3), g_2)
}

func Call_ifM(dictBind_0_loop *Constructor_Bind[gopurs_runtime.Value], cond_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var cond_1 gopurs_runtime.Value = cond_1_loop
_ = cond_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
return gopurs_runtime.Apply2(dictBind_0.V1, cond_1, gopurs_runtime.Func(func(cond_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_join(dictBind_0_loop *Constructor_Bind[gopurs_runtime.Value], m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
return gopurs_runtime.Apply2(dictBind_0.V1, m_1, Get_identity())
}

func Call_bind__4146772295(dict_0_loop *Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__737692327(dict_0_loop *Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bindFlipped__1485397639(dictBind_0_loop *Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2253242624(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Get_arrayBind() gopurs_runtime.Value {
	return _Gopurs_ArrayBind
}
