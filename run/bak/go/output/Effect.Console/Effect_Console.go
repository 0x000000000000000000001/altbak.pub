package Effect_Console

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_warnShow gopurs_runtime.Value
var once_warnShow sync.Once
func Get_warnShow() gopurs_runtime.Value {
	once_warnShow.Do(func() {
		cache_warnShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_warnShow(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
})
	})
	return cache_warnShow
}

var cache_logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		cache_logShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
})
	})
	return cache_logShow
}

var cache_infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		cache_infoShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_infoShow(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
})
	})
	return cache_infoShow
}

var cache_grouped gopurs_runtime.Value
var once_grouped sync.Once
func Get_grouped() gopurs_runtime.Value {
	once_grouped.Do(func() {
		cache_grouped = gopurs_runtime.Func2(func(name_0_box gopurs_runtime.Value, inner_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_grouped(name_0_box.StrVal(), inner_1_box)
})
	})
	return cache_grouped
}

var cache_errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		cache_errorShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_errorShow(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
})
	})
	return cache_errorShow
}

var cache_debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		cache_debugShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_debugShow(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
})
	})
	return cache_debugShow
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__1641029622 gopurs_runtime.Value
var once_pure__1641029622 sync.Once
func Get_pure__1641029622() gopurs_runtime.Value {
	once_pure__1641029622.Do(func() {
		cache_pure__1641029622 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1641029622(__eta0_0_box)
})
	})
	return cache_pure__1641029622
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__1598467489 gopurs_runtime.Value
var once_bind__1598467489 sync.Once
func Get_bind__1598467489() gopurs_runtime.Value {
	once_bind__1598467489.Do(func() {
		cache_bind__1598467489 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1598467489(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bind__1598467489
}

var cache_discard__1979268384 gopurs_runtime.Value
var once_discard__1979268384 sync.Once
func Get_discard__1979268384() gopurs_runtime.Value {
	once_discard__1979268384.Do(func() {
		cache_discard__1979268384 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__1979268384(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_discard__1979268384
}

var cache_discard__317162198 gopurs_runtime.Value
var once_discard__317162198 sync.Once
func Get_discard__317162198() gopurs_runtime.Value {
	once_discard__317162198.Do(func() {
		cache_discard__317162198 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__317162198(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_discard__317162198
}

var cache_discardUnit__2687062302 gopurs_runtime.Value
var once_discardUnit__2687062302 sync.Once
func Get_discardUnit__2687062302() gopurs_runtime.Value {
	once_discardUnit__2687062302.Do(func() {
		cache_discardUnit__2687062302 = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_discardUnit__2687062302
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_applicativeEffect__284161122 gopurs_runtime.Value
var once_applicativeEffect__284161122 sync.Once
func Get_applicativeEffect__284161122() gopurs_runtime.Value {
	once_applicativeEffect__284161122.Do(func() {
		cache_applicativeEffect__284161122 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_pureE())
	})
	return cache_applicativeEffect__284161122
}

var cache_applyEffect__2014400020 gopurs_runtime.Value
var once_applyEffect__2014400020 sync.Once
func Get_applyEffect__2014400020() gopurs_runtime.Value {
	once_applyEffect__2014400020.Do(func() {
		cache_applyEffect__2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__2014400020
}

var cache_bindEffect__2113658466 gopurs_runtime.Value
var once_bindEffect__2113658466 sync.Once
func Get_bindEffect__2113658466() gopurs_runtime.Value {
	once_bindEffect__2113658466.Do(func() {
		cache_bindEffect__2113658466 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_bindE())
	})
	return cache_bindEffect__2113658466
}

var cache_functorEffect__3107547953 gopurs_runtime.Value
var once_functorEffect__3107547953 sync.Once
func Get_functorEffect__3107547953() gopurs_runtime.Value {
	once_functorEffect__3107547953.Do(func() {
		cache_functorEffect__3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__3107547953
}

func Call_warnShow(dictShow_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_warn(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_0.V0, a_1).StrVal()))
}

func Call_logShow(dictShow_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_0.V0, a_1).StrVal()))
}

func Call_infoShow(dictShow_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_info(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_0.V0, a_1).StrVal()))
}

func Call_grouped(name_0_loop string, inner_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var name_0 string = name_0_loop
_ = name_0
var inner_1 gopurs_runtime.Value = inner_1_loop
_ = inner_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Effect.Get_bindEffect()))}, gopurs_runtime.Apply(Get_group(), gopurs_runtime.Str(name_0)), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), inner_1, gopurs_runtime.Func(func(result_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Effect.Get_bindEffect()))}, Get_groupEnd(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), result_3)
}))
}))
}))
}

func Call_errorShow(dictShow_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_error(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_0.V0, a_1).StrVal()))
}

func Call_debugShow(dictShow_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_debug(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_0.V0, a_1).StrVal()))
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1641029622(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1598467489(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_discard__1979268384(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), __eta0_0, __eta1_1)
}

func Call_discard__317162198(dict_0_loop *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_clear() gopurs_runtime.Value {
	return _Gopurs_Clear
}

func Get_debug() gopurs_runtime.Value {
	return _Gopurs_Debug
}

func Get_error() gopurs_runtime.Value {
	return _Gopurs_Error
}

func Get_group() gopurs_runtime.Value {
	return _Gopurs_Group
}

func Get_groupCollapsed() gopurs_runtime.Value {
	return _Gopurs_GroupCollapsed
}

func Get_groupEnd() gopurs_runtime.Value {
	return _Gopurs_GroupEnd
}

func Get_info() gopurs_runtime.Value {
	return _Gopurs_Info
}

func Get_log() gopurs_runtime.Value {
	return _Gopurs_Log
}

func Get_time() gopurs_runtime.Value {
	return _Gopurs_Time
}

func Get_timeEnd() gopurs_runtime.Value {
	return _Gopurs_TimeEnd
}

func Get_timeLog() gopurs_runtime.Value {
	return _Gopurs_TimeLog
}

func Get_warn() gopurs_runtime.Value {
	return _Gopurs_Warn
}
