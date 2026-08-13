package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Maybe_Last_Last gopurs_runtime.Value
var once_Data_Maybe_Last_Last sync.Once
func Get_Data_Maybe_Last_Last() gopurs_runtime.Value {
	once_Data_Maybe_Last_Last.Do(func() {
		cache_Data_Maybe_Last_Last = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_Last_Last(x_0_box)
})
	})
	return cache_Data_Maybe_Last_Last
}

var cache_Data_Maybe_Last_showLast gopurs_runtime.Value
var once_Data_Maybe_Last_showLast sync.Once
func Get_Data_Maybe_Last_showLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_showLast.Do(func() {
		cache_Data_Maybe_Last_showLast = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_Last_showLast(dictShow_0_box)
})
	})
	return cache_Data_Maybe_Last_showLast
}

var cache_Data_Maybe_Last_semigroupLast gopurs_runtime.Value
var once_Data_Maybe_Last_semigroupLast sync.Once
func Get_Data_Maybe_Last_semigroupLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_semigroupLast.Do(func() {
		cache_Data_Maybe_Last_semigroupLast = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1)
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0)
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
	return cache_Data_Maybe_Last_semigroupLast
}

var cache_Data_Maybe_Last_ordLast gopurs_runtime.Value
var once_Data_Maybe_Last_ordLast sync.Once
func Get_Data_Maybe_Last_ordLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_ordLast.Do(func() {
		cache_Data_Maybe_Last_ordLast = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_Last_ordLast(dictOrd_0_box)
})
	})
	return cache_Data_Maybe_Last_ordLast
}

var cache_Data_Maybe_Last_ord1Last gopurs_runtime.Value
var once_Data_Maybe_Last_ord1Last sync.Once
func Get_Data_Maybe_Last_ord1Last() gopurs_runtime.Value {
	once_Data_Maybe_Last_ord1Last.Do(func() {
		cache_Data_Maybe_Last_ord1Last = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1](Get_Data_Maybe_ord1Maybe()))}
	})
	return cache_Data_Maybe_Last_ord1Last
}

var cache_Data_Maybe_Last_newtypeLast gopurs_runtime.Value
var once_Data_Maybe_Last_newtypeLast sync.Once
func Get_Data_Maybe_Last_newtypeLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_newtypeLast.Do(func() {
		cache_Data_Maybe_Last_newtypeLast = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_Maybe_Last_newtypeLast
}

var cache_Data_Maybe_Last_monoidLast gopurs_runtime.Value
var once_Data_Maybe_Last_monoidLast sync.Once
func Get_Data_Maybe_Last_monoidLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_monoidLast.Do(func() {
		cache_Data_Maybe_Last_monoidLast = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Maybe_Last_semigroupLast()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})}
	})
	return cache_Data_Maybe_Last_monoidLast
}

var cache_Data_Maybe_Last_monadLast gopurs_runtime.Value
var once_Data_Maybe_Last_monadLast sync.Once
func Get_Data_Maybe_Last_monadLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_monadLast.Do(func() {
		cache_Data_Maybe_Last_monadLast = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Data_Maybe_monadMaybe()))}
	})
	return cache_Data_Maybe_Last_monadLast
}

var cache_Data_Maybe_Last_invariantLast gopurs_runtime.Value
var once_Data_Maybe_Last_invariantLast sync.Once
func Get_Data_Maybe_Last_invariantLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_invariantLast.Do(func() {
		cache_Data_Maybe_Last_invariantLast = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant](Get_Data_Maybe_invariantMaybe()))}
	})
	return cache_Data_Maybe_Last_invariantLast
}

var cache_Data_Maybe_Last_functorLast gopurs_runtime.Value
var once_Data_Maybe_Last_functorLast sync.Once
func Get_Data_Maybe_Last_functorLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_functorLast.Do(func() {
		cache_Data_Maybe_Last_functorLast = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Maybe_functorMaybe()))}
	})
	return cache_Data_Maybe_Last_functorLast
}

var cache_Data_Maybe_Last_extendLast gopurs_runtime.Value
var once_Data_Maybe_Last_extendLast sync.Once
func Get_Data_Maybe_Last_extendLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_extendLast.Do(func() {
		cache_Data_Maybe_Last_extendLast = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](Get_Data_Maybe_extendMaybe()))}
	})
	return cache_Data_Maybe_Last_extendLast
}

var cache_Data_Maybe_Last_eqLast gopurs_runtime.Value
var once_Data_Maybe_Last_eqLast sync.Once
func Get_Data_Maybe_Last_eqLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_eqLast.Do(func() {
		cache_Data_Maybe_Last_eqLast = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_Last_eqLast(dictEq_0_box)
})
	})
	return cache_Data_Maybe_Last_eqLast
}

var cache_Data_Maybe_Last_eq1Last gopurs_runtime.Value
var once_Data_Maybe_Last_eq1Last sync.Once
func Get_Data_Maybe_Last_eq1Last() gopurs_runtime.Value {
	once_Data_Maybe_Last_eq1Last.Do(func() {
		cache_Data_Maybe_Last_eq1Last = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Maybe_eq1Maybe()))}
	})
	return cache_Data_Maybe_Last_eq1Last
}

var cache_Data_Maybe_Last_boundedLast gopurs_runtime.Value
var once_Data_Maybe_Last_boundedLast sync.Once
func Get_Data_Maybe_Last_boundedLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_boundedLast.Do(func() {
		cache_Data_Maybe_Last_boundedLast = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_Last_boundedLast(dictBounded_0_box)
})
	})
	return cache_Data_Maybe_Last_boundedLast
}

