package Control_Apply

import (
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_applyProxy gopurs_runtime.Value
var once_applyProxy sync.Once
func Get_applyProxy() gopurs_runtime.Value {
	once_applyProxy.Do(func() {
		cache_applyProxy = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_applyProxy
}

var cache_applyProxy__gopurs_runtime_Value_315643445 gopurs_runtime.Value
var once_applyProxy__gopurs_runtime_Value_315643445 sync.Once
func Get_applyProxy__gopurs_runtime_Value_315643445() gopurs_runtime.Value {
	once_applyProxy__gopurs_runtime_Value_315643445.Do(func() {
		cache_applyProxy__gopurs_runtime_Value_315643445 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorProxy__gopurs_runtime_Value_711768561()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_applyProxy__gopurs_runtime_Value_315643445
}

var cache_applyFn gopurs_runtime.Value
var once_applyFn sync.Once
func Get_applyFn() gopurs_runtime.Value {
	once_applyFn.Do(func() {
		cache_applyFn = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_applyFn
}

var cache_applyFn__ptrConstructor_Apply_gopurs_runtime_Value__2722791663 gopurs_runtime.Value
var once_applyFn__ptrConstructor_Apply_gopurs_runtime_Value__2722791663 sync.Once
func Get_applyFn__ptrConstructor_Apply_gopurs_runtime_Value__2722791663() gopurs_runtime.Value {
	once_applyFn__ptrConstructor_Apply_gopurs_runtime_Value__2722791663.Do(func() {
		cache_applyFn__ptrConstructor_Apply_gopurs_runtime_Value__2722791663 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
})})}
	})
	return cache_applyFn__ptrConstructor_Apply_gopurs_runtime_Value__2722791663
}

var cache_applyFn__gopurs_runtime_Value_4042184691 gopurs_runtime.Value
var once_applyFn__gopurs_runtime_Value_4042184691 sync.Once
func Get_applyFn__gopurs_runtime_Value_4042184691() gopurs_runtime.Value {
	once_applyFn__gopurs_runtime_Value_4042184691.Do(func() {
		cache_applyFn__gopurs_runtime_Value_4042184691 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn__gopurs_runtime_Value_20325936()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_applyFn__gopurs_runtime_Value_4042184691
}

var cache_applyArray gopurs_runtime.Value
var once_applyArray sync.Once
func Get_applyArray() gopurs_runtime.Value {
	once_applyArray.Do(func() {
		cache_applyArray = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), Get_arrayApply())
	})
	return cache_applyArray
}

var cache_applyArray__gopurs_runtime_Value_2998472828 gopurs_runtime.Value
var once_applyArray__gopurs_runtime_Value_2998472828 sync.Once
func Get_applyArray__gopurs_runtime_Value_2998472828() gopurs_runtime.Value {
	once_applyArray__gopurs_runtime_Value_2998472828.Do(func() {
		cache_applyArray__gopurs_runtime_Value_2998472828 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray__gopurs_runtime_Value_361387505()
}), Get_arrayApply())
	})
	return cache_applyArray__gopurs_runtime_Value_2998472828
}

var cache_apply gopurs_runtime.Value
var once_apply sync.Once
func Get_apply() gopurs_runtime.Value {
	once_apply.Do(func() {
		cache_apply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply
}

var cache_apply__gopurs_runtime_Value_353515660 gopurs_runtime.Value
var once_apply__gopurs_runtime_Value_353515660 sync.Once
func Get_apply__gopurs_runtime_Value_353515660() gopurs_runtime.Value {
	once_apply__gopurs_runtime_Value_353515660.Do(func() {
		cache_apply__gopurs_runtime_Value_353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__gopurs_runtime_Value_353515660(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__gopurs_runtime_Value_353515660
}

var cache_applyFirst gopurs_runtime.Value
var once_applyFirst sync.Once
func Get_applyFirst() gopurs_runtime.Value {
	once_applyFirst.Do(func() {
		cache_applyFirst = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyFirst(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_applyFirst
}

var cache_applySecond gopurs_runtime.Value
var once_applySecond sync.Once
func Get_applySecond() gopurs_runtime.Value {
	once_applySecond.Do(func() {
		cache_applySecond = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_applySecond
}

var cache_applySecond__gopurs_runtime_Value_1627424644 gopurs_runtime.Value
var once_applySecond__gopurs_runtime_Value_1627424644 sync.Once
func Get_applySecond__gopurs_runtime_Value_1627424644() gopurs_runtime.Value {
	once_applySecond__gopurs_runtime_Value_1627424644.Do(func() {
		cache_applySecond__gopurs_runtime_Value_1627424644 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applySecond__gopurs_runtime_Value_1627424644(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_applySecond__gopurs_runtime_Value_1627424644
}

var cache_lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		cache_lift2 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2
}

var cache_lift2__gopurs_runtime_Value_2762258480 gopurs_runtime.Value
var once_lift2__gopurs_runtime_Value_2762258480 sync.Once
func Get_lift2__gopurs_runtime_Value_2762258480() gopurs_runtime.Value {
	once_lift2__gopurs_runtime_Value_2762258480.Do(func() {
		cache_lift2__gopurs_runtime_Value_2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__gopurs_runtime_Value_2762258480(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__gopurs_runtime_Value_2762258480
}

var cache_lift3 gopurs_runtime.Value
var once_lift3 sync.Once
func Get_lift3() gopurs_runtime.Value {
	once_lift3.Do(func() {
		cache_lift3 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift3(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift3
}

var cache_lift4 gopurs_runtime.Value
var once_lift4 sync.Once
func Get_lift4() gopurs_runtime.Value {
	once_lift4.Do(func() {
		cache_lift4 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift4(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift4
}

var cache_lift5 gopurs_runtime.Value
var once_lift5 sync.Once
func Get_lift5() gopurs_runtime.Value {
	once_lift5.Do(func() {
		cache_lift5 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift5(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift5
}

type Constructor_Apply[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3032403085] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Apply[gopurs_runtime.Value])(ptr)
		switch key {
		case "Functor0": return c.V0
		case "apply": return c.V1
		default: panic("Key not found in dictionary Constructor_Apply: " + key)
		}
	}
}


func Call_apply(dict_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__gopurs_runtime_Value_353515660(dict_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_applyFirst(dictApply_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Function.Get_go__const(), a_2), b_3)
})
})
}

func Call_applySecond(dictApply_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), a_2), b_3)
})
})
}

func Call_applySecond__gopurs_runtime_Value_1627424644(dictApply_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), a_2), b_3)
})
})
}

func Call_lift2(dictApply_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__gopurs_runtime_Value_2762258480(dictApply_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift3(dictApply_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4), c_5)
})
})
})
})
}

func Call_lift4(dictApply_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4), c_5), d_6)
})
})
})
})
})
}

func Call_lift5(dictApply_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4), c_5), d_6), e_7)
})
})
})
})
})
})
}

func Get_arrayApply() gopurs_runtime.Value {
	return _Gopurs_ArrayApply
}
