package Data_Functor_Product_Nested

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	unsafe "unsafe"
)

var product9 gopurs_runtime.Value
var once_product9 sync.Once
func Get_product9() gopurs_runtime.Value {
	once_product9.Do(func() {
		product9 = gopurs_runtime.Func9(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, i_8_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product9(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box, i_8_box)
})
	})
	return product9
}

var product8 gopurs_runtime.Value
var once_product8 sync.Once
func Get_product8() gopurs_runtime.Value {
	once_product8.Do(func() {
		product8 = gopurs_runtime.Func8(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product8(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box)
})
	})
	return product8
}

var product7 gopurs_runtime.Value
var once_product7 sync.Once
func Get_product7() gopurs_runtime.Value {
	once_product7.Do(func() {
		product7 = gopurs_runtime.Func7(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product7(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box)
})
	})
	return product7
}

var product6 gopurs_runtime.Value
var once_product6 sync.Once
func Get_product6() gopurs_runtime.Value {
	once_product6.Do(func() {
		product6 = gopurs_runtime.Func6(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product6(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box)
})
	})
	return product6
}

var product5 gopurs_runtime.Value
var once_product5 sync.Once
func Get_product5() gopurs_runtime.Value {
	once_product5.Do(func() {
		product5 = gopurs_runtime.Func5(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product5(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box)
})
	})
	return product5
}

var product4 gopurs_runtime.Value
var once_product4 sync.Once
func Get_product4() gopurs_runtime.Value {
	once_product4.Do(func() {
		product4 = gopurs_runtime.Func4(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product4(a_0_box, b_1_box, c_2_box, d_3_box)
})
	})
	return product4
}

var product3 gopurs_runtime.Value
var once_product3 sync.Once
func Get_product3() gopurs_runtime.Value {
	once_product3.Do(func() {
		product3 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product3(a_0_box, b_1_box, c_2_box)
})
	})
	return product3
}

var product2 gopurs_runtime.Value
var once_product2 sync.Once
func Get_product2() gopurs_runtime.Value {
	once_product2.Do(func() {
		product2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product2(a_0_box, b_1_box)
})
	})
	return product2
}

var product10 gopurs_runtime.Value
var once_product10 sync.Once
func Get_product10() gopurs_runtime.Value {
	once_product10.Do(func() {
		product10 = gopurs_runtime.Func10(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, i_8_box gopurs_runtime.Value, j_9_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product10(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box, i_8_box, j_9_box)
})
	})
	return product10
}

var product1 gopurs_runtime.Value
var once_product1 sync.Once
func Get_product1() gopurs_runtime.Value {
	once_product1.Do(func() {
		product1 = gopurs_runtime.Func(func(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, pkg_Data_Unit.Get_unit()})}
}()
})
	})
	return product1
}

var get9 gopurs_runtime.Value
var once_get9 sync.Once
func Get_get9() gopurs_runtime.Value {
	once_get9.Do(func() {
		get9 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get9
}

var get8 gopurs_runtime.Value
var once_get8 sync.Once
func Get_get8() gopurs_runtime.Value {
	once_get8.Do(func() {
		get8 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get8
}

var get7 gopurs_runtime.Value
var once_get7 sync.Once
func Get_get7() gopurs_runtime.Value {
	once_get7.Do(func() {
		get7 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get7
}

var get6 gopurs_runtime.Value
var once_get6 sync.Once
func Get_get6() gopurs_runtime.Value {
	once_get6.Do(func() {
		get6 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get6
}

var get5 gopurs_runtime.Value
var once_get5 sync.Once
func Get_get5() gopurs_runtime.Value {
	once_get5.Do(func() {
		get5 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get5
}

var get4 gopurs_runtime.Value
var once_get4 sync.Once
func Get_get4() gopurs_runtime.Value {
	once_get4.Do(func() {
		get4 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get4
}

var get3 gopurs_runtime.Value
var once_get3 sync.Once
func Get_get3() gopurs_runtime.Value {
	once_get3.Do(func() {
		get3 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get3
}

var get2 gopurs_runtime.Value
var once_get2 sync.Once
func Get_get2() gopurs_runtime.Value {
	once_get2.Do(func() {
		get2 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get2
}

var get10 gopurs_runtime.Value
var once_get10 sync.Once
func Get_get10() gopurs_runtime.Value {
	once_get10.Do(func() {
		get10 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}()
})
	})
	return get10
}

var get1 gopurs_runtime.Value
var once_get1 sync.Once
func Get_get1() gopurs_runtime.Value {
	once_get1.Do(func() {
		get1 = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
}()
})
	})
	return get1
}

func Call_product9(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, h_7_loop gopurs_runtime.Value, i_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{c_2, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{d_3, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{e_4, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{f_5, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{g_6, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{h_7, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{i_8, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}})}})}})}
}

func Call_product8(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, h_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{c_2, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{d_3, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{e_4, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{f_5, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{g_6, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{h_7, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}})}})}
}

func Call_product7(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{c_2, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{d_3, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{e_4, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{f_5, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{g_6, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}})}
}

func Call_product6(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{c_2, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{d_3, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{e_4, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{f_5, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}
}

func Call_product5(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{c_2, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{d_3, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{e_4, pkg_Data_Unit.Get_unit()})}})}})}})}})}
}

func Call_product4(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
var d_3 gopurs_runtime.Value = d_3_loop
_ = d_3
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{c_2, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{d_3, pkg_Data_Unit.Get_unit()})}})}})}})}
}

func Call_product3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{c_2, pkg_Data_Unit.Get_unit()})}})}})}
}

func Call_product2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, pkg_Data_Unit.Get_unit()})}})}
}

func Call_product10(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value, d_3_loop gopurs_runtime.Value, e_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value, g_6_loop gopurs_runtime.Value, h_7_loop gopurs_runtime.Value, i_8_loop gopurs_runtime.Value, j_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{b_1, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{c_2, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{d_3, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{e_4, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{f_5, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{g_6, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{h_7, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{i_8, gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{j_9, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}})}})}})}})}
}


