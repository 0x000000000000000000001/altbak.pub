package Data_Unfoldable

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_unfoldr gopurs_runtime.Value
var once_unfoldr sync.Once
func Get_unfoldr() gopurs_runtime.Value {
	once_unfoldr.Do(func() {
		cache_unfoldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr
}

var cache_unfoldr__gopurs_runtime_Value_1128708256 gopurs_runtime.Value
var once_unfoldr__gopurs_runtime_Value_1128708256 sync.Once
func Get_unfoldr__gopurs_runtime_Value_1128708256() gopurs_runtime.Value {
	once_unfoldr__gopurs_runtime_Value_1128708256.Do(func() {
		cache_unfoldr__gopurs_runtime_Value_1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__gopurs_runtime_Value_1128708256(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__gopurs_runtime_Value_1128708256
}

var cache_unfoldableMaybe gopurs_runtime.Value
var once_unfoldableMaybe sync.Once
func Get_unfoldableMaybe() gopurs_runtime.Value {
	once_unfoldableMaybe.Do(func() {
		cache_unfoldableMaybe = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Maybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_fst(), gopurs_runtime.Apply(f_0, b_1))))}
})
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
}), gopurs_runtime.Apply4(Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldableArray
}

var cache_unfoldableArray__ptrConstructor_Unfoldable_gopurs_runtime_Value__644327338 gopurs_runtime.Value
var once_unfoldableArray__ptrConstructor_Unfoldable_gopurs_runtime_Value__644327338 sync.Once
func Get_unfoldableArray__ptrConstructor_Unfoldable_gopurs_runtime_Value__644327338() gopurs_runtime.Value {
	once_unfoldableArray__ptrConstructor_Unfoldable_gopurs_runtime_Value__644327338.Do(func() {
		cache_unfoldableArray__ptrConstructor_Unfoldable_gopurs_runtime_Value__644327338 = gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(&Constructor_Unfoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Array()
}), gopurs_runtime.Apply4(Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd())})}
	})
	return cache_unfoldableArray__ptrConstructor_Unfoldable_gopurs_runtime_Value__644327338
}

var cache_replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		cache_replicate = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_replicate
}

var cache_replicate__gopurs_runtime_Value_2136688459 gopurs_runtime.Value
var once_replicate__gopurs_runtime_Value_2136688459 sync.Once
func Get_replicate__gopurs_runtime_Value_2136688459() gopurs_runtime.Value {
	once_replicate__gopurs_runtime_Value_2136688459.Do(func() {
		cache_replicate__gopurs_runtime_Value_2136688459 = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate__gopurs_runtime_Value_2136688459(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_replicate__gopurs_runtime_Value_2136688459
}

var cache_replicateA gopurs_runtime.Value
var once_replicateA sync.Once
func Get_replicateA() gopurs_runtime.Value {
	once_replicateA.Do(func() {
		cache_replicateA = gopurs_runtime.Func5(func(dictApplicative_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value, n_3_box gopurs_runtime.Value, m_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicateA(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_2_box), n_3_box.IntVal, m_4_box)
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
return Call_fromMaybe(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_fromMaybe
}

type Constructor_Unfoldable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2670894170] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Unfoldable[gopurs_runtime.Value])(ptr)
		switch key {
		case "Unfoldable10": return c.V0
		case "unfoldr": return c.V1
		default: panic("Key not found in dictionary Constructor_Unfoldable: " + key)
		}
	}
}


func Call_unfoldr(dict_0_loop *Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__gopurs_runtime_Value_1128708256(dict_0_loop *Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_replicate(dictUnfoldable_0_loop *Constructor_Unfoldable[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictUnfoldable_0.V1, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (i_3.IntVal) > (0) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
if (__t1.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](__t0))}
}), gopurs_runtime.Int(n_1))
}

func Call_replicate__gopurs_runtime_Value_2136688459(dictUnfoldable_0_loop *Constructor_Unfoldable[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictUnfoldable_0.V1, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (i_3.IntVal) > (0) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
if (__t1.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](__t0))}
}), gopurs_runtime.Int(n_1))
}

func Call_replicateA(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictUnfoldable_1_loop *Constructor_Unfoldable[gopurs_runtime.Value], dictTraversable_2_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value], n_3_loop int64, m_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var dictUnfoldable_1 *Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_1_loop
_ = dictUnfoldable_1
var dictTraversable_2 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_2_loop
_ = dictTraversable_2
var n_3 int64 = n_3_loop
_ = n_3
var m_4 gopurs_runtime.Value = m_4_loop
_ = m_4
return gopurs_runtime.Apply2(dictTraversable_2.V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Apply2(dictUnfoldable_1.V1, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (i_5.IntVal) > (0) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
if (__t1.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, m_4, gopurs_runtime.Int((i_5.IntVal) - (1))})}})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](__t0))}
}), gopurs_runtime.Int(n_3)))
}

func Call_none(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}), pkg_Data_Unit.Get_unit())
}

func Call_fromMaybe(dictUnfoldable_0_loop *Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}
}), b_1)))}
}))
}

func Get_unfoldrArrayImpl() gopurs_runtime.Value {
	return _Gopurs_UnfoldrArrayImpl
}
