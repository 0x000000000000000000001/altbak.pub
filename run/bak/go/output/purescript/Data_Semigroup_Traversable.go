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
		cache_Data_Semigroup_Traversable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_identity(x_0_box)
})
	})
	return cache_Data_Semigroup_Traversable_identity
}

var cache_Data_Semigroup_Traversable_Traversable1_dollarDict gopurs_runtime.Value
var once_Data_Semigroup_Traversable_Traversable1_dollarDict sync.Once
func Get_Data_Semigroup_Traversable_Traversable1_dollarDict() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_Traversable1_dollarDict.Do(func() {
		cache_Data_Semigroup_Traversable_Traversable1_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_Traversable1_dollarDict(x_0_box)
})
	})
	return cache_Data_Semigroup_Traversable_Traversable1_dollarDict
}

var cache_Data_Semigroup_Traversable_traverse1 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1 sync.Once
func Get_Data_Semigroup_Traversable_traverse1() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_traverse1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_traverse1
}

var cache_Data_Semigroup_Traversable_traversableTuple gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableTuple sync.Once
func Get_Data_Semigroup_Traversable_traversableTuple() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableTuple.Do(func() {
		cache_Data_Semigroup_Traversable_traversableTuple = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Foldable_foldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableTuple()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1))
})
})
}))
	})
	return cache_Data_Semigroup_Traversable_traversableTuple
}

var cache_Data_Semigroup_Traversable_traversableIdentity gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableIdentity sync.Once
func Get_Data_Semigroup_Traversable_traversableIdentity() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableIdentity.Do(func() {
		cache_Data_Semigroup_Traversable_traversableIdentity = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Foldable_foldableIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableIdentity()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Identity_Identity(), v_2)
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Identity_Identity(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Semigroup_Traversable_traversableIdentity
}

var cache_Data_Semigroup_Traversable_sequence1Default gopurs_runtime.Value
var once_Data_Semigroup_Traversable_sequence1Default sync.Once
func Get_Data_Semigroup_Traversable_sequence1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_sequence1Default.Do(func() {
		cache_Data_Semigroup_Traversable_sequence1Default = gopurs_runtime.Func2(func(dictTraversable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_sequence1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dictTraversable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Traversable_sequence1Default
}

var cache_Data_Semigroup_Traversable_traversableDual gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableDual sync.Once
func Get_Data_Semigroup_Traversable_traversableDual() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableDual.Do(func() {
		cache_Data_Semigroup_Traversable_traversableDual = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Foldable_foldableDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableDual()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_Traversable_traversableDual(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0))}, Get_Data_Semigroup_Traversable_identity())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Semigroup_Traversable_traversableDual
}

var cache_Data_Semigroup_Traversable_traversableMultiplicative gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableMultiplicative sync.Once
func Get_Data_Semigroup_Traversable_traversableMultiplicative() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableMultiplicative.Do(func() {
		cache_Data_Semigroup_Traversable_traversableMultiplicative = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Foldable_foldableMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableMultiplicative()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_Traversable_traversableMultiplicative(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0))}, Get_Data_Semigroup_Traversable_identity())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Multiplicative_Multiplicative(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Semigroup_Traversable_traversableMultiplicative
}

var cache_Data_Semigroup_Traversable_sequence1 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_sequence1 sync.Once
func Get_Data_Semigroup_Traversable_sequence1() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_sequence1.Do(func() {
		cache_Data_Semigroup_Traversable_sequence1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_sequence1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_sequence1
}

var cache_Data_Semigroup_Traversable_traverse1Default gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1Default sync.Once
func Get_Data_Semigroup_Traversable_traverse1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1Default.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1Default = gopurs_runtime.Func(func(dictTraversable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_traverse1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dictTraversable1_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_traverse1Default
}

var cache_Data_Semigroup_Traversable_sequence1__4078930642 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_sequence1__4078930642 sync.Once
func Get_Data_Semigroup_Traversable_sequence1__4078930642() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_sequence1__4078930642.Do(func() {
		cache_Data_Semigroup_Traversable_sequence1__4078930642 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_sequence1__4078930642(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_sequence1__4078930642
}

var cache_Data_Semigroup_Traversable_sequence1__2726498490 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_sequence1__2726498490 sync.Once
func Get_Data_Semigroup_Traversable_sequence1__2726498490() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_sequence1__2726498490.Do(func() {
		cache_Data_Semigroup_Traversable_sequence1__2726498490 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_sequence1__2726498490(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](dictApply_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_sequence1__2726498490
}

var cache_Data_Semigroup_Traversable_traverse1__236758920 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1__236758920 sync.Once
func Get_Data_Semigroup_Traversable_traverse1__236758920() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1__236758920.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1__236758920 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_traverse1__236758920(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_traverse1__236758920
}

var cache_Data_Semigroup_Traversable_traverse1__157785005 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1__157785005 sync.Once
func Get_Data_Semigroup_Traversable_traverse1__157785005() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1__157785005.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1__157785005 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_traverse1__157785005(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_traverse1__157785005
}

var cache_Data_Semigroup_Traversable_traverse1__4157595309 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1__4157595309 sync.Once
func Get_Data_Semigroup_Traversable_traverse1__4157595309() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1__4157595309.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1__4157595309 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_traverse1__4157595309(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Traversable_traverse1__4157595309
}

var cache_Data_Semigroup_Traversable_traverse1__42886725 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1__42886725 sync.Once
func Get_Data_Semigroup_Traversable_traverse1__42886725() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1__42886725.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1__42886725 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Traversable_traverse1__42886725(dictApply_0_box)
})
	})
	return cache_Data_Semigroup_Traversable_traverse1__42886725
}

type Constructor_Data_Semigroup_Traversable_Traversable1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1596088409] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_Traversable_Traversable1)(ptr)
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


