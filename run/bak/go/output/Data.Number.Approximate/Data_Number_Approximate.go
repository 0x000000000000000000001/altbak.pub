package Data_Number_Approximate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
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
var __t1 gopurs_runtime.Value
{
if (v1_1) == (0.0) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(v2_2)).FloatVal()) > (v_0) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(true)
}
end_branch_2:
__t1 = gopurs_runtime.Bool((__t2.IntVal) != (0))
goto end_branch_1
} else {

}
}
{
if (v2_2) == (0.0) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(v1_1)).FloatVal()) > (v_0) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(true)
}
end_branch_3:
__t1 = gopurs_runtime.Bool((__t3.IntVal) != (0))
goto end_branch_1
} else {

}
}
{
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float((v1_1) - (v2_2))).FloatVal()) > (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), gopurs_runtime.Float((v_0) * (gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float((v1_1) + (v2_2))).FloatVal())), gopurs_runtime.Float(2.0)).FloatVal()) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(true)
}
end_branch_0:
__t1 = gopurs_runtime.Bool((__t0.IntVal) != (0))
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_neqApproximate(x_0_loop float64, y_1_loop float64) bool {
var x_0 float64 = x_0_loop
_ = x_0
var y_1 float64 = y_1_loop
_ = y_1
return (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Bool(Call_eqRelative(0.000001, x_0, y_1))).IntVal) != (0)
}

func Call_eqAbsolute(v_0_loop float64, x_1_loop float64, y_2_loop float64) bool {
var v_0 float64 = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
var y_2 float64 = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float((x_1) - (y_2))).FloatVal()) > (v_0) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(true)
}
end_branch_0:
return (__t0.IntVal) != (0)
}


