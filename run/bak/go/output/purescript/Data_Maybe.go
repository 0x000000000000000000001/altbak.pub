package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Maybe_identity gopurs_runtime.Value
var once_Data_Maybe_identity sync.Once
func Get_Data_Maybe_identity() gopurs_runtime.Value {
	once_Data_Maybe_identity.Do(func() {
		cache_Data_Maybe_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_identity(x_0_box)
})
	})
	return cache_Data_Maybe_identity
}

var cache_Data_Maybe_Nothing gopurs_runtime.Value
var once_Data_Maybe_Nothing sync.Once
func Get_Data_Maybe_Nothing() gopurs_runtime.Value {
	once_Data_Maybe_Nothing.Do(func() {
		cache_Data_Maybe_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
	})
	return cache_Data_Maybe_Nothing
}

var cache_Data_Maybe_Just gopurs_runtime.Value
var once_Data_Maybe_Just sync.Once
func Get_Data_Maybe_Just() gopurs_runtime.Value {
	once_Data_Maybe_Just.Do(func() {
		cache_Data_Maybe_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, value0})}
})
	})
	return cache_Data_Maybe_Just
}

var cache_Data_Maybe_showMaybe gopurs_runtime.Value
var once_Data_Maybe_showMaybe sync.Once
func Get_Data_Maybe_showMaybe() gopurs_runtime.Value {
	once_Data_Maybe_showMaybe.Do(func() {
		cache_Data_Maybe_showMaybe = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_showMaybe(dictShow_0_box)
})
	})
	return cache_Data_Maybe_showMaybe
}

var cache_Data_Maybe_semigroupMaybe gopurs_runtime.Value
var once_Data_Maybe_semigroupMaybe sync.Once
func Get_Data_Maybe_semigroupMaybe() gopurs_runtime.Value {
	once_Data_Maybe_semigroupMaybe.Do(func() {
		cache_Data_Maybe_semigroupMaybe = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_semigroupMaybe(dictSemigroup_0_box)
})
	})
	return cache_Data_Maybe_semigroupMaybe
}

var cache_Data_Maybe_optional gopurs_runtime.Value
var once_Data_Maybe_optional sync.Once
func Get_Data_Maybe_optional() gopurs_runtime.Value {
	once_Data_Maybe_optional.Do(func() {
		cache_Data_Maybe_optional = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_optional(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](dictAlt_0_box))
})
	})
	return cache_Data_Maybe_optional
}

var cache_Data_Maybe_monoidMaybe gopurs_runtime.Value
var once_Data_Maybe_monoidMaybe sync.Once
func Get_Data_Maybe_monoidMaybe() gopurs_runtime.Value {
	once_Data_Maybe_monoidMaybe.Do(func() {
		cache_Data_Maybe_monoidMaybe = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_monoidMaybe(dictSemigroup_0_box)
})
	})
	return cache_Data_Maybe_monoidMaybe
}

var cache_Data_Maybe_maybe_prime gopurs_runtime.Value
var once_Data_Maybe_maybe_prime sync.Once
func Get_Data_Maybe_maybe_prime() gopurs_runtime.Value {
	once_Data_Maybe_maybe_prime.Do(func() {
		cache_Data_Maybe_maybe_prime = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe_prime(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe_prime
}

var cache_Data_Maybe_maybe gopurs_runtime.Value
var once_Data_Maybe_maybe sync.Once
func Get_Data_Maybe_maybe() gopurs_runtime.Value {
	once_Data_Maybe_maybe.Do(func() {
		cache_Data_Maybe_maybe = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe
}

var cache_Data_Maybe_isNothing gopurs_runtime.Value
var once_Data_Maybe_isNothing sync.Once
func Get_Data_Maybe_isNothing() gopurs_runtime.Value {
	once_Data_Maybe_isNothing.Do(func() {
		cache_Data_Maybe_isNothing = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing
}

var cache_Data_Maybe_isJust gopurs_runtime.Value
var once_Data_Maybe_isJust sync.Once
func Get_Data_Maybe_isJust() gopurs_runtime.Value {
	once_Data_Maybe_isJust.Do(func() {
		cache_Data_Maybe_isJust = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust
}

var cache_Data_Maybe_genericMaybe gopurs_runtime.Value
var once_Data_Maybe_genericMaybe sync.Once
func Get_Data_Maybe_genericMaybe() gopurs_runtime.Value {
	once_Data_Maybe_genericMaybe.Do(func() {
		cache_Data_Maybe_genericMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Generic{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr{1, (*Constructor_Data_Maybe_Just)(x_0.UnsafePtr).V0})}
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
var __t1 *Constructor_Data_Maybe_Just
{
if (x_0.Type == 9 && x_0.IntVal == 3478632216) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 492034566) {
__t1 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Generic_Rep_Inr)(x_0.UnsafePtr).V0}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})}
	})
	return cache_Data_Maybe_genericMaybe
}

var cache_Data_Maybe_functorMaybe gopurs_runtime.Value
var once_Data_Maybe_functorMaybe sync.Once
func Get_Data_Maybe_functorMaybe() gopurs_runtime.Value {
	once_Data_Maybe_functorMaybe.Do(func() {
		cache_Data_Maybe_functorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Maybe_Just)(v1_1.UnsafePtr).V0)}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
})})}
	})
	return cache_Data_Maybe_functorMaybe
}

var cache_Data_Maybe_invariantMaybe gopurs_runtime.Value
var once_Data_Maybe_invariantMaybe sync.Once
func Get_Data_Maybe_invariantMaybe() gopurs_runtime.Value {
	once_Data_Maybe_invariantMaybe.Do(func() {
		cache_Data_Maybe_invariantMaybe = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Invariant_Invariant{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Maybe_Just)(v1_2.UnsafePtr).V0)}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
})
})})}
	})
	return cache_Data_Maybe_invariantMaybe
}

var cache_Data_Maybe_fromMaybe_prime gopurs_runtime.Value
var once_Data_Maybe_fromMaybe_prime sync.Once
func Get_Data_Maybe_fromMaybe_prime() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe_prime.Do(func() {
		cache_Data_Maybe_fromMaybe_prime = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromMaybe_prime(a_0_box)
})
	})
	return cache_Data_Maybe_fromMaybe_prime
}

