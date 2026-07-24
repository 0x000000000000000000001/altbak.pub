package Data_List_Types

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Control_Category "gopurs/output/Control.Category"
)

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
		NonEmptyList = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return NonEmptyList
}

var toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		toList = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
}()
})
	})
	return toList
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

var nelCons gopurs_runtime.Value
var once_nelCons sync.Once
func Get_nelCons() gopurs_runtime.Value {
	once_nelCons.Do(func() {
		nelCons = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nelCons(a_0_box, v_1_box)
})
	})
	return nelCons
}

var listMap gopurs_runtime.Value
var once_listMap sync.Once
func Get_listMap() gopurs_runtime.Value {
	once_listMap.Do(func() {
		listMap = gopurs_runtime.Func(func(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0 gopurs_runtime.Value
chunkedRevMap_1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
chunkedRevMap_1_0:
for {
if false { continue chunkedRevMap_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1].UnsafePtr)[1].StrVal == "Cons").IntVal != 0 {
v_2_loop = gopurs_runtime.Constructor2("Cons", v1_3, v_2)
v1_3_loop = (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1].UnsafePtr)[1].UnsafePtr)[1]
continue chunkedRevMap_1_0
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
var reverseUnrolledMap_4_1 gopurs_runtime.Value
reverseUnrolledMap_4_1 = gopurs_runtime.Func(func(v2_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
reverseUnrolledMap_4_1:
for {
if false { continue reverseUnrolledMap_4_1 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 gopurs_runtime.Value = v3_6_loop
_ = v3_6
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_5.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0].UnsafePtr)[1].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0].UnsafePtr)[1].UnsafePtr)[1].StrVal == "Cons").IntVal != 0 {
v2_5_loop = (*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[1]
v3_6_loop = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0].UnsafePtr)[0]), gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0].UnsafePtr)[1].UnsafePtr)[0]), gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0].UnsafePtr)[1].UnsafePtr)[1].UnsafePtr)[0]), v3_6)))
continue reverseUnrolledMap_4_1
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = v3_6
}
end_branch_2:
return __t2
}
}()
})
})
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Cons").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1].StrVal == "Cons").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1].UnsafePtr)[1].StrVal == "Nil").IntVal != 0 {
__t5 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]), gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1].UnsafePtr)[0]), gopurs_runtime.Constructor0("Nil")))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("Nil")
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1].StrVal == "Nil").IntVal != 0 {
__t4 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]), gopurs_runtime.Constructor0("Nil"))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("Nil")
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nil")
}
end_branch_3:
__t6 = gopurs_runtime.Apply2(reverseUnrolledMap_4_1, v_2, __t3)
}
end_branch_6:
return __t6
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0, gopurs_runtime.Constructor0("Nil"))
}()
})
	})
	return listMap
}

var functorList gopurs_runtime.Value
var once_functorList sync.Once
func Get_functorList() gopurs_runtime.Value {
	once_functorList.Do(func() {
		functorList = gopurs_runtime.RecordDict1("map", Get_listMap())
	})
	return functorList
}

var functorNonEmptyList gopurs_runtime.Value
var once_functorNonEmptyList sync.Once
func Get_functorNonEmptyList() gopurs_runtime.Value {
	once_functorNonEmptyList.Do(func() {
		functorNonEmptyList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(m_1.UnsafePtr)[0]), gopurs_runtime.Apply2(Get_listMap(), f_0, (*[1024]gopurs_runtime.Value)(m_1.UnsafePtr)[1]))
}))
	})
	return functorNonEmptyList
}

var foldableList gopurs_runtime.Value
var once_foldableList sync.Once
func Get_foldableList() gopurs_runtime.Value {
	once_foldableList.Do(func() {
		foldableList = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
}), b_1)
_ = __local_var_2_0
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_1:
for {
if false { continue go__3_1 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 {
__t2 = v_4
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 {
v_4_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0], v_4)
v1_5_loop = (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]
continue go__3_1
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
__local_var_4_3 := gopurs_runtime.Apply(go__3_1, gopurs_runtime.Constructor0("Nil"))
_ = __local_var_4_3
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_4_3, x_5))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_4 gopurs_runtime.Value
go__1_4 = gopurs_runtime.Func(func(b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_4:
for {
if false { continue go__1_4 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Nil").IntVal != 0 {
__t5 = b_2
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Cons").IntVal != 0 {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0])
v_3_loop = (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]
continue go__1_4
__t5 = gopurs_runtime.Value{}
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
return go__1_4
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_6 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_6
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_3)
_ = __local_var_4_7
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_7, gopurs_runtime.Apply(f_2, x_5))
})
}), mempty_1_6)
})
}))
	})
	return foldableList
}

