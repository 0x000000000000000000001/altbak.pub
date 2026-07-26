package Data_Semigroup_Traversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
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
return Call_traverse1(dict_0_box)
})
	})
	return cache_traverse1
}

var cache_traversableTuple gopurs_runtime.Value
var once_traversableTuple sync.Once
func Get_traversableTuple() gopurs_runtime.Value {
	once_traversableTuple.Do(func() {
		cache_traversableTuple = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableTuple()
}), gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
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
}), gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), v_1)
}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply(f_1, v_2))
}))
	})
	return cache_traversableIdentity
}

var cache_sequence1Default gopurs_runtime.Value
var once_sequence1Default sync.Once
func Get_sequence1Default() gopurs_runtime.Value {
	once_sequence1Default.Do(func() {
		cache_sequence1Default = gopurs_runtime.Func2(func(dictTraversable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1Default(dictTraversable1_0_box, dictApply_1_box)
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
}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Dual.Get_Dual(), gopurs_runtime.Apply(f_1, v_2))
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
}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), gopurs_runtime.Apply(f_1, v_2))
}))
	})
	return cache_traversableMultiplicative
}

var cache_sequence1 gopurs_runtime.Value
var once_sequence1 sync.Once
func Get_sequence1() gopurs_runtime.Value {
	once_sequence1.Do(func() {
		cache_sequence1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1(dict_0_box)
})
	})
	return cache_sequence1
}

var cache_traverse1Default gopurs_runtime.Value
var once_traverse1Default sync.Once
func Get_traverse1Default() gopurs_runtime.Value {
	once_traverse1Default.Do(func() {
		cache_traverse1Default = gopurs_runtime.Func2(func(dictTraversable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1Default(dictTraversable1_0_box, dictApply_1_box)
})
	})
	return cache_traverse1Default
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_traverse1(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "traverse1")
}

func Call_sequence1Default(dictTraversable1_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable1_0 gopurs_runtime.Value = dictTraversable1_0_loop
_ = dictTraversable1_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_0, "traverse1"), dictApply_1, Get_identity())
}

func Call_sequence1(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "sequence1")
}

func Call_traverse1Default(dictTraversable1_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable1_0 gopurs_runtime.Value = dictTraversable1_0_loop
_ = dictTraversable1_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
sequence12_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_0, "sequence1"), dictApply_1)
_ = sequence12_2_0
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence12_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_0, "Traversable1"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), f_3, ta_4))
})
}


