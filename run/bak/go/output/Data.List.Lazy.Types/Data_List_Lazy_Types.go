package Data_List_Lazy_Types

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var List gopurs_runtime.Value
var once_List sync.Once
func Get_List() gopurs_runtime.Value {
	once_List.Do(func() {
		List = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return List
}

var Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		Nil = gopurs_runtime.Constructor0("Nil")
	})
	return Nil
}

var Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", value0, value1)
})
})
	})
	return Cons
}

var NonEmptyList gopurs_runtime.Value
var once_NonEmptyList sync.Once
func Get_NonEmptyList() gopurs_runtime.Value {
	once_NonEmptyList.Do(func() {
		NonEmptyList = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return NonEmptyList
}

var nil gopurs_runtime.Value
var once_nil sync.Once
func Get_nil() gopurs_runtime.Value {
	once_nil.Do(func() {
		nil = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nil")
}))
	})
	return nil
}

var newtypeNonEmptyList gopurs_runtime.Value
var once_newtypeNonEmptyList sync.Once
func Get_newtypeNonEmptyList() gopurs_runtime.Value {
	once_newtypeNonEmptyList.Do(func() {
		newtypeNonEmptyList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeNonEmptyList
}

var newtypeList gopurs_runtime.Value
var once_newtypeList sync.Once
func Get_newtypeList() gopurs_runtime.Value {
	once_newtypeList.Do(func() {
		newtypeList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeList
}

var step gopurs_runtime.Value
var once_step sync.Once
func Get_step() gopurs_runtime.Value {
	once_step.Do(func() {
		step = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0)
})
	})
	return step
}

var semigroupList gopurs_runtime.Value
var once_semigroupList sync.Once
func Get_semigroupList() gopurs_runtime.Value {
	once_semigroupList.Do(func() {
		semigroupList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.ConstructorGet(__local_var_3_0, 0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.ConstructorGet(__local_var_3_0, 1), ys_1))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}))
	})
	return semigroupList
}

var monoidList gopurs_runtime.Value
var once_monoidList sync.Once
func Get_monoidList() gopurs_runtime.Value {
	once_monoidList.Do(func() {
		monoidList = gopurs_runtime.RecordDict2("mempty", "Semigroup0", Get_nil(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupList()
}))
	})
	return monoidList
}

var lazyList gopurs_runtime.Value
var once_lazyList sync.Once
func Get_lazyList() gopurs_runtime.Value {
	once_lazyList.Do(func() {
		lazyList = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(f_0, x_1))
}))
}))
	})
	return lazyList
}

var functorList gopurs_runtime.Value
var once_functorList sync.Once
func Get_functorList() gopurs_runtime.Value {
	once_functorList.Do(func() {
		functorList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nil")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(f_0, gopurs_runtime.ConstructorGet(__local_var_3_0, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, gopurs_runtime.ConstructorGet(__local_var_3_0, 1)))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}))
	})
	return functorList
}

var functorNonEmptyList gopurs_runtime.Value
var once_functorNonEmptyList sync.Once
func Get_functorNonEmptyList() gopurs_runtime.Value {
	once_functorNonEmptyList.Do(func() {
		functorNonEmptyList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1)
_ = __local_var_3_0
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_0, gopurs_runtime.ConstructorGet(__local_var_3_0, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, gopurs_runtime.ConstructorGet(__local_var_3_0, 1)))
}))
}))
	})
	return functorNonEmptyList
}

var eq1List gopurs_runtime.Value
var once_eq1List sync.Once
func Get_eq1List() gopurs_runtime.Value {
	once_eq1List.Do(func() {
		eq1List = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(v1_5.StrVal == "Nil")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.ConstructorGet(v_4, 0), gopurs_runtime.ConstructorGet(v1_5, 0)).IntVal != 0 && gopurs_runtime.Apply2(go__3_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.ConstructorGet(v_4, 1)), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.ConstructorGet(v1_5, 1))).IntVal != 0).IntVal != 0).IntVal != 0)
}
end_branch_1:
return __t1
})
return gopurs_runtime.Apply2(go__3_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_2))
}))
	})
	return eq1List
}

