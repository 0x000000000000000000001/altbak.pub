package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Unfoldable_Unfoldable_dollarDict gopurs_runtime.Value
var once_Data_Unfoldable_Unfoldable_dollarDict sync.Once
func Get_Data_Unfoldable_Unfoldable_dollarDict() gopurs_runtime.Value {
	once_Data_Unfoldable_Unfoldable_dollarDict.Do(func() {
		cache_Data_Unfoldable_Unfoldable_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_Unfoldable_dollarDict(x_0_box)
})
	})
	return cache_Data_Unfoldable_Unfoldable_dollarDict
}

var cache_Data_Unfoldable_unfoldr gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr sync.Once
func Get_Data_Unfoldable_unfoldr() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr.Do(func() {
		cache_Data_Unfoldable_unfoldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr
}

var cache_Data_Unfoldable_unfoldableMaybe gopurs_runtime.Value
var once_Data_Unfoldable_unfoldableMaybe sync.Once
func Get_Data_Unfoldable_unfoldableMaybe() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldableMaybe.Do(func() {
		cache_Data_Unfoldable_unfoldableMaybe = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unfoldable1_unfoldable1Maybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Tuple_fst(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, b_1)))})))}
})
}))
	})
	return cache_Data_Unfoldable_unfoldableMaybe
}

var cache_Data_Unfoldable_unfoldableArray gopurs_runtime.Value
var once_Data_Unfoldable_unfoldableArray sync.Once
func Get_Data_Unfoldable_unfoldableArray() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldableArray.Do(func() {
		cache_Data_Unfoldable_unfoldableArray = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unfoldable1_unfoldable1Array()
}), gopurs_runtime.Apply4(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd()))
	})
	return cache_Data_Unfoldable_unfoldableArray
}

