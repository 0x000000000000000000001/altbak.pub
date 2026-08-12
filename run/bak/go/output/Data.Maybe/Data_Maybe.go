package Data_Maybe

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_Nothing gopurs_runtime.Value
var once_Nothing sync.Once
func Get_Nothing() gopurs_runtime.Value {
	once_Nothing.Do(func() {
		cache_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
	})
	return cache_Nothing
}

var cache_Just gopurs_runtime.Value
var once_Just sync.Once
func Get_Just() gopurs_runtime.Value {
	once_Just.Do(func() {
		cache_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Just
}

var cache_showMaybe gopurs_runtime.Value
var once_showMaybe sync.Once
func Get_showMaybe() gopurs_runtime.Value {
	once_showMaybe.Do(func() {
		cache_showMaybe = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showMaybe(dictShow_0_box)
})
	})
	return cache_showMaybe
}

var cache_semigroupMaybe gopurs_runtime.Value
var once_semigroupMaybe sync.Once
func Get_semigroupMaybe() gopurs_runtime.Value {
	once_semigroupMaybe.Do(func() {
		cache_semigroupMaybe = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupMaybe(dictSemigroup_0_box)
})
	})
	return cache_semigroupMaybe
}

var cache_optional gopurs_runtime.Value
var once_optional sync.Once
func Get_optional() gopurs_runtime.Value {
	once_optional.Do(func() {
		cache_optional = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_optional(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dictAlt_0_box))
})
	})
	return cache_optional
}

var cache_monoidMaybe gopurs_runtime.Value
var once_monoidMaybe sync.Once
func Get_monoidMaybe() gopurs_runtime.Value {
	once_monoidMaybe.Do(func() {
		cache_monoidMaybe = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidMaybe(dictSemigroup_0_box)
})
	})
	return cache_monoidMaybe
}

