package Bench

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Effect_Class "gopurs/output/Effect.Class"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_runBenchAff gopurs_runtime.Value
var once_runBenchAff sync.Once
func Get_runBenchAff() gopurs_runtime.Value {
	once_runBenchAff.Do(func() {
		cache_runBenchAff = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runBenchAff(describe_0_box, act_1_box)
})
	})
	return cache_runBenchAff
}

var cache_runBench gopurs_runtime.Value
var once_runBench sync.Once
func Get_runBench() gopurs_runtime.Value {
	once_runBench.Do(func() {
		cache_runBench = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runBench(describe_0_box, act_1_box)
})
	})
	return cache_runBench
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

var cache_pure__1895379222 gopurs_runtime.Value
var once_pure__1895379222 sync.Once
func Get_pure__1895379222() gopurs_runtime.Value {
	once_pure__1895379222.Do(func() {
		cache_pure__1895379222 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1895379222(__eta0_0_box)
})
	})
	return cache_pure__1895379222
}

var cache_pure__629383158 gopurs_runtime.Value
var once_pure__629383158 sync.Once
func Get_pure__629383158() gopurs_runtime.Value {
	once_pure__629383158.Do(func() {
		cache_pure__629383158 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__629383158(__eta0_0_box)
})
	})
	return cache_pure__629383158
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

var cache_bind__490123073 gopurs_runtime.Value
var once_bind__490123073 sync.Once
func Get_bind__490123073() gopurs_runtime.Value {
	once_bind__490123073.Do(func() {
		cache_bind__490123073 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__490123073(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bind__490123073
}

var cache_bind__1949526049 gopurs_runtime.Value
var once_bind__1949526049 sync.Once
func Get_bind__1949526049() gopurs_runtime.Value {
	once_bind__1949526049.Do(func() {
		cache_bind__1949526049 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1949526049(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bind__1949526049
}

var cache_discard__2596713024 gopurs_runtime.Value
var once_discard__2596713024 sync.Once
func Get_discard__2596713024() gopurs_runtime.Value {
	once_discard__2596713024.Do(func() {
		cache_discard__2596713024 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__2596713024(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_discard__2596713024
}

var cache_discard__203210016 gopurs_runtime.Value
var once_discard__203210016 sync.Once
func Get_discard__203210016() gopurs_runtime.Value {
	once_discard__203210016.Do(func() {
		cache_discard__203210016 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__203210016(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_discard__203210016
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

var cache_sub__1135378904 gopurs_runtime.Value
var once_sub__1135378904 sync.Once
func Get_sub__1135378904() gopurs_runtime.Value {
	once_sub__1135378904.Do(func() {
		cache_sub__1135378904 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1135378904(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1135378904
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_applicativeAff__3333162410 gopurs_runtime.Value
var once_applicativeAff__3333162410 sync.Once
func Get_applicativeAff__3333162410() gopurs_runtime.Value {
	once_applicativeAff__3333162410.Do(func() {
		cache_applicativeAff__3333162410 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_applyAff()
}), pkg_Effect_Aff.Get__pure())
	})
	return cache_applicativeAff__3333162410
}

var cache_applicativeAff__156155496 gopurs_runtime.Value
var once_applicativeAff__156155496 sync.Once
func Get_applicativeAff__156155496() gopurs_runtime.Value {
	once_applicativeAff__156155496.Do(func() {
		cache_applicativeAff__156155496 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_applyAff()
}), pkg_Effect_Aff.Get__pure())
	})
	return cache_applicativeAff__156155496
}

var cache_applyAff__2964533948 gopurs_runtime.Value
var once_applyAff__2964533948 sync.Once
func Get_applyAff__2964533948() gopurs_runtime.Value {
	once_applyAff__2964533948.Do(func() {
		cache_applyAff__2964533948 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadAff(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadAff(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_functorAff()
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
	return cache_applyAff__2964533948
}

var cache_bindAff__1273005738 gopurs_runtime.Value
var once_bindAff__1273005738 sync.Once
func Get_bindAff__1273005738() gopurs_runtime.Value {
	once_bindAff__1273005738.Do(func() {
		cache_bindAff__1273005738 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_applyAff()
}), pkg_Effect_Aff.Get__bind())
	})
	return cache_bindAff__1273005738
}

var cache_bindAff__1025486311 gopurs_runtime.Value
var once_bindAff__1025486311 sync.Once
func Get_bindAff__1025486311() gopurs_runtime.Value {
	once_bindAff__1025486311.Do(func() {
		cache_bindAff__1025486311 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_applyAff()
}), pkg_Effect_Aff.Get__bind())
	})
	return cache_bindAff__1025486311
}

var cache_functorAff__2378915857 gopurs_runtime.Value
var once_functorAff__2378915857 sync.Once
func Get_functorAff__2378915857() gopurs_runtime.Value {
	once_functorAff__2378915857.Do(func() {
		cache_functorAff__2378915857 = gopurs_runtime.RecordDict1("map", pkg_Effect_Aff.Get__map())
	})
	return cache_functorAff__2378915857
}

var cache_monadAff__2914113427 gopurs_runtime.Value
var once_monadAff__2914113427 sync.Once
func Get_monadAff__2914113427() gopurs_runtime.Value {
	once_monadAff__2914113427.Do(func() {
		cache_monadAff__2914113427 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_applicativeAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_bindAff()
}))
	})
	return cache_monadAff__2914113427
}

var cache_monadEffectAff__2194637066 gopurs_runtime.Value
var once_monadEffectAff__2194637066 sync.Once
func Get_monadEffectAff__2194637066() gopurs_runtime.Value {
	once_monadEffectAff__2194637066.Do(func() {
		cache_monadEffectAff__2194637066 = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_monadAff()
}), pkg_Effect_Aff.Get__liftEffect())
	})
	return cache_monadEffectAff__2194637066
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

var cache_liftEffect__273534483 gopurs_runtime.Value
var once_liftEffect__273534483 sync.Once
func Get_liftEffect__273534483() gopurs_runtime.Value {
	once_liftEffect__273534483.Do(func() {
		cache_liftEffect__273534483 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__273534483(__eta0_0_box)
})
	})
	return cache_liftEffect__273534483
}

var cache_liftEffect__3226494803 gopurs_runtime.Value
var once_liftEffect__3226494803 sync.Once
func Get_liftEffect__3226494803() gopurs_runtime.Value {
	once_liftEffect__3226494803.Do(func() {
		cache_liftEffect__3226494803 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__3226494803(__eta0_0_box)
})
	})
	return cache_liftEffect__3226494803
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

func Call_runBenchAff(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return Call_discard__2596713024(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a"))), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__2596713024(Call_liftEffect__3226494803(describe_0), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__2596713024(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("\x0a(Output)\x0a"))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__490123073(Call_liftEffect__273534483(Get_benchNow()), gopurs_runtime.Func(func(t1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__2596713024(act_1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__490123073(Call_liftEffect__273534483(Get_benchNow()), gopurs_runtime.Func(func(t2_7 gopurs_runtime.Value) gopurs_runtime.Value {
dt_8_0 := Call_sub__1135378904(gopurs_runtime.Float(t2_7.FloatVal()), gopurs_runtime.Float(t1_5.FloatVal()))
_ = dt_8_0
return Call_discard__2596713024(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("\x0a(Execution time)\x0a\x0a"), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(Get_formatNumber(), gopurs_runtime.Float(dt_8_0.FloatVal())).StrVal()), gopurs_runtime.Str(" μs\x0a")).StrVal())).StrVal()))), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1895379222(gopurs_runtime.Float(dt_8_0.FloatVal()))
}))
}))
}))
}))
}))
}))
}))
}

