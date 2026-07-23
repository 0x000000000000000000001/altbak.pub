package Data_List_Lazy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
)

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = func() gopurs_runtime.Value {
semigroupDisj1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(v_0.IntVal != 0 || v1_1.IntVal != 0)
}))
_ = semigroupDisj1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Bool(false), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_0_0
})))
}()
	})
	return any
}

var Pattern gopurs_runtime.Value
var once_Pattern sync.Once
func Get_Pattern() gopurs_runtime.Value {
	once_Pattern.Do(func() {
		Pattern = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Pattern
}

var zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		zipWith = gopurs_runtime.Func(func(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
zipWith:
for {
if false { continue zipWith }
var f_0 = f_0_loop
_ = f_0
var xs_1 = xs_1_loop
_ = xs_1
var ys_2 = ys_2_loop
_ = ys_2
__local_var_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_1, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons").IntVal != 0)).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(__local_var_4_1, "value0"), gopurs_runtime.RecordGet(v1_5, "value0")), gopurs_runtime.Apply3(Get_zipWith(), f_0, gopurs_runtime.RecordGet(__local_var_4_1, "value1"), gopurs_runtime.RecordGet(v1_5, "value1")))
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
}))
_ = __local_var_3_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Lazy.Get_force(), __local_var_3_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_2))
}))
}
}()
})
})
})
	})
	return zipWith
}

var zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		zipWithA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "traverse"), dictApplicative_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
_ = sequence1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply3(Get_zipWith(), f_2, xs_3, ys_4))
})
})
	})
	return zipWithA
}

var zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		zip = gopurs_runtime.Apply(Get_zipWith(), pkg_Data_Tuple.Get_Tuple())
	})
	return zip
}

var updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		updateAt = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
updateAt:
for {
if false { continue updateAt }
var n_0 = n_0_loop
_ = n_0
var x_1 = x_1_loop
_ = x_1
var xs_2 = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(n_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_1, gopurs_runtime.RecordGet(__local_var_4_0, "value1"))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(__local_var_4_0, "value0"), gopurs_runtime.Apply3(Get_updateAt(), gopurs_runtime.Int(n_0.IntVal - gopurs_runtime.Int(1).IntVal), x_1, gopurs_runtime.RecordGet(__local_var_4_0, "value1")))
}
end_branch_2:
__t1 = __t2
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
}
}()
})
})
})
	})
	return updateAt
}

var unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		unzip = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.RecordGet(v_0, "value0")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(v_0, "value1")
_ = __local_var_2_1
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.RecordGet(v1_3, "value0")
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.RecordGet(v1_3, "value1")
_ = __local_var_5_3
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_1_0, __local_var_4_2)
})), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_2_1, __local_var_5_3)
})))
})
}), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), pkg_Data_List_Lazy_Types.Get_nil(), pkg_Data_List_Lazy_Types.Get_nil()))
	})
	return unzip
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
v_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0)
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1_0, "_tag").StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.RecordGet(v_1_0, "value0"), gopurs_runtime.RecordGet(v_1_0, "value1")))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
	})
	return uncons
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(Get_uncons(), xs_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_2_0, "value0"), "head"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_2_0, "value0"), "tail")))
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
	return toUnfoldable
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var p_0 = p_0_loop
_ = p_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet(v_1, "value0")).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v_1, "value0"), gopurs_runtime.Apply2(Get_takeWhile(), p_0, gopurs_runtime.RecordGet(v_1, "value1")))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
}
end_branch_0:
return __t0
}))
}
}()
})
	})
	return takeWhile
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
take:
for {
if false { continue take }
var n_0 = n_0_loop
_ = n_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(n_0.IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_nil()
})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_1, "value0"), gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(n_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordGet(v1_1, "value1")))
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
}
end_branch_1:
return __t1
}
}()
})
	})
	return take
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_uncons(), xs_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_1_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_1_0, "value0"), "tail"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_1:
return __t1
})
	})
	return tail
}

var stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		stripPrefix = gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Value{PtrVal: func(prefix_3 gopurs_runtime.Value, input_4 gopurs_runtime.Value) gopurs_runtime.Value {
v1_5_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), prefix_3)
_ = v1_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), input_4))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5_1, "_tag").StrVal == "Cons")).IntVal != 0 {
v2_6_3 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), input_4)
_ = v2_6_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_6_3, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.RecordGet(v1_5_1, "value0"), gopurs_runtime.RecordGet(v2_6_3, "value0")).IntVal != 0)).IntVal != 0 {
__t4 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Loop"), gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.RecordGet(v1_5_1, "value1"), gopurs_runtime.RecordGet(v2_6_3, "value1"))))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_4:
__t2 = __t4
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}}
_ = __local_var_3_0
__local_var_4_5 := gopurs_runtime.Value{PtrVal: func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Just")).IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "_tag").StrVal == "Loop")).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Loop"), gopurs_runtime.UncurriedApp2(__local_var_3_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "value0"), "a"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "value0"), "b")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "_tag").StrVal == "Done")).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "value0")))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}}
_ = __local_var_4_5
var go__5_8 gopurs_runtime.Value
go__5_8 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_8:
for {
if false { continue go__5_8 }
var v_6 = v_6_loop
_ = v_6
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Loop")).IntVal != 0 {
v_6_loop = gopurs_runtime.UncurriedApp(__local_var_4_5, gopurs_runtime.RecordGet(v_6, "value0"))
continue go__5_8
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Done")).IntVal != 0 {
__t9 = gopurs_runtime.RecordGet(v_6, "value0")
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
return gopurs_runtime.Apply(go__5_8, gopurs_runtime.UncurriedApp(__local_var_4_5, gopurs_runtime.UncurriedApp2(__local_var_3_0, v_1, s_2)))
})
	})
	return stripPrefix
}

var span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		span = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
span:
for {
if false { continue span }
var p_0 = p_0_loop
_ = p_0
var xs_1 = xs_1_loop
_ = xs_1
v_2_0 := gopurs_runtime.Apply(Get_uncons(), xs_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2_0, "_tag").StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_2_0, "value0"), "head")).IntVal != 0)).IntVal != 0 {
__local_var_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_2_0, "value0"), "head")
_ = __local_var_3_2
v1_4_3 := gopurs_runtime.Apply2(Get_span(), p_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_2_0, "value0"), "tail"))
_ = v1_4_3
__local_var_5_4 := gopurs_runtime.RecordGet(v1_4_3, "init")
_ = __local_var_5_4
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_3_2, __local_var_5_4)
})), gopurs_runtime.RecordGet(v1_4_3, "rest"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", pkg_Data_List_Lazy_Types.Get_nil(), xs_1)
}
end_branch_1:
return __t1
}
}()
})
})
	})
	return span
}

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_1, pkg_Data_List_Lazy_Types.Get_nil())
})), xs_0)
})
	})
	return snoc
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), a_0, pkg_Data_List_Lazy_Types.Get_nil())
}))
})
	})
	return singleton
}

var showPattern gopurs_runtime.Value
var once_showPattern sync.Once
func Get_showPattern() gopurs_runtime.Value {
	once_showPattern.Do(func() {
		showPattern = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Pattern ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_showList(), dictShow_0), "show"), v_1).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
})
	})
	return showPattern
}

var scanlLazy gopurs_runtime.Value
var once_scanlLazy sync.Once
func Get_scanlLazy() gopurs_runtime.Value {
	once_scanlLazy.Do(func() {
		scanlLazy = gopurs_runtime.Func(func(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
scanlLazy:
for {
if false { continue scanlLazy }
var f_0 = f_0_loop
_ = f_0
var acc_1 = acc_1_loop
_ = acc_1
var xs_2 = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Cons")).IntVal != 0 {
acc_prime_5_2 := gopurs_runtime.Apply2(f_0, acc_1, gopurs_runtime.RecordGet(__local_var_4_0, "value0"))
_ = acc_prime_5_2
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), acc_prime_5_2, gopurs_runtime.Apply3(Get_scanlLazy(), f_0, acc_prime_5_2, gopurs_runtime.RecordGet(__local_var_4_0, "value1")))
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
}
}()
})
})
})
	})
	return scanlLazy
}

var reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		reverse = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), a_3, b_2)
}))
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_0))
}))
})
	})
	return reverse
}

var replicateM gopurs_runtime.Value
var once_replicateM sync.Once
func Get_replicateM() gopurs_runtime.Value {
	once_replicateM.Do(func() {
		replicateM = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
replicateM:
for {
if false { continue replicateM }
var dictMonad_0 = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(n_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(n_3.IntVal < gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply3(Get_replicateM(), dictMonad_0, gopurs_runtime.Int(n_3.IntVal - gopurs_runtime.Int(1).IntVal), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), a_5, as_6)
})))
}))
}))
}
end_branch_2:
return __t2
})
}
}()
})
	})
	return replicateM
}

var repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		repeat = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_0, go__1_0)
})))
}))
return go__1_0
})
	})
	return repeat
}

var replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		replicate = gopurs_runtime.Func2(func(i_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_take(), i_0, gopurs_runtime.Apply(Get_repeat(), xs_1))
})
	})
	return replicate
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func2(func(start_0 gopurs_runtime.Value, end_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(start_0.IntVal > end_1.IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_unfoldableList(), "unfoldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_2.IntVal >= end_1.IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), x_2, gopurs_runtime.Int(x_2.IntVal - gopurs_runtime.Int(1).IntVal)))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_2:
return __t2
}), start_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_unfoldableList(), "unfoldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_2.IntVal <= end_1.IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), x_2, gopurs_runtime.Int(x_2.IntVal + gopurs_runtime.Int(1).IntVal)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_0:
return __t0
}), start_0)
}
end_branch_1:
return __t1
})
	})
	return range_
}

var partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		partition = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, x_1)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("yes", "no", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_1, gopurs_runtime.RecordGet(v_2, "yes"))
})), gopurs_runtime.RecordGet(v_2, "no"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("yes", "no", gopurs_runtime.RecordGet(v_2, "yes"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_1, gopurs_runtime.RecordGet(v_2, "no"))
})))
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict2("yes", "no", pkg_Data_List_Lazy_Types.Get_nil(), pkg_Data_List_Lazy_Types.Get_nil()))
})
	})
	return partition
}

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_uncons(), x_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_1_0, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_1_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
	})
	return null
}

var nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		nubBy = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var goStep_1_0 gopurs_runtime.Value
_ = goStep_1_0
var go__1_1 gopurs_runtime.Value
_ = go__1_1
goStep_1_0 = gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons")).IntVal != 0 {
v2_4_3 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, gopurs_runtime.RecordGet(v1_3, "value0"), v_2)
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v2_4_3, "found")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(go__1_1, gopurs_runtime.RecordGet(v2_4_3, "result"), gopurs_runtime.RecordGet(v1_3, "value1")))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_3, "value0"), gopurs_runtime.Apply2(go__1_1, gopurs_runtime.RecordGet(v2_4_3, "result"), gopurs_runtime.RecordGet(v1_3, "value1")))
}
end_branch_4:
__t2 = __t4
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
go__1_1 = gopurs_runtime.Func2(func(s_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_5 := gopurs_runtime.Apply(goStep_1_0, s_2)
_ = __local_var_4_5
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_3))
}))
})
return gopurs_runtime.Apply(go__1_1, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf")))
})
	})
	return nubBy
}

var nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		nub = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_nubBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
})
	})
	return nub
}

var newtypePattern gopurs_runtime.Value
var once_newtypePattern sync.Once
func Get_newtypePattern() gopurs_runtime.Value {
	once_newtypePattern.Do(func() {
		newtypePattern = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypePattern
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var f_0 = f_0_loop
_ = f_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Cons")).IntVal != 0 {
v1_3_2 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_2, "value0"))
_ = v1_3_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(v_2, "value1"))
continue go__1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_3_2, "value0"), gopurs_runtime.Apply2(Get_mapMaybe(), f_0, gopurs_runtime.RecordGet(v_2, "value1")))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
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
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_2))
}))
})
}
}()
})
	})
	return mapMaybe
}

