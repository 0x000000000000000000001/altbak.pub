package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Unfoldable_Unfoldable_dollar_Dict gopurs_runtime.Value
var once_Data_Unfoldable_Unfoldable_dollar_Dict sync.Once
func Get_Data_Unfoldable_Unfoldable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Unfoldable_Unfoldable_dollar_Dict.Do(func() {
		cache_Data_Unfoldable_Unfoldable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Call_Data_Unfoldable_Unfoldable_dollar_Dict(func() struct{
	Unfoldable10 gopurs_runtime.Value
	unfoldr gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Unfoldable10 gopurs_runtime.Value
	unfoldr gopurs_runtime.Value
}{}
					clone.Unfoldable10 = gopurs_runtime.RecordGet(orig, "Unfoldable10")
					clone.unfoldr = gopurs_runtime.RecordGet(orig, "unfoldr")
					return clone
				}()))}
})
	})
	return cache_Data_Unfoldable_Unfoldable_dollar_Dict
}

var cache_Data_Unfoldable_unfoldr gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr sync.Once
func Get_Data_Unfoldable_unfoldr() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr.Do(func() {
		cache_Data_Unfoldable_unfoldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr
}

var cache_Data_Unfoldable_unfoldr__3438611141 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__3438611141 sync.Once
func Get_Data_Unfoldable_unfoldr__3438611141() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__3438611141.Do(func() {
		cache_Data_Unfoldable_unfoldr__3438611141 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__3438611141(__eta_norm_0_unused_0_box, b_1_box)
})
	})
	return cache_Data_Unfoldable_unfoldr__3438611141
}

var cache_Data_Unfoldable_unfoldr__158397599 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__158397599 sync.Once
func Get_Data_Unfoldable_unfoldr__158397599() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__158397599.Do(func() {
		cache_Data_Unfoldable_unfoldr__158397599 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__158397599(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Unfoldable_unfoldr__158397599
}

var cache_Data_Unfoldable_unfoldableMaybe gopurs_runtime.Value
var once_Data_Unfoldable_unfoldableMaybe sync.Once
func Get_Data_Unfoldable_unfoldableMaybe() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldableMaybe.Do(func() {
		cache_Data_Unfoldable_unfoldableMaybe = gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_1401495535_2738507278((&Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_4212083983_2187088110(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Unfoldable1_unfoldable1Maybe())))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
__local_var_2_0 := Rebox_Data_Unfoldable_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, b_1)))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{((__local_var_2_0).V0).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Unfoldable_unfoldableMaybe
}

var cache_Data_Unfoldable_unfoldableArray gopurs_runtime.Value
var once_Data_Unfoldable_unfoldableArray sync.Once
func Get_Data_Unfoldable_unfoldableArray() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldableArray.Do(func() {
		cache_Data_Unfoldable_unfoldableArray = gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array()))}
}), gopurs_runtime.Apply4(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_0 != nil) {
__t1 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), Get_Data_Tuple_fst(), Get_Data_Tuple_snd())}))}
	})
	return cache_Data_Unfoldable_unfoldableArray
}

var cache_Data_Unfoldable_replicate gopurs_runtime.Value
var once_Data_Unfoldable_replicate sync.Once
func Get_Data_Unfoldable_replicate() gopurs_runtime.Value {
	once_Data_Unfoldable_replicate.Do(func() {
		cache_Data_Unfoldable_replicate = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_replicate(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_Data_Unfoldable_replicate
}

var cache_Data_Unfoldable_replicateA gopurs_runtime.Value
var once_Data_Unfoldable_replicateA sync.Once
func Get_Data_Unfoldable_replicateA() gopurs_runtime.Value {
	once_Data_Unfoldable_replicateA.Do(func() {
		cache_Data_Unfoldable_replicateA = gopurs_runtime.Func5(func(dictApplicative_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value, n_3_box gopurs_runtime.Value, m_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_replicateA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_2_box), n_3_box.IntVal, m_4_box)
})
	})
	return cache_Data_Unfoldable_replicateA
}

var cache_Data_Unfoldable_none gopurs_runtime.Value
var once_Data_Unfoldable_none sync.Once
func Get_Data_Unfoldable_none() gopurs_runtime.Value {
	once_Data_Unfoldable_none.Do(func() {
		cache_Data_Unfoldable_none = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_none(dictUnfoldable_0_box)
})
	})
	return cache_Data_Unfoldable_none
}

