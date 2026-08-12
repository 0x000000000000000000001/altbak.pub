package Data_Functor_Coproduct_Nested

import (
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor_Coproduct "gopurs/output/Data.Functor.Coproduct"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
		cache_in1 = pkg_Data_Functor_Coproduct.Get_left()
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
var __t28 gopurs_runtime.Value
{
if (y_10.Type == 9 && y_10.IntVal == 3711209382) {
__t28 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0)
goto end_branch_28
} else {

}
}
{
if (y_10.Type == 9 && y_10.IntVal == 2465973597) {
var __t27 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t27 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_27
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var __t26 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t26 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_26
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t25 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382) {
__t25 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_25
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {
var __t24 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3711209382) {
__t24 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_24
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {
var __t23 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t23 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t22 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 3711209382) {
__t22 = gopurs_runtime.Apply(g_6, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 2465973597) {
var __t21 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 3711209382) {
__t21 = gopurs_runtime.Apply(h_7, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_21
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 2465973597) {
var __t20 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(i_8, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_20
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 2465973597) {
var __t19 gopurs_runtime.Value
{
var __t_tag_16 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(j_9, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
var __t_tag_17 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 2465973597) {
var spin_11_18_8 gopurs_runtime.Value
spin_11_18_8 = gopurs_runtime.Func(func(v_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_12_loop gopurs_runtime.Value = v_12_loop_val
spin_11_18_8:
for {
if false { continue spin_11_18_8 }
var v_12 gopurs_runtime.Value = v_12_loop
_ = v_12
v_12_loop = v_12
continue spin_11_18_8
return gopurs_runtime.Value{}
}
}()
})
__t19 = gopurs_runtime.Apply(spin_11_18_8, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t20 = __t19
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
__t21 = __t20
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t22 = __t21
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
__t23 = __t22
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t24 = __t23
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = __t24
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
__t27 = __t26
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
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

var cache_left__786465523 gopurs_runtime.Value
var once_left__786465523 sync.Once
func Get_left__786465523() gopurs_runtime.Value {
	once_left__786465523.Do(func() {
		cache_left__786465523 = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_left__786465523(fa_0_box)
})
	})
	return cache_left__786465523
}

var cache_left__1785278009 gopurs_runtime.Value
var once_left__1785278009 sync.Once
func Get_left__1785278009() gopurs_runtime.Value {
	once_left__1785278009.Do(func() {
		cache_left__1785278009 = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_left__1785278009(fa_0_box)
})
	})
	return cache_left__1785278009
}

var cache_left__4018477375 gopurs_runtime.Value
var once_left__4018477375 sync.Once
func Get_left__4018477375() gopurs_runtime.Value {
	once_left__4018477375.Do(func() {
		cache_left__4018477375 = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_left__4018477375(fa_0_box)
})
	})
	return cache_left__4018477375
}

var cache_right__1193983057 gopurs_runtime.Value
var once_right__1193983057 sync.Once
func Get_right__1193983057() gopurs_runtime.Value {
	once_right__1193983057.Do(func() {
		cache_right__1193983057 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__1193983057(ga_0_box)
})
	})
	return cache_right__1193983057
}

var cache_right__1623275833 gopurs_runtime.Value
var once_right__1623275833 sync.Once
func Get_right__1623275833() gopurs_runtime.Value {
	once_right__1623275833.Do(func() {
		cache_right__1623275833 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__1623275833(ga_0_box)
})
	})
	return cache_right__1623275833
}

var cache_right__856638295 gopurs_runtime.Value
var once_right__856638295 sync.Once
func Get_right__856638295() gopurs_runtime.Value {
	once_right__856638295.Do(func() {
		cache_right__856638295 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__856638295(ga_0_box)
})
	})
	return cache_right__856638295
}

var cache_right__1175236921 gopurs_runtime.Value
var once_right__1175236921 sync.Once
func Get_right__1175236921() gopurs_runtime.Value {
	once_right__1175236921.Do(func() {
		cache_right__1175236921 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__1175236921(ga_0_box)
})
	})
	return cache_right__1175236921
}

var cache_right__3235459543 gopurs_runtime.Value
var once_right__3235459543 sync.Once
func Get_right__3235459543() gopurs_runtime.Value {
	once_right__3235459543.Do(func() {
		cache_right__3235459543 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__3235459543(ga_0_box)
})
	})
	return cache_right__3235459543
}