var cache_Data_Maybe_fromMaybe gopurs_runtime.Value
var once_Data_Maybe_fromMaybe sync.Once
func Get_Data_Maybe_fromMaybe() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe.Do(func() {
		cache_Data_Maybe_fromMaybe = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromMaybe(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_1_box))
})
	})
	return cache_Data_Maybe_fromMaybe
}

var cache_Data_Maybe_fromJust gopurs_runtime.Value
var once_Data_Maybe_fromJust sync.Once
func Get_Data_Maybe_fromJust() gopurs_runtime.Value {
	once_Data_Maybe_fromJust.Do(func() {
		cache_Data_Maybe_fromJust = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromJust(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_1_box))
})
	})
	return cache_Data_Maybe_fromJust
}

var cache_Data_Maybe_extendMaybe gopurs_runtime.Value
var once_Data_Maybe_extendMaybe sync.Once
func Get_Data_Maybe_extendMaybe() gopurs_runtime.Value {
	once_Data_Maybe_extendMaybe.Do(func() {
		cache_Data_Maybe_extendMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Maybe_functorMaybe()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1))})}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
})})}
	})
	return cache_Data_Maybe_extendMaybe
}

var cache_Data_Maybe_eqMaybe gopurs_runtime.Value
var once_Data_Maybe_eqMaybe sync.Once
func Get_Data_Maybe_eqMaybe() gopurs_runtime.Value {
	once_Data_Maybe_eqMaybe.Do(func() {
		cache_Data_Maybe_eqMaybe = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_eqMaybe(dictEq_0_box)
})
	})
	return cache_Data_Maybe_eqMaybe
}

var cache_Data_Maybe_ordMaybe gopurs_runtime.Value
var once_Data_Maybe_ordMaybe sync.Once
func Get_Data_Maybe_ordMaybe() gopurs_runtime.Value {
	once_Data_Maybe_ordMaybe.Do(func() {
		cache_Data_Maybe_ordMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_ordMaybe(dictOrd_0_box)
})
	})
	return cache_Data_Maybe_ordMaybe
}

var cache_Data_Maybe_eq1Maybe gopurs_runtime.Value
var once_Data_Maybe_eq1Maybe sync.Once
func Get_Data_Maybe_eq1Maybe() gopurs_runtime.Value {
	once_Data_Maybe_eq1Maybe.Do(func() {
		cache_Data_Maybe_eq1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t0 bool
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t1 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Maybe_Just)(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_2.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
})
})})}
	})
	return cache_Data_Maybe_eq1Maybe
}

var cache_Data_Maybe_ord1Maybe gopurs_runtime.Value
var once_Data_Maybe_ord1Maybe sync.Once
func Get_Data_Maybe_ord1Maybe() gopurs_runtime.Value {
	once_Data_Maybe_ord1Maybe.Do(func() {
		cache_Data_Maybe_ord1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Maybe_eq1Maybe()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 uint32
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t0 uint32
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t1 = 380165415
goto end_branch_1
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t1 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Maybe_Just)(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_2.UnsafePtr).V0).IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
})
})
})})}
	})
	return cache_Data_Maybe_ord1Maybe
}

var cache_Data_Maybe_boundedMaybe gopurs_runtime.Value
var once_Data_Maybe_boundedMaybe sync.Once
func Get_Data_Maybe_boundedMaybe() gopurs_runtime.Value {
	once_Data_Maybe_boundedMaybe.Do(func() {
		cache_Data_Maybe_boundedMaybe = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_boundedMaybe(dictBounded_0_box)
})
	})
	return cache_Data_Maybe_boundedMaybe
}

var cache_Data_Maybe_applyMaybe gopurs_runtime.Value
var once_Data_Maybe_applyMaybe sync.Once
func Get_Data_Maybe_applyMaybe() gopurs_runtime.Value {
	once_Data_Maybe_applyMaybe.Do(func() {
		cache_Data_Maybe_applyMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Maybe_functorMaybe()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1)
if (__t_tag_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((*Constructor_Data_Maybe_Just)(v_0.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(v1_1.UnsafePtr).V0)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
})
})})}
	})
	return cache_Data_Maybe_applyMaybe
}

var cache_Data_Maybe_bindMaybe gopurs_runtime.Value
var once_Data_Maybe_bindMaybe sync.Once
func Get_Data_Maybe_bindMaybe() gopurs_runtime.Value {
	once_Data_Maybe_bindMaybe.Do(func() {
		cache_Data_Maybe_bindMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Maybe_applyMaybe()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (*Constructor_Data_Maybe_Just)(v_0.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
})})}
	})
	return cache_Data_Maybe_bindMaybe
}

var cache_Data_Maybe_semiringMaybe gopurs_runtime.Value
var once_Data_Maybe_semiringMaybe sync.Once
func Get_Data_Maybe_semiringMaybe() gopurs_runtime.Value {
	once_Data_Maybe_semiringMaybe.Do(func() {
		cache_Data_Maybe_semiringMaybe = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_semiringMaybe(dictSemiring_0_box)
})
	})
	return cache_Data_Maybe_semiringMaybe
}

var cache_Data_Maybe_applicativeMaybe gopurs_runtime.Value
var once_Data_Maybe_applicativeMaybe sync.Once
func Get_Data_Maybe_applicativeMaybe() gopurs_runtime.Value {
	once_Data_Maybe_applicativeMaybe.Do(func() {
		cache_Data_Maybe_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Maybe_applyMaybe()))}
}), Get_Data_Maybe_Just()})}
	})
	return cache_Data_Maybe_applicativeMaybe
}

var cache_Data_Maybe_monadMaybe gopurs_runtime.Value
var once_Data_Maybe_monadMaybe sync.Once
func Get_Data_Maybe_monadMaybe() gopurs_runtime.Value {
	once_Data_Maybe_monadMaybe.Do(func() {
		cache_Data_Maybe_monadMaybe = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Maybe_applicativeMaybe()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_Maybe_bindMaybe()))}
})})}
	})
	return cache_Data_Maybe_monadMaybe
}