var foldableNonEmptyList gopurs_runtime.Value
var once_foldableNonEmptyList sync.Once
func Get_foldableNonEmptyList() gopurs_runtime.Value {
	once_foldableNonEmptyList.Do(func() {
		foldableNonEmptyList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableList(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), gopurs_runtime.Apply2(foldMap1_1_0, f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(b_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_1:
for {
if false { continue go__3_1 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5.StrVal == "Nil").IntVal != 0 {
__t2 = b_4
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Cons").IntVal != 0 {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0])
v_5_loop = (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]
continue go__3_1
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
return gopurs_runtime.Apply2(go__3_1, gopurs_runtime.Apply2(f_0, b_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1])
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), f_0, b_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
}))
	})
	return foldableNonEmptyList
}

var foldableWithIndexList gopurs_runtime.Value
var once_foldableWithIndexList sync.Once
func Get_foldableWithIndexList() gopurs_runtime.Value {
	once_foldableWithIndexList.Do(func() {
		foldableWithIndexList = gopurs_runtime.RecordDict4("foldrWithIndex", "foldlWithIndex", "foldMapWithIndex", "Foldable0", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(b_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5.StrVal == "Nil").IntVal != 0 {
__t1 = b_4
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Cons").IntVal != 0 {
b_4_loop = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(b_4.UnsafePtr)[0].IntVal + 1), gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(b_4.UnsafePtr)[1]))
v_5_loop = (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]
continue go__3_0
__t1 = gopurs_runtime.Value{}
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
v_4_2 := gopurs_runtime.Apply2(go__3_0, gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int(0), gopurs_runtime.Constructor0("Nil")), xs_2)
_ = v_4_2
var go__5_3 gopurs_runtime.Value
go__5_3 = gopurs_runtime.Func(func(b_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_3:
for {
if false { continue go__5_3 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_7.StrVal == "Nil").IntVal != 0 {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_7.StrVal == "Cons").IntVal != 0 {
b_6_loop = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(b_6.UnsafePtr)[0].IntVal - 1), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(b_6.UnsafePtr)[0].IntVal - 1), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(b_6.UnsafePtr)[1]))
v_7_loop = (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1]
continue go__5_3
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply2(go__5_3, gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v_4_2.UnsafePtr)[0], b_1), (*[1024]gopurs_runtime.Value)(v_4_2.UnsafePtr)[1]).UnsafePtr)[1]
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_5 gopurs_runtime.Value
go__2_5 = gopurs_runtime.Func(func(b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_5:
for {
if false { continue go__2_5 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4.StrVal == "Nil").IntVal != 0 {
__t6 = b_3
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 {
b_3_loop = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(b_3.UnsafePtr)[0].IntVal + 1), gopurs_runtime.Apply3(f_0, (*[1024]gopurs_runtime.Value)(b_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(b_3.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]))
v_4_loop = (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]
continue go__2_5
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}()
})
})
__local_var_3_7 := gopurs_runtime.Apply(go__2_5, gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int(0), acc_1))
_ = __local_var_3_7
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(__local_var_3_7, x_4).UnsafePtr)[1]
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
if gopurs_runtime.Bool(x_4.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(x_4.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0].IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, __t1)
}), v_3)
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableWithIndexNonEmpty(), "foldlWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_3.StrVal == "Nothing").IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(x_3.StrVal == "Just").IntVal != 0 {
__t2 = gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[0].IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, __t2)
}), b_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableWithIndexNonEmpty(), "foldrWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_3.StrVal == "Nothing").IntVal != 0 {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(x_3.StrVal == "Just").IntVal != 0 {
__t3 = gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[0].IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Apply(f_0, __t3)
}), b_1, v_2)
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
return gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply2(f_0, i_1, x_2), acc_3)
}), gopurs_runtime.Constructor0("Nil"))
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
		mapWithIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex(f_0_box, v_1_box)
})
	})
	return mapWithIndex
}

