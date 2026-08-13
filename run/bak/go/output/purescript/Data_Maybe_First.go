package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Maybe_First_First gopurs_runtime.Value
var once_Data_Maybe_First_First sync.Once
func Get_Data_Maybe_First_First() gopurs_runtime.Value {
	once_Data_Maybe_First_First.Do(func() {
		cache_Data_Maybe_First_First = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_First(x_0_box)
})
	})
	return cache_Data_Maybe_First_First
}

var cache_Data_Maybe_First_showFirst gopurs_runtime.Value
var once_Data_Maybe_First_showFirst sync.Once
func Get_Data_Maybe_First_showFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_showFirst.Do(func() {
		cache_Data_Maybe_First_showFirst = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_showFirst(dictShow_0_box)
})
	})
	return cache_Data_Maybe_First_showFirst
}

var cache_Data_Maybe_First_semigroupFirst gopurs_runtime.Value
var once_Data_Maybe_First_semigroupFirst sync.Once
func Get_Data_Maybe_First_semigroupFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_semigroupFirst.Do(func() {
		cache_Data_Maybe_First_semigroupFirst = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
}))
	})
	return cache_Data_Maybe_First_semigroupFirst
}

var cache_Data_Maybe_First_ordFirst gopurs_runtime.Value
var once_Data_Maybe_First_ordFirst sync.Once
func Get_Data_Maybe_First_ordFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_ordFirst.Do(func() {
		cache_Data_Maybe_First_ordFirst = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_ordFirst(dictOrd_0_box)
})
	})
	return cache_Data_Maybe_First_ordFirst
}

var cache_Data_Maybe_First_ord1First gopurs_runtime.Value
var once_Data_Maybe_First_ord1First sync.Once
func Get_Data_Maybe_First_ord1First() gopurs_runtime.Value {
	once_Data_Maybe_First_ord1First.Do(func() {
		cache_Data_Maybe_First_ord1First = Get_Data_Maybe_ord1Maybe()
	})
	return cache_Data_Maybe_First_ord1First
}

var cache_Data_Maybe_First_newtypeFirst gopurs_runtime.Value
var once_Data_Maybe_First_newtypeFirst sync.Once
func Get_Data_Maybe_First_newtypeFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_newtypeFirst.Do(func() {
		cache_Data_Maybe_First_newtypeFirst = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Maybe_First_newtypeFirst
}

var cache_Data_Maybe_First_monoidFirst gopurs_runtime.Value
var once_Data_Maybe_First_monoidFirst sync.Once
func Get_Data_Maybe_First_monoidFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_monoidFirst.Do(func() {
		cache_Data_Maybe_First_monoidFirst = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_First_semigroupFirst()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
	})
	return cache_Data_Maybe_First_monoidFirst
}

var cache_Data_Maybe_First_monadFirst gopurs_runtime.Value
var once_Data_Maybe_First_monadFirst sync.Once
func Get_Data_Maybe_First_monadFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_monadFirst.Do(func() {
		cache_Data_Maybe_First_monadFirst = Get_Data_Maybe_monadMaybe()
	})
	return cache_Data_Maybe_First_monadFirst
}

var cache_Data_Maybe_First_invariantFirst gopurs_runtime.Value
var once_Data_Maybe_First_invariantFirst sync.Once
func Get_Data_Maybe_First_invariantFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_invariantFirst.Do(func() {
		cache_Data_Maybe_First_invariantFirst = Get_Data_Maybe_invariantMaybe()
	})
	return cache_Data_Maybe_First_invariantFirst
}

var cache_Data_Maybe_First_functorFirst gopurs_runtime.Value
var once_Data_Maybe_First_functorFirst sync.Once
func Get_Data_Maybe_First_functorFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_functorFirst.Do(func() {
		cache_Data_Maybe_First_functorFirst = Get_Data_Maybe_functorMaybe()
	})
	return cache_Data_Maybe_First_functorFirst
}

