package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Unfoldable1_Unfoldable1_dollar_Dict gopurs_runtime.Value
var once_Data_Unfoldable1_Unfoldable1_dollar_Dict sync.Once
func Get_Data_Unfoldable1_Unfoldable1_dollar_Dict() gopurs_runtime.Value {
	once_Data_Unfoldable1_Unfoldable1_dollar_Dict.Do(func() {
		cache_Data_Unfoldable1_Unfoldable1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Call_Data_Unfoldable1_Unfoldable1_dollar_Dict(func() struct{
	unfoldr1 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	unfoldr1 gopurs_runtime.Value
}{}
					clone.unfoldr1 = gopurs_runtime.RecordGet(orig, "unfoldr1")
					return clone
				}()))}
})
	})
	return cache_Data_Unfoldable1_Unfoldable1_dollar_Dict
}

var cache_Data_Unfoldable1_unfoldr1 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1 sync.Once
func Get_Data_Unfoldable1_unfoldr1() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1.Do(func() {
		cache_Data_Unfoldable1_unfoldr1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1
}

var cache_Data_Unfoldable1_unfoldable1Maybe gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldable1Maybe sync.Once
func Get_Data_Unfoldable1_unfoldable1Maybe() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldable1Maybe.Do(func() {
		cache_Data_Unfoldable1_unfoldable1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_4212083983_2187088110((&Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_0, b_1).UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
})})))}
	})
	return cache_Data_Unfoldable1_unfoldable1Maybe
}

var cache_Data_Unfoldable1_unfoldable1Array gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldable1Array sync.Once
func Get_Data_Unfoldable1_unfoldable1Array() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldable1Array.Do(func() {
		cache_Data_Unfoldable1_unfoldable1Array = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Apply4(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Unfoldable1_unfoldable1Array
}

var cache_Data_Unfoldable1_replicate1 gopurs_runtime.Value
var once_Data_Unfoldable1_replicate1 sync.Once
func Get_Data_Unfoldable1_replicate1() gopurs_runtime.Value {
	once_Data_Unfoldable1_replicate1.Do(func() {
		cache_Data_Unfoldable1_replicate1 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_replicate1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_Data_Unfoldable1_replicate1
}

var cache_Data_Unfoldable1_replicate1A gopurs_runtime.Value
var once_Data_Unfoldable1_replicate1A sync.Once
func Get_Data_Unfoldable1_replicate1A() gopurs_runtime.Value {
	once_Data_Unfoldable1_replicate1A.Do(func() {
		cache_Data_Unfoldable1_replicate1A = gopurs_runtime.Func5(func(dictApply_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value, dictTraversable1_2_box gopurs_runtime.Value, n_3_box gopurs_runtime.Value, m_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_replicate1A(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dictTraversable1_2_box), n_3_box.IntVal, m_4_box)
})
	})
	return cache_Data_Unfoldable1_replicate1A
}

var cache_Data_Unfoldable1_singleton gopurs_runtime.Value
var once_Data_Unfoldable1_singleton sync.Once
func Get_Data_Unfoldable1_singleton() gopurs_runtime.Value {
	once_Data_Unfoldable1_singleton.Do(func() {
		cache_Data_Unfoldable1_singleton = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_singleton(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), v_1_box)
})
	})
	return cache_Data_Unfoldable1_singleton
}

var cache_Data_Unfoldable1_go__range gopurs_runtime.Value
var once_Data_Unfoldable1_go__range sync.Once
func Get_Data_Unfoldable1_go__range() gopurs_runtime.Value {
	once_Data_Unfoldable1_go__range.Do(func() {
		cache_Data_Unfoldable1_go__range = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_go__range(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), start_1_box.IntVal, end_2_box.IntVal)
})
	})
	return cache_Data_Unfoldable1_go__range
}

var cache_Data_Unfoldable1_iterateN gopurs_runtime.Value
var once_Data_Unfoldable1_iterateN sync.Once
func Get_Data_Unfoldable1_iterateN() gopurs_runtime.Value {
	once_Data_Unfoldable1_iterateN.Do(func() {
		cache_Data_Unfoldable1_iterateN = gopurs_runtime.Func4(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_iterateN(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), n_1_box.IntVal, f_2_box, s_3_box)
})
	})
	return cache_Data_Unfoldable1_iterateN
}

