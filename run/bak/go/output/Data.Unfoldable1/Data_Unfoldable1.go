package Data_Unfoldable1

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup_Traversable "gopurs/output/Data.Semigroup.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_unfoldr1 gopurs_runtime.Value
var once_unfoldr1 sync.Once
func Get_unfoldr1() gopurs_runtime.Value {
	once_unfoldr1.Do(func() {
		cache_unfoldr1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr1(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr1
}

var cache_unfoldr1__gopurs_runtime_Value_2402610528 gopurs_runtime.Value
var once_unfoldr1__gopurs_runtime_Value_2402610528 sync.Once
func Get_unfoldr1__gopurs_runtime_Value_2402610528() gopurs_runtime.Value {
	once_unfoldr1__gopurs_runtime_Value_2402610528.Do(func() {
		cache_unfoldr1__gopurs_runtime_Value_2402610528 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr1__gopurs_runtime_Value_2402610528(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr1__gopurs_runtime_Value_2402610528
}

var cache_unfoldable1Maybe gopurs_runtime.Value
var once_unfoldable1Maybe sync.Once
func Get_unfoldable1Maybe() gopurs_runtime.Value {
	once_unfoldable1Maybe.Do(func() {
		cache_unfoldable1Maybe = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_0, b_1).UnsafePtr).V0})}))}
})
}))
	})
	return cache_unfoldable1Maybe
}

var cache_unfoldable1Maybe__gopurs_runtime_Value_3214541052 gopurs_runtime.Value
var once_unfoldable1Maybe__gopurs_runtime_Value_3214541052 sync.Once
func Get_unfoldable1Maybe__gopurs_runtime_Value_3214541052() gopurs_runtime.Value {
	once_unfoldable1Maybe__gopurs_runtime_Value_3214541052.Do(func() {
		cache_unfoldable1Maybe__gopurs_runtime_Value_3214541052 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_0, b_1).UnsafePtr).V0})}))}
})
}))
	})
	return cache_unfoldable1Maybe__gopurs_runtime_Value_3214541052
}

var cache_unfoldable1Array gopurs_runtime.Value
var once_unfoldable1Array sync.Once
func Get_unfoldable1Array() gopurs_runtime.Value {
	once_unfoldable1Array.Do(func() {
		cache_unfoldable1Array = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_unfoldable1Array
}

var cache_unfoldable1Array__ptrConstructor_Unfoldable1_gopurs_runtime_Value__2415700810 gopurs_runtime.Value
var once_unfoldable1Array__ptrConstructor_Unfoldable1_gopurs_runtime_Value__2415700810 sync.Once
func Get_unfoldable1Array__ptrConstructor_Unfoldable1_gopurs_runtime_Value__2415700810() gopurs_runtime.Value {
	once_unfoldable1Array__ptrConstructor_Unfoldable1_gopurs_runtime_Value__2415700810.Do(func() {
		cache_unfoldable1Array__ptrConstructor_Unfoldable1_gopurs_runtime_Value__2415700810 = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Apply4(Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_unfoldable1Array__ptrConstructor_Unfoldable1_gopurs_runtime_Value__2415700810
}

var cache_unfoldable1Array__gopurs_runtime_Value_4196906331 gopurs_runtime.Value
var once_unfoldable1Array__gopurs_runtime_Value_4196906331 sync.Once
func Get_unfoldable1Array__gopurs_runtime_Value_4196906331() gopurs_runtime.Value {
	once_unfoldable1Array__gopurs_runtime_Value_4196906331.Do(func() {
		cache_unfoldable1Array__gopurs_runtime_Value_4196906331 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_unfoldable1Array__gopurs_runtime_Value_4196906331
}

var cache_replicate1 gopurs_runtime.Value
var once_replicate1 sync.Once
func Get_replicate1() gopurs_runtime.Value {
	once_replicate1.Do(func() {
		cache_replicate1 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate1(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_replicate1
}

var cache_replicate1__gopurs_runtime_Value_3169098027 gopurs_runtime.Value
var once_replicate1__gopurs_runtime_Value_3169098027 sync.Once
func Get_replicate1__gopurs_runtime_Value_3169098027() gopurs_runtime.Value {
	once_replicate1__gopurs_runtime_Value_3169098027.Do(func() {
		cache_replicate1__gopurs_runtime_Value_3169098027 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate1__gopurs_runtime_Value_3169098027(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_replicate1__gopurs_runtime_Value_3169098027
}

var cache_replicate1A gopurs_runtime.Value
var once_replicate1A sync.Once
func Get_replicate1A() gopurs_runtime.Value {
	once_replicate1A.Do(func() {
		cache_replicate1A = gopurs_runtime.Func5(func(dictApply_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value, dictTraversable1_2_box gopurs_runtime.Value, n_3_box gopurs_runtime.Value, m_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate1A(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box), gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value]](dictTraversable1_2_box), n_3_box.IntVal, m_4_box)
})
	})
	return cache_replicate1A
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), v_1_box)
})
	})
	return cache_singleton
}

var cache_singleton__gopurs_runtime_Value_1620623815 gopurs_runtime.Value
var once_singleton__gopurs_runtime_Value_1620623815 sync.Once
func Get_singleton__gopurs_runtime_Value_1620623815() gopurs_runtime.Value {
	once_singleton__gopurs_runtime_Value_1620623815.Do(func() {
		cache_singleton__gopurs_runtime_Value_1620623815 = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton__gopurs_runtime_Value_1620623815(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), v_1_box)
})
	})
	return cache_singleton__gopurs_runtime_Value_1620623815
}

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__range(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), start_1_box.IntVal, end_2_box.IntVal)
})
	})
	return cache_go__range
}