var eqNonEmpty gopurs_runtime.Value
var once_eqNonEmpty sync.Once
func Get_eqNonEmpty() gopurs_runtime.Value {
	once_eqNonEmpty.Do(func() {
		eqNonEmpty = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.ConstructorGet(x_1, 0), gopurs_runtime.ConstructorGet(y_2, 0)).IntVal != 0 && gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_eq1List(), "eq1"), dictEq_0, gopurs_runtime.ConstructorGet(x_1, 1), gopurs_runtime.ConstructorGet(y_2, 1)).IntVal != 0)
}))
})
	})
	return eqNonEmpty
}

var eq1NonEmptyList gopurs_runtime.Value
var once_eq1NonEmptyList sync.Once
func Get_eq1NonEmptyList() gopurs_runtime.Value {
	once_eq1NonEmptyList.Do(func() {
		eq1NonEmptyList = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_eq1Lazy(), "eq1"), gopurs_runtime.Apply(Get_eqNonEmpty(), dictEq_0))
}))
	})
	return eq1NonEmptyList
}

var eqList gopurs_runtime.Value
var once_eqList sync.Once
func Get_eqList() gopurs_runtime.Value {
	once_eqList.Do(func() {
		eqList = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_eq1List(), "eq1"), dictEq_0))
})
	})
	return eqList
}

var eqNonEmptyList gopurs_runtime.Value
var once_eqNonEmptyList sync.Once
func Get_eqNonEmptyList() gopurs_runtime.Value {
	once_eqNonEmptyList.Do(func() {
		eqNonEmptyList = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_eqNonEmpty(), dictEq_0)
_ = __local_var_1_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_2), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), y_3))
}))
})
	})
	return eqNonEmptyList
}

var ord1List gopurs_runtime.Value
var once_ord1List sync.Once
func Get_ord1List() gopurs_runtime.Value {
	once_ord1List.Do(func() {
		ord1List = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.StrVal == "Nil")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_5.StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("EQ")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("LT")
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_5.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("GT")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0)).IntVal != 0 {
v2_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.ConstructorGet(v_4, 0), gopurs_runtime.ConstructorGet(v1_5, 0))
_ = v2_6_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_6_3.StrVal == "EQ")).IntVal != 0 {
v_4_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.ConstructorGet(v_4, 1))
v1_5_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.ConstructorGet(v1_5, 1))
continue go__3_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = v2_6_3
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__3_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_2))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1List()
}))
	})
	return ord1List
}

var ordNonEmpty gopurs_runtime.Value
var once_ordNonEmpty sync.Once
func Get_ordNonEmpty() gopurs_runtime.Value {
	once_ordNonEmpty.Do(func() {
		ordNonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_ordNonEmpty(), Get_ord1List())
	})
	return ordNonEmpty
}

var ord1NonEmptyList gopurs_runtime.Value
var once_ord1NonEmptyList sync.Once
func Get_ord1NonEmptyList() gopurs_runtime.Value {
	once_ord1NonEmptyList.Do(func() {
		ord1NonEmptyList = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_ordLazy(), gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd_0)), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1NonEmptyList()
}))
	})
	return ord1NonEmptyList
}

var ordList gopurs_runtime.Value
var once_ordList sync.Once
func Get_ordList() gopurs_runtime.Value {
	once_ordList.Do(func() {
		ordList = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_ord1List(), "compare1"), dictOrd_0), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_eq1List(), "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
}))
})
	})
	return ordList
}

var ordNonEmptyList gopurs_runtime.Value
var once_ordNonEmptyList sync.Once
func Get_ordNonEmptyList() gopurs_runtime.Value {
	once_ordNonEmptyList.Do(func() {
		ordNonEmptyList = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_ordLazy(), gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd_0))
})
	})
	return ordNonEmptyList
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", x_0, xs_1)
}))
})
	})
	return cons
}