type Constructor_Data_Unfoldable1_Unfoldable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3553002490] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Unfoldable1_Unfoldable1[any])(ptr)
		_ = c
		switch key {
		case "unfoldr1": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Unfoldable1_Unfoldable1: " + key)
		}
	}
}


func Call_Data_Unfoldable1_Unfoldable1_dollar_Dict(x_0_loop struct{
	unfoldr1 gopurs_runtime.Value
}) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
var x_0 struct{
	unfoldr1 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"unfoldr1"}, []gopurs_runtime.Value{orig.unfoldr1})
				}())
}

func Call_Data_Unfoldable1_unfoldr1(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_replicate1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]
{
if (i_3.IntVal) <= (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_3.IntVal) - (int64(1)))}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_3418986898_138441832(__t0))}
}), gopurs_runtime.Int((n_1) - (int64(1))))
}

func Call_Data_Unfoldable1_replicate1A(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], dictUnfoldable1_1_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], dictTraversable1_2_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value], n_3_loop int64, m_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
var dictUnfoldable1_1 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
var dictTraversable1_2 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dictTraversable1_2_loop
_ = dictTraversable1_2
var n_3 int64 = n_3_loop
_ = n_3
var m_4 gopurs_runtime.Value = m_4_loop
_ = m_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable1_2.V2), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_0)}, gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_1.V0), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]
{
if (i_5.IntVal) <= (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{m_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{m_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_5.IntVal) - (int64(1)))}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_3418986898_138441832(__t0))}
}), gopurs_runtime.Int((n_3) - (int64(1)))))
}

func Call_Data_Unfoldable1_singleton(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]
{
if (i_2.IntVal) <= (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_2.IntVal) - (int64(1)))}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_3418986898_138441832(__t0))}
}), gopurs_runtime.Int(int64(0)))
}

func Call_Data_Unfoldable1_go__range(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], start_1_loop int64, end_2_loop int64) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var start_1 int64 = start_1_loop
_ = start_1
var end_2 int64 = end_2_loop
_ = end_2
var __t1 int64
{
if (end_2) >= (start_1) {
__t1 = int64(1)
goto end_branch_1
} else {

}
}
{
__t1 = int64(-1)
}
end_branch_1:
// TAST (Let): __local_var_3_0 shape=Branch(LitInt, def=LitInt) bindingType=Int
__local_var_3_0 := __t1
_ = __local_var_3_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): i_prime__5_2 shape=Other bindingType=Int
i_prime__5_2 := (i_4.IntVal) + (__local_var_3_0)
_ = i_prime__5_2
var __t3 *Constructor_Data_Maybe_Just[int64]
{
if (i_4.IntVal) == (end_2) {
__t3 = Rebox_Data_Unfoldable1_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(i_prime__5_2), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_2874766729_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{i_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_1170268447_3094389156(__t3))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
}), gopurs_runtime.Int(start_1))
}

func Call_Data_Unfoldable1_iterateN(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], n_1_loop int64, f_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal) > (int64(0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal) - (int64(1)))}))}}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_1413047506_3094389156(Rebox_Data_Unfoldable1_3094389156_1413047506(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_4156246015_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, __t0}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_3415943795_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{s_3, gopurs_runtime.Int((n_1) - (int64(1)))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))})
}

func Rebox_Data_Unfoldable1_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Unfoldable1_138441832_3415943795(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
		out.V0 = in.V0
		out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Unfoldable1_1413047506_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_3415943795_138441832(in.V0))}
	return out
}

func Rebox_Data_Unfoldable1_2874766729_138441832(in *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_1170268447_3094389156(in.V1))}
	return out
}

func Rebox_Data_Unfoldable1_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Unfoldable1_3094389156_1413047506(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]{}
		out.V0 = Rebox_Data_Unfoldable1_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Unfoldable1_3415943795_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Unfoldable1_3418986898_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_1170268447_3094389156(in.V1))}
	return out
}

func Rebox_Data_Unfoldable1_4156246015_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Unfoldable1_1413047506_3094389156(in.V1))}
	return out
}

func Rebox_Data_Unfoldable1_4212083983_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Get_Data_Unfoldable1_unfoldr1ArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Unfoldable1_Unfoldr1ArrayImpl
}
