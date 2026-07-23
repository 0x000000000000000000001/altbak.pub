package Data_Foldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Maybe_First "gopurs/output/Data.Maybe.First"
)

var monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		monoidEndo = func() gopurs_runtime.Value {
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(v1_1, x_2))
}))
_ = semigroupEndo1_0_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}))
}()
	})
	return monoidEndo
}

var Empty gopurs_runtime.Value
var once_Empty sync.Once
func Get_Empty() gopurs_runtime.Value {
	once_Empty.Do(func() {
		Empty = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Empty"))
	})
	return Empty
}

var Node gopurs_runtime.Value
var once_Node sync.Once
func Get_Node() gopurs_runtime.Value {
	once_Node.Do(func() {
		Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Node"), value0)
})
	})
	return Node
}

var Append gopurs_runtime.Value
var once_Append sync.Once
func Get_Append() gopurs_runtime.Value {
	once_Append.Do(func() {
		Append = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Append"), value0, value1)
})
})
	})
	return Append
}

var semigroupFreeMonoidTree gopurs_runtime.Value
var once_semigroupFreeMonoidTree sync.Once
func Get_semigroupFreeMonoidTree() gopurs_runtime.Value {
	once_semigroupFreeMonoidTree.Do(func() {
		semigroupFreeMonoidTree = gopurs_runtime.RecordDict1("append", Get_Append())
	})
	return semigroupFreeMonoidTree
}

var monoidFreeMonoidTree gopurs_runtime.Value
var once_monoidFreeMonoidTree sync.Once
func Get_monoidFreeMonoidTree() gopurs_runtime.Value {
	once_monoidFreeMonoidTree.Do(func() {
		monoidFreeMonoidTree = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Empty")), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupFreeMonoidTree()
}))
	})
	return monoidFreeMonoidTree
}

var foldr gopurs_runtime.Value
var once_foldr sync.Once
func Get_foldr() gopurs_runtime.Value {
	once_foldr.Do(func() {
		foldr = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "foldr")
})
	})
	return foldr
}

var indexr gopurs_runtime.Value
var once_indexr sync.Once
func Get_indexr() gopurs_runtime.Value {
	once_indexr.Do(func() {
		indexr = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, idx_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, cursor_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(cursor_3, "elem"), "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = cursor_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(cursor_3, "pos").IntVal == idx_1.IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), a_2), gopurs_runtime.RecordGet(cursor_3, "pos"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("pos", "elem", gopurs_runtime.Int(gopurs_runtime.RecordGet(cursor_3, "pos").IntVal + gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordGet(cursor_3, "elem"))
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")), gopurs_runtime.Int(0)))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_2_0, x_3), "elem")
})
})
	})
	return indexr
}

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
}), gopurs_runtime.Bool(true))
})
	})
	return null
}

var oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		oneOf = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictPlus_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_1, "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.RecordGet(dictPlus_1, "empty"))
})
	})
	return oneOf
}

var oneOfMap gopurs_runtime.Value
var once_oneOfMap sync.Once
func Get_oneOfMap() gopurs_runtime.Value {
	once_oneOfMap.Do(func() {
		oneOfMap = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictPlus_1 gopurs_runtime.Value) gopurs_runtime.Value {
empty_2_0 := gopurs_runtime.RecordGet(dictPlus_1, "empty")
_ = empty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_1, "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply(f_3, x_4))
}), empty_2_0)
})
})
	})
	return oneOfMap
}

var traverse_ gopurs_runtime.Value
var once_traverse_ sync.Once
func Get_traverse_() gopurs_runtime.Value {
	once_traverse_.Do(func() {
		traverse_ = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
applySecond_1_0 := gopurs_runtime.Apply(pkg_Control_Apply.Get_applySecond(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = applySecond_1_0
return gopurs_runtime.Func2(func(dictFoldable_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(f_3, x_4))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_Unit.Get_unit()))
})
})
	})
	return traverse_
}

var for_ gopurs_runtime.Value
var once_for_ sync.Once
func Get_for_() gopurs_runtime.Value {
	once_for_.Do(func() {
		for_ = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictApplicative_0)
_ = traverse_1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(traverse_1_1_0, dictFoldable_2)
_ = __local_var_3_1
return gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_3_1, a_5, b_4)
})
})
})
	})
	return for_
}

