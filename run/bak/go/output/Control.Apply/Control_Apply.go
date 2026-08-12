package Control_Apply

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_applyProxy
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

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__4046394764 gopurs_runtime.Value
var once_apply__4046394764 sync.Once
func Get_apply__4046394764() gopurs_runtime.Value {
	once_apply__4046394764.Do(func() {
		cache_apply__4046394764 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__4046394764(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__4046394764
}

var cache_apply__2021897708 gopurs_runtime.Value
var once_apply__2021897708 sync.Once
func Get_apply__2021897708() gopurs_runtime.Value {
	once_apply__2021897708.Do(func() {
		cache_apply__2021897708 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2021897708(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2021897708
}

var cache_apply__2908293516 gopurs_runtime.Value
var once_apply__2908293516 sync.Once
func Get_apply__2908293516() gopurs_runtime.Value {
	once_apply__2908293516.Do(func() {
		cache_apply__2908293516 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2908293516(gopurs_runtime.CoerceToStruct[Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2908293516
}

var cache_categoryFn__3492036198 gopurs_runtime.Value
var once_categoryFn__3492036198 sync.Once
func Get_categoryFn__3492036198() gopurs_runtime.Value {
	once_categoryFn__3492036198.Do(func() {
		cache_categoryFn__3492036198 = gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Semigroupoid.Get_semigroupoidFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_categoryFn__3492036198
}

var cache_identity__2527656589 gopurs_runtime.Value
var once_identity__2527656589 sync.Once
func Get_identity__2527656589() gopurs_runtime.Value {
	once_identity__2527656589.Do(func() {
		cache_identity__2527656589 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity__2527656589(dict_0_box)
})
	})
	return cache_identity__2527656589
}

var cache_semigroupoidFn__3002128382 gopurs_runtime.Value
var once_semigroupoidFn__3002128382 sync.Once
func Get_semigroupoidFn__3002128382() gopurs_runtime.Value {
	once_semigroupoidFn__3002128382.Do(func() {
		cache_semigroupoidFn__3002128382 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__3002128382
}

var cache_const__1496134642 gopurs_runtime.Value
var once_const__1496134642 sync.Once
func Get_const__1496134642() gopurs_runtime.Value {
	once_const__1496134642.Do(func() {
		cache_const__1496134642 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1496134642(a_0_box, v_1_box)
})
	})
	return cache_const__1496134642
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
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

var cache_functorProxy__711768561 gopurs_runtime.Value
var once_functorProxy__711768561 sync.Once
func Get_functorProxy__711768561() gopurs_runtime.Value {
	once_functorProxy__711768561.Do(func() {
		cache_functorProxy__711768561 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_functorProxy__711768561
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__2749506004 gopurs_runtime.Value
var once_map__2749506004 sync.Once
func Get_map__2749506004() gopurs_runtime.Value {
	once_map__2749506004.Do(func() {
		cache_map__2749506004 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2749506004(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2749506004
}

var cache_map__2458357236 gopurs_runtime.Value
var once_map__2458357236 sync.Once
func Get_map__2458357236() gopurs_runtime.Value {
	once_map__2458357236.Do(func() {
		cache_map__2458357236 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2458357236(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2458357236
}

var cache_map__3172880212 gopurs_runtime.Value
var once_map__3172880212 sync.Once
func Get_map__3172880212() gopurs_runtime.Value {
	once_map__3172880212.Do(func() {
		cache_map__3172880212 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3172880212(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3172880212
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

func Call_apply__353515660(dict_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__4046394764(dict_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2021897708(dict_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2908293516(dict_0_loop *Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_identity__2527656589(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "identity")
}

func Call_const__1496134642(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2749506004(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2458357236(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3172880212(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_arrayApply() gopurs_runtime.Value {
	return _Gopurs_ArrayApply
}