var foldableList gopurs_runtime.Value
var once_foldableList sync.Once
func Get_foldableList() gopurs_runtime.Value {
	once_foldableList.Do(func() {
		foldableList = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(op_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", a_4, b_3)
}))
}), Get_nil(), xs_2))
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var b_2 = b_2_loop
_ = b_2
var xs_3 = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4_1.StrVal == "Nil")).IntVal != 0 {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_4_1.StrVal == "Cons")).IntVal != 0 {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, gopurs_runtime.ConstructorGet(v_4_1, 0))
xs_3_loop = gopurs_runtime.ConstructorGet(v_4_1, 1)
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return go__1_0
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_3 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_3
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), b_3, gopurs_runtime.Apply(f_2, a_4))
}), mempty_1_3)
})
}))
	})
	return foldableList
}

var foldableNonEmpty gopurs_runtime.Value
var once_foldableNonEmpty sync.Once
func Get_foldableNonEmpty() gopurs_runtime.Value {
	once_foldableNonEmpty.Do(func() {
		foldableNonEmpty = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableList(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_2, gopurs_runtime.ConstructorGet(v_3, 0)), gopurs_runtime.Apply2(foldMap1_1_0, f_2, gopurs_runtime.ConstructorGet(v_3, 1)))
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), f_0, gopurs_runtime.Apply2(f_0, b_1, gopurs_runtime.ConstructorGet(v_2, 0)), gopurs_runtime.ConstructorGet(v_2, 1))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.ConstructorGet(v_2, 0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), f_0, b_1, gopurs_runtime.ConstructorGet(v_2, 1)))
}))
	})
	return foldableNonEmpty
}

var extendList gopurs_runtime.Value
var once_extendList sync.Once
func Get_extendList() gopurs_runtime.Value {
	once_extendList.Do(func() {
		extendList = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2_0.StrVal == "Nil")).IntVal != 0 {
__t1 = Get_nil()
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2_0.StrVal == "Cons")).IntVal != 0 {
__local_var_3_2 := gopurs_runtime.Apply(f_0, l_1)
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.RecordGet(v_5, "acc")
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.RecordGet(v_5, "val")
_ = __local_var_7_5
acc_prime_8_6 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", a_4, __local_var_6_4)
}))
_ = acc_prime_8_6
__local_var_9_7 := gopurs_runtime.Apply(f_0, acc_prime_8_6)
_ = __local_var_9_7
return gopurs_runtime.RecordDict2("val", "acc", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_9_7, __local_var_7_5)
})), acc_prime_8_6)
}), gopurs_runtime.RecordDict2("val", "acc", Get_nil(), Get_nil()), gopurs_runtime.ConstructorGet(v_2_0, 1)), "val")
_ = __local_var_4_3
__t1 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_3_2, __local_var_4_3)
}))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}))
	})
	return extendList
}

var extendNonEmptyList gopurs_runtime.Value
var once_extendNonEmptyList sync.Once
func Get_extendNonEmptyList() gopurs_runtime.Value {
	once_extendNonEmptyList.Do(func() {
		extendNonEmptyList = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.ConstructorGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1), 1)
_ = __local_var_2_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.RecordGet(v1_5, "acc")
_ = __local_var_6_1
__local_var_7_2 := gopurs_runtime.RecordGet(v1_5, "val")
_ = __local_var_7_2
__local_var_8_3 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", a_4, __local_var_6_1)
})))
_ = __local_var_8_3
return gopurs_runtime.RecordDict2("val", "acc", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_8_3, __local_var_7_2)
})), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", a_4, __local_var_6_1)
})))
}), gopurs_runtime.RecordDict2("val", "acc", Get_nil(), Get_nil()), __local_var_2_0), "val"))
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}))
	})
	return extendNonEmptyList
}