var sequence_ gopurs_runtime.Value
var once_sequence_ sync.Once
func Get_sequence_() gopurs_runtime.Value {
	once_sequence_.Do(func() {
		sequence_ = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictApplicative_0)
_ = traverse_1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(traverse_1_1_0, dictFoldable_2, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
})
	})
	return sequence_
}

var foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		foldl = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "foldl")
})
	})
	return foldl
}

var indexl gopurs_runtime.Value
var once_indexl sync.Once
func Get_indexl() gopurs_runtime.Value {
	once_indexl.Do(func() {
		indexl = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, idx_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(cursor_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(cursor_2, "elem"), "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = cursor_2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(cursor_2, "pos").IntVal == idx_1.IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), a_3), gopurs_runtime.RecordGet(cursor_2, "pos"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("pos", "elem", gopurs_runtime.Int(gopurs_runtime.RecordGet(cursor_2, "pos").IntVal + gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordGet(cursor_2, "elem"))
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")), gopurs_runtime.Int(0)))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_2_0, x_3), "elem")
})
})
	})
	return indexl
}

var intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		intercalate = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_0
mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_3_1
return gopurs_runtime.Func2(func(sep_4 gopurs_runtime.Value, xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_6, "init")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), v1_7)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), sep_4, v1_7)))
}
end_branch_2:
return __t2
}), gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(true), mempty_3_1), xs_5), "acc")
})
})
	})
	return intercalate
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
one_2_0 := gopurs_runtime.RecordGet(dictSemiring_1, "one")
_ = one_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(c_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_1, "add"), one_2_0, c_3)
}), gopurs_runtime.RecordGet(dictSemiring_1, "zero"))
})
	})
	return length
}

var maximumBy gopurs_runtime.Value
var once_maximumBy sync.Once
func Get_maximumBy() gopurs_runtime.Value {
	once_maximumBy.Do(func() {
		maximumBy = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, cmp_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), v1_3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(cmp_1, gopurs_runtime.RecordGet(v_2, "value0"), v1_3), "_tag").StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(v_2, "value0")
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), __t1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
})
	})
	return maximumBy
}

var maximum gopurs_runtime.Value
var once_maximum sync.Once
func Get_maximum() gopurs_runtime.Value {
	once_maximum.Do(func() {
		maximum = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_1, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), v1_3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.RecordGet(v_2, "value0"), v1_3), "_tag").StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(v_2, "value0")
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), __t1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
})
	})
	return maximum
}

var minimumBy gopurs_runtime.Value
var once_minimumBy sync.Once
func Get_minimumBy() gopurs_runtime.Value {
	once_minimumBy.Do(func() {
		minimumBy = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, cmp_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), v1_3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(cmp_1, gopurs_runtime.RecordGet(v_2, "value0"), v1_3), "_tag").StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(v_2, "value0")
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), __t1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
})
	})
	return minimumBy
}

var minimum gopurs_runtime.Value
var once_minimum sync.Once
func Get_minimum() gopurs_runtime.Value {
	once_minimum.Do(func() {
		minimum = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_1, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), v1_3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.RecordGet(v_2, "value0"), v1_3), "_tag").StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(v_2, "value0")
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), __t1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
})
	})
	return minimum
}

var product gopurs_runtime.Value
var once_product sync.Once
func Get_product() gopurs_runtime.Value {
	once_product.Do(func() {
		product = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.RecordGet(dictSemiring_1, "mul"), gopurs_runtime.RecordGet(dictSemiring_1, "one"))
})
	})
	return product
}

var sum gopurs_runtime.Value
var once_sum sync.Once
func Get_sum() gopurs_runtime.Value {
	once_sum.Do(func() {
		sum = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.RecordGet(dictSemiring_1, "add"), gopurs_runtime.RecordGet(dictSemiring_1, "zero"))
})
	})
	return sum
}

var foldableTuple gopurs_runtime.Value
var once_foldableTuple sync.Once
func Get_foldableTuple() gopurs_runtime.Value {
	once_foldableTuple.Do(func() {
		foldableTuple = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(v_2, "value1"), z_1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.RecordGet(v_2, "value1"))
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v_2, "value1"))
}))
	})
	return foldableTuple
}

