package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Lazy_NonEmpty_uncons gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_uncons sync.Once
func Get_Data_List_Lazy_NonEmpty_uncons() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_uncons.Do(func() {
		cache_Data_List_Lazy_NonEmpty_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_uncons(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_uncons
}

var cache_Data_List_Lazy_NonEmpty_toList gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_toList sync.Once
func Get_Data_List_Lazy_NonEmpty_toList() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_toList.Do(func() {
		cache_Data_List_Lazy_NonEmpty_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_toList(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_toList
}

var cache_Data_List_Lazy_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_toUnfoldable sync.Once
func Get_Data_List_Lazy_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_List_Lazy_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_List_Lazy_NonEmpty_toUnfoldable
}

var cache_Data_List_Lazy_NonEmpty_tail gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_tail sync.Once
func Get_Data_List_Lazy_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_tail.Do(func() {
		cache_Data_List_Lazy_NonEmpty_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_tail(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_tail
}

var cache_Data_List_Lazy_NonEmpty_singleton gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_singleton sync.Once
func Get_Data_List_Lazy_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_singleton.Do(func() {
		cache_Data_List_Lazy_NonEmpty_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_singleton(a_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_singleton
}

var cache_Data_List_Lazy_NonEmpty_repeat gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_repeat sync.Once
func Get_Data_List_Lazy_NonEmpty_repeat() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_repeat.Do(func() {
		cache_Data_List_Lazy_NonEmpty_repeat = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_repeat(x_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_repeat
}

var cache_Data_List_Lazy_NonEmpty_length gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_length sync.Once
func Get_Data_List_Lazy_NonEmpty_length() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_length.Do(func() {
		cache_Data_List_Lazy_NonEmpty_length = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_List_Lazy_NonEmpty_length(v_0_box))
})
	})
	return cache_Data_List_Lazy_NonEmpty_length
}

var cache_Data_List_Lazy_NonEmpty_last gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_last sync.Once
func Get_Data_List_Lazy_NonEmpty_last() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_last.Do(func() {
		cache_Data_List_Lazy_NonEmpty_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_last(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_last
}

var cache_Data_List_Lazy_NonEmpty_iterate gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_iterate sync.Once
func Get_Data_List_Lazy_NonEmpty_iterate() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_iterate.Do(func() {
		cache_Data_List_Lazy_NonEmpty_iterate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_iterate(f_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_iterate
}

var cache_Data_List_Lazy_NonEmpty_init gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_init sync.Once
func Get_Data_List_Lazy_NonEmpty_init() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_init.Do(func() {
		cache_Data_List_Lazy_NonEmpty_init = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_init(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_init
}

var cache_Data_List_Lazy_NonEmpty_head gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_head sync.Once
func Get_Data_List_Lazy_NonEmpty_head() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_head.Do(func() {
		cache_Data_List_Lazy_NonEmpty_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_head(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_head
}

var cache_Data_List_Lazy_NonEmpty_fromList gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_fromList sync.Once
func Get_Data_List_Lazy_NonEmpty_fromList() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_fromList.Do(func() {
		cache_Data_List_Lazy_NonEmpty_fromList = gopurs_runtime.Func(func(l_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_NonEmpty_fromList(l_0_box))}
})
	})
	return cache_Data_List_Lazy_NonEmpty_fromList
}

var cache_Data_List_Lazy_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_fromFoldable sync.Once
func Get_Data_List_Lazy_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_fromFoldable.Do(func() {
		cache_Data_List_Lazy_NonEmpty_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_List_Lazy_NonEmpty_fromFoldable
}

var cache_Data_List_Lazy_NonEmpty_cons gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_cons sync.Once
func Get_Data_List_Lazy_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_cons.Do(func() {
		cache_Data_List_Lazy_NonEmpty_cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_cons(y_0_box, v_1_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_cons
}

var cache_Data_List_Lazy_NonEmpty_concatMap gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_concatMap sync.Once
func Get_Data_List_Lazy_NonEmpty_concatMap() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_concatMap.Do(func() {
		cache_Data_List_Lazy_NonEmpty_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_concatMap(b_0_box, a_1_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_concatMap
}

var cache_Data_List_Lazy_NonEmpty_appendFoldable gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_appendFoldable sync.Once
func Get_Data_List_Lazy_NonEmpty_appendFoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_appendFoldable.Do(func() {
		cache_Data_List_Lazy_NonEmpty_appendFoldable = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, nel_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_appendFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), nel_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_appendFoldable
}

var cache_Data_List_Lazy_NonEmpty_fromList__1361791219 gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_fromList__1361791219 sync.Once
func Get_Data_List_Lazy_NonEmpty_fromList__1361791219() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_fromList__1361791219.Do(func() {
		cache_Data_List_Lazy_NonEmpty_fromList__1361791219 = gopurs_runtime.Func(func(l_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_NonEmpty_fromList__1361791219(l_0_box))}
})
	})
	return cache_Data_List_Lazy_NonEmpty_fromList__1361791219
}

var cache_Data_List_Lazy_NonEmpty_head__3204870332 gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_head__3204870332 sync.Once
func Get_Data_List_Lazy_NonEmpty_head__3204870332() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_head__3204870332.Do(func() {
		cache_Data_List_Lazy_NonEmpty_head__3204870332 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_head__3204870332(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_head__3204870332
}

var cache_Data_List_Lazy_NonEmpty_tail__4101396777 gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_tail__4101396777 sync.Once
func Get_Data_List_Lazy_NonEmpty_tail__4101396777() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_tail__4101396777.Do(func() {
		cache_Data_List_Lazy_NonEmpty_tail__4101396777 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_tail__4101396777(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_tail__4101396777
}

var cache_Data_List_Lazy_NonEmpty_toList__1017592434 gopurs_runtime.Value
var once_Data_List_Lazy_NonEmpty_toList__1017592434 sync.Once
func Get_Data_List_Lazy_NonEmpty_toList__1017592434() gopurs_runtime.Value {
	once_Data_List_Lazy_NonEmpty_toList__1017592434.Do(func() {
		cache_Data_List_Lazy_NonEmpty_toList__1017592434 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_NonEmpty_toList__1017592434(v_0_box)
})
	})
	return cache_Data_List_Lazy_NonEmpty_toList__1017592434
}

func Call_Data_List_Lazy_NonEmpty_uncons(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 -> gopurs_runtime.Value
v1_1_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_1_0
return gopurs_runtime.RecordDict2("head", "tail", (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1)
}

func Call_Data_List_Lazy_NonEmpty_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 -> gopurs_runtime.Value
v1_1_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_1_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1})}
}))
}

func Call_Data_List_Lazy_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Maybe_Just
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_List_Lazy_uncons(), xs_1))
_ = __local_var_2_1
var __t2 *Constructor_Data_Maybe_Just
{
if (__local_var_2_1 != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet((__local_var_2_1).V0, "head"), gopurs_runtime.RecordGet((__local_var_2_1).V0, "tail")})}}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_3_3 -> gopurs_runtime.Value
v1_3_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2)
_ = v1_3_3
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_3_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v1_3_3.UnsafePtr).V1})}
})))
})
}