var cache_Data_Maybe_Last_bindLast gopurs_runtime.Value
var once_Data_Maybe_Last_bindLast sync.Once
func Get_Data_Maybe_Last_bindLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_bindLast.Do(func() {
		cache_Data_Maybe_Last_bindLast = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_Maybe_bindMaybe()))}
	})
	return cache_Data_Maybe_Last_bindLast
}

var cache_Data_Maybe_Last_applyLast gopurs_runtime.Value
var once_Data_Maybe_Last_applyLast sync.Once
func Get_Data_Maybe_Last_applyLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_applyLast.Do(func() {
		cache_Data_Maybe_Last_applyLast = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Maybe_applyMaybe()))}
	})
	return cache_Data_Maybe_Last_applyLast
}

var cache_Data_Maybe_Last_applicativeLast gopurs_runtime.Value
var once_Data_Maybe_Last_applicativeLast sync.Once
func Get_Data_Maybe_Last_applicativeLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_applicativeLast.Do(func() {
		cache_Data_Maybe_Last_applicativeLast = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Maybe_applicativeMaybe()))}
	})
	return cache_Data_Maybe_Last_applicativeLast
}

var cache_Data_Maybe_Last_altLast gopurs_runtime.Value
var once_Data_Maybe_Last_altLast sync.Once
func Get_Data_Maybe_Last_altLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_altLast.Do(func() {
		cache_Data_Maybe_Last_altLast = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Maybe_functorMaybe()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1)
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0)
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
	return cache_Data_Maybe_Last_altLast
}

var cache_Data_Maybe_Last_plusLast gopurs_runtime.Value
var once_Data_Maybe_Last_plusLast sync.Once
func Get_Data_Maybe_Last_plusLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_plusLast.Do(func() {
		cache_Data_Maybe_Last_plusLast = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Data_Maybe_Last_altLast()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})}
	})
	return cache_Data_Maybe_Last_plusLast
}

var cache_Data_Maybe_Last_alternativeLast gopurs_runtime.Value
var once_Data_Maybe_Last_alternativeLast sync.Once
func Get_Data_Maybe_Last_alternativeLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_alternativeLast.Do(func() {
		cache_Data_Maybe_Last_alternativeLast = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Maybe_applicativeMaybe()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_Maybe_Last_plusLast()))}
})})}
	})
	return cache_Data_Maybe_Last_alternativeLast
}

var cache_Data_Maybe_Last_altLast__1966529924 gopurs_runtime.Value
var once_Data_Maybe_Last_altLast__1966529924 sync.Once
func Get_Data_Maybe_Last_altLast__1966529924() gopurs_runtime.Value {
	once_Data_Maybe_Last_altLast__1966529924.Do(func() {
		cache_Data_Maybe_Last_altLast__1966529924 = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Maybe_functorMaybe()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1)
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0)
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
	return cache_Data_Maybe_Last_altLast__1966529924
}

var cache_Data_Maybe_Last_applicativeLast__2906236036 gopurs_runtime.Value
var once_Data_Maybe_Last_applicativeLast__2906236036 sync.Once
func Get_Data_Maybe_Last_applicativeLast__2906236036() gopurs_runtime.Value {
	once_Data_Maybe_Last_applicativeLast__2906236036.Do(func() {
		cache_Data_Maybe_Last_applicativeLast__2906236036 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Maybe_applicativeMaybe()))}
	})
	return cache_Data_Maybe_Last_applicativeLast__2906236036
}

var cache_Data_Maybe_Last_functorLast__930783699 gopurs_runtime.Value
var once_Data_Maybe_Last_functorLast__930783699 sync.Once
func Get_Data_Maybe_Last_functorLast__930783699() gopurs_runtime.Value {
	once_Data_Maybe_Last_functorLast__930783699.Do(func() {
		cache_Data_Maybe_Last_functorLast__930783699 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Maybe_functorMaybe()))}
	})
	return cache_Data_Maybe_Last_functorLast__930783699
}

var cache_Data_Maybe_Last_plusLast__918036452 gopurs_runtime.Value
var once_Data_Maybe_Last_plusLast__918036452 sync.Once
func Get_Data_Maybe_Last_plusLast__918036452() gopurs_runtime.Value {
	once_Data_Maybe_Last_plusLast__918036452.Do(func() {
		cache_Data_Maybe_Last_plusLast__918036452 = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Data_Maybe_Last_altLast()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})}
	})
	return cache_Data_Maybe_Last_plusLast__918036452
}

var cache_Data_Maybe_Last_semigroupLast__2956196563 gopurs_runtime.Value
var once_Data_Maybe_Last_semigroupLast__2956196563 sync.Once
func Get_Data_Maybe_Last_semigroupLast__2956196563() gopurs_runtime.Value {
	once_Data_Maybe_Last_semigroupLast__2956196563.Do(func() {
		cache_Data_Maybe_Last_semigroupLast__2956196563 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1)
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0)
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
	return cache_Data_Maybe_Last_semigroupLast__2956196563
}

func Call_Data_Maybe_Last_Last(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Maybe_Last_showLast(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showMaybe_1_0 -> *Constructor_Data_Show_Show
showMaybe_1_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 string
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t1 = (("(Just ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Maybe_Just)(v_1.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t1 = "Nothing"
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_1:
return gopurs_runtime.Str(__t1)
})}
_ = showMaybe_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Last ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showMaybe_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2))}).StrVal())) + (")"))
})})}
}

func Call_Data_Maybe_Last_ordLast(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Maybe_Last_eqLast(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Maybe_Last_boundedLast(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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


