package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bifunctor_identity gopurs_runtime.Value
var once_Data_Bifunctor_identity sync.Once
func Get_Data_Bifunctor_identity() gopurs_runtime.Value {
	once_Data_Bifunctor_identity.Do(func() {
		cache_Data_Bifunctor_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Bifunctor_identity
}

var cache_Data_Bifunctor_identity1 gopurs_runtime.Value
var once_Data_Bifunctor_identity1 sync.Once
func Get_Data_Bifunctor_identity1() gopurs_runtime.Value {
	once_Data_Bifunctor_identity1.Do(func() {
		cache_Data_Bifunctor_identity1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Bifunctor_identity1
}

var cache_Data_Bifunctor_Bifunctor_dollar_Dict gopurs_runtime.Value
var once_Data_Bifunctor_Bifunctor_dollar_Dict sync.Once
func Get_Data_Bifunctor_Bifunctor_dollar_Dict() gopurs_runtime.Value {
	once_Data_Bifunctor_Bifunctor_dollar_Dict.Do(func() {
		cache_Data_Bifunctor_Bifunctor_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Call_Data_Bifunctor_Bifunctor_dollar_Dict(func() struct{
	bimap gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	bimap gopurs_runtime.Value
}{}
					clone.bimap = gopurs_runtime.RecordGet(orig, "bimap")
					return clone
				}()))}
})
	})
	return cache_Data_Bifunctor_Bifunctor_dollar_Dict
}

var cache_Data_Bifunctor_bimap gopurs_runtime.Value
var once_Data_Bifunctor_bimap sync.Once
func Get_Data_Bifunctor_bimap() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap.Do(func() {
		cache_Data_Bifunctor_bimap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifunctor_bimap
}

var cache_Data_Bifunctor_bivoid gopurs_runtime.Value
var once_Data_Bifunctor_bivoid sync.Once
func Get_Data_Bifunctor_bivoid() gopurs_runtime.Value {
	once_Data_Bifunctor_bivoid.Do(func() {
		cache_Data_Bifunctor_bivoid = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bivoid(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](dictBifunctor_0_box))
})
	})
	return cache_Data_Bifunctor_bivoid
}

var cache_Data_Bifunctor_lmap gopurs_runtime.Value
var once_Data_Bifunctor_lmap sync.Once
func Get_Data_Bifunctor_lmap() gopurs_runtime.Value {
	once_Data_Bifunctor_lmap.Do(func() {
		cache_Data_Bifunctor_lmap = gopurs_runtime.Func2(func(dictBifunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_lmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](dictBifunctor_0_box), f_1_box)
})
	})
	return cache_Data_Bifunctor_lmap
}

var cache_Data_Bifunctor_rmap gopurs_runtime.Value
var once_Data_Bifunctor_rmap sync.Once
func Get_Data_Bifunctor_rmap() gopurs_runtime.Value {
	once_Data_Bifunctor_rmap.Do(func() {
		cache_Data_Bifunctor_rmap = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_rmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](dictBifunctor_0_box))
})
	})
	return cache_Data_Bifunctor_rmap
}

var cache_Data_Bifunctor_bifunctorTuple gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorTuple sync.Once
func Get_Data_Bifunctor_bifunctorTuple() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorTuple.Do(func() {
		cache_Data_Bifunctor_bifunctorTuple = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Bifunctor_1495429347_1688994542((&Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
})})))}
	})
	return cache_Data_Bifunctor_bifunctorTuple
}

var cache_Data_Bifunctor_bifunctorEither gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorEither sync.Once
func Get_Data_Bifunctor_bifunctorEither() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorEither.Do(func() {
		cache_Data_Bifunctor_bifunctorEither = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
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
	})
	return cache_Data_Bifunctor_bifunctorEither
}

var cache_Data_Bifunctor_bifunctorConst gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorConst sync.Once
func Get_Data_Bifunctor_bifunctorConst() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorConst.Do(func() {
		cache_Data_Bifunctor_bifunctorConst = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})}))}
	})
	return cache_Data_Bifunctor_bifunctorConst
}

type Constructor_Data_Bifunctor_Bifunctor[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4141114362] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bifunctor_Bifunctor[any])(ptr)
		_ = c
		switch key {
		case "bimap": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Bifunctor_Bifunctor: " + key)
		}
	}
}


func Call_Data_Bifunctor_Bifunctor_dollar_Dict(x_0_loop struct{
	bimap gopurs_runtime.Value
}) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
var x_0 struct{
	bimap gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"bimap"}, []gopurs_runtime.Value{orig.bimap})
				}())
}

func Call_Data_Bifunctor_bimap(dict_0_loop *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifunctor_bivoid(dictBifunctor_0_loop *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBifunctor_0.V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
}

func Call_Data_Bifunctor_lmap(dictBifunctor_0_loop *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] = dictBifunctor_0_loop
_ = dictBifunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBifunctor_0.V0), f_1, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Bifunctor_rmap(dictBifunctor_0_loop *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictBifunctor_0.V0), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Rebox_Data_Bifunctor_1495429347_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}