var foldableMultiplicative gopurs_runtime.Value
var once_foldableMultiplicative sync.Once
func Get_foldableMultiplicative() gopurs_runtime.Value {
	once_foldableMultiplicative.Do(func() {
		foldableMultiplicative = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}))
	})
	return foldableMultiplicative
}

var foldableMaybe gopurs_runtime.Value
var once_foldableMaybe sync.Once
func Get_foldableMaybe() gopurs_runtime.Value {
	once_foldableMaybe.Do(func() {
		foldableMaybe = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply2(v_0, gopurs_runtime.RecordGet(v2_2, "value0"), v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = v1_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(v_0, v1_1, gopurs_runtime.RecordGet(v2_2, "value0"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_2
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t3 = mempty_1_2
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(v_2, gopurs_runtime.RecordGet(v1_3, "value0"))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}))
	})
	return foldableMaybe
}

var foldableIdentity gopurs_runtime.Value
var once_foldableIdentity sync.Once
func Get_foldableIdentity() gopurs_runtime.Value {
	once_foldableIdentity.Do(func() {
		foldableIdentity = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}))
	})
	return foldableIdentity
}

var foldableEither gopurs_runtime.Value
var once_foldableEither sync.Once
func Get_foldableEither() gopurs_runtime.Value {
	once_foldableEither.Do(func() {
		foldableEither = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Left")).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply2(v_0, gopurs_runtime.RecordGet(v2_2, "value0"), v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Left")).IntVal != 0 {
__t1 = v1_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(v_0, v1_1, gopurs_runtime.RecordGet(v2_2, "value0"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_2
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Left")).IntVal != 0 {
__t3 = mempty_1_2
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Right")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(v_2, gopurs_runtime.RecordGet(v1_3, "value0"))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}))
	})
	return foldableEither
}

var foldableDual gopurs_runtime.Value
var once_foldableDual sync.Once
func Get_foldableDual() gopurs_runtime.Value {
	once_foldableDual.Do(func() {
		foldableDual = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}))
	})
	return foldableDual
}

var foldableDisj gopurs_runtime.Value
var once_foldableDisj sync.Once
func Get_foldableDisj() gopurs_runtime.Value {
	once_foldableDisj.Do(func() {
		foldableDisj = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}))
	})
	return foldableDisj
}

var foldableConst gopurs_runtime.Value
var once_foldableConst sync.Once
func Get_foldableConst() gopurs_runtime.Value {
	once_foldableConst.Do(func() {
		foldableConst = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
}))
	})
	return foldableConst
}

var foldableConj gopurs_runtime.Value
var once_foldableConj sync.Once
func Get_foldableConj() gopurs_runtime.Value {
	once_foldableConj.Do(func() {
		foldableConj = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}))
	})
	return foldableConj
}

var foldableAdditive gopurs_runtime.Value
var once_foldableAdditive sync.Once
func Get_foldableAdditive() gopurs_runtime.Value {
	once_foldableAdditive.Do(func() {
		foldableAdditive = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}))
	})
	return foldableAdditive
}

var foldMapDefaultR gopurs_runtime.Value
var once_foldMapDefaultR sync.Once
func Get_foldMapDefaultR() gopurs_runtime.Value {
	once_foldMapDefaultR.Do(func() {
		foldMapDefaultR = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_3, x_4), acc_5)
}), mempty_2_0)
})
})
	})
	return foldMapDefaultR
}

var foldableArray gopurs_runtime.Value
var once_foldableArray sync.Once
func Get_foldableArray() gopurs_runtime.Value {
	once_foldableArray.Do(func() {
		foldableArray = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", Get_foldrArray(), Get_foldlArray(), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableArray(), "foldr"), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_2, x_3), acc_4)
}), mempty_1_0)
})
}))
	})
	return foldableArray
}