var foldableNonEmptyList gopurs_runtime.Value
var once_foldableNonEmptyList sync.Once
func Get_foldableNonEmptyList() gopurs_runtime.Value {
	once_foldableNonEmptyList.Do(func() {
		foldableNonEmptyList = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)
_ = __local_var_3_0
return gopurs_runtime.Apply2(f_0, gopurs_runtime.ConstructorGet(__local_var_3_0, 0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), f_0, b_1, gopurs_runtime.ConstructorGet(__local_var_3_0, 1)))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)
_ = __local_var_3_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), f_0, gopurs_runtime.Apply2(f_0, b_1, gopurs_runtime.ConstructorGet(__local_var_3_1, 0)), gopurs_runtime.ConstructorGet(__local_var_3_1, 1))
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableNonEmpty(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_2
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_1_2, f_2, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_3))
})
}))
	})
	return foldableNonEmptyList
}

var showList gopurs_runtime.Value
var once_showList sync.Once
func Get_showList() gopurs_runtime.Value {
	once_showList.Do(func() {
		showList = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2_0.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Str("(fromFoldable [])")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2_0.StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(fromFoldable [").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.ConstructorGet(v_2_0, 0)).StrVal).StrVal + gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(shown_3 gopurs_runtime.Value, x_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(shown_3.StrVal + gopurs_runtime.Str(",").StrVal).StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), x_prime_4).StrVal)
}), gopurs_runtime.Str(""), gopurs_runtime.ConstructorGet(v_2_0, 1)).StrVal).StrVal + gopurs_runtime.Str("])").StrVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
})
	})
	return showList
}

var showNonEmptyList gopurs_runtime.Value
var once_showNonEmptyList sync.Once
func Get_showNonEmptyList() gopurs_runtime.Value {
	once_showNonEmptyList.Do(func() {
		showNonEmptyList = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_showList(), dictShow_0)
_ = __local_var_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)
_ = __local_var_3_1
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(NonEmptyList (defer \\_ -> (NonEmpty ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.ConstructorGet(__local_var_3_1, 0)).StrVal).StrVal + gopurs_runtime.Str(" ").StrVal).StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "show"), gopurs_runtime.ConstructorGet(__local_var_3_1, 1)).StrVal).StrVal + gopurs_runtime.Str(")))").StrVal)
}))
})
	})
	return showNonEmptyList
}

var showStep gopurs_runtime.Value
var once_showStep sync.Once
func Get_showStep() gopurs_runtime.Value {
	once_showStep.Do(func() {
		showStep = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Nil")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.ConstructorGet(v_1, 0)).StrVal).StrVal + gopurs_runtime.Str(" : ").StrVal).StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_showList(), dictShow_0), "show"), gopurs_runtime.ConstructorGet(v_1, 1)).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
	})
	return showStep
}

var foldableWithIndexList gopurs_runtime.Value
var once_foldableWithIndexList sync.Once
func Get_foldableWithIndexList() gopurs_runtime.Value {
	once_foldableWithIndexList.Do(func() {
		foldableWithIndexList = gopurs_runtime.RecordDict4("foldrWithIndex", "foldlWithIndex", "foldMapWithIndex", "Foldable0", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.ConstructorGet(v1_3, 1)
_ = __local_var_4_1
__local_var_5_2 := gopurs_runtime.ConstructorGet(v1_3, 0)
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int(__local_var_5_2.IntVal + gopurs_runtime.Int(1).IntVal), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", a_6, __local_var_4_1)
})))
})
}), gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int(0), Get_nil()), xs_2)
_ = v_3_0
return gopurs_runtime.ConstructorGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.ConstructorGet(v1_4, 1)
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.ConstructorGet(v1_4, 0)
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int(__local_var_6_4.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_6_4.IntVal - gopurs_runtime.Int(1).IntVal), a_7, __local_var_5_3))
})
}), gopurs_runtime.Constructor2("Tuple", gopurs_runtime.ConstructorGet(v_3_0, 0), b_1), gopurs_runtime.ConstructorGet(v_3_0, 1)), 1)
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_6 := gopurs_runtime.ConstructorGet(v_2, 1)
_ = __local_var_3_6
__local_var_4_7 := gopurs_runtime.ConstructorGet(v_2, 0)
_ = __local_var_4_7
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int(__local_var_4_7.IntVal + gopurs_runtime.Int(1).IntVal), gopurs_runtime.Apply3(f_0, __local_var_4_7, __local_var_3_6, a_5))
})
}), gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int(0), acc_1))
_ = __local_var_2_5
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ConstructorGet(gopurs_runtime.Apply(__local_var_2_5, x_3), 1)
})
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_8 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_8
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func2(func(i_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_4)
_ = __local_var_5_9
__local_var_6_10 := gopurs_runtime.Apply(f_2, i_3)
_ = __local_var_6_10
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_9, gopurs_runtime.Apply(__local_var_6_10, x_7))
})
}), mempty_1_8)
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}))
	})
	return foldableWithIndexList
}