var functorWithIndexNonEmptyList gopurs_runtime.Value
var once_functorWithIndexNonEmptyList sync.Once
func Get_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyList.Do(func() {
		functorWithIndexNonEmptyList = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(fn_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_2.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_2.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0].IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(fn_0, __t0)
}), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}))
	})
	return functorWithIndexNonEmptyList
}

var semigroupList gopurs_runtime.Value
var once_semigroupList sync.Once
func Get_semigroupList() gopurs_runtime.Value {
	once_semigroupList.Do(func() {
		semigroupList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), ys_1, xs_0)
}))
	})
	return semigroupList
}

var monoidList gopurs_runtime.Value
var once_monoidList sync.Once
func Get_monoidList() gopurs_runtime.Value {
	once_monoidList.Do(func() {
		monoidList = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Constructor0("Nil"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupList()
}))
	})
	return monoidList
}

var semigroupNonEmptyList gopurs_runtime.Value
var once_semigroupNonEmptyList sync.Once
func Get_semigroupNonEmptyList() gopurs_runtime.Value {
	once_semigroupNonEmptyList.Do(func() {
		semigroupNonEmptyList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(as_prime_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(as_prime_1.UnsafePtr)[1]), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]))
}))
	})
	return semigroupNonEmptyList
}

var showList gopurs_runtime.Value
var once_showList sync.Once
func Get_showList() gopurs_runtime.Value {
	once_showList.Do(func() {
		showList = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
show_1_0 := gopurs_runtime.RecordGet(dictShow_0, "show")
_ = show_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Nil").IntVal != 0 {
__t4 = gopurs_runtime.Str("Nil")
goto end_branch_4
} else {

}
}
{
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(b_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_1:
for {
if false { continue go__3_1 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5.StrVal == "Nil").IntVal != 0 {
__t2 = b_4
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Cons").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet(b_4, "init").IntVal != 0 {
__t3 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), gopurs_runtime.Str(gopurs_runtime.RecordGet(b_4, "acc").StrVal + " : " + (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0].StrVal))
}
end_branch_3:
b_4_loop = __t3
v_5_loop = (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]
continue go__3_1
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
__t4 = "(" + gopurs_runtime.RecordGet(gopurs_runtime.Apply2(go__3_1, gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(true), gopurs_runtime.Str("")), gopurs_runtime.Apply2(Get_listMap(), show_1_0, v_2)), "acc").StrVal + " : Nil)"
}
end_branch_4:
return __t4
}))
}()
})
	})
	return showList
}

var showNonEmptyList gopurs_runtime.Value
var once_showNonEmptyList sync.Once
func Get_showNonEmptyList() gopurs_runtime.Value {
	once_showNonEmptyList.Do(func() {
		showNonEmptyList = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
__local_var_1_0 := gopurs_runtime.Apply(Get_showList(), dictShow_0)
_ = __local_var_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(NonEmptyList (NonEmpty " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]).StrVal + "))")
}))
}()
})
	})
	return showNonEmptyList
}

var traversableList gopurs_runtime.Value
var once_traversableList sync.Once
func Get_traversableList() gopurs_runtime.Value {
	once_traversableList.Do(func() {
		traversableList = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(b_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_2:
for {
if false { continue go__3_2 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5.StrVal == "Nil").IntVal != 0 {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Cons").IntVal != 0 {
b_4_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], b_4)
v_5_loop = (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]
continue go__3_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__3_2, gopurs_runtime.Constructor0("Nil")))
_ = __local_var_3_1
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(b_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_4:
for {
if false { continue go__4_4 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6.StrVal == "Nil").IntVal != 0 {
__t5 = b_5
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Cons").IntVal != 0 {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", a_8, b_7)
}), b_5), gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]))
v_6_loop = (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1]
continue go__4_4
__t5 = gopurs_runtime.Value{}
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
__local_var_5_6 := gopurs_runtime.Apply(go__4_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor0("Nil")))
_ = __local_var_5_6
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(__local_var_5_6, x_6))
})
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

