package Data_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var NonEmpty gopurs_runtime.Value
var once_NonEmpty sync.Once
func Get_NonEmpty() gopurs_runtime.Value {
	once_NonEmpty.Do(func() {
		NonEmpty = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", value0, value1)
})
})
	})
	return NonEmpty
}

var unfoldable1NonEmpty gopurs_runtime.Value
var once_unfoldable1NonEmpty sync.Once
func Get_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_unfoldable1NonEmpty.Do(func() {
		unfoldable1NonEmpty = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(f_1, b_2)
_ = __local_var_3_0
return gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[1]))
}))
})
	})
	return unfoldable1NonEmpty
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
})
	})
	return tail
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", a_2, empty_1_0)
})
})
	})
	return singleton
}

var showNonEmpty gopurs_runtime.Value
var once_showNonEmpty sync.Once
func Get_showNonEmpty() gopurs_runtime.Value {
	once_showNonEmpty.Do(func() {
		showNonEmpty = gopurs_runtime.Func2(func(dictShow_0 gopurs_runtime.Value, dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(NonEmpty " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]).StrVal + ")")
}))
})
	})
	return showNonEmpty
}

var semigroupNonEmpty gopurs_runtime.Value
var once_semigroupNonEmpty sync.Once
func Get_semigroupNonEmpty() gopurs_runtime.Value {
	once_semigroupNonEmpty.Do(func() {
		semigroupNonEmpty = gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1], gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1])))
}))
})
	})
	return semigroupNonEmpty
}

var oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		oneOf = gopurs_runtime.Func2(func(dictAlternative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1])
})
	})
	return oneOf
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
})
	})
	return head
}

var functorNonEmpty gopurs_runtime.Value
var once_functorNonEmpty sync.Once
func Get_functorNonEmpty() gopurs_runtime.Value {
	once_functorNonEmpty.Do(func() {
		functorNonEmpty = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[1]))
}))
})
	})
	return functorNonEmpty
}

var functorWithIndex gopurs_runtime.Value
var once_functorWithIndex sync.Once
func Get_functorWithIndex() gopurs_runtime.Value {
	once_functorWithIndex.Do(func() {
		functorWithIndex = gopurs_runtime.Func(func(dictFunctorWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorNonEmpty1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(m_3.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, (*[1024]gopurs_runtime.Value)(m_3.UnsafePtr)[1]))
}))
_ = functorNonEmpty1_2_1
return gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply2(f_3, gopurs_runtime.Constructor0("Nothing"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Constructor1("Just", x_5))
}), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_2_1
}))
})
	})
	return functorWithIndex
}

var fromNonEmpty gopurs_runtime.Value
var once_fromNonEmpty sync.Once
func Get_fromNonEmpty() gopurs_runtime.Value {
	once_fromNonEmpty.Do(func() {
		fromNonEmpty = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1])
})
	})
	return fromNonEmpty
}

var foldableNonEmpty gopurs_runtime.Value
var once_foldableNonEmpty sync.Once
func Get_foldableNonEmpty() gopurs_runtime.Value {
	once_foldableNonEmpty.Do(func() {
		foldableNonEmpty = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
_ = foldMap1_2_0
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), gopurs_runtime.Apply2(foldMap1_2_0, f_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
})
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1])
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
}))
})
	})
	return foldableNonEmpty
}

var foldableWithIndexNonEmpty gopurs_runtime.Value
var once_foldableWithIndexNonEmpty sync.Once
func Get_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_foldableWithIndexNonEmpty.Do(func() {
		foldableWithIndexNonEmpty = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_0
foldableNonEmpty1_2_1 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "foldMap"), dictMonoid_2)
_ = foldMap1_3_2
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]), gopurs_runtime.Apply2(foldMap1_3_2, f_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]))
})
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), f_2, gopurs_runtime.Apply2(f_2, b_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1])
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldr"), f_2, b_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
}))
_ = foldableNonEmpty1_2_1
return gopurs_runtime.RecordDict4("foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", "Foldable0", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), dictMonoid_3)
_ = foldMapWithIndex1_4_3
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_5, gopurs_runtime.Constructor0("Nothing"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]), gopurs_runtime.Apply2(foldMapWithIndex1_4_3, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Constructor1("Just", x_7))
}), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1]))
})
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Constructor1("Just", x_6))
}), gopurs_runtime.Apply3(f_3, gopurs_runtime.Constructor0("Nothing"), b_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1])
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_3, gopurs_runtime.Constructor0("Nothing"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Constructor1("Just", x_6))
}), b_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_2_1
}))
})
	})
	return foldableWithIndexNonEmpty
}

