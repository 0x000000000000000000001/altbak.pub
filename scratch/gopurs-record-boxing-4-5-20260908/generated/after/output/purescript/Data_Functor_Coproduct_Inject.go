package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Coproduct_Inject_Inject_dollar_Dict gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_Inject_dollar_Dict sync.Once
func Get_Data_Functor_Coproduct_Inject_Inject_dollar_Dict() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_Inject_dollar_Dict.Do(func() {
		cache_Data_Functor_Coproduct_Inject_Inject_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3819687178, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Coproduct_Inject_Inject_dollar_Dict(func() struct{
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
}{}
					clone.inj = gopurs_runtime.RecordGet(orig, "inj")
					clone.prj = gopurs_runtime.RecordGet(orig, "prj")
					return clone
				}()))}
})
	})
	return cache_Data_Functor_Coproduct_Inject_Inject_dollar_Dict
}

var cache_Data_Functor_Coproduct_Inject_prj gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_prj sync.Once
func Get_Data_Functor_Coproduct_Inject_prj() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_prj.Do(func() {
		cache_Data_Functor_Coproduct_Inject_prj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_Coproduct_Inject_prj(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Functor_Coproduct_Inject_prj
}

var cache_Data_Functor_Coproduct_Inject_injectReflexive gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_injectReflexive sync.Once
func Get_Data_Functor_Coproduct_Inject_injectReflexive() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_injectReflexive.Do(func() {
		cache_Data_Functor_Coproduct_Inject_injectReflexive = gopurs_runtime.Value{Type: 9, IntVal: 3819687178, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), Get_Data_Maybe_Just()}))}
	})
	return cache_Data_Functor_Coproduct_Inject_injectReflexive
}

var cache_Data_Functor_Coproduct_Inject_injectLeft gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_injectLeft sync.Once
func Get_Data_Functor_Coproduct_Inject_injectLeft() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_injectLeft.Do(func() {
		cache_Data_Functor_Coproduct_Inject_injectLeft = gopurs_runtime.Value{Type: 9, IntVal: 3819687178, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0})
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})}))}
	})
	return cache_Data_Functor_Coproduct_Inject_injectLeft
}

var cache_Data_Functor_Coproduct_Inject_inj gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_inj sync.Once
func Get_Data_Functor_Coproduct_Inject_inj() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_inj.Do(func() {
		cache_Data_Functor_Coproduct_Inject_inj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_Inject_inj(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Functor_Coproduct_Inject_inj
}

var cache_Data_Functor_Coproduct_Inject_injectRight gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_injectRight sync.Once
func Get_Data_Functor_Coproduct_Inject_injectRight() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_injectRight.Do(func() {
		cache_Data_Functor_Coproduct_Inject_injectRight = gopurs_runtime.Func(func(dictInject_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_Inject_injectRight(dictInject_0_box)
})
	})
	return cache_Data_Functor_Coproduct_Inject_injectRight
}

type Constructor_Data_Functor_Coproduct_Inject_Inject[T_f any, T_g any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3819687178] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Functor_Coproduct_Inject_Inject[any, any])(ptr)
		_ = c
		switch key {
		case "inj": return gopurs_runtime.Box(c.V0)
		case "prj": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Functor_Coproduct_Inject_Inject: " + key)
		}
	}
}


func Call_Data_Functor_Coproduct_Inject_Inject_dollar_Dict(x_0_loop struct{
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
}) *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("inj", "prj", orig.inj, orig.prj)
				}())
}

func Call_Data_Functor_Coproduct_Inject_prj(dict_0_loop *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dict_0 *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Box(dict_0.V1)
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Functor_Coproduct_Inject_inj(dict_0_loop *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_Coproduct_Inject_injectRight(dictInject_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInject_0 gopurs_runtime.Value = dictInject_0_loop
_ = dictInject_0
return gopurs_runtime.Value{Type: 9, IntVal: 3819687178, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0, "inj"), x_1)}))}
}), gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0, "prj"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
}


