package Effect_Aff_Compat

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
)

var cache_discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		cache_discard = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard
}

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
return gopurs_runtime.UncurriedApp6(pkg_Effect_Aff.Get__makeAff(), pkg_Effect_Aff.Get_isLeft(), pkg_Effect_Aff.Get_unsafeFromLeft(), pkg_Effect_Aff.Get_unsafeFromRight(), pkg_Data_Either.Get_Left(), pkg_Data_Either.Get_Right(), gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(v_0, gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Semigroupoid.Get_composeImpl(), k_1, pkg_Data_Either.Get_Left(), __local_var_2)
}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Semigroupoid.Get_composeImpl(), k_1, pkg_Data_Either.Get_Right(), __local_var_2)
}))
}), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp6(pkg_Effect_Aff.Get__makeAff(), pkg_Effect_Aff.Get_isLeft(), pkg_Effect_Aff.Get_unsafeFromLeft(), pkg_Effect_Aff.Get_unsafeFromRight(), pkg_Data_Either.Get_Left(), pkg_Data_Either.Get_Right(), gopurs_runtime.Func(func(k2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(v1_2, e_3, gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Semigroupoid.Get_composeImpl(), k2_4, pkg_Data_Either.Get_Left(), __local_var_5)
}), gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Semigroupoid.Get_composeImpl(), k2_4, pkg_Data_Either.Get_Right(), __local_var_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), pkg_Effect_Aff.Get_nonCanceler())
}))
}))
}))
}))
}))
}