var cache_maybe_prime gopurs_runtime.Value
var once_maybe_prime sync.Once
func Get_maybe_prime() gopurs_runtime.Value {
	once_maybe_prime.Do(func() {
		cache_maybe_prime = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe_prime(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe_prime
}

var cache_maybe_prime__gopurs_runtime_Value_2328206764 gopurs_runtime.Value
var once_maybe_prime__gopurs_runtime_Value_2328206764 sync.Once
func Get_maybe_prime__gopurs_runtime_Value_2328206764() gopurs_runtime.Value {
	once_maybe_prime__gopurs_runtime_Value_2328206764.Do(func() {
		cache_maybe_prime__gopurs_runtime_Value_2328206764 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe_prime__gopurs_runtime_Value_2328206764(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe_prime__gopurs_runtime_Value_2328206764
}

var cache_maybe gopurs_runtime.Value
var once_maybe sync.Once
func Get_maybe() gopurs_runtime.Value {
	once_maybe.Do(func() {
		cache_maybe = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe
}

var cache_maybe__gopurs_runtime_Value_3658316244 gopurs_runtime.Value
var once_maybe__gopurs_runtime_Value_3658316244 sync.Once
func Get_maybe__gopurs_runtime_Value_3658316244() gopurs_runtime.Value {
	once_maybe__gopurs_runtime_Value_3658316244.Do(func() {
		cache_maybe__gopurs_runtime_Value_3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__gopurs_runtime_Value_3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__gopurs_runtime_Value_3658316244
}

var cache_isNothing gopurs_runtime.Value
var once_isNothing sync.Once
func Get_isNothing() gopurs_runtime.Value {
	once_isNothing.Do(func() {
		cache_isNothing = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing
}

var cache_isNothing__gopurs_runtime_Value_2514352589 gopurs_runtime.Value
var once_isNothing__gopurs_runtime_Value_2514352589 sync.Once
func Get_isNothing__gopurs_runtime_Value_2514352589() gopurs_runtime.Value {
	once_isNothing__gopurs_runtime_Value_2514352589.Do(func() {
		cache_isNothing__gopurs_runtime_Value_2514352589 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__gopurs_runtime_Value_2514352589(gopurs_runtime.CoerceToStruct[Constructor_Just[int64]](v2_0_box)))
})
	})
	return cache_isNothing__gopurs_runtime_Value_2514352589
}

var cache_isNothing__gopurs_runtime_Value_2591355336 gopurs_runtime.Value
var once_isNothing__gopurs_runtime_Value_2591355336 sync.Once
func Get_isNothing__gopurs_runtime_Value_2591355336() gopurs_runtime.Value {
	once_isNothing__gopurs_runtime_Value_2591355336.Do(func() {
		cache_isNothing__gopurs_runtime_Value_2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__gopurs_runtime_Value_2591355336(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__gopurs_runtime_Value_2591355336
}

var cache_isNothing__gopurs_runtime_Value_4206805139 gopurs_runtime.Value
var once_isNothing__gopurs_runtime_Value_4206805139 sync.Once
func Get_isNothing__gopurs_runtime_Value_4206805139() gopurs_runtime.Value {
	once_isNothing__gopurs_runtime_Value_4206805139.Do(func() {
		cache_isNothing__gopurs_runtime_Value_4206805139 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__gopurs_runtime_Value_4206805139(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__gopurs_runtime_Value_4206805139
}

var cache_isJust gopurs_runtime.Value
var once_isJust sync.Once
func Get_isJust() gopurs_runtime.Value {
	once_isJust.Do(func() {
		cache_isJust = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isJust
}

var cache_isJust__gopurs_runtime_Value_2514352589 gopurs_runtime.Value
var once_isJust__gopurs_runtime_Value_2514352589 sync.Once
func Get_isJust__gopurs_runtime_Value_2514352589() gopurs_runtime.Value {
	once_isJust__gopurs_runtime_Value_2514352589.Do(func() {
		cache_isJust__gopurs_runtime_Value_2514352589 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__gopurs_runtime_Value_2514352589(gopurs_runtime.CoerceToStruct[Constructor_Just[int64]](v2_0_box)))
})
	})
	return cache_isJust__gopurs_runtime_Value_2514352589
}

var cache_isJust__gopurs_runtime_Value_2475527019 gopurs_runtime.Value
var once_isJust__gopurs_runtime_Value_2475527019 sync.Once
func Get_isJust__gopurs_runtime_Value_2475527019() gopurs_runtime.Value {
	once_isJust__gopurs_runtime_Value_2475527019.Do(func() {
		cache_isJust__gopurs_runtime_Value_2475527019 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__gopurs_runtime_Value_2475527019(gopurs_runtime.CoerceToStruct[Constructor_Just[string]](v2_0_box)))
})
	})
	return cache_isJust__gopurs_runtime_Value_2475527019
}

var cache_isJust__gopurs_runtime_Value_2591355336 gopurs_runtime.Value
var once_isJust__gopurs_runtime_Value_2591355336 sync.Once
func Get_isJust__gopurs_runtime_Value_2591355336() gopurs_runtime.Value {
	once_isJust__gopurs_runtime_Value_2591355336.Do(func() {
		cache_isJust__gopurs_runtime_Value_2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__gopurs_runtime_Value_2591355336(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isJust__gopurs_runtime_Value_2591355336
}

var cache_isJust__gopurs_runtime_Value_4206805139 gopurs_runtime.Value
var once_isJust__gopurs_runtime_Value_4206805139 sync.Once
func Get_isJust__gopurs_runtime_Value_4206805139() gopurs_runtime.Value {
	once_isJust__gopurs_runtime_Value_4206805139.Do(func() {
		cache_isJust__gopurs_runtime_Value_4206805139 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__gopurs_runtime_Value_4206805139(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isJust__gopurs_runtime_Value_4206805139
}

var cache_genericMaybe gopurs_runtime.Value
var once_genericMaybe sync.Once
func Get_genericMaybe() gopurs_runtime.Value {
	once_genericMaybe.Do(func() {
		cache_genericMaybe = gopurs_runtime.RecordDict2("from", "to", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3478632216) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 492034566) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t1))}
}))
	})
	return cache_genericMaybe
}

var cache_functorMaybe gopurs_runtime.Value
var once_functorMaybe sync.Once
func Get_functorMaybe() gopurs_runtime.Value {
	once_functorMaybe.Do(func() {
		cache_functorMaybe = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe
}

var cache_functorMaybe__ptrData_Functor_Constructor_Functor_ptrConstructor_Just_gopurs_runtime_Value___2569569018 gopurs_runtime.Value
var once_functorMaybe__ptrData_Functor_Constructor_Functor_ptrConstructor_Just_gopurs_runtime_Value___2569569018 sync.Once
func Get_functorMaybe__ptrData_Functor_Constructor_Functor_ptrConstructor_Just_gopurs_runtime_Value___2569569018() gopurs_runtime.Value {
	once_functorMaybe__ptrData_Functor_Constructor_Functor_ptrConstructor_Just_gopurs_runtime_Value___2569569018.Do(func() {
		cache_functorMaybe__ptrData_Functor_Constructor_Functor_ptrConstructor_Just_gopurs_runtime_Value___2569569018 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&pkg_Data_Functor.Constructor_Functor[*Constructor_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
})})}
	})
	return cache_functorMaybe__ptrData_Functor_Constructor_Functor_ptrConstructor_Just_gopurs_runtime_Value___2569569018
}

var cache_functorMaybe__gopurs_runtime_Value_2097654001 gopurs_runtime.Value
var once_functorMaybe__gopurs_runtime_Value_2097654001 sync.Once
func Get_functorMaybe__gopurs_runtime_Value_2097654001() gopurs_runtime.Value {
	once_functorMaybe__gopurs_runtime_Value_2097654001.Do(func() {
		cache_functorMaybe__gopurs_runtime_Value_2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__gopurs_runtime_Value_2097654001
}

var cache_invariantMaybe gopurs_runtime.Value
var once_invariantMaybe sync.Once
func Get_invariantMaybe() gopurs_runtime.Value {
	once_invariantMaybe.Do(func() {
		cache_invariantMaybe = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorMaybe(), "map"), f_0)
})
}))
	})
	return cache_invariantMaybe
}

var cache_invariantMaybe__gopurs_runtime_Value_3070183577 gopurs_runtime.Value
var once_invariantMaybe__gopurs_runtime_Value_3070183577 sync.Once
func Get_invariantMaybe__gopurs_runtime_Value_3070183577() gopurs_runtime.Value {
	once_invariantMaybe__gopurs_runtime_Value_3070183577.Do(func() {
		cache_invariantMaybe__gopurs_runtime_Value_3070183577 = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorMaybe(), "map"), f_0)
})
}))
	})
	return cache_invariantMaybe__gopurs_runtime_Value_3070183577
}

