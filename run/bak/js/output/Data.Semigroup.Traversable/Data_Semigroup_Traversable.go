package Data_Semigroup_Traversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
)

var traverse1 gopurs_runtime.Value
var once_traverse1 sync.Once
func Get_traverse1() gopurs_runtime.Value {
	once_traverse1.Do(func() {
		traverse1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "traverse1")
})
	})
	return traverse1
}

var traversableTuple gopurs_runtime.Value
var once_traversableTuple sync.Once
func Get_traversableTuple() gopurs_runtime.Value {
	once_traversableTuple.Do(func() {
		traversableTuple = gopurs_runtime.RecordDict4("traverse1", "sequence1", "Foldable10", "Traversable1", gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]), gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
}), gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1])
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableTuple()
}))
	})
	return traversableTuple
}

var traversableIdentity gopurs_runtime.Value
var once_traversableIdentity sync.Once
func Get_traversableIdentity() gopurs_runtime.Value {
	once_traversableIdentity.Do(func() {
		traversableIdentity = gopurs_runtime.RecordDict4("traverse1", "sequence1", "Foldable10", "Traversable1", gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableIdentity()
}))
	})
	return traversableIdentity
}

var sequence1Default gopurs_runtime.Value
var once_sequence1Default sync.Once
func Get_sequence1Default() gopurs_runtime.Value {
	once_sequence1Default.Do(func() {
		sequence1Default = gopurs_runtime.Func2(func(dictTraversable1_0 gopurs_runtime.Value, dictApply_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_0, "traverse1"), dictApply_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return sequence1Default
}

var traversableDual gopurs_runtime.Value
var once_traversableDual sync.Once
func Get_traversableDual() gopurs_runtime.Value {
	once_traversableDual.Do(func() {
		traversableDual = gopurs_runtime.RecordDict4("traverse1", "sequence1", "Foldable10", "Traversable1", gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Dual.Get_Dual(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableDual(), "traverse1"), dictApply_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableDual()
}))
	})
	return traversableDual
}

var traversableMultiplicative gopurs_runtime.Value
var once_traversableMultiplicative sync.Once
func Get_traversableMultiplicative() gopurs_runtime.Value {
	once_traversableMultiplicative.Do(func() {
		traversableMultiplicative = gopurs_runtime.RecordDict4("traverse1", "sequence1", "Foldable10", "Traversable1", gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMultiplicative(), "traverse1"), dictApply_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableMultiplicative()
}))
	})
	return traversableMultiplicative
}

var sequence1 gopurs_runtime.Value
var once_sequence1 sync.Once
func Get_sequence1() gopurs_runtime.Value {
	once_sequence1.Do(func() {
		sequence1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "sequence1")
})
	})
	return sequence1
}

var traverse1Default gopurs_runtime.Value
var once_traverse1Default sync.Once
func Get_traverse1Default() gopurs_runtime.Value {
	once_traverse1Default.Do(func() {
		traverse1Default = gopurs_runtime.Func2(func(dictTraversable1_0 gopurs_runtime.Value, dictApply_1 gopurs_runtime.Value) gopurs_runtime.Value {
sequence12_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_0, "sequence1"), dictApply_1)
_ = sequence12_2_0
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence12_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_0, "Traversable1"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), f_3, ta_4))
})
})
	})
	return traverse1Default
}




