package Data_Either_Inject

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_prj gopurs_runtime.Value
var once_prj sync.Once
func Get_prj() gopurs_runtime.Value {
	once_prj.Do(func() {
		cache_prj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prj(gopurs_runtime.CoerceToStruct[Constructor_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_prj
}

var cache_injectReflexive gopurs_runtime.Value
var once_injectReflexive sync.Once
func Get_injectReflexive() gopurs_runtime.Value {
	once_injectReflexive.Do(func() {
		cache_injectReflexive = gopurs_runtime.RecordDict2("inj", "prj", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}), pkg_Data_Maybe.Get_Just())
	})
	return cache_injectReflexive
}

var cache_injectLeft gopurs_runtime.Value
var once_injectLeft sync.Once
func Get_injectLeft() gopurs_runtime.Value {
	once_injectLeft.Do(func() {
		cache_injectLeft = gopurs_runtime.RecordDict2("inj", "prj", pkg_Data_Either.Get_Left(), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
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
	return cache_injectLeft
}

var cache_inj gopurs_runtime.Value
var once_inj sync.Once
func Get_inj() gopurs_runtime.Value {
	once_inj.Do(func() {
		cache_inj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_inj(gopurs_runtime.CoerceToStruct[Constructor_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_inj
}

var cache_injectRight gopurs_runtime.Value
var once_injectRight sync.Once
func Get_injectRight() gopurs_runtime.Value {
	once_injectRight.Do(func() {
		cache_injectRight = gopurs_runtime.Func(func(dictInject_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_injectRight(dictInject_0_box)
})
	})
	return cache_injectRight
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

var cache_either__2158544585 gopurs_runtime.Value
var once_either__2158544585 sync.Once
func Get_either__2158544585() gopurs_runtime.Value {
	once_either__2158544585.Do(func() {
		cache_either__2158544585 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__2158544585(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__2158544585
}

var cache_either__3836941544 gopurs_runtime.Value
var once_either__3836941544 sync.Once
func Get_either__3836941544() gopurs_runtime.Value {
	once_either__3836941544.Do(func() {
		cache_either__3836941544 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__3836941544(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__3836941544
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_const__3525124788 gopurs_runtime.Value
var once_const__3525124788 sync.Once
func Get_const__3525124788() gopurs_runtime.Value {
	once_const__3525124788.Do(func() {
		cache_const__3525124788 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__3525124788(a_0_box, v_1_box)
})
	})
	return cache_const__3525124788
}

type Constructor_Inject[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[89024546] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Inject[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "inj": return c.V0
		case "prj": return c.V1
		default: panic("Key not found in dictionary Constructor_Inject: " + key)
		}
	}
}


func Call_prj(dict_0_loop *Constructor_Inject[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Inject[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_inj(dict_0_loop *Constructor_Inject[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Inject[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_injectRight(dictInject_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInject_0 gopurs_runtime.Value = dictInject_0_loop
_ = dictInject_0
return gopurs_runtime.RecordDict2("inj", "prj", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0, "inj"), x_1)})}
}), gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0, "prj"), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0)
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
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_either__2158544585(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_either__3836941544(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__3525124788(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}


