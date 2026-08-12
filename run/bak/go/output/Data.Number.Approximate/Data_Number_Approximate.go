package Data_Number_Approximate

import (
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Tolerance gopurs_runtime.Value
var once_Tolerance sync.Once
func Get_Tolerance() gopurs_runtime.Value {
	once_Tolerance.Do(func() {
		cache_Tolerance = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Tolerance(x_0_box)
})
	})
	return cache_Tolerance
}

var cache_Fraction gopurs_runtime.Value
var once_Fraction sync.Once
func Get_Fraction() gopurs_runtime.Value {
	once_Fraction.Do(func() {
		cache_Fraction = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Fraction(x_0_box)
})
	})
	return cache_Fraction
}

var cache_eqRelative gopurs_runtime.Value
var once_eqRelative sync.Once
func Get_eqRelative() gopurs_runtime.Value {
	once_eqRelative.Do(func() {
		cache_eqRelative = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_eqRelative(v_0_box.FloatVal(), v1_1_box.FloatVal(), v2_2_box.FloatVal()))
})
	})
	return cache_eqRelative
}

var cache_eqApproximate gopurs_runtime.Value
var once_eqApproximate sync.Once
func Get_eqApproximate() gopurs_runtime.Value {
	once_eqApproximate.Do(func() {
		cache_eqApproximate = gopurs_runtime.Apply(Get_eqRelative(), gopurs_runtime.Float(0.000001))
	})
	return cache_eqApproximate
}

var cache_neqApproximate gopurs_runtime.Value
var once_neqApproximate sync.Once
func Get_neqApproximate() gopurs_runtime.Value {
	once_neqApproximate.Do(func() {
		cache_neqApproximate = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_neqApproximate(x_0_box.FloatVal(), y_1_box.FloatVal()))
})
	})
	return cache_neqApproximate
}

var cache_eqAbsolute gopurs_runtime.Value
var once_eqAbsolute sync.Once
func Get_eqAbsolute() gopurs_runtime.Value {
	once_eqAbsolute.Do(func() {
		cache_eqAbsolute = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_eqAbsolute(v_0_box.FloatVal(), x_1_box.FloatVal(), y_2_box.FloatVal()))
})
	})
	return cache_eqAbsolute
}

var cache_div__1002719800 gopurs_runtime.Value
var once_div__1002719800 sync.Once
func Get_div__1002719800() gopurs_runtime.Value {
	once_div__1002719800.Do(func() {
		cache_div__1002719800 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div")
	})
	return cache_div__1002719800
}

var cache_div__2579358968 gopurs_runtime.Value
var once_div__2579358968 sync.Once
func Get_div__2579358968() gopurs_runtime.Value {
	once_div__2579358968.Do(func() {
		cache_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__2579358968
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
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
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
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

var cache_lessThanOrEq__1061005983 gopurs_runtime.Value
var once_lessThanOrEq__1061005983 sync.Once
func Get_lessThanOrEq__1061005983() gopurs_runtime.Value {
	once_lessThanOrEq__1061005983.Do(func() {
		cache_lessThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__1061005983
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_sub__1135378904 gopurs_runtime.Value
var once_sub__1135378904 sync.Once
func Get_sub__1135378904() gopurs_runtime.Value {
	once_sub__1135378904.Do(func() {
		cache_sub__1135378904 = pkg_Data_Ring.Get_numSub()
	})
	return cache_sub__1135378904
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

var cache_add__137136408 gopurs_runtime.Value
var once_add__137136408 sync.Once
func Get_add__137136408() gopurs_runtime.Value {
	once_add__137136408.Do(func() {
		cache_add__137136408 = pkg_Data_Semiring.Get_numAdd()
	})
	return cache_add__137136408
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

var cache_mul__137136408 gopurs_runtime.Value
var once_mul__137136408 sync.Once
func Get_mul__137136408() gopurs_runtime.Value {
	once_mul__137136408.Do(func() {
		cache_mul__137136408 = pkg_Data_Semiring.Get_numMul()
	})
	return cache_mul__137136408
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

func Call_Tolerance(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Fraction(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_eqRelative(v_0_loop float64, v1_1_loop float64, v2_2_loop float64) bool {
var v_0 float64 = v_0_loop
_ = v_0
var v1_1 float64 = v1_1_loop
_ = v1_1
var v2_2 float64 = v2_2_loop
_ = v2_2
var __t0 bool
{
if (v1_1) == (0.0) {
__t0 = (gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(v2_2)).FloatVal()), gopurs_runtime.Float(v_0))).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
if (v2_2) == (0.0) {
__t0 = (gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(v1_1)).FloatVal()), gopurs_runtime.Float(v_0))).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = (gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_sub__1135378904(), gopurs_runtime.Float(v1_1), gopurs_runtime.Float(v2_2)).FloatVal())).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_div__1002719800(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(v_0), gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_add__137136408(), gopurs_runtime.Float(v1_1), gopurs_runtime.Float(v2_2)).FloatVal())).FloatVal())).FloatVal()), gopurs_runtime.Float(2.0)).FloatVal()))).IntVal) != (0)
}
end_branch_0:
return __t0
}

func Call_neqApproximate(x_0_loop float64, y_1_loop float64) bool {
var x_0 float64 = x_0_loop
_ = x_0
var y_1 float64 = y_1_loop
_ = y_1
return (gopurs_runtime.Apply(Get_not__3201284355(), gopurs_runtime.Bool((gopurs_runtime.Bool(Call_eqRelative(0.000001, x_0, y_1)).IntVal) != (0))).IntVal) != (0)
}

func Call_eqAbsolute(v_0_loop float64, x_1_loop float64, y_2_loop float64) bool {
var v_0 float64 = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
var y_2 float64 = y_2_loop
_ = y_2
return (gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_sub__1135378904(), gopurs_runtime.Float(x_1), gopurs_runtime.Float(y_2)).FloatVal())).FloatVal()), gopurs_runtime.Float(v_0))).IntVal) != (0)
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_lessThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) > (a2_1.FloatVal()) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


