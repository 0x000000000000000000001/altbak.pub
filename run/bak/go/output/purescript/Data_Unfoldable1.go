package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Unfoldable1_Unfoldable1_dollarDict gopurs_runtime.Value
var once_Data_Unfoldable1_Unfoldable1_dollarDict sync.Once
func Get_Data_Unfoldable1_Unfoldable1_dollarDict() gopurs_runtime.Value {
	once_Data_Unfoldable1_Unfoldable1_dollarDict.Do(func() {
		cache_Data_Unfoldable1_Unfoldable1_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_Unfoldable1_dollarDict(x_0_box)
})
	})
	return cache_Data_Unfoldable1_Unfoldable1_dollarDict
}

var cache_Data_Unfoldable1_unfoldr1 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1 sync.Once
func Get_Data_Unfoldable1_unfoldr1() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1.Do(func() {
		cache_Data_Unfoldable1_unfoldr1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1
}

var cache_Data_Unfoldable1_unfoldable1Maybe gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldable1Maybe sync.Once
func Get_Data_Unfoldable1_unfoldable1Maybe() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldable1Maybe.Do(func() {
		cache_Data_Unfoldable1_unfoldable1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_0, b_1).UnsafePtr).V0})}
})
})})}
	})
	return cache_Data_Unfoldable1_unfoldable1Maybe
}

var cache_Data_Unfoldable1_unfoldable1Array gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldable1Array sync.Once
func Get_Data_Unfoldable1_unfoldable1Array() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldable1Array.Do(func() {
		cache_Data_Unfoldable1_unfoldable1Array = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, gopurs_runtime.Apply4(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = (*Constructor_Data_Maybe_Just)(v_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd())})}
	})
	return cache_Data_Unfoldable1_unfoldable1Array
}

var cache_Data_Unfoldable1_replicate1 gopurs_runtime.Value
var once_Data_Unfoldable1_replicate1 sync.Once
func Get_Data_Unfoldable1_replicate1() gopurs_runtime.Value {
	once_Data_Unfoldable1_replicate1.Do(func() {
		cache_Data_Unfoldable1_replicate1 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_replicate1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_Data_Unfoldable1_replicate1
}

var cache_Data_Unfoldable1_replicate1A gopurs_runtime.Value
var once_Data_Unfoldable1_replicate1A sync.Once
func Get_Data_Unfoldable1_replicate1A() gopurs_runtime.Value {
	once_Data_Unfoldable1_replicate1A.Do(func() {
		cache_Data_Unfoldable1_replicate1A = gopurs_runtime.Func5(func(dictApply_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value, dictTraversable1_2_box gopurs_runtime.Value, n_3_box gopurs_runtime.Value, m_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_replicate1A(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](dictTraversable1_2_box), n_3_box.IntVal, m_4_box)
})
	})
	return cache_Data_Unfoldable1_replicate1A
}

var cache_Data_Unfoldable1_singleton gopurs_runtime.Value
var once_Data_Unfoldable1_singleton sync.Once
func Get_Data_Unfoldable1_singleton() gopurs_runtime.Value {
	once_Data_Unfoldable1_singleton.Do(func() {
		cache_Data_Unfoldable1_singleton = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_singleton(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box), v_1_box)
})
	})
	return cache_Data_Unfoldable1_singleton
}

var cache_Data_Unfoldable1_go__range gopurs_runtime.Value
var once_Data_Unfoldable1_go__range sync.Once
func Get_Data_Unfoldable1_go__range() gopurs_runtime.Value {
	once_Data_Unfoldable1_go__range.Do(func() {
		cache_Data_Unfoldable1_go__range = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_go__range(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box), start_1_box.IntVal, end_2_box.IntVal)
})
	})
	return cache_Data_Unfoldable1_go__range
}

var cache_Data_Unfoldable1_iterateN gopurs_runtime.Value
var once_Data_Unfoldable1_iterateN sync.Once
func Get_Data_Unfoldable1_iterateN() gopurs_runtime.Value {
	once_Data_Unfoldable1_iterateN.Do(func() {
		cache_Data_Unfoldable1_iterateN = gopurs_runtime.Func4(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_iterateN(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box), n_1_box.IntVal, f_2_box, s_3_box)
})
	})
	return cache_Data_Unfoldable1_iterateN
}