var cache_iterateN gopurs_runtime.Value
var once_iterateN sync.Once
func Get_iterateN() gopurs_runtime.Value {
	once_iterateN.Do(func() {
		cache_iterateN = gopurs_runtime.Func4(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterateN(gopurs_runtime.CoerceToStruct[Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), n_1_box.IntVal, f_2_box, s_3_box)
})
	})
	return cache_iterateN
}

type Constructor_Unfoldable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3553002490] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Unfoldable1[gopurs_runtime.Value])(ptr)
		switch key {
		case "unfoldr1": return c.V0
		default: panic("Key not found in dictionary Constructor_Unfoldable1: " + key)
		}
	}
}


func Call_unfoldr1(dict_0_loop *Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unfoldr1__gopurs_runtime_Value_2402610528(dict_0_loop *Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_replicate1(dictUnfoldable1_0_loop *Constructor_Unfoldable1[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](__t0))}
}), gopurs_runtime.Int((n_1) - (1)))
}

func Call_replicate1__gopurs_runtime_Value_3169098027(dictUnfoldable1_0_loop *Constructor_Unfoldable1[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](__t0))}
}), gopurs_runtime.Int((n_1) - (1)))
}

func Call_replicate1A(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value], dictUnfoldable1_1_loop *Constructor_Unfoldable1[gopurs_runtime.Value], dictTraversable1_2_loop *pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value], n_3_loop int64, m_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
var dictUnfoldable1_1 *Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
var dictTraversable1_2 *pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value] = dictTraversable1_2_loop
_ = dictTraversable1_2
var n_3 int64 = n_3_loop
_ = n_3
var m_4 gopurs_runtime.Value = m_4_loop
_ = m_4
return gopurs_runtime.Apply2(dictTraversable1_2.V2, gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_0)}, gopurs_runtime.Apply2(dictUnfoldable1_1.V0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, m_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, m_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_5.IntVal) - (1))})}})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](__t0))}
}), gopurs_runtime.Int((n_3) - (1))))
}

func Call_singleton(dictUnfoldable1_0_loop *Constructor_Unfoldable1[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (i_2.IntVal) > (0) {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_2.IntVal) - (1))})}})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](__t0))}
}), gopurs_runtime.Int(0))
}

func Call_singleton__gopurs_runtime_Value_1620623815(dictUnfoldable1_0_loop *Constructor_Unfoldable1[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (i_2.IntVal) > (0) {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_2.IntVal) - (1))})}})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[int64]]](__t0))}
}), gopurs_runtime.Int(0))
}

func Call_go__range(dictUnfoldable1_0_loop *Constructor_Unfoldable1[gopurs_runtime.Value], start_1_loop int64, end_2_loop int64) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var start_1 int64 = start_1_loop
_ = start_1
var end_2 int64 = end_2_loop
_ = end_2
var __t1 gopurs_runtime.Value
{
var __t2 gopurs_runtime.Value
{
if (end_2) < (start_1) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(true)
}
end_branch_2:
if (__t2.IntVal) != (0) {
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
return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
i_prime_5_3 := (i_4.IntVal) + (__local_var_3_0.IntVal)
_ = i_prime_5_3
var __t4 gopurs_runtime.Value
{
if (i_4.IntVal) == (end_2) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(i_prime_5_3)})}))}
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, i_4, __t4})}))}
}), gopurs_runtime.Int(start_1))
}

func Call_iterateN(dictUnfoldable1_0_loop *Constructor_Unfoldable1[gopurs_runtime.Value], n_1_loop int64, f_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if ((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal) > (0) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
if (__t1.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Int(((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal) - (1))})}})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, __t0})}))}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_3, gopurs_runtime.Int((n_1) - (1))})})
}

func Get_unfoldr1ArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Unfoldr1ArrayImpl
}