var traversableNonEmptyList gopurs_runtime.Value
var once_traversableNonEmptyList sync.Once
func Get_traversableNonEmptyList() gopurs_runtime.Value {
	once_traversableNonEmptyList.Do(func() {
		traversableNonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableNonEmpty(), Get_traversableList())
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
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(b_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_2:
for {
if false { continue go__3_2 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5.StrVal == "Nil").IntVal != 0 {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Cons").IntVal != 0 {
b_4_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], b_4)
v_5_loop = (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]
continue go__3_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__3_2, gopurs_runtime.Constructor0("Nil")))
_ = __local_var_3_1
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(b_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_4:
for {
if false { continue go__4_4 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6.StrVal == "Nil").IntVal != 0 {
__t5 = b_5
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Cons").IntVal != 0 {
b_5_loop = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(b_5.UnsafePtr)[0].IntVal + 1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", a_8, b_7)
}), (*[1024]gopurs_runtime.Value)(b_5.UnsafePtr)[1]), gopurs_runtime.Apply2(f_2, (*[1024]gopurs_runtime.Value)(b_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0])))
v_6_loop = (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1]
continue go__4_4
__t5 = gopurs_runtime.Value{}
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
__local_var_5_6 := gopurs_runtime.Apply(go__4_4, gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Int(0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor0("Nil"))))
_ = __local_var_5_6
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(__local_var_5_6, x_6).UnsafePtr)[1])
})
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_NonEmptyList(), gopurs_runtime.Apply2(traverseWithIndex1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_4.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(x_4.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0].IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, __t1)
}), v_3))
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
		unfoldable1List = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(source_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 gopurs_runtime.Value = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.Apply(f_0, source_3)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[1].StrVal == "Just").IntVal != 0 {
source_3_loop = (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[1].UnsafePtr)[0]
memo_4_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[0], memo_4)
continue go__2_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[1].StrVal == "Nothing").IntVal != 0 {
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(b_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_3:
for {
if false { continue go__6_3 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_8.StrVal == "Nil").IntVal != 0 {
__t4 = b_7
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_8.StrVal == "Cons").IntVal != 0 {
b_7_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0], b_7)
v_8_loop = (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1]
continue go__6_3
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
__t2 = gopurs_runtime.Apply2(go__6_3, gopurs_runtime.Constructor0("Nil"), gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[0], memo_4))
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
return gopurs_runtime.Apply2(go__2_0, b_1, gopurs_runtime.Constructor0("Nil"))
}))
	})
	return unfoldable1List
}

var unfoldableList gopurs_runtime.Value
var once_unfoldableList sync.Once
func Get_unfoldableList() gopurs_runtime.Value {
	once_unfoldableList.Do(func() {
		unfoldableList = gopurs_runtime.RecordDict2("unfoldr", "Unfoldable10", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(source_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 gopurs_runtime.Value = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.Apply(f_0, source_3)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5_1.StrVal == "Nothing").IntVal != 0 {
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(b_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_3:
for {
if false { continue go__6_3 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_8.StrVal == "Nil").IntVal != 0 {
__t4 = b_7
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_8.StrVal == "Cons").IntVal != 0 {
b_7_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0], b_7)
v_8_loop = (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1]
continue go__6_3
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
__t2 = gopurs_runtime.Apply2(go__6_3, gopurs_runtime.Constructor0("Nil"), memo_4)
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_5_1.StrVal == "Just").IntVal != 0 {
source_3_loop = (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[0].UnsafePtr)[1]
memo_4_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[0].UnsafePtr)[0], memo_4)
continue go__2_0
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
return gopurs_runtime.Apply2(go__2_0, b_1, gopurs_runtime.Constructor0("Nil"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_unfoldable1List()
}))
	})
	return unfoldableList
}

