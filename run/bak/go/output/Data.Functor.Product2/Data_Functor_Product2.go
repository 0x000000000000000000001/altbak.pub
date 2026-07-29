package Data_Functor_Product2

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	unsafe "unsafe"
)

var cache_Product2 gopurs_runtime.Value
var once_Product2 sync.Once
func Get_Product2() gopurs_runtime.Value {
	once_Product2.Do(func() {
		cache_Product2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Product2
}

var cache_showProduct2 gopurs_runtime.Value
var once_showProduct2 sync.Once
func Get_showProduct2() gopurs_runtime.Value {
	once_showProduct2.Do(func() {
		cache_showProduct2 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showProduct2(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showProduct2
}

var cache_profunctorProduct2 gopurs_runtime.Value
var once_profunctorProduct2 sync.Once
func Get_profunctorProduct2() gopurs_runtime.Value {
	once_profunctorProduct2.Do(func() {
		cache_profunctorProduct2 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, dictProfunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_profunctorProduct2(dictProfunctor_0_box, dictProfunctor1_1_box)
})
	})
	return cache_profunctorProduct2
}

var cache_functorProduct2 gopurs_runtime.Value
var once_functorProduct2 sync.Once
func Get_functorProduct2() gopurs_runtime.Value {
	once_functorProduct2.Do(func() {
		cache_functorProduct2 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorProduct2(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_functorProduct2
}

var cache_eqProduct2 gopurs_runtime.Value
var once_eqProduct2 sync.Once
func Get_eqProduct2() gopurs_runtime.Value {
	once_eqProduct2.Do(func() {
		cache_eqProduct2 = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqProduct2(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_eqProduct2
}

var cache_ordProduct2 gopurs_runtime.Value
var once_ordProduct2 sync.Once
func Get_ordProduct2() gopurs_runtime.Value {
	once_ordProduct2.Do(func() {
		cache_ordProduct2 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordProduct2(dictOrd_0_box)
})
	})
	return cache_ordProduct2
}

var cache_bifunctorProduct2 gopurs_runtime.Value
var once_bifunctorProduct2 sync.Once
func Get_bifunctorProduct2() gopurs_runtime.Value {
	once_bifunctorProduct2.Do(func() {
		cache_bifunctorProduct2 = gopurs_runtime.Func2(func(dictBifunctor_0_box gopurs_runtime.Value, dictBifunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifunctorProduct2(dictBifunctor_0_box, dictBifunctor1_1_box)
})
	})
	return cache_bifunctorProduct2
}

var cache_biapplyProduct2 gopurs_runtime.Value
var once_biapplyProduct2 sync.Once
func Get_biapplyProduct2() gopurs_runtime.Value {
	once_biapplyProduct2.Do(func() {
		cache_biapplyProduct2 = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplyProduct2(dictBiapply_0_box)
})
	})
	return cache_biapplyProduct2
}

var cache_biapplicativeProduct2 gopurs_runtime.Value
var once_biapplicativeProduct2 sync.Once
func Get_biapplicativeProduct2() gopurs_runtime.Value {
	once_biapplicativeProduct2.Do(func() {
		cache_biapplicativeProduct2 = gopurs_runtime.Func(func(dictBiapplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplicativeProduct2(dictBiapplicative_0_box)
})
	})
	return cache_biapplicativeProduct2
}

type Constructor_Product2[T_f any, T_g any, T_a any, T_b any] struct {
	Rc uint32
	V0 T_f
	V1 T_g
}


func Call_showProduct2(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Product2 "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Str(")")))))
}))
}

func Call_profunctorProduct2(dictProfunctor_0_loop gopurs_runtime.Value, dictProfunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var dictProfunctor1_1 gopurs_runtime.Value = dictProfunctor1_1_loop
_ = dictProfunctor1_1
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_2, g_3, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor1_1, "dimap"), f_2, g_3, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
}))
}

func Call_functorProduct2(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
}))
}

func Call_eqProduct2(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1))
}))
}

func Call_ordProduct2(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqProduct21_1_0 := gopurs_runtime.Apply(Get_eqProduct2(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqProduct21_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqProduct22_3_1 := gopurs_runtime.Apply(eqProduct21_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqProduct22_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqProduct22_3_1
}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 380165415) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1)
}
end_branch_3:
return __t3
}))
})
}

func Call_bifunctorProduct2(dictBifunctor_0_loop gopurs_runtime.Value, dictBifunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
var dictBifunctor1_1 gopurs_runtime.Value = dictBifunctor1_1_loop
_ = dictBifunctor1_1
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_2, g_3, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_1, "bimap"), f_2, g_3, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
}))
}

func Call_biapplyProduct2(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictBiapply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply1_2, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
bifunctorProduct22_4_2 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((__local_var_1_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_1_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_4, g_5, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V1 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "bimap"), f_4, g_5, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
return __local_var_1_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_4, g_5, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "bimap"), f_4, g_5, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)})}
}
}()
}))
_ = bifunctorProduct22_4_2
return gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_4_2
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((__local_var_1_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_1_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0)
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply1_2, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
return __local_var_1_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply1_2, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1)})}
}
}()
}))
})
}

func Call_biapplicativeProduct2(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
_ = dictBiapplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictBiapplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative1_3, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_5_3
bifunctorProduct22_6_5 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, g_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((__local_var_1_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_1_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "bimap"), f_6, g_7, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V1 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_3, "bimap"), f_6, g_7, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1)
return __local_var_1_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "bimap"), f_6, g_7, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_3, "bimap"), f_6, g_7, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1)})}
}
}()
}))
_ = bifunctorProduct22_6_5
biapplyProduct22_6_4 := gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_6_5
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((__local_var_1_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_1_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0)
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1)
return __local_var_1_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1)})}
}
}()
}))
_ = biapplyProduct22_6_4
return gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyProduct22_6_4
}), gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((__local_var_1_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_1_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), a_7, b_8)
(*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative1_3, "bipure"), a_7, b_8)
return __local_var_1_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), a_7, b_8), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative1_3, "bipure"), a_7, b_8)})}
}
}()
}))
})
}


