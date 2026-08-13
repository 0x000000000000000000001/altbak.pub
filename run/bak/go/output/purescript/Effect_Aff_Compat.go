package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Aff_Compat_pure gopurs_runtime.Value
var once_Effect_Aff_Compat_pure sync.Once
func Get_Effect_Aff_Compat_pure() gopurs_runtime.Value {
	once_Effect_Aff_Compat_pure.Do(func() {
		cache_Effect_Aff_Compat_pure = gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure")
	})
	return cache_Effect_Aff_Compat_pure
}

var cache_Effect_Aff_Compat_EffectFnCanceler gopurs_runtime.Value
var once_Effect_Aff_Compat_EffectFnCanceler sync.Once
func Get_Effect_Aff_Compat_EffectFnCanceler() gopurs_runtime.Value {
	once_Effect_Aff_Compat_EffectFnCanceler.Do(func() {
		cache_Effect_Aff_Compat_EffectFnCanceler = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Compat_EffectFnCanceler(x_0_box)
})
	})
	return cache_Effect_Aff_Compat_EffectFnCanceler
}

var cache_Effect_Aff_Compat_EffectFnAff gopurs_runtime.Value
var once_Effect_Aff_Compat_EffectFnAff sync.Once
func Get_Effect_Aff_Compat_EffectFnAff() gopurs_runtime.Value {
	once_Effect_Aff_Compat_EffectFnAff.Do(func() {
		cache_Effect_Aff_Compat_EffectFnAff = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Compat_EffectFnAff(x_0_box)
})
	})
	return cache_Effect_Aff_Compat_EffectFnAff
}

var cache_Effect_Aff_Compat_fromEffectFnAff gopurs_runtime.Value
var once_Effect_Aff_Compat_fromEffectFnAff sync.Once
func Get_Effect_Aff_Compat_fromEffectFnAff() gopurs_runtime.Value {
	once_Effect_Aff_Compat_fromEffectFnAff.Do(func() {
		cache_Effect_Aff_Compat_fromEffectFnAff = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Compat_fromEffectFnAff(v_0_box)
})
	})
	return cache_Effect_Aff_Compat_fromEffectFnAff
}

func Call_Effect_Aff_Compat_EffectFnCanceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Effect_Aff_Compat_EffectFnAff(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Effect_Aff_Compat_fromEffectFnAff(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Effect_Aff_makeAff(), gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.UncurriedApp2(v_0, gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2})}), gopurs_runtime.Value{})
}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2})}), gopurs_runtime.Value{})
}))
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Effect_Aff_makeAff__3447620704(), gopurs_runtime.Func(func(k2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(__local_var_2_0, e_3, gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k2_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5})}), gopurs_runtime.Value{})
}), gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k2_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5})}), gopurs_runtime.Value{})
}))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_nonCanceler()
})
}))
}))
})), gopurs_runtime.Value{})
})
}))
}