var foldableFreeMonoidTree gopurs_runtime.Value
var once_foldableFreeMonoidTree sync.Once
func Get_foldableFreeMonoidTree() gopurs_runtime.Value {
	once_foldableFreeMonoidTree.Do(func() {
		foldableFreeMonoidTree = gopurs_runtime.RecordDict3("foldl", "foldr", "foldMap", gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(acc_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var acc_2 = acc_2_loop
_ = acc_2
var lhs_3 = lhs_3_loop
_ = lhs_3
var rhs_4 = rhs_4_loop
_ = rhs_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(lhs_3, "_tag").StrVal == "Node")).IntVal != 0 {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, gopurs_runtime.RecordGet(lhs_3, "value0"))
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Empty"))
continue go__1_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(lhs_3, "_tag").StrVal == "Append")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(lhs_3, "value1"), "_tag").StrVal == "Empty")).IntVal != 0 {
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.RecordGet(lhs_3, "value0")
rhs_4_loop = rhs_4
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(rhs_4, "_tag").StrVal == "Empty")).IntVal != 0 {
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.RecordGet(lhs_3, "value0")
rhs_4_loop = gopurs_runtime.RecordGet(lhs_3, "value1")
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.RecordGet(lhs_3, "value0")
rhs_4_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Append"), gopurs_runtime.RecordGet(lhs_3, "value1"), rhs_4)
continue go__1_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(lhs_3, "_tag").StrVal == "Empty")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(rhs_4, "_tag").StrVal == "Empty")).IntVal != 0 {
__t3 = acc_2
goto end_branch_3
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Empty"))
continue go__1_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t1 = __t3
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
})
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__1_0, a_2, b_3, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Empty")))
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_4 gopurs_runtime.Value
go__1_4 = gopurs_runtime.Func(func(acc_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_4:
for {
if false { continue go__1_4 }
var acc_2 = acc_2_loop
_ = acc_2
var lhs_3 = lhs_3_loop
_ = lhs_3
var rhs_4 = rhs_4_loop
_ = rhs_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(rhs_4, "_tag").StrVal == "Node")).IntVal != 0 {
acc_2_loop = gopurs_runtime.Apply2(fn_0, gopurs_runtime.RecordGet(rhs_4, "value0"), acc_2)
lhs_3_loop = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Empty"))
rhs_4_loop = lhs_3
continue go__1_4
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(rhs_4, "_tag").StrVal == "Append")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(rhs_4, "value0"), "_tag").StrVal == "Empty")).IntVal != 0 {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = gopurs_runtime.RecordGet(rhs_4, "value1")
continue go__1_4
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(lhs_3, "_tag").StrVal == "Empty")).IntVal != 0 {
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.RecordGet(rhs_4, "value0")
rhs_4_loop = gopurs_runtime.RecordGet(rhs_4, "value1")
continue go__1_4
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Append"), lhs_3, gopurs_runtime.RecordGet(rhs_4, "value0"))
rhs_4_loop = gopurs_runtime.RecordGet(rhs_4, "value1")
continue go__1_4
__t6 = gopurs_runtime.Value{}
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(rhs_4, "_tag").StrVal == "Empty")).IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(lhs_3, "_tag").StrVal == "Empty")).IntVal != 0 {
__t7 = acc_2
goto end_branch_7
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Empty"))
rhs_4_loop = lhs_3
continue go__1_4
__t7 = gopurs_runtime.Value{}
}
end_branch_7:
__t5 = __t7
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
})
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__1_4, a_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Empty")), b_3)
})
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_8 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_8
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_2, x_3), acc_4)
}), mempty_1_8)
})
}))
	})
	return foldableFreeMonoidTree
}

var foldMapDefaultL gopurs_runtime.Value
var once_foldMapDefaultL sync.Once
func Get_foldMapDefaultL() gopurs_runtime.Value {
	once_foldMapDefaultL.Do(func() {
		foldMapDefaultL = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_4, gopurs_runtime.Apply(f_3, x_5))
}), mempty_2_0)
})
})
	})
	return foldMapDefaultL
}

var foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		foldMap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "foldMap")
})
	})
	return foldMap
}

var foldableApp gopurs_runtime.Value
var once_foldableApp sync.Once
func Get_foldableApp() gopurs_runtime.Value {
	once_foldableApp.Do(func() {
		foldableApp = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, i_2, v_3)
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, i_2, v_3)
}), gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
}))
})
	})
	return foldableApp
}

var foldableCompose gopurs_runtime.Value
var once_foldableCompose sync.Once
func Get_foldableCompose() gopurs_runtime.Value {
	once_foldableCompose.Do(func() {
		foldableCompose = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2)
_ = __local_var_5_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_5_0, a_7, b_6)
}), i_3, v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2), i_3, v_4)
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2)
_ = foldMap4_3_1
foldMap5_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2)
_ = foldMap5_4_2
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap4_3_1, gopurs_runtime.Apply(foldMap5_4_2, f_5), v_6)
})
}))
})
	})
	return foldableCompose
}