var unfoldable1NonEmptyList gopurs_runtime.Value
var once_unfoldable1NonEmptyList sync.Once
func Get_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_unfoldable1NonEmptyList.Do(func() {
		unfoldable1NonEmptyList = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, b_1)
_ = __local_var_2_0
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(source_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_1:
for {
if false { continue go__3_1 }
var source_4 gopurs_runtime.Value = source_4_loop
_ = source_4
var memo_5 gopurs_runtime.Value = memo_5_loop
_ = memo_5
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(source_4.StrVal == "Just").IntVal != 0 {
source_4_loop = (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(source_4.UnsafePtr)[0]).UnsafePtr)[1]
memo_5_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(source_4.UnsafePtr)[0]).UnsafePtr)[0], memo_5)
continue go__3_1
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
var go__6_2 gopurs_runtime.Value
go__6_2 = gopurs_runtime.Func(func(b_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_2:
for {
if false { continue go__6_2 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_8.StrVal == "Nil").IntVal != 0 {
__t3 = b_7
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_8.StrVal == "Cons").IntVal != 0 {
b_7_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0], b_7)
v_8_loop = (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1]
continue go__6_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__t4 = gopurs_runtime.Apply2(go__6_2, gopurs_runtime.Constructor0("Nil"), memo_5)
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0], gopurs_runtime.Apply2(go__3_1, (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[1], gopurs_runtime.Constructor0("Nil")))
}))
	})
	return unfoldable1NonEmptyList
}

var foldable1NonEmptyList gopurs_runtime.Value
var once_foldable1NonEmptyList sync.Once
func Get_foldable1NonEmptyList() gopurs_runtime.Value {
	once_foldable1NonEmptyList.Do(func() {
		foldable1NonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), Get_foldableList())
	})
	return foldable1NonEmptyList
}

var extendNonEmptyList gopurs_runtime.Value
var once_extendNonEmptyList sync.Once
func Get_extendNonEmptyList() gopurs_runtime.Value {
	once_extendNonEmptyList.Do(func() {
		extendNonEmptyList = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("val", "acc", gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(f_0, gopurs_runtime.Constructor2("NonEmpty", a_2, gopurs_runtime.RecordGet(v1_3, "acc"))), gopurs_runtime.RecordGet(v1_3, "val")), gopurs_runtime.Constructor2("Cons", a_2, gopurs_runtime.RecordGet(v1_3, "acc")))
}), gopurs_runtime.RecordDict2("val", "acc", gopurs_runtime.Constructor0("Nil"), gopurs_runtime.Constructor0("Nil")), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]), "val"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}))
	})
	return extendNonEmptyList
}

var extendList gopurs_runtime.Value
var once_extendList sync.Once
func Get_extendList() gopurs_runtime.Value {
	once_extendList.Do(func() {
		extendList = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1.StrVal == "Nil").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nil")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v1_1.StrVal == "Cons").IntVal != 0 {
__t0 = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(a_prime_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("val", "acc", gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply(v_0, gopurs_runtime.Constructor2("Cons", a_prime_2, gopurs_runtime.RecordGet(v2_3, "acc"))), gopurs_runtime.RecordGet(v2_3, "val")), gopurs_runtime.Constructor2("Cons", a_prime_2, gopurs_runtime.RecordGet(v2_3, "acc")))
}), gopurs_runtime.RecordDict2("val", "acc", gopurs_runtime.Constructor0("Nil"), gopurs_runtime.Constructor0("Nil")), (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[1]), "val"))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}))
	})
	return extendList
}

var eq1List gopurs_runtime.Value
var once_eq1List sync.Once
func Get_eq1List() gopurs_runtime.Value {
	once_eq1List.Do(func() {
		eq1List = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if v2_6.IntVal != 0 != true {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 && v2_6.IntVal != 0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply3(go__3_0, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1], gopurs_runtime.Bool(v2_6.IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]).IntVal != 0)).IntVal != 0)
}
end_branch_1:
return __t1
})
return gopurs_runtime.Apply3(go__3_0, xs_1, ys_2, gopurs_runtime.Bool(true))
}))
	})
	return eq1List
}

