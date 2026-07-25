package Data_Unfoldable1

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		cache_fromJust = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136) {
__t0 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return cache_fromJust
}

var cache_unfoldr1 gopurs_runtime.Value
var once_unfoldr1 sync.Once
func Get_unfoldr1() gopurs_runtime.Value {
	once_unfoldr1.Do(func() {
		cache_unfoldr1 = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "unfoldr1")
}()
})
	})
	return cache_unfoldr1
}

var cache_unfoldable1Maybe gopurs_runtime.Value
var once_unfoldable1Maybe sync.Once
func Get_unfoldable1Maybe() gopurs_runtime.Value {
	once_unfoldable1Maybe.Do(func() {
		cache_unfoldable1Maybe = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_0, b_1).UnsafePtr).V0})}
}))
	})
	return cache_unfoldable1Maybe
}

var cache_unfoldable1Array gopurs_runtime.Value
var once_unfoldable1Array sync.Once
func Get_unfoldable1Array() gopurs_runtime.Value {
	once_unfoldable1Array.Do(func() {
		cache_unfoldable1Array = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), Get_fromJust(), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldable1Array
}

var cache_replicate1 gopurs_runtime.Value
var once_replicate1 sync.Once
func Get_replicate1() gopurs_runtime.Value {
	once_replicate1.Do(func() {
		cache_replicate1 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate1(dictUnfoldable1_0_box, n_1_box, v_2_box)
})
	})
	return cache_replicate1
}

var cache_replicate1A gopurs_runtime.Value
var once_replicate1A sync.Once
func Get_replicate1A() gopurs_runtime.Value {
	once_replicate1A.Do(func() {
		cache_replicate1A = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value, dictTraversable1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate1A(dictApply_0_box, dictUnfoldable1_1_box, dictTraversable1_2_box)
})
	})
	return cache_replicate1A
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(dictUnfoldable1_0_box, v_1_box)
})
	})
	return cache_singleton
}

var cache_range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		cache_range_ = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_range_(dictUnfoldable1_0_box, start_1_box, end_2_box)
})
	})
	return cache_range_
}

var cache_iterateN gopurs_runtime.Value
var once_iterateN sync.Once
func Get_iterateN() gopurs_runtime.Value {
	once_iterateN.Do(func() {
		cache_iterateN = gopurs_runtime.Func4(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterateN(dictUnfoldable1_0_box, n_1_box, f_2_box, s_3_box)
})
	})
	return cache_iterateN
}

func Call_replicate1(dictUnfoldable1_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 gopurs_runtime.Value = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (i_3.IntVal) <= (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{v_2, gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int((i_3.IntVal) - (1))})}})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Int((n_1.IntVal) - (1)))
}

func Call_replicate1A(dictApply_0_loop gopurs_runtime.Value, dictUnfoldable1_1_loop gopurs_runtime.Value, dictTraversable1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var dictUnfoldable1_1 gopurs_runtime.Value = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
var dictTraversable1_2 gopurs_runtime.Value = dictTraversable1_2_loop
_ = dictTraversable1_2
sequence1_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_2, "sequence1"), dictApply_0)
_ = sequence1_3_0
return gopurs_runtime.Func2(func(n_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_3_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_1, "unfoldr1"), gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (i_6.IntVal) <= (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{m_5, gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{m_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int((i_6.IntVal) - (1))})}})}
}
end_branch_1:
return __t1
}), gopurs_runtime.Int((n_4.IntVal) - (1))))
})
}

func Call_singleton(dictUnfoldable1_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 gopurs_runtime.Value = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (i_2.IntVal) <= (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{v_1, gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int((i_2.IntVal) - (1))})}})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Int(0))
}

func Call_range_(dictUnfoldable1_0_loop gopurs_runtime.Value, start_1_loop gopurs_runtime.Value, end_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 gopurs_runtime.Value = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var start_1 gopurs_runtime.Value = start_1_loop
_ = start_1
var end_2 gopurs_runtime.Value = end_2_loop
_ = end_2
var __t1 gopurs_runtime.Value
{
if (end_2.IntVal) >= (start_1.IntVal) {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Int(-1)
}
end_branch_1:
__local_var_3_0 := __t1
_ = __local_var_3_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
i_prime_5_2 := (i_4.IntVal) + (__local_var_3_0.IntVal)
_ = i_prime_5_2
var __t3 gopurs_runtime.Value
{
if (i_4.IntVal) == (end_2.IntVal) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(i_prime_5_2)})}
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{i_4, __t3})}
}), start_1)
}

func Call_iterateN(dictUnfoldable1_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 gopurs_runtime.Value = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1.IntVal) > (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), gopurs_runtime.Int(((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1.IntVal) - (1))})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, __t0})}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{s_3, gopurs_runtime.Int((n_1.IntVal) - (1))})})
}

func Get_unfoldr1ArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Unfoldr1ArrayImpl
}
