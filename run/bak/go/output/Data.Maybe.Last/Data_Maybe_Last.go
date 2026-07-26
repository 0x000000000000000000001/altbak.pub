package Data_Maybe_Last

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_Last gopurs_runtime.Value
var once_Last sync.Once
func Get_Last() gopurs_runtime.Value {
	once_Last.Do(func() {
		cache_Last = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Last(x_0_box)
})
	})
	return cache_Last
}

var cache_showLast gopurs_runtime.Value
var once_showLast sync.Once
func Get_showLast() gopurs_runtime.Value {
	once_showLast.Do(func() {
		cache_showLast = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showLast(dictShow_0_box)
})
	})
	return cache_showLast
}

var cache_semigroupLast gopurs_runtime.Value
var once_semigroupLast sync.Once
func Get_semigroupLast() gopurs_runtime.Value {
	once_semigroupLast.Do(func() {
		cache_semigroupLast = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return cache_semigroupLast
}

var cache_ordLast gopurs_runtime.Value
var once_ordLast sync.Once
func Get_ordLast() gopurs_runtime.Value {
	once_ordLast.Do(func() {
		cache_ordLast = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordLast(dictOrd_0_box)
})
	})
	return cache_ordLast
}

var cache_ord1Last gopurs_runtime.Value
var once_ord1Last sync.Once
func Get_ord1Last() gopurs_runtime.Value {
	once_ord1Last.Do(func() {
		cache_ord1Last = pkg_Data_Maybe.Get_ord1Maybe()
	})
	return cache_ord1Last
}

var cache_newtypeLast gopurs_runtime.Value
var once_newtypeLast sync.Once
func Get_newtypeLast() gopurs_runtime.Value {
	once_newtypeLast.Do(func() {
		cache_newtypeLast = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeLast
}

var cache_monoidLast gopurs_runtime.Value
var once_monoidLast sync.Once
func Get_monoidLast() gopurs_runtime.Value {
	once_monoidLast.Do(func() {
		cache_monoidLast = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupLast()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
	})
	return cache_monoidLast
}

var cache_monadLast gopurs_runtime.Value
var once_monadLast sync.Once
func Get_monadLast() gopurs_runtime.Value {
	once_monadLast.Do(func() {
		cache_monadLast = pkg_Data_Maybe.Get_monadMaybe()
	})
	return cache_monadLast
}

var cache_invariantLast gopurs_runtime.Value
var once_invariantLast sync.Once
func Get_invariantLast() gopurs_runtime.Value {
	once_invariantLast.Do(func() {
		cache_invariantLast = pkg_Data_Maybe.Get_invariantMaybe()
	})
	return cache_invariantLast
}

var cache_functorLast gopurs_runtime.Value
var once_functorLast sync.Once
func Get_functorLast() gopurs_runtime.Value {
	once_functorLast.Do(func() {
		cache_functorLast = pkg_Data_Maybe.Get_functorMaybe()
	})
	return cache_functorLast
}

var cache_extendLast gopurs_runtime.Value
var once_extendLast sync.Once
func Get_extendLast() gopurs_runtime.Value {
	once_extendLast.Do(func() {
		cache_extendLast = pkg_Data_Maybe.Get_extendMaybe()
	})
	return cache_extendLast
}

var cache_eqLast gopurs_runtime.Value
var once_eqLast sync.Once
func Get_eqLast() gopurs_runtime.Value {
	once_eqLast.Do(func() {
		cache_eqLast = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqLast(dictEq_0_box)
})
	})
	return cache_eqLast
}

var cache_eq1Last gopurs_runtime.Value
var once_eq1Last sync.Once
func Get_eq1Last() gopurs_runtime.Value {
	once_eq1Last.Do(func() {
		cache_eq1Last = pkg_Data_Maybe.Get_eq1Maybe()
	})
	return cache_eq1Last
}

var cache_boundedLast gopurs_runtime.Value
var once_boundedLast sync.Once
func Get_boundedLast() gopurs_runtime.Value {
	once_boundedLast.Do(func() {
		cache_boundedLast = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedLast(dictBounded_0_box)
})
	})
	return cache_boundedLast
}

var cache_bindLast gopurs_runtime.Value
var once_bindLast sync.Once
func Get_bindLast() gopurs_runtime.Value {
	once_bindLast.Do(func() {
		cache_bindLast = pkg_Data_Maybe.Get_bindMaybe()
	})
	return cache_bindLast
}

var cache_applyLast gopurs_runtime.Value
var once_applyLast sync.Once
func Get_applyLast() gopurs_runtime.Value {
	once_applyLast.Do(func() {
		cache_applyLast = pkg_Data_Maybe.Get_applyMaybe()
	})
	return cache_applyLast
}

var cache_applicativeLast gopurs_runtime.Value
var once_applicativeLast sync.Once
func Get_applicativeLast() gopurs_runtime.Value {
	once_applicativeLast.Do(func() {
		cache_applicativeLast = pkg_Data_Maybe.Get_applicativeMaybe()
	})
	return cache_applicativeLast
}

var cache_altLast gopurs_runtime.Value
var once_altLast sync.Once
func Get_altLast() gopurs_runtime.Value {
	once_altLast.Do(func() {
		cache_altLast = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.RecordGet(Get_semigroupLast(), "append"))
	})
	return cache_altLast
}

var cache_plusLast gopurs_runtime.Value
var once_plusLast sync.Once
func Get_plusLast() gopurs_runtime.Value {
	once_plusLast.Do(func() {
		cache_plusLast = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altLast()
}), gopurs_runtime.RecordGet(Get_monoidLast(), "mempty"))
	})
	return cache_plusLast
}

var cache_alternativeLast gopurs_runtime.Value
var once_alternativeLast sync.Once
func Get_alternativeLast() gopurs_runtime.Value {
	once_alternativeLast.Do(func() {
		cache_alternativeLast = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusLast()
}))
	})
	return cache_alternativeLast
}

func Call_Last(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showLast(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Last "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Maybe.Get_showMaybe(), dictShow_0), "show"), v_1), gopurs_runtime.Str(")")))
}))
}

func Call_ordLast(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqMaybe1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(((x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil)) && (((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_2:
return __t2
}))
_ = eqMaybe1_2_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_2_1
}), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
var __t4 gopurs_runtime.Value
{
if (y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if (y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if ((x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil)) && ((y_4.Type == 9 && y_4.IntVal == 930809136 && y_4.UnsafePtr != nil)) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_4.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
}

func Call_eqLast(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && (((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_0:
return __t0
}))
}

func Call_boundedLast(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Apply(pkg_Data_Maybe.Get_boundedMaybe(), dictBounded_0)
}


