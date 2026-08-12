package Effect_Class_Console

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect_Class "gopurs/output/Effect.Class"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_warnShow gopurs_runtime.Value
var once_warnShow sync.Once
func Get_warnShow() gopurs_runtime.Value {
	once_warnShow.Do(func() {
		cache_warnShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_warnShow(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_1_box), x_2_box)
})
	})
	return cache_warnShow
}

var cache_warn gopurs_runtime.Value
var once_warn sync.Once
func Get_warn() gopurs_runtime.Value {
	once_warn.Do(func() {
		cache_warn = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_warn(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_warn
}

var cache_timeLog gopurs_runtime.Value
var once_timeLog sync.Once
func Get_timeLog() gopurs_runtime.Value {
	once_timeLog.Do(func() {
		cache_timeLog = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_timeLog(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_timeLog
}

var cache_timeEnd gopurs_runtime.Value
var once_timeEnd sync.Once
func Get_timeEnd() gopurs_runtime.Value {
	once_timeEnd.Do(func() {
		cache_timeEnd = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_timeEnd(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_timeEnd
}

var cache_time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		cache_time = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_time(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_time
}

var cache_logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		cache_logShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_1_box), x_2_box)
})
	})
	return cache_logShow
}

var cache_log gopurs_runtime.Value
var once_log sync.Once
func Get_log() gopurs_runtime.Value {
	once_log.Do(func() {
		cache_log = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_log(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_log
}

var cache_infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		cache_infoShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_infoShow(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_1_box), x_2_box)
})
	})
	return cache_infoShow
}

var cache_info gopurs_runtime.Value
var once_info sync.Once
func Get_info() gopurs_runtime.Value {
	once_info.Do(func() {
		cache_info = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_info(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_info
}

var cache_groupEnd gopurs_runtime.Value
var once_groupEnd sync.Once
func Get_groupEnd() gopurs_runtime.Value {
	once_groupEnd.Do(func() {
		cache_groupEnd = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupEnd(dictMonadEffect_0_box)
})
	})
	return cache_groupEnd
}

var cache_groupCollapsed gopurs_runtime.Value
var once_groupCollapsed sync.Once
func Get_groupCollapsed() gopurs_runtime.Value {
	once_groupCollapsed.Do(func() {
		cache_groupCollapsed = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupCollapsed(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_groupCollapsed
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_group
}

var cache_grouped gopurs_runtime.Value
var once_grouped sync.Once
func Get_grouped() gopurs_runtime.Value {
	once_grouped.Do(func() {
		cache_grouped = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_grouped(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box))
})
	})
	return cache_grouped
}

var cache_errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		cache_errorShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_errorShow(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_1_box), x_2_box)
})
	})
	return cache_errorShow
}

var cache_error gopurs_runtime.Value
var once_error sync.Once
func Get_error() gopurs_runtime.Value {
	once_error.Do(func() {
		cache_error = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_error(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_error
}

var cache_debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		cache_debugShow = gopurs_runtime.Func3(func(dictMonadEffect_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_debugShow(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_1_box), x_2_box)
})
	})
	return cache_debugShow
}

var cache_debug gopurs_runtime.Value
var once_debug sync.Once
func Get_debug() gopurs_runtime.Value {
	once_debug.Do(func() {
		cache_debug = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_debug(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_debug
}

var cache_clear gopurs_runtime.Value
var once_clear sync.Once
func Get_clear() gopurs_runtime.Value {
	once_clear.Do(func() {
		cache_clear = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_clear(dictMonadEffect_0_box)
})
	})
	return cache_clear
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

var cache_discard__2561459590 gopurs_runtime.Value
var once_discard__2561459590 sync.Once
func Get_discard__2561459590() gopurs_runtime.Value {
	once_discard__2561459590.Do(func() {
		cache_discard__2561459590 = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__2561459590(dictBind_0_box)
})
	})
	return cache_discard__2561459590
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

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_group__227222742 gopurs_runtime.Value
var once_group__227222742 sync.Once
func Get_group__227222742() gopurs_runtime.Value {
	once_group__227222742.Do(func() {
		cache_group__227222742 = gopurs_runtime.Func2(func(dictMonadEffect_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group__227222742(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0_box), x_1_box.StrVal())
})
	})
	return cache_group__227222742
}

var cache_liftEffect__3456588885 gopurs_runtime.Value
var once_liftEffect__3456588885 sync.Once
func Get_liftEffect__3456588885() gopurs_runtime.Value {
	once_liftEffect__3456588885.Do(func() {
		cache_liftEffect__3456588885 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__3456588885(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__3456588885
}

var cache_liftEffect__1892566677 gopurs_runtime.Value
var once_liftEffect__1892566677 sync.Once
func Get_liftEffect__1892566677() gopurs_runtime.Value {
	once_liftEffect__1892566677.Do(func() {
		cache_liftEffect__1892566677 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__1892566677(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__1892566677
}

func Call_warnShow(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], dictShow_1_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_warn(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_1.V0, x_2).StrVal())))
}

func Call_warn(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_warn(), gopurs_runtime.Str(x_1)))
}

func Call_timeLog(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_timeLog(), gopurs_runtime.Str(x_1)))
}

func Call_timeEnd(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_timeEnd(), gopurs_runtime.Str(x_1)))
}

func Call_time(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_time(), gopurs_runtime.Str(x_1)))
}

func Call_logShow(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], dictShow_1_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_1.V0, x_2).StrVal())))
}

func Call_log(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(x_1)))
}

func Call_infoShow(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], dictShow_1_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_info(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_1.V0, x_2).StrVal())))
}

func Call_info(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_info(), gopurs_runtime.Str(x_1)))
}

func Call_groupEnd(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), pkg_Effect_Console.Get_groupEnd())
}

func Call_groupCollapsed(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_groupCollapsed(), gopurs_runtime.Str(x_1)))
}

func Call_group(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_group(), gopurs_runtime.Str(x_1)))
}

func Call_grouped(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadEffect_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
groupEnd1_3_2 := gopurs_runtime.Apply(dictMonadEffect_0.V1, pkg_Effect_Console.Get_groupEnd())
_ = groupEnd1_3_2
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.Func(func(name_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(inner_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Bind1_2_1)}, gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_group(), gopurs_runtime.Str(name_5.StrVal()))), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, inner_6, gopurs_runtime.Func(func(result_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Bind1_2_1)}, groupEnd1_3_2, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_4_3.V1, result_8)
}))
}))
}))
})
})
}

func Call_errorShow(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], dictShow_1_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_1.V0, x_2).StrVal())))
}

func Call_error(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), gopurs_runtime.Str(x_1)))
}

func Call_debugShow(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], dictShow_1_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var dictShow_1 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_1_loop
_ = dictShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_debug(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_1.V0, x_2).StrVal())))
}

func Call_debug(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_debug(), gopurs_runtime.Str(x_1)))
}

func Call_clear(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), pkg_Effect_Console.Get_clear())
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_discard__2561459590(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}

func Call_discard__317162198(dict_0_loop *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_group__227222742(dictMonadEffect_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value], x_1_loop string) gopurs_runtime.Value {
var dictMonadEffect_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dictMonadEffect_0_loop
_ = dictMonadEffect_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictMonadEffect_0.V1, gopurs_runtime.Apply(pkg_Effect_Console.Get_group(), gopurs_runtime.Str(x_1)))
}

func Call_liftEffect__3456588885(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__1892566677(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


