package Data_Maybe_First

import (
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_First gopurs_runtime.Value
var once_First sync.Once
func Get_First() gopurs_runtime.Value {
	once_First.Do(func() {
		cache_First = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_First(x_0_box)
})
	})
	return cache_First
}

var cache_showFirst gopurs_runtime.Value
var once_showFirst sync.Once
func Get_showFirst() gopurs_runtime.Value {
	once_showFirst.Do(func() {
		cache_showFirst = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showFirst(dictShow_0_box)
})
	})
	return cache_showFirst
}

var cache_semigroupFirst gopurs_runtime.Value
var once_semigroupFirst sync.Once
func Get_semigroupFirst() gopurs_runtime.Value {
	once_semigroupFirst.Do(func() {
		cache_semigroupFirst = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_semigroupFirst
}

var cache_semigroupFirst__gopurs_runtime_Value_3179391308 gopurs_runtime.Value
var once_semigroupFirst__gopurs_runtime_Value_3179391308 sync.Once
func Get_semigroupFirst__gopurs_runtime_Value_3179391308() gopurs_runtime.Value {
	once_semigroupFirst__gopurs_runtime_Value_3179391308.Do(func() {
		cache_semigroupFirst__gopurs_runtime_Value_3179391308 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_semigroupFirst__gopurs_runtime_Value_3179391308
}

var cache_ordFirst gopurs_runtime.Value
var once_ordFirst sync.Once
func Get_ordFirst() gopurs_runtime.Value {
	once_ordFirst.Do(func() {
		cache_ordFirst = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordFirst(dictOrd_0_box)
})
	})
	return cache_ordFirst
}

var cache_ord1First gopurs_runtime.Value
var once_ord1First sync.Once
func Get_ord1First() gopurs_runtime.Value {
	once_ord1First.Do(func() {
		cache_ord1First = pkg_Data_Maybe.Get_ord1Maybe()
	})
	return cache_ord1First
}

var cache_newtypeFirst gopurs_runtime.Value
var once_newtypeFirst sync.Once
func Get_newtypeFirst() gopurs_runtime.Value {
	once_newtypeFirst.Do(func() {
		cache_newtypeFirst = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeFirst
}

var cache_monoidFirst gopurs_runtime.Value
var once_monoidFirst sync.Once
func Get_monoidFirst() gopurs_runtime.Value {
	once_monoidFirst.Do(func() {
		cache_monoidFirst = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupFirst()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))})
	})
	return cache_monoidFirst
}

var cache_monoidFirst__ptrData_Monoid_Constructor_Monoid_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___4133827187 gopurs_runtime.Value
var once_monoidFirst__ptrData_Monoid_Constructor_Monoid_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___4133827187 sync.Once
func Get_monoidFirst__ptrData_Monoid_Constructor_Monoid_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___4133827187() gopurs_runtime.Value {
	once_monoidFirst__ptrData_Monoid_Constructor_Monoid_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___4133827187.Do(func() {
		cache_monoidFirst__ptrData_Monoid_Constructor_Monoid_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___4133827187 = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&pkg_Data_Monoid.Constructor_Monoid[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupFirst()
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})})}
	})
	return cache_monoidFirst__ptrData_Monoid_Constructor_Monoid_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___4133827187
}

var cache_monadFirst gopurs_runtime.Value
var once_monadFirst sync.Once
func Get_monadFirst() gopurs_runtime.Value {
	once_monadFirst.Do(func() {
		cache_monadFirst = pkg_Data_Maybe.Get_monadMaybe()
	})
	return cache_monadFirst
}

var cache_invariantFirst gopurs_runtime.Value
var once_invariantFirst sync.Once
func Get_invariantFirst() gopurs_runtime.Value {
	once_invariantFirst.Do(func() {
		cache_invariantFirst = pkg_Data_Maybe.Get_invariantMaybe()
	})
	return cache_invariantFirst
}

var cache_functorFirst gopurs_runtime.Value
var once_functorFirst sync.Once
func Get_functorFirst() gopurs_runtime.Value {
	once_functorFirst.Do(func() {
		cache_functorFirst = pkg_Data_Maybe.Get_functorMaybe()
	})
	return cache_functorFirst
}

var cache_functorFirst__gopurs_runtime_Value_2097654001 gopurs_runtime.Value
var once_functorFirst__gopurs_runtime_Value_2097654001 sync.Once
func Get_functorFirst__gopurs_runtime_Value_2097654001() gopurs_runtime.Value {
	once_functorFirst__gopurs_runtime_Value_2097654001.Do(func() {
		cache_functorFirst__gopurs_runtime_Value_2097654001 = pkg_Data_Maybe.Get_functorMaybe()
	})
	return cache_functorFirst__gopurs_runtime_Value_2097654001
}

var cache_extendFirst gopurs_runtime.Value
var once_extendFirst sync.Once
func Get_extendFirst() gopurs_runtime.Value {
	once_extendFirst.Do(func() {
		cache_extendFirst = pkg_Data_Maybe.Get_extendMaybe()
	})
	return cache_extendFirst
}