func Call_Data_Semigroup_Traversable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Traversable_Traversable1_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Traversable_traverse1(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semigroup_Traversable_sequence1Default(dictTraversable1_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1, dictApply_1_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictTraversable1_0 *Constructor_Data_Semigroup_Traversable_Traversable1 = dictTraversable1_0_loop
_ = dictTraversable1_0
var dictApply_1 *Constructor_Control_Apply_Apply = dictApply_1_loop
_ = dictApply_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable1_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_1)}, Get_Data_Semigroup_Traversable_identity())
}

func Call_Data_Semigroup_Traversable_sequence1(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semigroup_Traversable_traverse1Default(dictTraversable1_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1) gopurs_runtime.Value {
var dictTraversable1_0 *Constructor_Data_Semigroup_Traversable_Traversable1 = dictTraversable1_0_loop
_ = dictTraversable1_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversable1_0.V1), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable1_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_2))}, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_3, ta_4))
})
})
})
}

func Call_Data_Semigroup_Traversable_sequence1__4078930642(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semigroup_Traversable_sequence1__2726498490(dictApply_0_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var dictApply_0 *Constructor_Data_NonEmpty_NonEmpty = dictApply_0_loop
_ = dictApply_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversable1NonEmptyList(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(dictApply_0)}, Get_Data_List_Types_identity1())
}

func Call_Data_Semigroup_Traversable_traverse1__236758920(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semigroup_Traversable_traverse1__157785005(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semigroup_Traversable_traverse1__4157595309(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semigroup_Traversable_traverse1__42886725(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Types_nelCons(), a_6, b_5)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_applicativeNonEmptyList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_4.UnsafePtr).V0))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_2 -> *Constructor_Data_Functor_Functor
Functor0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_2.V0), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Types_nelCons(), a_8, b_7)
})
}), acc_4), b_6)
})
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(f_2, x_6))
})
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_List_Types_pure(), gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0))})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))}))
})
})
}