var cache_Data_Unfoldable_replicate gopurs_runtime.Value
var once_Data_Unfoldable_replicate sync.Once
func Get_Data_Unfoldable_replicate() gopurs_runtime.Value {
	once_Data_Unfoldable_replicate.Do(func() {
		cache_Data_Unfoldable_replicate = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_replicate(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_Data_Unfoldable_replicate
}

var cache_Data_Unfoldable_replicateA gopurs_runtime.Value
var once_Data_Unfoldable_replicateA sync.Once
func Get_Data_Unfoldable_replicateA() gopurs_runtime.Value {
	once_Data_Unfoldable_replicateA.Do(func() {
		cache_Data_Unfoldable_replicateA = gopurs_runtime.Func5(func(dictApplicative_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value, dictTraversable_2_box gopurs_runtime.Value, n_3_box gopurs_runtime.Value, m_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_replicateA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_2_box), n_3_box.IntVal, m_4_box)
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
return Call_Data_Unfoldable_fromMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_Unfoldable_fromMaybe
}

var cache_Data_Unfoldable_replicate__2136688459 gopurs_runtime.Value
var once_Data_Unfoldable_replicate__2136688459 sync.Once
func Get_Data_Unfoldable_replicate__2136688459() gopurs_runtime.Value {
	once_Data_Unfoldable_replicate__2136688459.Do(func() {
		cache_Data_Unfoldable_replicate__2136688459 = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_replicate__2136688459(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_Data_Unfoldable_replicate__2136688459
}

var cache_Data_Unfoldable_replicate__997517451 gopurs_runtime.Value
var once_Data_Unfoldable_replicate__997517451 sync.Once
func Get_Data_Unfoldable_replicate__997517451() gopurs_runtime.Value {
	once_Data_Unfoldable_replicate__997517451.Do(func() {
		cache_Data_Unfoldable_replicate__997517451 = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_replicate__997517451(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_Data_Unfoldable_replicate__997517451
}

var cache_Data_Unfoldable_unfoldableArray__644327338 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldableArray__644327338 sync.Once
func Get_Data_Unfoldable_unfoldableArray__644327338() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldableArray__644327338.Do(func() {
		cache_Data_Unfoldable_unfoldableArray__644327338 = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unfoldable1_unfoldable1Array()
}), gopurs_runtime.Apply4(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd()))
	})
	return cache_Data_Unfoldable_unfoldableArray__644327338
}

var cache_Data_Unfoldable_unfoldr__2235715281 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__2235715281 sync.Once
func Get_Data_Unfoldable_unfoldr__2235715281() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__2235715281.Do(func() {
		cache_Data_Unfoldable_unfoldr__2235715281 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__2235715281(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__2235715281
}

var cache_Data_Unfoldable_unfoldr__3341150410 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__3341150410 sync.Once
func Get_Data_Unfoldable_unfoldr__3341150410() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__3341150410.Do(func() {
		cache_Data_Unfoldable_unfoldr__3341150410 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__3341150410(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__3341150410
}

var cache_Data_Unfoldable_unfoldr__1426313893 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__1426313893 sync.Once
func Get_Data_Unfoldable_unfoldr__1426313893() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__1426313893.Do(func() {
		cache_Data_Unfoldable_unfoldr__1426313893 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__1426313893(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__1426313893
}

var cache_Data_Unfoldable_unfoldr__1786322085 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__1786322085 sync.Once
func Get_Data_Unfoldable_unfoldr__1786322085() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__1786322085.Do(func() {
		cache_Data_Unfoldable_unfoldr__1786322085 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__1786322085(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__1786322085
}

var cache_Data_Unfoldable_unfoldr__2777844709 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__2777844709 sync.Once
func Get_Data_Unfoldable_unfoldr__2777844709() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__2777844709.Do(func() {
		cache_Data_Unfoldable_unfoldr__2777844709 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__2777844709(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__2777844709
}

var cache_Data_Unfoldable_unfoldr__1842498883 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__1842498883 sync.Once
func Get_Data_Unfoldable_unfoldr__1842498883() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__1842498883.Do(func() {
		cache_Data_Unfoldable_unfoldr__1842498883 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__1842498883(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Unfoldable_unfoldr__1842498883
}

var cache_Data_Unfoldable_unfoldr__2604413936 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__2604413936 sync.Once
func Get_Data_Unfoldable_unfoldr__2604413936() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__2604413936.Do(func() {
		cache_Data_Unfoldable_unfoldr__2604413936 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__2604413936(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__2604413936
}

var cache_Data_Unfoldable_unfoldr__3587480224 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__3587480224 sync.Once
func Get_Data_Unfoldable_unfoldr__3587480224() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__3587480224.Do(func() {
		cache_Data_Unfoldable_unfoldr__3587480224 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__3587480224(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__3587480224
}

var cache_Data_Unfoldable_unfoldr__1128708256 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__1128708256 sync.Once
func Get_Data_Unfoldable_unfoldr__1128708256() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__1128708256.Do(func() {
		cache_Data_Unfoldable_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__1128708256
}

var cache_Data_Unfoldable_unfoldr__327534368 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__327534368 sync.Once
func Get_Data_Unfoldable_unfoldr__327534368() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__327534368.Do(func() {
		cache_Data_Unfoldable_unfoldr__327534368 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__327534368(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__327534368
}

var cache_Data_Unfoldable_unfoldr__3736536416 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__3736536416 sync.Once
func Get_Data_Unfoldable_unfoldr__3736536416() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__3736536416.Do(func() {
		cache_Data_Unfoldable_unfoldr__3736536416 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__3736536416(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__3736536416
}

var cache_Data_Unfoldable_unfoldr__3647947554 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__3647947554 sync.Once
func Get_Data_Unfoldable_unfoldr__3647947554() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__3647947554.Do(func() {
		cache_Data_Unfoldable_unfoldr__3647947554 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__3647947554(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__3647947554
}

var cache_Data_Unfoldable_unfoldr__457386988 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__457386988 sync.Once
func Get_Data_Unfoldable_unfoldr__457386988() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__457386988.Do(func() {
		cache_Data_Unfoldable_unfoldr__457386988 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__457386988(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__457386988
}

var cache_Data_Unfoldable_unfoldr__1779806517 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__1779806517 sync.Once
func Get_Data_Unfoldable_unfoldr__1779806517() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__1779806517.Do(func() {
		cache_Data_Unfoldable_unfoldr__1779806517 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__1779806517(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__1779806517
}

var cache_Data_Unfoldable_unfoldr__3827943605 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__3827943605 sync.Once
func Get_Data_Unfoldable_unfoldr__3827943605() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__3827943605.Do(func() {
		cache_Data_Unfoldable_unfoldr__3827943605 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__3827943605(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__3827943605
}

var cache_Data_Unfoldable_unfoldr__1519733018 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__1519733018 sync.Once
func Get_Data_Unfoldable_unfoldr__1519733018() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__1519733018.Do(func() {
		cache_Data_Unfoldable_unfoldr__1519733018 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__1519733018(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__1519733018
}

var cache_Data_Unfoldable_unfoldr__3132297377 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__3132297377 sync.Once
func Get_Data_Unfoldable_unfoldr__3132297377() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__3132297377.Do(func() {
		cache_Data_Unfoldable_unfoldr__3132297377 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__3132297377(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__3132297377
}

var cache_Data_Unfoldable_unfoldr__193332035 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__193332035 sync.Once
func Get_Data_Unfoldable_unfoldr__193332035() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__193332035.Do(func() {
		cache_Data_Unfoldable_unfoldr__193332035 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__193332035(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Unfoldable_unfoldr__193332035
}

var cache_Data_Unfoldable_unfoldr__553159458 gopurs_runtime.Value
var once_Data_Unfoldable_unfoldr__553159458 sync.Once
func Get_Data_Unfoldable_unfoldr__553159458() gopurs_runtime.Value {
	once_Data_Unfoldable_unfoldr__553159458.Do(func() {
		cache_Data_Unfoldable_unfoldr__553159458 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Unfoldable_unfoldr__553159458(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dict_0_box))
})
	})
	return cache_Data_Unfoldable_unfoldr__553159458
}

type Constructor_Data_Unfoldable_Unfoldable struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2670894170] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Unfoldable_Unfoldable)(ptr)
		_ = c
		switch key {
		case "Unfoldable10": return gopurs_runtime.Box(c.V0)
		case "unfoldr": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Unfoldable_Unfoldable: " + key)
		}
	}
}


func Call_Data_Unfoldable_Unfoldable_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Unfoldable_unfoldr(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_replicate(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))}
}), gopurs_runtime.Int(n_1))
}

func Call_Data_Unfoldable_replicateA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, dictUnfoldable_1_loop *Constructor_Data_Unfoldable_Unfoldable, dictTraversable_2_loop *Constructor_Data_Traversable_Traversable, n_3_loop int64, m_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var dictUnfoldable_1 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_1_loop
_ = dictUnfoldable_1
var dictTraversable_2 *Constructor_Data_Traversable_Traversable = dictTraversable_2_loop
_ = dictTraversable_2
var n_3 int64 = n_3_loop
_ = n_3
var m_4 gopurs_runtime.Value = m_4_loop
_ = m_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_2.V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_1.V1), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, m_4, gopurs_runtime.Int((i_5.IntVal) - (1))})}})}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))}
}), gopurs_runtime.Int(n_3)))
}

func Call_Data_Unfoldable_none(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}), Get_Data_Unit_unit())
}

func Call_Data_Unfoldable_fromMaybe(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](b_1))})))}
}))
}

func Call_Data_Unfoldable_replicate__2136688459(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))}
}), gopurs_runtime.Int(n_1))
}

