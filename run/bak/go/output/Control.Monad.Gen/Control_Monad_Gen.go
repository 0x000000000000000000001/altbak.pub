package Control_Monad_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Semigroup_Last "gopurs/output/Data.Semigroup.Last"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var monoidAdditive gopurs_runtime.Value
var once_monoidAdditive sync.Once
func Get_monoidAdditive() gopurs_runtime.Value {
	once_monoidAdditive.Do(func() {
		monoidAdditive = func() gopurs_runtime.Value {
semigroupAdditive1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.FloatAdd(v_0, v1_1)
}))
_ = semigroupAdditive1_0_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_0_0
}))
}()
	})
	return monoidAdditive
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

var Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		Nil = gopurs_runtime.Constructor0("Nil")
	})
	return Nil
}

var unfoldable gopurs_runtime.Value
var once_unfoldable sync.Once
func Get_unfoldable() gopurs_runtime.Value {
	once_unfoldable.Do(func() {
		unfoldable = gopurs_runtime.Func2(func(dictMonadRec_0 gopurs_runtime.Value, dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_1
Bind1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_4_2
return gopurs_runtime.Func2(func(dictUnfoldable_5 gopurs_runtime.Value, gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1].IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.Constructor1("Done", (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0]))
goto end_branch_7
} else {

}
}
{
__local_var_8_5 := (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0]
_ = __local_var_8_5
__local_var_9_6 := (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1]
_ = __local_var_9_6
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_4_2, "bind"), gen_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.Constructor1("Loop", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor2("Cons", x_10, __local_var_8_5), gopurs_runtime.Int(__local_var_9_6.IntVal - gopurs_runtime.Int(1).IntVal))))
}))
}
end_branch_7:
return __t7
}))
_ = __local_var_7_4
__local_var_8_8 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Constructor0("Nil"))
_ = __local_var_8_8
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_4_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_5, "unfoldr"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7.StrVal == "Nil")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_7.StrVal == "Cons")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1]))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(__local_var_8_8, x_9))
})))
})
})
	})
	return unfoldable
}

var semigroupFreqSemigroup gopurs_runtime.Value
var once_semigroupFreqSemigroup sync.Once
func Get_semigroupFreqSemigroup() gopurs_runtime.Value {
	once_semigroupFreqSemigroup.Do(func() {
		semigroupFreqSemigroup = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.Apply(v_0, pos_2)
_ = v2_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v2_3_0.UnsafePtr)[0].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(v1_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v2_3_0.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = v2_3_0
}
end_branch_1:
return __t1
}))
	})
	return semigroupFreqSemigroup
}

var fromIndex gopurs_runtime.Value
var once_fromIndex sync.Once
func Get_fromIndex() gopurs_runtime.Value {
	once_fromIndex.Do(func() {
		fromIndex = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), pkg_Data_Semigroup_Last.Get_semigroupLast())
_ = foldMap1_1_0
return gopurs_runtime.Func2(func(i_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_1:
for {
if false { continue go__4_1 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_6.StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1].StrVal == "Nil")).IntVal != 0 {
__t3 = (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_5.IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t3 = (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Int(v_5.IntVal - gopurs_runtime.Int(1).IntVal)
v1_6_loop = (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]
continue go__4_1
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v1_6.StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.Apply2(foldMap1_1_0, pkg_Data_Semigroup_Last.Get_Last(), xs_3)
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
return gopurs_runtime.Apply2(go__4_1, i_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "Foldable0"), gopurs_runtime.Value{}), "foldr"), Get_Cons(), gopurs_runtime.Constructor0("Nil"), xs_3))
})
})
	})
	return fromIndex
}

var oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		oneOf = gopurs_runtime.Func2(func(dictMonadGen_0 gopurs_runtime.Value, dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
length_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "Foldable0"), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func2(func(c_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal + c_2.IntVal)
}), gopurs_runtime.Int(0))
_ = length_2_0
fromIndex1_3_1 := gopurs_runtime.Apply(Get_fromIndex(), dictFoldable1_1)
_ = fromIndex1_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply(length_2_0, xs_4).IntVal - gopurs_runtime.Int(1).IntVal)), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(fromIndex1_3_1, n_5, xs_4)
}))
})
})
	})
	return oneOf
}

