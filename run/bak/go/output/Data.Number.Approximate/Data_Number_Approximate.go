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
		Tolerance = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Tolerance
}

var Fraction gopurs_runtime.Value
var once_Fraction sync.Once
func Get_Fraction() gopurs_runtime.Value {
	once_Fraction.Do(func() {
		Fraction = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Fraction
}

var eqRelative gopurs_runtime.Value
var once_eqRelative sync.Once
func Get_eqRelative() gopurs_runtime.Value {
	once_eqRelative.Do(func() {
		eqRelative = gopurs_runtime.Func3(Call_eqRelative)
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
		neqApproximate = gopurs_runtime.Func2(Call_neqApproximate)
	})
	return neqApproximate
}

var eqAbsolute gopurs_runtime.Value
var once_eqAbsolute sync.Once
func Get_eqAbsolute() gopurs_runtime.Value {
	once_eqAbsolute.Do(func() {
		eqAbsolute = gopurs_runtime.Func3(Call_eqAbsolute)
	})
	return eqAbsolute
}

func Call_eqRelative(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if v1_1_loop.FloatVal() == 0.0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), v2_2_loop).FloatVal() <= v_0_loop.FloatVal())
goto end_branch_0
} else {

}
}
{
if v2_2_loop.FloatVal() == 0.0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), v1_1_loop).FloatVal() <= v_0_loop.FloatVal())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(v1_1_loop.FloatVal() - v2_2_loop.FloatVal())).FloatVal() <= v_0_loop.FloatVal() * gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(v1_1_loop.FloatVal() + v2_2_loop.FloatVal())).FloatVal() / 2.0)
}
end_branch_0:
return __t0
}

func Call_neqApproximate(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return gopurs_runtime.Bool(Call_eqRelative(gopurs_runtime.Float(0.000001), x_0_loop, y_1_loop).IntVal != 0 != true)
}

func Call_eqAbsolute(v_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Float(x_1_loop.FloatVal() - y_2_loop.FloatVal())).FloatVal() <= v_0_loop.FloatVal())
}