func Call_Data_Unfoldable_replicate__997517451(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, v_2, gopurs_runtime.Int((i_3.IntVal) - (1))})}})}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))}
}), gopurs_runtime.Int(n_1))
}

func Call_Data_Unfoldable_unfoldr__2235715281(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__3341150410(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__1426313893(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__1786322085(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__2777844709(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__1842498883(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply6(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil) {
__t0 = (*Constructor_Data_Maybe_Just)(v_3.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), __eta0_0, __eta1_1)
}

func Call_Data_Unfoldable_unfoldr__2604413936(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__3587480224(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__1128708256(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__327534368(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__3736536416(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__3647947554(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__457386988(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__1779806517(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__3827943605(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__1519733018(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__3132297377(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Unfoldable_unfoldr__193332035(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var go__go_2_0_0 gopurs_runtime.Value
_ = go__go_2_0_0
go__go_2_0_0 = gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_6_1 -> *Constructor_Data_Maybe_Just
v1_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_3, gopurs_runtime.Int(b_4.IntVal)))
_ = v1_6_1
var __t2 gopurs_runtime.Value
{
if (v1_6_1 == nil) {
__t2 = Get_Data_List_Lazy_Types_nil()
goto end_branch_2
} else {

}
}
{
if (v1_6_1 != nil) {
__t2 = gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((*Constructor_Data_Tuple_Tuple)((v1_6_1).V0.UnsafePtr).V0))}, gopurs_runtime.Apply2(go__go_2_0_0, f_3, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)((v1_6_1).V0.UnsafePtr).V1.IntVal)))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
})
return gopurs_runtime.Apply2(go__go_2_0_0, __eta0_0, __eta1_1)
}

func Call_Data_Unfoldable_unfoldr__553159458(dict_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Unfoldable_Unfoldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Get_Data_Unfoldable_unfoldrArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Unfoldable_UnfoldrArrayImpl
}