var cache_Data_Maybe_altMaybe gopurs_runtime.Value
var once_Data_Maybe_altMaybe sync.Once
func Get_Data_Maybe_altMaybe() gopurs_runtime.Value {
	once_Data_Maybe_altMaybe.Do(func() {
		cache_Data_Maybe_altMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Maybe_functorMaybe()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
})})}
	})
	return cache_Data_Maybe_altMaybe
}

var cache_Data_Maybe_plusMaybe gopurs_runtime.Value
var once_Data_Maybe_plusMaybe sync.Once
func Get_Data_Maybe_plusMaybe() gopurs_runtime.Value {
	once_Data_Maybe_plusMaybe.Do(func() {
		cache_Data_Maybe_plusMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Data_Maybe_altMaybe()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})}
	})
	return cache_Data_Maybe_plusMaybe
}

var cache_Data_Maybe_alternativeMaybe gopurs_runtime.Value
var once_Data_Maybe_alternativeMaybe sync.Once
func Get_Data_Maybe_alternativeMaybe() gopurs_runtime.Value {
	once_Data_Maybe_alternativeMaybe.Do(func() {
		cache_Data_Maybe_alternativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Maybe_applicativeMaybe()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_Maybe_plusMaybe()))}
})})}
	})
	return cache_Data_Maybe_alternativeMaybe
}

var cache_Data_Maybe_fromJust__2181618881 gopurs_runtime.Value
var once_Data_Maybe_fromJust__2181618881 sync.Once
func Get_Data_Maybe_fromJust__2181618881() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__2181618881.Do(func() {
		cache_Data_Maybe_fromJust__2181618881 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_fromJust__2181618881(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_1_box)))
})
	})
	return cache_Data_Maybe_fromJust__2181618881
}

var cache_Data_Maybe_fromJust__1577979644 gopurs_runtime.Value
var once_Data_Maybe_fromJust__1577979644 sync.Once
func Get_Data_Maybe_fromJust__1577979644() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__1577979644.Do(func() {
		cache_Data_Maybe_fromJust__1577979644 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_fromJust__1577979644(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box)))
})
	})
	return cache_Data_Maybe_fromJust__1577979644
}

var cache_Data_Maybe_fromJust__4121089788 gopurs_runtime.Value
var once_Data_Maybe_fromJust__4121089788 sync.Once
func Get_Data_Maybe_fromJust__4121089788() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__4121089788.Do(func() {
		cache_Data_Maybe_fromJust__4121089788 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Maybe_fromJust__4121089788(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box)))
})
	})
	return cache_Data_Maybe_fromJust__4121089788
}

var cache_Data_Maybe_fromJust__1791383420 gopurs_runtime.Value
var once_Data_Maybe_fromJust__1791383420 sync.Once
func Get_Data_Maybe_fromJust__1791383420() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__1791383420.Do(func() {
		cache_Data_Maybe_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_1_box))
})
	})
	return cache_Data_Maybe_fromJust__1791383420
}

var cache_Data_Maybe_fromJust__911089788 gopurs_runtime.Value
var once_Data_Maybe_fromJust__911089788 sync.Once
func Get_Data_Maybe_fromJust__911089788() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__911089788.Do(func() {
		cache_Data_Maybe_fromJust__911089788 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Maybe_fromJust__911089788(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Maybe_fromJust__911089788
}

var cache_Data_Maybe_fromJust__3897574428 gopurs_runtime.Value
var once_Data_Maybe_fromJust__3897574428 sync.Once
func Get_Data_Maybe_fromJust__3897574428() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__3897574428.Do(func() {
		cache_Data_Maybe_fromJust__3897574428 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Maybe_fromJust__3897574428(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Maybe_fromJust__3897574428
}

var cache_Data_Maybe_fromJust__1478027324 gopurs_runtime.Value
var once_Data_Maybe_fromJust__1478027324 sync.Once
func Get_Data_Maybe_fromJust__1478027324() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__1478027324.Do(func() {
		cache_Data_Maybe_fromJust__1478027324 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Maybe_fromJust__1478027324(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box)))
})
	})
	return cache_Data_Maybe_fromJust__1478027324
}

var cache_Data_Maybe_fromJust__4142563260 gopurs_runtime.Value
var once_Data_Maybe_fromJust__4142563260 sync.Once
func Get_Data_Maybe_fromJust__4142563260() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__4142563260.Do(func() {
		cache_Data_Maybe_fromJust__4142563260 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromJust__4142563260(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromJust__4142563260
}

var cache_Data_Maybe_fromJust__3809843644 gopurs_runtime.Value
var once_Data_Maybe_fromJust__3809843644 sync.Once
func Get_Data_Maybe_fromJust__3809843644() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__3809843644.Do(func() {
		cache_Data_Maybe_fromJust__3809843644 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromJust__3809843644(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromJust__3809843644
}

var cache_Data_Maybe_fromJust__965748316 gopurs_runtime.Value
var once_Data_Maybe_fromJust__965748316 sync.Once
func Get_Data_Maybe_fromJust__965748316() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__965748316.Do(func() {
		cache_Data_Maybe_fromJust__965748316 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_fromJust__965748316(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box)))}
})
	})
	return cache_Data_Maybe_fromJust__965748316
}

var cache_Data_Maybe_fromJust__755886620 gopurs_runtime.Value
var once_Data_Maybe_fromJust__755886620 sync.Once
func Get_Data_Maybe_fromJust__755886620() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__755886620.Do(func() {
		cache_Data_Maybe_fromJust__755886620 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_fromJust__755886620(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box)))}
})
	})
	return cache_Data_Maybe_fromJust__755886620
}

var cache_Data_Maybe_fromMaybe__1972796397 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__1972796397 sync.Once
func Get_Data_Maybe_fromMaybe__1972796397() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__1972796397.Do(func() {
		cache_Data_Maybe_fromMaybe__1972796397 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_fromMaybe__1972796397(a_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_1_box)))
})
	})
	return cache_Data_Maybe_fromMaybe__1972796397
}

