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
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_106768939_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Product2 ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
})})))}
}

func Call_Data_Functor_Product2_profunctorProduct2(dictProfunctor_0_loop gopurs_runtime.Value, dictProfunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var dictProfunctor1_1 gopurs_runtime.Value = dictProfunctor1_1_loop
_ = dictProfunctor1_1
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_3352668427_4074365774((&Constructor_Data_Profunctor_Profunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_2, g_3, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor1_1, "dimap"), f_2, g_3, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)}))}
})})))}
}

func Call_Data_Functor_Product2_functorProduct2(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_3354170411_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)}))}
})})))}
}

func Call_Data_Functor_Product2_eqProduct2(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_4204058379_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
})})))}
}

func Call_Data_Functor_Product2_ordProduct2(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqProduct21__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
eqProduct21__193435443_1_0 := gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_4204058379_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
})})))}
})
_ = eqProduct21__193435443_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqProduct22_3_2 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Functor","Product2","Product2"] [(TypeVar f), (TypeVar g), (TypeVar a), (TypeVar b)])])
eqProduct22_3_2 := Rebox_Data_Functor_Product2_3790796878_4204058379(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqProduct21__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))))
_ = eqProduct22_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_686435819_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_4204058379_3790796878(eqProduct22_3_2))}
}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_6_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
_ = v_6_3
var __t4 uint32
{
if (v_6_3 == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (v_6_3 == 380165415) {
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
})})))}
})
}

func Call_Data_Functor_Product2_bifunctorProduct2(dictBifunctor_0_loop gopurs_runtime.Value, dictBifunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
var dictBifunctor1_1 gopurs_runtime.Value = dictBifunctor1_1_loop
_ = dictBifunctor1_1
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_2864580459_1688994542((&Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_2, g_3, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_1, "bimap"), f_2, g_3, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)}))}
})})))}
}

func Call_Data_Functor_Product2_biapplyProduct2(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
_ = dictBiapply_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorProduct21__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
bifunctorProduct21__193435443_1_0 := gopurs_runtime.Func(func(dictBifunctor1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_2864580459_1688994542((&Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), f_3, g_4, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_2, "bimap"), f_3, g_4, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)}))}
})})))}
})
_ = bifunctorProduct21__193435443_1_0
return gopurs_runtime.Func(func(dictBiapply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): bifunctorProduct22_3_2 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(ADT ["Data","Functor","Product2","Product2"] [(TypeVar f), (TypeVar g)])])
bifunctorProduct22_3_2 := Rebox_Data_Functor_Product2_1688994542_2864580459(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(bifunctorProduct21__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply1_2, "Bifunctor0"), gopurs_runtime.Value{}))))
_ = bifunctorProduct22_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_382698396_2448899513((&Constructor_Control_Biapply_Biapply[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_2864580459_1688994542(bifunctorProduct22_3_2))}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply1_2, "biapply"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}))}
})})))}
})
}

func Call_Data_Functor_Product2_biapplicativeProduct2(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
_ = dictBiapplicative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifunctorProduct21__193435443_2_2 shape=Let(Abs(LitRecord)) bindingType=Any
bifunctorProduct21__193435443_2_2 := gopurs_runtime.Func(func(dictBifunctor1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_2864580459_1688994542((&Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "bimap"), f_4, g_5, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_3, "bimap"), f_4, g_5, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)}))}
})})))}
})
_ = bifunctorProduct21__193435443_2_2
// TAST (Let): biapplyProduct21__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
biapplyProduct21__193435443_1_0 := gopurs_runtime.Func(func(dictBiapply1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): bifunctorProduct22_4_4 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(ADT ["Data","Functor","Product2","Product2"] [(TypeVar f), (TypeVar g)])])
bifunctorProduct22_4_4 := Rebox_Data_Functor_Product2_1688994542_2864580459(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(bifunctorProduct21__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply1_3, "Bifunctor0"), gopurs_runtime.Value{}))))
_ = bifunctorProduct22_4_4
return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_382698396_2448899513((&Constructor_Control_Biapply_Biapply[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_2864580459_1688994542(bifunctorProduct22_4_4))}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "biapply"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply1_3, "biapply"), (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}))}
})})))}
})
_ = biapplyProduct21__193435443_1_0
return gopurs_runtime.Func(func(dictBiapplicative1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): biapplyProduct22_3_5 shape=App(Other) bindingType=(ADT ["Control","Biapply","Biapply"] [(ADT ["Data","Functor","Product2","Product2"] [(TypeVar f), (TypeVar g)])])
biapplyProduct22_3_5 := Rebox_Data_Functor_Product2_2448899513_382698396(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](gopurs_runtime.Apply(biapplyProduct21__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative1_2, "Biapply0"), gopurs_runtime.Value{}))))
_ = biapplyProduct22_3_5
return gopurs_runtime.Value{Type: 9, IntVal: 3949191309, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_1380128092_3565318777((&Constructor_Control_Biapplicative_Biapplicative[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product2_382698396_2448899513(biapplyProduct22_3_5))}
}), gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), a_4, b_5), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative1_2, "bipure"), a_4, b_5)}))}
})})))}
})
}

func Rebox_Data_Functor_Product2_106768939_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Functor_Product2_1380128092_3565318777(in *Constructor_Control_Biapplicative_Biapplicative[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Functor_Product2_1688994542_2864580459(in *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Functor_Product2_2448899513_382698396(in *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]) *Constructor_Control_Biapply_Biapply[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Biapply_Biapply[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Functor_Product2_2864580459_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Functor_Product2_3352668427_4074365774(in *Constructor_Data_Profunctor_Profunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Functor_Product2_3354170411_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Functor_Product2_3790796878_4204058379(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Functor_Product2_382698396_2448899513(in *Constructor_Control_Biapply_Biapply[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Functor_Product2_4204058379_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Functor_Product2_686435819_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