var freqSemigroup gopurs_runtime.Value
var once_freqSemigroup sync.Once
func Get_freqSemigroup() gopurs_runtime.Value {
	once_freqSemigroup.Do(func() {
		freqSemigroup = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
_ = __local_var_1_0
__local_var_2_1 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
_ = __local_var_2_1
return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.FloatGte(pos_3, __local_var_1_0)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor1("Just", gopurs_runtime.FloatSub(pos_3, __local_var_1_0)), __local_var_2_1)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor0("Nothing"), __local_var_2_1)
}
end_branch_2:
return __t2
})
})
	})
	return freqSemigroup
}

var frequency gopurs_runtime.Value
var once_frequency sync.Once
func Get_frequency() gopurs_runtime.Value {
	once_frequency.Do(func() {
		frequency = gopurs_runtime.Func2(func(dictMonadGen_0 gopurs_runtime.Value, dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "Foldable0"), gopurs_runtime.Value{}), "foldMap"), Get_monoidAdditive())
_ = foldMap_2_0
foldMap1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap1"), Get_semigroupFreqSemigroup())
_ = foldMap1_3_1
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply2(foldMap1_3_1, Get_freqSemigroup(), xs_4)
_ = __local_var_5_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Apply2(foldMap_2_0, pkg_Data_Tuple.Get_fst(), xs_4)), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(__local_var_5_2, x_6).UnsafePtr)[1]
}))
})
})
	})
	return frequency
}

var filtered gopurs_runtime.Value
var once_filtered sync.Once
func Get_filtered() gopurs_runtime.Value {
	once_filtered.Do(func() {
		filtered = gopurs_runtime.Func2(func(dictMonadRec_0 gopurs_runtime.Value, dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Func(func(gen_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "map"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(a_5.StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Loop", pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(a_5.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Done", (*[1024]gopurs_runtime.Value)(a_5.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gen_3)
}), pkg_Data_Unit.Get_unit())
})
})
	})
	return filtered
}

var suchThat gopurs_runtime.Value
var once_suchThat sync.Once
func Get_suchThat() gopurs_runtime.Value {
	once_suchThat.Do(func() {
		suchThat = gopurs_runtime.Func2(func(dictMonadRec_0 gopurs_runtime.Value, dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
filtered2_2_0 := gopurs_runtime.Apply2(Get_filtered(), dictMonadRec_0, dictMonadGen_1)
_ = filtered2_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Func2(func(gen_4 gopurs_runtime.Value, pred_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(filtered2_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "map"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pred_5, a_6)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", a_6)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_2:
return __t2
}), gen_4))
})
})
	})
	return suchThat
}

var elements gopurs_runtime.Value
var once_elements sync.Once
func Get_elements() gopurs_runtime.Value {
	once_elements.Do(func() {
		elements = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
length_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "Foldable0"), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func2(func(c_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal + c_3.IntVal)
}), gopurs_runtime.Int(0))
_ = length_3_1
fromIndex1_4_2 := gopurs_runtime.Apply(Get_fromIndex(), dictFoldable1_2)
_ = fromIndex1_4_2
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply(length_3_1, xs_5).IntVal - gopurs_runtime.Int(1).IntVal)), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply2(fromIndex1_4_2, n_6, xs_5))
}))
})
})
})
	})
	return elements
}

var choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		choose = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
chooseBool_1_0 := gopurs_runtime.RecordGet(dictMonadGen_0, "chooseBool")
_ = chooseBool_1_0
return gopurs_runtime.Func2(func(genA_2 gopurs_runtime.Value, genB_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), chooseBool_1_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4).IntVal != 0 {
__t1 = genA_2
goto end_branch_1
} else {

}
}
{
__t1 = genB_3
}
end_branch_1:
return __t1
}))
})
})
	})
	return choose
}