func Call_Data_List_Lazy_NonEmpty_tail(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V1
}

func Call_Data_List_Lazy_NonEmpty_singleton(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_0, Get_Data_List_Lazy_Types_nil()})}
}))
}

func Call_Data_List_Lazy_NonEmpty_repeat(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_0, gopurs_runtime.Apply(Get_Data_List_Lazy_repeat(), x_0)})}
}))
}

func Call_Data_List_Lazy_NonEmpty_length(v_0_loop gopurs_runtime.Value) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (1) + (gopurs_runtime.Apply(Get_Data_List_Lazy_length(), (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V1).IntVal)
}

func Call_Data_List_Lazy_NonEmpty_last(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 -> gopurs_runtime.Value
v1_1_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_1_0
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Maybe_Just
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_List_Lazy_last(), (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1))
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1 == nil) {
__t2 = (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1 != nil) {
__t2 = (__local_var_2_1).V0
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

func Call_Data_List_Lazy_NonEmpty_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_1, gopurs_runtime.Apply2(Get_Data_List_Lazy_iterate(), f_0, gopurs_runtime.Apply(f_0, x_1))})}
}))
}

func Call_Data_List_Lazy_NonEmpty_init(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 -> gopurs_runtime.Value
v1_1_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_1_0
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Maybe_Just
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_List_Lazy_init(), (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1))
_ = __local_var_2_1
var __t3 gopurs_runtime.Value
{
if (__local_var_2_1 == nil) {
__t3 = Get_Data_List_Lazy_Types_nil()
goto end_branch_3
} else {

}
}
{
if (__local_var_2_1 != nil) {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := (__local_var_2_1).V0
_ = __local_var_3_2
__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0, __local_var_3_2})}
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
}

func Call_Data_List_Lazy_NonEmpty_head(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V0
}