var cache_fromMaybe_prime gopurs_runtime.Value
var once_fromMaybe_prime sync.Once
func Get_fromMaybe_prime() gopurs_runtime.Value {
	once_fromMaybe_prime.Do(func() {
		cache_fromMaybe_prime = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe_prime(a_0_box)
})
	})
	return cache_fromMaybe_prime
}

var cache_fromMaybe gopurs_runtime.Value
var once_fromMaybe sync.Once
func Get_fromMaybe() gopurs_runtime.Value {
	once_fromMaybe.Do(func() {
		cache_fromMaybe = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe
}

var cache_fromMaybe__gopurs_runtime_Value_430429096 gopurs_runtime.Value
var once_fromMaybe__gopurs_runtime_Value_430429096 sync.Once
func Get_fromMaybe__gopurs_runtime_Value_430429096() gopurs_runtime.Value {
	once_fromMaybe__gopurs_runtime_Value_430429096.Do(func() {
		cache_fromMaybe__gopurs_runtime_Value_430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__gopurs_runtime_Value_430429096(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__gopurs_runtime_Value_430429096
}

var cache_fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		cache_fromJust = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust
}

var cache_fromJust__gopurs_runtime_Value_1791383420 gopurs_runtime.Value
var once_fromJust__gopurs_runtime_Value_1791383420 sync.Once
func Get_fromJust__gopurs_runtime_Value_1791383420() gopurs_runtime.Value {
	once_fromJust__gopurs_runtime_Value_1791383420.Do(func() {
		cache_fromJust__gopurs_runtime_Value_1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__gopurs_runtime_Value_1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__gopurs_runtime_Value_1791383420
}

var cache_extendMaybe gopurs_runtime.Value
var once_extendMaybe sync.Once
func Get_extendMaybe() gopurs_runtime.Value {
	once_extendMaybe.Do(func() {
		cache_extendMaybe = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1)})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_extendMaybe
}

var cache_extendMaybe__gopurs_runtime_Value_3503256297 gopurs_runtime.Value
var once_extendMaybe__gopurs_runtime_Value_3503256297 sync.Once
func Get_extendMaybe__gopurs_runtime_Value_3503256297() gopurs_runtime.Value {
	once_extendMaybe__gopurs_runtime_Value_3503256297.Do(func() {
		cache_extendMaybe__gopurs_runtime_Value_3503256297 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1)})}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_extendMaybe__gopurs_runtime_Value_3503256297
}

var cache_eqMaybe gopurs_runtime.Value
var once_eqMaybe sync.Once
func Get_eqMaybe() gopurs_runtime.Value {
	once_eqMaybe.Do(func() {
		cache_eqMaybe = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqMaybe(dictEq_0_box)
})
	})
	return cache_eqMaybe
}

var cache_ordMaybe gopurs_runtime.Value
var once_ordMaybe sync.Once
func Get_ordMaybe() gopurs_runtime.Value {
	once_ordMaybe.Do(func() {
		cache_ordMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordMaybe(dictOrd_0_box)
})
	})
	return cache_ordMaybe
}

