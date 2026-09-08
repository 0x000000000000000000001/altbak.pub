package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Number_Approximate_Tolerance gopurs_runtime.Value
var once_Data_Number_Approximate_Tolerance sync.Once
func Get_Data_Number_Approximate_Tolerance() gopurs_runtime.Value {
	once_Data_Number_Approximate_Tolerance.Do(func() {
		cache_Data_Number_Approximate_Tolerance = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Number_Approximate_Tolerance(x_0_box.FloatVal()))
})
	})
	return cache_Data_Number_Approximate_Tolerance
}

var cache_Data_Number_Approximate_Fraction gopurs_runtime.Value
var once_Data_Number_Approximate_Fraction sync.Once
func Get_Data_Number_Approximate_Fraction() gopurs_runtime.Value {
	once_Data_Number_Approximate_Fraction.Do(func() {
		cache_Data_Number_Approximate_Fraction = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Number_Approximate_Fraction(x_0_box.FloatVal()))
})
	})
	return cache_Data_Number_Approximate_Fraction
}

var cache_Data_Number_Approximate_eqRelative gopurs_runtime.Value
var once_Data_Number_Approximate_eqRelative sync.Once
func Get_Data_Number_Approximate_eqRelative() gopurs_runtime.Value {
	once_Data_Number_Approximate_eqRelative.Do(func() {
		cache_Data_Number_Approximate_eqRelative = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Number_Approximate_eqRelative(v_0_box.FloatVal(), v1_1_box.FloatVal(), v2_2_box.FloatVal()))
})
	})
	return cache_Data_Number_Approximate_eqRelative
}

var cache_Data_Number_Approximate_eqApproximate gopurs_runtime.Value
var once_Data_Number_Approximate_eqApproximate sync.Once
func Get_Data_Number_Approximate_eqApproximate() gopurs_runtime.Value {
	once_Data_Number_Approximate_eqApproximate.Do(func() {
		cache_Data_Number_Approximate_eqApproximate = gopurs_runtime.Apply(Get_Data_Number_Approximate_eqRelative(), gopurs_runtime.Float(0.000001))
	})
	return cache_Data_Number_Approximate_eqApproximate
}

var cache_Data_Number_Approximate_neqApproximate gopurs_runtime.Value
var once_Data_Number_Approximate_neqApproximate sync.Once
func Get_Data_Number_Approximate_neqApproximate() gopurs_runtime.Value {
	once_Data_Number_Approximate_neqApproximate.Do(func() {
		cache_Data_Number_Approximate_neqApproximate = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Number_Approximate_neqApproximate(x_0_box.FloatVal(), y_1_box.FloatVal()))
})
	})
	return cache_Data_Number_Approximate_neqApproximate
}

var cache_Data_Number_Approximate_eqAbsolute gopurs_runtime.Value
var once_Data_Number_Approximate_eqAbsolute sync.Once
func Get_Data_Number_Approximate_eqAbsolute() gopurs_runtime.Value {
	once_Data_Number_Approximate_eqAbsolute.Do(func() {
		cache_Data_Number_Approximate_eqAbsolute = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Number_Approximate_eqAbsolute(v_0_box.FloatVal(), x_1_box.FloatVal(), y_2_box.FloatVal()))
})
	})
	return cache_Data_Number_Approximate_eqAbsolute
}

func Call_Data_Number_Approximate_Tolerance(x_0_loop float64) float64 {
var x_0 float64 = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Number_Approximate_Fraction(x_0_loop float64) float64 {
var x_0 float64 = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Number_Approximate_eqRelative(v_0_loop float64, v1_1_loop float64, v2_2_loop float64) bool {
var v_0 float64 = v_0_loop
_ = v_0
var v1_1 float64 = v1_1_loop
_ = v1_1
var v2_2 float64 = v2_2_loop
_ = v2_2
var __t3 bool
{
if (v1_1) == (0.0) {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Apply(Get_Data_Number_abs(), gopurs_runtime.Float(v2_2)), gopurs_runtime.Float(v_0))
__t3 = ((uint32(__t_tag_1.IntVal) == 380165415)) != (true)
goto end_branch_3
} else {

}
}
{
if (v2_2) == (0.0) {
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Apply(Get_Data_Number_abs(), gopurs_runtime.Float(v1_1)), gopurs_runtime.Float(v_0))
__t3 = ((uint32(__t_tag_2.IntVal) == 380165415)) != (true)
goto end_branch_3
} else {

}
}
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Apply(Get_Data_Number_abs(), gopurs_runtime.Float((v1_1) - (v2_2))), gopurs_runtime.Float(((v_0) * (gopurs_runtime.Apply(Get_Data_Number_abs(), gopurs_runtime.Float((v1_1) + (v2_2))).FloatVal())) / (2.0)))
__t3 = ((uint32(__t_tag_0.IntVal) == 380165415)) != (true)
}
end_branch_3:
return __t3
}

func Call_Data_Number_Approximate_neqApproximate(x_0_loop float64, y_1_loop float64) bool {
var x_0 float64 = x_0_loop
_ = x_0
var y_1 float64 = y_1_loop
_ = y_1
return (Call_Data_Number_Approximate_eqRelative(0.000001, x_0, y_1)) != (true)
}

func Call_Data_Number_Approximate_eqAbsolute(v_0_loop float64, x_1_loop float64, y_2_loop float64) bool {
var v_0 float64 = v_0_loop
_ = v_0
var x_1 float64 = x_1_loop
_ = x_1
var y_2 float64 = y_2_loop
_ = y_2
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Apply(Get_Data_Number_abs(), gopurs_runtime.Float((x_1) - (y_2))), gopurs_runtime.Float(v_0))
return ((uint32(__t_tag_0.IntVal) == 380165415)) != (true)
}