var cache_Data_Maybe_First_extendFirst gopurs_runtime.Value
var once_Data_Maybe_First_extendFirst sync.Once
func Get_Data_Maybe_First_extendFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_extendFirst.Do(func() {
		cache_Data_Maybe_First_extendFirst = Get_Data_Maybe_extendMaybe()
	})
	return cache_Data_Maybe_First_extendFirst
}

var cache_Data_Maybe_First_eqFirst gopurs_runtime.Value
var once_Data_Maybe_First_eqFirst sync.Once
func Get_Data_Maybe_First_eqFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_eqFirst.Do(func() {
		cache_Data_Maybe_First_eqFirst = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_eqFirst(dictEq_0_box)
})
	})
	return cache_Data_Maybe_First_eqFirst
}

var cache_Data_Maybe_First_eq1First gopurs_runtime.Value
var once_Data_Maybe_First_eq1First sync.Once
func Get_Data_Maybe_First_eq1First() gopurs_runtime.Value {
	once_Data_Maybe_First_eq1First.Do(func() {
		cache_Data_Maybe_First_eq1First = Get_Data_Maybe_eq1Maybe()
	})
	return cache_Data_Maybe_First_eq1First
}

var cache_Data_Maybe_First_boundedFirst gopurs_runtime.Value
var once_Data_Maybe_First_boundedFirst sync.Once
func Get_Data_Maybe_First_boundedFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_boundedFirst.Do(func() {
		cache_Data_Maybe_First_boundedFirst = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_First_boundedFirst(dictBounded_0_box)
})
	})
	return cache_Data_Maybe_First_boundedFirst
}

var cache_Data_Maybe_First_bindFirst gopurs_runtime.Value
var once_Data_Maybe_First_bindFirst sync.Once
func Get_Data_Maybe_First_bindFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_bindFirst.Do(func() {
		cache_Data_Maybe_First_bindFirst = Get_Data_Maybe_bindMaybe()
	})
	return cache_Data_Maybe_First_bindFirst
}

var cache_Data_Maybe_First_applyFirst gopurs_runtime.Value
var once_Data_Maybe_First_applyFirst sync.Once
func Get_Data_Maybe_First_applyFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_applyFirst.Do(func() {
		cache_Data_Maybe_First_applyFirst = Get_Data_Maybe_applyMaybe()
	})
	return cache_Data_Maybe_First_applyFirst
}

var cache_Data_Maybe_First_applicativeFirst gopurs_runtime.Value
var once_Data_Maybe_First_applicativeFirst sync.Once
func Get_Data_Maybe_First_applicativeFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_applicativeFirst.Do(func() {
		cache_Data_Maybe_First_applicativeFirst = Get_Data_Maybe_applicativeMaybe()
	})
	return cache_Data_Maybe_First_applicativeFirst
}

var cache_Data_Maybe_First_altFirst gopurs_runtime.Value
var once_Data_Maybe_First_altFirst sync.Once
func Get_Data_Maybe_First_altFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_altFirst.Do(func() {
		cache_Data_Maybe_First_altFirst = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.RecordGet(Get_Data_Maybe_First_semigroupFirst(), "append"))
	})
	return cache_Data_Maybe_First_altFirst
}

var cache_Data_Maybe_First_plusFirst gopurs_runtime.Value
var once_Data_Maybe_First_plusFirst sync.Once
func Get_Data_Maybe_First_plusFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_plusFirst.Do(func() {
		cache_Data_Maybe_First_plusFirst = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_First_altFirst()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(Get_Data_Maybe_First_monoidFirst(), "mempty")))})
	})
	return cache_Data_Maybe_First_plusFirst
}

var cache_Data_Maybe_First_alternativeFirst gopurs_runtime.Value
var once_Data_Maybe_First_alternativeFirst sync.Once
func Get_Data_Maybe_First_alternativeFirst() gopurs_runtime.Value {
	once_Data_Maybe_First_alternativeFirst.Do(func() {
		cache_Data_Maybe_First_alternativeFirst = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_First_plusFirst()
}))
	})
	return cache_Data_Maybe_First_alternativeFirst
}

