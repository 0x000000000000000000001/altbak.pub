package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semigroup_Traversable_identity gopurs_runtime.Value
var once_Data_Semigroup_Traversable_identity sync.Once
func Get_Data_Semigroup_Traversable_identity() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_identity.Do(func() {
		cache_Data_Semigroup_Traversable_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Semigroup_Traversable_identity
}

var cache_Data_Semigroup_Traversable_Traversable1_dollar_Dict gopurs_runtime.Value
var once_Data_Semigroup_Traversable_Traversable1_dollar_Dict sync.Once
func Get_Data_Semigroup_Traversable_Traversable1_dollar_Dict() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_Traversable1_dollar_Dict.Do(func() {
		cache_Data_Semigroup_Traversable_Traversable1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(Call_Data_Semigroup_Traversable_Traversable1_dollar_Dict(func() struct{
	Foldable10 gopurs_runtime.Value
	Traversable1 gopurs_runtime.Value
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Foldable10 gopurs_runtime.Value
	Traversable1 gopurs_runtime.Value
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
}{}
					clone.Foldable10 = gopurs_runtime.RecordGet(orig, "Foldable10")
					clone.Traversable1 = gopurs_runtime.RecordGet(orig, "Traversable1")
					clone.sequence1 = gopurs_runtime.RecordGet(orig, "sequence1")
					clone.traverse1 = gopurs_runtime.RecordGet(orig, "traverse1")
					return clone
				}()))}
})
	})
	return cache_Data_Semigroup_Traversable_Traversable1_dollar_Dict
}

var cache_Data_Semigroup_Traversable_traverse1 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1 sync.Once
func Get_Data_Semigroup_Traversable_traverse1() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_traverse1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_traverse1
}

var cache_Data_Semigroup_Traversable_traversableTuple gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableTuple sync.Once
func Get_Data_Semigroup_Traversable_traversableTuple() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableTuple.Do(func() {
		cache_Data_Semigroup_Traversable_traversableTuple = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_Traversable_1487937088_306175789((&Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_Traversable_1136416832_4151366573(Rebox_Data_Semigroup_Traversable_4151366573_1136416832(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Semigroup_Foldable_foldableTuple()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_Traversable_3543431075_3043886126(Rebox_Data_Semigroup_Traversable_3043886126_3543431075(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableTuple()))))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})})))}
	})
	return cache_Data_Semigroup_Traversable_traversableTuple
}

var cache_Data_Semigroup_Traversable_traversableIdentity gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableIdentity sync.Once
func Get_Data_Semigroup_Traversable_traversableIdentity() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableIdentity.Do(func() {
		cache_Data_Semigroup_Traversable_traversableIdentity = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Semigroup_Foldable_foldableIdentity()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableIdentity()))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Identity_Identity(), v_2)
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Identity_Identity(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
	})
	return cache_Data_Semigroup_Traversable_traversableIdentity
}

var cache_Data_Semigroup_Traversable_sequence1Default gopurs_runtime.Value
var once_Data_Semigroup_Traversable_sequence1Default sync.Once
func Get_Data_Semigroup_Traversable_sequence1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_sequence1Default.Do(func() {
		cache_Data_Semigroup_Traversable_sequence1Default = gopurs_runtime.Func2(func(dictTraversable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_sequence1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dictTraversable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Traversable_sequence1Default
}

var cache_Data_Semigroup_Traversable_traversableDual gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableDual sync.Once
func Get_Data_Semigroup_Traversable_traversableDual() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableDual.Do(func() {
		cache_Data_Semigroup_Traversable_traversableDual = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Semigroup_Foldable_foldableDual()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableDual()))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](Get_Data_Semigroup_Traversable_traversableDual()).V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
	})
	return cache_Data_Semigroup_Traversable_traversableDual
}

var cache_Data_Semigroup_Traversable_traversableMultiplicative gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableMultiplicative sync.Once
func Get_Data_Semigroup_Traversable_traversableMultiplicative() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableMultiplicative.Do(func() {
		cache_Data_Semigroup_Traversable_traversableMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Semigroup_Foldable_foldableMultiplicative()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMultiplicative()))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](Get_Data_Semigroup_Traversable_traversableMultiplicative()).V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Multiplicative_Multiplicative(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
	})
	return cache_Data_Semigroup_Traversable_traversableMultiplicative
}

var cache_Data_Semigroup_Traversable_sequence1 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_sequence1 sync.Once
func Get_Data_Semigroup_Traversable_sequence1() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_sequence1.Do(func() {
		cache_Data_Semigroup_Traversable_sequence1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_sequence1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_sequence1
}

var cache_Data_Semigroup_Traversable_traverse1Default gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1Default sync.Once
func Get_Data_Semigroup_Traversable_traverse1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1Default.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1Default = gopurs_runtime.Func(func(dictTraversable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_traverse1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dictTraversable1_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_traverse1Default
}

type Constructor_Data_Semigroup_Traversable_Traversable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1596088409] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_Traversable_Traversable1[any])(ptr)
		_ = c
		switch key {
		case "Foldable10": return gopurs_runtime.Box(c.V0)
		case "Traversable1": return gopurs_runtime.Box(c.V1)
		case "sequence1": return gopurs_runtime.Box(c.V2)
		case "traverse1": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_Semigroup_Traversable_Traversable1: " + key)
		}
	}
}


func Call_Data_Semigroup_Traversable_Traversable1_dollar_Dict(x_0_loop struct{
	Foldable10 gopurs_runtime.Value
	Traversable1 gopurs_runtime.Value
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
}) *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] {
var x_0 struct{
	Foldable10 gopurs_runtime.Value
	Traversable1 gopurs_runtime.Value
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", orig.Foldable10, orig.Traversable1, orig.sequence1, orig.traverse1)
				}())
}

func Call_Data_Semigroup_Traversable_traverse1(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semigroup_Traversable_sequence1Default(dictTraversable1_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value], dictApply_1_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversable1_0 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dictTraversable1_0_loop
_ = dictTraversable1_0
var dictApply_1 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable1_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_1)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Semigroup_Traversable_sequence1(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semigroup_Traversable_traverse1Default(dictTraversable1_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversable1_0 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dictTraversable1_0_loop
_ = dictTraversable1_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar t)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversable1_0.V1), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func3(func(dictApply_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable1_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_2))}, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_3, ta_4))
})
}

func Rebox_Data_Semigroup_Traversable_1136416832_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Semigroup_Traversable_1487937088_306175789(in *Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Semigroup_Traversable_3043886126_3543431075(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Semigroup_Traversable_3543431075_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Semigroup_Traversable_4151366573_1136416832(in *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}