var cache_eqFirst gopurs_runtime.Value
var once_eqFirst sync.Once
func Get_eqFirst() gopurs_runtime.Value {
	once_eqFirst.Do(func() {
		cache_eqFirst = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqFirst(dictEq_0_box)
})
	})
	return cache_eqFirst
}

var cache_eq1First gopurs_runtime.Value
var once_eq1First sync.Once
func Get_eq1First() gopurs_runtime.Value {
	once_eq1First.Do(func() {
		cache_eq1First = pkg_Data_Maybe.Get_eq1Maybe()
	})
	return cache_eq1First
}

var cache_boundedFirst gopurs_runtime.Value
var once_boundedFirst sync.Once
func Get_boundedFirst() gopurs_runtime.Value {
	once_boundedFirst.Do(func() {
		cache_boundedFirst = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedFirst(dictBounded_0_box)
})
	})
	return cache_boundedFirst
}

var cache_bindFirst gopurs_runtime.Value
var once_bindFirst sync.Once
func Get_bindFirst() gopurs_runtime.Value {
	once_bindFirst.Do(func() {
		cache_bindFirst = pkg_Data_Maybe.Get_bindMaybe()
	})
	return cache_bindFirst
}

var cache_applyFirst gopurs_runtime.Value
var once_applyFirst sync.Once
func Get_applyFirst() gopurs_runtime.Value {
	once_applyFirst.Do(func() {
		cache_applyFirst = pkg_Data_Maybe.Get_applyMaybe()
	})
	return cache_applyFirst
}

var cache_applicativeFirst gopurs_runtime.Value
var once_applicativeFirst sync.Once
func Get_applicativeFirst() gopurs_runtime.Value {
	once_applicativeFirst.Do(func() {
		cache_applicativeFirst = pkg_Data_Maybe.Get_applicativeMaybe()
	})
	return cache_applicativeFirst
}

var cache_applicativeFirst__gopurs_runtime_Value_500933224 gopurs_runtime.Value
var once_applicativeFirst__gopurs_runtime_Value_500933224 sync.Once
func Get_applicativeFirst__gopurs_runtime_Value_500933224() gopurs_runtime.Value {
	once_applicativeFirst__gopurs_runtime_Value_500933224.Do(func() {
		cache_applicativeFirst__gopurs_runtime_Value_500933224 = pkg_Data_Maybe.Get_applicativeMaybe()
	})
	return cache_applicativeFirst__gopurs_runtime_Value_500933224
}

var cache_altFirst gopurs_runtime.Value
var once_altFirst sync.Once
func Get_altFirst() gopurs_runtime.Value {
	once_altFirst.Do(func() {
		cache_altFirst = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.RecordGet(Get_semigroupFirst(), "append"))
	})
	return cache_altFirst
}

var cache_altFirst__gopurs_runtime_Value_4201091523 gopurs_runtime.Value
var once_altFirst__gopurs_runtime_Value_4201091523 sync.Once
func Get_altFirst__gopurs_runtime_Value_4201091523() gopurs_runtime.Value {
	once_altFirst__gopurs_runtime_Value_4201091523.Do(func() {
		cache_altFirst__gopurs_runtime_Value_4201091523 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.RecordGet(Get_semigroupFirst(), "append"))
	})
	return cache_altFirst__gopurs_runtime_Value_4201091523
}

var cache_plusFirst gopurs_runtime.Value
var once_plusFirst sync.Once
func Get_plusFirst() gopurs_runtime.Value {
	once_plusFirst.Do(func() {
		cache_plusFirst = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altFirst()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(Get_monoidFirst(), "mempty")))})
	})
	return cache_plusFirst
}

var cache_plusFirst__gopurs_runtime_Value_400696082 gopurs_runtime.Value
var once_plusFirst__gopurs_runtime_Value_400696082 sync.Once
func Get_plusFirst__gopurs_runtime_Value_400696082() gopurs_runtime.Value {
	once_plusFirst__gopurs_runtime_Value_400696082.Do(func() {
		cache_plusFirst__gopurs_runtime_Value_400696082 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altFirst()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(Get_monoidFirst(), "mempty")))})
	})
	return cache_plusFirst__gopurs_runtime_Value_400696082
}

var cache_alternativeFirst gopurs_runtime.Value
var once_alternativeFirst sync.Once
func Get_alternativeFirst() gopurs_runtime.Value {
	once_alternativeFirst.Do(func() {
		cache_alternativeFirst = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusFirst()
}))
	})
	return cache_alternativeFirst
}

func Call_First(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showFirst(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showMaybe_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_Maybe.Get_showMaybe(), dictShow_0))
_ = showMaybe_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("First ("), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(showMaybe_1_0.V0, v_2), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_ordFirst(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
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
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
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

func Call_eqFirst(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))
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

func Call_boundedFirst(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Apply(pkg_Data_Maybe.Get_boundedMaybe(), dictBounded_0)
}


