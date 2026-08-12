package Data_Profunctor_Choice

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_right gopurs_runtime.Value
var once_right sync.Once
func Get_right() gopurs_runtime.Value {
	once_right.Do(func() {
		cache_right = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right(gopurs_runtime.CoerceToStruct[Constructor_Choice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_right
}

var cache_left gopurs_runtime.Value
var once_left sync.Once
func Get_left() gopurs_runtime.Value {
	once_left.Do(func() {
		cache_left = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_left(gopurs_runtime.CoerceToStruct[Constructor_Choice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_left
}

var cache_splitChoice gopurs_runtime.Value
var once_splitChoice sync.Once
func Get_splitChoice() gopurs_runtime.Value {
	once_splitChoice.Do(func() {
		cache_splitChoice = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictChoice_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitChoice(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Choice[gopurs_runtime.Value]](dictChoice_1_box), l_2_box, r_3_box)
})
	})
	return cache_splitChoice
}

var cache_fanin gopurs_runtime.Value
var once_fanin sync.Once
func Get_fanin() gopurs_runtime.Value {
	once_fanin.Do(func() {
		cache_fanin = gopurs_runtime.Func2(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictChoice_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fanin(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Choice[gopurs_runtime.Value]](dictChoice_1_box))
})
	})
	return cache_fanin
}

var cache_choiceFn gopurs_runtime.Value
var once_choiceFn sync.Once
func Get_choiceFn() gopurs_runtime.Value {
	once_choiceFn.Do(func() {
		cache_choiceFn = gopurs_runtime.RecordDict3("Profunctor0", "left", "right", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0})}
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
}), gopurs_runtime.RecordGet(pkg_Data_Either.Get_functorEither(), "map"))
	})
	return cache_choiceFn
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

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
}

var cache_composeFlipped__4057377183 gopurs_runtime.Value
var once_composeFlipped__4057377183 sync.Once
func Get_composeFlipped__4057377183() gopurs_runtime.Value {
	once_composeFlipped__4057377183.Do(func() {
		cache_composeFlipped__4057377183 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__4057377183(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__4057377183
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

var cache_left__3408498362 gopurs_runtime.Value
var once_left__3408498362 sync.Once
func Get_left__3408498362() gopurs_runtime.Value {
	once_left__3408498362.Do(func() {
		cache_left__3408498362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_left__3408498362(gopurs_runtime.CoerceToStruct[Constructor_Choice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_left__3408498362
}

var cache_right__3408498362 gopurs_runtime.Value
var once_right__3408498362 sync.Once
func Get_right__3408498362() gopurs_runtime.Value {
	once_right__3408498362.Do(func() {
		cache_right__3408498362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__3408498362(gopurs_runtime.CoerceToStruct[Constructor_Choice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_right__3408498362
}

var cache_splitChoice__1826977503 gopurs_runtime.Value
var once_splitChoice__1826977503 sync.Once
func Get_splitChoice__1826977503() gopurs_runtime.Value {
	once_splitChoice__1826977503.Do(func() {
		cache_splitChoice__1826977503 = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictChoice_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitChoice__1826977503(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), gopurs_runtime.CoerceToStruct[Constructor_Choice[gopurs_runtime.Value]](dictChoice_1_box), l_2_box, r_3_box)
})
	})
	return cache_splitChoice__1826977503
}

var cache_dimap__1466332548 gopurs_runtime.Value
var once_dimap__1466332548 sync.Once
func Get_dimap__1466332548() gopurs_runtime.Value {
	once_dimap__1466332548.Do(func() {
		cache_dimap__1466332548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dimap__1466332548(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_dimap__1466332548
}

var cache_profunctorFn__3736629211 gopurs_runtime.Value
var once_profunctorFn__3736629211 sync.Once
func Get_profunctorFn__3736629211() gopurs_runtime.Value {
	once_profunctorFn__3736629211.Do(func() {
		cache_profunctorFn__3736629211 = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
}))
	})
	return cache_profunctorFn__3736629211
}

var cache_rmap__1762133278 gopurs_runtime.Value
var once_rmap__1762133278 sync.Once
func Get_rmap__1762133278() gopurs_runtime.Value {
	once_rmap__1762133278.Do(func() {
		cache_rmap__1762133278 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, b2c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rmap__1762133278(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), b2c_1_box)
})
	})
	return cache_rmap__1762133278
}

var cache_rmap__2905617982 gopurs_runtime.Value
var once_rmap__2905617982 sync.Once
func Get_rmap__2905617982() gopurs_runtime.Value {
	once_rmap__2905617982.Do(func() {
		cache_rmap__2905617982 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, b2c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rmap__2905617982(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), b2c_1_box)
})
	})
	return cache_rmap__2905617982
}

type Constructor_Choice[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3666633887] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Choice[gopurs_runtime.Value])(ptr)
		switch key {
		case "Profunctor0": return c.V0
		case "left": return c.V1
		case "right": return c.V2
		default: panic("Key not found in dictionary Constructor_Choice: " + key)
		}
	}
}


func Call_right(dict_0_loop *Constructor_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Choice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_left(dict_0_loop *Constructor_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Choice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_splitChoice(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], dictChoice_1_loop *Constructor_Choice[gopurs_runtime.Value], l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictChoice_1 *Constructor_Choice[gopurs_runtime.Value] = dictChoice_1_loop
_ = dictChoice_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, gopurs_runtime.Apply(dictChoice_1.V2, r_3), gopurs_runtime.Apply(dictChoice_1.V1, l_2))
}

func Call_fanin(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], dictChoice_1_loop *Constructor_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictChoice_1 *Constructor_Choice[gopurs_runtime.Value] = dictChoice_1_loop
_ = dictChoice_1
Profunctor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictChoice_1.V0, gopurs_runtime.Value{}))
_ = Profunctor0_2_0
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Profunctor0_2_0.V0, pkg_Data_Profunctor.Get_identity1(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 3711209382) {
__t1 = (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2465973597) {
__t1 = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Apply2(dictSemigroupoid_0.V0, gopurs_runtime.Apply(dictChoice_1.V2, r_4), gopurs_runtime.Apply(dictChoice_1.V1, l_3)))
})
})
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
}

func Call_composeFlipped__4057377183(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
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

func Call_left__3408498362(dict_0_loop *Constructor_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Choice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_right__3408498362(dict_0_loop *Constructor_Choice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Choice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_splitChoice__1826977503(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], dictChoice_1_loop *Constructor_Choice[gopurs_runtime.Value], l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictChoice_1 *Constructor_Choice[gopurs_runtime.Value] = dictChoice_1_loop
_ = dictChoice_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, gopurs_runtime.Apply(dictChoice_1.V2, r_3), gopurs_runtime.Apply(dictChoice_1.V1, l_2))
}

func Call_dimap__1466332548(dict_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_rmap__1762133278(dictProfunctor_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value], b2c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var b2c_1 gopurs_runtime.Value = b2c_1_loop
_ = b2c_1
return gopurs_runtime.Apply2(dictProfunctor_0.V0, pkg_Data_Profunctor.Get_identity1(), b2c_1)
}

func Call_rmap__2905617982(dictProfunctor_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value], b2c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var b2c_1 gopurs_runtime.Value = b2c_1_loop
_ = b2c_1
return gopurs_runtime.Apply2(dictProfunctor_0.V0, pkg_Data_Profunctor.Get_identity1(), b2c_1)
}


