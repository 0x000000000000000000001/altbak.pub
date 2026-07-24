package Data_List_Lazy_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_List_Lazy "gopurs/output/Data.List.Lazy"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
return gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1)
}()
})
	})
	return uncons
}

var toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		toList = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0
_ = __local_var_2_1
__local_var_3_2 := (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1
_ = __local_var_3_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 597046057, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Data_Data_List_Lazy_Types_Cons{__local_var_2_1, __local_var_3_2})}
}))
}()
})
	})
	return toList
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_uncons(), xs_1)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 1354639136) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.RecordGet((*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_1.UnsafePtr).V0, "head"), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_1.UnsafePtr).V0, "tail")})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_2:
return __t2
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(Get_toList(), x_2))
})
}()
})
	})
	return toUnfoldable
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V1
}()
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
		repeat = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1104112642, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{x_0, gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_repeat(), x_0)})}
}))
}()
})
	})
	return repeat
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Int(1 + gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_length(), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V1).IntVal)
}()
})
	})
	return length
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_last(), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 42808261) {
__t2 = (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 1354639136) {
__t2 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}()
})
	})
	return last
}

var iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		iterate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate(f_0_box, x_1_box)
})
	})
	return iterate
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_init_(), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 42808261) {
__t2 = pkg_Data_List_Lazy_Types.Get_nil()
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 1354639136) {
__local_var_3_3 := (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_1.UnsafePtr).V0
_ = __local_var_3_3
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 597046057, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Data_Data_List_Lazy_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0, __local_var_3_3})}
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
}()
})
	})
	return init_
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V0
}()
})
	})
	return head
}

var fromList gopurs_runtime.Value
var once_fromList sync.Once
func Get_fromList() gopurs_runtime.Value {
	once_fromList.Do(func() {
		fromList = gopurs_runtime.Func(func(l_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var l_0 gopurs_runtime.Value = l_0_loop
_ = l_0
v_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_0)
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 1330577875) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 597046057) {
__local_var_2_2 := (*pkg_Data_List_Lazy_Types.Data_Data_List_Lazy_Types_Cons)(v_1_0.UnsafePtr).V0
_ = __local_var_2_2
__local_var_3_3 := (*pkg_Data_List_Lazy_Types.Data_Data_List_Lazy_Types_Cons)(v_1_0.UnsafePtr).V1
_ = __local_var_3_3
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1104112642, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{__local_var_2_2, __local_var_3_3})}
}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}()
})
	})
	return fromList
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_fromList(), gopurs_runtime.Apply(__local_var_1_0, x_2))
})
}()
})
	})
	return fromFoldable
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons(y_0_box, v_1_box)
})
	})
	return cons
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concatMap(b_0_box, a_1_box)
})
	})
	return concatMap
}

var appendFoldable gopurs_runtime.Value
var once_appendFoldable sync.Once
func Get_appendFoldable() gopurs_runtime.Value {
	once_appendFoldable.Do(func() {
		appendFoldable = gopurs_runtime.Func(func(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
fromFoldable1_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
_ = fromFoldable1_1_0
return gopurs_runtime.Func2(func(nel_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1104112642, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), nel_2).UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), nel_2).UnsafePtr).V1, gopurs_runtime.Apply(fromFoldable1_1_0, ys_3))})}
}))
})
}()
})
	})
	return appendFoldable
}

func Call_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1104112642, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{x_1, gopurs_runtime.Apply2(pkg_Data_List_Lazy.Get_iterate(), f_0, gopurs_runtime.Apply(f_0, x_1))})}
}))
}

func Call_cons(y_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1)
_ = v2_3_0
__local_var_4_1 := (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v2_3_0.UnsafePtr).V0
_ = __local_var_4_1
__local_var_5_2 := (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v2_3_0.UnsafePtr).V1
_ = __local_var_5_2
return gopurs_runtime.Value{Type: 9, IntVal: 1104112642, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{y_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 597046057, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Data_Data_List_Lazy_Types_Cons{__local_var_4_1, __local_var_5_2})}
}))})}
}))
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindNonEmptyList(), "bind"), a_1, b_0)
}