func Call_Data_List_Lazy_NonEmpty_fromList(l_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var l_0 gopurs_runtime.Value = l_0_loop
_ = l_0
// TAST (Let): v_1_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_0))
_ = v_1_0
var __t3 *Constructor_Data_Maybe_Just
{
if (v_1_0 == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (v_1_0 != nil) {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := (v_1_0).V0
_ = __local_var_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := (v_1_0).V1
_ = __local_var_3_2
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_2_1, __local_var_3_2})}
}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return __t3
}

func Call_Data_List_Lazy_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Lazy_Types_cons(), Get_Data_List_Lazy_Types_nil())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_NonEmpty_fromList(gopurs_runtime.Apply(__local_var_1_0, x_2)))}
})
}

func Call_Data_List_Lazy_NonEmpty_cons(y_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_3_0 -> gopurs_runtime.Value
v2_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1)
_ = v2_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, y_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v2_3_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v2_3_0.UnsafePtr).V1})}
}))})}
}))
}

func Call_Data_List_Lazy_NonEmpty_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1)
_ = v1_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): v2_4_2 -> *Constructor_Data_NonEmpty_NonEmpty
v2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(b_0, (*Constructor_Data_NonEmpty_NonEmpty)(v1_2_0.UnsafePtr).V0)))
_ = v2_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (v2_4_2).V0
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (v2_4_2).V1
_ = __local_var_6_4
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1)
_ = __local_var_9_6
var __t27 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_9_6.Type == 9 && __local_var_9_6.IntVal == 218341868 && __local_var_9_6.UnsafePtr == nil) {
__t27 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_27
} else {

}
}
{
if (__local_var_9_6.Type == 9 && __local_var_9_6.IntVal == 218341868 && __local_var_9_6.UnsafePtr != nil) {
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := gopurs_runtime.Apply(Get_Data_List_Lazy_Types_toList(), gopurs_runtime.Apply(b_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_9_6.UnsafePtr).V0))
_ = __local_var_10_7
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_9_6.UnsafePtr).V1
_ = __local_var_11_9
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_10 -> gopurs_runtime.Value
__local_var_13_10 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_9)
_ = __local_var_13_10
var __t18 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_10.Type == 9 && __local_var_13_10.IntVal == 218341868 && __local_var_13_10.UnsafePtr == nil) {
__t18 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_18
} else {

}
}
{
if (__local_var_13_10.Type == 9 && __local_var_13_10.IntVal == 218341868 && __local_var_13_10.UnsafePtr != nil) {
// TAST (Let): __local_var_14_11 -> gopurs_runtime.Value
__local_var_14_11 := gopurs_runtime.Apply(Get_Data_List_Lazy_Types_toList(), gopurs_runtime.Apply(b_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_10.UnsafePtr).V0))
_ = __local_var_14_11
// TAST (Let): __local_var_15_12 -> gopurs_runtime.Value
__local_var_15_12 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_10.UnsafePtr).V1, gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_List_Lazy_Types_toList(), gopurs_runtime.Apply(b_0, x_15))
}))
_ = __local_var_15_12
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_13 -> gopurs_runtime.Value
__local_var_17_13 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_11)
_ = __local_var_17_13
var __t17 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_17_13.Type == 9 && __local_var_17_13.IntVal == 218341868 && __local_var_17_13.UnsafePtr == nil) {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_12))
goto end_branch_17
} else {

}
}
{
if (__local_var_17_13.Type == 9 && __local_var_17_13.IntVal == 218341868 && __local_var_17_13.UnsafePtr != nil) {
// TAST (Let): __local_var_18_14 -> gopurs_runtime.Value
__local_var_18_14 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_13.UnsafePtr).V1
_ = __local_var_18_14
__t17 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_17_13.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_15 -> gopurs_runtime.Value
__local_var_20_15 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_18_14)
_ = __local_var_20_15
var __t16 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_20_15.Type == 9 && __local_var_20_15.IntVal == 218341868 && __local_var_20_15.UnsafePtr == nil) {
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_15_12))
goto end_branch_16
} else {

}
}
{
if (__local_var_20_15.Type == 9 && __local_var_20_15.IntVal == 218341868 && __local_var_20_15.UnsafePtr != nil) {
__t16 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_15.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_20_15.UnsafePtr).V1, __local_var_15_12)}
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t16)}
}))}
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t17)}
}))))
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t18)}
}))
_ = __local_var_11_8
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_19 -> gopurs_runtime.Value
__local_var_13_19 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_7)
_ = __local_var_13_19
var __t26 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_19.Type == 9 && __local_var_13_19.IntVal == 218341868 && __local_var_13_19.UnsafePtr == nil) {
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_26
} else {

}
}
{
if (__local_var_13_19.Type == 9 && __local_var_13_19.IntVal == 218341868 && __local_var_13_19.UnsafePtr != nil) {
// TAST (Let): __local_var_14_20 -> gopurs_runtime.Value
__local_var_14_20 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_19.UnsafePtr).V1
_ = __local_var_14_20
__t26 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_19.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_21 -> gopurs_runtime.Value
__local_var_16_21 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_14_20)
_ = __local_var_16_21
var __t25 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_16_21.Type == 9 && __local_var_16_21.IntVal == 218341868 && __local_var_16_21.UnsafePtr == nil) {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_25
} else {

}
}
{
if (__local_var_16_21.Type == 9 && __local_var_16_21.IntVal == 218341868 && __local_var_16_21.UnsafePtr != nil) {
// TAST (Let): __local_var_17_22 -> gopurs_runtime.Value
__local_var_17_22 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_21.UnsafePtr).V1
_ = __local_var_17_22
__t25 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_16_21.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_23 -> gopurs_runtime.Value
__local_var_19_23 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_17_22)
_ = __local_var_19_23
var __t24 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_19_23.Type == 9 && __local_var_19_23.IntVal == 218341868 && __local_var_19_23.UnsafePtr == nil) {
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_8))
goto end_branch_24
} else {

}
}
{
if (__local_var_19_23.Type == 9 && __local_var_19_23.IntVal == 218341868 && __local_var_19_23.UnsafePtr != nil) {
__t24 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_23.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_19_23.UnsafePtr).V1, __local_var_11_8)}
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_24:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t24)}
}))}
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_25:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t25)}
}))}
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_26:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t26)}
}))))
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t27)}
}))
_ = __local_var_8_5
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_5_3, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_28 -> gopurs_runtime.Value
__local_var_10_28 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4)
_ = __local_var_10_28
var __t32 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_10_28.Type == 9 && __local_var_10_28.IntVal == 218341868 && __local_var_10_28.UnsafePtr == nil) {
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_5))
goto end_branch_32
} else {

}
}
{
if (__local_var_10_28.Type == 9 && __local_var_10_28.IntVal == 218341868 && __local_var_10_28.UnsafePtr != nil) {
// TAST (Let): __local_var_11_29 -> gopurs_runtime.Value
__local_var_11_29 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_10_28.UnsafePtr).V1
_ = __local_var_11_29
__t32 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_10_28.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_30 -> gopurs_runtime.Value
__local_var_13_30 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_11_29)
_ = __local_var_13_30
var __t31 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_13_30.Type == 9 && __local_var_13_30.IntVal == 218341868 && __local_var_13_30.UnsafePtr == nil) {
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_5))
goto end_branch_31
} else {

}
}
{
if (__local_var_13_30.Type == 9 && __local_var_13_30.IntVal == 218341868 && __local_var_13_30.UnsafePtr != nil) {
__t31 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_30.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_13_30.UnsafePtr).V1, __local_var_8_5)}
goto end_branch_31
} else {

}
}
{
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_31:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t31)}
}))}
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_32:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t32)}
}))})}
}))
}

