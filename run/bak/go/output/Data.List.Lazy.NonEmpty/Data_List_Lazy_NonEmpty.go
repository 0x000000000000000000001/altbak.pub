package Data_List_Lazy_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_List_Lazy "gopurs/output/Data.List.Lazy"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
)

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
return gopurs_runtime.RecordDict2("head", "tail", (*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[1])
})
	})
	return uncons
}

var toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		toList = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := (*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[0]
_ = __local_var_2_1
__local_var_3_2 := (*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[1]
_ = __local_var_3_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_2_1, __local_var_3_2)
}))
})
	})
	return toList
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_uncons(), xs_1)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_2_1.StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_2_1.UnsafePtr)[0], "head"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_2_1.UnsafePtr)[0], "tail")))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_2:
return __t2
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(Get_toList(), x_2))
})
})
	})
	return toUnfoldable
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr)[1]
})
	})
	return tail
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_applicativeNonEmptyList(), "pure")
	})
	return singleton
}

var repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		repeat = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", x_0, gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_repeat(), x_0))
}))
})
	})
	return repeat
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal + gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_length(), (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr)[1]).IntVal)
})
	})
	return length
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_last(), (*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[1])
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_2_1.StrVal == "Nothing")).IntVal != 0 {
__t2 = (*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[0]
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_2_1.StrVal == "Just")).IntVal != 0 {
__t2 = (*[1024]gopurs_runtime.Value)(__local_var_2_1.UnsafePtr)[0]
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
	})
	return last
}

var iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		iterate = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", x_1, gopurs_runtime.Apply2(pkg_Data_List_Lazy.Get_iterate(), f_0, gopurs_runtime.Apply(f_0, x_1)))
}))
})
	})
	return iterate
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_init_(), (*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[1])
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_2_1.StrVal == "Nothing")).IntVal != 0 {
__t2 = pkg_Data_List_Lazy_Types.Get_nil()
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_2_1.StrVal == "Just")).IntVal != 0 {
__local_var_3_3 := (*[1024]gopurs_runtime.Value)(__local_var_2_1.UnsafePtr)[0]
_ = __local_var_3_3
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[0], __local_var_3_3)
}))
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
	})
	return init_
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr)[0]
})
	})
	return head
}

var fromList gopurs_runtime.Value
var once_fromList sync.Once
func Get_fromList() gopurs_runtime.Value {
	once_fromList.Do(func() {
		fromList = gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
v_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_0)
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1_0.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1_0.StrVal == "Cons")).IntVal != 0 {
__local_var_2_2 := (*[1024]gopurs_runtime.Value)(v_1_0.UnsafePtr)[0]
_ = __local_var_2_2
__local_var_3_3 := (*[1024]gopurs_runtime.Value)(v_1_0.UnsafePtr)[1]
_ = __local_var_3_3
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", __local_var_2_2, __local_var_3_3)
})))
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
	return fromList
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_fromList(), gopurs_runtime.Apply(__local_var_1_0, x_2))
})
})
	})
	return fromFoldable
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func2(func(y_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1)
_ = v2_3_0
__local_var_4_1 := (*[1024]gopurs_runtime.Value)(v2_3_0.UnsafePtr)[0]
_ = __local_var_4_1
__local_var_5_2 := (*[1024]gopurs_runtime.Value)(v2_3_0.UnsafePtr)[1]
_ = __local_var_5_2
return gopurs_runtime.Constructor2("NonEmpty", y_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", __local_var_4_1, __local_var_5_2)
})))
}))
})
	})
	return cons
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func2(func(b_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindNonEmptyList(), "bind"), a_1, b_0)
})
	})
	return concatMap
}

var appendFoldable gopurs_runtime.Value
var once_appendFoldable sync.Once
func Get_appendFoldable() gopurs_runtime.Value {
	once_appendFoldable.Do(func() {
		appendFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
fromFoldable1_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
_ = fromFoldable1_1_0
return gopurs_runtime.Func2(func(nel_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), nel_2).UnsafePtr)[0], gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), nel_2).UnsafePtr)[1], gopurs_runtime.Apply(fromFoldable1_1_0, ys_3)))
}))
})
})
	})
	return appendFoldable
}