var eq1NonEmptyList gopurs_runtime.Value
var once_eq1NonEmptyList sync.Once
func Get_eq1NonEmptyList() gopurs_runtime.Value {
	once_eq1NonEmptyList.Do(func() {
		eq1NonEmptyList = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if v2_6.IntVal != 0 != true {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 && v2_6.IntVal != 0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply3(go__3_0, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1], gopurs_runtime.Bool(v2_6.IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]).IntVal != 0)).IntVal != 0)
}
end_branch_1:
return __t1
})
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply3(go__3_0, (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[1], gopurs_runtime.Bool(true)).IntVal != 0)
}))
	})
	return eq1NonEmptyList
}

var eqList gopurs_runtime.Value
var once_eqList sync.Once
func Get_eqList() gopurs_runtime.Value {
	once_eqList.Do(func() {
		eqList = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if v2_6.IntVal != 0 != true {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 && v2_6.IntVal != 0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply3(go__3_0, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1], gopurs_runtime.Bool(v2_6.IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]).IntVal != 0)).IntVal != 0)
}
end_branch_1:
return __t1
})
return gopurs_runtime.Apply3(go__3_0, xs_1, ys_2, gopurs_runtime.Bool(true))
}))
}()
})
	})
	return eqList
}

var eqNonEmptyList gopurs_runtime.Value
var once_eqNonEmptyList sync.Once
func Get_eqNonEmptyList() gopurs_runtime.Value {
	once_eqNonEmptyList.Do(func() {
		eqNonEmptyList = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if v2_6.IntVal != 0 != true {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 && v2_6.IntVal != 0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply3(go__3_0, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1], gopurs_runtime.Bool(v2_6.IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]).IntVal != 0)).IntVal != 0)
}
end_branch_1:
return __t1
})
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply3(go__3_0, (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[1], gopurs_runtime.Bool(true)).IntVal != 0)
}))
}()
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
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4.StrVal == "Nil").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 {
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
if gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("GT")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 {
v2_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0])
_ = v2_6_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_6_3.StrVal == "EQ").IntVal != 0 {
v_4_loop = (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]
v1_5_loop = (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]
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
return gopurs_runtime.Apply2(go__3_0, xs_1, ys_2)
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
		ord1NonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_ord1NonEmpty(), Get_ord1List())
	})
	return ord1NonEmptyList
}

var ordList gopurs_runtime.Value
var once_ordList sync.Once
func Get_ordList() gopurs_runtime.Value {
	once_ordList.Do(func() {
		ordList = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqList1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_2 gopurs_runtime.Value
_ = go__4_2
go__4_2 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if v2_7.IntVal != 0 != true {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Nil").IntVal != 0 {
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.StrVal == "Nil").IntVal != 0 && v2_7.IntVal != 0)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_5.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_6.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply3(go__4_2, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1], gopurs_runtime.Bool(v2_7.IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]).IntVal != 0)).IntVal != 0)
}
end_branch_3:
return __t3
})
return gopurs_runtime.Apply3(go__4_2, xs_2, ys_3, gopurs_runtime.Bool(true))
}))
_ = eqList1_2_1
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_4 gopurs_runtime.Value
go__5_4 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_4:
for {
if false { continue go__5_4 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6.StrVal == "Nil").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Nil").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("EQ")
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Constructor0("LT")
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v1_7.StrVal == "Nil").IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("GT")
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_7.StrVal == "Cons").IntVal != 0 {
v2_8_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0])
_ = v2_8_7
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_8_7.StrVal == "EQ").IntVal != 0 {
v_6_loop = (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1]
v1_7_loop = (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[1]
continue go__5_4
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
__t8 = v2_8_7
}
end_branch_8:
__t5 = __t8
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
return gopurs_runtime.Apply2(go__5_4, xs_3, ys_4)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqList1_2_1
}))
}()
})
	})
	return ordList
}

var ordNonEmptyList gopurs_runtime.Value
var once_ordNonEmptyList sync.Once
func Get_ordNonEmptyList() gopurs_runtime.Value {
	once_ordNonEmptyList.Do(func() {
		ordNonEmptyList = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd_0)
}()
})
	})
	return ordNonEmptyList
}

