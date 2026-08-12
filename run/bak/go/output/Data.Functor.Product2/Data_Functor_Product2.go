package Data_Functor_Product2

import (
	pkg_Control_Biapplicative "gopurs/output/Control.Biapplicative"
	pkg_Control_Biapply "gopurs/output/Control.Biapply"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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

var cache_bipure__1449949980 gopurs_runtime.Value
var once_bipure__1449949980 sync.Once
func Get_bipure__1449949980() gopurs_runtime.Value {
	once_bipure__1449949980.Do(func() {
		cache_bipure__1449949980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bipure__1449949980(gopurs_runtime.CoerceToStruct[pkg_Control_Biapplicative.Constructor_Biapplicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bipure__1449949980
}

var cache_biapply__3394381979 gopurs_runtime.Value
var once_biapply__3394381979 sync.Once
func Get_biapply__3394381979() gopurs_runtime.Value {
	once_biapply__3394381979.Do(func() {
		cache_biapply__3394381979 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapply__3394381979(gopurs_runtime.CoerceToStruct[pkg_Control_Biapply.Constructor_Biapply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_biapply__3394381979
}

var cache_bimap__132457202 gopurs_runtime.Value
var once_bimap__132457202 sync.Once
func Get_bimap__132457202() gopurs_runtime.Value {
	once_bimap__132457202.Do(func() {
		cache_bimap__132457202 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__132457202(gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__132457202
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq__163522700 gopurs_runtime.Value
var once_eq__163522700 sync.Once
func Get_eq__163522700() gopurs_runtime.Value {
	once_eq__163522700.Do(func() {
		cache_eq__163522700 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__163522700(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__163522700
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

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_compare__3215000822 gopurs_runtime.Value
var once_compare__3215000822 sync.Once
func Get_compare__3215000822() gopurs_runtime.Value {
	once_compare__3215000822.Do(func() {
		cache_compare__3215000822 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__3215000822(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__3215000822
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

var cache_show__255526802 gopurs_runtime.Value
var once_show__255526802 sync.Once
func Get_show__255526802() gopurs_runtime.Value {
	once_show__255526802.Do(func() {
		cache_show__255526802 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__255526802(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__255526802
}

type Constructor_Product2[T_f any, T_g any, T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_showProduct2(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Product2 "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())
}))
}

func Call_profunctorProduct2(dictProfunctor_0_loop gopurs_runtime.Value, dictProfunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var dictProfunctor1_1 gopurs_runtime.Value = dictProfunctor1_1_loop
_ = dictProfunctor1_1
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_2, g_3, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor1_1, "dimap"), f_2, g_3, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
})
}))
}

func Call_functorProduct2(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
})
}))
}

func Call_eqProduct2(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0))
})
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
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
_ = v_6_2
var __t3 uint32
{
if (uint32(v_6_2.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_6_2.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
}))
})
}

func Call_bifunctorProduct2(dictBifunctor_0_loop gopurs_runtime.Value, dictBifunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
var dictBifunctor1_1 gopurs_runtime.Value = dictBifunctor1_1_loop
_ = dictBifunctor1_1
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_2, g_3, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_1, "bimap"), f_2, g_3, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
})
}))
}

func Call_biapplyProduct2(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictBiapply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply1_2, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
bifunctorProduct22_3_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_4, g_5, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "bimap"), f_4, g_5, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)})}
})
})
}))
_ = bifunctorProduct22_3_1
return gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply1_2, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
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
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative1_3, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_4_3
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
bifunctorProduct22_5_4 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "bimap"), f_6, g_7, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_5, "bimap"), f_6, g_7, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1)})}
})
})
}))
_ = bifunctorProduct22_5_4
biapplyProduct22_4_2 := gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "biapply"), (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1, (*Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1)})}
})
}))
_ = biapplyProduct22_4_2
return gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyProduct22_4_2
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), a_5, b_6), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative1_3, "bipure"), a_5, b_6)})}
})
}))
})
}

func Call_bipure__1449949980(dict_0_loop *pkg_Control_Biapplicative.Constructor_Biapplicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Biapplicative.Constructor_Biapplicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_biapply__3394381979(dict_0_loop *pkg_Control_Biapply.Constructor_Biapply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Biapply.Constructor_Biapply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bimap__132457202(dict_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__163522700(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__3215000822(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_dimap__1466332548(dict_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_show__255526802(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


