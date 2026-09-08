package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Biapply_Biapply_dollar_Dict gopurs_runtime.Value
var once_Control_Biapply_Biapply_dollar_Dict sync.Once
func Get_Control_Biapply_Biapply_dollar_Dict() gopurs_runtime.Value {
	once_Control_Biapply_Biapply_dollar_Dict.Do(func() {
		cache_Control_Biapply_Biapply_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(Call_Control_Biapply_Biapply_dollar_Dict(func() struct{
	Bifunctor0 gopurs_runtime.Value
	biapply gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bifunctor0 gopurs_runtime.Value
	biapply gopurs_runtime.Value
}{}
					clone.Bifunctor0 = gopurs_runtime.RecordGet(orig, "Bifunctor0")
					clone.biapply = gopurs_runtime.RecordGet(orig, "biapply")
					return clone
				}()))}
})
	})
	return cache_Control_Biapply_Biapply_dollar_Dict
}

var cache_Control_Biapply_biapplyTuple gopurs_runtime.Value
var once_Control_Biapply_biapplyTuple sync.Once
func Get_Control_Biapply_biapplyTuple() gopurs_runtime.Value {
	once_Control_Biapply_biapplyTuple.Do(func() {
		cache_Control_Biapply_biapplyTuple = gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(Rebox_Control_Biapply_4265826836_2448899513((&Constructor_Control_Biapply_Biapply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Control_Biapply_1495429347_1688994542(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Bifunctor_bifunctorTuple())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
})})))}
	})
	return cache_Control_Biapply_biapplyTuple
}

var cache_Control_Biapply_biapply gopurs_runtime.Value
var once_Control_Biapply_biapply sync.Once
func Get_Control_Biapply_biapply() gopurs_runtime.Value {
	once_Control_Biapply_biapply.Do(func() {
		cache_Control_Biapply_biapply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_biapply(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Biapply_biapply
}

var cache_Control_Biapply_biapplyFirst gopurs_runtime.Value
var once_Control_Biapply_biapplyFirst sync.Once
func Get_Control_Biapply_biapplyFirst() gopurs_runtime.Value {
	once_Control_Biapply_biapplyFirst.Do(func() {
		cache_Control_Biapply_biapplyFirst = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_biapplyFirst(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](dictBiapply_0_box))
})
	})
	return cache_Control_Biapply_biapplyFirst
}

var cache_Control_Biapply_biapplySecond gopurs_runtime.Value
var once_Control_Biapply_biapplySecond sync.Once
func Get_Control_Biapply_biapplySecond() gopurs_runtime.Value {
	once_Control_Biapply_biapplySecond.Do(func() {
		cache_Control_Biapply_biapplySecond = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_biapplySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](dictBiapply_0_box))
})
	})
	return cache_Control_Biapply_biapplySecond
}

var cache_Control_Biapply_bilift2 gopurs_runtime.Value
var once_Control_Biapply_bilift2 sync.Once
func Get_Control_Biapply_bilift2() gopurs_runtime.Value {
	once_Control_Biapply_bilift2.Do(func() {
		cache_Control_Biapply_bilift2 = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_bilift2(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](dictBiapply_0_box))
})
	})
	return cache_Control_Biapply_bilift2
}

var cache_Control_Biapply_bilift3 gopurs_runtime.Value
var once_Control_Biapply_bilift3 sync.Once
func Get_Control_Biapply_bilift3() gopurs_runtime.Value {
	once_Control_Biapply_bilift3.Do(func() {
		cache_Control_Biapply_bilift3 = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapply_bilift3(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](dictBiapply_0_box))
})
	})
	return cache_Control_Biapply_bilift3
}

type Constructor_Control_Biapply_Biapply[T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3774602829] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Biapply_Biapply[any])(ptr)
		_ = c
		switch key {
		case "Bifunctor0": return gopurs_runtime.Box(c.V0)
		case "biapply": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Biapply_Biapply: " + key)
		}
	}
}


func Call_Control_Biapply_Biapply_dollar_Dict(x_0_loop struct{
	Bifunctor0 gopurs_runtime.Value
	biapply gopurs_runtime.Value
}) *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] {
var x_0 struct{
	Bifunctor0 gopurs_runtime.Value
	biapply gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Bifunctor0", "biapply"}, []gopurs_runtime.Value{orig.Bifunctor0, orig.biapply})
				}())
}

func Call_Control_Biapply_biapply(dict_0_loop *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Biapply_biapplyFirst(dictBiapply_0_loop *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): Bifunctor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeVar w)])
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictBiapply_0.V0), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
}), a_2), b_3)
})
}

func Call_Control_Biapply_biapplySecond(dictBiapply_0_loop *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): Bifunctor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeVar w)])
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictBiapply_0.V0), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), Get_Data_Function_go__const(), Get_Data_Function_go__const(), a_2), b_3)
})
}

func Call_Control_Biapply_bilift2(dictBiapply_0_loop *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): Bifunctor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeVar w)])
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictBiapply_0.V0), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), f_2, g_3, a_4), b_5)
})
}

func Call_Control_Biapply_bilift3(dictBiapply_0_loop *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBiapply_0 *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): Bifunctor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeVar w)])
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictBiapply_0.V0), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func5(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictBiapply_0.V1), gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), f_2, g_3, a_4), b_5), c_6)
})
}

func Rebox_Control_Biapply_1495429347_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Control_Biapply_4265826836_2448899513(in *Constructor_Control_Biapply_Biapply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