var comonadNonEmptyList gopurs_runtime.Value
var once_comonadNonEmptyList sync.Once
func Get_comonadNonEmptyList() gopurs_runtime.Value {
	once_comonadNonEmptyList.Do(func() {
		comonadNonEmptyList = gopurs_runtime.RecordDict2("extract", "Extend0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendNonEmptyList()
}))
	})
	return comonadNonEmptyList
}

var applyList gopurs_runtime.Value
var once_applyList sync.Once
func Get_applyList() gopurs_runtime.Value {
	once_applyList.Do(func() {
		applyList = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Nil").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nil")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Cons").IntVal != 0 {
__t0 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1], v1_1), gopurs_runtime.Apply2(Get_listMap(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], v1_1))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}))
	})
	return applyList
}

var applyNonEmptyList gopurs_runtime.Value
var once_applyNonEmptyList sync.Once
func Get_applyNonEmptyList() gopurs_runtime.Value {
	once_applyNonEmptyList.Do(func() {
		applyNonEmptyList = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0]), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]), (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[1]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1], gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0], gopurs_runtime.Constructor0("Nil")))))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}))
	})
	return applyNonEmptyList
}

var bindList gopurs_runtime.Value
var once_bindList sync.Once
func Get_bindList() gopurs_runtime.Value {
	once_bindList.Do(func() {
		bindList = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Nil").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nil")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Cons").IntVal != 0 {
__t0 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1], v1_1), gopurs_runtime.Apply(v1_1, (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}))
	})
	return bindList
}

var bindNonEmptyList gopurs_runtime.Value
var once_bindNonEmptyList sync.Once
func Get_bindNonEmptyList() gopurs_runtime.Value {
	once_bindNonEmptyList.Do(func() {
		bindNonEmptyList = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
_ = v1_2_0
return gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v1_2_0.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1], gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(__local_var_4_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_4_1.UnsafePtr)[1])
})), (*[1024]gopurs_runtime.Value)(v1_2_0.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}))
	})
	return bindNonEmptyList
}

var applicativeList gopurs_runtime.Value
var once_applicativeList sync.Once
func Get_applicativeList() gopurs_runtime.Value {
	once_applicativeList.Do(func() {
		applicativeList = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", a_0, gopurs_runtime.Constructor0("Nil"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}))
	})
	return applicativeList
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
		plusList = gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Constructor0("Nil"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
		applicativeNonEmptyList = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", x_0, gopurs_runtime.Constructor0("Nil"))
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

var traversable1NonEmptyList gopurs_runtime.Value
var once_traversable1NonEmptyList sync.Once
func Get_traversable1NonEmptyList() gopurs_runtime.Value {
	once_traversable1NonEmptyList.Do(func() {
		traversable1NonEmptyList = gopurs_runtime.RecordDict4("traverse1", "sequence1", "Foldable10", "Traversable1", gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(b_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_1:
for {
if false { continue go__4_1 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6.StrVal == "Nil").IntVal != 0 {
__t2 = b_5
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Cons").IntVal != 0 {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", a_8, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(b_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(b_7.UnsafePtr)[1]))
}), b_5), gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]))
v_6_loop = (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1]
continue go__4_1
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(b_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_3:
for {
if false { continue go__6_3 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_8.StrVal == "Nil").IntVal != 0 {
__t4 = b_7
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_8.StrVal == "Cons").IntVal != 0 {
b_7_loop = gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0], gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(b_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(b_7.UnsafePtr)[1]))
v_8_loop = (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1]
continue go__6_3
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Apply2(go__6_3, gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0], gopurs_runtime.Constructor0("Nil")), (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1])
}), gopurs_runtime.Apply2(go__4_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.RecordGet(Get_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0])), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversable1NonEmptyList(), "traverse1"), dictApply_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1NonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableNonEmptyList()
}))
	})
	return traversable1NonEmptyList
}

func Call_nelCons(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Constructor2("NonEmpty", a_0, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]))
}

func Call_mapWithIndex(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply2(f_0, gopurs_runtime.Constructor0("Nothing"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func3(func(i_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply2(f_0, gopurs_runtime.Constructor1("Just", i_2), x_3), acc_4)
}), gopurs_runtime.Constructor0("Nil"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]))
}


