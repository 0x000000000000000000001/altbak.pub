package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Coproduct_Coproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_Coproduct sync.Once
func Get_Data_Functor_Coproduct_Coproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Coproduct.Do(func() {
		cache_Data_Functor_Coproduct_Coproduct = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_Coproduct_Coproduct(x_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
})
	})
	return cache_Data_Functor_Coproduct_Coproduct
}

var cache_Data_Functor_Coproduct_showCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_showCoproduct sync.Once
func Get_Data_Functor_Coproduct_showCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_showCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_showCoproduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_showCoproduct(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Functor_Coproduct_showCoproduct
}

var cache_Data_Functor_Coproduct_right gopurs_runtime.Value
var once_Data_Functor_Coproduct_right sync.Once
func Get_Data_Functor_Coproduct_right() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right.Do(func() {
		cache_Data_Functor_Coproduct_right = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_Coproduct_right(ga_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
})
	})
	return cache_Data_Functor_Coproduct_right
}

var cache_Data_Functor_Coproduct_newtypeCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_newtypeCoproduct sync.Once
func Get_Data_Functor_Coproduct_newtypeCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_newtypeCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_newtypeCoproduct = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Data_Functor_Coproduct_newtypeCoproduct
}

var cache_Data_Functor_Coproduct_left gopurs_runtime.Value
var once_Data_Functor_Coproduct_left sync.Once
func Get_Data_Functor_Coproduct_left() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_left.Do(func() {
		cache_Data_Functor_Coproduct_left = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_Coproduct_left(fa_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
})
	})
	return cache_Data_Functor_Coproduct_left
}

var cache_Data_Functor_Coproduct_functorCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_functorCoproduct sync.Once
func Get_Data_Functor_Coproduct_functorCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_functorCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_functorCoproduct = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_functorCoproduct(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_Data_Functor_Coproduct_functorCoproduct
}

var cache_Data_Functor_Coproduct_eq1Coproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_eq1Coproduct sync.Once
func Get_Data_Functor_Coproduct_eq1Coproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_eq1Coproduct.Do(func() {
		cache_Data_Functor_Coproduct_eq1Coproduct = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_eq1Coproduct(dictEq1_0_box, dictEq11_1_box)
})
	})
	return cache_Data_Functor_Coproduct_eq1Coproduct
}

var cache_Data_Functor_Coproduct_eqCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_eqCoproduct sync.Once
func Get_Data_Functor_Coproduct_eqCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_eqCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_eqCoproduct = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_eqCoproduct(dictEq1_0_box, dictEq11_1_box)
})
	})
	return cache_Data_Functor_Coproduct_eqCoproduct
}

var cache_Data_Functor_Coproduct_ord1Coproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_ord1Coproduct sync.Once
func Get_Data_Functor_Coproduct_ord1Coproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_ord1Coproduct.Do(func() {
		cache_Data_Functor_Coproduct_ord1Coproduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_ord1Coproduct(dictOrd1_0_box)
})
	})
	return cache_Data_Functor_Coproduct_ord1Coproduct
}

var cache_Data_Functor_Coproduct_ordCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_ordCoproduct sync.Once
func Get_Data_Functor_Coproduct_ordCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_ordCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_ordCoproduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_ordCoproduct(dictOrd1_0_box)
})
	})
	return cache_Data_Functor_Coproduct_ordCoproduct
}

var cache_Data_Functor_Coproduct_coproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_coproduct sync.Once
func Get_Data_Functor_Coproduct_coproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_coproduct.Do(func() {
		cache_Data_Functor_Coproduct_coproduct = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_coproduct(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Functor_Coproduct_coproduct
}

var cache_Data_Functor_Coproduct_extendCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_extendCoproduct sync.Once
func Get_Data_Functor_Coproduct_extendCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_extendCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_extendCoproduct = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_extendCoproduct(dictExtend_0_box)
})
	})
	return cache_Data_Functor_Coproduct_extendCoproduct
}

var cache_Data_Functor_Coproduct_comonadCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_comonadCoproduct sync.Once
func Get_Data_Functor_Coproduct_comonadCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_comonadCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_comonadCoproduct = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_comonadCoproduct(dictComonad_0_box)
})
	})
	return cache_Data_Functor_Coproduct_comonadCoproduct
}

var cache_Data_Functor_Coproduct_bihoistCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_bihoistCoproduct sync.Once
func Get_Data_Functor_Coproduct_bihoistCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_bihoistCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_bihoistCoproduct = gopurs_runtime.Func3(func(natF_0_box gopurs_runtime.Value, natG_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_bihoistCoproduct(natF_0_box, natG_1_box, v_2_box)
})
	})
	return cache_Data_Functor_Coproduct_bihoistCoproduct
}

func Call_Data_Functor_Coproduct_Coproduct(x_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := x_0
				if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}

func Call_Data_Functor_Coproduct_showCoproduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = (("(left ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = (("(right ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
__t0 = func() string { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
})}))}
}

func Call_Data_Functor_Coproduct_right(ga_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, ga_0, true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
				if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}