func Call_Data_List_Lazy_NonEmpty_appendFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, nel_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var nel_1 gopurs_runtime.Value = nel_1_loop
_ = nel_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), nel_1).UnsafePtr).V1
_ = __local_var_4_0
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Lazy_Types_cons(), Get_Data_List_Lazy_Types_nil(), ys_2)
_ = __local_var_5_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), nel_1).UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_2 -> gopurs_runtime.Value
__local_var_7_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_0)
_ = __local_var_7_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 218341868 && __local_var_7_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_1))
goto end_branch_3
} else {

}
}
{
if (__local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 218341868 && __local_var_7_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_2.UnsafePtr).V1, __local_var_5_1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))})}
}))
}

func Call_Data_List_Lazy_NonEmpty_fromList__1361791219(l_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var l_0 gopurs_runtime.Value = l_0_loop
_ = l_0
// TAST (Let): v_1_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_0))
_ = v_1_0
var __t3 *Constructor_Data_Maybe_Just
{
if (v_1_0 == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (v_1_0 != nil) {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := (v_1_0).V0
_ = __local_var_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := (v_1_0).V1
_ = __local_var_3_2
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_2_1, __local_var_3_2})}
}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return __t3
}

func Call_Data_List_Lazy_NonEmpty_head__3204870332(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V0
}

func Call_Data_List_Lazy_NonEmpty_tail__4101396777(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_Data_NonEmpty_NonEmpty)(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0).UnsafePtr).V1
}

func Call_Data_List_Lazy_NonEmpty_toList__1017592434(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 -> gopurs_runtime.Value
v1_1_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = v1_1_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1_0.UnsafePtr).V1})}
}))
}