var foldableWithIndexNonEmpty gopurs_runtime.Value
var once_foldableWithIndexNonEmpty sync.Once
func Get_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_foldableWithIndexNonEmpty.Do(func() {
		foldableWithIndexNonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldableWithIndexNonEmpty(), Get_foldableWithIndexList())
	})
	return foldableWithIndexNonEmpty
}

var foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_foldableWithIndexNonEmptyList sync.Once
func Get_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyList.Do(func() {
		foldableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", "Foldable0", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableWithIndexNonEmpty(), "foldMapWithIndex"), dictMonoid_0)
_ = foldMapWithIndex1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_4.StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_4.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal + gopurs_runtime.ConstructorGet(x_4, 0).IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, __t1)
}), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_3))
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableWithIndexNonEmpty(), "foldlWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_3.StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(x_3.StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal + gopurs_runtime.ConstructorGet(x_3, 0).IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, __t2)
}), b_1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableWithIndexNonEmpty(), "foldrWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_3.StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(x_3.StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal + gopurs_runtime.ConstructorGet(x_3, 0).IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Apply(f_0, __t3)
}), b_1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
}))
	})
	return foldableWithIndexNonEmptyList
}

var functorWithIndexList gopurs_runtime.Value
var once_functorWithIndexList sync.Once
func Get_functorWithIndexList() gopurs_runtime.Value {
	once_functorWithIndexList.Do(func() {
		functorWithIndexList = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func3(func(i_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply2(f_0, i_1, x_2)
_ = __local_var_4_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_4_0, acc_3)
}))
}), Get_nil())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}))
	})
	return functorWithIndexList
}

var mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		mapWithIndex = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply2(f_0, gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.ConstructorGet(v_1, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorWithIndexList(), "mapWithIndex"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Constructor1("Just", x_2))
}), gopurs_runtime.ConstructorGet(v_1, 1)))
})
	})
	return mapWithIndex
}

var functorWithIndexNonEmptyList gopurs_runtime.Value
var once_functorWithIndexNonEmptyList sync.Once
func Get_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyList.Do(func() {
		functorWithIndexNonEmptyList = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_mapWithIndex(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_3.StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_3.StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal + gopurs_runtime.ConstructorGet(x_3, 0).IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_0, __t0)
}), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1))
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}))
	})
	return functorWithIndexNonEmptyList
}

var toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		toList = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
v2_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v2_2_0
__local_var_3_1 := gopurs_runtime.ConstructorGet(v2_2_0, 0)
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.ConstructorGet(v2_2_0, 1)
_ = __local_var_4_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_3_1, __local_var_4_2)
})))
}))
})
	})
	return toList
}