var cache_eq1Maybe gopurs_runtime.Value
var once_eq1Maybe sync.Once
func Get_eq1Maybe() gopurs_runtime.Value {
	once_eq1Maybe.Do(func() {
		cache_eq1Maybe = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t1 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
})
})
}))
	})
	return cache_eq1Maybe
}

var cache_eq1Maybe__gopurs_runtime_Value_1662522654 gopurs_runtime.Value
var once_eq1Maybe__gopurs_runtime_Value_1662522654 sync.Once
func Get_eq1Maybe__gopurs_runtime_Value_1662522654() gopurs_runtime.Value {
	once_eq1Maybe__gopurs_runtime_Value_1662522654.Do(func() {
		cache_eq1Maybe__gopurs_runtime_Value_1662522654 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t1 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
})
})
}))
	})
	return cache_eq1Maybe__gopurs_runtime_Value_1662522654
}

var cache_ord1Maybe gopurs_runtime.Value
var once_ord1Maybe sync.Once
func Get_ord1Maybe() gopurs_runtime.Value {
	once_ord1Maybe.Do(func() {
		cache_ord1Maybe = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Maybe()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t1 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t0.IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_ord1Maybe
}

var cache_ord1Maybe__gopurs_runtime_Value_1052104681 gopurs_runtime.Value
var once_ord1Maybe__gopurs_runtime_Value_1052104681 sync.Once
func Get_ord1Maybe__gopurs_runtime_Value_1052104681() gopurs_runtime.Value {
	once_ord1Maybe__gopurs_runtime_Value_1052104681.Do(func() {
		cache_ord1Maybe__gopurs_runtime_Value_1052104681 = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Maybe()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t1 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t0.IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_ord1Maybe__gopurs_runtime_Value_1052104681
}

var cache_boundedMaybe gopurs_runtime.Value
var once_boundedMaybe sync.Once
func Get_boundedMaybe() gopurs_runtime.Value {
	once_boundedMaybe.Do(func() {
		cache_boundedMaybe = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedMaybe(dictBounded_0_box)
})
	})
	return cache_boundedMaybe
}

var cache_applyMaybe gopurs_runtime.Value
var once_applyMaybe sync.Once
func Get_applyMaybe() gopurs_runtime.Value {
	once_applyMaybe.Do(func() {
		cache_applyMaybe = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorMaybe(), "map"), (*Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe
}

var cache_applyMaybe__ptrControl_Apply_Constructor_Apply_ptrConstructor_Just_gopurs_runtime_Value___3561700045 gopurs_runtime.Value
var once_applyMaybe__ptrControl_Apply_Constructor_Apply_ptrConstructor_Just_gopurs_runtime_Value___3561700045 sync.Once
func Get_applyMaybe__ptrControl_Apply_Constructor_Apply_ptrConstructor_Just_gopurs_runtime_Value___3561700045() gopurs_runtime.Value {
	once_applyMaybe__ptrControl_Apply_Constructor_Apply_ptrConstructor_Just_gopurs_runtime_Value___3561700045.Do(func() {
		cache_applyMaybe__ptrControl_Apply_Constructor_Apply_ptrConstructor_Just_gopurs_runtime_Value___3561700045 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&pkg_Control_Apply.Constructor_Apply[*Constructor_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorMaybe(), "map"), (*Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
})})}
	})
	return cache_applyMaybe__ptrControl_Apply_Constructor_Apply_ptrConstructor_Just_gopurs_runtime_Value___3561700045
}