var cache_Data_Maybe_fromMaybe__430429096 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__430429096 sync.Once
func Get_Data_Maybe_fromMaybe__430429096() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__430429096.Do(func() {
		cache_Data_Maybe_fromMaybe__430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromMaybe__430429096(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_1_box))
})
	})
	return cache_Data_Maybe_fromMaybe__430429096
}

var cache_Data_Maybe_fromMaybe__656947263 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__656947263 sync.Once
func Get_Data_Maybe_fromMaybe__656947263() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__656947263.Do(func() {
		cache_Data_Maybe_fromMaybe__656947263 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromMaybe__656947263(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_1_box))
})
	})
	return cache_Data_Maybe_fromMaybe__656947263
}

var cache_Data_Maybe_fromMaybe__18840980 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__18840980 sync.Once
func Get_Data_Maybe_fromMaybe__18840980() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__18840980.Do(func() {
		cache_Data_Maybe_fromMaybe__18840980 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromMaybe__18840980(uint32(a_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_1_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromMaybe__18840980
}

var cache_Data_Maybe_fromMaybe__737056608 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__737056608 sync.Once
func Get_Data_Maybe_fromMaybe__737056608() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__737056608.Do(func() {
		cache_Data_Maybe_fromMaybe__737056608 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_fromMaybe__737056608(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_1_box)))}
})
	})
	return cache_Data_Maybe_fromMaybe__737056608
}

var cache_Data_Maybe_fromMaybe__2067158953 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__2067158953 sync.Once
func Get_Data_Maybe_fromMaybe__2067158953() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__2067158953.Do(func() {
		cache_Data_Maybe_fromMaybe__2067158953 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_fromMaybe__2067158953(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_1_box)))}
})
	})
	return cache_Data_Maybe_fromMaybe__2067158953
}