var semigroupNonEmptyList gopurs_runtime.Value
var once_semigroupNonEmptyList sync.Once
func Get_semigroupNonEmptyList() gopurs_runtime.Value {
	once_semigroupNonEmptyList.Do(func() {
		semigroupNonEmptyList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_2_0
__local_var_3_1 := gopurs_runtime.ConstructorGet(v1_2_0, 0)
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.ConstructorGet(v1_2_0, 1)
_ = __local_var_4_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", __local_var_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), __local_var_4_2, gopurs_runtime.Apply(Get_toList(), as_prime_1)))
}))
}))
	})
	return semigroupNonEmptyList
}

var traversableList gopurs_runtime.Value
var once_traversableList sync.Once
func Get_traversableList() gopurs_runtime.Value {
	once_traversableList.Do(func() {
		traversableList = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), Get_cons(), gopurs_runtime.Apply(f_2, a_3)), b_4)
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_nil()))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableList(), "traverse"), dictApplicative_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}))
	})
	return traversableList
}

var traversableNonEmpty gopurs_runtime.Value
var once_traversableNonEmpty sync.Once
func Get_traversableNonEmpty() gopurs_runtime.Value {
	once_traversableNonEmpty.Do(func() {
		traversableNonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableNonEmpty(), Get_traversableList())
	})
	return traversableNonEmpty
}

var traversableNonEmptyList gopurs_runtime.Value
var once_traversableNonEmptyList sync.Once
func Get_traversableNonEmptyList() gopurs_runtime.Value {
	once_traversableNonEmptyList.Do(func() {
		traversableNonEmptyList = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableNonEmpty(), "traverse"), dictApplicative_0)
_ = traverse1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return xxs_4
}))
}), gopurs_runtime.Apply2(traverse1_1_0, f_2, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_3)))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableNonEmpty(), "sequence"), dictApplicative_0)
_ = sequence1_1_1
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return xxs_3
}))
}), gopurs_runtime.Apply(sequence1_1_1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_2)))
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
}))
	})
	return traversableNonEmptyList
}

var traversableWithIndexList gopurs_runtime.Value
var once_traversableWithIndexList sync.Once
func Get_traversableWithIndexList() gopurs_runtime.Value {
	once_traversableWithIndexList.Do(func() {
		traversableWithIndexList = gopurs_runtime.RecordDict4("traverseWithIndex", "FunctorWithIndex0", "FoldableWithIndex1", "Traversable2", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func3(func(i_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), Get_cons(), gopurs_runtime.Apply2(f_2, i_3, a_4)), b_5)
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_nil()))
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableList()
}))
	})
	return traversableWithIndexList
}

var traverseWithIndex gopurs_runtime.Value
var once_traverseWithIndex sync.Once
func Get_traverseWithIndex() gopurs_runtime.Value {
	once_traverseWithIndex.Do(func() {
		traverseWithIndex = gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableWithIndexNonEmpty(), Get_traversableWithIndexList()), "traverseWithIndex")
	})
	return traverseWithIndex
}

var traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_traversableWithIndexNonEmptyList sync.Once
func Get_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_traversableWithIndexNonEmptyList.Do(func() {
		traversableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("traverseWithIndex", "FunctorWithIndex0", "FoldableWithIndex1", "Traversable2", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex1_1_0 := gopurs_runtime.Apply(Get_traverseWithIndex(), dictApplicative_0)
_ = traverseWithIndex1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return xxs_4
}))
}), gopurs_runtime.Apply2(traverseWithIndex1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_4.StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_4.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal + gopurs_runtime.ConstructorGet(x_4, 0).IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, __t1)
}), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_3)))
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableNonEmptyList()
}))
	})
	return traversableWithIndexNonEmptyList
}

