package Data_Number_Approximate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Number "gopurs/output/Data.Number"
)

var Tolerance gopurs_runtime.Value
var once_Tolerance sync.Once
func Get_Tolerance() gopurs_runtime.Value {
	once_Tolerance.Do(func() {
		Tolerance = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Tolerance
}

var Fraction gopurs_runtime.Value
var once_Fraction sync.Once
func Get_Fraction() gopurs_runtime.Value {
	once_Fraction.Do(func() {
		Fraction = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Fraction
}

var eqRelative gopurs_runtime.Value
var once_eqRelative sync.Once
func Get_eqRelative() gopurs_runtime.Value {
	once_eqRelative.Do(func() {
		eqRelative = gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.FloatEq(v1_1, gopurs_runtime.Float(0.0))).IntVal != 0 {
__t0 = gopurs_runtime.FloatLte(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), v2_2), v_0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.FloatEq(v2_2, gopurs_runtime.Float(0.0))).IntVal != 0 {
__t0 = gopurs_runtime.FloatLte(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), v1_1), v_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.FloatLte(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.FloatSub(v1_1, v2_2)), gopurs_runtime.FloatDiv(gopurs_runtime.FloatMul(v_0, gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.FloatAdd(v1_1, v2_2))), gopurs_runtime.Float(2.0)))
}
end_branch_0:
return __t0
})
	})
	return eqRelative
}

var eqApproximate gopurs_runtime.Value
var once_eqApproximate sync.Once
func Get_eqApproximate() gopurs_runtime.Value {
	once_eqApproximate.Do(func() {
		eqApproximate = gopurs_runtime.Apply(Get_eqRelative(), gopurs_runtime.Float(0.000001))
	})
	return eqApproximate
}

var neqApproximate gopurs_runtime.Value
var once_neqApproximate sync.Once
func Get_neqApproximate() gopurs_runtime.Value {
	once_neqApproximate.Do(func() {
		neqApproximate = gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply3(Get_eqRelative(), gopurs_runtime.Float(0.000001), x_0, y_1).IntVal == 0)
})
	})
	return neqApproximate
}

var eqAbsolute gopurs_runtime.Value
var once_eqAbsolute sync.Once
func Get_eqAbsolute() gopurs_runtime.Value {
	once_eqAbsolute.Do(func() {
		eqAbsolute = gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.FloatLte(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.FloatSub(x_1, y_2)), v_0)
})
	})
	return eqAbsolute
}