var cache_applyMaybe__gopurs_runtime_Value_3698865467 gopurs_runtime.Value
var once_applyMaybe__gopurs_runtime_Value_3698865467 sync.Once
func Get_applyMaybe__gopurs_runtime_Value_3698865467() gopurs_runtime.Value {
	once_applyMaybe__gopurs_runtime_Value_3698865467.Do(func() {
		cache_applyMaybe__gopurs_runtime_Value_3698865467 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorMaybe(), "map"), (*Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__gopurs_runtime_Value_3698865467
}

var cache_bindMaybe gopurs_runtime.Value
var once_bindMaybe sync.Once
func Get_bindMaybe() gopurs_runtime.Value {
	once_bindMaybe.Do(func() {
		cache_bindMaybe = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindMaybe
}

var cache_bindMaybe__ptrControl_Bind_Constructor_Bind_ptrConstructor_Just_gopurs_runtime_Value___1910292045 gopurs_runtime.Value
var once_bindMaybe__ptrControl_Bind_Constructor_Bind_ptrConstructor_Just_gopurs_runtime_Value___1910292045 sync.Once
func Get_bindMaybe__ptrControl_Bind_Constructor_Bind_ptrConstructor_Just_gopurs_runtime_Value___1910292045() gopurs_runtime.Value {
	once_bindMaybe__ptrControl_Bind_Constructor_Bind_ptrConstructor_Just_gopurs_runtime_Value___1910292045.Do(func() {
		cache_bindMaybe__ptrControl_Bind_Constructor_Bind_ptrConstructor_Just_gopurs_runtime_Value___1910292045 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&pkg_Control_Bind.Constructor_Bind[*Constructor_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
})})}
	})
	return cache_bindMaybe__ptrControl_Bind_Constructor_Bind_ptrConstructor_Just_gopurs_runtime_Value___1910292045
}

var cache_bindMaybe__gopurs_runtime_Value_3591110311 gopurs_runtime.Value
var once_bindMaybe__gopurs_runtime_Value_3591110311 sync.Once
func Get_bindMaybe__gopurs_runtime_Value_3591110311() gopurs_runtime.Value {
	once_bindMaybe__gopurs_runtime_Value_3591110311.Do(func() {
		cache_bindMaybe__gopurs_runtime_Value_3591110311 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindMaybe__gopurs_runtime_Value_3591110311
}

var cache_semiringMaybe gopurs_runtime.Value
var once_semiringMaybe sync.Once
func Get_semiringMaybe() gopurs_runtime.Value {
	once_semiringMaybe.Do(func() {
		cache_semiringMaybe = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semiringMaybe(dictSemiring_0_box)
})
	})
	return cache_semiringMaybe
}

var cache_applicativeMaybe gopurs_runtime.Value
var once_applicativeMaybe sync.Once
func Get_applicativeMaybe() gopurs_runtime.Value {
	once_applicativeMaybe.Do(func() {
		cache_applicativeMaybe = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
}), Get_Just())
	})
	return cache_applicativeMaybe
}

var cache_applicativeMaybe__ptrControl_Applicative_Constructor_Applicative_ptrConstructor_Just_gopurs_runtime_Value___3016118221 gopurs_runtime.Value
var once_applicativeMaybe__ptrControl_Applicative_Constructor_Applicative_ptrConstructor_Just_gopurs_runtime_Value___3016118221 sync.Once
func Get_applicativeMaybe__ptrControl_Applicative_Constructor_Applicative_ptrConstructor_Just_gopurs_runtime_Value___3016118221() gopurs_runtime.Value {
	once_applicativeMaybe__ptrControl_Applicative_Constructor_Applicative_ptrConstructor_Just_gopurs_runtime_Value___3016118221.Do(func() {
		cache_applicativeMaybe__ptrControl_Applicative_Constructor_Applicative_ptrConstructor_Just_gopurs_runtime_Value___3016118221 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&pkg_Control_Applicative.Constructor_Applicative[*Constructor_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
}), Get_Just()})}
	})
	return cache_applicativeMaybe__ptrControl_Applicative_Constructor_Applicative_ptrConstructor_Just_gopurs_runtime_Value___3016118221
}