var cache_Data_Maybe_First_altFirst__4201091523 gopurs_runtime.Value
var once_Data_Maybe_First_altFirst__4201091523 sync.Once
func Get_Data_Maybe_First_altFirst__4201091523() gopurs_runtime.Value {
	once_Data_Maybe_First_altFirst__4201091523.Do(func() {
		cache_Data_Maybe_First_altFirst__4201091523 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.RecordGet(Get_Data_Maybe_First_semigroupFirst(), "append"))
	})
	return cache_Data_Maybe_First_altFirst__4201091523
}

var cache_Data_Maybe_First_applicativeFirst__500933224 gopurs_runtime.Value
var once_Data_Maybe_First_applicativeFirst__500933224 sync.Once
func Get_Data_Maybe_First_applicativeFirst__500933224() gopurs_runtime.Value {
	once_Data_Maybe_First_applicativeFirst__500933224.Do(func() {
		cache_Data_Maybe_First_applicativeFirst__500933224 = Get_Data_Maybe_applicativeMaybe()
	})
	return cache_Data_Maybe_First_applicativeFirst__500933224
}

var cache_Data_Maybe_First_functorFirst__2097654001 gopurs_runtime.Value
var once_Data_Maybe_First_functorFirst__2097654001 sync.Once
func Get_Data_Maybe_First_functorFirst__2097654001() gopurs_runtime.Value {
	once_Data_Maybe_First_functorFirst__2097654001.Do(func() {
		cache_Data_Maybe_First_functorFirst__2097654001 = Get_Data_Maybe_functorMaybe()
	})
	return cache_Data_Maybe_First_functorFirst__2097654001
}

var cache_Data_Maybe_First_monoidFirst__4133827187 gopurs_runtime.Value
var once_Data_Maybe_First_monoidFirst__4133827187 sync.Once
func Get_Data_Maybe_First_monoidFirst__4133827187() gopurs_runtime.Value {
	once_Data_Maybe_First_monoidFirst__4133827187.Do(func() {
		cache_Data_Maybe_First_monoidFirst__4133827187 = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_First_semigroupFirst()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
	})
	return cache_Data_Maybe_First_monoidFirst__4133827187
}

var cache_Data_Maybe_First_plusFirst__400696082 gopurs_runtime.Value
var once_Data_Maybe_First_plusFirst__400696082 sync.Once
func Get_Data_Maybe_First_plusFirst__400696082() gopurs_runtime.Value {
	once_Data_Maybe_First_plusFirst__400696082.Do(func() {
		cache_Data_Maybe_First_plusFirst__400696082 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_First_altFirst()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(Get_Data_Maybe_First_monoidFirst(), "mempty")))})
	})
	return cache_Data_Maybe_First_plusFirst__400696082
}

var cache_Data_Maybe_First_semigroupFirst__3179391308 gopurs_runtime.Value
var once_Data_Maybe_First_semigroupFirst__3179391308 sync.Once
func Get_Data_Maybe_First_semigroupFirst__3179391308() gopurs_runtime.Value {
	once_Data_Maybe_First_semigroupFirst__3179391308.Do(func() {
		cache_Data_Maybe_First_semigroupFirst__3179391308 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
}))
	})
	return cache_Data_Maybe_First_semigroupFirst__3179391308
}

func Call_Data_Maybe_First_First(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Maybe_First_showFirst(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showMaybe_1_0 -> *Constructor_Data_Show_Show
showMaybe_1_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.Str((("(Just ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Maybe_Just)(v_1.UnsafePtr).V0).StrVal())) + (")"))
goto end_branch_1
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Str("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Str(__t1.StrVal())
})}
_ = showMaybe_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("First (") + (gopurs_runtime.Apply(gopurs_runtime.Box(showMaybe_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2))}).StrVal())) + (")"))
}))
}

func Call_Data_Maybe_First_ordFirst(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqMaybe1_1_0 -> gopurs_runtime.Value
eqMaybe1_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
_ = eqMaybe1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
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
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil)) && ((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr != nil)) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Maybe_Just)(x_2.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_3.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t5.IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Maybe_First_eqFirst(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
}

func Call_Data_Maybe_First_boundedFirst(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Apply(Get_Data_Maybe_boundedMaybe(), dictBounded_0)
}