var cache_right__2015984185 gopurs_runtime.Value
var once_right__2015984185 sync.Once
func Get_right__2015984185() gopurs_runtime.Value {
	once_right__2015984185.Do(func() {
		cache_right__2015984185 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__2015984185(ga_0_box)
})
	})
	return cache_right__2015984185
}

var cache_right__176310615 gopurs_runtime.Value
var once_right__176310615 sync.Once
func Get_right__176310615() gopurs_runtime.Value {
	once_right__176310615.Do(func() {
		cache_right__176310615 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__176310615(ga_0_box)
})
	})
	return cache_right__176310615
}

var cache_right__1052251193 gopurs_runtime.Value
var once_right__1052251193 sync.Once
func Get_right__1052251193() gopurs_runtime.Value {
	once_right__1052251193.Do(func() {
		cache_right__1052251193 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__1052251193(ga_0_box)
})
	})
	return cache_right__1052251193
}

var cache_right__2641088471 gopurs_runtime.Value
var once_right__2641088471 sync.Once
func Get_right__2641088471() gopurs_runtime.Value {
	once_right__2641088471.Do(func() {
		cache_right__2641088471 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__2641088471(ga_0_box)
})
	})
	return cache_right__2641088471
}

var cache_right__3220060985 gopurs_runtime.Value
var once_right__3220060985 sync.Once
func Get_right__3220060985() gopurs_runtime.Value {
	once_right__3220060985.Do(func() {
		cache_right__3220060985 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_right__3220060985(ga_0_box)
})
	})
	return cache_right__3220060985
}

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

var cache_unwrap__3741956339 gopurs_runtime.Value
var once_unwrap__3741956339 sync.Once
func Get_unwrap__3741956339() gopurs_runtime.Value {
	once_unwrap__3741956339.Do(func() {
		cache_unwrap__3741956339 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3741956339(__eta0_0_box)
})
	})
	return cache_unwrap__3741956339
}

var cache_absurd__3279552488 gopurs_runtime.Value
var once_absurd__3279552488 sync.Once
func Get_absurd__3279552488() gopurs_runtime.Value {
	once_absurd__3279552488.Do(func() {
		cache_absurd__3279552488 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__3279552488(a_0_box)
})
	})
	return cache_absurd__3279552488
}

var cache_absurd__943547431 gopurs_runtime.Value
var once_absurd__943547431 sync.Once
func Get_absurd__943547431() gopurs_runtime.Value {
	once_absurd__943547431.Do(func() {
		cache_absurd__943547431 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__943547431(a_0_box)
})
	})
	return cache_absurd__943547431
}

func Call_in9(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}})}})}})}})}})}})}})}
}

func Call_in8(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}})}})}})}})}})}})}
}

func Call_in7(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}})}})}})}})}})}
}

func Call_in6(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}})}})}})}})}
}

func Call_in5(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}})}})}})}
}

func Call_in4(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}})}})}
}

func Call_in3(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}})}
}

func Call_in2(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}
}