var cache_applicativeMaybe__gopurs_runtime_Value_500933224 gopurs_runtime.Value
var once_applicativeMaybe__gopurs_runtime_Value_500933224 sync.Once
func Get_applicativeMaybe__gopurs_runtime_Value_500933224() gopurs_runtime.Value {
	once_applicativeMaybe__gopurs_runtime_Value_500933224.Do(func() {
		cache_applicativeMaybe__gopurs_runtime_Value_500933224 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
}), Get_Just())
	})
	return cache_applicativeMaybe__gopurs_runtime_Value_500933224
}

var cache_monadMaybe gopurs_runtime.Value
var once_monadMaybe sync.Once
func Get_monadMaybe() gopurs_runtime.Value {
	once_monadMaybe.Do(func() {
		cache_monadMaybe = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindMaybe()
}))
	})
	return cache_monadMaybe
}

var cache_monadMaybe__gopurs_runtime_Value_3072900051 gopurs_runtime.Value
var once_monadMaybe__gopurs_runtime_Value_3072900051 sync.Once
func Get_monadMaybe__gopurs_runtime_Value_3072900051() gopurs_runtime.Value {
	once_monadMaybe__gopurs_runtime_Value_3072900051.Do(func() {
		cache_monadMaybe__gopurs_runtime_Value_3072900051 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindMaybe()
}))
	})
	return cache_monadMaybe__gopurs_runtime_Value_3072900051
}

var cache_altMaybe gopurs_runtime.Value
var once_altMaybe sync.Once
func Get_altMaybe() gopurs_runtime.Value {
	once_altMaybe.Do(func() {
		cache_altMaybe = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v1_1))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v_0))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_altMaybe
}

var cache_altMaybe__gopurs_runtime_Value_4201091523 gopurs_runtime.Value
var once_altMaybe__gopurs_runtime_Value_4201091523 sync.Once
func Get_altMaybe__gopurs_runtime_Value_4201091523() gopurs_runtime.Value {
	once_altMaybe__gopurs_runtime_Value_4201091523.Do(func() {
		cache_altMaybe__gopurs_runtime_Value_4201091523 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v1_1))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v_0))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_altMaybe__gopurs_runtime_Value_4201091523
}

var cache_plusMaybe gopurs_runtime.Value
var once_plusMaybe sync.Once
func Get_plusMaybe() gopurs_runtime.Value {
	once_plusMaybe.Do(func() {
		cache_plusMaybe = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altMaybe()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))})
	})
	return cache_plusMaybe
}

var cache_plusMaybe__gopurs_runtime_Value_400696082 gopurs_runtime.Value
var once_plusMaybe__gopurs_runtime_Value_400696082 sync.Once
func Get_plusMaybe__gopurs_runtime_Value_400696082() gopurs_runtime.Value {
	once_plusMaybe__gopurs_runtime_Value_400696082.Do(func() {
		cache_plusMaybe__gopurs_runtime_Value_400696082 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altMaybe()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))})
	})
	return cache_plusMaybe__gopurs_runtime_Value_400696082
}

var cache_alternativeMaybe gopurs_runtime.Value
var once_alternativeMaybe sync.Once
func Get_alternativeMaybe() gopurs_runtime.Value {
	once_alternativeMaybe.Do(func() {
		cache_alternativeMaybe = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusMaybe()
}))
	})
	return cache_alternativeMaybe
}

var cache_alternativeMaybe__ptrControl_Alternative_Constructor_Alternative_ptrConstructor_Just_gopurs_runtime_Value___196315533 gopurs_runtime.Value
var once_alternativeMaybe__ptrControl_Alternative_Constructor_Alternative_ptrConstructor_Just_gopurs_runtime_Value___196315533 sync.Once
func Get_alternativeMaybe__ptrControl_Alternative_Constructor_Alternative_ptrConstructor_Just_gopurs_runtime_Value___196315533() gopurs_runtime.Value {
	once_alternativeMaybe__ptrControl_Alternative_Constructor_Alternative_ptrConstructor_Just_gopurs_runtime_Value___196315533.Do(func() {
		cache_alternativeMaybe__ptrControl_Alternative_Constructor_Alternative_ptrConstructor_Just_gopurs_runtime_Value___196315533 = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&pkg_Control_Alternative.Constructor_Alternative[*Constructor_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusMaybe()
})})}
	})
	return cache_alternativeMaybe__ptrControl_Alternative_Constructor_Alternative_ptrConstructor_Just_gopurs_runtime_Value___196315533
}