var unfoldable1List gopurs_runtime.Value
var once_unfoldable1List sync.Once
func Get_unfoldable1List() gopurs_runtime.Value {
	once_unfoldable1List.Do(func() {
		unfoldable1List = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
_ = go__0_0
go__0_0 = gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.Apply(f_1, b_2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v1_4_1, 1).StrVal == "Just")).IntVal != 0 {
__local_var_5_3 := gopurs_runtime.ConstructorGet(v1_4_1, 0)
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.Apply2(go__0_0, f_1, gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v1_4_1, 1), 0))
_ = __local_var_6_4
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_5_3, __local_var_6_4)
}))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v1_4_1, 1).StrVal == "Nothing")).IntVal != 0 {
__local_var_5_5 := gopurs_runtime.ConstructorGet(v1_4_1, 0)
_ = __local_var_5_5
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_5_5, Get_nil())
}))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __t2)
}))
})
return gopurs_runtime.RecordDict1("unfoldr1", go__0_0)
}()
	})
	return unfoldable1List
}

var unfoldableList gopurs_runtime.Value
var once_unfoldableList sync.Once
func Get_unfoldableList() gopurs_runtime.Value {
	once_unfoldableList.Do(func() {
		unfoldableList = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
_ = go__0_0
go__0_0 = gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.Apply(f_1, b_2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_1.StrVal == "Nothing")).IntVal != 0 {
__t2 = Get_nil()
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_1.StrVal == "Just")).IntVal != 0 {
__local_var_5_3 := gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v1_4_1, 0), 0)
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.Apply2(go__0_0, f_1, gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v1_4_1, 0), 1))
_ = __local_var_6_4
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_5_3, __local_var_6_4)
}))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __t2)
}))
})
return gopurs_runtime.RecordDict2("unfoldr", "Unfoldable10", go__0_0, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_unfoldable1List()
}))
}()
	})
	return unfoldableList
}

var unfoldr1 gopurs_runtime.Value
var once_unfoldr1 sync.Once
func Get_unfoldr1() gopurs_runtime.Value {
	once_unfoldr1.Do(func() {
		unfoldr1 = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, b_1)
_ = __local_var_2_0
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.ConstructorGet(__local_var_2_0, 0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_unfoldableList(), "unfoldr"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_0, gopurs_runtime.ConstructorGet(v1_3, 0)))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.ConstructorGet(__local_var_2_0, 1)))
})
	})
	return unfoldr1
}

var unfoldable1NonEmptyList gopurs_runtime.Value
var once_unfoldable1NonEmptyList sync.Once
func Get_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_unfoldable1NonEmptyList.Do(func() {
		unfoldable1NonEmptyList = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_unfoldr1(), f_0, b_1)
}))
}))
	})
	return unfoldable1NonEmptyList
}

var comonadNonEmptyList gopurs_runtime.Value
var once_comonadNonEmptyList sync.Once
func Get_comonadNonEmptyList() gopurs_runtime.Value {
	once_comonadNonEmptyList.Do(func() {
		comonadNonEmptyList = gopurs_runtime.RecordDict2("extract", "Extend0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ConstructorGet(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0), 0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendNonEmptyList()
}))
	})
	return comonadNonEmptyList
}

var monadList gopurs_runtime.Value
var once_monadList sync.Once
func Get_monadList() gopurs_runtime.Value {
	once_monadList.Do(func() {
		monadList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindList()
}))
	})
	return monadList
}

var bindList gopurs_runtime.Value
var once_bindList sync.Once
func Get_bindList() gopurs_runtime.Value {
	once_bindList.Do(func() {
		bindList = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nil")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply(f_1, gopurs_runtime.ConstructorGet(__local_var_3_0, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), gopurs_runtime.ConstructorGet(__local_var_3_0, 1), f_1)))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}))
	})
	return bindList
}

var applyList gopurs_runtime.Value
var once_applyList sync.Once
func Get_applyList() gopurs_runtime.Value {
	once_applyList.Do(func() {
		applyList = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), f_0, gopurs_runtime.Func(func(f_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), a_1, gopurs_runtime.Func(func(a_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeList(), "pure"), gopurs_runtime.Apply(f_prime_2, a_prime_3))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}))
	})
	return applyList
}

var applicativeList gopurs_runtime.Value
var once_applicativeList sync.Once
func Get_applicativeList() gopurs_runtime.Value {
	once_applicativeList.Do(func() {
		applicativeList = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", a_0, Get_nil())
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}))
	})
	return applicativeList
}