var cache_Data_Maybe_isJust__2514352589 gopurs_runtime.Value
var once_Data_Maybe_isJust__2514352589 sync.Once
func Get_Data_Maybe_isJust__2514352589() gopurs_runtime.Value {
	once_Data_Maybe_isJust__2514352589.Do(func() {
		cache_Data_Maybe_isJust__2514352589 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__2514352589(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust__2514352589
}

var cache_Data_Maybe_isJust__2475527019 gopurs_runtime.Value
var once_Data_Maybe_isJust__2475527019 sync.Once
func Get_Data_Maybe_isJust__2475527019() gopurs_runtime.Value {
	once_Data_Maybe_isJust__2475527019.Do(func() {
		cache_Data_Maybe_isJust__2475527019 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__2475527019(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust__2475527019
}

var cache_Data_Maybe_isJust__2591355336 gopurs_runtime.Value
var once_Data_Maybe_isJust__2591355336 sync.Once
func Get_Data_Maybe_isJust__2591355336() gopurs_runtime.Value {
	once_Data_Maybe_isJust__2591355336.Do(func() {
		cache_Data_Maybe_isJust__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__2591355336(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust__2591355336
}

var cache_Data_Maybe_isJust__1358705270 gopurs_runtime.Value
var once_Data_Maybe_isJust__1358705270 sync.Once
func Get_Data_Maybe_isJust__1358705270() gopurs_runtime.Value {
	once_Data_Maybe_isJust__1358705270.Do(func() {
		cache_Data_Maybe_isJust__1358705270 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__1358705270(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust__1358705270
}

var cache_Data_Maybe_isJust__4165351782 gopurs_runtime.Value
var once_Data_Maybe_isJust__4165351782 sync.Once
func Get_Data_Maybe_isJust__4165351782() gopurs_runtime.Value {
	once_Data_Maybe_isJust__4165351782.Do(func() {
		cache_Data_Maybe_isJust__4165351782 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__4165351782(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust__4165351782
}

var cache_Data_Maybe_isJust__4206805139 gopurs_runtime.Value
var once_Data_Maybe_isJust__4206805139 sync.Once
func Get_Data_Maybe_isJust__4206805139() gopurs_runtime.Value {
	once_Data_Maybe_isJust__4206805139.Do(func() {
		cache_Data_Maybe_isJust__4206805139 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__4206805139(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust__4206805139
}

var cache_Data_Maybe_isNothing__2514352589 gopurs_runtime.Value
var once_Data_Maybe_isNothing__2514352589 sync.Once
func Get_Data_Maybe_isNothing__2514352589() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__2514352589.Do(func() {
		cache_Data_Maybe_isNothing__2514352589 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__2514352589(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__2514352589
}

var cache_Data_Maybe_isNothing__2591355336 gopurs_runtime.Value
var once_Data_Maybe_isNothing__2591355336 sync.Once
func Get_Data_Maybe_isNothing__2591355336() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__2591355336.Do(func() {
		cache_Data_Maybe_isNothing__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__2591355336(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__2591355336
}

var cache_Data_Maybe_isNothing__1401305026 gopurs_runtime.Value
var once_Data_Maybe_isNothing__1401305026 sync.Once
func Get_Data_Maybe_isNothing__1401305026() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__1401305026.Do(func() {
		cache_Data_Maybe_isNothing__1401305026 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__1401305026(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__1401305026
}

var cache_Data_Maybe_isNothing__1358705270 gopurs_runtime.Value
var once_Data_Maybe_isNothing__1358705270 sync.Once
func Get_Data_Maybe_isNothing__1358705270() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__1358705270.Do(func() {
		cache_Data_Maybe_isNothing__1358705270 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__1358705270(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__1358705270
}

var cache_Data_Maybe_isNothing__4206805139 gopurs_runtime.Value
var once_Data_Maybe_isNothing__4206805139 sync.Once
func Get_Data_Maybe_isNothing__4206805139() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__4206805139.Do(func() {
		cache_Data_Maybe_isNothing__4206805139 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__4206805139(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__4206805139
}

var cache_Data_Maybe_isNothing__2787066607 gopurs_runtime.Value
var once_Data_Maybe_isNothing__2787066607 sync.Once
func Get_Data_Maybe_isNothing__2787066607() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__2787066607.Do(func() {
		cache_Data_Maybe_isNothing__2787066607 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__2787066607(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__2787066607
}

var cache_Data_Maybe_isNothing__323776123 gopurs_runtime.Value
var once_Data_Maybe_isNothing__323776123 sync.Once
func Get_Data_Maybe_isNothing__323776123() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__323776123.Do(func() {
		cache_Data_Maybe_isNothing__323776123 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__323776123(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing__323776123
}

var cache_Data_Maybe_maybe__919206801 gopurs_runtime.Value
var once_Data_Maybe_maybe__919206801 sync.Once
func Get_Data_Maybe_maybe__919206801() gopurs_runtime.Value {
	once_Data_Maybe_maybe__919206801.Do(func() {
		cache_Data_Maybe_maybe__919206801 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_maybe__919206801(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__919206801
}

var cache_Data_Maybe_maybe__3735358641 gopurs_runtime.Value
var once_Data_Maybe_maybe__3735358641 sync.Once
func Get_Data_Maybe_maybe__3735358641() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3735358641.Do(func() {
		cache_Data_Maybe_maybe__3735358641 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_maybe__3735358641(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__3735358641
}

var cache_Data_Maybe_maybe__3078346790 gopurs_runtime.Value
var once_Data_Maybe_maybe__3078346790 sync.Once
func Get_Data_Maybe_maybe__3078346790() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3078346790.Do(func() {
		cache_Data_Maybe_maybe__3078346790 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__3078346790((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__3078346790
}

var cache_Data_Maybe_maybe__487722278 gopurs_runtime.Value
var once_Data_Maybe_maybe__487722278 sync.Once
func Get_Data_Maybe_maybe__487722278() gopurs_runtime.Value {
	once_Data_Maybe_maybe__487722278.Do(func() {
		cache_Data_Maybe_maybe__487722278 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__487722278((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__487722278
}

var cache_Data_Maybe_maybe__1510464358 gopurs_runtime.Value
var once_Data_Maybe_maybe__1510464358 sync.Once
func Get_Data_Maybe_maybe__1510464358() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1510464358.Do(func() {
		cache_Data_Maybe_maybe__1510464358 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__1510464358((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__1510464358
}

var cache_Data_Maybe_maybe__1594528518 gopurs_runtime.Value
var once_Data_Maybe_maybe__1594528518 sync.Once
func Get_Data_Maybe_maybe__1594528518() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1594528518.Do(func() {
		cache_Data_Maybe_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__1594528518((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__1594528518
}

var cache_Data_Maybe_maybe__3407128198 gopurs_runtime.Value
var once_Data_Maybe_maybe__3407128198 sync.Once
func Get_Data_Maybe_maybe__3407128198() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3407128198.Do(func() {
		cache_Data_Maybe_maybe__3407128198 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__3407128198((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__3407128198
}

var cache_Data_Maybe_maybe__2158452262 gopurs_runtime.Value
var once_Data_Maybe_maybe__2158452262 sync.Once
func Get_Data_Maybe_maybe__2158452262() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2158452262.Do(func() {
		cache_Data_Maybe_maybe__2158452262 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__2158452262((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__2158452262
}

var cache_Data_Maybe_maybe__2641488518 gopurs_runtime.Value
var once_Data_Maybe_maybe__2641488518 sync.Once
func Get_Data_Maybe_maybe__2641488518() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2641488518.Do(func() {
		cache_Data_Maybe_maybe__2641488518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__2641488518((v_0_box.IntVal) != (0), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__2641488518
}

var cache_Data_Maybe_maybe__3718989812 gopurs_runtime.Value
var once_Data_Maybe_maybe__3718989812 sync.Once
func Get_Data_Maybe_maybe__3718989812() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3718989812.Do(func() {
		cache_Data_Maybe_maybe__3718989812 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__3718989812(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__3718989812
}

var cache_Data_Maybe_maybe__316277428 gopurs_runtime.Value
var once_Data_Maybe_maybe__316277428 sync.Once
func Get_Data_Maybe_maybe__316277428() gopurs_runtime.Value {
	once_Data_Maybe_maybe__316277428.Do(func() {
		cache_Data_Maybe_maybe__316277428 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__316277428(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__316277428
}

var cache_Data_Maybe_maybe__1647364852 gopurs_runtime.Value
var once_Data_Maybe_maybe__1647364852 sync.Once
func Get_Data_Maybe_maybe__1647364852() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1647364852.Do(func() {
		cache_Data_Maybe_maybe__1647364852 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__1647364852(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__1647364852
}

var cache_Data_Maybe_maybe__3658316244 gopurs_runtime.Value
var once_Data_Maybe_maybe__3658316244 sync.Once
func Get_Data_Maybe_maybe__3658316244() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3658316244.Do(func() {
		cache_Data_Maybe_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__3658316244
}

var cache_Data_Maybe_maybe__1726410932 gopurs_runtime.Value
var once_Data_Maybe_maybe__1726410932 sync.Once
func Get_Data_Maybe_maybe__1726410932() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1726410932.Do(func() {
		cache_Data_Maybe_maybe__1726410932 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__1726410932(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__1726410932
}

var cache_Data_Maybe_maybe__1732740436 gopurs_runtime.Value
var once_Data_Maybe_maybe__1732740436 sync.Once
func Get_Data_Maybe_maybe__1732740436() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1732740436.Do(func() {
		cache_Data_Maybe_maybe__1732740436 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__1732740436(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__1732740436
}

var cache_Data_Maybe_maybe__3234875316 gopurs_runtime.Value
var once_Data_Maybe_maybe__3234875316 sync.Once
func Get_Data_Maybe_maybe__3234875316() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3234875316.Do(func() {
		cache_Data_Maybe_maybe__3234875316 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__3234875316(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__3234875316
}

var cache_Data_Maybe_maybe__4043979444 gopurs_runtime.Value
var once_Data_Maybe_maybe__4043979444 sync.Once
func Get_Data_Maybe_maybe__4043979444() gopurs_runtime.Value {
	once_Data_Maybe_maybe__4043979444.Do(func() {
		cache_Data_Maybe_maybe__4043979444 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__4043979444(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__4043979444
}

var cache_Data_Maybe_maybe__1061305652 gopurs_runtime.Value
var once_Data_Maybe_maybe__1061305652 sync.Once
func Get_Data_Maybe_maybe__1061305652() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1061305652.Do(func() {
		cache_Data_Maybe_maybe__1061305652 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__1061305652(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__1061305652
}

var cache_Data_Maybe_maybe__3931678292 gopurs_runtime.Value
var once_Data_Maybe_maybe__3931678292 sync.Once
func Get_Data_Maybe_maybe__3931678292() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3931678292.Do(func() {
		cache_Data_Maybe_maybe__3931678292 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__3931678292(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__3931678292
}

var cache_Data_Maybe_maybe__2925953714 gopurs_runtime.Value
var once_Data_Maybe_maybe__2925953714 sync.Once
func Get_Data_Maybe_maybe__2925953714() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2925953714.Do(func() {
		cache_Data_Maybe_maybe__2925953714 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Maybe_maybe__2925953714(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__2925953714
}

var cache_Data_Maybe_maybe__1408653394 gopurs_runtime.Value
var once_Data_Maybe_maybe__1408653394 sync.Once
func Get_Data_Maybe_maybe__1408653394() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1408653394.Do(func() {
		cache_Data_Maybe_maybe__1408653394 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Maybe_maybe__1408653394(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__1408653394
}

var cache_Data_Maybe_maybe__727024722 gopurs_runtime.Value
var once_Data_Maybe_maybe__727024722 sync.Once
func Get_Data_Maybe_maybe__727024722() gopurs_runtime.Value {
	once_Data_Maybe_maybe__727024722.Do(func() {
		cache_Data_Maybe_maybe__727024722 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Maybe_maybe__727024722(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))
})
	})
	return cache_Data_Maybe_maybe__727024722
}

var cache_Data_Maybe_maybe__2340146595 gopurs_runtime.Value
var once_Data_Maybe_maybe__2340146595 sync.Once
func Get_Data_Maybe_maybe__2340146595() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2340146595.Do(func() {
		cache_Data_Maybe_maybe__2340146595 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__2340146595(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__2340146595
}

var cache_Data_Maybe_maybe__47782440 gopurs_runtime.Value
var once_Data_Maybe_maybe__47782440 sync.Once
func Get_Data_Maybe_maybe__47782440() gopurs_runtime.Value {
	once_Data_Maybe_maybe__47782440.Do(func() {
		cache_Data_Maybe_maybe__47782440 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_maybe__47782440(uint32(v_0_box.IntVal), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_maybe__47782440
}

var cache_Data_Maybe_maybe__2472412892 gopurs_runtime.Value
var once_Data_Maybe_maybe__2472412892 sync.Once
func Get_Data_Maybe_maybe__2472412892() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2472412892.Do(func() {
		cache_Data_Maybe_maybe__2472412892 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_maybe__2472412892(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))}
})
	})
	return cache_Data_Maybe_maybe__2472412892
}

var cache_Data_Maybe_maybe__4159800284 gopurs_runtime.Value
var once_Data_Maybe_maybe__4159800284 sync.Once
func Get_Data_Maybe_maybe__4159800284() gopurs_runtime.Value {
	once_Data_Maybe_maybe__4159800284.Do(func() {
		cache_Data_Maybe_maybe__4159800284 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__4159800284(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__4159800284
}

var cache_Data_Maybe_maybe__1845466785 gopurs_runtime.Value
var once_Data_Maybe_maybe__1845466785 sync.Once
func Get_Data_Maybe_maybe__1845466785() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1845466785.Do(func() {
		cache_Data_Maybe_maybe__1845466785 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__1845466785(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__1845466785
}

var cache_Data_Maybe_maybe__1246883617 gopurs_runtime.Value
var once_Data_Maybe_maybe__1246883617 sync.Once
func Get_Data_Maybe_maybe__1246883617() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1246883617.Do(func() {
		cache_Data_Maybe_maybe__1246883617 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_maybe__1246883617(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))}
})
	})
	return cache_Data_Maybe_maybe__1246883617
}

var cache_Data_Maybe_maybe__2251533876 gopurs_runtime.Value
var once_Data_Maybe_maybe__2251533876 sync.Once
func Get_Data_Maybe_maybe__2251533876() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2251533876.Do(func() {
		cache_Data_Maybe_maybe__2251533876 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_maybe__2251533876(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))}
})
	})
	return cache_Data_Maybe_maybe__2251533876
}

var cache_Data_Maybe_maybe__563787205 gopurs_runtime.Value
var once_Data_Maybe_maybe__563787205 sync.Once
func Get_Data_Maybe_maybe__563787205() gopurs_runtime.Value {
	once_Data_Maybe_maybe__563787205.Do(func() {
		cache_Data_Maybe_maybe__563787205 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_maybe__563787205(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))}
})
	})
	return cache_Data_Maybe_maybe__563787205
}

var cache_Data_Maybe_maybe__316004085 gopurs_runtime.Value
var once_Data_Maybe_maybe__316004085 sync.Once
func Get_Data_Maybe_maybe__316004085() gopurs_runtime.Value {
	once_Data_Maybe_maybe__316004085.Do(func() {
		cache_Data_Maybe_maybe__316004085 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_maybe__316004085(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))}
})
	})
	return cache_Data_Maybe_maybe__316004085
}

var cache_Data_Maybe_maybe__1936890643 gopurs_runtime.Value
var once_Data_Maybe_maybe__1936890643 sync.Once
func Get_Data_Maybe_maybe__1936890643() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1936890643.Do(func() {
		cache_Data_Maybe_maybe__1936890643 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_maybe__1936890643(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))}
})
	})
	return cache_Data_Maybe_maybe__1936890643
}

var cache_Data_Maybe_maybe__2748933053 gopurs_runtime.Value
var once_Data_Maybe_maybe__2748933053 sync.Once
func Get_Data_Maybe_maybe__2748933053() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2748933053.Do(func() {
		cache_Data_Maybe_maybe__2748933053 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_maybe__2748933053(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))}
})
	})
	return cache_Data_Maybe_maybe__2748933053
}

var cache_Data_Maybe_maybe__137256317 gopurs_runtime.Value
var once_Data_Maybe_maybe__137256317 sync.Once
func Get_Data_Maybe_maybe__137256317() gopurs_runtime.Value {
	once_Data_Maybe_maybe__137256317.Do(func() {
		cache_Data_Maybe_maybe__137256317 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_maybe__137256317(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box)))}
})
	})
	return cache_Data_Maybe_maybe__137256317
}

var cache_Data_Maybe_maybe__3305810139 gopurs_runtime.Value
var once_Data_Maybe_maybe__3305810139 sync.Once
func Get_Data_Maybe_maybe__3305810139() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3305810139.Do(func() {
		cache_Data_Maybe_maybe__3305810139 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe__3305810139(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe__3305810139
}

var cache_Data_Maybe_maybe_prime__2328206764 gopurs_runtime.Value
var once_Data_Maybe_maybe_prime__2328206764 sync.Once
func Get_Data_Maybe_maybe_prime__2328206764() gopurs_runtime.Value {
	once_Data_Maybe_maybe_prime__2328206764.Do(func() {
		cache_Data_Maybe_maybe_prime__2328206764 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe_prime__2328206764(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe_prime__2328206764
}

var cache_Data_Maybe_maybe_prime__4209968548 gopurs_runtime.Value
var once_Data_Maybe_maybe_prime__4209968548 sync.Once
func Get_Data_Maybe_maybe_prime__4209968548() gopurs_runtime.Value {
	once_Data_Maybe_maybe_prime__4209968548.Do(func() {
		cache_Data_Maybe_maybe_prime__4209968548 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe_prime__4209968548(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe_prime__4209968548
}

type Constructor_Data_Maybe_Nothing struct {
	Rc uint32
}


type Constructor_Data_Maybe_Just struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func Call_Data_Maybe_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Maybe_showMaybe(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = (("(Just ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Maybe_Just)(v_1.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t0 = "Nothing"
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
})})}
}

func Call_Data_Maybe_semigroupMaybe(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_2)
goto end_branch_0
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_1)
goto end_branch_0
} else {

}
}
{
if ((v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil)) && ((v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil)) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Maybe_Just)(v_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(v1_2.UnsafePtr).V0)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
})})}
}

func Call_Data_Maybe_optional(dictAlt_0_loop *Constructor_Control_Alt_Alt) gopurs_runtime.Value {
var dictAlt_0 *Constructor_Control_Alt_Alt = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlt_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictAlt_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Just(), a_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}))
})
})
}

func Call_Data_Maybe_monoidMaybe(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): semigroupMaybe1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupMaybe1_1_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_2)
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_1)
goto end_branch_1
} else {

}
}
{
if ((v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil)) && ((v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil)) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Maybe_Just)(v_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(v1_2.UnsafePtr).V0)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})
})}
_ = semigroupMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupMaybe1_1_0)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})}
}

func Call_Data_Maybe_maybe_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_isNothing(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isJust(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_fromMaybe_prime(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply2(Get_Data_Maybe_maybe_prime(), a_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Maybe_fromMaybe(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1 == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = (v2_1).V0
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

func Call_Data_Maybe_fromJust(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_Maybe_Just = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (v_1 != nil) {
__t0 = (v_1).V0
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

func Call_Data_Maybe_eqMaybe(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t0 bool
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t1 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Maybe_Just)(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_2.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
})})}
}

func Call_Data_Maybe_ordMaybe(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqMaybe1_1_0 -> *Constructor_Data_Eq_Eq
eqMaybe1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
var __t2 bool
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil)) && ((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr != nil)) {
__t3 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_Maybe_Just)(x_2.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
})))
_ = eqMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMaybe1_1_0)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 uint32
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
var __t4 uint32
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil)) && ((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr != nil)) {
__t5 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Maybe_Just)(x_2.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_3.UnsafePtr).V0).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})
})})}
}

func Call_Data_Maybe_boundedMaybe(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): eqMaybe1_2_2 -> *Constructor_Data_Eq_Eq
eqMaybe1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 bool
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
var __t4 bool
{
if (y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr == nil) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if ((x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil)) && ((y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr != nil)) {
__t5 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "eq"), (*Constructor_Data_Maybe_Just)(x_3.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})
})))
_ = eqMaybe1_2_2
// TAST (Let): ordMaybe1_1_0 -> *Constructor_Data_Ord_Ord
ordMaybe1_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMaybe1_2_2)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 uint32
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
var __t6 uint32
{
if (y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr == nil) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
if (y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr == nil) {
__t7 = 380165415
goto end_branch_7
} else {

}
}
{
if ((x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil)) && ((y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr != nil)) {
__t7 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compare"), (*Constructor_Data_Maybe_Just)(x_3.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_4.UnsafePtr).V0).IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
})
})}
_ = ordMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bounded_Bounded{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordMaybe1_1_0)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet(dictBounded_0, "top")})}})}
}