var cache_Data_Unfoldable_fromMaybe gopurs_runtime.Value
var once_Data_Unfoldable_fromMaybe sync.Once
func Get_Data_Unfoldable_fromMaybe() gopurs_runtime.Value {
	once_Data_Unfoldable_fromMaybe.Do(func() {
		cache_Data_Unfoldable_fromMaybe = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_fromMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_Data_Unfoldable_fromMaybe
}

type Constructor_Data_Unfoldable_Unfoldable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2670894170] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Unfoldable_Unfoldable[any])(ptr)
		_ = c
		switch key {
		case "Unfoldable10": return gopurs_runtime.Box(c.V0)
		case "unfoldr": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Unfoldable_Unfoldable: " + key)
		}
	}
}


func Call_Data_Unfoldable_Unfoldable_dollar_Dict(x_0_loop struct{
	Unfoldable10 gopurs_runtime.Value
	unfoldr gopurs_runtime.Value
}) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
var x_0 struct{
	Unfoldable10 gopurs_runtime.Value
	unfoldr gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", orig.Unfoldable10, orig.unfoldr)
				}())
}

func Call_Data_Unfoldable_unfoldr(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__3438611141(__eta_norm_0_unused_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
unfoldr__3438611141:
for {
if false { continue unfoldr__3438611141 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var Call_local_Data_Unfoldable_go__go_2_0_0 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_Unfoldable_go__go_2_0_0
var go__go_2_0_0 gopurs_runtime.Value
_ = go__go_2_0_0
Call_local_Data_Unfoldable_go__go_2_0_0 = func(source_3_loop gopurs_runtime.Value, memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_0_0:
for {
if false { continue go__go_2_0_0 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
v_5_1 := Rebox_Data_Unfoldable_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepUnfoldr(), source_3)))
_ = v_5_1
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_5_1 == nil) {
var Call_local_Data_Unfoldable_go__go_6_2_1 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Unfoldable_go__go_6_2_1
var go__go_6_2_1 gopurs_runtime.Value
_ = go__go_6_2_1
Call_local_Data_Unfoldable_go__go_6_2_1 = func(b_7_loop gopurs_runtime.Value, v_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_6_2_1:
for {
if false { continue go__go_6_2_1 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_8_loop
_ = v_8
var __t3 gopurs_runtime.Value
{
if (v_8 == nil) {
__t3 = b_7
goto end_branch_3
} else {

}
}
{
if (v_8 != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_8).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_7)}))}
v_8_loop = (v_8).V1
continue go__go_6_2_1
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_6_2_1 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Unfoldable_go__go_6_2_1(b_7_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_8_loop_val))
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_Unfoldable_go__go_6_2_1(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, memo_4))
goto end_branch_4
} else {

}
}
{
if (v_5_1 != nil) {
source_3_loop = ((v_5_1).V0).V1
memo_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, ((v_5_1).V0).V0, memo_4})
continue go__go_2_0_0
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_2_0_0 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_Unfoldable_go__go_2_0_0(source_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_Unfoldable_go__go_2_0_0(b_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))}
}
}

func Call_Data_Unfoldable_unfoldr__158397599(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
unfoldr__158397599:
for {
if false { continue unfoldr__158397599 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()).V1), Get_Data_String_CodePoints_unconsButWithTuple(), __eta_norm_0_1)
}
}

func Call_Data_Unfoldable_replicate(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]
{
if (i_3.IntVal) <= (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Int((i_3.IntVal) - (int64(1)))}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_1413047506_3094389156(__t0))}
}), gopurs_runtime.Int(n_1))
}

func Call_Data_Unfoldable_replicateA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictUnfoldable_1_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], dictTraversable_2_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], n_3_loop int64, m_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var dictUnfoldable_1 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_1_loop
_ = dictUnfoldable_1
var dictTraversable_2 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
_ = dictTraversable_2
var n_3 int64 = n_3_loop
_ = n_3
var m_4 gopurs_runtime.Value = m_4_loop
_ = m_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_2.V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_1.V1), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]
{
if (i_5.IntVal) <= (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, m_4, gopurs_runtime.Int((i_5.IntVal) - (int64(1)))}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_1413047506_3094389156(__t0))}
}), gopurs_runtime.Int(n_3)))
}

func Call_Data_Unfoldable_none(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
}), Get_Data_Unit_unit())
}

func Call_Data_Unfoldable_fromMaybe(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](b_1)
if (__t_tag_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(b_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
}

func Rebox_Data_Unfoldable_1401495535_2738507278(in *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Unfoldable_1413047506_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable_3415943795_138441832(in.V0))}
	return out
}

func Rebox_Data_Unfoldable_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Unfoldable_3415943795_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Unfoldable_4212083983_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Get_Data_Unfoldable_unfoldrArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Unfoldable_UnfoldrArrayImpl
}
