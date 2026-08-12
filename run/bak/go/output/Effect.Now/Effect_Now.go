package Effect_Now

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_DateTime "gopurs/output/Data.DateTime"
	pkg_Data_DateTime_Instant "gopurs/output/Data.DateTime.Instant"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Effect "gopurs/output/Effect"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_nowTime gopurs_runtime.Value
var once_nowTime sync.Once
func Get_nowTime() gopurs_runtime.Value {
	once_nowTime.Do(func() {
		cache_nowTime = Call_map__3644121587(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Apply(pkg_Data_DateTime_Instant.Get_toDateTime(), x_0).UnsafePtr).V1)}
}), Get_now())
	})
	return cache_nowTime
}

var cache_nowDateTime gopurs_runtime.Value
var once_nowDateTime sync.Once
func Get_nowDateTime() gopurs_runtime.Value {
	once_nowDateTime.Do(func() {
		cache_nowDateTime = Call_map__2261469235(pkg_Data_DateTime_Instant.Get_toDateTime(), Get_now())
	})
	return cache_nowDateTime
}

var cache_nowDate gopurs_runtime.Value
var once_nowDate sync.Once
func Get_nowDate() gopurs_runtime.Value {
	once_nowDate.Do(func() {
		cache_nowDate = Call_map__2919116915(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Apply(pkg_Data_DateTime_Instant.Get_toDateTime(), x_0).UnsafePtr).V0)}
}), Get_now())
	})
	return cache_nowDate
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_fromEnum__1637084359 gopurs_runtime.Value
var once_fromEnum__1637084359 sync.Once
func Get_fromEnum__1637084359() gopurs_runtime.Value {
	once_fromEnum__1637084359.Do(func() {
		cache_fromEnum__1637084359 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__1637084359(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_fromEnum__1637084359
}

var cache_fromEnum__1196942535 gopurs_runtime.Value
var once_fromEnum__1196942535 sync.Once
func Get_fromEnum__1196942535() gopurs_runtime.Value {
	once_fromEnum__1196942535.Do(func() {
		cache_fromEnum__1196942535 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_fromEnum__1196942535(uint32(v_0_box.IntVal)))
})
	})
	return cache_fromEnum__1196942535
}

var cache_toEnum__3317293286 gopurs_runtime.Value
var once_toEnum__3317293286 sync.Once
func Get_toEnum__3317293286() gopurs_runtime.Value {
	once_toEnum__3317293286.Do(func() {
		cache_toEnum__3317293286 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum__3317293286(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_toEnum__3317293286
}

var cache_toEnum__2309750950 gopurs_runtime.Value
var once_toEnum__2309750950 sync.Once
func Get_toEnum__2309750950() gopurs_runtime.Value {
	once_toEnum__2309750950.Do(func() {
		cache_toEnum__2309750950 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_toEnum__2309750950(v_0_box.IntVal))}
})
	})
	return cache_toEnum__2309750950
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

var cache_map__2919116915 gopurs_runtime.Value
var once_map__2919116915 sync.Once
func Get_map__2919116915() gopurs_runtime.Value {
	once_map__2919116915.Do(func() {
		cache_map__2919116915 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2919116915(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__2919116915
}

var cache_map__2261469235 gopurs_runtime.Value
var once_map__2261469235 sync.Once
func Get_map__2261469235() gopurs_runtime.Value {
	once_map__2261469235.Do(func() {
		cache_map__2261469235 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2261469235(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__2261469235
}

var cache_map__3644121587 gopurs_runtime.Value
var once_map__3644121587 sync.Once
func Get_map__3644121587() gopurs_runtime.Value {
	once_map__3644121587.Do(func() {
		cache_map__3644121587 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3644121587(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__3644121587
}

var cache_fromJust__1791383420 gopurs_runtime.Value
var once_fromJust__1791383420 sync.Once
func Get_fromJust__1791383420() gopurs_runtime.Value {
	once_fromJust__1791383420.Do(func() {
		cache_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__1791383420
}

var cache_fromJust__4142563260 gopurs_runtime.Value
var once_fromJust__4142563260 sync.Once
func Get_fromJust__4142563260() gopurs_runtime.Value {
	once_fromJust__4142563260.Do(func() {
		cache_fromJust__4142563260 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__4142563260(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__4142563260
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_functorEffect__347161653 gopurs_runtime.Value
var once_functorEffect__347161653 sync.Once
func Get_functorEffect__347161653() gopurs_runtime.Value {
	once_functorEffect__347161653.Do(func() {
		cache_functorEffect__347161653 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__347161653
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1130268957 gopurs_runtime.Value
var once_unsafePartial__1130268957 sync.Once
func Get_unsafePartial__1130268957() gopurs_runtime.Value {
	once_unsafePartial__1130268957.Do(func() {
		cache_unsafePartial__1130268957 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1130268957
}

var cache_unsafePartial__3962928989 gopurs_runtime.Value
var once_unsafePartial__3962928989 sync.Once
func Get_unsafePartial__3962928989() gopurs_runtime.Value {
	once_unsafePartial__3962928989.Do(func() {
		cache_unsafePartial__3962928989 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__3962928989
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromEnum__1637084359(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_fromEnum__1196942535(v_0_loop uint32) int64 {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == 1908470532) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (v_0 == 2455627378) {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if (v_0 == 4162469099) {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if (v_0 == 1692989816) {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if (v_0 == 330658827) {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if (v_0 == 4067355978) {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if (v_0 == 2276710548) {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if (v_0 == 243771071) {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if (v_0 == 215731793) {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if (v_0 == 8639228) {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if (v_0 == 49471444) {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if (v_0 == 3889233761) {
__t0 = gopurs_runtime.Int(12)
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

func Call_toEnum__3317293286(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_toEnum__2309750950(v_0_loop int64) *pkg_Data_Maybe.Constructor_Just[uint32] {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (1) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (2) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (3) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (4) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (5) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (6) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (7) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (8) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (9) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (10) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (11) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (12) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](__t0)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2919116915(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), __eta0_0), __eta1_1)
}

func Call_map__2261469235(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), __eta0_0), __eta1_1)
}

func Call_map__3644121587(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), __eta0_0), __eta1_1)
}

func Call_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
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

func Call_fromJust__4142563260(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
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

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_getTimezoneOffset() gopurs_runtime.Value {
	return _Gopurs_GetTimezoneOffset
}

func Get_now() gopurs_runtime.Value {
	return _Gopurs_Now
}
