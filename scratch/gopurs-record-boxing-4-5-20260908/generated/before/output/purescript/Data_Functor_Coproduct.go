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
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
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
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
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
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
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
		cache_Data_Functor_Coproduct_eqCoproduct = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_eqCoproduct(dictEq1_0_box, dictEq11_1_box, dictEq_2_box)
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
				if _v.Type == 9 && _v.IntVal == 3234899973 && _v.UnsafePtr != nil {
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
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
				if _v.Type == 9 && _v.IntVal == 3234899973 && _v.UnsafePtr != nil {
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
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
				if _v.Type == 9 && _v.IntVal == 3234899973 && _v.UnsafePtr != nil {
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
// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=Any
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2)
_ = __local_var_4_0
// TAST (Let): __local_var_5_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2)
_ = __local_var_5_1
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 3711209382) {
__t2 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply(__local_var_4_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
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
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
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

func Call_Data_Functor_Coproduct_eqCoproduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Functor_Coproduct_ord1Coproduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eq1Coproduct1__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
eq1Coproduct1__193435443_1_0 := gopurs_runtime.Func(func(dictEq11_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictEq_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 bool
{
if (v_4.Type == 9 && v_4.IntVal == 3711209382) {
__t2 = ((v1_5.Type == 9 && v1_5.IntVal == 3711209382)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_3))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0))
goto end_branch_2
} else {

}
}
{
__t2 = ((v_4.Type == 9 && v_4.IntVal == 2465973597)) && (((v1_5.Type == 9 && v1_5.IntVal == 2465973597)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_3))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})}))}
})
_ = eq1Coproduct1__193435443_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eq1Coproduct2_3_3 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
eq1Coproduct2_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](gopurs_runtime.Apply(eq1Coproduct1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})))
_ = eq1Coproduct2_3_3
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Coproduct2_3_3)}
}), gopurs_runtime.Func3(func(dictOrd_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 uint32
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
var __t4 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t4 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 2465973597)) && ((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) {
__t5 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})}))}
})
}

func Call_Data_Functor_Coproduct_ordCoproduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eq1Coproduct1__193435443_1_1 shape=Let(Abs(LitRecord)) bindingType=Any
eq1Coproduct1__193435443_1_1 := gopurs_runtime.Func(func(dictEq11_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictEq_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
if (v_4.Type == 9 && v_4.IntVal == 3711209382) {
__t3 = ((v1_5.Type == 9 && v1_5.IntVal == 3711209382)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_3))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0))
goto end_branch_3
} else {

}
}
{
__t3 = ((v_4.Type == 9 && v_4.IntVal == 2465973597)) && (((v1_5.Type == 9 && v1_5.IntVal == 2465973597)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_3))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})}))}
})
_ = eq1Coproduct1__193435443_1_1
// TAST (Let): ord1Coproduct1__193435443_1_0 shape=Let(Abs(Let(LitRecord))) bindingType=Any
ord1Coproduct1__193435443_1_0 := gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eq1Coproduct2_3_4 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
eq1Coproduct2_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](gopurs_runtime.Apply(eq1Coproduct1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})))
_ = eq1Coproduct2_3_4
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Coproduct2_3_4)}
}), gopurs_runtime.Func3(func(dictOrd_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 uint32
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
var __t5 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t5 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 2465973597)) && ((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) {
__t6 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)
goto end_branch_6
} else {

}
}
{
__t6 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t6), UnsafePtr: nil}
})}))}
})
_ = ord1Coproduct1__193435443_1_0
// TAST (Let): __local_var_2_8 shape=App(Other) bindingType=Any
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_8
// TAST (Let): eqCoproduct1__193435443_2_7 shape=Let(Abs(Abs(LitRecord))) bindingType=Any
eqCoproduct1__193435443_2_7 := gopurs_runtime.Func2(func(dictEq11_3 gopurs_runtime.Value, dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 bool
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
__t9 = ((v1_6.Type == 9 && v1_6.IntVal == 3711209382)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_8, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_4))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0))
goto end_branch_9
} else {

}
}
{
__t9 = ((v_5.Type == 9 && v_5.IntVal == 2465973597)) && (((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_3, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_4))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
})}))}
})
_ = eqCoproduct1__193435443_2_7
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): compare1__193435443_4_10 shape=Other bindingType=Any
compare1__193435443_4_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(ord1Coproduct1__193435443_1_0, dictOrd11_3), "compare1")
_ = compare1__193435443_4_10
// TAST (Let): eqCoproduct2__193435443_5_11 shape=App(Other) bindingType=Any
eqCoproduct2__193435443_5_11 := gopurs_runtime.Apply(eqCoproduct1__193435443_2_7, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{}))
_ = eqCoproduct2__193435443_5_11
return gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqCoproduct3_7_12 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar a)])])
eqCoproduct3_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqCoproduct2__193435443_5_11, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_6, "Eq0"), gopurs_runtime.Value{})))
_ = eqCoproduct3_7_12
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqCoproduct3_7_12)}
}), gopurs_runtime.Apply(compare1__193435443_4_10, dictOrd_6)}))}
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
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorCoproduct1__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
functorCoproduct1__193435443_1_0 := gopurs_runtime.Func(func(dictFunctor1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=Any
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_3)
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_2, "map"), f_3)
_ = __local_var_6_3
var __t4 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 3711209382) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply(__local_var_5_2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2465973597) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(__local_var_6_3, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})}))}
})
_ = functorCoproduct1__193435443_1_0
return gopurs_runtime.Func(func(dictExtend1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCoproduct2_3_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCoproduct2_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCoproduct1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "Functor0"), gopurs_runtime.Value{})))
_ = functorCoproduct2_3_5
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCoproduct2_3_5)}
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)]))
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_5, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
_ = __local_var_5_7
// TAST (Let): __local_var_6_9 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_6, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
_ = __local_var_6_9
// TAST (Let): __local_var_6_8 shape=Let(Abs(Other)) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar b)]), (TypeApp (TypeVar g) [(TypeVar b)])]))
__local_var_6_8 := gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_9, x_7)}))}
})
_ = __local_var_6_8
// TAST (Let): __local_var_5_6 shape=Let(Let(Abs(Branch(Other, App(Other), def=Other)))) bindingType=(Func [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar a)])] (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar b)]), (TypeApp (TypeVar g) [(TypeVar b)])]))
__local_var_5_6 := gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_5_7, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)}))}
goto end_branch_10
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t10 = gopurs_runtime.Apply(__local_var_6_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
_ = __local_var_5_6
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_6, x_6)
})
})}))}
})
}

