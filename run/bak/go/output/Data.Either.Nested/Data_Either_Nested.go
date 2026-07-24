package Data_Either_Nested

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Either "gopurs/output/Data.Either"
)

var in9 gopurs_runtime.Value
var once_in9 sync.Once
func Get_in9() gopurs_runtime.Value {
	once_in9.Do(func() {
		in9 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0)))))))))
}()
})
	})
	return in9
}

var in8 gopurs_runtime.Value
var once_in8 sync.Once
func Get_in8() gopurs_runtime.Value {
	once_in8.Do(func() {
		in8 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0))))))))
}()
})
	})
	return in8
}

var in7 gopurs_runtime.Value
var once_in7 sync.Once
func Get_in7() gopurs_runtime.Value {
	once_in7.Do(func() {
		in7 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0)))))))
}()
})
	})
	return in7
}

var in6 gopurs_runtime.Value
var once_in6 sync.Once
func Get_in6() gopurs_runtime.Value {
	once_in6.Do(func() {
		in6 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0))))))
}()
})
	})
	return in6
}

var in5 gopurs_runtime.Value
var once_in5 sync.Once
func Get_in5() gopurs_runtime.Value {
	once_in5.Do(func() {
		in5 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0)))))
}()
})
	})
	return in5
}

var in4 gopurs_runtime.Value
var once_in4 sync.Once
func Get_in4() gopurs_runtime.Value {
	once_in4.Do(func() {
		in4 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0))))
}()
})
	})
	return in4
}

var in3 gopurs_runtime.Value
var once_in3 sync.Once
func Get_in3() gopurs_runtime.Value {
	once_in3.Do(func() {
		in3 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0)))
}()
})
	})
	return in3
}

var in2 gopurs_runtime.Value
var once_in2 sync.Once
func Get_in2() gopurs_runtime.Value {
	once_in2.Do(func() {
		in2 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0))
}()
})
	})
	return in2
}

var in10 gopurs_runtime.Value
var once_in10 sync.Once
func Get_in10() gopurs_runtime.Value {
	once_in10.Do(func() {
		in10 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor1("Left", v_0))))))))))
}()
})
	})
	return in10
}

var in1 gopurs_runtime.Value
var once_in1 sync.Once
func Get_in1() gopurs_runtime.Value {
	once_in1.Do(func() {
		in1 = pkg_Data_Either.Get_Left()
	})
	return in1
}

var either9 gopurs_runtime.Value
var once_either9 sync.Once
func Get_either9() gopurs_runtime.Value {
	once_either9.Do(func() {
		either9 = gopurs_runtime.Func10(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, i_8_box gopurs_runtime.Value, y_9_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either9(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box, i_8_box, y_9_box)
})
	})
	return either9
}

var either8 gopurs_runtime.Value
var once_either8 sync.Once
func Get_either8() gopurs_runtime.Value {
	once_either8.Do(func() {
		either8 = gopurs_runtime.Func9(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, y_8_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either8(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box, y_8_box)
})
	})
	return either8
}

var either7 gopurs_runtime.Value
var once_either7 sync.Once
func Get_either7() gopurs_runtime.Value {
	once_either7.Do(func() {
		either7 = gopurs_runtime.Func8(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, y_7_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either7(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, y_7_box)
})
	})
	return either7
}

var either6 gopurs_runtime.Value
var once_either6 sync.Once
func Get_either6() gopurs_runtime.Value {
	once_either6.Do(func() {
		either6 = gopurs_runtime.Func7(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, y_6_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either6(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, y_6_box)
})
	})
	return either6
}

var either5 gopurs_runtime.Value
var once_either5 sync.Once
func Get_either5() gopurs_runtime.Value {
	once_either5.Do(func() {
		either5 = gopurs_runtime.Func6(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, y_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either5(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, y_5_box)
})
	})
	return either5
}

var either4 gopurs_runtime.Value
var once_either4 sync.Once
func Get_either4() gopurs_runtime.Value {
	once_either4.Do(func() {
		either4 = gopurs_runtime.Func5(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, y_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either4(a_0_box, b_1_box, c_2_box, d_3_box, y_4_box)
})
	})
	return either4
}

var either3 gopurs_runtime.Value
var once_either3 sync.Once
func Get_either3() gopurs_runtime.Value {
	once_either3.Do(func() {
		either3 = gopurs_runtime.Func4(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either3(a_0_box, b_1_box, c_2_box, y_3_box)
})
	})
	return either3
}

