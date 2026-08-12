package Data_Maybe_Last

import (
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
		cache_semigroupLast = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
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
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
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
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(Get_monoidLast(), "mempty")))})
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

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__901270812(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__901270812
}

var cache_altLast__4201091523 gopurs_runtime.Value
var once_altLast__4201091523 sync.Once
func Get_altLast__4201091523() gopurs_runtime.Value {
	once_altLast__4201091523.Do(func() {
		cache_altLast__4201091523 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.RecordGet(Get_semigroupLast(), "append"))
	})
	return cache_altLast__4201091523
}

var cache_applicativeLast__500933224 gopurs_runtime.Value
var once_applicativeLast__500933224 sync.Once
func Get_applicativeLast__500933224() gopurs_runtime.Value {
	once_applicativeLast__500933224.Do(func() {
		cache_applicativeLast__500933224 = pkg_Data_Maybe.Get_applicativeMaybe()
	})
	return cache_applicativeLast__500933224
}

var cache_functorLast__2097654001 gopurs_runtime.Value
var once_functorLast__2097654001 sync.Once
func Get_functorLast__2097654001() gopurs_runtime.Value {
	once_functorLast__2097654001.Do(func() {
		cache_functorLast__2097654001 = pkg_Data_Maybe.Get_functorMaybe()
	})
	return cache_functorLast__2097654001
}

var cache_plusLast__400696082 gopurs_runtime.Value
var once_plusLast__400696082 sync.Once
func Get_plusLast__400696082() gopurs_runtime.Value {
	once_plusLast__400696082.Do(func() {
		cache_plusLast__400696082 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altLast()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(Get_monoidLast(), "mempty")))})
	})
	return cache_plusLast__400696082
}

var cache_semigroupLast__3179391308 gopurs_runtime.Value
var once_semigroupLast__3179391308 sync.Once
func Get_semigroupLast__3179391308() gopurs_runtime.Value {
	once_semigroupLast__3179391308.Do(func() {
		cache_semigroupLast__3179391308 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_semigroupLast__3179391308
}

var cache_applicativeMaybe__500933224 gopurs_runtime.Value
var once_applicativeMaybe__500933224 sync.Once
func Get_applicativeMaybe__500933224() gopurs_runtime.Value {
	once_applicativeMaybe__500933224.Do(func() {
		cache_applicativeMaybe__500933224 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), pkg_Data_Maybe.Get_Just())
	})
	return cache_applicativeMaybe__500933224
}

var cache_applyMaybe__3698865467 gopurs_runtime.Value
var once_applyMaybe__3698865467 sync.Once
func Get_applyMaybe__3698865467() gopurs_runtime.Value {
	once_applyMaybe__3698865467.Do(func() {
		cache_applyMaybe__3698865467 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))})))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3698865467
}

var cache_bindMaybe__3591110311 gopurs_runtime.Value
var once_bindMaybe__3591110311 sync.Once
func Get_bindMaybe__3591110311() gopurs_runtime.Value {
	once_bindMaybe__3591110311.Do(func() {
		cache_bindMaybe__3591110311 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindMaybe__3591110311
}

var cache_eq1Maybe__1662522654 gopurs_runtime.Value
var once_eq1Maybe__1662522654 sync.Once
func Get_eq1Maybe__1662522654() gopurs_runtime.Value {
	once_eq1Maybe__1662522654.Do(func() {
		cache_eq1Maybe__1662522654 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0)
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
}))
	})
	return cache_eq1Maybe__1662522654
}

var cache_extendMaybe__3503256297 gopurs_runtime.Value
var once_extendMaybe__3503256297 sync.Once
func Get_extendMaybe__3503256297() gopurs_runtime.Value {
	once_extendMaybe__3503256297.Do(func() {
		cache_extendMaybe__3503256297 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))})})}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_extendMaybe__3503256297
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_invariantMaybe__3070183577 gopurs_runtime.Value
var once_invariantMaybe__3070183577 sync.Once
func Get_invariantMaybe__3070183577() gopurs_runtime.Value {
	once_invariantMaybe__3070183577.Do(func() {
		cache_invariantMaybe__3070183577 = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), f_0)
})
}))
	})
	return cache_invariantMaybe__3070183577
}

var cache_monadMaybe__3072900051 gopurs_runtime.Value
var once_monadMaybe__3072900051 sync.Once
func Get_monadMaybe__3072900051() gopurs_runtime.Value {
	once_monadMaybe__3072900051.Do(func() {
		cache_monadMaybe__3072900051 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_bindMaybe()
}))
	})
	return cache_monadMaybe__3072900051
}

var cache_ord1Maybe__1052104681 gopurs_runtime.Value
var once_ord1Maybe__1052104681 sync.Once
func Get_ord1Maybe__1052104681() gopurs_runtime.Value {
	once_ord1Maybe__1052104681.Do(func() {
		cache_ord1Maybe__1052104681 = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_eq1Maybe()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t1.IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_ord1Maybe__1052104681
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__2183537362 gopurs_runtime.Value
var once_show__2183537362 sync.Once
func Get_show__2183537362() gopurs_runtime.Value {
	once_show__2183537362.Do(func() {
		cache_show__2183537362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2183537362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_show__2183537362
}

func Call_Last(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showLast(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showMaybe_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_Maybe.Get_showMaybe(), dictShow_0))
_ = showMaybe_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Last "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(showMaybe_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))}).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_ordLast(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
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
__t3 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)
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
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
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

func Call_eqLast(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0)
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

func Call_boundedLast(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Apply(pkg_Data_Maybe.Get_boundedMaybe(), dictBounded_0)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__901270812(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2183537362(dict_0_loop *pkg_Data_Show.Constructor_Show[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}