type Constructor_Nothing[T_a any] struct {
	Rc uint32
}


type Constructor_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showMaybe(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Just "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0), gopurs_runtime.Str(")"))).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Str("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
}

func Call_semigroupMaybe(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v1_2))}
goto end_branch_0
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v_1))}
goto end_branch_0
} else {

}
}
{
if ((v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil)) && ((v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0)})}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
}

func Call_optional(dictAlt_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlt_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dictAlt_0_loop
_ = dictAlt_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlt_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictAlt_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Just(), a_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))
})
})
}

func Call_monoidMaybe(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
semigroupMaybe1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v1_2))}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v_1))}
goto end_branch_1
} else {

}
}
{
if ((v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil)) && ((v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0)})}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t1))}
})
}))
_ = semigroupMaybe1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMaybe1_1_0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_maybe_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_maybe_prime__gopurs_runtime_Value_2328206764(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_maybe(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_maybe__gopurs_runtime_Value_3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_isNothing(v2_0_loop *Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isNothing__gopurs_runtime_Value_2514352589(v2_0_loop *Constructor_Just[int64]) bool {
var v2_0 *Constructor_Just[int64] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isNothing__gopurs_runtime_Value_2591355336(v2_0_loop *Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isNothing__gopurs_runtime_Value_4206805139(v2_0_loop *Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isJust(v2_0_loop *Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isJust__gopurs_runtime_Value_2514352589(v2_0_loop *Constructor_Just[int64]) bool {
var v2_0 *Constructor_Just[int64] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isJust__gopurs_runtime_Value_2475527019(v2_0_loop *Constructor_Just[string]) bool {
var v2_0 *Constructor_Just[string] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isJust__gopurs_runtime_Value_2591355336(v2_0_loop *Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isJust__gopurs_runtime_Value_4206805139(v2_0_loop *Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_fromMaybe_prime(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply2(Get_maybe_prime(), a_0, Get_identity())
}

func Call_fromMaybe(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
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

func Call_fromMaybe__gopurs_runtime_Value_430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
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

func Call_fromJust(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
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

func Call_fromJust__gopurs_runtime_Value_1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
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

func Call_eqMaybe(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t1 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
})
}))
}

func Call_ordMaybe(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
eqMaybe1_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(false)
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil)) && ((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr != nil)) {
__t2 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(false)
}
end_branch_2:
return gopurs_runtime.Bool((__t2.IntVal) != (0))
})
}))
_ = eqMaybe1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
var __t5 gopurs_runtime.Value
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil)) && ((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr != nil)) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t4.IntVal)), UnsafePtr: nil}
})
}))
}

func Call_boundedMaybe(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_3
eqMaybe1_2_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
var __t5 gopurs_runtime.Value
{
if (y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Bool(true)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Bool(false)
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if ((x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil)) && ((y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr != nil)) {
__t4 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "eq"), (*Constructor_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
return gopurs_runtime.Bool((__t4.IntVal) != (0))
})
}))
_ = eqMaybe1_2_2
ordMaybe1_1_0 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_2_2
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
var __t7 gopurs_runtime.Value
{
if (y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
if (y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if ((x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil)) && ((y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr != nil)) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compare"), (*Constructor_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t6.IntVal)), UnsafePtr: nil}
})
}))
_ = ordMaybe1_1_0
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ordMaybe1_1_0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictBounded_0, "top")})}))})
}

func Call_semiringMaybe(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
mul_1_0 := gopurs_runtime.RecordGet(dictSemiring_0, "mul")
_ = mul_1_0
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v1_3))}
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](v_2))}
goto end_branch_1
} else {

}
}
{
if ((v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil)) && ((v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)})}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](__t1))}
})
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorMaybe(), "map"), mul_1_0, x_2), y_3)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictSemiring_0, "one")})}))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}


