package Data_Unfoldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	unsafe "unsafe"
)

var cache_fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		cache_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0_box.UnsafePtr))
})
	})
	return cache_fromJust
}

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415)) != (true))
})
}()
	})
	return cache_lessThanOrEq
}

var cache_unfoldr gopurs_runtime.Value
var once_unfoldr sync.Once
func Get_unfoldr() gopurs_runtime.Value {
	once_unfoldr.Do(func() {
		cache_unfoldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr(dict_0_box)
})
	})
	return cache_unfoldr
}

var cache_unfoldr__gopurs_runtime_Value_956138623 gopurs_runtime.Value
var once_unfoldr__gopurs_runtime_Value_956138623 sync.Once
func Get_unfoldr__gopurs_runtime_Value_956138623() gopurs_runtime.Value {
	once_unfoldr__gopurs_runtime_Value_956138623.Do(func() {
		cache_unfoldr__gopurs_runtime_Value_956138623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__gopurs_runtime_Value_956138623(dict_0_box)
})
	})
	return cache_unfoldr__gopurs_runtime_Value_956138623
}

var cache_unfoldableMaybe gopurs_runtime.Value
var once_unfoldableMaybe sync.Once
func Get_unfoldableMaybe() gopurs_runtime.Value {
	once_unfoldableMaybe.Do(func() {
		cache_unfoldableMaybe = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Maybe()
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_fst(), gopurs_runtime.Apply(f_0, b_1))
}))
	})
	return cache_unfoldableMaybe
}

var cache_unfoldableArray gopurs_runtime.Value
var once_unfoldableArray sync.Once
func Get_unfoldableArray() gopurs_runtime.Value {
	once_unfoldableArray.Do(func() {
		cache_unfoldableArray = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Array()
}), gopurs_runtime.Apply4(Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), Get_fromJust(), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldableArray
}

var cache_replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		cache_replicate = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate(dictUnfoldable_0_box, n_1_box.IntVal, v_2_box)
})
	})
	return cache_replicate
}

var cache_replicate__gopurs_runtime_Value_1763173703 gopurs_runtime.Value
var once_replicate__gopurs_runtime_Value_1763173703 sync.Once
func Get_replicate__gopurs_runtime_Value_1763173703() gopurs_runtime.Value {
	once_replicate__gopurs_runtime_Value_1763173703.Do(func() {
		cache_replicate__gopurs_runtime_Value_1763173703 = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate__gopurs_runtime_Value_1763173703(dictUnfoldable_0_box, n_1_box.IntVal, v_2_box)
})
	})
	return cache_replicate__gopurs_runtime_Value_1763173703
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
		cache_none = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_none(dictUnfoldable_0_box)
})
	})
	return cache_none
}

var cache_fromMaybe gopurs_runtime.Value
var once_fromMaybe sync.Once
func Get_fromMaybe() gopurs_runtime.Value {
	once_fromMaybe.Do(func() {
		cache_fromMaybe = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe(dictUnfoldable_0_box)
})
	})
	return cache_fromMaybe
}

func Call_fromJust(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_unfoldr(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "unfoldr")
}

func Call_unfoldr__gopurs_runtime_Value_956138623(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "unfoldr")
}

func Call_replicate(dictUnfoldable_0_loop gopurs_runtime.Value, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), i_3, gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Int(n_1))
}

func Call_replicate__gopurs_runtime_Value_1763173703(dictUnfoldable_0_loop gopurs_runtime.Value, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), i_3, gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Int(n_1))
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
return gopurs_runtime.Apply(sequence_3_0, Call_replicate(dictUnfoldable_1, n_4.IntVal, m_5))
})
}

func Call_none(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}), pkg_Data_Unit.Get_unit())
}

func Call_fromMaybe(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}
}), b_1)
}))
}

func Get_unfoldrArrayImpl() gopurs_runtime.Value {
	return _Gopurs_UnfoldrArrayImpl
}
