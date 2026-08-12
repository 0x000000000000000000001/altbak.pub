package Effect_Aff_Compat

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_EffectFnCanceler gopurs_runtime.Value
var once_EffectFnCanceler sync.Once
func Get_EffectFnCanceler() gopurs_runtime.Value {
	once_EffectFnCanceler.Do(func() {
		cache_EffectFnCanceler = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_EffectFnCanceler(x_0_box)
})
	})
	return cache_EffectFnCanceler
}

var cache_EffectFnAff gopurs_runtime.Value
var once_EffectFnAff sync.Once
func Get_EffectFnAff() gopurs_runtime.Value {
	once_EffectFnAff.Do(func() {
		cache_EffectFnAff = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_EffectFnAff(x_0_box)
})
	})
	return cache_EffectFnAff
}

var cache_fromEffectFnAff gopurs_runtime.Value
var once_fromEffectFnAff sync.Once
func Get_fromEffectFnAff() gopurs_runtime.Value {
	once_fromEffectFnAff.Do(func() {
		cache_fromEffectFnAff = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEffectFnAff(v_0_box)
})
	})
	return cache_fromEffectFnAff
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

var cache_pure__3514127574 gopurs_runtime.Value
var once_pure__3514127574 sync.Once
func Get_pure__3514127574() gopurs_runtime.Value {
	once_pure__3514127574.Do(func() {
		cache_pure__3514127574 = gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_applicativeAff(), "pure")
	})
	return cache_pure__3514127574
}

var cache_pure__2106705590 gopurs_runtime.Value
var once_pure__2106705590 sync.Once
func Get_pure__2106705590() gopurs_runtime.Value {
	once_pure__2106705590.Do(func() {
		cache_pure__2106705590 = gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure")
	})
	return cache_pure__2106705590
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

var cache_bind__1922668001 gopurs_runtime.Value
var once_bind__1922668001 sync.Once
func Get_bind__1922668001() gopurs_runtime.Value {
	once_bind__1922668001.Do(func() {
		cache_bind__1922668001 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__1922668001
}

var cache_discard__2966453088 gopurs_runtime.Value
var once_discard__2966453088 sync.Once
func Get_discard__2966453088() gopurs_runtime.Value {
	once_discard__2966453088.Do(func() {
		cache_discard__2966453088 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard__2966453088
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

var cache_const__3415939124 gopurs_runtime.Value
var once_const__3415939124 sync.Once
func Get_const__3415939124() gopurs_runtime.Value {
	once_const__3415939124.Do(func() {
		cache_const__3415939124 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__3415939124(a_0_box, v_1_box)
})
	})
	return cache_const__3415939124
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

var cache_functorAff__2378915857 gopurs_runtime.Value
var once_functorAff__2378915857 sync.Once
func Get_functorAff__2378915857() gopurs_runtime.Value {
	once_functorAff__2378915857.Do(func() {
		cache_functorAff__2378915857 = gopurs_runtime.RecordDict1("map", pkg_Effect_Aff.Get__map())
	})
	return cache_functorAff__2378915857
}

var cache_makeAff__3447620704 gopurs_runtime.Value
var once_makeAff__3447620704 sync.Once
func Get_makeAff__3447620704() gopurs_runtime.Value {
	once_makeAff__3447620704.Do(func() {
		cache_makeAff__3447620704 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff__3447620704(build_0_box)
})
	})
	return cache_makeAff__3447620704
}

var cache_makeAff__3958971776 gopurs_runtime.Value
var once_makeAff__3958971776 sync.Once
func Get_makeAff__3958971776() gopurs_runtime.Value {
	once_makeAff__3958971776.Do(func() {
		cache_makeAff__3958971776 = gopurs_runtime.Func(func(build_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff__3958971776(build_0_box)
})
	})
	return cache_makeAff__3958971776
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

func Call_EffectFnCanceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_EffectFnAff(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_fromEffectFnAff(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(pkg_Effect_Aff.Get_makeAff(), gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_bind__1922668001(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(v_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2})}), gopurs_runtime.Value{})
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2})}), gopurs_runtime.Value{})
}))
}), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeAff__3447620704(gopurs_runtime.Func(func(k2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard__2966453088(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(v1_2, e_3, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k2_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_5})}), gopurs_runtime.Value{})
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k2_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_5})}), gopurs_runtime.Value{})
}))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_pure__2106705590(), pkg_Effect_Aff.Get_nonCanceler())
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

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__3415939124(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_makeAff__3447620704(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(pkg_Effect_Aff.Get__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (either_3.Type == 9 && either_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(onError_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (either_3.Type == 9 && either_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(onSuccess_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
}))
}

func Call_makeAff__3958971776(build_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var build_0 gopurs_runtime.Value = build_0_loop
_ = build_0
return gopurs_runtime.Apply(pkg_Effect_Aff.Get__makeAffImpl(), gopurs_runtime.Func(func(onError_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(onSuccess_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(build_0, gopurs_runtime.Func(func(either_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (either_3.Type == 9 && either_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(onError_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (either_3.Type == 9 && either_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(onSuccess_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(either_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
}))
}