var applyNonEmptyList gopurs_runtime.Value
var once_applyNonEmptyList sync.Once
func Get_applyNonEmptyList() gopurs_runtime.Value {
	once_applyNonEmptyList.Do(func() {
		applyNonEmptyList = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
v2_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v1_1)
_ = v2_2_0
v3_3_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v3_3_1
__local_var_4_2 := gopurs_runtime.ConstructorGet(v2_2_0, 0)
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.ConstructorGet(v2_2_0, 1)
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.ConstructorGet(v3_3_1, 0)
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.ConstructorGet(v3_3_1, 1)
_ = __local_var_7_5
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), __local_var_7_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_4_2, Get_nil())
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_6_4, __local_var_7_5)
})), __local_var_5_3)))
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}))
	})
	return applyNonEmptyList
}

var bindNonEmptyList gopurs_runtime.Value
var once_bindNonEmptyList sync.Once
func Get_bindNonEmptyList() gopurs_runtime.Value {
	once_bindNonEmptyList.Do(func() {
		bindNonEmptyList = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_2_0
__local_var_3_1 := gopurs_runtime.ConstructorGet(v1_2_0, 1)
_ = __local_var_3_1
v2_4_2 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(f_1, gopurs_runtime.ConstructorGet(v1_2_0, 0)))
_ = v2_4_2
__local_var_5_3 := gopurs_runtime.ConstructorGet(v2_4_2, 0)
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.ConstructorGet(v2_4_2, 1)
_ = __local_var_6_4
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", __local_var_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), __local_var_6_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), __local_var_3_1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_toList(), gopurs_runtime.Apply(f_1, x_8))
}))))
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}))
	})
	return bindNonEmptyList
}

var altNonEmptyList gopurs_runtime.Value
var once_altNonEmptyList sync.Once
func Get_altNonEmptyList() gopurs_runtime.Value {
	once_altNonEmptyList.Do(func() {
		altNonEmptyList = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.RecordGet(Get_semigroupNonEmptyList(), "append"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}))
	})
	return altNonEmptyList
}

var altList gopurs_runtime.Value
var once_altList sync.Once
func Get_altList() gopurs_runtime.Value {
	once_altList.Do(func() {
		altList = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}))
	})
	return altList
}

var plusList gopurs_runtime.Value
var once_plusList sync.Once
func Get_plusList() gopurs_runtime.Value {
	once_plusList.Do(func() {
		plusList = gopurs_runtime.RecordDict2("empty", "Alt0", Get_nil(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altList()
}))
	})
	return plusList
}

var alternativeList gopurs_runtime.Value
var once_alternativeList sync.Once
func Get_alternativeList() gopurs_runtime.Value {
	once_alternativeList.Do(func() {
		alternativeList = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusList()
}))
	})
	return alternativeList
}

var monadPlusList gopurs_runtime.Value
var once_monadPlusList sync.Once
func Get_monadPlusList() gopurs_runtime.Value {
	once_monadPlusList.Do(func() {
		monadPlusList = gopurs_runtime.RecordDict2("Monad0", "Alternative1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_alternativeList()
}))
	})
	return monadPlusList
}

var applicativeNonEmptyList gopurs_runtime.Value
var once_applicativeNonEmptyList sync.Once
func Get_applicativeNonEmptyList() gopurs_runtime.Value {
	once_applicativeNonEmptyList.Do(func() {
		applicativeNonEmptyList = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", a_0, Get_nil())
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}))
	})
	return applicativeNonEmptyList
}

var monadNonEmptyList gopurs_runtime.Value
var once_monadNonEmptyList sync.Once
func Get_monadNonEmptyList() gopurs_runtime.Value {
	once_monadNonEmptyList.Do(func() {
		monadNonEmptyList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindNonEmptyList()
}))
	})
	return monadNonEmptyList
}