var cache_Data_Unfoldable1_replicate1__3169098027 gopurs_runtime.Value
var once_Data_Unfoldable1_replicate1__3169098027 sync.Once
func Get_Data_Unfoldable1_replicate1__3169098027() gopurs_runtime.Value {
	once_Data_Unfoldable1_replicate1__3169098027.Do(func() {
		cache_Data_Unfoldable1_replicate1__3169098027 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_replicate1__3169098027(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_Data_Unfoldable1_replicate1__3169098027
}

var cache_Data_Unfoldable1_replicate1__3087621739 gopurs_runtime.Value
var once_Data_Unfoldable1_replicate1__3087621739 sync.Once
func Get_Data_Unfoldable1_replicate1__3087621739() gopurs_runtime.Value {
	once_Data_Unfoldable1_replicate1__3087621739.Do(func() {
		cache_Data_Unfoldable1_replicate1__3087621739 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_replicate1__3087621739(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_Data_Unfoldable1_replicate1__3087621739
}

var cache_Data_Unfoldable1_singleton__1620623815 gopurs_runtime.Value
var once_Data_Unfoldable1_singleton__1620623815 sync.Once
func Get_Data_Unfoldable1_singleton__1620623815() gopurs_runtime.Value {
	once_Data_Unfoldable1_singleton__1620623815.Do(func() {
		cache_Data_Unfoldable1_singleton__1620623815 = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_singleton__1620623815(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box), v_1_box)
})
	})
	return cache_Data_Unfoldable1_singleton__1620623815
}

var cache_Data_Unfoldable1_unfoldr1__2739418437 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1__2739418437 sync.Once
func Get_Data_Unfoldable1_unfoldr1__2739418437() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1__2739418437.Do(func() {
		cache_Data_Unfoldable1_unfoldr1__2739418437 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1__2739418437(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1__2739418437
}

var cache_Data_Unfoldable1_unfoldr1__169691813 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1__169691813 sync.Once
func Get_Data_Unfoldable1_unfoldr1__169691813() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1__169691813.Do(func() {
		cache_Data_Unfoldable1_unfoldr1__169691813 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1__169691813(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1__169691813
}

var cache_Data_Unfoldable1_unfoldr1__3410377797 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1__3410377797 sync.Once
func Get_Data_Unfoldable1_unfoldr1__3410377797() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1__3410377797.Do(func() {
		cache_Data_Unfoldable1_unfoldr1__3410377797 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1__3410377797(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1__3410377797
}

var cache_Data_Unfoldable1_unfoldr1__2402610528 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1__2402610528 sync.Once
func Get_Data_Unfoldable1_unfoldr1__2402610528() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1__2402610528.Do(func() {
		cache_Data_Unfoldable1_unfoldr1__2402610528 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1__2402610528(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1__2402610528
}

var cache_Data_Unfoldable1_unfoldr1__630091776 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1__630091776 sync.Once
func Get_Data_Unfoldable1_unfoldr1__630091776() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1__630091776.Do(func() {
		cache_Data_Unfoldable1_unfoldr1__630091776 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1__630091776(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1__630091776
}

var cache_Data_Unfoldable1_unfoldr1__1597580941 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1__1597580941 sync.Once
func Get_Data_Unfoldable1_unfoldr1__1597580941() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1__1597580941.Do(func() {
		cache_Data_Unfoldable1_unfoldr1__1597580941 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1__1597580941(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1__1597580941
}

var cache_Data_Unfoldable1_unfoldr1__78405858 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1__78405858 sync.Once
func Get_Data_Unfoldable1_unfoldr1__78405858() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1__78405858.Do(func() {
		cache_Data_Unfoldable1_unfoldr1__78405858 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1__78405858(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dict_0_box))
})
	})
	return cache_Data_Unfoldable1_unfoldr1__78405858
}

var cache_Data_Unfoldable1_unfoldr1__2387656390 gopurs_runtime.Value
var once_Data_Unfoldable1_unfoldr1__2387656390 sync.Once
func Get_Data_Unfoldable1_unfoldr1__2387656390() gopurs_runtime.Value {
	once_Data_Unfoldable1_unfoldr1__2387656390.Do(func() {
		cache_Data_Unfoldable1_unfoldr1__2387656390 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable1_unfoldr1__2387656390(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Unfoldable1_unfoldr1__2387656390
}

type Constructor_Data_Unfoldable1_Unfoldable1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3553002490] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Unfoldable1_Unfoldable1)(ptr)
		_ = c
		switch key {
		case "unfoldr1": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Unfoldable1_Unfoldable1: " + key)
		}
	}
}


func Call_Data_Unfoldable1_Unfoldable1_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Unfoldable1_unfoldr1(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_replicate1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Tuple_Tuple
{
var __t0 bool
{
if (i_3.IntVal) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_3.IntVal) - (1))})}}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Int((n_1) - (1)))
}

func Call_Data_Unfoldable1_replicate1A(dictApply_0_loop *Constructor_Control_Apply_Apply, dictUnfoldable1_1_loop *Constructor_Data_Unfoldable1_Unfoldable1, dictTraversable1_2_loop *Constructor_Data_Semigroup_Traversable_Traversable1, n_3_loop int64, m_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
var dictUnfoldable1_1 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
var dictTraversable1_2 *Constructor_Data_Semigroup_Traversable_Traversable1 = dictTraversable1_2_loop
_ = dictTraversable1_2
var n_3 int64 = n_3_loop
_ = n_3
var m_4 gopurs_runtime.Value = m_4_loop
_ = m_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable1_2.V2), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_0)}, gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_1.V0), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Tuple_Tuple
{
var __t0 bool
{
if (i_5.IntVal) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Tuple_Tuple{1, m_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Tuple_Tuple{1, m_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_5.IntVal) - (1))})}}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Int((n_3) - (1))))
}