var some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		some = gopurs_runtime.Func3(func(dictAlternative_0 gopurs_runtime.Value, dictLazy_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Lazy_Types.Get_cons(), v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_1, "defer"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_many(), dictAlternative_0, dictLazy_1, v_2)
})))
})
	})
	return some
}

var many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		many = gopurs_runtime.Func3(func(dictAlternative_0 gopurs_runtime.Value, dictLazy_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply3(Get_some(), dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
})
	})
	return many
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(l_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(l_0.IntVal + gopurs_runtime.Int(1).IntVal)
}), gopurs_runtime.Int(0))
	})
	return length
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var v_1 = v_1_loop
_ = v_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
__local_var_2_3 := gopurs_runtime.Apply(Get_uncons(), gopurs_runtime.RecordGet(v_1, "value1"))
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_3, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_3, "_tag").StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Bool(false)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
if (__t4).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(v_1, "value0"))
goto end_branch_2
} else {

}
}
{
v_1_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(v_1, "value1"))
continue go__0_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return last
}

var iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		iterate = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, go__2_0)
_ = __local_var_4_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_1, __local_var_4_1)
})))
}))
return go__2_0
})
	})
	return iterate
}

var insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		insertAt = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
insertAt:
for {
if false { continue insertAt }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var v2_2 = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v1_1, v2_2)
}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v2_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v1_1, pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(__local_var_4_0, "value0"), gopurs_runtime.Apply3(Get_insertAt(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), v1_1, gopurs_runtime.RecordGet(__local_var_4_0, "value1")))
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
}
end_branch_2:
return __t2
}
}()
})
})
})
	})
	return insertAt
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
_ = go__0_0
go__0_0 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
__local_var_2_7 := gopurs_runtime.Apply(Get_uncons(), gopurs_runtime.RecordGet(v_1, "value1"))
_ = __local_var_2_7
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_7, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t8 = gopurs_runtime.Bool(true)
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_7, "_tag").StrVal == "Just")).IntVal != 0 {
__t8 = gopurs_runtime.Bool(false)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
if (__t8).IntVal != 0 {
__t6 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_6
} else {

}
}
{
__local_var_2_2 := gopurs_runtime.Apply(go__0_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(v_1, "value1")))
_ = __local_var_2_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_2, "_tag").StrVal == "Just")).IntVal != 0 {
__local_var_3_4 := gopurs_runtime.RecordGet(v_1, "value0")
_ = __local_var_3_4
__local_var_4_5 := gopurs_runtime.RecordGet(__local_var_2_2, "value0")
_ = __local_var_4_5
__t3 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_3_4, __local_var_4_5)
})))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_3:
__t6 = __t3
}
end_branch_6:
__t1 = __t6
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_1:
return __t1
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return init_
}

var index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		index = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(v_2, "value0"))
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(v_2, "value1"))
v1_3_loop = gopurs_runtime.Int(v1_3.IntVal - gopurs_runtime.Int(1).IntVal)
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
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
})
	})
	return index
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_uncons(), xs_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_1_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_1_0, "value0"), "head"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_1:
return __t1
})
	})
	return head
}

var transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		transpose = gopurs_runtime.Func(func(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
transpose:
for {
if false { continue transpose }
var xs_0 = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.Apply(Get_uncons(), xs_0)
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1_0, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = xs_0
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1_0, "_tag").StrVal == "Just")).IntVal != 0 {
v1_2_2 := gopurs_runtime.Apply(Get_uncons(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_1_0, "value0"), "head"))
_ = v1_2_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(Get_transpose(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_1_0, "value0"), "tail"))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2_2, "_tag").StrVal == "Just")).IntVal != 0 {
__local_var_3_4 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_2_2, "value0"), "head")
_ = __local_var_3_4
__local_var_4_5 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_2_2, "value0"), "tail")
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply2(Get_mapMaybe(), Get_head(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_1_0, "value0"), "tail"))
_ = __local_var_5_6
__local_var_6_7 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_3_4, __local_var_5_6)
}))
_ = __local_var_6_7
__local_var_7_9 := gopurs_runtime.Apply2(Get_mapMaybe(), Get_tail(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_1_0, "value0"), "tail"))
_ = __local_var_7_9
__local_var_7_8 := gopurs_runtime.Apply(Get_transpose(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_4_5, __local_var_7_9)
})))
_ = __local_var_7_8
__t3 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_6_7, __local_var_7_8)
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
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
	return transpose
}

var groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		groupBy = gopurs_runtime.Func(func(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
groupBy:
for {
if false { continue groupBy }
var eq_0 = eq_0_loop
_ = eq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Cons")).IntVal != 0 {
__local_var_2_1 := gopurs_runtime.RecordGet(v_1, "value0")
_ = __local_var_2_1
v1_3_2 := gopurs_runtime.Apply2(Get_span(), gopurs_runtime.Apply(eq_0, __local_var_2_1), gopurs_runtime.RecordGet(v_1, "value1"))
_ = v1_3_2
__local_var_4_3 := gopurs_runtime.RecordGet(v1_3_2, "init")
_ = __local_var_4_3
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("NonEmpty"), __local_var_2_1, __local_var_4_3)
})), gopurs_runtime.Apply2(Get_groupBy(), eq_0, gopurs_runtime.RecordGet(v1_3_2, "rest")))
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
}
}()
})
	})
	return groupBy
}

var group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		group = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return group
}

var insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		insertBy = gopurs_runtime.Func(func(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
insertBy:
for {
if false { continue insertBy }
var cmp_0 = cmp_0_loop
_ = cmp_0
var x_1 = x_1_loop
_ = x_1
var xs_2 = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_1, pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(cmp_0, x_1, gopurs_runtime.RecordGet(__local_var_4_0, "value0")), "_tag").StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(__local_var_4_0, "value0"), gopurs_runtime.Apply3(Get_insertBy(), cmp_0, x_1, gopurs_runtime.RecordGet(__local_var_4_0, "value1")))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_0
})))
}
end_branch_2:
__t1 = __t2
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
}
}()
})
})
})
	})
	return insertBy
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_insertBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
})
	})
	return insert
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
})
	})
	return fromFoldable
}

var foldrLazy gopurs_runtime.Value
var once_foldrLazy sync.Once
func Get_foldrLazy() gopurs_runtime.Value {
	once_foldrLazy.Do(func() {
		foldrLazy = gopurs_runtime.Func3(func(dictLazy_0 gopurs_runtime.Value, op_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_4)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5_1, "_tag").StrVal == "Cons")).IntVal != 0 {
__local_var_6_3 := gopurs_runtime.RecordGet(v_5_1, "value0")
_ = __local_var_6_3
__local_var_7_4 := gopurs_runtime.RecordGet(v_5_1, "value1")
_ = __local_var_7_4
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_0, "defer"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_1, __local_var_6_3, gopurs_runtime.Apply(go__3_0, __local_var_7_4))
}))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = z_2
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
return go__3_0
})
	})
	return foldrLazy
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 = dictMonad_0_loop
_ = dictMonad_0
var f_1 = f_1_loop
_ = f_1
var b_2 = b_2_loop
_ = b_2
var xs_3 = xs_3_loop
_ = xs_3
v_4_0 := gopurs_runtime.Apply(Get_uncons(), xs_3)
_ = v_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4_0, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4_0, "_tag").StrVal == "Just")).IntVal != 0 {
__local_var_5_2 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4_0, "value0"), "tail")
_ = __local_var_5_2
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4_0, "value0"), "head")), gopurs_runtime.Func(func(b_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_foldM(), dictMonad_0, f_1, b_prime_6, __local_var_5_2)
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
}
}()
})
})
})
})
	})
	return foldM
}

var findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		findIndex = gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(n_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var n_2 = n_2_loop
_ = n_2
var list_3 = list_3_loop
_ = list_3
__local_var_4_1 := gopurs_runtime.Apply(Get_uncons(), list_3)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_1, "_tag").StrVal == "Just")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_4_1, "value0"), "head"))).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), n_2)
goto end_branch_3
} else {

}
}
{
n_2_loop = gopurs_runtime.Int(n_2.IntVal + gopurs_runtime.Int(1).IntVal)
list_3_loop = gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_4_1, "value0"), "tail")
continue go__1_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_1, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Int(0))
})
	})
	return findIndex
}

var findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		findLastIndex = gopurs_runtime.Func2(func(fn_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(Get_findIndex(), fn_0, gopurs_runtime.Apply(Get_reverse(), xs_1))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Int(gopurs_runtime.Int(gopurs_runtime.Apply(Get_length(), xs_1).IntVal - gopurs_runtime.Int(1).IntVal).IntVal - gopurs_runtime.RecordGet(__local_var_2_0, "value0").IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_1:
return __t1
})
	})
	return findLastIndex
}

var filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		filterM = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
filterM:
for {
if false { continue filterM }
var dictMonad_0 = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(p_3 gopurs_runtime.Value, list_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_2 := gopurs_runtime.Apply(Get_uncons(), list_4)
_ = v_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil())
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5_2, "_tag").StrVal == "Just")).IntVal != 0 {
__local_var_6_4 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_5_2, "value0"), "head")
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_5_2, "value0"), "tail")
_ = __local_var_7_5
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply(p_3, __local_var_6_4), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply3(Get_filterM(), dictMonad_0, p_3, __local_var_7_5), gopurs_runtime.Func(func(xs_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (b_8).IntVal != 0 {
__t6 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_6_4, xs_prime_9)
}))
goto end_branch_6
} else {

}
}
{
__t6 = xs_prime_9
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), __t6)
}))
}))
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
}
}()
})
	})
	return filterM
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var p_0 = p_0_loop
_ = p_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet(v_2, "value0"))).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v_2, "value0"), gopurs_runtime.Apply2(Get_filter(), p_0, gopurs_runtime.RecordGet(v_2, "value1")))
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(v_2, "value1"))
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
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_2))
}))
})
}
}()
})
	})
	return filter
}

var intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		intersectBy = gopurs_runtime.Func3(func(eq_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_filter(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_any(), gopurs_runtime.Apply(eq_0, x_3), ys_2)
}), xs_1)
})
	})
	return intersectBy
}

var intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		intersect = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_intersectBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return intersect
}

var nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		nubByEq = gopurs_runtime.Func(func(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
nubByEq:
for {
if false { continue nubByEq }
var eq_0 = eq_0_loop
_ = eq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Cons")).IntVal != 0 {
__local_var_2_1 := gopurs_runtime.RecordGet(v_1, "value0")
_ = __local_var_2_1
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_2_1, gopurs_runtime.Apply2(Get_nubByEq(), eq_0, gopurs_runtime.Apply2(Get_filter(), gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(eq_0, __local_var_2_1, y_3).IntVal == 0)
}), gopurs_runtime.RecordGet(v_1, "value1"))))
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
}
}()
})
	})
	return nubByEq
}

var nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		nubEq = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_nubByEq(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return nubEq
}

var eqPattern gopurs_runtime.Value
var once_eqPattern sync.Once
func Get_eqPattern() gopurs_runtime.Value {
	once_eqPattern.Do(func() {
		eqPattern = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_eq1List(), "eq1"), dictEq_0, x_1, y_2)
}))
})
	})
	return eqPattern
}

var ordPattern gopurs_runtime.Value
var once_ordPattern sync.Once
func Get_ordPattern() gopurs_runtime.Value {
	once_ordPattern.Do(func() {
		ordPattern = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqPattern1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_eq1List(), "eq1"), __local_var_1_0, x_2, y_3)
}))
_ = eqPattern1_2_1
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_ordList(), dictOrd_0), "compare"), x_3, y_4)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqPattern1_2_1
}))
})
	})
	return ordPattern
}

var elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		elemLastIndex = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1)
}))
})
	})
	return elemLastIndex
}

var elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		elemIndex = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1)
}))
})
	})
	return elemIndex
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet(v_2, "value0")).IntVal != 0)).IntVal != 0 {
v_2_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(v_2, "value1"))
continue go__1_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return v_2
}))
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_2))
})
})
	})
	return dropWhile
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons")).IntVal != 0 {
v_2_loop = gopurs_runtime.Int(v_2.IntVal - gopurs_runtime.Int(1).IntVal)
v1_3_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(v1_3, "value1"))
continue go__1_0
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
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(go__1_0, n_0))
})
	})
	return drop
}

var slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		slice = gopurs_runtime.Func3(func(start_0 gopurs_runtime.Value, end_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(end_1.IntVal - start_0.IntVal), gopurs_runtime.Apply2(Get_drop(), start_0, xs_2))
})
	})
	return slice
}

var deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		deleteBy = gopurs_runtime.Func(func(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
deleteBy:
for {
if false { continue deleteBy }
var eq_0 = eq_0_loop
_ = eq_0
var x_1 = x_1_loop
_ = x_1
var xs_2 = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(eq_0, x_1, gopurs_runtime.RecordGet(__local_var_4_0, "value0"))).IntVal != 0 {
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(__local_var_4_0, "value1"))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(__local_var_4_0, "value0"), gopurs_runtime.Apply3(Get_deleteBy(), eq_0, x_1, gopurs_runtime.RecordGet(__local_var_4_0, "value1")))
}
end_branch_2:
__t1 = __t2
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
}
}()
})
})
})
	})
	return deleteBy
}

var unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		unionBy = gopurs_runtime.Func3(func(eq_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), xs_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_deleteBy(), eq_0, a_4, b_3)
}), gopurs_runtime.Apply2(Get_nubByEq(), eq_0, ys_2), xs_1))
})
	})
	return unionBy
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unionBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return union
}

var deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		deleteAt = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
deleteAt:
for {
if false { continue deleteAt }
var n_0 = n_0_loop
_ = n_0
var xs_1 = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_1)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_3_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_3_0, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(n_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(__local_var_3_0, "value1"))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(__local_var_3_0, "value0"), gopurs_runtime.Apply2(Get_deleteAt(), gopurs_runtime.Int(n_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordGet(__local_var_3_0, "value1")))
}
end_branch_2:
__t1 = __t2
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
}
}()
})
})
	})
	return deleteAt
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return delete_
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"), a_2, b_1)
}))
})
	})
	return difference
}

var cycle gopurs_runtime.Value
var once_cycle sync.Once
func Get_cycle() gopurs_runtime.Value {
	once_cycle.Do(func() {
		cycle = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), xs_0, go__1_0))
}))
return go__1_0
})
	})
	return cycle
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func2(func(b_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindList(), "bind"), a_1, b_0)
})
	})
	return concatMap
}

var concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		concat = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindList(), "bind"), v_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return concat
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Apply(Get_mapMaybe(), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
	})
	return catMaybes
}

var alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		alterAt = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
alterAt:
for {
if false { continue alterAt }
var n_0 = n_0_loop
_ = n_0
var f_1 = f_1_loop
_ = f_1
var xs_2 = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(n_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
v2_5_3 := gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(__local_var_4_0, "value0"))
_ = v2_5_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_5_3, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.RecordGet(__local_var_4_0, "value1"))
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_5_3, "_tag").StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_5_3, "value0"), gopurs_runtime.RecordGet(__local_var_4_0, "value1"))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t2 = __t4
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(__local_var_4_0, "value0"), gopurs_runtime.Apply3(Get_alterAt(), gopurs_runtime.Int(n_0.IntVal - gopurs_runtime.Int(1).IntVal), f_1, gopurs_runtime.RecordGet(__local_var_4_0, "value1")))
}
end_branch_2:
__t1 = __t2
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
}
}()
})
})
})
	})
	return alterAt
}

var modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		modifyAt = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_alterAt(), n_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Apply(f_1, x_2))
}))
})
	})
	return modifyAt
}