var either2 gopurs_runtime.Value
var once_either2 sync.Once
func Get_either2() gopurs_runtime.Value {
	once_either2.Do(func() {
		either2 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either2(a_0_box, b_1_box, y_2_box)
})
	})
	return either2
}

var either10 gopurs_runtime.Value
var once_either10 sync.Once
func Get_either10() gopurs_runtime.Value {
	once_either10.Do(func() {
		either10 = gopurs_runtime.Func(func(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(y_10.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_10.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(c_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Apply(d_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Apply(e_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t5 = gopurs_runtime.Apply(f_5, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t6 = gopurs_runtime.Apply(g_6, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t7 = gopurs_runtime.Apply(h_7, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t8 = gopurs_runtime.Apply(i_8, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_8
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t9 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t9 = gopurs_runtime.Apply(j_9, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_9
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_11_10 gopurs_runtime.Value
spin_11_10 = gopurs_runtime.Func(func(v_12_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_11_10:
for {
if false { continue spin_11_10 }
var v_12 gopurs_runtime.Value = v_12_loop
_ = v_12
v_12_loop = v_12
continue spin_11_10
return gopurs_runtime.Value{}
}
}()
})
__t9 = gopurs_runtime.Apply(spin_11_10, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t8 = __t9
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t7 = __t8
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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
	return either10
}

var either1 gopurs_runtime.Value
var once_either1 sync.Once
func Get_either1() gopurs_runtime.Value {
	once_either1.Do(func() {
		either1 = gopurs_runtime.Func(func(y_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_0.StrVal == "Left").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(y_0.UnsafePtr)[0]
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_0.StrVal == "Right").IntVal != 0 {
var spin_1_1 gopurs_runtime.Value
spin_1_1 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Apply(spin_1_1, (*[1024]gopurs_runtime.Value)(y_0.UnsafePtr)[0])
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
	return either1
}

var at9 gopurs_runtime.Value
var once_at9 sync.Once
func Get_at9() gopurs_runtime.Value {
	once_at9.Do(func() {
		at9 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at9(b_0_box, f_1_box, y_2_box)
})
	})
	return at9
}

var at8 gopurs_runtime.Value
var once_at8 sync.Once
func Get_at8() gopurs_runtime.Value {
	once_at8.Do(func() {
		at8 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at8(b_0_box, f_1_box, y_2_box)
})
	})
	return at8
}

var at7 gopurs_runtime.Value
var once_at7 sync.Once
func Get_at7() gopurs_runtime.Value {
	once_at7.Do(func() {
		at7 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at7(b_0_box, f_1_box, y_2_box)
})
	})
	return at7
}

var at6 gopurs_runtime.Value
var once_at6 sync.Once
func Get_at6() gopurs_runtime.Value {
	once_at6.Do(func() {
		at6 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at6(b_0_box, f_1_box, y_2_box)
})
	})
	return at6
}

var at5 gopurs_runtime.Value
var once_at5 sync.Once
func Get_at5() gopurs_runtime.Value {
	once_at5.Do(func() {
		at5 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at5(b_0_box, f_1_box, y_2_box)
})
	})
	return at5
}

var at4 gopurs_runtime.Value
var once_at4 sync.Once
func Get_at4() gopurs_runtime.Value {
	once_at4.Do(func() {
		at4 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at4(b_0_box, f_1_box, y_2_box)
})
	})
	return at4
}

var at3 gopurs_runtime.Value
var once_at3 sync.Once
func Get_at3() gopurs_runtime.Value {
	once_at3.Do(func() {
		at3 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at3(b_0_box, f_1_box, y_2_box)
})
	})
	return at3
}

var at2 gopurs_runtime.Value
var once_at2 sync.Once
func Get_at2() gopurs_runtime.Value {
	once_at2.Do(func() {
		at2 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at2(b_0_box, f_1_box, y_2_box)
})
	})
	return at2
}

var at10 gopurs_runtime.Value
var once_at10 sync.Once
func Get_at10() gopurs_runtime.Value {
	once_at10.Do(func() {
		at10 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at10(b_0_box, f_1_box, y_2_box)
})
	})
	return at10
}