func Call_Data_Functor_Coproduct_comonadCoproduct(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorCoproduct1__193435443_2_2 shape=Let(Abs(LitRecord)) bindingType=Any
functorCoproduct1__193435443_2_2 := gopurs_runtime.Func(func(dictFunctor1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=Any
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_4)
_ = __local_var_6_4
// TAST (Let): __local_var_7_5 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_3, "map"), f_4)
_ = __local_var_7_5
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
__t6 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply(__local_var_6_4, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 2465973597) {
__t6 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(__local_var_7_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})}))}
})
_ = functorCoproduct1__193435443_2_2
// TAST (Let): extendCoproduct1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
extendCoproduct1__193435443_1_0 := gopurs_runtime.Func(func(dictExtend1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCoproduct2_4_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCoproduct2_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCoproduct1__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_3, "Functor0"), gopurs_runtime.Value{})))
_ = functorCoproduct2_4_7
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCoproduct2_4_7)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_9 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)]))
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_6, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
_ = __local_var_6_9
// TAST (Let): __local_var_7_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
__local_var_7_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_3, "extend"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
_ = __local_var_7_11
// TAST (Let): __local_var_7_10 shape=Let(Abs(Other)) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar b)]), (TypeApp (TypeVar g) [(TypeVar b)])]))
__local_var_7_10 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_7_11, x_8)}))}
})
_ = __local_var_7_10
// TAST (Let): __local_var_6_8 shape=Let(Let(Abs(Branch(Other, App(Other), def=Other)))) bindingType=(Func [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar a)])] (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar b)]), (TypeApp (TypeVar g) [(TypeVar b)])]))
__local_var_6_8 := gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_9, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)}))}
goto end_branch_12
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t12 = gopurs_runtime.Apply(__local_var_7_10, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
_ = __local_var_6_8
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_8, x_7)
})
})}))}
})
_ = extendCoproduct1__193435443_1_0
return gopurs_runtime.Func(func(dictComonad1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): extendCoproduct2_3_13 shape=App(Other) bindingType=(ADT ["Control","Extend","Extend"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
extendCoproduct2_3_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](gopurs_runtime.Apply(extendCoproduct1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "Extend0"), gopurs_runtime.Value{})))
_ = extendCoproduct2_3_13
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendCoproduct2_3_13)}
}), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "extract"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
})}))}
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
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
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
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
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