func Call_Data_Functor_Coproduct_left(fa_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{fa_0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
				if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}

func Call_Data_Functor_Coproduct_functorCoproduct(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)]))
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2)
_ = __local_var_4_0
// TAST (Let): __local_var_5_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g$scope57) [(TypeVar a$scope61)])] (TypeApp (TypeVar g$scope57) [(TypeVar b$scope62)]))
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2)
_ = __local_var_5_1
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 3711209382) {
__t2 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply(__local_var_4_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2465973597) {
__t2 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(__local_var_5_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})}))}
}

func Call_Data_Functor_Coproduct_eq1Coproduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictEq_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 bool
{
if (v_3.Type == 9 && v_3.IntVal == 3711209382) {
__t0 = ((v1_4.Type == 9 && v1_4.IntVal == 3711209382)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = ((v_3.Type == 9 && v_3.IntVal == 2465973597)) && (((v1_4.Type == 9 && v1_4.IntVal == 2465973597)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_0:
return gopurs_runtime.Bool(__t0)
})}))}
}

func Call_Data_Functor_Coproduct_eqCoproduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
// TAST (Let): eq1_2_0 shape=App(Var) bindingType=Any
eq1_2_0 := Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Data_Functor_Coproduct_eq1Coproduct(dictEq1_0, dictEq11_1)))
_ = eq1_2_0
return gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(eq1_2_0, dictEq_3)}))}
})
}

func Call_Data_Functor_Coproduct_ord1Coproduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): eq1Coproduct1_1_0 shape=App(Var) bindingType=Any
eq1Coproduct1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Coproduct_eq1Coproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
_ = eq1Coproduct1_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eq1Coproduct2_3_1 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope28), (TypeVar g$scope29)])])
eq1Coproduct2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](gopurs_runtime.Apply(eq1Coproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})))
_ = eq1Coproduct2_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Coproduct2_3_1)}
}), gopurs_runtime.Func3(func(dictOrd_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 uint32
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
var __t2 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t2 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 2465973597)) && ((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) {
__t3 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})}))}
})
}

func Call_Data_Functor_Coproduct_ordCoproduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): ord1Coproduct1_1_0 shape=App(Var) bindingType=Any
ord1Coproduct1_1_0 := Call_Data_Functor_Coproduct_ord1Coproduct(dictOrd1_0)
_ = ord1Coproduct1_1_0
// TAST (Let): eqCoproduct1_2_1 shape=App(Var) bindingType=Any
eqCoproduct1_2_1 := gopurs_runtime.Apply(Get_Data_Functor_Coproduct_eqCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
_ = eqCoproduct1_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): compare1_4_2 shape=App(Var) bindingType=Any
compare1_4_2 := Call_Data_Ord_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](gopurs_runtime.Apply(ord1Coproduct1_1_0, dictOrd11_3)))
_ = compare1_4_2
// TAST (Let): eqCoproduct2_5_3 shape=App(Other) bindingType=Any
eqCoproduct2_5_3 := gopurs_runtime.Apply(eqCoproduct1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{}))
_ = eqCoproduct2_5_3
return gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqCoproduct3_7_4 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope22), (TypeVar g$scope23), (TypeVar a$scope24)])])
eqCoproduct3_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqCoproduct2_5_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_6, "Eq0"), gopurs_runtime.Value{})))
_ = eqCoproduct3_7_4
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqCoproduct3_7_4)}
}), gopurs_runtime.Apply(compare1_4_2, dictOrd_6)}))}
})
})
}

func Call_Data_Functor_Coproduct_coproduct(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Functor_Coproduct_extendCoproduct(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
// TAST (Let): functorCoproduct1_1_0 shape=App(Var) bindingType=Any
functorCoproduct1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Coproduct_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct1_1_0
return gopurs_runtime.Func(func(dictExtend1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCoproduct2_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope91), (TypeVar g$scope92)])])
functorCoproduct2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCoproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "Functor0"), gopurs_runtime.Value{})))
_ = functorCoproduct2_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCoproduct2_3_1)}
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Either_Left(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), Get_Data_Either_Left())))), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Either_Right(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "extend"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), Get_Data_Either_Right()))))))
})}))}
})
}

func Call_Data_Functor_Coproduct_comonadCoproduct(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): extract_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope128) [(TypeVar a$scope132)])] (TypeVar a$scope132))
extract_1_0 := Call_Control_Comonad_extract(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dictComonad_0))
_ = extract_1_0
// TAST (Let): extendCoproduct1_2_1 shape=App(Var) bindingType=Any
extendCoproduct1_2_1 := Call_Data_Functor_Coproduct_extendCoproduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}))
_ = extendCoproduct1_2_1
return gopurs_runtime.Func(func(dictComonad1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): extendCoproduct2_4_2 shape=App(Other) bindingType=(ADT ["Control","Extend","Extend"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope128), (TypeVar g$scope129)])])
extendCoproduct2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](gopurs_runtime.Apply(extendCoproduct1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_3, "Extend0"), gopurs_runtime.Value{})))
_ = extendCoproduct2_4_2
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendCoproduct2_4_2)}
}), gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), extract_1_0, Call_Control_Comonad_extract(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dictComonad1_3)))}))}
})
}

func Call_Data_Functor_Coproduct_bihoistCoproduct(natF_0_loop gopurs_runtime.Value, natG_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var natF_0 gopurs_runtime.Value = natF_0_loop
_ = natF_0
var natG_1 gopurs_runtime.Value = natG_1_loop
_ = natG_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply(natF_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(natG_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
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