func Call_runBench(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return Call_discard__203210016(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a")), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__203210016(describe_0, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__203210016(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("\x0a(Output)\x0a")), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1949526049(Get_benchNow(), gopurs_runtime.Func(func(t1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__203210016(act_1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1949526049(Get_benchNow(), gopurs_runtime.Func(func(t2_7 gopurs_runtime.Value) gopurs_runtime.Value {
dt_8_0 := Call_sub__1135378904(gopurs_runtime.Float(t2_7.FloatVal()), gopurs_runtime.Float(t1_5.FloatVal()))
_ = dt_8_0
return Call_discard__203210016(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("\x0a(Execution time)\x0a\x0a"), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(Get_formatNumber(), gopurs_runtime.Float(dt_8_0.FloatVal())).StrVal()), gopurs_runtime.Str(" μs\x0a")).StrVal())).StrVal())), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__629383158(gopurs_runtime.Float(dt_8_0.FloatVal()))
}))
}))
}))
}))
}))
}))
}))
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1895379222(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(pkg_Effect_Aff.Get__pure(), __eta0_0)
}

func Call_pure__629383158(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_bind__490123073(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(pkg_Effect_Aff.Get__bind(), __eta0_0, __eta1_1)
}

func Call_bind__1949526049(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_discard__2596713024(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_bindAff(), "bind"), __eta0_0, __eta1_1)
}

func Call_discard__203210016(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_sub__1135378904(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) - (__eta1_1.FloatVal()))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_liftEffect__1892566677(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__273534483(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(pkg_Effect_Aff.Get__liftEffect(), __eta0_0)
}

func Call_liftEffect__3226494803(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(pkg_Effect_Aff.Get__liftEffect(), __eta0_0)
}

func Get_benchNow() gopurs_runtime.Value {
	return _Gopurs_BenchNow
}

func Get_formatNumber() gopurs_runtime.Value {
	return _Gopurs_FormatNumber
}

func Get_opaque() gopurs_runtime.Value {
	return _Gopurs_Opaque
}