var foldableCoproduct gopurs_runtime.Value
var once_foldableCoproduct sync.Once
func Get_foldableCoproduct() gopurs_runtime.Value {
	once_foldableCoproduct.Do(func() {
		foldableCoproduct = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, z_3)
_ = __local_var_4_0
__local_var_5_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3)
_ = __local_var_5_1
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_6, "_tag").StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.RecordGet(v2_6, "value0"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_6, "_tag").StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.RecordGet(v2_6, "value0"))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3)
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, z_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_6, "_tag").StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.RecordGet(v2_6, "value0"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_6, "_tag").StrVal == "Right")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.RecordGet(v2_6, "value0"))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2)
_ = foldMap4_3_6
foldMap5_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2)
_ = foldMap5_4_7
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_8 := gopurs_runtime.Apply(foldMap4_3_6, f_5)
_ = __local_var_6_8
__local_var_7_9 := gopurs_runtime.Apply(foldMap5_4_7, f_5)
_ = __local_var_7_9
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_8, "_tag").StrVal == "Left")).IntVal != 0 {
__t10 = gopurs_runtime.Apply(__local_var_6_8, gopurs_runtime.RecordGet(v2_8, "value0"))
goto end_branch_10
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_8, "_tag").StrVal == "Right")).IntVal != 0 {
__t10 = gopurs_runtime.Apply(__local_var_7_9, gopurs_runtime.RecordGet(v2_8, "value0"))
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
})
}))
})
	})
	return foldableCoproduct
}

var foldableFirst gopurs_runtime.Value
var once_foldableFirst sync.Once
func Get_foldableFirst() gopurs_runtime.Value {
	once_foldableFirst.Do(func() {
		foldableFirst = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = z_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(v_2, "value0"), z_1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = z_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.RecordGet(v_2, "value0"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_2
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t3 = mempty_1_2
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(v_2, gopurs_runtime.RecordGet(v1_3, "value0"))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}))
	})
	return foldableFirst
}

var foldableLast gopurs_runtime.Value
var once_foldableLast sync.Once
func Get_foldableLast() gopurs_runtime.Value {
	once_foldableLast.Do(func() {
		foldableLast = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = z_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(v_2, "value0"), z_1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = z_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.RecordGet(v_2, "value0"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_2
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t3 = mempty_1_2
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(v_2, gopurs_runtime.RecordGet(v1_3, "value0"))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}))
	})
	return foldableLast
}

var foldableProduct gopurs_runtime.Value
var once_foldableProduct sync.Once
func Get_foldableProduct() gopurs_runtime.Value {
	once_foldableProduct.Do(func() {
		foldableProduct = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3, gopurs_runtime.RecordGet(v_4, "value1")), gopurs_runtime.RecordGet(v_4, "value0"))
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3, gopurs_runtime.RecordGet(v_4, "value0")), gopurs_runtime.RecordGet(v_4, "value1"))
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2)
_ = foldMap4_3_0
foldMap5_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2)
_ = foldMap5_4_1
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(foldMap4_3_0, f_5, gopurs_runtime.RecordGet(v_6, "value0")), gopurs_runtime.Apply2(foldMap5_4_1, f_5, gopurs_runtime.RecordGet(v_6, "value1")))
})
}))
})
	})
	return foldableProduct
}

var foldlDefault gopurs_runtime.Value
var once_foldlDefault sync.Once
func Get_foldlDefault() gopurs_runtime.Value {
	once_foldlDefault.Do(func() {
		foldlDefault = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidFreeMonoidTree())
_ = foldMap2_1_0
return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldl"), c_2, u_3, gopurs_runtime.Apply2(foldMap2_1_0, Get_Node(), xs_4))
})
})
	})
	return foldlDefault
}

var foldrDefault gopurs_runtime.Value
var once_foldrDefault sync.Once
func Get_foldrDefault() gopurs_runtime.Value {
	once_foldrDefault.Do(func() {
		foldrDefault = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidFreeMonoidTree())
_ = foldMap2_1_0
return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), c_2, u_3, gopurs_runtime.Apply2(foldMap2_1_0, Get_Node(), xs_4))
})
})
	})
	return foldrDefault
}

var lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		lookup = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), pkg_Data_Maybe_First.Get_monoidFirst())
_ = foldMap2_1_0
return gopurs_runtime.Func2(func(dictEq_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap2_1_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_2, "eq"), a_3, gopurs_runtime.RecordGet(v_4, "value0"))).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(v_4, "value1"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_1:
return __t1
}))
})
})
	})
	return lookup
}

var surroundMap gopurs_runtime.Value
var once_surroundMap sync.Once
func Get_surroundMap() gopurs_runtime.Value {
	once_surroundMap.Do(func() {
		surroundMap = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), Get_monoidEndo())
_ = foldMap2_1_0
return gopurs_runtime.Func4(func(dictSemigroup_2 gopurs_runtime.Value, d_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMap2_1_0, gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), d_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), gopurs_runtime.Apply(t_4, a_6), m_7))
}), f_5, d_3)
})
})
	})
	return surroundMap
}

var surround gopurs_runtime.Value
var once_surround sync.Once
func Get_surround() gopurs_runtime.Value {
	once_surround.Do(func() {
		surround = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
surroundMap1_1_0 := gopurs_runtime.Apply(Get_surroundMap(), dictFoldable_0)
_ = surroundMap1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
surroundMap2_3_1 := gopurs_runtime.Apply(surroundMap1_1_0, dictSemigroup_2)
_ = surroundMap2_3_1
return gopurs_runtime.Func(func(d_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(surroundMap2_3_1, d_4, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
})
})
	})
	return surround
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func4(func(dictFoldable_0 gopurs_runtime.Value, dictMonad_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value, b0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}), "bind"), b_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_6, a_5)
}))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), b0_3))
})
	})
	return foldM
}

var fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		fold = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return fold
}

var findMap gopurs_runtime.Value
var once_findMap sync.Once
func Get_findMap() gopurs_runtime.Value {
	once_findMap.Do(func() {
		findMap = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(p_1, v1_3)
goto end_branch_0
} else {

}
}
{
__t0 = v_2
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
})
	})
	return findMap
}

var find gopurs_runtime.Value
var once_find sync.Once
func Get_find() gopurs_runtime.Value {
	once_find.Do(func() {
		find = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nothing").IntVal != 0 && gopurs_runtime.Apply(p_1, v1_3).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), v1_3)
goto end_branch_0
} else {

}
}
{
__t0 = v_2
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
})
	})
	return find
}

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "disj"), v_2, v1_3)
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "ff"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
})))
})
	})
	return any
}

var elem gopurs_runtime.Value
var once_elem sync.Once
func Get_elem() gopurs_runtime.Value {
	once_elem.Do(func() {
		elem = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(v_1.IntVal != 0 || v1_2.IntVal != 0)
}))
_ = semigroupDisj1_1_1
any1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Bool(false), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_1
})))
_ = any1_1_0
return gopurs_runtime.Func2(func(dictEq_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq_2, "eq"), x_3))
})
})
	})
	return elem
}

var notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		notElem = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(v_1.IntVal != 0 || v1_2.IntVal != 0)
}))
_ = semigroupDisj1_1_1
any1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Bool(false), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_1
})))
_ = any1_1_0
return gopurs_runtime.Func2(func(dictEq_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq_2, "eq"), x_3))
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(__local_var_4_2, x_5).IntVal == 0)
})
})
})
	})
	return notElem
}

var or gopurs_runtime.Value
var once_or sync.Once
func Get_or() gopurs_runtime.Value {
	once_or.Do(func() {
		or = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "disj"), v_2, v1_3)
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "ff"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
})), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return or
}

var all gopurs_runtime.Value
var once_all sync.Once
func Get_all() gopurs_runtime.Value {
	once_all.Do(func() {
		all = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "conj"), v_2, v1_3)
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "tt"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
})))
})
	})
	return all
}

var and gopurs_runtime.Value
var once_and sync.Once
func Get_and() gopurs_runtime.Value {
	once_and.Do(func() {
		and = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "conj"), v_2, v1_3)
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "tt"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
})), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return and
}

func Get_foldlArray() gopurs_runtime.Value {
	return _Gopurs_FoldlArray
}

func Get_foldrArray() gopurs_runtime.Value {
	return _Gopurs_FoldrArray
}
