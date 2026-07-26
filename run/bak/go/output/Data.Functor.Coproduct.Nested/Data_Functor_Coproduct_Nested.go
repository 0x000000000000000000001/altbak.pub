package Data_Functor_Coproduct_Nested

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Functor_Coproduct "gopurs/output/Data.Functor.Coproduct"
	pkg_Data_Either "gopurs/output/Data.Either"
	unsafe "unsafe"
)

var cache_in9 gopurs_runtime.Value
var once_in9 sync.Once
func Get_in9() gopurs_runtime.Value {
	once_in9.Do(func() {
		cache_in9 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in9(v_0_box)
})
	})
	return cache_in9
}

var cache_in8 gopurs_runtime.Value
var once_in8 sync.Once
func Get_in8() gopurs_runtime.Value {
	once_in8.Do(func() {
		cache_in8 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in8(v_0_box)
})
	})
	return cache_in8
}

var cache_in7 gopurs_runtime.Value
var once_in7 sync.Once
func Get_in7() gopurs_runtime.Value {
	once_in7.Do(func() {
		cache_in7 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in7(v_0_box)
})
	})
	return cache_in7
}

var cache_in6 gopurs_runtime.Value
var once_in6 sync.Once
func Get_in6() gopurs_runtime.Value {
	once_in6.Do(func() {
		cache_in6 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in6(v_0_box)
})
	})
	return cache_in6
}

var cache_in5 gopurs_runtime.Value
var once_in5 sync.Once
func Get_in5() gopurs_runtime.Value {
	once_in5.Do(func() {
		cache_in5 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in5(v_0_box)
})
	})
	return cache_in5
}

var cache_in4 gopurs_runtime.Value
var once_in4 sync.Once
func Get_in4() gopurs_runtime.Value {
	once_in4.Do(func() {
		cache_in4 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in4(v_0_box)
})
	})
	return cache_in4
}

var cache_in3 gopurs_runtime.Value
var once_in3 sync.Once
func Get_in3() gopurs_runtime.Value {
	once_in3.Do(func() {
		cache_in3 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in3(v_0_box)
})
	})
	return cache_in3
}

var cache_in2 gopurs_runtime.Value
var once_in2 sync.Once
func Get_in2() gopurs_runtime.Value {
	once_in2.Do(func() {
		cache_in2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in2(v_0_box)
})
	})
	return cache_in2
}

var cache_in10 gopurs_runtime.Value
var once_in10 sync.Once
func Get_in10() gopurs_runtime.Value {
	once_in10.Do(func() {
		cache_in10 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_in10(v_0_box)
})
	})
	return cache_in10
}

var cache_in1 gopurs_runtime.Value
var once_in1 sync.Once
func Get_in1() gopurs_runtime.Value {
	once_in1.Do(func() {
		cache_in1 = pkg_Data_Functor_Coproduct.Get_left__gopurs_runtime_Value()
	})
	return cache_in1
}

var cache_coproduct9 gopurs_runtime.Value
var once_coproduct9 sync.Once
func Get_coproduct9() gopurs_runtime.Value {
	once_coproduct9.Do(func() {
		cache_coproduct9 = gopurs_runtime.Func10(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, i_8_box gopurs_runtime.Value, y_9_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct9(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box, i_8_box, y_9_box)
})
	})
	return cache_coproduct9
}

var cache_coproduct8 gopurs_runtime.Value
var once_coproduct8 sync.Once
func Get_coproduct8() gopurs_runtime.Value {
	once_coproduct8.Do(func() {
		cache_coproduct8 = gopurs_runtime.Func9(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, y_8_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct8(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box, y_8_box)
})
	})
	return cache_coproduct8
}

var cache_coproduct7 gopurs_runtime.Value
var once_coproduct7 sync.Once
func Get_coproduct7() gopurs_runtime.Value {
	once_coproduct7.Do(func() {
		cache_coproduct7 = gopurs_runtime.Func8(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, y_7_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct7(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, y_7_box)
})
	})
	return cache_coproduct7
}

var cache_coproduct6 gopurs_runtime.Value
var once_coproduct6 sync.Once
func Get_coproduct6() gopurs_runtime.Value {
	once_coproduct6.Do(func() {
		cache_coproduct6 = gopurs_runtime.Func7(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, y_6_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct6(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, y_6_box)
})
	})
	return cache_coproduct6
}

var cache_coproduct5 gopurs_runtime.Value
var once_coproduct5 sync.Once
func Get_coproduct5() gopurs_runtime.Value {
	once_coproduct5.Do(func() {
		cache_coproduct5 = gopurs_runtime.Func6(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, y_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct5(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, y_5_box)
})
	})
	return cache_coproduct5
}