func Call_Data_Unfoldable1_singleton(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Tuple_Tuple
{
var __t0 bool
{
if (i_2.IntVal) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Tuple_Tuple{1, v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Tuple_Tuple{1, v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_2.IntVal) - (1))})}}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Int(0))
}

func Call_Data_Unfoldable1_go__range(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1, start_1_loop int64, end_2_loop int64) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var start_1 int64 = start_1_loop
_ = start_1
var end_2 int64 = end_2_loop
_ = end_2
var __t2 int64
{
var __t1 bool
{
if (end_2) < (start_1) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
if __t1 {
__t2 = 1
goto end_branch_2
} else {

}
}
{
__t2 = -1
}
end_branch_2:
// TAST (Let): __local_var_3_0 -> int64
__local_var_3_0 := __t2
_ = __local_var_3_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): i_prime_5_3 -> int64
i_prime_5_3 := (i_4.IntVal) + (__local_var_3_0)
_ = i_prime_5_3
var __t4 *Constructor_Data_Maybe_Just
{
if (i_4.IntVal) == (end_2) {
__t4 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_4
} else {

}
}
{
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(i_prime_5_3)}
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(i_4.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}})}
}), gopurs_runtime.Int(start_1))
}

func Call_Data_Unfoldable1_iterateN(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1, n_1_loop int64, f_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1.IntVal) > (0) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1.IntVal) - (1))})}}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, s_3, gopurs_runtime.Int((n_1) - (1))})})
}

func Call_Data_Unfoldable1_replicate1__3169098027(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Tuple_Tuple
{
var __t0 bool
{
if (i_3.IntVal) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_3.IntVal) - (1))})}}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Int((n_1) - (1)))
}

func Call_Data_Unfoldable1_replicate1__3087621739(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Tuple_Tuple
{
var __t0 bool
{
if (i_3.IntVal) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_3.IntVal) - (1))})}}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Int((n_1) - (1)))
}

func Call_Data_Unfoldable1_singleton__1620623815(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Tuple_Tuple
{
var __t0 bool
{
if (i_2.IntVal) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Tuple_Tuple{1, v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Tuple_Tuple{1, v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_2.IntVal) - (1))})}}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Int(0))
}

func Call_Data_Unfoldable1_unfoldr1__2739418437(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_unfoldr1__169691813(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_unfoldr1__3410377797(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_unfoldr1__2402610528(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_unfoldr1__630091776(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_unfoldr1__1597580941(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_unfoldr1__78405858(dict_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Unfoldable1_unfoldr1__2387656390(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Lazy_Types_unfoldable1NonEmpty()).V0), __eta0_0, __eta1_1)
}

func Get_Data_Unfoldable1_unfoldr1ArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Unfoldable1_Unfoldr1ArrayImpl
}
