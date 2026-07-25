package Data_Unfoldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Unit "gopurs/output/Data.Unit"
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

var cache_unfoldr gopurs_runtime.Value
var once_unfoldr sync.Once
func Get_unfoldr() gopurs_runtime.Value {
	once_unfoldr.Do(func() {
		cache_unfoldr = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "unfoldr")
}()
})
	})
	return cache_unfoldr
}

var cache_unfoldableMaybe gopurs_runtime.Value
var once_unfoldableMaybe sync.Once
func Get_unfoldableMaybe() gopurs_runtime.Value {
	once_unfoldableMaybe.Do(func() {
		cache_unfoldableMaybe = gopurs_runtime.RecordDict2("unfoldr", "Unfoldable10", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, b_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Maybe()
}))
	})
	return cache_unfoldableMaybe
}

var cache_unfoldableArray gopurs_runtime.Value
var once_unfoldableArray sync.Once
func Get_unfoldableArray() gopurs_runtime.Value {
	once_unfoldableArray.Do(func() {
		cache_unfoldableArray = gopurs_runtime.RecordDict2("unfoldr", "Unfoldable10", gopurs_runtime.Apply4(Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), Get_fromJust(), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Array()
}))
	})
	return cache_unfoldableArray
}

var cache_replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		cache_replicate = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate(dictUnfoldable_0_box, n_1_box, v_2_box)
})
	})
	return cache_replicate
}

var cache_replicateA gopurs_runtime.Value
var once_replicateA sync.Once
func Get_replicateA() gopurs_runtime.Value {
	once_replicateA.Do(func() {
		cache_replicateA = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicateA(dictApplicative_0_box, dictUnfoldable_1_box, dictTraversable_2_box)
})
	})
	return cache_replicateA
}

var cache_none gopurs_runtime.Value
var once_none sync.Once
func Get_none() gopurs_runtime.Value {
	once_none.Do(func() {
		cache_none = gopurs_runtime.Func(func(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}), pkg_Data_Unit.Get_unit())
}()
})
	})
	return cache_none
}

var cache_fromMaybe gopurs_runtime.Value
var once_fromMaybe sync.Once
func Get_fromMaybe() gopurs_runtime.Value {
	once_fromMaybe.Do(func() {
		cache_fromMaybe = gopurs_runtime.Func(func(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (b_1.Type == 9 && b_1.IntVal == 930809136) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(b_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}))
}()
})
	})
	return cache_fromMaybe
}

func Call_replicate(dictUnfoldable_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (i_3.IntVal) <= (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{v_2, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}
}
end_branch_0:
return __t0
}), n_1)
}

func Call_replicateA(dictApplicative_0_loop gopurs_runtime.Value, dictUnfoldable_1_loop gopurs_runtime.Value, dictTraversable_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var dictUnfoldable_1 gopurs_runtime.Value = dictUnfoldable_1_loop
_ = dictUnfoldable_1
var dictTraversable_2 gopurs_runtime.Value = dictTraversable_2_loop
_ = dictTraversable_2
sequence_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_2, "sequence"), dictApplicative_0)
_ = sequence_3_0
return gopurs_runtime.Func2(func(n_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence_3_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_1, "unfoldr"), gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (i_6.IntVal) <= (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{m_5, gopurs_runtime.Int((i_6.IntVal) - (1))})}})}
}
end_branch_1:
return __t1
}), n_4))
})
}

func Get_unfoldrArrayImpl() gopurs_runtime.Value {
	return _Gopurs_UnfoldrArrayImpl
}