var at1 gopurs_runtime.Value
var once_at1 sync.Once
func Get_at1() gopurs_runtime.Value {
	once_at1.Do(func() {
		at1 = gopurs_runtime.Func3(func(b_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_at1(b_0_box, f_1_box, y_2_box)
})
	})
	return at1
}

func Call_either9(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, h_7_loop gopurs_runtime.Value, i_8_loop gopurs_runtime.Value, y_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(y_9.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_9.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(c_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Apply(d_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Apply(e_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t5 = gopurs_runtime.Apply(f_5, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t6 = gopurs_runtime.Apply(g_6, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t7 = gopurs_runtime.Apply(h_7, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t8 = gopurs_runtime.Apply(i_8, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_8
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_10_9 gopurs_runtime.Value
spin_10_9 = gopurs_runtime.Func(func(v_11_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_10_9:
for {
if false { continue spin_10_9 }
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
v_11_loop = v_11
continue spin_10_9
return gopurs_runtime.Value{}
}
}()
})
__t8 = gopurs_runtime.Apply(spin_10_9, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t7 = __t8
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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

func Call_either8(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, h_7_loop gopurs_runtime.Value, y_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(y_8.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_8.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(c_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Apply(d_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Apply(e_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t5 = gopurs_runtime.Apply(f_5, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t6 = gopurs_runtime.Apply(g_6, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t7 = gopurs_runtime.Apply(h_7, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_9_8 gopurs_runtime.Value
spin_9_8 = gopurs_runtime.Func(func(v_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_9_8:
for {
if false { continue spin_9_8 }
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
v_10_loop = v_10
continue spin_9_8
return gopurs_runtime.Value{}
}
}()
})
__t7 = gopurs_runtime.Apply(spin_9_8, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_8.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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

func Call_either7(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, y_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(y_7.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_7.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(c_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Apply(d_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Apply(e_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t5 = gopurs_runtime.Apply(f_5, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t6 = gopurs_runtime.Apply(g_6, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_8_7 gopurs_runtime.Value
spin_8_7 = gopurs_runtime.Func(func(v_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_8_7:
for {
if false { continue spin_8_7 }
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
v_9_loop = v_9
continue spin_8_7
return gopurs_runtime.Value{}
}
}()
})
__t6 = gopurs_runtime.Apply(spin_8_7, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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

func Call_either6(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, y_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(y_6.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_6.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(c_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Apply(d_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Apply(e_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t5 = gopurs_runtime.Apply(f_5, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_7_6 gopurs_runtime.Value
spin_7_6 = gopurs_runtime.Func(func(v_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_7_6:
for {
if false { continue spin_7_6 }
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
v_8_loop = v_8
continue spin_7_6
return gopurs_runtime.Value{}
}
}()
})
__t5 = gopurs_runtime.Apply(spin_7_6, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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

func Call_either5(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, y_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(y_5.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_5.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(c_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Apply(d_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Apply(e_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_6_5 gopurs_runtime.Value
spin_6_5 = gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_6_5:
for {
if false { continue spin_6_5 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
v_7_loop = v_7
continue spin_6_5
return gopurs_runtime.Value{}
}
}()
})
__t4 = gopurs_runtime.Apply(spin_6_5, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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

func Call_either4(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, y_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(y_4.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_4.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(c_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Apply(d_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_5_4 gopurs_runtime.Value
spin_5_4 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_5_4:
for {
if false { continue spin_5_4 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
v_6_loop = v_6
continue spin_5_4
return gopurs_runtime.Value{}
}
}()
})
__t3 = gopurs_runtime.Apply(spin_5_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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

func Call_either3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(y_3.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_3.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(c_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_4_3 gopurs_runtime.Value
spin_4_3 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_4_3:
for {
if false { continue spin_4_3 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
v_5_loop = v_5
continue spin_4_3
return gopurs_runtime.Value{}
}
}()
})
__t2 = gopurs_runtime.Apply(spin_4_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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

func Call_either2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_2.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(a_0, (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(b_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
var spin_3_2 gopurs_runtime.Value
spin_3_2 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_3_2:
for {
if false { continue spin_3_2 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
v_4_loop = v_4
continue spin_3_2
return gopurs_runtime.Value{}
}
}()
})
__t1 = gopurs_runtime.Apply(spin_3_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0])
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

func Call_at9(b_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0].UnsafePtr)[0])
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
if gopurs_runtime.Bool(y_2.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0])
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


