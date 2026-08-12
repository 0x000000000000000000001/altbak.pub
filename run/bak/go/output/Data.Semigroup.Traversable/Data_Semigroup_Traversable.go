package Data_Semigroup_Traversable

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_traverse1 gopurs_runtime.Value
var once_traverse1 sync.Once
func Get_traverse1() gopurs_runtime.Value {
	once_traverse1.Do(func() {
		cache_traverse1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1(gopurs_runtime.CoerceToStruct[Constructor_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse1
}

var cache_traverse1__gopurs_runtime_Value_157785005 gopurs_runtime.Value
var once_traverse1__gopurs_runtime_Value_157785005 sync.Once
func Get_traverse1__gopurs_runtime_Value_157785005() gopurs_runtime.Value {
	once_traverse1__gopurs_runtime_Value_157785005.Do(func() {
		cache_traverse1__gopurs_runtime_Value_157785005 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1__gopurs_runtime_Value_157785005(gopurs_runtime.CoerceToStruct[Constructor_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse1__gopurs_runtime_Value_157785005
}

var cache_traversableTuple gopurs_runtime.Value
var once_traversableTuple sync.Once
func Get_traversableTuple() gopurs_runtime.Value {
	once_traversableTuple.Do(func() {
		cache_traversableTuple = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableTuple()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})
}))
	})
	return cache_traversableTuple
}

var cache_traversableIdentity gopurs_runtime.Value
var once_traversableIdentity sync.Once
func Get_traversableIdentity() gopurs_runtime.Value {
	once_traversableIdentity.Do(func() {
		cache_traversableIdentity = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableIdentity()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Identity.Get_Identity(), v_2)
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableIdentity
}

var cache_sequence1Default gopurs_runtime.Value
var once_sequence1Default sync.Once
func Get_sequence1Default() gopurs_runtime.Value {
	once_sequence1Default.Do(func() {
		cache_sequence1Default = gopurs_runtime.Func2(func(dictTraversable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1Default(gopurs_runtime.CoerceToStruct[Constructor_Traversable1[gopurs_runtime.Value]](dictTraversable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_sequence1Default
}

var cache_traversableDual gopurs_runtime.Value
var once_traversableDual sync.Once
func Get_traversableDual() gopurs_runtime.Value {
	once_traversableDual.Do(func() {
		cache_traversableDual = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableDual()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableDual(), "traverse1"), dictApply_0, Get_identity())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Dual.Get_Dual(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableDual
}

var cache_traversableMultiplicative gopurs_runtime.Value
var once_traversableMultiplicative sync.Once
func Get_traversableMultiplicative() gopurs_runtime.Value {
	once_traversableMultiplicative.Do(func() {
		cache_traversableMultiplicative = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableMultiplicative()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMultiplicative(), "traverse1"), dictApply_0, Get_identity())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableMultiplicative
}

var cache_sequence1 gopurs_runtime.Value
var once_sequence1 sync.Once
func Get_sequence1() gopurs_runtime.Value {
	once_sequence1.Do(func() {
		cache_sequence1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1(gopurs_runtime.CoerceToStruct[Constructor_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence1
}

var cache_sequence1__gopurs_runtime_Value_4078930642 gopurs_runtime.Value
var once_sequence1__gopurs_runtime_Value_4078930642 sync.Once
func Get_sequence1__gopurs_runtime_Value_4078930642() gopurs_runtime.Value {
	once_sequence1__gopurs_runtime_Value_4078930642.Do(func() {
		cache_sequence1__gopurs_runtime_Value_4078930642 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1__gopurs_runtime_Value_4078930642(gopurs_runtime.CoerceToStruct[Constructor_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence1__gopurs_runtime_Value_4078930642
}

var cache_traverse1Default gopurs_runtime.Value
var once_traverse1Default sync.Once
func Get_traverse1Default() gopurs_runtime.Value {
	once_traverse1Default.Do(func() {
		cache_traverse1Default = gopurs_runtime.Func(func(dictTraversable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1Default(gopurs_runtime.CoerceToStruct[Constructor_Traversable1[gopurs_runtime.Value]](dictTraversable1_0_box))
})
	})
	return cache_traverse1Default
}

type Constructor_Traversable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1596088409] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Traversable1[gopurs_runtime.Value])(ptr)
		switch key {
		case "Foldable10": return c.V0
		case "Traversable1": return c.V1
		case "sequence1": return c.V2
		case "traverse1": return c.V3
		default: panic("Key not found in dictionary Constructor_Traversable1: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_traverse1(dict_0_loop *Constructor_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse1__gopurs_runtime_Value_157785005(dict_0_loop *Constructor_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_sequence1Default(dictTraversable1_0_loop *Constructor_Traversable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversable1_0 *Constructor_Traversable1[gopurs_runtime.Value] = dictTraversable1_0_loop
_ = dictTraversable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
return gopurs_runtime.Apply2(dictTraversable1_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_1)}, Get_identity())
}

func Call_sequence1(dict_0_loop *Constructor_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_sequence1__gopurs_runtime_Value_4078930642(dict_0_loop *Constructor_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traverse1Default(dictTraversable1_0_loop *Constructor_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversable1_0 *Constructor_Traversable1[gopurs_runtime.Value] = dictTraversable1_0_loop
_ = dictTraversable1_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictTraversable1_0.V1, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictTraversable1_0.V2, dictApply_2, gopurs_runtime.Apply2(Functor0_1_0.V0, f_3, ta_4))
})
})
})
}