func Call_in10(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_0})}})}})}})}})}})}})}})}})}})}
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
var __t25 gopurs_runtime.Value
{
if (y_9.Type == 9 && y_9.IntVal == 3711209382) {
__t25 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0)
goto end_branch_25
} else {

}
}
{
if (y_9.Type == 9 && y_9.IntVal == 2465973597) {
var __t24 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t24 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_24
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var __t23 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t23 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t22 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382) {
__t22 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {
var __t21 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3711209382) {
__t21 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_21
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {
var __t20 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_20
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t19 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(g_6, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 2465973597) {
var __t18 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(h_7, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 2465973597) {
var __t17 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(i_8, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 2465973597) {
var spin_10_16_0 gopurs_runtime.Value
spin_10_16_0 = gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_11_loop gopurs_runtime.Value = v_11_loop_val
spin_10_16_0:
for {
if false { continue spin_10_16_0 }
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
v_11_loop = v_11
continue spin_10_16_0
return gopurs_runtime.Value{}
}
}()
})
__t17 = gopurs_runtime.Apply(spin_10_16_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
__t18 = __t17
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t20 = __t19
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
__t21 = __t20
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t22 = __t21
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
__t23 = __t22
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t24 = __t23
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = __t24
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
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
var __t22 gopurs_runtime.Value
{
if (y_8.Type == 9 && y_8.IntVal == 3711209382) {
__t22 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
if (y_8.Type == 9 && y_8.IntVal == 2465973597) {
var __t21 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t21 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_21
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var __t20 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_20
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t19 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {
var __t18 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {
var __t17 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t16 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(g_6, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 2465973597) {
var __t15 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(h_7, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 2465973597) {
var spin_9_14_1 gopurs_runtime.Value
spin_9_14_1 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_10_loop gopurs_runtime.Value = v_10_loop_val
spin_9_14_1:
for {
if false { continue spin_9_14_1 }
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
v_10_loop = v_10
continue spin_9_14_1
return gopurs_runtime.Value{}
}
}()
})
__t15 = gopurs_runtime.Apply(spin_9_14_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t17 = __t16
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
__t18 = __t17
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t20 = __t19
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
__t21 = __t20
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t22 = __t21
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
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
var __t19 gopurs_runtime.Value
{
if (y_7.Type == 9 && y_7.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
if (y_7.Type == 9 && y_7.IntVal == 2465973597) {
var __t18 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var __t17 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t16 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {
var __t15 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {
var __t14 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t14 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var __t13 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(g_6, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 2465973597) {
var spin_8_12_2 gopurs_runtime.Value
spin_8_12_2 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_9_loop gopurs_runtime.Value = v_9_loop_val
spin_8_12_2:
for {
if false { continue spin_8_12_2 }
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
v_9_loop = v_9
continue spin_8_12_2
return gopurs_runtime.Value{}
}
}()
})
__t13 = gopurs_runtime.Apply(spin_8_12_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t17 = __t16
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
__t18 = __t17
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
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
var __t16 gopurs_runtime.Value
{
if (y_6.Type == 9 && y_6.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 2465973597) {
var __t15 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var __t14 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t14 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t13 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {
var __t12 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3711209382) {
__t12 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {
var __t11 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3711209382) {
__t11 = gopurs_runtime.Apply(f_5, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 2465973597) {
var spin_7_10_3 gopurs_runtime.Value
spin_7_10_3 = gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_8_loop gopurs_runtime.Value = v_8_loop_val
spin_7_10_3:
for {
if false { continue spin_7_10_3 }
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
v_8_loop = v_8
continue spin_7_10_3
return gopurs_runtime.Value{}
}
}()
})
__t11 = gopurs_runtime.Apply(spin_7_10_3, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t12 = __t11
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
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
var __t13 gopurs_runtime.Value
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 2465973597) {
var __t12 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t12 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var __t11 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t11 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t10 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {
var __t9 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3711209382) {
__t9 = gopurs_runtime.Apply(e_4, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_9
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {
var spin_6_8_4 gopurs_runtime.Value
spin_6_8_4 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop gopurs_runtime.Value = v_7_loop_val
spin_6_8_4:
for {
if false { continue spin_6_8_4 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
v_7_loop = v_7
continue spin_6_8_4
return gopurs_runtime.Value{}
}
}()
})
__t9 = gopurs_runtime.Apply(spin_6_8_4, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t12 = __t11
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
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
var __t10 gopurs_runtime.Value
{
if (y_4.Type == 9 && y_4.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
if (y_4.Type == 9 && y_4.IntVal == 2465973597) {
var __t9 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t9 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_9
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var __t8 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t8 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(d_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {
var spin_5_6_5 gopurs_runtime.Value
spin_5_6_5 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
spin_5_6_5:
for {
if false { continue spin_5_6_5 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
v_6_loop = v_6
continue spin_5_6_5
return gopurs_runtime.Value{}
}
}()
})
__t7 = gopurs_runtime.Apply(spin_5_6_5, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
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
var __t7 gopurs_runtime.Value
{
if (y_3.Type == 9 && y_3.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
if (y_3.Type == 9 && y_3.IntVal == 2465973597) {
var __t6 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t6 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382) {
__t5 = gopurs_runtime.Apply(c_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {
var spin_4_4_6 gopurs_runtime.Value
spin_4_4_6 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
spin_4_4_6:
for {
if false { continue spin_4_4_6 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
v_5_loop = v_5
continue spin_4_4_6
return gopurs_runtime.Value{}
}
}()
})
__t5 = gopurs_runtime.Apply(spin_4_4_6, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}

func Call_coproduct2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t4 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(a_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {
var __t3 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply(b_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {
var spin_3_2_7 gopurs_runtime.Value
spin_3_2_7 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
spin_3_2_7:
for {
if false { continue spin_3_2_7 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
v_4_loop = v_4
continue spin_3_2_7
return gopurs_runtime.Value{}
}
}()
})
__t3 = gopurs_runtime.Apply(spin_3_2_7, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}

func Call_coproduct1(y_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var __t1 gopurs_runtime.Value
{
if (y_0.Type == 9 && y_0.IntVal == 3711209382) {
__t1 = (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
if (y_0.Type == 9 && y_0.IntVal == 2465973597) {
var spin_1_0_9 gopurs_runtime.Value
spin_1_0_9 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_9:
for {
if false { continue spin_1_0_9 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_9
return gopurs_runtime.Value{}
}
}()
})
__t1 = gopurs_runtime.Apply(spin_1_0_9, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_at9(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t16 gopurs_runtime.Value
{
var __t_and_15 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
var __t_and_14 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_13 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_12 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_11 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_9 bool = false
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {

var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_8 bool = false
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {

var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
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
__t_and_15 = __t_and_14
}
if __t_and_15 {
__t16 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = b_0
}
end_branch_16:
return __t16
}

func Call_at8(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t14 gopurs_runtime.Value
{
var __t_and_13 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
var __t_and_12 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_11 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_9 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_8 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_7 bool = false
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {

var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
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
__t_and_13 = __t_and_12
}
if __t_and_13 {
__t14 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = b_0
}
end_branch_14:
return __t14
}

func Call_at7(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t12 gopurs_runtime.Value
{
var __t_and_11 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_9 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_8 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_7 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
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
__t_and_11 = __t_and_10
}
if __t_and_11 {
__t12 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
__t12 = b_0
}
end_branch_12:
return __t12
}

func Call_at6(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t10 gopurs_runtime.Value
{
var __t_and_9 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
var __t_and_8 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_7 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3711209382)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
__t_and_8 = __t_and_7
}
__t_and_9 = __t_and_8
}
if __t_and_9 {
__t10 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = b_0
}
end_branch_10:
return __t10
}

func Call_at5(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t8 gopurs_runtime.Value
{
var __t_and_7 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_5 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_4 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_4 = (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 3711209382)
}
__t_and_5 = __t_and_4
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
__t8 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = b_0
}
end_branch_8:
return __t8
}

func Call_at4(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t6 gopurs_runtime.Value
{
var __t_and_5 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
var __t_and_4 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_3 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
__t_and_3 = (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 3711209382)
}
__t_and_4 = __t_and_3
}
__t_and_5 = __t_and_4
}
if __t_and_5 {
__t6 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = b_0
}
end_branch_6:
return __t6
}

func Call_at3(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t4 gopurs_runtime.Value
{
var __t_and_3 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
var __t_and_2 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0
__t_and_2 = (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 3711209382)
}
__t_and_3 = __t_and_2
}
if __t_and_3 {
__t4 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
__t4 = b_0
}
end_branch_4:
return __t4
}

func Call_at2(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t2 gopurs_runtime.Value
{
var __t_and_1 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
__t_and_1 = (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 3711209382)
}
if __t_and_1 {
__t2 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = b_0
}
end_branch_2:
return __t2
}

func Call_at10(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t18 gopurs_runtime.Value
{
var __t_and_17 bool = false
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {

var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0
var __t_and_16 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 2465973597) {

var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0
var __t_and_15 bool = false
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2465973597) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_14 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2465973597) {

var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_13 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2465973597) {

var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_12 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2465973597) {

var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_11 bool = false
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2465973597) {

var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2465973597) {

var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
var __t_and_9 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 2465973597) {

var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0
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
__t_and_17 = __t_and_16
}
if __t_and_17 {
__t18 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = b_0
}
end_branch_18:
return __t18
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
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0)
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

func Call_left__786465523(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0})}
}

func Call_left__1785278009(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0})}
}

func Call_left__4018477375(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0})}
}

func Call_right__1193983057(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__1623275833(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__856638295(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__1175236921(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__3235459543(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__2015984185(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__176310615(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__1052251193(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__2641088471(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_right__3220060985(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__3741956339(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_absurd__3279552488(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_10 gopurs_runtime.Value
spin_1_0_10 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_10:
for {
if false { continue spin_1_0_10 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_10
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_10, a_0)
}

func Call_absurd__943547431(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_11 gopurs_runtime.Value
spin_1_0_11 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_11:
for {
if false { continue spin_1_0_11 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_11
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_11, a_0)
}