func Call_Data_Maybe_semiringMaybe(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semiring_Semiring{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_2)
goto end_branch_0
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_1)
goto end_branch_0
} else {

}
}
{
if ((v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil)) && ((v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil)) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*Constructor_Data_Maybe_Just)(v_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(v1_2.UnsafePtr).V0)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
var __t_tag_2 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](x_1)
if (__t_tag_2 != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), (*Constructor_Data_Maybe_Just)(x_1.UnsafePtr).V0)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
// TAST (Let): __local_var_3_1 -> *Constructor_Data_Maybe_Just
var __local_var_3_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})
var __t6 *Constructor_Data_Maybe_Just
{
if (__local_var_3_1 != nil) {
var __t5 *Constructor_Data_Maybe_Just
{
var __t_tag_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](y_2)
if (__t_tag_4 != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_3_1).V0, (*Constructor_Data_Maybe_Just)(y_2.UnsafePtr).V0)}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (__local_var_3_1 == nil) {
__t6 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet(dictSemiring_0, "one")})}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})}
}

func Call_Data_Maybe_fromJust__2181618881(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Maybe_Just) int64 {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_Maybe_Just = v_1_loop
_ = v_1
var __t0 int64
{
if (v_1 != nil) {
__t0 = (v_1).V0.IntVal
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return gopurs_runtime.Int(__t0).IntVal
}

func Call_Data_Maybe_fromJust__1577979644(v_0_loop *Constructor_Data_Maybe_Just) int64 {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 != nil) {
__t0 = (v_0).V0.IntVal
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__4121089788(v_0_loop *Constructor_Data_Maybe_Just) string {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 string
{
if (v_0 != nil) {
__t0 = (v_0).V0.StrVal()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_Maybe_Just = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (v_1 != nil) {
__t0 = (v_1).V0
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

func Call_Data_Maybe_fromJust__911089788(v_0_loop *Constructor_Data_Maybe_Just) []int64 {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 []int64
{
if (v_0 != nil) {
__t0 = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)((v_0).V0.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
goto end_branch_0
} else {

}
}
{
__t0 = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__3897574428(v_0_loop *Constructor_Data_Maybe_Just) []string {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 []string
{
if (v_0 != nil) {
__t0 = func() []string {
					arr := *(*[]gopurs_runtime.Value)((v_0).V0.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
goto end_branch_0
} else {

}
}
{
__t0 = func() []string {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__1478027324(v_0_loop *Constructor_Data_Maybe_Just) []gopurs_runtime.Value {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 []gopurs_runtime.Value
{
if (v_0 != nil) {
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((v_0).V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__4142563260(v_0_loop *Constructor_Data_Maybe_Just) uint32 {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 uint32
{
if (v_0 != nil) {
__t0 = uint32((v_0).V0.IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__3809843644(v_0_loop *Constructor_Data_Maybe_Just) uint32 {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 uint32
{
if (v_0 != nil) {
__t0 = uint32((v_0).V0.IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__965748316(v_0_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Date_Date {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Date_Date
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((v_0).V0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__755886620(v_0_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Time_Time {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Time_Time
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time]((v_0).V0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromMaybe__1972796397(a_0_loop int64, v2_1_loop *Constructor_Data_Maybe_Just) int64 {
var a_0 int64 = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1 == nil) {
__t0 = gopurs_runtime.Int(a_0)
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = (v2_1).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_Data_Maybe_fromMaybe__430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1 == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = (v2_1).V0
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

func Call_Data_Maybe_fromMaybe__656947263(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1 == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = (v2_1).V0
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

func Call_Data_Maybe_fromMaybe__18840980(a_0_loop uint32, v2_1_loop *Constructor_Data_Maybe_Just) uint32 {
var a_0 uint32 = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(a_0), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = (v2_1).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(__t0.IntVal)
}

func Call_Data_Maybe_fromMaybe__737056608(a_0_loop *Constructor_Data_Date_Date, v2_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Date_Date {
var a_0 *Constructor_Data_Date_Date = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just = v2_1_loop
_ = v2_1
var __t0 *Constructor_Data_Date_Date
{
if (v2_1 == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((v2_1).V0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromMaybe__2067158953(a_0_loop *Constructor_Data_Maybe_Just, v2_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just = v2_1_loop
_ = v2_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v2_1 == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((v2_1).V0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_isJust__2514352589(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isJust__2475527019(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isJust__2591355336(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isJust__1358705270(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isJust__4165351782(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isJust__4206805139(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isNothing__2514352589(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isNothing__2591355336(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isNothing__1401305026(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isNothing__1358705270(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isNothing__4206805139(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isNothing__2787066607(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_isNothing__323776123(v2_0_loop *Constructor_Data_Maybe_Just) bool {
var v2_0 *Constructor_Data_Maybe_Just = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
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

func Call_Data_Maybe_maybe__919206801(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) int64 {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 int64
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0.IntVal)).IntVal
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__3735358641(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) int64 {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 int64
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0).IntVal
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__3078346790(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) bool {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0.IntVal)).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__487722278(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) bool {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, gopurs_runtime.Float((v2_2).V0.FloatVal())).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__1510464358(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) bool {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, gopurs_runtime.Str((v2_2).V0.StrVal())).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__1594528518(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) bool {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, (v2_2).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__3407128198(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) bool {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((v2_2).V0.IntVal)), UnsafePtr: nil}).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__2158452262(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) bool {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((v2_2).V0))}).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__2641488518(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) bool {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = (gopurs_runtime.Apply(v1_1, (v2_2).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__3718989812(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0.IntVal))
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

func Call_Data_Maybe_maybe__316277428(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Float((v2_2).V0.FloatVal()))
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

func Call_Data_Maybe_maybe__1647364852(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Str((v2_2).V0.StrVal()))
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

func Call_Data_Maybe_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe__1726410932(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe__1732740436(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((v2_2).V0.IntVal)), UnsafePtr: nil})
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

func Call_Data_Maybe_maybe__3234875316(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((v2_2).V0))})
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

func Call_Data_Maybe_maybe__4043979444(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((v2_2).V0))})
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

func Call_Data_Maybe_maybe__1061305652(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((v2_2).V0))})
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

func Call_Data_Maybe_maybe__3931678292(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe__2925953714(v_0_loop []gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) []gopurs_runtime.Value {
var v_0 []gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 []gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0.IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__1408653394(v_0_loop []gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) []gopurs_runtime.Value {
var v_0 []gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 []gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(v1_1, (v2_2).V0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__727024722(v_0_loop []gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) []gopurs_runtime.Value {
var v_0 []gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 []gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(v1_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((v2_2).V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__2340146595(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe__47782440(v_0_loop uint32, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) uint32 {
var v_0 uint32 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 uint32
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = uint32(gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((v2_2).V0.IntVal)), UnsafePtr: nil}).IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__2472412892(v_0_loop *Constructor_Data_Date_Date, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Date_Date {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_Date_Date
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((v2_2).V0))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__4159800284(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe__1845466785(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe__1246883617(v_0_loop *Constructor_Data_List_Types_Cons, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_List_Types_Cons
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v2_2).V0))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__2251533876(v_0_loop *Constructor_Data_Map_Internal_Node, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Map_Internal_Node {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_Map_Internal_Node
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(v1_1, (v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__563787205(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_Maybe_Just
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__316004085(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_Maybe_Just
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((v2_2).V0))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__1936890643(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_Maybe_Just
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__2748933053(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_Maybe_Just
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__137256317(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 *Constructor_Data_Maybe_Just
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__3305810139(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe_prime__2328206764(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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

func Call_Data_Maybe_maybe_prime__4209968548(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
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


