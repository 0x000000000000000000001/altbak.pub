package Effect_Aff_Compat

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Uncurried "gopurs/output/Effect.Uncurried"
)

var EffectFnCanceler gopurs_runtime.Value
var once_EffectFnCanceler sync.Once
func Get_EffectFnCanceler() gopurs_runtime.Value {
	once_EffectFnCanceler.Do(func() {
		EffectFnCanceler = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return EffectFnCanceler
}

var EffectFnAff gopurs_runtime.Value
var once_EffectFnAff sync.Once
func Get_EffectFnAff() gopurs_runtime.Value {
	once_EffectFnAff.Do(func() {
		EffectFnAff = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return EffectFnAff
}

var fromEffectFnAff gopurs_runtime.Value
var once_fromEffectFnAff sync.Once
func Get_fromEffectFnAff() gopurs_runtime.Value {
	once_fromEffectFnAff.Do(func() {
		fromEffectFnAff = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Aff.Get_makeAff(), gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Uncurried.Get_runEffectFn2(), v_0), gopurs_runtime.Apply(pkg_Effect_Uncurried.Get_mkEffectFn1(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_2}))
}))), gopurs_runtime.Apply(pkg_Effect_Uncurried.Get_mkEffectFn1(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_2}))
})))), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect.Get_pureE(), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Aff.Get_makeAff(), gopurs_runtime.Func(func(k2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Uncurried.Get_runEffectFn3(), v1_2), e_3), gopurs_runtime.Apply(pkg_Effect_Uncurried.Get_mkEffectFn1(), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k2_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_5}))
}))), gopurs_runtime.Apply(pkg_Effect_Uncurried.Get_mkEffectFn1(), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k2_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_5}))
})))), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect.Get_pureE(), pkg_Effect_Aff.Get_nonCanceler())
}))
}))
}))
}))
}))
})
	})
	return fromEffectFnAff
}