var traversableNonEmpty gopurs_runtime.Value
var once_traversableNonEmpty sync.Once
func Get_traversableNonEmpty() gopurs_runtime.Value {
	once_traversableNonEmpty.Do(func() {
		traversableNonEmpty = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorNonEmpty1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(m_3.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, (*[1024]gopurs_runtime.Value)(m_3.UnsafePtr)[1]))
}))
_ = functorNonEmpty1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_3_2
foldableNonEmpty1_4_3 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldMap"), dictMonoid_4)
_ = foldMap1_5_4
return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0]), gopurs_runtime.Apply2(foldMap1_5_4, f_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1]))
})
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldl"), f_4, gopurs_runtime.Apply2(f_4, b_5, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1])
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_4, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldr"), f_4, b_5, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1]))
}))
_ = foldableNonEmpty1_4_3
return gopurs_runtime.RecordDict4("sequence", "traverse", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_5
sequence1_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_5)
_ = sequence1_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_6_5, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_5, "Functor0"), gopurs_runtime.Value{}), "map"), Get_NonEmpty(), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0]), gopurs_runtime.Apply(sequence1_7_6, (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1]))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_7
traverse1_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_5)
_ = traverse1_7_8
return gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_6_7, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_7, "Functor0"), gopurs_runtime.Value{}), "map"), Get_NonEmpty(), gopurs_runtime.Apply(f_8, (*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[0])), gopurs_runtime.Apply2(traverse1_7_8, f_8, (*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[1]))
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_4_3
}))
})
	})
	return traversableNonEmpty
}

var traversableWithIndexNonEmpty gopurs_runtime.Value
var once_traversableWithIndexNonEmpty sync.Once
func Get_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_traversableWithIndexNonEmpty.Do(func() {
		traversableWithIndexNonEmpty = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorNonEmpty1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(m_4.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_3, (*[1024]gopurs_runtime.Value)(m_4.UnsafePtr)[1]))
}))
_ = functorNonEmpty1_3_3
functorWithIndex1_3_2 := gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply2(f_4, gopurs_runtime.Constructor0("Nothing"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Constructor1("Just", x_6))
}), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_3_3
}))
_ = functorWithIndex1_3_2
foldableWithIndexNonEmpty1_4_4 := gopurs_runtime.Apply(Get_foldableWithIndexNonEmpty(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexNonEmpty1_4_4
traversableNonEmpty1_5_5 := gopurs_runtime.Apply(Get_traversableNonEmpty(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableNonEmpty1_5_5
return gopurs_runtime.RecordDict4("traverseWithIndex", "FunctorWithIndex0", "FoldableWithIndex1", "Traversable2", gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_6
traverseWithIndex1_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_6)
_ = traverseWithIndex1_8_7
return gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_6, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_6, "Functor0"), gopurs_runtime.Value{}), "map"), Get_NonEmpty(), gopurs_runtime.Apply2(f_9, gopurs_runtime.Constructor0("Nothing"), (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[0])), gopurs_runtime.Apply2(traverseWithIndex1_8_7, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_9, gopurs_runtime.Constructor1("Just", x_11))
}), (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[1]))
})
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndex1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexNonEmpty1_4_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableNonEmpty1_5_5
}))
})
	})
	return traversableWithIndexNonEmpty
}

var foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		foldable1NonEmpty = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldableNonEmpty1_1_0 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
_ = foldMap1_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), gopurs_runtime.Apply2(foldMap1_2_1, f_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
})
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1])
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
}))
_ = foldableNonEmpty1_1_0
return gopurs_runtime.RecordDict4("foldMap1", "foldr1", "foldl1", "Foldable0", gopurs_runtime.Func3(func(dictSemigroup_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(s_5 gopurs_runtime.Value, a1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), s_5, gopurs_runtime.Apply(f_3, a1_6))
}), gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1])
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0])
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(f_2, a1_5)
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_7.StrVal == "Nothing").IntVal != 0 {
__t5 = a1_5
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(x_7.StrVal == "Just").IntVal != 0 {
__t5 = gopurs_runtime.Apply(__local_var_6_4, (*[1024]gopurs_runtime.Value)(x_7.UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Constructor1("Just", __t5)
})
}), gopurs_runtime.Constructor0("Nothing"), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1])
_ = __local_var_5_3
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_3.StrVal == "Nothing").IntVal != 0 {
__t6 = (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_3.StrVal == "Just").IntVal != 0 {
__t6 = gopurs_runtime.Apply(__local_var_4_2, (*[1024]gopurs_runtime.Value)(__local_var_5_3.UnsafePtr)[0])
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1])
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_1_0
}))
})
	})
	return foldable1NonEmpty
}

var foldl1 gopurs_runtime.Value
var once_foldl1 sync.Once
func Get_foldl1() gopurs_runtime.Value {
	once_foldl1.Do(func() {
		foldl1 = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_foldable1NonEmpty(), dictFoldable_0), "foldl1")
})
	})
	return foldl1
}

var eqNonEmpty gopurs_runtime.Value
var once_eqNonEmpty sync.Once
func Get_eqNonEmpty() gopurs_runtime.Value {
	once_eqNonEmpty.Do(func() {
		eqNonEmpty = gopurs_runtime.Func2(func(dictEq1_0 gopurs_runtime.Value, dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_1)
_ = eq11_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(eq11_2_0, (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[1]).IntVal != 0)
}))
})
	})
	return eqNonEmpty
}

var ordNonEmpty gopurs_runtime.Value
var once_ordNonEmpty sync.Once
func Get_ordNonEmpty() gopurs_runtime.Value {
	once_ordNonEmpty.Do(func() {
		ordNonEmpty = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), dictOrd_2)
_ = compare11_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_2
eq11_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), __local_var_4_2)
_ = eq11_5_3
eqNonEmpty2_6_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "eq"), (*[1024]gopurs_runtime.Value)(x_6.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(eq11_5_3, (*[1024]gopurs_runtime.Value)(x_6.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[1]).IntVal != 0)
}))
_ = eqNonEmpty2_6_4
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
v_9_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_2, "compare"), (*[1024]gopurs_runtime.Value)(x_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0])
_ = v_9_5
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_9_5.StrVal == "LT").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("LT")
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(v_9_5.StrVal == "GT").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("GT")
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply2(compare11_3_1, (*[1024]gopurs_runtime.Value)(x_7.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[1])
}
end_branch_6:
return __t6
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqNonEmpty2_6_4
}))
})
})
	})
	return ordNonEmpty
}

var eq1NonEmpty gopurs_runtime.Value
var once_eq1NonEmpty sync.Once
func Get_eq1NonEmpty() gopurs_runtime.Value {
	once_eq1NonEmpty.Do(func() {
		eq1NonEmpty = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_1)
_ = eq11_2_0
return gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(eq11_2_0, (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[1]).IntVal != 0)
})
}))
})
	})
	return eq1NonEmpty
}

var ord1NonEmpty gopurs_runtime.Value
var once_ord1NonEmpty sync.Once
func Get_ord1NonEmpty() gopurs_runtime.Value {
	once_ord1NonEmpty.Do(func() {
		ord1NonEmpty = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordNonEmpty1_1_0 := gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd1_0)
_ = ordNonEmpty1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1NonEmpty1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), dictEq_3)
_ = eq11_4_3
return gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_3, "eq"), (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(eq11_4_3, (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[1]).IntVal != 0)
})
}))
_ = eq1NonEmpty1_3_2
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordNonEmpty1_1_0, dictOrd_4), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1NonEmpty1_3_2
}))
})
	})
	return ord1NonEmpty
}




