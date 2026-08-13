package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Product2_Product2 gopurs_runtime.Value
var once_Data_Functor_Product2_Product2 sync.Once
func Get_Data_Functor_Product2_Product2() gopurs_runtime.Value {
	once_Data_Functor_Product2_Product2.Do(func() {
		cache_Data_Functor_Product2_Product2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Data_Functor_Product2_Product2
}

var cache_Data_Functor_Product2_showProduct2 gopurs_runtime.Value
var once_Data_Functor_Product2_showProduct2 sync.Once
func Get_Data_Functor_Product2_showProduct2() gopurs_runtime.Value {
	once_Data_Functor_Product2_showProduct2.Do(func() {
		cache_Data_Functor_Product2_showProduct2 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product2_showProduct2(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Functor_Product2_showProduct2
}

var cache_Data_Functor_Product2_profunctorProduct2 gopurs_runtime.Value
var once_Data_Functor_Product2_profunctorProduct2 sync.Once
func Get_Data_Functor_Product2_profunctorProduct2() gopurs_runtime.Value {
	once_Data_Functor_Product2_profunctorProduct2.Do(func() {
		cache_Data_Functor_Product2_profunctorProduct2 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, dictProfunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product2_profunctorProduct2(dictProfunctor_0_box, dictProfunctor1_1_box)
})
	})
	return cache_Data_Functor_Product2_profunctorProduct2
}

var cache_Data_Functor_Product2_functorProduct2 gopurs_runtime.Value
var once_Data_Functor_Product2_functorProduct2 sync.Once
func Get_Data_Functor_Product2_functorProduct2() gopurs_runtime.Value {
	once_Data_Functor_Product2_functorProduct2.Do(func() {
		cache_Data_Functor_Product2_functorProduct2 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product2_functorProduct2(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_Data_Functor_Product2_functorProduct2
}

var cache_Data_Functor_Product2_eqProduct2 gopurs_runtime.Value
var once_Data_Functor_Product2_eqProduct2 sync.Once
func Get_Data_Functor_Product2_eqProduct2() gopurs_runtime.Value {
	once_Data_Functor_Product2_eqProduct2.Do(func() {
		cache_Data_Functor_Product2_eqProduct2 = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product2_eqProduct2(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_Data_Functor_Product2_eqProduct2
}

var cache_Data_Functor_Product2_ordProduct2 gopurs_runtime.Value
var once_Data_Functor_Product2_ordProduct2 sync.Once
func Get_Data_Functor_Product2_ordProduct2() gopurs_runtime.Value {
	once_Data_Functor_Product2_ordProduct2.Do(func() {
		cache_Data_Functor_Product2_ordProduct2 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product2_ordProduct2(dictOrd_0_box)
})
	})
	return cache_Data_Functor_Product2_ordProduct2
}

var cache_Data_Functor_Product2_bifunctorProduct2 gopurs_runtime.Value
var once_Data_Functor_Product2_bifunctorProduct2 sync.Once
func Get_Data_Functor_Product2_bifunctorProduct2() gopurs_runtime.Value {
	once_Data_Functor_Product2_bifunctorProduct2.Do(func() {
		cache_Data_Functor_Product2_bifunctorProduct2 = gopurs_runtime.Func2(func(dictBifunctor_0_box gopurs_runtime.Value, dictBifunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product2_bifunctorProduct2(dictBifunctor_0_box, dictBifunctor1_1_box)
})
	})
	return cache_Data_Functor_Product2_bifunctorProduct2
}

var cache_Data_Functor_Product2_biapplyProduct2 gopurs_runtime.Value
var once_Data_Functor_Product2_biapplyProduct2 sync.Once
func Get_Data_Functor_Product2_biapplyProduct2() gopurs_runtime.Value {
	once_Data_Functor_Product2_biapplyProduct2.Do(func() {
		cache_Data_Functor_Product2_biapplyProduct2 = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product2_biapplyProduct2(dictBiapply_0_box)
})
	})
	return cache_Data_Functor_Product2_biapplyProduct2
}

var cache_Data_Functor_Product2_biapplicativeProduct2 gopurs_runtime.Value
var once_Data_Functor_Product2_biapplicativeProduct2 sync.Once
func Get_Data_Functor_Product2_biapplicativeProduct2() gopurs_runtime.Value {
	once_Data_Functor_Product2_biapplicativeProduct2.Do(func() {
		cache_Data_Functor_Product2_biapplicativeProduct2 = gopurs_runtime.Func(func(dictBiapplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Product2_biapplicativeProduct2(dictBiapplicative_0_box)
})
	})
	return cache_Data_Functor_Product2_biapplicativeProduct2
}

type Constructor_Data_Functor_Product2_Product2[T_f any, T_g any, T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_Data_Functor_Product2_showProduct2(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Product2 ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))
}

func Call_Data_Functor_Product2_profunctorProduct2(dictProfunctor_0_loop gopurs_runtime.Value, dictProfunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var dictProfunctor1_1 gopurs_runtime.Value = dictProfunctor1_1_loop
_ = dictProfunctor1_1
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_2, g_3, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor1_1, "dimap"), f_2, g_3, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
})
}))
}

func Call_Data_Functor_Product2_functorProduct2(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
})
}))
}

func Call_Data_Functor_Product2_eqProduct2(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
})
}))
}

func Call_Data_Functor_Product2_ordProduct2(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): eqProduct22_3_1 -> gopurs_runtime.Value
eqProduct22_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal) != (0)))
})
}))
_ = eqProduct22_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqProduct22_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_3 -> gopurs_runtime.Value
v_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
_ = v_6_3
var __t4 uint32
{
if (uint32(v_6_3.IntVal) == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (uint32(v_6_3.IntVal) == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
__t4 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
}))
})
}

func Call_Data_Functor_Product2_bifunctorProduct2(dictBifunctor_0_loop gopurs_runtime.Value, dictBifunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
var dictBifunctor1_1 gopurs_runtime.Value = dictBifunctor1_1_loop
_ = dictBifunctor1_1
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_2, g_3, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_1, "bimap"), f_2, g_3, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
})
}))
}

func Call_Data_Functor_Product2_biapplyProduct2(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictBiapply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply1_2, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): bifunctorProduct22_3_1 -> gopurs_runtime.Value
bifunctorProduct22_3_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_4, g_5, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "bimap"), f_4, g_5, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)})}
})
})
}))
_ = bifunctorProduct22_3_1
return gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply1_2, "biapply"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
}))
})
}

func Call_Data_Functor_Product2_biapplicativeProduct2(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
_ = dictBiapplicative_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictBiapplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative1_3, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): bifunctorProduct22_5_4 -> gopurs_runtime.Value
bifunctorProduct22_5_4 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "bimap"), f_6, g_7, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_5, "bimap"), f_6, g_7, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1)})}
})
})
}))
_ = bifunctorProduct22_5_4
// TAST (Let): biapplyProduct22_4_2 -> gopurs_runtime.Value
biapplyProduct22_4_2 := gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "biapply"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "biapply"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1)})}
})
}))
_ = biapplyProduct22_4_2
return gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyProduct22_4_2
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), a_5, b_6), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative1_3, "bipure"), a_5, b_6)})}
})
}))
})
}