var cache_coproduct4 gopurs_runtime.Value
var once_coproduct4 sync.Once
func Get_coproduct4() gopurs_runtime.Value {
	once_coproduct4.Do(func() {
		cache_coproduct4 = gopurs_runtime.Func5(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, y_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct4(a_0_box, b_1_box, c_2_box, d_3_box, y_4_box)
})
	})
	return cache_coproduct4
}

var cache_coproduct3 gopurs_runtime.Value
var once_coproduct3 sync.Once
func Get_coproduct3() gopurs_runtime.Value {
	once_coproduct3.Do(func() {
		cache_coproduct3 = gopurs_runtime.Func4(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct3(a_0_box, b_1_box, c_2_box, y_3_box)
})
	})
	return cache_coproduct3
}

var cache_coproduct2 gopurs_runtime.Value
var once_coproduct2 sync.Once
func Get_coproduct2() gopurs_runtime.Value {
	once_coproduct2.Do(func() {
		cache_coproduct2 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct2(a_0_box, b_1_box, y_2_box)
})
	})
	return cache_coproduct2
}

var cache_coproduct10 gopurs_runtime.Value
var once_coproduct10 sync.Once
func Get_coproduct10() gopurs_runtime.Value {
	once_coproduct10.Do(func() {
		cache_coproduct10 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
var a_0_loop gopurs_runtime.Value = a_0_box
return gopurs_runtime.Func(func(b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
var b_1_loop gopurs_runtime.Value = b_1_box
return gopurs_runtime.Func(func(c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
var c_2_loop gopurs_runtime.Value = c_2_box
return gopurs_runtime.Func(func(d_3_box gopurs_runtime.Value) gopurs_runtime.Value {
var d_3_loop gopurs_runtime.Value = d_3_box
return gopurs_runtime.Func(func(e_4_box gopurs_runtime.Value) gopurs_runtime.Value {
var e_4_loop gopurs_runtime.Value = e_4_box
return gopurs_runtime.Func(func(f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
var f_5_loop gopurs_runtime.Value = f_5_box
return gopurs_runtime.Func(func(g_6_box gopurs_runtime.Value) gopurs_runtime.Value {
var g_6_loop gopurs_runtime.Value = g_6_box
return gopurs_runtime.Func(func(h_7_box gopurs_runtime.Value) gopurs_runtime.Value {
var h_7_loop gopurs_runtime.Value = h_7_box
return gopurs_runtime.Func(func(i_8_box gopurs_runtime.Value) gopurs_runtime.Value {
var i_8_loop gopurs_runtime.Value = i_8_box
return gopurs_runtime.Func(func(j_9_box gopurs_runtime.Value) gopurs_runtime.Value {
var j_9_loop gopurs_runtime.Value = j_9_box
return gopurs_runtime.Func(func(y_10_box gopurs_runtime.Value) gopurs_runtime.Value {
var y_10_loop gopurs_runtime.Value = y_10_box
return func() gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var d_3 gopurs_runtime.Value = d_3_loop
_ = d_3
var e_4 gopurs_runtime.Value = e_4_loop
_ = e_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var g_6 gopurs_runtime.Value = g_6_loop
_ = g_6
var h_7 gopurs_runtime.Value = h_7_loop
_ = h_7
var i_8 gopurs_runtime.Value = i_8_loop
_ = i_8
var j_9 gopurs_runtime.Value = j_9_loop
_ = j_9
var y_10 gopurs_runtime.Value = y_10_loop
_ = y_10
var __t0 gopurs_runtime.Value
{
if (y_10.Type == 9 && y_10.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_10.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_10.Type == 9 && y_10.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t10 gopurs_runtime.Value
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2465973597) {
var __t13 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 2465973597) {
var __t16 gopurs_runtime.Value
{
var __t_tag_17 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(g_6, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
var __t_tag_18 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_18.Type == 9 && __t_tag_18.IntVal == 2465973597) {
var __t19 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(h_7, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 2465973597) {
var __t22 gopurs_runtime.Value
{
var __t_tag_23 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_23.Type == 9 && __t_tag_23.IntVal == 3711209382) {
__t22 = gopurs_runtime.Apply(i_8, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
var __t_tag_24 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_24.Type == 9 && __t_tag_24.IntVal == 2465973597) {
var __t25 gopurs_runtime.Value
{
var __t_tag_26 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_26.Type == 9 && __t_tag_26.IntVal == 3711209382) {
__t25 = gopurs_runtime.Apply(j_9, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_25
} else {

}
}
{
var __t_tag_27 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_27.Type == 9 && __t_tag_27.IntVal == 2465973597) {
var spin_11_28 gopurs_runtime.Value
spin_11_28 = gopurs_runtime.Func(func(v_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_12_loop gopurs_runtime.Value = v_12_loop_val
spin_11_28:
for {
if false { continue spin_11_28 }
var v_12 gopurs_runtime.Value = v_12_loop
_ = v_12
v_12_loop = v_12
continue spin_11_28
return gopurs_runtime.Value{}
}
}()
})
__t25 = gopurs_runtime.Apply(spin_11_28, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
__t22 = __t25
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
__t19 = __t22
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t16 = __t19
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t13 = __t16
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t10 = __t13
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t7 = __t10
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
})
})
})
})
})
})
})
})
})
})
	})
	return cache_coproduct10
}

var cache_coproduct1 gopurs_runtime.Value
var once_coproduct1 sync.Once
func Get_coproduct1() gopurs_runtime.Value {
	once_coproduct1.Do(func() {
		cache_coproduct1 = gopurs_runtime.Func(func(y_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct1(y_0_box)
})
	})
	return cache_coproduct1
}

var cache_at9 gopurs_runtime.Value
var once_at9 sync.Once
func Get_at9() gopurs_runtime.Value {
	once_at9.Do(func() {
		cache_at9 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at9(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at9
}

var cache_at8 gopurs_runtime.Value
var once_at8 sync.Once
func Get_at8() gopurs_runtime.Value {
	once_at8.Do(func() {
		cache_at8 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at8(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at8
}

var cache_at7 gopurs_runtime.Value
var once_at7 sync.Once
func Get_at7() gopurs_runtime.Value {
	once_at7.Do(func() {
		cache_at7 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at7(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at7
}

var cache_at6 gopurs_runtime.Value
var once_at6 sync.Once
func Get_at6() gopurs_runtime.Value {
	once_at6.Do(func() {
		cache_at6 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at6(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at6
}

var cache_at5 gopurs_runtime.Value
var once_at5 sync.Once
func Get_at5() gopurs_runtime.Value {
	once_at5.Do(func() {
		cache_at5 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at5(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at5
}

var cache_at4 gopurs_runtime.Value
var once_at4 sync.Once
func Get_at4() gopurs_runtime.Value {
	once_at4.Do(func() {
		cache_at4 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at4(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at4
}

var cache_at3 gopurs_runtime.Value
var once_at3 sync.Once
func Get_at3() gopurs_runtime.Value {
	once_at3.Do(func() {
		cache_at3 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at3(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at3
}

var cache_at2 gopurs_runtime.Value
var once_at2 sync.Once
func Get_at2() gopurs_runtime.Value {
	once_at2.Do(func() {
		cache_at2 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at2(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at2
}

var cache_at10 gopurs_runtime.Value
var once_at10 sync.Once
func Get_at10() gopurs_runtime.Value {
	once_at10.Do(func() {
		cache_at10 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at10(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at10
}

var cache_at1 gopurs_runtime.Value
var once_at1 sync.Once
func Get_at1() gopurs_runtime.Value {
	once_at1.Do(func() {
		cache_at1 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at1(b_0_box, f_1_box, y_2_box)
})
	})
	return cache_at1
}

type Record_alt_gopurs_runtime_Value struct {
	alt gopurs_runtime.Value
}

type Record_ struct {
	
}

type Record_pure_gopurs_runtime_Value struct {
	pure gopurs_runtime.Value
}

type Record_apply_gopurs_runtime_Value struct {
	apply gopurs_runtime.Value
}

type Record_bipure_gopurs_runtime_Value struct {
	bipure gopurs_runtime.Value
}

type Record_biapply_gopurs_runtime_Value struct {
	biapply gopurs_runtime.Value
}

type Record_bind_gopurs_runtime_Value struct {
	bind gopurs_runtime.Value
}

type Record_discard_gopurs_runtime_Value struct {
	discard gopurs_runtime.Value
}

type Record_identity_gopurs_runtime_Value struct {
	identity gopurs_runtime.Value
}

type Record_ask_gopurs_runtime_Value struct {
	ask gopurs_runtime.Value
}

type Record_local_gopurs_runtime_Value struct {
	local gopurs_runtime.Value
}

type Record_peek_gopurs_runtime_Value_pos_gopurs_runtime_Value struct {
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
}

type Record_track_gopurs_runtime_Value struct {
	track gopurs_runtime.Value
}

type Record_extract_gopurs_runtime_Value struct {
	extract gopurs_runtime.Value
}

type Record_extend_gopurs_runtime_Value struct {
	extend gopurs_runtime.Value
}

type Record_defer__gopurs_runtime_Value struct {
	defer_ gopurs_runtime.Value
}

type Record_callCC_gopurs_runtime_Value struct {
	callCC gopurs_runtime.Value
}

type Record_catchError_gopurs_runtime_Value struct {
	catchError gopurs_runtime.Value
}

type Record_throwError_gopurs_runtime_Value struct {
	throwError gopurs_runtime.Value
}

type Record_chooseBool_gopurs_runtime_Value_chooseFloat_gopurs_runtime_Value_chooseInt_gopurs_runtime_Value_resize_gopurs_runtime_Value_sized_gopurs_runtime_Value struct {
	chooseBool gopurs_runtime.Value
	chooseFloat gopurs_runtime.Value
	chooseInt gopurs_runtime.Value
	resize gopurs_runtime.Value
	sized gopurs_runtime.Value
}

type Record_foldMap1_gopurs_runtime_Value_foldl1_gopurs_runtime_Value_foldr1_gopurs_runtime_Value struct {
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
}

type Record_append__gopurs_runtime_Value struct {
	append_ gopurs_runtime.Value
}

type Record_tailRecM_gopurs_runtime_Value struct {
	tailRecM gopurs_runtime.Value
}

type Record_unfoldr_gopurs_runtime_Value struct {
	unfoldr gopurs_runtime.Value
}

type Record_map__gopurs_runtime_Value struct {
	map_ gopurs_runtime.Value
}

type Record_state_gopurs_runtime_Value struct {
	state gopurs_runtime.Value
}

type Record_lift_gopurs_runtime_Value struct {
	lift gopurs_runtime.Value
}

type Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value struct {
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
}

type Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value struct {
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
}

type Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value struct {
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
}

type Record_mempty_gopurs_runtime_Value struct {
	mempty gopurs_runtime.Value
}

type Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value struct {
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
}

type Record_empty_gopurs_runtime_Value struct {
	empty gopurs_runtime.Value
}

type Record_compose_gopurs_runtime_Value struct {
	compose gopurs_runtime.Value
}

type Record_eq_gopurs_runtime_Value struct {
	eq gopurs_runtime.Value
}

type Record_compare_gopurs_runtime_Value struct {
	compare gopurs_runtime.Value
}

type Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value struct {
	bifoldMap gopurs_runtime.Value
	bifoldl gopurs_runtime.Value
	bifoldr gopurs_runtime.Value
}

type Record_bimap_gopurs_runtime_Value struct {
	bimap gopurs_runtime.Value
}

type Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value struct {
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
}

type Record_genericBottom_prime_gopurs_runtime_Value struct {
	genericBottom_prime gopurs_runtime.Value
}

type Record_genericTop_prime_gopurs_runtime_Value struct {
	genericTop_prime gopurs_runtime.Value
}

type Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value struct {
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
}

type Record_lose_gopurs_runtime_Value struct {
	lose gopurs_runtime.Value
}

type Record_choose_gopurs_runtime_Value struct {
	choose gopurs_runtime.Value
}

type Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value struct {
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
}

type Record_divide_gopurs_runtime_Value struct {
	divide gopurs_runtime.Value
}

type Record_recip_gopurs_runtime_Value struct {
	recip gopurs_runtime.Value
}

type Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value struct {
	genericCardinality_prime gopurs_runtime.Value
	genericFromEnum_prime gopurs_runtime.Value
	genericToEnum_prime gopurs_runtime.Value
}

type Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value struct {
	genericPred_prime gopurs_runtime.Value
	genericSucc_prime gopurs_runtime.Value
}

type Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value struct {
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}

type Record_unfoldr1_gopurs_runtime_Value struct {
	unfoldr1 gopurs_runtime.Value
}

type Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value struct {
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}

type Record_genericEq_prime_gopurs_runtime_Value struct {
	genericEq_prime gopurs_runtime.Value
}

type Record_eq1_gopurs_runtime_Value struct {
	eq1 gopurs_runtime.Value
}

type Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value struct {
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff gopurs_runtime.Value
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt gopurs_runtime.Value
}

type Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value struct {
	add gopurs_runtime.Value
	mul gopurs_runtime.Value
	one gopurs_runtime.Value
	zero gopurs_runtime.Value
}

type Record_foldMapWithIndex_gopurs_runtime_Value_foldlWithIndex_gopurs_runtime_Value_foldrWithIndex_gopurs_runtime_Value struct {
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
}

type Record_cmap_gopurs_runtime_Value struct {
	cmap gopurs_runtime.Value
}

type Record_imap_gopurs_runtime_Value struct {
	imap gopurs_runtime.Value
}

type Record_mapWithIndex_gopurs_runtime_Value struct {
	mapWithIndex gopurs_runtime.Value
}

type Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value struct {
	from gopurs_runtime.Value
	to gopurs_runtime.Value
}

type Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value struct {
	genericConj_prime gopurs_runtime.Value
	genericDisj_prime gopurs_runtime.Value
	genericFF_prime gopurs_runtime.Value
	genericImplies_prime gopurs_runtime.Value
	genericNot_prime gopurs_runtime.Value
	genericTT_prime gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_bool_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_bool struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff bool
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt bool
}

type Record_genericMempty_prime_gopurs_runtime_Value struct {
	genericMempty_prime gopurs_runtime.Value
}

type Record_genericCompare_prime_gopurs_runtime_Value struct {
	genericCompare_prime gopurs_runtime.Value
}

type Record_sub_gopurs_runtime_Value struct {
	sub gopurs_runtime.Value
}

type Record_compare1_gopurs_runtime_Value struct {
	compare1 gopurs_runtime.Value
}

type Record_left_gopurs_runtime_Value_right_gopurs_runtime_Value struct {
	left gopurs_runtime.Value
	right gopurs_runtime.Value
}

type Record_first_gopurs_runtime_Value_second_gopurs_runtime_Value struct {
	first gopurs_runtime.Value
	second gopurs_runtime.Value
}

type Record_dimap_gopurs_runtime_Value struct {
	dimap gopurs_runtime.Value
}

type Record_genericSub_prime_gopurs_runtime_Value struct {
	genericSub_prime gopurs_runtime.Value
}

type Record_genericAppend_prime_gopurs_runtime_Value struct {
	genericAppend_prime gopurs_runtime.Value
}

type Record_sequence1_gopurs_runtime_Value_traverse1_gopurs_runtime_Value struct {
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
}

type Record_genericAdd_prime_gopurs_runtime_Value_genericMul_prime_gopurs_runtime_Value_genericOne_prime_gopurs_runtime_Value_genericZero_prime_gopurs_runtime_Value struct {
	genericAdd_prime gopurs_runtime.Value
	genericMul_prime gopurs_runtime.Value
	genericOne_prime gopurs_runtime.Value
	genericZero_prime gopurs_runtime.Value
}

type Record_genericShow_prime_gopurs_runtime_Value struct {
	genericShow_prime gopurs_runtime.Value
}

type Record_genericShowArgs_gopurs_runtime_Value struct {
	genericShowArgs gopurs_runtime.Value
}

type Record_show_gopurs_runtime_Value struct {
	show gopurs_runtime.Value
}

type Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value struct {
	fromDuration gopurs_runtime.Value
	toDuration gopurs_runtime.Value
}

type Record_traverseWithIndex_gopurs_runtime_Value struct {
	traverseWithIndex gopurs_runtime.Value
}

type Record_liftEffect_gopurs_runtime_Value struct {
	liftEffect gopurs_runtime.Value
}

type Record_mappend__gopurs_runtime_Value_mempty__gopurs_runtime_Value struct {
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
}

type Record_proof_gopurs_runtime_Value struct {
	proof gopurs_runtime.Value
}

type Record_lower_gopurs_runtime_Value struct {
	lower gopurs_runtime.Value
}

type Record_liftST_gopurs_runtime_Value struct {
	liftST gopurs_runtime.Value
}

type Record_tell_gopurs_runtime_Value struct {
	tell gopurs_runtime.Value
}

type Record_reflectSymbol_gopurs_runtime_Value struct {
	reflectSymbol gopurs_runtime.Value
}

type Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value struct {
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
}

type Record_conquer_gopurs_runtime_Value struct {
	conquer gopurs_runtime.Value
}

type Record_inj_gopurs_runtime_Value_prj_gopurs_runtime_Value struct {
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
}

type Record_eqRecord_gopurs_runtime_Value struct {
	eqRecord gopurs_runtime.Value
}

type Record_conjRecord_gopurs_runtime_Value_disjRecord_gopurs_runtime_Value_ffRecord_gopurs_runtime_Value_impliesRecord_gopurs_runtime_Value_notRecord_gopurs_runtime_Value_ttRecord_gopurs_runtime_Value struct {
	conjRecord gopurs_runtime.Value
	disjRecord gopurs_runtime.Value
	ffRecord gopurs_runtime.Value
	impliesRecord gopurs_runtime.Value
	notRecord gopurs_runtime.Value
	ttRecord gopurs_runtime.Value
}

type Record_memptyRecord_gopurs_runtime_Value struct {
	memptyRecord gopurs_runtime.Value
}

type Record_compareRecord_gopurs_runtime_Value struct {
	compareRecord gopurs_runtime.Value
}

type Record_closed_gopurs_runtime_Value struct {
	closed gopurs_runtime.Value
}

type Record_unleft_gopurs_runtime_Value_unright_gopurs_runtime_Value struct {
	unleft gopurs_runtime.Value
	unright gopurs_runtime.Value
}

type Record_unfirst_gopurs_runtime_Value_unsecond_gopurs_runtime_Value struct {
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
}

type Record_reflectType_gopurs_runtime_Value struct {
	reflectType gopurs_runtime.Value
}

type Record_subRecord_gopurs_runtime_Value struct {
	subRecord gopurs_runtime.Value
}

type Record_appendRecord_gopurs_runtime_Value struct {
	appendRecord gopurs_runtime.Value
}

type Record_addRecord_gopurs_runtime_Value_mulRecord_gopurs_runtime_Value_oneRecord_gopurs_runtime_Value_zeroRecord_gopurs_runtime_Value struct {
	addRecord gopurs_runtime.Value
	mulRecord gopurs_runtime.Value
	oneRecord gopurs_runtime.Value
	zeroRecord gopurs_runtime.Value
}

type Record_showRecordFields_gopurs_runtime_Value struct {
	showRecordFields gopurs_runtime.Value
}

type Record_nes_gopurs_runtime_Value struct {
	nes gopurs_runtime.Value
}

type Record_liftAff_gopurs_runtime_Value struct {
	liftAff gopurs_runtime.Value
}

func Call_in9(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}})}})}})}})}})}})}})}
}

func Call_in8(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}})}})}})}})}})}})}
}

func Call_in7(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}})}})}})}})}})}
}

func Call_in6(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}})}})}})}})}
}

func Call_in5(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}})}})}})}
}

func Call_in4(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}})}})}
}

func Call_in3(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}})}
}

func Call_in2(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}
}

func Call_in10(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{v_0})}})}})}})}})}})}})}})}})}})}
}

func Call_coproduct9(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, h_7_loop gopurs_runtime.Value, i_8_loop gopurs_runtime.Value, y_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var d_3 gopurs_runtime.Value = d_3_loop
_ = d_3
var e_4 gopurs_runtime.Value = e_4_loop
_ = e_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var g_6 gopurs_runtime.Value = g_6_loop
_ = g_6
var h_7 gopurs_runtime.Value = h_7_loop
_ = h_7
var i_8 gopurs_runtime.Value = i_8_loop
_ = i_8
var y_9 gopurs_runtime.Value = y_9_loop
_ = y_9
var __t0 gopurs_runtime.Value
{
if (y_9.Type == 9 && y_9.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_9.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_9.Type == 9 && y_9.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t10 gopurs_runtime.Value
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2465973597) {
var __t13 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 2465973597) {
var __t16 gopurs_runtime.Value
{
var __t_tag_17 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(g_6, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
var __t_tag_18 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_18.Type == 9 && __t_tag_18.IntVal == 2465973597) {
var __t19 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(h_7, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 2465973597) {
var __t22 gopurs_runtime.Value
{
var __t_tag_23 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_23.Type == 9 && __t_tag_23.IntVal == 3711209382) {
__t22 = gopurs_runtime.Apply(i_8, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
var __t_tag_24 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_24.Type == 9 && __t_tag_24.IntVal == 2465973597) {
var spin_10_25 gopurs_runtime.Value
spin_10_25 = gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_11_loop gopurs_runtime.Value = v_11_loop_val
spin_10_25:
for {
if false { continue spin_10_25 }
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
v_11_loop = v_11
continue spin_10_25
return gopurs_runtime.Value{}
}
}()
})
__t22 = gopurs_runtime.Apply(spin_10_25, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
__t19 = __t22
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t16 = __t19
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t13 = __t16
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t10 = __t13
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t7 = __t10
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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

func Call_coproduct8(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, h_7_loop gopurs_runtime.Value, y_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var d_3 gopurs_runtime.Value = d_3_loop
_ = d_3
var e_4 gopurs_runtime.Value = e_4_loop
_ = e_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var g_6 gopurs_runtime.Value = g_6_loop
_ = g_6
var h_7 gopurs_runtime.Value = h_7_loop
_ = h_7
var y_8 gopurs_runtime.Value = y_8_loop
_ = y_8
var __t0 gopurs_runtime.Value
{
if (y_8.Type == 9 && y_8.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_8.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_8.Type == 9 && y_8.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t10 gopurs_runtime.Value
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2465973597) {
var __t13 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 2465973597) {
var __t16 gopurs_runtime.Value
{
var __t_tag_17 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(g_6, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
var __t_tag_18 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_18.Type == 9 && __t_tag_18.IntVal == 2465973597) {
var __t19 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(h_7, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 2465973597) {
var spin_9_22 gopurs_runtime.Value
spin_9_22 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_10_loop gopurs_runtime.Value = v_10_loop_val
spin_9_22:
for {
if false { continue spin_9_22 }
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
v_10_loop = v_10
continue spin_9_22
return gopurs_runtime.Value{}
}
}()
})
__t19 = gopurs_runtime.Apply(spin_9_22, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t16 = __t19
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t13 = __t16
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t10 = __t13
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t7 = __t10
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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

func Call_coproduct7(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, y_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var d_3 gopurs_runtime.Value = d_3_loop
_ = d_3
var e_4 gopurs_runtime.Value = e_4_loop
_ = e_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var g_6 gopurs_runtime.Value = g_6_loop
_ = g_6
var y_7 gopurs_runtime.Value = y_7_loop
_ = y_7
var __t0 gopurs_runtime.Value
{
if (y_7.Type == 9 && y_7.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_7.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_7.Type == 9 && y_7.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t10 gopurs_runtime.Value
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2465973597) {
var __t13 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 2465973597) {
var __t16 gopurs_runtime.Value
{
var __t_tag_17 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(g_6, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
var __t_tag_18 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_18.Type == 9 && __t_tag_18.IntVal == 2465973597) {
var spin_8_19 gopurs_runtime.Value
spin_8_19 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_9_loop gopurs_runtime.Value = v_9_loop_val
spin_8_19:
for {
if false { continue spin_8_19 }
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
v_9_loop = v_9
continue spin_8_19
return gopurs_runtime.Value{}
}
}()
})
__t16 = gopurs_runtime.Apply(spin_8_19, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t13 = __t16
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t10 = __t13
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t7 = __t10
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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

func Call_coproduct6(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, y_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var d_3 gopurs_runtime.Value = d_3_loop
_ = d_3
var e_4 gopurs_runtime.Value = e_4_loop
_ = e_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var y_6 gopurs_runtime.Value = y_6_loop
_ = y_6
var __t0 gopurs_runtime.Value
{
if (y_6.Type == 9 && y_6.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_6.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t10 gopurs_runtime.Value
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2465973597) {
var __t13 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 2465973597) {
var spin_7_16 gopurs_runtime.Value
spin_7_16 = gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_8_loop gopurs_runtime.Value = v_8_loop_val
spin_7_16:
for {
if false { continue spin_7_16 }
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
v_8_loop = v_8
continue spin_7_16
return gopurs_runtime.Value{}
}
}()
})
__t13 = gopurs_runtime.Apply(spin_7_16, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t10 = __t13
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t7 = __t10
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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

func Call_coproduct5(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, y_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var d_3 gopurs_runtime.Value = d_3_loop
_ = d_3
var e_4 gopurs_runtime.Value = e_4_loop
_ = e_4
var y_5 gopurs_runtime.Value = y_5_loop
_ = y_5
var __t0 gopurs_runtime.Value
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_5.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t10 gopurs_runtime.Value
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2465973597) {
var spin_6_13 gopurs_runtime.Value
spin_6_13 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop gopurs_runtime.Value = v_7_loop_val
spin_6_13:
for {
if false { continue spin_6_13 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
v_7_loop = v_7
continue spin_6_13
return gopurs_runtime.Value{}
}
}()
})
__t10 = gopurs_runtime.Apply(spin_6_13, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t7 = __t10
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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

func Call_coproduct4(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, y_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var d_3 gopurs_runtime.Value = d_3_loop
_ = d_3
var y_4 gopurs_runtime.Value = y_4_loop
_ = y_4
var __t0 gopurs_runtime.Value
{
if (y_4.Type == 9 && y_4.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_4.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_4.Type == 9 && y_4.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var spin_5_10 gopurs_runtime.Value
spin_5_10 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
spin_5_10:
for {
if false { continue spin_5_10 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
v_6_loop = v_6
continue spin_5_10
return gopurs_runtime.Value{}
}
}()
})
__t7 = gopurs_runtime.Apply(spin_5_10, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t4 = __t7
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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

func Call_coproduct3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
var __t0 gopurs_runtime.Value
{
if (y_3.Type == 9 && y_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_3.Type == 9 && y_3.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_3.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_3.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_3.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_3.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_3.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_3.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {
var spin_4_7 gopurs_runtime.Value
spin_4_7 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
spin_4_7:
for {
if false { continue spin_4_7 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
v_5_loop = v_5
continue spin_4_7
return gopurs_runtime.Value{}
}
}()
})
__t4 = gopurs_runtime.Apply(spin_4_7, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_3.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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

func Call_coproduct2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Data_Data_Either_Left)(y_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var spin_3_4 gopurs_runtime.Value
spin_3_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
spin_3_4:
for {
if false { continue spin_3_4 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
v_4_loop = v_4
continue spin_3_4
return gopurs_runtime.Value{}
}
}()
})
__t1 = gopurs_runtime.Apply(spin_3_4, (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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

func Call_coproduct1(y_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var __t0 gopurs_runtime.Value
{
if (y_0.Type == 9 && y_0.IntVal == 3711209382) {
__t0 = (*pkg_Data_Either.Data_Data_Either_Left)(y_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
if (y_0.Type == 9 && y_0.IntVal == 2465973597) {
var spin_1_1 gopurs_runtime.Value
spin_1_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_1:
for {
if false { continue spin_1_1 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_1
return gopurs_runtime.Value{}
}
}()
})
__t0 = gopurs_runtime.Apply(spin_1_1, (*pkg_Data_Either.Data_Data_Either_Right)(y_0.UnsafePtr).V0)
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

func Call_at9(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_16 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
var __t_and_15 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_14 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_13 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_12 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_11 bool = false
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {

var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {

var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_9 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {

var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_9 = (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382)
}
__t_and_10 = __t_and_9
}
__t_and_11 = __t_and_10
}
__t_and_12 = __t_and_11
}
__t_and_13 = __t_and_12
}
__t_and_14 = __t_and_13
}
__t_and_15 = __t_and_14
}
__t_and_16 = __t_and_15
}
if __t_and_16 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at8(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_14 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
var __t_and_13 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_12 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_11 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_9 bool = false
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {

var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_8 bool = false
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {

var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_8 = (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 3711209382)
}
__t_and_9 = __t_and_8
}
__t_and_10 = __t_and_9
}
__t_and_11 = __t_and_10
}
__t_and_12 = __t_and_11
}
__t_and_13 = __t_and_12
}
__t_and_14 = __t_and_13
}
if __t_and_14 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at7(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_12 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
var __t_and_11 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_9 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_8 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_7 bool = false
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {

var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_7 = (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3711209382)
}
__t_and_8 = __t_and_7
}
__t_and_9 = __t_and_8
}
__t_and_10 = __t_and_9
}
__t_and_11 = __t_and_10
}
__t_and_12 = __t_and_11
}
if __t_and_12 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at6(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_10 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
var __t_and_9 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_8 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_7 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_6 = (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3711209382)
}
__t_and_7 = __t_and_6
}
__t_and_8 = __t_and_7
}
__t_and_9 = __t_and_8
}
__t_and_10 = __t_and_9
}
if __t_and_10 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at5(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_8 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
var __t_and_7 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
__t_and_8 = __t_and_7
}
if __t_and_8 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at4(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_6 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
var __t_and_5 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_4 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_4 = (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 3711209382)
}
__t_and_5 = __t_and_4
}
__t_and_6 = __t_and_5
}
if __t_and_6 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at3(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_4 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
var __t_and_3 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0
__t_and_3 = (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382)
}
__t_and_4 = __t_and_3
}
if __t_and_4 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at2(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_2 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
__t_and_2 = (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 3711209382)
}
if __t_and_2 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at10(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
var __t_and_18 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0
var __t_and_17 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_16 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_15 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_14 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_13 bool = false
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {

var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_12 bool = false
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {

var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_11 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {

var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 2465973597) {

var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_10 = (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 3711209382)
}
__t_and_11 = __t_and_10
}
__t_and_12 = __t_and_11
}
__t_and_13 = __t_and_12
}
__t_and_14 = __t_and_13
}
__t_and_15 = __t_and_14
}
__t_and_16 = __t_and_15
}
__t_and_17 = __t_and_16
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)((*pkg_Data_Either.Data_Data_Either_Right)(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}

func Call_at1(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Data_Data_Either_Left)(y_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = b_0
}
end_branch_0:
return __t0
}


